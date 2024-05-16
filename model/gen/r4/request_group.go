package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A group of related requests that can be used to capture intended activities that have inter-dependencies such as "give this medication after that one".
type RequestGroup struct {
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
	// Allows a service to provide a unique, business identifier for the request.
	Identifier []Identifier
	// A canonical URL referencing a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this request.
	InstantiatesCanonical []Canonical
	// A URL referencing an externally defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this request.
	InstantiatesUri []Uri
	// A plan, proposal or order that is fulfilled in whole or in part by this request.
	BasedOn []Reference
	// Completed or terminated request(s) whose function is taken by this new request.
	Replaces []Reference
	// A shared identifier common to all requests that were authorized more or less simultaneously by a single author, representing the identifier of the requisition, prescription or similar form.
	GroupIdentifier *Identifier
	// The current state of the request. For request groups, the status reflects the status of all the requests in the group.
	Status Code
	// Indicates the level of authority/intentionality associated with the request and where the request fits into the workflow chain.
	Intent Code
	// Indicates how quickly the request should be addressed with respect to other requests.
	Priority *Code
	// A code that identifies what the overall request group is.
	Code *CodeableConcept
	// The subject for which the request group was created.
	Subject *Reference
	// Describes the context of the request group, if any.
	Encounter *Reference
	// Indicates when the request group was created.
	AuthoredOn *DateTime
	// Provides a reference to the author of the request group.
	Author *Reference
	// Describes the reason for the request group in coded or textual form.
	ReasonCode []CodeableConcept
	// Indicates another resource whose existence justifies this request group.
	ReasonReference []Reference
	// Provides a mechanism to communicate additional information about the response.
	Note []Annotation
	// The actions, if any, produced by the evaluation of the artifact.
	Action []RequestGroupAction
}

