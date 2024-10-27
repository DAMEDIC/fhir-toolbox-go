package r4

import "encoding/json"

// Base StructureDefinition for uri Type: String of characters used to identify a name or a resource
type Uri struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The actual value
	Value *string
}

func (r Uri) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Value)
}
func (r *Uri) UnmarshalJSON(b []byte) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	*r = Uri{Value: &value}
	return nil
}
