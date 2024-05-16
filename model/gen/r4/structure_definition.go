package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A definition of a FHIR structure. This resource is used to describe the underlying resources, data types defined in FHIR, and also for describing extensions and constraints on resources and data types.
type StructureDefinition struct {
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
	// An absolute URI that is used to identify this structure definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this structure definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the structure definition is stored on different servers.
	Url Uri
	// A formal identifier that is used to identify this structure definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the structure definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the structure definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the structure definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name String
	// A short, descriptive, user-friendly title for the structure definition.
	Title *String
	// The status of this structure definition. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this structure definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the structure definition was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the structure definition changes.
	Date *DateTime
	// The name of the organization or individual that published the structure definition.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the structure definition from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate structure definition instances.
	UseContext []UsageContext
	// A legal or geographic region in which the structure definition is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this structure definition is needed and why it has been designed as it has.
	Purpose *Markdown
	// A copyright statement relating to the structure definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the structure definition.
	Copyright *Markdown
	// A set of key words or terms from external terminologies that may be used to assist with indexing and searching of templates nby describing the use of this structure definition, or the content it describes.
	Keyword []Coding
	// The version of the FHIR specification on which this StructureDefinition is based - this is the formal version of the specification, without the revision number, e.g. [publication].[major].[minor], which is 4.0.1. for this version.
	FhirVersion *Code
	// An external specification that the content is mapped to.
	Mapping []StructureDefinitionMapping
	// Defines the kind of structure that this definition is describing.
	Kind Code
	// Whether structure this definition describes is abstract or not  - that is, whether the structure is not intended to be instantiated. For Resources and Data types, abstract types will never be exchanged  between systems.
	Abstract Boolean
	// Identifies the types of resource or data type elements to which the extension can be applied.
	Context []StructureDefinitionContext
	// A set of rules as FHIRPath Invariants about when the extension can be used (e.g. co-occurrence variants for the extension). All the rules must be true.
	ContextInvariant []String
	// The type this structure describes. If the derivation kind is 'specialization' then this is the master definition for a type, and there is always one of these (a data type, an extension, a resource, including abstract ones). Otherwise the structure definition is a constraint on the stated type (and in this case, the type cannot be an abstract type).  References are URLs that are relative to http://hl7.org/fhir/StructureDefinition e.g. "string" is a reference to http://hl7.org/fhir/StructureDefinition/string. Absolute URLs are only allowed in logical models.
	Type Uri
	// An absolute URI that is the base structure from which this type is derived, either by specialization or constraint.
	BaseDefinition *Canonical
	// How the type relates to the baseDefinition.
	Derivation *Code
	// A snapshot view is expressed in a standalone form that can be used and interpreted without considering the base StructureDefinition.
	Snapshot *StructureDefinitionSnapshot
	// A differential view is expressed relative to the base StructureDefinition - a statement of differences that it applies.
	Differential *StructureDefinitionDifferential
}

