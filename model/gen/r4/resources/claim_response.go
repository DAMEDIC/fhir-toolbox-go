package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Errors encountered during the processing of the adjudication.
type ClaimResponseError struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The sequence number of the line item submitted which contains the error. This value is omitted when the error occurs outside of the item structure.
	ItemSequence *types.PositiveInt
	// The sequence number of the detail within the line item submitted which contains the error. This value is omitted when the error occurs outside of the item structure.
	DetailSequence *types.PositiveInt
	// The sequence number of the sub-detail within the detail within the line item submitted which contains the error. This value is omitted when the error occurs outside of the item structure.
	SubDetailSequence *types.PositiveInt
	// An error code, from a specified code system, which details why the claim could not be adjudicated.
	Code types.CodeableConcept
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
}

func (s ClaimResponseError) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Categorized monetary totals for the adjudication.
type ClaimResponseTotal struct {
	// Monetary total amount associated with the category.
	Amount types.Money
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A code to indicate the information type of this adjudication record. Information types may include: the value submitted, maximum values or percentages allowed or payable under the plan, amounts that the patient is responsible for in aggregate or pertaining to this item, amounts paid by other coverages, and the benefit payable for this item.
	Category types.CodeableConcept
}

func (s ClaimResponseTotal) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// If this item is a group then the values here are a summary of the adjudication of the detail items. If this item is a simple product or service then this is the result of the adjudication of this item.
type ClaimResponseItemAdjudication struct {
	// A code to indicate the information type of this adjudication record. Information types may include the value submitted, maximum values or percentages allowed or payable under the plan, amounts that: the patient is responsible for in aggregate or pertaining to this item; amounts paid by other coverages; and, the benefit payable for this item.
	Category types.CodeableConcept
	// A code supporting the understanding of the adjudication result and explaining variance from expected amount.
	Reason *types.CodeableConcept
	// Monetary amount associated with the category.
	Amount *types.Money
	// A non-monetary value associated with the category. Mutually exclusive to the amount element above.
	Value *types.Decimal
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
}

func (s ClaimResponseItemAdjudication) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A sub-detail adjudication of a simple product or service.
type ClaimResponseItemDetailSubDetail struct {
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A number to uniquely reference the claim sub-detail entry.
	SubDetailSequence types.PositiveInt
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []types.PositiveInt
}

func (s ClaimResponseItemDetailSubDetail) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A claim detail. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
type ClaimResponseItemDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A number to uniquely reference the claim detail entry.
	DetailSequence types.PositiveInt
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []types.PositiveInt
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
	// A sub-detail adjudication of a simple product or service.
	SubDetail []ClaimResponseItemDetailSubDetail
}

func (s ClaimResponseItemDetail) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
type ClaimResponseItem struct {
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A number to uniquely reference the claim item entries.
	ItemSequence types.PositiveInt
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []types.PositiveInt
	// If this item is a group then the values here are a summary of the adjudication of the detail items. If this item is a simple product or service then this is the result of the adjudication of this item.
	Adjudication []ClaimResponseItemAdjudication
	// A claim detail. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
	Detail []ClaimResponseItemDetail
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
}

func (s ClaimResponseItem) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The third-tier service adjudications for payor added services.
type ClaimResponseAddItemDetailSubDetail struct {
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *types.Money
	// The quantity times the unit price for an additional service or product or charge.
	Net *types.Money
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService types.CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []types.CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *types.Quantity
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *types.Decimal
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []types.PositiveInt
}

func (s ClaimResponseAddItemDetailSubDetail) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The second-tier service adjudications for payor added services.
type ClaimResponseAddItemDetail struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService types.CodeableConcept
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
	// The quantity times the unit price for an additional service or product or charge.
	Net *types.Money
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []types.PositiveInt
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []types.CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *types.Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *types.Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *types.Decimal
	// The third-tier service adjudications for payor added services.
	SubDetail []ClaimResponseAddItemDetailSubDetail
}

func (s ClaimResponseAddItemDetail) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The first-tier service adjudications for payor added product or service lines.
type ClaimResponseAddItem struct {
	// The providers who are authorized for the services rendered to the patient.
	Provider []types.Reference
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []types.CodeableConcept
	// Where the product or service was provided.
	Location r4.Element
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *types.Money
	// Physical service site on the patient (limb, tooth, etc.).
	BodySite *types.CodeableConcept
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []types.PositiveInt
	// Claim items which this service line is intended to replace.
	ItemSequence []types.PositiveInt
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// The sequence number of the details within the claim item which this line is intended to replace.
	DetailSequence []types.PositiveInt
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *types.Decimal
	// A region or surface of the bodySite, e.g. limb region or tooth surface(s).
	SubSite []types.CodeableConcept
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// The number of repetitions of a service or product.
	Quantity *types.Quantity
	// The second-tier service adjudications for payor added services.
	Detail []ClaimResponseAddItemDetail
	// The date or dates when the service or product was supplied, performed or completed.
	Serviced r4.Element
	// The sequence number of the sub-details within the details within the claim item which this line is intended to replace.
	SubdetailSequence []types.PositiveInt
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService types.CodeableConcept
	// Identifies the program under which this may be recovered.
	ProgramCode []types.CodeableConcept
	// The quantity times the unit price for an additional service or product or charge.
	Net *types.Money
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
}

