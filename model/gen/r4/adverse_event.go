package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// Actual or  potential/avoided event causing unintended physical injury resulting from or contributed to by medical care, a research study or other healthcare setting factors that requires additional monitoring, treatment, or hospitalization, or that results in death.
type AdverseEvent struct {
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
	// Business identifiers assigned to this adverse event by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier *Identifier
	// Whether the event actually happened, or just had the potential to. Note that this is independent of whether anyone was affected or harmed or how severely.
	Actuality Code
	// The overall type of event, intended for search and filtering purposes.
	Category []CodeableConcept
	// This element defines the specific type of event that occurred or that was prevented from occurring.
	Event *CodeableConcept
	// This subject or group impacted by the event.
	Subject Reference
	// The Encounter during which AdverseEvent was created or to which the creation of this record is tightly associated.
	Encounter *Reference
	// The date (and perhaps time) when the adverse event occurred.
	Date *DateTime
	// Estimated or actual date the AdverseEvent began, in the opinion of the reporter.
	Detected *DateTime
	// The date on which the existence of the AdverseEvent was first recorded.
	RecordedDate *DateTime
	// Includes information about the reaction that occurred as a result of exposure to a substance (for example, a drug or a chemical).
	ResultingCondition []Reference
	// The information about where the adverse event occurred.
	Location *Reference
	// Assessment whether this event was of real importance.
	Seriousness *CodeableConcept
	// Describes the severity of the adverse event, in relation to the subject. Contrast to AdverseEvent.seriousness - a severe rash might not be serious, but a mild heart problem is.
	Severity *CodeableConcept
	// Describes the type of outcome from the adverse event.
	Outcome *CodeableConcept
	// Information on who recorded the adverse event.  May be the patient or a practitioner.
	Recorder *Reference
	// Parties that may or should contribute or have contributed information to the adverse event, which can consist of one or more activities.  Such information includes information leading to the decision to perform the activity and how to perform the activity (e.g. consultant), information that the activity itself seeks to reveal (e.g. informant of clinical history), or information about what activity was performed (e.g. informant witness).
	Contributor []Reference
	// Describes the entity that is suspected to have caused the adverse event.
	SuspectEntity []AdverseEventSuspectEntity
	// AdverseEvent.subjectMedicalHistory.
	SubjectMedicalHistory []Reference
	// AdverseEvent.referenceDocument.
	ReferenceDocument []Reference
	// AdverseEvent.study.
	Study []Reference
}

func (r AdverseEvent) ResourceType() string {
	return "AdverseEvent"
}
func (r AdverseEvent) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonAdverseEvent struct {
	ResourceType                  string                      `json:"resourceType"`
	Id                            *Id                         `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement           `json:"_id,omitempty"`
	Meta                          *Meta                       `json:"meta,omitempty"`
	ImplicitRules                 *Uri                        `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement           `json:"_implicitRules,omitempty"`
	Language                      *Code                       `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement           `json:"_language,omitempty"`
	Text                          *Narrative                  `json:"text,omitempty"`
	Contained                     []ContainedResource         `json:"contained,omitempty"`
	Extension                     []Extension                 `json:"extension,omitempty"`
	ModifierExtension             []Extension                 `json:"modifierExtension,omitempty"`
	Identifier                    *Identifier                 `json:"identifier,omitempty"`
	Actuality                     Code                        `json:"actuality,omitempty"`
	ActualityPrimitiveElement     *primitiveElement           `json:"_actuality,omitempty"`
	Category                      []CodeableConcept           `json:"category,omitempty"`
	Event                         *CodeableConcept            `json:"event,omitempty"`
	Subject                       Reference                   `json:"subject,omitempty"`
	Encounter                     *Reference                  `json:"encounter,omitempty"`
	Date                          *DateTime                   `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement           `json:"_date,omitempty"`
	Detected                      *DateTime                   `json:"detected,omitempty"`
	DetectedPrimitiveElement      *primitiveElement           `json:"_detected,omitempty"`
	RecordedDate                  *DateTime                   `json:"recordedDate,omitempty"`
	RecordedDatePrimitiveElement  *primitiveElement           `json:"_recordedDate,omitempty"`
	ResultingCondition            []Reference                 `json:"resultingCondition,omitempty"`
	Location                      *Reference                  `json:"location,omitempty"`
	Seriousness                   *CodeableConcept            `json:"seriousness,omitempty"`
	Severity                      *CodeableConcept            `json:"severity,omitempty"`
	Outcome                       *CodeableConcept            `json:"outcome,omitempty"`
	Recorder                      *Reference                  `json:"recorder,omitempty"`
	Contributor                   []Reference                 `json:"contributor,omitempty"`
	SuspectEntity                 []AdverseEventSuspectEntity `json:"suspectEntity,omitempty"`
	SubjectMedicalHistory         []Reference                 `json:"subjectMedicalHistory,omitempty"`
	ReferenceDocument             []Reference                 `json:"referenceDocument,omitempty"`
	Study                         []Reference                 `json:"study,omitempty"`
}

