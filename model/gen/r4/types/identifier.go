package types

import "encoding/json"

// Base StructureDefinition for Identifier Type: An identifier - identifies some entity uniquely and unambiguously. Typically this is used for business identifiers.
//
// Need to be able to identify things with confidence and be sure that the identification is not subject to misinterpretation.
type Identifier struct {
	// Time period during which identifier is/was valid for use.
	Period *Period
	// Organization that issued/manages the identifier.
	Assigner *Reference
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The purpose of this identifier.
	Use *Code
	// A coded type for the identifier that can be used to determine which identifier to use for a specific purpose.
	Type *CodeableConcept
	// Establishes the namespace for the value - that is, a URL that describes a set values that are unique.
	System *Uri
	// The portion of the identifier typically relevant to the user and which is unique within the context of the system.
	Value *String
}

func (s Identifier) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
