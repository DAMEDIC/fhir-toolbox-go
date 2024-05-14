package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// Financial instrument which may be used to reimburse or pay for health care products and services. Includes both insurance and self-payment.
//
// Coverage provides a link between covered parties (patients) and the payors of their healthcare costs (both insurance and self-pay).
type Coverage struct {
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
	// A unique identifier assigned to this coverage.
	Identifier []Identifier
	// The status of the resource instance.
	Status Code
	// The type of coverage: social program, medical plan, accident coverage (workers compensation, auto), group health or payment by an individual or organization.
	Type *CodeableConcept
	// The party who 'owns' the insurance policy.
	PolicyHolder *Reference
	// The party who has signed-up for or 'owns' the contractual relationship to the policy or to whom the benefit of the policy for services rendered to them or their family is due.
	Subscriber *Reference
	// The insurer assigned ID for the Subscriber.
	SubscriberId *String
	// The party who benefits from the insurance coverage; the patient when products and/or services are provided.
	Beneficiary Reference
	// A unique identifier for a dependent under the coverage.
	Dependent *String
	// The relationship of beneficiary (patient) to the subscriber.
	Relationship *CodeableConcept
	// Time period during which the coverage is in force. A missing start date indicates the start date isn't known, a missing end date means the coverage is continuing to be in force.
	Period *Period
	// The program or plan underwriter or payor including both insurance and non-insurance agreements, such as patient-pay agreements.
	Payor []Reference
	// A suite of underwriter specific classifiers.
	Class []CoverageClass
	// The order of applicability of this coverage relative to other coverages which are currently in force. Note, there may be gaps in the numbering and this does not imply primary, secondary etc. as the specific positioning of coverages depends upon the episode of care.
	Order *PositiveInt
	// The insurer-specific identifier for the insurer-defined network of providers to which the beneficiary may seek treatment which will be covered at the 'in-network' rate, otherwise 'out of network' terms and conditions apply.
	Network *String
	// A suite of codes indicating the cost category and associated amount which have been detailed in the policy and may have been  included on the health card.
	CostToBeneficiary []CoverageCostToBeneficiary
	// When 'subrogation=true' this insurance instance has been included not for adjudication but to provide insurers with the details to recover costs.
	Subrogation *Boolean
	// The policy(s) which constitute this insurance coverage.
	Contract []Reference
}

func (r Coverage) ResourceType() string {
	return "Coverage"
}

