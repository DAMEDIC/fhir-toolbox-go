package r5

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	model "github.com/DAMEDIC/fhir-toolbox-go/model"
	"io"
	"slices"
	"unsafe"
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
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, nor can they have their own independent transaction scope. This is allowed to be a Parameters resource if and only if it is referenced by a resource that provides context/meaning.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
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
	// Information code for an event with a corresponding date or period.
	Event []CoverageEligibilityResponseEvent
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
type isCoverageEligibilityResponseServiced interface {
	model.Element
	isCoverageEligibilityResponseServiced()
}

func (r Date) isCoverageEligibilityResponseServiced()   {}
func (r Period) isCoverageEligibilityResponseServiced() {}

// Information code for an event with a corresponding date or period.
type CoverageEligibilityResponseEvent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A coded event such as when a service is expected or a card printed.
	Type CodeableConcept
	// A date or period in the past or future indicating when the event occurred or is expectd to occur.
	When isCoverageEligibilityResponseEventWhen
}
type isCoverageEligibilityResponseEventWhen interface {
	model.Element
	isCoverageEligibilityResponseEventWhen()
}

func (r DateTime) isCoverageEligibilityResponseEventWhen() {}
func (r Period) isCoverageEligibilityResponseEventWhen()   {}

// Financial instruments for reimbursement for the health care products and services.
type CoverageEligibilityResponseInsurance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
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

// Benefits and optionally current balances, and authorization details by category or service.
type CoverageEligibilityResponseInsuranceItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
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

// Benefits used to date.
type CoverageEligibilityResponseInsuranceItemBenefit struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
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
	model.Element
	isCoverageEligibilityResponseInsuranceItemBenefitAllowed()
}

func (r UnsignedInt) isCoverageEligibilityResponseInsuranceItemBenefitAllowed() {}
func (r String) isCoverageEligibilityResponseInsuranceItemBenefitAllowed()      {}
func (r Money) isCoverageEligibilityResponseInsuranceItemBenefitAllowed()       {}

type isCoverageEligibilityResponseInsuranceItemBenefitUsed interface {
	model.Element
	isCoverageEligibilityResponseInsuranceItemBenefitUsed()
}

func (r UnsignedInt) isCoverageEligibilityResponseInsuranceItemBenefitUsed() {}
func (r String) isCoverageEligibilityResponseInsuranceItemBenefitUsed()      {}
func (r Money) isCoverageEligibilityResponseInsuranceItemBenefitUsed()       {}

