package r4

import "encoding/json"

// Base StructureDefinition for boolean Type: Value of "true" or "false"
type Boolean struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The actual value
	Value *bool
}
type jsonBoolean struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Value     *bool       `json:"value,omitempty"`
}

func (r Boolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Value)
}
func (r *Boolean) UnmarshalJSON(b []byte) error {
	var value bool
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	*r = Boolean{Value: &value}
	return nil
}
