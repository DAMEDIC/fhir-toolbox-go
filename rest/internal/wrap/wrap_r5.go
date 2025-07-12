//go:build r5 || !(r4 || r4b || r5)

package wrap

import (
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	capabilitiesR5 "github.com/DAMEDIC/fhir-toolbox-go/capabilities/gen/r5"
)

func init() {
	genericR5 = func(api capabilities.ConcreteCapabilities) (capabilities.GenericCapabilities, error) {
		return capabilitiesR5.Generic{Concrete: api}, nil
	}
}
