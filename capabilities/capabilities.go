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
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/create"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/deletion"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/read"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/update"
)

// Capabilities is a description of all capabilities that an implementation provides.
type Capabilities struct {
	// Create is a list of supported resources.
	Create map[string]create.Capabilities
	// Read is a list of supported resources.
	Read map[string]read.Capabilities
	// Update capabilities, indexed by the resource type.
	Update map[string]update.Capabilities
	// Delete is a list of supported resources.
	Delete map[string]deletion.Capabilities
	// Search capabilities, indexed by the resource type.
	Search map[string]search.Capabilities[search.Parameter]
}
