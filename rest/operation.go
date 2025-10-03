package rest

import (
	"context"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/rest/internal/encoding"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
)

// level represents the invocation level of an operation.
type level string

const (
	levelSystem   level = "system"
	levelType     level = "type"
	levelInstance level = "instance"
)

// handleOperation handles all operation invocations routed via patterns with '$' outside braces.
func (s *Server[R]) handleOperation(w http.ResponseWriter, r *http.Request) {
	requestFormat := s.detectFormat(r, "Content-Type")
	responseFormat := s.detectFormat(r, "Accept")

	level, resourceType, resourceID, code, _ := parseOperationRoute(r.URL.Path)

	// Build Parameters from request
	var params basic.Parameters
	switch r.Method {
	case http.MethodGet:
		params = parseOperationQuery(r.URL.Query())
	case http.MethodPost:
		if r.ContentLength == 0 {
			params = basic.Parameters{}
		} else {
			p, err := encoding.Decode[basic.Parameters](r.Body, encoding.Format(requestFormat))
			if err != nil {
				returnErr(w, err, responseFormat)
				return
			}
			params = p
		}
	default:
		returnErr(w, createOperationOutcome("fatal", "processing", "unsupported method for operation"), responseFormat)
		return
	}

	out, err := s.dispatchOperation(r.Context(), r.Method, level, resourceType, resourceID, code, params)
	if err != nil {
		returnErr(w, err, responseFormat)
		return
	}
	if out == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	returnResult(w, out, http.StatusOK, responseFormat)
}

