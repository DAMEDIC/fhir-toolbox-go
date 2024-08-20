package wrap

import (
	"fhir-toolbox/capabilities"
	capabilitiesR4 "fhir-toolbox/capabilities/gen/r4"
	concreteR4 "fhir-toolbox/capabilities/wrap/gen/r4/concrete"
)

func ConcreteR4(api capabilities.GenericAPI) capabilitiesR4.FullAPI {
	return concreteR4.InternalWrapper{api}
}
