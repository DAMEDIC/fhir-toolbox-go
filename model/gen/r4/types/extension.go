package types

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
)

// Base StructureDefinition for Extension Type: Optional Extension Element - found in all resources.
//
// The ability to add extensions in a structured way is what keeps FHIR resources simple.
type Extension struct {
	// Source of the definition for the extension code - a logical name or a URL.
	Url string
	// Value of extension - must be one of a constrained set of the data types (see [Extensibility](extensibility.html) for a list).
	Value r4.Element
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
}

func (s Extension) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
