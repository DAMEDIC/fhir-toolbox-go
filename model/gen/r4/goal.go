package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// Describes the intended objective(s) for a patient, group or organization care, for example, weight loss, restoring an activity of daily living, obtaining herd immunity via immunization, meeting a process improvement objective, etc.
type Goal struct {
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
	// Business identifiers assigned to this goal by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// The state of the goal throughout its lifecycle.
	LifecycleStatus Code
	// Describes the progression, or lack thereof, towards the goal against the target.
	AchievementStatus *CodeableConcept
	// Indicates a category the goal falls within.
	Category []CodeableConcept
	// Identifies the mutually agreed level of importance associated with reaching/sustaining the goal.
	Priority *CodeableConcept
	// Human-readable and/or coded description of a specific desired objective of care, such as "control blood pressure" or "negotiate an obstacle course" or "dance with child at wedding".
	Description CodeableConcept
	// Identifies the patient, group or organization for whom the goal is being established.
	Subject Reference
	// The date or event after which the goal should begin being pursued.
	Start isGoalStart
	// Indicates what should be done by when.
	Target []GoalTarget
	// Identifies when the current status.  I.e. When initially created, when achieved, when cancelled, etc.
	StatusDate *Date
	// Captures the reason for the current status.
	StatusReason *String
	// Indicates whose goal this is - patient goal, practitioner goal, etc.
	ExpressedBy *Reference
	// The identified conditions and other health record elements that are intended to be addressed by the goal.
	Addresses []Reference
	// Any comments related to the goal.
	Note []Annotation
	// Identifies the change (or lack of change) at the point when the status of the goal is assessed.
	OutcomeCode []CodeableConcept
	// Details of what's changed (or not changed).
	OutcomeReference []Reference
}

func (r Goal) ResourceType() string {
	return "Goal"
}
func (r Goal) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isGoalStart interface {
	isGoalStart()
}

func (r Date) isGoalStart()            {}
func (r CodeableConcept) isGoalStart() {}

type jsonGoal struct {
	ResourceType                    string              `json:"resourceType"`
	Id                              *Id                 `json:"id,omitempty"`
	IdPrimitiveElement              *primitiveElement   `json:"_id,omitempty"`
	Meta                            *Meta               `json:"meta,omitempty"`
	ImplicitRules                   *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement   *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                        *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement        *primitiveElement   `json:"_language,omitempty"`
	Text                            *Narrative          `json:"text,omitempty"`
	Contained                       []ContainedResource `json:"contained,omitempty"`
	Extension                       []Extension         `json:"extension,omitempty"`
	ModifierExtension               []Extension         `json:"modifierExtension,omitempty"`
	Identifier                      []Identifier        `json:"identifier,omitempty"`
	LifecycleStatus                 Code                `json:"lifecycleStatus,omitempty"`
	LifecycleStatusPrimitiveElement *primitiveElement   `json:"_lifecycleStatus,omitempty"`
	AchievementStatus               *CodeableConcept    `json:"achievementStatus,omitempty"`
	Category                        []CodeableConcept   `json:"category,omitempty"`
	Priority                        *CodeableConcept    `json:"priority,omitempty"`
	Description                     CodeableConcept     `json:"description,omitempty"`
	Subject                         Reference           `json:"subject,omitempty"`
	StartDate                       *Date               `json:"startDate,omitempty"`
	StartDatePrimitiveElement       *primitiveElement   `json:"_startDate,omitempty"`
	StartCodeableConcept            *CodeableConcept    `json:"startCodeableConcept,omitempty"`
	Target                          []GoalTarget        `json:"target,omitempty"`
	StatusDate                      *Date               `json:"statusDate,omitempty"`
	StatusDatePrimitiveElement      *primitiveElement   `json:"_statusDate,omitempty"`
	StatusReason                    *String             `json:"statusReason,omitempty"`
	StatusReasonPrimitiveElement    *primitiveElement   `json:"_statusReason,omitempty"`
	ExpressedBy                     *Reference          `json:"expressedBy,omitempty"`
	Addresses                       []Reference         `json:"addresses,omitempty"`
	Note                            []Annotation        `json:"note,omitempty"`
	OutcomeCode                     []CodeableConcept   `json:"outcomeCode,omitempty"`
	OutcomeReference                []Reference         `json:"outcomeReference,omitempty"`
}

