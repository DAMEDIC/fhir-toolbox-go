package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// A formal computable definition of an operation (on the RESTful interface) or a named query (using the search interaction).
type OperationDefinition struct {
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
	// An absolute URI that is used to identify this operation definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this operation definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the operation definition is stored on different servers.
	Url *Uri
	// The identifier that is used to identify this version of the operation definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the operation definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the operation definition. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name String
	// A short, descriptive, user-friendly title for the operation definition.
	Title *String
	// The status of this operation definition. Enables tracking the life-cycle of the content.
	Status Code
	// Whether this is an operation or a named query.
	Kind Code
	// A Boolean value to indicate that this operation definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the operation definition was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the operation definition changes.
	Date *DateTime
	// The name of the organization or individual that published the operation definition.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the operation definition from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate operation definition instances.
	UseContext []UsageContext
	// A legal or geographic region in which the operation definition is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this operation definition is needed and why it has been designed as it has.
	Purpose *Markdown
	// Whether the operation affects state. Side effects such as producing audit trail entries do not count as 'affecting  state'.
	AffectsState *Boolean
	// The name used to invoke the operation.
	Code Code
	// Additional information about how to use this operation or named query.
	Comment *Markdown
	// Indicates that this operation definition is a constraining profile on the base.
	Base *Canonical
	// The types on which this operation can be executed.
	Resource []Code
	// Indicates whether this operation or named query can be invoked at the system level (e.g. without needing to choose a resource type for the context).
	System Boolean
	// Indicates whether this operation or named query can be invoked at the resource type level for any given resource type level (e.g. without needing to choose a specific resource id for the context).
	Type Boolean
	// Indicates whether this operation can be invoked on a particular instance of one of the given types.
	Instance Boolean
	// Additional validation information for the in parameters - a single profile that covers all the parameters. The profile is a constraint on the parameters resource as a whole.
	InputProfile *Canonical
	// Additional validation information for the out parameters - a single profile that covers all the parameters. The profile is a constraint on the parameters resource.
	OutputProfile *Canonical
	// The parameters for the operation/query.
	Parameter []OperationDefinitionParameter
	// Defines an appropriate combination of parameters to use when invoking this operation, to help code generators when generating overloaded parameter sets for this operation.
	Overload []OperationDefinitionOverload
}

