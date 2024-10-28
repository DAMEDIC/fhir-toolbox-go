package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// A clinical condition, problem, diagnosis, or other event, situation, issue, or clinical concept that has risen to a level of concern.
type Condition struct {
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
	// Business identifiers assigned to this condition by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// The clinical status of the condition.
	ClinicalStatus *CodeableConcept
	// The verification status to support the clinical status of the condition.
	VerificationStatus *CodeableConcept
	// A category assigned to the condition.
	Category []CodeableConcept
	// A subjective assessment of the severity of the condition as evaluated by the clinician.
	Severity *CodeableConcept
	// Identification of the condition, problem or diagnosis.
	Code *CodeableConcept
	// The anatomical location where this condition manifests itself.
	BodySite []CodeableConcept
	// Indicates the patient or group who the condition record is associated with.
	Subject Reference
	// The Encounter during which this Condition was created or to which the creation of this record is tightly associated.
	Encounter *Reference
	// Estimated or actual date or date-time  the condition began, in the opinion of the clinician.
	Onset isConditionOnset
	// The date or estimated date that the condition resolved or went into remission. This is called "abatement" because of the many overloaded connotations associated with "remission" or "resolution" - Conditions are never really resolved, but they can abate.
	Abatement isConditionAbatement
	// The recordedDate represents when this particular Condition record was created in the system, which is often a system-generated date.
	RecordedDate *DateTime
	// Individual who recorded the record and takes responsibility for its content.
	Recorder *Reference
	// Individual who is making the condition statement.
	Asserter *Reference
	// Clinical stage or grade of a condition. May include formal severity assessments.
	Stage []ConditionStage
	// Supporting evidence / manifestations that are the basis of the Condition's verification status, such as evidence that confirmed or refuted the condition.
	Evidence []ConditionEvidence
	// Additional information about the Condition. This is a general notes/comments entry  for description of the Condition, its diagnosis and prognosis.
	Note []Annotation
}

func (r Condition) ResourceType() string {
	return "Condition"
}
func (r Condition) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isConditionOnset interface {
	isConditionOnset()
}

func (r DateTime) isConditionOnset() {}
func (r Age) isConditionOnset()      {}
func (r Period) isConditionOnset()   {}
func (r Range) isConditionOnset()    {}
func (r String) isConditionOnset()   {}

type isConditionAbatement interface {
	isConditionAbatement()
}

func (r DateTime) isConditionAbatement() {}
func (r Age) isConditionAbatement()      {}
func (r Period) isConditionAbatement()   {}
func (r Range) isConditionAbatement()    {}
func (r String) isConditionAbatement()   {}

type jsonCondition struct {
	ResourceType                      string              `json:"resourceType"`
	Id                                *Id                 `json:"id,omitempty"`
	IdPrimitiveElement                *primitiveElement   `json:"_id,omitempty"`
	Meta                              *Meta               `json:"meta,omitempty"`
	ImplicitRules                     *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement     *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                          *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement          *primitiveElement   `json:"_language,omitempty"`
	Text                              *Narrative          `json:"text,omitempty"`
	Contained                         []ContainedResource `json:"contained,omitempty"`
	Extension                         []Extension         `json:"extension,omitempty"`
	ModifierExtension                 []Extension         `json:"modifierExtension,omitempty"`
	Identifier                        []Identifier        `json:"identifier,omitempty"`
	ClinicalStatus                    *CodeableConcept    `json:"clinicalStatus,omitempty"`
	VerificationStatus                *CodeableConcept    `json:"verificationStatus,omitempty"`
	Category                          []CodeableConcept   `json:"category,omitempty"`
	Severity                          *CodeableConcept    `json:"severity,omitempty"`
	Code                              *CodeableConcept    `json:"code,omitempty"`
	BodySite                          []CodeableConcept   `json:"bodySite,omitempty"`
	Subject                           Reference           `json:"subject,omitempty"`
	Encounter                         *Reference          `json:"encounter,omitempty"`
	OnsetDateTime                     *DateTime           `json:"onsetDateTime,omitempty"`
	OnsetDateTimePrimitiveElement     *primitiveElement   `json:"_onsetDateTime,omitempty"`
	OnsetAge                          *Age                `json:"onsetAge,omitempty"`
	OnsetPeriod                       *Period             `json:"onsetPeriod,omitempty"`
	OnsetRange                        *Range              `json:"onsetRange,omitempty"`
	OnsetString                       *String             `json:"onsetString,omitempty"`
	OnsetStringPrimitiveElement       *primitiveElement   `json:"_onsetString,omitempty"`
	AbatementDateTime                 *DateTime           `json:"abatementDateTime,omitempty"`
	AbatementDateTimePrimitiveElement *primitiveElement   `json:"_abatementDateTime,omitempty"`
	AbatementAge                      *Age                `json:"abatementAge,omitempty"`
	AbatementPeriod                   *Period             `json:"abatementPeriod,omitempty"`
	AbatementRange                    *Range              `json:"abatementRange,omitempty"`
	AbatementString                   *String             `json:"abatementString,omitempty"`
	AbatementStringPrimitiveElement   *primitiveElement   `json:"_abatementString,omitempty"`
	RecordedDate                      *DateTime           `json:"recordedDate,omitempty"`
	RecordedDatePrimitiveElement      *primitiveElement   `json:"_recordedDate,omitempty"`
	Recorder                          *Reference          `json:"recorder,omitempty"`
	Asserter                          *Reference          `json:"asserter,omitempty"`
	Stage                             []ConditionStage    `json:"stage,omitempty"`
	Evidence                          []ConditionEvidence `json:"evidence,omitempty"`
	Note                              []Annotation        `json:"note,omitempty"`
}

