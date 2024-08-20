package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// The subscription resource is used to define a push-based subscription from a server to another system. Once a subscription is registered with the server, the server checks every resource that is created or updated, and if the resource matches the given criteria, it sends a message on the defined "channel" so that another system can take an appropriate action.
type Subscription struct {
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
	// The status of the subscription, which marks the server state for managing the subscription.
	Status Code
	// Contact details for a human to contact about the subscription. The primary use of this for system administrator troubleshooting.
	Contact []ContactPoint
	// The time for the server to turn the subscription off.
	End *Instant
	// A description of why this subscription is defined.
	Reason String
	// The rules that the server should use to determine when to generate notifications for this subscription.
	Criteria String
	// A record of the last error that occurred when the server processed a notification.
	Error *String
	// Details where to send notifications when resources are received that meet the criteria.
	Channel SubscriptionChannel
}

func (r Subscription) ResourceType() string {
	return "Subscription"
}
func (r Subscription) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonSubscription struct {
	ResourceType                  string              `json:"resourceType"`
	Id                            *Id                 `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement   `json:"_id,omitempty"`
	Meta                          *Meta               `json:"meta,omitempty"`
	ImplicitRules                 *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                      *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement   `json:"_language,omitempty"`
	Text                          *Narrative          `json:"text,omitempty"`
	Contained                     []ContainedResource `json:"contained,omitempty"`
	Extension                     []Extension         `json:"extension,omitempty"`
	ModifierExtension             []Extension         `json:"modifierExtension,omitempty"`
	Status                        Code                `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement   `json:"_status,omitempty"`
	Contact                       []ContactPoint      `json:"contact,omitempty"`
	End                           *Instant            `json:"end,omitempty"`
	EndPrimitiveElement           *primitiveElement   `json:"_end,omitempty"`
	Reason                        String              `json:"reason,omitempty"`
	ReasonPrimitiveElement        *primitiveElement   `json:"_reason,omitempty"`
	Criteria                      String              `json:"criteria,omitempty"`
	CriteriaPrimitiveElement      *primitiveElement   `json:"_criteria,omitempty"`
	Error                         *String             `json:"error,omitempty"`
	ErrorPrimitiveElement         *primitiveElement   `json:"_error,omitempty"`
	Channel                       SubscriptionChannel `json:"channel,omitempty"`
}

func (r Subscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Subscription) marshalJSON() jsonSubscription {
	m := jsonSubscription{}
	m.ResourceType = "Subscription"
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
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Contact = r.Contact
	m.End = r.End
	if r.End != nil && (r.End.Id != nil || r.End.Extension != nil) {
		m.EndPrimitiveElement = &primitiveElement{Id: r.End.Id, Extension: r.End.Extension}
	}
	m.Reason = r.Reason
	if r.Reason.Id != nil || r.Reason.Extension != nil {
		m.ReasonPrimitiveElement = &primitiveElement{Id: r.Reason.Id, Extension: r.Reason.Extension}
	}
	m.Criteria = r.Criteria
	if r.Criteria.Id != nil || r.Criteria.Extension != nil {
		m.CriteriaPrimitiveElement = &primitiveElement{Id: r.Criteria.Id, Extension: r.Criteria.Extension}
	}
	m.Error = r.Error
	if r.Error != nil && (r.Error.Id != nil || r.Error.Extension != nil) {
		m.ErrorPrimitiveElement = &primitiveElement{Id: r.Error.Id, Extension: r.Error.Extension}
	}
	m.Channel = r.Channel
	return m
}
func (r *Subscription) UnmarshalJSON(b []byte) error {
	var m jsonSubscription
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Subscription) unmarshalJSON(m jsonSubscription) error {
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
		r.Contained = append(r.Contained, v.Resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Contact = m.Contact
	r.End = m.End
	if m.EndPrimitiveElement != nil {
		r.End.Id = m.EndPrimitiveElement.Id
		r.End.Extension = m.EndPrimitiveElement.Extension
	}
	r.Reason = m.Reason
	if m.ReasonPrimitiveElement != nil {
		r.Reason.Id = m.ReasonPrimitiveElement.Id
		r.Reason.Extension = m.ReasonPrimitiveElement.Extension
	}
	r.Criteria = m.Criteria
	if m.CriteriaPrimitiveElement != nil {
		r.Criteria.Id = m.CriteriaPrimitiveElement.Id
		r.Criteria.Extension = m.CriteriaPrimitiveElement.Extension
	}
	r.Error = m.Error
	if m.ErrorPrimitiveElement != nil {
		r.Error.Id = m.ErrorPrimitiveElement.Id
		r.Error.Extension = m.ErrorPrimitiveElement.Extension
	}
	r.Channel = m.Channel
	return nil
}
func (r Subscription) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Details where to send notifications when resources are received that meet the criteria.
type SubscriptionChannel struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of channel to send notifications on.
	Type Code
	// The url that describes the actual end-point to send messages to.
	Endpoint *Url
	// The mime type to send the payload in - either application/fhir+xml, or application/fhir+json. If the payload is not present, then there is no payload in the notification, just a notification. The mime type "text/plain" may also be used for Email and SMS subscriptions.
	Payload *Code
	// Additional headers / information to send as part of the notification.
	Header []String
}
type jsonSubscriptionChannel struct {
	Id                       *string             `json:"id,omitempty"`
	Extension                []Extension         `json:"extension,omitempty"`
	ModifierExtension        []Extension         `json:"modifierExtension,omitempty"`
	Type                     Code                `json:"type,omitempty"`
	TypePrimitiveElement     *primitiveElement   `json:"_type,omitempty"`
	Endpoint                 *Url                `json:"endpoint,omitempty"`
	EndpointPrimitiveElement *primitiveElement   `json:"_endpoint,omitempty"`
	Payload                  *Code               `json:"payload,omitempty"`
	PayloadPrimitiveElement  *primitiveElement   `json:"_payload,omitempty"`
	Header                   []String            `json:"header,omitempty"`
	HeaderPrimitiveElement   []*primitiveElement `json:"_header,omitempty"`
}

