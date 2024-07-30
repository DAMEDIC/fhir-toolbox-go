package capabilities

import (
	"context"
	"fhir-toolbox/capabilities/search"
	"fhir-toolbox/model"
)

type GenericAPI interface {
	GenericRead
	GenericSearch
}

type GenericRead interface {
	Read(ctx context.Context, resourceType string, id string) (model.Resource, FHIRError)
}

type GenericSearch interface {
	SearchCapabilities(resourceType string) (search.Capabilities, FHIRError)
	Search(ctx context.Context, resourceType string, options search.Options) (search.Result, FHIRError)
}
