package r4

import "encoding/json"

// Base StructureDefinition for integer Type: A whole number
type Integer struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The actual value
	Value *int32
}

func (r Integer) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Value)
}
func (r *Integer) UnmarshalJSON(b []byte) error {
	var value int32
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	*r = Integer{Value: &value}
	return nil
}
