package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A provider issued list of professional services and products which have been provided, or are to be provided, to a patient which is sent to an insurer for reimbursement.
//
// The Claim resource is used by providers to exchange services and products rendered to patients or planned to be rendered with insurers for reimbuserment. It is also used by insurers to exchange claims information with statutory reporting and data analytics firms.
type Claim struct {
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
	// A unique identifier assigned to this claim.
	Identifier []Identifier
	// The status of the resource instance.
	Status Code
	// The category of claim, e.g. oral, pharmacy, vision, institutional, professional.
	Type CodeableConcept
	// A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
	SubType *CodeableConcept
	// A code to indicate whether the nature of the request is: to request adjudication of products and services previously rendered; or requesting authorization and adjudication for provision in the future; or requesting the non-binding adjudication of the listed products and services which could be provided in the future.
	Use Code
	// The party to whom the professional services and/or products have been supplied or are being considered and for whom actual or forecast reimbursement is sought.
	Patient Reference
	// The period for which charges are being submitted.
	BillablePeriod *Period
	// The date this resource was created.
	Created DateTime
	// Individual who created the claim, predetermination or preauthorization.
	Enterer *Reference
	// The Insurer who is target of the request.
	Insurer *Reference
	// The provider which is responsible for the claim, predetermination or preauthorization.
	Provider Reference
	// The provider-required urgency of processing the request. Typical values include: stat, routine deferred.
	Priority CodeableConcept
	// A code to indicate whether and for whom funds are to be reserved for future claims.
	FundsReserve *CodeableConcept
	// Other claims which are related to this claim such as prior submissions or claims for related services or for the same event.
	Related []ClaimRelated
	// Prescription to support the dispensing of pharmacy, device or vision products.
	Prescription *Reference
	// Original prescription which has been superseded by this prescription to support the dispensing of pharmacy services, medications or products.
	OriginalPrescription *Reference
	// The party to be reimbursed for cost of the products and services according to the terms of the policy.
	Payee *ClaimPayee
	// A reference to a referral resource.
	Referral *Reference
	// Facility where the services were provided.
	Facility *Reference
	// The members of the team who provided the products and services.
	CareTeam []ClaimCareTeam
	// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
	SupportingInfo []ClaimSupportingInfo
	// Information about diagnoses relevant to the claim items.
	Diagnosis []ClaimDiagnosis
	// Procedures performed on the patient relevant to the billing items with the claim.
	Procedure []ClaimProcedure
	// Financial instruments for reimbursement for the health care products and services specified on the claim.
	Insurance []ClaimInsurance
	// Details of an accident which resulted in injuries which required the products and services listed in the claim.
	Accident *ClaimAccident
	// A claim line. Either a simple  product or service or a 'group' of details which can each be a simple items or groups of sub-details.
	Item []ClaimItem
	// The total value of the all the items in the claim.
	Total *Money
}

func (r Claim) ResourceType() string {
	return "Claim"
}
func (r Claim) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonClaim struct {
	ResourceType                  string                `json:"resourceType"`
	Id                            *Id                   `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement     `json:"_id,omitempty"`
	Meta                          *Meta                 `json:"meta,omitempty"`
	ImplicitRules                 *Uri                  `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement     `json:"_implicitRules,omitempty"`
	Language                      *Code                 `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement     `json:"_language,omitempty"`
	Text                          *Narrative            `json:"text,omitempty"`
	Contained                     []containedResource   `json:"contained,omitempty"`
	Extension                     []Extension           `json:"extension,omitempty"`
	ModifierExtension             []Extension           `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier          `json:"identifier,omitempty"`
	Status                        Code                  `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement     `json:"_status,omitempty"`
	Type                          CodeableConcept       `json:"type,omitempty"`
	SubType                       *CodeableConcept      `json:"subType,omitempty"`
	Use                           Code                  `json:"use,omitempty"`
	UsePrimitiveElement           *primitiveElement     `json:"_use,omitempty"`
	Patient                       Reference             `json:"patient,omitempty"`
	BillablePeriod                *Period               `json:"billablePeriod,omitempty"`
	Created                       DateTime              `json:"created,omitempty"`
	CreatedPrimitiveElement       *primitiveElement     `json:"_created,omitempty"`
	Enterer                       *Reference            `json:"enterer,omitempty"`
	Insurer                       *Reference            `json:"insurer,omitempty"`
	Provider                      Reference             `json:"provider,omitempty"`
	Priority                      CodeableConcept       `json:"priority,omitempty"`
	FundsReserve                  *CodeableConcept      `json:"fundsReserve,omitempty"`
	Related                       []ClaimRelated        `json:"related,omitempty"`
	Prescription                  *Reference            `json:"prescription,omitempty"`
	OriginalPrescription          *Reference            `json:"originalPrescription,omitempty"`
	Payee                         *ClaimPayee           `json:"payee,omitempty"`
	Referral                      *Reference            `json:"referral,omitempty"`
	Facility                      *Reference            `json:"facility,omitempty"`
	CareTeam                      []ClaimCareTeam       `json:"careTeam,omitempty"`
	SupportingInfo                []ClaimSupportingInfo `json:"supportingInfo,omitempty"`
	Diagnosis                     []ClaimDiagnosis      `json:"diagnosis,omitempty"`
	Procedure                     []ClaimProcedure      `json:"procedure,omitempty"`
	Insurance                     []ClaimInsurance      `json:"insurance,omitempty"`
	Accident                      *ClaimAccident        `json:"accident,omitempty"`
	Item                          []ClaimItem           `json:"item,omitempty"`
	Total                         *Money                `json:"total,omitempty"`
}

func (r Claim) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Claim) marshalJSON() jsonClaim {
	m := jsonClaim{}
	m.ResourceType = "Claim"
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
	m.SubType = r.SubType
	m.Use = r.Use
	if r.Use.Id != nil || r.Use.Extension != nil {
		m.UsePrimitiveElement = &primitiveElement{Id: r.Use.Id, Extension: r.Use.Extension}
	}
	m.Patient = r.Patient
	m.BillablePeriod = r.BillablePeriod
	m.Created = r.Created
	if r.Created.Id != nil || r.Created.Extension != nil {
		m.CreatedPrimitiveElement = &primitiveElement{Id: r.Created.Id, Extension: r.Created.Extension}
	}
	m.Enterer = r.Enterer
	m.Insurer = r.Insurer
	m.Provider = r.Provider
	m.Priority = r.Priority
	m.FundsReserve = r.FundsReserve
	m.Related = r.Related
	m.Prescription = r.Prescription
	m.OriginalPrescription = r.OriginalPrescription
	m.Payee = r.Payee
	m.Referral = r.Referral
	m.Facility = r.Facility
	m.CareTeam = r.CareTeam
	m.SupportingInfo = r.SupportingInfo
	m.Diagnosis = r.Diagnosis
	m.Procedure = r.Procedure
	m.Insurance = r.Insurance
	m.Accident = r.Accident
	m.Item = r.Item
	m.Total = r.Total
	return m
}
func (r *Claim) UnmarshalJSON(b []byte) error {
	var m jsonClaim
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Claim) unmarshalJSON(m jsonClaim) error {
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
	r.SubType = m.SubType
	r.Use = m.Use
	if m.UsePrimitiveElement != nil {
		r.Use.Id = m.UsePrimitiveElement.Id
		r.Use.Extension = m.UsePrimitiveElement.Extension
	}
	r.Patient = m.Patient
	r.BillablePeriod = m.BillablePeriod
	r.Created = m.Created
	if m.CreatedPrimitiveElement != nil {
		r.Created.Id = m.CreatedPrimitiveElement.Id
		r.Created.Extension = m.CreatedPrimitiveElement.Extension
	}
	r.Enterer = m.Enterer
	r.Insurer = m.Insurer
	r.Provider = m.Provider
	r.Priority = m.Priority
	r.FundsReserve = m.FundsReserve
	r.Related = m.Related
	r.Prescription = m.Prescription
	r.OriginalPrescription = m.OriginalPrescription
	r.Payee = m.Payee
	r.Referral = m.Referral
	r.Facility = m.Facility
	r.CareTeam = m.CareTeam
	r.SupportingInfo = m.SupportingInfo
	r.Diagnosis = m.Diagnosis
	r.Procedure = m.Procedure
	r.Insurance = m.Insurance
	r.Accident = m.Accident
	r.Item = m.Item
	r.Total = m.Total
	return nil
}
func (r Claim) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Other claims which are related to this claim such as prior submissions or claims for related services or for the same event.
type ClaimRelated struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Reference to a related claim.
	Claim *Reference
	// A code to convey how the claims are related.
	Relationship *CodeableConcept
	// An alternate organizational reference to the case or file to which this particular claim pertains.
	Reference *Identifier
}
type jsonClaimRelated struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Claim             *Reference       `json:"claim,omitempty"`
	Relationship      *CodeableConcept `json:"relationship,omitempty"`
	Reference         *Identifier      `json:"reference,omitempty"`
}

func (r ClaimRelated) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimRelated) marshalJSON() jsonClaimRelated {
	m := jsonClaimRelated{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Claim = r.Claim
	m.Relationship = r.Relationship
	m.Reference = r.Reference
	return m
}
func (r *ClaimRelated) UnmarshalJSON(b []byte) error {
	var m jsonClaimRelated
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimRelated) unmarshalJSON(m jsonClaimRelated) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Claim = m.Claim
	r.Relationship = m.Relationship
	r.Reference = m.Reference
	return nil
}
func (r ClaimRelated) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The party to be reimbursed for cost of the products and services according to the terms of the policy.
type ClaimPayee struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Type of Party to be reimbursed: subscriber, provider, other.
	Type CodeableConcept
	// Reference to the individual or organization to whom any payment will be made.
	Party *Reference
}
type jsonClaimPayee struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type,omitempty"`
	Party             *Reference      `json:"party,omitempty"`
}

