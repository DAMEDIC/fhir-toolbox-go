//go:build r4b || !(r4 || r4b || r5)

package wrap

import (
	"fhir-toolbox/capabilities"
	capabilitiesR4B "fhir-toolbox/capabilities/gen/r4b"
)

func init() {
	genericR4B = func(api any) (capabilities.GenericAPI, error) {
		return capabilitiesR4B.Generic{Concrete: api}, nil
	}
}

// The ConcreteR4B function wraps a generic API with concrete capabilities methods.
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
//	concreteAPI := wrap.ConcreteR4B(genericAPI)
//	patient, err := concreteAPI.ReadPatient(ctx, id)
//
// ```
//
// The function is concrete vs. generic over the FHIR release, because it is not possible to return different
// interfaces depending on generic parameters.
func ConcreteR4B(api capabilities.GenericAPI) capabilitiesR4B.ConcreteAPI {
	return capabilitiesR4B.Concrete{Generic: api}
}
