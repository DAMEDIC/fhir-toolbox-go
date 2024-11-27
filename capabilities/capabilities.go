package capabilities

import (
	"fhir-toolbox/capabilities/search"
)

type Capabilities struct {
	ReadInteractions   []string
	SearchCapabilities map[string]search.Capabilities
}