// dispatchOperation performs the complete operation invocation flow.
func (s *Server[R]) dispatchOperation(
	ctx context.Context,
	method string,
	level level,
	resourceType, resourceID, code string,
	params basic.Parameters,
) (model.Resource, error) {
	anyBackend, err := s.genericBackend()
	if err != nil {
		slog.Error("error in backend configuration", "err", err)
		return nil, err
	}

	backend, impl := anyBackend.(capabilities.GenericOperation)
	if !impl {
		return nil, notImplementedError("operation")
	}

	// Get CapabilityStatement and resolve OperationDefinition
	cs, err := anyBackend.CapabilityStatement(ctx)
	if err != nil {
		slog.Error("error getting metadata", "err", err)
		return nil, err
	}

	canonical, err := resolveOperationCanonical(cs, resourceType, code)
	if err != nil {
		return nil, searchError(err.Error())
	}
	opID := extractIDFromCanonical(canonical)

	// We need GenericRead to fetch the OperationDefinition
	readBackend, ok := anyBackend.(capabilities.GenericRead)
	if !ok {
		return nil, notImplementedError("read for OperationDefinition")
	}
	opDef, err := readBackend.Read(ctx, "OperationDefinition", opID)
	if err != nil {
		return nil, err
	}

	// Validate allowed level
	if err := checkAllowedLevel(opDef, level, resourceType); err != nil {
		return nil, createOperationOutcome("fatal", "not-supported", err.Error())
	}

	// If GET is used for an operation that affects state, reject
	if method == http.MethodGet {
		if affectsState(opDef) {
			return nil, createOperationOutcome("fatal", "not-supported", "operation with affectsState=true must be invoked via POST")
		}
	}

	// Invoke backend
	out, err := backend.Invoke(ctx, resourceType, resourceID, code, params)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// affectsState inspects an OperationDefinition to determine if affectsState is true.
func affectsState(opDef model.Resource) bool {
	e, ok := opDef.(fhirpath.Element)
	if !ok {
		return false
	}
	b, ok, err := fhirpath.Singleton[fhirpath.Boolean](e.Children("affectsState"))
	if err != nil || !ok {
		return false
	}
	return bool(b)
}

// resolveOperationCanonical locates the OperationDefinition canonical URL.
func resolveOperationCanonical(cs basic.CapabilityStatement, resourceType string, code string) (string, error) {
	matches := func(name *string) bool {
		if name == nil {
			return false
		}
		n := *name
		return n == code || n == "$"+code
	}
	if resourceType != "" {
		for _, rest := range cs.Rest {
			for _, res := range rest.Resource {
				if res.Type.Value != nil && *res.Type.Value == resourceType {
					for _, op := range res.Operation {
						if matches(op.Name.Value) && op.Definition.Value != nil {
							return *op.Definition.Value, nil
						}
					}
				}
			}
		}
	}
	for _, rest := range cs.Rest {
		for _, op := range rest.Operation {
			if matches(op.Name.Value) && op.Definition.Value != nil {
				return *op.Definition.Value, nil
			}
		}
	}
	return "", fmt.Errorf("operation '%s' not declared in CapabilityStatement", code)
}

// parseOperationQuery assembles a basic.Parameters resource from URL query values.
//
// Values are captured as valueString entries. Repeating parameters are represented
// by multiple Parameter entries with the same Name.
// Modifiers (e.g., ":in") are kept as part of the Name to match FHIR rules.
func parseOperationQuery(query url.Values) basic.Parameters {
	var params []basic.ParametersParameter
	for k, values := range query {
		// Skip special result-modifying or formatting parameters if present
		if k == "_format" {
			continue
		}
		for _, v := range values {
			vv := v // create stable pointer
			name := k
			params = append(params, basic.ParametersParameter{
				Name:  basic.String{Value: &name},
				Value: basic.String{Value: &vv},
			})
		}
	}
	return basic.Parameters{Parameter: params}
}

// parseOperationRoute extracts operation invocation info from a URL path.
// Supports:
//   - /$code
//   - /{type}/$code
//   - /{type}/{resourceID}/$code
func parseOperationRoute(path string) (level, string, string, string, bool) {
	// trim spaces and trailing slashes (but keep root slash semantics)
	p := strings.Trim(path, " ")
	if p == "" || p == "/" {
		return "", "", "", "", false
	}
	// split and drop empty segs due to leading slash
	raw := strings.Split(p, "/")
	segs := make([]string, 0, len(raw))
	for _, s := range raw {
		if s != "" {
			segs = append(segs, s)
		}
	}
	if len(segs) == 0 {
		return "", "", "", "", false
	}
	last := segs[len(segs)-1]
	if !strings.HasPrefix(last, "$") || len(last) < 2 {
		return "", "", "", "", false
	}
	code := last[1:]
	switch len(segs) {
	case 1:
		return levelSystem, "", "", code, true
	case 2:
		return levelType, segs[0], "", code, true
	case 3:
		return levelInstance, segs[0], segs[1], code, true
	default:
		return "", "", "", "", false
	}
}

// valuesFromParameters converts a Parameters resource into url.Values by
// extracting string representations of value[x] fields where possible.
// valuesFromParameters was removed; operations use Parameters directly.

// checkAllowedLevel validates cross-version via FHIRPath.
func checkAllowedLevel(def model.Resource, level level, resourceType string) error {
	elem, ok := def.(fhirpath.Element)
	if !ok {
		return fmt.Errorf("unsupported OperationDefinition element")
	}
	system := boolField(elem, "system")
	typeLevel := boolField(elem, "type")
	instance := boolField(elem, "instance")
	allowed := stringList(elem, "resource")

	switch level {
	case levelSystem:
		if !system {
			return fmt.Errorf("operation not allowed at system level")
		}
	case levelType:
		if !typeLevel {
			return fmt.Errorf("operation not allowed at type level")
		}
		if !allowedOnResource(resourceType, allowed) {
			return fmt.Errorf("operation not allowed for resource type %s", resourceType)
		}
	case levelInstance:
		if !instance {
			return fmt.Errorf("operation not allowed at instance level")
		}
		if !allowedOnResource(resourceType, allowed) {
			return fmt.Errorf("operation not allowed for resource type %s", resourceType)
		}
	}
	return nil
}

func boolField(e fhirpath.Element, name string) bool {
	v, ok, err := fhirpath.Singleton[fhirpath.Boolean](e.Children(name))
	if err != nil || !ok {
		return false
	}
	return bool(v)
}

func stringList(e fhirpath.Element, name string) []string {
	var out []string
	for _, c := range e.Children(name) {
		s, ok, err := c.ToString(false)
		if err == nil && ok {
			out = append(out, string(s))
		}
	}
	return out
}

func allowedOnResource(resource string, allowed []string) bool {
	if len(allowed) == 0 {
		return true
	}
	for _, a := range allowed {
		if a == resource {
			return true
		}
	}
	return false
}
