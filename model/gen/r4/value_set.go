package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A ValueSet resource instance specifies a set of codes drawn from one or more code systems, intended for use in a particular context. Value sets link between [[[CodeSystem]]] definitions and their use in [coded elements](terminologies.html).
type ValueSet struct {
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
	// An absolute URI that is used to identify this value set when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this value set is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the value set is stored on different servers.
	Url *Uri
	// A formal identifier that is used to identify this value set when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the value set when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the value set author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the value set. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the value set.
	Title *String
	// The status of this value set. Enables tracking the life-cycle of the content. The status of the value set applies to the value set definition (ValueSet.compose) and the associated ValueSet metadata. Expansions do not have a state.
	Status Code
	// A Boolean value to indicate that this value set is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date (and optionally time) when the value set was created or revised (e.g. the 'content logical definition').
	Date *DateTime
	// The name of the organization or individual that published the value set.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the value set from a consumer's perspective. The textual description specifies the span of meanings for concepts to be included within the Value Set Expansion, and also may specify the intended use and limitations of the Value Set.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate value set instances.
	UseContext []UsageContext
	// A legal or geographic region in which the value set is intended to be used.
	Jurisdiction []CodeableConcept
	// If this is set to 'true', then no new versions of the content logical definition can be created.  Note: Other metadata might still change.
	Immutable *Boolean
	// Explanation of why this value set is needed and why it has been designed as it has.
	Purpose *Markdown
	// A copyright statement relating to the value set and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the value set.
	Copyright *Markdown
	// A set of criteria that define the contents of the value set by including or excluding codes selected from the specified code system(s) that the value set draws from. This is also known as the Content Logical Definition (CLD).
	Compose *ValueSetCompose
	// A value set can also be "expanded", where the value set is turned into a simple collection of enumerated codes. This element holds the expansion, if it has been performed.
	Expansion *ValueSetExpansion
}

