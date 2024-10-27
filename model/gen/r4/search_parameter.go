package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A search parameter that defines a named search item that can be used to search/filter on a resource.
type SearchParameter struct {
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
	// An absolute URI that is used to identify this search parameter when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this search parameter is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the search parameter is stored on different servers.
	Url Uri
	// The identifier that is used to identify this version of the search parameter when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the search parameter author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the search parameter. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name String
	// Where this search parameter is originally defined. If a derivedFrom is provided, then the details in the search parameter must be consistent with the definition from which it is defined. i.e. the parameter should have the same meaning, and (usually) the functionality should be a proper subset of the underlying search parameter.
	DerivedFrom *Canonical
	// The status of this search parameter. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this search parameter is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the search parameter was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the search parameter changes.
	Date *DateTime
	// The name of the organization or individual that published the search parameter.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// And how it used.
	Description Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate search parameter instances.
	UseContext []UsageContext
	// A legal or geographic region in which the search parameter is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this search parameter is needed and why it has been designed as it has.
	Purpose *Markdown
	// The code used in the URL or the parameter name in a parameters resource for this search parameter.
	Code Code
	// The base resource type(s) that this search parameter can be used against.
	Base []Code
	// The type of value that a search parameter may contain, and how the content is interpreted.
	Type Code
	// A FHIRPath expression that returns a set of elements for the search parameter.
	Expression *String
	// An XPath expression that returns a set of elements for the search parameter.
	Xpath *String
	// How the search parameter relates to the set of elements returned by evaluating the xpath query.
	XpathUsage *Code
	// Types of resource (if a resource is referenced).
	Target []Code
	// Whether multiple values are allowed for each time the parameter exists. Values are separated by commas, and the parameter matches if any of the values match.
	MultipleOr *Boolean
	// Whether multiple parameters are allowed - e.g. more than one parameter with the same name. The search matches if all the parameters match.
	MultipleAnd *Boolean
	// Comparators supported for the search parameter.
	Comparator []Code
	// A modifier supported for the search parameter.
	Modifier []Code
	// Contains the names of any search parameters which may be chained to the containing search parameter. Chained parameters may be added to search parameters of type reference and specify that resources will only be returned if they contain a reference to a resource which matches the chained parameter value. Values for this field should be drawn from SearchParameter.code for a parameter on the target resource type.
	Chain []String
	// Used to define the parts of a composite search parameter.
	Component []SearchParameterComponent
}

