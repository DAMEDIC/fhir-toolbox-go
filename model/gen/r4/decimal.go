package r4

// Base StructureDefinition for decimal Type: A rational number with implicit precision
type Decimal struct {
	// unique id for the element within a resource (for internal references)
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The actual value
	Value *string
}
type jsonDecimal struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Value     *string     `json:"value,omitempty"`
}

func (r Decimal) MarshalJSON() ([]byte, error) {
	return []byte(*r.Value), nil
}
func (r *Decimal) UnmarshalJSON(b []byte) error {
	var value string
	value = string(b)
	*r = Decimal{Value: &value}
	return nil
}
