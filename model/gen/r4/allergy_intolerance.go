package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// Risk of harmful or undesirable, physiological response which is unique to an individual and associated with exposure to a substance.
//
// To record a clinical assessment of a propensity, or potential risk to an individual, of an adverse reaction upon future exposure to the specified substance, or class of substance.
type AllergyIntolerance struct {
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
	// Business identifiers assigned to this AllergyIntolerance by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// The clinical status of the allergy or intolerance.
	ClinicalStatus *CodeableConcept
	// Assertion about certainty associated with the propensity, or potential risk, of a reaction to the identified substance (including pharmaceutical product).
	VerificationStatus *CodeableConcept
	// Identification of the underlying physiological mechanism for the reaction risk.
	Type *Code
	// Category of the identified substance.
	Category []Code
	// Estimate of the potential clinical harm, or seriousness, of the reaction to the identified substance.
	Criticality *Code
	// Code for an allergy or intolerance statement (either a positive or a negated/excluded statement).  This may be a code for a substance or pharmaceutical product that is considered to be responsible for the adverse reaction risk (e.g., "Latex"), an allergy or intolerance condition (e.g., "Latex allergy"), or a negated/excluded code for a specific substance or class (e.g., "No latex allergy") or a general or categorical negated statement (e.g.,  "No known allergy", "No known drug allergies").  Note: the substance for a specific reaction may be different from the substance identified as the cause of the risk, but it must be consistent with it. For instance, it may be a more specific substance (e.g. a brand medication) or a composite product that includes the identified substance. It must be clinically safe to only process the 'code' and ignore the 'reaction.substance'.  If a receiving system is unable to confirm that AllergyIntolerance.reaction.substance falls within the semantic scope of AllergyIntolerance.code, then the receiving system should ignore AllergyIntolerance.reaction.substance.
	Code *CodeableConcept
	// The patient who has the allergy or intolerance.
	Patient Reference
	// The encounter when the allergy or intolerance was asserted.
	Encounter *Reference
	// Estimated or actual date,  date-time, or age when allergy or intolerance was identified.
	Onset isAllergyIntoleranceOnset
	// The recordedDate represents when this particular AllergyIntolerance record was created in the system, which is often a system-generated date.
	RecordedDate *DateTime
	// Individual who recorded the record and takes responsibility for its content.
	Recorder *Reference
	// The source of the information about the allergy that is recorded.
	Asserter *Reference
	// Represents the date and/or time of the last known occurrence of a reaction event.
	LastOccurrence *DateTime
	// Additional narrative about the propensity for the Adverse Reaction, not captured in other fields.
	Note []Annotation
	// Details about each adverse reaction event linked to exposure to the identified substance.
	Reaction []AllergyIntoleranceReaction
}

func (r AllergyIntolerance) ResourceType() string {
	return "AllergyIntolerance"
}
func (r AllergyIntolerance) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isAllergyIntoleranceOnset interface {
	isAllergyIntoleranceOnset()
}

func (r DateTime) isAllergyIntoleranceOnset() {}
func (r Age) isAllergyIntoleranceOnset()      {}
func (r Period) isAllergyIntoleranceOnset()   {}
func (r Range) isAllergyIntoleranceOnset()    {}
func (r String) isAllergyIntoleranceOnset()   {}

