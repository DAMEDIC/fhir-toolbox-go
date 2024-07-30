package generic

import (
	"fhir-toolbox/capabilities"
	genericR4 "fhir-toolbox/generic/gen/r4"
	"fhir-toolbox/model"
)

func Wrap[R model.Release](api any) capabilities.GenericAPI {
	generic, ok := api.(capabilities.GenericAPI)
	if ok {
		return generic
	}

	var r R
	switch any(r).(type) {
	case model.R4:
		return genericR4.Wrap(api)
	default:
		panic("unsupported release")
	}
}
