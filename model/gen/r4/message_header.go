package r4

import (
	"encoding/json"
	"fmt"
)

// The header for a message exchange that is either requesting or responding to an action.  The reference(s) that are the subject of the action as well as other information related to the action are typically transmitted in a bundle in which the MessageHeader resource instance is the first resource in the bundle.
//
// Many implementations are not prepared to use REST and need a messaging based infrastructure.
type MessageHeader struct {
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
	// Code that identifies the event this message represents and connects it with its definition. Events defined as part of the FHIR specification have the system value "http://terminology.hl7.org/CodeSystem/message-events".  Alternatively uri to the EventDefinition.
	Event isMessageHeaderEvent
	// The destination application which the message is intended for.
	Destination []MessageHeaderDestination
	// Identifies the sending system to allow the use of a trust relationship.
	Sender *Reference
	// The person or device that performed the data entry leading to this message. When there is more than one candidate, pick the most proximal to the message. Can provide other enterers in extensions.
	Enterer *Reference
	// The logical author of the message - the person or device that decided the described event should happen. When there is more than one candidate, pick the most proximal to the MessageHeader. Can provide other authors in extensions.
	Author *Reference
	// The source application from which this message originated.
	Source MessageHeaderSource
	// The person or organization that accepts overall responsibility for the contents of the message. The implication is that the message event happened under the policies of the responsible party.
	Responsible *Reference
	// Coded indication of the cause for the event - indicates  a reason for the occurrence of the event that is a focus of this message.
	Reason *CodeableConcept
	// Information about the message that this message is a response to.  Only present if this message is a response.
	Response *MessageHeaderResponse
	// The actual data of the message - a reference to the root/focus class of the event.
	Focus []Reference
	// Permanent link to the MessageDefinition for this message.
	Definition *Canonical
}
type isMessageHeaderEvent interface {
	isMessageHeaderEvent()
}

func (r Coding) isMessageHeaderEvent() {}
func (r Uri) isMessageHeaderEvent()    {}

