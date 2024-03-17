package types

import "encoding/json"

// Base StructureDefinition for oid type: An OID represented as a URI
type Oid struct {
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Primitive value for oid
	Value *string
	// unique id for the element within a resource (for internal references)
	Id *string
}

func (s Oid) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