func (r SearchParameter) ResourceType() string {
	return "SearchParameter"
}
func (r SearchParameter) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonSearchParameter struct {
	ResourceType                  string                     `json:"resourceType"`
	Id                            *Id                        `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement          `json:"_id,omitempty"`
	Meta                          *Meta                      `json:"meta,omitempty"`
	ImplicitRules                 *Uri                       `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement          `json:"_implicitRules,omitempty"`
	Language                      *Code                      `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement          `json:"_language,omitempty"`
	Text                          *Narrative                 `json:"text,omitempty"`
	Contained                     []ContainedResource        `json:"contained,omitempty"`
	Extension                     []Extension                `json:"extension,omitempty"`
	ModifierExtension             []Extension                `json:"modifierExtension,omitempty"`
	Url                           Uri                        `json:"url,omitempty"`
	UrlPrimitiveElement           *primitiveElement          `json:"_url,omitempty"`
	Version                       *String                    `json:"version,omitempty"`
	VersionPrimitiveElement       *primitiveElement          `json:"_version,omitempty"`
	Name                          String                     `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement          `json:"_name,omitempty"`
	DerivedFrom                   *Canonical                 `json:"derivedFrom,omitempty"`
	DerivedFromPrimitiveElement   *primitiveElement          `json:"_derivedFrom,omitempty"`
	Status                        Code                       `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement          `json:"_status,omitempty"`
	Experimental                  *Boolean                   `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement  *primitiveElement          `json:"_experimental,omitempty"`
	Date                          *DateTime                  `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement          `json:"_date,omitempty"`
	Publisher                     *String                    `json:"publisher,omitempty"`
	PublisherPrimitiveElement     *primitiveElement          `json:"_publisher,omitempty"`
	Contact                       []ContactDetail            `json:"contact,omitempty"`
	Description                   Markdown                   `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement          `json:"_description,omitempty"`
	UseContext                    []UsageContext             `json:"useContext,omitempty"`
	Jurisdiction                  []CodeableConcept          `json:"jurisdiction,omitempty"`
	Purpose                       *Markdown                  `json:"purpose,omitempty"`
	PurposePrimitiveElement       *primitiveElement          `json:"_purpose,omitempty"`
	Code                          Code                       `json:"code,omitempty"`
	CodePrimitiveElement          *primitiveElement          `json:"_code,omitempty"`
	Base                          []Code                     `json:"base,omitempty"`
	BasePrimitiveElement          []*primitiveElement        `json:"_base,omitempty"`
	Type                          Code                       `json:"type,omitempty"`
	TypePrimitiveElement          *primitiveElement          `json:"_type,omitempty"`
	Expression                    *String                    `json:"expression,omitempty"`
	ExpressionPrimitiveElement    *primitiveElement          `json:"_expression,omitempty"`
	Xpath                         *String                    `json:"xpath,omitempty"`
	XpathPrimitiveElement         *primitiveElement          `json:"_xpath,omitempty"`
	XpathUsage                    *Code                      `json:"xpathUsage,omitempty"`
	XpathUsagePrimitiveElement    *primitiveElement          `json:"_xpathUsage,omitempty"`
	Target                        []Code                     `json:"target,omitempty"`
	TargetPrimitiveElement        []*primitiveElement        `json:"_target,omitempty"`
	MultipleOr                    *Boolean                   `json:"multipleOr,omitempty"`
	MultipleOrPrimitiveElement    *primitiveElement          `json:"_multipleOr,omitempty"`
	MultipleAnd                   *Boolean                   `json:"multipleAnd,omitempty"`
	MultipleAndPrimitiveElement   *primitiveElement          `json:"_multipleAnd,omitempty"`
	Comparator                    []Code                     `json:"comparator,omitempty"`
	ComparatorPrimitiveElement    []*primitiveElement        `json:"_comparator,omitempty"`
	Modifier                      []Code                     `json:"modifier,omitempty"`
	ModifierPrimitiveElement      []*primitiveElement        `json:"_modifier,omitempty"`
	Chain                         []String                   `json:"chain,omitempty"`
	ChainPrimitiveElement         []*primitiveElement        `json:"_chain,omitempty"`
	Component                     []SearchParameterComponent `json:"component,omitempty"`
}

