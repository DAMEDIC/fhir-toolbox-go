package types

import "encoding/json"

// Base StructureDefinition for HumanName Type: A human's name with the ability to identify parts and usage.
//
// Need to be able to record names, along with notes about their use.
type HumanName struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Specifies the entire name as it should be displayed e.g. on an application UI. This may be provided instead of or as well as the specific parts.
	Text *String
	// Given name.
	Given []String
	// Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the start of the name.
	Prefix []String
	// Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the end of the name.
	Suffix []String
	// Indicates the period of time when this name was valid for the named person.
	Period *Period
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Identifies the purpose for this name.
	Use *Code
	// The part of a name that links to the genealogy. In some cultures (e.g. Eritrea) the family name of a son is the first name of his father.
	Family *String
}

func (s HumanName) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
