package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A formal computable definition of a graph of resources - that is, a coherent set of resources that form a graph by following references. The Graph Definition resource defines a set and makes rules about the set.
type GraphDefinition struct {
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
	// An absolute URI that is used to identify this graph definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this graph definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the graph definition is stored on different servers.
	Url *Uri
	// The identifier that is used to identify this version of the graph definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the graph definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the graph definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name String
	// The status of this graph definition. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this graph definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the graph definition was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the graph definition changes.
	Date *DateTime
	// The name of the organization or individual that published the graph definition.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the graph definition from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate graph definition instances.
	UseContext []UsageContext
	// A legal or geographic region in which the graph definition is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this graph definition is needed and why it has been designed as it has.
	Purpose *Markdown
	// The type of FHIR resource at which instances of this graph start.
	Start Code
	// The profile that describes the use of the base resource.
	Profile *Canonical
	// Links this graph makes rules about.
	Link []GraphDefinitionLink
}

func (r GraphDefinition) ResourceType() string {
	return "GraphDefinition"
}
func (r GraphDefinition) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonGraphDefinition struct {
	ResourceType                  string                `json:"resourceType"`
	Id                            *Id                   `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement     `json:"_id,omitempty"`
	Meta                          *Meta                 `json:"meta,omitempty"`
	ImplicitRules                 *Uri                  `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement     `json:"_implicitRules,omitempty"`
	Language                      *Code                 `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement     `json:"_language,omitempty"`
	Text                          *Narrative            `json:"text,omitempty"`
	Contained                     []containedResource   `json:"contained,omitempty"`
	Extension                     []Extension           `json:"extension,omitempty"`
	ModifierExtension             []Extension           `json:"modifierExtension,omitempty"`
	Url                           *Uri                  `json:"url,omitempty"`
	UrlPrimitiveElement           *primitiveElement     `json:"_url,omitempty"`
	Version                       *String               `json:"version,omitempty"`
	VersionPrimitiveElement       *primitiveElement     `json:"_version,omitempty"`
	Name                          String                `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement     `json:"_name,omitempty"`
	Status                        Code                  `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement     `json:"_status,omitempty"`
	Experimental                  *Boolean              `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement  *primitiveElement     `json:"_experimental,omitempty"`
	Date                          *DateTime             `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement     `json:"_date,omitempty"`
	Publisher                     *String               `json:"publisher,omitempty"`
	PublisherPrimitiveElement     *primitiveElement     `json:"_publisher,omitempty"`
	Contact                       []ContactDetail       `json:"contact,omitempty"`
	Description                   *Markdown             `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement     `json:"_description,omitempty"`
	UseContext                    []UsageContext        `json:"useContext,omitempty"`
	Jurisdiction                  []CodeableConcept     `json:"jurisdiction,omitempty"`
	Purpose                       *Markdown             `json:"purpose,omitempty"`
	PurposePrimitiveElement       *primitiveElement     `json:"_purpose,omitempty"`
	Start                         Code                  `json:"start,omitempty"`
	StartPrimitiveElement         *primitiveElement     `json:"_start,omitempty"`
	Profile                       *Canonical            `json:"profile,omitempty"`
	ProfilePrimitiveElement       *primitiveElement     `json:"_profile,omitempty"`
	Link                          []GraphDefinitionLink `json:"link,omitempty"`
}

