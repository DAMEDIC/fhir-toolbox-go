package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// Details of a Health Insurance product/plan provided by an organization.
type InsurancePlan struct {
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
	// Business identifiers assigned to this health insurance product which remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// The current state of the health insurance product.
	Status *Code
	// The kind of health insurance product.
	Type []CodeableConcept
	// Official name of the health insurance product as designated by the owner.
	Name *String
	// A list of alternate names that the product is known as, or was known as in the past.
	Alias []String
	// The period of time that the health insurance product is available.
	Period *Period
	// The entity that is providing  the health insurance product and underwriting the risk.  This is typically an insurance carriers, other third-party payers, or health plan sponsors comonly referred to as 'payers'.
	OwnedBy *Reference
	// An organization which administer other services such as underwriting, customer service and/or claims processing on behalf of the health insurance product owner.
	AdministeredBy *Reference
	// The geographic region in which a health insurance product's benefits apply.
	CoverageArea []Reference
	// The contact for the health insurance product for a certain purpose.
	Contact []InsurancePlanContact
	// The technical endpoints providing access to services operated for the health insurance product.
	Endpoint []Reference
	// Reference to the network included in the health insurance product.
	Network []Reference
	// Details about the coverage offered by the insurance product.
	Coverage []InsurancePlanCoverage
	// Details about an insurance plan.
	Plan []InsurancePlanPlan
}

func (r InsurancePlan) ResourceType() string {
	return "InsurancePlan"
}

type jsonInsurancePlan struct {
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
	Identifier                    []Identifier            `json:"identifier,omitempty"`
	Status                        *Code                   `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement       `json:"_status,omitempty"`
	Type                          []CodeableConcept       `json:"type,omitempty"`
	Name                          *String                 `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement       `json:"_name,omitempty"`
	Alias                         []String                `json:"alias,omitempty"`
	AliasPrimitiveElement         []*primitiveElement     `json:"_alias,omitempty"`
	Period                        *Period                 `json:"period,omitempty"`
	OwnedBy                       *Reference              `json:"ownedBy,omitempty"`
	AdministeredBy                *Reference              `json:"administeredBy,omitempty"`
	CoverageArea                  []Reference             `json:"coverageArea,omitempty"`
	Contact                       []InsurancePlanContact  `json:"contact,omitempty"`
	Endpoint                      []Reference             `json:"endpoint,omitempty"`
	Network                       []Reference             `json:"network,omitempty"`
	Coverage                      []InsurancePlanCoverage `json:"coverage,omitempty"`
	Plan                          []InsurancePlanPlan     `json:"plan,omitempty"`
}