func (r ValueSet) ResourceType() string {
	return "ValueSet"
}
func (r ValueSet) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonValueSet struct {
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
	Url                           *Uri                `json:"url,omitempty"`
	UrlPrimitiveElement           *primitiveElement   `json:"_url,omitempty"`
	Identifier                    []Identifier        `json:"identifier,omitempty"`
	Version                       *String             `json:"version,omitempty"`
	VersionPrimitiveElement       *primitiveElement   `json:"_version,omitempty"`
	Name                          *String             `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement   `json:"_name,omitempty"`
	Title                         *String             `json:"title,omitempty"`
	TitlePrimitiveElement         *primitiveElement   `json:"_title,omitempty"`
	Status                        Code                `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement   `json:"_status,omitempty"`
	Experimental                  *Boolean            `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement  *primitiveElement   `json:"_experimental,omitempty"`
	Date                          *DateTime           `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement   `json:"_date,omitempty"`
	Publisher                     *String             `json:"publisher,omitempty"`
	PublisherPrimitiveElement     *primitiveElement   `json:"_publisher,omitempty"`
	Contact                       []ContactDetail     `json:"contact,omitempty"`
	Description                   *Markdown           `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement   `json:"_description,omitempty"`
	UseContext                    []UsageContext      `json:"useContext,omitempty"`
	Jurisdiction                  []CodeableConcept   `json:"jurisdiction,omitempty"`
	Immutable                     *Boolean            `json:"immutable,omitempty"`
	ImmutablePrimitiveElement     *primitiveElement   `json:"_immutable,omitempty"`
	Purpose                       *Markdown           `json:"purpose,omitempty"`
	PurposePrimitiveElement       *primitiveElement   `json:"_purpose,omitempty"`
	Copyright                     *Markdown           `json:"copyright,omitempty"`
	CopyrightPrimitiveElement     *primitiveElement   `json:"_copyright,omitempty"`
	Compose                       *ValueSetCompose    `json:"compose,omitempty"`
	Expansion                     *ValueSetExpansion  `json:"expansion,omitempty"`
}

func (r ValueSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ValueSet) marshalJSON() jsonValueSet {
	m := jsonValueSet{}
	m.ResourceType = "ValueSet"
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
	m.Immutable = r.Immutable
	if r.Immutable != nil && (r.Immutable.Id != nil || r.Immutable.Extension != nil) {
		m.ImmutablePrimitiveElement = &primitiveElement{Id: r.Immutable.Id, Extension: r.Immutable.Extension}
	}
	m.Purpose = r.Purpose
	if r.Purpose != nil && (r.Purpose.Id != nil || r.Purpose.Extension != nil) {
		m.PurposePrimitiveElement = &primitiveElement{Id: r.Purpose.Id, Extension: r.Purpose.Extension}
	}
	m.Copyright = r.Copyright
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		m.CopyrightPrimitiveElement = &primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
	}
	m.Compose = r.Compose
	m.Expansion = r.Expansion
	return m
}
func (r *ValueSet) UnmarshalJSON(b []byte) error {
	var m jsonValueSet
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ValueSet) unmarshalJSON(m jsonValueSet) error {
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
	r.Immutable = m.Immutable
	if m.ImmutablePrimitiveElement != nil {
		r.Immutable.Id = m.ImmutablePrimitiveElement.Id
		r.Immutable.Extension = m.ImmutablePrimitiveElement.Extension
	}
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
	r.Compose = m.Compose
	r.Expansion = m.Expansion
	return nil
}
func (r ValueSet) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A set of criteria that define the contents of the value set by including or excluding codes selected from the specified code system(s) that the value set draws from. This is also known as the Content Logical Definition (CLD).
type ValueSetCompose struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The Locked Date is  the effective date that is used to determine the version of all referenced Code Systems and Value Set Definitions included in the compose that are not already tied to a specific version.
	LockedDate *Date
	// Whether inactive codes - codes that are not approved for current use - are in the value set. If inactive = true, inactive codes are to be included in the expansion, if inactive = false, the inactive codes will not be included in the expansion. If absent, the behavior is determined by the implementation, or by the applicable $expand parameters (but generally, inactive codes would be expected to be included).
	Inactive *Boolean
	// Include one or more codes from a code system or other value set(s).
	Include []ValueSetComposeInclude
	// Exclude one or more codes from the value set based on code system filters and/or other value sets.
	Exclude []ValueSetComposeInclude
}
type jsonValueSetCompose struct {
	Id                         *string                  `json:"id,omitempty"`
	Extension                  []Extension              `json:"extension,omitempty"`
	ModifierExtension          []Extension              `json:"modifierExtension,omitempty"`
	LockedDate                 *Date                    `json:"lockedDate,omitempty"`
	LockedDatePrimitiveElement *primitiveElement        `json:"_lockedDate,omitempty"`
	Inactive                   *Boolean                 `json:"inactive,omitempty"`
	InactivePrimitiveElement   *primitiveElement        `json:"_inactive,omitempty"`
	Include                    []ValueSetComposeInclude `json:"include,omitempty"`
	Exclude                    []ValueSetComposeInclude `json:"exclude,omitempty"`
}

func (r ValueSetCompose) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ValueSetCompose) marshalJSON() jsonValueSetCompose {
	m := jsonValueSetCompose{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.LockedDate = r.LockedDate
	if r.LockedDate != nil && (r.LockedDate.Id != nil || r.LockedDate.Extension != nil) {
		m.LockedDatePrimitiveElement = &primitiveElement{Id: r.LockedDate.Id, Extension: r.LockedDate.Extension}
	}
	m.Inactive = r.Inactive
	if r.Inactive != nil && (r.Inactive.Id != nil || r.Inactive.Extension != nil) {
		m.InactivePrimitiveElement = &primitiveElement{Id: r.Inactive.Id, Extension: r.Inactive.Extension}
	}
	m.Include = r.Include
	m.Exclude = r.Exclude
	return m
}
func (r *ValueSetCompose) UnmarshalJSON(b []byte) error {
	var m jsonValueSetCompose
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ValueSetCompose) unmarshalJSON(m jsonValueSetCompose) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.LockedDate = m.LockedDate
	if m.LockedDatePrimitiveElement != nil {
		r.LockedDate.Id = m.LockedDatePrimitiveElement.Id
		r.LockedDate.Extension = m.LockedDatePrimitiveElement.Extension
	}
	r.Inactive = m.Inactive
	if m.InactivePrimitiveElement != nil {
		r.Inactive.Id = m.InactivePrimitiveElement.Id
		r.Inactive.Extension = m.InactivePrimitiveElement.Extension
	}
	r.Include = m.Include
	r.Exclude = m.Exclude
	return nil
}
func (r ValueSetCompose) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Include one or more codes from a code system or other value set(s).
type ValueSetComposeInclude struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An absolute URI which is the code system from which the selected codes come from.
	System *Uri
	// The version of the code system that the codes are selected from, or the special version '*' for all versions.
	Version *String
	// Specifies a concept to be included or excluded.
	Concept []ValueSetComposeIncludeConcept
	// Select concepts by specify a matching criterion based on the properties (including relationships) defined by the system, or on filters defined by the system. If multiple filters are specified, they SHALL all be true.
	Filter []ValueSetComposeIncludeFilter
	// Selects the concepts found in this value set (based on its value set definition). This is an absolute URI that is a reference to ValueSet.url.  If multiple value sets are specified this includes the union of the contents of all of the referenced value sets.
	ValueSet []Canonical
}
type jsonValueSetComposeInclude struct {
	Id                       *string                         `json:"id,omitempty"`
	Extension                []Extension                     `json:"extension,omitempty"`
	ModifierExtension        []Extension                     `json:"modifierExtension,omitempty"`
	System                   *Uri                            `json:"system,omitempty"`
	SystemPrimitiveElement   *primitiveElement               `json:"_system,omitempty"`
	Version                  *String                         `json:"version,omitempty"`
	VersionPrimitiveElement  *primitiveElement               `json:"_version,omitempty"`
	Concept                  []ValueSetComposeIncludeConcept `json:"concept,omitempty"`
	Filter                   []ValueSetComposeIncludeFilter  `json:"filter,omitempty"`
	ValueSet                 []Canonical                     `json:"valueSet,omitempty"`
	ValueSetPrimitiveElement []*primitiveElement             `json:"_valueSet,omitempty"`
}

func (r ValueSetComposeInclude) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ValueSetComposeInclude) marshalJSON() jsonValueSetComposeInclude {
	m := jsonValueSetComposeInclude{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.System = r.System
	if r.System != nil && (r.System.Id != nil || r.System.Extension != nil) {
		m.SystemPrimitiveElement = &primitiveElement{Id: r.System.Id, Extension: r.System.Extension}
	}
	m.Version = r.Version
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	m.Concept = r.Concept
	m.Filter = r.Filter
	m.ValueSet = r.ValueSet
	anyValueSetIdOrExtension := false
	for _, e := range r.ValueSet {
		if e.Id != nil || e.Extension != nil {
			anyValueSetIdOrExtension = true
			break
		}
	}
	if anyValueSetIdOrExtension {
		m.ValueSetPrimitiveElement = make([]*primitiveElement, 0, len(r.ValueSet))
		for _, e := range r.ValueSet {
			if e.Id != nil || e.Extension != nil {
				m.ValueSetPrimitiveElement = append(m.ValueSetPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ValueSetPrimitiveElement = append(m.ValueSetPrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *ValueSetComposeInclude) UnmarshalJSON(b []byte) error {
	var m jsonValueSetComposeInclude
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ValueSetComposeInclude) unmarshalJSON(m jsonValueSetComposeInclude) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.System = m.System
	if m.SystemPrimitiveElement != nil {
		r.System.Id = m.SystemPrimitiveElement.Id
		r.System.Extension = m.SystemPrimitiveElement.Extension
	}
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	r.Concept = m.Concept
	r.Filter = m.Filter
	r.ValueSet = m.ValueSet
	for i, e := range m.ValueSetPrimitiveElement {
		if len(r.ValueSet) > i {
			r.ValueSet[i].Id = e.Id
			r.ValueSet[i].Extension = e.Extension
		} else {
			r.ValueSet = append(r.ValueSet, Canonical{Id: e.Id, Extension: e.Extension})
		}
	}
	return nil
}
func (r ValueSetComposeInclude) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Specifies a concept to be included or excluded.
type ValueSetComposeIncludeConcept struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Specifies a code for the concept to be included or excluded.
	Code Code
	// The text to display to the user for this concept in the context of this valueset. If no display is provided, then applications using the value set use the display specified for the code by the system.
	Display *String
	// Additional representations for this concept when used in this value set - other languages, aliases, specialized purposes, used for particular purposes, etc.
	Designation []ValueSetComposeIncludeConceptDesignation
}
type jsonValueSetComposeIncludeConcept struct {
	Id                      *string                                    `json:"id,omitempty"`
	Extension               []Extension                                `json:"extension,omitempty"`
	ModifierExtension       []Extension                                `json:"modifierExtension,omitempty"`
	Code                    Code                                       `json:"code,omitempty"`
	CodePrimitiveElement    *primitiveElement                          `json:"_code,omitempty"`
	Display                 *String                                    `json:"display,omitempty"`
	DisplayPrimitiveElement *primitiveElement                          `json:"_display,omitempty"`
	Designation             []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty"`
}

