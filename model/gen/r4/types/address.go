package types

import "encoding/json"

// Base StructureDefinition for Address Type: An address expressed using postal conventions (as opposed to GPS or other location definition formats).  This data type may be used to convey addresses for use in delivering mail as well as for visiting locations which might not be valid for mail delivery.  There are a variety of postal address formats defined around the world.
//
// Need to be able to record postal addresses, along with notes about their use.
type Address struct {
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Distinguishes between physical addresses (those you can visit) and mailing addresses (e.g. PO Boxes and care-of addresses). Most addresses are both.
	Type *Code
	// This component contains the house number, apartment number, street name, street direction,  P.O. Box number, delivery hints, and similar address information.
	Line []String
	// The name of the administrative area (county).
	District *String
	// A postal code designating a region defined by the postal service.
	PostalCode *String
	// Country - a nation as commonly understood or generally accepted.
	Country *String
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// The purpose of this address.
	Use *Code
	// Specifies the entire address as it should be displayed e.g. on a postal label. This may be provided instead of or as well as the specific parts.
	Text *String
	// The name of the city, town, suburb, village or other community or delivery center.
	City *String
	// Sub-unit of a country with limited sovereignty in a federally organized country. A code may be used if codes are in common use (e.g. US 2 letter state codes).
	State *String
	// Time period when address was/is in use.
	Period *Period
}

func (s Address) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