func (r SubscriptionChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubscriptionChannel) marshalJSON() jsonSubscriptionChannel {
	m := jsonSubscriptionChannel{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Endpoint = r.Endpoint
	if r.Endpoint != nil && (r.Endpoint.Id != nil || r.Endpoint.Extension != nil) {
		m.EndpointPrimitiveElement = &primitiveElement{Id: r.Endpoint.Id, Extension: r.Endpoint.Extension}
	}
	m.Payload = r.Payload
	if r.Payload != nil && (r.Payload.Id != nil || r.Payload.Extension != nil) {
		m.PayloadPrimitiveElement = &primitiveElement{Id: r.Payload.Id, Extension: r.Payload.Extension}
	}
	m.Header = r.Header
	anyHeaderIdOrExtension := false
	for _, e := range r.Header {
		if e.Id != nil || e.Extension != nil {
			anyHeaderIdOrExtension = true
			break
		}
	}
	if anyHeaderIdOrExtension {
		m.HeaderPrimitiveElement = make([]*primitiveElement, 0, len(r.Header))
		for _, e := range r.Header {
			if e.Id != nil || e.Extension != nil {
				m.HeaderPrimitiveElement = append(m.HeaderPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.HeaderPrimitiveElement = append(m.HeaderPrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *SubscriptionChannel) UnmarshalJSON(b []byte) error {
	var m jsonSubscriptionChannel
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubscriptionChannel) unmarshalJSON(m jsonSubscriptionChannel) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Endpoint = m.Endpoint
	if m.EndpointPrimitiveElement != nil {
		r.Endpoint.Id = m.EndpointPrimitiveElement.Id
		r.Endpoint.Extension = m.EndpointPrimitiveElement.Extension
	}
	r.Payload = m.Payload
	if m.PayloadPrimitiveElement != nil {
		r.Payload.Id = m.PayloadPrimitiveElement.Id
		r.Payload.Extension = m.PayloadPrimitiveElement.Extension
	}
	r.Header = m.Header
	for i, e := range m.HeaderPrimitiveElement {
		if len(r.Header) > i {
			r.Header[i].Id = e.Id
			r.Header[i].Extension = e.Extension
		} else {
			r.Header = append(r.Header, String{Id: e.Id, Extension: e.Extension})
		}
	}
	return nil
}
func (r SubscriptionChannel) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