func (r OperationDefinition) ResourceType() string {
	return "OperationDefinition"
}
func (r OperationDefinition) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonOperationDefinition struct {
	ResourceType                  string                         `json:"resourceType"`
	Id                            *Id                            `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement              `json:"_id,omitempty"`
	Meta                          *Meta                          `json:"meta,omitempty"`
	ImplicitRules                 *Uri                           `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement              `json:"_implicitRules,omitempty"`
	Language                      *Code                          `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement              `json:"_language,omitempty"`
	Text                          *Narrative                     `json:"text,omitempty"`
	Contained                     []ContainedResource            `json:"contained,omitempty"`
	Extension                     []Extension                    `json:"extension,omitempty"`
	ModifierExtension             []Extension                    `json:"modifierExtension,omitempty"`
	Url                           *Uri                           `json:"url,omitempty"`
	UrlPrimitiveElement           *primitiveElement              `json:"_url,omitempty"`
	Version                       *String                        `json:"version,omitempty"`
	VersionPrimitiveElement       *primitiveElement              `json:"_version,omitempty"`
	Name                          String                         `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement              `json:"_name,omitempty"`
	Title                         *String                        `json:"title,omitempty"`
	TitlePrimitiveElement         *primitiveElement              `json:"_title,omitempty"`
	Status                        Code                           `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement              `json:"_status,omitempty"`
	Kind                          Code                           `json:"kind,omitempty"`
	KindPrimitiveElement          *primitiveElement              `json:"_kind,omitempty"`
	Experimental                  *Boolean                       `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement  *primitiveElement              `json:"_experimental,omitempty"`
	Date                          *DateTime                      `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement              `json:"_date,omitempty"`
	Publisher                     *String                        `json:"publisher,omitempty"`
	PublisherPrimitiveElement     *primitiveElement              `json:"_publisher,omitempty"`
	Contact                       []ContactDetail                `json:"contact,omitempty"`
	Description                   *Markdown                      `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement              `json:"_description,omitempty"`
	UseContext                    []UsageContext                 `json:"useContext,omitempty"`
	Jurisdiction                  []CodeableConcept              `json:"jurisdiction,omitempty"`
	Purpose                       *Markdown                      `json:"purpose,omitempty"`
	PurposePrimitiveElement       *primitiveElement              `json:"_purpose,omitempty"`
	AffectsState                  *Boolean                       `json:"affectsState,omitempty"`
	AffectsStatePrimitiveElement  *primitiveElement              `json:"_affectsState,omitempty"`
	Code                          Code                           `json:"code,omitempty"`
	CodePrimitiveElement          *primitiveElement              `json:"_code,omitempty"`
	Comment                       *Markdown                      `json:"comment,omitempty"`
	CommentPrimitiveElement       *primitiveElement              `json:"_comment,omitempty"`
	Base                          *Canonical                     `json:"base,omitempty"`
	BasePrimitiveElement          *primitiveElement              `json:"_base,omitempty"`
	Resource                      []Code                         `json:"resource,omitempty"`
	ResourcePrimitiveElement      []*primitiveElement            `json:"_resource,omitempty"`
	System                        Boolean                        `json:"system,omitempty"`
	SystemPrimitiveElement        *primitiveElement              `json:"_system,omitempty"`
	Type                          Boolean                        `json:"type,omitempty"`
	TypePrimitiveElement          *primitiveElement              `json:"_type,omitempty"`
	Instance                      Boolean                        `json:"instance,omitempty"`
	InstancePrimitiveElement      *primitiveElement              `json:"_instance,omitempty"`
	InputProfile                  *Canonical                     `json:"inputProfile,omitempty"`
	InputProfilePrimitiveElement  *primitiveElement              `json:"_inputProfile,omitempty"`
	OutputProfile                 *Canonical                     `json:"outputProfile,omitempty"`
	OutputProfilePrimitiveElement *primitiveElement              `json:"_outputProfile,omitempty"`
	Parameter                     []OperationDefinitionParameter `json:"parameter,omitempty"`
	Overload                      []OperationDefinitionOverload  `json:"overload,omitempty"`
}

