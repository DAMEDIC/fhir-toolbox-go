//go:build r5 || !(r4 || r4b || r5)

package wrap

import (
	"fhir-toolbox/capabilities"
	capabilitiesR5 "fhir-toolbox/capabilities/gen/r5"
)

func init() {
	genericR5 = func(api any) (capabilities.GenericAPI, error) {
		return capabilitiesR5.Generic{Concrete: api}, nil
	}
}

// The ConcreteR5 function wraps a generic API with concrete capabilities methods.
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
//	concreteAPI := wrap.ConcreteR5(genericAPI)
//	patient, err := concreteAPI.ReadPatient(ctx, id)
//
// ```
//
// The function is concrete vs. generic over the FHIR release, because it is not possible to return different
// interfaces depending on generic parameters.
func ConcreteR5(api capabilities.GenericAPI) capabilitiesR5.ConcreteAPI {
	return capabilitiesR5.Concrete{Generic: api}
}