type jsonMessageHeader struct {
	ResourceType                  string                     `json:"resourceType"`
	Id                            *Id                        `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement          `json:"_id,omitempty"`
	Meta                          *Meta                      `json:"meta,omitempty"`
	ImplicitRules                 *Uri                       `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement          `json:"_implicitRules,omitempty"`
	Language                      *Code                      `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement          `json:"_language,omitempty"`
	Text                          *Narrative                 `json:"text,omitempty"`
	Contained                     []containedResource        `json:"contained,omitempty"`
	Extension                     []Extension                `json:"extension,omitempty"`
	ModifierExtension             []Extension                `json:"modifierExtension,omitempty"`
	EventCoding                   *Coding                    `json:"eventCoding,omitempty"`
	EventUri                      *Uri                       `json:"eventUri,omitempty"`
	EventUriPrimitiveElement      *primitiveElement          `json:"_eventUri,omitempty"`
	Destination                   []MessageHeaderDestination `json:"destination,omitempty"`
	Sender                        *Reference                 `json:"sender,omitempty"`
	Enterer                       *Reference                 `json:"enterer,omitempty"`
	Author                        *Reference                 `json:"author,omitempty"`
	Source                        MessageHeaderSource        `json:"source,omitempty"`
	Responsible                   *Reference                 `json:"responsible,omitempty"`
	Reason                        *CodeableConcept           `json:"reason,omitempty"`
	Response                      *MessageHeaderResponse     `json:"response,omitempty"`
	Focus                         []Reference                `json:"focus,omitempty"`
	Definition                    *Canonical                 `json:"definition,omitempty"`
	DefinitionPrimitiveElement    *primitiveElement          `json:"_definition,omitempty"`
}

func (r MessageHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MessageHeader) marshalJSON() jsonMessageHeader {
	m := jsonMessageHeader{}
	m.ResourceType = "MessageHeader"
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
	switch v := r.Event.(type) {
	case Coding:
		m.EventCoding = &v
	case *Coding:
		m.EventCoding = v
	case Uri:
		m.EventUri = &v
		if v.Id != nil || v.Extension != nil {
			m.EventUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uri:
		m.EventUri = v
		if v.Id != nil || v.Extension != nil {
			m.EventUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.Destination = r.Destination
	m.Sender = r.Sender
	m.Enterer = r.Enterer
	m.Author = r.Author
	m.Source = r.Source
	m.Responsible = r.Responsible
	m.Reason = r.Reason
	m.Response = r.Response
	m.Focus = r.Focus
	m.Definition = r.Definition
	if r.Definition != nil && (r.Definition.Id != nil || r.Definition.Extension != nil) {
		m.DefinitionPrimitiveElement = &primitiveElement{Id: r.Definition.Id, Extension: r.Definition.Extension}
	}
	return m
}
func (r *MessageHeader) UnmarshalJSON(b []byte) error {
	var m jsonMessageHeader
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MessageHeader) unmarshalJSON(m jsonMessageHeader) error {
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
	if m.EventCoding != nil {
		if r.Event != nil {
			return fmt.Errorf("multiple values for field \"Event\"")
		}
		v := m.EventCoding
		r.Event = v
	}
	if m.EventUri != nil || m.EventUriPrimitiveElement != nil {
		if r.Event != nil {
			return fmt.Errorf("multiple values for field \"Event\"")
		}
		v := m.EventUri
		if m.EventUriPrimitiveElement != nil {
			if v == nil {
				v = &Uri{}
			}
			v.Id = m.EventUriPrimitiveElement.Id
			v.Extension = m.EventUriPrimitiveElement.Extension
		}
		r.Event = v
	}
	r.Destination = m.Destination
	r.Sender = m.Sender
	r.Enterer = m.Enterer
	r.Author = m.Author
	r.Source = m.Source
	r.Responsible = m.Responsible
	r.Reason = m.Reason
	r.Response = m.Response
	r.Focus = m.Focus
	r.Definition = m.Definition
	if m.DefinitionPrimitiveElement != nil {
		r.Definition.Id = m.DefinitionPrimitiveElement.Id
		r.Definition.Extension = m.DefinitionPrimitiveElement.Extension
	}
	return nil
}
func (r MessageHeader) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The destination application which the message is intended for.
type MessageHeaderDestination struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Human-readable name for the target system.
	Name *String
	// Identifies the target end system in situations where the initial message transmission is to an intermediary system.
	Target *Reference
	// Indicates where the message should be routed to.
	Endpoint Url
	// Allows data conveyed by a message to be addressed to a particular person or department when routing to a specific application isn't sufficient.
	Receiver *Reference
}
type jsonMessageHeaderDestination struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Name                     *String           `json:"name,omitempty"`
	NamePrimitiveElement     *primitiveElement `json:"_name,omitempty"`
	Target                   *Reference        `json:"target,omitempty"`
	Endpoint                 Url               `json:"endpoint,omitempty"`
	EndpointPrimitiveElement *primitiveElement `json:"_endpoint,omitempty"`
	Receiver                 *Reference        `json:"receiver,omitempty"`
}

func (r MessageHeaderDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MessageHeaderDestination) marshalJSON() jsonMessageHeaderDestination {
	m := jsonMessageHeaderDestination{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Target = r.Target
	m.Endpoint = r.Endpoint
	if r.Endpoint.Id != nil || r.Endpoint.Extension != nil {
		m.EndpointPrimitiveElement = &primitiveElement{Id: r.Endpoint.Id, Extension: r.Endpoint.Extension}
	}
	m.Receiver = r.Receiver
	return m
}
func (r *MessageHeaderDestination) UnmarshalJSON(b []byte) error {
	var m jsonMessageHeaderDestination
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MessageHeaderDestination) unmarshalJSON(m jsonMessageHeaderDestination) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Target = m.Target
	r.Endpoint = m.Endpoint
	if m.EndpointPrimitiveElement != nil {
		r.Endpoint.Id = m.EndpointPrimitiveElement.Id
		r.Endpoint.Extension = m.EndpointPrimitiveElement.Extension
	}
	r.Receiver = m.Receiver
	return nil
}
func (r MessageHeaderDestination) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The source application from which this message originated.
type MessageHeaderSource struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Human-readable name for the source system.
	Name *String
	// May include configuration or other information useful in debugging.
	Software *String
	// Can convey versions of multiple systems in situations where a message passes through multiple hands.
	Version *String
	// An e-mail, phone, website or other contact point to use to resolve issues with message communications.
	Contact *ContactPoint
	// Identifies the routing target to send acknowledgements to.
	Endpoint Url
}
type jsonMessageHeaderSource struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Name                     *String           `json:"name,omitempty"`
	NamePrimitiveElement     *primitiveElement `json:"_name,omitempty"`
	Software                 *String           `json:"software,omitempty"`
	SoftwarePrimitiveElement *primitiveElement `json:"_software,omitempty"`
	Version                  *String           `json:"version,omitempty"`
	VersionPrimitiveElement  *primitiveElement `json:"_version,omitempty"`
	Contact                  *ContactPoint     `json:"contact,omitempty"`
	Endpoint                 Url               `json:"endpoint,omitempty"`
	EndpointPrimitiveElement *primitiveElement `json:"_endpoint,omitempty"`
}

func (r MessageHeaderSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MessageHeaderSource) marshalJSON() jsonMessageHeaderSource {
	m := jsonMessageHeaderSource{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Software = r.Software
	if r.Software != nil && (r.Software.Id != nil || r.Software.Extension != nil) {
		m.SoftwarePrimitiveElement = &primitiveElement{Id: r.Software.Id, Extension: r.Software.Extension}
	}
	m.Version = r.Version
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	m.Contact = r.Contact
	m.Endpoint = r.Endpoint
	if r.Endpoint.Id != nil || r.Endpoint.Extension != nil {
		m.EndpointPrimitiveElement = &primitiveElement{Id: r.Endpoint.Id, Extension: r.Endpoint.Extension}
	}
	return m
}
func (r *MessageHeaderSource) UnmarshalJSON(b []byte) error {
	var m jsonMessageHeaderSource
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MessageHeaderSource) unmarshalJSON(m jsonMessageHeaderSource) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Software = m.Software
	if m.SoftwarePrimitiveElement != nil {
		r.Software.Id = m.SoftwarePrimitiveElement.Id
		r.Software.Extension = m.SoftwarePrimitiveElement.Extension
	}
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	r.Contact = m.Contact
	r.Endpoint = m.Endpoint
	if m.EndpointPrimitiveElement != nil {
		r.Endpoint.Id = m.EndpointPrimitiveElement.Id
		r.Endpoint.Extension = m.EndpointPrimitiveElement.Extension
	}
	return nil
}
func (r MessageHeaderSource) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about the message that this message is a response to.  Only present if this message is a response.
type MessageHeaderResponse struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The MessageHeader.id of the message to which this message is a response.
	Identifier Id
	// Code that identifies the type of response to the message - whether it was successful or not, and whether it should be resent or not.
	Code Code
	// Full details of any issues found in the message.
	Details *Reference
}
type jsonMessageHeaderResponse struct {
	Id                         *string           `json:"id,omitempty"`
	Extension                  []Extension       `json:"extension,omitempty"`
	ModifierExtension          []Extension       `json:"modifierExtension,omitempty"`
	Identifier                 Id                `json:"identifier,omitempty"`
	IdentifierPrimitiveElement *primitiveElement `json:"_identifier,omitempty"`
	Code                       Code              `json:"code,omitempty"`
	CodePrimitiveElement       *primitiveElement `json:"_code,omitempty"`
	Details                    *Reference        `json:"details,omitempty"`
}

func (r MessageHeaderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MessageHeaderResponse) marshalJSON() jsonMessageHeaderResponse {
	m := jsonMessageHeaderResponse{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	if r.Identifier.Id != nil || r.Identifier.Extension != nil {
		m.IdentifierPrimitiveElement = &primitiveElement{Id: r.Identifier.Id, Extension: r.Identifier.Extension}
	}
	m.Code = r.Code
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Details = r.Details
	return m
}
func (r *MessageHeaderResponse) UnmarshalJSON(b []byte) error {
	var m jsonMessageHeaderResponse
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MessageHeaderResponse) unmarshalJSON(m jsonMessageHeaderResponse) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	if m.IdentifierPrimitiveElement != nil {
		r.Identifier.Id = m.IdentifierPrimitiveElement.Id
		r.Identifier.Extension = m.IdentifierPrimitiveElement.Extension
	}
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Details = m.Details
	return nil
}
func (r MessageHeaderResponse) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
