package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit struct {
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
	// A unique identifier assigned to this explanation of benefit.
	Identifier []Identifier
	// The status of the resource instance.
	Status Code
	// The category of claim, e.g. oral, pharmacy, vision, institutional, professional.
	Type CodeableConcept
	// A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
	SubType *CodeableConcept
	// A code to indicate whether the nature of the request is: to request adjudication of products and services previously rendered; or requesting authorization and adjudication for provision in the future; or requesting the non-binding adjudication of the listed products and services which could be provided in the future.
	Use Code
	// The party to whom the professional services and/or products have been supplied or are being considered and for whom actual for forecast reimbursement is sought.
	Patient Reference
	// The period for which charges are being submitted.
	BillablePeriod *Period
	// The date this resource was created.
	Created DateTime
	// Individual who created the claim, predetermination or preauthorization.
	Enterer *Reference
	// The party responsible for authorization, adjudication and reimbursement.
	Insurer Reference
	// The provider which is responsible for the claim, predetermination or preauthorization.
	Provider Reference
	// The provider-required urgency of processing the request. Typical values include: stat, routine deferred.
	Priority *CodeableConcept
	// A code to indicate whether and for whom funds are to be reserved for future claims.
	FundsReserveRequested *CodeableConcept
	// A code, used only on a response to a preauthorization, to indicate whether the benefits payable have been reserved and for whom.
	FundsReserve *CodeableConcept
	// Other claims which are related to this claim such as prior submissions or claims for related services or for the same event.
	Related []ExplanationOfBenefitRelated
	// Prescription to support the dispensing of pharmacy, device or vision products.
	Prescription *Reference
	// Original prescription which has been superseded by this prescription to support the dispensing of pharmacy services, medications or products.
	OriginalPrescription *Reference
	// The party to be reimbursed for cost of the products and services according to the terms of the policy.
	Payee *ExplanationOfBenefitPayee
	// A reference to a referral resource.
	Referral *Reference
	// Facility where the services were provided.
	Facility *Reference
	// The business identifier for the instance of the adjudication request: claim predetermination or preauthorization.
	Claim *Reference
	// The business identifier for the instance of the adjudication response: claim, predetermination or preauthorization response.
	ClaimResponse *Reference
	// The outcome of the claim, predetermination, or preauthorization processing.
	Outcome Code
	// A human readable description of the status of the adjudication.
	Disposition *String
	// Reference from the Insurer which is used in later communications which refers to this adjudication.
	PreAuthRef []String
	// The timeframe during which the supplied preauthorization reference may be quoted on claims to obtain the adjudication as provided.
	PreAuthRefPeriod []Period
	// The members of the team who provided the products and services.
	CareTeam []ExplanationOfBenefitCareTeam
	// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
	SupportingInfo []ExplanationOfBenefitSupportingInfo
	// Information about diagnoses relevant to the claim items.
	Diagnosis []ExplanationOfBenefitDiagnosis
	// Procedures performed on the patient relevant to the billing items with the claim.
	Procedure []ExplanationOfBenefitProcedure
	// This indicates the relative order of a series of EOBs related to different coverages for the same suite of services.
	Precedence *PositiveInt
	// Financial instruments for reimbursement for the health care products and services specified on the claim.
	Insurance []ExplanationOfBenefitInsurance
	// Details of a accident which resulted in injuries which required the products and services listed in the claim.
	Accident *ExplanationOfBenefitAccident
	// A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
	Item []ExplanationOfBenefitItem
	// The first-tier service adjudications for payor added product or service lines.
	AddItem []ExplanationOfBenefitAddItem
	// The adjudication results which are presented at the header level rather than at the line-item or add-item levels.
	Adjudication []ExplanationOfBenefitItemAdjudication
	// Categorized monetary totals for the adjudication.
	Total []ExplanationOfBenefitTotal
	// Payment details for the adjudication of the claim.
	Payment *ExplanationOfBenefitPayment
	// A code for the form to be used for printing the content.
	FormCode *CodeableConcept
	// The actual form, by reference or inclusion, for printing the content or an EOB.
	Form *Attachment
	// A note that describes or explains adjudication results in a human readable form.
	ProcessNote []ExplanationOfBenefitProcessNote
	// The term of the benefits documented in this response.
	BenefitPeriod *Period
	// Balance by Benefit Category.
	BenefitBalance []ExplanationOfBenefitBenefitBalance
}
type jsonExplanationOfBenefit struct {
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
	Identifier                    []Identifier                           `json:"identifier,omitempty"`
	Status                        Code                                   `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement                      `json:"_status,omitempty"`
	Type                          CodeableConcept                        `json:"type,omitempty"`
	SubType                       *CodeableConcept                       `json:"subType,omitempty"`
	Use                           Code                                   `json:"use,omitempty"`
	UsePrimitiveElement           *primitiveElement                      `json:"_use,omitempty"`
	Patient                       Reference                              `json:"patient,omitempty"`
	BillablePeriod                *Period                                `json:"billablePeriod,omitempty"`
	Created                       DateTime                               `json:"created,omitempty"`
	CreatedPrimitiveElement       *primitiveElement                      `json:"_created,omitempty"`
	Enterer                       *Reference                             `json:"enterer,omitempty"`
	Insurer                       Reference                              `json:"insurer,omitempty"`
	Provider                      Reference                              `json:"provider,omitempty"`
	Priority                      *CodeableConcept                       `json:"priority,omitempty"`
	FundsReserveRequested         *CodeableConcept                       `json:"fundsReserveRequested,omitempty"`
	FundsReserve                  *CodeableConcept                       `json:"fundsReserve,omitempty"`
	Related                       []ExplanationOfBenefitRelated          `json:"related,omitempty"`
	Prescription                  *Reference                             `json:"prescription,omitempty"`
	OriginalPrescription          *Reference                             `json:"originalPrescription,omitempty"`
	Payee                         *ExplanationOfBenefitPayee             `json:"payee,omitempty"`
	Referral                      *Reference                             `json:"referral,omitempty"`
	Facility                      *Reference                             `json:"facility,omitempty"`
	Claim                         *Reference                             `json:"claim,omitempty"`
	ClaimResponse                 *Reference                             `json:"claimResponse,omitempty"`
	Outcome                       Code                                   `json:"outcome,omitempty"`
	OutcomePrimitiveElement       *primitiveElement                      `json:"_outcome,omitempty"`
	Disposition                   *String                                `json:"disposition,omitempty"`
	DispositionPrimitiveElement   *primitiveElement                      `json:"_disposition,omitempty"`
	PreAuthRef                    []String                               `json:"preAuthRef,omitempty"`
	PreAuthRefPrimitiveElement    []*primitiveElement                    `json:"_preAuthRef,omitempty"`
	PreAuthRefPeriod              []Period                               `json:"preAuthRefPeriod,omitempty"`
	CareTeam                      []ExplanationOfBenefitCareTeam         `json:"careTeam,omitempty"`
	SupportingInfo                []ExplanationOfBenefitSupportingInfo   `json:"supportingInfo,omitempty"`
	Diagnosis                     []ExplanationOfBenefitDiagnosis        `json:"diagnosis,omitempty"`
	Procedure                     []ExplanationOfBenefitProcedure        `json:"procedure,omitempty"`
	Precedence                    *PositiveInt                           `json:"precedence,omitempty"`
	PrecedencePrimitiveElement    *primitiveElement                      `json:"_precedence,omitempty"`
	Insurance                     []ExplanationOfBenefitInsurance        `json:"insurance,omitempty"`
	Accident                      *ExplanationOfBenefitAccident          `json:"accident,omitempty"`
	Item                          []ExplanationOfBenefitItem             `json:"item,omitempty"`
	AddItem                       []ExplanationOfBenefitAddItem          `json:"addItem,omitempty"`
	Adjudication                  []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`
	Total                         []ExplanationOfBenefitTotal            `json:"total,omitempty"`
	Payment                       *ExplanationOfBenefitPayment           `json:"payment,omitempty"`
	FormCode                      *CodeableConcept                       `json:"formCode,omitempty"`
	Form                          *Attachment                            `json:"form,omitempty"`
	ProcessNote                   []ExplanationOfBenefitProcessNote      `json:"processNote,omitempty"`
	BenefitPeriod                 *Period                                `json:"benefitPeriod,omitempty"`
	BenefitBalance                []ExplanationOfBenefitBenefitBalance   `json:"benefitBalance,omitempty"`
}

