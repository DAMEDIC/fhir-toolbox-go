package wrap

import (
	"fhir-toolbox/capabilities"
	concreteR4 "fhir-toolbox/capabilities/wrap/gen/r4/concrete"
	"fhir-toolbox/model"
)

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
