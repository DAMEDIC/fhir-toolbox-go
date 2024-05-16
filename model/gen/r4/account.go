package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A financial tool for tracking value accrued for a particular purpose.  In the healthcare field, used to track charges for a patient, cost centers, etc.
type Account struct {
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
	// Unique identifier used to reference the account.  Might or might not be intended for human use (e.g. credit card number).
	Identifier []Identifier
	// Indicates whether the account is presently used/usable or not.
	Status Code
	// Categorizes the account for reporting and searching purposes.
	Type *CodeableConcept
	// Name used for the account when displaying it to humans in reports, etc.
	Name *String
	// Identifies the entity which incurs the expenses. While the immediate recipients of services or goods might be entities related to the subject, the expenses were ultimately incurred by the subject of the Account.
	Subject []Reference
	// The date range of services associated with this account.
	ServicePeriod *Period
	// The party(s) that are responsible for covering the payment of this account, and what order should they be applied to the account.
	Coverage []AccountCoverage
	// Indicates the service area, hospital, department, etc. with responsibility for managing the Account.
	Owner *Reference
	// Provides additional information about what the account tracks and how it is used.
	Description *String
	// The parties responsible for balancing the account if other payment options fall short.
	Guarantor []AccountGuarantor
	// Reference to a parent Account.
	PartOf *Reference
}

func (r Account) ResourceType() string {
	return "Account"
}
func (r Account) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonAccount struct {
	ResourceType                  string              `json:"resourceType"`
	Id                            *Id                 `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement   `json:"_id,omitempty"`
	Meta                          *Meta               `json:"meta,omitempty"`
	ImplicitRules                 *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                      *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement   `json:"_language,omitempty"`
	Text                          *Narrative          `json:"text,omitempty"`
	Contained                     []containedResource `json:"contained,omitempty"`
	Extension                     []Extension         `json:"extension,omitempty"`
	ModifierExtension             []Extension         `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier        `json:"identifier,omitempty"`
	Status                        Code                `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement   `json:"_status,omitempty"`
	Type                          *CodeableConcept    `json:"type,omitempty"`
	Name                          *String             `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement   `json:"_name,omitempty"`
	Subject                       []Reference         `json:"subject,omitempty"`
	ServicePeriod                 *Period             `json:"servicePeriod,omitempty"`
	Coverage                      []AccountCoverage   `json:"coverage,omitempty"`
	Owner                         *Reference          `json:"owner,omitempty"`
	Description                   *String             `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement   `json:"_description,omitempty"`
	Guarantor                     []AccountGuarantor  `json:"guarantor,omitempty"`
	PartOf                        *Reference          `json:"partOf,omitempty"`
}

func (r Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Account) marshalJSON() jsonAccount {
	m := jsonAccount{}
	m.ResourceType = "Account"
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
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Subject = r.Subject
	m.ServicePeriod = r.ServicePeriod
	m.Coverage = r.Coverage
	m.Owner = r.Owner
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Guarantor = r.Guarantor
	m.PartOf = r.PartOf
	return m
}
func (r *Account) UnmarshalJSON(b []byte) error {
	var m jsonAccount
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Account) unmarshalJSON(m jsonAccount) error {
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
	r.Subject = m.Subject
	r.ServicePeriod = m.ServicePeriod
	r.Coverage = m.Coverage
	r.Owner = m.Owner
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Guarantor = m.Guarantor
	r.PartOf = m.PartOf
	return nil
}
func (r Account) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The party(s) that are responsible for covering the payment of this account, and what order should they be applied to the account.
type AccountCoverage struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The party(s) that contribute to payment (or part of) of the charges applied to this account (including self-pay).
	//
	// A coverage may only be responsible for specific types of charges, and the sequence of the coverages in the account could be important when processing billing.
	Coverage Reference
	// The priority of the coverage in the context of this account.
	Priority *PositiveInt
}
type jsonAccountCoverage struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Coverage                 Reference         `json:"coverage,omitempty"`
	Priority                 *PositiveInt      `json:"priority,omitempty"`
	PriorityPrimitiveElement *primitiveElement `json:"_priority,omitempty"`
}

func (r AccountCoverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AccountCoverage) marshalJSON() jsonAccountCoverage {
	m := jsonAccountCoverage{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Coverage = r.Coverage
	m.Priority = r.Priority
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		m.PriorityPrimitiveElement = &primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
	}
	return m
}
func (r *AccountCoverage) UnmarshalJSON(b []byte) error {
	var m jsonAccountCoverage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AccountCoverage) unmarshalJSON(m jsonAccountCoverage) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Coverage = m.Coverage
	r.Priority = m.Priority
	if m.PriorityPrimitiveElement != nil {
		r.Priority.Id = m.PriorityPrimitiveElement.Id
		r.Priority.Extension = m.PriorityPrimitiveElement.Extension
	}
	return nil
}
func (r AccountCoverage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The parties responsible for balancing the account if other payment options fall short.
type AccountGuarantor struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The entity who is responsible.
	Party Reference
	// A guarantor may be placed on credit hold or otherwise have their role temporarily suspended.
	OnHold *Boolean
	// The timeframe during which the guarantor accepts responsibility for the account.
	Period *Period
}
type jsonAccountGuarantor struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Party                  Reference         `json:"party,omitempty"`
	OnHold                 *Boolean          `json:"onHold,omitempty"`
	OnHoldPrimitiveElement *primitiveElement `json:"_onHold,omitempty"`
	Period                 *Period           `json:"period,omitempty"`
}

func (r AccountGuarantor) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r AccountGuarantor) marshalJSON() jsonAccountGuarantor {
	m := jsonAccountGuarantor{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Party = r.Party
	m.OnHold = r.OnHold
	if r.OnHold != nil && (r.OnHold.Id != nil || r.OnHold.Extension != nil) {
		m.OnHoldPrimitiveElement = &primitiveElement{Id: r.OnHold.Id, Extension: r.OnHold.Extension}
	}
	m.Period = r.Period
	return m
}
func (r *AccountGuarantor) UnmarshalJSON(b []byte) error {
	var m jsonAccountGuarantor
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *AccountGuarantor) unmarshalJSON(m jsonAccountGuarantor) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Party = m.Party
	r.OnHold = m.OnHold
	if m.OnHoldPrimitiveElement != nil {
		r.OnHold.Id = m.OnHoldPrimitiveElement.Id
		r.OnHold.Extension = m.OnHoldPrimitiveElement.Extension
	}
	r.Period = m.Period
	return nil
}
func (r AccountGuarantor) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