func (r ExplanationOfBenefit) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefit) marshalJSON() jsonExplanationOfBenefit {
	m := jsonExplanationOfBenefit{}
	m.ResourceType = "ExplanationOfBenefit"
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
	m.FundsReserveRequested = r.FundsReserveRequested
	m.FundsReserve = r.FundsReserve
	m.Related = r.Related
	m.Prescription = r.Prescription
	m.OriginalPrescription = r.OriginalPrescription
	m.Payee = r.Payee
	m.Referral = r.Referral
	m.Facility = r.Facility
	m.Claim = r.Claim
	m.ClaimResponse = r.ClaimResponse
	m.Outcome = r.Outcome
	if r.Outcome.Id != nil || r.Outcome.Extension != nil {
		m.OutcomePrimitiveElement = &primitiveElement{Id: r.Outcome.Id, Extension: r.Outcome.Extension}
	}
	m.Disposition = r.Disposition
	if r.Disposition != nil && (r.Disposition.Id != nil || r.Disposition.Extension != nil) {
		m.DispositionPrimitiveElement = &primitiveElement{Id: r.Disposition.Id, Extension: r.Disposition.Extension}
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
	m.PreAuthRefPeriod = r.PreAuthRefPeriod
	m.CareTeam = r.CareTeam
	m.SupportingInfo = r.SupportingInfo
	m.Diagnosis = r.Diagnosis
	m.Procedure = r.Procedure
	m.Precedence = r.Precedence
	if r.Precedence != nil && (r.Precedence.Id != nil || r.Precedence.Extension != nil) {
		m.PrecedencePrimitiveElement = &primitiveElement{Id: r.Precedence.Id, Extension: r.Precedence.Extension}
	}
	m.Insurance = r.Insurance
	m.Accident = r.Accident
	m.Item = r.Item
	m.AddItem = r.AddItem
	m.Adjudication = r.Adjudication
	m.Total = r.Total
	m.Payment = r.Payment
	m.FormCode = r.FormCode
	m.Form = r.Form
	m.ProcessNote = r.ProcessNote
	m.BenefitPeriod = r.BenefitPeriod
	m.BenefitBalance = r.BenefitBalance
	return m
}
func (r *ExplanationOfBenefit) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefit
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefit) unmarshalJSON(m jsonExplanationOfBenefit) error {
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
	r.FundsReserveRequested = m.FundsReserveRequested
	r.FundsReserve = m.FundsReserve
	r.Related = m.Related
	r.Prescription = m.Prescription
	r.OriginalPrescription = m.OriginalPrescription
	r.Payee = m.Payee
	r.Referral = m.Referral
	r.Facility = m.Facility
	r.Claim = m.Claim
	r.ClaimResponse = m.ClaimResponse
	r.Outcome = m.Outcome
	if m.OutcomePrimitiveElement != nil {
		r.Outcome.Id = m.OutcomePrimitiveElement.Id
		r.Outcome.Extension = m.OutcomePrimitiveElement.Extension
	}
	r.Disposition = m.Disposition
	if m.DispositionPrimitiveElement != nil {
		r.Disposition.Id = m.DispositionPrimitiveElement.Id
		r.Disposition.Extension = m.DispositionPrimitiveElement.Extension
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
	r.PreAuthRefPeriod = m.PreAuthRefPeriod
	r.CareTeam = m.CareTeam
	r.SupportingInfo = m.SupportingInfo
	r.Diagnosis = m.Diagnosis
	r.Procedure = m.Procedure
	r.Precedence = m.Precedence
	if m.PrecedencePrimitiveElement != nil {
		r.Precedence.Id = m.PrecedencePrimitiveElement.Id
		r.Precedence.Extension = m.PrecedencePrimitiveElement.Extension
	}
	r.Insurance = m.Insurance
	r.Accident = m.Accident
	r.Item = m.Item
	r.AddItem = m.AddItem
	r.Adjudication = m.Adjudication
	r.Total = m.Total
	r.Payment = m.Payment
	r.FormCode = m.FormCode
	r.Form = m.Form
	r.ProcessNote = m.ProcessNote
	r.BenefitPeriod = m.BenefitPeriod
	r.BenefitBalance = m.BenefitBalance
	return nil
}
func (r ExplanationOfBenefit) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Other claims which are related to this claim such as prior submissions or claims for related services or for the same event.
type ExplanationOfBenefitRelated struct {
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
type jsonExplanationOfBenefitRelated struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Claim             *Reference       `json:"claim,omitempty"`
	Relationship      *CodeableConcept `json:"relationship,omitempty"`
	Reference         *Identifier      `json:"reference,omitempty"`
}

func (r ExplanationOfBenefitRelated) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitRelated) marshalJSON() jsonExplanationOfBenefitRelated {
	m := jsonExplanationOfBenefitRelated{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Claim = r.Claim
	m.Relationship = r.Relationship
	m.Reference = r.Reference
	return m
}
func (r *ExplanationOfBenefitRelated) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitRelated
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitRelated) unmarshalJSON(m jsonExplanationOfBenefitRelated) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Claim = m.Claim
	r.Relationship = m.Relationship
	r.Reference = m.Reference
	return nil
}
func (r ExplanationOfBenefitRelated) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The party to be reimbursed for cost of the products and services according to the terms of the policy.
type ExplanationOfBenefitPayee struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Type of Party to be reimbursed: Subscriber, provider, other.
	Type *CodeableConcept
	// Reference to the individual or organization to whom any payment will be made.
	Party *Reference
}
type jsonExplanationOfBenefitPayee struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Party             *Reference       `json:"party,omitempty"`
}

