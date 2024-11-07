package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// An occurrence of information being transmitted; e.g. an alert that was sent to a responsible provider, a public health agency that was notified about a reportable condition.
type Communication struct {
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
	// Business identifiers assigned to this communication by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Communication.
	InstantiatesCanonical []Canonical
	// The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this Communication.
	InstantiatesUri []Uri
	// An order, proposal or plan fulfilled in whole or in part by this Communication.
	BasedOn []Reference
	// Part of this action.
	PartOf []Reference
	// Prior communication that this communication is in response to.
	InResponseTo []Reference
	// The status of the transmission.
	Status Code
	// Captures the reason for the current state of the Communication.
	StatusReason *CodeableConcept
	// The type of message conveyed such as alert, notification, reminder, instruction, etc.
	Category []CodeableConcept
	// Characterizes how quickly the planned or in progress communication must be addressed. Includes concepts such as stat, urgent, routine.
	Priority *Code
	// A channel that was used for this communication (e.g. email, fax).
	Medium []CodeableConcept
	// The patient or group that was the focus of this communication.
	Subject *Reference
	// Description of the purpose/content, similar to a subject line in an email.
	Topic *CodeableConcept
	// Other resources that pertain to this communication and to which this communication should be associated.
	About []Reference
	// The Encounter during which this Communication was created or to which the creation of this record is tightly associated.
	Encounter *Reference
	// The time when this communication was sent.
	Sent *DateTime
	// The time when this communication arrived at the destination.
	Received *DateTime
	// The entity (e.g. person, organization, clinical information system, care team or device) which was the target of the communication. If receipts need to be tracked by an individual, a separate resource instance will need to be created for each recipient.  Multiple recipient communications are intended where either receipts are not tracked (e.g. a mass mail-out) or a receipt is captured in aggregate (all emails confirmed received by a particular time).
	Recipient []Reference
	// The entity (e.g. person, organization, clinical information system, or device) which was the source of the communication.
	Sender *Reference
	// The reason or justification for the communication.
	ReasonCode []CodeableConcept
	// Indicates another resource whose existence justifies this communication.
	ReasonReference []Reference
	// Text, attachment(s), or resource(s) that was communicated to the recipient.
	Payload []CommunicationPayload
	// Additional notes or commentary about the communication by the sender, receiver or other interested parties.
	Note []Annotation
}

