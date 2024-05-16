package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// Describes the intention of how one or more practitioners intend to deliver care for a particular patient, group or community for a period of time, possibly limited to care for a specific condition or set of conditions.
type CarePlan struct {
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
	// Business identifiers assigned to this care plan by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// The URL pointing to a FHIR-defined protocol, guideline, questionnaire or other definition that is adhered to in whole or in part by this CarePlan.
	InstantiatesCanonical []Canonical
	// The URL pointing to an externally maintained protocol, guideline, questionnaire or other definition that is adhered to in whole or in part by this CarePlan.
	InstantiatesUri []Uri
	// A care plan that is fulfilled in whole or in part by this care plan.
	BasedOn []Reference
	// Completed or terminated care plan whose function is taken by this new care plan.
	Replaces []Reference
	// A larger care plan of which this particular care plan is a component or step.
	PartOf []Reference
	// Indicates whether the plan is currently being acted upon, represents future intentions or is now a historical record.
	Status Code
	// Indicates the level of authority/intentionality associated with the care plan and where the care plan fits into the workflow chain.
	Intent Code
	// Identifies what "kind" of plan this is to support differentiation between multiple co-existing plans; e.g. "Home health", "psychiatric", "asthma", "disease management", "wellness plan", etc.
	Category []CodeableConcept
	// Human-friendly name for the care plan.
	Title *String
	// A description of the scope and nature of the plan.
	Description *String
	// Identifies the patient or group whose intended care is described by the plan.
	Subject Reference
	// The Encounter during which this CarePlan was created or to which the creation of this record is tightly associated.
	Encounter *Reference
	// Indicates when the plan did (or is intended to) come into effect and end.
	Period *Period
	// Represents when this particular CarePlan record was created in the system, which is often a system-generated date.
	Created *DateTime
	// When populated, the author is responsible for the care plan.  The care plan is attributed to the author.
	Author *Reference
	// Identifies the individual(s) or organization who provided the contents of the care plan.
	Contributor []Reference
	// Identifies all people and organizations who are expected to be involved in the care envisioned by this plan.
	CareTeam []Reference
	// Identifies the conditions/problems/concerns/diagnoses/etc. whose management and/or mitigation are handled by this plan.
	Addresses []Reference
	// Identifies portions of the patient's record that specifically influenced the formation of the plan.  These might include comorbidities, recent procedures, limitations, recent assessments, etc.
	SupportingInfo []Reference
	// Describes the intended objective(s) of carrying out the care plan.
	Goal []Reference
	// Identifies a planned action to occur as part of the plan.  For example, a medication to be used, lab tests to perform, self-monitoring, education, etc.
	Activity []CarePlanActivity
	// General notes about the care plan not covered elsewhere.
	Note []Annotation
}