func (r ExplanationOfBenefitPayee) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitPayee) marshalJSON() jsonExplanationOfBenefitPayee {
	m := jsonExplanationOfBenefitPayee{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Party = r.Party
	return m
}
func (r *ExplanationOfBenefitPayee) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitPayee
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitPayee) unmarshalJSON(m jsonExplanationOfBenefitPayee) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Party = m.Party
	return nil
}
func (r ExplanationOfBenefitPayee) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The members of the team who provided the products and services.
type ExplanationOfBenefitCareTeam struct {
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
type jsonExplanationOfBenefitCareTeam struct {
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

func (r ExplanationOfBenefitCareTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitCareTeam) marshalJSON() jsonExplanationOfBenefitCareTeam {
	m := jsonExplanationOfBenefitCareTeam{}
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
func (r *ExplanationOfBenefitCareTeam) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitCareTeam
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitCareTeam) unmarshalJSON(m jsonExplanationOfBenefitCareTeam) error {
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
func (r ExplanationOfBenefitCareTeam) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
type ExplanationOfBenefitSupportingInfo struct {
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
	Timing isExplanationOfBenefitSupportingInfoTiming
	// Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
	Value isExplanationOfBenefitSupportingInfoValue
	// Provides the reason in the situation where a reason code is required in addition to the content.
	Reason *Coding
}
type isExplanationOfBenefitSupportingInfoTiming interface {
	isExplanationOfBenefitSupportingInfoTiming()
}

func (r Date) isExplanationOfBenefitSupportingInfoTiming()   {}
func (r Period) isExplanationOfBenefitSupportingInfoTiming() {}

type isExplanationOfBenefitSupportingInfoValue interface {
	isExplanationOfBenefitSupportingInfoValue()
}

func (r Boolean) isExplanationOfBenefitSupportingInfoValue()    {}
func (r String) isExplanationOfBenefitSupportingInfoValue()     {}
func (r Quantity) isExplanationOfBenefitSupportingInfoValue()   {}
func (r Attachment) isExplanationOfBenefitSupportingInfoValue() {}
func (r Reference) isExplanationOfBenefitSupportingInfoValue()  {}

type jsonExplanationOfBenefitSupportingInfo struct {
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
	Reason                       *Coding           `json:"reason,omitempty"`
}

func (r ExplanationOfBenefitSupportingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitSupportingInfo) marshalJSON() jsonExplanationOfBenefitSupportingInfo {
	m := jsonExplanationOfBenefitSupportingInfo{}
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
func (r *ExplanationOfBenefitSupportingInfo) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitSupportingInfo
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitSupportingInfo) unmarshalJSON(m jsonExplanationOfBenefitSupportingInfo) error {
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
func (r ExplanationOfBenefitSupportingInfo) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about diagnoses relevant to the claim items.
type ExplanationOfBenefitDiagnosis struct {
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
	Diagnosis isExplanationOfBenefitDiagnosisDiagnosis
	// When the condition was observed or the relative ranking.
	Type []CodeableConcept
	// Indication of whether the diagnosis was present on admission to a facility.
	OnAdmission *CodeableConcept
	// A package billing code or bundle code used to group products and services to a particular health condition (such as heart attack) which is based on a predetermined grouping code system.
	PackageCode *CodeableConcept
}
type isExplanationOfBenefitDiagnosisDiagnosis interface {
	isExplanationOfBenefitDiagnosisDiagnosis()
}

func (r CodeableConcept) isExplanationOfBenefitDiagnosisDiagnosis() {}
func (r Reference) isExplanationOfBenefitDiagnosisDiagnosis()       {}

type jsonExplanationOfBenefitDiagnosis struct {
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

func (r ExplanationOfBenefitDiagnosis) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitDiagnosis) marshalJSON() jsonExplanationOfBenefitDiagnosis {
	m := jsonExplanationOfBenefitDiagnosis{}
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
func (r *ExplanationOfBenefitDiagnosis) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitDiagnosis
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitDiagnosis) unmarshalJSON(m jsonExplanationOfBenefitDiagnosis) error {
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
func (r ExplanationOfBenefitDiagnosis) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Procedures performed on the patient relevant to the billing items with the claim.
type ExplanationOfBenefitProcedure struct {
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
	Procedure isExplanationOfBenefitProcedureProcedure
	// Unique Device Identifiers associated with this line item.
	Udi []Reference
}
type isExplanationOfBenefitProcedureProcedure interface {
	isExplanationOfBenefitProcedureProcedure()
}

func (r CodeableConcept) isExplanationOfBenefitProcedureProcedure() {}
func (r Reference) isExplanationOfBenefitProcedureProcedure()       {}

type jsonExplanationOfBenefitProcedure struct {
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

func (r ExplanationOfBenefitProcedure) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitProcedure) marshalJSON() jsonExplanationOfBenefitProcedure {
	m := jsonExplanationOfBenefitProcedure{}
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
func (r *ExplanationOfBenefitProcedure) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitProcedure
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitProcedure) unmarshalJSON(m jsonExplanationOfBenefitProcedure) error {
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
func (r ExplanationOfBenefitProcedure) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Financial instruments for reimbursement for the health care products and services specified on the claim.
type ExplanationOfBenefitInsurance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A flag to indicate that this Coverage is to be used for adjudication of this claim when set to true.
	Focal Boolean
	// Reference to the insurance card level information contained in the Coverage resource. The coverage issuing insurer will use these details to locate the patient's actual coverage within the insurer's information system.
	Coverage Reference
	// Reference numbers previously provided by the insurer to the provider to be quoted on subsequent claims containing services or products related to the prior authorization.
	PreAuthRef []String
}
type jsonExplanationOfBenefitInsurance struct {
	Id                         *string             `json:"id,omitempty"`
	Extension                  []Extension         `json:"extension,omitempty"`
	ModifierExtension          []Extension         `json:"modifierExtension,omitempty"`
	Focal                      Boolean             `json:"focal,omitempty"`
	FocalPrimitiveElement      *primitiveElement   `json:"_focal,omitempty"`
	Coverage                   Reference           `json:"coverage,omitempty"`
	PreAuthRef                 []String            `json:"preAuthRef,omitempty"`
	PreAuthRefPrimitiveElement []*primitiveElement `json:"_preAuthRef,omitempty"`
}

func (r ExplanationOfBenefitInsurance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitInsurance) marshalJSON() jsonExplanationOfBenefitInsurance {
	m := jsonExplanationOfBenefitInsurance{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Focal = r.Focal
	if r.Focal.Id != nil || r.Focal.Extension != nil {
		m.FocalPrimitiveElement = &primitiveElement{Id: r.Focal.Id, Extension: r.Focal.Extension}
	}
	m.Coverage = r.Coverage
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
	return m
}
func (r *ExplanationOfBenefitInsurance) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitInsurance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitInsurance) unmarshalJSON(m jsonExplanationOfBenefitInsurance) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Focal = m.Focal
	if m.FocalPrimitiveElement != nil {
		r.Focal.Id = m.FocalPrimitiveElement.Id
		r.Focal.Extension = m.FocalPrimitiveElement.Extension
	}
	r.Coverage = m.Coverage
	r.PreAuthRef = m.PreAuthRef
	for i, e := range m.PreAuthRefPrimitiveElement {
		if len(r.PreAuthRef) > i {
			r.PreAuthRef[i].Id = e.Id
			r.PreAuthRef[i].Extension = e.Extension
		} else {
			r.PreAuthRef = append(r.PreAuthRef, String{Id: e.Id, Extension: e.Extension})
		}
	}
	return nil
}
func (r ExplanationOfBenefitInsurance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Details of a accident which resulted in injuries which required the products and services listed in the claim.
type ExplanationOfBenefitAccident struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Date of an accident event  related to the products and services contained in the claim.
	Date *Date
	// The type or context of the accident event for the purposes of selection of potential insurance coverages and determination of coordination between insurers.
	Type *CodeableConcept
	// The physical location of the accident event.
	Location isExplanationOfBenefitAccidentLocation
}
type isExplanationOfBenefitAccidentLocation interface {
	isExplanationOfBenefitAccidentLocation()
}

func (r Address) isExplanationOfBenefitAccidentLocation()   {}
func (r Reference) isExplanationOfBenefitAccidentLocation() {}

type jsonExplanationOfBenefitAccident struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Date                 *Date             `json:"date,omitempty"`
	DatePrimitiveElement *primitiveElement `json:"_date,omitempty"`
	Type                 *CodeableConcept  `json:"type,omitempty"`
	LocationAddress      *Address          `json:"locationAddress,omitempty"`
	LocationReference    *Reference        `json:"locationReference,omitempty"`
}

func (r ExplanationOfBenefitAccident) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitAccident) marshalJSON() jsonExplanationOfBenefitAccident {
	m := jsonExplanationOfBenefitAccident{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
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
func (r *ExplanationOfBenefitAccident) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitAccident
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitAccident) unmarshalJSON(m jsonExplanationOfBenefitAccident) error {
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
func (r ExplanationOfBenefitAccident) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
type ExplanationOfBenefitItem struct {
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
	// Care team members related to this service or product.
	CareTeamSequence []PositiveInt
	// Diagnoses applicable for this service or product.
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
	Serviced isExplanationOfBenefitItemServiced
	// Where the product or service was provided.
	Location isExplanationOfBenefitItemLocation
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
	// A billed item may include goods or services provided in multiple encounters.
	Encounter []Reference
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// If this item is a group then the values here are a summary of the adjudication of the detail items. If this item is a simple product or service then this is the result of the adjudication of this item.
	Adjudication []ExplanationOfBenefitItemAdjudication
	// Second-tier of goods and services.
	Detail []ExplanationOfBenefitItemDetail
}
type isExplanationOfBenefitItemServiced interface {
	isExplanationOfBenefitItemServiced()
}

func (r Date) isExplanationOfBenefitItemServiced()   {}
func (r Period) isExplanationOfBenefitItemServiced() {}

type isExplanationOfBenefitItemLocation interface {
	isExplanationOfBenefitItemLocation()
}

func (r CodeableConcept) isExplanationOfBenefitItemLocation() {}
func (r Address) isExplanationOfBenefitItemLocation()         {}
func (r Reference) isExplanationOfBenefitItemLocation()       {}

type jsonExplanationOfBenefitItem struct {
	Id                                  *string                                `json:"id,omitempty"`
	Extension                           []Extension                            `json:"extension,omitempty"`
	ModifierExtension                   []Extension                            `json:"modifierExtension,omitempty"`
	Sequence                            PositiveInt                            `json:"sequence,omitempty"`
	SequencePrimitiveElement            *primitiveElement                      `json:"_sequence,omitempty"`
	CareTeamSequence                    []PositiveInt                          `json:"careTeamSequence,omitempty"`
	CareTeamSequencePrimitiveElement    []*primitiveElement                    `json:"_careTeamSequence,omitempty"`
	DiagnosisSequence                   []PositiveInt                          `json:"diagnosisSequence,omitempty"`
	DiagnosisSequencePrimitiveElement   []*primitiveElement                    `json:"_diagnosisSequence,omitempty"`
	ProcedureSequence                   []PositiveInt                          `json:"procedureSequence,omitempty"`
	ProcedureSequencePrimitiveElement   []*primitiveElement                    `json:"_procedureSequence,omitempty"`
	InformationSequence                 []PositiveInt                          `json:"informationSequence,omitempty"`
	InformationSequencePrimitiveElement []*primitiveElement                    `json:"_informationSequence,omitempty"`
	Revenue                             *CodeableConcept                       `json:"revenue,omitempty"`
	Category                            *CodeableConcept                       `json:"category,omitempty"`
	ProductOrService                    CodeableConcept                        `json:"productOrService,omitempty"`
	Modifier                            []CodeableConcept                      `json:"modifier,omitempty"`
	ProgramCode                         []CodeableConcept                      `json:"programCode,omitempty"`
	ServicedDate                        *Date                                  `json:"servicedDate,omitempty"`
	ServicedDatePrimitiveElement        *primitiveElement                      `json:"_servicedDate,omitempty"`
	ServicedPeriod                      *Period                                `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept             *CodeableConcept                       `json:"locationCodeableConcept,omitempty"`
	LocationAddress                     *Address                               `json:"locationAddress,omitempty"`
	LocationReference                   *Reference                             `json:"locationReference,omitempty"`
	Quantity                            *Quantity                              `json:"quantity,omitempty"`
	UnitPrice                           *Money                                 `json:"unitPrice,omitempty"`
	Factor                              *Decimal                               `json:"factor,omitempty"`
	FactorPrimitiveElement              *primitiveElement                      `json:"_factor,omitempty"`
	Net                                 *Money                                 `json:"net,omitempty"`
	Udi                                 []Reference                            `json:"udi,omitempty"`
	BodySite                            *CodeableConcept                       `json:"bodySite,omitempty"`
	SubSite                             []CodeableConcept                      `json:"subSite,omitempty"`
	Encounter                           []Reference                            `json:"encounter,omitempty"`
	NoteNumber                          []PositiveInt                          `json:"noteNumber,omitempty"`
	NoteNumberPrimitiveElement          []*primitiveElement                    `json:"_noteNumber,omitempty"`
	Adjudication                        []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`
	Detail                              []ExplanationOfBenefitItemDetail       `json:"detail,omitempty"`
}

func (r ExplanationOfBenefitItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitItem) marshalJSON() jsonExplanationOfBenefitItem {
	m := jsonExplanationOfBenefitItem{}
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
	m.NoteNumber = r.NoteNumber
	anyNoteNumberIdOrExtension := false
	for _, e := range r.NoteNumber {
		if e.Id != nil || e.Extension != nil {
			anyNoteNumberIdOrExtension = true
			break
		}
	}
	if anyNoteNumberIdOrExtension {
		m.NoteNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.NoteNumber))
		for _, e := range r.NoteNumber {
			if e.Id != nil || e.Extension != nil {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, nil)
			}
		}
	}
	m.Adjudication = r.Adjudication
	m.Detail = r.Detail
	return m
}
func (r *ExplanationOfBenefitItem) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitItem) unmarshalJSON(m jsonExplanationOfBenefitItem) error {
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
	r.NoteNumber = m.NoteNumber
	for i, e := range m.NoteNumberPrimitiveElement {
		if len(r.NoteNumber) > i {
			r.NoteNumber[i].Id = e.Id
			r.NoteNumber[i].Extension = e.Extension
		} else {
			r.NoteNumber = append(r.NoteNumber, PositiveInt{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Adjudication = m.Adjudication
	r.Detail = m.Detail
	return nil
}
func (r ExplanationOfBenefitItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// If this item is a group then the values here are a summary of the adjudication of the detail items. If this item is a simple product or service then this is the result of the adjudication of this item.
type ExplanationOfBenefitItemAdjudication struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code to indicate the information type of this adjudication record. Information types may include: the value submitted, maximum values or percentages allowed or payable under the plan, amounts that the patient is responsible for in-aggregate or pertaining to this item, amounts paid by other coverages, and the benefit payable for this item.
	Category CodeableConcept
	// A code supporting the understanding of the adjudication result and explaining variance from expected amount.
	Reason *CodeableConcept
	// Monetary amount associated with the category.
	Amount *Money
	// A non-monetary value associated with the category. Mutually exclusive to the amount element above.
	Value *Decimal
}
type jsonExplanationOfBenefitItemAdjudication struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Category              CodeableConcept   `json:"category,omitempty"`
	Reason                *CodeableConcept  `json:"reason,omitempty"`
	Amount                *Money            `json:"amount,omitempty"`
	Value                 *Decimal          `json:"value,omitempty"`
	ValuePrimitiveElement *primitiveElement `json:"_value,omitempty"`
}

func (r ExplanationOfBenefitItemAdjudication) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitItemAdjudication) marshalJSON() jsonExplanationOfBenefitItemAdjudication {
	m := jsonExplanationOfBenefitItemAdjudication{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Category = r.Category
	m.Reason = r.Reason
	m.Amount = r.Amount
	m.Value = r.Value
	if r.Value != nil && (r.Value.Id != nil || r.Value.Extension != nil) {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	return m
}
func (r *ExplanationOfBenefitItemAdjudication) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitItemAdjudication
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitItemAdjudication) unmarshalJSON(m jsonExplanationOfBenefitItemAdjudication) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Category = m.Category
	r.Reason = m.Reason
	r.Amount = m.Amount
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	return nil
}
func (r ExplanationOfBenefitItemAdjudication) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Second-tier of goods and services.
type ExplanationOfBenefitItemDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
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
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The adjudication results.
	Adjudication []ExplanationOfBenefitItemAdjudication
	// Third-tier of goods and services.
	SubDetail []ExplanationOfBenefitItemDetailSubDetail
}
type jsonExplanationOfBenefitItemDetail struct {
	Id                         *string                                   `json:"id,omitempty"`
	Extension                  []Extension                               `json:"extension,omitempty"`
	ModifierExtension          []Extension                               `json:"modifierExtension,omitempty"`
	Sequence                   PositiveInt                               `json:"sequence,omitempty"`
	SequencePrimitiveElement   *primitiveElement                         `json:"_sequence,omitempty"`
	Revenue                    *CodeableConcept                          `json:"revenue,omitempty"`
	Category                   *CodeableConcept                          `json:"category,omitempty"`
	ProductOrService           CodeableConcept                           `json:"productOrService,omitempty"`
	Modifier                   []CodeableConcept                         `json:"modifier,omitempty"`
	ProgramCode                []CodeableConcept                         `json:"programCode,omitempty"`
	Quantity                   *Quantity                                 `json:"quantity,omitempty"`
	UnitPrice                  *Money                                    `json:"unitPrice,omitempty"`
	Factor                     *Decimal                                  `json:"factor,omitempty"`
	FactorPrimitiveElement     *primitiveElement                         `json:"_factor,omitempty"`
	Net                        *Money                                    `json:"net,omitempty"`
	Udi                        []Reference                               `json:"udi,omitempty"`
	NoteNumber                 []PositiveInt                             `json:"noteNumber,omitempty"`
	NoteNumberPrimitiveElement []*primitiveElement                       `json:"_noteNumber,omitempty"`
	Adjudication               []ExplanationOfBenefitItemAdjudication    `json:"adjudication,omitempty"`
	SubDetail                  []ExplanationOfBenefitItemDetailSubDetail `json:"subDetail,omitempty"`
}

func (r ExplanationOfBenefitItemDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitItemDetail) marshalJSON() jsonExplanationOfBenefitItemDetail {
	m := jsonExplanationOfBenefitItemDetail{}
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
	m.NoteNumber = r.NoteNumber
	anyNoteNumberIdOrExtension := false
	for _, e := range r.NoteNumber {
		if e.Id != nil || e.Extension != nil {
			anyNoteNumberIdOrExtension = true
			break
		}
	}
	if anyNoteNumberIdOrExtension {
		m.NoteNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.NoteNumber))
		for _, e := range r.NoteNumber {
			if e.Id != nil || e.Extension != nil {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, nil)
			}
		}
	}
	m.Adjudication = r.Adjudication
	m.SubDetail = r.SubDetail
	return m
}
func (r *ExplanationOfBenefitItemDetail) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitItemDetail
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitItemDetail) unmarshalJSON(m jsonExplanationOfBenefitItemDetail) error {
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
	r.NoteNumber = m.NoteNumber
	for i, e := range m.NoteNumberPrimitiveElement {
		if len(r.NoteNumber) > i {
			r.NoteNumber[i].Id = e.Id
			r.NoteNumber[i].Extension = e.Extension
		} else {
			r.NoteNumber = append(r.NoteNumber, PositiveInt{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Adjudication = m.Adjudication
	r.SubDetail = m.SubDetail
	return nil
}
func (r ExplanationOfBenefitItemDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Third-tier of goods and services.
type ExplanationOfBenefitItemDetailSubDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
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
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The adjudication results.
	Adjudication []ExplanationOfBenefitItemAdjudication
}
type jsonExplanationOfBenefitItemDetailSubDetail struct {
	Id                         *string                                `json:"id,omitempty"`
	Extension                  []Extension                            `json:"extension,omitempty"`
	ModifierExtension          []Extension                            `json:"modifierExtension,omitempty"`
	Sequence                   PositiveInt                            `json:"sequence,omitempty"`
	SequencePrimitiveElement   *primitiveElement                      `json:"_sequence,omitempty"`
	Revenue                    *CodeableConcept                       `json:"revenue,omitempty"`
	Category                   *CodeableConcept                       `json:"category,omitempty"`
	ProductOrService           CodeableConcept                        `json:"productOrService,omitempty"`
	Modifier                   []CodeableConcept                      `json:"modifier,omitempty"`
	ProgramCode                []CodeableConcept                      `json:"programCode,omitempty"`
	Quantity                   *Quantity                              `json:"quantity,omitempty"`
	UnitPrice                  *Money                                 `json:"unitPrice,omitempty"`
	Factor                     *Decimal                               `json:"factor,omitempty"`
	FactorPrimitiveElement     *primitiveElement                      `json:"_factor,omitempty"`
	Net                        *Money                                 `json:"net,omitempty"`
	Udi                        []Reference                            `json:"udi,omitempty"`
	NoteNumber                 []PositiveInt                          `json:"noteNumber,omitempty"`
	NoteNumberPrimitiveElement []*primitiveElement                    `json:"_noteNumber,omitempty"`
	Adjudication               []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`
}

func (r ExplanationOfBenefitItemDetailSubDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitItemDetailSubDetail) marshalJSON() jsonExplanationOfBenefitItemDetailSubDetail {
	m := jsonExplanationOfBenefitItemDetailSubDetail{}
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
	m.NoteNumber = r.NoteNumber
	anyNoteNumberIdOrExtension := false
	for _, e := range r.NoteNumber {
		if e.Id != nil || e.Extension != nil {
			anyNoteNumberIdOrExtension = true
			break
		}
	}
	if anyNoteNumberIdOrExtension {
		m.NoteNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.NoteNumber))
		for _, e := range r.NoteNumber {
			if e.Id != nil || e.Extension != nil {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, nil)
			}
		}
	}
	m.Adjudication = r.Adjudication
	return m
}
func (r *ExplanationOfBenefitItemDetailSubDetail) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitItemDetailSubDetail
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitItemDetailSubDetail) unmarshalJSON(m jsonExplanationOfBenefitItemDetailSubDetail) error {
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
	r.NoteNumber = m.NoteNumber
	for i, e := range m.NoteNumberPrimitiveElement {
		if len(r.NoteNumber) > i {
			r.NoteNumber[i].Id = e.Id
			r.NoteNumber[i].Extension = e.Extension
		} else {
			r.NoteNumber = append(r.NoteNumber, PositiveInt{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Adjudication = m.Adjudication
	return nil
}
func (r ExplanationOfBenefitItemDetailSubDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The first-tier service adjudications for payor added product or service lines.
type ExplanationOfBenefitAddItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Claim items which this service line is intended to replace.
	ItemSequence []PositiveInt
	// The sequence number of the details within the claim item which this line is intended to replace.
	DetailSequence []PositiveInt
	// The sequence number of the sub-details woithin the details within the claim item which this line is intended to replace.
	SubDetailSequence []PositiveInt
	// The providers who are authorized for the services rendered to the patient.
	Provider []Reference
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []CodeableConcept
	// Identifies the program under which this may be recovered.
	ProgramCode []CodeableConcept
	// The date or dates when the service or product was supplied, performed or completed.
	Serviced isExplanationOfBenefitAddItemServiced
	// Where the product or service was provided.
	Location isExplanationOfBenefitAddItemLocation
	// The number of repetitions of a service or product.
	Quantity *Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *Decimal
	// The quantity times the unit price for an additional service or product or charge.
	Net *Money
	// Physical service site on the patient (limb, tooth, etc.).
	BodySite *CodeableConcept
	// A region or surface of the bodySite, e.g. limb region or tooth surface(s).
	SubSite []CodeableConcept
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The adjudication results.
	Adjudication []ExplanationOfBenefitItemAdjudication
	// The second-tier service adjudications for payor added services.
	Detail []ExplanationOfBenefitAddItemDetail
}
type isExplanationOfBenefitAddItemServiced interface {
	isExplanationOfBenefitAddItemServiced()
}

func (r Date) isExplanationOfBenefitAddItemServiced()   {}
func (r Period) isExplanationOfBenefitAddItemServiced() {}

type isExplanationOfBenefitAddItemLocation interface {
	isExplanationOfBenefitAddItemLocation()
}

func (r CodeableConcept) isExplanationOfBenefitAddItemLocation() {}
func (r Address) isExplanationOfBenefitAddItemLocation()         {}
func (r Reference) isExplanationOfBenefitAddItemLocation()       {}

type jsonExplanationOfBenefitAddItem struct {
	Id                                *string                                `json:"id,omitempty"`
	Extension                         []Extension                            `json:"extension,omitempty"`
	ModifierExtension                 []Extension                            `json:"modifierExtension,omitempty"`
	ItemSequence                      []PositiveInt                          `json:"itemSequence,omitempty"`
	ItemSequencePrimitiveElement      []*primitiveElement                    `json:"_itemSequence,omitempty"`
	DetailSequence                    []PositiveInt                          `json:"detailSequence,omitempty"`
	DetailSequencePrimitiveElement    []*primitiveElement                    `json:"_detailSequence,omitempty"`
	SubDetailSequence                 []PositiveInt                          `json:"subDetailSequence,omitempty"`
	SubDetailSequencePrimitiveElement []*primitiveElement                    `json:"_subDetailSequence,omitempty"`
	Provider                          []Reference                            `json:"provider,omitempty"`
	ProductOrService                  CodeableConcept                        `json:"productOrService,omitempty"`
	Modifier                          []CodeableConcept                      `json:"modifier,omitempty"`
	ProgramCode                       []CodeableConcept                      `json:"programCode,omitempty"`
	ServicedDate                      *Date                                  `json:"servicedDate,omitempty"`
	ServicedDatePrimitiveElement      *primitiveElement                      `json:"_servicedDate,omitempty"`
	ServicedPeriod                    *Period                                `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept           *CodeableConcept                       `json:"locationCodeableConcept,omitempty"`
	LocationAddress                   *Address                               `json:"locationAddress,omitempty"`
	LocationReference                 *Reference                             `json:"locationReference,omitempty"`
	Quantity                          *Quantity                              `json:"quantity,omitempty"`
	UnitPrice                         *Money                                 `json:"unitPrice,omitempty"`
	Factor                            *Decimal                               `json:"factor,omitempty"`
	FactorPrimitiveElement            *primitiveElement                      `json:"_factor,omitempty"`
	Net                               *Money                                 `json:"net,omitempty"`
	BodySite                          *CodeableConcept                       `json:"bodySite,omitempty"`
	SubSite                           []CodeableConcept                      `json:"subSite,omitempty"`
	NoteNumber                        []PositiveInt                          `json:"noteNumber,omitempty"`
	NoteNumberPrimitiveElement        []*primitiveElement                    `json:"_noteNumber,omitempty"`
	Adjudication                      []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`
	Detail                            []ExplanationOfBenefitAddItemDetail    `json:"detail,omitempty"`
}

func (r ExplanationOfBenefitAddItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitAddItem) marshalJSON() jsonExplanationOfBenefitAddItem {
	m := jsonExplanationOfBenefitAddItem{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.ItemSequence = r.ItemSequence
	anyItemSequenceIdOrExtension := false
	for _, e := range r.ItemSequence {
		if e.Id != nil || e.Extension != nil {
			anyItemSequenceIdOrExtension = true
			break
		}
	}
	if anyItemSequenceIdOrExtension {
		m.ItemSequencePrimitiveElement = make([]*primitiveElement, 0, len(r.ItemSequence))
		for _, e := range r.ItemSequence {
			if e.Id != nil || e.Extension != nil {
				m.ItemSequencePrimitiveElement = append(m.ItemSequencePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ItemSequencePrimitiveElement = append(m.ItemSequencePrimitiveElement, nil)
			}
		}
	}
	m.DetailSequence = r.DetailSequence
	anyDetailSequenceIdOrExtension := false
	for _, e := range r.DetailSequence {
		if e.Id != nil || e.Extension != nil {
			anyDetailSequenceIdOrExtension = true
			break
		}
	}
	if anyDetailSequenceIdOrExtension {
		m.DetailSequencePrimitiveElement = make([]*primitiveElement, 0, len(r.DetailSequence))
		for _, e := range r.DetailSequence {
			if e.Id != nil || e.Extension != nil {
				m.DetailSequencePrimitiveElement = append(m.DetailSequencePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.DetailSequencePrimitiveElement = append(m.DetailSequencePrimitiveElement, nil)
			}
		}
	}
	m.SubDetailSequence = r.SubDetailSequence
	anySubDetailSequenceIdOrExtension := false
	for _, e := range r.SubDetailSequence {
		if e.Id != nil || e.Extension != nil {
			anySubDetailSequenceIdOrExtension = true
			break
		}
	}
	if anySubDetailSequenceIdOrExtension {
		m.SubDetailSequencePrimitiveElement = make([]*primitiveElement, 0, len(r.SubDetailSequence))
		for _, e := range r.SubDetailSequence {
			if e.Id != nil || e.Extension != nil {
				m.SubDetailSequencePrimitiveElement = append(m.SubDetailSequencePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SubDetailSequencePrimitiveElement = append(m.SubDetailSequencePrimitiveElement, nil)
			}
		}
	}
	m.Provider = r.Provider
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
	m.BodySite = r.BodySite
	m.SubSite = r.SubSite
	m.NoteNumber = r.NoteNumber
	anyNoteNumberIdOrExtension := false
	for _, e := range r.NoteNumber {
		if e.Id != nil || e.Extension != nil {
			anyNoteNumberIdOrExtension = true
			break
		}
	}
	if anyNoteNumberIdOrExtension {
		m.NoteNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.NoteNumber))
		for _, e := range r.NoteNumber {
			if e.Id != nil || e.Extension != nil {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, nil)
			}
		}
	}
	m.Adjudication = r.Adjudication
	m.Detail = r.Detail
	return m
}
func (r *ExplanationOfBenefitAddItem) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitAddItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitAddItem) unmarshalJSON(m jsonExplanationOfBenefitAddItem) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ItemSequence = m.ItemSequence
	for i, e := range m.ItemSequencePrimitiveElement {
		if len(r.ItemSequence) > i {
			r.ItemSequence[i].Id = e.Id
			r.ItemSequence[i].Extension = e.Extension
		} else {
			r.ItemSequence = append(r.ItemSequence, PositiveInt{Id: e.Id, Extension: e.Extension})
		}
	}
	r.DetailSequence = m.DetailSequence
	for i, e := range m.DetailSequencePrimitiveElement {
		if len(r.DetailSequence) > i {
			r.DetailSequence[i].Id = e.Id
			r.DetailSequence[i].Extension = e.Extension
		} else {
			r.DetailSequence = append(r.DetailSequence, PositiveInt{Id: e.Id, Extension: e.Extension})
		}
	}
	r.SubDetailSequence = m.SubDetailSequence
	for i, e := range m.SubDetailSequencePrimitiveElement {
		if len(r.SubDetailSequence) > i {
			r.SubDetailSequence[i].Id = e.Id
			r.SubDetailSequence[i].Extension = e.Extension
		} else {
			r.SubDetailSequence = append(r.SubDetailSequence, PositiveInt{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Provider = m.Provider
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
	r.BodySite = m.BodySite
	r.SubSite = m.SubSite
	r.NoteNumber = m.NoteNumber
	for i, e := range m.NoteNumberPrimitiveElement {
		if len(r.NoteNumber) > i {
			r.NoteNumber[i].Id = e.Id
			r.NoteNumber[i].Extension = e.Extension
		} else {
			r.NoteNumber = append(r.NoteNumber, PositiveInt{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Adjudication = m.Adjudication
	r.Detail = m.Detail
	return nil
}
func (r ExplanationOfBenefitAddItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The second-tier service adjudications for payor added services.
type ExplanationOfBenefitAddItemDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *Decimal
	// The quantity times the unit price for an additional service or product or charge.
	Net *Money
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The adjudication results.
	Adjudication []ExplanationOfBenefitItemAdjudication
	// The third-tier service adjudications for payor added services.
	SubDetail []ExplanationOfBenefitAddItemDetailSubDetail
}
type jsonExplanationOfBenefitAddItemDetail struct {
	Id                         *string                                      `json:"id,omitempty"`
	Extension                  []Extension                                  `json:"extension,omitempty"`
	ModifierExtension          []Extension                                  `json:"modifierExtension,omitempty"`
	ProductOrService           CodeableConcept                              `json:"productOrService,omitempty"`
	Modifier                   []CodeableConcept                            `json:"modifier,omitempty"`
	Quantity                   *Quantity                                    `json:"quantity,omitempty"`
	UnitPrice                  *Money                                       `json:"unitPrice,omitempty"`
	Factor                     *Decimal                                     `json:"factor,omitempty"`
	FactorPrimitiveElement     *primitiveElement                            `json:"_factor,omitempty"`
	Net                        *Money                                       `json:"net,omitempty"`
	NoteNumber                 []PositiveInt                                `json:"noteNumber,omitempty"`
	NoteNumberPrimitiveElement []*primitiveElement                          `json:"_noteNumber,omitempty"`
	Adjudication               []ExplanationOfBenefitItemAdjudication       `json:"adjudication,omitempty"`
	SubDetail                  []ExplanationOfBenefitAddItemDetailSubDetail `json:"subDetail,omitempty"`
}

func (r ExplanationOfBenefitAddItemDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitAddItemDetail) marshalJSON() jsonExplanationOfBenefitAddItemDetail {
	m := jsonExplanationOfBenefitAddItemDetail{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.ProductOrService = r.ProductOrService
	m.Modifier = r.Modifier
	m.Quantity = r.Quantity
	m.UnitPrice = r.UnitPrice
	m.Factor = r.Factor
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		m.FactorPrimitiveElement = &primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
	}
	m.Net = r.Net
	m.NoteNumber = r.NoteNumber
	anyNoteNumberIdOrExtension := false
	for _, e := range r.NoteNumber {
		if e.Id != nil || e.Extension != nil {
			anyNoteNumberIdOrExtension = true
			break
		}
	}
	if anyNoteNumberIdOrExtension {
		m.NoteNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.NoteNumber))
		for _, e := range r.NoteNumber {
			if e.Id != nil || e.Extension != nil {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, nil)
			}
		}
	}
	m.Adjudication = r.Adjudication
	m.SubDetail = r.SubDetail
	return m
}
func (r *ExplanationOfBenefitAddItemDetail) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitAddItemDetail
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitAddItemDetail) unmarshalJSON(m jsonExplanationOfBenefitAddItemDetail) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ProductOrService = m.ProductOrService
	r.Modifier = m.Modifier
	r.Quantity = m.Quantity
	r.UnitPrice = m.UnitPrice
	r.Factor = m.Factor
	if m.FactorPrimitiveElement != nil {
		r.Factor.Id = m.FactorPrimitiveElement.Id
		r.Factor.Extension = m.FactorPrimitiveElement.Extension
	}
	r.Net = m.Net
	r.NoteNumber = m.NoteNumber
	for i, e := range m.NoteNumberPrimitiveElement {
		if len(r.NoteNumber) > i {
			r.NoteNumber[i].Id = e.Id
			r.NoteNumber[i].Extension = e.Extension
		} else {
			r.NoteNumber = append(r.NoteNumber, PositiveInt{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Adjudication = m.Adjudication
	r.SubDetail = m.SubDetail
	return nil
}
func (r ExplanationOfBenefitAddItemDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The third-tier service adjudications for payor added services.
type ExplanationOfBenefitAddItemDetailSubDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *Decimal
	// The quantity times the unit price for an additional service or product or charge.
	Net *Money
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The adjudication results.
	Adjudication []ExplanationOfBenefitItemAdjudication
}
type jsonExplanationOfBenefitAddItemDetailSubDetail struct {
	Id                         *string                                `json:"id,omitempty"`
	Extension                  []Extension                            `json:"extension,omitempty"`
	ModifierExtension          []Extension                            `json:"modifierExtension,omitempty"`
	ProductOrService           CodeableConcept                        `json:"productOrService,omitempty"`
	Modifier                   []CodeableConcept                      `json:"modifier,omitempty"`
	Quantity                   *Quantity                              `json:"quantity,omitempty"`
	UnitPrice                  *Money                                 `json:"unitPrice,omitempty"`
	Factor                     *Decimal                               `json:"factor,omitempty"`
	FactorPrimitiveElement     *primitiveElement                      `json:"_factor,omitempty"`
	Net                        *Money                                 `json:"net,omitempty"`
	NoteNumber                 []PositiveInt                          `json:"noteNumber,omitempty"`
	NoteNumberPrimitiveElement []*primitiveElement                    `json:"_noteNumber,omitempty"`
	Adjudication               []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`
}

func (r ExplanationOfBenefitAddItemDetailSubDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitAddItemDetailSubDetail) marshalJSON() jsonExplanationOfBenefitAddItemDetailSubDetail {
	m := jsonExplanationOfBenefitAddItemDetailSubDetail{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.ProductOrService = r.ProductOrService
	m.Modifier = r.Modifier
	m.Quantity = r.Quantity
	m.UnitPrice = r.UnitPrice
	m.Factor = r.Factor
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		m.FactorPrimitiveElement = &primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
	}
	m.Net = r.Net
	m.NoteNumber = r.NoteNumber
	anyNoteNumberIdOrExtension := false
	for _, e := range r.NoteNumber {
		if e.Id != nil || e.Extension != nil {
			anyNoteNumberIdOrExtension = true
			break
		}
	}
	if anyNoteNumberIdOrExtension {
		m.NoteNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.NoteNumber))
		for _, e := range r.NoteNumber {
			if e.Id != nil || e.Extension != nil {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, nil)
			}
		}
	}
	m.Adjudication = r.Adjudication
	return m
}
func (r *ExplanationOfBenefitAddItemDetailSubDetail) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitAddItemDetailSubDetail
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitAddItemDetailSubDetail) unmarshalJSON(m jsonExplanationOfBenefitAddItemDetailSubDetail) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ProductOrService = m.ProductOrService
	r.Modifier = m.Modifier
	r.Quantity = m.Quantity
	r.UnitPrice = m.UnitPrice
	r.Factor = m.Factor
	if m.FactorPrimitiveElement != nil {
		r.Factor.Id = m.FactorPrimitiveElement.Id
		r.Factor.Extension = m.FactorPrimitiveElement.Extension
	}
	r.Net = m.Net
	r.NoteNumber = m.NoteNumber
	for i, e := range m.NoteNumberPrimitiveElement {
		if len(r.NoteNumber) > i {
			r.NoteNumber[i].Id = e.Id
			r.NoteNumber[i].Extension = e.Extension
		} else {
			r.NoteNumber = append(r.NoteNumber, PositiveInt{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Adjudication = m.Adjudication
	return nil
}
func (r ExplanationOfBenefitAddItemDetailSubDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Categorized monetary totals for the adjudication.
type ExplanationOfBenefitTotal struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code to indicate the information type of this adjudication record. Information types may include: the value submitted, maximum values or percentages allowed or payable under the plan, amounts that the patient is responsible for in aggregate or pertaining to this item, amounts paid by other coverages, and the benefit payable for this item.
	Category CodeableConcept
	// Monetary total amount associated with the category.
	Amount Money
}
type jsonExplanationOfBenefitTotal struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Category          CodeableConcept `json:"category,omitempty"`
	Amount            Money           `json:"amount,omitempty"`
}

func (r ExplanationOfBenefitTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitTotal) marshalJSON() jsonExplanationOfBenefitTotal {
	m := jsonExplanationOfBenefitTotal{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Category = r.Category
	m.Amount = r.Amount
	return m
}
func (r *ExplanationOfBenefitTotal) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitTotal
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitTotal) unmarshalJSON(m jsonExplanationOfBenefitTotal) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Category = m.Category
	r.Amount = m.Amount
	return nil
}
func (r ExplanationOfBenefitTotal) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Payment details for the adjudication of the claim.
type ExplanationOfBenefitPayment struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Whether this represents partial or complete payment of the benefits payable.
	Type *CodeableConcept
	// Total amount of all adjustments to this payment included in this transaction which are not related to this claim's adjudication.
	Adjustment *Money
	// Reason for the payment adjustment.
	AdjustmentReason *CodeableConcept
	// Estimated date the payment will be issued or the actual issue date of payment.
	Date *Date
	// Benefits payable less any payment adjustment.
	Amount *Money
	// Issuer's unique identifier for the payment instrument.
	Identifier *Identifier
}
type jsonExplanationOfBenefitPayment struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Type                 *CodeableConcept  `json:"type,omitempty"`
	Adjustment           *Money            `json:"adjustment,omitempty"`
	AdjustmentReason     *CodeableConcept  `json:"adjustmentReason,omitempty"`
	Date                 *Date             `json:"date,omitempty"`
	DatePrimitiveElement *primitiveElement `json:"_date,omitempty"`
	Amount               *Money            `json:"amount,omitempty"`
	Identifier           *Identifier       `json:"identifier,omitempty"`
}