func (r InsurancePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r InsurancePlan) marshalJSON() jsonInsurancePlan {
	m := jsonInsurancePlan{}
	m.ResourceType = "InsurancePlan"
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
	m.Identifier = r.Identifier
	m.Status = r.Status
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Type = r.Type
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Alias = r.Alias
	anyAliasIdOrExtension := false
	for _, e := range r.Alias {
		if e.Id != nil || e.Extension != nil {
			anyAliasIdOrExtension = true
			break
		}
	}
	if anyAliasIdOrExtension {
		m.AliasPrimitiveElement = make([]*primitiveElement, 0, len(r.Alias))
		for _, e := range r.Alias {
			if e.Id != nil || e.Extension != nil {
				m.AliasPrimitiveElement = append(m.AliasPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.AliasPrimitiveElement = append(m.AliasPrimitiveElement, nil)
			}
		}
	}
	m.Period = r.Period
	m.OwnedBy = r.OwnedBy
	m.AdministeredBy = r.AdministeredBy
	m.CoverageArea = r.CoverageArea
	m.Contact = r.Contact
	m.Endpoint = r.Endpoint
	m.Network = r.Network
	m.Coverage = r.Coverage
	m.Plan = r.Plan
	return m
}
func (r *InsurancePlan) UnmarshalJSON(b []byte) error {
	var m jsonInsurancePlan
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *InsurancePlan) unmarshalJSON(m jsonInsurancePlan) error {
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
	r.Identifier = m.Identifier
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Alias = m.Alias
	for i, e := range m.AliasPrimitiveElement {
		if len(r.Alias) > i {
			r.Alias[i].Id = e.Id
			r.Alias[i].Extension = e.Extension
		} else {
			r.Alias = append(r.Alias, String{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Period = m.Period
	r.OwnedBy = m.OwnedBy
	r.AdministeredBy = m.AdministeredBy
	r.CoverageArea = m.CoverageArea
	r.Contact = m.Contact
	r.Endpoint = m.Endpoint
	r.Network = m.Network
	r.Coverage = m.Coverage
	r.Plan = m.Plan
	return nil
}
func (r InsurancePlan) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The contact for the health insurance product for a certain purpose.
type InsurancePlanContact struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates a purpose for which the contact can be reached.
	Purpose *CodeableConcept
	// A name associated with the contact.
	Name *HumanName
	// A contact detail (e.g. a telephone number or an email address) by which the party may be contacted.
	Telecom []ContactPoint
	// Visiting or postal addresses for the contact.
	Address *Address
}
type jsonInsurancePlanContact struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Purpose           *CodeableConcept `json:"purpose,omitempty"`
	Name              *HumanName       `json:"name,omitempty"`
	Telecom           []ContactPoint   `json:"telecom,omitempty"`
	Address           *Address         `json:"address,omitempty"`
}

func (r InsurancePlanContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r InsurancePlanContact) marshalJSON() jsonInsurancePlanContact {
	m := jsonInsurancePlanContact{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Purpose = r.Purpose
	m.Name = r.Name
	m.Telecom = r.Telecom
	m.Address = r.Address
	return m
}
func (r *InsurancePlanContact) UnmarshalJSON(b []byte) error {
	var m jsonInsurancePlanContact
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *InsurancePlanContact) unmarshalJSON(m jsonInsurancePlanContact) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Purpose = m.Purpose
	r.Name = m.Name
	r.Telecom = m.Telecom
	r.Address = m.Address
	return nil
}
func (r InsurancePlanContact) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Details about the coverage offered by the insurance product.
type InsurancePlanCoverage struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Type of coverage  (Medical; Dental; Mental Health; Substance Abuse; Vision; Drug; Short Term; Long Term Care; Hospice; Home Health).
	Type CodeableConcept
	// Reference to the network that providing the type of coverage.
	Network []Reference
	// Specific benefits under this type of coverage.
	Benefit []InsurancePlanCoverageBenefit
}
type jsonInsurancePlanCoverage struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Type              CodeableConcept                `json:"type,omitempty"`
	Network           []Reference                    `json:"network,omitempty"`
	Benefit           []InsurancePlanCoverageBenefit `json:"benefit,omitempty"`
}

func (r InsurancePlanCoverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r InsurancePlanCoverage) marshalJSON() jsonInsurancePlanCoverage {
	m := jsonInsurancePlanCoverage{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Network = r.Network
	m.Benefit = r.Benefit
	return m
}
func (r *InsurancePlanCoverage) UnmarshalJSON(b []byte) error {
	var m jsonInsurancePlanCoverage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *InsurancePlanCoverage) unmarshalJSON(m jsonInsurancePlanCoverage) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Network = m.Network
	r.Benefit = m.Benefit
	return nil
}
func (r InsurancePlanCoverage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Specific benefits under this type of coverage.
type InsurancePlanCoverageBenefit struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Type of benefit (primary care; speciality care; inpatient; outpatient).
	Type CodeableConcept
	// The referral requirements to have access/coverage for this benefit.
	Requirement *String
	// The specific limits on the benefit.
	Limit []InsurancePlanCoverageBenefitLimit
}
type jsonInsurancePlanCoverageBenefit struct {
	Id                          *string                             `json:"id,omitempty"`
	Extension                   []Extension                         `json:"extension,omitempty"`
	ModifierExtension           []Extension                         `json:"modifierExtension,omitempty"`
	Type                        CodeableConcept                     `json:"type,omitempty"`
	Requirement                 *String                             `json:"requirement,omitempty"`
	RequirementPrimitiveElement *primitiveElement                   `json:"_requirement,omitempty"`
	Limit                       []InsurancePlanCoverageBenefitLimit `json:"limit,omitempty"`
}

func (r InsurancePlanCoverageBenefit) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r InsurancePlanCoverageBenefit) marshalJSON() jsonInsurancePlanCoverageBenefit {
	m := jsonInsurancePlanCoverageBenefit{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Requirement = r.Requirement
	if r.Requirement != nil && (r.Requirement.Id != nil || r.Requirement.Extension != nil) {
		m.RequirementPrimitiveElement = &primitiveElement{Id: r.Requirement.Id, Extension: r.Requirement.Extension}
	}
	m.Limit = r.Limit
	return m
}
func (r *InsurancePlanCoverageBenefit) UnmarshalJSON(b []byte) error {
	var m jsonInsurancePlanCoverageBenefit
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *InsurancePlanCoverageBenefit) unmarshalJSON(m jsonInsurancePlanCoverageBenefit) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Requirement = m.Requirement
	if m.RequirementPrimitiveElement != nil {
		r.Requirement.Id = m.RequirementPrimitiveElement.Id
		r.Requirement.Extension = m.RequirementPrimitiveElement.Extension
	}
	r.Limit = m.Limit
	return nil
}
func (r InsurancePlanCoverageBenefit) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The specific limits on the benefit.
type InsurancePlanCoverageBenefitLimit struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The maximum amount of a service item a plan will pay for a covered benefit.  For examples. wellness visits, or eyeglasses.
	Value *Quantity
	// The specific limit on the benefit.
	Code *CodeableConcept
}
type jsonInsurancePlanCoverageBenefitLimit struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Value             *Quantity        `json:"value,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
}

func (r InsurancePlanCoverageBenefitLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r InsurancePlanCoverageBenefitLimit) marshalJSON() jsonInsurancePlanCoverageBenefitLimit {
	m := jsonInsurancePlanCoverageBenefitLimit{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Value = r.Value
	m.Code = r.Code
	return m
}
func (r *InsurancePlanCoverageBenefitLimit) UnmarshalJSON(b []byte) error {
	var m jsonInsurancePlanCoverageBenefitLimit
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *InsurancePlanCoverageBenefitLimit) unmarshalJSON(m jsonInsurancePlanCoverageBenefitLimit) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Value = m.Value
	r.Code = m.Code
	return nil
}
func (r InsurancePlanCoverageBenefitLimit) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Details about an insurance plan.
type InsurancePlanPlan struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Business identifiers assigned to this health insurance plan which remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// Type of plan. For example, "Platinum" or "High Deductable".
	Type *CodeableConcept
	// The geographic region in which a health insurance plan's benefits apply.
	CoverageArea []Reference
	// Reference to the network that providing the type of coverage.
	Network []Reference
	// Overall costs associated with the plan.
	GeneralCost []InsurancePlanPlanGeneralCost
	// Costs associated with the coverage provided by the product.
	SpecificCost []InsurancePlanPlanSpecificCost
}
type jsonInsurancePlanPlan struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                    `json:"identifier,omitempty"`
	Type              *CodeableConcept                `json:"type,omitempty"`
	CoverageArea      []Reference                     `json:"coverageArea,omitempty"`
	Network           []Reference                     `json:"network,omitempty"`
	GeneralCost       []InsurancePlanPlanGeneralCost  `json:"generalCost,omitempty"`
	SpecificCost      []InsurancePlanPlanSpecificCost `json:"specificCost,omitempty"`
}

func (r InsurancePlanPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r InsurancePlanPlan) marshalJSON() jsonInsurancePlanPlan {
	m := jsonInsurancePlanPlan{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Type = r.Type
	m.CoverageArea = r.CoverageArea
	m.Network = r.Network
	m.GeneralCost = r.GeneralCost
	m.SpecificCost = r.SpecificCost
	return m
}
func (r *InsurancePlanPlan) UnmarshalJSON(b []byte) error {
	var m jsonInsurancePlanPlan
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *InsurancePlanPlan) unmarshalJSON(m jsonInsurancePlanPlan) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Type = m.Type
	r.CoverageArea = m.CoverageArea
	r.Network = m.Network
	r.GeneralCost = m.GeneralCost
	r.SpecificCost = m.SpecificCost
	return nil
}
func (r InsurancePlanPlan) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Overall costs associated with the plan.
type InsurancePlanPlanGeneralCost struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Type of cost.
	Type *CodeableConcept
	// Number of participants enrolled in the plan.
	GroupSize *PositiveInt
	// Value of the cost.
	Cost *Money
	// Additional information about the general costs associated with this plan.
	Comment *String
}
type jsonInsurancePlanPlanGeneralCost struct {
	Id                        *string           `json:"id,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Type                      *CodeableConcept  `json:"type,omitempty"`
	GroupSize                 *PositiveInt      `json:"groupSize,omitempty"`
	GroupSizePrimitiveElement *primitiveElement `json:"_groupSize,omitempty"`
	Cost                      *Money            `json:"cost,omitempty"`
	Comment                   *String           `json:"comment,omitempty"`
	CommentPrimitiveElement   *primitiveElement `json:"_comment,omitempty"`
}

func (r InsurancePlanPlanGeneralCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r InsurancePlanPlanGeneralCost) marshalJSON() jsonInsurancePlanPlanGeneralCost {
	m := jsonInsurancePlanPlanGeneralCost{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.GroupSize = r.GroupSize
	if r.GroupSize != nil && (r.GroupSize.Id != nil || r.GroupSize.Extension != nil) {
		m.GroupSizePrimitiveElement = &primitiveElement{Id: r.GroupSize.Id, Extension: r.GroupSize.Extension}
	}
	m.Cost = r.Cost
	m.Comment = r.Comment
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	return m
}
func (r *InsurancePlanPlanGeneralCost) UnmarshalJSON(b []byte) error {
	var m jsonInsurancePlanPlanGeneralCost
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *InsurancePlanPlanGeneralCost) unmarshalJSON(m jsonInsurancePlanPlanGeneralCost) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.GroupSize = m.GroupSize
	if m.GroupSizePrimitiveElement != nil {
		r.GroupSize.Id = m.GroupSizePrimitiveElement.Id
		r.GroupSize.Extension = m.GroupSizePrimitiveElement.Extension
	}
	r.Cost = m.Cost
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	return nil
}
func (r InsurancePlanPlanGeneralCost) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Costs associated with the coverage provided by the product.
type InsurancePlanPlanSpecificCost struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// General category of benefit (Medical; Dental; Vision; Drug; Mental Health; Substance Abuse; Hospice, Home Health).
	Category CodeableConcept
	// List of the specific benefits under this category of benefit.
	Benefit []InsurancePlanPlanSpecificCostBenefit
}
type jsonInsurancePlanPlanSpecificCost struct {
	Id                *string                                `json:"id,omitempty"`
	Extension         []Extension                            `json:"extension,omitempty"`
	ModifierExtension []Extension                            `json:"modifierExtension,omitempty"`
	Category          CodeableConcept                        `json:"category,omitempty"`
	Benefit           []InsurancePlanPlanSpecificCostBenefit `json:"benefit,omitempty"`
}

func (r InsurancePlanPlanSpecificCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r InsurancePlanPlanSpecificCost) marshalJSON() jsonInsurancePlanPlanSpecificCost {
	m := jsonInsurancePlanPlanSpecificCost{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Category = r.Category
	m.Benefit = r.Benefit
	return m
}
func (r *InsurancePlanPlanSpecificCost) UnmarshalJSON(b []byte) error {
	var m jsonInsurancePlanPlanSpecificCost
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *InsurancePlanPlanSpecificCost) unmarshalJSON(m jsonInsurancePlanPlanSpecificCost) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Category = m.Category
	r.Benefit = m.Benefit
	return nil
}
func (r InsurancePlanPlanSpecificCost) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// List of the specific benefits under this category of benefit.
type InsurancePlanPlanSpecificCostBenefit struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Type of specific benefit (preventative; primary care office visit; speciality office visit; hospitalization; emergency room; urgent care).
	Type CodeableConcept
	// List of the costs associated with a specific benefit.
	Cost []InsurancePlanPlanSpecificCostBenefitCost
}
type jsonInsurancePlanPlanSpecificCostBenefit struct {
	Id                *string                                    `json:"id,omitempty"`
	Extension         []Extension                                `json:"extension,omitempty"`
	ModifierExtension []Extension                                `json:"modifierExtension,omitempty"`
	Type              CodeableConcept                            `json:"type,omitempty"`
	Cost              []InsurancePlanPlanSpecificCostBenefitCost `json:"cost,omitempty"`
}

