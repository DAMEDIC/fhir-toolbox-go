package capabilities

import (
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
)

type Capabilities struct {
	ReadInteractions   []string
	SearchCapabilities map[string]search.Capabilities
}