func (r ExplanationOfBenefitPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitPayment) marshalJSON() jsonExplanationOfBenefitPayment {
	m := jsonExplanationOfBenefitPayment{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Adjustment = r.Adjustment
	m.AdjustmentReason = r.AdjustmentReason
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Amount = r.Amount
	m.Identifier = r.Identifier
	return m
}
func (r *ExplanationOfBenefitPayment) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitPayment
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitPayment) unmarshalJSON(m jsonExplanationOfBenefitPayment) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Adjustment = m.Adjustment
	r.AdjustmentReason = m.AdjustmentReason
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Amount = m.Amount
	r.Identifier = m.Identifier
	return nil
}
func (r ExplanationOfBenefitPayment) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A note that describes or explains adjudication results in a human readable form.
type ExplanationOfBenefitProcessNote struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely identify a note entry.
	Number *PositiveInt
	// The business purpose of the note text.
	Type *Code
	// The explanation or description associated with the processing.
	Text *String
	// A code to define the language used in the text of the note.
	Language *CodeableConcept
}
type jsonExplanationOfBenefitProcessNote struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Number                 *PositiveInt      `json:"number,omitempty"`
	NumberPrimitiveElement *primitiveElement `json:"_number,omitempty"`
	Type                   *Code             `json:"type,omitempty"`
	TypePrimitiveElement   *primitiveElement `json:"_type,omitempty"`
	Text                   *String           `json:"text,omitempty"`
	TextPrimitiveElement   *primitiveElement `json:"_text,omitempty"`
	Language               *CodeableConcept  `json:"language,omitempty"`
}

func (r ExplanationOfBenefitProcessNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitProcessNote) marshalJSON() jsonExplanationOfBenefitProcessNote {
	m := jsonExplanationOfBenefitProcessNote{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Number = r.Number
	if r.Number != nil && (r.Number.Id != nil || r.Number.Extension != nil) {
		m.NumberPrimitiveElement = &primitiveElement{Id: r.Number.Id, Extension: r.Number.Extension}
	}
	m.Type = r.Type
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Text = r.Text
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	m.Language = r.Language
	return m
}
func (r *ExplanationOfBenefitProcessNote) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitProcessNote
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitProcessNote) unmarshalJSON(m jsonExplanationOfBenefitProcessNote) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Number = m.Number
	if m.NumberPrimitiveElement != nil {
		r.Number.Id = m.NumberPrimitiveElement.Id
		r.Number.Extension = m.NumberPrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	r.Language = m.Language
	return nil
}
func (r ExplanationOfBenefitProcessNote) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Balance by Benefit Category.
type ExplanationOfBenefitBenefitBalance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Code to identify the general type of benefits under which products and services are provided.
	Category CodeableConcept
	// True if the indicated class of service is excluded from the plan, missing or False indicates the product or service is included in the coverage.
	Excluded *Boolean
	// A short name or tag for the benefit.
	Name *String
	// A richer description of the benefit or services covered.
	Description *String
	// Is a flag to indicate whether the benefits refer to in-network providers or out-of-network providers.
	Network *CodeableConcept
	// Indicates if the benefits apply to an individual or to the family.
	Unit *CodeableConcept
	// The term or period of the values such as 'maximum lifetime benefit' or 'maximum annual visits'.
	Term *CodeableConcept
	// Benefits Used to date.
	Financial []ExplanationOfBenefitBenefitBalanceFinancial
}
type jsonExplanationOfBenefitBenefitBalance struct {
	Id                          *string                                       `json:"id,omitempty"`
	Extension                   []Extension                                   `json:"extension,omitempty"`
	ModifierExtension           []Extension                                   `json:"modifierExtension,omitempty"`
	Category                    CodeableConcept                               `json:"category,omitempty"`
	Excluded                    *Boolean                                      `json:"excluded,omitempty"`
	ExcludedPrimitiveElement    *primitiveElement                             `json:"_excluded,omitempty"`
	Name                        *String                                       `json:"name,omitempty"`
	NamePrimitiveElement        *primitiveElement                             `json:"_name,omitempty"`
	Description                 *String                                       `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement                             `json:"_description,omitempty"`
	Network                     *CodeableConcept                              `json:"network,omitempty"`
	Unit                        *CodeableConcept                              `json:"unit,omitempty"`
	Term                        *CodeableConcept                              `json:"term,omitempty"`
	Financial                   []ExplanationOfBenefitBenefitBalanceFinancial `json:"financial,omitempty"`
}

func (r ExplanationOfBenefitBenefitBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitBenefitBalance) marshalJSON() jsonExplanationOfBenefitBenefitBalance {
	m := jsonExplanationOfBenefitBenefitBalance{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Category = r.Category
	m.Excluded = r.Excluded
	if r.Excluded != nil && (r.Excluded.Id != nil || r.Excluded.Extension != nil) {
		m.ExcludedPrimitiveElement = &primitiveElement{Id: r.Excluded.Id, Extension: r.Excluded.Extension}
	}
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Network = r.Network
	m.Unit = r.Unit
	m.Term = r.Term
	m.Financial = r.Financial
	return m
}
func (r *ExplanationOfBenefitBenefitBalance) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitBenefitBalance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitBenefitBalance) unmarshalJSON(m jsonExplanationOfBenefitBenefitBalance) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Category = m.Category
	r.Excluded = m.Excluded
	if m.ExcludedPrimitiveElement != nil {
		r.Excluded.Id = m.ExcludedPrimitiveElement.Id
		r.Excluded.Extension = m.ExcludedPrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Network = m.Network
	r.Unit = m.Unit
	r.Term = m.Term
	r.Financial = m.Financial
	return nil
}
func (r ExplanationOfBenefitBenefitBalance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Benefits Used to date.
type ExplanationOfBenefitBenefitBalanceFinancial struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Classification of benefit being provided.
	Type CodeableConcept
	// The quantity of the benefit which is permitted under the coverage.
	Allowed isExplanationOfBenefitBenefitBalanceFinancialAllowed
	// The quantity of the benefit which have been consumed to date.
	Used isExplanationOfBenefitBenefitBalanceFinancialUsed
}
type isExplanationOfBenefitBenefitBalanceFinancialAllowed interface {
	isExplanationOfBenefitBenefitBalanceFinancialAllowed()
}

func (r UnsignedInt) isExplanationOfBenefitBenefitBalanceFinancialAllowed() {}
func (r String) isExplanationOfBenefitBenefitBalanceFinancialAllowed()      {}
func (r Money) isExplanationOfBenefitBenefitBalanceFinancialAllowed()       {}

type isExplanationOfBenefitBenefitBalanceFinancialUsed interface {
	isExplanationOfBenefitBenefitBalanceFinancialUsed()
}

func (r UnsignedInt) isExplanationOfBenefitBenefitBalanceFinancialUsed() {}
func (r Money) isExplanationOfBenefitBenefitBalanceFinancialUsed()       {}

type jsonExplanationOfBenefitBenefitBalanceFinancial struct {
	Id                                 *string           `json:"id,omitempty"`
	Extension                          []Extension       `json:"extension,omitempty"`
	ModifierExtension                  []Extension       `json:"modifierExtension,omitempty"`
	Type                               CodeableConcept   `json:"type,omitempty"`
	AllowedUnsignedInt                 *UnsignedInt      `json:"allowedUnsignedInt,omitempty"`
	AllowedUnsignedIntPrimitiveElement *primitiveElement `json:"_allowedUnsignedInt,omitempty"`
	AllowedString                      *String           `json:"allowedString,omitempty"`
	AllowedStringPrimitiveElement      *primitiveElement `json:"_allowedString,omitempty"`
	AllowedMoney                       *Money            `json:"allowedMoney,omitempty"`
	UsedUnsignedInt                    *UnsignedInt      `json:"usedUnsignedInt,omitempty"`
	UsedUnsignedIntPrimitiveElement    *primitiveElement `json:"_usedUnsignedInt,omitempty"`
	UsedMoney                          *Money            `json:"usedMoney,omitempty"`
}

func (r ExplanationOfBenefitBenefitBalanceFinancial) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExplanationOfBenefitBenefitBalanceFinancial) marshalJSON() jsonExplanationOfBenefitBenefitBalanceFinancial {
	m := jsonExplanationOfBenefitBenefitBalanceFinancial{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	switch v := r.Allowed.(type) {
	case UnsignedInt:
		m.AllowedUnsignedInt = &v
		if v.Id != nil || v.Extension != nil {
			m.AllowedUnsignedIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *UnsignedInt:
		m.AllowedUnsignedInt = v
		if v.Id != nil || v.Extension != nil {
			m.AllowedUnsignedIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case String:
		m.AllowedString = &v
		if v.Id != nil || v.Extension != nil {
			m.AllowedStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.AllowedString = v
		if v.Id != nil || v.Extension != nil {
			m.AllowedStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Money:
		m.AllowedMoney = &v
	case *Money:
		m.AllowedMoney = v
	}
	switch v := r.Used.(type) {
	case UnsignedInt:
		m.UsedUnsignedInt = &v
		if v.Id != nil || v.Extension != nil {
			m.UsedUnsignedIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *UnsignedInt:
		m.UsedUnsignedInt = v
		if v.Id != nil || v.Extension != nil {
			m.UsedUnsignedIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Money:
		m.UsedMoney = &v
	case *Money:
		m.UsedMoney = v
	}
	return m
}
func (r *ExplanationOfBenefitBenefitBalanceFinancial) UnmarshalJSON(b []byte) error {
	var m jsonExplanationOfBenefitBenefitBalanceFinancial
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExplanationOfBenefitBenefitBalanceFinancial) unmarshalJSON(m jsonExplanationOfBenefitBenefitBalanceFinancial) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.AllowedUnsignedInt != nil || m.AllowedUnsignedIntPrimitiveElement != nil {
		if r.Allowed != nil {
			return fmt.Errorf("multiple values for field \"Allowed\"")
		}
		v := m.AllowedUnsignedInt
		if m.AllowedUnsignedIntPrimitiveElement != nil {
			if v == nil {
				v = &UnsignedInt{}
			}
			v.Id = m.AllowedUnsignedIntPrimitiveElement.Id
			v.Extension = m.AllowedUnsignedIntPrimitiveElement.Extension
		}
		r.Allowed = v
	}
	if m.AllowedString != nil || m.AllowedStringPrimitiveElement != nil {
		if r.Allowed != nil {
			return fmt.Errorf("multiple values for field \"Allowed\"")
		}
		v := m.AllowedString
		if m.AllowedStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.AllowedStringPrimitiveElement.Id
			v.Extension = m.AllowedStringPrimitiveElement.Extension
		}
		r.Allowed = v
	}
	if m.AllowedMoney != nil {
		if r.Allowed != nil {
			return fmt.Errorf("multiple values for field \"Allowed\"")
		}
		v := m.AllowedMoney
		r.Allowed = v
	}
	if m.UsedUnsignedInt != nil || m.UsedUnsignedIntPrimitiveElement != nil {
		if r.Used != nil {
			return fmt.Errorf("multiple values for field \"Used\"")
		}
		v := m.UsedUnsignedInt
		if m.UsedUnsignedIntPrimitiveElement != nil {
			if v == nil {
				v = &UnsignedInt{}
			}
			v.Id = m.UsedUnsignedIntPrimitiveElement.Id
			v.Extension = m.UsedUnsignedIntPrimitiveElement.Extension
		}
		r.Used = v
	}
	if m.UsedMoney != nil {
		if r.Used != nil {
			return fmt.Errorf("multiple values for field \"Used\"")
		}
		v := m.UsedMoney
		r.Used = v
	}
	return nil
}
func (r ExplanationOfBenefitBenefitBalanceFinancial) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