func (r CarePlan) ResourceType() string {
	return "CarePlan"
}
func (r CarePlan) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonCarePlan struct {
	ResourceType                          string              `json:"resourceType"`
	Id                                    *Id                 `json:"id,omitempty"`
	IdPrimitiveElement                    *primitiveElement   `json:"_id,omitempty"`
	Meta                                  *Meta               `json:"meta,omitempty"`
	ImplicitRules                         *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement         *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                              *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement              *primitiveElement   `json:"_language,omitempty"`
	Text                                  *Narrative          `json:"text,omitempty"`
	Contained                             []containedResource `json:"contained,omitempty"`
	Extension                             []Extension         `json:"extension,omitempty"`
	ModifierExtension                     []Extension         `json:"modifierExtension,omitempty"`
	Identifier                            []Identifier        `json:"identifier,omitempty"`
	InstantiatesCanonical                 []Canonical         `json:"instantiatesCanonical,omitempty"`
	InstantiatesCanonicalPrimitiveElement []*primitiveElement `json:"_instantiatesCanonical,omitempty"`
	InstantiatesUri                       []Uri               `json:"instantiatesUri,omitempty"`
	InstantiatesUriPrimitiveElement       []*primitiveElement `json:"_instantiatesUri,omitempty"`
	BasedOn                               []Reference         `json:"basedOn,omitempty"`
	Replaces                              []Reference         `json:"replaces,omitempty"`
	PartOf                                []Reference         `json:"partOf,omitempty"`
	Status                                Code                `json:"status,omitempty"`
	StatusPrimitiveElement                *primitiveElement   `json:"_status,omitempty"`
	Intent                                Code                `json:"intent,omitempty"`
	IntentPrimitiveElement                *primitiveElement   `json:"_intent,omitempty"`
	Category                              []CodeableConcept   `json:"category,omitempty"`
	Title                                 *String             `json:"title,omitempty"`
	TitlePrimitiveElement                 *primitiveElement   `json:"_title,omitempty"`
	Description                           *String             `json:"description,omitempty"`
	DescriptionPrimitiveElement           *primitiveElement   `json:"_description,omitempty"`
	Subject                               Reference           `json:"subject,omitempty"`
	Encounter                             *Reference          `json:"encounter,omitempty"`
	Period                                *Period             `json:"period,omitempty"`
	Created                               *DateTime           `json:"created,omitempty"`
	CreatedPrimitiveElement               *primitiveElement   `json:"_created,omitempty"`
	Author                                *Reference          `json:"author,omitempty"`
	Contributor                           []Reference         `json:"contributor,omitempty"`
	CareTeam                              []Reference         `json:"careTeam,omitempty"`
	Addresses                             []Reference         `json:"addresses,omitempty"`
	SupportingInfo                        []Reference         `json:"supportingInfo,omitempty"`
	Goal                                  []Reference         `json:"goal,omitempty"`
	Activity                              []CarePlanActivity  `json:"activity,omitempty"`
	Note                                  []Annotation        `json:"note,omitempty"`
}