func (r Goal) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Goal) marshalJSON() jsonGoal {
	m := jsonGoal{}
	m.ResourceType = "Goal"
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
	if r.LifecycleStatus.Value != nil {
		m.LifecycleStatus = r.LifecycleStatus
	}
	if r.LifecycleStatus.Id != nil || r.LifecycleStatus.Extension != nil {
		m.LifecycleStatusPrimitiveElement = &primitiveElement{Id: r.LifecycleStatus.Id, Extension: r.LifecycleStatus.Extension}
	}
	m.AchievementStatus = r.AchievementStatus
	m.Category = r.Category
	m.Priority = r.Priority
	m.Description = r.Description
	m.Subject = r.Subject
	switch v := r.Start.(type) {
	case Date:
		if v.Value != nil {
			m.StartDate = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.StartDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		if v.Value != nil {
			m.StartDate = v
		}
		if v.Id != nil || v.Extension != nil {
			m.StartDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case CodeableConcept:
		m.StartCodeableConcept = &v
	case *CodeableConcept:
		m.StartCodeableConcept = v
	}
	m.Target = r.Target
	if r.StatusDate != nil && r.StatusDate.Value != nil {
		m.StatusDate = r.StatusDate
	}
	if r.StatusDate != nil && (r.StatusDate.Id != nil || r.StatusDate.Extension != nil) {
		m.StatusDatePrimitiveElement = &primitiveElement{Id: r.StatusDate.Id, Extension: r.StatusDate.Extension}
	}
	if r.StatusReason != nil && r.StatusReason.Value != nil {
		m.StatusReason = r.StatusReason
	}
	if r.StatusReason != nil && (r.StatusReason.Id != nil || r.StatusReason.Extension != nil) {
		m.StatusReasonPrimitiveElement = &primitiveElement{Id: r.StatusReason.Id, Extension: r.StatusReason.Extension}
	}
	m.ExpressedBy = r.ExpressedBy
	m.Addresses = r.Addresses
	m.Note = r.Note
	m.OutcomeCode = r.OutcomeCode
	m.OutcomeReference = r.OutcomeReference
	return m
}
func (r *Goal) UnmarshalJSON(b []byte) error {
	var m jsonGoal
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Goal) unmarshalJSON(m jsonGoal) error {
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
	r.LifecycleStatus = m.LifecycleStatus
	if m.LifecycleStatusPrimitiveElement != nil {
		r.LifecycleStatus.Id = m.LifecycleStatusPrimitiveElement.Id
		r.LifecycleStatus.Extension = m.LifecycleStatusPrimitiveElement.Extension
	}
	r.AchievementStatus = m.AchievementStatus
	r.Category = m.Category
	r.Priority = m.Priority
	r.Description = m.Description
	r.Subject = m.Subject
	if m.StartDate != nil || m.StartDatePrimitiveElement != nil {
		if r.Start != nil {
			return fmt.Errorf("multiple values for field \"Start\"")
		}
		v := m.StartDate
		if m.StartDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.StartDatePrimitiveElement.Id
			v.Extension = m.StartDatePrimitiveElement.Extension
		}
		r.Start = v
	}
	if m.StartCodeableConcept != nil {
		if r.Start != nil {
			return fmt.Errorf("multiple values for field \"Start\"")
		}
		v := m.StartCodeableConcept
		r.Start = v
	}
	r.Target = m.Target
	r.StatusDate = m.StatusDate
	if m.StatusDatePrimitiveElement != nil {
		if r.StatusDate == nil {
			r.StatusDate = &Date{}
		}
		r.StatusDate.Id = m.StatusDatePrimitiveElement.Id
		r.StatusDate.Extension = m.StatusDatePrimitiveElement.Extension
	}
	r.StatusReason = m.StatusReason
	if m.StatusReasonPrimitiveElement != nil {
		if r.StatusReason == nil {
			r.StatusReason = &String{}
		}
		r.StatusReason.Id = m.StatusReasonPrimitiveElement.Id
		r.StatusReason.Extension = m.StatusReasonPrimitiveElement.Extension
	}
	r.ExpressedBy = m.ExpressedBy
	r.Addresses = m.Addresses
	r.Note = m.Note
	r.OutcomeCode = m.OutcomeCode
	r.OutcomeReference = m.OutcomeReference
	return nil
}
func (r Goal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Goal"
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
	err = e.EncodeElement(r.LifecycleStatus, xml.StartElement{Name: xml.Name{Local: "lifecycleStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AchievementStatus, xml.StartElement{Name: xml.Name{Local: "achievementStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Category, xml.StartElement{Name: xml.Name{Local: "category"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Priority, xml.StartElement{Name: xml.Name{Local: "priority"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Subject, xml.StartElement{Name: xml.Name{Local: "subject"}})
	if err != nil {
		return err
	}
	switch v := r.Start.(type) {
	case Date, *Date:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "startDate"}})
		if err != nil {
			return err
		}
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "startCodeableConcept"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Target, xml.StartElement{Name: xml.Name{Local: "target"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.StatusDate, xml.StartElement{Name: xml.Name{Local: "statusDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.StatusReason, xml.StartElement{Name: xml.Name{Local: "statusReason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ExpressedBy, xml.StartElement{Name: xml.Name{Local: "expressedBy"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Addresses, xml.StartElement{Name: xml.Name{Local: "addresses"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OutcomeCode, xml.StartElement{Name: xml.Name{Local: "outcomeCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OutcomeReference, xml.StartElement{Name: xml.Name{Local: "outcomeReference"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Goal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "lifecycleStatus":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LifecycleStatus = v
			case "achievementStatus":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AchievementStatus = &v
			case "category":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = append(r.Category, v)
			case "priority":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Priority = &v
			case "description":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = v
			case "subject":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = v
			case "startDate":
				if r.Start != nil {
					return fmt.Errorf("multiple values for field \"Start\"")
				}
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Start = v
			case "startCodeableConcept":
				if r.Start != nil {
					return fmt.Errorf("multiple values for field \"Start\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Start = v
			case "target":
				var v GoalTarget
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Target = append(r.Target, v)
			case "statusDate":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StatusDate = &v
			case "statusReason":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StatusReason = &v
			case "expressedBy":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ExpressedBy = &v
			case "addresses":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Addresses = append(r.Addresses, v)
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			case "outcomeCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OutcomeCode = append(r.OutcomeCode, v)
			case "outcomeReference":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OutcomeReference = append(r.OutcomeReference, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Goal) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Indicates what should be done by when.
type GoalTarget struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The parameter whose value is being tracked, e.g. body weight, blood pressure, or hemoglobin A1c level.
	Measure *CodeableConcept
	// The target value of the focus to be achieved to signify the fulfillment of the goal, e.g. 150 pounds, 7.0%. Either the high or low or both values of the range can be specified. When a low value is missing, it indicates that the goal is achieved at any focus value at or below the high value. Similarly, if the high value is missing, it indicates that the goal is achieved at any focus value at or above the low value.
	Detail isGoalTargetDetail
	// Indicates either the date or the duration after start by which the goal should be met.
	Due isGoalTargetDue
}
type isGoalTargetDetail interface {
	isGoalTargetDetail()
}

func (r Quantity) isGoalTargetDetail()        {}
func (r Range) isGoalTargetDetail()           {}
func (r CodeableConcept) isGoalTargetDetail() {}
func (r String) isGoalTargetDetail()          {}
func (r Boolean) isGoalTargetDetail()         {}
func (r Integer) isGoalTargetDetail()         {}
func (r Ratio) isGoalTargetDetail()           {}

type isGoalTargetDue interface {
	isGoalTargetDue()
}

func (r Date) isGoalTargetDue()     {}
func (r Duration) isGoalTargetDue() {}

type jsonGoalTarget struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	Measure                       *CodeableConcept  `json:"measure,omitempty"`
	DetailQuantity                *Quantity         `json:"detailQuantity,omitempty"`
	DetailRange                   *Range            `json:"detailRange,omitempty"`
	DetailCodeableConcept         *CodeableConcept  `json:"detailCodeableConcept,omitempty"`
	DetailString                  *String           `json:"detailString,omitempty"`
	DetailStringPrimitiveElement  *primitiveElement `json:"_detailString,omitempty"`
	DetailBoolean                 *Boolean          `json:"detailBoolean,omitempty"`
	DetailBooleanPrimitiveElement *primitiveElement `json:"_detailBoolean,omitempty"`
	DetailInteger                 *Integer          `json:"detailInteger,omitempty"`
	DetailIntegerPrimitiveElement *primitiveElement `json:"_detailInteger,omitempty"`
	DetailRatio                   *Ratio            `json:"detailRatio,omitempty"`
	DueDate                       *Date             `json:"dueDate,omitempty"`
	DueDatePrimitiveElement       *primitiveElement `json:"_dueDate,omitempty"`
	DueDuration                   *Duration         `json:"dueDuration,omitempty"`
}

func (r GoalTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r GoalTarget) marshalJSON() jsonGoalTarget {
	m := jsonGoalTarget{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Measure = r.Measure
	switch v := r.Detail.(type) {
	case Quantity:
		m.DetailQuantity = &v
	case *Quantity:
		m.DetailQuantity = v
	case Range:
		m.DetailRange = &v
	case *Range:
		m.DetailRange = v
	case CodeableConcept:
		m.DetailCodeableConcept = &v
	case *CodeableConcept:
		m.DetailCodeableConcept = v
	case String:
		if v.Value != nil {
			m.DetailString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.DetailStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.DetailString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.DetailStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Boolean:
		if v.Value != nil {
			m.DetailBoolean = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.DetailBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		if v.Value != nil {
			m.DetailBoolean = v
		}
		if v.Id != nil || v.Extension != nil {
			m.DetailBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Integer:
		if v.Value != nil {
			m.DetailInteger = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.DetailIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Integer:
		if v.Value != nil {
			m.DetailInteger = v
		}
		if v.Id != nil || v.Extension != nil {
			m.DetailIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Ratio:
		m.DetailRatio = &v
	case *Ratio:
		m.DetailRatio = v
	}
	switch v := r.Due.(type) {
	case Date:
		if v.Value != nil {
			m.DueDate = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.DueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		if v.Value != nil {
			m.DueDate = v
		}
		if v.Id != nil || v.Extension != nil {
			m.DueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Duration:
		m.DueDuration = &v
	case *Duration:
		m.DueDuration = v
	}
	return m
}
func (r *GoalTarget) UnmarshalJSON(b []byte) error {
	var m jsonGoalTarget
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *GoalTarget) unmarshalJSON(m jsonGoalTarget) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Measure = m.Measure
	if m.DetailQuantity != nil {
		if r.Detail != nil {
			return fmt.Errorf("multiple values for field \"Detail\"")
		}
		v := m.DetailQuantity
		r.Detail = v
	}
	if m.DetailRange != nil {
		if r.Detail != nil {
			return fmt.Errorf("multiple values for field \"Detail\"")
		}
		v := m.DetailRange
		r.Detail = v
	}
	if m.DetailCodeableConcept != nil {
		if r.Detail != nil {
			return fmt.Errorf("multiple values for field \"Detail\"")
		}
		v := m.DetailCodeableConcept
		r.Detail = v
	}
	if m.DetailString != nil || m.DetailStringPrimitiveElement != nil {
		if r.Detail != nil {
			return fmt.Errorf("multiple values for field \"Detail\"")
		}
		v := m.DetailString
		if m.DetailStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.DetailStringPrimitiveElement.Id
			v.Extension = m.DetailStringPrimitiveElement.Extension
		}
		r.Detail = v
	}
	if m.DetailBoolean != nil || m.DetailBooleanPrimitiveElement != nil {
		if r.Detail != nil {
			return fmt.Errorf("multiple values for field \"Detail\"")
		}
		v := m.DetailBoolean
		if m.DetailBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.DetailBooleanPrimitiveElement.Id
			v.Extension = m.DetailBooleanPrimitiveElement.Extension
		}
		r.Detail = v
	}
	if m.DetailInteger != nil || m.DetailIntegerPrimitiveElement != nil {
		if r.Detail != nil {
			return fmt.Errorf("multiple values for field \"Detail\"")
		}
		v := m.DetailInteger
		if m.DetailIntegerPrimitiveElement != nil {
			if v == nil {
				v = &Integer{}
			}
			v.Id = m.DetailIntegerPrimitiveElement.Id
			v.Extension = m.DetailIntegerPrimitiveElement.Extension
		}
		r.Detail = v
	}
	if m.DetailRatio != nil {
		if r.Detail != nil {
			return fmt.Errorf("multiple values for field \"Detail\"")
		}
		v := m.DetailRatio
		r.Detail = v
	}
	if m.DueDate != nil || m.DueDatePrimitiveElement != nil {
		if r.Due != nil {
			return fmt.Errorf("multiple values for field \"Due\"")
		}
		v := m.DueDate
		if m.DueDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.DueDatePrimitiveElement.Id
			v.Extension = m.DueDatePrimitiveElement.Extension
		}
		r.Due = v
	}
	if m.DueDuration != nil {
		if r.Due != nil {
			return fmt.Errorf("multiple values for field \"Due\"")
		}
		v := m.DueDuration
		r.Due = v
	}
	return nil
}
func (r GoalTarget) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Measure, xml.StartElement{Name: xml.Name{Local: "measure"}})
	if err != nil {
		return err
	}
	switch v := r.Detail.(type) {
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "detailQuantity"}})
		if err != nil {
			return err
		}
	case Range, *Range:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "detailRange"}})
		if err != nil {
			return err
		}
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "detailCodeableConcept"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "detailString"}})
		if err != nil {
			return err
		}
	case Boolean, *Boolean:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "detailBoolean"}})
		if err != nil {
			return err
		}
	case Integer, *Integer:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "detailInteger"}})
		if err != nil {
			return err
		}
	case Ratio, *Ratio:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "detailRatio"}})
		if err != nil {
			return err
		}
	}
	switch v := r.Due.(type) {
	case Date, *Date:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "dueDate"}})
		if err != nil {
			return err
		}
	case Duration, *Duration:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "dueDuration"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *GoalTarget) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "measure":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Measure = &v
			case "detailQuantity":
				if r.Detail != nil {
					return fmt.Errorf("multiple values for field \"Detail\"")
				}
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = v
			case "detailRange":
				if r.Detail != nil {
					return fmt.Errorf("multiple values for field \"Detail\"")
				}
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = v
			case "detailCodeableConcept":
				if r.Detail != nil {
					return fmt.Errorf("multiple values for field \"Detail\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = v
			case "detailString":
				if r.Detail != nil {
					return fmt.Errorf("multiple values for field \"Detail\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = v
			case "detailBoolean":
				if r.Detail != nil {
					return fmt.Errorf("multiple values for field \"Detail\"")
				}
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = v
			case "detailInteger":
				if r.Detail != nil {
					return fmt.Errorf("multiple values for field \"Detail\"")
				}
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = v
			case "detailRatio":
				if r.Detail != nil {
					return fmt.Errorf("multiple values for field \"Detail\"")
				}
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = v
			case "dueDate":
				if r.Due != nil {
					return fmt.Errorf("multiple values for field \"Due\"")
				}
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Due = v
			case "dueDuration":
				if r.Due != nil {
					return fmt.Errorf("multiple values for field \"Due\"")
				}
				var v Duration
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Due = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r GoalTarget) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
