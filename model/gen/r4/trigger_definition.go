package r4

import (
	"encoding/json"
	"fmt"
)

// Base StructureDefinition for TriggerDefinition Type: A description of a triggering event. Triggering events can be named events, data events, or periodic, as determined by the type element.
type TriggerDefinition struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The type of triggering event.
	Type Code
	// A formal name for the event. This may be an absolute URI that identifies the event formally (e.g. from a trigger registry), or a simple relative URI that identifies the event in a local context.
	Name *String
	// The timing of the event (if this is a periodic trigger).
	Timing isTriggerDefinitionTiming
	// The triggering data of the event (if this is a data trigger). If more than one data is requirement is specified, then all the data requirements must be true.
	Data []DataRequirement
	// A boolean-valued expression that is evaluated in the context of the container of the trigger definition and returns whether or not the trigger fires.
	Condition *Expression
}
type isTriggerDefinitionTiming interface {
	isTriggerDefinitionTiming()
}

func (r Timing) isTriggerDefinitionTiming()    {}
func (r Reference) isTriggerDefinitionTiming() {}
func (r Date) isTriggerDefinitionTiming()      {}
func (r DateTime) isTriggerDefinitionTiming()  {}

type jsonTriggerDefinition struct {
	Id                             *string           `json:"id,omitempty"`
	Extension                      []Extension       `json:"extension,omitempty"`
	Type                           Code              `json:"type,omitempty"`
	TypePrimitiveElement           *primitiveElement `json:"_type,omitempty"`
	Name                           *String           `json:"name,omitempty"`
	NamePrimitiveElement           *primitiveElement `json:"_name,omitempty"`
	TimingTiming                   *Timing           `json:"timingTiming,omitempty"`
	TimingReference                *Reference        `json:"timingReference,omitempty"`
	TimingDate                     *Date             `json:"timingDate,omitempty"`
	TimingDatePrimitiveElement     *primitiveElement `json:"_timingDate,omitempty"`
	TimingDateTime                 *DateTime         `json:"timingDateTime,omitempty"`
	TimingDateTimePrimitiveElement *primitiveElement `json:"_timingDateTime,omitempty"`
	Data                           []DataRequirement `json:"data,omitempty"`
	Condition                      *Expression       `json:"condition,omitempty"`
}

func (r TriggerDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TriggerDefinition) marshalJSON() jsonTriggerDefinition {
	m := jsonTriggerDefinition{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	switch v := r.Timing.(type) {
	case Timing:
		m.TimingTiming = &v
	case *Timing:
		m.TimingTiming = v
	case Reference:
		m.TimingReference = &v
	case *Reference:
		m.TimingReference = v
	case Date:
		m.TimingDate = &v
		if v.Id != nil || v.Extension != nil {
			m.TimingDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		m.TimingDate = v
		if v.Id != nil || v.Extension != nil {
			m.TimingDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case DateTime:
		m.TimingDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.TimingDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.TimingDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.TimingDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.Data = r.Data
	m.Condition = r.Condition
	return m
}
func (r *TriggerDefinition) UnmarshalJSON(b []byte) error {
	var m jsonTriggerDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TriggerDefinition) unmarshalJSON(m jsonTriggerDefinition) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	if m.TimingTiming != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingTiming
		r.Timing = v
	}
	if m.TimingReference != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingReference
		r.Timing = v
	}
	if m.TimingDate != nil || m.TimingDatePrimitiveElement != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingDate
		if m.TimingDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.TimingDatePrimitiveElement.Id
			v.Extension = m.TimingDatePrimitiveElement.Extension
		}
		r.Timing = v
	}
	if m.TimingDateTime != nil || m.TimingDateTimePrimitiveElement != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingDateTime
		if m.TimingDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.TimingDateTimePrimitiveElement.Id
			v.Extension = m.TimingDateTimePrimitiveElement.Extension
		}
		r.Timing = v
	}
	r.Data = m.Data
	r.Condition = m.Condition
	return nil
}
func (r TriggerDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