func (r CarePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CarePlan) marshalJSON() jsonCarePlan {
	m := jsonCarePlan{}
	m.ResourceType = "CarePlan"
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
	m.InstantiatesCanonical = r.InstantiatesCanonical
	anyInstantiatesCanonicalIdOrExtension := false
	for _, e := range r.InstantiatesCanonical {
		if e.Id != nil || e.Extension != nil {
			anyInstantiatesCanonicalIdOrExtension = true
			break
		}
	}
	if anyInstantiatesCanonicalIdOrExtension {
		m.InstantiatesCanonicalPrimitiveElement = make([]*primitiveElement, 0, len(r.InstantiatesCanonical))
		for _, e := range r.InstantiatesCanonical {
			if e.Id != nil || e.Extension != nil {
				m.InstantiatesCanonicalPrimitiveElement = append(m.InstantiatesCanonicalPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.InstantiatesCanonicalPrimitiveElement = append(m.InstantiatesCanonicalPrimitiveElement, nil)
			}
		}
	}
	m.InstantiatesUri = r.InstantiatesUri
	anyInstantiatesUriIdOrExtension := false
	for _, e := range r.InstantiatesUri {
		if e.Id != nil || e.Extension != nil {
			anyInstantiatesUriIdOrExtension = true
			break
		}
	}
	if anyInstantiatesUriIdOrExtension {
		m.InstantiatesUriPrimitiveElement = make([]*primitiveElement, 0, len(r.InstantiatesUri))
		for _, e := range r.InstantiatesUri {
			if e.Id != nil || e.Extension != nil {
				m.InstantiatesUriPrimitiveElement = append(m.InstantiatesUriPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.InstantiatesUriPrimitiveElement = append(m.InstantiatesUriPrimitiveElement, nil)
			}
		}
	}
	m.BasedOn = r.BasedOn
	m.Replaces = r.Replaces
	m.PartOf = r.PartOf
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Intent = r.Intent
	if r.Intent.Id != nil || r.Intent.Extension != nil {
		m.IntentPrimitiveElement = &primitiveElement{Id: r.Intent.Id, Extension: r.Intent.Extension}
	}
	m.Category = r.Category
	m.Title = r.Title
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	m.Period = r.Period
	m.Created = r.Created
	if r.Created != nil && (r.Created.Id != nil || r.Created.Extension != nil) {
		m.CreatedPrimitiveElement = &primitiveElement{Id: r.Created.Id, Extension: r.Created.Extension}
	}
	m.Author = r.Author
	m.Contributor = r.Contributor
	m.CareTeam = r.CareTeam
	m.Addresses = r.Addresses
	m.SupportingInfo = r.SupportingInfo
	m.Goal = r.Goal
	m.Activity = r.Activity
	m.Note = r.Note
	return m
}
func (r *CarePlan) UnmarshalJSON(b []byte) error {
	var m jsonCarePlan
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CarePlan) unmarshalJSON(m jsonCarePlan) error {
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
	r.InstantiatesCanonical = m.InstantiatesCanonical
	for i, e := range m.InstantiatesCanonicalPrimitiveElement {
		if len(r.InstantiatesCanonical) > i {
			r.InstantiatesCanonical[i].Id = e.Id
			r.InstantiatesCanonical[i].Extension = e.Extension
		} else {
			r.InstantiatesCanonical = append(r.InstantiatesCanonical, Canonical{Id: e.Id, Extension: e.Extension})
		}
	}
	r.InstantiatesUri = m.InstantiatesUri
	for i, e := range m.InstantiatesUriPrimitiveElement {
		if len(r.InstantiatesUri) > i {
			r.InstantiatesUri[i].Id = e.Id
			r.InstantiatesUri[i].Extension = e.Extension
		} else {
			r.InstantiatesUri = append(r.InstantiatesUri, Uri{Id: e.Id, Extension: e.Extension})
		}
	}
	r.BasedOn = m.BasedOn
	r.Replaces = m.Replaces
	r.PartOf = m.PartOf
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Intent = m.Intent
	if m.IntentPrimitiveElement != nil {
		r.Intent.Id = m.IntentPrimitiveElement.Id
		r.Intent.Extension = m.IntentPrimitiveElement.Extension
	}
	r.Category = m.Category
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Subject = m.Subject
	r.Encounter = m.Encounter
	r.Period = m.Period
	r.Created = m.Created
	if m.CreatedPrimitiveElement != nil {
		r.Created.Id = m.CreatedPrimitiveElement.Id
		r.Created.Extension = m.CreatedPrimitiveElement.Extension
	}
	r.Author = m.Author
	r.Contributor = m.Contributor
	r.CareTeam = m.CareTeam
	r.Addresses = m.Addresses
	r.SupportingInfo = m.SupportingInfo
	r.Goal = m.Goal
	r.Activity = m.Activity
	r.Note = m.Note
	return nil
}
func (r CarePlan) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Identifies a planned action to occur as part of the plan.  For example, a medication to be used, lab tests to perform, self-monitoring, education, etc.
type CarePlanActivity struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identifies the outcome at the point when the status of the activity is assessed.  For example, the outcome of an education activity could be patient understands (or not).
	OutcomeCodeableConcept []CodeableConcept
	// Details of the outcome or action resulting from the activity.  The reference to an "event" resource, such as Procedure or Encounter or Observation, is the result/outcome of the activity itself.  The activity can be conveyed using CarePlan.activity.detail OR using the CarePlan.activity.reference (a reference to a “request” resource).
	OutcomeReference []Reference
	// Notes about the adherence/status/progress of the activity.
	Progress []Annotation
	// The details of the proposed activity represented in a specific resource.
	Reference *Reference
	// A simple summary of a planned activity suitable for a general care plan system (e.g. form driven) that doesn't know about specific resources such as procedure etc.
	Detail *CarePlanActivityDetail
}
type jsonCarePlanActivity struct {
	Id                     *string                 `json:"id,omitempty"`
	Extension              []Extension             `json:"extension,omitempty"`
	ModifierExtension      []Extension             `json:"modifierExtension,omitempty"`
	OutcomeCodeableConcept []CodeableConcept       `json:"outcomeCodeableConcept,omitempty"`
	OutcomeReference       []Reference             `json:"outcomeReference,omitempty"`
	Progress               []Annotation            `json:"progress,omitempty"`
	Reference              *Reference              `json:"reference,omitempty"`
	Detail                 *CarePlanActivityDetail `json:"detail,omitempty"`
}

func (r CarePlanActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CarePlanActivity) marshalJSON() jsonCarePlanActivity {
	m := jsonCarePlanActivity{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.OutcomeCodeableConcept = r.OutcomeCodeableConcept
	m.OutcomeReference = r.OutcomeReference
	m.Progress = r.Progress
	m.Reference = r.Reference
	m.Detail = r.Detail
	return m
}
func (r *CarePlanActivity) UnmarshalJSON(b []byte) error {
	var m jsonCarePlanActivity
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CarePlanActivity) unmarshalJSON(m jsonCarePlanActivity) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.OutcomeCodeableConcept = m.OutcomeCodeableConcept
	r.OutcomeReference = m.OutcomeReference
	r.Progress = m.Progress
	r.Reference = m.Reference
	r.Detail = m.Detail
	return nil
}
func (r CarePlanActivity) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A simple summary of a planned activity suitable for a general care plan system (e.g. form driven) that doesn't know about specific resources such as procedure etc.
type CarePlanActivityDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A description of the kind of resource the in-line definition of a care plan activity is representing.  The CarePlan.activity.detail is an in-line definition when a resource is not referenced using CarePlan.activity.reference.  For example, a MedicationRequest, a ServiceRequest, or a CommunicationRequest.
	Kind *Code
	// The URL pointing to a FHIR-defined protocol, guideline, questionnaire or other definition that is adhered to in whole or in part by this CarePlan activity.
	InstantiatesCanonical []Canonical
	// The URL pointing to an externally maintained protocol, guideline, questionnaire or other definition that is adhered to in whole or in part by this CarePlan activity.
	InstantiatesUri []Uri
	// Detailed description of the type of planned activity; e.g. what lab test, what procedure, what kind of encounter.
	Code *CodeableConcept
	// Provides the rationale that drove the inclusion of this particular activity as part of the plan or the reason why the activity was prohibited.
	ReasonCode []CodeableConcept
	// Indicates another resource, such as the health condition(s), whose existence justifies this request and drove the inclusion of this particular activity as part of the plan.
	ReasonReference []Reference
	// Internal reference that identifies the goals that this activity is intended to contribute towards meeting.
	Goal []Reference
	// Identifies what progress is being made for the specific activity.
	Status Code
	// Provides reason why the activity isn't yet started, is on hold, was cancelled, etc.
	StatusReason *CodeableConcept
	// If true, indicates that the described activity is one that must NOT be engaged in when following the plan.  If false, or missing, indicates that the described activity is one that should be engaged in when following the plan.
	DoNotPerform *Boolean
	// The period, timing or frequency upon which the described activity is to occur.
	Scheduled isCarePlanActivityDetailScheduled
	// Identifies the facility where the activity will occur; e.g. home, hospital, specific clinic, etc.
	Location *Reference
	// Identifies who's expected to be involved in the activity.
	Performer []Reference
	// Identifies the food, drug or other product to be consumed or supplied in the activity.
	Product isCarePlanActivityDetailProduct
	// Identifies the quantity expected to be consumed in a given day.
	DailyAmount *Quantity
	// Identifies the quantity expected to be supplied, administered or consumed by the subject.
	Quantity *Quantity
	// This provides a textual description of constraints on the intended activity occurrence, including relation to other activities.  It may also include objectives, pre-conditions and end-conditions.  Finally, it may convey specifics about the activity such as body site, method, route, etc.
	Description *String
}
type isCarePlanActivityDetailScheduled interface {
	isCarePlanActivityDetailScheduled()
}

