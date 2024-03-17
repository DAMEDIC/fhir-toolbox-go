package types

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
)

// Base StructureDefinition for Annotation Type: A  text note which also  contains information about who made the statement and when.
type Annotation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The individual responsible for making the annotation.
	Author r4.Element
	// Indicates when this particular annotation was made.
	Time *DateTime
	// The text of the annotation in markdown format.
	Text Markdown
}

func (s Annotation) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
