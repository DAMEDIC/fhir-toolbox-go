package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// Defines the characteristics of a message that can be shared between systems, including the type of event that initiates the message, the content to be transmitted and what response(s), if any, are permitted.
//
// Allows messages to be defined once and re-used across systems.
type MessageDefinition struct {
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
	// The business identifier that is used to reference the MessageDefinition and *is* expected to be consistent from server to server.
	Url *Uri
	// A formal identifier that is used to identify this message definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the message definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the message definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the message definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the message definition.
	Title *String
	// A MessageDefinition that is superseded by this definition.
	Replaces []Canonical
	// The status of this message definition. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this message definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the message definition was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the message definition changes.
	Date DateTime
	// The name of the organization or individual that published the message definition.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the message definition from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate message definition instances.
	UseContext []UsageContext
	// A legal or geographic region in which the message definition is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this message definition is needed and why it has been designed as it has.
	Purpose *Markdown
	// A copyright statement relating to the message definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the message definition.
	Copyright *Markdown
	// The MessageDefinition that is the basis for the contents of this resource.
	Base *Canonical
	// Identifies a protocol or workflow that this MessageDefinition represents a step in.
	Parent []Canonical
	// Event code or link to the EventDefinition.
	Event isMessageDefinitionEvent
	// The impact of the content of the message.
	Category *Code
	// Identifies the resource (or resources) that are being addressed by the event.  For example, the Encounter for an admit message or two Account records for a merge.
	Focus []MessageDefinitionFocus
	// Declare at a message definition level whether a response is required or only upon error or success, or never.
	ResponseRequired *Code
	// Indicates what types of messages may be sent as an application-level response to this message.
	AllowedResponse []MessageDefinitionAllowedResponse
	// Canonical reference to a GraphDefinition. If a URL is provided, it is the canonical reference to a [GraphDefinition](graphdefinition.html) that it controls what resources are to be added to the bundle when building the document. The GraphDefinition can also specify profiles that apply to the various resources.
	Graph []Canonical
}

func (r MessageDefinition) ResourceType() string {
	return "MessageDefinition"
}
func (r MessageDefinition) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isMessageDefinitionEvent interface {
	isMessageDefinitionEvent()
}

func (r Coding) isMessageDefinitionEvent() {}
func (r Uri) isMessageDefinitionEvent()    {}

type jsonMessageDefinition struct {
	ResourceType                     string                             `json:"resourceType"`
	Id                               *Id                                `json:"id,omitempty"`
	IdPrimitiveElement               *primitiveElement                  `json:"_id,omitempty"`
	Meta                             *Meta                              `json:"meta,omitempty"`
	ImplicitRules                    *Uri                               `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement    *primitiveElement                  `json:"_implicitRules,omitempty"`
	Language                         *Code                              `json:"language,omitempty"`
	LanguagePrimitiveElement         *primitiveElement                  `json:"_language,omitempty"`
	Text                             *Narrative                         `json:"text,omitempty"`
	Contained                        []ContainedResource                `json:"contained,omitempty"`
	Extension                        []Extension                        `json:"extension,omitempty"`
	ModifierExtension                []Extension                        `json:"modifierExtension,omitempty"`
	Url                              *Uri                               `json:"url,omitempty"`
	UrlPrimitiveElement              *primitiveElement                  `json:"_url,omitempty"`
	Identifier                       []Identifier                       `json:"identifier,omitempty"`
	Version                          *String                            `json:"version,omitempty"`
	VersionPrimitiveElement          *primitiveElement                  `json:"_version,omitempty"`
	Name                             *String                            `json:"name,omitempty"`
	NamePrimitiveElement             *primitiveElement                  `json:"_name,omitempty"`
	Title                            *String                            `json:"title,omitempty"`
	TitlePrimitiveElement            *primitiveElement                  `json:"_title,omitempty"`
	Replaces                         []Canonical                        `json:"replaces,omitempty"`
	ReplacesPrimitiveElement         []*primitiveElement                `json:"_replaces,omitempty"`
	Status                           Code                               `json:"status,omitempty"`
	StatusPrimitiveElement           *primitiveElement                  `json:"_status,omitempty"`
	Experimental                     *Boolean                           `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement     *primitiveElement                  `json:"_experimental,omitempty"`
	Date                             DateTime                           `json:"date,omitempty"`
	DatePrimitiveElement             *primitiveElement                  `json:"_date,omitempty"`
	Publisher                        *String                            `json:"publisher,omitempty"`
	PublisherPrimitiveElement        *primitiveElement                  `json:"_publisher,omitempty"`
	Contact                          []ContactDetail                    `json:"contact,omitempty"`
	Description                      *Markdown                          `json:"description,omitempty"`
	DescriptionPrimitiveElement      *primitiveElement                  `json:"_description,omitempty"`
	UseContext                       []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction                     []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Purpose                          *Markdown                          `json:"purpose,omitempty"`
	PurposePrimitiveElement          *primitiveElement                  `json:"_purpose,omitempty"`
	Copyright                        *Markdown                          `json:"copyright,omitempty"`
	CopyrightPrimitiveElement        *primitiveElement                  `json:"_copyright,omitempty"`
	Base                             *Canonical                         `json:"base,omitempty"`
	BasePrimitiveElement             *primitiveElement                  `json:"_base,omitempty"`
	Parent                           []Canonical                        `json:"parent,omitempty"`
	ParentPrimitiveElement           []*primitiveElement                `json:"_parent,omitempty"`
	EventCoding                      *Coding                            `json:"eventCoding,omitempty"`
	EventUri                         *Uri                               `json:"eventUri,omitempty"`
	EventUriPrimitiveElement         *primitiveElement                  `json:"_eventUri,omitempty"`
	Category                         *Code                              `json:"category,omitempty"`
	CategoryPrimitiveElement         *primitiveElement                  `json:"_category,omitempty"`
	Focus                            []MessageDefinitionFocus           `json:"focus,omitempty"`
	ResponseRequired                 *Code                              `json:"responseRequired,omitempty"`
	ResponseRequiredPrimitiveElement *primitiveElement                  `json:"_responseRequired,omitempty"`
	AllowedResponse                  []MessageDefinitionAllowedResponse `json:"allowedResponse,omitempty"`
	Graph                            []Canonical                        `json:"graph,omitempty"`
	GraphPrimitiveElement            []*primitiveElement                `json:"_graph,omitempty"`
}