type jsonAllergyIntolerance struct {
	ResourceType                   string                       `json:"resourceType"`
	Id                             *Id                          `json:"id,omitempty"`
	IdPrimitiveElement             *primitiveElement            `json:"_id,omitempty"`
	Meta                           *Meta                        `json:"meta,omitempty"`
	ImplicitRules                  *Uri                         `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement  *primitiveElement            `json:"_implicitRules,omitempty"`
	Language                       *Code                        `json:"language,omitempty"`
	LanguagePrimitiveElement       *primitiveElement            `json:"_language,omitempty"`
	Text                           *Narrative                   `json:"text,omitempty"`
	Contained                      []ContainedResource          `json:"contained,omitempty"`
	Extension                      []Extension                  `json:"extension,omitempty"`
	ModifierExtension              []Extension                  `json:"modifierExtension,omitempty"`
	Identifier                     []Identifier                 `json:"identifier,omitempty"`
	ClinicalStatus                 *CodeableConcept             `json:"clinicalStatus,omitempty"`
	VerificationStatus             *CodeableConcept             `json:"verificationStatus,omitempty"`
	Type                           *Code                        `json:"type,omitempty"`
	TypePrimitiveElement           *primitiveElement            `json:"_type,omitempty"`
	Category                       []Code                       `json:"category,omitempty"`
	CategoryPrimitiveElement       []*primitiveElement          `json:"_category,omitempty"`
	Criticality                    *Code                        `json:"criticality,omitempty"`
	CriticalityPrimitiveElement    *primitiveElement            `json:"_criticality,omitempty"`
	Code                           *CodeableConcept             `json:"code,omitempty"`
	Patient                        Reference                    `json:"patient,omitempty"`
	Encounter                      *Reference                   `json:"encounter,omitempty"`
	OnsetDateTime                  *DateTime                    `json:"onsetDateTime,omitempty"`
	OnsetDateTimePrimitiveElement  *primitiveElement            `json:"_onsetDateTime,omitempty"`
	OnsetAge                       *Age                         `json:"onsetAge,omitempty"`
	OnsetPeriod                    *Period                      `json:"onsetPeriod,omitempty"`
	OnsetRange                     *Range                       `json:"onsetRange,omitempty"`
	OnsetString                    *String                      `json:"onsetString,omitempty"`
	OnsetStringPrimitiveElement    *primitiveElement            `json:"_onsetString,omitempty"`
	RecordedDate                   *DateTime                    `json:"recordedDate,omitempty"`
	RecordedDatePrimitiveElement   *primitiveElement            `json:"_recordedDate,omitempty"`
	Recorder                       *Reference                   `json:"recorder,omitempty"`
	Asserter                       *Reference                   `json:"asserter,omitempty"`
	LastOccurrence                 *DateTime                    `json:"lastOccurrence,omitempty"`
	LastOccurrencePrimitiveElement *primitiveElement            `json:"_lastOccurrence,omitempty"`
	Note                           []Annotation                 `json:"note,omitempty"`
	Reaction                       []AllergyIntoleranceReaction `json:"reaction,omitempty"`
}

func (r AllergyIntolerance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AllergyIntolerance) marshalJSON() jsonAllergyIntolerance {
	m := jsonAllergyIntolerance{}
	m.ResourceType = "AllergyIntolerance"
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
	if r.Type != nil && r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	anyCategoryValue := false
	for _, e := range r.Category {
		if e.Value != nil {
			anyCategoryValue = true
			break
		}
	}
	if anyCategoryValue {
		m.Category = r.Category
	}
	anyCategoryIdOrExtension := false
	for _, e := range r.Category {
		if e.Id != nil || e.Extension != nil {
			anyCategoryIdOrExtension = true
			break
		}
	}
	if anyCategoryIdOrExtension {
		m.CategoryPrimitiveElement = make([]*primitiveElement, 0, len(r.Category))
		for _, e := range r.Category {
			if e.Id != nil || e.Extension != nil {
				m.CategoryPrimitiveElement = append(m.CategoryPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.CategoryPrimitiveElement = append(m.CategoryPrimitiveElement, nil)
			}
		}
	}
	if r.Criticality != nil && r.Criticality.Value != nil {
		m.Criticality = r.Criticality
	}
	if r.Criticality != nil && (r.Criticality.Id != nil || r.Criticality.Extension != nil) {
		m.CriticalityPrimitiveElement = &primitiveElement{Id: r.Criticality.Id, Extension: r.Criticality.Extension}
	}
	m.Code = r.Code
	m.Patient = r.Patient
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
	if r.RecordedDate != nil && r.RecordedDate.Value != nil {
		m.RecordedDate = r.RecordedDate
	}
	if r.RecordedDate != nil && (r.RecordedDate.Id != nil || r.RecordedDate.Extension != nil) {
		m.RecordedDatePrimitiveElement = &primitiveElement{Id: r.RecordedDate.Id, Extension: r.RecordedDate.Extension}
	}
	m.Recorder = r.Recorder
	m.Asserter = r.Asserter
	if r.LastOccurrence != nil && r.LastOccurrence.Value != nil {
		m.LastOccurrence = r.LastOccurrence
	}
	if r.LastOccurrence != nil && (r.LastOccurrence.Id != nil || r.LastOccurrence.Extension != nil) {
		m.LastOccurrencePrimitiveElement = &primitiveElement{Id: r.LastOccurrence.Id, Extension: r.LastOccurrence.Extension}
	}
	m.Note = r.Note
	m.Reaction = r.Reaction
	return m
}
func (r *AllergyIntolerance) UnmarshalJSON(b []byte) error {
	var m jsonAllergyIntolerance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AllergyIntolerance) unmarshalJSON(m jsonAllergyIntolerance) error {
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
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		if r.Type == nil {
			r.Type = &Code{}
		}
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Category = m.Category
	for i, e := range m.CategoryPrimitiveElement {
		if len(r.Category) <= i {
			r.Category = append(r.Category, Code{})
		}
		if e != nil {
			r.Category[i].Id = e.Id
			r.Category[i].Extension = e.Extension
		}
	}
	r.Criticality = m.Criticality
	if m.CriticalityPrimitiveElement != nil {
		if r.Criticality == nil {
			r.Criticality = &Code{}
		}
		r.Criticality.Id = m.CriticalityPrimitiveElement.Id
		r.Criticality.Extension = m.CriticalityPrimitiveElement.Extension
	}
	r.Code = m.Code
	r.Patient = m.Patient
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
	r.LastOccurrence = m.LastOccurrence
	if m.LastOccurrencePrimitiveElement != nil {
		if r.LastOccurrence == nil {
			r.LastOccurrence = &DateTime{}
		}
		r.LastOccurrence.Id = m.LastOccurrencePrimitiveElement.Id
		r.LastOccurrence.Extension = m.LastOccurrencePrimitiveElement.Extension
	}
	r.Note = m.Note
	r.Reaction = m.Reaction
	return nil
}
func (r AllergyIntolerance) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "AllergyIntolerance"
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Category, xml.StartElement{Name: xml.Name{Local: "category"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Criticality, xml.StartElement{Name: xml.Name{Local: "criticality"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Patient, xml.StartElement{Name: xml.Name{Local: "patient"}})
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
	err = e.EncodeElement(r.LastOccurrence, xml.StartElement{Name: xml.Name{Local: "lastOccurrence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Reaction, xml.StartElement{Name: xml.Name{Local: "reaction"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *AllergyIntolerance) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "category":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = append(r.Category, v)
			case "criticality":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Criticality = &v
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = &v
			case "patient":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Patient = v
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
			case "lastOccurrence":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LastOccurrence = &v
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			case "reaction":
				var v AllergyIntoleranceReaction
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reaction = append(r.Reaction, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r AllergyIntolerance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Details about each adverse reaction event linked to exposure to the identified substance.
type AllergyIntoleranceReaction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identification of the specific substance (or pharmaceutical product) considered to be responsible for the Adverse Reaction event. Note: the substance for a specific reaction may be different from the substance identified as the cause of the risk, but it must be consistent with it. For instance, it may be a more specific substance (e.g. a brand medication) or a composite product that includes the identified substance. It must be clinically safe to only process the 'code' and ignore the 'reaction.substance'.  If a receiving system is unable to confirm that AllergyIntolerance.reaction.substance falls within the semantic scope of AllergyIntolerance.code, then the receiving system should ignore AllergyIntolerance.reaction.substance.
	Substance *CodeableConcept
	// Clinical symptoms and/or signs that are observed or associated with the adverse reaction event.
	Manifestation []CodeableConcept
	// Text description about the reaction as a whole, including details of the manifestation if required.
	Description *String
	// Record of the date and/or time of the onset of the Reaction.
	Onset *DateTime
	// Clinical assessment of the severity of the reaction event as a whole, potentially considering multiple different manifestations.
	Severity *Code
	// Identification of the route by which the subject was exposed to the substance.
	ExposureRoute *CodeableConcept
	// Additional text about the adverse reaction event not captured in other fields.
	Note []Annotation
}
type jsonAllergyIntoleranceReaction struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Substance                   *CodeableConcept  `json:"substance,omitempty"`
	Manifestation               []CodeableConcept `json:"manifestation,omitempty"`
	Description                 *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
	Onset                       *DateTime         `json:"onset,omitempty"`
	OnsetPrimitiveElement       *primitiveElement `json:"_onset,omitempty"`
	Severity                    *Code             `json:"severity,omitempty"`
	SeverityPrimitiveElement    *primitiveElement `json:"_severity,omitempty"`
	ExposureRoute               *CodeableConcept  `json:"exposureRoute,omitempty"`
	Note                        []Annotation      `json:"note,omitempty"`
}

func (r AllergyIntoleranceReaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AllergyIntoleranceReaction) marshalJSON() jsonAllergyIntoleranceReaction {
	m := jsonAllergyIntoleranceReaction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Substance = r.Substance
	m.Manifestation = r.Manifestation
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	if r.Onset != nil && r.Onset.Value != nil {
		m.Onset = r.Onset
	}
	if r.Onset != nil && (r.Onset.Id != nil || r.Onset.Extension != nil) {
		m.OnsetPrimitiveElement = &primitiveElement{Id: r.Onset.Id, Extension: r.Onset.Extension}
	}
	if r.Severity != nil && r.Severity.Value != nil {
		m.Severity = r.Severity
	}
	if r.Severity != nil && (r.Severity.Id != nil || r.Severity.Extension != nil) {
		m.SeverityPrimitiveElement = &primitiveElement{Id: r.Severity.Id, Extension: r.Severity.Extension}
	}
	m.ExposureRoute = r.ExposureRoute
	m.Note = r.Note
	return m
}
func (r *AllergyIntoleranceReaction) UnmarshalJSON(b []byte) error {
	var m jsonAllergyIntoleranceReaction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AllergyIntoleranceReaction) unmarshalJSON(m jsonAllergyIntoleranceReaction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Substance = m.Substance
	r.Manifestation = m.Manifestation
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Onset = m.Onset
	if m.OnsetPrimitiveElement != nil {
		if r.Onset == nil {
			r.Onset = &DateTime{}
		}
		r.Onset.Id = m.OnsetPrimitiveElement.Id
		r.Onset.Extension = m.OnsetPrimitiveElement.Extension
	}
	r.Severity = m.Severity
	if m.SeverityPrimitiveElement != nil {
		if r.Severity == nil {
			r.Severity = &Code{}
		}
		r.Severity.Id = m.SeverityPrimitiveElement.Id
		r.Severity.Extension = m.SeverityPrimitiveElement.Extension
	}
	r.ExposureRoute = m.ExposureRoute
	r.Note = m.Note
	return nil
}
func (r AllergyIntoleranceReaction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Substance, xml.StartElement{Name: xml.Name{Local: "substance"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Manifestation, xml.StartElement{Name: xml.Name{Local: "manifestation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Onset, xml.StartElement{Name: xml.Name{Local: "onset"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Severity, xml.StartElement{Name: xml.Name{Local: "severity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ExposureRoute, xml.StartElement{Name: xml.Name{Local: "exposureRoute"}})
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
func (r *AllergyIntoleranceReaction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "substance":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Substance = &v
			case "manifestation":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Manifestation = append(r.Manifestation, v)
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "onset":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Onset = &v
			case "severity":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Severity = &v
			case "exposureRoute":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ExposureRoute = &v
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
func (r AllergyIntoleranceReaction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
