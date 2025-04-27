//go:build r4b || !(r4 || r4b || r5)

package wrap

import (
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	capabilitiesR4B "github.com/DAMEDIC/fhir-toolbox-go/capabilities/gen/r4b"
)

func init() {
	genericR4B = func(api any) (capabilities.GenericCapabilities, error) {
		return capabilitiesR4B.Generic{Concrete: api}, nil
	}
}
