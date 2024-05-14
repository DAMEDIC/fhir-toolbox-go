package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A record of a medication that is being consumed by a patient.   A MedicationStatement may indicate that the patient may be taking the medication now or has taken the medication in the past or will be taking the medication in the future.  The source of this information can be the patient, significant other (such as a family member or spouse), or a clinician.  A common scenario where this information is captured is during the history taking process during a patient visit or stay.   The medication information may come from sources such as the patient's memory, from a prescription bottle,  or from a list of medications the patient, clinician or other party maintains.
//
// The primary difference between a medication statement and a medication administration is that the medication administration has complete administration information and is based on actual administration information from the person who administered the medication.  A medication statement is often, if not always, less specific.  There is no required date/time when the medication was administered, in fact we only know that a source has reported the patient is taking this medication, where details such as time, quantity, or rate or even medication product may be incomplete or missing or less precise.  As stated earlier, the medication statement information may come from the patient's memory, from a prescription bottle or from a list of medications the patient, clinician or other party maintains.  Medication administration is more formal and is not missing detailed information.
type MedicationStatement struct {
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
	// Identifiers associated with this Medication Statement that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate. They are business identifiers assigned to this resource by the performer or other systems and remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// A plan, proposal or order that is fulfilled in whole or in part by this event.
	BasedOn []Reference
	// A larger event of which this particular event is a component or step.
	PartOf []Reference
	// A code representing the patient or other source's judgment about the state of the medication used that this statement is about.  Generally, this will be active or completed.
	Status Code
	// Captures the reason for the current state of the MedicationStatement.
	StatusReason []CodeableConcept
	// Indicates where the medication is expected to be consumed or administered.
	Category *CodeableConcept
	// Identifies the medication being administered. This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
	Medication isMedicationStatementMedication
	// The person, animal or group who is/was taking the medication.
	Subject Reference
	// The encounter or episode of care that establishes the context for this MedicationStatement.
	Context *Reference
	// The interval of time during which it is being asserted that the patient is/was/will be taking the medication (or was not taking, when the MedicationStatement.taken element is No).
	Effective isMedicationStatementEffective
	// The date when the medication statement was asserted by the information source.
	DateAsserted *DateTime
	// The person or organization that provided the information about the taking of this medication. Note: Use derivedFrom when a MedicationStatement is derived from other resources, e.g. Claim or MedicationRequest.
	InformationSource *Reference
	// Allows linking the MedicationStatement to the underlying MedicationRequest, or to other information that supports or is used to derive the MedicationStatement.
	DerivedFrom []Reference
	// A reason for why the medication is being/was taken.
	ReasonCode []CodeableConcept
	// Condition or observation that supports why the medication is being/was taken.
	ReasonReference []Reference
	// Provides extra information about the medication statement that is not conveyed by the other attributes.
	Note []Annotation
	// Indicates how the medication is/was or should be taken by the patient.
	Dosage []Dosage
}

func (r MedicationStatement) ResourceType() string {
	return "MedicationStatement"
}

type isMedicationStatementMedication interface {
	isMedicationStatementMedication()
}

func (r CodeableConcept) isMedicationStatementMedication() {}
func (r Reference) isMedicationStatementMedication()       {}

type isMedicationStatementEffective interface {
	isMedicationStatementEffective()
}

func (r DateTime) isMedicationStatementEffective() {}
func (r Period) isMedicationStatementEffective()   {}

type jsonMedicationStatement struct {
	ResourceType                      string              `json:"resourceType"`
	Id                                *Id                 `json:"id,omitempty"`
	IdPrimitiveElement                *primitiveElement   `json:"_id,omitempty"`
	Meta                              *Meta               `json:"meta,omitempty"`
	ImplicitRules                     *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement     *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                          *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement          *primitiveElement   `json:"_language,omitempty"`
	Text                              *Narrative          `json:"text,omitempty"`
	Contained                         []containedResource `json:"contained,omitempty"`
	Extension                         []Extension         `json:"extension,omitempty"`
	ModifierExtension                 []Extension         `json:"modifierExtension,omitempty"`
	Identifier                        []Identifier        `json:"identifier,omitempty"`
	BasedOn                           []Reference         `json:"basedOn,omitempty"`
	PartOf                            []Reference         `json:"partOf,omitempty"`
	Status                            Code                `json:"status,omitempty"`
	StatusPrimitiveElement            *primitiveElement   `json:"_status,omitempty"`
	StatusReason                      []CodeableConcept   `json:"statusReason,omitempty"`
	Category                          *CodeableConcept    `json:"category,omitempty"`
	MedicationCodeableConcept         *CodeableConcept    `json:"medicationCodeableConcept,omitempty"`
	MedicationReference               *Reference          `json:"medicationReference,omitempty"`
	Subject                           Reference           `json:"subject,omitempty"`
	Context                           *Reference          `json:"context,omitempty"`
	EffectiveDateTime                 *DateTime           `json:"effectiveDateTime,omitempty"`
	EffectiveDateTimePrimitiveElement *primitiveElement   `json:"_effectiveDateTime,omitempty"`
	EffectivePeriod                   *Period             `json:"effectivePeriod,omitempty"`
	DateAsserted                      *DateTime           `json:"dateAsserted,omitempty"`
	DateAssertedPrimitiveElement      *primitiveElement   `json:"_dateAsserted,omitempty"`
	InformationSource                 *Reference          `json:"informationSource,omitempty"`
	DerivedFrom                       []Reference         `json:"derivedFrom,omitempty"`
	ReasonCode                        []CodeableConcept   `json:"reasonCode,omitempty"`
	ReasonReference                   []Reference         `json:"reasonReference,omitempty"`
	Note                              []Annotation        `json:"note,omitempty"`
	Dosage                            []Dosage            `json:"dosage,omitempty"`
}

