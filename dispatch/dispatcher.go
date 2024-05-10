package dispatch

import (
	"context"
	"fhir-toolbox/capabilities"
	"fhir-toolbox/model"

	dispatchR4 "fhir-toolbox/dispatch/gen/r4"
)

type Dispatcher struct {
	Read               func(ctx context.Context, api any, resourceType string, id string) (any, error)
	SearchCapabilities func(api any, resourceType string) (capabilities.SearchCapabilities, error)
	Search             func(ctx context.Context, api any, resourceType string, parameters capabilities.SearchParameters) ([]any, error)
}

func DispatcherFor[R model.Release]() Dispatcher {
	var r R
	switch any(r).(type) {
	case model.R4:
		return Dispatcher{
			Read:               dispatchR4.Read,
			SearchCapabilities: dispatchR4.SearchCapabilities,
			Search:             dispatchR4.Search,
		}
	default:
		panic("unsupported release")
	}
}
