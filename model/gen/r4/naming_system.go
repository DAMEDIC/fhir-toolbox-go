package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A curated namespace that issues unique symbols within that namespace for the identification of concepts, people, devices, etc.  Represents a "System" used within the Identifier and Coding data types.
type NamingSystem struct {
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
	// A natural language name identifying the naming system. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name String
	// The status of this naming system. Enables tracking the life-cycle of the content.
	Status Code
	// Indicates the purpose for the naming system - what kinds of things does it make unique?
	Kind Code
	// The date  (and optionally time) when the naming system was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the naming system changes.
	Date DateTime
	// The name of the organization or individual that published the naming system.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// The name of the organization that is responsible for issuing identifiers or codes for this namespace and ensuring their non-collision.
	Responsible *String
	// Categorizes a naming system for easier search by grouping related naming systems.
	Type *CodeableConcept
	// A free text natural language description of the naming system from a consumer's perspective. Details about what the namespace identifies including scope, granularity, version labeling, etc.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate naming system instances.
	UseContext []UsageContext
	// A legal or geographic region in which the naming system is intended to be used.
	Jurisdiction []CodeableConcept
	// Provides guidance on the use of the namespace, including the handling of formatting characters, use of upper vs. lower case, etc.
	Usage *String
	// Indicates how the system may be identified when referenced in electronic exchange.
	UniqueId []NamingSystemUniqueId
}

