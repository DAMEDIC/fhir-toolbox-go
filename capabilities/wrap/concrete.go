package wrap

import (
	"fhir-toolbox/capabilities"
	concreteR4 "fhir-toolbox/capabilities/wrap/gen/r4/concrete"
	"fhir-toolbox/model"
)

// The Concrete function wraps a generic API with concrete capabilities methods.
//
// # Example
//
// A given a generic implementation
// ```Go
//
//	func (a myAPI) Read(ctx context.Context, resourceType, id string) (r4.Patient, capabilities.FHIRError) {}
//
// ```
// can be wrapped and called by its concrete methods
// ```Go
//
//	concreteAPI := wrap.Concrete[model.R4](genericAPI)
//	patient, err := concreteAPI.ReadPatient(ctx, id)
//
// ```
func Concrete[R model.Release](api capabilities.GenericAPI) capabilities.GenericAPI {
	generic, ok := api.(capabilities.GenericAPI)
	if ok {
		return generic
	}

	var r R
	switch any(r).(type) {
	case model.R4:
		return concreteR4.InternalWrapper{api}
	default:
		panic("unsupported release")
	}
}
