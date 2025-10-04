// Example server exposing operations at system, type, and instance levels
// using the concrete API. It demonstrates how the REST server discovers
// operations from CapabilityStatement and dispatches to concrete methods
// named Invoke{Name}[Operation] with appropriate arity.
package main

import (
	"context"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/rest"
	"log"
	"net/http"
	"time"

	basic "github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	r5 "github.com/DAMEDIC/fhir-toolbox-go/model/gen/r5"
	"github.com/DAMEDIC/fhir-toolbox-go/utils/ptr"
)

// opBackend implements concrete operations and minimal reads needed by the server.
type opBackend struct{}

// CapabilityBase declares system/type/instance operations via CapabilityStatement.
func (b *opBackend) CapabilityBase(ctx context.Context) (basic.CapabilityStatement, error) {
	// Keep base metadata minimal; operations are detected via Definition methods.
	return basic.CapabilityStatement{
		Status:      basic.Code{Value: ptr.To("active")},
		Date:        basic.DateTime{Value: ptr.To(time.Now().Format(time.RFC3339))},
		Kind:        basic.Code{Value: ptr.To("instance")},
		FhirVersion: basic.Code{Value: ptr.To("5.0")},
		Format:      []basic.Code{{Value: ptr.To("json")}},
		Software: &basic.CapabilityStatementSoftware{
			Name:    basic.String{Value: ptr.To("operations-demo")},
			Version: &basic.String{Value: ptr.To("0.0.1")},
		},
		Implementation: &basic.CapabilityStatementImplementation{
			Description: basic.String{Value: ptr.To("Demo operations server (concrete API)")},
			Url:         &basic.Url{Value: ptr.To("http://localhost:8080")},
		},
	}, nil
}

// Operation Definitions exposed via concrete API
func (b *opBackend) PingOperationDefinition(ctx context.Context /* is optional */) r5.OperationDefinition {
	return r5.OperationDefinition{
		Id:     &r5.Id{Value: ptr.To("ping")},
		Code:   r5.Code{Value: ptr.To("ping")},
		System: r5.Boolean{Value: ptr.To(true)},
	}
}

// System: POST/GET /$ping -> returns Parameters echoing inputs and a server timestamp.
func (b *opBackend) InvokePing(ctx context.Context, parameters basic.Parameters) (basic.Parameters, error) {
	now := time.Now().Format(time.RFC3339)
	ps := append([]basic.ParametersParameter{}, parameters.Parameter...)
	name := "timestamp"
	ps = append(ps, basic.ParametersParameter{Name: basic.String{Value: &name}, Value: basic.String{Value: &now}})
	return basic.Parameters{Parameter: ps}, nil
}

func (b *opBackend) EchoOperationDefinition() r5.OperationDefinition {
	return r5.OperationDefinition{
		Id:       &r5.Id{Value: ptr.To("echo")},
		Code:     r5.Code{Value: ptr.To("echo")},
		Type:     r5.Boolean{Value: ptr.To(true)},
		Resource: []r5.Code{{Value: ptr.To("Patient")}},
	}
}

// Type: POST/GET /Patient/$echo -> returns a Patient with name from parameters["name"].
func (b *opBackend) InvokeEcho(ctx context.Context, resourceType string, parameters basic.Parameters) (r5.Patient, error) {
	var given string
	for _, p := range parameters.Parameter {
		if p.Name.Value != nil && *p.Name.Value == "name" {
			s, ok, err := p.Value.ToString(false)
			if err == nil && ok && s != "" {
				given = string(s)
				break
			}
		}
	}
	// Build a simple Patient resource
	pat := r5.Patient{
		Name: []r5.HumanName{{Given: []r5.String{{Value: &given}}}},
	}
	return pat, nil
}

func (b *opBackend) HelloOperationDefinition() r5.OperationDefinition {
	return r5.OperationDefinition{
		Id:       &r5.Id{Value: ptr.To("hello")},
		Code:     r5.Code{Value: ptr.To("hello")},
		Instance: r5.Boolean{Value: ptr.To(true)},
		Resource: []r5.Code{{Value: ptr.To("Patient")}},
	}
}

// Instance: POST/GET /Patient/{id}/$hello -> returns Parameters greeting the id.
func (b *opBackend) InvokeHello(ctx context.Context, resourceType string, resourceID string, parameters basic.Parameters) (basic.Parameters, error) {
	msg := "hello, " + resourceType + " " + resourceID
	name := "message"
	return basic.Parameters{Parameter: []basic.ParametersParameter{{
		Name:  basic.String{Value: &name},
		Value: basic.String{Value: &msg},
	}}}, nil
}

func main() {
	backend := &opBackend{}
	server := &rest.Server[model.R5]{Backend: backend}

	log.Println("listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