func (r MedicationStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationStatement) marshalJSON() jsonMedicationStatement {
	m := jsonMedicationStatement{}
	m.ResourceType = "MedicationStatement"
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
	m.Contained = make([]containedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, containedResource{resource: c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.BasedOn = r.BasedOn
	m.PartOf = r.PartOf
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.StatusReason = r.StatusReason
	m.Category = r.Category
	switch v := r.Medication.(type) {
	case CodeableConcept:
		m.MedicationCodeableConcept = &v
	case *CodeableConcept:
		m.MedicationCodeableConcept = v
	case Reference:
		m.MedicationReference = &v
	case *Reference:
		m.MedicationReference = v
	}
	m.Subject = r.Subject
	m.Context = r.Context
	switch v := r.Effective.(type) {
	case DateTime:
		m.EffectiveDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.EffectiveDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.EffectiveDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.EffectiveDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.EffectivePeriod = &v
	case *Period:
		m.EffectivePeriod = v
	}
	m.DateAsserted = r.DateAsserted
	if r.DateAsserted != nil && (r.DateAsserted.Id != nil || r.DateAsserted.Extension != nil) {
		m.DateAssertedPrimitiveElement = &primitiveElement{Id: r.DateAsserted.Id, Extension: r.DateAsserted.Extension}
	}
	m.InformationSource = r.InformationSource
	m.DerivedFrom = r.DerivedFrom
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Note = r.Note
	m.Dosage = r.Dosage
	return m
}
func (r *MedicationStatement) UnmarshalJSON(b []byte) error {
	var m jsonMedicationStatement
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationStatement) unmarshalJSON(m jsonMedicationStatement) error {
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
		r.Contained = append(r.Contained, v.resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.BasedOn = m.BasedOn
	r.PartOf = m.PartOf
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.StatusReason = m.StatusReason
	r.Category = m.Category
	if m.MedicationCodeableConcept != nil {
		if r.Medication != nil {
			return fmt.Errorf("multiple values for field \"Medication\"")
		}
		v := m.MedicationCodeableConcept
		r.Medication = v
	}
	if m.MedicationReference != nil {
		if r.Medication != nil {
			return fmt.Errorf("multiple values for field \"Medication\"")
		}
		v := m.MedicationReference
		r.Medication = v
	}
	r.Subject = m.Subject
	r.Context = m.Context
	if m.EffectiveDateTime != nil || m.EffectiveDateTimePrimitiveElement != nil {
		if r.Effective != nil {
			return fmt.Errorf("multiple values for field \"Effective\"")
		}
		v := m.EffectiveDateTime
		if m.EffectiveDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.EffectiveDateTimePrimitiveElement.Id
			v.Extension = m.EffectiveDateTimePrimitiveElement.Extension
		}
		r.Effective = v
	}
	if m.EffectivePeriod != nil {
		if r.Effective != nil {
			return fmt.Errorf("multiple values for field \"Effective\"")
		}
		v := m.EffectivePeriod
		r.Effective = v
	}
	r.DateAsserted = m.DateAsserted
	if m.DateAssertedPrimitiveElement != nil {
		r.DateAsserted.Id = m.DateAssertedPrimitiveElement.Id
		r.DateAsserted.Extension = m.DateAssertedPrimitiveElement.Extension
	}
	r.InformationSource = m.InformationSource
	r.DerivedFrom = m.DerivedFrom
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Note = m.Note
	r.Dosage = m.Dosage
	return nil
}
func (r MedicationStatement) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