func (r InsurancePlanPlanSpecificCostBenefit) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r InsurancePlanPlanSpecificCostBenefit) marshalJSON() jsonInsurancePlanPlanSpecificCostBenefit {
	m := jsonInsurancePlanPlanSpecificCostBenefit{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Cost = r.Cost
	return m
}
func (r *InsurancePlanPlanSpecificCostBenefit) UnmarshalJSON(b []byte) error {
	var m jsonInsurancePlanPlanSpecificCostBenefit
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *InsurancePlanPlanSpecificCostBenefit) unmarshalJSON(m jsonInsurancePlanPlanSpecificCostBenefit) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Cost = m.Cost
	return nil
}
func (r InsurancePlanPlanSpecificCostBenefit) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// List of the costs associated with a specific benefit.
type InsurancePlanPlanSpecificCostBenefitCost struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Type of cost (copay; individual cap; family cap; coinsurance; deductible).
	Type CodeableConcept
	// Whether the cost applies to in-network or out-of-network providers (in-network; out-of-network; other).
	Applicability *CodeableConcept
	// Additional information about the cost, such as information about funding sources (e.g. HSA, HRA, FSA, RRA).
	Qualifiers []CodeableConcept
	// The actual cost value. (some of the costs may be represented as percentages rather than currency, e.g. 10% coinsurance).
	Value *Quantity
}
type jsonInsurancePlanPlanSpecificCostBenefitCost struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `json:"type,omitempty"`
	Applicability     *CodeableConcept  `json:"applicability,omitempty"`
	Qualifiers        []CodeableConcept `json:"qualifiers,omitempty"`
	Value             *Quantity         `json:"value,omitempty"`
}

func (r InsurancePlanPlanSpecificCostBenefitCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r InsurancePlanPlanSpecificCostBenefitCost) marshalJSON() jsonInsurancePlanPlanSpecificCostBenefitCost {
	m := jsonInsurancePlanPlanSpecificCostBenefitCost{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Applicability = r.Applicability
	m.Qualifiers = r.Qualifiers
	m.Value = r.Value
	return m
}
func (r *InsurancePlanPlanSpecificCostBenefitCost) UnmarshalJSON(b []byte) error {
	var m jsonInsurancePlanPlanSpecificCostBenefitCost
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *InsurancePlanPlanSpecificCostBenefitCost) unmarshalJSON(m jsonInsurancePlanPlanSpecificCostBenefitCost) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Applicability = m.Applicability
	r.Qualifiers = m.Qualifiers
	r.Value = m.Value
	return nil
}
func (r InsurancePlanPlanSpecificCostBenefitCost) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
