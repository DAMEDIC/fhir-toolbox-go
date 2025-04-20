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
	GenericRead
	GenericSearch
	AllCapabilities(ctx context.Context) (Capabilities, FHIRError)
}

// The GenericRead interface provides a generic read capability by passing the `resourceType` as string.
type GenericRead interface {
	Read(ctx context.Context, resourceType, id string) (model.Resource, FHIRError)
}

// The GenericSearch interface provides a generic search capability by passing the `resourceType` as string.
type GenericSearch interface {
	// Generic SearchCapabilities is not required as it can be deduced from AllCapabilities.

	Search(ctx context.Context, resourceType string, options search.Options) (search.Result, FHIRError)
}
