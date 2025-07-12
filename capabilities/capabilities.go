// Package capabilities provides interfaces modeling capabilities.
// This flexible architecture allows different use cases, such as
//
//   - building FHIR® facades to legacy systems by implementing a custom backend
//   - using this library as a FHIR® client (by leveraging a - still to be build - REST backend)
//
// # Concrete vs. Generic API
//
// The library provides two API styles.
// The concrete API:
//
//	func (a myAPI) ReadPatient(ctx context.Context, id string) (r4.Patient, error) {}
//
//	func (a myAPI) SearchPatient(ctx context.Context, options search.Options) (search.Result, error) {}
//
// and the generic API:
//
//	func (a myAPI) Read(ctx context.Context, resourceType, id string) (r4.Patient, error) {}
//
//	func (a myAPI) Search(ctx context.Context, resourceType string, options search.Options) (search.Result, error) {}
//
// You can implement your custom backend or client either way.
// The concrete API is ideal for building custom FHIR® facades where a limited set of resources is used.
// The generic API is better suited for e.g. building FHIR® clients or standalone FHIR® servers.
//
// # Interoperability
//
// The [capabilities/wrap] package provides wrapper functions to wrap a concrete into the generic API:
//
//	genericAPI := wrap.Generic[model.R4](concreteAPI)
//
// and vice versa:
//
//	concreteAPI := wrap.ConcreteR4(genericAPI)
package capabilities

import (
	"context"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
)

// The ConcreteCapabilities interface allows concrete implementations to provide a base CapabilityStatement
// that will be enhanced with the detected concrete capabilities. This is useful for setting implementation
// details, base URLs, and other metadata that should be preserved in the final CapabilityStatement.
type ConcreteCapabilities interface {
	CapabilityBase(ctx context.Context) (basic.CapabilityStatement, error)
}
