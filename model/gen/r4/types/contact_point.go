package types

import "encoding/json"

// Base StructureDefinition for ContactPoint Type: Details for all kinds of technology mediated contact points for a person or organization, including telephone, email, etc.
//
// Need to track phone, fax, mobile, sms numbers, email addresses, twitter tags, etc.
type ContactPoint struct {
	// Identifies the purpose for the contact point.
	Use *Code
	// Specifies a preferred order in which to use a set of contacts. ContactPoints with lower rank values are more preferred than those with higher rank values.
	Rank *PositiveInt
	// Time period when the contact point was/is in use.
	Period *Period
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Telecommunications form for contact point - what communications system is required to make use of the contact.
	System *Code
	// The actual contact point details, in a form that is meaningful to the designated communication system (i.e. phone number or email address).
	Value *String
}

func (s ContactPoint) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
