package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// A Capability Statement documents a set of capabilities (behaviors) of a FHIR Server for a particular version of FHIR that may be used as a statement of actual server functionality or a statement of required or desired server implementation.
type CapabilityStatement struct {
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
	// An absolute URI that is used to identify this capability statement when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this capability statement is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the capability statement is stored on different servers.
	Url *Uri
	// The identifier that is used to identify this version of the capability statement when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the capability statement author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the capability statement. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the capability statement.
	Title *String
	// The status of this capability statement. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this capability statement is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the capability statement was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the capability statement changes.
	Date DateTime
	// The name of the organization or individual that published the capability statement.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the capability statement from a consumer's perspective. Typically, this is used when the capability statement describes a desired rather than an actual solution, for example as a formal expression of requirements as part of an RFP.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate capability statement instances.
	UseContext []UsageContext
	// A legal or geographic region in which the capability statement is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this capability statement is needed and why it has been designed as it has.
	Purpose *Markdown
	// A copyright statement relating to the capability statement and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the capability statement.
	Copyright *Markdown
	// The way that this statement is intended to be used, to describe an actual running instance of software, a particular product (kind, not instance of software) or a class of implementation (e.g. a desired purchase).
	Kind Code
	// Reference to a canonical URL of another CapabilityStatement that this software implements. This capability statement is a published API description that corresponds to a business service. The server may actually implement a subset of the capability statement it claims to implement, so the capability statement must specify the full capability details.
	Instantiates []Canonical
	// Reference to a canonical URL of another CapabilityStatement that this software adds to. The capability statement automatically includes everything in the other statement, and it is not duplicated, though the server may repeat the same resources, interactions and operations to add additional details to them.
	Imports []Canonical
	// Software that is covered by this capability statement.  It is used when the capability statement describes the capabilities of a particular software version, independent of an installation.
	Software *CapabilityStatementSoftware
	// Identifies a specific implementation instance that is described by the capability statement - i.e. a particular installation, rather than the capabilities of a software program.
	Implementation *CapabilityStatementImplementation
	// The version of the FHIR specification that this CapabilityStatement describes (which SHALL be the same as the FHIR version of the CapabilityStatement itself). There is no default value.
	FhirVersion Code
	// A list of the formats supported by this implementation using their content types.
	Format []Code
	// A list of the patch formats supported by this implementation using their content types.
	PatchFormat []Code
	// A list of implementation guides that the server does (or should) support in their entirety.
	ImplementationGuide []Canonical
	// A definition of the restful capabilities of the solution, if any.
	Rest []CapabilityStatementRest
	// A description of the messaging capabilities of the solution.
	Messaging []CapabilityStatementMessaging
	// A document definition.
	Document []CapabilityStatementDocument
}

func (r CapabilityStatement) ResourceType() string {
	return "CapabilityStatement"
}
func (r CapabilityStatement) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonCapabilityStatement struct {
	ResourceType                        string                             `json:"resourceType"`
	Id                                  *Id                                `json:"id,omitempty"`
	IdPrimitiveElement                  *primitiveElement                  `json:"_id,omitempty"`
	Meta                                *Meta                              `json:"meta,omitempty"`
	ImplicitRules                       *Uri                               `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement       *primitiveElement                  `json:"_implicitRules,omitempty"`
	Language                            *Code                              `json:"language,omitempty"`
	LanguagePrimitiveElement            *primitiveElement                  `json:"_language,omitempty"`
	Text                                *Narrative                         `json:"text,omitempty"`
	Contained                           []ContainedResource                `json:"contained,omitempty"`
	Extension                           []Extension                        `json:"extension,omitempty"`
	ModifierExtension                   []Extension                        `json:"modifierExtension,omitempty"`
	Url                                 *Uri                               `json:"url,omitempty"`
	UrlPrimitiveElement                 *primitiveElement                  `json:"_url,omitempty"`
	Version                             *String                            `json:"version,omitempty"`
	VersionPrimitiveElement             *primitiveElement                  `json:"_version,omitempty"`
	Name                                *String                            `json:"name,omitempty"`
	NamePrimitiveElement                *primitiveElement                  `json:"_name,omitempty"`
	Title                               *String                            `json:"title,omitempty"`
	TitlePrimitiveElement               *primitiveElement                  `json:"_title,omitempty"`
	Status                              Code                               `json:"status,omitempty"`
	StatusPrimitiveElement              *primitiveElement                  `json:"_status,omitempty"`
	Experimental                        *Boolean                           `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement        *primitiveElement                  `json:"_experimental,omitempty"`
	Date                                DateTime                           `json:"date,omitempty"`
	DatePrimitiveElement                *primitiveElement                  `json:"_date,omitempty"`
	Publisher                           *String                            `json:"publisher,omitempty"`
	PublisherPrimitiveElement           *primitiveElement                  `json:"_publisher,omitempty"`
	Contact                             []ContactDetail                    `json:"contact,omitempty"`
	Description                         *Markdown                          `json:"description,omitempty"`
	DescriptionPrimitiveElement         *primitiveElement                  `json:"_description,omitempty"`
	UseContext                          []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction                        []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Purpose                             *Markdown                          `json:"purpose,omitempty"`
	PurposePrimitiveElement             *primitiveElement                  `json:"_purpose,omitempty"`
	Copyright                           *Markdown                          `json:"copyright,omitempty"`
	CopyrightPrimitiveElement           *primitiveElement                  `json:"_copyright,omitempty"`
	Kind                                Code                               `json:"kind,omitempty"`
	KindPrimitiveElement                *primitiveElement                  `json:"_kind,omitempty"`
	Instantiates                        []Canonical                        `json:"instantiates,omitempty"`
	InstantiatesPrimitiveElement        []*primitiveElement                `json:"_instantiates,omitempty"`
	Imports                             []Canonical                        `json:"imports,omitempty"`
	ImportsPrimitiveElement             []*primitiveElement                `json:"_imports,omitempty"`
	Software                            *CapabilityStatementSoftware       `json:"software,omitempty"`
	Implementation                      *CapabilityStatementImplementation `json:"implementation,omitempty"`
	FhirVersion                         Code                               `json:"fhirVersion,omitempty"`
	FhirVersionPrimitiveElement         *primitiveElement                  `json:"_fhirVersion,omitempty"`
	Format                              []Code                             `json:"format,omitempty"`
	FormatPrimitiveElement              []*primitiveElement                `json:"_format,omitempty"`
	PatchFormat                         []Code                             `json:"patchFormat,omitempty"`
	PatchFormatPrimitiveElement         []*primitiveElement                `json:"_patchFormat,omitempty"`
	ImplementationGuide                 []Canonical                        `json:"implementationGuide,omitempty"`
	ImplementationGuidePrimitiveElement []*primitiveElement                `json:"_implementationGuide,omitempty"`
	Rest                                []CapabilityStatementRest          `json:"rest,omitempty"`
	Messaging                           []CapabilityStatementMessaging     `json:"messaging,omitempty"`
	Document                            []CapabilityStatementDocument      `json:"document,omitempty"`
}