func (r RequestGroup) ResourceType() string {
	return "RequestGroup"
}
func (r RequestGroup) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonRequestGroup struct {
	ResourceType                          string               `json:"resourceType"`
	Id                                    *Id                  `json:"id,omitempty"`
	IdPrimitiveElement                    *primitiveElement    `json:"_id,omitempty"`
	Meta                                  *Meta                `json:"meta,omitempty"`
	ImplicitRules                         *Uri                 `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement         *primitiveElement    `json:"_implicitRules,omitempty"`
	Language                              *Code                `json:"language,omitempty"`
	LanguagePrimitiveElement              *primitiveElement    `json:"_language,omitempty"`
	Text                                  *Narrative           `json:"text,omitempty"`
	Contained                             []containedResource  `json:"contained,omitempty"`
	Extension                             []Extension          `json:"extension,omitempty"`
	ModifierExtension                     []Extension          `json:"modifierExtension,omitempty"`
	Identifier                            []Identifier         `json:"identifier,omitempty"`
	InstantiatesCanonical                 []Canonical          `json:"instantiatesCanonical,omitempty"`
	InstantiatesCanonicalPrimitiveElement []*primitiveElement  `json:"_instantiatesCanonical,omitempty"`
	InstantiatesUri                       []Uri                `json:"instantiatesUri,omitempty"`
	InstantiatesUriPrimitiveElement       []*primitiveElement  `json:"_instantiatesUri,omitempty"`
	BasedOn                               []Reference          `json:"basedOn,omitempty"`
	Replaces                              []Reference          `json:"replaces,omitempty"`
	GroupIdentifier                       *Identifier          `json:"groupIdentifier,omitempty"`
	Status                                Code                 `json:"status,omitempty"`
	StatusPrimitiveElement                *primitiveElement    `json:"_status,omitempty"`
	Intent                                Code                 `json:"intent,omitempty"`
	IntentPrimitiveElement                *primitiveElement    `json:"_intent,omitempty"`
	Priority                              *Code                `json:"priority,omitempty"`
	PriorityPrimitiveElement              *primitiveElement    `json:"_priority,omitempty"`
	Code                                  *CodeableConcept     `json:"code,omitempty"`
	Subject                               *Reference           `json:"subject,omitempty"`
	Encounter                             *Reference           `json:"encounter,omitempty"`
	AuthoredOn                            *DateTime            `json:"authoredOn,omitempty"`
	AuthoredOnPrimitiveElement            *primitiveElement    `json:"_authoredOn,omitempty"`
	Author                                *Reference           `json:"author,omitempty"`
	ReasonCode                            []CodeableConcept    `json:"reasonCode,omitempty"`
	ReasonReference                       []Reference          `json:"reasonReference,omitempty"`
	Note                                  []Annotation         `json:"note,omitempty"`
	Action                                []RequestGroupAction `json:"action,omitempty"`
}

func (r RequestGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RequestGroup) marshalJSON() jsonRequestGroup {
	m := jsonRequestGroup{}
	m.ResourceType = "RequestGroup"
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
	m.GroupIdentifier = r.GroupIdentifier
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Intent = r.Intent
	if r.Intent.Id != nil || r.Intent.Extension != nil {
		m.IntentPrimitiveElement = &primitiveElement{Id: r.Intent.Id, Extension: r.Intent.Extension}
	}
	m.Priority = r.Priority
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		m.PriorityPrimitiveElement = &primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
	}
	m.Code = r.Code
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	m.AuthoredOn = r.AuthoredOn
	if r.AuthoredOn != nil && (r.AuthoredOn.Id != nil || r.AuthoredOn.Extension != nil) {
		m.AuthoredOnPrimitiveElement = &primitiveElement{Id: r.AuthoredOn.Id, Extension: r.AuthoredOn.Extension}
	}
	m.Author = r.Author
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Note = r.Note
	m.Action = r.Action
	return m
}
func (r *RequestGroup) UnmarshalJSON(b []byte) error {
	var m jsonRequestGroup
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RequestGroup) unmarshalJSON(m jsonRequestGroup) error {
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
	r.GroupIdentifier = m.GroupIdentifier
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
	r.Priority = m.Priority
	if m.PriorityPrimitiveElement != nil {
		r.Priority.Id = m.PriorityPrimitiveElement.Id
		r.Priority.Extension = m.PriorityPrimitiveElement.Extension
	}
	r.Code = m.Code
	r.Subject = m.Subject
	r.Encounter = m.Encounter
	r.AuthoredOn = m.AuthoredOn
	if m.AuthoredOnPrimitiveElement != nil {
		r.AuthoredOn.Id = m.AuthoredOnPrimitiveElement.Id
		r.AuthoredOn.Extension = m.AuthoredOnPrimitiveElement.Extension
	}
	r.Author = m.Author
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Note = m.Note
	r.Action = m.Action
	return nil
}
func (r RequestGroup) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The actions, if any, produced by the evaluation of the artifact.
type RequestGroupAction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A user-visible prefix for the action.
	Prefix *String
	// The title of the action displayed to a user.
	Title *String
	// A short description of the action used to provide a summary to display to the user.
	Description *String
	// A text equivalent of the action to be performed. This provides a human-interpretable description of the action when the definition is consumed by a system that might not be capable of interpreting it dynamically.
	TextEquivalent *String
	// Indicates how quickly the action should be addressed with respect to other actions.
	Priority *Code
	// A code that provides meaning for the action or action group. For example, a section may have a LOINC code for a section of a documentation template.
	Code []CodeableConcept
	// Didactic or other informational resources associated with the action that can be provided to the CDS recipient. Information resources can include inline text commentary and links to web resources.
	Documentation []RelatedArtifact
	// An expression that describes applicability criteria, or start/stop conditions for the action.
	Condition []RequestGroupActionCondition
	// A relationship to another action such as "before" or "30-60 minutes after start of".
	RelatedAction []RequestGroupActionRelatedAction
	// An optional value describing when the action should be performed.
	Timing isRequestGroupActionTiming
	// The participant that should perform or be responsible for this action.
	Participant []Reference
	// The type of action to perform (create, update, remove).
	Type *CodeableConcept
	// Defines the grouping behavior for the action and its children.
	GroupingBehavior *Code
	// Defines the selection behavior for the action and its children.
	SelectionBehavior *Code
	// Defines expectations around whether an action is required.
	RequiredBehavior *Code
	// Defines whether the action should usually be preselected.
	PrecheckBehavior *Code
	// Defines whether the action can be selected multiple times.
	CardinalityBehavior *Code
	// The resource that is the target of the action (e.g. CommunicationRequest).
	Resource *Reference
	// Sub actions.
	Action []RequestGroupAction
}
type isRequestGroupActionTiming interface {
	isRequestGroupActionTiming()
}

func (r DateTime) isRequestGroupActionTiming() {}
func (r Age) isRequestGroupActionTiming()      {}
func (r Period) isRequestGroupActionTiming()   {}
func (r Duration) isRequestGroupActionTiming() {}
func (r Range) isRequestGroupActionTiming()    {}
func (r Timing) isRequestGroupActionTiming()   {}

type jsonRequestGroupAction struct {
	Id                                  *string                           `json:"id,omitempty"`
	Extension                           []Extension                       `json:"extension,omitempty"`
	ModifierExtension                   []Extension                       `json:"modifierExtension,omitempty"`
	Prefix                              *String                           `json:"prefix,omitempty"`
	PrefixPrimitiveElement              *primitiveElement                 `json:"_prefix,omitempty"`
	Title                               *String                           `json:"title,omitempty"`
	TitlePrimitiveElement               *primitiveElement                 `json:"_title,omitempty"`
	Description                         *String                           `json:"description,omitempty"`
	DescriptionPrimitiveElement         *primitiveElement                 `json:"_description,omitempty"`
	TextEquivalent                      *String                           `json:"textEquivalent,omitempty"`
	TextEquivalentPrimitiveElement      *primitiveElement                 `json:"_textEquivalent,omitempty"`
	Priority                            *Code                             `json:"priority,omitempty"`
	PriorityPrimitiveElement            *primitiveElement                 `json:"_priority,omitempty"`
	Code                                []CodeableConcept                 `json:"code,omitempty"`
	Documentation                       []RelatedArtifact                 `json:"documentation,omitempty"`
	Condition                           []RequestGroupActionCondition     `json:"condition,omitempty"`
	RelatedAction                       []RequestGroupActionRelatedAction `json:"relatedAction,omitempty"`
	TimingDateTime                      *DateTime                         `json:"timingDateTime,omitempty"`
	TimingDateTimePrimitiveElement      *primitiveElement                 `json:"_timingDateTime,omitempty"`
	TimingAge                           *Age                              `json:"timingAge,omitempty"`
	TimingPeriod                        *Period                           `json:"timingPeriod,omitempty"`
	TimingDuration                      *Duration                         `json:"timingDuration,omitempty"`
	TimingRange                         *Range                            `json:"timingRange,omitempty"`
	TimingTiming                        *Timing                           `json:"timingTiming,omitempty"`
	Participant                         []Reference                       `json:"participant,omitempty"`
	Type                                *CodeableConcept                  `json:"type,omitempty"`
	GroupingBehavior                    *Code                             `json:"groupingBehavior,omitempty"`
	GroupingBehaviorPrimitiveElement    *primitiveElement                 `json:"_groupingBehavior,omitempty"`
	SelectionBehavior                   *Code                             `json:"selectionBehavior,omitempty"`
	SelectionBehaviorPrimitiveElement   *primitiveElement                 `json:"_selectionBehavior,omitempty"`
	RequiredBehavior                    *Code                             `json:"requiredBehavior,omitempty"`
	RequiredBehaviorPrimitiveElement    *primitiveElement                 `json:"_requiredBehavior,omitempty"`
	PrecheckBehavior                    *Code                             `json:"precheckBehavior,omitempty"`
	PrecheckBehaviorPrimitiveElement    *primitiveElement                 `json:"_precheckBehavior,omitempty"`
	CardinalityBehavior                 *Code                             `json:"cardinalityBehavior,omitempty"`
	CardinalityBehaviorPrimitiveElement *primitiveElement                 `json:"_cardinalityBehavior,omitempty"`
	Resource                            *Reference                        `json:"resource,omitempty"`
	Action                              []RequestGroupAction              `json:"action,omitempty"`
}

func (r RequestGroupAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RequestGroupAction) marshalJSON() jsonRequestGroupAction {
	m := jsonRequestGroupAction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Prefix = r.Prefix
	if r.Prefix != nil && (r.Prefix.Id != nil || r.Prefix.Extension != nil) {
		m.PrefixPrimitiveElement = &primitiveElement{Id: r.Prefix.Id, Extension: r.Prefix.Extension}
	}
	m.Title = r.Title
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.TextEquivalent = r.TextEquivalent
	if r.TextEquivalent != nil && (r.TextEquivalent.Id != nil || r.TextEquivalent.Extension != nil) {
		m.TextEquivalentPrimitiveElement = &primitiveElement{Id: r.TextEquivalent.Id, Extension: r.TextEquivalent.Extension}
	}
	m.Priority = r.Priority
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		m.PriorityPrimitiveElement = &primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
	}
	m.Code = r.Code
	m.Documentation = r.Documentation
	m.Condition = r.Condition
	m.RelatedAction = r.RelatedAction
	switch v := r.Timing.(type) {
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
	case Age:
		m.TimingAge = &v
	case *Age:
		m.TimingAge = v
	case Period:
		m.TimingPeriod = &v
	case *Period:
		m.TimingPeriod = v
	case Duration:
		m.TimingDuration = &v
	case *Duration:
		m.TimingDuration = v
	case Range:
		m.TimingRange = &v
	case *Range:
		m.TimingRange = v
	case Timing:
		m.TimingTiming = &v
	case *Timing:
		m.TimingTiming = v
	}
	m.Participant = r.Participant
	m.Type = r.Type
	m.GroupingBehavior = r.GroupingBehavior
	if r.GroupingBehavior != nil && (r.GroupingBehavior.Id != nil || r.GroupingBehavior.Extension != nil) {
		m.GroupingBehaviorPrimitiveElement = &primitiveElement{Id: r.GroupingBehavior.Id, Extension: r.GroupingBehavior.Extension}
	}
	m.SelectionBehavior = r.SelectionBehavior
	if r.SelectionBehavior != nil && (r.SelectionBehavior.Id != nil || r.SelectionBehavior.Extension != nil) {
		m.SelectionBehaviorPrimitiveElement = &primitiveElement{Id: r.SelectionBehavior.Id, Extension: r.SelectionBehavior.Extension}
	}
	m.RequiredBehavior = r.RequiredBehavior
	if r.RequiredBehavior != nil && (r.RequiredBehavior.Id != nil || r.RequiredBehavior.Extension != nil) {
		m.RequiredBehaviorPrimitiveElement = &primitiveElement{Id: r.RequiredBehavior.Id, Extension: r.RequiredBehavior.Extension}
	}
	m.PrecheckBehavior = r.PrecheckBehavior
	if r.PrecheckBehavior != nil && (r.PrecheckBehavior.Id != nil || r.PrecheckBehavior.Extension != nil) {
		m.PrecheckBehaviorPrimitiveElement = &primitiveElement{Id: r.PrecheckBehavior.Id, Extension: r.PrecheckBehavior.Extension}
	}
	m.CardinalityBehavior = r.CardinalityBehavior
	if r.CardinalityBehavior != nil && (r.CardinalityBehavior.Id != nil || r.CardinalityBehavior.Extension != nil) {
		m.CardinalityBehaviorPrimitiveElement = &primitiveElement{Id: r.CardinalityBehavior.Id, Extension: r.CardinalityBehavior.Extension}
	}
	m.Resource = r.Resource
	m.Action = r.Action
	return m
}
func (r *RequestGroupAction) UnmarshalJSON(b []byte) error {
	var m jsonRequestGroupAction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RequestGroupAction) unmarshalJSON(m jsonRequestGroupAction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Prefix = m.Prefix
	if m.PrefixPrimitiveElement != nil {
		r.Prefix.Id = m.PrefixPrimitiveElement.Id
		r.Prefix.Extension = m.PrefixPrimitiveElement.Extension
	}
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
	r.TextEquivalent = m.TextEquivalent
	if m.TextEquivalentPrimitiveElement != nil {
		r.TextEquivalent.Id = m.TextEquivalentPrimitiveElement.Id
		r.TextEquivalent.Extension = m.TextEquivalentPrimitiveElement.Extension
	}
	r.Priority = m.Priority
	if m.PriorityPrimitiveElement != nil {
		r.Priority.Id = m.PriorityPrimitiveElement.Id
		r.Priority.Extension = m.PriorityPrimitiveElement.Extension
	}
	r.Code = m.Code
	r.Documentation = m.Documentation
	r.Condition = m.Condition
	r.RelatedAction = m.RelatedAction
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
	if m.TimingAge != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingAge
		r.Timing = v
	}
	if m.TimingPeriod != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingPeriod
		r.Timing = v
	}
	if m.TimingDuration != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingDuration
		r.Timing = v
	}
	if m.TimingRange != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingRange
		r.Timing = v
	}
	if m.TimingTiming != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingTiming
		r.Timing = v
	}
	r.Participant = m.Participant
	r.Type = m.Type
	r.GroupingBehavior = m.GroupingBehavior
	if m.GroupingBehaviorPrimitiveElement != nil {
		r.GroupingBehavior.Id = m.GroupingBehaviorPrimitiveElement.Id
		r.GroupingBehavior.Extension = m.GroupingBehaviorPrimitiveElement.Extension
	}
	r.SelectionBehavior = m.SelectionBehavior
	if m.SelectionBehaviorPrimitiveElement != nil {
		r.SelectionBehavior.Id = m.SelectionBehaviorPrimitiveElement.Id
		r.SelectionBehavior.Extension = m.SelectionBehaviorPrimitiveElement.Extension
	}
	r.RequiredBehavior = m.RequiredBehavior
	if m.RequiredBehaviorPrimitiveElement != nil {
		r.RequiredBehavior.Id = m.RequiredBehaviorPrimitiveElement.Id
		r.RequiredBehavior.Extension = m.RequiredBehaviorPrimitiveElement.Extension
	}
	r.PrecheckBehavior = m.PrecheckBehavior
	if m.PrecheckBehaviorPrimitiveElement != nil {
		r.PrecheckBehavior.Id = m.PrecheckBehaviorPrimitiveElement.Id
		r.PrecheckBehavior.Extension = m.PrecheckBehaviorPrimitiveElement.Extension
	}
	r.CardinalityBehavior = m.CardinalityBehavior
	if m.CardinalityBehaviorPrimitiveElement != nil {
		r.CardinalityBehavior.Id = m.CardinalityBehaviorPrimitiveElement.Id
		r.CardinalityBehavior.Extension = m.CardinalityBehaviorPrimitiveElement.Extension
	}
	r.Resource = m.Resource
	r.Action = m.Action
	return nil
}
func (r RequestGroupAction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// An expression that describes applicability criteria, or start/stop conditions for the action.
type RequestGroupActionCondition struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The kind of condition.
	Kind Code
	// An expression that returns true or false, indicating whether or not the condition is satisfied.
	Expression *Expression
}
type jsonRequestGroupActionCondition struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Kind                 Code              `json:"kind,omitempty"`
	KindPrimitiveElement *primitiveElement `json:"_kind,omitempty"`
	Expression           *Expression       `json:"expression,omitempty"`
}

func (r RequestGroupActionCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RequestGroupActionCondition) marshalJSON() jsonRequestGroupActionCondition {
	m := jsonRequestGroupActionCondition{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Kind = r.Kind
	if r.Kind.Id != nil || r.Kind.Extension != nil {
		m.KindPrimitiveElement = &primitiveElement{Id: r.Kind.Id, Extension: r.Kind.Extension}
	}
	m.Expression = r.Expression
	return m
}
func (r *RequestGroupActionCondition) UnmarshalJSON(b []byte) error {
	var m jsonRequestGroupActionCondition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RequestGroupActionCondition) unmarshalJSON(m jsonRequestGroupActionCondition) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Kind = m.Kind
	if m.KindPrimitiveElement != nil {
		r.Kind.Id = m.KindPrimitiveElement.Id
		r.Kind.Extension = m.KindPrimitiveElement.Extension
	}
	r.Expression = m.Expression
	return nil
}
func (r RequestGroupActionCondition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A relationship to another action such as "before" or "30-60 minutes after start of".
type RequestGroupActionRelatedAction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The element id of the action this is related to.
	ActionId Id
	// The relationship of this action to the related action.
	Relationship Code
	// A duration or range of durations to apply to the relationship. For example, 30-60 minutes before.
	Offset isRequestGroupActionRelatedActionOffset
}
type isRequestGroupActionRelatedActionOffset interface {
	isRequestGroupActionRelatedActionOffset()
}

func (r Duration) isRequestGroupActionRelatedActionOffset() {}
func (r Range) isRequestGroupActionRelatedActionOffset()    {}

type jsonRequestGroupActionRelatedAction struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	ActionId                     Id                `json:"actionId,omitempty"`
	ActionIdPrimitiveElement     *primitiveElement `json:"_actionId,omitempty"`
	Relationship                 Code              `json:"relationship,omitempty"`
	RelationshipPrimitiveElement *primitiveElement `json:"_relationship,omitempty"`
	OffsetDuration               *Duration         `json:"offsetDuration,omitempty"`
	OffsetRange                  *Range            `json:"offsetRange,omitempty"`
}

func (r RequestGroupActionRelatedAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RequestGroupActionRelatedAction) marshalJSON() jsonRequestGroupActionRelatedAction {
	m := jsonRequestGroupActionRelatedAction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.ActionId = r.ActionId
	if r.ActionId.Id != nil || r.ActionId.Extension != nil {
		m.ActionIdPrimitiveElement = &primitiveElement{Id: r.ActionId.Id, Extension: r.ActionId.Extension}
	}
	m.Relationship = r.Relationship
	if r.Relationship.Id != nil || r.Relationship.Extension != nil {
		m.RelationshipPrimitiveElement = &primitiveElement{Id: r.Relationship.Id, Extension: r.Relationship.Extension}
	}
	switch v := r.Offset.(type) {
	case Duration:
		m.OffsetDuration = &v
	case *Duration:
		m.OffsetDuration = v
	case Range:
		m.OffsetRange = &v
	case *Range:
		m.OffsetRange = v
	}
	return m
}
func (r *RequestGroupActionRelatedAction) UnmarshalJSON(b []byte) error {
	var m jsonRequestGroupActionRelatedAction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RequestGroupActionRelatedAction) unmarshalJSON(m jsonRequestGroupActionRelatedAction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ActionId = m.ActionId
	if m.ActionIdPrimitiveElement != nil {
		r.ActionId.Id = m.ActionIdPrimitiveElement.Id
		r.ActionId.Extension = m.ActionIdPrimitiveElement.Extension
	}
	r.Relationship = m.Relationship
	if m.RelationshipPrimitiveElement != nil {
		r.Relationship.Id = m.RelationshipPrimitiveElement.Id
		r.Relationship.Extension = m.RelationshipPrimitiveElement.Extension
	}
	if m.OffsetDuration != nil {
		if r.Offset != nil {
			return fmt.Errorf("multiple values for field \"Offset\"")
		}
		v := m.OffsetDuration
		r.Offset = v
	}
	if m.OffsetRange != nil {
		if r.Offset != nil {
			return fmt.Errorf("multiple values for field \"Offset\"")
		}
		v := m.OffsetRange
		r.Offset = v
	}
	return nil
}
func (r RequestGroupActionRelatedAction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
