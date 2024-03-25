package r4

import (
	"encoding/json"
	"fmt"
)

// Base StructureDefinition for UsageContext Type: Specifies clinical/business/etc. metadata that can be used to retrieve, index and/or categorize an artifact. This metadata can either be specific to the applicable population (e.g., age category, DRG) or the specific context of care (e.g., venue, care setting, provider of care).
//
// Consumers of the resource must be able to determine the intended applicability for the resource. Ideally, this information would be used programmatically to determine when and how it should be incorporated or exposed.
type UsageContext struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// A code that identifies the type of context being specified by this usage context.
	Code Coding
	// A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
	Value isUsageContextValue
}
type isUsageContextValue interface {
	isUsageContextValue()
}

func (r CodeableConcept) isUsageContextValue() {}
func (r Quantity) isUsageContextValue()        {}
func (r Range) isUsageContextValue()           {}
func (r Reference) isUsageContextValue()       {}

type jsonUsageContext struct {
	Id                   *string          `json:"id,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	Code                 Coding           `json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range           `json:"valueRange,omitempty"`
	ValueReference       *Reference       `json:"valueReference,omitempty"`
}

func (r UsageContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r UsageContext) marshalJSON() jsonUsageContext {
	m := jsonUsageContext{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Code = r.Code
	switch v := r.Value.(type) {
	case CodeableConcept:
		m.ValueCodeableConcept = &v
	case *CodeableConcept:
		m.ValueCodeableConcept = v
	case Quantity:
		m.ValueQuantity = &v
	case *Quantity:
		m.ValueQuantity = v
	case Range:
		m.ValueRange = &v
	case *Range:
		m.ValueRange = v
	case Reference:
		m.ValueReference = &v
	case *Reference:
		m.ValueReference = v
	}
	return m
}
func (r *UsageContext) UnmarshalJSON(b []byte) error {
	var m jsonUsageContext
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *UsageContext) unmarshalJSON(m jsonUsageContext) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Code = m.Code
	if m.ValueCodeableConcept != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCodeableConcept
		r.Value = v
	}
	if m.ValueQuantity != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueQuantity
		r.Value = v
	}
	if m.ValueRange != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueRange
		r.Value = v
	}
	if m.ValueReference != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueReference
		r.Value = v
	}
	return nil
}
func (r UsageContext) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
