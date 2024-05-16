package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A TerminologyCapabilities resource documents a set of capabilities (behaviors) of a FHIR Terminology Server that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type TerminologyCapabilities struct {
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
	// An absolute URI that is used to identify this terminology capabilities when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this terminology capabilities is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the terminology capabilities is stored on different servers.
	Url *Uri
	// The identifier that is used to identify this version of the terminology capabilities when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the terminology capabilities author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the terminology capabilities. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the terminology capabilities.
	Title *String
	// The status of this terminology capabilities. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this terminology capabilities is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the terminology capabilities was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the terminology capabilities changes.
	Date DateTime
	// The name of the organization or individual that published the terminology capabilities.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the terminology capabilities from a consumer's perspective. Typically, this is used when the capability statement describes a desired rather than an actual solution, for example as a formal expression of requirements as part of an RFP.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate terminology capabilities instances.
	UseContext []UsageContext
	// A legal or geographic region in which the terminology capabilities is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this terminology capabilities is needed and why it has been designed as it has.
	Purpose *Markdown
	// A copyright statement relating to the terminology capabilities and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the terminology capabilities.
	Copyright *Markdown
	// The way that this statement is intended to be used, to describe an actual running instance of software, a particular product (kind, not instance of software) or a class of implementation (e.g. a desired purchase).
	Kind Code
	// Software that is covered by this terminology capability statement.  It is used when the statement describes the capabilities of a particular software version, independent of an installation.
	Software *TerminologyCapabilitiesSoftware
	// Identifies a specific implementation instance that is described by the terminology capability statement - i.e. a particular installation, rather than the capabilities of a software program.
	Implementation *TerminologyCapabilitiesImplementation
	// Whether the server supports lockedDate.
	LockedDate *Boolean
	// Identifies a code system that is supported by the server. If there is a no code system URL, then this declares the general assumptions a client can make about support for any CodeSystem resource.
	CodeSystem []TerminologyCapabilitiesCodeSystem
	// Information about the [ValueSet/$expand](valueset-operation-expand.html) operation.
	Expansion *TerminologyCapabilitiesExpansion
	// The degree to which the server supports the code search parameter on ValueSet, if it is supported.
	CodeSearch *Code
	// Information about the [ValueSet/$validate-code](valueset-operation-validate-code.html) operation.
	ValidateCode *TerminologyCapabilitiesValidateCode
	// Information about the [ConceptMap/$translate](conceptmap-operation-translate.html) operation.
	Translation *TerminologyCapabilitiesTranslation
	// Whether the $closure operation is supported.
	Closure *TerminologyCapabilitiesClosure
}