func (r SearchParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SearchParameter) marshalJSON() jsonSearchParameter {
	m := jsonSearchParameter{}
	m.ResourceType = "SearchParameter"
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
	if r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url.Id != nil || r.Url.Extension != nil {
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
	if r.DerivedFrom != nil && r.DerivedFrom.Value != nil {
		m.DerivedFrom = r.DerivedFrom
	}
	if r.DerivedFrom != nil && (r.DerivedFrom.Id != nil || r.DerivedFrom.Extension != nil) {
		m.DerivedFromPrimitiveElement = &primitiveElement{Id: r.DerivedFrom.Id, Extension: r.DerivedFrom.Extension}
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
	if r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description.Id != nil || r.Description.Extension != nil {
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
	if r.Code.Value != nil {
		m.Code = r.Code
	}
	if r.Code.Id != nil || r.Code.Extension != nil {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	anyBaseValue := false
	for _, e := range r.Base {
		if e.Value != nil {
			anyBaseValue = true
			break
		}
	}
	if anyBaseValue {
		m.Base = r.Base
	}
	anyBaseIdOrExtension := false
	for _, e := range r.Base {
		if e.Id != nil || e.Extension != nil {
			anyBaseIdOrExtension = true
			break
		}
	}
	if anyBaseIdOrExtension {
		m.BasePrimitiveElement = make([]*primitiveElement, 0, len(r.Base))
		for _, e := range r.Base {
			if e.Id != nil || e.Extension != nil {
				m.BasePrimitiveElement = append(m.BasePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.BasePrimitiveElement = append(m.BasePrimitiveElement, nil)
			}
		}
	}
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.Expression != nil && r.Expression.Value != nil {
		m.Expression = r.Expression
	}
	if r.Expression != nil && (r.Expression.Id != nil || r.Expression.Extension != nil) {
		m.ExpressionPrimitiveElement = &primitiveElement{Id: r.Expression.Id, Extension: r.Expression.Extension}
	}
	if r.Xpath != nil && r.Xpath.Value != nil {
		m.Xpath = r.Xpath
	}
	if r.Xpath != nil && (r.Xpath.Id != nil || r.Xpath.Extension != nil) {
		m.XpathPrimitiveElement = &primitiveElement{Id: r.Xpath.Id, Extension: r.Xpath.Extension}
	}
	if r.XpathUsage != nil && r.XpathUsage.Value != nil {
		m.XpathUsage = r.XpathUsage
	}
	if r.XpathUsage != nil && (r.XpathUsage.Id != nil || r.XpathUsage.Extension != nil) {
		m.XpathUsagePrimitiveElement = &primitiveElement{Id: r.XpathUsage.Id, Extension: r.XpathUsage.Extension}
	}
	anyTargetValue := false
	for _, e := range r.Target {
		if e.Value != nil {
			anyTargetValue = true
			break
		}
	}
	if anyTargetValue {
		m.Target = r.Target
	}
	anyTargetIdOrExtension := false
	for _, e := range r.Target {
		if e.Id != nil || e.Extension != nil {
			anyTargetIdOrExtension = true
			break
		}
	}
	if anyTargetIdOrExtension {
		m.TargetPrimitiveElement = make([]*primitiveElement, 0, len(r.Target))
		for _, e := range r.Target {
			if e.Id != nil || e.Extension != nil {
				m.TargetPrimitiveElement = append(m.TargetPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.TargetPrimitiveElement = append(m.TargetPrimitiveElement, nil)
			}
		}
	}
	if r.MultipleOr != nil && r.MultipleOr.Value != nil {
		m.MultipleOr = r.MultipleOr
	}
	if r.MultipleOr != nil && (r.MultipleOr.Id != nil || r.MultipleOr.Extension != nil) {
		m.MultipleOrPrimitiveElement = &primitiveElement{Id: r.MultipleOr.Id, Extension: r.MultipleOr.Extension}
	}
	if r.MultipleAnd != nil && r.MultipleAnd.Value != nil {
		m.MultipleAnd = r.MultipleAnd
	}
	if r.MultipleAnd != nil && (r.MultipleAnd.Id != nil || r.MultipleAnd.Extension != nil) {
		m.MultipleAndPrimitiveElement = &primitiveElement{Id: r.MultipleAnd.Id, Extension: r.MultipleAnd.Extension}
	}
	anyComparatorValue := false
	for _, e := range r.Comparator {
		if e.Value != nil {
			anyComparatorValue = true
			break
		}
	}
	if anyComparatorValue {
		m.Comparator = r.Comparator
	}
	anyComparatorIdOrExtension := false
	for _, e := range r.Comparator {
		if e.Id != nil || e.Extension != nil {
			anyComparatorIdOrExtension = true
			break
		}
	}
	if anyComparatorIdOrExtension {
		m.ComparatorPrimitiveElement = make([]*primitiveElement, 0, len(r.Comparator))
		for _, e := range r.Comparator {
			if e.Id != nil || e.Extension != nil {
				m.ComparatorPrimitiveElement = append(m.ComparatorPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ComparatorPrimitiveElement = append(m.ComparatorPrimitiveElement, nil)
			}
		}
	}
	anyModifierValue := false
	for _, e := range r.Modifier {
		if e.Value != nil {
			anyModifierValue = true
			break
		}
	}
	if anyModifierValue {
		m.Modifier = r.Modifier
	}
	anyModifierIdOrExtension := false
	for _, e := range r.Modifier {
		if e.Id != nil || e.Extension != nil {
			anyModifierIdOrExtension = true
			break
		}
	}
	if anyModifierIdOrExtension {
		m.ModifierPrimitiveElement = make([]*primitiveElement, 0, len(r.Modifier))
		for _, e := range r.Modifier {
			if e.Id != nil || e.Extension != nil {
				m.ModifierPrimitiveElement = append(m.ModifierPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ModifierPrimitiveElement = append(m.ModifierPrimitiveElement, nil)
			}
		}
	}
	anyChainValue := false
	for _, e := range r.Chain {
		if e.Value != nil {
			anyChainValue = true
			break
		}
	}
	if anyChainValue {
		m.Chain = r.Chain
	}
	anyChainIdOrExtension := false
	for _, e := range r.Chain {
		if e.Id != nil || e.Extension != nil {
			anyChainIdOrExtension = true
			break
		}
	}
	if anyChainIdOrExtension {
		m.ChainPrimitiveElement = make([]*primitiveElement, 0, len(r.Chain))
		for _, e := range r.Chain {
			if e.Id != nil || e.Extension != nil {
				m.ChainPrimitiveElement = append(m.ChainPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ChainPrimitiveElement = append(m.ChainPrimitiveElement, nil)
			}
		}
	}
	m.Component = r.Component
	return m
}
func (r *SearchParameter) UnmarshalJSON(b []byte) error {
	var m jsonSearchParameter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SearchParameter) unmarshalJSON(m jsonSearchParameter) error {
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
	r.DerivedFrom = m.DerivedFrom
	if m.DerivedFromPrimitiveElement != nil {
		if r.DerivedFrom == nil {
			r.DerivedFrom = &Canonical{}
		}
		r.DerivedFrom.Id = m.DerivedFromPrimitiveElement.Id
		r.DerivedFrom.Extension = m.DerivedFromPrimitiveElement.Extension
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
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Base = m.Base
	for i, e := range m.BasePrimitiveElement {
		if len(r.Base) <= i {
			r.Base = append(r.Base, Code{})
		}
		if e != nil {
			r.Base[i].Id = e.Id
			r.Base[i].Extension = e.Extension
		}
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Expression = m.Expression
	if m.ExpressionPrimitiveElement != nil {
		if r.Expression == nil {
			r.Expression = &String{}
		}
		r.Expression.Id = m.ExpressionPrimitiveElement.Id
		r.Expression.Extension = m.ExpressionPrimitiveElement.Extension
	}
	r.Xpath = m.Xpath
	if m.XpathPrimitiveElement != nil {
		if r.Xpath == nil {
			r.Xpath = &String{}
		}
		r.Xpath.Id = m.XpathPrimitiveElement.Id
		r.Xpath.Extension = m.XpathPrimitiveElement.Extension
	}
	r.XpathUsage = m.XpathUsage
	if m.XpathUsagePrimitiveElement != nil {
		if r.XpathUsage == nil {
			r.XpathUsage = &Code{}
		}
		r.XpathUsage.Id = m.XpathUsagePrimitiveElement.Id
		r.XpathUsage.Extension = m.XpathUsagePrimitiveElement.Extension
	}
	r.Target = m.Target
	for i, e := range m.TargetPrimitiveElement {
		if len(r.Target) <= i {
			r.Target = append(r.Target, Code{})
		}
		if e != nil {
			r.Target[i].Id = e.Id
			r.Target[i].Extension = e.Extension
		}
	}
	r.MultipleOr = m.MultipleOr
	if m.MultipleOrPrimitiveElement != nil {
		if r.MultipleOr == nil {
			r.MultipleOr = &Boolean{}
		}
		r.MultipleOr.Id = m.MultipleOrPrimitiveElement.Id
		r.MultipleOr.Extension = m.MultipleOrPrimitiveElement.Extension
	}
	r.MultipleAnd = m.MultipleAnd
	if m.MultipleAndPrimitiveElement != nil {
		if r.MultipleAnd == nil {
			r.MultipleAnd = &Boolean{}
		}
		r.MultipleAnd.Id = m.MultipleAndPrimitiveElement.Id
		r.MultipleAnd.Extension = m.MultipleAndPrimitiveElement.Extension
	}
	r.Comparator = m.Comparator
	for i, e := range m.ComparatorPrimitiveElement {
		if len(r.Comparator) <= i {
			r.Comparator = append(r.Comparator, Code{})
		}
		if e != nil {
			r.Comparator[i].Id = e.Id
			r.Comparator[i].Extension = e.Extension
		}
	}
	r.Modifier = m.Modifier
	for i, e := range m.ModifierPrimitiveElement {
		if len(r.Modifier) <= i {
			r.Modifier = append(r.Modifier, Code{})
		}
		if e != nil {
			r.Modifier[i].Id = e.Id
			r.Modifier[i].Extension = e.Extension
		}
	}
	r.Chain = m.Chain
	for i, e := range m.ChainPrimitiveElement {
		if len(r.Chain) <= i {
			r.Chain = append(r.Chain, String{})
		}
		if e != nil {
			r.Chain[i].Id = e.Id
			r.Chain[i].Extension = e.Extension
		}
	}
	r.Component = m.Component
	return nil
}
func (r SearchParameter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Used to define the parts of a composite search parameter.
type SearchParameterComponent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The definition of the search parameter that describes this part.
	Definition Canonical
	// A sub-expression that defines how to extract values for this component from the output of the main SearchParameter.expression.
	Expression String
}
type jsonSearchParameterComponent struct {
	Id                         *string           `json:"id,omitempty"`
	Extension                  []Extension       `json:"extension,omitempty"`
	ModifierExtension          []Extension       `json:"modifierExtension,omitempty"`
	Definition                 Canonical         `json:"definition,omitempty"`
	DefinitionPrimitiveElement *primitiveElement `json:"_definition,omitempty"`
	Expression                 String            `json:"expression,omitempty"`
	ExpressionPrimitiveElement *primitiveElement `json:"_expression,omitempty"`
}

func (r SearchParameterComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SearchParameterComponent) marshalJSON() jsonSearchParameterComponent {
	m := jsonSearchParameterComponent{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Definition.Value != nil {
		m.Definition = r.Definition
	}
	if r.Definition.Id != nil || r.Definition.Extension != nil {
		m.DefinitionPrimitiveElement = &primitiveElement{Id: r.Definition.Id, Extension: r.Definition.Extension}
	}
	if r.Expression.Value != nil {
		m.Expression = r.Expression
	}
	if r.Expression.Id != nil || r.Expression.Extension != nil {
		m.ExpressionPrimitiveElement = &primitiveElement{Id: r.Expression.Id, Extension: r.Expression.Extension}
	}
	return m
}
func (r *SearchParameterComponent) UnmarshalJSON(b []byte) error {
	var m jsonSearchParameterComponent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SearchParameterComponent) unmarshalJSON(m jsonSearchParameterComponent) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Definition = m.Definition
	if m.DefinitionPrimitiveElement != nil {
		r.Definition.Id = m.DefinitionPrimitiveElement.Id
		r.Definition.Extension = m.DefinitionPrimitiveElement.Extension
	}
	r.Expression = m.Expression
	if m.ExpressionPrimitiveElement != nil {
		r.Expression.Id = m.ExpressionPrimitiveElement.Id
		r.Expression.Extension = m.ExpressionPrimitiveElement.Extension
	}
	return nil
}
func (r SearchParameterComponent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
