package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// The ChargeItemDefinition resource provides the properties that apply to the (billing) codes necessary to calculate costs and prices. The properties may differ largely depending on type and realm, therefore this resource gives only a rough structure and requires profiling for each type of billing code system.
type ChargeItemDefinition struct {
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
	// An absolute URI that is used to identify this charge item definition when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this charge item definition is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the charge item definition is stored on different servers.
	Url Uri
	// A formal identifier that is used to identify this charge item definition when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the charge item definition when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the charge item definition author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active assets.
	Version *String
	// A short, descriptive, user-friendly title for the charge item definition.
	Title *String
	// The URL pointing to an externally-defined charge item definition that is adhered to in whole or in part by this definition.
	DerivedFromUri []Uri
	// A larger definition of which this particular definition is a component or step.
	PartOf []Canonical
	// As new versions of a protocol or guideline are defined, allows identification of what versions are replaced by a new instance.
	Replaces []Canonical
	// The current state of the ChargeItemDefinition.
	Status Code
	// A Boolean value to indicate that this charge item definition is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the charge item definition was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the charge item definition changes.
	Date *DateTime
	// The name of the organization or individual that published the charge item definition.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the charge item definition from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate charge item definition instances.
	UseContext []UsageContext
	// A legal or geographic region in which the charge item definition is intended to be used.
	Jurisdiction []CodeableConcept
	// A copyright statement relating to the charge item definition and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the charge item definition.
	Copyright *Markdown
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *Date
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *Date
	// The period during which the charge item definition content was or is planned to be in active use.
	EffectivePeriod *Period
	// The defined billing details in this resource pertain to the given billing code.
	Code *CodeableConcept
	// The defined billing details in this resource pertain to the given product instance(s).
	Instance []Reference
	// Expressions that describe applicability criteria for the billing code.
	Applicability []ChargeItemDefinitionApplicability
	// Group of properties which are applicable under the same conditions. If no applicability rules are established for the group, then all properties always apply.
	PropertyGroup []ChargeItemDefinitionPropertyGroup
}

