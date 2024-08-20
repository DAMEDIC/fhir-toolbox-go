package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A record of a device being used by a patient where the record is the result of a report from the patient or another clinician.
type DeviceUseStatement struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *Id
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *Uri
	// The base language in which the resource is written.
	Language *Code
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *Narrative
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An external identifier for this statement such as an IRI.
	Identifier []Identifier
	// A plan, proposal or order that is fulfilled in whole or in part by this DeviceUseStatement.
	BasedOn []Reference
	// A code representing the patient or other source's judgment about the state of the device used that this statement is about.  Generally this will be active or completed.
	Status Code
	// The patient who used the device.
	Subject Reference
	// Allows linking the DeviceUseStatement to the underlying Request, or to other information that supports or is used to derive the DeviceUseStatement.
	DerivedFrom []Reference
	// How often the device was used.
	Timing isDeviceUseStatementTiming
	// The time at which the statement was made/recorded.
	RecordedOn *DateTime
	// Who reported the device was being used by the patient.
	Source *Reference
	// The details of the device used.
	Device Reference
	// Reason or justification for the use of the device.
	ReasonCode []CodeableConcept
	// Indicates another resource whose existence justifies this DeviceUseStatement.
	ReasonReference []Reference
	// Indicates the anotomic location on the subject's body where the device was used ( i.e. the target).
	BodySite *CodeableConcept
	// Details about the device statement that were not represented at all or sufficiently in one of the attributes provided in a class. These may include for example a comment, an instruction, or a note associated with the statement.
	Note []Annotation
}

func (r DeviceUseStatement) ResourceType() string {
	return "DeviceUseStatement"
}
func (r DeviceUseStatement) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isDeviceUseStatementTiming interface {
	isDeviceUseStatementTiming()
}

func (r Timing) isDeviceUseStatementTiming()   {}
func (r Period) isDeviceUseStatementTiming()   {}
func (r DateTime) isDeviceUseStatementTiming() {}

type jsonDeviceUseStatement struct {
	ResourceType                   string              `json:"resourceType"`
	Id                             *Id                 `json:"id,omitempty"`
	IdPrimitiveElement             *primitiveElement   `json:"_id,omitempty"`
	Meta                           *Meta               `json:"meta,omitempty"`
	ImplicitRules                  *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement  *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                       *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement       *primitiveElement   `json:"_language,omitempty"`
	Text                           *Narrative          `json:"text,omitempty"`
	Contained                      []ContainedResource `json:"contained,omitempty"`
	Extension                      []Extension         `json:"extension,omitempty"`
	ModifierExtension              []Extension         `json:"modifierExtension,omitempty"`
	Identifier                     []Identifier        `json:"identifier,omitempty"`
	BasedOn                        []Reference         `json:"basedOn,omitempty"`
	Status                         Code                `json:"status,omitempty"`
	StatusPrimitiveElement         *primitiveElement   `json:"_status,omitempty"`
	Subject                        Reference           `json:"subject,omitempty"`
	DerivedFrom                    []Reference         `json:"derivedFrom,omitempty"`
	TimingTiming                   *Timing             `json:"timingTiming,omitempty"`
	TimingPeriod                   *Period             `json:"timingPeriod,omitempty"`
	TimingDateTime                 *DateTime           `json:"timingDateTime,omitempty"`
	TimingDateTimePrimitiveElement *primitiveElement   `json:"_timingDateTime,omitempty"`
	RecordedOn                     *DateTime           `json:"recordedOn,omitempty"`
	RecordedOnPrimitiveElement     *primitiveElement   `json:"_recordedOn,omitempty"`
	Source                         *Reference          `json:"source,omitempty"`
	Device                         Reference           `json:"device,omitempty"`
	ReasonCode                     []CodeableConcept   `json:"reasonCode,omitempty"`
	ReasonReference                []Reference         `json:"reasonReference,omitempty"`
	BodySite                       *CodeableConcept    `json:"bodySite,omitempty"`
	Note                           []Annotation        `json:"note,omitempty"`
}

func (r DeviceUseStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceUseStatement) marshalJSON() jsonDeviceUseStatement {
	m := jsonDeviceUseStatement{}
	m.ResourceType = "DeviceUseStatement"
	m.Id = r.Id
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	m.ImplicitRules = r.ImplicitRules
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	m.Language = r.Language
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.Text = r.Text
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.BasedOn = r.BasedOn
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Subject = r.Subject
	m.DerivedFrom = r.DerivedFrom
	switch v := r.Timing.(type) {
	case Timing:
		m.TimingTiming = &v
	case *Timing:
		m.TimingTiming = v
	case Period:
		m.TimingPeriod = &v
	case *Period:
		m.TimingPeriod = v
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
	m.RecordedOn = r.RecordedOn
	if r.RecordedOn != nil && (r.RecordedOn.Id != nil || r.RecordedOn.Extension != nil) {
		m.RecordedOnPrimitiveElement = &primitiveElement{Id: r.RecordedOn.Id, Extension: r.RecordedOn.Extension}
	}
	m.Source = r.Source
	m.Device = r.Device
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.BodySite = r.BodySite
	m.Note = r.Note
	return m
}
func (r *DeviceUseStatement) UnmarshalJSON(b []byte) error {
	var m jsonDeviceUseStatement
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceUseStatement) unmarshalJSON(m jsonDeviceUseStatement) error {
	r.Id = m.Id
	if m.IdPrimitiveElement != nil {
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Text = m.Text
	r.Contained = make([]model.Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.Resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.BasedOn = m.BasedOn
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Subject = m.Subject
	r.DerivedFrom = m.DerivedFrom
	if m.TimingTiming != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingTiming
		r.Timing = v
	}
	if m.TimingPeriod != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingPeriod
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
	r.RecordedOn = m.RecordedOn
	if m.RecordedOnPrimitiveElement != nil {
		r.RecordedOn.Id = m.RecordedOnPrimitiveElement.Id
		r.RecordedOn.Extension = m.RecordedOnPrimitiveElement.Extension
	}
	r.Source = m.Source
	r.Device = m.Device
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.BodySite = m.BodySite
	r.Note = m.Note
	return nil
}
func (r DeviceUseStatement) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