func (r Timing) isCarePlanActivityDetailScheduled() {}
func (r Period) isCarePlanActivityDetailScheduled() {}
func (r String) isCarePlanActivityDetailScheduled() {}

type isCarePlanActivityDetailProduct interface {
	isCarePlanActivityDetailProduct()
}

func (r CodeableConcept) isCarePlanActivityDetailProduct() {}
func (r Reference) isCarePlanActivityDetailProduct()       {}

type jsonCarePlanActivityDetail struct {
	Id                                    *string             `json:"id,omitempty"`
	Extension                             []Extension         `json:"extension,omitempty"`
	ModifierExtension                     []Extension         `json:"modifierExtension,omitempty"`
	Kind                                  *Code               `json:"kind,omitempty"`
	KindPrimitiveElement                  *primitiveElement   `json:"_kind,omitempty"`
	InstantiatesCanonical                 []Canonical         `json:"instantiatesCanonical,omitempty"`
	InstantiatesCanonicalPrimitiveElement []*primitiveElement `json:"_instantiatesCanonical,omitempty"`
	InstantiatesUri                       []Uri               `json:"instantiatesUri,omitempty"`
	InstantiatesUriPrimitiveElement       []*primitiveElement `json:"_instantiatesUri,omitempty"`
	Code                                  *CodeableConcept    `json:"code,omitempty"`
	ReasonCode                            []CodeableConcept   `json:"reasonCode,omitempty"`
	ReasonReference                       []Reference         `json:"reasonReference,omitempty"`
	Goal                                  []Reference         `json:"goal,omitempty"`
	Status                                Code                `json:"status,omitempty"`
	StatusPrimitiveElement                *primitiveElement   `json:"_status,omitempty"`
	StatusReason                          *CodeableConcept    `json:"statusReason,omitempty"`
	DoNotPerform                          *Boolean            `json:"doNotPerform,omitempty"`
	DoNotPerformPrimitiveElement          *primitiveElement   `json:"_doNotPerform,omitempty"`
	ScheduledTiming                       *Timing             `json:"scheduledTiming,omitempty"`
	ScheduledPeriod                       *Period             `json:"scheduledPeriod,omitempty"`
	ScheduledString                       *String             `json:"scheduledString,omitempty"`
	ScheduledStringPrimitiveElement       *primitiveElement   `json:"_scheduledString,omitempty"`
	Location                              *Reference          `json:"location,omitempty"`
	Performer                             []Reference         `json:"performer,omitempty"`
	ProductCodeableConcept                *CodeableConcept    `json:"productCodeableConcept,omitempty"`
	ProductReference                      *Reference          `json:"productReference,omitempty"`
	DailyAmount                           *Quantity           `json:"dailyAmount,omitempty"`
	Quantity                              *Quantity           `json:"quantity,omitempty"`
	Description                           *String             `json:"description,omitempty"`
	DescriptionPrimitiveElement           *primitiveElement   `json:"_description,omitempty"`
}

