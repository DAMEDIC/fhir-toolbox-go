package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions for a patient; e.g. Drug-drug interaction, Ineffective treatment frequency, Procedure-condition conflict, etc.
type DetectedIssue struct {
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
	// Business identifier associated with the detected issue record.
	Identifier []Identifier
	// Indicates the status of the detected issue.
	Status Code
	// Identifies the general type of issue identified.
	Code *CodeableConcept
	// Indicates the degree of importance associated with the identified issue based on the potential impact on the patient.
	Severity *Code
	// Indicates the patient whose record the detected issue is associated with.
	Patient *Reference
	// The date or period when the detected issue was initially identified.
	Identified isDetectedIssueIdentified
	// Individual or device responsible for the issue being raised.  For example, a decision support application or a pharmacist conducting a medication review.
	Author *Reference
	// Indicates the resource representing the current activity or proposed activity that is potentially problematic.
	Implicated []Reference
	// Supporting evidence or manifestations that provide the basis for identifying the detected issue such as a GuidanceResponse or MeasureReport.
	Evidence []DetectedIssueEvidence
	// A textual explanation of the detected issue.
	Detail *String
	// The literature, knowledge-base or similar reference that describes the propensity for the detected issue identified.
	Reference *Uri
	// Indicates an action that has been taken or is committed to reduce or eliminate the likelihood of the risk identified by the detected issue from manifesting.  Can also reflect an observation of known mitigating factors that may reduce/eliminate the need for any action.
	Mitigation []DetectedIssueMitigation
}

func (r DetectedIssue) ResourceType() string {
	return "DetectedIssue"
}
func (r DetectedIssue) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isDetectedIssueIdentified interface {
	isDetectedIssueIdentified()
}

func (r DateTime) isDetectedIssueIdentified() {}
func (r Period) isDetectedIssueIdentified()   {}