func (r ValueSetComposeIncludeConcept) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ValueSetComposeIncludeConcept) marshalJSON() jsonValueSetComposeIncludeConcept {
	m := jsonValueSetComposeIncludeConcept{}
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
	m.Designation = r.Designation
	return m
}
func (r *ValueSetComposeIncludeConcept) UnmarshalJSON(b []byte) error {
	var m jsonValueSetComposeIncludeConcept
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ValueSetComposeIncludeConcept) unmarshalJSON(m jsonValueSetComposeIncludeConcept) error {
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
	r.Designation = m.Designation
	return nil
}
func (r ValueSetComposeIncludeConcept) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Additional representations for this concept when used in this value set - other languages, aliases, specialized purposes, used for particular purposes, etc.
type ValueSetComposeIncludeConceptDesignation struct {
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
	// A code that represents types of uses of designations.
	Use *Coding
	// The text value for this designation.
	Value String
}
type jsonValueSetComposeIncludeConceptDesignation struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Language                 *Code             `json:"language,omitempty"`
	LanguagePrimitiveElement *primitiveElement `json:"_language,omitempty"`
	Use                      *Coding           `json:"use,omitempty"`
	Value                    String            `json:"value,omitempty"`
	ValuePrimitiveElement    *primitiveElement `json:"_value,omitempty"`
}

