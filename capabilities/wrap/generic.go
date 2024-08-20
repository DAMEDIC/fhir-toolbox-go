package wrap

import (
	"fhir-toolbox/capabilities"
	genericR4 "fhir-toolbox/capabilities/wrap/gen/r4/generic"
	"fhir-toolbox/model"
)

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