func (r MessageDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MessageDefinition) marshalJSON() jsonMessageDefinition {
	m := jsonMessageDefinition{}
	m.ResourceType = "MessageDefinition"
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
	m.Url = r.Url
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	m.Identifier = r.Identifier
	m.Version = r.Version
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Title = r.Title
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	m.Replaces = r.Replaces
	anyReplacesIdOrExtension := false
	for _, e := range r.Replaces {
		if e.Id != nil || e.Extension != nil {
			anyReplacesIdOrExtension = true
			break
		}
	}
	if anyReplacesIdOrExtension {
		m.ReplacesPrimitiveElement = make([]*primitiveElement, 0, len(r.Replaces))
		for _, e := range r.Replaces {
			if e.Id != nil || e.Extension != nil {
				m.ReplacesPrimitiveElement = append(m.ReplacesPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ReplacesPrimitiveElement = append(m.ReplacesPrimitiveElement, nil)
			}
		}
	}
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Experimental = r.Experimental
	if r.Experimental != nil && (r.Experimental.Id != nil || r.Experimental.Extension != nil) {
		m.ExperimentalPrimitiveElement = &primitiveElement{Id: r.Experimental.Id, Extension: r.Experimental.Extension}
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
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.UseContext = r.UseContext
	m.Jurisdiction = r.Jurisdiction
	m.Purpose = r.Purpose
	if r.Purpose != nil && (r.Purpose.Id != nil || r.Purpose.Extension != nil) {
		m.PurposePrimitiveElement = &primitiveElement{Id: r.Purpose.Id, Extension: r.Purpose.Extension}
	}
	m.Copyright = r.Copyright
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		m.CopyrightPrimitiveElement = &primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
	}
	m.Base = r.Base
	if r.Base != nil && (r.Base.Id != nil || r.Base.Extension != nil) {
		m.BasePrimitiveElement = &primitiveElement{Id: r.Base.Id, Extension: r.Base.Extension}
	}
	m.Parent = r.Parent
	anyParentIdOrExtension := false
	for _, e := range r.Parent {
		if e.Id != nil || e.Extension != nil {
			anyParentIdOrExtension = true
			break
		}
	}
	if anyParentIdOrExtension {
		m.ParentPrimitiveElement = make([]*primitiveElement, 0, len(r.Parent))
		for _, e := range r.Parent {
			if e.Id != nil || e.Extension != nil {
				m.ParentPrimitiveElement = append(m.ParentPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ParentPrimitiveElement = append(m.ParentPrimitiveElement, nil)
			}
		}
	}
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
	m.Category = r.Category
	if r.Category != nil && (r.Category.Id != nil || r.Category.Extension != nil) {
		m.CategoryPrimitiveElement = &primitiveElement{Id: r.Category.Id, Extension: r.Category.Extension}
	}
	m.Focus = r.Focus
	m.ResponseRequired = r.ResponseRequired
	if r.ResponseRequired != nil && (r.ResponseRequired.Id != nil || r.ResponseRequired.Extension != nil) {
		m.ResponseRequiredPrimitiveElement = &primitiveElement{Id: r.ResponseRequired.Id, Extension: r.ResponseRequired.Extension}
	}
	m.AllowedResponse = r.AllowedResponse
	m.Graph = r.Graph
	anyGraphIdOrExtension := false
	for _, e := range r.Graph {
		if e.Id != nil || e.Extension != nil {
			anyGraphIdOrExtension = true
			break
		}
	}
	if anyGraphIdOrExtension {
		m.GraphPrimitiveElement = make([]*primitiveElement, 0, len(r.Graph))
		for _, e := range r.Graph {
			if e.Id != nil || e.Extension != nil {
				m.GraphPrimitiveElement = append(m.GraphPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.GraphPrimitiveElement = append(m.GraphPrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *MessageDefinition) UnmarshalJSON(b []byte) error {
	var m jsonMessageDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MessageDefinition) unmarshalJSON(m jsonMessageDefinition) error {
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
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Identifier = m.Identifier
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Replaces = m.Replaces
	for i, e := range m.ReplacesPrimitiveElement {
		if len(r.Replaces) > i {
			r.Replaces[i].Id = e.Id
			r.Replaces[i].Extension = e.Extension
		} else {
			r.Replaces = append(r.Replaces, Canonical{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Experimental = m.Experimental
	if m.ExperimentalPrimitiveElement != nil {
		r.Experimental.Id = m.ExperimentalPrimitiveElement.Id
		r.Experimental.Extension = m.ExperimentalPrimitiveElement.Extension
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
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.UseContext = m.UseContext
	r.Jurisdiction = m.Jurisdiction
	r.Purpose = m.Purpose
	if m.PurposePrimitiveElement != nil {
		r.Purpose.Id = m.PurposePrimitiveElement.Id
		r.Purpose.Extension = m.PurposePrimitiveElement.Extension
	}
	r.Copyright = m.Copyright
	if m.CopyrightPrimitiveElement != nil {
		r.Copyright.Id = m.CopyrightPrimitiveElement.Id
		r.Copyright.Extension = m.CopyrightPrimitiveElement.Extension
	}
	r.Base = m.Base
	if m.BasePrimitiveElement != nil {
		r.Base.Id = m.BasePrimitiveElement.Id
		r.Base.Extension = m.BasePrimitiveElement.Extension
	}
	r.Parent = m.Parent
	for i, e := range m.ParentPrimitiveElement {
		if len(r.Parent) > i {
			r.Parent[i].Id = e.Id
			r.Parent[i].Extension = e.Extension
		} else {
			r.Parent = append(r.Parent, Canonical{Id: e.Id, Extension: e.Extension})
		}
	}
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
	r.Category = m.Category
	if m.CategoryPrimitiveElement != nil {
		r.Category.Id = m.CategoryPrimitiveElement.Id
		r.Category.Extension = m.CategoryPrimitiveElement.Extension
	}
	r.Focus = m.Focus
	r.ResponseRequired = m.ResponseRequired
	if m.ResponseRequiredPrimitiveElement != nil {
		r.ResponseRequired.Id = m.ResponseRequiredPrimitiveElement.Id
		r.ResponseRequired.Extension = m.ResponseRequiredPrimitiveElement.Extension
	}
	r.AllowedResponse = m.AllowedResponse
	r.Graph = m.Graph
	for i, e := range m.GraphPrimitiveElement {
		if len(r.Graph) > i {
			r.Graph[i].Id = e.Id
			r.Graph[i].Extension = e.Extension
		} else {
			r.Graph = append(r.Graph, Canonical{Id: e.Id, Extension: e.Extension})
		}
	}
	return nil
}
func (r MessageDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Identifies the resource (or resources) that are being addressed by the event.  For example, the Encounter for an admit message or two Account records for a merge.
type MessageDefinitionFocus struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The kind of resource that must be the focus for this message.
	Code Code
	// A profile that reflects constraints for the focal resource (and potentially for related resources).
	Profile *Canonical
	// Identifies the minimum number of resources of this type that must be pointed to by a message in order for it to be valid against this MessageDefinition.
	Min UnsignedInt
	// Identifies the maximum number of resources of this type that must be pointed to by a message in order for it to be valid against this MessageDefinition.
	Max *String
}
type jsonMessageDefinitionFocus struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Code                    Code              `json:"code,omitempty"`
	CodePrimitiveElement    *primitiveElement `json:"_code,omitempty"`
	Profile                 *Canonical        `json:"profile,omitempty"`
	ProfilePrimitiveElement *primitiveElement `json:"_profile,omitempty"`
	Min                     UnsignedInt       `json:"min,omitempty"`
	MinPrimitiveElement     *primitiveElement `json:"_min,omitempty"`
	Max                     *String           `json:"max,omitempty"`
	MaxPrimitiveElement     *primitiveElement `json:"_max,omitempty"`
}

func (r MessageDefinitionFocus) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MessageDefinitionFocus) marshalJSON() jsonMessageDefinitionFocus {
	m := jsonMessageDefinitionFocus{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Profile = r.Profile
	if r.Profile != nil && (r.Profile.Id != nil || r.Profile.Extension != nil) {
		m.ProfilePrimitiveElement = &primitiveElement{Id: r.Profile.Id, Extension: r.Profile.Extension}
	}
	m.Min = r.Min
	if r.Min.Id != nil || r.Min.Extension != nil {
		m.MinPrimitiveElement = &primitiveElement{Id: r.Min.Id, Extension: r.Min.Extension}
	}
	m.Max = r.Max
	if r.Max != nil && (r.Max.Id != nil || r.Max.Extension != nil) {
		m.MaxPrimitiveElement = &primitiveElement{Id: r.Max.Id, Extension: r.Max.Extension}
	}
	return m
}
func (r *MessageDefinitionFocus) UnmarshalJSON(b []byte) error {
	var m jsonMessageDefinitionFocus
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MessageDefinitionFocus) unmarshalJSON(m jsonMessageDefinitionFocus) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Profile = m.Profile
	if m.ProfilePrimitiveElement != nil {
		r.Profile.Id = m.ProfilePrimitiveElement.Id
		r.Profile.Extension = m.ProfilePrimitiveElement.Extension
	}
	r.Min = m.Min
	if m.MinPrimitiveElement != nil {
		r.Min.Id = m.MinPrimitiveElement.Id
		r.Min.Extension = m.MinPrimitiveElement.Extension
	}
	r.Max = m.Max
	if m.MaxPrimitiveElement != nil {
		r.Max.Id = m.MaxPrimitiveElement.Id
		r.Max.Extension = m.MaxPrimitiveElement.Extension
	}
	return nil
}
func (r MessageDefinitionFocus) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates what types of messages may be sent as an application-level response to this message.
type MessageDefinitionAllowedResponse struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A reference to the message definition that must be adhered to by this supported response.
	Message Canonical
	// Provides a description of the circumstances in which this response should be used (as opposed to one of the alternative responses).
	Situation *Markdown
}
type jsonMessageDefinitionAllowedResponse struct {
	Id                        *string           `json:"id,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Message                   Canonical         `json:"message,omitempty"`
	MessagePrimitiveElement   *primitiveElement `json:"_message,omitempty"`
	Situation                 *Markdown         `json:"situation,omitempty"`
	SituationPrimitiveElement *primitiveElement `json:"_situation,omitempty"`
}

func (r MessageDefinitionAllowedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MessageDefinitionAllowedResponse) marshalJSON() jsonMessageDefinitionAllowedResponse {
	m := jsonMessageDefinitionAllowedResponse{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Message = r.Message
	if r.Message.Id != nil || r.Message.Extension != nil {
		m.MessagePrimitiveElement = &primitiveElement{Id: r.Message.Id, Extension: r.Message.Extension}
	}
	m.Situation = r.Situation
	if r.Situation != nil && (r.Situation.Id != nil || r.Situation.Extension != nil) {
		m.SituationPrimitiveElement = &primitiveElement{Id: r.Situation.Id, Extension: r.Situation.Extension}
	}
	return m
}
func (r *MessageDefinitionAllowedResponse) UnmarshalJSON(b []byte) error {
	var m jsonMessageDefinitionAllowedResponse
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MessageDefinitionAllowedResponse) unmarshalJSON(m jsonMessageDefinitionAllowedResponse) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Message = m.Message
	if m.MessagePrimitiveElement != nil {
		r.Message.Id = m.MessagePrimitiveElement.Id
		r.Message.Extension = m.MessagePrimitiveElement.Extension
	}
	r.Situation = m.Situation
	if m.SituationPrimitiveElement != nil {
		r.Situation.Id = m.SituationPrimitiveElement.Id
		r.Situation.Extension = m.SituationPrimitiveElement.Extension
	}
	return nil
}
func (r MessageDefinitionAllowedResponse) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
