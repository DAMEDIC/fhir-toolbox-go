package r4

import "encoding/json"

// Base StructureDefinition for unsignedInt type: An integer with a value that is not negative (e.g. >= 0)
type UnsignedInt struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Primitive value for unsignedInt
	Value *uint32
}

func (r UnsignedInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Value)
}
func (r *UnsignedInt) UnmarshalJSON(b []byte) error {
	var value uint32
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	*r = UnsignedInt{Value: &value}
	return nil
}