func (r ClaimPayee) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimPayee) marshalJSON() jsonClaimPayee {
	m := jsonClaimPayee{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Party = r.Party
	return m
}
func (r *ClaimPayee) UnmarshalJSON(b []byte) error {
	var m jsonClaimPayee
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimPayee) unmarshalJSON(m jsonClaimPayee) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Party = m.Party
	return nil
}
func (r ClaimPayee) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The members of the team who provided the products and services.
type ClaimCareTeam struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely identify care team entries.
	Sequence PositiveInt
	// Member of the team who provided the product or service.
	Provider Reference
	// The party who is billing and/or responsible for the claimed products or services.
	Responsible *Boolean
	// The lead, assisting or supervising practitioner and their discipline if a multidisciplinary team.
	Role *CodeableConcept
	// The qualification of the practitioner which is applicable for this service.
	Qualification *CodeableConcept
}
type jsonClaimCareTeam struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Sequence                    PositiveInt       `json:"sequence,omitempty"`
	SequencePrimitiveElement    *primitiveElement `json:"_sequence,omitempty"`
	Provider                    Reference         `json:"provider,omitempty"`
	Responsible                 *Boolean          `json:"responsible,omitempty"`
	ResponsiblePrimitiveElement *primitiveElement `json:"_responsible,omitempty"`
	Role                        *CodeableConcept  `json:"role,omitempty"`
	Qualification               *CodeableConcept  `json:"qualification,omitempty"`
}

func (r ClaimCareTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimCareTeam) marshalJSON() jsonClaimCareTeam {
	m := jsonClaimCareTeam{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Sequence = r.Sequence
	if r.Sequence.Id != nil || r.Sequence.Extension != nil {
		m.SequencePrimitiveElement = &primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
	}
	m.Provider = r.Provider
	m.Responsible = r.Responsible
	if r.Responsible != nil && (r.Responsible.Id != nil || r.Responsible.Extension != nil) {
		m.ResponsiblePrimitiveElement = &primitiveElement{Id: r.Responsible.Id, Extension: r.Responsible.Extension}
	}
	m.Role = r.Role
	m.Qualification = r.Qualification
	return m
}
func (r *ClaimCareTeam) UnmarshalJSON(b []byte) error {
	var m jsonClaimCareTeam
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimCareTeam) unmarshalJSON(m jsonClaimCareTeam) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Sequence = m.Sequence
	if m.SequencePrimitiveElement != nil {
		r.Sequence.Id = m.SequencePrimitiveElement.Id
		r.Sequence.Extension = m.SequencePrimitiveElement.Extension
	}
	r.Provider = m.Provider
	r.Responsible = m.Responsible
	if m.ResponsiblePrimitiveElement != nil {
		r.Responsible.Id = m.ResponsiblePrimitiveElement.Id
		r.Responsible.Extension = m.ResponsiblePrimitiveElement.Extension
	}
	r.Role = m.Role
	r.Qualification = m.Qualification
	return nil
}
func (r ClaimCareTeam) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
type ClaimSupportingInfo struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely identify supporting information entries.
	Sequence PositiveInt
	// The general class of the information supplied: information; exception; accident, employment; onset, etc.
	Category CodeableConcept
	// System and code pertaining to the specific information regarding special conditions relating to the setting, treatment or patient  for which care is sought.
	Code *CodeableConcept
	// The date when or period to which this information refers.
	Timing isClaimSupportingInfoTiming
	// Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
	Value isClaimSupportingInfoValue
	// Provides the reason in the situation where a reason code is required in addition to the content.
	Reason *CodeableConcept
}
type isClaimSupportingInfoTiming interface {
	isClaimSupportingInfoTiming()
}

