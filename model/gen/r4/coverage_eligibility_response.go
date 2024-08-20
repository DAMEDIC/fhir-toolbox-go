package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// This resource provides eligibility and plan details from the processing of an CoverageEligibilityRequest resource.
type CoverageEligibilityResponse struct {
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
	// A unique identifier assigned to this coverage eligiblity request.
	Identifier []Identifier
	// The status of the resource instance.
	Status Code
	// Code to specify whether requesting: prior authorization requirements for some service categories or billing codes; benefits for coverages specified or discovered; discovery and return of coverages for the patient; and/or validation that the specified coverage is in-force at the date/period specified or 'now' if not specified.
	Purpose []Code
	// The party who is the beneficiary of the supplied coverage and for whom eligibility is sought.
	Patient Reference
	// The date or dates when the enclosed suite of services were performed or completed.
	Serviced isCoverageEligibilityResponseServiced
	// The date this resource was created.
	Created DateTime
	// The provider which is responsible for the request.
	Requestor *Reference
	// Reference to the original request resource.
	Request Reference
	// The outcome of the request processing.
	Outcome Code
	// A human readable description of the status of the adjudication.
	Disposition *String
	// The Insurer who issued the coverage in question and is the author of the response.
	Insurer Reference
	// Financial instruments for reimbursement for the health care products and services.
	Insurance []CoverageEligibilityResponseInsurance
	// A reference from the Insurer to which these services pertain to be used on further communication and as proof that the request occurred.
	PreAuthRef *String
	// A code for the form to be used for printing the content.
	Form *CodeableConcept
	// Errors encountered during the processing of the request.
	Error []CoverageEligibilityResponseError
}

func (r CoverageEligibilityResponse) ResourceType() string {
	return "CoverageEligibilityResponse"
}
func (r CoverageEligibilityResponse) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isCoverageEligibilityResponseServiced interface {
	isCoverageEligibilityResponseServiced()
}

func (r Date) isCoverageEligibilityResponseServiced()   {}
func (r Period) isCoverageEligibilityResponseServiced() {}

