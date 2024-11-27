package wrap

import (
	"fhir-toolbox/capabilities"
	capabilitiesR4 "fhir-toolbox/capabilities/gen/r4"
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
func Generic[R model.Release](api any) capabilitiesR4.Generic {
	// if already generic do not wrap it again
	generic, ok := api.(capabilitiesR4.Generic)
	if ok {
		return generic
	}

	var r R
	switch any(r).(type) {
	case model.R4:
		return capabilitiesR4.Generic{Concrete: api}
	default:
		panic("unsupported release")
	}
}

// The ConcreteR4 function wraps a generic API with concrete capabilities methods.
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
//	concreteAPI := wrap.ConcreteR4(genericAPI)
//	patient, err := concreteAPI.ReadPatient(ctx, id)
//
// ```
//
// The function is concrete vs. generic over the FHIR release, because it is not possible to return different
// interfaces depending on generic parameters.
func ConcreteR4(api capabilities.GenericAPI) capabilitiesR4.ConcreteAPI {
	return capabilitiesR4.Concrete{Generic: api}
}