func (r Date) isClaimSupportingInfoTiming()   {}
func (r Period) isClaimSupportingInfoTiming() {}

type isClaimSupportingInfoValue interface {
	isClaimSupportingInfoValue()
}

func (r Boolean) isClaimSupportingInfoValue()    {}
func (r String) isClaimSupportingInfoValue()     {}
func (r Quantity) isClaimSupportingInfoValue()   {}
func (r Attachment) isClaimSupportingInfoValue() {}
func (r Reference) isClaimSupportingInfoValue()  {}

type jsonClaimSupportingInfo struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Sequence                     PositiveInt       `json:"sequence,omitempty"`
	SequencePrimitiveElement     *primitiveElement `json:"_sequence,omitempty"`
	Category                     CodeableConcept   `json:"category,omitempty"`
	Code                         *CodeableConcept  `json:"code,omitempty"`
	TimingDate                   *Date             `json:"timingDate,omitempty"`
	TimingDatePrimitiveElement   *primitiveElement `json:"_timingDate,omitempty"`
	TimingPeriod                 *Period           `json:"timingPeriod,omitempty"`
	ValueBoolean                 *Boolean          `json:"valueBoolean,omitempty"`
	ValueBooleanPrimitiveElement *primitiveElement `json:"_valueBoolean,omitempty"`
	ValueString                  *String           `json:"valueString,omitempty"`
	ValueStringPrimitiveElement  *primitiveElement `json:"_valueString,omitempty"`
	ValueQuantity                *Quantity         `json:"valueQuantity,omitempty"`
	ValueAttachment              *Attachment       `json:"valueAttachment,omitempty"`
	ValueReference               *Reference        `json:"valueReference,omitempty"`
	Reason                       *CodeableConcept  `json:"reason,omitempty"`
}

func (r ClaimSupportingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimSupportingInfo) marshalJSON() jsonClaimSupportingInfo {
	m := jsonClaimSupportingInfo{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Sequence = r.Sequence
	if r.Sequence.Id != nil || r.Sequence.Extension != nil {
		m.SequencePrimitiveElement = &primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
	}
	m.Category = r.Category
	m.Code = r.Code
	switch v := r.Timing.(type) {
	case Date:
		m.TimingDate = &v
		if v.Id != nil || v.Extension != nil {
			m.TimingDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		m.TimingDate = v
		if v.Id != nil || v.Extension != nil {
			m.TimingDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.TimingPeriod = &v
	case *Period:
		m.TimingPeriod = v
	}
	switch v := r.Value.(type) {
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
	case Quantity:
		m.ValueQuantity = &v
	case *Quantity:
		m.ValueQuantity = v
	case Attachment:
		m.ValueAttachment = &v
	case *Attachment:
		m.ValueAttachment = v
	case Reference:
		m.ValueReference = &v
	case *Reference:
		m.ValueReference = v
	}
	m.Reason = r.Reason
	return m
}
func (r *ClaimSupportingInfo) UnmarshalJSON(b []byte) error {
	var m jsonClaimSupportingInfo
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimSupportingInfo) unmarshalJSON(m jsonClaimSupportingInfo) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Sequence = m.Sequence
	if m.SequencePrimitiveElement != nil {
		r.Sequence.Id = m.SequencePrimitiveElement.Id
		r.Sequence.Extension = m.SequencePrimitiveElement.Extension
	}
	r.Category = m.Category
	r.Code = m.Code
	if m.TimingDate != nil || m.TimingDatePrimitiveElement != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingDate
		if m.TimingDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.TimingDatePrimitiveElement.Id
			v.Extension = m.TimingDatePrimitiveElement.Extension
		}
		r.Timing = v
	}
	if m.TimingPeriod != nil {
		if r.Timing != nil {
			return fmt.Errorf("multiple values for field \"Timing\"")
		}
		v := m.TimingPeriod
		r.Timing = v
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
	if m.ValueQuantity != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueQuantity
		r.Value = v
	}
	if m.ValueAttachment != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAttachment
		r.Value = v
	}
	if m.ValueReference != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueReference
		r.Value = v
	}
	r.Reason = m.Reason
	return nil
}
func (r ClaimSupportingInfo) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about diagnoses relevant to the claim items.
type ClaimDiagnosis struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely identify diagnosis entries.
	Sequence PositiveInt
	// The nature of illness or problem in a coded form or as a reference to an external defined Condition.
	Diagnosis isClaimDiagnosisDiagnosis
	// When the condition was observed or the relative ranking.
	Type []CodeableConcept
	// Indication of whether the diagnosis was present on admission to a facility.
	OnAdmission *CodeableConcept
	// A package billing code or bundle code used to group products and services to a particular health condition (such as heart attack) which is based on a predetermined grouping code system.
	PackageCode *CodeableConcept
}
type isClaimDiagnosisDiagnosis interface {
	isClaimDiagnosisDiagnosis()
}

func (r CodeableConcept) isClaimDiagnosisDiagnosis() {}
func (r Reference) isClaimDiagnosisDiagnosis()       {}

