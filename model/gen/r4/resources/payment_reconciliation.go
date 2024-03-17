package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// A note that describes or explains the processing in a human readable form.
type PaymentReconciliationProcessNote struct {
	// The explanation or description associated with the processing.
	Text *types.String
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The business purpose of the note text.
	Type *types.Code
}

func (s PaymentReconciliationProcessNote) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Distribution of the payment amount for a previously acknowledged payable.
type PaymentReconciliationDetail struct {
	// The party which is receiving the payment.
	Payee *types.Reference
	// The monetary amount allocated from the total payment to the payable.
	Amount *types.Money
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A resource, such as a Claim, the evaluation of which could lead to payment.
	Request *types.Reference
	// The date from the response resource containing a commitment to pay.
	Date *types.Date
	// Unique identifier for the prior payment item for the referenced payable.
	Predecessor *types.Identifier
	// Code to indicate the nature of the payment.
	Type types.CodeableConcept
	// The party which submitted the claim or financial transaction.
	Submitter *types.Reference
	// A resource, such as a ClaimResponse, which contains a commitment to payment.
	Response *types.Reference
	// A reference to the individual who is responsible for inquiries regarding the response and its payment.
	Responsible *types.Reference
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Unique identifier for the current payment item for the referenced payable.
	Identifier *types.Identifier
}

func (s PaymentReconciliationDetail) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// This resource provides the details including amount of a payment and allocates the payment items being paid.
type PaymentReconciliation struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// The status of the resource instance.
	Status types.Code
	// The date of payment as indicated on the financial instrument.
	PaymentDate types.Date
	// Issuer's unique identifier for the payment instrument.
	PaymentIdentifier *types.Identifier
	// A note that describes or explains the processing in a human readable form.
	ProcessNote []PaymentReconciliationProcessNote
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The period of time for which payments have been gathered into this bulk payment for settlement.
	Period *types.Period
	// The date when the resource was created.
	Created types.DateTime
	// Original request resource reference.
	Request *types.Reference
	// A code for the form to be used for printing the content.
	FormCode *types.CodeableConcept
	// Total payment amount as indicated on the financial instrument.
	PaymentAmount types.Money
	// Distribution of the payment amount for a previously acknowledged payable.
	Detail []PaymentReconciliationDetail
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// The practitioner who is responsible for the services rendered to the patient.
	Requestor *types.Reference
	// The outcome of a request for a reconciliation.
	Outcome *types.Code
	// A human readable description of the status of the request for the reconciliation.
	Disposition *types.String
	// The base language in which the resource is written.
	Language *types.Code
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// A unique identifier assigned to this payment reconciliation.
	Identifier []types.Identifier
	// The party who generated the payment.
	PaymentIssuer *types.Reference
}

func (s PaymentReconciliation) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
