package r4

import (
	"encoding/json"
	"fmt"
)

// Base StructureDefinition for SubstanceAmount Type: Chemical substances are a single substance type whose primary defining element is the molecular structure. Chemical substances shall be defined on the basis of their complete covalent molecular structure; the presence of a salt (counter-ion) and/or solvates (water, alcohols) is also captured. Purity, grade, physical form or particle size are not taken into account in the definition of a chemical substance or in the assignment of a Substance ID.
type SubstanceAmount struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Used to capture quantitative values for a variety of elements. If only limits are given, the arithmetic mean would be the average. If only a single definite value for a given element is given, it would be captured in this field.
	Amount isSubstanceAmountAmount
	// Most elements that require a quantitative value will also have a field called amount type. Amount type should always be specified because the actual value of the amount is often dependent on it. EXAMPLE: In capturing the actual relative amounts of substances or molecular fragments it is essential to indicate whether the amount refers to a mole ratio or weight ratio. For any given element an effort should be made to use same the amount type for all related definitional elements.
	AmountType *CodeableConcept
	// A textual comment on a numeric value.
	AmountText *String
	// Reference range of possible or expected values.
	ReferenceRange *SubstanceAmountReferenceRange
}
type isSubstanceAmountAmount interface {
	isSubstanceAmountAmount()
}

func (r Quantity) isSubstanceAmountAmount() {}
func (r Range) isSubstanceAmountAmount()    {}
func (r String) isSubstanceAmountAmount()   {}

type jsonSubstanceAmount struct {
	Id                           *string                        `json:"id,omitempty"`
	Extension                    []Extension                    `json:"extension,omitempty"`
	ModifierExtension            []Extension                    `json:"modifierExtension,omitempty"`
	AmountQuantity               *Quantity                      `json:"amountQuantity,omitempty"`
	AmountRange                  *Range                         `json:"amountRange,omitempty"`
	AmountString                 *String                        `json:"amountString,omitempty"`
	AmountStringPrimitiveElement *primitiveElement              `json:"_amountString,omitempty"`
	AmountType                   *CodeableConcept               `json:"amountType,omitempty"`
	AmountText                   *String                        `json:"amountText,omitempty"`
	AmountTextPrimitiveElement   *primitiveElement              `json:"_amountText,omitempty"`
	ReferenceRange               *SubstanceAmountReferenceRange `json:"referenceRange,omitempty"`
}

func (r SubstanceAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceAmount) marshalJSON() jsonSubstanceAmount {
	m := jsonSubstanceAmount{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Amount.(type) {
	case Quantity:
		m.AmountQuantity = &v
	case *Quantity:
		m.AmountQuantity = v
	case Range:
		m.AmountRange = &v
	case *Range:
		m.AmountRange = v
	case String:
		m.AmountString = &v
		if v.Id != nil || v.Extension != nil {
			m.AmountStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.AmountString = v
		if v.Id != nil || v.Extension != nil {
			m.AmountStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.AmountType = r.AmountType
	m.AmountText = r.AmountText
	if r.AmountText != nil && (r.AmountText.Id != nil || r.AmountText.Extension != nil) {
		m.AmountTextPrimitiveElement = &primitiveElement{Id: r.AmountText.Id, Extension: r.AmountText.Extension}
	}
	m.ReferenceRange = r.ReferenceRange
	return m
}
func (r *SubstanceAmount) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceAmount
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceAmount) unmarshalJSON(m jsonSubstanceAmount) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.AmountQuantity != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountQuantity
		r.Amount = v
	}
	if m.AmountRange != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountRange
		r.Amount = v
	}
	if m.AmountString != nil || m.AmountStringPrimitiveElement != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountString
		if m.AmountStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.AmountStringPrimitiveElement.Id
			v.Extension = m.AmountStringPrimitiveElement.Extension
		}
		r.Amount = v
	}
	r.AmountType = m.AmountType
	r.AmountText = m.AmountText
	if m.AmountTextPrimitiveElement != nil {
		r.AmountText.Id = m.AmountTextPrimitiveElement.Id
		r.AmountText.Extension = m.AmountTextPrimitiveElement.Extension
	}
	r.ReferenceRange = m.ReferenceRange
	return nil
}
func (r SubstanceAmount) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Reference range of possible or expected values.
type SubstanceAmountReferenceRange struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Lower limit possible or expected.
	LowLimit *Quantity
	// Upper limit possible or expected.
	HighLimit *Quantity
}
type jsonSubstanceAmountReferenceRange struct {
	Id        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	LowLimit  *Quantity   `json:"lowLimit,omitempty"`
	HighLimit *Quantity   `json:"highLimit,omitempty"`
}

func (r SubstanceAmountReferenceRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceAmountReferenceRange) marshalJSON() jsonSubstanceAmountReferenceRange {
	m := jsonSubstanceAmountReferenceRange{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.LowLimit = r.LowLimit
	m.HighLimit = r.HighLimit
	return m
}
func (r *SubstanceAmountReferenceRange) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceAmountReferenceRange
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceAmountReferenceRange) unmarshalJSON(m jsonSubstanceAmountReferenceRange) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.LowLimit = m.LowLimit
	r.HighLimit = m.HighLimit
	return nil
}
func (r SubstanceAmountReferenceRange) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
