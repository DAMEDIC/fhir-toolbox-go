package r4

import (
	"encoding/json"
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
	Contained                       []containedResource `json:"contained,omitempty"`
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
	m.LifecycleStatus = r.LifecycleStatus
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
		m.StartDate = &v
		if v.Id != nil || v.Extension != nil {
			m.StartDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		m.StartDate = v
		if v.Id != nil || v.Extension != nil {
			m.StartDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case CodeableConcept:
		m.StartCodeableConcept = &v
	case *CodeableConcept:
		m.StartCodeableConcept = v
	}
	m.Target = r.Target
	m.StatusDate = r.StatusDate
	if r.StatusDate != nil && (r.StatusDate.Id != nil || r.StatusDate.Extension != nil) {
		m.StatusDatePrimitiveElement = &primitiveElement{Id: r.StatusDate.Id, Extension: r.StatusDate.Extension}
	}
	m.StatusReason = r.StatusReason
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
		r.StatusDate.Id = m.StatusDatePrimitiveElement.Id
		r.StatusDate.Extension = m.StatusDatePrimitiveElement.Extension
	}
	r.StatusReason = m.StatusReason
	if m.StatusReasonPrimitiveElement != nil {
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
func (r Goal) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
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
		m.DetailString = &v
		if v.Id != nil || v.Extension != nil {
			m.DetailStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.DetailString = v
		if v.Id != nil || v.Extension != nil {
			m.DetailStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Boolean:
		m.DetailBoolean = &v
		if v.Id != nil || v.Extension != nil {
			m.DetailBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		m.DetailBoolean = v
		if v.Id != nil || v.Extension != nil {
			m.DetailBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Integer:
		m.DetailInteger = &v
		if v.Id != nil || v.Extension != nil {
			m.DetailIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Integer:
		m.DetailInteger = v
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
		m.DueDate = &v
		if v.Id != nil || v.Extension != nil {
			m.DueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		m.DueDate = v
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
func (r GoalTarget) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