func (r Communication) ResourceType() string {
	return "Communication"
}
func (r Communication) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonCommunication struct {
	ResourceType                          string                 `json:"resourceType"`
	Id                                    *Id                    `json:"id,omitempty"`
	IdPrimitiveElement                    *primitiveElement      `json:"_id,omitempty"`
	Meta                                  *Meta                  `json:"meta,omitempty"`
	ImplicitRules                         *Uri                   `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement         *primitiveElement      `json:"_implicitRules,omitempty"`
	Language                              *Code                  `json:"language,omitempty"`
	LanguagePrimitiveElement              *primitiveElement      `json:"_language,omitempty"`
	Text                                  *Narrative             `json:"text,omitempty"`
	Contained                             []ContainedResource    `json:"contained,omitempty"`
	Extension                             []Extension            `json:"extension,omitempty"`
	ModifierExtension                     []Extension            `json:"modifierExtension,omitempty"`
	Identifier                            []Identifier           `json:"identifier,omitempty"`
	InstantiatesCanonical                 []Canonical            `json:"instantiatesCanonical,omitempty"`
	InstantiatesCanonicalPrimitiveElement []*primitiveElement    `json:"_instantiatesCanonical,omitempty"`
	InstantiatesUri                       []Uri                  `json:"instantiatesUri,omitempty"`
	InstantiatesUriPrimitiveElement       []*primitiveElement    `json:"_instantiatesUri,omitempty"`
	BasedOn                               []Reference            `json:"basedOn,omitempty"`
	PartOf                                []Reference            `json:"partOf,omitempty"`
	InResponseTo                          []Reference            `json:"inResponseTo,omitempty"`
	Status                                Code                   `json:"status,omitempty"`
	StatusPrimitiveElement                *primitiveElement      `json:"_status,omitempty"`
	StatusReason                          *CodeableConcept       `json:"statusReason,omitempty"`
	Category                              []CodeableConcept      `json:"category,omitempty"`
	Priority                              *Code                  `json:"priority,omitempty"`
	PriorityPrimitiveElement              *primitiveElement      `json:"_priority,omitempty"`
	Medium                                []CodeableConcept      `json:"medium,omitempty"`
	Subject                               *Reference             `json:"subject,omitempty"`
	Topic                                 *CodeableConcept       `json:"topic,omitempty"`
	About                                 []Reference            `json:"about,omitempty"`
	Encounter                             *Reference             `json:"encounter,omitempty"`
	Sent                                  *DateTime              `json:"sent,omitempty"`
	SentPrimitiveElement                  *primitiveElement      `json:"_sent,omitempty"`
	Received                              *DateTime              `json:"received,omitempty"`
	ReceivedPrimitiveElement              *primitiveElement      `json:"_received,omitempty"`
	Recipient                             []Reference            `json:"recipient,omitempty"`
	Sender                                *Reference             `json:"sender,omitempty"`
	ReasonCode                            []CodeableConcept      `json:"reasonCode,omitempty"`
	ReasonReference                       []Reference            `json:"reasonReference,omitempty"`
	Payload                               []CommunicationPayload `json:"payload,omitempty"`
	Note                                  []Annotation           `json:"note,omitempty"`
}

func (r Communication) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Communication) marshalJSON() jsonCommunication {
	m := jsonCommunication{}
	m.ResourceType = "Communication"
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
	anyInstantiatesCanonicalValue := false
	for _, e := range r.InstantiatesCanonical {
		if e.Value != nil {
			anyInstantiatesCanonicalValue = true
			break
		}
	}
	if anyInstantiatesCanonicalValue {
		m.InstantiatesCanonical = r.InstantiatesCanonical
	}
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
	anyInstantiatesUriValue := false
	for _, e := range r.InstantiatesUri {
		if e.Value != nil {
			anyInstantiatesUriValue = true
			break
		}
	}
	if anyInstantiatesUriValue {
		m.InstantiatesUri = r.InstantiatesUri
	}
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
	m.PartOf = r.PartOf
	m.InResponseTo = r.InResponseTo
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.StatusReason = r.StatusReason
	m.Category = r.Category
	if r.Priority != nil && r.Priority.Value != nil {
		m.Priority = r.Priority
	}
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		m.PriorityPrimitiveElement = &primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
	}
	m.Medium = r.Medium
	m.Subject = r.Subject
	m.Topic = r.Topic
	m.About = r.About
	m.Encounter = r.Encounter
	if r.Sent != nil && r.Sent.Value != nil {
		m.Sent = r.Sent
	}
	if r.Sent != nil && (r.Sent.Id != nil || r.Sent.Extension != nil) {
		m.SentPrimitiveElement = &primitiveElement{Id: r.Sent.Id, Extension: r.Sent.Extension}
	}
	if r.Received != nil && r.Received.Value != nil {
		m.Received = r.Received
	}
	if r.Received != nil && (r.Received.Id != nil || r.Received.Extension != nil) {
		m.ReceivedPrimitiveElement = &primitiveElement{Id: r.Received.Id, Extension: r.Received.Extension}
	}
	m.Recipient = r.Recipient
	m.Sender = r.Sender
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Payload = r.Payload
	m.Note = r.Note
	return m
}
func (r *Communication) UnmarshalJSON(b []byte) error {
	var m jsonCommunication
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Communication) unmarshalJSON(m jsonCommunication) error {
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
	r.InstantiatesCanonical = m.InstantiatesCanonical
	for i, e := range m.InstantiatesCanonicalPrimitiveElement {
		if len(r.InstantiatesCanonical) <= i {
			r.InstantiatesCanonical = append(r.InstantiatesCanonical, Canonical{})
		}
		if e != nil {
			r.InstantiatesCanonical[i].Id = e.Id
			r.InstantiatesCanonical[i].Extension = e.Extension
		}
	}
	r.InstantiatesUri = m.InstantiatesUri
	for i, e := range m.InstantiatesUriPrimitiveElement {
		if len(r.InstantiatesUri) <= i {
			r.InstantiatesUri = append(r.InstantiatesUri, Uri{})
		}
		if e != nil {
			r.InstantiatesUri[i].Id = e.Id
			r.InstantiatesUri[i].Extension = e.Extension
		}
	}
	r.BasedOn = m.BasedOn
	r.PartOf = m.PartOf
	r.InResponseTo = m.InResponseTo
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.StatusReason = m.StatusReason
	r.Category = m.Category
	r.Priority = m.Priority
	if m.PriorityPrimitiveElement != nil {
		if r.Priority == nil {
			r.Priority = &Code{}
		}
		r.Priority.Id = m.PriorityPrimitiveElement.Id
		r.Priority.Extension = m.PriorityPrimitiveElement.Extension
	}
	r.Medium = m.Medium
	r.Subject = m.Subject
	r.Topic = m.Topic
	r.About = m.About
	r.Encounter = m.Encounter
	r.Sent = m.Sent
	if m.SentPrimitiveElement != nil {
		if r.Sent == nil {
			r.Sent = &DateTime{}
		}
		r.Sent.Id = m.SentPrimitiveElement.Id
		r.Sent.Extension = m.SentPrimitiveElement.Extension
	}
	r.Received = m.Received
	if m.ReceivedPrimitiveElement != nil {
		if r.Received == nil {
			r.Received = &DateTime{}
		}
		r.Received.Id = m.ReceivedPrimitiveElement.Id
		r.Received.Extension = m.ReceivedPrimitiveElement.Extension
	}
	r.Recipient = m.Recipient
	r.Sender = m.Sender
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Payload = m.Payload
	r.Note = m.Note
	return nil
}
func (r Communication) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Communication"
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
	err = e.EncodeElement(r.InstantiatesCanonical, xml.StartElement{Name: xml.Name{Local: "instantiatesCanonical"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.InstantiatesUri, xml.StartElement{Name: xml.Name{Local: "instantiatesUri"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BasedOn, xml.StartElement{Name: xml.Name{Local: "basedOn"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PartOf, xml.StartElement{Name: xml.Name{Local: "partOf"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.InResponseTo, xml.StartElement{Name: xml.Name{Local: "inResponseTo"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.StatusReason, xml.StartElement{Name: xml.Name{Local: "statusReason"}})
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
	err = e.EncodeElement(r.Medium, xml.StartElement{Name: xml.Name{Local: "medium"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Subject, xml.StartElement{Name: xml.Name{Local: "subject"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Topic, xml.StartElement{Name: xml.Name{Local: "topic"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.About, xml.StartElement{Name: xml.Name{Local: "about"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Encounter, xml.StartElement{Name: xml.Name{Local: "encounter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Sent, xml.StartElement{Name: xml.Name{Local: "sent"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Received, xml.StartElement{Name: xml.Name{Local: "received"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Recipient, xml.StartElement{Name: xml.Name{Local: "recipient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Sender, xml.StartElement{Name: xml.Name{Local: "sender"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReasonCode, xml.StartElement{Name: xml.Name{Local: "reasonCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReasonReference, xml.StartElement{Name: xml.Name{Local: "reasonReference"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Payload, xml.StartElement{Name: xml.Name{Local: "payload"}})
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
func (r *Communication) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "instantiatesCanonical":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.InstantiatesCanonical = append(r.InstantiatesCanonical, v)
			case "instantiatesUri":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.InstantiatesUri = append(r.InstantiatesUri, v)
			case "basedOn":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BasedOn = append(r.BasedOn, v)
			case "partOf":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PartOf = append(r.PartOf, v)
			case "inResponseTo":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.InResponseTo = append(r.InResponseTo, v)
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "statusReason":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StatusReason = &v
			case "category":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = append(r.Category, v)
			case "priority":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Priority = &v
			case "medium":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Medium = append(r.Medium, v)
			case "subject":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = &v
			case "topic":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Topic = &v
			case "about":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.About = append(r.About, v)
			case "encounter":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Encounter = &v
			case "sent":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Sent = &v
			case "received":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Received = &v
			case "recipient":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Recipient = append(r.Recipient, v)
			case "sender":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Sender = &v
			case "reasonCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReasonCode = append(r.ReasonCode, v)
			case "reasonReference":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReasonReference = append(r.ReasonReference, v)
			case "payload":
				var v CommunicationPayload
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Payload = append(r.Payload, v)
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
func (r Communication) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Text, attachment(s), or resource(s) that was communicated to the recipient.
type CommunicationPayload struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A communicated content (or for multi-part communications, one portion of the communication).
	Content isCommunicationPayloadContent
}
type isCommunicationPayloadContent interface {
	isCommunicationPayloadContent()
}

func (r String) isCommunicationPayloadContent()     {}
func (r Attachment) isCommunicationPayloadContent() {}
func (r Reference) isCommunicationPayloadContent()  {}

type jsonCommunicationPayload struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	ContentString                 *String           `json:"contentString,omitempty"`
	ContentStringPrimitiveElement *primitiveElement `json:"_contentString,omitempty"`
	ContentAttachment             *Attachment       `json:"contentAttachment,omitempty"`
	ContentReference              *Reference        `json:"contentReference,omitempty"`
}

func (r CommunicationPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CommunicationPayload) marshalJSON() jsonCommunicationPayload {
	m := jsonCommunicationPayload{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Content.(type) {
	case String:
		if v.Value != nil {
			m.ContentString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ContentStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.ContentString = v
		}
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
func (r *CommunicationPayload) UnmarshalJSON(b []byte) error {
	var m jsonCommunicationPayload
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CommunicationPayload) unmarshalJSON(m jsonCommunicationPayload) error {
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
func (r CommunicationPayload) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	switch v := r.Content.(type) {
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contentString"}})
		if err != nil {
			return err
		}
	case Attachment, *Attachment:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contentAttachment"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contentReference"}})
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
func (r *CommunicationPayload) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "contentString":
				if r.Content != nil {
					return fmt.Errorf("multiple values for field \"Content\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Content = v
			case "contentAttachment":
				if r.Content != nil {
					return fmt.Errorf("multiple values for field \"Content\"")
				}
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Content = v
			case "contentReference":
				if r.Content != nil {
					return fmt.Errorf("multiple values for field \"Content\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Content = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CommunicationPayload) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