type jsonClaimDiagnosis struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Sequence                 PositiveInt       `json:"sequence,omitempty"`
	SequencePrimitiveElement *primitiveElement `json:"_sequence,omitempty"`
	DiagnosisCodeableConcept *CodeableConcept  `json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *Reference        `json:"diagnosisReference,omitempty"`
	Type                     []CodeableConcept `json:"type,omitempty"`
	OnAdmission              *CodeableConcept  `json:"onAdmission,omitempty"`
	PackageCode              *CodeableConcept  `json:"packageCode,omitempty"`
}

func (r ClaimDiagnosis) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimDiagnosis) marshalJSON() jsonClaimDiagnosis {
	m := jsonClaimDiagnosis{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Sequence = r.Sequence
	if r.Sequence.Id != nil || r.Sequence.Extension != nil {
		m.SequencePrimitiveElement = &primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
	}
	switch v := r.Diagnosis.(type) {
	case CodeableConcept:
		m.DiagnosisCodeableConcept = &v
	case *CodeableConcept:
		m.DiagnosisCodeableConcept = v
	case Reference:
		m.DiagnosisReference = &v
	case *Reference:
		m.DiagnosisReference = v
	}
	m.Type = r.Type
	m.OnAdmission = r.OnAdmission
	m.PackageCode = r.PackageCode
	return m
}
func (r *ClaimDiagnosis) UnmarshalJSON(b []byte) error {
	var m jsonClaimDiagnosis
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimDiagnosis) unmarshalJSON(m jsonClaimDiagnosis) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Sequence = m.Sequence
	if m.SequencePrimitiveElement != nil {
		r.Sequence.Id = m.SequencePrimitiveElement.Id
		r.Sequence.Extension = m.SequencePrimitiveElement.Extension
	}
	if m.DiagnosisCodeableConcept != nil {
		if r.Diagnosis != nil {
			return fmt.Errorf("multiple values for field \"Diagnosis\"")
		}
		v := m.DiagnosisCodeableConcept
		r.Diagnosis = v
	}
	if m.DiagnosisReference != nil {
		if r.Diagnosis != nil {
			return fmt.Errorf("multiple values for field \"Diagnosis\"")
		}
		v := m.DiagnosisReference
		r.Diagnosis = v
	}
	r.Type = m.Type
	r.OnAdmission = m.OnAdmission
	r.PackageCode = m.PackageCode
	return nil
}
func (r ClaimDiagnosis) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Procedures performed on the patient relevant to the billing items with the claim.
type ClaimProcedure struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely identify procedure entries.
	Sequence PositiveInt
	// When the condition was observed or the relative ranking.
	Type []CodeableConcept
	// Date and optionally time the procedure was performed.
	Date *DateTime
	// The code or reference to a Procedure resource which identifies the clinical intervention performed.
	Procedure isClaimProcedureProcedure
	// Unique Device Identifiers associated with this line item.
	Udi []Reference
}
type isClaimProcedureProcedure interface {
	isClaimProcedureProcedure()
}

func (r CodeableConcept) isClaimProcedureProcedure() {}
func (r Reference) isClaimProcedureProcedure()       {}

type jsonClaimProcedure struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Sequence                 PositiveInt       `json:"sequence,omitempty"`
	SequencePrimitiveElement *primitiveElement `json:"_sequence,omitempty"`
	Type                     []CodeableConcept `json:"type,omitempty"`
	Date                     *DateTime         `json:"date,omitempty"`
	DatePrimitiveElement     *primitiveElement `json:"_date,omitempty"`
	ProcedureCodeableConcept *CodeableConcept  `json:"procedureCodeableConcept,omitempty"`
	ProcedureReference       *Reference        `json:"procedureReference,omitempty"`
	Udi                      []Reference       `json:"udi,omitempty"`
}

func (r ClaimProcedure) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimProcedure) marshalJSON() jsonClaimProcedure {
	m := jsonClaimProcedure{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Sequence = r.Sequence
	if r.Sequence.Id != nil || r.Sequence.Extension != nil {
		m.SequencePrimitiveElement = &primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
	}
	m.Type = r.Type
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	switch v := r.Procedure.(type) {
	case CodeableConcept:
		m.ProcedureCodeableConcept = &v
	case *CodeableConcept:
		m.ProcedureCodeableConcept = v
	case Reference:
		m.ProcedureReference = &v
	case *Reference:
		m.ProcedureReference = v
	}
	m.Udi = r.Udi
	return m
}
func (r *ClaimProcedure) UnmarshalJSON(b []byte) error {
	var m jsonClaimProcedure
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimProcedure) unmarshalJSON(m jsonClaimProcedure) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Sequence = m.Sequence
	if m.SequencePrimitiveElement != nil {
		r.Sequence.Id = m.SequencePrimitiveElement.Id
		r.Sequence.Extension = m.SequencePrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	if m.ProcedureCodeableConcept != nil {
		if r.Procedure != nil {
			return fmt.Errorf("multiple values for field \"Procedure\"")
		}
		v := m.ProcedureCodeableConcept
		r.Procedure = v
	}
	if m.ProcedureReference != nil {
		if r.Procedure != nil {
			return fmt.Errorf("multiple values for field \"Procedure\"")
		}
		v := m.ProcedureReference
		r.Procedure = v
	}
	r.Udi = m.Udi
	return nil
}
func (r ClaimProcedure) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Financial instruments for reimbursement for the health care products and services specified on the claim.
type ClaimInsurance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely identify insurance entries and provide a sequence of coverages to convey coordination of benefit order.
	Sequence PositiveInt
	// A flag to indicate that this Coverage is to be used for adjudication of this claim when set to true.
	Focal Boolean
	// The business identifier to be used when the claim is sent for adjudication against this insurance policy.
	Identifier *Identifier
	// Reference to the insurance card level information contained in the Coverage resource. The coverage issuing insurer will use these details to locate the patient's actual coverage within the insurer's information system.
	Coverage Reference
	// A business agreement number established between the provider and the insurer for special business processing purposes.
	BusinessArrangement *String
	// Reference numbers previously provided by the insurer to the provider to be quoted on subsequent claims containing services or products related to the prior authorization.
	PreAuthRef []String
	// The result of the adjudication of the line items for the Coverage specified in this insurance.
	ClaimResponse *Reference
}
type jsonClaimInsurance struct {
	Id                                  *string             `json:"id,omitempty"`
	Extension                           []Extension         `json:"extension,omitempty"`
	ModifierExtension                   []Extension         `json:"modifierExtension,omitempty"`
	Sequence                            PositiveInt         `json:"sequence,omitempty"`
	SequencePrimitiveElement            *primitiveElement   `json:"_sequence,omitempty"`
	Focal                               Boolean             `json:"focal,omitempty"`
	FocalPrimitiveElement               *primitiveElement   `json:"_focal,omitempty"`
	Identifier                          *Identifier         `json:"identifier,omitempty"`
	Coverage                            Reference           `json:"coverage,omitempty"`
	BusinessArrangement                 *String             `json:"businessArrangement,omitempty"`
	BusinessArrangementPrimitiveElement *primitiveElement   `json:"_businessArrangement,omitempty"`
	PreAuthRef                          []String            `json:"preAuthRef,omitempty"`
	PreAuthRefPrimitiveElement          []*primitiveElement `json:"_preAuthRef,omitempty"`
	ClaimResponse                       *Reference          `json:"claimResponse,omitempty"`
}

func (r ClaimInsurance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimInsurance) marshalJSON() jsonClaimInsurance {
	m := jsonClaimInsurance{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Sequence = r.Sequence
	if r.Sequence.Id != nil || r.Sequence.Extension != nil {
		m.SequencePrimitiveElement = &primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
	}
	m.Focal = r.Focal
	if r.Focal.Id != nil || r.Focal.Extension != nil {
		m.FocalPrimitiveElement = &primitiveElement{Id: r.Focal.Id, Extension: r.Focal.Extension}
	}
	m.Identifier = r.Identifier
	m.Coverage = r.Coverage
	m.BusinessArrangement = r.BusinessArrangement
	if r.BusinessArrangement != nil && (r.BusinessArrangement.Id != nil || r.BusinessArrangement.Extension != nil) {
		m.BusinessArrangementPrimitiveElement = &primitiveElement{Id: r.BusinessArrangement.Id, Extension: r.BusinessArrangement.Extension}
	}
	m.PreAuthRef = r.PreAuthRef
	anyPreAuthRefIdOrExtension := false
	for _, e := range r.PreAuthRef {
		if e.Id != nil || e.Extension != nil {
			anyPreAuthRefIdOrExtension = true
			break
		}
	}
	if anyPreAuthRefIdOrExtension {
		m.PreAuthRefPrimitiveElement = make([]*primitiveElement, 0, len(r.PreAuthRef))
		for _, e := range r.PreAuthRef {
			if e.Id != nil || e.Extension != nil {
				m.PreAuthRefPrimitiveElement = append(m.PreAuthRefPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.PreAuthRefPrimitiveElement = append(m.PreAuthRefPrimitiveElement, nil)
			}
		}
	}
	m.ClaimResponse = r.ClaimResponse
	return m
}
func (r *ClaimInsurance) UnmarshalJSON(b []byte) error {
	var m jsonClaimInsurance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimInsurance) unmarshalJSON(m jsonClaimInsurance) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Sequence = m.Sequence
	if m.SequencePrimitiveElement != nil {
		r.Sequence.Id = m.SequencePrimitiveElement.Id
		r.Sequence.Extension = m.SequencePrimitiveElement.Extension
	}
	r.Focal = m.Focal
	if m.FocalPrimitiveElement != nil {
		r.Focal.Id = m.FocalPrimitiveElement.Id
		r.Focal.Extension = m.FocalPrimitiveElement.Extension
	}
	r.Identifier = m.Identifier
	r.Coverage = m.Coverage
	r.BusinessArrangement = m.BusinessArrangement
	if m.BusinessArrangementPrimitiveElement != nil {
		r.BusinessArrangement.Id = m.BusinessArrangementPrimitiveElement.Id
		r.BusinessArrangement.Extension = m.BusinessArrangementPrimitiveElement.Extension
	}
	r.PreAuthRef = m.PreAuthRef
	for i, e := range m.PreAuthRefPrimitiveElement {
		if len(r.PreAuthRef) > i {
			r.PreAuthRef[i].Id = e.Id
			r.PreAuthRef[i].Extension = e.Extension
		} else {
			r.PreAuthRef = append(r.PreAuthRef, String{Id: e.Id, Extension: e.Extension})
		}
	}
	r.ClaimResponse = m.ClaimResponse
	return nil
}
func (r ClaimInsurance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Details of an accident which resulted in injuries which required the products and services listed in the claim.
type ClaimAccident struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Date of an accident event  related to the products and services contained in the claim.
	Date Date
	// The type or context of the accident event for the purposes of selection of potential insurance coverages and determination of coordination between insurers.
	Type *CodeableConcept
	// The physical location of the accident event.
	Location isClaimAccidentLocation
}
type isClaimAccidentLocation interface {
	isClaimAccidentLocation()
}

func (r Address) isClaimAccidentLocation()   {}
func (r Reference) isClaimAccidentLocation() {}

type jsonClaimAccident struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Date                 Date              `json:"date,omitempty"`
	DatePrimitiveElement *primitiveElement `json:"_date,omitempty"`
	Type                 *CodeableConcept  `json:"type,omitempty"`
	LocationAddress      *Address          `json:"locationAddress,omitempty"`
	LocationReference    *Reference        `json:"locationReference,omitempty"`
}

func (r ClaimAccident) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimAccident) marshalJSON() jsonClaimAccident {
	m := jsonClaimAccident{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Date = r.Date
	if r.Date.Id != nil || r.Date.Extension != nil {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Type = r.Type
	switch v := r.Location.(type) {
	case Address:
		m.LocationAddress = &v
	case *Address:
		m.LocationAddress = v
	case Reference:
		m.LocationReference = &v
	case *Reference:
		m.LocationReference = v
	}
	return m
}
func (r *ClaimAccident) UnmarshalJSON(b []byte) error {
	var m jsonClaimAccident
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimAccident) unmarshalJSON(m jsonClaimAccident) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.LocationAddress != nil {
		if r.Location != nil {
			return fmt.Errorf("multiple values for field \"Location\"")
		}
		v := m.LocationAddress
		r.Location = v
	}
	if m.LocationReference != nil {
		if r.Location != nil {
			return fmt.Errorf("multiple values for field \"Location\"")
		}
		v := m.LocationReference
		r.Location = v
	}
	return nil
}
func (r ClaimAccident) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A claim line. Either a simple  product or service or a 'group' of details which can each be a simple items or groups of sub-details.
type ClaimItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely identify item entries.
	Sequence PositiveInt
	// CareTeam members related to this service or product.
	CareTeamSequence []PositiveInt
	// Diagnosis applicable for this service or product.
	DiagnosisSequence []PositiveInt
	// Procedures applicable for this service or product.
	ProcedureSequence []PositiveInt
	// Exceptions, special conditions and supporting information applicable for this service or product.
	InformationSequence []PositiveInt
	// The type of revenue or cost center providing the product and/or service.
	Revenue *CodeableConcept
	// Code to identify the general type of benefits under which products and services are provided.
	Category *CodeableConcept
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []CodeableConcept
	// Identifies the program under which this may be recovered.
	ProgramCode []CodeableConcept
	// The date or dates when the service or product was supplied, performed or completed.
	Serviced isClaimItemServiced
	// Where the product or service was provided.
	Location isClaimItemLocation
	// The number of repetitions of a service or product.
	Quantity *Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *Decimal
	// The quantity times the unit price for an additional service or product or charge.
	Net *Money
	// Unique Device Identifiers associated with this line item.
	Udi []Reference
	// Physical service site on the patient (limb, tooth, etc.).
	BodySite *CodeableConcept
	// A region or surface of the bodySite, e.g. limb region or tooth surface(s).
	SubSite []CodeableConcept
	// The Encounters during which this Claim was created or to which the creation of this record is tightly associated.
	Encounter []Reference
	// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
	Detail []ClaimItemDetail
}
type isClaimItemServiced interface {
	isClaimItemServiced()
}

func (r Date) isClaimItemServiced()   {}
func (r Period) isClaimItemServiced() {}

type isClaimItemLocation interface {
	isClaimItemLocation()
}

func (r CodeableConcept) isClaimItemLocation() {}
func (r Address) isClaimItemLocation()         {}
func (r Reference) isClaimItemLocation()       {}

type jsonClaimItem struct {
	Id                                  *string             `json:"id,omitempty"`
	Extension                           []Extension         `json:"extension,omitempty"`
	ModifierExtension                   []Extension         `json:"modifierExtension,omitempty"`
	Sequence                            PositiveInt         `json:"sequence,omitempty"`
	SequencePrimitiveElement            *primitiveElement   `json:"_sequence,omitempty"`
	CareTeamSequence                    []PositiveInt       `json:"careTeamSequence,omitempty"`
	CareTeamSequencePrimitiveElement    []*primitiveElement `json:"_careTeamSequence,omitempty"`
	DiagnosisSequence                   []PositiveInt       `json:"diagnosisSequence,omitempty"`
	DiagnosisSequencePrimitiveElement   []*primitiveElement `json:"_diagnosisSequence,omitempty"`
	ProcedureSequence                   []PositiveInt       `json:"procedureSequence,omitempty"`
	ProcedureSequencePrimitiveElement   []*primitiveElement `json:"_procedureSequence,omitempty"`
	InformationSequence                 []PositiveInt       `json:"informationSequence,omitempty"`
	InformationSequencePrimitiveElement []*primitiveElement `json:"_informationSequence,omitempty"`
	Revenue                             *CodeableConcept    `json:"revenue,omitempty"`
	Category                            *CodeableConcept    `json:"category,omitempty"`
	ProductOrService                    CodeableConcept     `json:"productOrService,omitempty"`
	Modifier                            []CodeableConcept   `json:"modifier,omitempty"`
	ProgramCode                         []CodeableConcept   `json:"programCode,omitempty"`
	ServicedDate                        *Date               `json:"servicedDate,omitempty"`
	ServicedDatePrimitiveElement        *primitiveElement   `json:"_servicedDate,omitempty"`
	ServicedPeriod                      *Period             `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept             *CodeableConcept    `json:"locationCodeableConcept,omitempty"`
	LocationAddress                     *Address            `json:"locationAddress,omitempty"`
	LocationReference                   *Reference          `json:"locationReference,omitempty"`
	Quantity                            *Quantity           `json:"quantity,omitempty"`
	UnitPrice                           *Money              `json:"unitPrice,omitempty"`
	Factor                              *Decimal            `json:"factor,omitempty"`
	FactorPrimitiveElement              *primitiveElement   `json:"_factor,omitempty"`
	Net                                 *Money              `json:"net,omitempty"`
	Udi                                 []Reference         `json:"udi,omitempty"`
	BodySite                            *CodeableConcept    `json:"bodySite,omitempty"`
	SubSite                             []CodeableConcept   `json:"subSite,omitempty"`
	Encounter                           []Reference         `json:"encounter,omitempty"`
	Detail                              []ClaimItemDetail   `json:"detail,omitempty"`
}