func (r TerminologyCapabilities) ResourceType() string {
	return "TerminologyCapabilities"
}
func (r TerminologyCapabilities) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonTerminologyCapabilities struct {
	ResourceType                  string                                 `json:"resourceType"`
	Id                            *Id                                    `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement                      `json:"_id,omitempty"`
	Meta                          *Meta                                  `json:"meta,omitempty"`
	ImplicitRules                 *Uri                                   `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement                      `json:"_implicitRules,omitempty"`
	Language                      *Code                                  `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement                      `json:"_language,omitempty"`
	Text                          *Narrative                             `json:"text,omitempty"`
	Contained                     []containedResource                    `json:"contained,omitempty"`
	Extension                     []Extension                            `json:"extension,omitempty"`
	ModifierExtension             []Extension                            `json:"modifierExtension,omitempty"`
	Url                           *Uri                                   `json:"url,omitempty"`
	UrlPrimitiveElement           *primitiveElement                      `json:"_url,omitempty"`
	Version                       *String                                `json:"version,omitempty"`
	VersionPrimitiveElement       *primitiveElement                      `json:"_version,omitempty"`
	Name                          *String                                `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement                      `json:"_name,omitempty"`
	Title                         *String                                `json:"title,omitempty"`
	TitlePrimitiveElement         *primitiveElement                      `json:"_title,omitempty"`
	Status                        Code                                   `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement                      `json:"_status,omitempty"`
	Experimental                  *Boolean                               `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement  *primitiveElement                      `json:"_experimental,omitempty"`
	Date                          DateTime                               `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement                      `json:"_date,omitempty"`
	Publisher                     *String                                `json:"publisher,omitempty"`
	PublisherPrimitiveElement     *primitiveElement                      `json:"_publisher,omitempty"`
	Contact                       []ContactDetail                        `json:"contact,omitempty"`
	Description                   *Markdown                              `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement                      `json:"_description,omitempty"`
	UseContext                    []UsageContext                         `json:"useContext,omitempty"`
	Jurisdiction                  []CodeableConcept                      `json:"jurisdiction,omitempty"`
	Purpose                       *Markdown                              `json:"purpose,omitempty"`
	PurposePrimitiveElement       *primitiveElement                      `json:"_purpose,omitempty"`
	Copyright                     *Markdown                              `json:"copyright,omitempty"`
	CopyrightPrimitiveElement     *primitiveElement                      `json:"_copyright,omitempty"`
	Kind                          Code                                   `json:"kind,omitempty"`
	KindPrimitiveElement          *primitiveElement                      `json:"_kind,omitempty"`
	Software                      *TerminologyCapabilitiesSoftware       `json:"software,omitempty"`
	Implementation                *TerminologyCapabilitiesImplementation `json:"implementation,omitempty"`
	LockedDate                    *Boolean                               `json:"lockedDate,omitempty"`
	LockedDatePrimitiveElement    *primitiveElement                      `json:"_lockedDate,omitempty"`
	CodeSystem                    []TerminologyCapabilitiesCodeSystem    `json:"codeSystem,omitempty"`
	Expansion                     *TerminologyCapabilitiesExpansion      `json:"expansion,omitempty"`
	CodeSearch                    *Code                                  `json:"codeSearch,omitempty"`
	CodeSearchPrimitiveElement    *primitiveElement                      `json:"_codeSearch,omitempty"`
	ValidateCode                  *TerminologyCapabilitiesValidateCode   `json:"validateCode,omitempty"`
	Translation                   *TerminologyCapabilitiesTranslation    `json:"translation,omitempty"`
	Closure                       *TerminologyCapabilitiesClosure        `json:"closure,omitempty"`
}

func (r TerminologyCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TerminologyCapabilities) marshalJSON() jsonTerminologyCapabilities {
	m := jsonTerminologyCapabilities{}
	m.ResourceType = "TerminologyCapabilities"
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
	m.Kind = r.Kind
	if r.Kind.Id != nil || r.Kind.Extension != nil {
		m.KindPrimitiveElement = &primitiveElement{Id: r.Kind.Id, Extension: r.Kind.Extension}
	}
	m.Software = r.Software
	m.Implementation = r.Implementation
	m.LockedDate = r.LockedDate
	if r.LockedDate != nil && (r.LockedDate.Id != nil || r.LockedDate.Extension != nil) {
		m.LockedDatePrimitiveElement = &primitiveElement{Id: r.LockedDate.Id, Extension: r.LockedDate.Extension}
	}
	m.CodeSystem = r.CodeSystem
	m.Expansion = r.Expansion
	m.CodeSearch = r.CodeSearch
	if r.CodeSearch != nil && (r.CodeSearch.Id != nil || r.CodeSearch.Extension != nil) {
		m.CodeSearchPrimitiveElement = &primitiveElement{Id: r.CodeSearch.Id, Extension: r.CodeSearch.Extension}
	}
	m.ValidateCode = r.ValidateCode
	m.Translation = r.Translation
	m.Closure = r.Closure
	return m
}
func (r *TerminologyCapabilities) UnmarshalJSON(b []byte) error {
	var m jsonTerminologyCapabilities
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TerminologyCapabilities) unmarshalJSON(m jsonTerminologyCapabilities) error {
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
	r.Kind = m.Kind
	if m.KindPrimitiveElement != nil {
		r.Kind.Id = m.KindPrimitiveElement.Id
		r.Kind.Extension = m.KindPrimitiveElement.Extension
	}
	r.Software = m.Software
	r.Implementation = m.Implementation
	r.LockedDate = m.LockedDate
	if m.LockedDatePrimitiveElement != nil {
		r.LockedDate.Id = m.LockedDatePrimitiveElement.Id
		r.LockedDate.Extension = m.LockedDatePrimitiveElement.Extension
	}
	r.CodeSystem = m.CodeSystem
	r.Expansion = m.Expansion
	r.CodeSearch = m.CodeSearch
	if m.CodeSearchPrimitiveElement != nil {
		r.CodeSearch.Id = m.CodeSearchPrimitiveElement.Id
		r.CodeSearch.Extension = m.CodeSearchPrimitiveElement.Extension
	}
	r.ValidateCode = m.ValidateCode
	r.Translation = m.Translation
	r.Closure = m.Closure
	return nil
}
func (r TerminologyCapabilities) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Software that is covered by this terminology capability statement.  It is used when the statement describes the capabilities of a particular software version, independent of an installation.
type TerminologyCapabilitiesSoftware struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Name the software is known by.
	Name String
	// The version identifier for the software covered by this statement.
	Version *String
}
type jsonTerminologyCapabilitiesSoftware struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Name                    String            `json:"name,omitempty"`
	NamePrimitiveElement    *primitiveElement `json:"_name,omitempty"`
	Version                 *String           `json:"version,omitempty"`
	VersionPrimitiveElement *primitiveElement `json:"_version,omitempty"`
}

