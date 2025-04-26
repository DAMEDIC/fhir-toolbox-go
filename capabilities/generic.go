package capabilities

import (
	"context"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
)

// The GenericCapabilities interface provides a generic capabilities method that returns all capabilities of the underlying concrete implementation.
type GenericCapabilities interface {
	AllCapabilities(ctx context.Context) (Capabilities, error)
}

// The GenericCreate interface provides a generic create capability.
//
// The persisted resource is returned.
type GenericCreate interface {
	Create(ctx context.Context, resource model.Resource) (model.Resource, error)
}

// The GenericRead interface provides a generic read capability by passing the `resourceType` as string.
type GenericRead interface {
	Read(ctx context.Context, resourceType, id string) (model.Resource, error)
}

// The GenericUpdate interface provides a generic update capability.
//
// The persisted resource is returned.
type GenericUpdate interface {
	Update(ctx context.Context, resource model.Resource) (UpdateResult[model.Resource], error)
}

// UpdateResult is the result of an update operation.
//
// It contains the updated resource and a boolean indicating whether the resource was created or updated.
type UpdateResult[R model.Resource] struct {
	Resource R
	// Created indicates whether the resource was newly created (true) or an existing resource was updated (false).
	Created bool
}

// The GenericDelete interface provides a generic deletion capability by passing the `resourceType` as string.
type GenericDelete interface {
	Delete(ctx context.Context, resourceType, id string) error
}

// The GenericSearch interface provides a generic search capability by passing the `resourceType` as string.
type GenericSearch interface {
	// GenericCapabilities is required because it includes the search capabilities (parameters etc.).
	GenericCapabilities
	Search(ctx context.Context, resourceType string, options search.Options) (search.Result, error)
}

// The GenericAPI interface combines all generic interfaces to provide a complete API.
type GenericAPI interface {
	GenericCapabilities
	GenericCreate
	GenericRead
	GenericUpdate
	GenericDelete
	GenericSearch
}
