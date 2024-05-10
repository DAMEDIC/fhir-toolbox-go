package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// The CoverageEligibilityRequest provides patient and insurance coverage information to an insurer for them to respond, in the form of an CoverageEligibilityResponse, with information regarding whether the stated coverage is valid and in-force and optionally to provide the insurance details of the policy.
type CoverageEligibilityRequest struct {
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
	// When the requestor expects the processor to complete processing.
	Priority *CodeableConcept
	// Code to specify whether requesting: prior authorization requirements for some service categories or billing codes; benefits for coverages specified or discovered; discovery and return of coverages for the patient; and/or validation that the specified coverage is in-force at the date/period specified or 'now' if not specified.
	Purpose []Code
	// The party who is the beneficiary of the supplied coverage and for whom eligibility is sought.
	Patient Reference
	// The date or dates when the enclosed suite of services were performed or completed.
	Serviced isCoverageEligibilityRequestServiced
	// The date when this resource was created.
	Created DateTime
	// Person who created the request.
	Enterer *Reference
	// The provider which is responsible for the request.
	Provider *Reference
	// The Insurer who issued the coverage in question and is the recipient of the request.
	Insurer Reference
	// Facility where the services are intended to be provided.
	Facility *Reference
	// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
	SupportingInfo []CoverageEligibilityRequestSupportingInfo
	// Financial instruments for reimbursement for the health care products and services.
	Insurance []CoverageEligibilityRequestInsurance
	// Service categories or billable services for which benefit details and/or an authorization prior to service delivery may be required by the payor.
	Item []CoverageEligibilityRequestItem
}
type isCoverageEligibilityRequestServiced interface {
	isCoverageEligibilityRequestServiced()
}

func (r Date) isCoverageEligibilityRequestServiced()   {}
func (r Period) isCoverageEligibilityRequestServiced() {}

