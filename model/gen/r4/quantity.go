package r4

import "encoding/json"

// Base StructureDefinition for Quantity Type: A measured amount (or an amount that can potentially be measured). Note that measured amounts include amounts that are not precisely quantified, including amounts involving arbitrary units and floating currencies.
//
// Need to able to capture all sorts of measured values, even if the measured value are not precisely quantified. Values include exact measures such as 3.51g, customary units such as 3 tablets, and currencies such as $100.32USD.
type Quantity struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The value of the measured amount. The value includes an implicit precision in the presentation of the value.
	Value *Decimal
	// How the value should be understood and represented - whether the actual value is greater or less than the stated value due to measurement issues; e.g. if the comparator is "<" , then the real value is < stated value.
	Comparator *Code
	// A human-readable form of the unit.
	Unit *String
	// The identification of the system that provides the coded form of the unit.
	System *Uri
	// A computer processable form of the unit in some unit representation system.
	Code *Code
}
type jsonQuantity struct {
	Id                         *string           `json:"id,omitempty"`
	Extension                  []Extension       `json:"extension,omitempty"`
	Value                      *Decimal          `json:"value,omitempty"`
	ValuePrimitiveElement      *primitiveElement `json:"_value,omitempty"`
	Comparator                 *Code             `json:"comparator,omitempty"`
	ComparatorPrimitiveElement *primitiveElement `json:"_comparator,omitempty"`
	Unit                       *String           `json:"unit,omitempty"`
	UnitPrimitiveElement       *primitiveElement `json:"_unit,omitempty"`
	System                     *Uri              `json:"system,omitempty"`
	SystemPrimitiveElement     *primitiveElement `json:"_system,omitempty"`
	Code                       *Code             `json:"code,omitempty"`
	CodePrimitiveElement       *primitiveElement `json:"_code,omitempty"`
}

func (r Quantity) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Quantity) marshalJSON() jsonQuantity {
	m := jsonQuantity{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Value = r.Value
	if r.Value != nil && (r.Value.Id != nil || r.Value.Extension != nil) {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	m.Comparator = r.Comparator
	if r.Comparator != nil && (r.Comparator.Id != nil || r.Comparator.Extension != nil) {
		m.ComparatorPrimitiveElement = &primitiveElement{Id: r.Comparator.Id, Extension: r.Comparator.Extension}
	}
	m.Unit = r.Unit
	if r.Unit != nil && (r.Unit.Id != nil || r.Unit.Extension != nil) {
		m.UnitPrimitiveElement = &primitiveElement{Id: r.Unit.Id, Extension: r.Unit.Extension}
	}
	m.System = r.System
	if r.System != nil && (r.System.Id != nil || r.System.Extension != nil) {
		m.SystemPrimitiveElement = &primitiveElement{Id: r.System.Id, Extension: r.System.Extension}
	}
	m.Code = r.Code
	if r.Code != nil && (r.Code.Id != nil || r.Code.Extension != nil) {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	return m
}
func (r *Quantity) UnmarshalJSON(b []byte) error {
	var m jsonQuantity
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Quantity) unmarshalJSON(m jsonQuantity) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	r.Comparator = m.Comparator
	if m.ComparatorPrimitiveElement != nil {
		r.Comparator.Id = m.ComparatorPrimitiveElement.Id
		r.Comparator.Extension = m.ComparatorPrimitiveElement.Extension
	}
	r.Unit = m.Unit
	if m.UnitPrimitiveElement != nil {
		r.Unit.Id = m.UnitPrimitiveElement.Id
		r.Unit.Extension = m.UnitPrimitiveElement.Extension
	}
	r.System = m.System
	if m.SystemPrimitiveElement != nil {
		r.System.Id = m.SystemPrimitiveElement.Id
		r.System.Extension = m.SystemPrimitiveElement.Extension
	}
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	return nil
}
func (r Quantity) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
