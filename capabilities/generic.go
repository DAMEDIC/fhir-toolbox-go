package capabilities

import (
	"context"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
)

// A GenericAPI irrespective of the capabilities of the underlying concrete implementation.
//
// This can be used to build adapters, clients or generic servers.
// See [github.com/DAMEDIC/fhir-toolbox/capabilities/wrap] for conversion between concrete to generic APIs and vice versa.
type GenericAPI interface {
	GenericCapabilities
	GenericCreate
	GenericRead
	GenericUpdate
	GenericSearch
}

// The GenericCapabilities interface provides a generic capabilities method that returns all capabilities of the underlying concrete implementation.
type GenericCapabilities interface {
	AllCapabilities(ctx context.Context) (Capabilities, FHIRError)
}

// The GenericCreate interface provides a generic create capability.
//
// The persisted resource is returned.
type GenericCreate interface {
	Create(ctx context.Context, resource model.Resource) (model.Resource, FHIRError)
}

// The GenericRead interface provides a generic read capability by passing the `resourceType` as string.
type GenericRead interface {
	Read(ctx context.Context, resourceType, id string) (model.Resource, FHIRError)
}

// The GenericUpdate interface provides a generic update capability.
//
// The persisted resource is returned.
type GenericUpdate interface {
	Update(ctx context.Context, resource model.Resource) (UpdateResult[model.Resource], FHIRError)
}

// UpdateResult is the result of an update operation.
//
// It contains the updated resource and a boolean indicating whether the resource was created or updated.
type UpdateResult[R model.Resource] struct {
	Resource R
	Created  bool
}

// The GenericSearch interface provides a generic search capability by passing the `resourceType` as string.
type GenericSearch interface {
	// GenericCapabilities is required because it includes the search capabilities (parameters etc.).
	GenericCapabilities
	Search(ctx context.Context, resourceType string, options search.Options) (search.Result, FHIRError)
}