type jsonCoverageEligibilityRequest struct {
	ResourceType                  string                                     `json:"resourceType"`
	Id                            *Id                                        `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement                          `json:"_id,omitempty"`
	Meta                          *Meta                                      `json:"meta,omitempty"`
	ImplicitRules                 *Uri                                       `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement                          `json:"_implicitRules,omitempty"`
	Language                      *Code                                      `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement                          `json:"_language,omitempty"`
	Text                          *Narrative                                 `json:"text,omitempty"`
	Contained                     []containedResource                        `json:"contained,omitempty"`
	Extension                     []Extension                                `json:"extension,omitempty"`
	ModifierExtension             []Extension                                `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier                               `json:"identifier,omitempty"`
	Status                        Code                                       `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement                          `json:"_status,omitempty"`
	Priority                      *CodeableConcept                           `json:"priority,omitempty"`
	Purpose                       []Code                                     `json:"purpose,omitempty"`
	PurposePrimitiveElement       []*primitiveElement                        `json:"_purpose,omitempty"`
	Patient                       Reference                                  `json:"patient,omitempty"`
	ServicedDate                  *Date                                      `json:"servicedDate,omitempty"`
	ServicedDatePrimitiveElement  *primitiveElement                          `json:"_servicedDate,omitempty"`
	ServicedPeriod                *Period                                    `json:"servicedPeriod,omitempty"`
	Created                       DateTime                                   `json:"created,omitempty"`
	CreatedPrimitiveElement       *primitiveElement                          `json:"_created,omitempty"`
	Enterer                       *Reference                                 `json:"enterer,omitempty"`
	Provider                      *Reference                                 `json:"provider,omitempty"`
	Insurer                       Reference                                  `json:"insurer,omitempty"`
	Facility                      *Reference                                 `json:"facility,omitempty"`
	SupportingInfo                []CoverageEligibilityRequestSupportingInfo `json:"supportingInfo,omitempty"`
	Insurance                     []CoverageEligibilityRequestInsurance      `json:"insurance,omitempty"`
	Item                          []CoverageEligibilityRequestItem           `json:"item,omitempty"`
}

func (r CoverageEligibilityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CoverageEligibilityRequest) marshalJSON() jsonCoverageEligibilityRequest {
	m := jsonCoverageEligibilityRequest{}
	m.ResourceType = "CoverageEligibilityRequest"
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
	m.Priority = r.Priority
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
	m.Enterer = r.Enterer
	m.Provider = r.Provider
	m.Insurer = r.Insurer
	m.Facility = r.Facility
	m.SupportingInfo = r.SupportingInfo
	m.Insurance = r.Insurance
	m.Item = r.Item
	return m
}
func (r *CoverageEligibilityRequest) UnmarshalJSON(b []byte) error {
	var m jsonCoverageEligibilityRequest
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CoverageEligibilityRequest) unmarshalJSON(m jsonCoverageEligibilityRequest) error {
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
	r.Priority = m.Priority
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
	r.Enterer = m.Enterer
	r.Provider = m.Provider
	r.Insurer = m.Insurer
	r.Facility = m.Facility
	r.SupportingInfo = m.SupportingInfo
	r.Insurance = m.Insurance
	r.Item = m.Item
	return nil
}
func (r CoverageEligibilityRequest) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
type CoverageEligibilityRequestSupportingInfo struct {
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
	// Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
	Information Reference
	// The supporting materials are applicable for all detail items, product/servce categories and specific billing codes.
	AppliesToAll *Boolean
}
type jsonCoverageEligibilityRequestSupportingInfo struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Sequence                     PositiveInt       `json:"sequence,omitempty"`
	SequencePrimitiveElement     *primitiveElement `json:"_sequence,omitempty"`
	Information                  Reference         `json:"information,omitempty"`
	AppliesToAll                 *Boolean          `json:"appliesToAll,omitempty"`
	AppliesToAllPrimitiveElement *primitiveElement `json:"_appliesToAll,omitempty"`
}

func (r CoverageEligibilityRequestSupportingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CoverageEligibilityRequestSupportingInfo) marshalJSON() jsonCoverageEligibilityRequestSupportingInfo {
	m := jsonCoverageEligibilityRequestSupportingInfo{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Sequence = r.Sequence
	if r.Sequence.Id != nil || r.Sequence.Extension != nil {
		m.SequencePrimitiveElement = &primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
	}
	m.Information = r.Information
	m.AppliesToAll = r.AppliesToAll
	if r.AppliesToAll != nil && (r.AppliesToAll.Id != nil || r.AppliesToAll.Extension != nil) {
		m.AppliesToAllPrimitiveElement = &primitiveElement{Id: r.AppliesToAll.Id, Extension: r.AppliesToAll.Extension}
	}
	return m
}
func (r *CoverageEligibilityRequestSupportingInfo) UnmarshalJSON(b []byte) error {
	var m jsonCoverageEligibilityRequestSupportingInfo
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CoverageEligibilityRequestSupportingInfo) unmarshalJSON(m jsonCoverageEligibilityRequestSupportingInfo) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Sequence = m.Sequence
	if m.SequencePrimitiveElement != nil {
		r.Sequence.Id = m.SequencePrimitiveElement.Id
		r.Sequence.Extension = m.SequencePrimitiveElement.Extension
	}
	r.Information = m.Information
	r.AppliesToAll = m.AppliesToAll
	if m.AppliesToAllPrimitiveElement != nil {
		r.AppliesToAll.Id = m.AppliesToAllPrimitiveElement.Id
		r.AppliesToAll.Extension = m.AppliesToAllPrimitiveElement.Extension
	}
	return nil
}
func (r CoverageEligibilityRequestSupportingInfo) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Financial instruments for reimbursement for the health care products and services.
type CoverageEligibilityRequestInsurance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A flag to indicate that this Coverage is to be used for evaluation of this request when set to true.
	Focal *Boolean
	// Reference to the insurance card level information contained in the Coverage resource. The coverage issuing insurer will use these details to locate the patient's actual coverage within the insurer's information system.
	Coverage Reference
	// A business agreement number established between the provider and the insurer for special business processing purposes.
	BusinessArrangement *String
}
type jsonCoverageEligibilityRequestInsurance struct {
	Id                                  *string           `json:"id,omitempty"`
	Extension                           []Extension       `json:"extension,omitempty"`
	ModifierExtension                   []Extension       `json:"modifierExtension,omitempty"`
	Focal                               *Boolean          `json:"focal,omitempty"`
	FocalPrimitiveElement               *primitiveElement `json:"_focal,omitempty"`
	Coverage                            Reference         `json:"coverage,omitempty"`
	BusinessArrangement                 *String           `json:"businessArrangement,omitempty"`
	BusinessArrangementPrimitiveElement *primitiveElement `json:"_businessArrangement,omitempty"`
}

func (r CoverageEligibilityRequestInsurance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CoverageEligibilityRequestInsurance) marshalJSON() jsonCoverageEligibilityRequestInsurance {
	m := jsonCoverageEligibilityRequestInsurance{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Focal = r.Focal
	if r.Focal != nil && (r.Focal.Id != nil || r.Focal.Extension != nil) {
		m.FocalPrimitiveElement = &primitiveElement{Id: r.Focal.Id, Extension: r.Focal.Extension}
	}
	m.Coverage = r.Coverage
	m.BusinessArrangement = r.BusinessArrangement
	if r.BusinessArrangement != nil && (r.BusinessArrangement.Id != nil || r.BusinessArrangement.Extension != nil) {
		m.BusinessArrangementPrimitiveElement = &primitiveElement{Id: r.BusinessArrangement.Id, Extension: r.BusinessArrangement.Extension}
	}
	return m
}
func (r *CoverageEligibilityRequestInsurance) UnmarshalJSON(b []byte) error {
	var m jsonCoverageEligibilityRequestInsurance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CoverageEligibilityRequestInsurance) unmarshalJSON(m jsonCoverageEligibilityRequestInsurance) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Focal = m.Focal
	if m.FocalPrimitiveElement != nil {
		r.Focal.Id = m.FocalPrimitiveElement.Id
		r.Focal.Extension = m.FocalPrimitiveElement.Extension
	}
	r.Coverage = m.Coverage
	r.BusinessArrangement = m.BusinessArrangement
	if m.BusinessArrangementPrimitiveElement != nil {
		r.BusinessArrangement.Id = m.BusinessArrangementPrimitiveElement.Id
		r.BusinessArrangement.Extension = m.BusinessArrangementPrimitiveElement.Extension
	}
	return nil
}
func (r CoverageEligibilityRequestInsurance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Service categories or billable services for which benefit details and/or an authorization prior to service delivery may be required by the payor.
type CoverageEligibilityRequestItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Exceptions, special conditions and supporting information applicable for this service or product line.
	SupportingInfoSequence []PositiveInt
	// Code to identify the general type of benefits under which products and services are provided.
	Category *CodeableConcept
	// This contains the product, service, drug or other billing code for the item.
	ProductOrService *CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []CodeableConcept
	// The practitioner who is responsible for the product or service to be rendered to the patient.
	Provider *Reference
	// The number of repetitions of a service or product.
	Quantity *Quantity
	// The amount charged to the patient by the provider for a single unit.
	UnitPrice *Money
	// Facility where the services will be provided.
	Facility *Reference
	// Patient diagnosis for which care is sought.
	Diagnosis []CoverageEligibilityRequestItemDiagnosis
	// The plan/proposal/order describing the proposed service in detail.
	Detail []Reference
}
type jsonCoverageEligibilityRequestItem struct {
	Id                                     *string                                   `json:"id,omitempty"`
	Extension                              []Extension                               `json:"extension,omitempty"`
	ModifierExtension                      []Extension                               `json:"modifierExtension,omitempty"`
	SupportingInfoSequence                 []PositiveInt                             `json:"supportingInfoSequence,omitempty"`
	SupportingInfoSequencePrimitiveElement []*primitiveElement                       `json:"_supportingInfoSequence,omitempty"`
	Category                               *CodeableConcept                          `json:"category,omitempty"`
	ProductOrService                       *CodeableConcept                          `json:"productOrService,omitempty"`
	Modifier                               []CodeableConcept                         `json:"modifier,omitempty"`
	Provider                               *Reference                                `json:"provider,omitempty"`
	Quantity                               *Quantity                                 `json:"quantity,omitempty"`
	UnitPrice                              *Money                                    `json:"unitPrice,omitempty"`
	Facility                               *Reference                                `json:"facility,omitempty"`
	Diagnosis                              []CoverageEligibilityRequestItemDiagnosis `json:"diagnosis,omitempty"`
	Detail                                 []Reference                               `json:"detail,omitempty"`
}

func (r CoverageEligibilityRequestItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CoverageEligibilityRequestItem) marshalJSON() jsonCoverageEligibilityRequestItem {
	m := jsonCoverageEligibilityRequestItem{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.SupportingInfoSequence = r.SupportingInfoSequence
	anySupportingInfoSequenceIdOrExtension := false
	for _, e := range r.SupportingInfoSequence {
		if e.Id != nil || e.Extension != nil {
			anySupportingInfoSequenceIdOrExtension = true
			break
		}
	}
	if anySupportingInfoSequenceIdOrExtension {
		m.SupportingInfoSequencePrimitiveElement = make([]*primitiveElement, 0, len(r.SupportingInfoSequence))
		for _, e := range r.SupportingInfoSequence {
			if e.Id != nil || e.Extension != nil {
				m.SupportingInfoSequencePrimitiveElement = append(m.SupportingInfoSequencePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SupportingInfoSequencePrimitiveElement = append(m.SupportingInfoSequencePrimitiveElement, nil)
			}
		}
	}
	m.Category = r.Category
	m.ProductOrService = r.ProductOrService
	m.Modifier = r.Modifier
	m.Provider = r.Provider
	m.Quantity = r.Quantity
	m.UnitPrice = r.UnitPrice
	m.Facility = r.Facility
	m.Diagnosis = r.Diagnosis
	m.Detail = r.Detail
	return m
}
func (r *CoverageEligibilityRequestItem) UnmarshalJSON(b []byte) error {
	var m jsonCoverageEligibilityRequestItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CoverageEligibilityRequestItem) unmarshalJSON(m jsonCoverageEligibilityRequestItem) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.SupportingInfoSequence = m.SupportingInfoSequence
	for i, e := range m.SupportingInfoSequencePrimitiveElement {
		if len(r.SupportingInfoSequence) > i {
			r.SupportingInfoSequence[i].Id = e.Id
			r.SupportingInfoSequence[i].Extension = e.Extension
		} else {
			r.SupportingInfoSequence = append(r.SupportingInfoSequence, PositiveInt{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Category = m.Category
	r.ProductOrService = m.ProductOrService
	r.Modifier = m.Modifier
	r.Provider = m.Provider
	r.Quantity = m.Quantity
	r.UnitPrice = m.UnitPrice
	r.Facility = m.Facility
	r.Diagnosis = m.Diagnosis
	r.Detail = m.Detail
	return nil
}
func (r CoverageEligibilityRequestItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Patient diagnosis for which care is sought.
type CoverageEligibilityRequestItemDiagnosis struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The nature of illness or problem in a coded form or as a reference to an external defined Condition.
	Diagnosis isCoverageEligibilityRequestItemDiagnosisDiagnosis
}
type isCoverageEligibilityRequestItemDiagnosisDiagnosis interface {
	isCoverageEligibilityRequestItemDiagnosisDiagnosis()
}

func (r CodeableConcept) isCoverageEligibilityRequestItemDiagnosisDiagnosis() {}
func (r Reference) isCoverageEligibilityRequestItemDiagnosisDiagnosis()       {}

type jsonCoverageEligibilityRequestItemDiagnosis struct {
	Id                       *string          `json:"id,omitempty"`
	Extension                []Extension      `json:"extension,omitempty"`
	ModifierExtension        []Extension      `json:"modifierExtension,omitempty"`
	DiagnosisCodeableConcept *CodeableConcept `json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *Reference       `json:"diagnosisReference,omitempty"`
}

func (r CoverageEligibilityRequestItemDiagnosis) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CoverageEligibilityRequestItemDiagnosis) marshalJSON() jsonCoverageEligibilityRequestItemDiagnosis {
	m := jsonCoverageEligibilityRequestItemDiagnosis{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
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
	return m
}
func (r *CoverageEligibilityRequestItemDiagnosis) UnmarshalJSON(b []byte) error {
	var m jsonCoverageEligibilityRequestItemDiagnosis
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CoverageEligibilityRequestItemDiagnosis) unmarshalJSON(m jsonCoverageEligibilityRequestItemDiagnosis) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
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
	return nil
}
func (r CoverageEligibilityRequestItemDiagnosis) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
