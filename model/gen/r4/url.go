package r4

import "encoding/json"

// Base StructureDefinition for url type: A URI that is a literal reference
type Url struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Primitive value for url
	Value *string
}

func (r Url) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Value)
}
func (r *Url) UnmarshalJSON(b []byte) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	*r = Url{Value: &value}
	return nil
}
