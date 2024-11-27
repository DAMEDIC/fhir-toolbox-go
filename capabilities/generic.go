package capabilities

import (
	"context"
	"fhir-toolbox/capabilities/search"
	"fhir-toolbox/model"
)

// A GenericAPI irrespective of the capabilities of the underlying concrete implementation.
//
// This can be used to build adapters, clients or generic servers.
// See [fhir-toolbox/capabilities/wrap] for conversion between concrete to generic APIs and vice versa.
type GenericAPI interface {
	GenericRead
	GenericSearch

	AllCapabilities() Capabilities
}

// The GenericRead interface provides a generic read capability by passing the `resourceType` as string.
type GenericRead interface {
	Read(ctx context.Context, resourceType string, id string) (model.Resource, FHIRError)
}

// The GenericSearch interface provides a generic search capability by passing the `resourceType` as string.
type GenericSearch interface {
	SearchCapabilities(resourceType string) (search.Capabilities, FHIRError)
	Search(ctx context.Context, resourceType string, options search.Options) (search.Result, FHIRError)
}
