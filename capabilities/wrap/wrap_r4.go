//go:build r4 || !(r4 || r4b || r5)

package wrap

import (
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	capabilitiesR4 "github.com/DAMEDIC/fhir-toolbox-go/capabilities/gen/r4"
)

func init() {
	genericR4 = func(api any) (capabilities.GenericAPI, error) {
		return capabilitiesR4.Generic{Concrete: api}, nil
	}
}

// The ConcreteR4 function wraps a generic API with concrete capabilities methods.
//
// # Example
//
// A given a generic implementation
//
//	func (a myAPI) Read(ctx context.Context, resourceType, id string) (r4.Patient, capabilities.FHIRError) {}
//
// can be wrapped and called by its concrete methods
//
//	concreteAPI := wrap.ConcreteR4(genericAPI)
//	patient, err := concreteAPI.ReadPatient(ctx, id)
//
// The function is concrete vs. generic over the FHIR release, because it is not possible to return different
// interfaces depending on generic parameters.
func ConcreteR4(api capabilities.GenericAPI) capabilitiesR4.ConcreteAPI {
	return capabilitiesR4.Concrete{Generic: api}
}