type jsonCoverage struct {
	ResourceType                  string                      `json:"resourceType"`
	Id                            *Id                         `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement           `json:"_id,omitempty"`
	Meta                          *Meta                       `json:"meta,omitempty"`
	ImplicitRules                 *Uri                        `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement           `json:"_implicitRules,omitempty"`
	Language                      *Code                       `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement           `json:"_language,omitempty"`
	Text                          *Narrative                  `json:"text,omitempty"`
	Contained                     []containedResource         `json:"contained,omitempty"`
	Extension                     []Extension                 `json:"extension,omitempty"`
	ModifierExtension             []Extension                 `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier                `json:"identifier,omitempty"`
	Status                        Code                        `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement           `json:"_status,omitempty"`
	Type                          *CodeableConcept            `json:"type,omitempty"`
	PolicyHolder                  *Reference                  `json:"policyHolder,omitempty"`
	Subscriber                    *Reference                  `json:"subscriber,omitempty"`
	SubscriberId                  *String                     `json:"subscriberId,omitempty"`
	SubscriberIdPrimitiveElement  *primitiveElement           `json:"_subscriberId,omitempty"`
	Beneficiary                   Reference                   `json:"beneficiary,omitempty"`
	Dependent                     *String                     `json:"dependent,omitempty"`
	DependentPrimitiveElement     *primitiveElement           `json:"_dependent,omitempty"`
	Relationship                  *CodeableConcept            `json:"relationship,omitempty"`
	Period                        *Period                     `json:"period,omitempty"`
	Payor                         []Reference                 `json:"payor,omitempty"`
	Class                         []CoverageClass             `json:"class,omitempty"`
	Order                         *PositiveInt                `json:"order,omitempty"`
	OrderPrimitiveElement         *primitiveElement           `json:"_order,omitempty"`
	Network                       *String                     `json:"network,omitempty"`
	NetworkPrimitiveElement       *primitiveElement           `json:"_network,omitempty"`
	CostToBeneficiary             []CoverageCostToBeneficiary `json:"costToBeneficiary,omitempty"`
	Subrogation                   *Boolean                    `json:"subrogation,omitempty"`
	SubrogationPrimitiveElement   *primitiveElement           `json:"_subrogation,omitempty"`
	Contract                      []Reference                 `json:"contract,omitempty"`
}

func (r Coverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Coverage) marshalJSON() jsonCoverage {
	m := jsonCoverage{}
	m.ResourceType = "Coverage"
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
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Type = r.Type
	m.PolicyHolder = r.PolicyHolder
	m.Subscriber = r.Subscriber
	m.SubscriberId = r.SubscriberId
	if r.SubscriberId != nil && (r.SubscriberId.Id != nil || r.SubscriberId.Extension != nil) {
		m.SubscriberIdPrimitiveElement = &primitiveElement{Id: r.SubscriberId.Id, Extension: r.SubscriberId.Extension}
	}
	m.Beneficiary = r.Beneficiary
	m.Dependent = r.Dependent
	if r.Dependent != nil && (r.Dependent.Id != nil || r.Dependent.Extension != nil) {
		m.DependentPrimitiveElement = &primitiveElement{Id: r.Dependent.Id, Extension: r.Dependent.Extension}
	}
	m.Relationship = r.Relationship
	m.Period = r.Period
	m.Payor = r.Payor
	m.Class = r.Class
	m.Order = r.Order
	if r.Order != nil && (r.Order.Id != nil || r.Order.Extension != nil) {
		m.OrderPrimitiveElement = &primitiveElement{Id: r.Order.Id, Extension: r.Order.Extension}
	}
	m.Network = r.Network
	if r.Network != nil && (r.Network.Id != nil || r.Network.Extension != nil) {
		m.NetworkPrimitiveElement = &primitiveElement{Id: r.Network.Id, Extension: r.Network.Extension}
	}
	m.CostToBeneficiary = r.CostToBeneficiary
	m.Subrogation = r.Subrogation
	if r.Subrogation != nil && (r.Subrogation.Id != nil || r.Subrogation.Extension != nil) {
		m.SubrogationPrimitiveElement = &primitiveElement{Id: r.Subrogation.Id, Extension: r.Subrogation.Extension}
	}
	m.Contract = r.Contract
	return m
}
func (r *Coverage) UnmarshalJSON(b []byte) error {
	var m jsonCoverage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Coverage) unmarshalJSON(m jsonCoverage) error {
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
	r.PolicyHolder = m.PolicyHolder
	r.Subscriber = m.Subscriber
	r.SubscriberId = m.SubscriberId
	if m.SubscriberIdPrimitiveElement != nil {
		r.SubscriberId.Id = m.SubscriberIdPrimitiveElement.Id
		r.SubscriberId.Extension = m.SubscriberIdPrimitiveElement.Extension
	}
	r.Beneficiary = m.Beneficiary
	r.Dependent = m.Dependent
	if m.DependentPrimitiveElement != nil {
		r.Dependent.Id = m.DependentPrimitiveElement.Id
		r.Dependent.Extension = m.DependentPrimitiveElement.Extension
	}
	r.Relationship = m.Relationship
	r.Period = m.Period
	r.Payor = m.Payor
	r.Class = m.Class
	r.Order = m.Order
	if m.OrderPrimitiveElement != nil {
		r.Order.Id = m.OrderPrimitiveElement.Id
		r.Order.Extension = m.OrderPrimitiveElement.Extension
	}
	r.Network = m.Network
	if m.NetworkPrimitiveElement != nil {
		r.Network.Id = m.NetworkPrimitiveElement.Id
		r.Network.Extension = m.NetworkPrimitiveElement.Extension
	}
	r.CostToBeneficiary = m.CostToBeneficiary
	r.Subrogation = m.Subrogation
	if m.SubrogationPrimitiveElement != nil {
		r.Subrogation.Id = m.SubrogationPrimitiveElement.Id
		r.Subrogation.Extension = m.SubrogationPrimitiveElement.Extension
	}
	r.Contract = m.Contract
	return nil
}
func (r Coverage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A suite of underwriter specific classifiers.
type CoverageClass struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of classification for which an insurer-specific class label or number and optional name is provided, for example may be used to identify a class of coverage or employer group, Policy, Plan.
	Type CodeableConcept
	// The alphanumeric string value associated with the insurer issued label.
	Value String
	// A short description for the class.
	Name *String
}
type jsonCoverageClass struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Type                  CodeableConcept   `json:"type,omitempty"`
	Value                 String            `json:"value,omitempty"`
	ValuePrimitiveElement *primitiveElement `json:"_value,omitempty"`
	Name                  *String           `json:"name,omitempty"`
	NamePrimitiveElement  *primitiveElement `json:"_name,omitempty"`
}

func (r CoverageClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CoverageClass) marshalJSON() jsonCoverageClass {
	m := jsonCoverageClass{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Value = r.Value
	if r.Value.Id != nil || r.Value.Extension != nil {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	return m
}
func (r *CoverageClass) UnmarshalJSON(b []byte) error {
	var m jsonCoverageClass
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CoverageClass) unmarshalJSON(m jsonCoverageClass) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	return nil
}
func (r CoverageClass) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A suite of codes indicating the cost category and associated amount which have been detailed in the policy and may have been  included on the health card.
type CoverageCostToBeneficiary struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The category of patient centric costs associated with treatment.
	Type *CodeableConcept
	// The amount due from the patient for the cost category.
	Value isCoverageCostToBeneficiaryValue
	// A suite of codes indicating exceptions or reductions to patient costs and their effective periods.
	Exception []CoverageCostToBeneficiaryException
}
type isCoverageCostToBeneficiaryValue interface {
	isCoverageCostToBeneficiaryValue()
}

func (r Quantity) isCoverageCostToBeneficiaryValue() {}
func (r Money) isCoverageCostToBeneficiaryValue()    {}

type jsonCoverageCostToBeneficiary struct {
	Id                *string                              `json:"id,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept                     `json:"type,omitempty"`
	ValueQuantity     *Quantity                            `json:"valueQuantity,omitempty"`
	ValueMoney        *Money                               `json:"valueMoney,omitempty"`
	Exception         []CoverageCostToBeneficiaryException `json:"exception,omitempty"`
}

