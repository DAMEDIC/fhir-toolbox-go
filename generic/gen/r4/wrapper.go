package genericR4

import capabilities "fhir-toolbox/capabilities"

type wrapper struct {
	api any
}

func Wrap(api any) capabilities.GenericAPI {
	return wrapper{api}
}
