package wrap

import (
	"fhir-toolbox/capabilities"
	genericR4 "fhir-toolbox/capabilities/wrap/gen/r4/generic"
	"fhir-toolbox/model"
)

// The Generic function wraps a concrete API with generic capabilities methods.
//
// # Example
//
// A given a concrete implementation
// ```Go
//
//	func (a myAPI) ReadPatient(ctx context.Context, id string) (r4.Patient, capabilities.FHIRError) {}
//
// ```
// can be wrapped and called by its concrete methods
// ```Go
//
//	genericAPI := wrap.Generic[model.R4](concreteAPI)
//	patient, err := genericAPI.Read(ctx, "Patient", id)
//
// ```
func Generic[R model.Release](api any) capabilities.GenericAPI {
	generic, ok := api.(capabilities.GenericAPI)
	if ok {
		return generic
	}

	var r R
	switch any(r).(type) {
	case model.R4:
		return genericR4.InternalWrapper{api}
	default:
		panic("unsupported release")
	}
}