type jsonCoverageEligibilityResponse struct {
	ResourceType                  string                                 `json:"resourceType"`
	Id                            *Id                                    `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement                      `json:"_id,omitempty"`
	Meta                          *Meta                                  `json:"meta,omitempty"`
	ImplicitRules                 *Uri                                   `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement                      `json:"_implicitRules,omitempty"`
	Language                      *Code                                  `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement                      `json:"_language,omitempty"`
	Text                          *Narrative                             `json:"text,omitempty"`
	Contained                     []ContainedResource                    `json:"contained,omitempty"`
	Extension                     []Extension                            `json:"extension,omitempty"`
	ModifierExtension             []Extension                            `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier                           `json:"identifier,omitempty"`
	Status                        Code                                   `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement                      `json:"_status,omitempty"`
	Purpose                       []Code                                 `json:"purpose,omitempty"`
	PurposePrimitiveElement       []*primitiveElement                    `json:"_purpose,omitempty"`
	Patient                       Reference                              `json:"patient,omitempty"`
	ServicedDate                  *Date                                  `json:"servicedDate,omitempty"`
	ServicedDatePrimitiveElement  *primitiveElement                      `json:"_servicedDate,omitempty"`
	ServicedPeriod                *Period                                `json:"servicedPeriod,omitempty"`
	Created                       DateTime                               `json:"created,omitempty"`
	CreatedPrimitiveElement       *primitiveElement                      `json:"_created,omitempty"`
	Requestor                     *Reference                             `json:"requestor,omitempty"`
	Request                       Reference                              `json:"request,omitempty"`
	Outcome                       Code                                   `json:"outcome,omitempty"`
	OutcomePrimitiveElement       *primitiveElement                      `json:"_outcome,omitempty"`
	Disposition                   *String                                `json:"disposition,omitempty"`
	DispositionPrimitiveElement   *primitiveElement                      `json:"_disposition,omitempty"`
	Insurer                       Reference                              `json:"insurer,omitempty"`
	Insurance                     []CoverageEligibilityResponseInsurance `json:"insurance,omitempty"`
	PreAuthRef                    *String                                `json:"preAuthRef,omitempty"`
	PreAuthRefPrimitiveElement    *primitiveElement                      `json:"_preAuthRef,omitempty"`
	Form                          *CodeableConcept                       `json:"form,omitempty"`
	Error                         []CoverageEligibilityResponseError     `json:"error,omitempty"`
}

func (r CoverageEligibilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CoverageEligibilityResponse) marshalJSON() jsonCoverageEligibilityResponse {
	m := jsonCoverageEligibilityResponse{}
	m.ResourceType = "CoverageEligibilityResponse"
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
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Purpose = r.Purpose
	anyPurposeIdOrExtension := false
	for _, e := range r.Purpose {
		if e.Id != nil || e.Extension != nil {
			anyPurposeIdOrExtension = true
			break
		}
	}
	if anyPurposeIdOrExtension {
		m.PurposePrimitiveElement = make([]*primitiveElement, 0, len(r.Purpose))
		for _, e := range r.Purpose {
			if e.Id != nil || e.Extension != nil {
				m.PurposePrimitiveElement = append(m.PurposePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.PurposePrimitiveElement = append(m.PurposePrimitiveElement, nil)
			}
		}
	}
	m.Patient = r.Patient
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
	m.Created = r.Created
	if r.Created.Id != nil || r.Created.Extension != nil {
		m.CreatedPrimitiveElement = &primitiveElement{Id: r.Created.Id, Extension: r.Created.Extension}
	}
	m.Requestor = r.Requestor
	m.Request = r.Request
	m.Outcome = r.Outcome
	if r.Outcome.Id != nil || r.Outcome.Extension != nil {
		m.OutcomePrimitiveElement = &primitiveElement{Id: r.Outcome.Id, Extension: r.Outcome.Extension}
	}
	m.Disposition = r.Disposition
	if r.Disposition != nil && (r.Disposition.Id != nil || r.Disposition.Extension != nil) {
		m.DispositionPrimitiveElement = &primitiveElement{Id: r.Disposition.Id, Extension: r.Disposition.Extension}
	}
	m.Insurer = r.Insurer
	m.Insurance = r.Insurance
	m.PreAuthRef = r.PreAuthRef
	if r.PreAuthRef != nil && (r.PreAuthRef.Id != nil || r.PreAuthRef.Extension != nil) {
		m.PreAuthRefPrimitiveElement = &primitiveElement{Id: r.PreAuthRef.Id, Extension: r.PreAuthRef.Extension}
	}
	m.Form = r.Form
	m.Error = r.Error
	return m
}
func (r *CoverageEligibilityResponse) UnmarshalJSON(b []byte) error {
	var m jsonCoverageEligibilityResponse
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CoverageEligibilityResponse) unmarshalJSON(m jsonCoverageEligibilityResponse) error {
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
		r.Contained = append(r.Contained, v.Resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Purpose = m.Purpose
	for i, e := range m.PurposePrimitiveElement {
		if len(r.Purpose) > i {
			r.Purpose[i].Id = e.Id
			r.Purpose[i].Extension = e.Extension
		} else {
			r.Purpose = append(r.Purpose, Code{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Patient = m.Patient
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
	r.Created = m.Created
	if m.CreatedPrimitiveElement != nil {
		r.Created.Id = m.CreatedPrimitiveElement.Id
		r.Created.Extension = m.CreatedPrimitiveElement.Extension
	}
	r.Requestor = m.Requestor
	r.Request = m.Request
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
	r.Insurer = m.Insurer
	r.Insurance = m.Insurance
	r.PreAuthRef = m.PreAuthRef
	if m.PreAuthRefPrimitiveElement != nil {
		r.PreAuthRef.Id = m.PreAuthRefPrimitiveElement.Id
		r.PreAuthRef.Extension = m.PreAuthRefPrimitiveElement.Extension
	}
	r.Form = m.Form
	r.Error = m.Error
	return nil
}
func (r CoverageEligibilityResponse) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Financial instruments for reimbursement for the health care products and services.
type CoverageEligibilityResponseInsurance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Reference to the insurance card level information contained in the Coverage resource. The coverage issuing insurer will use these details to locate the patient's actual coverage within the insurer's information system.
	Coverage Reference
	// Flag indicating if the coverage provided is inforce currently if no service date(s) specified or for the whole duration of the service dates.
	Inforce *Boolean
	// The term of the benefits documented in this response.
	BenefitPeriod *Period
	// Benefits and optionally current balances, and authorization details by category or service.
	Item []CoverageEligibilityResponseInsuranceItem
}
type jsonCoverageEligibilityResponseInsurance struct {
	Id                      *string                                    `json:"id,omitempty"`
	Extension               []Extension                                `json:"extension,omitempty"`
	ModifierExtension       []Extension                                `json:"modifierExtension,omitempty"`
	Coverage                Reference                                  `json:"coverage,omitempty"`
	Inforce                 *Boolean                                   `json:"inforce,omitempty"`
	InforcePrimitiveElement *primitiveElement                          `json:"_inforce,omitempty"`
	BenefitPeriod           *Period                                    `json:"benefitPeriod,omitempty"`
	Item                    []CoverageEligibilityResponseInsuranceItem `json:"item,omitempty"`
}

func (r CoverageEligibilityResponseInsurance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CoverageEligibilityResponseInsurance) marshalJSON() jsonCoverageEligibilityResponseInsurance {
	m := jsonCoverageEligibilityResponseInsurance{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Coverage = r.Coverage
	m.Inforce = r.Inforce
	if r.Inforce != nil && (r.Inforce.Id != nil || r.Inforce.Extension != nil) {
		m.InforcePrimitiveElement = &primitiveElement{Id: r.Inforce.Id, Extension: r.Inforce.Extension}
	}
	m.BenefitPeriod = r.BenefitPeriod
	m.Item = r.Item
	return m
}
func (r *CoverageEligibilityResponseInsurance) UnmarshalJSON(b []byte) error {
	var m jsonCoverageEligibilityResponseInsurance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CoverageEligibilityResponseInsurance) unmarshalJSON(m jsonCoverageEligibilityResponseInsurance) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Coverage = m.Coverage
	r.Inforce = m.Inforce
	if m.InforcePrimitiveElement != nil {
		r.Inforce.Id = m.InforcePrimitiveElement.Id
		r.Inforce.Extension = m.InforcePrimitiveElement.Extension
	}
	r.BenefitPeriod = m.BenefitPeriod
	r.Item = m.Item
	return nil
}
func (r CoverageEligibilityResponseInsurance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Benefits and optionally current balances, and authorization details by category or service.
type CoverageEligibilityResponseInsuranceItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Code to identify the general type of benefits under which products and services are provided.
	Category *CodeableConcept
	// This contains the product, service, drug or other billing code for the item.
	ProductOrService *CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []CodeableConcept
	// The practitioner who is eligible for the provision of the product or service.
	Provider *Reference
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
	// Benefits used to date.
	Benefit []CoverageEligibilityResponseInsuranceItemBenefit
	// A boolean flag indicating whether a preauthorization is required prior to actual service delivery.
	AuthorizationRequired *Boolean
	// Codes or comments regarding information or actions associated with the preauthorization.
	AuthorizationSupporting []CodeableConcept
	// A web location for obtaining requirements or descriptive information regarding the preauthorization.
	AuthorizationUrl *Uri
}
type jsonCoverageEligibilityResponseInsuranceItem struct {
	Id                                    *string                                           `json:"id,omitempty"`
	Extension                             []Extension                                       `json:"extension,omitempty"`
	ModifierExtension                     []Extension                                       `json:"modifierExtension,omitempty"`
	Category                              *CodeableConcept                                  `json:"category,omitempty"`
	ProductOrService                      *CodeableConcept                                  `json:"productOrService,omitempty"`
	Modifier                              []CodeableConcept                                 `json:"modifier,omitempty"`
	Provider                              *Reference                                        `json:"provider,omitempty"`
	Excluded                              *Boolean                                          `json:"excluded,omitempty"`
	ExcludedPrimitiveElement              *primitiveElement                                 `json:"_excluded,omitempty"`
	Name                                  *String                                           `json:"name,omitempty"`
	NamePrimitiveElement                  *primitiveElement                                 `json:"_name,omitempty"`
	Description                           *String                                           `json:"description,omitempty"`
	DescriptionPrimitiveElement           *primitiveElement                                 `json:"_description,omitempty"`
	Network                               *CodeableConcept                                  `json:"network,omitempty"`
	Unit                                  *CodeableConcept                                  `json:"unit,omitempty"`
	Term                                  *CodeableConcept                                  `json:"term,omitempty"`
	Benefit                               []CoverageEligibilityResponseInsuranceItemBenefit `json:"benefit,omitempty"`
	AuthorizationRequired                 *Boolean                                          `json:"authorizationRequired,omitempty"`
	AuthorizationRequiredPrimitiveElement *primitiveElement                                 `json:"_authorizationRequired,omitempty"`
	AuthorizationSupporting               []CodeableConcept                                 `json:"authorizationSupporting,omitempty"`
	AuthorizationUrl                      *Uri                                              `json:"authorizationUrl,omitempty"`
	AuthorizationUrlPrimitiveElement      *primitiveElement                                 `json:"_authorizationUrl,omitempty"`
}

func (r CoverageEligibilityResponseInsuranceItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CoverageEligibilityResponseInsuranceItem) marshalJSON() jsonCoverageEligibilityResponseInsuranceItem {
	m := jsonCoverageEligibilityResponseInsuranceItem{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Category = r.Category
	m.ProductOrService = r.ProductOrService
	m.Modifier = r.Modifier
	m.Provider = r.Provider
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
	m.Benefit = r.Benefit
	m.AuthorizationRequired = r.AuthorizationRequired
	if r.AuthorizationRequired != nil && (r.AuthorizationRequired.Id != nil || r.AuthorizationRequired.Extension != nil) {
		m.AuthorizationRequiredPrimitiveElement = &primitiveElement{Id: r.AuthorizationRequired.Id, Extension: r.AuthorizationRequired.Extension}
	}
	m.AuthorizationSupporting = r.AuthorizationSupporting
	m.AuthorizationUrl = r.AuthorizationUrl
	if r.AuthorizationUrl != nil && (r.AuthorizationUrl.Id != nil || r.AuthorizationUrl.Extension != nil) {
		m.AuthorizationUrlPrimitiveElement = &primitiveElement{Id: r.AuthorizationUrl.Id, Extension: r.AuthorizationUrl.Extension}
	}
	return m
}
func (r *CoverageEligibilityResponseInsuranceItem) UnmarshalJSON(b []byte) error {
	var m jsonCoverageEligibilityResponseInsuranceItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CoverageEligibilityResponseInsuranceItem) unmarshalJSON(m jsonCoverageEligibilityResponseInsuranceItem) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Category = m.Category
	r.ProductOrService = m.ProductOrService
	r.Modifier = m.Modifier
	r.Provider = m.Provider
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
	r.Benefit = m.Benefit
	r.AuthorizationRequired = m.AuthorizationRequired
	if m.AuthorizationRequiredPrimitiveElement != nil {
		r.AuthorizationRequired.Id = m.AuthorizationRequiredPrimitiveElement.Id
		r.AuthorizationRequired.Extension = m.AuthorizationRequiredPrimitiveElement.Extension
	}
	r.AuthorizationSupporting = m.AuthorizationSupporting
	r.AuthorizationUrl = m.AuthorizationUrl
	if m.AuthorizationUrlPrimitiveElement != nil {
		r.AuthorizationUrl.Id = m.AuthorizationUrlPrimitiveElement.Id
		r.AuthorizationUrl.Extension = m.AuthorizationUrlPrimitiveElement.Extension
	}
	return nil
}
func (r CoverageEligibilityResponseInsuranceItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Benefits used to date.
type CoverageEligibilityResponseInsuranceItemBenefit struct {
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
	Allowed isCoverageEligibilityResponseInsuranceItemBenefitAllowed
	// The quantity of the benefit which have been consumed to date.
	Used isCoverageEligibilityResponseInsuranceItemBenefitUsed
}
type isCoverageEligibilityResponseInsuranceItemBenefitAllowed interface {
	isCoverageEligibilityResponseInsuranceItemBenefitAllowed()
}

func (r UnsignedInt) isCoverageEligibilityResponseInsuranceItemBenefitAllowed() {}
func (r String) isCoverageEligibilityResponseInsuranceItemBenefitAllowed()      {}
func (r Money) isCoverageEligibilityResponseInsuranceItemBenefitAllowed()       {}

type isCoverageEligibilityResponseInsuranceItemBenefitUsed interface {
	isCoverageEligibilityResponseInsuranceItemBenefitUsed()
}

func (r UnsignedInt) isCoverageEligibilityResponseInsuranceItemBenefitUsed() {}
func (r String) isCoverageEligibilityResponseInsuranceItemBenefitUsed()      {}
func (r Money) isCoverageEligibilityResponseInsuranceItemBenefitUsed()       {}

type jsonCoverageEligibilityResponseInsuranceItemBenefit struct {
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
	UsedString                         *String           `json:"usedString,omitempty"`
	UsedStringPrimitiveElement         *primitiveElement `json:"_usedString,omitempty"`
	UsedMoney                          *Money            `json:"usedMoney,omitempty"`
}

func (r CoverageEligibilityResponseInsuranceItemBenefit) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) marshalJSON() jsonCoverageEligibilityResponseInsuranceItemBenefit {
	m := jsonCoverageEligibilityResponseInsuranceItemBenefit{}
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
	case String:
		m.UsedString = &v
		if v.Id != nil || v.Extension != nil {
			m.UsedStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.UsedString = v
		if v.Id != nil || v.Extension != nil {
			m.UsedStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Money:
		m.UsedMoney = &v
	case *Money:
		m.UsedMoney = v
	}
	return m
}
func (r *CoverageEligibilityResponseInsuranceItemBenefit) UnmarshalJSON(b []byte) error {
	var m jsonCoverageEligibilityResponseInsuranceItemBenefit
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CoverageEligibilityResponseInsuranceItemBenefit) unmarshalJSON(m jsonCoverageEligibilityResponseInsuranceItemBenefit) error {
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
	if m.UsedString != nil || m.UsedStringPrimitiveElement != nil {
		if r.Used != nil {
			return fmt.Errorf("multiple values for field \"Used\"")
		}
		v := m.UsedString
		if m.UsedStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.UsedStringPrimitiveElement.Id
			v.Extension = m.UsedStringPrimitiveElement.Extension
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
func (r CoverageEligibilityResponseInsuranceItemBenefit) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Errors encountered during the processing of the request.
type CoverageEligibilityResponseError struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An error code,from a specified code system, which details why the eligibility check could not be performed.
	Code CodeableConcept
}
type jsonCoverageEligibilityResponseError struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code,omitempty"`
}

func (r CoverageEligibilityResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CoverageEligibilityResponseError) marshalJSON() jsonCoverageEligibilityResponseError {
	m := jsonCoverageEligibilityResponseError{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	return m
}
func (r *CoverageEligibilityResponseError) UnmarshalJSON(b []byte) error {
	var m jsonCoverageEligibilityResponseError
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CoverageEligibilityResponseError) unmarshalJSON(m jsonCoverageEligibilityResponseError) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	return nil
}
func (r CoverageEligibilityResponseError) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
