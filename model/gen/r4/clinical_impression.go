package r4

import (
	"encoding/json"
	"fmt"
)

// A record of a clinical assessment performed to determine what problem(s) may affect the patient and before planning the treatments or management strategies that are best to manage a patient's condition. Assessments are often 1:1 with a clinical consultation / encounter,  but this varies greatly depending on the clinical workflow. This resource is called "ClinicalImpression" rather than "ClinicalAssessment" to avoid confusion with the recording of assessment tools such as Apgar score.
type ClinicalImpression struct {
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
	Contained []Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Business identifiers assigned to this clinical impression by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// Identifies the workflow status of the assessment.
	Status Code
	// Captures the reason for the current state of the ClinicalImpression.
	StatusReason *CodeableConcept
	// Categorizes the type of clinical assessment performed.
	Code *CodeableConcept
	// A summary of the context and/or cause of the assessment - why / where it was performed, and what patient events/status prompted it.
	Description *String
	// The patient or group of individuals assessed as part of this record.
	Subject Reference
	// The Encounter during which this ClinicalImpression was created or to which the creation of this record is tightly associated.
	Encounter *Reference
	// The point in time or period over which the subject was assessed.
	Effective isClinicalImpressionEffective
	// Indicates when the documentation of the assessment was complete.
	Date *DateTime
	// The clinician performing the assessment.
	Assessor *Reference
	// A reference to the last assessment that was conducted on this patient. Assessments are often/usually ongoing in nature; a care provider (practitioner or team) will make new assessments on an ongoing basis as new data arises or the patient's conditions changes.
	Previous *Reference
	// A list of the relevant problems/conditions for a patient.
	Problem []Reference
	// One or more sets of investigations (signs, symptoms, etc.). The actual grouping of investigations varies greatly depending on the type and context of the assessment. These investigations may include data generated during the assessment process, or data previously generated and recorded that is pertinent to the outcomes.
	Investigation []ClinicalImpressionInvestigation
	// Reference to a specific published clinical protocol that was followed during this assessment, and/or that provides evidence in support of the diagnosis.
	Protocol []Uri
	// A text summary of the investigations and the diagnosis.
	Summary *String
	// Specific findings or diagnoses that were considered likely or relevant to ongoing treatment.
	Finding []ClinicalImpressionFinding
	// Estimate of likely outcome.
	PrognosisCodeableConcept []CodeableConcept
	// RiskAssessment expressing likely outcome.
	PrognosisReference []Reference
	// Information supporting the clinical impression.
	SupportingInfo []Reference
	// Commentary about the impression, typically recorded after the impression itself was made, though supplemental notes by the original author could also appear.
	Note []Annotation
}
type isClinicalImpressionEffective interface {
	isClinicalImpressionEffective()
}

func (r DateTime) isClinicalImpressionEffective() {}
func (r Period) isClinicalImpressionEffective()   {}

