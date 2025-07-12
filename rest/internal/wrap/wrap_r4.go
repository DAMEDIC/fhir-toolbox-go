//go:build r4 || !(r4 || r4b || r5)

package wrap

import (
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	capabilitiesR4 "github.com/DAMEDIC/fhir-toolbox-go/capabilities/gen/r4"
)

func init() {
	genericR4 = func(api capabilities.ConcreteCapabilities) (capabilities.GenericCapabilities, error) {
		return capabilitiesR4.Generic{Concrete: api}, nil
	}
}
