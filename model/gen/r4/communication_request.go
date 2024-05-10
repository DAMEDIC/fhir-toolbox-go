package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A request to convey information; e.g. the CDS system proposes that an alert be sent to a responsible provider, the CDS system proposes that the public health agency be notified about a reportable condition.
type CommunicationRequest struct {
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
	// Business identifiers assigned to this communication request by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// A plan or proposal that is fulfilled in whole or in part by this request.
	BasedOn []Reference
	// Completed or terminated request(s) whose function is taken by this new request.
	Replaces []Reference
	// A shared identifier common to all requests that were authorized more or less simultaneously by a single author, representing the identifier of the requisition, prescription or similar form.
	GroupIdentifier *Identifier
	// The status of the proposal or order.
	Status Code
	// Captures the reason for the current state of the CommunicationRequest.
	StatusReason *CodeableConcept
	// The type of message to be sent such as alert, notification, reminder, instruction, etc.
	Category []CodeableConcept
	// Characterizes how quickly the proposed act must be initiated. Includes concepts such as stat, urgent, routine.
	Priority *Code
	// If true indicates that the CommunicationRequest is asking for the specified action to *not* occur.
	DoNotPerform *Boolean
	// A channel that was used for this communication (e.g. email, fax).
	Medium []CodeableConcept
	// The patient or group that is the focus of this communication request.
	Subject *Reference
	// Other resources that pertain to this communication request and to which this communication request should be associated.
	About []Reference
	// The Encounter during which this CommunicationRequest was created or to which the creation of this record is tightly associated.
	Encounter *Reference
	// Text, attachment(s), or resource(s) to be communicated to the recipient.
	Payload []CommunicationRequestPayload
	// The time when this communication is to occur.
	Occurrence isCommunicationRequestOccurrence
	// For draft requests, indicates the date of initial creation.  For requests with other statuses, indicates the date of activation.
	AuthoredOn *DateTime
	// The device, individual, or organization who initiated the request and has responsibility for its activation.
	Requester *Reference
	// The entity (e.g. person, organization, clinical information system, device, group, or care team) which is the intended target of the communication.
	Recipient []Reference
	// The entity (e.g. person, organization, clinical information system, or device) which is to be the source of the communication.
	Sender *Reference
	// Describes why the request is being made in coded or textual form.
	ReasonCode []CodeableConcept
	// Indicates another resource whose existence justifies this request.
	ReasonReference []Reference
	// Comments made about the request by the requester, sender, recipient, subject or other participants.
	Note []Annotation
}
type isCommunicationRequestOccurrence interface {
	isCommunicationRequestOccurrence()
}

func (r DateTime) isCommunicationRequestOccurrence() {}
func (r Period) isCommunicationRequestOccurrence()   {}