type jsonClinicalImpression struct {
	ResourceType                      string                            `json:"resourceType"`
	Id                                *Id                               `json:"id,omitempty"`
	IdPrimitiveElement                *primitiveElement                 `json:"_id,omitempty"`
	Meta                              *Meta                             `json:"meta,omitempty"`
	ImplicitRules                     *Uri                              `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement     *primitiveElement                 `json:"_implicitRules,omitempty"`
	Language                          *Code                             `json:"language,omitempty"`
	LanguagePrimitiveElement          *primitiveElement                 `json:"_language,omitempty"`
	Text                              *Narrative                        `json:"text,omitempty"`
	Contained                         []containedResource               `json:"contained,omitempty"`
	Extension                         []Extension                       `json:"extension,omitempty"`
	ModifierExtension                 []Extension                       `json:"modifierExtension,omitempty"`
	Identifier                        []Identifier                      `json:"identifier,omitempty"`
	Status                            Code                              `json:"status,omitempty"`
	StatusPrimitiveElement            *primitiveElement                 `json:"_status,omitempty"`
	StatusReason                      *CodeableConcept                  `json:"statusReason,omitempty"`
	Code                              *CodeableConcept                  `json:"code,omitempty"`
	Description                       *String                           `json:"description,omitempty"`
	DescriptionPrimitiveElement       *primitiveElement                 `json:"_description,omitempty"`
	Subject                           Reference                         `json:"subject,omitempty"`
	Encounter                         *Reference                        `json:"encounter,omitempty"`
	EffectiveDateTime                 *DateTime                         `json:"effectiveDateTime,omitempty"`
	EffectiveDateTimePrimitiveElement *primitiveElement                 `json:"_effectiveDateTime,omitempty"`
	EffectivePeriod                   *Period                           `json:"effectivePeriod,omitempty"`
	Date                              *DateTime                         `json:"date,omitempty"`
	DatePrimitiveElement              *primitiveElement                 `json:"_date,omitempty"`
	Assessor                          *Reference                        `json:"assessor,omitempty"`
	Previous                          *Reference                        `json:"previous,omitempty"`
	Problem                           []Reference                       `json:"problem,omitempty"`
	Investigation                     []ClinicalImpressionInvestigation `json:"investigation,omitempty"`
	Protocol                          []Uri                             `json:"protocol,omitempty"`
	ProtocolPrimitiveElement          []*primitiveElement               `json:"_protocol,omitempty"`
	Summary                           *String                           `json:"summary,omitempty"`
	SummaryPrimitiveElement           *primitiveElement                 `json:"_summary,omitempty"`
	Finding                           []ClinicalImpressionFinding       `json:"finding,omitempty"`
	PrognosisCodeableConcept          []CodeableConcept                 `json:"prognosisCodeableConcept,omitempty"`
	PrognosisReference                []Reference                       `json:"prognosisReference,omitempty"`
	SupportingInfo                    []Reference                       `json:"supportingInfo,omitempty"`
	Note                              []Annotation                      `json:"note,omitempty"`
}

func (r ClinicalImpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClinicalImpression) marshalJSON() jsonClinicalImpression {
	m := jsonClinicalImpression{}
	m.ResourceType = "ClinicalImpression"
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
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.StatusReason = r.StatusReason
	m.Code = r.Code
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Subject = r.Subject
	m.Encounter = r.Encounter
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
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Assessor = r.Assessor
	m.Previous = r.Previous
	m.Problem = r.Problem
	m.Investigation = r.Investigation
	m.Protocol = r.Protocol
	anyProtocolIdOrExtension := false
	for _, e := range r.Protocol {
		if e.Id != nil || e.Extension != nil {
			anyProtocolIdOrExtension = true
			break
		}
	}
	if anyProtocolIdOrExtension {
		m.ProtocolPrimitiveElement = make([]*primitiveElement, 0, len(r.Protocol))
		for _, e := range r.Protocol {
			if e.Id != nil || e.Extension != nil {
				m.ProtocolPrimitiveElement = append(m.ProtocolPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ProtocolPrimitiveElement = append(m.ProtocolPrimitiveElement, nil)
			}
		}
	}
	m.Summary = r.Summary
	if r.Summary != nil && (r.Summary.Id != nil || r.Summary.Extension != nil) {
		m.SummaryPrimitiveElement = &primitiveElement{Id: r.Summary.Id, Extension: r.Summary.Extension}
	}
	m.Finding = r.Finding
	m.PrognosisCodeableConcept = r.PrognosisCodeableConcept
	m.PrognosisReference = r.PrognosisReference
	m.SupportingInfo = r.SupportingInfo
	m.Note = r.Note
	return m
}
func (r *ClinicalImpression) UnmarshalJSON(b []byte) error {
	var m jsonClinicalImpression
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClinicalImpression) unmarshalJSON(m jsonClinicalImpression) error {
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
	r.Contained = make([]Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.StatusReason = m.StatusReason
	r.Code = m.Code
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Subject = m.Subject
	r.Encounter = m.Encounter
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
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Assessor = m.Assessor
	r.Previous = m.Previous
	r.Problem = m.Problem
	r.Investigation = m.Investigation
	r.Protocol = m.Protocol
	for i, e := range m.ProtocolPrimitiveElement {
		if len(r.Protocol) > i {
			r.Protocol[i].Id = e.Id
			r.Protocol[i].Extension = e.Extension
		} else {
			r.Protocol = append(r.Protocol, Uri{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Summary = m.Summary
	if m.SummaryPrimitiveElement != nil {
		r.Summary.Id = m.SummaryPrimitiveElement.Id
		r.Summary.Extension = m.SummaryPrimitiveElement.Extension
	}
	r.Finding = m.Finding
	r.PrognosisCodeableConcept = m.PrognosisCodeableConcept
	r.PrognosisReference = m.PrognosisReference
	r.SupportingInfo = m.SupportingInfo
	r.Note = m.Note
	return nil
}
func (r ClinicalImpression) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// One or more sets of investigations (signs, symptoms, etc.). The actual grouping of investigations varies greatly depending on the type and context of the assessment. These investigations may include data generated during the assessment process, or data previously generated and recorded that is pertinent to the outcomes.
type ClinicalImpressionInvestigation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A name/code for the group ("set") of investigations. Typically, this will be something like "signs", "symptoms", "clinical", "diagnostic", but the list is not constrained, and others such groups such as (exposure|family|travel|nutritional) history may be used.
	Code CodeableConcept
	// A record of a specific investigation that was undertaken.
	Item []Reference
}
type jsonClinicalImpressionInvestigation struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code,omitempty"`
	Item              []Reference     `json:"item,omitempty"`
}

