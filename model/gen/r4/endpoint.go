package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// The technical details of an endpoint that can be used for electronic services, such as for web services providing XDS.b or a REST endpoint for another FHIR server. This may include any security context information.
type Endpoint struct {
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
	// Identifier for the organization that is used to identify the endpoint across multiple disparate systems.
	Identifier []Identifier
	// active | suspended | error | off | test.
	Status Code
	// A coded value that represents the technical details of the usage of this endpoint, such as what WSDLs should be used in what way. (e.g. XDS.b/DICOM/cds-hook).
	ConnectionType Coding
	// A friendly name that this endpoint can be referred to with.
	Name *String
	// The organization that manages this endpoint (even if technically another organization is hosting this in the cloud, it is the organization associated with the data).
	ManagingOrganization *Reference
	// Contact details for a human to contact about the subscription. The primary use of this for system administrator troubleshooting.
	Contact []ContactPoint
	// The interval during which the endpoint is expected to be operational.
	Period *Period
	// The payload type describes the acceptable content that can be communicated on the endpoint.
	PayloadType []CodeableConcept
	// The mime type to send the payload in - e.g. application/fhir+xml, application/fhir+json. If the mime type is not specified, then the sender could send any content (including no content depending on the connectionType).
	PayloadMimeType []Code
	// The uri that describes the actual end-point to connect to.
	Address Url
	// Additional headers / information to send as part of the notification.
	Header []String
}

func (r Endpoint) ResourceType() string {
	return "Endpoint"
}
func (r Endpoint) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonEndpoint struct {
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
	Status                          Code                `json:"status,omitempty"`
	StatusPrimitiveElement          *primitiveElement   `json:"_status,omitempty"`
	ConnectionType                  Coding              `json:"connectionType,omitempty"`
	Name                            *String             `json:"name,omitempty"`
	NamePrimitiveElement            *primitiveElement   `json:"_name,omitempty"`
	ManagingOrganization            *Reference          `json:"managingOrganization,omitempty"`
	Contact                         []ContactPoint      `json:"contact,omitempty"`
	Period                          *Period             `json:"period,omitempty"`
	PayloadType                     []CodeableConcept   `json:"payloadType,omitempty"`
	PayloadMimeType                 []Code              `json:"payloadMimeType,omitempty"`
	PayloadMimeTypePrimitiveElement []*primitiveElement `json:"_payloadMimeType,omitempty"`
	Address                         Url                 `json:"address,omitempty"`
	AddressPrimitiveElement         *primitiveElement   `json:"_address,omitempty"`
	Header                          []String            `json:"header,omitempty"`
	HeaderPrimitiveElement          []*primitiveElement `json:"_header,omitempty"`
}

func (r Endpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Endpoint) marshalJSON() jsonEndpoint {
	m := jsonEndpoint{}
	m.ResourceType = "Endpoint"
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
	m.Identifier = r.Identifier
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.ConnectionType = r.ConnectionType
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.ManagingOrganization = r.ManagingOrganization
	m.Contact = r.Contact
	m.Period = r.Period
	m.PayloadType = r.PayloadType
	m.PayloadMimeType = r.PayloadMimeType
	anyPayloadMimeTypeIdOrExtension := false
	for _, e := range r.PayloadMimeType {
		if e.Id != nil || e.Extension != nil {
			anyPayloadMimeTypeIdOrExtension = true
			break
		}
	}
	if anyPayloadMimeTypeIdOrExtension {
		m.PayloadMimeTypePrimitiveElement = make([]*primitiveElement, 0, len(r.PayloadMimeType))
		for _, e := range r.PayloadMimeType {
			if e.Id != nil || e.Extension != nil {
				m.PayloadMimeTypePrimitiveElement = append(m.PayloadMimeTypePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.PayloadMimeTypePrimitiveElement = append(m.PayloadMimeTypePrimitiveElement, nil)
			}
		}
	}
	m.Address = r.Address
	if r.Address.Id != nil || r.Address.Extension != nil {
		m.AddressPrimitiveElement = &primitiveElement{Id: r.Address.Id, Extension: r.Address.Extension}
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
func (r *Endpoint) UnmarshalJSON(b []byte) error {
	var m jsonEndpoint
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Endpoint) unmarshalJSON(m jsonEndpoint) error {
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
	r.Identifier = m.Identifier
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.ConnectionType = m.ConnectionType
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.ManagingOrganization = m.ManagingOrganization
	r.Contact = m.Contact
	r.Period = m.Period
	r.PayloadType = m.PayloadType
	r.PayloadMimeType = m.PayloadMimeType
	for i, e := range m.PayloadMimeTypePrimitiveElement {
		if len(r.PayloadMimeType) > i {
			r.PayloadMimeType[i].Id = e.Id
			r.PayloadMimeType[i].Extension = e.Extension
		} else {
			r.PayloadMimeType = append(r.PayloadMimeType, Code{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Address = m.Address
	if m.AddressPrimitiveElement != nil {
		r.Address.Id = m.AddressPrimitiveElement.Id
		r.Address.Extension = m.AddressPrimitiveElement.Extension
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
func (r Endpoint) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
