package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Benefits used to date.
type CoverageEligibilityResponseInsuranceItemBenefit struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Classification of benefit being provided.
	Type types.CodeableConcept
	// The quantity of the benefit which is permitted under the coverage.
	Allowed r4.Element
	// The quantity of the benefit which have been consumed to date.
	Used r4.Element
}

func (s CoverageEligibilityResponseInsuranceItemBenefit) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Benefits and optionally current balances, and authorization details by category or service.
type CoverageEligibilityResponseInsuranceItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// True if the indicated class of service is excluded from the plan, missing or False indicates the product or service is included in the coverage.
	Excluded *types.Boolean
	// A richer description of the benefit or services covered.
	Description *types.String
	// The term or period of the values such as 'maximum lifetime benefit' or 'maximum annual visits'.
	Term *types.CodeableConcept
	// Codes or comments regarding information or actions associated with the preauthorization.
	AuthorizationSupporting []types.CodeableConcept
	// A web location for obtaining requirements or descriptive information regarding the preauthorization.
	AuthorizationUrl *types.Uri
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Code to identify the general type of benefits under which products and services are provided.
	Category *types.CodeableConcept
	// A short name or tag for the benefit.
	Name *types.String
	// Is a flag to indicate whether the benefits refer to in-network providers or out-of-network providers.
	Network *types.CodeableConcept
	// A boolean flag indicating whether a preauthorization is required prior to actual service delivery.
	AuthorizationRequired *types.Boolean
	// This contains the product, service, drug or other billing code for the item.
	ProductOrService *types.CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []types.CodeableConcept
	// The practitioner who is eligible for the provision of the product or service.
	Provider *types.Reference
	// Indicates if the benefits apply to an individual or to the family.
	Unit *types.CodeableConcept
	// Benefits used to date.
	Benefit []CoverageEligibilityResponseInsuranceItemBenefit
}

func (s CoverageEligibilityResponseInsuranceItem) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Financial instruments for reimbursement for the health care products and services.
type CoverageEligibilityResponseInsurance struct {
	// The term of the benefits documented in this response.
	BenefitPeriod *types.Period
	// Benefits and optionally current balances, and authorization details by category or service.
	Item []CoverageEligibilityResponseInsuranceItem
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Reference to the insurance card level information contained in the Coverage resource. The coverage issuing insurer will use these details to locate the patient's actual coverage within the insurer's information system.
	Coverage types.Reference
	// Flag indicating if the coverage provided is inforce currently if no service date(s) specified or for the whole duration of the service dates.
	Inforce *types.Boolean
}

func (s CoverageEligibilityResponseInsurance) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
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
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// An error code,from a specified code system, which details why the eligibility check could not be performed.
	Code types.CodeableConcept
}

func (s CoverageEligibilityResponseError) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// This resource provides eligibility and plan details from the processing of an CoverageEligibilityRequest resource.
type CoverageEligibilityResponse struct {
	// Financial instruments for reimbursement for the health care products and services.
	Insurance []CoverageEligibilityResponseInsurance
	// A reference from the Insurer to which these services pertain to be used on further communication and as proof that the request occurred.
	PreAuthRef *types.String
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The status of the resource instance.
	Status types.Code
	// The outcome of the request processing.
	Outcome types.Code
	// A human readable description of the status of the adjudication.
	Disposition *types.String
	// The Insurer who issued the coverage in question and is the author of the response.
	Insurer types.Reference
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// The base language in which the resource is written.
	Language *types.Code
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Reference to the original request resource.
	Request types.Reference
	// A code for the form to be used for printing the content.
	Form *types.CodeableConcept
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// The party who is the beneficiary of the supplied coverage and for whom eligibility is sought.
	Patient types.Reference
	// The date or dates when the enclosed suite of services were performed or completed.
	Serviced r4.Element
	// The provider which is responsible for the request.
	Requestor *types.Reference
	// Errors encountered during the processing of the request.
	Error []CoverageEligibilityResponseError
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// A unique identifier assigned to this coverage eligiblity request.
	Identifier []types.Identifier
	// Code to specify whether requesting: prior authorization requirements for some service categories or billing codes; benefits for coverages specified or discovered; discovery and return of coverages for the patient; and/or validation that the specified coverage is in-force at the date/period specified or 'now' if not specified.
	Purpose []types.Code
	// The date this resource was created.
	Created types.DateTime
}

func (s CoverageEligibilityResponse) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
