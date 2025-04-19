package rest

import (
	"cmp"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/utils"
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
		Status: basic.Code{Value: utils.Ptr("active")},
		// the capability statement is created once at system start
		// we simply use the system start as last changed date for now
		Date: basic.DateTime{Value: utils.Ptr(date.Format(time.RFC3339))},
		Kind: basic.Code{Value: utils.Ptr("instance")},
		Software: &basic.CapabilityStatementSoftware{
			Name: basic.String{Value: utils.Ptr("fhir-toolbox-go")},
		},
		Implementation: &basic.CapabilityStatementImplementation{
			Description: basic.String{Value: utils.Ptr("a simple FHIR service built with fhir-toolbox-go")},
			Url:         &basic.Url{Value: &metadataURL},
		},
		FhirVersion: basic.Code{Value: utils.Ptr(model.ReleaseVersion[R]())},
		Format:      []basic.Code{{Value: utils.Ptr("xml")}, {Value: utils.Ptr("json")}},
		Rest:        []basic.CapabilityStatementRest{rest(capabilities)},
	}
}

func rest(
	capabilities capabilities.Capabilities,
) basic.CapabilityStatementRest {
	resourcesMap := map[string]basic.CapabilityStatementRestResource{}

	for _, name := range capabilities.ReadInteractions {
		r, ok := resourcesMap[name]
		if !ok {
			r = basic.CapabilityStatementRestResource{Type: basic.Code{Value: &name}}
		}
		r.Interaction = append(
			resourcesMap[name].Interaction,
			basic.CapabilityStatementRestResourceInteraction{Code: basic.Code{Value: utils.Ptr("read")}},
		)
		resourcesMap[name] = r
	}

	for name, capability := range capabilities.SearchCapabilities {
		r, ok := resourcesMap[name]
		if !ok {
			r = basic.CapabilityStatementRestResource{Type: basic.Code{Value: &name}}
		}
		r.Interaction = append(
			resourcesMap[name].Interaction,
			basic.CapabilityStatementRestResourceInteraction{Code: basic.Code{Value: utils.Ptr("search-type")}},
		)

		for _, include := range capability.Includes {
			r.SearchInclude = append(
				r.SearchInclude,
				basic.String{Value: &include},
			)
		}

		for paramName, paramProps := range capability.Parameters {
			r.SearchParam = append(
				r.SearchParam,
				basic.CapabilityStatementRestResourceSearchParam{
					Name: basic.String{Value: &paramName},
					Type: basic.Code{Value: utils.Ptr(string(paramProps.Type))},
				},
			)
		}
		resourcesMap[name] = r
	}

	resourcesList := make([]basic.CapabilityStatementRestResource, 0, len(resourcesMap))

	for _, r := range resourcesMap {
		// Sort for deterministic output. This makes writing tests much easier.
		slices.SortFunc(r.Interaction, func(a, b basic.CapabilityStatementRestResourceInteraction) int {
			return cmp.Compare(*a.Code.Value, *b.Code.Value)
		})
		slices.SortFunc(r.SearchInclude, func(a, b basic.String) int {
			return cmp.Compare(*a.Value, *b.Value)
		})
		slices.SortFunc(r.SearchParam, func(a, b basic.CapabilityStatementRestResourceSearchParam) int {
			return cmp.Or(cmp.Compare(*a.Name.Value, *b.Name.Value), cmp.Compare(*a.Type.Value, *b.Type.Value))
		})

		resourcesList = append(resourcesList, r)
	}

	slices.SortFunc(resourcesList, func(a, b basic.CapabilityStatementRestResource) int {
		return cmp.Compare(*a.Type.Value, *b.Type.Value)
	})

	return basic.CapabilityStatementRest{
		Mode:     basic.Code{Value: utils.Ptr("server")},
		Resource: resourcesList,
	}
}