func (r Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Condition) marshalJSON() jsonCondition {
	m := jsonCondition{}
	m.ResourceType = "Condition"
	if r.Id != nil && r.Id.Value != nil {
		m.Id = r.Id
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		m.ImplicitRules = r.ImplicitRules
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	if r.Language != nil && r.Language.Value != nil {
		m.Language = r.Language
	}
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
	m.ClinicalStatus = r.ClinicalStatus
	m.VerificationStatus = r.VerificationStatus
	m.Category = r.Category
	m.Severity = r.Severity
	m.Code = r.Code
	m.BodySite = r.BodySite
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	switch v := r.Onset.(type) {
	case DateTime:
		if v.Value != nil {
			m.OnsetDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.OnsetDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.OnsetDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.OnsetDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Age:
		m.OnsetAge = &v
	case *Age:
		m.OnsetAge = v
	case Period:
		m.OnsetPeriod = &v
	case *Period:
		m.OnsetPeriod = v
	case Range:
		m.OnsetRange = &v
	case *Range:
		m.OnsetRange = v
	case String:
		if v.Value != nil {
			m.OnsetString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.OnsetStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.OnsetString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.OnsetStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	switch v := r.Abatement.(type) {
	case DateTime:
		if v.Value != nil {
			m.AbatementDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.AbatementDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.AbatementDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.AbatementDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Age:
		m.AbatementAge = &v
	case *Age:
		m.AbatementAge = v
	case Period:
		m.AbatementPeriod = &v
	case *Period:
		m.AbatementPeriod = v
	case Range:
		m.AbatementRange = &v
	case *Range:
		m.AbatementRange = v
	case String:
		if v.Value != nil {
			m.AbatementString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.AbatementStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.AbatementString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.AbatementStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	if r.RecordedDate != nil && r.RecordedDate.Value != nil {
		m.RecordedDate = r.RecordedDate
	}
	if r.RecordedDate != nil && (r.RecordedDate.Id != nil || r.RecordedDate.Extension != nil) {
		m.RecordedDatePrimitiveElement = &primitiveElement{Id: r.RecordedDate.Id, Extension: r.RecordedDate.Extension}
	}
	m.Recorder = r.Recorder
	m.Asserter = r.Asserter
	m.Stage = r.Stage
	m.Evidence = r.Evidence
	m.Note = r.Note
	return m
}
func (r *Condition) UnmarshalJSON(b []byte) error {
	var m jsonCondition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Condition) unmarshalJSON(m jsonCondition) error {
	r.Id = m.Id
	if m.IdPrimitiveElement != nil {
		if r.Id == nil {
			r.Id = &Id{}
		}
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		if r.ImplicitRules == nil {
			r.ImplicitRules = &Uri{}
		}
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		if r.Language == nil {
			r.Language = &Code{}
		}
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
	r.ClinicalStatus = m.ClinicalStatus
	r.VerificationStatus = m.VerificationStatus
	r.Category = m.Category
	r.Severity = m.Severity
	r.Code = m.Code
	r.BodySite = m.BodySite
	r.Subject = m.Subject
	r.Encounter = m.Encounter
	if m.OnsetDateTime != nil || m.OnsetDateTimePrimitiveElement != nil {
		if r.Onset != nil {
			return fmt.Errorf("multiple values for field \"Onset\"")
		}
		v := m.OnsetDateTime
		if m.OnsetDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.OnsetDateTimePrimitiveElement.Id
			v.Extension = m.OnsetDateTimePrimitiveElement.Extension
		}
		r.Onset = v
	}
	if m.OnsetAge != nil {
		if r.Onset != nil {
			return fmt.Errorf("multiple values for field \"Onset\"")
		}
		v := m.OnsetAge
		r.Onset = v
	}
	if m.OnsetPeriod != nil {
		if r.Onset != nil {
			return fmt.Errorf("multiple values for field \"Onset\"")
		}
		v := m.OnsetPeriod
		r.Onset = v
	}
	if m.OnsetRange != nil {
		if r.Onset != nil {
			return fmt.Errorf("multiple values for field \"Onset\"")
		}
		v := m.OnsetRange
		r.Onset = v
	}
	if m.OnsetString != nil || m.OnsetStringPrimitiveElement != nil {
		if r.Onset != nil {
			return fmt.Errorf("multiple values for field \"Onset\"")
		}
		v := m.OnsetString
		if m.OnsetStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.OnsetStringPrimitiveElement.Id
			v.Extension = m.OnsetStringPrimitiveElement.Extension
		}
		r.Onset = v
	}
	if m.AbatementDateTime != nil || m.AbatementDateTimePrimitiveElement != nil {
		if r.Abatement != nil {
			return fmt.Errorf("multiple values for field \"Abatement\"")
		}
		v := m.AbatementDateTime
		if m.AbatementDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.AbatementDateTimePrimitiveElement.Id
			v.Extension = m.AbatementDateTimePrimitiveElement.Extension
		}
		r.Abatement = v
	}
	if m.AbatementAge != nil {
		if r.Abatement != nil {
			return fmt.Errorf("multiple values for field \"Abatement\"")
		}
		v := m.AbatementAge
		r.Abatement = v
	}
	if m.AbatementPeriod != nil {
		if r.Abatement != nil {
			return fmt.Errorf("multiple values for field \"Abatement\"")
		}
		v := m.AbatementPeriod
		r.Abatement = v
	}
	if m.AbatementRange != nil {
		if r.Abatement != nil {
			return fmt.Errorf("multiple values for field \"Abatement\"")
		}
		v := m.AbatementRange
		r.Abatement = v
	}
	if m.AbatementString != nil || m.AbatementStringPrimitiveElement != nil {
		if r.Abatement != nil {
			return fmt.Errorf("multiple values for field \"Abatement\"")
		}
		v := m.AbatementString
		if m.AbatementStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.AbatementStringPrimitiveElement.Id
			v.Extension = m.AbatementStringPrimitiveElement.Extension
		}
		r.Abatement = v
	}
	r.RecordedDate = m.RecordedDate
	if m.RecordedDatePrimitiveElement != nil {
		if r.RecordedDate == nil {
			r.RecordedDate = &DateTime{}
		}
		r.RecordedDate.Id = m.RecordedDatePrimitiveElement.Id
		r.RecordedDate.Extension = m.RecordedDatePrimitiveElement.Extension
	}
	r.Recorder = m.Recorder
	r.Asserter = m.Asserter
	r.Stage = m.Stage
	r.Evidence = m.Evidence
	r.Note = m.Note
	return nil
}
func (r Condition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Condition"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Id, xml.StartElement{Name: xml.Name{Local: "id"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Meta, xml.StartElement{Name: xml.Name{Local: "meta"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ImplicitRules, xml.StartElement{Name: xml.Name{Local: "implicitRules"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	v := make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		v = append(v, ContainedResource{c})
	}
	err = e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contained"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ClinicalStatus, xml.StartElement{Name: xml.Name{Local: "clinicalStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.VerificationStatus, xml.StartElement{Name: xml.Name{Local: "verificationStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Category, xml.StartElement{Name: xml.Name{Local: "category"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Severity, xml.StartElement{Name: xml.Name{Local: "severity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BodySite, xml.StartElement{Name: xml.Name{Local: "bodySite"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Subject, xml.StartElement{Name: xml.Name{Local: "subject"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Encounter, xml.StartElement{Name: xml.Name{Local: "encounter"}})
	if err != nil {
		return err
	}
	switch v := r.Onset.(type) {
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "onsetDateTime"}})
		if err != nil {
			return err
		}
	case Age, *Age:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "onsetAge"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "onsetPeriod"}})
		if err != nil {
			return err
		}
	case Range, *Range:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "onsetRange"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "onsetString"}})
		if err != nil {
			return err
		}
	}
	switch v := r.Abatement.(type) {
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "abatementDateTime"}})
		if err != nil {
			return err
		}
	case Age, *Age:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "abatementAge"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "abatementPeriod"}})
		if err != nil {
			return err
		}
	case Range, *Range:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "abatementRange"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "abatementString"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.RecordedDate, xml.StartElement{Name: xml.Name{Local: "recordedDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Recorder, xml.StartElement{Name: xml.Name{Local: "recorder"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Asserter, xml.StartElement{Name: xml.Name{Local: "asserter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Stage, xml.StartElement{Name: xml.Name{Local: "stage"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Evidence, xml.StartElement{Name: xml.Name{Local: "evidence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Condition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "id":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Id = &v
			case "meta":
				var v Meta
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Meta = &v
			case "implicitRules":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ImplicitRules = &v
			case "language":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = &v
			case "text":
				var v Narrative
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = &v
			case "contained":
				var c ContainedResource
				err := d.DecodeElement(&c, &t)
				if err != nil {
					return err
				}
				r.Contained = append(r.Contained, c.Resource)
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "clinicalStatus":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ClinicalStatus = &v
			case "verificationStatus":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.VerificationStatus = &v
			case "category":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = append(r.Category, v)
			case "severity":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Severity = &v
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = &v
			case "bodySite":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BodySite = append(r.BodySite, v)
			case "subject":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = v
			case "encounter":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Encounter = &v
			case "onsetDateTime":
				if r.Onset != nil {
					return fmt.Errorf("multiple values for field \"Onset\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Onset = v
			case "onsetAge":
				if r.Onset != nil {
					return fmt.Errorf("multiple values for field \"Onset\"")
				}
				var v Age
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Onset = v
			case "onsetPeriod":
				if r.Onset != nil {
					return fmt.Errorf("multiple values for field \"Onset\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Onset = v
			case "onsetRange":
				if r.Onset != nil {
					return fmt.Errorf("multiple values for field \"Onset\"")
				}
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Onset = v
			case "onsetString":
				if r.Onset != nil {
					return fmt.Errorf("multiple values for field \"Onset\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Onset = v
			case "abatementDateTime":
				if r.Abatement != nil {
					return fmt.Errorf("multiple values for field \"Abatement\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Abatement = v
			case "abatementAge":
				if r.Abatement != nil {
					return fmt.Errorf("multiple values for field \"Abatement\"")
				}
				var v Age
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Abatement = v
			case "abatementPeriod":
				if r.Abatement != nil {
					return fmt.Errorf("multiple values for field \"Abatement\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Abatement = v
			case "abatementRange":
				if r.Abatement != nil {
					return fmt.Errorf("multiple values for field \"Abatement\"")
				}
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Abatement = v
			case "abatementString":
				if r.Abatement != nil {
					return fmt.Errorf("multiple values for field \"Abatement\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Abatement = v
			case "recordedDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RecordedDate = &v
			case "recorder":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Recorder = &v
			case "asserter":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Asserter = &v
			case "stage":
				var v ConditionStage
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Stage = append(r.Stage, v)
			case "evidence":
				var v ConditionEvidence
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Evidence = append(r.Evidence, v)
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Condition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Clinical stage or grade of a condition. May include formal severity assessments.
type ConditionStage struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A simple summary of the stage such as "Stage 3". The determination of the stage is disease-specific.
	Summary *CodeableConcept
	// Reference to a formal record of the evidence on which the staging assessment is based.
	Assessment []Reference
	// The kind of staging, such as pathological or clinical staging.
	Type *CodeableConcept
}
type jsonConditionStage struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Summary           *CodeableConcept `json:"summary,omitempty"`
	Assessment        []Reference      `json:"assessment,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
}

func (r ConditionStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ConditionStage) marshalJSON() jsonConditionStage {
	m := jsonConditionStage{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Summary = r.Summary
	m.Assessment = r.Assessment
	m.Type = r.Type
	return m
}
func (r *ConditionStage) UnmarshalJSON(b []byte) error {
	var m jsonConditionStage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ConditionStage) unmarshalJSON(m jsonConditionStage) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Summary = m.Summary
	r.Assessment = m.Assessment
	r.Type = m.Type
	return nil
}
func (r ConditionStage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Summary, xml.StartElement{Name: xml.Name{Local: "summary"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Assessment, xml.StartElement{Name: xml.Name{Local: "assessment"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ConditionStage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "summary":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Summary = &v
			case "assessment":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Assessment = append(r.Assessment, v)
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ConditionStage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Supporting evidence / manifestations that are the basis of the Condition's verification status, such as evidence that confirmed or refuted the condition.
type ConditionEvidence struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A manifestation or symptom that led to the recording of this condition.
	Code []CodeableConcept
	// Links to other relevant information, including pathology reports.
	Detail []Reference
}
type jsonConditionEvidence struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Detail            []Reference       `json:"detail,omitempty"`
}

func (r ConditionEvidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ConditionEvidence) marshalJSON() jsonConditionEvidence {
	m := jsonConditionEvidence{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Detail = r.Detail
	return m
}
func (r *ConditionEvidence) UnmarshalJSON(b []byte) error {
	var m jsonConditionEvidence
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ConditionEvidence) unmarshalJSON(m jsonConditionEvidence) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Detail = m.Detail
	return nil
}
func (r ConditionEvidence) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Detail, xml.StartElement{Name: xml.Name{Local: "detail"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ConditionEvidence) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = append(r.Code, v)
			case "detail":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = append(r.Detail, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ConditionEvidence) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