func (r StructureDefinition) ResourceType() string {
	return "StructureDefinition"
}
func (r StructureDefinition) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonStructureDefinition struct {
	ResourceType                     string                           `json:"resourceType"`
	Id                               *Id                              `json:"id,omitempty"`
	IdPrimitiveElement               *primitiveElement                `json:"_id,omitempty"`
	Meta                             *Meta                            `json:"meta,omitempty"`
	ImplicitRules                    *Uri                             `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement    *primitiveElement                `json:"_implicitRules,omitempty"`
	Language                         *Code                            `json:"language,omitempty"`
	LanguagePrimitiveElement         *primitiveElement                `json:"_language,omitempty"`
	Text                             *Narrative                       `json:"text,omitempty"`
	Contained                        []containedResource              `json:"contained,omitempty"`
	Extension                        []Extension                      `json:"extension,omitempty"`
	ModifierExtension                []Extension                      `json:"modifierExtension,omitempty"`
	Url                              Uri                              `json:"url,omitempty"`
	UrlPrimitiveElement              *primitiveElement                `json:"_url,omitempty"`
	Identifier                       []Identifier                     `json:"identifier,omitempty"`
	Version                          *String                          `json:"version,omitempty"`
	VersionPrimitiveElement          *primitiveElement                `json:"_version,omitempty"`
	Name                             String                           `json:"name,omitempty"`
	NamePrimitiveElement             *primitiveElement                `json:"_name,omitempty"`
	Title                            *String                          `json:"title,omitempty"`
	TitlePrimitiveElement            *primitiveElement                `json:"_title,omitempty"`
	Status                           Code                             `json:"status,omitempty"`
	StatusPrimitiveElement           *primitiveElement                `json:"_status,omitempty"`
	Experimental                     *Boolean                         `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement     *primitiveElement                `json:"_experimental,omitempty"`
	Date                             *DateTime                        `json:"date,omitempty"`
	DatePrimitiveElement             *primitiveElement                `json:"_date,omitempty"`
	Publisher                        *String                          `json:"publisher,omitempty"`
	PublisherPrimitiveElement        *primitiveElement                `json:"_publisher,omitempty"`
	Contact                          []ContactDetail                  `json:"contact,omitempty"`
	Description                      *Markdown                        `json:"description,omitempty"`
	DescriptionPrimitiveElement      *primitiveElement                `json:"_description,omitempty"`
	UseContext                       []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction                     []CodeableConcept                `json:"jurisdiction,omitempty"`
	Purpose                          *Markdown                        `json:"purpose,omitempty"`
	PurposePrimitiveElement          *primitiveElement                `json:"_purpose,omitempty"`
	Copyright                        *Markdown                        `json:"copyright,omitempty"`
	CopyrightPrimitiveElement        *primitiveElement                `json:"_copyright,omitempty"`
	Keyword                          []Coding                         `json:"keyword,omitempty"`
	FhirVersion                      *Code                            `json:"fhirVersion,omitempty"`
	FhirVersionPrimitiveElement      *primitiveElement                `json:"_fhirVersion,omitempty"`
	Mapping                          []StructureDefinitionMapping     `json:"mapping,omitempty"`
	Kind                             Code                             `json:"kind,omitempty"`
	KindPrimitiveElement             *primitiveElement                `json:"_kind,omitempty"`
	Abstract                         Boolean                          `json:"abstract,omitempty"`
	AbstractPrimitiveElement         *primitiveElement                `json:"_abstract,omitempty"`
	Context                          []StructureDefinitionContext     `json:"context,omitempty"`
	ContextInvariant                 []String                         `json:"contextInvariant,omitempty"`
	ContextInvariantPrimitiveElement []*primitiveElement              `json:"_contextInvariant,omitempty"`
	Type                             Uri                              `json:"type,omitempty"`
	TypePrimitiveElement             *primitiveElement                `json:"_type,omitempty"`
	BaseDefinition                   *Canonical                       `json:"baseDefinition,omitempty"`
	BaseDefinitionPrimitiveElement   *primitiveElement                `json:"_baseDefinition,omitempty"`
	Derivation                       *Code                            `json:"derivation,omitempty"`
	DerivationPrimitiveElement       *primitiveElement                `json:"_derivation,omitempty"`
	Snapshot                         *StructureDefinitionSnapshot     `json:"snapshot,omitempty"`
	Differential                     *StructureDefinitionDifferential `json:"differential,omitempty"`
}

