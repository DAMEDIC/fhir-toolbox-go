package r4

import (
	"encoding/json"
	"fmt"
)

// A Map of relationships between 2 structures that can be used to transform data.
type StructureMap struct {
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
	// An absolute URI that is used to identify this structure map when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this structure map is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the structure map is stored on different servers.
	Url Uri
	// A formal identifier that is used to identify this structure map when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the structure map when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the structure map author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the structure map. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name String
	// A short, descriptive, user-friendly title for the structure map.
	Title *String
	// The status of this structure map. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this structure map is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the structure map was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the structure map changes.
	Date *DateTime
	// The name of the organization or individual that published the structure map.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the structure map from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate structure map instances.
	UseContext []UsageContext
	// A legal or geographic region in which the structure map is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this structure map is needed and why it has been designed as it has.
	Purpose *Markdown
	// A copyright statement relating to the structure map and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the structure map.
	Copyright *Markdown
	// A structure definition used by this map. The structure definition may describe instances that are converted, or the instances that are produced.
	Structure []StructureMapStructure
	// Other maps used by this map (canonical URLs).
	Import []Canonical
	// Organizes the mapping into manageable chunks for human review/ease of maintenance.
	Group []StructureMapGroup
}
type jsonStructureMap struct {
	ResourceType                  string                  `json:"resourceType"`
	Id                            *Id                     `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement       `json:"_id,omitempty"`
	Meta                          *Meta                   `json:"meta,omitempty"`
	ImplicitRules                 *Uri                    `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement       `json:"_implicitRules,omitempty"`
	Language                      *Code                   `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement       `json:"_language,omitempty"`
	Text                          *Narrative              `json:"text,omitempty"`
	Contained                     []containedResource     `json:"contained,omitempty"`
	Extension                     []Extension             `json:"extension,omitempty"`
	ModifierExtension             []Extension             `json:"modifierExtension,omitempty"`
	Url                           Uri                     `json:"url,omitempty"`
	UrlPrimitiveElement           *primitiveElement       `json:"_url,omitempty"`
	Identifier                    []Identifier            `json:"identifier,omitempty"`
	Version                       *String                 `json:"version,omitempty"`
	VersionPrimitiveElement       *primitiveElement       `json:"_version,omitempty"`
	Name                          String                  `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement       `json:"_name,omitempty"`
	Title                         *String                 `json:"title,omitempty"`
	TitlePrimitiveElement         *primitiveElement       `json:"_title,omitempty"`
	Status                        Code                    `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement       `json:"_status,omitempty"`
	Experimental                  *Boolean                `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement  *primitiveElement       `json:"_experimental,omitempty"`
	Date                          *DateTime               `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement       `json:"_date,omitempty"`
	Publisher                     *String                 `json:"publisher,omitempty"`
	PublisherPrimitiveElement     *primitiveElement       `json:"_publisher,omitempty"`
	Contact                       []ContactDetail         `json:"contact,omitempty"`
	Description                   *Markdown               `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement       `json:"_description,omitempty"`
	UseContext                    []UsageContext          `json:"useContext,omitempty"`
	Jurisdiction                  []CodeableConcept       `json:"jurisdiction,omitempty"`
	Purpose                       *Markdown               `json:"purpose,omitempty"`
	PurposePrimitiveElement       *primitiveElement       `json:"_purpose,omitempty"`
	Copyright                     *Markdown               `json:"copyright,omitempty"`
	CopyrightPrimitiveElement     *primitiveElement       `json:"_copyright,omitempty"`
	Structure                     []StructureMapStructure `json:"structure,omitempty"`
	Import                        []Canonical             `json:"import,omitempty"`
	ImportPrimitiveElement        []*primitiveElement     `json:"_import,omitempty"`
	Group                         []StructureMapGroup     `json:"group,omitempty"`
}