func (r ClinicalImpressionInvestigation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClinicalImpressionInvestigation) marshalJSON() jsonClinicalImpressionInvestigation {
	m := jsonClinicalImpressionInvestigation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Item = r.Item
	return m
}
func (r *ClinicalImpressionInvestigation) UnmarshalJSON(b []byte) error {
	var m jsonClinicalImpressionInvestigation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClinicalImpressionInvestigation) unmarshalJSON(m jsonClinicalImpressionInvestigation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Item = m.Item
	return nil
}
func (r ClinicalImpressionInvestigation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Specific findings or diagnoses that were considered likely or relevant to ongoing treatment.
type ClinicalImpressionFinding struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Specific text or code for finding or diagnosis, which may include ruled-out or resolved conditions.
	ItemCodeableConcept *CodeableConcept
	// Specific reference for finding or diagnosis, which may include ruled-out or resolved conditions.
	ItemReference *Reference
	// Which investigations support finding or diagnosis.
	Basis *String
}
type jsonClinicalImpressionFinding struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	ItemCodeableConcept   *CodeableConcept  `json:"itemCodeableConcept,omitempty"`
	ItemReference         *Reference        `json:"itemReference,omitempty"`
	Basis                 *String           `json:"basis,omitempty"`
	BasisPrimitiveElement *primitiveElement `json:"_basis,omitempty"`
}

func (r ClinicalImpressionFinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClinicalImpressionFinding) marshalJSON() jsonClinicalImpressionFinding {
	m := jsonClinicalImpressionFinding{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.ItemCodeableConcept = r.ItemCodeableConcept
	m.ItemReference = r.ItemReference
	m.Basis = r.Basis
	if r.Basis != nil && (r.Basis.Id != nil || r.Basis.Extension != nil) {
		m.BasisPrimitiveElement = &primitiveElement{Id: r.Basis.Id, Extension: r.Basis.Extension}
	}
	return m
}
func (r *ClinicalImpressionFinding) UnmarshalJSON(b []byte) error {
	var m jsonClinicalImpressionFinding
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClinicalImpressionFinding) unmarshalJSON(m jsonClinicalImpressionFinding) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ItemCodeableConcept = m.ItemCodeableConcept
	r.ItemReference = m.ItemReference
	r.Basis = m.Basis
	if m.BasisPrimitiveElement != nil {
		r.Basis.Id = m.BasisPrimitiveElement.Id
		r.Basis.Extension = m.BasisPrimitiveElement.Extension
	}
	return nil
}
func (r ClinicalImpressionFinding) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