func (r ClaimItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimItem) marshalJSON() jsonClaimItem {
	m := jsonClaimItem{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Sequence = r.Sequence
	if r.Sequence.Id != nil || r.Sequence.Extension != nil {
		m.SequencePrimitiveElement = &primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
	}
	m.CareTeamSequence = r.CareTeamSequence
	anyCareTeamSequenceIdOrExtension := false
	for _, e := range r.CareTeamSequence {
		if e.Id != nil || e.Extension != nil {
			anyCareTeamSequenceIdOrExtension = true
			break
		}
	}
	if anyCareTeamSequenceIdOrExtension {
		m.CareTeamSequencePrimitiveElement = make([]*primitiveElement, 0, len(r.CareTeamSequence))
		for _, e := range r.CareTeamSequence {
			if e.Id != nil || e.Extension != nil {
				m.CareTeamSequencePrimitiveElement = append(m.CareTeamSequencePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.CareTeamSequencePrimitiveElement = append(m.CareTeamSequencePrimitiveElement, nil)
			}
		}
	}
	m.DiagnosisSequence = r.DiagnosisSequence
	anyDiagnosisSequenceIdOrExtension := false
	for _, e := range r.DiagnosisSequence {
		if e.Id != nil || e.Extension != nil {
			anyDiagnosisSequenceIdOrExtension = true
			break
		}
	}
	if anyDiagnosisSequenceIdOrExtension {
		m.DiagnosisSequencePrimitiveElement = make([]*primitiveElement, 0, len(r.DiagnosisSequence))
		for _, e := range r.DiagnosisSequence {
			if e.Id != nil || e.Extension != nil {
				m.DiagnosisSequencePrimitiveElement = append(m.DiagnosisSequencePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.DiagnosisSequencePrimitiveElement = append(m.DiagnosisSequencePrimitiveElement, nil)
			}
		}
	}
	m.ProcedureSequence = r.ProcedureSequence
	anyProcedureSequenceIdOrExtension := false
	for _, e := range r.ProcedureSequence {
		if e.Id != nil || e.Extension != nil {
			anyProcedureSequenceIdOrExtension = true
			break
		}
	}
	if anyProcedureSequenceIdOrExtension {
		m.ProcedureSequencePrimitiveElement = make([]*primitiveElement, 0, len(r.ProcedureSequence))
		for _, e := range r.ProcedureSequence {
			if e.Id != nil || e.Extension != nil {
				m.ProcedureSequencePrimitiveElement = append(m.ProcedureSequencePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ProcedureSequencePrimitiveElement = append(m.ProcedureSequencePrimitiveElement, nil)
			}
		}
	}
	m.InformationSequence = r.InformationSequence
	anyInformationSequenceIdOrExtension := false
	for _, e := range r.InformationSequence {
		if e.Id != nil || e.Extension != nil {
			anyInformationSequenceIdOrExtension = true
			break
		}
	}
	if anyInformationSequenceIdOrExtension {
		m.InformationSequencePrimitiveElement = make([]*primitiveElement, 0, len(r.InformationSequence))
		for _, e := range r.InformationSequence {
			if e.Id != nil || e.Extension != nil {
				m.InformationSequencePrimitiveElement = append(m.InformationSequencePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.InformationSequencePrimitiveElement = append(m.InformationSequencePrimitiveElement, nil)
			}
		}
	}
	m.Revenue = r.Revenue
	m.Category = r.Category
	m.ProductOrService = r.ProductOrService
	m.Modifier = r.Modifier
	m.ProgramCode = r.ProgramCode
	switch v := r.Serviced.(type) {
	case Date:
		m.ServicedDate = &v
		if v.Id != nil || v.Extension != nil {
			m.ServicedDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		m.ServicedDate = v
		if v.Id != nil || v.Extension != nil {
			m.ServicedDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.ServicedPeriod = &v
	case *Period:
		m.ServicedPeriod = v
	}
	switch v := r.Location.(type) {
	case CodeableConcept:
		m.LocationCodeableConcept = &v
	case *CodeableConcept:
		m.LocationCodeableConcept = v
	case Address:
		m.LocationAddress = &v
	case *Address:
		m.LocationAddress = v
	case Reference:
		m.LocationReference = &v
	case *Reference:
		m.LocationReference = v
	}
	m.Quantity = r.Quantity
	m.UnitPrice = r.UnitPrice
	m.Factor = r.Factor
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		m.FactorPrimitiveElement = &primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
	}
	m.Net = r.Net
	m.Udi = r.Udi
	m.BodySite = r.BodySite
	m.SubSite = r.SubSite
	m.Encounter = r.Encounter
	m.Detail = r.Detail
	return m
}
func (r *ClaimItem) UnmarshalJSON(b []byte) error {
	var m jsonClaimItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimItem) unmarshalJSON(m jsonClaimItem) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Sequence = m.Sequence
	if m.SequencePrimitiveElement != nil {
		r.Sequence.Id = m.SequencePrimitiveElement.Id
		r.Sequence.Extension = m.SequencePrimitiveElement.Extension
	}
	r.CareTeamSequence = m.CareTeamSequence
	for i, e := range m.CareTeamSequencePrimitiveElement {
		if len(r.CareTeamSequence) > i {
			r.CareTeamSequence[i].Id = e.Id
			r.CareTeamSequence[i].Extension = e.Extension
		} else {
			r.CareTeamSequence = append(r.CareTeamSequence, PositiveInt{Id: e.Id, Extension: e.Extension})
		}
	}
	r.DiagnosisSequence = m.DiagnosisSequence
	for i, e := range m.DiagnosisSequencePrimitiveElement {
		if len(r.DiagnosisSequence) > i {
			r.DiagnosisSequence[i].Id = e.Id
			r.DiagnosisSequence[i].Extension = e.Extension
		} else {
			r.DiagnosisSequence = append(r.DiagnosisSequence, PositiveInt{Id: e.Id, Extension: e.Extension})
		}
	}
	r.ProcedureSequence = m.ProcedureSequence
	for i, e := range m.ProcedureSequencePrimitiveElement {
		if len(r.ProcedureSequence) > i {
			r.ProcedureSequence[i].Id = e.Id
			r.ProcedureSequence[i].Extension = e.Extension
		} else {
			r.ProcedureSequence = append(r.ProcedureSequence, PositiveInt{Id: e.Id, Extension: e.Extension})
		}
	}
	r.InformationSequence = m.InformationSequence
	for i, e := range m.InformationSequencePrimitiveElement {
		if len(r.InformationSequence) > i {
			r.InformationSequence[i].Id = e.Id
			r.InformationSequence[i].Extension = e.Extension
		} else {
			r.InformationSequence = append(r.InformationSequence, PositiveInt{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Revenue = m.Revenue
	r.Category = m.Category
	r.ProductOrService = m.ProductOrService
	r.Modifier = m.Modifier
	r.ProgramCode = m.ProgramCode
	if m.ServicedDate != nil || m.ServicedDatePrimitiveElement != nil {
		if r.Serviced != nil {
			return fmt.Errorf("multiple values for field \"Serviced\"")
		}
		v := m.ServicedDate
		if m.ServicedDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.ServicedDatePrimitiveElement.Id
			v.Extension = m.ServicedDatePrimitiveElement.Extension
		}
		r.Serviced = v
	}
	if m.ServicedPeriod != nil {
		if r.Serviced != nil {
			return fmt.Errorf("multiple values for field \"Serviced\"")
		}
		v := m.ServicedPeriod
		r.Serviced = v
	}
	if m.LocationCodeableConcept != nil {
		if r.Location != nil {
			return fmt.Errorf("multiple values for field \"Location\"")
		}
		v := m.LocationCodeableConcept
		r.Location = v
	}
	if m.LocationAddress != nil {
		if r.Location != nil {
			return fmt.Errorf("multiple values for field \"Location\"")
		}
		v := m.LocationAddress
		r.Location = v
	}
	if m.LocationReference != nil {
		if r.Location != nil {
			return fmt.Errorf("multiple values for field \"Location\"")
		}
		v := m.LocationReference
		r.Location = v
	}
	r.Quantity = m.Quantity
	r.UnitPrice = m.UnitPrice
	r.Factor = m.Factor
	if m.FactorPrimitiveElement != nil {
		r.Factor.Id = m.FactorPrimitiveElement.Id
		r.Factor.Extension = m.FactorPrimitiveElement.Extension
	}
	r.Net = m.Net
	r.Udi = m.Udi
	r.BodySite = m.BodySite
	r.SubSite = m.SubSite
	r.Encounter = m.Encounter
	r.Detail = m.Detail
	return nil
}
func (r ClaimItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
type ClaimItemDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely identify item entries.
	Sequence PositiveInt
	// The type of revenue or cost center providing the product and/or service.
	Revenue *CodeableConcept
	// Code to identify the general type of benefits under which products and services are provided.
	Category *CodeableConcept
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []CodeableConcept
	// Identifies the program under which this may be recovered.
	ProgramCode []CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *Decimal
	// The quantity times the unit price for an additional service or product or charge.
	Net *Money
	// Unique Device Identifiers associated with this line item.
	Udi []Reference
	// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
	SubDetail []ClaimItemDetailSubDetail
}
type jsonClaimItemDetail struct {
	Id                       *string                    `json:"id,omitempty"`
	Extension                []Extension                `json:"extension,omitempty"`
	ModifierExtension        []Extension                `json:"modifierExtension,omitempty"`
	Sequence                 PositiveInt                `json:"sequence,omitempty"`
	SequencePrimitiveElement *primitiveElement          `json:"_sequence,omitempty"`
	Revenue                  *CodeableConcept           `json:"revenue,omitempty"`
	Category                 *CodeableConcept           `json:"category,omitempty"`
	ProductOrService         CodeableConcept            `json:"productOrService,omitempty"`
	Modifier                 []CodeableConcept          `json:"modifier,omitempty"`
	ProgramCode              []CodeableConcept          `json:"programCode,omitempty"`
	Quantity                 *Quantity                  `json:"quantity,omitempty"`
	UnitPrice                *Money                     `json:"unitPrice,omitempty"`
	Factor                   *Decimal                   `json:"factor,omitempty"`
	FactorPrimitiveElement   *primitiveElement          `json:"_factor,omitempty"`
	Net                      *Money                     `json:"net,omitempty"`
	Udi                      []Reference                `json:"udi,omitempty"`
	SubDetail                []ClaimItemDetailSubDetail `json:"subDetail,omitempty"`
}

func (r ClaimItemDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimItemDetail) marshalJSON() jsonClaimItemDetail {
	m := jsonClaimItemDetail{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Sequence = r.Sequence
	if r.Sequence.Id != nil || r.Sequence.Extension != nil {
		m.SequencePrimitiveElement = &primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
	}
	m.Revenue = r.Revenue
	m.Category = r.Category
	m.ProductOrService = r.ProductOrService
	m.Modifier = r.Modifier
	m.ProgramCode = r.ProgramCode
	m.Quantity = r.Quantity
	m.UnitPrice = r.UnitPrice
	m.Factor = r.Factor
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		m.FactorPrimitiveElement = &primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
	}
	m.Net = r.Net
	m.Udi = r.Udi
	m.SubDetail = r.SubDetail
	return m
}
func (r *ClaimItemDetail) UnmarshalJSON(b []byte) error {
	var m jsonClaimItemDetail
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimItemDetail) unmarshalJSON(m jsonClaimItemDetail) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Sequence = m.Sequence
	if m.SequencePrimitiveElement != nil {
		r.Sequence.Id = m.SequencePrimitiveElement.Id
		r.Sequence.Extension = m.SequencePrimitiveElement.Extension
	}
	r.Revenue = m.Revenue
	r.Category = m.Category
	r.ProductOrService = m.ProductOrService
	r.Modifier = m.Modifier
	r.ProgramCode = m.ProgramCode
	r.Quantity = m.Quantity
	r.UnitPrice = m.UnitPrice
	r.Factor = m.Factor
	if m.FactorPrimitiveElement != nil {
		r.Factor.Id = m.FactorPrimitiveElement.Id
		r.Factor.Extension = m.FactorPrimitiveElement.Extension
	}
	r.Net = m.Net
	r.Udi = m.Udi
	r.SubDetail = m.SubDetail
	return nil
}
func (r ClaimItemDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
type ClaimItemDetailSubDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely identify item entries.
	Sequence PositiveInt
	// The type of revenue or cost center providing the product and/or service.
	Revenue *CodeableConcept
	// Code to identify the general type of benefits under which products and services are provided.
	Category *CodeableConcept
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []CodeableConcept
	// Identifies the program under which this may be recovered.
	ProgramCode []CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *Decimal
	// The quantity times the unit price for an additional service or product or charge.
	Net *Money
	// Unique Device Identifiers associated with this line item.
	Udi []Reference
}
type jsonClaimItemDetailSubDetail struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Sequence                 PositiveInt       `json:"sequence,omitempty"`
	SequencePrimitiveElement *primitiveElement `json:"_sequence,omitempty"`
	Revenue                  *CodeableConcept  `json:"revenue,omitempty"`
	Category                 *CodeableConcept  `json:"category,omitempty"`
	ProductOrService         CodeableConcept   `json:"productOrService,omitempty"`
	Modifier                 []CodeableConcept `json:"modifier,omitempty"`
	ProgramCode              []CodeableConcept `json:"programCode,omitempty"`
	Quantity                 *Quantity         `json:"quantity,omitempty"`
	UnitPrice                *Money            `json:"unitPrice,omitempty"`
	Factor                   *Decimal          `json:"factor,omitempty"`
	FactorPrimitiveElement   *primitiveElement `json:"_factor,omitempty"`
	Net                      *Money            `json:"net,omitempty"`
	Udi                      []Reference       `json:"udi,omitempty"`
}

func (r ClaimItemDetailSubDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimItemDetailSubDetail) marshalJSON() jsonClaimItemDetailSubDetail {
	m := jsonClaimItemDetailSubDetail{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Sequence = r.Sequence
	if r.Sequence.Id != nil || r.Sequence.Extension != nil {
		m.SequencePrimitiveElement = &primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
	}
	m.Revenue = r.Revenue
	m.Category = r.Category
	m.ProductOrService = r.ProductOrService
	m.Modifier = r.Modifier
	m.ProgramCode = r.ProgramCode
	m.Quantity = r.Quantity
	m.UnitPrice = r.UnitPrice
	m.Factor = r.Factor
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		m.FactorPrimitiveElement = &primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
	}
	m.Net = r.Net
	m.Udi = r.Udi
	return m
}
func (r *ClaimItemDetailSubDetail) UnmarshalJSON(b []byte) error {
	var m jsonClaimItemDetailSubDetail
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimItemDetailSubDetail) unmarshalJSON(m jsonClaimItemDetailSubDetail) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Sequence = m.Sequence
	if m.SequencePrimitiveElement != nil {
		r.Sequence.Id = m.SequencePrimitiveElement.Id
		r.Sequence.Extension = m.SequencePrimitiveElement.Extension
	}
	r.Revenue = m.Revenue
	r.Category = m.Category
	r.ProductOrService = m.ProductOrService
	r.Modifier = m.Modifier
	r.ProgramCode = m.ProgramCode
	r.Quantity = m.Quantity
	r.UnitPrice = m.UnitPrice
	r.Factor = m.Factor
	if m.FactorPrimitiveElement != nil {
		r.Factor.Id = m.FactorPrimitiveElement.Id
		r.Factor.Extension = m.FactorPrimitiveElement.Extension
	}
	r.Net = m.Net
	r.Udi = m.Udi
	return nil
}
func (r ClaimItemDetailSubDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