func (r ValueSetComposeIncludeConceptDesignation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ValueSetComposeIncludeConceptDesignation) marshalJSON() jsonValueSetComposeIncludeConceptDesignation {
	m := jsonValueSetComposeIncludeConceptDesignation{}
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
func (r *ValueSetComposeIncludeConceptDesignation) UnmarshalJSON(b []byte) error {
	var m jsonValueSetComposeIncludeConceptDesignation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ValueSetComposeIncludeConceptDesignation) unmarshalJSON(m jsonValueSetComposeIncludeConceptDesignation) error {
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
func (r ValueSetComposeIncludeConceptDesignation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Select concepts by specify a matching criterion based on the properties (including relationships) defined by the system, or on filters defined by the system. If multiple filters are specified, they SHALL all be true.
type ValueSetComposeIncludeFilter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code that identifies a property or a filter defined in the code system.
	Property Code
	// The kind of operation to perform as a part of the filter criteria.
	Op Code
	// The match value may be either a code defined by the system, or a string value, which is a regex match on the literal string of the property value  (if the filter represents a property defined in CodeSystem) or of the system filter value (if the filter represents a filter defined in CodeSystem) when the operation is 'regex', or one of the values (true and false), when the operation is 'exists'.
	Value String
}
type jsonValueSetComposeIncludeFilter struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Property                 Code              `json:"property,omitempty"`
	PropertyPrimitiveElement *primitiveElement `json:"_property,omitempty"`
	Op                       Code              `json:"op,omitempty"`
	OpPrimitiveElement       *primitiveElement `json:"_op,omitempty"`
	Value                    String            `json:"value,omitempty"`
	ValuePrimitiveElement    *primitiveElement `json:"_value,omitempty"`
}

func (r ValueSetComposeIncludeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ValueSetComposeIncludeFilter) marshalJSON() jsonValueSetComposeIncludeFilter {
	m := jsonValueSetComposeIncludeFilter{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Property = r.Property
	if r.Property.Id != nil || r.Property.Extension != nil {
		m.PropertyPrimitiveElement = &primitiveElement{Id: r.Property.Id, Extension: r.Property.Extension}
	}
	m.Op = r.Op
	if r.Op.Id != nil || r.Op.Extension != nil {
		m.OpPrimitiveElement = &primitiveElement{Id: r.Op.Id, Extension: r.Op.Extension}
	}
	m.Value = r.Value
	if r.Value.Id != nil || r.Value.Extension != nil {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	return m
}
func (r *ValueSetComposeIncludeFilter) UnmarshalJSON(b []byte) error {
	var m jsonValueSetComposeIncludeFilter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ValueSetComposeIncludeFilter) unmarshalJSON(m jsonValueSetComposeIncludeFilter) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Property = m.Property
	if m.PropertyPrimitiveElement != nil {
		r.Property.Id = m.PropertyPrimitiveElement.Id
		r.Property.Extension = m.PropertyPrimitiveElement.Extension
	}
	r.Op = m.Op
	if m.OpPrimitiveElement != nil {
		r.Op.Id = m.OpPrimitiveElement.Id
		r.Op.Extension = m.OpPrimitiveElement.Extension
	}
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	return nil
}
func (r ValueSetComposeIncludeFilter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A value set can also be "expanded", where the value set is turned into a simple collection of enumerated codes. This element holds the expansion, if it has been performed.
type ValueSetExpansion struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An identifier that uniquely identifies this expansion of the valueset, based on a unique combination of the provided parameters, the system default parameters, and the underlying system code system versions etc. Systems may re-use the same identifier as long as those factors remain the same, and the expansion is the same, but are not required to do so. This is a business identifier.
	Identifier *Uri
	// The time at which the expansion was produced by the expanding system.
	Timestamp DateTime
	// The total number of concepts in the expansion. If the number of concept nodes in this resource is less than the stated number, then the server can return more using the offset parameter.
	Total *Integer
	// If paging is being used, the offset at which this resource starts.  I.e. this resource is a partial view into the expansion. If paging is not being used, this element SHALL NOT be present.
	Offset *Integer
	// A parameter that controlled the expansion process. These parameters may be used by users of expanded value sets to check whether the expansion is suitable for a particular purpose, or to pick the correct expansion.
	Parameter []ValueSetExpansionParameter
	// The codes that are contained in the value set expansion.
	Contains []ValueSetExpansionContains
}
type jsonValueSetExpansion struct {
	Id                         *string                      `json:"id,omitempty"`
	Extension                  []Extension                  `json:"extension,omitempty"`
	ModifierExtension          []Extension                  `json:"modifierExtension,omitempty"`
	Identifier                 *Uri                         `json:"identifier,omitempty"`
	IdentifierPrimitiveElement *primitiveElement            `json:"_identifier,omitempty"`
	Timestamp                  DateTime                     `json:"timestamp,omitempty"`
	TimestampPrimitiveElement  *primitiveElement            `json:"_timestamp,omitempty"`
	Total                      *Integer                     `json:"total,omitempty"`
	TotalPrimitiveElement      *primitiveElement            `json:"_total,omitempty"`
	Offset                     *Integer                     `json:"offset,omitempty"`
	OffsetPrimitiveElement     *primitiveElement            `json:"_offset,omitempty"`
	Parameter                  []ValueSetExpansionParameter `json:"parameter,omitempty"`
	Contains                   []ValueSetExpansionContains  `json:"contains,omitempty"`
}

func (r ValueSetExpansion) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ValueSetExpansion) marshalJSON() jsonValueSetExpansion {
	m := jsonValueSetExpansion{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	if r.Identifier != nil && (r.Identifier.Id != nil || r.Identifier.Extension != nil) {
		m.IdentifierPrimitiveElement = &primitiveElement{Id: r.Identifier.Id, Extension: r.Identifier.Extension}
	}
	m.Timestamp = r.Timestamp
	if r.Timestamp.Id != nil || r.Timestamp.Extension != nil {
		m.TimestampPrimitiveElement = &primitiveElement{Id: r.Timestamp.Id, Extension: r.Timestamp.Extension}
	}
	m.Total = r.Total
	if r.Total != nil && (r.Total.Id != nil || r.Total.Extension != nil) {
		m.TotalPrimitiveElement = &primitiveElement{Id: r.Total.Id, Extension: r.Total.Extension}
	}
	m.Offset = r.Offset
	if r.Offset != nil && (r.Offset.Id != nil || r.Offset.Extension != nil) {
		m.OffsetPrimitiveElement = &primitiveElement{Id: r.Offset.Id, Extension: r.Offset.Extension}
	}
	m.Parameter = r.Parameter
	m.Contains = r.Contains
	return m
}
func (r *ValueSetExpansion) UnmarshalJSON(b []byte) error {
	var m jsonValueSetExpansion
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ValueSetExpansion) unmarshalJSON(m jsonValueSetExpansion) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	if m.IdentifierPrimitiveElement != nil {
		r.Identifier.Id = m.IdentifierPrimitiveElement.Id
		r.Identifier.Extension = m.IdentifierPrimitiveElement.Extension
	}
	r.Timestamp = m.Timestamp
	if m.TimestampPrimitiveElement != nil {
		r.Timestamp.Id = m.TimestampPrimitiveElement.Id
		r.Timestamp.Extension = m.TimestampPrimitiveElement.Extension
	}
	r.Total = m.Total
	if m.TotalPrimitiveElement != nil {
		r.Total.Id = m.TotalPrimitiveElement.Id
		r.Total.Extension = m.TotalPrimitiveElement.Extension
	}
	r.Offset = m.Offset
	if m.OffsetPrimitiveElement != nil {
		r.Offset.Id = m.OffsetPrimitiveElement.Id
		r.Offset.Extension = m.OffsetPrimitiveElement.Extension
	}
	r.Parameter = m.Parameter
	r.Contains = m.Contains
	return nil
}
func (r ValueSetExpansion) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A parameter that controlled the expansion process. These parameters may be used by users of expanded value sets to check whether the expansion is suitable for a particular purpose, or to pick the correct expansion.
type ValueSetExpansionParameter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Name of the input parameter to the $expand operation; may be a server-assigned name for additional default or other server-supplied parameters used to control the expansion process.
	Name String
	// The value of the parameter.
	Value isValueSetExpansionParameterValue
}
type isValueSetExpansionParameterValue interface {
	isValueSetExpansionParameterValue()
}

