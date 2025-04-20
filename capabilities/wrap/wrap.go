// Package wrap provides conversion between concrete and generic APIs.
//
// To wrap a concrete into the generic API:
//
//	genericAPI := wrap.Generic[model.R4](concreteAPI)
//
// and vice versa:
//
//	concreteAPI := wrap.ConcreteR4(genericAPI)
package wrap

import (
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"strings"
)

var genericR4 = func(api any) (capabilities.GenericAPI, error) { return nil, disabledErr[model.R4]() }
var genericR4B = func(api any) (capabilities.GenericAPI, error) { return nil, disabledErr[model.R4B]() }
var genericR5 = func(api any) (capabilities.GenericAPI, error) { return nil, disabledErr[model.R5]() }

func disabledErr[R model.Release]() error {
	r := model.ReleaseName[R]()

	return fmt.Errorf(
		"release %s disabled by build tag; remove all build tags or add %s",
		r, strings.ToLower(r),
	)
}

// The Generic function wraps a concrete API with generic capabilities methods.
//
// # Example
//
// A given a concrete implementation
//
//	func (a myAPI) ReadPatient(ctx context.Context, id string) (r4.Patient, capabilities.FHIRError) {}
//
// can be wrapped and called by its concrete methods
//
//	genericAPI := wrap.Generic[model.R4](concreteAPI)
//	patient, err := genericAPI.Read(ctx, "Patient", id)
func Generic[R model.Release](api any) (capabilities.GenericAPI, error) {
	// if already generic do not wrap it again
	generic, ok := api.(capabilities.GenericAPI)
	if ok {
		return generic, nil
	}

	var r R
	switch any(r).(type) {
	case model.R4:
		return genericR4(api)
	case model.R4B:
		return genericR4B(api)
	case model.R5:
		return genericR5(api)
	default:
		// This should never happen as long as we control all implementations of the Release interface.
		// This is achieved by sealing the interface. See the interface definition for more information.
		panic("unsupported release")
	}
}
