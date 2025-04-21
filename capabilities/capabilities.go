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
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
)

// Capabilities is a description of all capabilities that an implementation provides.
type Capabilities struct {
	CreateInteractions []string
	ReadInteractions   []string
	UpdateInteractions []string
	DeleteInteractions []string
	SearchCapabilities map[string]search.Capabilities
}