func (r String) isValueSetExpansionParameterValue()   {}
func (r Boolean) isValueSetExpansionParameterValue()  {}
func (r Integer) isValueSetExpansionParameterValue()  {}
func (r Decimal) isValueSetExpansionParameterValue()  {}
func (r Uri) isValueSetExpansionParameterValue()      {}
func (r Code) isValueSetExpansionParameterValue()     {}
func (r DateTime) isValueSetExpansionParameterValue() {}

type jsonValueSetExpansionParameter struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	Name                          String            `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement `json:"_name,omitempty"`
	ValueString                   *String           `json:"valueString,omitempty"`
	ValueStringPrimitiveElement   *primitiveElement `json:"_valueString,omitempty"`
	ValueBoolean                  *Boolean          `json:"valueBoolean,omitempty"`
	ValueBooleanPrimitiveElement  *primitiveElement `json:"_valueBoolean,omitempty"`
	ValueInteger                  *Integer          `json:"valueInteger,omitempty"`
	ValueIntegerPrimitiveElement  *primitiveElement `json:"_valueInteger,omitempty"`
	ValueDecimal                  *Decimal          `json:"valueDecimal,omitempty"`
	ValueDecimalPrimitiveElement  *primitiveElement `json:"_valueDecimal,omitempty"`
	ValueUri                      *Uri              `json:"valueUri,omitempty"`
	ValueUriPrimitiveElement      *primitiveElement `json:"_valueUri,omitempty"`
	ValueCode                     *Code             `json:"valueCode,omitempty"`
	ValueCodePrimitiveElement     *primitiveElement `json:"_valueCode,omitempty"`
	ValueDateTime                 *DateTime         `json:"valueDateTime,omitempty"`
	ValueDateTimePrimitiveElement *primitiveElement `json:"_valueDateTime,omitempty"`
}

func (r ValueSetExpansionParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ValueSetExpansionParameter) marshalJSON() jsonValueSetExpansionParameter {
	m := jsonValueSetExpansionParameter{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Name = r.Name
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	switch v := r.Value.(type) {
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
	case Uri:
		m.ValueUri = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uri:
		m.ValueUri = v
		if v.Id != nil || v.Extension != nil {
			m.ValueUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
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
	}
	return m
}
func (r *ValueSetExpansionParameter) UnmarshalJSON(b []byte) error {
	var m jsonValueSetExpansionParameter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ValueSetExpansionParameter) unmarshalJSON(m jsonValueSetExpansionParameter) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
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
	if m.ValueUri != nil || m.ValueUriPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUri
		if m.ValueUriPrimitiveElement != nil {
			if v == nil {
				v = &Uri{}
			}
			v.Id = m.ValueUriPrimitiveElement.Id
			v.Extension = m.ValueUriPrimitiveElement.Extension
		}
		r.Value = v
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
	return nil
}
func (r ValueSetExpansionParameter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The codes that are contained in the value set expansion.
type ValueSetExpansionContains struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An absolute URI which is the code system in which the code for this item in the expansion is defined.
	System *Uri
	// If true, this entry is included in the expansion for navigational purposes, and the user cannot select the code directly as a proper value.
	Abstract *Boolean
	// If the concept is inactive in the code system that defines it. Inactive codes are those that are no longer to be used, but are maintained by the code system for understanding legacy data. It might not be known or specified whether an concept is inactive (and it may depend on the context of use).
	Inactive *Boolean
	// The version of the code system from this code was taken. Note that a well-maintained code system does not need the version reported, because the meaning of codes is consistent across versions. However this cannot consistently be assured, and when the meaning is not guaranteed to be consistent, the version SHOULD be exchanged.
	Version *String
	// The code for this item in the expansion hierarchy. If this code is missing the entry in the hierarchy is a place holder (abstract) and does not represent a valid code in the value set.
	Code *Code
	// The recommended display for this item in the expansion.
	Display *String
	// Additional representations for this item - other languages, aliases, specialized purposes, used for particular purposes, etc. These are relevant when the conditions of the expansion do not fix to a single correct representation.
	Designation []ValueSetComposeIncludeConceptDesignation
	// Other codes and entries contained under this entry in the hierarchy.
	Contains []ValueSetExpansionContains
}
type jsonValueSetExpansionContains struct {
	Id                       *string                                    `json:"id,omitempty"`
	Extension                []Extension                                `json:"extension,omitempty"`
	ModifierExtension        []Extension                                `json:"modifierExtension,omitempty"`
	System                   *Uri                                       `json:"system,omitempty"`
	SystemPrimitiveElement   *primitiveElement                          `json:"_system,omitempty"`
	Abstract                 *Boolean                                   `json:"abstract,omitempty"`
	AbstractPrimitiveElement *primitiveElement                          `json:"_abstract,omitempty"`
	Inactive                 *Boolean                                   `json:"inactive,omitempty"`
	InactivePrimitiveElement *primitiveElement                          `json:"_inactive,omitempty"`
	Version                  *String                                    `json:"version,omitempty"`
	VersionPrimitiveElement  *primitiveElement                          `json:"_version,omitempty"`
	Code                     *Code                                      `json:"code,omitempty"`
	CodePrimitiveElement     *primitiveElement                          `json:"_code,omitempty"`
	Display                  *String                                    `json:"display,omitempty"`
	DisplayPrimitiveElement  *primitiveElement                          `json:"_display,omitempty"`
	Designation              []ValueSetComposeIncludeConceptDesignation `json:"designation,omitempty"`
	Contains                 []ValueSetExpansionContains                `json:"contains,omitempty"`
}

func (r ValueSetExpansionContains) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ValueSetExpansionContains) marshalJSON() jsonValueSetExpansionContains {
	m := jsonValueSetExpansionContains{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.System = r.System
	if r.System != nil && (r.System.Id != nil || r.System.Extension != nil) {
		m.SystemPrimitiveElement = &primitiveElement{Id: r.System.Id, Extension: r.System.Extension}
	}
	m.Abstract = r.Abstract
	if r.Abstract != nil && (r.Abstract.Id != nil || r.Abstract.Extension != nil) {
		m.AbstractPrimitiveElement = &primitiveElement{Id: r.Abstract.Id, Extension: r.Abstract.Extension}
	}
	m.Inactive = r.Inactive
	if r.Inactive != nil && (r.Inactive.Id != nil || r.Inactive.Extension != nil) {
		m.InactivePrimitiveElement = &primitiveElement{Id: r.Inactive.Id, Extension: r.Inactive.Extension}
	}
	m.Version = r.Version
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	m.Code = r.Code
	if r.Code != nil && (r.Code.Id != nil || r.Code.Extension != nil) {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Display = r.Display
	if r.Display != nil && (r.Display.Id != nil || r.Display.Extension != nil) {
		m.DisplayPrimitiveElement = &primitiveElement{Id: r.Display.Id, Extension: r.Display.Extension}
	}
	m.Designation = r.Designation
	m.Contains = r.Contains
	return m
}
func (r *ValueSetExpansionContains) UnmarshalJSON(b []byte) error {
	var m jsonValueSetExpansionContains
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ValueSetExpansionContains) unmarshalJSON(m jsonValueSetExpansionContains) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.System = m.System
	if m.SystemPrimitiveElement != nil {
		r.System.Id = m.SystemPrimitiveElement.Id
		r.System.Extension = m.SystemPrimitiveElement.Extension
	}
	r.Abstract = m.Abstract
	if m.AbstractPrimitiveElement != nil {
		r.Abstract.Id = m.AbstractPrimitiveElement.Id
		r.Abstract.Extension = m.AbstractPrimitiveElement.Extension
	}
	r.Inactive = m.Inactive
	if m.InactivePrimitiveElement != nil {
		r.Inactive.Id = m.InactivePrimitiveElement.Id
		r.Inactive.Extension = m.InactivePrimitiveElement.Extension
	}
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
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
	r.Designation = m.Designation
	r.Contains = m.Contains
	return nil
}
func (r ValueSetExpansionContains) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