func (r CoverageCostToBeneficiary) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CoverageCostToBeneficiary) marshalJSON() jsonCoverageCostToBeneficiary {
	m := jsonCoverageCostToBeneficiary{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	switch v := r.Value.(type) {
	case Quantity:
		m.ValueQuantity = &v
	case *Quantity:
		m.ValueQuantity = v
	case Money:
		m.ValueMoney = &v
	case *Money:
		m.ValueMoney = v
	}
	m.Exception = r.Exception
	return m
}
func (r *CoverageCostToBeneficiary) UnmarshalJSON(b []byte) error {
	var m jsonCoverageCostToBeneficiary
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CoverageCostToBeneficiary) unmarshalJSON(m jsonCoverageCostToBeneficiary) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.ValueQuantity != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueQuantity
		r.Value = v
	}
	if m.ValueMoney != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueMoney
		r.Value = v
	}
	r.Exception = m.Exception
	return nil
}
func (r CoverageCostToBeneficiary) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A suite of codes indicating exceptions or reductions to patient costs and their effective periods.
type CoverageCostToBeneficiaryException struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The code for the specific exception.
	Type CodeableConcept
	// The timeframe during when the exception is in force.
	Period *Period
}
type jsonCoverageCostToBeneficiaryException struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type,omitempty"`
	Period            *Period         `json:"period,omitempty"`
}

func (r CoverageCostToBeneficiaryException) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CoverageCostToBeneficiaryException) marshalJSON() jsonCoverageCostToBeneficiaryException {
	m := jsonCoverageCostToBeneficiaryException{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Period = r.Period
	return m
}
func (r *CoverageCostToBeneficiaryException) UnmarshalJSON(b []byte) error {
	var m jsonCoverageCostToBeneficiaryException
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CoverageCostToBeneficiaryException) unmarshalJSON(m jsonCoverageCostToBeneficiaryException) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Period = m.Period
	return nil
}
func (r CoverageCostToBeneficiaryException) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