func (r OperationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r OperationDefinition) marshalJSON() jsonOperationDefinition {
	m := jsonOperationDefinition{}
	m.ResourceType = "OperationDefinition"
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
	if r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name.Id != nil || r.Name.Extension != nil {
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
	if r.Kind.Value != nil {
		m.Kind = r.Kind
	}
	if r.Kind.Id != nil || r.Kind.Extension != nil {
		m.KindPrimitiveElement = &primitiveElement{Id: r.Kind.Id, Extension: r.Kind.Extension}
	}
	if r.Experimental != nil && r.Experimental.Value != nil {
		m.Experimental = r.Experimental
	}
	if r.Experimental != nil && (r.Experimental.Id != nil || r.Experimental.Extension != nil) {
		m.ExperimentalPrimitiveElement = &primitiveElement{Id: r.Experimental.Id, Extension: r.Experimental.Extension}
	}
	if r.Date != nil && r.Date.Value != nil {
		m.Date = r.Date
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
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
	if r.AffectsState != nil && r.AffectsState.Value != nil {
		m.AffectsState = r.AffectsState
	}
	if r.AffectsState != nil && (r.AffectsState.Id != nil || r.AffectsState.Extension != nil) {
		m.AffectsStatePrimitiveElement = &primitiveElement{Id: r.AffectsState.Id, Extension: r.AffectsState.Extension}
	}
	if r.Code.Value != nil {
		m.Code = r.Code
	}
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	if r.Comment != nil && r.Comment.Value != nil {
		m.Comment = r.Comment
	}
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	if r.Base != nil && r.Base.Value != nil {
		m.Base = r.Base
	}
	if r.Base != nil && (r.Base.Id != nil || r.Base.Extension != nil) {
		m.BasePrimitiveElement = &primitiveElement{Id: r.Base.Id, Extension: r.Base.Extension}
	}
	anyResourceValue := false
	for _, e := range r.Resource {
		if e.Value != nil {
			anyResourceValue = true
			break
		}
	}
	if anyResourceValue {
		m.Resource = r.Resource
	}
	anyResourceIdOrExtension := false
	for _, e := range r.Resource {
		if e.Id != nil || e.Extension != nil {
			anyResourceIdOrExtension = true
			break
		}
	}
	if anyResourceIdOrExtension {
		m.ResourcePrimitiveElement = make([]*primitiveElement, 0, len(r.Resource))
		for _, e := range r.Resource {
			if e.Id != nil || e.Extension != nil {
				m.ResourcePrimitiveElement = append(m.ResourcePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ResourcePrimitiveElement = append(m.ResourcePrimitiveElement, nil)
			}
		}
	}
	if r.System.Value != nil {
		m.System = r.System
	}
	if r.System.Id != nil || r.System.Extension != nil {
		m.SystemPrimitiveElement = &primitiveElement{Id: r.System.Id, Extension: r.System.Extension}
	}
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.Instance.Value != nil {
		m.Instance = r.Instance
	}
	if r.Instance.Id != nil || r.Instance.Extension != nil {
		m.InstancePrimitiveElement = &primitiveElement{Id: r.Instance.Id, Extension: r.Instance.Extension}
	}
	if r.InputProfile != nil && r.InputProfile.Value != nil {
		m.InputProfile = r.InputProfile
	}
	if r.InputProfile != nil && (r.InputProfile.Id != nil || r.InputProfile.Extension != nil) {
		m.InputProfilePrimitiveElement = &primitiveElement{Id: r.InputProfile.Id, Extension: r.InputProfile.Extension}
	}
	if r.OutputProfile != nil && r.OutputProfile.Value != nil {
		m.OutputProfile = r.OutputProfile
	}
	if r.OutputProfile != nil && (r.OutputProfile.Id != nil || r.OutputProfile.Extension != nil) {
		m.OutputProfilePrimitiveElement = &primitiveElement{Id: r.OutputProfile.Id, Extension: r.OutputProfile.Extension}
	}
	m.Parameter = r.Parameter
	m.Overload = r.Overload
	return m
}
func (r *OperationDefinition) UnmarshalJSON(b []byte) error {
	var m jsonOperationDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *OperationDefinition) unmarshalJSON(m jsonOperationDefinition) error {
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
	r.Kind = m.Kind
	if m.KindPrimitiveElement != nil {
		r.Kind.Id = m.KindPrimitiveElement.Id
		r.Kind.Extension = m.KindPrimitiveElement.Extension
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
		if r.Date == nil {
			r.Date = &DateTime{}
		}
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
	r.AffectsState = m.AffectsState
	if m.AffectsStatePrimitiveElement != nil {
		if r.AffectsState == nil {
			r.AffectsState = &Boolean{}
		}
		r.AffectsState.Id = m.AffectsStatePrimitiveElement.Id
		r.AffectsState.Extension = m.AffectsStatePrimitiveElement.Extension
	}
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		if r.Comment == nil {
			r.Comment = &Markdown{}
		}
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	r.Base = m.Base
	if m.BasePrimitiveElement != nil {
		if r.Base == nil {
			r.Base = &Canonical{}
		}
		r.Base.Id = m.BasePrimitiveElement.Id
		r.Base.Extension = m.BasePrimitiveElement.Extension
	}
	r.Resource = m.Resource
	for i, e := range m.ResourcePrimitiveElement {
		if len(r.Resource) <= i {
			r.Resource = append(r.Resource, Code{})
		}
		if e != nil {
			r.Resource[i].Id = e.Id
			r.Resource[i].Extension = e.Extension
		}
	}
	r.System = m.System
	if m.SystemPrimitiveElement != nil {
		r.System.Id = m.SystemPrimitiveElement.Id
		r.System.Extension = m.SystemPrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Instance = m.Instance
	if m.InstancePrimitiveElement != nil {
		r.Instance.Id = m.InstancePrimitiveElement.Id
		r.Instance.Extension = m.InstancePrimitiveElement.Extension
	}
	r.InputProfile = m.InputProfile
	if m.InputProfilePrimitiveElement != nil {
		if r.InputProfile == nil {
			r.InputProfile = &Canonical{}
		}
		r.InputProfile.Id = m.InputProfilePrimitiveElement.Id
		r.InputProfile.Extension = m.InputProfilePrimitiveElement.Extension
	}
	r.OutputProfile = m.OutputProfile
	if m.OutputProfilePrimitiveElement != nil {
		if r.OutputProfile == nil {
			r.OutputProfile = &Canonical{}
		}
		r.OutputProfile.Id = m.OutputProfilePrimitiveElement.Id
		r.OutputProfile.Extension = m.OutputProfilePrimitiveElement.Extension
	}
	r.Parameter = m.Parameter
	r.Overload = m.Overload
	return nil
}
func (r OperationDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "OperationDefinition"
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
	err = e.EncodeElement(r.Kind, xml.StartElement{Name: xml.Name{Local: "kind"}})
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
	err = e.EncodeElement(r.AffectsState, xml.StartElement{Name: xml.Name{Local: "affectsState"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Comment, xml.StartElement{Name: xml.Name{Local: "comment"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Base, xml.StartElement{Name: xml.Name{Local: "base"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Resource, xml.StartElement{Name: xml.Name{Local: "resource"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.System, xml.StartElement{Name: xml.Name{Local: "system"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Instance, xml.StartElement{Name: xml.Name{Local: "instance"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.InputProfile, xml.StartElement{Name: xml.Name{Local: "inputProfile"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OutputProfile, xml.StartElement{Name: xml.Name{Local: "outputProfile"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Parameter, xml.StartElement{Name: xml.Name{Local: "parameter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Overload, xml.StartElement{Name: xml.Name{Local: "overload"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *OperationDefinition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.Name = v
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
			case "kind":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Kind = v
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
				r.Date = &v
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
			case "affectsState":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AffectsState = &v
			case "code":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = v
			case "comment":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Comment = &v
			case "base":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Base = &v
			case "resource":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Resource = append(r.Resource, v)
			case "system":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.System = v
			case "type":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "instance":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Instance = v
			case "inputProfile":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.InputProfile = &v
			case "outputProfile":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OutputProfile = &v
			case "parameter":
				var v OperationDefinitionParameter
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Parameter = append(r.Parameter, v)
			case "overload":
				var v OperationDefinitionOverload
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Overload = append(r.Overload, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r OperationDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The parameters for the operation/query.
type OperationDefinitionParameter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The name of used to identify the parameter.
	Name Code
	// Whether this is an input or an output parameter.
	Use Code
	// The minimum number of times this parameter SHALL appear in the request or response.
	Min Integer
	// The maximum number of times this element is permitted to appear in the request or response.
	Max String
	// Describes the meaning or use of this parameter.
	Documentation *String
	// The type for this parameter.
	Type *Code
	// Used when the type is "Reference" or "canonical", and identifies a profile structure or implementation Guide that applies to the target of the reference this parameter refers to. If any profiles are specified, then the content must conform to at least one of them. The URL can be a local reference - to a contained StructureDefinition, or a reference to another StructureDefinition or Implementation Guide by a canonical URL. When an implementation guide is specified, the target resource SHALL conform to at least one profile defined in the implementation guide.
	TargetProfile []Canonical
	// How the parameter is understood as a search parameter. This is only used if the parameter type is 'string'.
	SearchType *Code
	// Binds to a value set if this parameter is coded (code, Coding, CodeableConcept).
	Binding *OperationDefinitionParameterBinding
	// Identifies other resource parameters within the operation invocation that are expected to resolve to this resource.
	ReferencedFrom []OperationDefinitionParameterReferencedFrom
	// The parts of a nested Parameter.
	Part []OperationDefinitionParameter
}
type jsonOperationDefinitionParameter struct {
	Id                            *string                                      `json:"id,omitempty"`
	Extension                     []Extension                                  `json:"extension,omitempty"`
	ModifierExtension             []Extension                                  `json:"modifierExtension,omitempty"`
	Name                          Code                                         `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement                            `json:"_name,omitempty"`
	Use                           Code                                         `json:"use,omitempty"`
	UsePrimitiveElement           *primitiveElement                            `json:"_use,omitempty"`
	Min                           Integer                                      `json:"min,omitempty"`
	MinPrimitiveElement           *primitiveElement                            `json:"_min,omitempty"`
	Max                           String                                       `json:"max,omitempty"`
	MaxPrimitiveElement           *primitiveElement                            `json:"_max,omitempty"`
	Documentation                 *String                                      `json:"documentation,omitempty"`
	DocumentationPrimitiveElement *primitiveElement                            `json:"_documentation,omitempty"`
	Type                          *Code                                        `json:"type,omitempty"`
	TypePrimitiveElement          *primitiveElement                            `json:"_type,omitempty"`
	TargetProfile                 []Canonical                                  `json:"targetProfile,omitempty"`
	TargetProfilePrimitiveElement []*primitiveElement                          `json:"_targetProfile,omitempty"`
	SearchType                    *Code                                        `json:"searchType,omitempty"`
	SearchTypePrimitiveElement    *primitiveElement                            `json:"_searchType,omitempty"`
	Binding                       *OperationDefinitionParameterBinding         `json:"binding,omitempty"`
	ReferencedFrom                []OperationDefinitionParameterReferencedFrom `json:"referencedFrom,omitempty"`
	Part                          []OperationDefinitionParameter               `json:"part,omitempty"`
}

func (r OperationDefinitionParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r OperationDefinitionParameter) marshalJSON() jsonOperationDefinitionParameter {
	m := jsonOperationDefinitionParameter{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Use.Value != nil {
		m.Use = r.Use
	}
	if r.Use.Id != nil || r.Use.Extension != nil {
		m.UsePrimitiveElement = &primitiveElement{Id: r.Use.Id, Extension: r.Use.Extension}
	}
	if r.Min.Value != nil {
		m.Min = r.Min
	}
	if r.Min.Id != nil || r.Min.Extension != nil {
		m.MinPrimitiveElement = &primitiveElement{Id: r.Min.Id, Extension: r.Min.Extension}
	}
	if r.Max.Value != nil {
		m.Max = r.Max
	}
	if r.Max.Id != nil || r.Max.Extension != nil {
		m.MaxPrimitiveElement = &primitiveElement{Id: r.Max.Id, Extension: r.Max.Extension}
	}
	if r.Documentation != nil && r.Documentation.Value != nil {
		m.Documentation = r.Documentation
	}
	if r.Documentation != nil && (r.Documentation.Id != nil || r.Documentation.Extension != nil) {
		m.DocumentationPrimitiveElement = &primitiveElement{Id: r.Documentation.Id, Extension: r.Documentation.Extension}
	}
	if r.Type != nil && r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	anyTargetProfileValue := false
	for _, e := range r.TargetProfile {
		if e.Value != nil {
			anyTargetProfileValue = true
			break
		}
	}
	if anyTargetProfileValue {
		m.TargetProfile = r.TargetProfile
	}
	anyTargetProfileIdOrExtension := false
	for _, e := range r.TargetProfile {
		if e.Id != nil || e.Extension != nil {
			anyTargetProfileIdOrExtension = true
			break
		}
	}
	if anyTargetProfileIdOrExtension {
		m.TargetProfilePrimitiveElement = make([]*primitiveElement, 0, len(r.TargetProfile))
		for _, e := range r.TargetProfile {
			if e.Id != nil || e.Extension != nil {
				m.TargetProfilePrimitiveElement = append(m.TargetProfilePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.TargetProfilePrimitiveElement = append(m.TargetProfilePrimitiveElement, nil)
			}
		}
	}
	if r.SearchType != nil && r.SearchType.Value != nil {
		m.SearchType = r.SearchType
	}
	if r.SearchType != nil && (r.SearchType.Id != nil || r.SearchType.Extension != nil) {
		m.SearchTypePrimitiveElement = &primitiveElement{Id: r.SearchType.Id, Extension: r.SearchType.Extension}
	}
	m.Binding = r.Binding
	m.ReferencedFrom = r.ReferencedFrom
	m.Part = r.Part
	return m
}
func (r *OperationDefinitionParameter) UnmarshalJSON(b []byte) error {
	var m jsonOperationDefinitionParameter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *OperationDefinitionParameter) unmarshalJSON(m jsonOperationDefinitionParameter) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Use = m.Use
	if m.UsePrimitiveElement != nil {
		r.Use.Id = m.UsePrimitiveElement.Id
		r.Use.Extension = m.UsePrimitiveElement.Extension
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
	r.Documentation = m.Documentation
	if m.DocumentationPrimitiveElement != nil {
		if r.Documentation == nil {
			r.Documentation = &String{}
		}
		r.Documentation.Id = m.DocumentationPrimitiveElement.Id
		r.Documentation.Extension = m.DocumentationPrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		if r.Type == nil {
			r.Type = &Code{}
		}
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.TargetProfile = m.TargetProfile
	for i, e := range m.TargetProfilePrimitiveElement {
		if len(r.TargetProfile) <= i {
			r.TargetProfile = append(r.TargetProfile, Canonical{})
		}
		if e != nil {
			r.TargetProfile[i].Id = e.Id
			r.TargetProfile[i].Extension = e.Extension
		}
	}
	r.SearchType = m.SearchType
	if m.SearchTypePrimitiveElement != nil {
		if r.SearchType == nil {
			r.SearchType = &Code{}
		}
		r.SearchType.Id = m.SearchTypePrimitiveElement.Id
		r.SearchType.Extension = m.SearchTypePrimitiveElement.Extension
	}
	r.Binding = m.Binding
	r.ReferencedFrom = m.ReferencedFrom
	r.Part = m.Part
	return nil
}
func (r OperationDefinitionParameter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Use, xml.StartElement{Name: xml.Name{Local: "use"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Min, xml.StartElement{Name: xml.Name{Local: "min"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Max, xml.StartElement{Name: xml.Name{Local: "max"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Documentation, xml.StartElement{Name: xml.Name{Local: "documentation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TargetProfile, xml.StartElement{Name: xml.Name{Local: "targetProfile"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SearchType, xml.StartElement{Name: xml.Name{Local: "searchType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Binding, xml.StartElement{Name: xml.Name{Local: "binding"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReferencedFrom, xml.StartElement{Name: xml.Name{Local: "referencedFrom"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Part, xml.StartElement{Name: xml.Name{Local: "part"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *OperationDefinitionParameter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = v
			case "use":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Use = v
			case "min":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Min = v
			case "max":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Max = v
			case "documentation":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Documentation = &v
			case "type":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "targetProfile":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TargetProfile = append(r.TargetProfile, v)
			case "searchType":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SearchType = &v
			case "binding":
				var v OperationDefinitionParameterBinding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Binding = &v
			case "referencedFrom":
				var v OperationDefinitionParameterReferencedFrom
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReferencedFrom = append(r.ReferencedFrom, v)
			case "part":
				var v OperationDefinitionParameter
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Part = append(r.Part, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r OperationDefinitionParameter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Binds to a value set if this parameter is coded (code, Coding, CodeableConcept).
type OperationDefinitionParameterBinding struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates the degree of conformance expectations associated with this binding - that is, the degree to which the provided value set must be adhered to in the instances.
	Strength Code
	// Points to the value set or external definition (e.g. implicit value set) that identifies the set of codes to be used.
	ValueSet Canonical
}
type jsonOperationDefinitionParameterBinding struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Strength                 Code              `json:"strength,omitempty"`
	StrengthPrimitiveElement *primitiveElement `json:"_strength,omitempty"`
	ValueSet                 Canonical         `json:"valueSet,omitempty"`
	ValueSetPrimitiveElement *primitiveElement `json:"_valueSet,omitempty"`
}

func (r OperationDefinitionParameterBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r OperationDefinitionParameterBinding) marshalJSON() jsonOperationDefinitionParameterBinding {
	m := jsonOperationDefinitionParameterBinding{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Strength.Value != nil {
		m.Strength = r.Strength
	}
	if r.Strength.Id != nil || r.Strength.Extension != nil {
		m.StrengthPrimitiveElement = &primitiveElement{Id: r.Strength.Id, Extension: r.Strength.Extension}
	}
	if r.ValueSet.Value != nil {
		m.ValueSet = r.ValueSet
	}
	if r.ValueSet.Id != nil || r.ValueSet.Extension != nil {
		m.ValueSetPrimitiveElement = &primitiveElement{Id: r.ValueSet.Id, Extension: r.ValueSet.Extension}
	}
	return m
}
func (r *OperationDefinitionParameterBinding) UnmarshalJSON(b []byte) error {
	var m jsonOperationDefinitionParameterBinding
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *OperationDefinitionParameterBinding) unmarshalJSON(m jsonOperationDefinitionParameterBinding) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Strength = m.Strength
	if m.StrengthPrimitiveElement != nil {
		r.Strength.Id = m.StrengthPrimitiveElement.Id
		r.Strength.Extension = m.StrengthPrimitiveElement.Extension
	}
	r.ValueSet = m.ValueSet
	if m.ValueSetPrimitiveElement != nil {
		r.ValueSet.Id = m.ValueSetPrimitiveElement.Id
		r.ValueSet.Extension = m.ValueSetPrimitiveElement.Extension
	}
	return nil
}
func (r OperationDefinitionParameterBinding) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Strength, xml.StartElement{Name: xml.Name{Local: "strength"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ValueSet, xml.StartElement{Name: xml.Name{Local: "valueSet"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *OperationDefinitionParameterBinding) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "strength":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Strength = v
			case "valueSet":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ValueSet = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r OperationDefinitionParameterBinding) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Identifies other resource parameters within the operation invocation that are expected to resolve to this resource.
type OperationDefinitionParameterReferencedFrom struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The name of the parameter or dot-separated path of parameter names pointing to the resource parameter that is expected to contain a reference to this resource.
	Source String
	// The id of the element in the referencing resource that is expected to resolve to this resource.
	SourceId *String
}
type jsonOperationDefinitionParameterReferencedFrom struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Source                   String            `json:"source,omitempty"`
	SourcePrimitiveElement   *primitiveElement `json:"_source,omitempty"`
	SourceId                 *String           `json:"sourceId,omitempty"`
	SourceIdPrimitiveElement *primitiveElement `json:"_sourceId,omitempty"`
}

func (r OperationDefinitionParameterReferencedFrom) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r OperationDefinitionParameterReferencedFrom) marshalJSON() jsonOperationDefinitionParameterReferencedFrom {
	m := jsonOperationDefinitionParameterReferencedFrom{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Source.Value != nil {
		m.Source = r.Source
	}
	if r.Source.Id != nil || r.Source.Extension != nil {
		m.SourcePrimitiveElement = &primitiveElement{Id: r.Source.Id, Extension: r.Source.Extension}
	}
	if r.SourceId != nil && r.SourceId.Value != nil {
		m.SourceId = r.SourceId
	}
	if r.SourceId != nil && (r.SourceId.Id != nil || r.SourceId.Extension != nil) {
		m.SourceIdPrimitiveElement = &primitiveElement{Id: r.SourceId.Id, Extension: r.SourceId.Extension}
	}
	return m
}
func (r *OperationDefinitionParameterReferencedFrom) UnmarshalJSON(b []byte) error {
	var m jsonOperationDefinitionParameterReferencedFrom
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *OperationDefinitionParameterReferencedFrom) unmarshalJSON(m jsonOperationDefinitionParameterReferencedFrom) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Source = m.Source
	if m.SourcePrimitiveElement != nil {
		r.Source.Id = m.SourcePrimitiveElement.Id
		r.Source.Extension = m.SourcePrimitiveElement.Extension
	}
	r.SourceId = m.SourceId
	if m.SourceIdPrimitiveElement != nil {
		if r.SourceId == nil {
			r.SourceId = &String{}
		}
		r.SourceId.Id = m.SourceIdPrimitiveElement.Id
		r.SourceId.Extension = m.SourceIdPrimitiveElement.Extension
	}
	return nil
}
func (r OperationDefinitionParameterReferencedFrom) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Source, xml.StartElement{Name: xml.Name{Local: "source"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SourceId, xml.StartElement{Name: xml.Name{Local: "sourceId"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *OperationDefinitionParameterReferencedFrom) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "source":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Source = v
			case "sourceId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SourceId = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r OperationDefinitionParameterReferencedFrom) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Defines an appropriate combination of parameters to use when invoking this operation, to help code generators when generating overloaded parameter sets for this operation.
type OperationDefinitionOverload struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Name of parameter to include in overload.
	ParameterName []String
	// Comments to go on overload.
	Comment *String
}
type jsonOperationDefinitionOverload struct {
	Id                            *string             `json:"id,omitempty"`
	Extension                     []Extension         `json:"extension,omitempty"`
	ModifierExtension             []Extension         `json:"modifierExtension,omitempty"`
	ParameterName                 []String            `json:"parameterName,omitempty"`
	ParameterNamePrimitiveElement []*primitiveElement `json:"_parameterName,omitempty"`
	Comment                       *String             `json:"comment,omitempty"`
	CommentPrimitiveElement       *primitiveElement   `json:"_comment,omitempty"`
}

func (r OperationDefinitionOverload) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r OperationDefinitionOverload) marshalJSON() jsonOperationDefinitionOverload {
	m := jsonOperationDefinitionOverload{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	anyParameterNameValue := false
	for _, e := range r.ParameterName {
		if e.Value != nil {
			anyParameterNameValue = true
			break
		}
	}
	if anyParameterNameValue {
		m.ParameterName = r.ParameterName
	}
	anyParameterNameIdOrExtension := false
	for _, e := range r.ParameterName {
		if e.Id != nil || e.Extension != nil {
			anyParameterNameIdOrExtension = true
			break
		}
	}
	if anyParameterNameIdOrExtension {
		m.ParameterNamePrimitiveElement = make([]*primitiveElement, 0, len(r.ParameterName))
		for _, e := range r.ParameterName {
			if e.Id != nil || e.Extension != nil {
				m.ParameterNamePrimitiveElement = append(m.ParameterNamePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ParameterNamePrimitiveElement = append(m.ParameterNamePrimitiveElement, nil)
			}
		}
	}
	if r.Comment != nil && r.Comment.Value != nil {
		m.Comment = r.Comment
	}
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	return m
}
func (r *OperationDefinitionOverload) UnmarshalJSON(b []byte) error {
	var m jsonOperationDefinitionOverload
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *OperationDefinitionOverload) unmarshalJSON(m jsonOperationDefinitionOverload) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ParameterName = m.ParameterName
	for i, e := range m.ParameterNamePrimitiveElement {
		if len(r.ParameterName) <= i {
			r.ParameterName = append(r.ParameterName, String{})
		}
		if e != nil {
			r.ParameterName[i].Id = e.Id
			r.ParameterName[i].Extension = e.Extension
		}
	}
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		if r.Comment == nil {
			r.Comment = &String{}
		}
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	return nil
}
func (r OperationDefinitionOverload) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ParameterName, xml.StartElement{Name: xml.Name{Local: "parameterName"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Comment, xml.StartElement{Name: xml.Name{Local: "comment"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *OperationDefinitionOverload) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "parameterName":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ParameterName = append(r.ParameterName, v)
			case "comment":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Comment = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r OperationDefinitionOverload) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
