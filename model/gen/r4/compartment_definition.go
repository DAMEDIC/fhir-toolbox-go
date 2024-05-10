package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A compartment definition that defines how resources are accessed on a server.
type CompartmentDefinition struct {
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
	// An absolute URI that is used to identify this compartment definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this compartment definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the compartment definition is stored on different servers.
	Url Uri
	// The identifier that is used to identify this version of the compartment definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the compartment definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the compartment definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name String
	// The status of this compartment definition. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this compartment definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the compartment definition was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the compartment definition changes.
	Date *DateTime
	// The name of the organization or individual that published the compartment definition.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the compartment definition from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate compartment definition instances.
	UseContext []UsageContext
	// Explanation of why this compartment definition is needed and why it has been designed as it has.
	Purpose *Markdown
	// Which compartment this definition describes.
	Code Code
	// Whether the search syntax is supported,.
	Search Boolean
	// Information about how a resource is related to the compartment.
	Resource []CompartmentDefinitionResource
}
type jsonCompartmentDefinition struct {
	ResourceType                  string                          `json:"resourceType"`
	Id                            *Id                             `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement               `json:"_id,omitempty"`
	Meta                          *Meta                           `json:"meta,omitempty"`
	ImplicitRules                 *Uri                            `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement               `json:"_implicitRules,omitempty"`
	Language                      *Code                           `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement               `json:"_language,omitempty"`
	Text                          *Narrative                      `json:"text,omitempty"`
	Contained                     []containedResource             `json:"contained,omitempty"`
	Extension                     []Extension                     `json:"extension,omitempty"`
	ModifierExtension             []Extension                     `json:"modifierExtension,omitempty"`
	Url                           Uri                             `json:"url,omitempty"`
	UrlPrimitiveElement           *primitiveElement               `json:"_url,omitempty"`
	Version                       *String                         `json:"version,omitempty"`
	VersionPrimitiveElement       *primitiveElement               `json:"_version,omitempty"`
	Name                          String                          `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement               `json:"_name,omitempty"`
	Status                        Code                            `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement               `json:"_status,omitempty"`
	Experimental                  *Boolean                        `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement  *primitiveElement               `json:"_experimental,omitempty"`
	Date                          *DateTime                       `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement               `json:"_date,omitempty"`
	Publisher                     *String                         `json:"publisher,omitempty"`
	PublisherPrimitiveElement     *primitiveElement               `json:"_publisher,omitempty"`
	Contact                       []ContactDetail                 `json:"contact,omitempty"`
	Description                   *Markdown                       `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement               `json:"_description,omitempty"`
	UseContext                    []UsageContext                  `json:"useContext,omitempty"`
	Purpose                       *Markdown                       `json:"purpose,omitempty"`
	PurposePrimitiveElement       *primitiveElement               `json:"_purpose,omitempty"`
	Code                          Code                            `json:"code,omitempty"`
	CodePrimitiveElement          *primitiveElement               `json:"_code,omitempty"`
	Search                        Boolean                         `json:"search,omitempty"`
	SearchPrimitiveElement        *primitiveElement               `json:"_search,omitempty"`
	Resource                      []CompartmentDefinitionResource `json:"resource,omitempty"`
}

func (r CompartmentDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CompartmentDefinition) marshalJSON() jsonCompartmentDefinition {
	m := jsonCompartmentDefinition{}
	m.ResourceType = "CompartmentDefinition"
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
	m.Url = r.Url
	if r.Url.Id != nil || r.Url.Extension != nil {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	m.Version = r.Version
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	m.Name = r.Name
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
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
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
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
	m.Purpose = r.Purpose
	if r.Purpose != nil && (r.Purpose.Id != nil || r.Purpose.Extension != nil) {
		m.PurposePrimitiveElement = &primitiveElement{Id: r.Purpose.Id, Extension: r.Purpose.Extension}
	}
	m.Code = r.Code
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Search = r.Search
	if r.Search.Id != nil || r.Search.Extension != nil {
		m.SearchPrimitiveElement = &primitiveElement{Id: r.Search.Id, Extension: r.Search.Extension}
	}
	m.Resource = r.Resource
	return m
}
func (r *CompartmentDefinition) UnmarshalJSON(b []byte) error {
	var m jsonCompartmentDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CompartmentDefinition) unmarshalJSON(m jsonCompartmentDefinition) error {
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
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
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
	r.Purpose = m.Purpose
	if m.PurposePrimitiveElement != nil {
		r.Purpose.Id = m.PurposePrimitiveElement.Id
		r.Purpose.Extension = m.PurposePrimitiveElement.Extension
	}
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Search = m.Search
	if m.SearchPrimitiveElement != nil {
		r.Search.Id = m.SearchPrimitiveElement.Id
		r.Search.Extension = m.SearchPrimitiveElement.Extension
	}
	r.Resource = m.Resource
	return nil
}
func (r CompartmentDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about how a resource is related to the compartment.
type CompartmentDefinitionResource struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The name of a resource supported by the server.
	Code Code
	// The name of a search parameter that represents the link to the compartment. More than one may be listed because a resource may be linked to a compartment in more than one way,.
	Param []String
	// Additional documentation about the resource and compartment.
	Documentation *String
}
type jsonCompartmentDefinitionResource struct {
	Id                            *string             `json:"id,omitempty"`
	Extension                     []Extension         `json:"extension,omitempty"`
	ModifierExtension             []Extension         `json:"modifierExtension,omitempty"`
	Code                          Code                `json:"code,omitempty"`
	CodePrimitiveElement          *primitiveElement   `json:"_code,omitempty"`
	Param                         []String            `json:"param,omitempty"`
	ParamPrimitiveElement         []*primitiveElement `json:"_param,omitempty"`
	Documentation                 *String             `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement   `json:"_documentation,omitempty"`
}

func (r CompartmentDefinitionResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CompartmentDefinitionResource) marshalJSON() jsonCompartmentDefinitionResource {
	m := jsonCompartmentDefinitionResource{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Param = r.Param
	anyParamIdOrExtension := false
	for _, e := range r.Param {
		if e.Id != nil || e.Extension != nil {
			anyParamIdOrExtension = true
			break
		}
	}
	if anyParamIdOrExtension {
		m.ParamPrimitiveElement = make([]*primitiveElement, 0, len(r.Param))
		for _, e := range r.Param {
			if e.Id != nil || e.Extension != nil {
				m.ParamPrimitiveElement = append(m.ParamPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ParamPrimitiveElement = append(m.ParamPrimitiveElement, nil)
			}
		}
	}
	m.Documentation = r.Documentation
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	return m
}
func (r *CompartmentDefinitionResource) UnmarshalJSON(b []byte) error {
	var m jsonCompartmentDefinitionResource
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CompartmentDefinitionResource) unmarshalJSON(m jsonCompartmentDefinitionResource) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Param = m.Param
	for i, e := range m.ParamPrimitiveElement {
		if len(r.Param) > i {
			r.Param[i].Id = e.Id
			r.Param[i].Extension = e.Extension
		} else {
			r.Param = append(r.Param, String{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	return nil
}
func (r CompartmentDefinitionResource) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