// Errors encountered during the processing of the request.
type CoverageEligibilityResponseError struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An error code,from a specified code system, which details why the eligibility check could not be performed.
	Code CodeableConcept
	// A [simple subset of FHIRPath](fhirpath.html#simple) limited to element names, repetition indicators and the default child accessor that identifies one of the elements in the resource that caused this issue to be raised.
	Expression []String
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
func (r CoverageEligibilityResponse) MemSize() int {
	var emptyIface any
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += r.Id.MemSize()
	}
	if r.Meta != nil {
		s += r.Meta.MemSize()
	}
	if r.ImplicitRules != nil {
		s += r.ImplicitRules.MemSize()
	}
	if r.Language != nil {
		s += r.Language.MemSize()
	}
	if r.Text != nil {
		s += r.Text.MemSize()
	}
	for _, i := range r.Contained {
		s += i.MemSize()
	}
	s += (cap(r.Contained) - len(r.Contained)) * int(unsafe.Sizeof(emptyIface))
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	s += r.Status.MemSize() - int(unsafe.Sizeof(r.Status))
	for _, i := range r.Purpose {
		s += i.MemSize()
	}
	s += (cap(r.Purpose) - len(r.Purpose)) * int(unsafe.Sizeof(Code{}))
	s += r.Patient.MemSize() - int(unsafe.Sizeof(r.Patient))
	for _, i := range r.Event {
		s += i.MemSize()
	}
	s += (cap(r.Event) - len(r.Event)) * int(unsafe.Sizeof(CoverageEligibilityResponseEvent{}))
	if r.Serviced != nil {
		s += r.Serviced.MemSize()
	}
	s += r.Created.MemSize() - int(unsafe.Sizeof(r.Created))
	if r.Requestor != nil {
		s += r.Requestor.MemSize()
	}
	s += r.Request.MemSize() - int(unsafe.Sizeof(r.Request))
	s += r.Outcome.MemSize() - int(unsafe.Sizeof(r.Outcome))
	if r.Disposition != nil {
		s += r.Disposition.MemSize()
	}
	s += r.Insurer.MemSize() - int(unsafe.Sizeof(r.Insurer))
	for _, i := range r.Insurance {
		s += i.MemSize()
	}
	s += (cap(r.Insurance) - len(r.Insurance)) * int(unsafe.Sizeof(CoverageEligibilityResponseInsurance{}))
	if r.PreAuthRef != nil {
		s += r.PreAuthRef.MemSize()
	}
	if r.Form != nil {
		s += r.Form.MemSize()
	}
	for _, i := range r.Error {
		s += i.MemSize()
	}
	s += (cap(r.Error) - len(r.Error)) * int(unsafe.Sizeof(CoverageEligibilityResponseError{}))
	return s
}
func (r CoverageEligibilityResponseEvent) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	s += r.Type.MemSize() - int(unsafe.Sizeof(r.Type))
	if r.When != nil {
		s += r.When.MemSize()
	}
	return s
}
func (r CoverageEligibilityResponseInsurance) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	s += r.Coverage.MemSize() - int(unsafe.Sizeof(r.Coverage))
	if r.Inforce != nil {
		s += r.Inforce.MemSize()
	}
	if r.BenefitPeriod != nil {
		s += r.BenefitPeriod.MemSize()
	}
	for _, i := range r.Item {
		s += i.MemSize()
	}
	s += (cap(r.Item) - len(r.Item)) * int(unsafe.Sizeof(CoverageEligibilityResponseInsuranceItem{}))
	return s
}
func (r CoverageEligibilityResponseInsuranceItem) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Category != nil {
		s += r.Category.MemSize()
	}
	if r.ProductOrService != nil {
		s += r.ProductOrService.MemSize()
	}
	for _, i := range r.Modifier {
		s += i.MemSize()
	}
	s += (cap(r.Modifier) - len(r.Modifier)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Provider != nil {
		s += r.Provider.MemSize()
	}
	if r.Excluded != nil {
		s += r.Excluded.MemSize()
	}
	if r.Name != nil {
		s += r.Name.MemSize()
	}
	if r.Description != nil {
		s += r.Description.MemSize()
	}
	if r.Network != nil {
		s += r.Network.MemSize()
	}
	if r.Unit != nil {
		s += r.Unit.MemSize()
	}
	if r.Term != nil {
		s += r.Term.MemSize()
	}
	for _, i := range r.Benefit {
		s += i.MemSize()
	}
	s += (cap(r.Benefit) - len(r.Benefit)) * int(unsafe.Sizeof(CoverageEligibilityResponseInsuranceItemBenefit{}))
	if r.AuthorizationRequired != nil {
		s += r.AuthorizationRequired.MemSize()
	}
	for _, i := range r.AuthorizationSupporting {
		s += i.MemSize()
	}
	s += (cap(r.AuthorizationSupporting) - len(r.AuthorizationSupporting)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.AuthorizationUrl != nil {
		s += r.AuthorizationUrl.MemSize()
	}
	return s
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	s += r.Type.MemSize() - int(unsafe.Sizeof(r.Type))
	if r.Allowed != nil {
		s += r.Allowed.MemSize()
	}
	if r.Used != nil {
		s += r.Used.MemSize()
	}
	return s
}
func (r CoverageEligibilityResponseError) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	s += r.Code.MemSize() - int(unsafe.Sizeof(r.Code))
	for _, i := range r.Expression {
		s += i.MemSize()
	}
	s += (cap(r.Expression) - len(r.Expression)) * int(unsafe.Sizeof(String{}))
	return s
}
func (r CoverageEligibilityResponse) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r CoverageEligibilityResponse) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CoverageEligibilityResponse) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"CoverageEligibilityResponse\""))
	if err != nil {
		return err
	}
	setComma := true
	if r.Id != nil && r.Id.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		p := primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_id\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Meta != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"meta\":"))
		if err != nil {
			return err
		}
		err = r.Meta.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"implicitRules\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ImplicitRules)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		p := primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_implicitRules\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && r.Language.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"language\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Language)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		p := primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_language\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Text != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"text\":"))
		if err != nil {
			return err
		}
		err = r.Text.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Contained) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"contained\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, c := range r.Contained {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = ContainedResource{c}.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Identifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"identifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Identifier {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"status\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Status)
		if err != nil {
			return err
		}
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		p := primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_status\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		anyValue := false
		for _, e := range r.Purpose {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"purpose\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.Purpose)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.Purpose {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_purpose\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.Purpose {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"patient\":"))
	if err != nil {
		return err
	}
	err = r.Patient.marshalJSON(w)
	if err != nil {
		return err
	}
	if len(r.Event) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"event\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Event {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	switch v := r.Serviced.(type) {
	case Date:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"servicedDate\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_servicedDate\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Date:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"servicedDate\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_servicedDate\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"servicedPeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"servicedPeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"created\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Created)
		if err != nil {
			return err
		}
	}
	if r.Created.Id != nil || r.Created.Extension != nil {
		p := primitiveElement{Id: r.Created.Id, Extension: r.Created.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_created\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Requestor != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"requestor\":"))
		if err != nil {
			return err
		}
		err = r.Requestor.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"request\":"))
	if err != nil {
		return err
	}
	err = r.Request.marshalJSON(w)
	if err != nil {
		return err
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"outcome\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Outcome)
		if err != nil {
			return err
		}
	}
	if r.Outcome.Id != nil || r.Outcome.Extension != nil {
		p := primitiveElement{Id: r.Outcome.Id, Extension: r.Outcome.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_outcome\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Disposition != nil && r.Disposition.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"disposition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Disposition)
		if err != nil {
			return err
		}
	}
	if r.Disposition != nil && (r.Disposition.Id != nil || r.Disposition.Extension != nil) {
		p := primitiveElement{Id: r.Disposition.Id, Extension: r.Disposition.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_disposition\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"insurer\":"))
	if err != nil {
		return err
	}
	err = r.Insurer.marshalJSON(w)
	if err != nil {
		return err
	}
	if len(r.Insurance) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"insurance\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Insurance {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.PreAuthRef != nil && r.PreAuthRef.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"preAuthRef\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.PreAuthRef)
		if err != nil {
			return err
		}
	}
	if r.PreAuthRef != nil && (r.PreAuthRef.Id != nil || r.PreAuthRef.Extension != nil) {
		p := primitiveElement{Id: r.PreAuthRef.Id, Extension: r.PreAuthRef.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_preAuthRef\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Form != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"form\":"))
		if err != nil {
			return err
		}
		err = r.Form.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Error) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"error\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Error {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r CoverageEligibilityResponseEvent) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CoverageEligibilityResponseEvent) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"type\":"))
	if err != nil {
		return err
	}
	err = r.Type.marshalJSON(w)
	if err != nil {
		return err
	}
	switch v := r.When.(type) {
	case DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"whenDateTime\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_whenDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"whenDateTime\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_whenDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"whenPeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"whenPeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r CoverageEligibilityResponseInsurance) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CoverageEligibilityResponseInsurance) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"coverage\":"))
	if err != nil {
		return err
	}
	err = r.Coverage.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.Inforce != nil && r.Inforce.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"inforce\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Inforce)
		if err != nil {
			return err
		}
	}
	if r.Inforce != nil && (r.Inforce.Id != nil || r.Inforce.Extension != nil) {
		p := primitiveElement{Id: r.Inforce.Id, Extension: r.Inforce.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_inforce\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.BenefitPeriod != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"benefitPeriod\":"))
		if err != nil {
			return err
		}
		err = r.BenefitPeriod.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Item) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"item\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Item {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r CoverageEligibilityResponseInsuranceItem) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CoverageEligibilityResponseInsuranceItem) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Category != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"category\":"))
		if err != nil {
			return err
		}
		err = r.Category.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ProductOrService != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"productOrService\":"))
		if err != nil {
			return err
		}
		err = r.ProductOrService.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Modifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Modifier {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Provider != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"provider\":"))
		if err != nil {
			return err
		}
		err = r.Provider.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Excluded != nil && r.Excluded.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"excluded\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Excluded)
		if err != nil {
			return err
		}
	}
	if r.Excluded != nil && (r.Excluded.Id != nil || r.Excluded.Extension != nil) {
		p := primitiveElement{Id: r.Excluded.Id, Extension: r.Excluded.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_excluded\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Name != nil && r.Name.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"name\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Name)
		if err != nil {
			return err
		}
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		p := primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_name\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Description != nil && r.Description.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"description\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Description)
		if err != nil {
			return err
		}
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		p := primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_description\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Network != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"network\":"))
		if err != nil {
			return err
		}
		err = r.Network.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Unit != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"unit\":"))
		if err != nil {
			return err
		}
		err = r.Unit.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Term != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"term\":"))
		if err != nil {
			return err
		}
		err = r.Term.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Benefit) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"benefit\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Benefit {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.AuthorizationRequired != nil && r.AuthorizationRequired.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"authorizationRequired\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.AuthorizationRequired)
		if err != nil {
			return err
		}
	}
	if r.AuthorizationRequired != nil && (r.AuthorizationRequired.Id != nil || r.AuthorizationRequired.Extension != nil) {
		p := primitiveElement{Id: r.AuthorizationRequired.Id, Extension: r.AuthorizationRequired.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_authorizationRequired\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.AuthorizationSupporting) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"authorizationSupporting\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.AuthorizationSupporting {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.AuthorizationUrl != nil && r.AuthorizationUrl.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"authorizationUrl\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.AuthorizationUrl)
		if err != nil {
			return err
		}
	}
	if r.AuthorizationUrl != nil && (r.AuthorizationUrl.Id != nil || r.AuthorizationUrl.Extension != nil) {
		p := primitiveElement{Id: r.AuthorizationUrl.Id, Extension: r.AuthorizationUrl.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_authorizationUrl\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"type\":"))
	if err != nil {
		return err
	}
	err = r.Type.marshalJSON(w)
	if err != nil {
		return err
	}
	switch v := r.Allowed.(type) {
	case UnsignedInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"allowedUnsignedInt\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_allowedUnsignedInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *UnsignedInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"allowedUnsignedInt\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_allowedUnsignedInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"allowedString\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_allowedString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"allowedString\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_allowedString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Money:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"allowedMoney\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Money:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"allowedMoney\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	}
	switch v := r.Used.(type) {
	case UnsignedInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"usedUnsignedInt\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_usedUnsignedInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *UnsignedInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"usedUnsignedInt\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_usedUnsignedInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"usedString\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_usedString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"usedString\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_usedString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Money:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"usedMoney\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Money:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"usedMoney\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r CoverageEligibilityResponseError) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r CoverageEligibilityResponseError) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"code\":"))
	if err != nil {
		return err
	}
	err = r.Code.marshalJSON(w)
	if err != nil {
		return err
	}
	{
		anyValue := false
		for _, e := range r.Expression {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"expression\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.Expression)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.Expression {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_expression\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.Expression {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r *CoverageEligibilityResponse) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *CoverageEligibilityResponse) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CoverageEligibilityResponse element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CoverageEligibilityResponse element", t)
		}
		switch f {
		case "resourceType":
			_, err := d.Token()
			if err != nil {
				return err
			}
		case "id":
			var v Id
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Value = v.Value
		case "_id":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Id = v.Id
			r.Id.Extension = v.Extension
		case "meta":
			var v Meta
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Meta = &v
		case "implicitRules":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Value = v.Value
		case "_implicitRules":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Id = v.Id
			r.ImplicitRules.Extension = v.Extension
		case "language":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Value = v.Value
		case "_language":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Id = v.Id
			r.Language.Extension = v.Extension
		case "text":
			var v Narrative
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Text = &v
		case "contained":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponse element", t)
			}
			for d.More() {
				var v ContainedResource
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Contained = append(r.Contained, v.Resource)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponse element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponse element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponse element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponse element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponse element", t)
			}
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponse element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponse element", t)
			}
		case "status":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Status.Value = v.Value
		case "_status":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Status.Id = v.Id
			r.Status.Extension = v.Extension
		case "purpose":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponse element", t)
			}
			for i := 0; d.More(); i++ {
				var v Code
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.Purpose) <= i {
					r.Purpose = append(r.Purpose, Code{})
				}
				r.Purpose[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponse element", t)
			}
		case "_purpose":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponse element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.Purpose) <= i {
					r.Purpose = append(r.Purpose, Code{})
				}
				r.Purpose[i].Id = v.Id
				r.Purpose[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponse element", t)
			}
		case "patient":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Patient = v
		case "event":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponse element", t)
			}
			for d.More() {
				var v CoverageEligibilityResponseEvent
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Event = append(r.Event, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponse element", t)
			}
		case "servicedDate":
			var v Date
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Serviced != nil {
				r.Serviced = Date{
					Extension: r.Serviced.(Date).Extension,
					Id:        r.Serviced.(Date).Id,
					Value:     v.Value,
				}
			} else {
				r.Serviced = v
			}
		case "_servicedDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Serviced != nil {
				r.Serviced = Date{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Serviced.(Date).Value,
				}
			} else {
				r.Serviced = Date{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "servicedPeriod":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Serviced = v
		case "created":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Created.Value = v.Value
		case "_created":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Created.Id = v.Id
			r.Created.Extension = v.Extension
		case "requestor":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Requestor = &v
		case "request":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Request = v
		case "outcome":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Outcome.Value = v.Value
		case "_outcome":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Outcome.Id = v.Id
			r.Outcome.Extension = v.Extension
		case "disposition":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Disposition == nil {
				r.Disposition = &String{}
			}
			r.Disposition.Value = v.Value
		case "_disposition":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Disposition == nil {
				r.Disposition = &String{}
			}
			r.Disposition.Id = v.Id
			r.Disposition.Extension = v.Extension
		case "insurer":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Insurer = v
		case "insurance":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponse element", t)
			}
			for d.More() {
				var v CoverageEligibilityResponseInsurance
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Insurance = append(r.Insurance, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponse element", t)
			}
		case "preAuthRef":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.PreAuthRef == nil {
				r.PreAuthRef = &String{}
			}
			r.PreAuthRef.Value = v.Value
		case "_preAuthRef":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.PreAuthRef == nil {
				r.PreAuthRef = &String{}
			}
			r.PreAuthRef.Id = v.Id
			r.PreAuthRef.Extension = v.Extension
		case "form":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Form = &v
		case "error":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponse element", t)
			}
			for d.More() {
				var v CoverageEligibilityResponseError
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Error = append(r.Error, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponse element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in CoverageEligibilityResponse", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CoverageEligibilityResponse element", t)
	}
	return nil
}
func (r *CoverageEligibilityResponseEvent) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CoverageEligibilityResponseEvent element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CoverageEligibilityResponseEvent element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseEvent element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseEvent element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseEvent element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseEvent element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = v
		case "whenDateTime":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.When != nil {
				r.When = DateTime{
					Extension: r.When.(DateTime).Extension,
					Id:        r.When.(DateTime).Id,
					Value:     v.Value,
				}
			} else {
				r.When = v
			}
		case "_whenDateTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.When != nil {
				r.When = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.When.(DateTime).Value,
				}
			} else {
				r.When = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "whenPeriod":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.When = v
		default:
			return fmt.Errorf("invalid field: %s in CoverageEligibilityResponseEvent", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CoverageEligibilityResponseEvent element", t)
	}
	return nil
}
func (r *CoverageEligibilityResponseInsurance) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CoverageEligibilityResponseInsurance element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CoverageEligibilityResponseInsurance element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseInsurance element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseInsurance element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseInsurance element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseInsurance element", t)
			}
		case "coverage":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Coverage = v
		case "inforce":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Inforce == nil {
				r.Inforce = &Boolean{}
			}
			r.Inforce.Value = v.Value
		case "_inforce":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Inforce == nil {
				r.Inforce = &Boolean{}
			}
			r.Inforce.Id = v.Id
			r.Inforce.Extension = v.Extension
		case "benefitPeriod":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.BenefitPeriod = &v
		case "item":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseInsurance element", t)
			}
			for d.More() {
				var v CoverageEligibilityResponseInsuranceItem
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Item = append(r.Item, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseInsurance element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in CoverageEligibilityResponseInsurance", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CoverageEligibilityResponseInsurance element", t)
	}
	return nil
}
func (r *CoverageEligibilityResponseInsuranceItem) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CoverageEligibilityResponseInsuranceItem element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CoverageEligibilityResponseInsuranceItem element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseInsuranceItem element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseInsuranceItem element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseInsuranceItem element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseInsuranceItem element", t)
			}
		case "category":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Category = &v
		case "productOrService":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ProductOrService = &v
		case "modifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseInsuranceItem element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Modifier = append(r.Modifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseInsuranceItem element", t)
			}
		case "provider":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Provider = &v
		case "excluded":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Excluded == nil {
				r.Excluded = &Boolean{}
			}
			r.Excluded.Value = v.Value
		case "_excluded":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Excluded == nil {
				r.Excluded = &Boolean{}
			}
			r.Excluded.Id = v.Id
			r.Excluded.Extension = v.Extension
		case "name":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Name == nil {
				r.Name = &String{}
			}
			r.Name.Value = v.Value
		case "_name":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Name == nil {
				r.Name = &String{}
			}
			r.Name.Id = v.Id
			r.Name.Extension = v.Extension
		case "description":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Value = v.Value
		case "_description":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Id = v.Id
			r.Description.Extension = v.Extension
		case "network":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Network = &v
		case "unit":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Unit = &v
		case "term":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Term = &v
		case "benefit":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseInsuranceItem element", t)
			}
			for d.More() {
				var v CoverageEligibilityResponseInsuranceItemBenefit
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Benefit = append(r.Benefit, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseInsuranceItem element", t)
			}
		case "authorizationRequired":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.AuthorizationRequired == nil {
				r.AuthorizationRequired = &Boolean{}
			}
			r.AuthorizationRequired.Value = v.Value
		case "_authorizationRequired":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.AuthorizationRequired == nil {
				r.AuthorizationRequired = &Boolean{}
			}
			r.AuthorizationRequired.Id = v.Id
			r.AuthorizationRequired.Extension = v.Extension
		case "authorizationSupporting":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseInsuranceItem element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.AuthorizationSupporting = append(r.AuthorizationSupporting, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseInsuranceItem element", t)
			}
		case "authorizationUrl":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.AuthorizationUrl == nil {
				r.AuthorizationUrl = &Uri{}
			}
			r.AuthorizationUrl.Value = v.Value
		case "_authorizationUrl":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.AuthorizationUrl == nil {
				r.AuthorizationUrl = &Uri{}
			}
			r.AuthorizationUrl.Id = v.Id
			r.AuthorizationUrl.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in CoverageEligibilityResponseInsuranceItem", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CoverageEligibilityResponseInsuranceItem element", t)
	}
	return nil
}
func (r *CoverageEligibilityResponseInsuranceItemBenefit) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CoverageEligibilityResponseInsuranceItemBenefit element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CoverageEligibilityResponseInsuranceItemBenefit element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseInsuranceItemBenefit element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseInsuranceItemBenefit element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseInsuranceItemBenefit element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseInsuranceItemBenefit element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = v
		case "allowedUnsignedInt":
			var v UnsignedInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Allowed != nil {
				r.Allowed = UnsignedInt{
					Extension: r.Allowed.(UnsignedInt).Extension,
					Id:        r.Allowed.(UnsignedInt).Id,
					Value:     v.Value,
				}
			} else {
				r.Allowed = v
			}
		case "_allowedUnsignedInt":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Allowed != nil {
				r.Allowed = UnsignedInt{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Allowed.(UnsignedInt).Value,
				}
			} else {
				r.Allowed = UnsignedInt{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "allowedString":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Allowed != nil {
				r.Allowed = String{
					Extension: r.Allowed.(String).Extension,
					Id:        r.Allowed.(String).Id,
					Value:     v.Value,
				}
			} else {
				r.Allowed = v
			}
		case "_allowedString":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Allowed != nil {
				r.Allowed = String{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Allowed.(String).Value,
				}
			} else {
				r.Allowed = String{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "allowedMoney":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Allowed = v
		case "usedUnsignedInt":
			var v UnsignedInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Used != nil {
				r.Used = UnsignedInt{
					Extension: r.Used.(UnsignedInt).Extension,
					Id:        r.Used.(UnsignedInt).Id,
					Value:     v.Value,
				}
			} else {
				r.Used = v
			}
		case "_usedUnsignedInt":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Used != nil {
				r.Used = UnsignedInt{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Used.(UnsignedInt).Value,
				}
			} else {
				r.Used = UnsignedInt{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "usedString":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Used != nil {
				r.Used = String{
					Extension: r.Used.(String).Extension,
					Id:        r.Used.(String).Id,
					Value:     v.Value,
				}
			} else {
				r.Used = v
			}
		case "_usedString":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Used != nil {
				r.Used = String{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Used.(String).Value,
				}
			} else {
				r.Used = String{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "usedMoney":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Used = v
		default:
			return fmt.Errorf("invalid field: %s in CoverageEligibilityResponseInsuranceItemBenefit", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CoverageEligibilityResponseInsuranceItemBenefit element", t)
	}
	return nil
}
func (r *CoverageEligibilityResponseError) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in CoverageEligibilityResponseError element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in CoverageEligibilityResponseError element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseError element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseError element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseError element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseError element", t)
			}
		case "code":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Code = v
		case "expression":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseError element", t)
			}
			for i := 0; d.More(); i++ {
				var v String
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.Expression) <= i {
					r.Expression = append(r.Expression, String{})
				}
				r.Expression[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseError element", t)
			}
		case "_expression":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in CoverageEligibilityResponseError element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.Expression) <= i {
					r.Expression = append(r.Expression, String{})
				}
				r.Expression[i].Id = v.Id
				r.Expression[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in CoverageEligibilityResponseError element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in CoverageEligibilityResponseError", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in CoverageEligibilityResponseError element", t)
	}
	return nil
}
func (r CoverageEligibilityResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "CoverageEligibilityResponse"
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
	for _, c := range r.Contained {
		err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "contained"}})
		if err != nil {
			return err
		}
		err = e.EncodeElement(c, xml.StartElement{Name: xml.Name{Local: "__contained__"}})
		if err != nil {
			return err
		}
		err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "contained"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Purpose, xml.StartElement{Name: xml.Name{Local: "purpose"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Patient, xml.StartElement{Name: xml.Name{Local: "patient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Event, xml.StartElement{Name: xml.Name{Local: "event"}})
	if err != nil {
		return err
	}
	switch v := r.Serviced.(type) {
	case Date, *Date:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "servicedDate"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "servicedPeriod"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Created, xml.StartElement{Name: xml.Name{Local: "created"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Requestor, xml.StartElement{Name: xml.Name{Local: "requestor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Request, xml.StartElement{Name: xml.Name{Local: "request"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Outcome, xml.StartElement{Name: xml.Name{Local: "outcome"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Disposition, xml.StartElement{Name: xml.Name{Local: "disposition"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Insurer, xml.StartElement{Name: xml.Name{Local: "insurer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Insurance, xml.StartElement{Name: xml.Name{Local: "insurance"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PreAuthRef, xml.StartElement{Name: xml.Name{Local: "preAuthRef"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Form, xml.StartElement{Name: xml.Name{Local: "form"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Error, xml.StartElement{Name: xml.Name{Local: "error"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CoverageEligibilityResponseEvent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	switch v := r.When.(type) {
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "whenDateTime"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "whenPeriod"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CoverageEligibilityResponseInsurance) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Coverage, xml.StartElement{Name: xml.Name{Local: "coverage"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Inforce, xml.StartElement{Name: xml.Name{Local: "inforce"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BenefitPeriod, xml.StartElement{Name: xml.Name{Local: "benefitPeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Item, xml.StartElement{Name: xml.Name{Local: "item"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CoverageEligibilityResponseInsuranceItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Category, xml.StartElement{Name: xml.Name{Local: "category"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProductOrService, xml.StartElement{Name: xml.Name{Local: "productOrService"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Modifier, xml.StartElement{Name: xml.Name{Local: "modifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Provider, xml.StartElement{Name: xml.Name{Local: "provider"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Excluded, xml.StartElement{Name: xml.Name{Local: "excluded"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Network, xml.StartElement{Name: xml.Name{Local: "network"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Unit, xml.StartElement{Name: xml.Name{Local: "unit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Term, xml.StartElement{Name: xml.Name{Local: "term"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Benefit, xml.StartElement{Name: xml.Name{Local: "benefit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AuthorizationRequired, xml.StartElement{Name: xml.Name{Local: "authorizationRequired"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AuthorizationSupporting, xml.StartElement{Name: xml.Name{Local: "authorizationSupporting"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AuthorizationUrl, xml.StartElement{Name: xml.Name{Local: "authorizationUrl"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	switch v := r.Allowed.(type) {
	case UnsignedInt, *UnsignedInt:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "allowedUnsignedInt"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "allowedString"}})
		if err != nil {
			return err
		}
	case Money, *Money:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "allowedMoney"}})
		if err != nil {
			return err
		}
	}
	switch v := r.Used.(type) {
	case UnsignedInt, *UnsignedInt:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "usedUnsignedInt"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "usedString"}})
		if err != nil {
			return err
		}
	case Money, *Money:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "usedMoney"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r CoverageEligibilityResponseError) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Expression, xml.StartElement{Name: xml.Name{Local: "expression"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CoverageEligibilityResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "purpose":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Purpose = append(r.Purpose, v)
			case "patient":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Patient = v
			case "event":
				var v CoverageEligibilityResponseEvent
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Event = append(r.Event, v)
			case "servicedDate":
				if r.Serviced != nil {
					return fmt.Errorf("multiple values for field \"Serviced\"")
				}
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Serviced = v
			case "servicedPeriod":
				if r.Serviced != nil {
					return fmt.Errorf("multiple values for field \"Serviced\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Serviced = v
			case "created":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Created = v
			case "requestor":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Requestor = &v
			case "request":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Request = v
			case "outcome":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Outcome = v
			case "disposition":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Disposition = &v
			case "insurer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Insurer = v
			case "insurance":
				var v CoverageEligibilityResponseInsurance
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Insurance = append(r.Insurance, v)
			case "preAuthRef":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PreAuthRef = &v
			case "form":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Form = &v
			case "error":
				var v CoverageEligibilityResponseError
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Error = append(r.Error, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CoverageEligibilityResponseEvent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "whenDateTime":
				if r.When != nil {
					return fmt.Errorf("multiple values for field \"When\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.When = v
			case "whenPeriod":
				if r.When != nil {
					return fmt.Errorf("multiple values for field \"When\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.When = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CoverageEligibilityResponseInsurance) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "coverage":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Coverage = v
			case "inforce":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Inforce = &v
			case "benefitPeriod":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BenefitPeriod = &v
			case "item":
				var v CoverageEligibilityResponseInsuranceItem
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Item = append(r.Item, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CoverageEligibilityResponseInsuranceItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "category":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = &v
			case "productOrService":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductOrService = &v
			case "modifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Modifier = append(r.Modifier, v)
			case "provider":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Provider = &v
			case "excluded":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Excluded = &v
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "network":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Network = &v
			case "unit":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Unit = &v
			case "term":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Term = &v
			case "benefit":
				var v CoverageEligibilityResponseInsuranceItemBenefit
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Benefit = append(r.Benefit, v)
			case "authorizationRequired":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AuthorizationRequired = &v
			case "authorizationSupporting":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AuthorizationSupporting = append(r.AuthorizationSupporting, v)
			case "authorizationUrl":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AuthorizationUrl = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CoverageEligibilityResponseInsuranceItemBenefit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "allowedUnsignedInt":
				if r.Allowed != nil {
					return fmt.Errorf("multiple values for field \"Allowed\"")
				}
				var v UnsignedInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Allowed = v
			case "allowedString":
				if r.Allowed != nil {
					return fmt.Errorf("multiple values for field \"Allowed\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Allowed = v
			case "allowedMoney":
				if r.Allowed != nil {
					return fmt.Errorf("multiple values for field \"Allowed\"")
				}
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Allowed = v
			case "usedUnsignedInt":
				if r.Used != nil {
					return fmt.Errorf("multiple values for field \"Used\"")
				}
				var v UnsignedInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Used = v
			case "usedString":
				if r.Used != nil {
					return fmt.Errorf("multiple values for field \"Used\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Used = v
			case "usedMoney":
				if r.Used != nil {
					return fmt.Errorf("multiple values for field \"Used\"")
				}
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Used = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *CoverageEligibilityResponseError) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = v
			case "expression":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Expression = append(r.Expression, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CoverageEligibilityResponse) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, *r.Id)
		}
	}
	if len(name) == 0 || slices.Contains(name, "meta") {
		if r.Meta != nil {
			children = append(children, *r.Meta)
		}
	}
	if len(name) == 0 || slices.Contains(name, "implicitRules") {
		if r.ImplicitRules != nil {
			children = append(children, *r.ImplicitRules)
		}
	}
	if len(name) == 0 || slices.Contains(name, "language") {
		if r.Language != nil {
			children = append(children, *r.Language)
		}
	}
	if len(name) == 0 || slices.Contains(name, "text") {
		if r.Text != nil {
			children = append(children, *r.Text)
		}
	}
	if len(name) == 0 || slices.Contains(name, "contained") {
		for _, v := range r.Contained {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "identifier") {
		for _, v := range r.Identifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "status") {
		children = append(children, r.Status)
	}
	if len(name) == 0 || slices.Contains(name, "purpose") {
		for _, v := range r.Purpose {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "patient") {
		children = append(children, r.Patient)
	}
	if len(name) == 0 || slices.Contains(name, "event") {
		for _, v := range r.Event {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "serviced") {
		if r.Serviced != nil {
			children = append(children, r.Serviced)
		}
	}
	if len(name) == 0 || slices.Contains(name, "created") {
		children = append(children, r.Created)
	}
	if len(name) == 0 || slices.Contains(name, "requestor") {
		if r.Requestor != nil {
			children = append(children, *r.Requestor)
		}
	}
	if len(name) == 0 || slices.Contains(name, "request") {
		children = append(children, r.Request)
	}
	if len(name) == 0 || slices.Contains(name, "outcome") {
		children = append(children, r.Outcome)
	}
	if len(name) == 0 || slices.Contains(name, "disposition") {
		if r.Disposition != nil {
			children = append(children, *r.Disposition)
		}
	}
	if len(name) == 0 || slices.Contains(name, "insurer") {
		children = append(children, r.Insurer)
	}
	if len(name) == 0 || slices.Contains(name, "insurance") {
		for _, v := range r.Insurance {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "preAuthRef") {
		if r.PreAuthRef != nil {
			children = append(children, *r.PreAuthRef)
		}
	}
	if len(name) == 0 || slices.Contains(name, "form") {
		if r.Form != nil {
			children = append(children, *r.Form)
		}
	}
	if len(name) == 0 || slices.Contains(name, "error") {
		for _, v := range r.Error {
			children = append(children, v)
		}
	}
	return children
}
func (r CoverageEligibilityResponse) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponse to Boolean")
}
func (r CoverageEligibilityResponse) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponse to String")
}
func (r CoverageEligibilityResponse) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponse to Integer")
}
func (r CoverageEligibilityResponse) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponse to Decimal")
}
func (r CoverageEligibilityResponse) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponse to Date")
}
func (r CoverageEligibilityResponse) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponse to Time")
}
func (r CoverageEligibilityResponse) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponse to DateTime")
}
func (r CoverageEligibilityResponse) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponse to Quantity")
}
func (r CoverageEligibilityResponse) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.Id",
		}, {
			Name: "Meta",
			Type: "FHIR.Meta",
		}, {
			Name: "ImplicitRules",
			Type: "FHIR.Uri",
		}, {
			Name: "Language",
			Type: "FHIR.Code",
		}, {
			Name: "Text",
			Type: "FHIR.Narrative",
		}, {
			Name: "Contained",
			Type: "List<FHIR.>",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Identifier",
			Type: "List<FHIR.Identifier>",
		}, {
			Name: "Status",
			Type: "FHIR.Code",
		}, {
			Name: "Purpose",
			Type: "List<FHIR.Code>",
		}, {
			Name: "Patient",
			Type: "FHIR.Reference",
		}, {
			Name: "Event",
			Type: "List<FHIR.CoverageEligibilityResponseEvent>",
		}, {
			Name: "Serviced",
			Type: "FHIR.PrimitiveElement",
		}, {
			Name: "Created",
			Type: "FHIR.DateTime",
		}, {
			Name: "Requestor",
			Type: "FHIR.Reference",
		}, {
			Name: "Request",
			Type: "FHIR.Reference",
		}, {
			Name: "Outcome",
			Type: "FHIR.Code",
		}, {
			Name: "Disposition",
			Type: "FHIR.String",
		}, {
			Name: "Insurer",
			Type: "FHIR.Reference",
		}, {
			Name: "Insurance",
			Type: "List<FHIR.CoverageEligibilityResponseInsurance>",
		}, {
			Name: "PreAuthRef",
			Type: "FHIR.String",
		}, {
			Name: "Form",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Error",
			Type: "List<FHIR.CoverageEligibilityResponseError>",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "CoverageEligibilityResponse",
			Namespace: "FHIR",
		},
	}
}
func (r CoverageEligibilityResponseEvent) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "type") {
		children = append(children, r.Type)
	}
	if len(name) == 0 || slices.Contains(name, "when") {
		children = append(children, r.When)
	}
	return children
}
func (r CoverageEligibilityResponseEvent) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseEvent to Boolean")
}
func (r CoverageEligibilityResponseEvent) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseEvent to String")
}
func (r CoverageEligibilityResponseEvent) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseEvent to Integer")
}
func (r CoverageEligibilityResponseEvent) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseEvent to Decimal")
}
func (r CoverageEligibilityResponseEvent) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseEvent to Date")
}
func (r CoverageEligibilityResponseEvent) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseEvent to Time")
}
func (r CoverageEligibilityResponseEvent) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseEvent to DateTime")
}
func (r CoverageEligibilityResponseEvent) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseEvent to Quantity")
}
func (r CoverageEligibilityResponseEvent) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Type",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "When",
			Type: "FHIR.PrimitiveElement",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CoverageEligibilityResponseEvent",
			Namespace: "FHIR",
		},
	}
}
func (r CoverageEligibilityResponseInsurance) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "coverage") {
		children = append(children, r.Coverage)
	}
	if len(name) == 0 || slices.Contains(name, "inforce") {
		if r.Inforce != nil {
			children = append(children, *r.Inforce)
		}
	}
	if len(name) == 0 || slices.Contains(name, "benefitPeriod") {
		if r.BenefitPeriod != nil {
			children = append(children, *r.BenefitPeriod)
		}
	}
	if len(name) == 0 || slices.Contains(name, "item") {
		for _, v := range r.Item {
			children = append(children, v)
		}
	}
	return children
}
func (r CoverageEligibilityResponseInsurance) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsurance to Boolean")
}
func (r CoverageEligibilityResponseInsurance) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsurance to String")
}
func (r CoverageEligibilityResponseInsurance) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsurance to Integer")
}
func (r CoverageEligibilityResponseInsurance) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsurance to Decimal")
}
func (r CoverageEligibilityResponseInsurance) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsurance to Date")
}
func (r CoverageEligibilityResponseInsurance) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsurance to Time")
}
func (r CoverageEligibilityResponseInsurance) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsurance to DateTime")
}
func (r CoverageEligibilityResponseInsurance) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsurance to Quantity")
}
func (r CoverageEligibilityResponseInsurance) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Coverage",
			Type: "FHIR.Reference",
		}, {
			Name: "Inforce",
			Type: "FHIR.Boolean",
		}, {
			Name: "BenefitPeriod",
			Type: "FHIR.Period",
		}, {
			Name: "Item",
			Type: "List<FHIR.CoverageEligibilityResponseInsuranceItem>",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CoverageEligibilityResponseInsurance",
			Namespace: "FHIR",
		},
	}
}
func (r CoverageEligibilityResponseInsuranceItem) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "category") {
		if r.Category != nil {
			children = append(children, *r.Category)
		}
	}
	if len(name) == 0 || slices.Contains(name, "productOrService") {
		if r.ProductOrService != nil {
			children = append(children, *r.ProductOrService)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifier") {
		for _, v := range r.Modifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "provider") {
		if r.Provider != nil {
			children = append(children, *r.Provider)
		}
	}
	if len(name) == 0 || slices.Contains(name, "excluded") {
		if r.Excluded != nil {
			children = append(children, *r.Excluded)
		}
	}
	if len(name) == 0 || slices.Contains(name, "name") {
		if r.Name != nil {
			children = append(children, *r.Name)
		}
	}
	if len(name) == 0 || slices.Contains(name, "description") {
		if r.Description != nil {
			children = append(children, *r.Description)
		}
	}
	if len(name) == 0 || slices.Contains(name, "network") {
		if r.Network != nil {
			children = append(children, *r.Network)
		}
	}
	if len(name) == 0 || slices.Contains(name, "unit") {
		if r.Unit != nil {
			children = append(children, *r.Unit)
		}
	}
	if len(name) == 0 || slices.Contains(name, "term") {
		if r.Term != nil {
			children = append(children, *r.Term)
		}
	}
	if len(name) == 0 || slices.Contains(name, "benefit") {
		for _, v := range r.Benefit {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "authorizationRequired") {
		if r.AuthorizationRequired != nil {
			children = append(children, *r.AuthorizationRequired)
		}
	}
	if len(name) == 0 || slices.Contains(name, "authorizationSupporting") {
		for _, v := range r.AuthorizationSupporting {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "authorizationUrl") {
		if r.AuthorizationUrl != nil {
			children = append(children, *r.AuthorizationUrl)
		}
	}
	return children
}
func (r CoverageEligibilityResponseInsuranceItem) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItem to Boolean")
}
func (r CoverageEligibilityResponseInsuranceItem) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItem to String")
}
func (r CoverageEligibilityResponseInsuranceItem) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItem to Integer")
}
func (r CoverageEligibilityResponseInsuranceItem) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItem to Decimal")
}
func (r CoverageEligibilityResponseInsuranceItem) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItem to Date")
}
func (r CoverageEligibilityResponseInsuranceItem) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItem to Time")
}
func (r CoverageEligibilityResponseInsuranceItem) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItem to DateTime")
}
func (r CoverageEligibilityResponseInsuranceItem) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItem to Quantity")
}
func (r CoverageEligibilityResponseInsuranceItem) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Category",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "ProductOrService",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Modifier",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "Provider",
			Type: "FHIR.Reference",
		}, {
			Name: "Excluded",
			Type: "FHIR.Boolean",
		}, {
			Name: "Name",
			Type: "FHIR.String",
		}, {
			Name: "Description",
			Type: "FHIR.String",
		}, {
			Name: "Network",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Unit",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Term",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Benefit",
			Type: "List<FHIR.CoverageEligibilityResponseInsuranceItemBenefit>",
		}, {
			Name: "AuthorizationRequired",
			Type: "FHIR.Boolean",
		}, {
			Name: "AuthorizationSupporting",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "AuthorizationUrl",
			Type: "FHIR.Uri",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CoverageEligibilityResponseInsuranceItem",
			Namespace: "FHIR",
		},
	}
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "type") {
		children = append(children, r.Type)
	}
	if len(name) == 0 || slices.Contains(name, "allowed") {
		if r.Allowed != nil {
			children = append(children, r.Allowed)
		}
	}
	if len(name) == 0 || slices.Contains(name, "used") {
		if r.Used != nil {
			children = append(children, r.Used)
		}
	}
	return children
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItemBenefit to Boolean")
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItemBenefit to String")
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItemBenefit to Integer")
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItemBenefit to Decimal")
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItemBenefit to Date")
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItemBenefit to Time")
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItemBenefit to DateTime")
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseInsuranceItemBenefit to Quantity")
}
func (r CoverageEligibilityResponseInsuranceItemBenefit) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Type",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Allowed",
			Type: "FHIR.PrimitiveElement",
		}, {
			Name: "Used",
			Type: "FHIR.PrimitiveElement",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CoverageEligibilityResponseInsuranceItemBenefit",
			Namespace: "FHIR",
		},
	}
}
func (r CoverageEligibilityResponseError) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "code") {
		children = append(children, r.Code)
	}
	if len(name) == 0 || slices.Contains(name, "expression") {
		for _, v := range r.Expression {
			children = append(children, v)
		}
	}
	return children
}
func (r CoverageEligibilityResponseError) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseError to Boolean")
}
func (r CoverageEligibilityResponseError) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseError to String")
}
func (r CoverageEligibilityResponseError) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseError to Integer")
}
func (r CoverageEligibilityResponseError) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseError to Decimal")
}
func (r CoverageEligibilityResponseError) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseError to Date")
}
func (r CoverageEligibilityResponseError) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseError to Time")
}
func (r CoverageEligibilityResponseError) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseError to DateTime")
}
func (r CoverageEligibilityResponseError) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert CoverageEligibilityResponseError to Quantity")
}
func (r CoverageEligibilityResponseError) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Code",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Expression",
			Type: "List<FHIR.String>",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "CoverageEligibilityResponseError",
			Namespace: "FHIR",
		},
	}
}
