package rest

import (
	"cmp"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/utils"
	"net/url"
	"slices"
	"time"
)

func CapabilityStatement[R model.Release](
	baseURL *url.URL,
	capabilities capabilities.Capabilities,
	date time.Time,
) r4.CapabilityStatement {
	metadataURL := baseURL.JoinPath("metadata").String()

	return r4.CapabilityStatement{
		Status: r4.Code{Value: utils.Ptr("active")},
		// the capability statement is created once at system start
		// we simply use the system start as last changed date for now
		Date: r4.DateTime{Value: utils.Ptr(date.Format(time.RFC3339))},
		Kind: r4.Code{Value: utils.Ptr("instance")},
		Software: &r4.CapabilityStatementSoftware{
			Name: r4.String{Value: utils.Ptr("fhir-toolbox-go")},
		},
		Implementation: &r4.CapabilityStatementImplementation{
			Description: r4.String{Value: utils.Ptr("a simple FHIR service built with fhir-toolbox-go")},
			Url:         &r4.Url{Value: &metadataURL},
		},
		FhirVersion: r4.Code{Value: utils.Ptr(model.ReleaseVersion[R]())},
		Format:      []r4.Code{{Value: utils.Ptr("xml")}, {Value: utils.Ptr("json")}},
		Rest:        []r4.CapabilityStatementRest{rest(capabilities)},
	}
}

func rest(
	capabilities capabilities.Capabilities,
) r4.CapabilityStatementRest {
	resourcesMap := map[string]r4.CapabilityStatementRestResource{}

	for _, name := range capabilities.ReadInteractions {
		r, ok := resourcesMap[name]
		if !ok {
			r = r4.CapabilityStatementRestResource{Type: r4.Code{Value: &name}}
		}
		r.Interaction = append(
			resourcesMap[name].Interaction,
			r4.CapabilityStatementRestResourceInteraction{Code: r4.Code{Value: utils.Ptr("read")}},
		)
		resourcesMap[name] = r
	}

	for name, capability := range capabilities.SearchCapabilities {
		r, ok := resourcesMap[name]
		if !ok {
			r = r4.CapabilityStatementRestResource{Type: r4.Code{Value: &name}}
		}
		r.Interaction = append(
			resourcesMap[name].Interaction,
			r4.CapabilityStatementRestResourceInteraction{Code: r4.Code{Value: utils.Ptr("search-type")}},
		)

		for _, include := range capability.Includes {
			r.SearchInclude = append(
				r.SearchInclude,
				r4.String{Value: &include},
			)
		}

		for paramName, paramProps := range capability.Params {
			r.SearchParam = append(
				r.SearchParam,
				r4.CapabilityStatementRestResourceSearchParam{
					Name: r4.String{Value: &paramName},
					Type: r4.Code{Value: utils.Ptr(string(paramProps.Type))},
				},
			)
		}
		resourcesMap[name] = r
	}

	resourcesList := make([]r4.CapabilityStatementRestResource, 0, len(resourcesMap))

	for _, r := range resourcesMap {
		// Sort for deterministic output. This makes writing tests much easier.
		slices.SortFunc(r.Interaction, func(a, b r4.CapabilityStatementRestResourceInteraction) int {
			return cmp.Compare(*a.Code.Value, *b.Code.Value)
		})
		slices.SortFunc(r.SearchInclude, func(a, b r4.String) int {
			return cmp.Compare(*a.Value, *b.Value)
		})
		slices.SortFunc(r.SearchParam, func(a, b r4.CapabilityStatementRestResourceSearchParam) int {
			return cmp.Or(cmp.Compare(*a.Name.Value, *b.Name.Value), cmp.Compare(*a.Type.Value, *b.Type.Value))
		})

		resourcesList = append(resourcesList, r)
	}

	slices.SortFunc(resourcesList, func(a, b r4.CapabilityStatementRestResource) int {
		return cmp.Compare(*a.Type.Value, *b.Type.Value)
	})

	return r4.CapabilityStatementRest{
		Mode:     r4.Code{Value: utils.Ptr("server")},
		Resource: resourcesList,
	}
}
