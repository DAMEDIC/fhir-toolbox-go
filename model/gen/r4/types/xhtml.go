package types

import "encoding/json"

// Base StructureDefinition for xhtml Type
type Xhtml struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// Actual xhtml
	Value string
}

func (s Xhtml) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