type jsonCommunicationRequest struct {
	ResourceType                       string                        `json:"resourceType"`
	Id                                 *Id                           `json:"id,omitempty"`
	IdPrimitiveElement                 *primitiveElement             `json:"_id,omitempty"`
	Meta                               *Meta                         `json:"meta,omitempty"`
	ImplicitRules                      *Uri                          `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement      *primitiveElement             `json:"_implicitRules,omitempty"`
	Language                           *Code                         `json:"language,omitempty"`
	LanguagePrimitiveElement           *primitiveElement             `json:"_language,omitempty"`
	Text                               *Narrative                    `json:"text,omitempty"`
	Contained                          []containedResource           `json:"contained,omitempty"`
	Extension                          []Extension                   `json:"extension,omitempty"`
	ModifierExtension                  []Extension                   `json:"modifierExtension,omitempty"`
	Identifier                         []Identifier                  `json:"identifier,omitempty"`
	BasedOn                            []Reference                   `json:"basedOn,omitempty"`
	Replaces                           []Reference                   `json:"replaces,omitempty"`
	GroupIdentifier                    *Identifier                   `json:"groupIdentifier,omitempty"`
	Status                             Code                          `json:"status,omitempty"`
	StatusPrimitiveElement             *primitiveElement             `json:"_status,omitempty"`
	StatusReason                       *CodeableConcept              `json:"statusReason,omitempty"`
	Category                           []CodeableConcept             `json:"category,omitempty"`
	Priority                           *Code                         `json:"priority,omitempty"`
	PriorityPrimitiveElement           *primitiveElement             `json:"_priority,omitempty"`
	DoNotPerform                       *Boolean                      `json:"doNotPerform,omitempty"`
	DoNotPerformPrimitiveElement       *primitiveElement             `json:"_doNotPerform,omitempty"`
	Medium                             []CodeableConcept             `json:"medium,omitempty"`
	Subject                            *Reference                    `json:"subject,omitempty"`
	About                              []Reference                   `json:"about,omitempty"`
	Encounter                          *Reference                    `json:"encounter,omitempty"`
	Payload                            []CommunicationRequestPayload `json:"payload,omitempty"`
	OccurrenceDateTime                 *DateTime                     `json:"occurrenceDateTime,omitempty"`
	OccurrenceDateTimePrimitiveElement *primitiveElement             `json:"_occurrenceDateTime,omitempty"`
	OccurrencePeriod                   *Period                       `json:"occurrencePeriod,omitempty"`
	AuthoredOn                         *DateTime                     `json:"authoredOn,omitempty"`
	AuthoredOnPrimitiveElement         *primitiveElement             `json:"_authoredOn,omitempty"`
	Requester                          *Reference                    `json:"requester,omitempty"`
	Recipient                          []Reference                   `json:"recipient,omitempty"`
	Sender                             *Reference                    `json:"sender,omitempty"`
	ReasonCode                         []CodeableConcept             `json:"reasonCode,omitempty"`
	ReasonReference                    []Reference                   `json:"reasonReference,omitempty"`
	Note                               []Annotation                  `json:"note,omitempty"`
}

func (r CommunicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CommunicationRequest) marshalJSON() jsonCommunicationRequest {
	m := jsonCommunicationRequest{}
	m.ResourceType = "CommunicationRequest"
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
	m.Replaces = r.Replaces
	m.GroupIdentifier = r.GroupIdentifier
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.StatusReason = r.StatusReason
	m.Category = r.Category
	m.Priority = r.Priority
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		m.PriorityPrimitiveElement = &primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
	}
	m.DoNotPerform = r.DoNotPerform
	if r.DoNotPerform != nil && (r.DoNotPerform.Id != nil || r.DoNotPerform.Extension != nil) {
		m.DoNotPerformPrimitiveElement = &primitiveElement{Id: r.DoNotPerform.Id, Extension: r.DoNotPerform.Extension}
	}
	m.Medium = r.Medium
	m.Subject = r.Subject
	m.About = r.About
	m.Encounter = r.Encounter
	m.Payload = r.Payload
	switch v := r.Occurrence.(type) {
	case DateTime:
		m.OccurrenceDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.OccurrenceDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.OccurrenceDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.OccurrenceDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.OccurrencePeriod = &v
	case *Period:
		m.OccurrencePeriod = v
	}
	m.AuthoredOn = r.AuthoredOn
	if r.AuthoredOn != nil && (r.AuthoredOn.Id != nil || r.AuthoredOn.Extension != nil) {
		m.AuthoredOnPrimitiveElement = &primitiveElement{Id: r.AuthoredOn.Id, Extension: r.AuthoredOn.Extension}
	}
	m.Requester = r.Requester
	m.Recipient = r.Recipient
	m.Sender = r.Sender
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Note = r.Note
	return m
}
func (r *CommunicationRequest) UnmarshalJSON(b []byte) error {
	var m jsonCommunicationRequest
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CommunicationRequest) unmarshalJSON(m jsonCommunicationRequest) error {
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
	r.Replaces = m.Replaces
	r.GroupIdentifier = m.GroupIdentifier
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.StatusReason = m.StatusReason
	r.Category = m.Category
	r.Priority = m.Priority
	if m.PriorityPrimitiveElement != nil {
		r.Priority.Id = m.PriorityPrimitiveElement.Id
		r.Priority.Extension = m.PriorityPrimitiveElement.Extension
	}
	r.DoNotPerform = m.DoNotPerform
	if m.DoNotPerformPrimitiveElement != nil {
		r.DoNotPerform.Id = m.DoNotPerformPrimitiveElement.Id
		r.DoNotPerform.Extension = m.DoNotPerformPrimitiveElement.Extension
	}
	r.Medium = m.Medium
	r.Subject = m.Subject
	r.About = m.About
	r.Encounter = m.Encounter
	r.Payload = m.Payload
	if m.OccurrenceDateTime != nil || m.OccurrenceDateTimePrimitiveElement != nil {
		if r.Occurrence != nil {
			return fmt.Errorf("multiple values for field \"Occurrence\"")
		}
		v := m.OccurrenceDateTime
		if m.OccurrenceDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.OccurrenceDateTimePrimitiveElement.Id
			v.Extension = m.OccurrenceDateTimePrimitiveElement.Extension
		}
		r.Occurrence = v
	}
	if m.OccurrencePeriod != nil {
		if r.Occurrence != nil {
			return fmt.Errorf("multiple values for field \"Occurrence\"")
		}
		v := m.OccurrencePeriod
		r.Occurrence = v
	}
	r.AuthoredOn = m.AuthoredOn
	if m.AuthoredOnPrimitiveElement != nil {
		r.AuthoredOn.Id = m.AuthoredOnPrimitiveElement.Id
		r.AuthoredOn.Extension = m.AuthoredOnPrimitiveElement.Extension
	}
	r.Requester = m.Requester
	r.Recipient = m.Recipient
	r.Sender = m.Sender
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Note = m.Note
	return nil
}
func (r CommunicationRequest) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Text, attachment(s), or resource(s) to be communicated to the recipient.
type CommunicationRequestPayload struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The communicated content (or for multi-part communications, one portion of the communication).
	Content isCommunicationRequestPayloadContent
}
type isCommunicationRequestPayloadContent interface {
	isCommunicationRequestPayloadContent()
}

func (r String) isCommunicationRequestPayloadContent()     {}
func (r Attachment) isCommunicationRequestPayloadContent() {}
func (r Reference) isCommunicationRequestPayloadContent()  {}

type jsonCommunicationRequestPayload struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	ContentString                 *String           `json:"contentString,omitempty"`
	ContentStringPrimitiveElement *primitiveElement `json:"_contentString,omitempty"`
	ContentAttachment             *Attachment       `json:"contentAttachment,omitempty"`
	ContentReference              *Reference        `json:"contentReference,omitempty"`
}

func (r CommunicationRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CommunicationRequestPayload) marshalJSON() jsonCommunicationRequestPayload {
	m := jsonCommunicationRequestPayload{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Content.(type) {
	case String:
		m.ContentString = &v
		if v.Id != nil || v.Extension != nil {
			m.ContentStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.ContentString = v
		if v.Id != nil || v.Extension != nil {
			m.ContentStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Attachment:
		m.ContentAttachment = &v
	case *Attachment:
		m.ContentAttachment = v
	case Reference:
		m.ContentReference = &v
	case *Reference:
		m.ContentReference = v
	}
	return m
}
func (r *CommunicationRequestPayload) UnmarshalJSON(b []byte) error {
	var m jsonCommunicationRequestPayload
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CommunicationRequestPayload) unmarshalJSON(m jsonCommunicationRequestPayload) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.ContentString != nil || m.ContentStringPrimitiveElement != nil {
		if r.Content != nil {
			return fmt.Errorf("multiple values for field \"Content\"")
		}
		v := m.ContentString
		if m.ContentStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.ContentStringPrimitiveElement.Id
			v.Extension = m.ContentStringPrimitiveElement.Extension
		}
		r.Content = v
	}
	if m.ContentAttachment != nil {
		if r.Content != nil {
			return fmt.Errorf("multiple values for field \"Content\"")
		}
		v := m.ContentAttachment
		r.Content = v
	}
	if m.ContentReference != nil {
		if r.Content != nil {
			return fmt.Errorf("multiple values for field \"Content\"")
		}
		v := m.ContentReference
		r.Content = v
	}
	return nil
}
func (r CommunicationRequestPayload) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
