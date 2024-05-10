package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// The CodeSystem resource is used to declare the existence of and describe a code system or code system supplement and its key properties, and optionally define a part or all of its content.
type CodeSystem struct {
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
	// An absolute URI that is used to identify this code system when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this code system is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the code system is stored on different servers. This is used in [Coding](datatypes.html#Coding).system.
	Url *Uri
	// A formal identifier that is used to identify this code system when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the code system when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the code system author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. This is used in [Coding](datatypes.html#Coding).version.
	Version *String
	// A natural language name identifying the code system. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the code system.
	Title *String
	// The date (and optionally time) when the code system resource was created or revised.
	Status Code
	// A Boolean value to indicate that this code system is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the code system was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the code system changes.
	Date *DateTime
	// The name of the organization or individual that published the code system.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the code system from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate code system instances.
	UseContext []UsageContext
	// A legal or geographic region in which the code system is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this code system is needed and why it has been designed as it has.
	Purpose *Markdown
	// A copyright statement relating to the code system and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the code system.
	Copyright *Markdown
	// If code comparison is case sensitive when codes within this system are compared to each other.
	CaseSensitive *Boolean
	// Canonical reference to the value set that contains the entire code system.
	ValueSet *Canonical
	// The meaning of the hierarchy of concepts as represented in this resource.
	HierarchyMeaning *Code
	// The code system defines a compositional (post-coordination) grammar.
	Compositional *Boolean
	// This flag is used to signify that the code system does not commit to concept permanence across versions. If true, a version must be specified when referencing this code system.
	VersionNeeded *Boolean
	// The extent of the content of the code system (the concepts and codes it defines) are represented in this resource instance.
	Content Code
	// The canonical URL of the code system that this code system supplement is adding designations and properties to.
	Supplements *Canonical
	// The total number of concepts defined by the code system. Where the code system has a compositional grammar, the basis of this count is defined by the system steward.
	Count *UnsignedInt
	// A filter that can be used in a value set compose statement when selecting concepts using a filter.
	Filter []CodeSystemFilter
	// A property defines an additional slot through which additional information can be provided about a concept.
	Property []CodeSystemProperty
	// Concepts that are in the code system. The concept definitions are inherently hierarchical, but the definitions must be consulted to determine what the meanings of the hierarchical relationships are.
	Concept []CodeSystemConcept
}
type jsonCodeSystem struct {
	ResourceType                     string               `json:"resourceType"`
	Id                               *Id                  `json:"id,omitempty"`
	IdPrimitiveElement               *primitiveElement    `json:"_id,omitempty"`
	Meta                             *Meta                `json:"meta,omitempty"`
	ImplicitRules                    *Uri                 `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement    *primitiveElement    `json:"_implicitRules,omitempty"`
	Language                         *Code                `json:"language,omitempty"`
	LanguagePrimitiveElement         *primitiveElement    `json:"_language,omitempty"`
	Text                             *Narrative           `json:"text,omitempty"`
	Contained                        []containedResource  `json:"contained,omitempty"`
	Extension                        []Extension          `json:"extension,omitempty"`
	ModifierExtension                []Extension          `json:"modifierExtension,omitempty"`
	Url                              *Uri                 `json:"url,omitempty"`
	UrlPrimitiveElement              *primitiveElement    `json:"_url,omitempty"`
	Identifier                       []Identifier         `json:"identifier,omitempty"`
	Version                          *String              `json:"version,omitempty"`
	VersionPrimitiveElement          *primitiveElement    `json:"_version,omitempty"`
	Name                             *String              `json:"name,omitempty"`
	NamePrimitiveElement             *primitiveElement    `json:"_name,omitempty"`
	Title                            *String              `json:"title,omitempty"`
	TitlePrimitiveElement            *primitiveElement    `json:"_title,omitempty"`
	Status                           Code                 `json:"status,omitempty"`
	StatusPrimitiveElement           *primitiveElement    `json:"_status,omitempty"`
	Experimental                     *Boolean             `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement     *primitiveElement    `json:"_experimental,omitempty"`
	Date                             *DateTime            `json:"date,omitempty"`
	DatePrimitiveElement             *primitiveElement    `json:"_date,omitempty"`
	Publisher                        *String              `json:"publisher,omitempty"`
	PublisherPrimitiveElement        *primitiveElement    `json:"_publisher,omitempty"`
	Contact                          []ContactDetail      `json:"contact,omitempty"`
	Description                      *Markdown            `json:"description,omitempty"`
	DescriptionPrimitiveElement      *primitiveElement    `json:"_description,omitempty"`
	UseContext                       []UsageContext       `json:"useContext,omitempty"`
	Jurisdiction                     []CodeableConcept    `json:"jurisdiction,omitempty"`
	Purpose                          *Markdown            `json:"purpose,omitempty"`
	PurposePrimitiveElement          *primitiveElement    `json:"_purpose,omitempty"`
	Copyright                        *Markdown            `json:"copyright,omitempty"`
	CopyrightPrimitiveElement        *primitiveElement    `json:"_copyright,omitempty"`
	CaseSensitive                    *Boolean             `json:"caseSensitive,omitempty"`
	CaseSensitivePrimitiveElement    *primitiveElement    `json:"_caseSensitive,omitempty"`
	ValueSet                         *Canonical           `json:"valueSet,omitempty"`
	ValueSetPrimitiveElement         *primitiveElement    `json:"_valueSet,omitempty"`
	HierarchyMeaning                 *Code                `json:"hierarchyMeaning,omitempty"`
	HierarchyMeaningPrimitiveElement *primitiveElement    `json:"_hierarchyMeaning,omitempty"`
	Compositional                    *Boolean             `json:"compositional,omitempty"`
	CompositionalPrimitiveElement    *primitiveElement    `json:"_compositional,omitempty"`
	VersionNeeded                    *Boolean             `json:"versionNeeded,omitempty"`
	VersionNeededPrimitiveElement    *primitiveElement    `json:"_versionNeeded,omitempty"`
	Content                          Code                 `json:"content,omitempty"`
	ContentPrimitiveElement          *primitiveElement    `json:"_content,omitempty"`
	Supplements                      *Canonical           `json:"supplements,omitempty"`
	SupplementsPrimitiveElement      *primitiveElement    `json:"_supplements,omitempty"`
	Count                            *UnsignedInt         `json:"count,omitempty"`
	CountPrimitiveElement            *primitiveElement    `json:"_count,omitempty"`
	Filter                           []CodeSystemFilter   `json:"filter,omitempty"`
	Property                         []CodeSystemProperty `json:"property,omitempty"`
	Concept                          []CodeSystemConcept  `json:"concept,omitempty"`
}