func (r CarePlanActivityDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CarePlanActivityDetail) marshalJSON() jsonCarePlanActivityDetail {
	m := jsonCarePlanActivityDetail{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Kind = r.Kind
	if r.Kind != nil && (r.Kind.Id != nil || r.Kind.Extension != nil) {
		m.KindPrimitiveElement = &primitiveElement{Id: r.Kind.Id, Extension: r.Kind.Extension}
	}
	m.InstantiatesCanonical = r.InstantiatesCanonical
	anyInstantiatesCanonicalIdOrExtension := false
	for _, e := range r.InstantiatesCanonical {
		if e.Id != nil || e.Extension != nil {
			anyInstantiatesCanonicalIdOrExtension = true
			break
		}
	}
	if anyInstantiatesCanonicalIdOrExtension {
		m.InstantiatesCanonicalPrimitiveElement = make([]*primitiveElement, 0, len(r.InstantiatesCanonical))
		for _, e := range r.InstantiatesCanonical {
			if e.Id != nil || e.Extension != nil {
				m.InstantiatesCanonicalPrimitiveElement = append(m.InstantiatesCanonicalPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.InstantiatesCanonicalPrimitiveElement = append(m.InstantiatesCanonicalPrimitiveElement, nil)
			}
		}
	}
	m.InstantiatesUri = r.InstantiatesUri
	anyInstantiatesUriIdOrExtension := false
	for _, e := range r.InstantiatesUri {
		if e.Id != nil || e.Extension != nil {
			anyInstantiatesUriIdOrExtension = true
			break
		}
	}
	if anyInstantiatesUriIdOrExtension {
		m.InstantiatesUriPrimitiveElement = make([]*primitiveElement, 0, len(r.InstantiatesUri))
		for _, e := range r.InstantiatesUri {
			if e.Id != nil || e.Extension != nil {
				m.InstantiatesUriPrimitiveElement = append(m.InstantiatesUriPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.InstantiatesUriPrimitiveElement = append(m.InstantiatesUriPrimitiveElement, nil)
			}
		}
	}
	m.Code = r.Code
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Goal = r.Goal
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.StatusReason = r.StatusReason
	m.DoNotPerform = r.DoNotPerform
	if r.DoNotPerform != nil && (r.DoNotPerform.Id != nil || r.DoNotPerform.Extension != nil) {
		m.DoNotPerformPrimitiveElement = &primitiveElement{Id: r.DoNotPerform.Id, Extension: r.DoNotPerform.Extension}
	}
	switch v := r.Scheduled.(type) {
	case Timing:
		m.ScheduledTiming = &v
	case *Timing:
		m.ScheduledTiming = v
	case Period:
		m.ScheduledPeriod = &v
	case *Period:
		m.ScheduledPeriod = v
	case String:
		m.ScheduledString = &v
		if v.Id != nil || v.Extension != nil {
			m.ScheduledStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.ScheduledString = v
		if v.Id != nil || v.Extension != nil {
			m.ScheduledStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.Location = r.Location
	m.Performer = r.Performer
	switch v := r.Product.(type) {
	case CodeableConcept:
		m.ProductCodeableConcept = &v
	case *CodeableConcept:
		m.ProductCodeableConcept = v
	case Reference:
		m.ProductReference = &v
	case *Reference:
		m.ProductReference = v
	}
	m.DailyAmount = r.DailyAmount
	m.Quantity = r.Quantity
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	return m
}
func (r *CarePlanActivityDetail) UnmarshalJSON(b []byte) error {
	var m jsonCarePlanActivityDetail
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CarePlanActivityDetail) unmarshalJSON(m jsonCarePlanActivityDetail) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Kind = m.Kind
	if m.KindPrimitiveElement != nil {
		r.Kind.Id = m.KindPrimitiveElement.Id
		r.Kind.Extension = m.KindPrimitiveElement.Extension
	}
	r.InstantiatesCanonical = m.InstantiatesCanonical
	for i, e := range m.InstantiatesCanonicalPrimitiveElement {
		if len(r.InstantiatesCanonical) > i {
			r.InstantiatesCanonical[i].Id = e.Id
			r.InstantiatesCanonical[i].Extension = e.Extension
		} else {
			r.InstantiatesCanonical = append(r.InstantiatesCanonical, Canonical{Id: e.Id, Extension: e.Extension})
		}
	}
	r.InstantiatesUri = m.InstantiatesUri
	for i, e := range m.InstantiatesUriPrimitiveElement {
		if len(r.InstantiatesUri) > i {
			r.InstantiatesUri[i].Id = e.Id
			r.InstantiatesUri[i].Extension = e.Extension
		} else {
			r.InstantiatesUri = append(r.InstantiatesUri, Uri{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Code = m.Code
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Goal = m.Goal
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.StatusReason = m.StatusReason
	r.DoNotPerform = m.DoNotPerform
	if m.DoNotPerformPrimitiveElement != nil {
		r.DoNotPerform.Id = m.DoNotPerformPrimitiveElement.Id
		r.DoNotPerform.Extension = m.DoNotPerformPrimitiveElement.Extension
	}
	if m.ScheduledTiming != nil {
		if r.Scheduled != nil {
			return fmt.Errorf("multiple values for field \"Scheduled\"")
		}
		v := m.ScheduledTiming
		r.Scheduled = v
	}
	if m.ScheduledPeriod != nil {
		if r.Scheduled != nil {
			return fmt.Errorf("multiple values for field \"Scheduled\"")
		}
		v := m.ScheduledPeriod
		r.Scheduled = v
	}
	if m.ScheduledString != nil || m.ScheduledStringPrimitiveElement != nil {
		if r.Scheduled != nil {
			return fmt.Errorf("multiple values for field \"Scheduled\"")
		}
		v := m.ScheduledString
		if m.ScheduledStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.ScheduledStringPrimitiveElement.Id
			v.Extension = m.ScheduledStringPrimitiveElement.Extension
		}
		r.Scheduled = v
	}
	r.Location = m.Location
	r.Performer = m.Performer
	if m.ProductCodeableConcept != nil {
		if r.Product != nil {
			return fmt.Errorf("multiple values for field \"Product\"")
		}
		v := m.ProductCodeableConcept
		r.Product = v
	}
	if m.ProductReference != nil {
		if r.Product != nil {
			return fmt.Errorf("multiple values for field \"Product\"")
		}
		v := m.ProductReference
		r.Product = v
	}
	r.DailyAmount = m.DailyAmount
	r.Quantity = m.Quantity
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	return nil
}
func (r CarePlanActivityDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
