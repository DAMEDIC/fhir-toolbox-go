package r4

import "encoding/json"

// Base StructureDefinition for time Type: A time during the day, with no date specified
type Time struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The actual value
	Value *string
}
type jsonTime struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Value     *string     `json:"value,omitempty"`
}

func (r Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Value)
}
func (r *Time) UnmarshalJSON(b []byte) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	*r = Time{Value: &value}
	return nil
}