func (r GraphDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r GraphDefinition) marshalJSON() jsonGraphDefinition {
	m := jsonGraphDefinition{}
	m.ResourceType = "GraphDefinition"
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
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
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
	m.Jurisdiction = r.Jurisdiction
	m.Purpose = r.Purpose
	if r.Purpose != nil && (r.Purpose.Id != nil || r.Purpose.Extension != nil) {
		m.PurposePrimitiveElement = &primitiveElement{Id: r.Purpose.Id, Extension: r.Purpose.Extension}
	}
	m.Start = r.Start
	if r.Start.Id != nil || r.Start.Extension != nil {
		m.StartPrimitiveElement = &primitiveElement{Id: r.Start.Id, Extension: r.Start.Extension}
	}
	m.Profile = r.Profile
	if r.Profile != nil && (r.Profile.Id != nil || r.Profile.Extension != nil) {
		m.ProfilePrimitiveElement = &primitiveElement{Id: r.Profile.Id, Extension: r.Profile.Extension}
	}
	m.Link = r.Link
	return m
}
func (r *GraphDefinition) UnmarshalJSON(b []byte) error {
	var m jsonGraphDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *GraphDefinition) unmarshalJSON(m jsonGraphDefinition) error {
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
	r.Jurisdiction = m.Jurisdiction
	r.Purpose = m.Purpose
	if m.PurposePrimitiveElement != nil {
		r.Purpose.Id = m.PurposePrimitiveElement.Id
		r.Purpose.Extension = m.PurposePrimitiveElement.Extension
	}
	r.Start = m.Start
	if m.StartPrimitiveElement != nil {
		r.Start.Id = m.StartPrimitiveElement.Id
		r.Start.Extension = m.StartPrimitiveElement.Extension
	}
	r.Profile = m.Profile
	if m.ProfilePrimitiveElement != nil {
		r.Profile.Id = m.ProfilePrimitiveElement.Id
		r.Profile.Extension = m.ProfilePrimitiveElement.Extension
	}
	r.Link = m.Link
	return nil
}
func (r GraphDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Links this graph makes rules about.
type GraphDefinitionLink struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A FHIR expression that identifies one of FHIR References to other resources.
	Path *String
	// Which slice (if profiled).
	SliceName *String
	// Minimum occurrences for this link.
	Min *Integer
	// Maximum occurrences for this link.
	Max *String
	// Information about why this link is of interest in this graph definition.
	Description *String
	// Potential target for the link.
	Target []GraphDefinitionLinkTarget
}
type jsonGraphDefinitionLink struct {
	Id                          *string                     `json:"id,omitempty"`
	Extension                   []Extension                 `json:"extension,omitempty"`
	ModifierExtension           []Extension                 `json:"modifierExtension,omitempty"`
	Path                        *String                     `json:"path,omitempty"`
	PathPrimitiveElement        *primitiveElement           `json:"_path,omitempty"`
	SliceName                   *String                     `json:"sliceName,omitempty"`
	SliceNamePrimitiveElement   *primitiveElement           `json:"_sliceName,omitempty"`
	Min                         *Integer                    `json:"min,omitempty"`
	MinPrimitiveElement         *primitiveElement           `json:"_min,omitempty"`
	Max                         *String                     `json:"max,omitempty"`
	MaxPrimitiveElement         *primitiveElement           `json:"_max,omitempty"`
	Description                 *String                     `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement           `json:"_description,omitempty"`
	Target                      []GraphDefinitionLinkTarget `json:"target,omitempty"`
}

func (r GraphDefinitionLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r GraphDefinitionLink) marshalJSON() jsonGraphDefinitionLink {
	m := jsonGraphDefinitionLink{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Path = r.Path
	if r.Path != nil && (r.Path.Id != nil || r.Path.Extension != nil) {
		m.PathPrimitiveElement = &primitiveElement{Id: r.Path.Id, Extension: r.Path.Extension}
	}
	m.SliceName = r.SliceName
	if r.SliceName != nil && (r.SliceName.Id != nil || r.SliceName.Extension != nil) {
		m.SliceNamePrimitiveElement = &primitiveElement{Id: r.SliceName.Id, Extension: r.SliceName.Extension}
	}
	m.Min = r.Min
	if r.Min != nil && (r.Min.Id != nil || r.Min.Extension != nil) {
		m.MinPrimitiveElement = &primitiveElement{Id: r.Min.Id, Extension: r.Min.Extension}
	}
	m.Max = r.Max
	if r.Max != nil && (r.Max.Id != nil || r.Max.Extension != nil) {
		m.MaxPrimitiveElement = &primitiveElement{Id: r.Max.Id, Extension: r.Max.Extension}
	}
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Target = r.Target
	return m
}
func (r *GraphDefinitionLink) UnmarshalJSON(b []byte) error {
	var m jsonGraphDefinitionLink
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *GraphDefinitionLink) unmarshalJSON(m jsonGraphDefinitionLink) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Path = m.Path
	if m.PathPrimitiveElement != nil {
		r.Path.Id = m.PathPrimitiveElement.Id
		r.Path.Extension = m.PathPrimitiveElement.Extension
	}
	r.SliceName = m.SliceName
	if m.SliceNamePrimitiveElement != nil {
		r.SliceName.Id = m.SliceNamePrimitiveElement.Id
		r.SliceName.Extension = m.SliceNamePrimitiveElement.Extension
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
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Target = m.Target
	return nil
}
func (r GraphDefinitionLink) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Potential target for the link.
type GraphDefinitionLinkTarget struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Type of resource this link refers to.
	Type Code
	// A set of parameters to look up.
	Params *String
	// Profile for the target resource.
	Profile *Canonical
	// Compartment Consistency Rules.
	Compartment []GraphDefinitionLinkTargetCompartment
	// Additional links from target resource.
	Link []GraphDefinitionLink
}
type jsonGraphDefinitionLinkTarget struct {
	Id                      *string                                `json:"id,omitempty"`
	Extension               []Extension                            `json:"extension,omitempty"`
	ModifierExtension       []Extension                            `json:"modifierExtension,omitempty"`
	Type                    Code                                   `json:"type,omitempty"`
	TypePrimitiveElement    *primitiveElement                      `json:"_type,omitempty"`
	Params                  *String                                `json:"params,omitempty"`
	ParamsPrimitiveElement  *primitiveElement                      `json:"_params,omitempty"`
	Profile                 *Canonical                             `json:"profile,omitempty"`
	ProfilePrimitiveElement *primitiveElement                      `json:"_profile,omitempty"`
	Compartment             []GraphDefinitionLinkTargetCompartment `json:"compartment,omitempty"`
	Link                    []GraphDefinitionLink                  `json:"link,omitempty"`
}

func (r GraphDefinitionLinkTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r GraphDefinitionLinkTarget) marshalJSON() jsonGraphDefinitionLinkTarget {
	m := jsonGraphDefinitionLinkTarget{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Params = r.Params
	if r.Params != nil && (r.Params.Id != nil || r.Params.Extension != nil) {
		m.ParamsPrimitiveElement = &primitiveElement{Id: r.Params.Id, Extension: r.Params.Extension}
	}
	m.Profile = r.Profile
	if r.Profile != nil && (r.Profile.Id != nil || r.Profile.Extension != nil) {
		m.ProfilePrimitiveElement = &primitiveElement{Id: r.Profile.Id, Extension: r.Profile.Extension}
	}
	m.Compartment = r.Compartment
	m.Link = r.Link
	return m
}
func (r *GraphDefinitionLinkTarget) UnmarshalJSON(b []byte) error {
	var m jsonGraphDefinitionLinkTarget
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *GraphDefinitionLinkTarget) unmarshalJSON(m jsonGraphDefinitionLinkTarget) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Params = m.Params
	if m.ParamsPrimitiveElement != nil {
		r.Params.Id = m.ParamsPrimitiveElement.Id
		r.Params.Extension = m.ParamsPrimitiveElement.Extension
	}
	r.Profile = m.Profile
	if m.ProfilePrimitiveElement != nil {
		r.Profile.Id = m.ProfilePrimitiveElement.Id
		r.Profile.Extension = m.ProfilePrimitiveElement.Extension
	}
	r.Compartment = m.Compartment
	r.Link = m.Link
	return nil
}
func (r GraphDefinitionLinkTarget) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Compartment Consistency Rules.
type GraphDefinitionLinkTargetCompartment struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Defines how the compartment rule is used - whether it it is used to test whether resources are subject to the rule, or whether it is a rule that must be followed.
	Use Code
	// Identifies the compartment.
	Code Code
	// identical | matching | different | no-rule | custom.
	Rule Code
	// Custom rule, as a FHIRPath expression.
	Expression *String
	// Documentation for FHIRPath expression.
	Description *String
}
type jsonGraphDefinitionLinkTargetCompartment struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Use                         Code              `json:"use,omitempty"`
	UsePrimitiveElement         *primitiveElement `json:"_use,omitempty"`
	Code                        Code              `json:"code,omitempty"`
	CodePrimitiveElement        *primitiveElement `json:"_code,omitempty"`
	Rule                        Code              `json:"rule,omitempty"`
	RulePrimitiveElement        *primitiveElement `json:"_rule,omitempty"`
	Expression                  *String           `json:"expression,omitempty"`
	ExpressionPrimitiveElement  *primitiveElement `json:"_expression,omitempty"`
	Description                 *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
}

func (r GraphDefinitionLinkTargetCompartment) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r GraphDefinitionLinkTargetCompartment) marshalJSON() jsonGraphDefinitionLinkTargetCompartment {
	m := jsonGraphDefinitionLinkTargetCompartment{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Use = r.Use
	if r.Use.Id != nil || r.Use.Extension != nil {
		m.UsePrimitiveElement = &primitiveElement{Id: r.Use.Id, Extension: r.Use.Extension}
	}
	m.Code = r.Code
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Rule = r.Rule
	if r.Rule.Id != nil || r.Rule.Extension != nil {
		m.RulePrimitiveElement = &primitiveElement{Id: r.Rule.Id, Extension: r.Rule.Extension}
	}
	m.Expression = r.Expression
	if r.Expression != nil && (r.Expression.Id != nil || r.Expression.Extension != nil) {
		m.ExpressionPrimitiveElement = &primitiveElement{Id: r.Expression.Id, Extension: r.Expression.Extension}
	}
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	return m
}
func (r *GraphDefinitionLinkTargetCompartment) UnmarshalJSON(b []byte) error {
	var m jsonGraphDefinitionLinkTargetCompartment
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *GraphDefinitionLinkTargetCompartment) unmarshalJSON(m jsonGraphDefinitionLinkTargetCompartment) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Use = m.Use
	if m.UsePrimitiveElement != nil {
		r.Use.Id = m.UsePrimitiveElement.Id
		r.Use.Extension = m.UsePrimitiveElement.Extension
	}
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Rule = m.Rule
	if m.RulePrimitiveElement != nil {
		r.Rule.Id = m.RulePrimitiveElement.Id
		r.Rule.Extension = m.RulePrimitiveElement.Extension
	}
	r.Expression = m.Expression
	if m.ExpressionPrimitiveElement != nil {
		r.Expression.Id = m.ExpressionPrimitiveElement.Id
		r.Expression.Extension = m.ExpressionPrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	return nil
}
func (r GraphDefinitionLinkTargetCompartment) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