type jsonDetectedIssue struct {
	ResourceType                       string                    `json:"resourceType"`
	Id                                 *Id                       `json:"id,omitempty"`
	IdPrimitiveElement                 *primitiveElement         `json:"_id,omitempty"`
	Meta                               *Meta                     `json:"meta,omitempty"`
	ImplicitRules                      *Uri                      `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement      *primitiveElement         `json:"_implicitRules,omitempty"`
	Language                           *Code                     `json:"language,omitempty"`
	LanguagePrimitiveElement           *primitiveElement         `json:"_language,omitempty"`
	Text                               *Narrative                `json:"text,omitempty"`
	Contained                          []containedResource       `json:"contained,omitempty"`
	Extension                          []Extension               `json:"extension,omitempty"`
	ModifierExtension                  []Extension               `json:"modifierExtension,omitempty"`
	Identifier                         []Identifier              `json:"identifier,omitempty"`
	Status                             Code                      `json:"status,omitempty"`
	StatusPrimitiveElement             *primitiveElement         `json:"_status,omitempty"`
	Code                               *CodeableConcept          `json:"code,omitempty"`
	Severity                           *Code                     `json:"severity,omitempty"`
	SeverityPrimitiveElement           *primitiveElement         `json:"_severity,omitempty"`
	Patient                            *Reference                `json:"patient,omitempty"`
	IdentifiedDateTime                 *DateTime                 `json:"identifiedDateTime,omitempty"`
	IdentifiedDateTimePrimitiveElement *primitiveElement         `json:"_identifiedDateTime,omitempty"`
	IdentifiedPeriod                   *Period                   `json:"identifiedPeriod,omitempty"`
	Author                             *Reference                `json:"author,omitempty"`
	Implicated                         []Reference               `json:"implicated,omitempty"`
	Evidence                           []DetectedIssueEvidence   `json:"evidence,omitempty"`
	Detail                             *String                   `json:"detail,omitempty"`
	DetailPrimitiveElement             *primitiveElement         `json:"_detail,omitempty"`
	Reference                          *Uri                      `json:"reference,omitempty"`
	ReferencePrimitiveElement          *primitiveElement         `json:"_reference,omitempty"`
	Mitigation                         []DetectedIssueMitigation `json:"mitigation,omitempty"`
}

func (r DetectedIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DetectedIssue) marshalJSON() jsonDetectedIssue {
	m := jsonDetectedIssue{}
	m.ResourceType = "DetectedIssue"
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
	m.Code = r.Code
	m.Severity = r.Severity
	if r.Severity != nil && (r.Severity.Id != nil || r.Severity.Extension != nil) {
		m.SeverityPrimitiveElement = &primitiveElement{Id: r.Severity.Id, Extension: r.Severity.Extension}
	}
	m.Patient = r.Patient
	switch v := r.Identified.(type) {
	case DateTime:
		m.IdentifiedDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.IdentifiedDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.IdentifiedDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.IdentifiedDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.IdentifiedPeriod = &v
	case *Period:
		m.IdentifiedPeriod = v
	}
	m.Author = r.Author
	m.Implicated = r.Implicated
	m.Evidence = r.Evidence
	m.Detail = r.Detail
	if r.Detail != nil && (r.Detail.Id != nil || r.Detail.Extension != nil) {
		m.DetailPrimitiveElement = &primitiveElement{Id: r.Detail.Id, Extension: r.Detail.Extension}
	}
	m.Reference = r.Reference
	if r.Reference != nil && (r.Reference.Id != nil || r.Reference.Extension != nil) {
		m.ReferencePrimitiveElement = &primitiveElement{Id: r.Reference.Id, Extension: r.Reference.Extension}
	}
	m.Mitigation = r.Mitigation
	return m
}
func (r *DetectedIssue) UnmarshalJSON(b []byte) error {
	var m jsonDetectedIssue
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DetectedIssue) unmarshalJSON(m jsonDetectedIssue) error {
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
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Code = m.Code
	r.Severity = m.Severity
	if m.SeverityPrimitiveElement != nil {
		r.Severity.Id = m.SeverityPrimitiveElement.Id
		r.Severity.Extension = m.SeverityPrimitiveElement.Extension
	}
	r.Patient = m.Patient
	if m.IdentifiedDateTime != nil || m.IdentifiedDateTimePrimitiveElement != nil {
		if r.Identified != nil {
			return fmt.Errorf("multiple values for field \"Identified\"")
		}
		v := m.IdentifiedDateTime
		if m.IdentifiedDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.IdentifiedDateTimePrimitiveElement.Id
			v.Extension = m.IdentifiedDateTimePrimitiveElement.Extension
		}
		r.Identified = v
	}
	if m.IdentifiedPeriod != nil {
		if r.Identified != nil {
			return fmt.Errorf("multiple values for field \"Identified\"")
		}
		v := m.IdentifiedPeriod
		r.Identified = v
	}
	r.Author = m.Author
	r.Implicated = m.Implicated
	r.Evidence = m.Evidence
	r.Detail = m.Detail
	if m.DetailPrimitiveElement != nil {
		r.Detail.Id = m.DetailPrimitiveElement.Id
		r.Detail.Extension = m.DetailPrimitiveElement.Extension
	}
	r.Reference = m.Reference
	if m.ReferencePrimitiveElement != nil {
		r.Reference.Id = m.ReferencePrimitiveElement.Id
		r.Reference.Extension = m.ReferencePrimitiveElement.Extension
	}
	r.Mitigation = m.Mitigation
	return nil
}
func (r DetectedIssue) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Supporting evidence or manifestations that provide the basis for identifying the detected issue such as a GuidanceResponse or MeasureReport.
type DetectedIssueEvidence struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A manifestation that led to the recording of this detected issue.
	Code []CodeableConcept
	// Links to resources that constitute evidence for the detected issue such as a GuidanceResponse or MeasureReport.
	Detail []Reference
}
type jsonDetectedIssueEvidence struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Detail            []Reference       `json:"detail,omitempty"`
}

func (r DetectedIssueEvidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DetectedIssueEvidence) marshalJSON() jsonDetectedIssueEvidence {
	m := jsonDetectedIssueEvidence{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Detail = r.Detail
	return m
}
func (r *DetectedIssueEvidence) UnmarshalJSON(b []byte) error {
	var m jsonDetectedIssueEvidence
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DetectedIssueEvidence) unmarshalJSON(m jsonDetectedIssueEvidence) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Detail = m.Detail
	return nil
}
func (r DetectedIssueEvidence) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates an action that has been taken or is committed to reduce or eliminate the likelihood of the risk identified by the detected issue from manifesting.  Can also reflect an observation of known mitigating factors that may reduce/eliminate the need for any action.
type DetectedIssueMitigation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Describes the action that was taken or the observation that was made that reduces/eliminates the risk associated with the identified issue.
	Action CodeableConcept
	// Indicates when the mitigating action was documented.
	Date *DateTime
	// Identifies the practitioner who determined the mitigation and takes responsibility for the mitigation step occurring.
	Author *Reference
}
type jsonDetectedIssueMitigation struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Action               CodeableConcept   `json:"action,omitempty"`
	Date                 *DateTime         `json:"date,omitempty"`
	DatePrimitiveElement *primitiveElement `json:"_date,omitempty"`
	Author               *Reference        `json:"author,omitempty"`
}

func (r DetectedIssueMitigation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DetectedIssueMitigation) marshalJSON() jsonDetectedIssueMitigation {
	m := jsonDetectedIssueMitigation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Action = r.Action
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Author = r.Author
	return m
}
func (r *DetectedIssueMitigation) UnmarshalJSON(b []byte) error {
	var m jsonDetectedIssueMitigation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DetectedIssueMitigation) unmarshalJSON(m jsonDetectedIssueMitigation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Action = m.Action
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Author = m.Author
	return nil
}
func (r DetectedIssueMitigation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