func (r ChargeItemDefinition) ResourceType() string {
	return "ChargeItemDefinition"
}
func (r ChargeItemDefinition) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonChargeItemDefinition struct {
	ResourceType                   string                              `json:"resourceType"`
	Id                             *Id                                 `json:"id,omitempty"`
	IdPrimitiveElement             *primitiveElement                   `json:"_id,omitempty"`
	Meta                           *Meta                               `json:"meta,omitempty"`
	ImplicitRules                  *Uri                                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement  *primitiveElement                   `json:"_implicitRules,omitempty"`
	Language                       *Code                               `json:"language,omitempty"`
	LanguagePrimitiveElement       *primitiveElement                   `json:"_language,omitempty"`
	Text                           *Narrative                          `json:"text,omitempty"`
	Contained                      []containedResource                 `json:"contained,omitempty"`
	Extension                      []Extension                         `json:"extension,omitempty"`
	ModifierExtension              []Extension                         `json:"modifierExtension,omitempty"`
	Url                            Uri                                 `json:"url,omitempty"`
	UrlPrimitiveElement            *primitiveElement                   `json:"_url,omitempty"`
	Identifier                     []Identifier                        `json:"identifier,omitempty"`
	Version                        *String                             `json:"version,omitempty"`
	VersionPrimitiveElement        *primitiveElement                   `json:"_version,omitempty"`
	Title                          *String                             `json:"title,omitempty"`
	TitlePrimitiveElement          *primitiveElement                   `json:"_title,omitempty"`
	DerivedFromUri                 []Uri                               `json:"derivedFromUri,omitempty"`
	DerivedFromUriPrimitiveElement []*primitiveElement                 `json:"_derivedFromUri,omitempty"`
	PartOf                         []Canonical                         `json:"partOf,omitempty"`
	PartOfPrimitiveElement         []*primitiveElement                 `json:"_partOf,omitempty"`
	Replaces                       []Canonical                         `json:"replaces,omitempty"`
	ReplacesPrimitiveElement       []*primitiveElement                 `json:"_replaces,omitempty"`
	Status                         Code                                `json:"status,omitempty"`
	StatusPrimitiveElement         *primitiveElement                   `json:"_status,omitempty"`
	Experimental                   *Boolean                            `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement   *primitiveElement                   `json:"_experimental,omitempty"`
	Date                           *DateTime                           `json:"date,omitempty"`
	DatePrimitiveElement           *primitiveElement                   `json:"_date,omitempty"`
	Publisher                      *String                             `json:"publisher,omitempty"`
	PublisherPrimitiveElement      *primitiveElement                   `json:"_publisher,omitempty"`
	Contact                        []ContactDetail                     `json:"contact,omitempty"`
	Description                    *Markdown                           `json:"description,omitempty"`
	DescriptionPrimitiveElement    *primitiveElement                   `json:"_description,omitempty"`
	UseContext                     []UsageContext                      `json:"useContext,omitempty"`
	Jurisdiction                   []CodeableConcept                   `json:"jurisdiction,omitempty"`
	Copyright                      *Markdown                           `json:"copyright,omitempty"`
	CopyrightPrimitiveElement      *primitiveElement                   `json:"_copyright,omitempty"`
	ApprovalDate                   *Date                               `json:"approvalDate,omitempty"`
	ApprovalDatePrimitiveElement   *primitiveElement                   `json:"_approvalDate,omitempty"`
	LastReviewDate                 *Date                               `json:"lastReviewDate,omitempty"`
	LastReviewDatePrimitiveElement *primitiveElement                   `json:"_lastReviewDate,omitempty"`
	EffectivePeriod                *Period                             `json:"effectivePeriod,omitempty"`
	Code                           *CodeableConcept                    `json:"code,omitempty"`
	Instance                       []Reference                         `json:"instance,omitempty"`
	Applicability                  []ChargeItemDefinitionApplicability `json:"applicability,omitempty"`
	PropertyGroup                  []ChargeItemDefinitionPropertyGroup `json:"propertyGroup,omitempty"`
}

func (r ChargeItemDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ChargeItemDefinition) marshalJSON() jsonChargeItemDefinition {
	m := jsonChargeItemDefinition{}
	m.ResourceType = "ChargeItemDefinition"
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
	m.Title = r.Title
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	m.DerivedFromUri = r.DerivedFromUri
	anyDerivedFromUriIdOrExtension := false
	for _, e := range r.DerivedFromUri {
		if e.Id != nil || e.Extension != nil {
			anyDerivedFromUriIdOrExtension = true
			break
		}
	}
	if anyDerivedFromUriIdOrExtension {
		m.DerivedFromUriPrimitiveElement = make([]*primitiveElement, 0, len(r.DerivedFromUri))
		for _, e := range r.DerivedFromUri {
			if e.Id != nil || e.Extension != nil {
				m.DerivedFromUriPrimitiveElement = append(m.DerivedFromUriPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.DerivedFromUriPrimitiveElement = append(m.DerivedFromUriPrimitiveElement, nil)
			}
		}
	}
	m.PartOf = r.PartOf
	anyPartOfIdOrExtension := false
	for _, e := range r.PartOf {
		if e.Id != nil || e.Extension != nil {
			anyPartOfIdOrExtension = true
			break
		}
	}
	if anyPartOfIdOrExtension {
		m.PartOfPrimitiveElement = make([]*primitiveElement, 0, len(r.PartOf))
		for _, e := range r.PartOf {
			if e.Id != nil || e.Extension != nil {
				m.PartOfPrimitiveElement = append(m.PartOfPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.PartOfPrimitiveElement = append(m.PartOfPrimitiveElement, nil)
			}
		}
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
	m.Copyright = r.Copyright
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		m.CopyrightPrimitiveElement = &primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
	}
	m.ApprovalDate = r.ApprovalDate
	if r.ApprovalDate != nil && (r.ApprovalDate.Id != nil || r.ApprovalDate.Extension != nil) {
		m.ApprovalDatePrimitiveElement = &primitiveElement{Id: r.ApprovalDate.Id, Extension: r.ApprovalDate.Extension}
	}
	m.LastReviewDate = r.LastReviewDate
	if r.LastReviewDate != nil && (r.LastReviewDate.Id != nil || r.LastReviewDate.Extension != nil) {
		m.LastReviewDatePrimitiveElement = &primitiveElement{Id: r.LastReviewDate.Id, Extension: r.LastReviewDate.Extension}
	}
	m.EffectivePeriod = r.EffectivePeriod
	m.Code = r.Code
	m.Instance = r.Instance
	m.Applicability = r.Applicability
	m.PropertyGroup = r.PropertyGroup
	return m
}
func (r *ChargeItemDefinition) UnmarshalJSON(b []byte) error {
	var m jsonChargeItemDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ChargeItemDefinition) unmarshalJSON(m jsonChargeItemDefinition) error {
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
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.DerivedFromUri = m.DerivedFromUri
	for i, e := range m.DerivedFromUriPrimitiveElement {
		if len(r.DerivedFromUri) > i {
			r.DerivedFromUri[i].Id = e.Id
			r.DerivedFromUri[i].Extension = e.Extension
		} else {
			r.DerivedFromUri = append(r.DerivedFromUri, Uri{Id: e.Id, Extension: e.Extension})
		}
	}
	r.PartOf = m.PartOf
	for i, e := range m.PartOfPrimitiveElement {
		if len(r.PartOf) > i {
			r.PartOf[i].Id = e.Id
			r.PartOf[i].Extension = e.Extension
		} else {
			r.PartOf = append(r.PartOf, Canonical{Id: e.Id, Extension: e.Extension})
		}
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
	r.Copyright = m.Copyright
	if m.CopyrightPrimitiveElement != nil {
		r.Copyright.Id = m.CopyrightPrimitiveElement.Id
		r.Copyright.Extension = m.CopyrightPrimitiveElement.Extension
	}
	r.ApprovalDate = m.ApprovalDate
	if m.ApprovalDatePrimitiveElement != nil {
		r.ApprovalDate.Id = m.ApprovalDatePrimitiveElement.Id
		r.ApprovalDate.Extension = m.ApprovalDatePrimitiveElement.Extension
	}
	r.LastReviewDate = m.LastReviewDate
	if m.LastReviewDatePrimitiveElement != nil {
		r.LastReviewDate.Id = m.LastReviewDatePrimitiveElement.Id
		r.LastReviewDate.Extension = m.LastReviewDatePrimitiveElement.Extension
	}
	r.EffectivePeriod = m.EffectivePeriod
	r.Code = m.Code
	r.Instance = m.Instance
	r.Applicability = m.Applicability
	r.PropertyGroup = m.PropertyGroup
	return nil
}
func (r ChargeItemDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Expressions that describe applicability criteria for the billing code.
type ChargeItemDefinitionApplicability struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A brief, natural language description of the condition that effectively communicates the intended semantics.
	Description *String
	// The media type of the language for the expression, e.g. "text/cql" for Clinical Query Language expressions or "text/fhirpath" for FHIRPath expressions.
	Language *String
	// An expression that returns true or false, indicating whether the condition is satisfied. When using FHIRPath expressions, the %context environment variable must be replaced at runtime with the ChargeItem resource to which this definition is applied.
	Expression *String
}
type jsonChargeItemDefinitionApplicability struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Description                 *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
	Language                    *String           `json:"language,omitempty"`
	LanguagePrimitiveElement    *primitiveElement `json:"_language,omitempty"`
	Expression                  *String           `json:"expression,omitempty"`
	ExpressionPrimitiveElement  *primitiveElement `json:"_expression,omitempty"`
}

func (r ChargeItemDefinitionApplicability) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ChargeItemDefinitionApplicability) marshalJSON() jsonChargeItemDefinitionApplicability {
	m := jsonChargeItemDefinitionApplicability{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Language = r.Language
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.Expression = r.Expression
	if r.Expression != nil && (r.Expression.Id != nil || r.Expression.Extension != nil) {
		m.ExpressionPrimitiveElement = &primitiveElement{Id: r.Expression.Id, Extension: r.Expression.Extension}
	}
	return m
}
func (r *ChargeItemDefinitionApplicability) UnmarshalJSON(b []byte) error {
	var m jsonChargeItemDefinitionApplicability
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ChargeItemDefinitionApplicability) unmarshalJSON(m jsonChargeItemDefinitionApplicability) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Expression = m.Expression
	if m.ExpressionPrimitiveElement != nil {
		r.Expression.Id = m.ExpressionPrimitiveElement.Id
		r.Expression.Extension = m.ExpressionPrimitiveElement.Extension
	}
	return nil
}
func (r ChargeItemDefinitionApplicability) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Group of properties which are applicable under the same conditions. If no applicability rules are established for the group, then all properties always apply.
type ChargeItemDefinitionPropertyGroup struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Expressions that describe applicability criteria for the priceComponent.
	Applicability []ChargeItemDefinitionApplicability
	// The price for a ChargeItem may be calculated as a base price with surcharges/deductions that apply in certain conditions. A ChargeItemDefinition resource that defines the prices, factors and conditions that apply to a billing code is currently under development. The priceComponent element can be used to offer transparency to the recipient of the Invoice of how the prices have been calculated.
	PriceComponent []ChargeItemDefinitionPropertyGroupPriceComponent
}
type jsonChargeItemDefinitionPropertyGroup struct {
	Id                *string                                           `json:"id,omitempty"`
	Extension         []Extension                                       `json:"extension,omitempty"`
	ModifierExtension []Extension                                       `json:"modifierExtension,omitempty"`
	Applicability     []ChargeItemDefinitionApplicability               `json:"applicability,omitempty"`
	PriceComponent    []ChargeItemDefinitionPropertyGroupPriceComponent `json:"priceComponent,omitempty"`
}

func (r ChargeItemDefinitionPropertyGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ChargeItemDefinitionPropertyGroup) marshalJSON() jsonChargeItemDefinitionPropertyGroup {
	m := jsonChargeItemDefinitionPropertyGroup{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Applicability = r.Applicability
	m.PriceComponent = r.PriceComponent
	return m
}
func (r *ChargeItemDefinitionPropertyGroup) UnmarshalJSON(b []byte) error {
	var m jsonChargeItemDefinitionPropertyGroup
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ChargeItemDefinitionPropertyGroup) unmarshalJSON(m jsonChargeItemDefinitionPropertyGroup) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Applicability = m.Applicability
	r.PriceComponent = m.PriceComponent
	return nil
}
func (r ChargeItemDefinitionPropertyGroup) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The price for a ChargeItem may be calculated as a base price with surcharges/deductions that apply in certain conditions. A ChargeItemDefinition resource that defines the prices, factors and conditions that apply to a billing code is currently under development. The priceComponent element can be used to offer transparency to the recipient of the Invoice of how the prices have been calculated.
type ChargeItemDefinitionPropertyGroupPriceComponent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// This code identifies the type of the component.
	Type Code
	// A code that identifies the component. Codes may be used to differentiate between kinds of taxes, surcharges, discounts etc.
	Code *CodeableConcept
	// The factor that has been applied on the base price for calculating this component.
	Factor *Decimal
	// The amount calculated for this component.
	Amount *Money
}
type jsonChargeItemDefinitionPropertyGroupPriceComponent struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Type                   Code              `json:"type,omitempty"`
	TypePrimitiveElement   *primitiveElement `json:"_type,omitempty"`
	Code                   *CodeableConcept  `json:"code,omitempty"`
	Factor                 *Decimal          `json:"factor,omitempty"`
	FactorPrimitiveElement *primitiveElement `json:"_factor,omitempty"`
	Amount                 *Money            `json:"amount,omitempty"`
}

func (r ChargeItemDefinitionPropertyGroupPriceComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ChargeItemDefinitionPropertyGroupPriceComponent) marshalJSON() jsonChargeItemDefinitionPropertyGroupPriceComponent {
	m := jsonChargeItemDefinitionPropertyGroupPriceComponent{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Code = r.Code
	m.Factor = r.Factor
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		m.FactorPrimitiveElement = &primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
	}
	m.Amount = r.Amount
	return m
}
func (r *ChargeItemDefinitionPropertyGroupPriceComponent) UnmarshalJSON(b []byte) error {
	var m jsonChargeItemDefinitionPropertyGroupPriceComponent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ChargeItemDefinitionPropertyGroupPriceComponent) unmarshalJSON(m jsonChargeItemDefinitionPropertyGroupPriceComponent) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Code = m.Code
	r.Factor = m.Factor
	if m.FactorPrimitiveElement != nil {
		r.Factor.Id = m.FactorPrimitiveElement.Id
		r.Factor.Extension = m.FactorPrimitiveElement.Extension
	}
	r.Amount = m.Amount
	return nil
}
func (r ChargeItemDefinitionPropertyGroupPriceComponent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