func (s ClaimResponseAddItem) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Financial instruments for reimbursement for the health care products and services specified on the claim.
type ClaimResponseInsurance struct {
	// A flag to indicate that this Coverage is to be used for adjudication of this claim when set to true.
	Focal types.Boolean
	// Reference to the insurance card level information contained in the Coverage resource. The coverage issuing insurer will use these details to locate the patient's actual coverage within the insurer's information system.
	Coverage types.Reference
	// A business agreement number established between the provider and the insurer for special business processing purposes.
	BusinessArrangement *types.String
	// The result of the adjudication of the line items for the Coverage specified in this insurance.
	ClaimResponse *types.Reference
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A number to uniquely identify insurance entries and provide a sequence of coverages to convey coordination of benefit order.
	Sequence types.PositiveInt
}

func (s ClaimResponseInsurance) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Payment details for the adjudication of the claim.
type ClaimResponsePayment struct {
	// Issuer's unique identifier for the payment instrument.
	Identifier *types.Identifier
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Total amount of all adjustments to this payment included in this transaction which are not related to this claim's adjudication.
	Adjustment *types.Money
	// Reason for the payment adjustment.
	AdjustmentReason *types.CodeableConcept
	// Benefits payable less any payment adjustment.
	Amount types.Money
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Whether this represents partial or complete payment of the benefits payable.
	Type types.CodeableConcept
	// Estimated date the payment will be issued or the actual issue date of payment.
	Date *types.Date
}

func (s ClaimResponsePayment) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A note that describes or explains adjudication results in a human readable form.
type ClaimResponseProcessNote struct {
	// The explanation or description associated with the processing.
	Text types.String
	// A code to define the language used in the text of the note.
	Language *types.CodeableConcept
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A number to uniquely identify a note entry.
	Number *types.PositiveInt
	// The business purpose of the note text.
	Type *types.Code
}

func (s ClaimResponseProcessNote) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse struct {
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
	Type types.CodeableConcept
	// Original request resource reference.
	Request *types.Reference
	// Type of Party to be reimbursed: subscriber, provider, other.
	PayeeType *types.CodeableConcept
	// The adjudication results which are presented at the header level rather than at the line-item or add-item levels.
	Adjudication []ClaimResponseItemAdjudication
	// A code, used only on a response to a preauthorization, to indicate whether the benefits payable have been reserved and for whom.
	FundsReserve *types.CodeableConcept
	// Errors encountered during the processing of the adjudication.
	Error []ClaimResponseError
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The date this resource was created.
	Created types.DateTime
	// The party responsible for authorization, adjudication and reimbursement.
	Insurer types.Reference
	// The actual form, by reference or inclusion, for printing the content or an EOB.
	Form *types.Attachment
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// The outcome of the claim, predetermination, or preauthorization processing.
	Outcome types.Code
	// Categorized monetary totals for the adjudication.
	Total []ClaimResponseTotal
	// A code for the form to be used for printing the content.
	FormCode *types.CodeableConcept
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// Reference from the Insurer which is used in later communications which refers to this adjudication.
	PreAuthRef *types.String
	// A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
	Item []ClaimResponseItem
	// The status of the resource instance.
	Status types.Code
	// A code to indicate whether the nature of the request is: to request adjudication of products and services previously rendered; or requesting authorization and adjudication for provision in the future; or requesting the non-binding adjudication of the listed products and services which could be provided in the future.
	Use types.Code
	// The first-tier service adjudications for payor added product or service lines.
	AddItem []ClaimResponseAddItem
	// Financial instruments for reimbursement for the health care products and services specified on the claim.
	Insurance []ClaimResponseInsurance
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// Payment details for the adjudication of the claim.
	Payment *ClaimResponsePayment
	// The party to whom the professional services and/or products have been supplied or are being considered and for whom actual for facast reimbursement is sought.
	Patient types.Reference
	// A unique identifier assigned to this claim response.
	Identifier []types.Identifier
	// A human readable description of the status of the adjudication.
	Disposition *types.String
	// A note that describes or explains adjudication results in a human readable form.
	ProcessNote []ClaimResponseProcessNote
	// Request for additional supporting or authorizing information.
	CommunicationRequest []types.Reference
	// The base language in which the resource is written.
	Language *types.Code
	// The provider which is responsible for the claim, predetermination or preauthorization.
	Requestor *types.Reference
	// The time frame during which this authorization is effective.
	PreAuthPeriod *types.Period
	// A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
	SubType *types.CodeableConcept
}

func (s ClaimResponse) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
