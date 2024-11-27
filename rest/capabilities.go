package rest

import (
	"cmp"
	"fhir-toolbox/capabilities"
	"fhir-toolbox/model"
	"fhir-toolbox/model/basic"
	"net/url"
	"slices"
	"time"
)

func CapabilityStatement[R model.Release](
	baseURL *url.URL,
	capabilities capabilities.Capabilities,
	date time.Time,
) basic.CapabilityStatement {
	return basic.CapabilityStatement{
		Status: "active",
		// the capability statement is created once at system start
		// we simply use the system start as last changed date for now
		Date: date.Format(time.RFC3339),
		Kind: "instance",
		Software: basic.CapabilityStatementSoftware{
			Name: "fhir-toolbox-go",
		},
		Implementation: basic.CapabilityStatementImplementation{
			Description: "a simple FHIR service built with fhir-toolbox-go",
			Url:         baseURL.String(),
		},
		FHIRVersion: model.ReleaseVersion[R](),
		Format:      []string{"xml", "json"},
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
			r = basic.CapabilityStatementRestResource{Type: name}
		}
		r.Interaction = append(
			resourcesMap[name].Interaction,
			basic.CapabilityStatementRestResourceInteraction{Code: "read"},
		)
		resourcesMap[name] = r
	}

	for name, capability := range capabilities.SearchCapabilities {
		r, ok := resourcesMap[name]
		if !ok {
			r = basic.CapabilityStatementRestResource{Type: name}
		}
		r.Interaction = append(
			resourcesMap[name].Interaction,
			basic.CapabilityStatementRestResourceInteraction{Code: "search-type"},
		)
		r.SearchInclude = capability.Includes
		for paramName, paramProps := range capability.Params {
			r.SearchParam = append(
				r.SearchParam,
				basic.CapabilityStatementRestResourceSearchParam{
					Name: paramName,
					Type: string(paramProps.Type),
				},
			)
		}
		resourcesMap[name] = r
	}

	resourcesList := make([]basic.CapabilityStatementRestResource, 0, len(resourcesMap))

	for _, r := range resourcesMap {
		// Sort for deterministic output. This makes writing tests much easier.
		slices.SortFunc(r.Interaction, func(a, b basic.CapabilityStatementRestResourceInteraction) int {
			return cmp.Compare(a.Code, b.Code)
		})
		slices.Sort(r.SearchInclude)
		slices.SortFunc(r.SearchParam, func(a, b basic.CapabilityStatementRestResourceSearchParam) int {
			return cmp.Or(cmp.Compare(a.Name, b.Name), cmp.Compare(a.Type, b.Type))
		})

		resourcesList = append(resourcesList, r)
	}

	slices.SortFunc(resourcesList, func(a, b basic.CapabilityStatementRestResource) int {
		return cmp.Compare(a.Type, b.Type)
	})

	return basic.CapabilityStatementRest{
		Mode:     "server",
		Resource: resourcesList,
	}
}