func (r NamingSystem) ResourceType() string {
	return "NamingSystem"
}
func (r NamingSystem) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonNamingSystem struct {
	ResourceType                  string                 `json:"resourceType"`
	Id                            *Id                    `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement      `json:"_id,omitempty"`
	Meta                          *Meta                  `json:"meta,omitempty"`
	ImplicitRules                 *Uri                   `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement      `json:"_implicitRules,omitempty"`
	Language                      *Code                  `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement      `json:"_language,omitempty"`
	Text                          *Narrative             `json:"text,omitempty"`
	Contained                     []containedResource    `json:"contained,omitempty"`
	Extension                     []Extension            `json:"extension,omitempty"`
	ModifierExtension             []Extension            `json:"modifierExtension,omitempty"`
	Name                          String                 `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement      `json:"_name,omitempty"`
	Status                        Code                   `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement      `json:"_status,omitempty"`
	Kind                          Code                   `json:"kind,omitempty"`
	KindPrimitiveElement          *primitiveElement      `json:"_kind,omitempty"`
	Date                          DateTime               `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement      `json:"_date,omitempty"`
	Publisher                     *String                `json:"publisher,omitempty"`
	PublisherPrimitiveElement     *primitiveElement      `json:"_publisher,omitempty"`
	Contact                       []ContactDetail        `json:"contact,omitempty"`
	Responsible                   *String                `json:"responsible,omitempty"`
	ResponsiblePrimitiveElement   *primitiveElement      `json:"_responsible,omitempty"`
	Type                          *CodeableConcept       `json:"type,omitempty"`
	Description                   *Markdown              `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement      `json:"_description,omitempty"`
	UseContext                    []UsageContext         `json:"useContext,omitempty"`
	Jurisdiction                  []CodeableConcept      `json:"jurisdiction,omitempty"`
	Usage                         *String                `json:"usage,omitempty"`
	UsagePrimitiveElement         *primitiveElement      `json:"_usage,omitempty"`
	UniqueId                      []NamingSystemUniqueId `json:"uniqueId,omitempty"`
}

func (r NamingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r NamingSystem) marshalJSON() jsonNamingSystem {
	m := jsonNamingSystem{}
	m.ResourceType = "NamingSystem"
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
	m.Name = r.Name
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Kind = r.Kind
	if r.Kind.Id != nil || r.Kind.Extension != nil {
		m.KindPrimitiveElement = &primitiveElement{Id: r.Kind.Id, Extension: r.Kind.Extension}
	}
	m.Date = r.Date
	if r.Date.Id != nil || r.Date.Extension != nil {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Publisher = r.Publisher
	if r.Publisher != nil && (r.Publisher.Id != nil || r.Publisher.Extension != nil) {
		m.PublisherPrimitiveElement = &primitiveElement{Id: r.Publisher.Id, Extension: r.Publisher.Extension}
	}
	m.Contact = r.Contact
	m.Responsible = r.Responsible
	if r.Responsible != nil && (r.Responsible.Id != nil || r.Responsible.Extension != nil) {
		m.ResponsiblePrimitiveElement = &primitiveElement{Id: r.Responsible.Id, Extension: r.Responsible.Extension}
	}
	m.Type = r.Type
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.UseContext = r.UseContext
	m.Jurisdiction = r.Jurisdiction
	m.Usage = r.Usage
	if r.Usage != nil && (r.Usage.Id != nil || r.Usage.Extension != nil) {
		m.UsagePrimitiveElement = &primitiveElement{Id: r.Usage.Id, Extension: r.Usage.Extension}
	}
	m.UniqueId = r.UniqueId
	return m
}
func (r *NamingSystem) UnmarshalJSON(b []byte) error {
	var m jsonNamingSystem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *NamingSystem) unmarshalJSON(m jsonNamingSystem) error {
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
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Kind = m.Kind
	if m.KindPrimitiveElement != nil {
		r.Kind.Id = m.KindPrimitiveElement.Id
		r.Kind.Extension = m.KindPrimitiveElement.Extension
	}
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Publisher = m.Publisher
	if m.PublisherPrimitiveElement != nil {
		r.Publisher.Id = m.PublisherPrimitiveElement.Id
		r.Publisher.Extension = m.PublisherPrimitiveElement.Extension
	}
	r.Contact = m.Contact
	r.Responsible = m.Responsible
	if m.ResponsiblePrimitiveElement != nil {
		r.Responsible.Id = m.ResponsiblePrimitiveElement.Id
		r.Responsible.Extension = m.ResponsiblePrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.UseContext = m.UseContext
	r.Jurisdiction = m.Jurisdiction
	r.Usage = m.Usage
	if m.UsagePrimitiveElement != nil {
		r.Usage.Id = m.UsagePrimitiveElement.Id
		r.Usage.Extension = m.UsagePrimitiveElement.Extension
	}
	r.UniqueId = m.UniqueId
	return nil
}
func (r NamingSystem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates how the system may be identified when referenced in electronic exchange.
type NamingSystemUniqueId struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identifies the unique identifier scheme used for this particular identifier.
	Type Code
	// The string that should be sent over the wire to identify the code system or identifier system.
	Value String
	// Indicates whether this identifier is the "preferred" identifier of this type.
	Preferred *Boolean
	// Notes about the past or intended usage of this identifier.
	Comment *String
	// Identifies the period of time over which this identifier is considered appropriate to refer to the naming system.  Outside of this window, the identifier might be non-deterministic.
	Period *Period
}
type jsonNamingSystemUniqueId struct {
	Id                        *string           `json:"id,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Type                      Code              `json:"type,omitempty"`
	TypePrimitiveElement      *primitiveElement `json:"_type,omitempty"`
	Value                     String            `json:"value,omitempty"`
	ValuePrimitiveElement     *primitiveElement `json:"_value,omitempty"`
	Preferred                 *Boolean          `json:"preferred,omitempty"`
	PreferredPrimitiveElement *primitiveElement `json:"_preferred,omitempty"`
	Comment                   *String           `json:"comment,omitempty"`
	CommentPrimitiveElement   *primitiveElement `json:"_comment,omitempty"`
	Period                    *Period           `json:"period,omitempty"`
}

func (r NamingSystemUniqueId) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r NamingSystemUniqueId) marshalJSON() jsonNamingSystemUniqueId {
	m := jsonNamingSystemUniqueId{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Value = r.Value
	if r.Value.Id != nil || r.Value.Extension != nil {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	m.Preferred = r.Preferred
	if r.Preferred != nil && (r.Preferred.Id != nil || r.Preferred.Extension != nil) {
		m.PreferredPrimitiveElement = &primitiveElement{Id: r.Preferred.Id, Extension: r.Preferred.Extension}
	}
	m.Comment = r.Comment
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	m.Period = r.Period
	return m
}
func (r *NamingSystemUniqueId) UnmarshalJSON(b []byte) error {
	var m jsonNamingSystemUniqueId
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *NamingSystemUniqueId) unmarshalJSON(m jsonNamingSystemUniqueId) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	r.Preferred = m.Preferred
	if m.PreferredPrimitiveElement != nil {
		r.Preferred.Id = m.PreferredPrimitiveElement.Id
		r.Preferred.Extension = m.PreferredPrimitiveElement.Extension
	}
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	r.Period = m.Period
	return nil
}
func (r NamingSystemUniqueId) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
