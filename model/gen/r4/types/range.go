package types

import "encoding/json"

// Base StructureDefinition for Range Type: A set of ordered Quantities defined by a low and high limit.
//
// Need to be able to specify ranges of values.
type Range struct {
	// The low limit. The boundary is inclusive.
	Low *Quantity
	// The high limit. The boundary is inclusive.
	High *Quantity
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
}

func (s Range) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