func (r CapabilityStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CapabilityStatement) marshalJSON() jsonCapabilityStatement {
	m := jsonCapabilityStatement{}
	m.ResourceType = "CapabilityStatement"
	if r.Id != nil && r.Id.Value != nil {
		m.Id = r.Id
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		m.ImplicitRules = r.ImplicitRules
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	if r.Language != nil && r.Language.Value != nil {
		m.Language = r.Language
	}
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
	if r.Url != nil && r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	if r.Version != nil && r.Version.Value != nil {
		m.Version = r.Version
	}
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Title != nil && r.Title.Value != nil {
		m.Title = r.Title
	}
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	if r.Experimental != nil && r.Experimental.Value != nil {
		m.Experimental = r.Experimental
	}
	if r.Experimental != nil && (r.Experimental.Id != nil || r.Experimental.Extension != nil) {
		m.ExperimentalPrimitiveElement = &primitiveElement{Id: r.Experimental.Id, Extension: r.Experimental.Extension}
	}
	if r.Date.Value != nil {
		m.Date = r.Date
	}
	if r.Date.Id != nil || r.Date.Extension != nil {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	if r.Publisher != nil && r.Publisher.Value != nil {
		m.Publisher = r.Publisher
	}
	if r.Publisher != nil && (r.Publisher.Id != nil || r.Publisher.Extension != nil) {
		m.PublisherPrimitiveElement = &primitiveElement{Id: r.Publisher.Id, Extension: r.Publisher.Extension}
	}
	m.Contact = r.Contact
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.UseContext = r.UseContext
	m.Jurisdiction = r.Jurisdiction
	if r.Purpose != nil && r.Purpose.Value != nil {
		m.Purpose = r.Purpose
	}
	if r.Purpose != nil && (r.Purpose.Id != nil || r.Purpose.Extension != nil) {
		m.PurposePrimitiveElement = &primitiveElement{Id: r.Purpose.Id, Extension: r.Purpose.Extension}
	}
	if r.Copyright != nil && r.Copyright.Value != nil {
		m.Copyright = r.Copyright
	}
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		m.CopyrightPrimitiveElement = &primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
	}
	if r.Kind.Value != nil {
		m.Kind = r.Kind
	}
	if r.Kind.Id != nil || r.Kind.Extension != nil {
		m.KindPrimitiveElement = &primitiveElement{Id: r.Kind.Id, Extension: r.Kind.Extension}
	}
	anyInstantiatesValue := false
	for _, e := range r.Instantiates {
		if e.Value != nil {
			anyInstantiatesValue = true
			break
		}
	}
	if anyInstantiatesValue {
		m.Instantiates = r.Instantiates
	}
	anyInstantiatesIdOrExtension := false
	for _, e := range r.Instantiates {
		if e.Id != nil || e.Extension != nil {
			anyInstantiatesIdOrExtension = true
			break
		}
	}
	if anyInstantiatesIdOrExtension {
		m.InstantiatesPrimitiveElement = make([]*primitiveElement, 0, len(r.Instantiates))
		for _, e := range r.Instantiates {
			if e.Id != nil || e.Extension != nil {
				m.InstantiatesPrimitiveElement = append(m.InstantiatesPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.InstantiatesPrimitiveElement = append(m.InstantiatesPrimitiveElement, nil)
			}
		}
	}
	anyImportsValue := false
	for _, e := range r.Imports {
		if e.Value != nil {
			anyImportsValue = true
			break
		}
	}
	if anyImportsValue {
		m.Imports = r.Imports
	}
	anyImportsIdOrExtension := false
	for _, e := range r.Imports {
		if e.Id != nil || e.Extension != nil {
			anyImportsIdOrExtension = true
			break
		}
	}
	if anyImportsIdOrExtension {
		m.ImportsPrimitiveElement = make([]*primitiveElement, 0, len(r.Imports))
		for _, e := range r.Imports {
			if e.Id != nil || e.Extension != nil {
				m.ImportsPrimitiveElement = append(m.ImportsPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ImportsPrimitiveElement = append(m.ImportsPrimitiveElement, nil)
			}
		}
	}
	m.Software = r.Software
	m.Implementation = r.Implementation
	if r.FhirVersion.Value != nil {
		m.FhirVersion = r.FhirVersion
	}
	if r.FhirVersion.Id != nil || r.FhirVersion.Extension != nil {
		m.FhirVersionPrimitiveElement = &primitiveElement{Id: r.FhirVersion.Id, Extension: r.FhirVersion.Extension}
	}
	anyFormatValue := false
	for _, e := range r.Format {
		if e.Value != nil {
			anyFormatValue = true
			break
		}
	}
	if anyFormatValue {
		m.Format = r.Format
	}
	anyFormatIdOrExtension := false
	for _, e := range r.Format {
		if e.Id != nil || e.Extension != nil {
			anyFormatIdOrExtension = true
			break
		}
	}
	if anyFormatIdOrExtension {
		m.FormatPrimitiveElement = make([]*primitiveElement, 0, len(r.Format))
		for _, e := range r.Format {
			if e.Id != nil || e.Extension != nil {
				m.FormatPrimitiveElement = append(m.FormatPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.FormatPrimitiveElement = append(m.FormatPrimitiveElement, nil)
			}
		}
	}
	anyPatchFormatValue := false
	for _, e := range r.PatchFormat {
		if e.Value != nil {
			anyPatchFormatValue = true
			break
		}
	}
	if anyPatchFormatValue {
		m.PatchFormat = r.PatchFormat
	}
	anyPatchFormatIdOrExtension := false
	for _, e := range r.PatchFormat {
		if e.Id != nil || e.Extension != nil {
			anyPatchFormatIdOrExtension = true
			break
		}
	}
	if anyPatchFormatIdOrExtension {
		m.PatchFormatPrimitiveElement = make([]*primitiveElement, 0, len(r.PatchFormat))
		for _, e := range r.PatchFormat {
			if e.Id != nil || e.Extension != nil {
				m.PatchFormatPrimitiveElement = append(m.PatchFormatPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.PatchFormatPrimitiveElement = append(m.PatchFormatPrimitiveElement, nil)
			}
		}
	}
	anyImplementationGuideValue := false
	for _, e := range r.ImplementationGuide {
		if e.Value != nil {
			anyImplementationGuideValue = true
			break
		}
	}
	if anyImplementationGuideValue {
		m.ImplementationGuide = r.ImplementationGuide
	}
	anyImplementationGuideIdOrExtension := false
	for _, e := range r.ImplementationGuide {
		if e.Id != nil || e.Extension != nil {
			anyImplementationGuideIdOrExtension = true
			break
		}
	}
	if anyImplementationGuideIdOrExtension {
		m.ImplementationGuidePrimitiveElement = make([]*primitiveElement, 0, len(r.ImplementationGuide))
		for _, e := range r.ImplementationGuide {
			if e.Id != nil || e.Extension != nil {
				m.ImplementationGuidePrimitiveElement = append(m.ImplementationGuidePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ImplementationGuidePrimitiveElement = append(m.ImplementationGuidePrimitiveElement, nil)
			}
		}
	}
	m.Rest = r.Rest
	m.Messaging = r.Messaging
	m.Document = r.Document
	return m
}
func (r *CapabilityStatement) UnmarshalJSON(b []byte) error {
	var m jsonCapabilityStatement
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CapabilityStatement) unmarshalJSON(m jsonCapabilityStatement) error {
	r.Id = m.Id
	if m.IdPrimitiveElement != nil {
		if r.Id == nil {
			r.Id = &Id{}
		}
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		if r.ImplicitRules == nil {
			r.ImplicitRules = &Uri{}
		}
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		if r.Language == nil {
			r.Language = &Code{}
		}
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
		if r.Url == nil {
			r.Url = &Uri{}
		}
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		if r.Version == nil {
			r.Version = &String{}
		}
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		if r.Title == nil {
			r.Title = &String{}
		}
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
		if r.Experimental == nil {
			r.Experimental = &Boolean{}
		}
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
		if r.Publisher == nil {
			r.Publisher = &String{}
		}
		r.Publisher.Id = m.PublisherPrimitiveElement.Id
		r.Publisher.Extension = m.PublisherPrimitiveElement.Extension
	}
	r.Contact = m.Contact
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &Markdown{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.UseContext = m.UseContext
	r.Jurisdiction = m.Jurisdiction
	r.Purpose = m.Purpose
	if m.PurposePrimitiveElement != nil {
		if r.Purpose == nil {
			r.Purpose = &Markdown{}
		}
		r.Purpose.Id = m.PurposePrimitiveElement.Id
		r.Purpose.Extension = m.PurposePrimitiveElement.Extension
	}
	r.Copyright = m.Copyright
	if m.CopyrightPrimitiveElement != nil {
		if r.Copyright == nil {
			r.Copyright = &Markdown{}
		}
		r.Copyright.Id = m.CopyrightPrimitiveElement.Id
		r.Copyright.Extension = m.CopyrightPrimitiveElement.Extension
	}
	r.Kind = m.Kind
	if m.KindPrimitiveElement != nil {
		r.Kind.Id = m.KindPrimitiveElement.Id
		r.Kind.Extension = m.KindPrimitiveElement.Extension
	}
	r.Instantiates = m.Instantiates
	for i, e := range m.InstantiatesPrimitiveElement {
		if len(r.Instantiates) <= i {
			r.Instantiates = append(r.Instantiates, Canonical{})
		}
		if e != nil {
			r.Instantiates[i].Id = e.Id
			r.Instantiates[i].Extension = e.Extension
		}
	}
	r.Imports = m.Imports
	for i, e := range m.ImportsPrimitiveElement {
		if len(r.Imports) <= i {
			r.Imports = append(r.Imports, Canonical{})
		}
		if e != nil {
			r.Imports[i].Id = e.Id
			r.Imports[i].Extension = e.Extension
		}
	}
	r.Software = m.Software
	r.Implementation = m.Implementation
	r.FhirVersion = m.FhirVersion
	if m.FhirVersionPrimitiveElement != nil {
		r.FhirVersion.Id = m.FhirVersionPrimitiveElement.Id
		r.FhirVersion.Extension = m.FhirVersionPrimitiveElement.Extension
	}
	r.Format = m.Format
	for i, e := range m.FormatPrimitiveElement {
		if len(r.Format) <= i {
			r.Format = append(r.Format, Code{})
		}
		if e != nil {
			r.Format[i].Id = e.Id
			r.Format[i].Extension = e.Extension
		}
	}
	r.PatchFormat = m.PatchFormat
	for i, e := range m.PatchFormatPrimitiveElement {
		if len(r.PatchFormat) <= i {
			r.PatchFormat = append(r.PatchFormat, Code{})
		}
		if e != nil {
			r.PatchFormat[i].Id = e.Id
			r.PatchFormat[i].Extension = e.Extension
		}
	}
	r.ImplementationGuide = m.ImplementationGuide
	for i, e := range m.ImplementationGuidePrimitiveElement {
		if len(r.ImplementationGuide) <= i {
			r.ImplementationGuide = append(r.ImplementationGuide, Canonical{})
		}
		if e != nil {
			r.ImplementationGuide[i].Id = e.Id
			r.ImplementationGuide[i].Extension = e.Extension
		}
	}
	r.Rest = m.Rest
	r.Messaging = m.Messaging
	r.Document = m.Document
	return nil
}
func (r CapabilityStatement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "CapabilityStatement"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Id, xml.StartElement{Name: xml.Name{Local: "id"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Meta, xml.StartElement{Name: xml.Name{Local: "meta"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ImplicitRules, xml.StartElement{Name: xml.Name{Local: "implicitRules"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	v := make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		v = append(v, ContainedResource{c})
	}
	err = e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contained"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Version, xml.StartElement{Name: xml.Name{Local: "version"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Title, xml.StartElement{Name: xml.Name{Local: "title"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Experimental, xml.StartElement{Name: xml.Name{Local: "experimental"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Publisher, xml.StartElement{Name: xml.Name{Local: "publisher"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Contact, xml.StartElement{Name: xml.Name{Local: "contact"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UseContext, xml.StartElement{Name: xml.Name{Local: "useContext"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Jurisdiction, xml.StartElement{Name: xml.Name{Local: "jurisdiction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Purpose, xml.StartElement{Name: xml.Name{Local: "purpose"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Copyright, xml.StartElement{Name: xml.Name{Local: "copyright"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Kind, xml.StartElement{Name: xml.Name{Local: "kind"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Instantiates, xml.StartElement{Name: xml.Name{Local: "instantiates"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Imports, xml.StartElement{Name: xml.Name{Local: "imports"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Software, xml.StartElement{Name: xml.Name{Local: "software"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Implementation, xml.StartElement{Name: xml.Name{Local: "implementation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FhirVersion, xml.StartElement{Name: xml.Name{Local: "fhirVersion"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Format, xml.StartElement{Name: xml.Name{Local: "format"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PatchFormat, xml.StartElement{Name: xml.Name{Local: "patchFormat"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ImplementationGuide, xml.StartElement{Name: xml.Name{Local: "implementationGuide"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Rest, xml.StartElement{Name: xml.Name{Local: "rest"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Messaging, xml.StartElement{Name: xml.Name{Local: "messaging"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Document, xml.StartElement{Name: xml.Name{Local: "document"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CapabilityStatement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "id":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Id = &v
			case "meta":
				var v Meta
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Meta = &v
			case "implicitRules":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ImplicitRules = &v
			case "language":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = &v
			case "text":
				var v Narrative
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = &v
			case "contained":
				var c ContainedResource
				err := d.DecodeElement(&c, &t)
				if err != nil {
					return err
				}
				r.Contained = append(r.Contained, c.Resource)
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "url":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Url = &v
			case "version":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Version = &v
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "title":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Title = &v
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "experimental":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Experimental = &v
			case "date":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = v
			case "publisher":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Publisher = &v
			case "contact":
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Contact = append(r.Contact, v)
			case "description":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "useContext":
				var v UsageContext
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UseContext = append(r.UseContext, v)
			case "jurisdiction":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Jurisdiction = append(r.Jurisdiction, v)
			case "purpose":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Purpose = &v
			case "copyright":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Copyright = &v
			case "kind":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Kind = v
			case "instantiates":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Instantiates = append(r.Instantiates, v)
			case "imports":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Imports = append(r.Imports, v)
			case "software":
				var v CapabilityStatementSoftware
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Software = &v
			case "implementation":
				var v CapabilityStatementImplementation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Implementation = &v
			case "fhirVersion":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FhirVersion = v
			case "format":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Format = append(r.Format, v)
			case "patchFormat":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PatchFormat = append(r.PatchFormat, v)
			case "implementationGuide":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ImplementationGuide = append(r.ImplementationGuide, v)
			case "rest":
				var v CapabilityStatementRest
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Rest = append(r.Rest, v)
			case "messaging":
				var v CapabilityStatementMessaging
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Messaging = append(r.Messaging, v)
			case "document":
				var v CapabilityStatementDocument
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Document = append(r.Document, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CapabilityStatement) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Software that is covered by this capability statement.  It is used when the capability statement describes the capabilities of a particular software version, independent of an installation.
type CapabilityStatementSoftware struct {
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
	// Date this version of the software was released.
	ReleaseDate *DateTime
}
type jsonCapabilityStatementSoftware struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Name                        String            `json:"name,omitempty"`
	NamePrimitiveElement        *primitiveElement `json:"_name,omitempty"`
	Version                     *String           `json:"version,omitempty"`
	VersionPrimitiveElement     *primitiveElement `json:"_version,omitempty"`
	ReleaseDate                 *DateTime         `json:"releaseDate,omitempty"`
	ReleaseDatePrimitiveElement *primitiveElement `json:"_releaseDate,omitempty"`
}

func (r CapabilityStatementSoftware) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CapabilityStatementSoftware) marshalJSON() jsonCapabilityStatementSoftware {
	m := jsonCapabilityStatementSoftware{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Version != nil && r.Version.Value != nil {
		m.Version = r.Version
	}
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	if r.ReleaseDate != nil && r.ReleaseDate.Value != nil {
		m.ReleaseDate = r.ReleaseDate
	}
	if r.ReleaseDate != nil && (r.ReleaseDate.Id != nil || r.ReleaseDate.Extension != nil) {
		m.ReleaseDatePrimitiveElement = &primitiveElement{Id: r.ReleaseDate.Id, Extension: r.ReleaseDate.Extension}
	}
	return m
}
func (r *CapabilityStatementSoftware) UnmarshalJSON(b []byte) error {
	var m jsonCapabilityStatementSoftware
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CapabilityStatementSoftware) unmarshalJSON(m jsonCapabilityStatementSoftware) error {
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
		if r.Version == nil {
			r.Version = &String{}
		}
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	r.ReleaseDate = m.ReleaseDate
	if m.ReleaseDatePrimitiveElement != nil {
		if r.ReleaseDate == nil {
			r.ReleaseDate = &DateTime{}
		}
		r.ReleaseDate.Id = m.ReleaseDatePrimitiveElement.Id
		r.ReleaseDate.Extension = m.ReleaseDatePrimitiveElement.Extension
	}
	return nil
}
func (r CapabilityStatementSoftware) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Version, xml.StartElement{Name: xml.Name{Local: "version"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReleaseDate, xml.StartElement{Name: xml.Name{Local: "releaseDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CapabilityStatementSoftware) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = v
			case "version":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Version = &v
			case "releaseDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReleaseDate = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CapabilityStatementSoftware) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Identifies a specific implementation instance that is described by the capability statement - i.e. a particular installation, rather than the capabilities of a software program.
type CapabilityStatementImplementation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Information about the specific installation that this capability statement relates to.
	Description String
	// An absolute base URL for the implementation.  This forms the base for REST interfaces as well as the mailbox and document interfaces.
	Url *Url
	// The organization responsible for the management of the instance and oversight of the data on the server at the specified URL.
	Custodian *Reference
}
type jsonCapabilityStatementImplementation struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Description                 String            `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
	Url                         *Url              `json:"url,omitempty"`
	UrlPrimitiveElement         *primitiveElement `json:"_url,omitempty"`
	Custodian                   *Reference        `json:"custodian,omitempty"`
}

func (r CapabilityStatementImplementation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CapabilityStatementImplementation) marshalJSON() jsonCapabilityStatementImplementation {
	m := jsonCapabilityStatementImplementation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description.Id != nil || r.Description.Extension != nil {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	if r.Url != nil && r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	m.Custodian = r.Custodian
	return m
}
func (r *CapabilityStatementImplementation) UnmarshalJSON(b []byte) error {
	var m jsonCapabilityStatementImplementation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CapabilityStatementImplementation) unmarshalJSON(m jsonCapabilityStatementImplementation) error {
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
		if r.Url == nil {
			r.Url = &Url{}
		}
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Custodian = m.Custodian
	return nil
}
func (r CapabilityStatementImplementation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Custodian, xml.StartElement{Name: xml.Name{Local: "custodian"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CapabilityStatementImplementation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = v
			case "url":
				var v Url
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Url = &v
			case "custodian":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Custodian = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CapabilityStatementImplementation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A definition of the restful capabilities of the solution, if any.
type CapabilityStatementRest struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identifies whether this portion of the statement is describing the ability to initiate or receive restful operations.
	Mode Code
	// Information about the system's restful capabilities that apply across all applications, such as security.
	Documentation *Markdown
	// Information about security implementation from an interface perspective - what a client needs to know.
	Security *CapabilityStatementRestSecurity
	// A specification of the restful capabilities of the solution for a specific resource type.
	Resource []CapabilityStatementRestResource
	// A specification of restful operations supported by the system.
	Interaction []CapabilityStatementRestInteraction
	// Search parameters that are supported for searching all resources for implementations to support and/or make use of - either references to ones defined in the specification, or additional ones defined for/by the implementation.
	SearchParam []CapabilityStatementRestResourceSearchParam
	// Definition of an operation or a named query together with its parameters and their meaning and type.
	Operation []CapabilityStatementRestResourceOperation
	// An absolute URI which is a reference to the definition of a compartment that the system supports. The reference is to a CompartmentDefinition resource by its canonical URL .
	Compartment []Canonical
}
type jsonCapabilityStatementRest struct {
	Id                            *string                                      `json:"id,omitempty"`
	Extension                     []Extension                                  `json:"extension,omitempty"`
	ModifierExtension             []Extension                                  `json:"modifierExtension,omitempty"`
	Mode                          Code                                         `json:"mode,omitempty"`
	ModePrimitiveElement          *primitiveElement                            `json:"_mode,omitempty"`
	Documentation                 *Markdown                                    `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement                            `json:"_documentation,omitempty"`
	Security                      *CapabilityStatementRestSecurity             `json:"security,omitempty"`
	Resource                      []CapabilityStatementRestResource            `json:"resource,omitempty"`
	Interaction                   []CapabilityStatementRestInteraction         `json:"interaction,omitempty"`
	SearchParam                   []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`
	Operation                     []CapabilityStatementRestResourceOperation   `json:"operation,omitempty"`
	Compartment                   []Canonical                                  `json:"compartment,omitempty"`
	CompartmentPrimitiveElement   []*primitiveElement                          `json:"_compartment,omitempty"`
}

func (r CapabilityStatementRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CapabilityStatementRest) marshalJSON() jsonCapabilityStatementRest {
	m := jsonCapabilityStatementRest{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Mode.Value != nil {
		m.Mode = r.Mode
	}
	if r.Mode.Id != nil || r.Mode.Extension != nil {
		m.ModePrimitiveElement = &primitiveElement{Id: r.Mode.Id, Extension: r.Mode.Extension}
	}
	if r.Documentation != nil && r.Documentation.Value != nil {
		m.Documentation = r.Documentation
	}
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	m.Security = r.Security
	m.Resource = r.Resource
	m.Interaction = r.Interaction
	m.SearchParam = r.SearchParam
	m.Operation = r.Operation
	anyCompartmentValue := false
	for _, e := range r.Compartment {
		if e.Value != nil {
			anyCompartmentValue = true
			break
		}
	}
	if anyCompartmentValue {
		m.Compartment = r.Compartment
	}
	anyCompartmentIdOrExtension := false
	for _, e := range r.Compartment {
		if e.Id != nil || e.Extension != nil {
			anyCompartmentIdOrExtension = true
			break
		}
	}
	if anyCompartmentIdOrExtension {
		m.CompartmentPrimitiveElement = make([]*primitiveElement, 0, len(r.Compartment))
		for _, e := range r.Compartment {
			if e.Id != nil || e.Extension != nil {
				m.CompartmentPrimitiveElement = append(m.CompartmentPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.CompartmentPrimitiveElement = append(m.CompartmentPrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *CapabilityStatementRest) UnmarshalJSON(b []byte) error {
	var m jsonCapabilityStatementRest
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CapabilityStatementRest) unmarshalJSON(m jsonCapabilityStatementRest) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Mode = m.Mode
	if m.ModePrimitiveElement != nil {
		r.Mode.Id = m.ModePrimitiveElement.Id
		r.Mode.Extension = m.ModePrimitiveElement.Extension
	}
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		if r.Documentation == nil {
			r.Documentation = &Markdown{}
		}
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	r.Security = m.Security
	r.Resource = m.Resource
	r.Interaction = m.Interaction
	r.SearchParam = m.SearchParam
	r.Operation = m.Operation
	r.Compartment = m.Compartment
	for i, e := range m.CompartmentPrimitiveElement {
		if len(r.Compartment) <= i {
			r.Compartment = append(r.Compartment, Canonical{})
		}
		if e != nil {
			r.Compartment[i].Id = e.Id
			r.Compartment[i].Extension = e.Extension
		}
	}
	return nil
}
func (r CapabilityStatementRest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Mode, xml.StartElement{Name: xml.Name{Local: "mode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Documentation, xml.StartElement{Name: xml.Name{Local: "documentation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Security, xml.StartElement{Name: xml.Name{Local: "security"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Resource, xml.StartElement{Name: xml.Name{Local: "resource"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Interaction, xml.StartElement{Name: xml.Name{Local: "interaction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SearchParam, xml.StartElement{Name: xml.Name{Local: "searchParam"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Operation, xml.StartElement{Name: xml.Name{Local: "operation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Compartment, xml.StartElement{Name: xml.Name{Local: "compartment"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CapabilityStatementRest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "mode":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Mode = v
			case "documentation":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Documentation = &v
			case "security":
				var v CapabilityStatementRestSecurity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Security = &v
			case "resource":
				var v CapabilityStatementRestResource
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Resource = append(r.Resource, v)
			case "interaction":
				var v CapabilityStatementRestInteraction
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Interaction = append(r.Interaction, v)
			case "searchParam":
				var v CapabilityStatementRestResourceSearchParam
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SearchParam = append(r.SearchParam, v)
			case "operation":
				var v CapabilityStatementRestResourceOperation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Operation = append(r.Operation, v)
			case "compartment":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Compartment = append(r.Compartment, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CapabilityStatementRest) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Information about security implementation from an interface perspective - what a client needs to know.
type CapabilityStatementRestSecurity struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Server adds CORS headers when responding to requests - this enables Javascript applications to use the server.
	Cors *Boolean
	// Types of security services that are supported/required by the system.
	Service []CodeableConcept
	// General description of how security works.
	Description *Markdown
}
type jsonCapabilityStatementRestSecurity struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Cors                        *Boolean          `json:"cors,omitempty"`
	CorsPrimitiveElement        *primitiveElement `json:"_cors,omitempty"`
	Service                     []CodeableConcept `json:"service,omitempty"`
	Description                 *Markdown         `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
}

func (r CapabilityStatementRestSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CapabilityStatementRestSecurity) marshalJSON() jsonCapabilityStatementRestSecurity {
	m := jsonCapabilityStatementRestSecurity{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Cors != nil && r.Cors.Value != nil {
		m.Cors = r.Cors
	}
	if r.Cors != nil && (r.Cors.Id != nil || r.Cors.Extension != nil) {
		m.CorsPrimitiveElement = &primitiveElement{Id: r.Cors.Id, Extension: r.Cors.Extension}
	}
	m.Service = r.Service
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	return m
}
func (r *CapabilityStatementRestSecurity) UnmarshalJSON(b []byte) error {
	var m jsonCapabilityStatementRestSecurity
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CapabilityStatementRestSecurity) unmarshalJSON(m jsonCapabilityStatementRestSecurity) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Cors = m.Cors
	if m.CorsPrimitiveElement != nil {
		if r.Cors == nil {
			r.Cors = &Boolean{}
		}
		r.Cors.Id = m.CorsPrimitiveElement.Id
		r.Cors.Extension = m.CorsPrimitiveElement.Extension
	}
	r.Service = m.Service
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &Markdown{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	return nil
}
func (r CapabilityStatementRestSecurity) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Cors, xml.StartElement{Name: xml.Name{Local: "cors"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Service, xml.StartElement{Name: xml.Name{Local: "service"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CapabilityStatementRestSecurity) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "cors":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Cors = &v
			case "service":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Service = append(r.Service, v)
			case "description":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CapabilityStatementRestSecurity) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A specification of the restful capabilities of the solution for a specific resource type.
type CapabilityStatementRestResource struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A type of resource exposed via the restful interface.
	Type Code
	// A specification of the profile that describes the solution's overall support for the resource, including any constraints on cardinality, bindings, lengths or other limitations. See further discussion in [Using Profiles](profiling.html#profile-uses).
	Profile *Canonical
	// A list of profiles that represent different use cases supported by the system. For a server, "supported by the system" means the system hosts/produces a set of resources that are conformant to a particular profile, and allows clients that use its services to search using this profile and to find appropriate data. For a client, it means the system will search by this profile and process data according to the guidance implicit in the profile. See further discussion in [Using Profiles](profiling.html#profile-uses).
	SupportedProfile []Canonical
	// Additional information about the resource type used by the system.
	Documentation *Markdown
	// Identifies a restful operation supported by the solution.
	Interaction []CapabilityStatementRestResourceInteraction
	// This field is set to no-version to specify that the system does not support (server) or use (client) versioning for this resource type. If this has some other value, the server must at least correctly track and populate the versionId meta-property on resources. If the value is 'versioned-update', then the server supports all the versioning features, including using e-tags for version integrity in the API.
	Versioning *Code
	// A flag for whether the server is able to return past versions as part of the vRead operation.
	ReadHistory *Boolean
	// A flag to indicate that the server allows or needs to allow the client to create new identities on the server (that is, the client PUTs to a location where there is no existing resource). Allowing this operation means that the server allows the client to create new identities on the server.
	UpdateCreate *Boolean
	// A flag that indicates that the server supports conditional create.
	ConditionalCreate *Boolean
	// A code that indicates how the server supports conditional read.
	ConditionalRead *Code
	// A flag that indicates that the server supports conditional update.
	ConditionalUpdate *Boolean
	// A code that indicates how the server supports conditional delete.
	ConditionalDelete *Code
	// A set of flags that defines how references are supported.
	ReferencePolicy []Code
	// A list of _include values supported by the server.
	SearchInclude []String
	// A list of _revinclude (reverse include) values supported by the server.
	SearchRevInclude []String
	// Search parameters for implementations to support and/or make use of - either references to ones defined in the specification, or additional ones defined for/by the implementation.
	SearchParam []CapabilityStatementRestResourceSearchParam
	// Definition of an operation or a named query together with its parameters and their meaning and type. Consult the definition of the operation for details about how to invoke the operation, and the parameters.
	Operation []CapabilityStatementRestResourceOperation
}
type jsonCapabilityStatementRestResource struct {
	Id                                *string                                      `json:"id,omitempty"`
	Extension                         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension                 []Extension                                  `json:"modifierExtension,omitempty"`
	Type                              Code                                         `json:"type,omitempty"`
	TypePrimitiveElement              *primitiveElement                            `json:"_type,omitempty"`
	Profile                           *Canonical                                   `json:"profile,omitempty"`
	ProfilePrimitiveElement           *primitiveElement                            `json:"_profile,omitempty"`
	SupportedProfile                  []Canonical                                  `json:"supportedProfile,omitempty"`
	SupportedProfilePrimitiveElement  []*primitiveElement                          `json:"_supportedProfile,omitempty"`
	Documentation                     *Markdown                                    `json:"documentation,omitempty"`
	DocumentationPrimitiveElement     *primitiveElement                            `json:"_documentation,omitempty"`
	Interaction                       []CapabilityStatementRestResourceInteraction `json:"interaction,omitempty"`
	Versioning                        *Code                                        `json:"versioning,omitempty"`
	VersioningPrimitiveElement        *primitiveElement                            `json:"_versioning,omitempty"`
	ReadHistory                       *Boolean                                     `json:"readHistory,omitempty"`
	ReadHistoryPrimitiveElement       *primitiveElement                            `json:"_readHistory,omitempty"`
	UpdateCreate                      *Boolean                                     `json:"updateCreate,omitempty"`
	UpdateCreatePrimitiveElement      *primitiveElement                            `json:"_updateCreate,omitempty"`
	ConditionalCreate                 *Boolean                                     `json:"conditionalCreate,omitempty"`
	ConditionalCreatePrimitiveElement *primitiveElement                            `json:"_conditionalCreate,omitempty"`
	ConditionalRead                   *Code                                        `json:"conditionalRead,omitempty"`
	ConditionalReadPrimitiveElement   *primitiveElement                            `json:"_conditionalRead,omitempty"`
	ConditionalUpdate                 *Boolean                                     `json:"conditionalUpdate,omitempty"`
	ConditionalUpdatePrimitiveElement *primitiveElement                            `json:"_conditionalUpdate,omitempty"`
	ConditionalDelete                 *Code                                        `json:"conditionalDelete,omitempty"`
	ConditionalDeletePrimitiveElement *primitiveElement                            `json:"_conditionalDelete,omitempty"`
	ReferencePolicy                   []Code                                       `json:"referencePolicy,omitempty"`
	ReferencePolicyPrimitiveElement   []*primitiveElement                          `json:"_referencePolicy,omitempty"`
	SearchInclude                     []String                                     `json:"searchInclude,omitempty"`
	SearchIncludePrimitiveElement     []*primitiveElement                          `json:"_searchInclude,omitempty"`
	SearchRevInclude                  []String                                     `json:"searchRevInclude,omitempty"`
	SearchRevIncludePrimitiveElement  []*primitiveElement                          `json:"_searchRevInclude,omitempty"`
	SearchParam                       []CapabilityStatementRestResourceSearchParam `json:"searchParam,omitempty"`
	Operation                         []CapabilityStatementRestResourceOperation   `json:"operation,omitempty"`
}

func (r CapabilityStatementRestResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CapabilityStatementRestResource) marshalJSON() jsonCapabilityStatementRestResource {
	m := jsonCapabilityStatementRestResource{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.Profile != nil && r.Profile.Value != nil {
		m.Profile = r.Profile
	}
	if r.Profile != nil && (r.Profile.Id != nil || r.Profile.Extension != nil) {
		m.ProfilePrimitiveElement = &primitiveElement{Id: r.Profile.Id, Extension: r.Profile.Extension}
	}
	anySupportedProfileValue := false
	for _, e := range r.SupportedProfile {
		if e.Value != nil {
			anySupportedProfileValue = true
			break
		}
	}
	if anySupportedProfileValue {
		m.SupportedProfile = r.SupportedProfile
	}
	anySupportedProfileIdOrExtension := false
	for _, e := range r.SupportedProfile {
		if e.Id != nil || e.Extension != nil {
			anySupportedProfileIdOrExtension = true
			break
		}
	}
	if anySupportedProfileIdOrExtension {
		m.SupportedProfilePrimitiveElement = make([]*primitiveElement, 0, len(r.SupportedProfile))
		for _, e := range r.SupportedProfile {
			if e.Id != nil || e.Extension != nil {
				m.SupportedProfilePrimitiveElement = append(m.SupportedProfilePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SupportedProfilePrimitiveElement = append(m.SupportedProfilePrimitiveElement, nil)
			}
		}
	}
	if r.Documentation != nil && r.Documentation.Value != nil {
		m.Documentation = r.Documentation
	}
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	m.Interaction = r.Interaction
	if r.Versioning != nil && r.Versioning.Value != nil {
		m.Versioning = r.Versioning
	}
	if r.Versioning != nil && (r.Versioning.Id != nil || r.Versioning.Extension != nil) {
		m.VersioningPrimitiveElement = &primitiveElement{Id: r.Versioning.Id, Extension: r.Versioning.Extension}
	}
	if r.ReadHistory != nil && r.ReadHistory.Value != nil {
		m.ReadHistory = r.ReadHistory
	}
	if r.ReadHistory != nil && (r.ReadHistory.Id != nil || r.ReadHistory.Extension != nil) {
		m.ReadHistoryPrimitiveElement = &primitiveElement{Id: r.ReadHistory.Id, Extension: r.ReadHistory.Extension}
	}
	if r.UpdateCreate != nil && r.UpdateCreate.Value != nil {
		m.UpdateCreate = r.UpdateCreate
	}
	if r.UpdateCreate != nil && (r.UpdateCreate.Id != nil || r.UpdateCreate.Extension != nil) {
		m.UpdateCreatePrimitiveElement = &primitiveElement{Id: r.UpdateCreate.Id, Extension: r.UpdateCreate.Extension}
	}
	if r.ConditionalCreate != nil && r.ConditionalCreate.Value != nil {
		m.ConditionalCreate = r.ConditionalCreate
	}
	if r.ConditionalCreate != nil && (r.ConditionalCreate.Id != nil || r.ConditionalCreate.Extension != nil) {
		m.ConditionalCreatePrimitiveElement = &primitiveElement{Id: r.ConditionalCreate.Id, Extension: r.ConditionalCreate.Extension}
	}
	if r.ConditionalRead != nil && r.ConditionalRead.Value != nil {
		m.ConditionalRead = r.ConditionalRead
	}
	if r.ConditionalRead != nil && (r.ConditionalRead.Id != nil || r.ConditionalRead.Extension != nil) {
		m.ConditionalReadPrimitiveElement = &primitiveElement{Id: r.ConditionalRead.Id, Extension: r.ConditionalRead.Extension}
	}
	if r.ConditionalUpdate != nil && r.ConditionalUpdate.Value != nil {
		m.ConditionalUpdate = r.ConditionalUpdate
	}
	if r.ConditionalUpdate != nil && (r.ConditionalUpdate.Id != nil || r.ConditionalUpdate.Extension != nil) {
		m.ConditionalUpdatePrimitiveElement = &primitiveElement{Id: r.ConditionalUpdate.Id, Extension: r.ConditionalUpdate.Extension}
	}
	if r.ConditionalDelete != nil && r.ConditionalDelete.Value != nil {
		m.ConditionalDelete = r.ConditionalDelete
	}
	if r.ConditionalDelete != nil && (r.ConditionalDelete.Id != nil || r.ConditionalDelete.Extension != nil) {
		m.ConditionalDeletePrimitiveElement = &primitiveElement{Id: r.ConditionalDelete.Id, Extension: r.ConditionalDelete.Extension}
	}
	anyReferencePolicyValue := false
	for _, e := range r.ReferencePolicy {
		if e.Value != nil {
			anyReferencePolicyValue = true
			break
		}
	}
	if anyReferencePolicyValue {
		m.ReferencePolicy = r.ReferencePolicy
	}
	anyReferencePolicyIdOrExtension := false
	for _, e := range r.ReferencePolicy {
		if e.Id != nil || e.Extension != nil {
			anyReferencePolicyIdOrExtension = true
			break
		}
	}
	if anyReferencePolicyIdOrExtension {
		m.ReferencePolicyPrimitiveElement = make([]*primitiveElement, 0, len(r.ReferencePolicy))
		for _, e := range r.ReferencePolicy {
			if e.Id != nil || e.Extension != nil {
				m.ReferencePolicyPrimitiveElement = append(m.ReferencePolicyPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ReferencePolicyPrimitiveElement = append(m.ReferencePolicyPrimitiveElement, nil)
			}
		}
	}
	anySearchIncludeValue := false
	for _, e := range r.SearchInclude {
		if e.Value != nil {
			anySearchIncludeValue = true
			break
		}
	}
	if anySearchIncludeValue {
		m.SearchInclude = r.SearchInclude
	}
	anySearchIncludeIdOrExtension := false
	for _, e := range r.SearchInclude {
		if e.Id != nil || e.Extension != nil {
			anySearchIncludeIdOrExtension = true
			break
		}
	}
	if anySearchIncludeIdOrExtension {
		m.SearchIncludePrimitiveElement = make([]*primitiveElement, 0, len(r.SearchInclude))
		for _, e := range r.SearchInclude {
			if e.Id != nil || e.Extension != nil {
				m.SearchIncludePrimitiveElement = append(m.SearchIncludePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SearchIncludePrimitiveElement = append(m.SearchIncludePrimitiveElement, nil)
			}
		}
	}
	anySearchRevIncludeValue := false
	for _, e := range r.SearchRevInclude {
		if e.Value != nil {
			anySearchRevIncludeValue = true
			break
		}
	}
	if anySearchRevIncludeValue {
		m.SearchRevInclude = r.SearchRevInclude
	}
	anySearchRevIncludeIdOrExtension := false
	for _, e := range r.SearchRevInclude {
		if e.Id != nil || e.Extension != nil {
			anySearchRevIncludeIdOrExtension = true
			break
		}
	}
	if anySearchRevIncludeIdOrExtension {
		m.SearchRevIncludePrimitiveElement = make([]*primitiveElement, 0, len(r.SearchRevInclude))
		for _, e := range r.SearchRevInclude {
			if e.Id != nil || e.Extension != nil {
				m.SearchRevIncludePrimitiveElement = append(m.SearchRevIncludePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SearchRevIncludePrimitiveElement = append(m.SearchRevIncludePrimitiveElement, nil)
			}
		}
	}
	m.SearchParam = r.SearchParam
	m.Operation = r.Operation
	return m
}
func (r *CapabilityStatementRestResource) UnmarshalJSON(b []byte) error {
	var m jsonCapabilityStatementRestResource
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CapabilityStatementRestResource) unmarshalJSON(m jsonCapabilityStatementRestResource) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Profile = m.Profile
	if m.ProfilePrimitiveElement != nil {
		if r.Profile == nil {
			r.Profile = &Canonical{}
		}
		r.Profile.Id = m.ProfilePrimitiveElement.Id
		r.Profile.Extension = m.ProfilePrimitiveElement.Extension
	}
	r.SupportedProfile = m.SupportedProfile
	for i, e := range m.SupportedProfilePrimitiveElement {
		if len(r.SupportedProfile) <= i {
			r.SupportedProfile = append(r.SupportedProfile, Canonical{})
		}
		if e != nil {
			r.SupportedProfile[i].Id = e.Id
			r.SupportedProfile[i].Extension = e.Extension
		}
	}
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		if r.Documentation == nil {
			r.Documentation = &Markdown{}
		}
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	r.Interaction = m.Interaction
	r.Versioning = m.Versioning
	if m.VersioningPrimitiveElement != nil {
		if r.Versioning == nil {
			r.Versioning = &Code{}
		}
		r.Versioning.Id = m.VersioningPrimitiveElement.Id
		r.Versioning.Extension = m.VersioningPrimitiveElement.Extension
	}
	r.ReadHistory = m.ReadHistory
	if m.ReadHistoryPrimitiveElement != nil {
		if r.ReadHistory == nil {
			r.ReadHistory = &Boolean{}
		}
		r.ReadHistory.Id = m.ReadHistoryPrimitiveElement.Id
		r.ReadHistory.Extension = m.ReadHistoryPrimitiveElement.Extension
	}
	r.UpdateCreate = m.UpdateCreate
	if m.UpdateCreatePrimitiveElement != nil {
		if r.UpdateCreate == nil {
			r.UpdateCreate = &Boolean{}
		}
		r.UpdateCreate.Id = m.UpdateCreatePrimitiveElement.Id
		r.UpdateCreate.Extension = m.UpdateCreatePrimitiveElement.Extension
	}
	r.ConditionalCreate = m.ConditionalCreate
	if m.ConditionalCreatePrimitiveElement != nil {
		if r.ConditionalCreate == nil {
			r.ConditionalCreate = &Boolean{}
		}
		r.ConditionalCreate.Id = m.ConditionalCreatePrimitiveElement.Id
		r.ConditionalCreate.Extension = m.ConditionalCreatePrimitiveElement.Extension
	}
	r.ConditionalRead = m.ConditionalRead
	if m.ConditionalReadPrimitiveElement != nil {
		if r.ConditionalRead == nil {
			r.ConditionalRead = &Code{}
		}
		r.ConditionalRead.Id = m.ConditionalReadPrimitiveElement.Id
		r.ConditionalRead.Extension = m.ConditionalReadPrimitiveElement.Extension
	}
	r.ConditionalUpdate = m.ConditionalUpdate
	if m.ConditionalUpdatePrimitiveElement != nil {
		if r.ConditionalUpdate == nil {
			r.ConditionalUpdate = &Boolean{}
		}
		r.ConditionalUpdate.Id = m.ConditionalUpdatePrimitiveElement.Id
		r.ConditionalUpdate.Extension = m.ConditionalUpdatePrimitiveElement.Extension
	}
	r.ConditionalDelete = m.ConditionalDelete
	if m.ConditionalDeletePrimitiveElement != nil {
		if r.ConditionalDelete == nil {
			r.ConditionalDelete = &Code{}
		}
		r.ConditionalDelete.Id = m.ConditionalDeletePrimitiveElement.Id
		r.ConditionalDelete.Extension = m.ConditionalDeletePrimitiveElement.Extension
	}
	r.ReferencePolicy = m.ReferencePolicy
	for i, e := range m.ReferencePolicyPrimitiveElement {
		if len(r.ReferencePolicy) <= i {
			r.ReferencePolicy = append(r.ReferencePolicy, Code{})
		}
		if e != nil {
			r.ReferencePolicy[i].Id = e.Id
			r.ReferencePolicy[i].Extension = e.Extension
		}
	}
	r.SearchInclude = m.SearchInclude
	for i, e := range m.SearchIncludePrimitiveElement {
		if len(r.SearchInclude) <= i {
			r.SearchInclude = append(r.SearchInclude, String{})
		}
		if e != nil {
			r.SearchInclude[i].Id = e.Id
			r.SearchInclude[i].Extension = e.Extension
		}
	}
	r.SearchRevInclude = m.SearchRevInclude
	for i, e := range m.SearchRevIncludePrimitiveElement {
		if len(r.SearchRevInclude) <= i {
			r.SearchRevInclude = append(r.SearchRevInclude, String{})
		}
		if e != nil {
			r.SearchRevInclude[i].Id = e.Id
			r.SearchRevInclude[i].Extension = e.Extension
		}
	}
	r.SearchParam = m.SearchParam
	r.Operation = m.Operation
	return nil
}
func (r CapabilityStatementRestResource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Profile, xml.StartElement{Name: xml.Name{Local: "profile"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SupportedProfile, xml.StartElement{Name: xml.Name{Local: "supportedProfile"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Documentation, xml.StartElement{Name: xml.Name{Local: "documentation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Interaction, xml.StartElement{Name: xml.Name{Local: "interaction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Versioning, xml.StartElement{Name: xml.Name{Local: "versioning"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReadHistory, xml.StartElement{Name: xml.Name{Local: "readHistory"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UpdateCreate, xml.StartElement{Name: xml.Name{Local: "updateCreate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ConditionalCreate, xml.StartElement{Name: xml.Name{Local: "conditionalCreate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ConditionalRead, xml.StartElement{Name: xml.Name{Local: "conditionalRead"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ConditionalUpdate, xml.StartElement{Name: xml.Name{Local: "conditionalUpdate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ConditionalDelete, xml.StartElement{Name: xml.Name{Local: "conditionalDelete"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReferencePolicy, xml.StartElement{Name: xml.Name{Local: "referencePolicy"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SearchInclude, xml.StartElement{Name: xml.Name{Local: "searchInclude"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SearchRevInclude, xml.StartElement{Name: xml.Name{Local: "searchRevInclude"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SearchParam, xml.StartElement{Name: xml.Name{Local: "searchParam"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Operation, xml.StartElement{Name: xml.Name{Local: "operation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CapabilityStatementRestResource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "type":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "profile":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Profile = &v
			case "supportedProfile":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SupportedProfile = append(r.SupportedProfile, v)
			case "documentation":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Documentation = &v
			case "interaction":
				var v CapabilityStatementRestResourceInteraction
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Interaction = append(r.Interaction, v)
			case "versioning":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Versioning = &v
			case "readHistory":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReadHistory = &v
			case "updateCreate":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UpdateCreate = &v
			case "conditionalCreate":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ConditionalCreate = &v
			case "conditionalRead":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ConditionalRead = &v
			case "conditionalUpdate":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ConditionalUpdate = &v
			case "conditionalDelete":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ConditionalDelete = &v
			case "referencePolicy":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReferencePolicy = append(r.ReferencePolicy, v)
			case "searchInclude":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SearchInclude = append(r.SearchInclude, v)
			case "searchRevInclude":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SearchRevInclude = append(r.SearchRevInclude, v)
			case "searchParam":
				var v CapabilityStatementRestResourceSearchParam
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SearchParam = append(r.SearchParam, v)
			case "operation":
				var v CapabilityStatementRestResourceOperation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Operation = append(r.Operation, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CapabilityStatementRestResource) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Identifies a restful operation supported by the solution.
type CapabilityStatementRestResourceInteraction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Coded identifier of the operation, supported by the system resource.
	Code Code
	// Guidance specific to the implementation of this operation, such as 'delete is a logical delete' or 'updates are only allowed with version id' or 'creates permitted from pre-authorized certificates only'.
	Documentation *Markdown
}
type jsonCapabilityStatementRestResourceInteraction struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	Code                          Code              `json:"code,omitempty"`
	CodePrimitiveElement          *primitiveElement `json:"_code,omitempty"`
	Documentation                 *Markdown         `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement `json:"_documentation,omitempty"`
}

func (r CapabilityStatementRestResourceInteraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CapabilityStatementRestResourceInteraction) marshalJSON() jsonCapabilityStatementRestResourceInteraction {
	m := jsonCapabilityStatementRestResourceInteraction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Code.Value != nil {
		m.Code = r.Code
	}
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	if r.Documentation != nil && r.Documentation.Value != nil {
		m.Documentation = r.Documentation
	}
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	return m
}
func (r *CapabilityStatementRestResourceInteraction) UnmarshalJSON(b []byte) error {
	var m jsonCapabilityStatementRestResourceInteraction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CapabilityStatementRestResourceInteraction) unmarshalJSON(m jsonCapabilityStatementRestResourceInteraction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		if r.Documentation == nil {
			r.Documentation = &Markdown{}
		}
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	return nil
}
func (r CapabilityStatementRestResourceInteraction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Documentation, xml.StartElement{Name: xml.Name{Local: "documentation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CapabilityStatementRestResourceInteraction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "code":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = v
			case "documentation":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Documentation = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CapabilityStatementRestResourceInteraction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Search parameters for implementations to support and/or make use of - either references to ones defined in the specification, or additional ones defined for/by the implementation.
type CapabilityStatementRestResourceSearchParam struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The name of the search parameter used in the interface.
	Name String
	// An absolute URI that is a formal reference to where this parameter was first defined, so that a client can be confident of the meaning of the search parameter (a reference to [SearchParameter.url](searchparameter-definitions.html#SearchParameter.url)). This element SHALL be populated if the search parameter refers to a SearchParameter defined by the FHIR core specification or externally defined IGs.
	Definition *Canonical
	// The type of value a search parameter refers to, and how the content is interpreted.
	Type Code
	// This allows documentation of any distinct behaviors about how the search parameter is used.  For example, text matching algorithms.
	Documentation *Markdown
}
type jsonCapabilityStatementRestResourceSearchParam struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	Name                          String            `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement `json:"_name,omitempty"`
	Definition                    *Canonical        `json:"definition,omitempty"`
	DefinitionPrimitiveElement    *primitiveElement `json:"_definition,omitempty"`
	Type                          Code              `json:"type,omitempty"`
	TypePrimitiveElement          *primitiveElement `json:"_type,omitempty"`
	Documentation                 *Markdown         `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement `json:"_documentation,omitempty"`
}

func (r CapabilityStatementRestResourceSearchParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CapabilityStatementRestResourceSearchParam) marshalJSON() jsonCapabilityStatementRestResourceSearchParam {
	m := jsonCapabilityStatementRestResourceSearchParam{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Definition != nil && r.Definition.Value != nil {
		m.Definition = r.Definition
	}
	if r.Definition != nil && (r.Definition.Id != nil || r.Definition.Extension != nil) {
		m.DefinitionPrimitiveElement = &primitiveElement{Id: r.Definition.Id, Extension: r.Definition.Extension}
	}
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.Documentation != nil && r.Documentation.Value != nil {
		m.Documentation = r.Documentation
	}
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	return m
}
func (r *CapabilityStatementRestResourceSearchParam) UnmarshalJSON(b []byte) error {
	var m jsonCapabilityStatementRestResourceSearchParam
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CapabilityStatementRestResourceSearchParam) unmarshalJSON(m jsonCapabilityStatementRestResourceSearchParam) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Definition = m.Definition
	if m.DefinitionPrimitiveElement != nil {
		if r.Definition == nil {
			r.Definition = &Canonical{}
		}
		r.Definition.Id = m.DefinitionPrimitiveElement.Id
		r.Definition.Extension = m.DefinitionPrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		if r.Documentation == nil {
			r.Documentation = &Markdown{}
		}
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	return nil
}
func (r CapabilityStatementRestResourceSearchParam) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Definition, xml.StartElement{Name: xml.Name{Local: "definition"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Documentation, xml.StartElement{Name: xml.Name{Local: "documentation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CapabilityStatementRestResourceSearchParam) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = v
			case "definition":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Definition = &v
			case "type":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "documentation":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Documentation = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CapabilityStatementRestResourceSearchParam) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Definition of an operation or a named query together with its parameters and their meaning and type. Consult the definition of the operation for details about how to invoke the operation, and the parameters.
type CapabilityStatementRestResourceOperation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The name of the operation or query. For an operation, this is the name  prefixed with $ and used in the URL. For a query, this is the name used in the _query parameter when the query is called.
	Name String
	// Where the formal definition can be found. If a server references the base definition of an Operation (i.e. from the specification itself such as ```http://hl7.org/fhir/OperationDefinition/ValueSet-expand```), that means it supports the full capabilities of the operation - e.g. both GET and POST invocation.  If it only supports a subset, it must define its own custom [OperationDefinition](operationdefinition.html#) with a 'base' of the original OperationDefinition.  The custom definition would describe the specific subset of functionality supported.
	Definition Canonical
	// Documentation that describes anything special about the operation behavior, possibly detailing different behavior for system, type and instance-level invocation of the operation.
	Documentation *Markdown
}
type jsonCapabilityStatementRestResourceOperation struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	Name                          String            `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement `json:"_name,omitempty"`
	Definition                    Canonical         `json:"definition,omitempty"`
	DefinitionPrimitiveElement    *primitiveElement `json:"_definition,omitempty"`
	Documentation                 *Markdown         `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement `json:"_documentation,omitempty"`
}

func (r CapabilityStatementRestResourceOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CapabilityStatementRestResourceOperation) marshalJSON() jsonCapabilityStatementRestResourceOperation {
	m := jsonCapabilityStatementRestResourceOperation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Definition.Value != nil {
		m.Definition = r.Definition
	}
	if r.Definition.Id != nil || r.Definition.Extension != nil {
		m.DefinitionPrimitiveElement = &primitiveElement{Id: r.Definition.Id, Extension: r.Definition.Extension}
	}
	if r.Documentation != nil && r.Documentation.Value != nil {
		m.Documentation = r.Documentation
	}
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	return m
}
func (r *CapabilityStatementRestResourceOperation) UnmarshalJSON(b []byte) error {
	var m jsonCapabilityStatementRestResourceOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CapabilityStatementRestResourceOperation) unmarshalJSON(m jsonCapabilityStatementRestResourceOperation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Definition = m.Definition
	if m.DefinitionPrimitiveElement != nil {
		r.Definition.Id = m.DefinitionPrimitiveElement.Id
		r.Definition.Extension = m.DefinitionPrimitiveElement.Extension
	}
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		if r.Documentation == nil {
			r.Documentation = &Markdown{}
		}
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	return nil
}
func (r CapabilityStatementRestResourceOperation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Definition, xml.StartElement{Name: xml.Name{Local: "definition"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Documentation, xml.StartElement{Name: xml.Name{Local: "documentation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CapabilityStatementRestResourceOperation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = v
			case "definition":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Definition = v
			case "documentation":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Documentation = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CapabilityStatementRestResourceOperation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A specification of restful operations supported by the system.
type CapabilityStatementRestInteraction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A coded identifier of the operation, supported by the system.
	Code Code
	// Guidance specific to the implementation of this operation, such as limitations on the kind of transactions allowed, or information about system wide search is implemented.
	Documentation *Markdown
}
type jsonCapabilityStatementRestInteraction struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	Code                          Code              `json:"code,omitempty"`
	CodePrimitiveElement          *primitiveElement `json:"_code,omitempty"`
	Documentation                 *Markdown         `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement `json:"_documentation,omitempty"`
}

func (r CapabilityStatementRestInteraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CapabilityStatementRestInteraction) marshalJSON() jsonCapabilityStatementRestInteraction {
	m := jsonCapabilityStatementRestInteraction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Code.Value != nil {
		m.Code = r.Code
	}
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	if r.Documentation != nil && r.Documentation.Value != nil {
		m.Documentation = r.Documentation
	}
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	return m
}
func (r *CapabilityStatementRestInteraction) UnmarshalJSON(b []byte) error {
	var m jsonCapabilityStatementRestInteraction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CapabilityStatementRestInteraction) unmarshalJSON(m jsonCapabilityStatementRestInteraction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		if r.Documentation == nil {
			r.Documentation = &Markdown{}
		}
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	return nil
}
func (r CapabilityStatementRestInteraction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Documentation, xml.StartElement{Name: xml.Name{Local: "documentation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CapabilityStatementRestInteraction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "code":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = v
			case "documentation":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Documentation = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CapabilityStatementRestInteraction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A description of the messaging capabilities of the solution.
type CapabilityStatementMessaging struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An endpoint (network accessible address) to which messages and/or replies are to be sent.
	Endpoint []CapabilityStatementMessagingEndpoint
	// Length if the receiver's reliable messaging cache in minutes (if a receiver) or how long the cache length on the receiver should be (if a sender).
	ReliableCache *UnsignedInt
	// Documentation about the system's messaging capabilities for this endpoint not otherwise documented by the capability statement.  For example, the process for becoming an authorized messaging exchange partner.
	Documentation *Markdown
	// References to message definitions for messages this system can send or receive.
	SupportedMessage []CapabilityStatementMessagingSupportedMessage
}
type jsonCapabilityStatementMessaging struct {
	Id                            *string                                        `json:"id,omitempty"`
	Extension                     []Extension                                    `json:"extension,omitempty"`
	ModifierExtension             []Extension                                    `json:"modifierExtension,omitempty"`
	Endpoint                      []CapabilityStatementMessagingEndpoint         `json:"endpoint,omitempty"`
	ReliableCache                 *UnsignedInt                                   `json:"reliableCache,omitempty"`
	ReliableCachePrimitiveElement *primitiveElement                              `json:"_reliableCache,omitempty"`
	Documentation                 *Markdown                                      `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement                              `json:"_documentation,omitempty"`
	SupportedMessage              []CapabilityStatementMessagingSupportedMessage `json:"supportedMessage,omitempty"`
}

func (r CapabilityStatementMessaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CapabilityStatementMessaging) marshalJSON() jsonCapabilityStatementMessaging {
	m := jsonCapabilityStatementMessaging{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Endpoint = r.Endpoint
	if r.ReliableCache != nil && r.ReliableCache.Value != nil {
		m.ReliableCache = r.ReliableCache
	}
	if r.ReliableCache != nil && (r.ReliableCache.Id != nil || r.ReliableCache.Extension != nil) {
		m.ReliableCachePrimitiveElement = &primitiveElement{Id: r.ReliableCache.Id, Extension: r.ReliableCache.Extension}
	}
	if r.Documentation != nil && r.Documentation.Value != nil {
		m.Documentation = r.Documentation
	}
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	m.SupportedMessage = r.SupportedMessage
	return m
}
func (r *CapabilityStatementMessaging) UnmarshalJSON(b []byte) error {
	var m jsonCapabilityStatementMessaging
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CapabilityStatementMessaging) unmarshalJSON(m jsonCapabilityStatementMessaging) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Endpoint = m.Endpoint
	r.ReliableCache = m.ReliableCache
	if m.ReliableCachePrimitiveElement != nil {
		if r.ReliableCache == nil {
			r.ReliableCache = &UnsignedInt{}
		}
		r.ReliableCache.Id = m.ReliableCachePrimitiveElement.Id
		r.ReliableCache.Extension = m.ReliableCachePrimitiveElement.Extension
	}
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		if r.Documentation == nil {
			r.Documentation = &Markdown{}
		}
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	r.SupportedMessage = m.SupportedMessage
	return nil
}
func (r CapabilityStatementMessaging) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Endpoint, xml.StartElement{Name: xml.Name{Local: "endpoint"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReliableCache, xml.StartElement{Name: xml.Name{Local: "reliableCache"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Documentation, xml.StartElement{Name: xml.Name{Local: "documentation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SupportedMessage, xml.StartElement{Name: xml.Name{Local: "supportedMessage"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CapabilityStatementMessaging) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "endpoint":
				var v CapabilityStatementMessagingEndpoint
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Endpoint = append(r.Endpoint, v)
			case "reliableCache":
				var v UnsignedInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReliableCache = &v
			case "documentation":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Documentation = &v
			case "supportedMessage":
				var v CapabilityStatementMessagingSupportedMessage
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SupportedMessage = append(r.SupportedMessage, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CapabilityStatementMessaging) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// An endpoint (network accessible address) to which messages and/or replies are to be sent.
type CapabilityStatementMessagingEndpoint struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A list of the messaging transport protocol(s) identifiers, supported by this endpoint.
	Protocol Coding
	// The network address of the endpoint. For solutions that do not use network addresses for routing, it can be just an identifier.
	Address Url
}
type jsonCapabilityStatementMessagingEndpoint struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Protocol                Coding            `json:"protocol,omitempty"`
	Address                 Url               `json:"address,omitempty"`
	AddressPrimitiveElement *primitiveElement `json:"_address,omitempty"`
}

func (r CapabilityStatementMessagingEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CapabilityStatementMessagingEndpoint) marshalJSON() jsonCapabilityStatementMessagingEndpoint {
	m := jsonCapabilityStatementMessagingEndpoint{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Protocol = r.Protocol
	if r.Address.Value != nil {
		m.Address = r.Address
	}
	if r.Address.Id != nil || r.Address.Extension != nil {
		m.AddressPrimitiveElement = &primitiveElement{Id: r.Address.Id, Extension: r.Address.Extension}
	}
	return m
}
func (r *CapabilityStatementMessagingEndpoint) UnmarshalJSON(b []byte) error {
	var m jsonCapabilityStatementMessagingEndpoint
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CapabilityStatementMessagingEndpoint) unmarshalJSON(m jsonCapabilityStatementMessagingEndpoint) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Protocol = m.Protocol
	r.Address = m.Address
	if m.AddressPrimitiveElement != nil {
		r.Address.Id = m.AddressPrimitiveElement.Id
		r.Address.Extension = m.AddressPrimitiveElement.Extension
	}
	return nil
}
func (r CapabilityStatementMessagingEndpoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Protocol, xml.StartElement{Name: xml.Name{Local: "protocol"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Address, xml.StartElement{Name: xml.Name{Local: "address"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CapabilityStatementMessagingEndpoint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "protocol":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Protocol = v
			case "address":
				var v Url
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Address = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CapabilityStatementMessagingEndpoint) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// References to message definitions for messages this system can send or receive.
type CapabilityStatementMessagingSupportedMessage struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The mode of this event declaration - whether application is sender or receiver.
	Mode Code
	// Points to a message definition that identifies the messaging event, message structure, allowed responses, etc.
	Definition Canonical
}
type jsonCapabilityStatementMessagingSupportedMessage struct {
	Id                         *string           `json:"id,omitempty"`
	Extension                  []Extension       `json:"extension,omitempty"`
	ModifierExtension          []Extension       `json:"modifierExtension,omitempty"`
	Mode                       Code              `json:"mode,omitempty"`
	ModePrimitiveElement       *primitiveElement `json:"_mode,omitempty"`
	Definition                 Canonical         `json:"definition,omitempty"`
	DefinitionPrimitiveElement *primitiveElement `json:"_definition,omitempty"`
}

func (r CapabilityStatementMessagingSupportedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CapabilityStatementMessagingSupportedMessage) marshalJSON() jsonCapabilityStatementMessagingSupportedMessage {
	m := jsonCapabilityStatementMessagingSupportedMessage{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Mode.Value != nil {
		m.Mode = r.Mode
	}
	if r.Mode.Id != nil || r.Mode.Extension != nil {
		m.ModePrimitiveElement = &primitiveElement{Id: r.Mode.Id, Extension: r.Mode.Extension}
	}
	if r.Definition.Value != nil {
		m.Definition = r.Definition
	}
	if r.Definition.Id != nil || r.Definition.Extension != nil {
		m.DefinitionPrimitiveElement = &primitiveElement{Id: r.Definition.Id, Extension: r.Definition.Extension}
	}
	return m
}
func (r *CapabilityStatementMessagingSupportedMessage) UnmarshalJSON(b []byte) error {
	var m jsonCapabilityStatementMessagingSupportedMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CapabilityStatementMessagingSupportedMessage) unmarshalJSON(m jsonCapabilityStatementMessagingSupportedMessage) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Mode = m.Mode
	if m.ModePrimitiveElement != nil {
		r.Mode.Id = m.ModePrimitiveElement.Id
		r.Mode.Extension = m.ModePrimitiveElement.Extension
	}
	r.Definition = m.Definition
	if m.DefinitionPrimitiveElement != nil {
		r.Definition.Id = m.DefinitionPrimitiveElement.Id
		r.Definition.Extension = m.DefinitionPrimitiveElement.Extension
	}
	return nil
}
func (r CapabilityStatementMessagingSupportedMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Mode, xml.StartElement{Name: xml.Name{Local: "mode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Definition, xml.StartElement{Name: xml.Name{Local: "definition"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CapabilityStatementMessagingSupportedMessage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "mode":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Mode = v
			case "definition":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Definition = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CapabilityStatementMessagingSupportedMessage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A document definition.
type CapabilityStatementDocument struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Mode of this document declaration - whether an application is a producer or consumer.
	Mode Code
	// A description of how the application supports or uses the specified document profile.  For example, when documents are created, what action is taken with consumed documents, etc.
	Documentation *Markdown
	// A profile on the document Bundle that constrains which resources are present, and their contents.
	Profile Canonical
}
type jsonCapabilityStatementDocument struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	Mode                          Code              `json:"mode,omitempty"`
	ModePrimitiveElement          *primitiveElement `json:"_mode,omitempty"`
	Documentation                 *Markdown         `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement `json:"_documentation,omitempty"`
	Profile                       Canonical         `json:"profile,omitempty"`
	ProfilePrimitiveElement       *primitiveElement `json:"_profile,omitempty"`
}

func (r CapabilityStatementDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CapabilityStatementDocument) marshalJSON() jsonCapabilityStatementDocument {
	m := jsonCapabilityStatementDocument{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Mode.Value != nil {
		m.Mode = r.Mode
	}
	if r.Mode.Id != nil || r.Mode.Extension != nil {
		m.ModePrimitiveElement = &primitiveElement{Id: r.Mode.Id, Extension: r.Mode.Extension}
	}
	if r.Documentation != nil && r.Documentation.Value != nil {
		m.Documentation = r.Documentation
	}
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	if r.Profile.Value != nil {
		m.Profile = r.Profile
	}
	if r.Profile.Id != nil || r.Profile.Extension != nil {
		m.ProfilePrimitiveElement = &primitiveElement{Id: r.Profile.Id, Extension: r.Profile.Extension}
	}
	return m
}
func (r *CapabilityStatementDocument) UnmarshalJSON(b []byte) error {
	var m jsonCapabilityStatementDocument
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CapabilityStatementDocument) unmarshalJSON(m jsonCapabilityStatementDocument) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Mode = m.Mode
	if m.ModePrimitiveElement != nil {
		r.Mode.Id = m.ModePrimitiveElement.Id
		r.Mode.Extension = m.ModePrimitiveElement.Extension
	}
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		if r.Documentation == nil {
			r.Documentation = &Markdown{}
		}
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	r.Profile = m.Profile
	if m.ProfilePrimitiveElement != nil {
		r.Profile.Id = m.ProfilePrimitiveElement.Id
		r.Profile.Extension = m.ProfilePrimitiveElement.Extension
	}
	return nil
}
func (r CapabilityStatementDocument) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Mode, xml.StartElement{Name: xml.Name{Local: "mode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Documentation, xml.StartElement{Name: xml.Name{Local: "documentation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Profile, xml.StartElement{Name: xml.Name{Local: "profile"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CapabilityStatementDocument) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "mode":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Mode = v
			case "documentation":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Documentation = &v
			case "profile":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Profile = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CapabilityStatementDocument) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