func (r AdverseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AdverseEvent) marshalJSON() jsonAdverseEvent {
	m := jsonAdverseEvent{}
	m.ResourceType = "AdverseEvent"
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
	if r.Actuality.Value != nil {
		m.Actuality = r.Actuality
	}
	if r.Actuality.Id != nil || r.Actuality.Extension != nil {
		m.ActualityPrimitiveElement = &primitiveElement{Id: r.Actuality.Id, Extension: r.Actuality.Extension}
	}
	m.Category = r.Category
	m.Event = r.Event
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	if r.Date != nil && r.Date.Value != nil {
		m.Date = r.Date
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	if r.Detected != nil && r.Detected.Value != nil {
		m.Detected = r.Detected
	}
	if r.Detected != nil && (r.Detected.Id != nil || r.Detected.Extension != nil) {
		m.DetectedPrimitiveElement = &primitiveElement{Id: r.Detected.Id, Extension: r.Detected.Extension}
	}
	if r.RecordedDate != nil && r.RecordedDate.Value != nil {
		m.RecordedDate = r.RecordedDate
	}
	if r.RecordedDate != nil && (r.RecordedDate.Id != nil || r.RecordedDate.Extension != nil) {
		m.RecordedDatePrimitiveElement = &primitiveElement{Id: r.RecordedDate.Id, Extension: r.RecordedDate.Extension}
	}
	m.ResultingCondition = r.ResultingCondition
	m.Location = r.Location
	m.Seriousness = r.Seriousness
	m.Severity = r.Severity
	m.Outcome = r.Outcome
	m.Recorder = r.Recorder
	m.Contributor = r.Contributor
	m.SuspectEntity = r.SuspectEntity
	m.SubjectMedicalHistory = r.SubjectMedicalHistory
	m.ReferenceDocument = r.ReferenceDocument
	m.Study = r.Study
	return m
}
func (r *AdverseEvent) UnmarshalJSON(b []byte) error {
	var m jsonAdverseEvent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AdverseEvent) unmarshalJSON(m jsonAdverseEvent) error {
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
	r.Actuality = m.Actuality
	if m.ActualityPrimitiveElement != nil {
		r.Actuality.Id = m.ActualityPrimitiveElement.Id
		r.Actuality.Extension = m.ActualityPrimitiveElement.Extension
	}
	r.Category = m.Category
	r.Event = m.Event
	r.Subject = m.Subject
	r.Encounter = m.Encounter
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		if r.Date == nil {
			r.Date = &DateTime{}
		}
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Detected = m.Detected
	if m.DetectedPrimitiveElement != nil {
		if r.Detected == nil {
			r.Detected = &DateTime{}
		}
		r.Detected.Id = m.DetectedPrimitiveElement.Id
		r.Detected.Extension = m.DetectedPrimitiveElement.Extension
	}
	r.RecordedDate = m.RecordedDate
	if m.RecordedDatePrimitiveElement != nil {
		if r.RecordedDate == nil {
			r.RecordedDate = &DateTime{}
		}
		r.RecordedDate.Id = m.RecordedDatePrimitiveElement.Id
		r.RecordedDate.Extension = m.RecordedDatePrimitiveElement.Extension
	}
	r.ResultingCondition = m.ResultingCondition
	r.Location = m.Location
	r.Seriousness = m.Seriousness
	r.Severity = m.Severity
	r.Outcome = m.Outcome
	r.Recorder = m.Recorder
	r.Contributor = m.Contributor
	r.SuspectEntity = m.SuspectEntity
	r.SubjectMedicalHistory = m.SubjectMedicalHistory
	r.ReferenceDocument = m.ReferenceDocument
	r.Study = m.Study
	return nil
}
func (r AdverseEvent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "AdverseEvent"
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
	err = e.EncodeElement(r.Actuality, xml.StartElement{Name: xml.Name{Local: "actuality"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Category, xml.StartElement{Name: xml.Name{Local: "category"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Event, xml.StartElement{Name: xml.Name{Local: "event"}})
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
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Detected, xml.StartElement{Name: xml.Name{Local: "detected"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RecordedDate, xml.StartElement{Name: xml.Name{Local: "recordedDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ResultingCondition, xml.StartElement{Name: xml.Name{Local: "resultingCondition"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Location, xml.StartElement{Name: xml.Name{Local: "location"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Seriousness, xml.StartElement{Name: xml.Name{Local: "seriousness"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Severity, xml.StartElement{Name: xml.Name{Local: "severity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Outcome, xml.StartElement{Name: xml.Name{Local: "outcome"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Recorder, xml.StartElement{Name: xml.Name{Local: "recorder"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Contributor, xml.StartElement{Name: xml.Name{Local: "contributor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SuspectEntity, xml.StartElement{Name: xml.Name{Local: "suspectEntity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubjectMedicalHistory, xml.StartElement{Name: xml.Name{Local: "subjectMedicalHistory"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReferenceDocument, xml.StartElement{Name: xml.Name{Local: "referenceDocument"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Study, xml.StartElement{Name: xml.Name{Local: "study"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *AdverseEvent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.Identifier = &v
			case "actuality":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Actuality = v
			case "category":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = append(r.Category, v)
			case "event":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Event = &v
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
			case "date":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = &v
			case "detected":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detected = &v
			case "recordedDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RecordedDate = &v
			case "resultingCondition":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ResultingCondition = append(r.ResultingCondition, v)
			case "location":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = &v
			case "seriousness":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Seriousness = &v
			case "severity":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Severity = &v
			case "outcome":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Outcome = &v
			case "recorder":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Recorder = &v
			case "contributor":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Contributor = append(r.Contributor, v)
			case "suspectEntity":
				var v AdverseEventSuspectEntity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SuspectEntity = append(r.SuspectEntity, v)
			case "subjectMedicalHistory":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubjectMedicalHistory = append(r.SubjectMedicalHistory, v)
			case "referenceDocument":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReferenceDocument = append(r.ReferenceDocument, v)
			case "study":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Study = append(r.Study, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r AdverseEvent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Describes the entity that is suspected to have caused the adverse event.
type AdverseEventSuspectEntity struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identifies the actual instance of what caused the adverse event.  May be a substance, medication, medication administration, medication statement or a device.
	Instance Reference
	// Information on the possible cause of the event.
	Causality []AdverseEventSuspectEntityCausality
}
type jsonAdverseEventSuspectEntity struct {
	Id                *string                              `json:"id,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	Instance          Reference                            `json:"instance,omitempty"`
	Causality         []AdverseEventSuspectEntityCausality `json:"causality,omitempty"`
}

func (r AdverseEventSuspectEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AdverseEventSuspectEntity) marshalJSON() jsonAdverseEventSuspectEntity {
	m := jsonAdverseEventSuspectEntity{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Instance = r.Instance
	m.Causality = r.Causality
	return m
}
func (r *AdverseEventSuspectEntity) UnmarshalJSON(b []byte) error {
	var m jsonAdverseEventSuspectEntity
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AdverseEventSuspectEntity) unmarshalJSON(m jsonAdverseEventSuspectEntity) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Instance = m.Instance
	r.Causality = m.Causality
	return nil
}
func (r AdverseEventSuspectEntity) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Instance, xml.StartElement{Name: xml.Name{Local: "instance"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Causality, xml.StartElement{Name: xml.Name{Local: "causality"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *AdverseEventSuspectEntity) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "instance":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Instance = v
			case "causality":
				var v AdverseEventSuspectEntityCausality
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Causality = append(r.Causality, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r AdverseEventSuspectEntity) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Information on the possible cause of the event.
type AdverseEventSuspectEntityCausality struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Assessment of if the entity caused the event.
	Assessment *CodeableConcept
	// AdverseEvent.suspectEntity.causalityProductRelatedness.
	ProductRelatedness *String
	// AdverseEvent.suspectEntity.causalityAuthor.
	Author *Reference
	// ProbabilityScale | Bayesian | Checklist.
	Method *CodeableConcept
}
type jsonAdverseEventSuspectEntityCausality struct {
	Id                                 *string           `json:"id,omitempty"`
	Extension                          []Extension       `json:"extension,omitempty"`
	ModifierExtension                  []Extension       `json:"modifierExtension,omitempty"`
	Assessment                         *CodeableConcept  `json:"assessment,omitempty"`
	ProductRelatedness                 *String           `json:"productRelatedness,omitempty"`
	ProductRelatednessPrimitiveElement *primitiveElement `json:"_productRelatedness,omitempty"`
	Author                             *Reference        `json:"author,omitempty"`
	Method                             *CodeableConcept  `json:"method,omitempty"`
}

func (r AdverseEventSuspectEntityCausality) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AdverseEventSuspectEntityCausality) marshalJSON() jsonAdverseEventSuspectEntityCausality {
	m := jsonAdverseEventSuspectEntityCausality{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Assessment = r.Assessment
	if r.ProductRelatedness != nil && r.ProductRelatedness.Value != nil {
		m.ProductRelatedness = r.ProductRelatedness
	}
	if r.ProductRelatedness != nil && (r.ProductRelatedness.Id != nil || r.ProductRelatedness.Extension != nil) {
		m.ProductRelatednessPrimitiveElement = &primitiveElement{Id: r.ProductRelatedness.Id, Extension: r.ProductRelatedness.Extension}
	}
	m.Author = r.Author
	m.Method = r.Method
	return m
}
func (r *AdverseEventSuspectEntityCausality) UnmarshalJSON(b []byte) error {
	var m jsonAdverseEventSuspectEntityCausality
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AdverseEventSuspectEntityCausality) unmarshalJSON(m jsonAdverseEventSuspectEntityCausality) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Assessment = m.Assessment
	r.ProductRelatedness = m.ProductRelatedness
	if m.ProductRelatednessPrimitiveElement != nil {
		if r.ProductRelatedness == nil {
			r.ProductRelatedness = &String{}
		}
		r.ProductRelatedness.Id = m.ProductRelatednessPrimitiveElement.Id
		r.ProductRelatedness.Extension = m.ProductRelatednessPrimitiveElement.Extension
	}
	r.Author = m.Author
	r.Method = m.Method
	return nil
}
func (r AdverseEventSuspectEntityCausality) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Assessment, xml.StartElement{Name: xml.Name{Local: "assessment"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProductRelatedness, xml.StartElement{Name: xml.Name{Local: "productRelatedness"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Author, xml.StartElement{Name: xml.Name{Local: "author"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Method, xml.StartElement{Name: xml.Name{Local: "method"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *AdverseEventSuspectEntityCausality) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "assessment":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Assessment = &v
			case "productRelatedness":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductRelatedness = &v
			case "author":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Author = &v
			case "method":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Method = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r AdverseEventSuspectEntityCausality) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