func (r StructureDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r StructureDefinition) marshalJSON() jsonStructureDefinition {
	m := jsonStructureDefinition{}
	m.ResourceType = "StructureDefinition"
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
	m.Keyword = r.Keyword
	m.FhirVersion = r.FhirVersion
	if r.FhirVersion != nil && (r.FhirVersion.Id != nil || r.FhirVersion.Extension != nil) {
		m.FhirVersionPrimitiveElement = &primitiveElement{Id: r.FhirVersion.Id, Extension: r.FhirVersion.Extension}
	}
	m.Mapping = r.Mapping
	m.Kind = r.Kind
	if r.Kind.Id != nil || r.Kind.Extension != nil {
		m.KindPrimitiveElement = &primitiveElement{Id: r.Kind.Id, Extension: r.Kind.Extension}
	}
	m.Abstract = r.Abstract
	if r.Abstract.Id != nil || r.Abstract.Extension != nil {
		m.AbstractPrimitiveElement = &primitiveElement{Id: r.Abstract.Id, Extension: r.Abstract.Extension}
	}
	m.Context = r.Context
	m.ContextInvariant = r.ContextInvariant
	anyContextInvariantIdOrExtension := false
	for _, e := range r.ContextInvariant {
		if e.Id != nil || e.Extension != nil {
			anyContextInvariantIdOrExtension = true
			break
		}
	}
	if anyContextInvariantIdOrExtension {
		m.ContextInvariantPrimitiveElement = make([]*primitiveElement, 0, len(r.ContextInvariant))
		for _, e := range r.ContextInvariant {
			if e.Id != nil || e.Extension != nil {
				m.ContextInvariantPrimitiveElement = append(m.ContextInvariantPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ContextInvariantPrimitiveElement = append(m.ContextInvariantPrimitiveElement, nil)
			}
		}
	}
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.BaseDefinition = r.BaseDefinition
	if r.BaseDefinition != nil && (r.BaseDefinition.Id != nil || r.BaseDefinition.Extension != nil) {
		m.BaseDefinitionPrimitiveElement = &primitiveElement{Id: r.BaseDefinition.Id, Extension: r.BaseDefinition.Extension}
	}
	m.Derivation = r.Derivation
	if r.Derivation != nil && (r.Derivation.Id != nil || r.Derivation.Extension != nil) {
		m.DerivationPrimitiveElement = &primitiveElement{Id: r.Derivation.Id, Extension: r.Derivation.Extension}
	}
	m.Snapshot = r.Snapshot
	m.Differential = r.Differential
	return m
}
func (r *StructureDefinition) UnmarshalJSON(b []byte) error {
	var m jsonStructureDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *StructureDefinition) unmarshalJSON(m jsonStructureDefinition) error {
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
	r.Keyword = m.Keyword
	r.FhirVersion = m.FhirVersion
	if m.FhirVersionPrimitiveElement != nil {
		r.FhirVersion.Id = m.FhirVersionPrimitiveElement.Id
		r.FhirVersion.Extension = m.FhirVersionPrimitiveElement.Extension
	}
	r.Mapping = m.Mapping
	r.Kind = m.Kind
	if m.KindPrimitiveElement != nil {
		r.Kind.Id = m.KindPrimitiveElement.Id
		r.Kind.Extension = m.KindPrimitiveElement.Extension
	}
	r.Abstract = m.Abstract
	if m.AbstractPrimitiveElement != nil {
		r.Abstract.Id = m.AbstractPrimitiveElement.Id
		r.Abstract.Extension = m.AbstractPrimitiveElement.Extension
	}
	r.Context = m.Context
	r.ContextInvariant = m.ContextInvariant
	for i, e := range m.ContextInvariantPrimitiveElement {
		if len(r.ContextInvariant) > i {
			r.ContextInvariant[i].Id = e.Id
			r.ContextInvariant[i].Extension = e.Extension
		} else {
			r.ContextInvariant = append(r.ContextInvariant, String{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.BaseDefinition = m.BaseDefinition
	if m.BaseDefinitionPrimitiveElement != nil {
		r.BaseDefinition.Id = m.BaseDefinitionPrimitiveElement.Id
		r.BaseDefinition.Extension = m.BaseDefinitionPrimitiveElement.Extension
	}
	r.Derivation = m.Derivation
	if m.DerivationPrimitiveElement != nil {
		r.Derivation.Id = m.DerivationPrimitiveElement.Id
		r.Derivation.Extension = m.DerivationPrimitiveElement.Extension
	}
	r.Snapshot = m.Snapshot
	r.Differential = m.Differential
	return nil
}
func (r StructureDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// An external specification that the content is mapped to.
type StructureDefinitionMapping struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An Internal id that is used to identify this mapping set when specific mappings are made.
	Identity Id
	// An absolute URI that identifies the specification that this mapping is expressed to.
	Uri *Uri
	// A name for the specification that is being mapped to.
	Name *String
	// Comments about this mapping, including version notes, issues, scope limitations, and other important notes for usage.
	Comment *String
}
type jsonStructureDefinitionMapping struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Identity                 Id                `json:"identity,omitempty"`
	IdentityPrimitiveElement *primitiveElement `json:"_identity,omitempty"`
	Uri                      *Uri              `json:"uri,omitempty"`
	UriPrimitiveElement      *primitiveElement `json:"_uri,omitempty"`
	Name                     *String           `json:"name,omitempty"`
	NamePrimitiveElement     *primitiveElement `json:"_name,omitempty"`
	Comment                  *String           `json:"comment,omitempty"`
	CommentPrimitiveElement  *primitiveElement `json:"_comment,omitempty"`
}

func (r StructureDefinitionMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r StructureDefinitionMapping) marshalJSON() jsonStructureDefinitionMapping {
	m := jsonStructureDefinitionMapping{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identity = r.Identity
	if r.Identity.Id != nil || r.Identity.Extension != nil {
		m.IdentityPrimitiveElement = &primitiveElement{Id: r.Identity.Id, Extension: r.Identity.Extension}
	}
	m.Uri = r.Uri
	if r.Uri != nil && (r.Uri.Id != nil || r.Uri.Extension != nil) {
		m.UriPrimitiveElement = &primitiveElement{Id: r.Uri.Id, Extension: r.Uri.Extension}
	}
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Comment = r.Comment
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	return m
}
func (r *StructureDefinitionMapping) UnmarshalJSON(b []byte) error {
	var m jsonStructureDefinitionMapping
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *StructureDefinitionMapping) unmarshalJSON(m jsonStructureDefinitionMapping) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identity = m.Identity
	if m.IdentityPrimitiveElement != nil {
		r.Identity.Id = m.IdentityPrimitiveElement.Id
		r.Identity.Extension = m.IdentityPrimitiveElement.Extension
	}
	r.Uri = m.Uri
	if m.UriPrimitiveElement != nil {
		r.Uri.Id = m.UriPrimitiveElement.Id
		r.Uri.Extension = m.UriPrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	return nil
}
func (r StructureDefinitionMapping) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Identifies the types of resource or data type elements to which the extension can be applied.
type StructureDefinitionContext struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Defines how to interpret the expression that defines what the context of the extension is.
	Type Code
	// An expression that defines where an extension can be used in resources.
	Expression String
}
type jsonStructureDefinitionContext struct {
	Id                         *string           `json:"id,omitempty"`
	Extension                  []Extension       `json:"extension,omitempty"`
	ModifierExtension          []Extension       `json:"modifierExtension,omitempty"`
	Type                       Code              `json:"type,omitempty"`
	TypePrimitiveElement       *primitiveElement `json:"_type,omitempty"`
	Expression                 String            `json:"expression,omitempty"`
	ExpressionPrimitiveElement *primitiveElement `json:"_expression,omitempty"`
}

func (r StructureDefinitionContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r StructureDefinitionContext) marshalJSON() jsonStructureDefinitionContext {
	m := jsonStructureDefinitionContext{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Expression = r.Expression
	if r.Expression.Id != nil || r.Expression.Extension != nil {
		m.ExpressionPrimitiveElement = &primitiveElement{Id: r.Expression.Id, Extension: r.Expression.Extension}
	}
	return m
}
func (r *StructureDefinitionContext) UnmarshalJSON(b []byte) error {
	var m jsonStructureDefinitionContext
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *StructureDefinitionContext) unmarshalJSON(m jsonStructureDefinitionContext) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Expression = m.Expression
	if m.ExpressionPrimitiveElement != nil {
		r.Expression.Id = m.ExpressionPrimitiveElement.Id
		r.Expression.Extension = m.ExpressionPrimitiveElement.Extension
	}
	return nil
}
func (r StructureDefinitionContext) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A snapshot view is expressed in a standalone form that can be used and interpreted without considering the base StructureDefinition.
type StructureDefinitionSnapshot struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Captures constraints on each element within the resource.
	Element []ElementDefinition
}
type jsonStructureDefinitionSnapshot struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Element           []ElementDefinition `json:"element,omitempty"`
}

func (r StructureDefinitionSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r StructureDefinitionSnapshot) marshalJSON() jsonStructureDefinitionSnapshot {
	m := jsonStructureDefinitionSnapshot{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Element = r.Element
	return m
}
func (r *StructureDefinitionSnapshot) UnmarshalJSON(b []byte) error {
	var m jsonStructureDefinitionSnapshot
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *StructureDefinitionSnapshot) unmarshalJSON(m jsonStructureDefinitionSnapshot) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Element = m.Element
	return nil
}
func (r StructureDefinitionSnapshot) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A differential view is expressed relative to the base StructureDefinition - a statement of differences that it applies.
type StructureDefinitionDifferential struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Captures constraints on each element within the resource.
	Element []ElementDefinition
}
type jsonStructureDefinitionDifferential struct {
	Id                *string             `json:"id,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Element           []ElementDefinition `json:"element,omitempty"`
}

func (r StructureDefinitionDifferential) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r StructureDefinitionDifferential) marshalJSON() jsonStructureDefinitionDifferential {
	m := jsonStructureDefinitionDifferential{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Element = r.Element
	return m
}
func (r *StructureDefinitionDifferential) UnmarshalJSON(b []byte) error {
	var m jsonStructureDefinitionDifferential
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *StructureDefinitionDifferential) unmarshalJSON(m jsonStructureDefinitionDifferential) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Element = m.Element
	return nil
}
func (r StructureDefinitionDifferential) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