func (r StructureMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r StructureMap) marshalJSON() jsonStructureMap {
	m := jsonStructureMap{}
	m.ResourceType = "StructureMap"
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
	m.Identifier = r.Identifier
	m.Version = r.Version
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	m.Name = r.Name
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Title = r.Title
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
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
	m.Copyright = r.Copyright
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		m.CopyrightPrimitiveElement = &primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
	}
	m.Structure = r.Structure
	m.Import = r.Import
	anyImportIdOrExtension := false
	for _, e := range r.Import {
		if e.Id != nil || e.Extension != nil {
			anyImportIdOrExtension = true
			break
		}
	}
	if anyImportIdOrExtension {
		m.ImportPrimitiveElement = make([]*primitiveElement, 0, len(r.Import))
		for _, e := range r.Import {
			if e.Id != nil || e.Extension != nil {
				m.ImportPrimitiveElement = append(m.ImportPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ImportPrimitiveElement = append(m.ImportPrimitiveElement, nil)
			}
		}
	}
	m.Group = r.Group
	return m
}
func (r *StructureMap) UnmarshalJSON(b []byte) error {
	var m jsonStructureMap
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *StructureMap) unmarshalJSON(m jsonStructureMap) error {
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
	r.Structure = m.Structure
	r.Import = m.Import
	for i, e := range m.ImportPrimitiveElement {
		if len(r.Import) > i {
			r.Import[i].Id = e.Id
			r.Import[i].Extension = e.Extension
		} else {
			r.Import = append(r.Import, Canonical{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Group = m.Group
	return nil
}
func (r StructureMap) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A structure definition used by this map. The structure definition may describe instances that are converted, or the instances that are produced.
type StructureMapStructure struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The canonical reference to the structure.
	Url Canonical
	// How the referenced structure is used in this mapping.
	Mode Code
	// The name used for this type in the map.
	Alias *String
	// Documentation that describes how the structure is used in the mapping.
	Documentation *String
}
type jsonStructureMapStructure struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	Url                           Canonical         `json:"url,omitempty"`
	UrlPrimitiveElement           *primitiveElement `json:"_url,omitempty"`
	Mode                          Code              `json:"mode,omitempty"`
	ModePrimitiveElement          *primitiveElement `json:"_mode,omitempty"`
	Alias                         *String           `json:"alias,omitempty"`
	AliasPrimitiveElement         *primitiveElement `json:"_alias,omitempty"`
	Documentation                 *String           `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement `json:"_documentation,omitempty"`
}

func (r StructureMapStructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r StructureMapStructure) marshalJSON() jsonStructureMapStructure {
	m := jsonStructureMapStructure{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Url = r.Url
	if r.Url.Id != nil || r.Url.Extension != nil {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	m.Mode = r.Mode
	if r.Mode.Id != nil || r.Mode.Extension != nil {
		m.ModePrimitiveElement = &primitiveElement{Id: r.Mode.Id, Extension: r.Mode.Extension}
	}
	m.Alias = r.Alias
	if r.Alias != nil && (r.Alias.Id != nil || r.Alias.Extension != nil) {
		m.AliasPrimitiveElement = &primitiveElement{Id: r.Alias.Id, Extension: r.Alias.Extension}
	}
	m.Documentation = r.Documentation
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	return m
}
func (r *StructureMapStructure) UnmarshalJSON(b []byte) error {
	var m jsonStructureMapStructure
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *StructureMapStructure) unmarshalJSON(m jsonStructureMapStructure) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Mode = m.Mode
	if m.ModePrimitiveElement != nil {
		r.Mode.Id = m.ModePrimitiveElement.Id
		r.Mode.Extension = m.ModePrimitiveElement.Extension
	}
	r.Alias = m.Alias
	if m.AliasPrimitiveElement != nil {
		r.Alias.Id = m.AliasPrimitiveElement.Id
		r.Alias.Extension = m.AliasPrimitiveElement.Extension
	}
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	return nil
}
func (r StructureMapStructure) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Organizes the mapping into manageable chunks for human review/ease of maintenance.
type StructureMapGroup struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A unique name for the group for the convenience of human readers.
	Name Id
	// Another group that this group adds rules to.
	Extends *Id
	// If this is the default rule set to apply for the source type or this combination of types.
	TypeMode Code
	// Additional supporting documentation that explains the purpose of the group and the types of mappings within it.
	Documentation *String
	// A name assigned to an instance of data. The instance must be provided when the mapping is invoked.
	Input []StructureMapGroupInput
	// Transform Rule from source to target.
	Rule []StructureMapGroupRule
}
type jsonStructureMapGroup struct {
	Id                            *string                  `json:"id,omitempty"`
	Extension                     []Extension              `json:"extension,omitempty"`
	ModifierExtension             []Extension              `json:"modifierExtension,omitempty"`
	Name                          Id                       `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement        `json:"_name,omitempty"`
	Extends                       *Id                      `json:"extends,omitempty"`
	ExtendsPrimitiveElement       *primitiveElement        `json:"_extends,omitempty"`
	TypeMode                      Code                     `json:"typeMode,omitempty"`
	TypeModePrimitiveElement      *primitiveElement        `json:"_typeMode,omitempty"`
	Documentation                 *String                  `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement        `json:"_documentation,omitempty"`
	Input                         []StructureMapGroupInput `json:"input,omitempty"`
	Rule                          []StructureMapGroupRule  `json:"rule,omitempty"`
}

func (r StructureMapGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r StructureMapGroup) marshalJSON() jsonStructureMapGroup {
	m := jsonStructureMapGroup{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Name = r.Name
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Extends = r.Extends
	if r.Extends != nil && (r.Extends.Id != nil || r.Extends.Extension != nil) {
		m.ExtendsPrimitiveElement = &primitiveElement{Id: r.Extends.Id, Extension: r.Extends.Extension}
	}
	m.TypeMode = r.TypeMode
	if r.TypeMode.Id != nil || r.TypeMode.Extension != nil {
		m.TypeModePrimitiveElement = &primitiveElement{Id: r.TypeMode.Id, Extension: r.TypeMode.Extension}
	}
	m.Documentation = r.Documentation
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	m.Input = r.Input
	m.Rule = r.Rule
	return m
}
func (r *StructureMapGroup) UnmarshalJSON(b []byte) error {
	var m jsonStructureMapGroup
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *StructureMapGroup) unmarshalJSON(m jsonStructureMapGroup) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Extends = m.Extends
	if m.ExtendsPrimitiveElement != nil {
		r.Extends.Id = m.ExtendsPrimitiveElement.Id
		r.Extends.Extension = m.ExtendsPrimitiveElement.Extension
	}
	r.TypeMode = m.TypeMode
	if m.TypeModePrimitiveElement != nil {
		r.TypeMode.Id = m.TypeModePrimitiveElement.Id
		r.TypeMode.Extension = m.TypeModePrimitiveElement.Extension
	}
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	r.Input = m.Input
	r.Rule = m.Rule
	return nil
}
func (r StructureMapGroup) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A name assigned to an instance of data. The instance must be provided when the mapping is invoked.
type StructureMapGroupInput struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Name for this instance of data.
	Name Id
	// Type for this instance of data.
	Type *String
	// Mode for this instance of data.
	Mode Code
	// Documentation for this instance of data.
	Documentation *String
}
type jsonStructureMapGroupInput struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	Name                          Id                `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement `json:"_name,omitempty"`
	Type                          *String           `json:"type,omitempty"`
	TypePrimitiveElement          *primitiveElement `json:"_type,omitempty"`
	Mode                          Code              `json:"mode,omitempty"`
	ModePrimitiveElement          *primitiveElement `json:"_mode,omitempty"`
	Documentation                 *String           `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement `json:"_documentation,omitempty"`
}

func (r StructureMapGroupInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r StructureMapGroupInput) marshalJSON() jsonStructureMapGroupInput {
	m := jsonStructureMapGroupInput{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Name = r.Name
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Type = r.Type
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Mode = r.Mode
	if r.Mode.Id != nil || r.Mode.Extension != nil {
		m.ModePrimitiveElement = &primitiveElement{Id: r.Mode.Id, Extension: r.Mode.Extension}
	}
	m.Documentation = r.Documentation
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	return m
}
func (r *StructureMapGroupInput) UnmarshalJSON(b []byte) error {
	var m jsonStructureMapGroupInput
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *StructureMapGroupInput) unmarshalJSON(m jsonStructureMapGroupInput) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Mode = m.Mode
	if m.ModePrimitiveElement != nil {
		r.Mode.Id = m.ModePrimitiveElement.Id
		r.Mode.Extension = m.ModePrimitiveElement.Extension
	}
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	return nil
}
func (r StructureMapGroupInput) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Transform Rule from source to target.
type StructureMapGroupRule struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Name of the rule for internal references.
	Name Id
	// Source inputs to the mapping.
	Source []StructureMapGroupRuleSource
	// Content to create because of this mapping rule.
	Target []StructureMapGroupRuleTarget
	// Rules contained in this rule.
	Rule []StructureMapGroupRule
	// Which other rules to apply in the context of this rule.
	Dependent []StructureMapGroupRuleDependent
	// Documentation for this instance of data.
	Documentation *String
}
type jsonStructureMapGroupRule struct {
	Id                            *string                          `json:"id,omitempty"`
	Extension                     []Extension                      `json:"extension,omitempty"`
	ModifierExtension             []Extension                      `json:"modifierExtension,omitempty"`
	Name                          Id                               `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement                `json:"_name,omitempty"`
	Source                        []StructureMapGroupRuleSource    `json:"source,omitempty"`
	Target                        []StructureMapGroupRuleTarget    `json:"target,omitempty"`
	Rule                          []StructureMapGroupRule          `json:"rule,omitempty"`
	Dependent                     []StructureMapGroupRuleDependent `json:"dependent,omitempty"`
	Documentation                 *String                          `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement                `json:"_documentation,omitempty"`
}

func (r StructureMapGroupRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r StructureMapGroupRule) marshalJSON() jsonStructureMapGroupRule {
	m := jsonStructureMapGroupRule{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Name = r.Name
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Source = r.Source
	m.Target = r.Target
	m.Rule = r.Rule
	m.Dependent = r.Dependent
	m.Documentation = r.Documentation
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	return m
}
func (r *StructureMapGroupRule) UnmarshalJSON(b []byte) error {
	var m jsonStructureMapGroupRule
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *StructureMapGroupRule) unmarshalJSON(m jsonStructureMapGroupRule) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Source = m.Source
	r.Target = m.Target
	r.Rule = m.Rule
	r.Dependent = m.Dependent
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	return nil
}
func (r StructureMapGroupRule) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Source inputs to the mapping.
type StructureMapGroupRuleSource struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Type or variable this rule applies to.
	Context Id
	// Specified minimum cardinality for the element. This is optional; if present, it acts an implicit check on the input content.
	Min *Integer
	// Specified maximum cardinality for the element - a number or a "*". This is optional; if present, it acts an implicit check on the input content (* just serves as documentation; it's the default value).
	Max *String
	// Specified type for the element. This works as a condition on the mapping - use for polymorphic elements.
	Type *String
	// A value to use if there is no existing value in the source object.
	DefaultValue isStructureMapGroupRuleSourceDefaultValue
	// Optional field for this source.
	Element *String
	// How to handle the list mode for this element.
	ListMode *Code
	// Named context for field, if a field is specified.
	Variable *Id
	// FHIRPath expression  - must be true or the rule does not apply.
	Condition *String
	// FHIRPath expression  - must be true or the mapping engine throws an error instead of completing.
	Check *String
	// A FHIRPath expression which specifies a message to put in the transform log when content matching the source rule is found.
	LogMessage *String
}
type isStructureMapGroupRuleSourceDefaultValue interface {
	isStructureMapGroupRuleSourceDefaultValue()
}

func (r Base64Binary) isStructureMapGroupRuleSourceDefaultValue()        {}
func (r Boolean) isStructureMapGroupRuleSourceDefaultValue()             {}
func (r Canonical) isStructureMapGroupRuleSourceDefaultValue()           {}
func (r Code) isStructureMapGroupRuleSourceDefaultValue()                {}
func (r Date) isStructureMapGroupRuleSourceDefaultValue()                {}
func (r DateTime) isStructureMapGroupRuleSourceDefaultValue()            {}
func (r Decimal) isStructureMapGroupRuleSourceDefaultValue()             {}
func (r Id) isStructureMapGroupRuleSourceDefaultValue()                  {}
func (r Instant) isStructureMapGroupRuleSourceDefaultValue()             {}
func (r Integer) isStructureMapGroupRuleSourceDefaultValue()             {}
func (r Markdown) isStructureMapGroupRuleSourceDefaultValue()            {}
func (r Oid) isStructureMapGroupRuleSourceDefaultValue()                 {}
func (r PositiveInt) isStructureMapGroupRuleSourceDefaultValue()         {}
func (r String) isStructureMapGroupRuleSourceDefaultValue()              {}
func (r Time) isStructureMapGroupRuleSourceDefaultValue()                {}
func (r UnsignedInt) isStructureMapGroupRuleSourceDefaultValue()         {}
func (r Uri) isStructureMapGroupRuleSourceDefaultValue()                 {}
func (r Url) isStructureMapGroupRuleSourceDefaultValue()                 {}
func (r Uuid) isStructureMapGroupRuleSourceDefaultValue()                {}
func (r Address) isStructureMapGroupRuleSourceDefaultValue()             {}
func (r Age) isStructureMapGroupRuleSourceDefaultValue()                 {}
func (r Annotation) isStructureMapGroupRuleSourceDefaultValue()          {}
func (r Attachment) isStructureMapGroupRuleSourceDefaultValue()          {}
func (r CodeableConcept) isStructureMapGroupRuleSourceDefaultValue()     {}
func (r Coding) isStructureMapGroupRuleSourceDefaultValue()              {}
func (r ContactPoint) isStructureMapGroupRuleSourceDefaultValue()        {}
func (r Count) isStructureMapGroupRuleSourceDefaultValue()               {}
func (r Distance) isStructureMapGroupRuleSourceDefaultValue()            {}
func (r Duration) isStructureMapGroupRuleSourceDefaultValue()            {}
func (r HumanName) isStructureMapGroupRuleSourceDefaultValue()           {}
func (r Identifier) isStructureMapGroupRuleSourceDefaultValue()          {}
func (r Money) isStructureMapGroupRuleSourceDefaultValue()               {}
func (r Period) isStructureMapGroupRuleSourceDefaultValue()              {}
func (r Quantity) isStructureMapGroupRuleSourceDefaultValue()            {}
func (r Range) isStructureMapGroupRuleSourceDefaultValue()               {}
func (r Ratio) isStructureMapGroupRuleSourceDefaultValue()               {}
func (r Reference) isStructureMapGroupRuleSourceDefaultValue()           {}
func (r SampledData) isStructureMapGroupRuleSourceDefaultValue()         {}
func (r Signature) isStructureMapGroupRuleSourceDefaultValue()           {}
func (r Timing) isStructureMapGroupRuleSourceDefaultValue()              {}
func (r ContactDetail) isStructureMapGroupRuleSourceDefaultValue()       {}
func (r Contributor) isStructureMapGroupRuleSourceDefaultValue()         {}
func (r DataRequirement) isStructureMapGroupRuleSourceDefaultValue()     {}
func (r Expression) isStructureMapGroupRuleSourceDefaultValue()          {}
func (r ParameterDefinition) isStructureMapGroupRuleSourceDefaultValue() {}
func (r RelatedArtifact) isStructureMapGroupRuleSourceDefaultValue()     {}
func (r TriggerDefinition) isStructureMapGroupRuleSourceDefaultValue()   {}
func (r UsageContext) isStructureMapGroupRuleSourceDefaultValue()        {}
func (r Dosage) isStructureMapGroupRuleSourceDefaultValue()              {}
func (r Meta) isStructureMapGroupRuleSourceDefaultValue()                {}

type jsonStructureMapGroupRuleSource struct {
	Id                                       *string              `json:"id,omitempty"`
	Extension                                []Extension          `json:"extension,omitempty"`
	ModifierExtension                        []Extension          `json:"modifierExtension,omitempty"`
	Context                                  Id                   `json:"context,omitempty"`
	ContextPrimitiveElement                  *primitiveElement    `json:"_context,omitempty"`
	Min                                      *Integer             `json:"min,omitempty"`
	MinPrimitiveElement                      *primitiveElement    `json:"_min,omitempty"`
	Max                                      *String              `json:"max,omitempty"`
	MaxPrimitiveElement                      *primitiveElement    `json:"_max,omitempty"`
	Type                                     *String              `json:"type,omitempty"`
	TypePrimitiveElement                     *primitiveElement    `json:"_type,omitempty"`
	DefaultValueBase64Binary                 *Base64Binary        `json:"defaultValueBase64Binary,omitempty"`
	DefaultValueBase64BinaryPrimitiveElement *primitiveElement    `json:"_defaultValueBase64Binary,omitempty"`
	DefaultValueBoolean                      *Boolean             `json:"defaultValueBoolean,omitempty"`
	DefaultValueBooleanPrimitiveElement      *primitiveElement    `json:"_defaultValueBoolean,omitempty"`
	DefaultValueCanonical                    *Canonical           `json:"defaultValueCanonical,omitempty"`
	DefaultValueCanonicalPrimitiveElement    *primitiveElement    `json:"_defaultValueCanonical,omitempty"`
	DefaultValueCode                         *Code                `json:"defaultValueCode,omitempty"`
	DefaultValueCodePrimitiveElement         *primitiveElement    `json:"_defaultValueCode,omitempty"`
	DefaultValueDate                         *Date                `json:"defaultValueDate,omitempty"`
	DefaultValueDatePrimitiveElement         *primitiveElement    `json:"_defaultValueDate,omitempty"`
	DefaultValueDateTime                     *DateTime            `json:"defaultValueDateTime,omitempty"`
	DefaultValueDateTimePrimitiveElement     *primitiveElement    `json:"_defaultValueDateTime,omitempty"`
	DefaultValueDecimal                      *Decimal             `json:"defaultValueDecimal,omitempty"`
	DefaultValueDecimalPrimitiveElement      *primitiveElement    `json:"_defaultValueDecimal,omitempty"`
	DefaultValueId                           *Id                  `json:"defaultValueId,omitempty"`
	DefaultValueIdPrimitiveElement           *primitiveElement    `json:"_defaultValueId,omitempty"`
	DefaultValueInstant                      *Instant             `json:"defaultValueInstant,omitempty"`
	DefaultValueInstantPrimitiveElement      *primitiveElement    `json:"_defaultValueInstant,omitempty"`
	DefaultValueInteger                      *Integer             `json:"defaultValueInteger,omitempty"`
	DefaultValueIntegerPrimitiveElement      *primitiveElement    `json:"_defaultValueInteger,omitempty"`
	DefaultValueMarkdown                     *Markdown            `json:"defaultValueMarkdown,omitempty"`
	DefaultValueMarkdownPrimitiveElement     *primitiveElement    `json:"_defaultValueMarkdown,omitempty"`
	DefaultValueOid                          *Oid                 `json:"defaultValueOid,omitempty"`
	DefaultValueOidPrimitiveElement          *primitiveElement    `json:"_defaultValueOid,omitempty"`
	DefaultValuePositiveInt                  *PositiveInt         `json:"defaultValuePositiveInt,omitempty"`
	DefaultValuePositiveIntPrimitiveElement  *primitiveElement    `json:"_defaultValuePositiveInt,omitempty"`
	DefaultValueString                       *String              `json:"defaultValueString,omitempty"`
	DefaultValueStringPrimitiveElement       *primitiveElement    `json:"_defaultValueString,omitempty"`
	DefaultValueTime                         *Time                `json:"defaultValueTime,omitempty"`
	DefaultValueTimePrimitiveElement         *primitiveElement    `json:"_defaultValueTime,omitempty"`
	DefaultValueUnsignedInt                  *UnsignedInt         `json:"defaultValueUnsignedInt,omitempty"`
	DefaultValueUnsignedIntPrimitiveElement  *primitiveElement    `json:"_defaultValueUnsignedInt,omitempty"`
	DefaultValueUri                          *Uri                 `json:"defaultValueUri,omitempty"`
	DefaultValueUriPrimitiveElement          *primitiveElement    `json:"_defaultValueUri,omitempty"`
	DefaultValueUrl                          *Url                 `json:"defaultValueUrl,omitempty"`
	DefaultValueUrlPrimitiveElement          *primitiveElement    `json:"_defaultValueUrl,omitempty"`
	DefaultValueUuid                         *Uuid                `json:"defaultValueUuid,omitempty"`
	DefaultValueUuidPrimitiveElement         *primitiveElement    `json:"_defaultValueUuid,omitempty"`
	DefaultValueAddress                      *Address             `json:"defaultValueAddress,omitempty"`
	DefaultValueAge                          *Age                 `json:"defaultValueAge,omitempty"`
	DefaultValueAnnotation                   *Annotation          `json:"defaultValueAnnotation,omitempty"`
	DefaultValueAttachment                   *Attachment          `json:"defaultValueAttachment,omitempty"`
	DefaultValueCodeableConcept              *CodeableConcept     `json:"defaultValueCodeableConcept,omitempty"`
	DefaultValueCoding                       *Coding              `json:"defaultValueCoding,omitempty"`
	DefaultValueContactPoint                 *ContactPoint        `json:"defaultValueContactPoint,omitempty"`
	DefaultValueCount                        *Count               `json:"defaultValueCount,omitempty"`
	DefaultValueDistance                     *Distance            `json:"defaultValueDistance,omitempty"`
	DefaultValueDuration                     *Duration            `json:"defaultValueDuration,omitempty"`
	DefaultValueHumanName                    *HumanName           `json:"defaultValueHumanName,omitempty"`
	DefaultValueIdentifier                   *Identifier          `json:"defaultValueIdentifier,omitempty"`
	DefaultValueMoney                        *Money               `json:"defaultValueMoney,omitempty"`
	DefaultValuePeriod                       *Period              `json:"defaultValuePeriod,omitempty"`
	DefaultValueQuantity                     *Quantity            `json:"defaultValueQuantity,omitempty"`
	DefaultValueRange                        *Range               `json:"defaultValueRange,omitempty"`
	DefaultValueRatio                        *Ratio               `json:"defaultValueRatio,omitempty"`
	DefaultValueReference                    *Reference           `json:"defaultValueReference,omitempty"`
	DefaultValueSampledData                  *SampledData         `json:"defaultValueSampledData,omitempty"`
	DefaultValueSignature                    *Signature           `json:"defaultValueSignature,omitempty"`
	DefaultValueTiming                       *Timing              `json:"defaultValueTiming,omitempty"`
	DefaultValueContactDetail                *ContactDetail       `json:"defaultValueContactDetail,omitempty"`
	DefaultValueContributor                  *Contributor         `json:"defaultValueContributor,omitempty"`
	DefaultValueDataRequirement              *DataRequirement     `json:"defaultValueDataRequirement,omitempty"`
	DefaultValueExpression                   *Expression          `json:"defaultValueExpression,omitempty"`
	DefaultValueParameterDefinition          *ParameterDefinition `json:"defaultValueParameterDefinition,omitempty"`
	DefaultValueRelatedArtifact              *RelatedArtifact     `json:"defaultValueRelatedArtifact,omitempty"`
	DefaultValueTriggerDefinition            *TriggerDefinition   `json:"defaultValueTriggerDefinition,omitempty"`
	DefaultValueUsageContext                 *UsageContext        `json:"defaultValueUsageContext,omitempty"`
	DefaultValueDosage                       *Dosage              `json:"defaultValueDosage,omitempty"`
	DefaultValueMeta                         *Meta                `json:"defaultValueMeta,omitempty"`
	Element                                  *String              `json:"element,omitempty"`
	ElementPrimitiveElement                  *primitiveElement    `json:"_element,omitempty"`
	ListMode                                 *Code                `json:"listMode,omitempty"`
	ListModePrimitiveElement                 *primitiveElement    `json:"_listMode,omitempty"`
	Variable                                 *Id                  `json:"variable,omitempty"`
	VariablePrimitiveElement                 *primitiveElement    `json:"_variable,omitempty"`
	Condition                                *String              `json:"condition,omitempty"`
	ConditionPrimitiveElement                *primitiveElement    `json:"_condition,omitempty"`
	Check                                    *String              `json:"check,omitempty"`
	CheckPrimitiveElement                    *primitiveElement    `json:"_check,omitempty"`
	LogMessage                               *String              `json:"logMessage,omitempty"`
	LogMessagePrimitiveElement               *primitiveElement    `json:"_logMessage,omitempty"`
}

func (r StructureMapGroupRuleSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r StructureMapGroupRuleSource) marshalJSON() jsonStructureMapGroupRuleSource {
	m := jsonStructureMapGroupRuleSource{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Context = r.Context
	if r.Context.Id != nil || r.Context.Extension != nil {
		m.ContextPrimitiveElement = &primitiveElement{Id: r.Context.Id, Extension: r.Context.Extension}
	}
	m.Min = r.Min
	if r.Min != nil && (r.Min.Id != nil || r.Min.Extension != nil) {
		m.MinPrimitiveElement = &primitiveElement{Id: r.Min.Id, Extension: r.Min.Extension}
	}
	m.Max = r.Max
	if r.Max != nil && (r.Max.Id != nil || r.Max.Extension != nil) {
		m.MaxPrimitiveElement = &primitiveElement{Id: r.Max.Id, Extension: r.Max.Extension}
	}
	m.Type = r.Type
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	switch v := r.DefaultValue.(type) {
	case Base64Binary:
		m.DefaultValueBase64Binary = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueBase64BinaryPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Base64Binary:
		m.DefaultValueBase64Binary = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueBase64BinaryPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Boolean:
		m.DefaultValueBoolean = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		m.DefaultValueBoolean = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Canonical:
		m.DefaultValueCanonical = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Canonical:
		m.DefaultValueCanonical = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Code:
		m.DefaultValueCode = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueCodePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Code:
		m.DefaultValueCode = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueCodePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Date:
		m.DefaultValueDate = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		m.DefaultValueDate = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case DateTime:
		m.DefaultValueDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.DefaultValueDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Decimal:
		m.DefaultValueDecimal = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Decimal:
		m.DefaultValueDecimal = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Id:
		m.DefaultValueId = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueIdPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Id:
		m.DefaultValueId = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueIdPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Instant:
		m.DefaultValueInstant = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueInstantPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Instant:
		m.DefaultValueInstant = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueInstantPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Integer:
		m.DefaultValueInteger = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Integer:
		m.DefaultValueInteger = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Markdown:
		m.DefaultValueMarkdown = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueMarkdownPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Markdown:
		m.DefaultValueMarkdown = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueMarkdownPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Oid:
		m.DefaultValueOid = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueOidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Oid:
		m.DefaultValueOid = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueOidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case PositiveInt:
		m.DefaultValuePositiveInt = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValuePositiveIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *PositiveInt:
		m.DefaultValuePositiveInt = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValuePositiveIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case String:
		m.DefaultValueString = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.DefaultValueString = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Time:
		m.DefaultValueTime = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Time:
		m.DefaultValueTime = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case UnsignedInt:
		m.DefaultValueUnsignedInt = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueUnsignedIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *UnsignedInt:
		m.DefaultValueUnsignedInt = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueUnsignedIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Uri:
		m.DefaultValueUri = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uri:
		m.DefaultValueUri = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Url:
		m.DefaultValueUrl = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueUrlPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Url:
		m.DefaultValueUrl = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueUrlPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Uuid:
		m.DefaultValueUuid = &v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueUuidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uuid:
		m.DefaultValueUuid = v
		if v.Id != nil || v.Extension != nil {
			m.DefaultValueUuidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Address:
		m.DefaultValueAddress = &v
	case *Address:
		m.DefaultValueAddress = v
	case Age:
		m.DefaultValueAge = &v
	case *Age:
		m.DefaultValueAge = v
	case Annotation:
		m.DefaultValueAnnotation = &v
	case *Annotation:
		m.DefaultValueAnnotation = v
	case Attachment:
		m.DefaultValueAttachment = &v
	case *Attachment:
		m.DefaultValueAttachment = v
	case CodeableConcept:
		m.DefaultValueCodeableConcept = &v
	case *CodeableConcept:
		m.DefaultValueCodeableConcept = v
	case Coding:
		m.DefaultValueCoding = &v
	case *Coding:
		m.DefaultValueCoding = v
	case ContactPoint:
		m.DefaultValueContactPoint = &v
	case *ContactPoint:
		m.DefaultValueContactPoint = v
	case Count:
		m.DefaultValueCount = &v
	case *Count:
		m.DefaultValueCount = v
	case Distance:
		m.DefaultValueDistance = &v
	case *Distance:
		m.DefaultValueDistance = v
	case Duration:
		m.DefaultValueDuration = &v
	case *Duration:
		m.DefaultValueDuration = v
	case HumanName:
		m.DefaultValueHumanName = &v
	case *HumanName:
		m.DefaultValueHumanName = v
	case Identifier:
		m.DefaultValueIdentifier = &v
	case *Identifier:
		m.DefaultValueIdentifier = v
	case Money:
		m.DefaultValueMoney = &v
	case *Money:
		m.DefaultValueMoney = v
	case Period:
		m.DefaultValuePeriod = &v
	case *Period:
		m.DefaultValuePeriod = v
	case Quantity:
		m.DefaultValueQuantity = &v
	case *Quantity:
		m.DefaultValueQuantity = v
	case Range:
		m.DefaultValueRange = &v
	case *Range:
		m.DefaultValueRange = v
	case Ratio:
		m.DefaultValueRatio = &v
	case *Ratio:
		m.DefaultValueRatio = v
	case Reference:
		m.DefaultValueReference = &v
	case *Reference:
		m.DefaultValueReference = v
	case SampledData:
		m.DefaultValueSampledData = &v
	case *SampledData:
		m.DefaultValueSampledData = v
	case Signature:
		m.DefaultValueSignature = &v
	case *Signature:
		m.DefaultValueSignature = v
	case Timing:
		m.DefaultValueTiming = &v
	case *Timing:
		m.DefaultValueTiming = v
	case ContactDetail:
		m.DefaultValueContactDetail = &v
	case *ContactDetail:
		m.DefaultValueContactDetail = v
	case Contributor:
		m.DefaultValueContributor = &v
	case *Contributor:
		m.DefaultValueContributor = v
	case DataRequirement:
		m.DefaultValueDataRequirement = &v
	case *DataRequirement:
		m.DefaultValueDataRequirement = v
	case Expression:
		m.DefaultValueExpression = &v
	case *Expression:
		m.DefaultValueExpression = v
	case ParameterDefinition:
		m.DefaultValueParameterDefinition = &v
	case *ParameterDefinition:
		m.DefaultValueParameterDefinition = v
	case RelatedArtifact:
		m.DefaultValueRelatedArtifact = &v
	case *RelatedArtifact:
		m.DefaultValueRelatedArtifact = v
	case TriggerDefinition:
		m.DefaultValueTriggerDefinition = &v
	case *TriggerDefinition:
		m.DefaultValueTriggerDefinition = v
	case UsageContext:
		m.DefaultValueUsageContext = &v
	case *UsageContext:
		m.DefaultValueUsageContext = v
	case Dosage:
		m.DefaultValueDosage = &v
	case *Dosage:
		m.DefaultValueDosage = v
	case Meta:
		m.DefaultValueMeta = &v
	case *Meta:
		m.DefaultValueMeta = v
	}
	m.Element = r.Element
	if r.Element != nil && (r.Element.Id != nil || r.Element.Extension != nil) {
		m.ElementPrimitiveElement = &primitiveElement{Id: r.Element.Id, Extension: r.Element.Extension}
	}
	m.ListMode = r.ListMode
	if r.ListMode != nil && (r.ListMode.Id != nil || r.ListMode.Extension != nil) {
		m.ListModePrimitiveElement = &primitiveElement{Id: r.ListMode.Id, Extension: r.ListMode.Extension}
	}
	m.Variable = r.Variable
	if r.Variable != nil && (r.Variable.Id != nil || r.Variable.Extension != nil) {
		m.VariablePrimitiveElement = &primitiveElement{Id: r.Variable.Id, Extension: r.Variable.Extension}
	}
	m.Condition = r.Condition
	if r.Condition != nil && (r.Condition.Id != nil || r.Condition.Extension != nil) {
		m.ConditionPrimitiveElement = &primitiveElement{Id: r.Condition.Id, Extension: r.Condition.Extension}
	}
	m.Check = r.Check
	if r.Check != nil && (r.Check.Id != nil || r.Check.Extension != nil) {
		m.CheckPrimitiveElement = &primitiveElement{Id: r.Check.Id, Extension: r.Check.Extension}
	}
	m.LogMessage = r.LogMessage
	if r.LogMessage != nil && (r.LogMessage.Id != nil || r.LogMessage.Extension != nil) {
		m.LogMessagePrimitiveElement = &primitiveElement{Id: r.LogMessage.Id, Extension: r.LogMessage.Extension}
	}
	return m
}
func (r *StructureMapGroupRuleSource) UnmarshalJSON(b []byte) error {
	var m jsonStructureMapGroupRuleSource
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *StructureMapGroupRuleSource) unmarshalJSON(m jsonStructureMapGroupRuleSource) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Context = m.Context
	if m.ContextPrimitiveElement != nil {
		r.Context.Id = m.ContextPrimitiveElement.Id
		r.Context.Extension = m.ContextPrimitiveElement.Extension
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
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	if m.DefaultValueBase64Binary != nil || m.DefaultValueBase64BinaryPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueBase64Binary
		if m.DefaultValueBase64BinaryPrimitiveElement != nil {
			if v == nil {
				v = &Base64Binary{}
			}
			v.Id = m.DefaultValueBase64BinaryPrimitiveElement.Id
			v.Extension = m.DefaultValueBase64BinaryPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueBoolean != nil || m.DefaultValueBooleanPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueBoolean
		if m.DefaultValueBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.DefaultValueBooleanPrimitiveElement.Id
			v.Extension = m.DefaultValueBooleanPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueCanonical != nil || m.DefaultValueCanonicalPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueCanonical
		if m.DefaultValueCanonicalPrimitiveElement != nil {
			if v == nil {
				v = &Canonical{}
			}
			v.Id = m.DefaultValueCanonicalPrimitiveElement.Id
			v.Extension = m.DefaultValueCanonicalPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueCode != nil || m.DefaultValueCodePrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueCode
		if m.DefaultValueCodePrimitiveElement != nil {
			if v == nil {
				v = &Code{}
			}
			v.Id = m.DefaultValueCodePrimitiveElement.Id
			v.Extension = m.DefaultValueCodePrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueDate != nil || m.DefaultValueDatePrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueDate
		if m.DefaultValueDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.DefaultValueDatePrimitiveElement.Id
			v.Extension = m.DefaultValueDatePrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueDateTime != nil || m.DefaultValueDateTimePrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueDateTime
		if m.DefaultValueDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.DefaultValueDateTimePrimitiveElement.Id
			v.Extension = m.DefaultValueDateTimePrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueDecimal != nil || m.DefaultValueDecimalPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueDecimal
		if m.DefaultValueDecimalPrimitiveElement != nil {
			if v == nil {
				v = &Decimal{}
			}
			v.Id = m.DefaultValueDecimalPrimitiveElement.Id
			v.Extension = m.DefaultValueDecimalPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueId != nil || m.DefaultValueIdPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueId
		if m.DefaultValueIdPrimitiveElement != nil {
			if v == nil {
				v = &Id{}
			}
			v.Id = m.DefaultValueIdPrimitiveElement.Id
			v.Extension = m.DefaultValueIdPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueInstant != nil || m.DefaultValueInstantPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueInstant
		if m.DefaultValueInstantPrimitiveElement != nil {
			if v == nil {
				v = &Instant{}
			}
			v.Id = m.DefaultValueInstantPrimitiveElement.Id
			v.Extension = m.DefaultValueInstantPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueInteger != nil || m.DefaultValueIntegerPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueInteger
		if m.DefaultValueIntegerPrimitiveElement != nil {
			if v == nil {
				v = &Integer{}
			}
			v.Id = m.DefaultValueIntegerPrimitiveElement.Id
			v.Extension = m.DefaultValueIntegerPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueMarkdown != nil || m.DefaultValueMarkdownPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueMarkdown
		if m.DefaultValueMarkdownPrimitiveElement != nil {
			if v == nil {
				v = &Markdown{}
			}
			v.Id = m.DefaultValueMarkdownPrimitiveElement.Id
			v.Extension = m.DefaultValueMarkdownPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueOid != nil || m.DefaultValueOidPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueOid
		if m.DefaultValueOidPrimitiveElement != nil {
			if v == nil {
				v = &Oid{}
			}
			v.Id = m.DefaultValueOidPrimitiveElement.Id
			v.Extension = m.DefaultValueOidPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValuePositiveInt != nil || m.DefaultValuePositiveIntPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValuePositiveInt
		if m.DefaultValuePositiveIntPrimitiveElement != nil {
			if v == nil {
				v = &PositiveInt{}
			}
			v.Id = m.DefaultValuePositiveIntPrimitiveElement.Id
			v.Extension = m.DefaultValuePositiveIntPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueString != nil || m.DefaultValueStringPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueString
		if m.DefaultValueStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.DefaultValueStringPrimitiveElement.Id
			v.Extension = m.DefaultValueStringPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueTime != nil || m.DefaultValueTimePrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueTime
		if m.DefaultValueTimePrimitiveElement != nil {
			if v == nil {
				v = &Time{}
			}
			v.Id = m.DefaultValueTimePrimitiveElement.Id
			v.Extension = m.DefaultValueTimePrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueUnsignedInt != nil || m.DefaultValueUnsignedIntPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueUnsignedInt
		if m.DefaultValueUnsignedIntPrimitiveElement != nil {
			if v == nil {
				v = &UnsignedInt{}
			}
			v.Id = m.DefaultValueUnsignedIntPrimitiveElement.Id
			v.Extension = m.DefaultValueUnsignedIntPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueUri != nil || m.DefaultValueUriPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueUri
		if m.DefaultValueUriPrimitiveElement != nil {
			if v == nil {
				v = &Uri{}
			}
			v.Id = m.DefaultValueUriPrimitiveElement.Id
			v.Extension = m.DefaultValueUriPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueUrl != nil || m.DefaultValueUrlPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueUrl
		if m.DefaultValueUrlPrimitiveElement != nil {
			if v == nil {
				v = &Url{}
			}
			v.Id = m.DefaultValueUrlPrimitiveElement.Id
			v.Extension = m.DefaultValueUrlPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueUuid != nil || m.DefaultValueUuidPrimitiveElement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueUuid
		if m.DefaultValueUuidPrimitiveElement != nil {
			if v == nil {
				v = &Uuid{}
			}
			v.Id = m.DefaultValueUuidPrimitiveElement.Id
			v.Extension = m.DefaultValueUuidPrimitiveElement.Extension
		}
		r.DefaultValue = v
	}
	if m.DefaultValueAddress != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueAddress
		r.DefaultValue = v
	}
	if m.DefaultValueAge != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueAge
		r.DefaultValue = v
	}
	if m.DefaultValueAnnotation != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueAnnotation
		r.DefaultValue = v
	}
	if m.DefaultValueAttachment != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueAttachment
		r.DefaultValue = v
	}
	if m.DefaultValueCodeableConcept != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueCodeableConcept
		r.DefaultValue = v
	}
	if m.DefaultValueCoding != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueCoding
		r.DefaultValue = v
	}
	if m.DefaultValueContactPoint != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueContactPoint
		r.DefaultValue = v
	}
	if m.DefaultValueCount != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueCount
		r.DefaultValue = v
	}
	if m.DefaultValueDistance != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueDistance
		r.DefaultValue = v
	}
	if m.DefaultValueDuration != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueDuration
		r.DefaultValue = v
	}
	if m.DefaultValueHumanName != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueHumanName
		r.DefaultValue = v
	}
	if m.DefaultValueIdentifier != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueIdentifier
		r.DefaultValue = v
	}
	if m.DefaultValueMoney != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueMoney
		r.DefaultValue = v
	}
	if m.DefaultValuePeriod != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValuePeriod
		r.DefaultValue = v
	}
	if m.DefaultValueQuantity != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueQuantity
		r.DefaultValue = v
	}
	if m.DefaultValueRange != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueRange
		r.DefaultValue = v
	}
	if m.DefaultValueRatio != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueRatio
		r.DefaultValue = v
	}
	if m.DefaultValueReference != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueReference
		r.DefaultValue = v
	}
	if m.DefaultValueSampledData != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueSampledData
		r.DefaultValue = v
	}
	if m.DefaultValueSignature != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueSignature
		r.DefaultValue = v
	}
	if m.DefaultValueTiming != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueTiming
		r.DefaultValue = v
	}
	if m.DefaultValueContactDetail != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueContactDetail
		r.DefaultValue = v
	}
	if m.DefaultValueContributor != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueContributor
		r.DefaultValue = v
	}
	if m.DefaultValueDataRequirement != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueDataRequirement
		r.DefaultValue = v
	}
	if m.DefaultValueExpression != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueExpression
		r.DefaultValue = v
	}
	if m.DefaultValueParameterDefinition != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueParameterDefinition
		r.DefaultValue = v
	}
	if m.DefaultValueRelatedArtifact != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueRelatedArtifact
		r.DefaultValue = v
	}
	if m.DefaultValueTriggerDefinition != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueTriggerDefinition
		r.DefaultValue = v
	}
	if m.DefaultValueUsageContext != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueUsageContext
		r.DefaultValue = v
	}
	if m.DefaultValueDosage != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueDosage
		r.DefaultValue = v
	}
	if m.DefaultValueMeta != nil {
		if r.DefaultValue != nil {
			return fmt.Errorf("multiple values for field \"DefaultValue\"")
		}
		v := m.DefaultValueMeta
		r.DefaultValue = v
	}
	r.Element = m.Element
	if m.ElementPrimitiveElement != nil {
		r.Element.Id = m.ElementPrimitiveElement.Id
		r.Element.Extension = m.ElementPrimitiveElement.Extension
	}
	r.ListMode = m.ListMode
	if m.ListModePrimitiveElement != nil {
		r.ListMode.Id = m.ListModePrimitiveElement.Id
		r.ListMode.Extension = m.ListModePrimitiveElement.Extension
	}
	r.Variable = m.Variable
	if m.VariablePrimitiveElement != nil {
		r.Variable.Id = m.VariablePrimitiveElement.Id
		r.Variable.Extension = m.VariablePrimitiveElement.Extension
	}
	r.Condition = m.Condition
	if m.ConditionPrimitiveElement != nil {
		r.Condition.Id = m.ConditionPrimitiveElement.Id
		r.Condition.Extension = m.ConditionPrimitiveElement.Extension
	}
	r.Check = m.Check
	if m.CheckPrimitiveElement != nil {
		r.Check.Id = m.CheckPrimitiveElement.Id
		r.Check.Extension = m.CheckPrimitiveElement.Extension
	}
	r.LogMessage = m.LogMessage
	if m.LogMessagePrimitiveElement != nil {
		r.LogMessage.Id = m.LogMessagePrimitiveElement.Id
		r.LogMessage.Extension = m.LogMessagePrimitiveElement.Extension
	}
	return nil
}
func (r StructureMapGroupRuleSource) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Content to create because of this mapping rule.
type StructureMapGroupRuleTarget struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Type or variable this rule applies to.
	Context *Id
	// How to interpret the context.
	ContextType *Code
	// Field to create in the context.
	Element *String
	// Named context for field, if desired, and a field is specified.
	Variable *Id
	// If field is a list, how to manage the list.
	ListMode []Code
	// Internal rule reference for shared list items.
	ListRuleId *Id
	// How the data is copied / created.
	Transform *Code
	// Parameters to the transform.
	Parameter []StructureMapGroupRuleTargetParameter
}
type jsonStructureMapGroupRuleTarget struct {
	Id                          *string                                `json:"id,omitempty"`
	Extension                   []Extension                            `json:"extension,omitempty"`
	ModifierExtension           []Extension                            `json:"modifierExtension,omitempty"`
	Context                     *Id                                    `json:"context,omitempty"`
	ContextPrimitiveElement     *primitiveElement                      `json:"_context,omitempty"`
	ContextType                 *Code                                  `json:"contextType,omitempty"`
	ContextTypePrimitiveElement *primitiveElement                      `json:"_contextType,omitempty"`
	Element                     *String                                `json:"element,omitempty"`
	ElementPrimitiveElement     *primitiveElement                      `json:"_element,omitempty"`
	Variable                    *Id                                    `json:"variable,omitempty"`
	VariablePrimitiveElement    *primitiveElement                      `json:"_variable,omitempty"`
	ListMode                    []Code                                 `json:"listMode,omitempty"`
	ListModePrimitiveElement    []*primitiveElement                    `json:"_listMode,omitempty"`
	ListRuleId                  *Id                                    `json:"listRuleId,omitempty"`
	ListRuleIdPrimitiveElement  *primitiveElement                      `json:"_listRuleId,omitempty"`
	Transform                   *Code                                  `json:"transform,omitempty"`
	TransformPrimitiveElement   *primitiveElement                      `json:"_transform,omitempty"`
	Parameter                   []StructureMapGroupRuleTargetParameter `json:"parameter,omitempty"`
}

func (r StructureMapGroupRuleTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r StructureMapGroupRuleTarget) marshalJSON() jsonStructureMapGroupRuleTarget {
	m := jsonStructureMapGroupRuleTarget{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Context = r.Context
	if r.Context != nil && (r.Context.Id != nil || r.Context.Extension != nil) {
		m.ContextPrimitiveElement = &primitiveElement{Id: r.Context.Id, Extension: r.Context.Extension}
	}
	m.ContextType = r.ContextType
	if r.ContextType != nil && (r.ContextType.Id != nil || r.ContextType.Extension != nil) {
		m.ContextTypePrimitiveElement = &primitiveElement{Id: r.ContextType.Id, Extension: r.ContextType.Extension}
	}
	m.Element = r.Element
	if r.Element != nil && (r.Element.Id != nil || r.Element.Extension != nil) {
		m.ElementPrimitiveElement = &primitiveElement{Id: r.Element.Id, Extension: r.Element.Extension}
	}
	m.Variable = r.Variable
	if r.Variable != nil && (r.Variable.Id != nil || r.Variable.Extension != nil) {
		m.VariablePrimitiveElement = &primitiveElement{Id: r.Variable.Id, Extension: r.Variable.Extension}
	}
	m.ListMode = r.ListMode
	anyListModeIdOrExtension := false
	for _, e := range r.ListMode {
		if e.Id != nil || e.Extension != nil {
			anyListModeIdOrExtension = true
			break
		}
	}
	if anyListModeIdOrExtension {
		m.ListModePrimitiveElement = make([]*primitiveElement, 0, len(r.ListMode))
		for _, e := range r.ListMode {
			if e.Id != nil || e.Extension != nil {
				m.ListModePrimitiveElement = append(m.ListModePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ListModePrimitiveElement = append(m.ListModePrimitiveElement, nil)
			}
		}
	}
	m.ListRuleId = r.ListRuleId
	if r.ListRuleId != nil && (r.ListRuleId.Id != nil || r.ListRuleId.Extension != nil) {
		m.ListRuleIdPrimitiveElement = &primitiveElement{Id: r.ListRuleId.Id, Extension: r.ListRuleId.Extension}
	}
	m.Transform = r.Transform
	if r.Transform != nil && (r.Transform.Id != nil || r.Transform.Extension != nil) {
		m.TransformPrimitiveElement = &primitiveElement{Id: r.Transform.Id, Extension: r.Transform.Extension}
	}
	m.Parameter = r.Parameter
	return m
}
func (r *StructureMapGroupRuleTarget) UnmarshalJSON(b []byte) error {
	var m jsonStructureMapGroupRuleTarget
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *StructureMapGroupRuleTarget) unmarshalJSON(m jsonStructureMapGroupRuleTarget) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Context = m.Context
	if m.ContextPrimitiveElement != nil {
		r.Context.Id = m.ContextPrimitiveElement.Id
		r.Context.Extension = m.ContextPrimitiveElement.Extension
	}
	r.ContextType = m.ContextType
	if m.ContextTypePrimitiveElement != nil {
		r.ContextType.Id = m.ContextTypePrimitiveElement.Id
		r.ContextType.Extension = m.ContextTypePrimitiveElement.Extension
	}
	r.Element = m.Element
	if m.ElementPrimitiveElement != nil {
		r.Element.Id = m.ElementPrimitiveElement.Id
		r.Element.Extension = m.ElementPrimitiveElement.Extension
	}
	r.Variable = m.Variable
	if m.VariablePrimitiveElement != nil {
		r.Variable.Id = m.VariablePrimitiveElement.Id
		r.Variable.Extension = m.VariablePrimitiveElement.Extension
	}
	r.ListMode = m.ListMode
	for i, e := range m.ListModePrimitiveElement {
		if len(r.ListMode) > i {
			r.ListMode[i].Id = e.Id
			r.ListMode[i].Extension = e.Extension
		} else {
			r.ListMode = append(r.ListMode, Code{Id: e.Id, Extension: e.Extension})
		}
	}
	r.ListRuleId = m.ListRuleId
	if m.ListRuleIdPrimitiveElement != nil {
		r.ListRuleId.Id = m.ListRuleIdPrimitiveElement.Id
		r.ListRuleId.Extension = m.ListRuleIdPrimitiveElement.Extension
	}
	r.Transform = m.Transform
	if m.TransformPrimitiveElement != nil {
		r.Transform.Id = m.TransformPrimitiveElement.Id
		r.Transform.Extension = m.TransformPrimitiveElement.Extension
	}
	r.Parameter = m.Parameter
	return nil
}
func (r StructureMapGroupRuleTarget) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Parameters to the transform.
type StructureMapGroupRuleTargetParameter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Parameter value - variable or literal.
	Value isStructureMapGroupRuleTargetParameterValue
}
type isStructureMapGroupRuleTargetParameterValue interface {
	isStructureMapGroupRuleTargetParameterValue()
}

func (r Id) isStructureMapGroupRuleTargetParameterValue()      {}
func (r String) isStructureMapGroupRuleTargetParameterValue()  {}
func (r Boolean) isStructureMapGroupRuleTargetParameterValue() {}
func (r Integer) isStructureMapGroupRuleTargetParameterValue() {}
func (r Decimal) isStructureMapGroupRuleTargetParameterValue() {}

type jsonStructureMapGroupRuleTargetParameter struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	ValueId                      *Id               `json:"valueId,omitempty"`
	ValueIdPrimitiveElement      *primitiveElement `json:"_valueId,omitempty"`
	ValueString                  *String           `json:"valueString,omitempty"`
	ValueStringPrimitiveElement  *primitiveElement `json:"_valueString,omitempty"`
	ValueBoolean                 *Boolean          `json:"valueBoolean,omitempty"`
	ValueBooleanPrimitiveElement *primitiveElement `json:"_valueBoolean,omitempty"`
	ValueInteger                 *Integer          `json:"valueInteger,omitempty"`
	ValueIntegerPrimitiveElement *primitiveElement `json:"_valueInteger,omitempty"`
	ValueDecimal                 *Decimal          `json:"valueDecimal,omitempty"`
	ValueDecimalPrimitiveElement *primitiveElement `json:"_valueDecimal,omitempty"`
}

func (r StructureMapGroupRuleTargetParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r StructureMapGroupRuleTargetParameter) marshalJSON() jsonStructureMapGroupRuleTargetParameter {
	m := jsonStructureMapGroupRuleTargetParameter{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Value.(type) {
	case Id:
		m.ValueId = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueIdPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Id:
		m.ValueId = v
		if v.Id != nil || v.Extension != nil {
			m.ValueIdPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case String:
		m.ValueString = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.ValueString = v
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Boolean:
		m.ValueBoolean = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		m.ValueBoolean = v
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Integer:
		m.ValueInteger = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Integer:
		m.ValueInteger = v
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Decimal:
		m.ValueDecimal = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Decimal:
		m.ValueDecimal = v
		if v.Id != nil || v.Extension != nil {
			m.ValueDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	return m
}
func (r *StructureMapGroupRuleTargetParameter) UnmarshalJSON(b []byte) error {
	var m jsonStructureMapGroupRuleTargetParameter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *StructureMapGroupRuleTargetParameter) unmarshalJSON(m jsonStructureMapGroupRuleTargetParameter) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.ValueId != nil || m.ValueIdPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueId
		if m.ValueIdPrimitiveElement != nil {
			if v == nil {
				v = &Id{}
			}
			v.Id = m.ValueIdPrimitiveElement.Id
			v.Extension = m.ValueIdPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueString != nil || m.ValueStringPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueString
		if m.ValueStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.ValueStringPrimitiveElement.Id
			v.Extension = m.ValueStringPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueBoolean != nil || m.ValueBooleanPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueBoolean
		if m.ValueBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.ValueBooleanPrimitiveElement.Id
			v.Extension = m.ValueBooleanPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueInteger != nil || m.ValueIntegerPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueInteger
		if m.ValueIntegerPrimitiveElement != nil {
			if v == nil {
				v = &Integer{}
			}
			v.Id = m.ValueIntegerPrimitiveElement.Id
			v.Extension = m.ValueIntegerPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueDecimal != nil || m.ValueDecimalPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDecimal
		if m.ValueDecimalPrimitiveElement != nil {
			if v == nil {
				v = &Decimal{}
			}
			v.Id = m.ValueDecimalPrimitiveElement.Id
			v.Extension = m.ValueDecimalPrimitiveElement.Extension
		}
		r.Value = v
	}
	return nil
}
func (r StructureMapGroupRuleTargetParameter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Which other rules to apply in the context of this rule.
type StructureMapGroupRuleDependent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Name of a rule or group to apply.
	Name Id
	// Variable to pass to the rule or group.
	Variable []String
}
type jsonStructureMapGroupRuleDependent struct {
	Id                       *string             `json:"id,omitempty"`
	Extension                []Extension         `json:"extension,omitempty"`
	ModifierExtension        []Extension         `json:"modifierExtension,omitempty"`
	Name                     Id                  `json:"name,omitempty"`
	NamePrimitiveElement     *primitiveElement   `json:"_name,omitempty"`
	Variable                 []String            `json:"variable,omitempty"`
	VariablePrimitiveElement []*primitiveElement `json:"_variable,omitempty"`
}

func (r StructureMapGroupRuleDependent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r StructureMapGroupRuleDependent) marshalJSON() jsonStructureMapGroupRuleDependent {
	m := jsonStructureMapGroupRuleDependent{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Name = r.Name
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Variable = r.Variable
	anyVariableIdOrExtension := false
	for _, e := range r.Variable {
		if e.Id != nil || e.Extension != nil {
			anyVariableIdOrExtension = true
			break
		}
	}
	if anyVariableIdOrExtension {
		m.VariablePrimitiveElement = make([]*primitiveElement, 0, len(r.Variable))
		for _, e := range r.Variable {
			if e.Id != nil || e.Extension != nil {
				m.VariablePrimitiveElement = append(m.VariablePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.VariablePrimitiveElement = append(m.VariablePrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *StructureMapGroupRuleDependent) UnmarshalJSON(b []byte) error {
	var m jsonStructureMapGroupRuleDependent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *StructureMapGroupRuleDependent) unmarshalJSON(m jsonStructureMapGroupRuleDependent) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Variable = m.Variable
	for i, e := range m.VariablePrimitiveElement {
		if len(r.Variable) > i {
			r.Variable[i].Id = e.Id
			r.Variable[i].Extension = e.Extension
		} else {
			r.Variable = append(r.Variable, String{Id: e.Id, Extension: e.Extension})
		}
	}
	return nil
}
func (r StructureMapGroupRuleDependent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