func (r CodeSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CodeSystem) marshalJSON() jsonCodeSystem {
	m := jsonCodeSystem{}
	m.ResourceType = "CodeSystem"
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
	m.CaseSensitive = r.CaseSensitive
	if r.CaseSensitive != nil && (r.CaseSensitive.Id != nil || r.CaseSensitive.Extension != nil) {
		m.CaseSensitivePrimitiveElement = &primitiveElement{Id: r.CaseSensitive.Id, Extension: r.CaseSensitive.Extension}
	}
	m.ValueSet = r.ValueSet
	if r.ValueSet != nil && (r.ValueSet.Id != nil || r.ValueSet.Extension != nil) {
		m.ValueSetPrimitiveElement = &primitiveElement{Id: r.ValueSet.Id, Extension: r.ValueSet.Extension}
	}
	m.HierarchyMeaning = r.HierarchyMeaning
	if r.HierarchyMeaning != nil && (r.HierarchyMeaning.Id != nil || r.HierarchyMeaning.Extension != nil) {
		m.HierarchyMeaningPrimitiveElement = &primitiveElement{Id: r.HierarchyMeaning.Id, Extension: r.HierarchyMeaning.Extension}
	}
	m.Compositional = r.Compositional
	if r.Compositional != nil && (r.Compositional.Id != nil || r.Compositional.Extension != nil) {
		m.CompositionalPrimitiveElement = &primitiveElement{Id: r.Compositional.Id, Extension: r.Compositional.Extension}
	}
	m.VersionNeeded = r.VersionNeeded
	if r.VersionNeeded != nil && (r.VersionNeeded.Id != nil || r.VersionNeeded.Extension != nil) {
		m.VersionNeededPrimitiveElement = &primitiveElement{Id: r.VersionNeeded.Id, Extension: r.VersionNeeded.Extension}
	}
	m.Content = r.Content
	if r.Content.Id != nil || r.Content.Extension != nil {
		m.ContentPrimitiveElement = &primitiveElement{Id: r.Content.Id, Extension: r.Content.Extension}
	}
	m.Supplements = r.Supplements
	if r.Supplements != nil && (r.Supplements.Id != nil || r.Supplements.Extension != nil) {
		m.SupplementsPrimitiveElement = &primitiveElement{Id: r.Supplements.Id, Extension: r.Supplements.Extension}
	}
	m.Count = r.Count
	if r.Count != nil && (r.Count.Id != nil || r.Count.Extension != nil) {
		m.CountPrimitiveElement = &primitiveElement{Id: r.Count.Id, Extension: r.Count.Extension}
	}
	m.Filter = r.Filter
	m.Property = r.Property
	m.Concept = r.Concept
	return m
}
func (r *CodeSystem) UnmarshalJSON(b []byte) error {
	var m jsonCodeSystem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CodeSystem) unmarshalJSON(m jsonCodeSystem) error {
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
	r.CaseSensitive = m.CaseSensitive
	if m.CaseSensitivePrimitiveElement != nil {
		r.CaseSensitive.Id = m.CaseSensitivePrimitiveElement.Id
		r.CaseSensitive.Extension = m.CaseSensitivePrimitiveElement.Extension
	}
	r.ValueSet = m.ValueSet
	if m.ValueSetPrimitiveElement != nil {
		r.ValueSet.Id = m.ValueSetPrimitiveElement.Id
		r.ValueSet.Extension = m.ValueSetPrimitiveElement.Extension
	}
	r.HierarchyMeaning = m.HierarchyMeaning
	if m.HierarchyMeaningPrimitiveElement != nil {
		r.HierarchyMeaning.Id = m.HierarchyMeaningPrimitiveElement.Id
		r.HierarchyMeaning.Extension = m.HierarchyMeaningPrimitiveElement.Extension
	}
	r.Compositional = m.Compositional
	if m.CompositionalPrimitiveElement != nil {
		r.Compositional.Id = m.CompositionalPrimitiveElement.Id
		r.Compositional.Extension = m.CompositionalPrimitiveElement.Extension
	}
	r.VersionNeeded = m.VersionNeeded
	if m.VersionNeededPrimitiveElement != nil {
		r.VersionNeeded.Id = m.VersionNeededPrimitiveElement.Id
		r.VersionNeeded.Extension = m.VersionNeededPrimitiveElement.Extension
	}
	r.Content = m.Content
	if m.ContentPrimitiveElement != nil {
		r.Content.Id = m.ContentPrimitiveElement.Id
		r.Content.Extension = m.ContentPrimitiveElement.Extension
	}
	r.Supplements = m.Supplements
	if m.SupplementsPrimitiveElement != nil {
		r.Supplements.Id = m.SupplementsPrimitiveElement.Id
		r.Supplements.Extension = m.SupplementsPrimitiveElement.Extension
	}
	r.Count = m.Count
	if m.CountPrimitiveElement != nil {
		r.Count.Id = m.CountPrimitiveElement.Id
		r.Count.Extension = m.CountPrimitiveElement.Extension
	}
	r.Filter = m.Filter
	r.Property = m.Property
	r.Concept = m.Concept
	return nil
}
func (r CodeSystem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A filter that can be used in a value set compose statement when selecting concepts using a filter.
type CodeSystemFilter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The code that identifies this filter when it is used as a filter in [ValueSet](valueset.html#).compose.include.filter.
	Code Code
	// A description of how or why the filter is used.
	Description *String
	// A list of operators that can be used with the filter.
	Operator []Code
	// A description of what the value for the filter should be.
	Value String
}
type jsonCodeSystemFilter struct {
	Id                          *string             `json:"id,omitempty"`
	Extension                   []Extension         `json:"extension,omitempty"`
	ModifierExtension           []Extension         `json:"modifierExtension,omitempty"`
	Code                        Code                `json:"code,omitempty"`
	CodePrimitiveElement        *primitiveElement   `json:"_code,omitempty"`
	Description                 *String             `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement   `json:"_description,omitempty"`
	Operator                    []Code              `json:"operator,omitempty"`
	OperatorPrimitiveElement    []*primitiveElement `json:"_operator,omitempty"`
	Value                       String              `json:"value,omitempty"`
	ValuePrimitiveElement       *primitiveElement   `json:"_value,omitempty"`
}

func (r CodeSystemFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CodeSystemFilter) marshalJSON() jsonCodeSystemFilter {
	m := jsonCodeSystemFilter{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Operator = r.Operator
	anyOperatorIdOrExtension := false
	for _, e := range r.Operator {
		if e.Id != nil || e.Extension != nil {
			anyOperatorIdOrExtension = true
			break
		}
	}
	if anyOperatorIdOrExtension {
		m.OperatorPrimitiveElement = make([]*primitiveElement, 0, len(r.Operator))
		for _, e := range r.Operator {
			if e.Id != nil || e.Extension != nil {
				m.OperatorPrimitiveElement = append(m.OperatorPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.OperatorPrimitiveElement = append(m.OperatorPrimitiveElement, nil)
			}
		}
	}
	m.Value = r.Value
	if r.Value.Id != nil || r.Value.Extension != nil {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	return m
}
func (r *CodeSystemFilter) UnmarshalJSON(b []byte) error {
	var m jsonCodeSystemFilter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CodeSystemFilter) unmarshalJSON(m jsonCodeSystemFilter) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Operator = m.Operator
	for i, e := range m.OperatorPrimitiveElement {
		if len(r.Operator) > i {
			r.Operator[i].Id = e.Id
			r.Operator[i].Extension = e.Extension
		} else {
			r.Operator = append(r.Operator, Code{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	return nil
}
func (r CodeSystemFilter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A property defines an additional slot through which additional information can be provided about a concept.
type CodeSystemProperty struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code that is used to identify the property. The code is used internally (in CodeSystem.concept.property.code) and also externally, such as in property filters.
	Code Code
	// Reference to the formal meaning of the property. One possible source of meaning is the [Concept Properties](codesystem-concept-properties.html) code system.
	Uri *Uri
	// A description of the property- why it is defined, and how its value might be used.
	Description *String
	// The type of the property value. Properties of type "code" contain a code defined by the code system (e.g. a reference to another defined concept).
	Type Code
}
type jsonCodeSystemProperty struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Code                        Code              `json:"code,omitempty"`
	CodePrimitiveElement        *primitiveElement `json:"_code,omitempty"`
	Uri                         *Uri              `json:"uri,omitempty"`
	UriPrimitiveElement         *primitiveElement `json:"_uri,omitempty"`
	Description                 *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
	Type                        Code              `json:"type,omitempty"`
	TypePrimitiveElement        *primitiveElement `json:"_type,omitempty"`
}

func (r CodeSystemProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CodeSystemProperty) marshalJSON() jsonCodeSystemProperty {
	m := jsonCodeSystemProperty{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Uri = r.Uri
	if r.Uri != nil && (r.Uri.Id != nil || r.Uri.Extension != nil) {
		m.UriPrimitiveElement = &primitiveElement{Id: r.Uri.Id, Extension: r.Uri.Extension}
	}
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	return m
}
func (r *CodeSystemProperty) UnmarshalJSON(b []byte) error {
	var m jsonCodeSystemProperty
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CodeSystemProperty) unmarshalJSON(m jsonCodeSystemProperty) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Uri = m.Uri
	if m.UriPrimitiveElement != nil {
		r.Uri.Id = m.UriPrimitiveElement.Id
		r.Uri.Extension = m.UriPrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	return nil
}
func (r CodeSystemProperty) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Concepts that are in the code system. The concept definitions are inherently hierarchical, but the definitions must be consulted to determine what the meanings of the hierarchical relationships are.
type CodeSystemConcept struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code - a text symbol - that uniquely identifies the concept within the code system.
	Code Code
	// A human readable string that is the recommended default way to present this concept to a user.
	Display *String
	// The formal definition of the concept. The code system resource does not make formal definitions required, because of the prevalence of legacy systems. However, they are highly recommended, as without them there is no formal meaning associated with the concept.
	Definition *String
	// Additional representations for the concept - other languages, aliases, specialized purposes, used for particular purposes, etc.
	Designation []CodeSystemConceptDesignation
	// A property value for this concept.
	Property []CodeSystemConceptProperty
	// Defines children of a concept to produce a hierarchy of concepts. The nature of the relationships is variable (is-a/contains/categorizes) - see hierarchyMeaning.
	Concept []CodeSystemConcept
}
type jsonCodeSystemConcept struct {
	Id                         *string                        `json:"id,omitempty"`
	Extension                  []Extension                    `json:"extension,omitempty"`
	ModifierExtension          []Extension                    `json:"modifierExtension,omitempty"`
	Code                       Code                           `json:"code,omitempty"`
	CodePrimitiveElement       *primitiveElement              `json:"_code,omitempty"`
	Display                    *String                        `json:"display,omitempty"`
	DisplayPrimitiveElement    *primitiveElement              `json:"_display,omitempty"`
	Definition                 *String                        `json:"definition,omitempty"`
	DefinitionPrimitiveElement *primitiveElement              `json:"_definition,omitempty"`
	Designation                []CodeSystemConceptDesignation `json:"designation,omitempty"`
	Property                   []CodeSystemConceptProperty    `json:"property,omitempty"`
	Concept                    []CodeSystemConcept            `json:"concept,omitempty"`
}

func (r CodeSystemConcept) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CodeSystemConcept) marshalJSON() jsonCodeSystemConcept {
	m := jsonCodeSystemConcept{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Display = r.Display
	if r.Display != nil && (r.Display.Id != nil || r.Display.Extension != nil) {
		m.DisplayPrimitiveElement = &primitiveElement{Id: r.Display.Id, Extension: r.Display.Extension}
	}
	m.Definition = r.Definition
	if r.Definition != nil && (r.Definition.Id != nil || r.Definition.Extension != nil) {
		m.DefinitionPrimitiveElement = &primitiveElement{Id: r.Definition.Id, Extension: r.Definition.Extension}
	}
	m.Designation = r.Designation
	m.Property = r.Property
	m.Concept = r.Concept
	return m
}
func (r *CodeSystemConcept) UnmarshalJSON(b []byte) error {
	var m jsonCodeSystemConcept
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CodeSystemConcept) unmarshalJSON(m jsonCodeSystemConcept) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Display = m.Display
	if m.DisplayPrimitiveElement != nil {
		r.Display.Id = m.DisplayPrimitiveElement.Id
		r.Display.Extension = m.DisplayPrimitiveElement.Extension
	}
	r.Definition = m.Definition
	if m.DefinitionPrimitiveElement != nil {
		r.Definition.Id = m.DefinitionPrimitiveElement.Id
		r.Definition.Extension = m.DefinitionPrimitiveElement.Extension
	}
	r.Designation = m.Designation
	r.Property = m.Property
	r.Concept = m.Concept
	return nil
}
func (r CodeSystemConcept) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Additional representations for the concept - other languages, aliases, specialized purposes, used for particular purposes, etc.
type CodeSystemConceptDesignation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The language this designation is defined for.
	Language *Code
	// A code that details how this designation would be used.
	Use *Coding
	// The text value for this designation.
	Value String
}
type jsonCodeSystemConceptDesignation struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Language                 *Code             `json:"language,omitempty"`
	LanguagePrimitiveElement *primitiveElement `json:"_language,omitempty"`
	Use                      *Coding           `json:"use,omitempty"`
	Value                    String            `json:"value,omitempty"`
	ValuePrimitiveElement    *primitiveElement `json:"_value,omitempty"`
}

func (r CodeSystemConceptDesignation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CodeSystemConceptDesignation) marshalJSON() jsonCodeSystemConceptDesignation {
	m := jsonCodeSystemConceptDesignation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Language = r.Language
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.Use = r.Use
	m.Value = r.Value
	if r.Value.Id != nil || r.Value.Extension != nil {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	return m
}
func (r *CodeSystemConceptDesignation) UnmarshalJSON(b []byte) error {
	var m jsonCodeSystemConceptDesignation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CodeSystemConceptDesignation) unmarshalJSON(m jsonCodeSystemConceptDesignation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Use = m.Use
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	return nil
}
func (r CodeSystemConceptDesignation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A property value for this concept.
type CodeSystemConceptProperty struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code that is a reference to CodeSystem.property.code.
	Code Code
	// The value of this property.
	Value isCodeSystemConceptPropertyValue
}
type isCodeSystemConceptPropertyValue interface {
	isCodeSystemConceptPropertyValue()
}

func (r Code) isCodeSystemConceptPropertyValue()     {}
func (r Coding) isCodeSystemConceptPropertyValue()   {}
func (r String) isCodeSystemConceptPropertyValue()   {}
func (r Integer) isCodeSystemConceptPropertyValue()  {}
func (r Boolean) isCodeSystemConceptPropertyValue()  {}
func (r DateTime) isCodeSystemConceptPropertyValue() {}
func (r Decimal) isCodeSystemConceptPropertyValue()  {}

type jsonCodeSystemConceptProperty struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	Code                          Code              `json:"code,omitempty"`
	CodePrimitiveElement          *primitiveElement `json:"_code,omitempty"`
	ValueCode                     *Code             `json:"valueCode,omitempty"`
	ValueCodePrimitiveElement     *primitiveElement `json:"_valueCode,omitempty"`
	ValueCoding                   *Coding           `json:"valueCoding,omitempty"`
	ValueString                   *String           `json:"valueString,omitempty"`
	ValueStringPrimitiveElement   *primitiveElement `json:"_valueString,omitempty"`
	ValueInteger                  *Integer          `json:"valueInteger,omitempty"`
	ValueIntegerPrimitiveElement  *primitiveElement `json:"_valueInteger,omitempty"`
	ValueBoolean                  *Boolean          `json:"valueBoolean,omitempty"`
	ValueBooleanPrimitiveElement  *primitiveElement `json:"_valueBoolean,omitempty"`
	ValueDateTime                 *DateTime         `json:"valueDateTime,omitempty"`
	ValueDateTimePrimitiveElement *primitiveElement `json:"_valueDateTime,omitempty"`
	ValueDecimal                  *Decimal          `json:"valueDecimal,omitempty"`
	ValueDecimalPrimitiveElement  *primitiveElement `json:"_valueDecimal,omitempty"`
}

func (r CodeSystemConceptProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CodeSystemConceptProperty) marshalJSON() jsonCodeSystemConceptProperty {
	m := jsonCodeSystemConceptProperty{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	switch v := r.Value.(type) {
	case Code:
		m.ValueCode = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueCodePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Code:
		m.ValueCode = v
		if v.Id != nil || v.Extension != nil {
			m.ValueCodePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Coding:
		m.ValueCoding = &v
	case *Coding:
		m.ValueCoding = v
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
	case DateTime:
		m.ValueDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.ValueDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.ValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
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
func (r *CodeSystemConceptProperty) UnmarshalJSON(b []byte) error {
	var m jsonCodeSystemConceptProperty
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CodeSystemConceptProperty) unmarshalJSON(m jsonCodeSystemConceptProperty) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	if m.ValueCode != nil || m.ValueCodePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCode
		if m.ValueCodePrimitiveElement != nil {
			if v == nil {
				v = &Code{}
			}
			v.Id = m.ValueCodePrimitiveElement.Id
			v.Extension = m.ValueCodePrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueCoding != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCoding
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
	if m.ValueDateTime != nil || m.ValueDateTimePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDateTime
		if m.ValueDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.ValueDateTimePrimitiveElement.Id
			v.Extension = m.ValueDateTimePrimitiveElement.Extension
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
func (r CodeSystemConceptProperty) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