func (r TerminologyCapabilitiesSoftware) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TerminologyCapabilitiesSoftware) marshalJSON() jsonTerminologyCapabilitiesSoftware {
	m := jsonTerminologyCapabilitiesSoftware{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Name = r.Name
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Version = r.Version
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	return m
}
func (r *TerminologyCapabilitiesSoftware) UnmarshalJSON(b []byte) error {
	var m jsonTerminologyCapabilitiesSoftware
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TerminologyCapabilitiesSoftware) unmarshalJSON(m jsonTerminologyCapabilitiesSoftware) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	return nil
}
func (r TerminologyCapabilitiesSoftware) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Identifies a specific implementation instance that is described by the terminology capability statement - i.e. a particular installation, rather than the capabilities of a software program.
type TerminologyCapabilitiesImplementation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Information about the specific installation that this terminology capability statement relates to.
	Description String
	// An absolute base URL for the implementation.
	Url *Url
}
type jsonTerminologyCapabilitiesImplementation struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Description                 String            `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
	Url                         *Url              `json:"url,omitempty"`
	UrlPrimitiveElement         *primitiveElement `json:"_url,omitempty"`
}

func (r TerminologyCapabilitiesImplementation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TerminologyCapabilitiesImplementation) marshalJSON() jsonTerminologyCapabilitiesImplementation {
	m := jsonTerminologyCapabilitiesImplementation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Description = r.Description
	if r.Description.Id != nil || r.Description.Extension != nil {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Url = r.Url
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	return m
}
func (r *TerminologyCapabilitiesImplementation) UnmarshalJSON(b []byte) error {
	var m jsonTerminologyCapabilitiesImplementation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TerminologyCapabilitiesImplementation) unmarshalJSON(m jsonTerminologyCapabilitiesImplementation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	return nil
}
func (r TerminologyCapabilitiesImplementation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Identifies a code system that is supported by the server. If there is a no code system URL, then this declares the general assumptions a client can make about support for any CodeSystem resource.
type TerminologyCapabilitiesCodeSystem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// URI for the Code System.
	Uri *Canonical
	// For the code system, a list of versions that are supported by the server.
	Version []TerminologyCapabilitiesCodeSystemVersion
	// True if subsumption is supported for this version of the code system.
	Subsumption *Boolean
}
type jsonTerminologyCapabilitiesCodeSystem struct {
	Id                          *string                                    `json:"id,omitempty"`
	Extension                   []Extension                                `json:"extension,omitempty"`
	ModifierExtension           []Extension                                `json:"modifierExtension,omitempty"`
	Uri                         *Canonical                                 `json:"uri,omitempty"`
	UriPrimitiveElement         *primitiveElement                          `json:"_uri,omitempty"`
	Version                     []TerminologyCapabilitiesCodeSystemVersion `json:"version,omitempty"`
	Subsumption                 *Boolean                                   `json:"subsumption,omitempty"`
	SubsumptionPrimitiveElement *primitiveElement                          `json:"_subsumption,omitempty"`
}

func (r TerminologyCapabilitiesCodeSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TerminologyCapabilitiesCodeSystem) marshalJSON() jsonTerminologyCapabilitiesCodeSystem {
	m := jsonTerminologyCapabilitiesCodeSystem{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Uri = r.Uri
	if r.Uri != nil && (r.Uri.Id != nil || r.Uri.Extension != nil) {
		m.UriPrimitiveElement = &primitiveElement{Id: r.Uri.Id, Extension: r.Uri.Extension}
	}
	m.Version = r.Version
	m.Subsumption = r.Subsumption
	if r.Subsumption != nil && (r.Subsumption.Id != nil || r.Subsumption.Extension != nil) {
		m.SubsumptionPrimitiveElement = &primitiveElement{Id: r.Subsumption.Id, Extension: r.Subsumption.Extension}
	}
	return m
}
func (r *TerminologyCapabilitiesCodeSystem) UnmarshalJSON(b []byte) error {
	var m jsonTerminologyCapabilitiesCodeSystem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TerminologyCapabilitiesCodeSystem) unmarshalJSON(m jsonTerminologyCapabilitiesCodeSystem) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Uri = m.Uri
	if m.UriPrimitiveElement != nil {
		r.Uri.Id = m.UriPrimitiveElement.Id
		r.Uri.Extension = m.UriPrimitiveElement.Extension
	}
	r.Version = m.Version
	r.Subsumption = m.Subsumption
	if m.SubsumptionPrimitiveElement != nil {
		r.Subsumption.Id = m.SubsumptionPrimitiveElement.Id
		r.Subsumption.Extension = m.SubsumptionPrimitiveElement.Extension
	}
	return nil
}
func (r TerminologyCapabilitiesCodeSystem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// For the code system, a list of versions that are supported by the server.
type TerminologyCapabilitiesCodeSystemVersion struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// For version-less code systems, there should be a single version with no identifier.
	Code *String
	// If this is the default version for this code system.
	IsDefault *Boolean
	// If the compositional grammar defined by the code system is supported.
	Compositional *Boolean
	// Language Displays supported.
	Language []Code
	// Filter Properties supported.
	Filter []TerminologyCapabilitiesCodeSystemVersionFilter
	// Properties supported for $lookup.
	Property []Code
}
type jsonTerminologyCapabilitiesCodeSystemVersion struct {
	Id                            *string                                          `json:"id,omitempty"`
	Extension                     []Extension                                      `json:"extension,omitempty"`
	ModifierExtension             []Extension                                      `json:"modifierExtension,omitempty"`
	Code                          *String                                          `json:"code,omitempty"`
	CodePrimitiveElement          *primitiveElement                                `json:"_code,omitempty"`
	IsDefault                     *Boolean                                         `json:"isDefault,omitempty"`
	IsDefaultPrimitiveElement     *primitiveElement                                `json:"_isDefault,omitempty"`
	Compositional                 *Boolean                                         `json:"compositional,omitempty"`
	CompositionalPrimitiveElement *primitiveElement                                `json:"_compositional,omitempty"`
	Language                      []Code                                           `json:"language,omitempty"`
	LanguagePrimitiveElement      []*primitiveElement                              `json:"_language,omitempty"`
	Filter                        []TerminologyCapabilitiesCodeSystemVersionFilter `json:"filter,omitempty"`
	Property                      []Code                                           `json:"property,omitempty"`
	PropertyPrimitiveElement      []*primitiveElement                              `json:"_property,omitempty"`
}

func (r TerminologyCapabilitiesCodeSystemVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TerminologyCapabilitiesCodeSystemVersion) marshalJSON() jsonTerminologyCapabilitiesCodeSystemVersion {
	m := jsonTerminologyCapabilitiesCodeSystemVersion{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	if r.Code != nil && (r.Code.Id != nil || r.Code.Extension != nil) {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.IsDefault = r.IsDefault
	if r.IsDefault != nil && (r.IsDefault.Id != nil || r.IsDefault.Extension != nil) {
		m.IsDefaultPrimitiveElement = &primitiveElement{Id: r.IsDefault.Id, Extension: r.IsDefault.Extension}
	}
	m.Compositional = r.Compositional
	if r.Compositional != nil && (r.Compositional.Id != nil || r.Compositional.Extension != nil) {
		m.CompositionalPrimitiveElement = &primitiveElement{Id: r.Compositional.Id, Extension: r.Compositional.Extension}
	}
	m.Language = r.Language
	anyLanguageIdOrExtension := false
	for _, e := range r.Language {
		if e.Id != nil || e.Extension != nil {
			anyLanguageIdOrExtension = true
			break
		}
	}
	if anyLanguageIdOrExtension {
		m.LanguagePrimitiveElement = make([]*primitiveElement, 0, len(r.Language))
		for _, e := range r.Language {
			if e.Id != nil || e.Extension != nil {
				m.LanguagePrimitiveElement = append(m.LanguagePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.LanguagePrimitiveElement = append(m.LanguagePrimitiveElement, nil)
			}
		}
	}
	m.Filter = r.Filter
	m.Property = r.Property
	anyPropertyIdOrExtension := false
	for _, e := range r.Property {
		if e.Id != nil || e.Extension != nil {
			anyPropertyIdOrExtension = true
			break
		}
	}
	if anyPropertyIdOrExtension {
		m.PropertyPrimitiveElement = make([]*primitiveElement, 0, len(r.Property))
		for _, e := range r.Property {
			if e.Id != nil || e.Extension != nil {
				m.PropertyPrimitiveElement = append(m.PropertyPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.PropertyPrimitiveElement = append(m.PropertyPrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *TerminologyCapabilitiesCodeSystemVersion) UnmarshalJSON(b []byte) error {
	var m jsonTerminologyCapabilitiesCodeSystemVersion
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TerminologyCapabilitiesCodeSystemVersion) unmarshalJSON(m jsonTerminologyCapabilitiesCodeSystemVersion) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.IsDefault = m.IsDefault
	if m.IsDefaultPrimitiveElement != nil {
		r.IsDefault.Id = m.IsDefaultPrimitiveElement.Id
		r.IsDefault.Extension = m.IsDefaultPrimitiveElement.Extension
	}
	r.Compositional = m.Compositional
	if m.CompositionalPrimitiveElement != nil {
		r.Compositional.Id = m.CompositionalPrimitiveElement.Id
		r.Compositional.Extension = m.CompositionalPrimitiveElement.Extension
	}
	r.Language = m.Language
	for i, e := range m.LanguagePrimitiveElement {
		if len(r.Language) > i {
			r.Language[i].Id = e.Id
			r.Language[i].Extension = e.Extension
		} else {
			r.Language = append(r.Language, Code{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Filter = m.Filter
	r.Property = m.Property
	for i, e := range m.PropertyPrimitiveElement {
		if len(r.Property) > i {
			r.Property[i].Id = e.Id
			r.Property[i].Extension = e.Extension
		} else {
			r.Property = append(r.Property, Code{Id: e.Id, Extension: e.Extension})
		}
	}
	return nil
}
func (r TerminologyCapabilitiesCodeSystemVersion) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Filter Properties supported.
type TerminologyCapabilitiesCodeSystemVersionFilter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Code of the property supported.
	Code Code
	// Operations supported for the property.
	Op []Code
}
type jsonTerminologyCapabilitiesCodeSystemVersionFilter struct {
	Id                   *string             `json:"id,omitempty"`
	Extension            []Extension         `json:"extension,omitempty"`
	ModifierExtension    []Extension         `json:"modifierExtension,omitempty"`
	Code                 Code                `json:"code,omitempty"`
	CodePrimitiveElement *primitiveElement   `json:"_code,omitempty"`
	Op                   []Code              `json:"op,omitempty"`
	OpPrimitiveElement   []*primitiveElement `json:"_op,omitempty"`
}

func (r TerminologyCapabilitiesCodeSystemVersionFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TerminologyCapabilitiesCodeSystemVersionFilter) marshalJSON() jsonTerminologyCapabilitiesCodeSystemVersionFilter {
	m := jsonTerminologyCapabilitiesCodeSystemVersionFilter{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Op = r.Op
	anyOpIdOrExtension := false
	for _, e := range r.Op {
		if e.Id != nil || e.Extension != nil {
			anyOpIdOrExtension = true
			break
		}
	}
	if anyOpIdOrExtension {
		m.OpPrimitiveElement = make([]*primitiveElement, 0, len(r.Op))
		for _, e := range r.Op {
			if e.Id != nil || e.Extension != nil {
				m.OpPrimitiveElement = append(m.OpPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.OpPrimitiveElement = append(m.OpPrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *TerminologyCapabilitiesCodeSystemVersionFilter) UnmarshalJSON(b []byte) error {
	var m jsonTerminologyCapabilitiesCodeSystemVersionFilter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TerminologyCapabilitiesCodeSystemVersionFilter) unmarshalJSON(m jsonTerminologyCapabilitiesCodeSystemVersionFilter) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Op = m.Op
	for i, e := range m.OpPrimitiveElement {
		if len(r.Op) > i {
			r.Op[i].Id = e.Id
			r.Op[i].Extension = e.Extension
		} else {
			r.Op = append(r.Op, Code{Id: e.Id, Extension: e.Extension})
		}
	}
	return nil
}
func (r TerminologyCapabilitiesCodeSystemVersionFilter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about the [ValueSet/$expand](valueset-operation-expand.html) operation.
type TerminologyCapabilitiesExpansion struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Whether the server can return nested value sets.
	Hierarchical *Boolean
	// Whether the server supports paging on expansion.
	Paging *Boolean
	// Allow request for incomplete expansions?
	Incomplete *Boolean
	// Supported expansion parameter.
	Parameter []TerminologyCapabilitiesExpansionParameter
	// Documentation about text searching works.
	TextFilter *Markdown
}
type jsonTerminologyCapabilitiesExpansion struct {
	Id                           *string                                     `json:"id,omitempty"`
	Extension                    []Extension                                 `json:"extension,omitempty"`
	ModifierExtension            []Extension                                 `json:"modifierExtension,omitempty"`
	Hierarchical                 *Boolean                                    `json:"hierarchical,omitempty"`
	HierarchicalPrimitiveElement *primitiveElement                           `json:"_hierarchical,omitempty"`
	Paging                       *Boolean                                    `json:"paging,omitempty"`
	PagingPrimitiveElement       *primitiveElement                           `json:"_paging,omitempty"`
	Incomplete                   *Boolean                                    `json:"incomplete,omitempty"`
	IncompletePrimitiveElement   *primitiveElement                           `json:"_incomplete,omitempty"`
	Parameter                    []TerminologyCapabilitiesExpansionParameter `json:"parameter,omitempty"`
	TextFilter                   *Markdown                                   `json:"textFilter,omitempty"`
	TextFilterPrimitiveElement   *primitiveElement                           `json:"_textFilter,omitempty"`
}

func (r TerminologyCapabilitiesExpansion) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TerminologyCapabilitiesExpansion) marshalJSON() jsonTerminologyCapabilitiesExpansion {
	m := jsonTerminologyCapabilitiesExpansion{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Hierarchical = r.Hierarchical
	if r.Hierarchical != nil && (r.Hierarchical.Id != nil || r.Hierarchical.Extension != nil) {
		m.HierarchicalPrimitiveElement = &primitiveElement{Id: r.Hierarchical.Id, Extension: r.Hierarchical.Extension}
	}
	m.Paging = r.Paging
	if r.Paging != nil && (r.Paging.Id != nil || r.Paging.Extension != nil) {
		m.PagingPrimitiveElement = &primitiveElement{Id: r.Paging.Id, Extension: r.Paging.Extension}
	}
	m.Incomplete = r.Incomplete
	if r.Incomplete != nil && (r.Incomplete.Id != nil || r.Incomplete.Extension != nil) {
		m.IncompletePrimitiveElement = &primitiveElement{Id: r.Incomplete.Id, Extension: r.Incomplete.Extension}
	}
	m.Parameter = r.Parameter
	m.TextFilter = r.TextFilter
	if r.TextFilter != nil && (r.TextFilter.Id != nil || r.TextFilter.Extension != nil) {
		m.TextFilterPrimitiveElement = &primitiveElement{Id: r.TextFilter.Id, Extension: r.TextFilter.Extension}
	}
	return m
}
func (r *TerminologyCapabilitiesExpansion) UnmarshalJSON(b []byte) error {
	var m jsonTerminologyCapabilitiesExpansion
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TerminologyCapabilitiesExpansion) unmarshalJSON(m jsonTerminologyCapabilitiesExpansion) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Hierarchical = m.Hierarchical
	if m.HierarchicalPrimitiveElement != nil {
		r.Hierarchical.Id = m.HierarchicalPrimitiveElement.Id
		r.Hierarchical.Extension = m.HierarchicalPrimitiveElement.Extension
	}
	r.Paging = m.Paging
	if m.PagingPrimitiveElement != nil {
		r.Paging.Id = m.PagingPrimitiveElement.Id
		r.Paging.Extension = m.PagingPrimitiveElement.Extension
	}
	r.Incomplete = m.Incomplete
	if m.IncompletePrimitiveElement != nil {
		r.Incomplete.Id = m.IncompletePrimitiveElement.Id
		r.Incomplete.Extension = m.IncompletePrimitiveElement.Extension
	}
	r.Parameter = m.Parameter
	r.TextFilter = m.TextFilter
	if m.TextFilterPrimitiveElement != nil {
		r.TextFilter.Id = m.TextFilterPrimitiveElement.Id
		r.TextFilter.Extension = m.TextFilterPrimitiveElement.Extension
	}
	return nil
}
func (r TerminologyCapabilitiesExpansion) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Supported expansion parameter.
type TerminologyCapabilitiesExpansionParameter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Expansion Parameter name.
	Name Code
	// Description of support for parameter.
	Documentation *String
}
type jsonTerminologyCapabilitiesExpansionParameter struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	Name                          Code              `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement `json:"_name,omitempty"`
	Documentation                 *String           `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement `json:"_documentation,omitempty"`
}

func (r TerminologyCapabilitiesExpansionParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TerminologyCapabilitiesExpansionParameter) marshalJSON() jsonTerminologyCapabilitiesExpansionParameter {
	m := jsonTerminologyCapabilitiesExpansionParameter{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Name = r.Name
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Documentation = r.Documentation
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	return m
}
func (r *TerminologyCapabilitiesExpansionParameter) UnmarshalJSON(b []byte) error {
	var m jsonTerminologyCapabilitiesExpansionParameter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TerminologyCapabilitiesExpansionParameter) unmarshalJSON(m jsonTerminologyCapabilitiesExpansionParameter) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	return nil
}
func (r TerminologyCapabilitiesExpansionParameter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about the [ValueSet/$validate-code](valueset-operation-validate-code.html) operation.
type TerminologyCapabilitiesValidateCode struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Whether translations are validated.
	Translations Boolean
}
type jsonTerminologyCapabilitiesValidateCode struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Translations                 Boolean           `json:"translations,omitempty"`
	TranslationsPrimitiveElement *primitiveElement `json:"_translations,omitempty"`
}

func (r TerminologyCapabilitiesValidateCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TerminologyCapabilitiesValidateCode) marshalJSON() jsonTerminologyCapabilitiesValidateCode {
	m := jsonTerminologyCapabilitiesValidateCode{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Translations = r.Translations
	if r.Translations.Id != nil || r.Translations.Extension != nil {
		m.TranslationsPrimitiveElement = &primitiveElement{Id: r.Translations.Id, Extension: r.Translations.Extension}
	}
	return m
}
func (r *TerminologyCapabilitiesValidateCode) UnmarshalJSON(b []byte) error {
	var m jsonTerminologyCapabilitiesValidateCode
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TerminologyCapabilitiesValidateCode) unmarshalJSON(m jsonTerminologyCapabilitiesValidateCode) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Translations = m.Translations
	if m.TranslationsPrimitiveElement != nil {
		r.Translations.Id = m.TranslationsPrimitiveElement.Id
		r.Translations.Extension = m.TranslationsPrimitiveElement.Extension
	}
	return nil
}
func (r TerminologyCapabilitiesValidateCode) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about the [ConceptMap/$translate](conceptmap-operation-translate.html) operation.
type TerminologyCapabilitiesTranslation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Whether the client must identify the map.
	NeedsMap Boolean
}
type jsonTerminologyCapabilitiesTranslation struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	NeedsMap                 Boolean           `json:"needsMap,omitempty"`
	NeedsMapPrimitiveElement *primitiveElement `json:"_needsMap,omitempty"`
}

func (r TerminologyCapabilitiesTranslation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TerminologyCapabilitiesTranslation) marshalJSON() jsonTerminologyCapabilitiesTranslation {
	m := jsonTerminologyCapabilitiesTranslation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.NeedsMap = r.NeedsMap
	if r.NeedsMap.Id != nil || r.NeedsMap.Extension != nil {
		m.NeedsMapPrimitiveElement = &primitiveElement{Id: r.NeedsMap.Id, Extension: r.NeedsMap.Extension}
	}
	return m
}
func (r *TerminologyCapabilitiesTranslation) UnmarshalJSON(b []byte) error {
	var m jsonTerminologyCapabilitiesTranslation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TerminologyCapabilitiesTranslation) unmarshalJSON(m jsonTerminologyCapabilitiesTranslation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.NeedsMap = m.NeedsMap
	if m.NeedsMapPrimitiveElement != nil {
		r.NeedsMap.Id = m.NeedsMapPrimitiveElement.Id
		r.NeedsMap.Extension = m.NeedsMapPrimitiveElement.Extension
	}
	return nil
}
func (r TerminologyCapabilitiesTranslation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Whether the $closure operation is supported.
type TerminologyCapabilitiesClosure struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// If cross-system closure is supported.
	Translation *Boolean
}
type jsonTerminologyCapabilitiesClosure struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Translation                 *Boolean          `json:"translation,omitempty"`
	TranslationPrimitiveElement *primitiveElement `json:"_translation,omitempty"`
}

func (r TerminologyCapabilitiesClosure) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TerminologyCapabilitiesClosure) marshalJSON() jsonTerminologyCapabilitiesClosure {
	m := jsonTerminologyCapabilitiesClosure{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Translation = r.Translation
	if r.Translation != nil && (r.Translation.Id != nil || r.Translation.Extension != nil) {
		m.TranslationPrimitiveElement = &primitiveElement{Id: r.Translation.Id, Extension: r.Translation.Extension}
	}
	return m
}
func (r *TerminologyCapabilitiesClosure) UnmarshalJSON(b []byte) error {
	var m jsonTerminologyCapabilitiesClosure
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TerminologyCapabilitiesClosure) unmarshalJSON(m jsonTerminologyCapabilitiesClosure) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Translation = m.Translation
	if m.TranslationPrimitiveElement != nil {
		r.Translation.Id = m.TranslationPrimitiveElement.Id
		r.Translation.Extension = m.TranslationPrimitiveElement.Extension
	}
	return nil
}
func (r TerminologyCapabilitiesClosure) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
