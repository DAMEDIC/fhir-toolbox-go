package rest

import (
	"cmp"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/utils/ptr"
	"net/url"
	"slices"
	"time"
)

// capabilityStatement building from provided capabilities.
func capabilityStatement[R model.Release](
	baseURL *url.URL,
	capabilities capabilities.Capabilities,
	date time.Time,
) basic.CapabilityStatement {
	metadataURL := baseURL.JoinPath("metadata").String()

	return basic.CapabilityStatement{
		Status: basic.Code{Value: ptr.To("active")},
		// the capability statement is created once at system start
		// we simply use the system start as last changed date for now
		Date: basic.DateTime{Value: ptr.To(date.Format(time.RFC3339))},
		Kind: basic.Code{Value: ptr.To("instance")},
		Software: &basic.CapabilityStatementSoftware{
			Name: basic.String{Value: ptr.To("fhir-toolbox-go")},
		},
		Implementation: &basic.CapabilityStatementImplementation{
			Description: basic.String{Value: ptr.To("a simple FHIR service built with fhir-toolbox-go")},
			Url:         &basic.Url{Value: &metadataURL},
		},
		FhirVersion: basic.Code{Value: ptr.To(model.ReleaseVersion[R]())},
		Format:      []basic.Code{{Value: ptr.To("xml")}, {Value: ptr.To("json")}},
		Rest:        []basic.CapabilityStatementRest{rest(capabilities)},
	}
}

func rest(
	capabilities capabilities.Capabilities,
) basic.CapabilityStatementRest {
	resourcesMap := map[string]basic.CapabilityStatementRestResource{}

	// Helper function to get or create a resource and add an interaction
	addInteraction := func(name string, interactionCode string) basic.CapabilityStatementRestResource {
		r, ok := resourcesMap[name]
		if !ok {
			r = basic.CapabilityStatementRestResource{Type: basic.Code{Value: &name}}
		}
		r.Interaction = append(
			r.Interaction,
			basic.CapabilityStatementRestResourceInteraction{Code: basic.Code{Value: ptr.To(interactionCode)}},
		)
		return r
	}

	for name := range capabilities.Create {
		resourcesMap[name] = addInteraction(name, "create")
	}

	for name := range capabilities.Read {
		resourcesMap[name] = addInteraction(name, "read")
	}

	for name, capability := range capabilities.Update {
		r := addInteraction(name, "update")
		r.UpdateCreate = &basic.Boolean{Value: ptr.To(capability.UpdateCreate)}
		resourcesMap[name] = r
	}

	for name := range capabilities.Delete {
		resourcesMap[name] = addInteraction(name, "delete")
	}

	for name, capability := range capabilities.Search {
		r := addInteraction(name, "search-type")

		for _, include := range capability.Includes {
			r.SearchInclude = append(
				r.SearchInclude,
				basic.String{Value: &include},
			)
		}

		for spName, sp := range capability.Parameters {
			fhirpathType, ok, err := fhirpath.Singleton[fhirpath.String](sp.Children("type"))
			if !ok || err != nil {
				// should never happen
				continue
			}
			resolvedType := string(fhirpathType)

			r.SearchParam = append(
				r.SearchParam,
				basic.CapabilityStatementRestResourceSearchParam{
					Name: basic.String{Value: &spName},
					Type: basic.Code{Value: &resolvedType},
				},
			)
		}
		resourcesMap[name] = r
	}

	resourcesList := make([]basic.CapabilityStatementRestResource, 0, len(resourcesMap))

	for _, r := range resourcesMap {
		// sort search params by name (they come from a map, so we can't rely on the order)
		// this makes testing easier
		slices.SortStableFunc(r.SearchParam, func(a, b basic.CapabilityStatementRestResourceSearchParam) int {
			return cmp.Compare(*a.Name.Value, *b.Name.Value)
		})

		resourcesList = append(resourcesList, r)
	}

	slices.SortFunc(resourcesList, func(a, b basic.CapabilityStatementRestResource) int {
		return cmp.Compare(*a.Type.Value, *b.Type.Value)
	})

	return basic.CapabilityStatementRest{
		Mode:     basic.Code{Value: ptr.To("server")},
		Resource: resourcesList,
	}
}
