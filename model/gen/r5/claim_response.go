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

// This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse struct {
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
	// A unique identifier assigned to this claim response.
	Identifier []Identifier
	// Trace number for tracking purposes. May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier
	// The status of the resource instance.
	Status Code
	// A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
	Type CodeableConcept
	// A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
	SubType *CodeableConcept
	// A code to indicate whether the nature of the request is: Claim - A request to an Insurer to adjudicate the supplied charges for health care goods and services under the identified policy and to pay the determined Benefit amount, if any; Preauthorization - A request to an Insurer to adjudicate the supplied proposed future charges for health care goods and services under the identified policy and to approve the services and provide the expected benefit amounts and potentially to reserve funds to pay the benefits when Claims for the indicated services are later submitted; or, Pre-determination - A request to an Insurer to adjudicate the supplied 'what if' charges for health care goods and services under the identified policy and report back what the Benefit payable would be had the services actually been provided.
	Use Code
	// The party to whom the professional services and/or products have been supplied or are being considered and for whom actual for facast reimbursement is sought.
	Patient Reference
	// The date this resource was created.
	Created DateTime
	// The party responsible for authorization, adjudication and reimbursement.
	Insurer *Reference
	// The provider which is responsible for the claim, predetermination or preauthorization.
	Requestor *Reference
	// Original request resource reference.
	Request *Reference
	// The outcome of the claim, predetermination, or preauthorization processing.
	Outcome Code
	// The result of the claim, predetermination, or preauthorization adjudication.
	Decision *CodeableConcept
	// A human readable description of the status of the adjudication.
	Disposition *String
	// Reference from the Insurer which is used in later communications which refers to this adjudication.
	PreAuthRef *String
	// The time frame during which this authorization is effective.
	PreAuthPeriod *Period
	// Information code for an event with a corresponding date or period.
	Event []ClaimResponseEvent
	// Type of Party to be reimbursed: subscriber, provider, other.
	PayeeType *CodeableConcept
	// Healthcare encounters related to this claim.
	Encounter []Reference
	// A package billing code or bundle code used to group products and services to a particular health condition (such as heart attack) which is based on a predetermined grouping code system.
	DiagnosisRelatedGroup *CodeableConcept
	// A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
	Item []ClaimResponseItem
	// The first-tier service adjudications for payor added product or service lines.
	AddItem []ClaimResponseAddItem
	// The adjudication results which are presented at the header level rather than at the line-item or add-item levels.
	Adjudication []ClaimResponseItemAdjudication
	// Categorized monetary totals for the adjudication.
	Total []ClaimResponseTotal
	// Payment details for the adjudication of the claim.
	Payment *ClaimResponsePayment
	// A code, used only on a response to a preauthorization, to indicate whether the benefits payable have been reserved and for whom.
	FundsReserve *CodeableConcept
	// A code for the form to be used for printing the content.
	FormCode *CodeableConcept
	// The actual form, by reference or inclusion, for printing the content or an EOB.
	Form *Attachment
	// A note that describes or explains adjudication results in a human readable form.
	ProcessNote []ClaimResponseProcessNote
	// Request for additional supporting or authorizing information.
	CommunicationRequest []Reference
	// Financial instruments for reimbursement for the health care products and services specified on the claim.
	Insurance []ClaimResponseInsurance
	// Errors encountered during the processing of the adjudication.
	Error []ClaimResponseError
}

// Information code for an event with a corresponding date or period.
type ClaimResponseEvent struct {
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
	When isClaimResponseEventWhen
}
type isClaimResponseEventWhen interface {
	model.Element
	isClaimResponseEventWhen()
}

func (r DateTime) isClaimResponseEventWhen() {}
func (r Period) isClaimResponseEventWhen()   {}

// A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
type ClaimResponseItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely reference the claim item entries.
	ItemSequence PositiveInt
	// Trace number for tracking purposes. May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The high-level results of the adjudication if adjudication has been performed.
	ReviewOutcome *ClaimResponseItemReviewOutcome
	// If this item is a group then the values here are a summary of the adjudication of the detail items. If this item is a simple product or service then this is the result of the adjudication of this item.
	Adjudication []ClaimResponseItemAdjudication
	// A claim detail. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
	Detail []ClaimResponseItemDetail
}

// The high-level results of the adjudication if adjudication has been performed.
type ClaimResponseItemReviewOutcome struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The result of the claim, predetermination, or preauthorization adjudication.
	Decision *CodeableConcept
	// The reasons for the result of the claim, predetermination, or preauthorization adjudication.
	Reason []CodeableConcept
	// Reference from the Insurer which is used in later communications which refers to this adjudication.
	PreAuthRef *String
	// The time frame during which this authorization is effective.
	PreAuthPeriod *Period
}

// If this item is a group then the values here are a summary of the adjudication of the detail items. If this item is a simple product or service then this is the result of the adjudication of this item.
type ClaimResponseItemAdjudication struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code to indicate the information type of this adjudication record. Information types may include the value submitted, maximum values or percentages allowed or payable under the plan, amounts that: the patient is responsible for in aggregate or pertaining to this item; amounts paid by other coverages; and, the benefit payable for this item.
	Category CodeableConcept
	// A code supporting the understanding of the adjudication result and explaining variance from expected amount.
	Reason *CodeableConcept
	// Monetary amount associated with the category.
	Amount *Money
	// A non-monetary value associated with the category. Mutually exclusive to the amount element above.
	Quantity *Quantity
}

// A claim detail. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
type ClaimResponseItemDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely reference the claim detail entry.
	DetailSequence PositiveInt
	// Trace number for tracking purposes. May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The high-level results of the adjudication if adjudication has been performed.
	ReviewOutcome *ClaimResponseItemReviewOutcome
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
	// A sub-detail adjudication of a simple product or service.
	SubDetail []ClaimResponseItemDetailSubDetail
}

// A sub-detail adjudication of a simple product or service.
type ClaimResponseItemDetailSubDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely reference the claim sub-detail entry.
	SubDetailSequence PositiveInt
	// Trace number for tracking purposes. May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The high-level results of the adjudication if adjudication has been performed.
	ReviewOutcome *ClaimResponseItemReviewOutcome
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
}

// The first-tier service adjudications for payor added product or service lines.
type ClaimResponseAddItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Claim items which this service line is intended to replace.
	ItemSequence []PositiveInt
	// The sequence number of the details within the claim item which this line is intended to replace.
	DetailSequence []PositiveInt
	// The sequence number of the sub-details within the details within the claim item which this line is intended to replace.
	SubdetailSequence []PositiveInt
	// Trace number for tracking purposes. May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier
	// The providers who are authorized for the services rendered to the patient.
	Provider []Reference
	// The type of revenue or cost center providing the product and/or service.
	Revenue *CodeableConcept
	// When the value is a group code then this item collects a set of related item details, otherwise this contains the product, service, drug or other billing code for the item. This element may be the start of a range of .productOrService codes used in conjunction with .productOrServiceEnd or it may be a solo element where .productOrServiceEnd is not used.
	ProductOrService *CodeableConcept
	// This contains the end of a range of product, service, drug or other billing codes for the item. This element is not used when the .productOrService is a group code. This value may only be present when a .productOfService code has been provided to convey the start of the range. Typically this value may be used only with preauthorizations and not with claims.
	ProductOrServiceEnd *CodeableConcept
	// Request or Referral for Goods or Service to be rendered.
	Request []Reference
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []CodeableConcept
	// Identifies the program under which this may be recovered.
	ProgramCode []CodeableConcept
	// The date or dates when the service or product was supplied, performed or completed.
	Serviced isClaimResponseAddItemServiced
	// Where the product or service was provided.
	Location isClaimResponseAddItemLocation
	// The number of repetitions of a service or product.
	Quantity *Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *Decimal
	// The total of taxes applicable for this product or service.
	Tax *Money
	// The total amount claimed for the group (if a grouper) or the addItem. Net = unit price * quantity * factor.
	Net *Money
	// Physical location where the service is performed or applies.
	BodySite []ClaimResponseAddItemBodySite
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The high-level results of the adjudication if adjudication has been performed.
	ReviewOutcome *ClaimResponseItemReviewOutcome
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
	// The second-tier service adjudications for payor added services.
	Detail []ClaimResponseAddItemDetail
}
type isClaimResponseAddItemServiced interface {
	model.Element
	isClaimResponseAddItemServiced()
}

func (r Date) isClaimResponseAddItemServiced()   {}
func (r Period) isClaimResponseAddItemServiced() {}

type isClaimResponseAddItemLocation interface {
	model.Element
	isClaimResponseAddItemLocation()
}

func (r CodeableConcept) isClaimResponseAddItemLocation() {}
func (r Address) isClaimResponseAddItemLocation()         {}
func (r Reference) isClaimResponseAddItemLocation()       {}

// Physical location where the service is performed or applies.
type ClaimResponseAddItemBodySite struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Physical service site on the patient (limb, tooth, etc.).
	Site []CodeableReference
	// A region or surface of the bodySite, e.g. limb region or tooth surface(s).
	SubSite []CodeableConcept
}

// The second-tier service adjudications for payor added services.
type ClaimResponseAddItemDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Trace number for tracking purposes. May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier
	// The type of revenue or cost center providing the product and/or service.
	Revenue *CodeableConcept
	// When the value is a group code then this item collects a set of related item details, otherwise this contains the product, service, drug or other billing code for the item. This element may be the start of a range of .productOrService codes used in conjunction with .productOrServiceEnd or it may be a solo element where .productOrServiceEnd is not used.
	ProductOrService *CodeableConcept
	// This contains the end of a range of product, service, drug or other billing codes for the item. This element is not used when the .productOrService is a group code. This value may only be present when a .productOfService code has been provided to convey the start of the range. Typically this value may be used only with preauthorizations and not with claims.
	ProductOrServiceEnd *CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *Decimal
	// The total of taxes applicable for this product or service.
	Tax *Money
	// The total amount claimed for the group (if a grouper) or the addItem.detail. Net = unit price * quantity * factor.
	Net *Money
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The high-level results of the adjudication if adjudication has been performed.
	ReviewOutcome *ClaimResponseItemReviewOutcome
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
	// The third-tier service adjudications for payor added services.
	SubDetail []ClaimResponseAddItemDetailSubDetail
}

// The third-tier service adjudications for payor added services.
type ClaimResponseAddItemDetailSubDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Trace number for tracking purposes. May be defined at the jurisdiction level or between trading partners.
	TraceNumber []Identifier
	// The type of revenue or cost center providing the product and/or service.
	Revenue *CodeableConcept
	// When the value is a group code then this item collects a set of related item details, otherwise this contains the product, service, drug or other billing code for the item. This element may be the start of a range of .productOrService codes used in conjunction with .productOrServiceEnd or it may be a solo element where .productOrServiceEnd is not used.
	ProductOrService *CodeableConcept
	// This contains the end of a range of product, service, drug or other billing codes for the item. This element is not used when the .productOrService is a group code. This value may only be present when a .productOfService code has been provided to convey the start of the range. Typically this value may be used only with preauthorizations and not with claims.
	ProductOrServiceEnd *CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *Decimal
	// The total of taxes applicable for this product or service.
	Tax *Money
	// The total amount claimed for the addItem.detail.subDetail. Net = unit price * quantity * factor.
	Net *Money
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The high-level results of the adjudication if adjudication has been performed.
	ReviewOutcome *ClaimResponseItemReviewOutcome
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
}

// Categorized monetary totals for the adjudication.
type ClaimResponseTotal struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code to indicate the information type of this adjudication record. Information types may include: the value submitted, maximum values or percentages allowed or payable under the plan, amounts that the patient is responsible for in aggregate or pertaining to this item, amounts paid by other coverages, and the benefit payable for this item.
	Category CodeableConcept
	// Monetary total amount associated with the category.
	Amount Money
}

// Payment details for the adjudication of the claim.
type ClaimResponsePayment struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Whether this represents partial or complete payment of the benefits payable.
	Type CodeableConcept
	// Total amount of all adjustments to this payment included in this transaction which are not related to this claim's adjudication.
	Adjustment *Money
	// Reason for the payment adjustment.
	AdjustmentReason *CodeableConcept
	// Estimated date the payment will be issued or the actual issue date of payment.
	Date *Date
	// Benefits payable less any payment adjustment.
	Amount Money
	// Issuer's unique identifier for the payment instrument.
	Identifier *Identifier
}

// A note that describes or explains adjudication results in a human readable form.
type ClaimResponseProcessNote struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely identify a note entry.
	Number *PositiveInt
	// The business purpose of the note text.
	Type *CodeableConcept
	// The explanation or description associated with the processing.
	Text String
	// A code to define the language used in the text of the note.
	Language *CodeableConcept
}

// Financial instruments for reimbursement for the health care products and services specified on the claim.
type ClaimResponseInsurance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely identify insurance entries and provide a sequence of coverages to convey coordination of benefit order.
	Sequence PositiveInt
	// A flag to indicate that this Coverage is to be used for adjudication of this claim when set to true.
	Focal Boolean
	// Reference to the insurance card level information contained in the Coverage resource. The coverage issuing insurer will use these details to locate the patient's actual coverage within the insurer's information system.
	Coverage Reference
	// A business agreement number established between the provider and the insurer for special business processing purposes.
	BusinessArrangement *String
	// The result of the adjudication of the line items for the Coverage specified in this insurance.
	ClaimResponse *Reference
}

// Errors encountered during the processing of the adjudication.
type ClaimResponseError struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The sequence number of the line item submitted which contains the error. This value is omitted when the error occurs outside of the item structure.
	ItemSequence *PositiveInt
	// The sequence number of the detail within the line item submitted which contains the error. This value is omitted when the error occurs outside of the item structure.
	DetailSequence *PositiveInt
	// The sequence number of the sub-detail within the detail within the line item submitted which contains the error. This value is omitted when the error occurs outside of the item structure.
	SubDetailSequence *PositiveInt
	// An error code, from a specified code system, which details why the claim could not be adjudicated.
	Code CodeableConcept
	// A [simple subset of FHIRPath](fhirpath.html#simple) limited to element names, repetition indicators and the default child accessor that identifies one of the elements in the resource that caused this issue to be raised.
	Expression []String
}

func (r ClaimResponse) ResourceType() string {
	return "ClaimResponse"
}
func (r ClaimResponse) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r ClaimResponse) MemSize() int {
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
	for _, i := range r.TraceNumber {
		s += i.MemSize()
	}
	s += (cap(r.TraceNumber) - len(r.TraceNumber)) * int(unsafe.Sizeof(Identifier{}))
	s += r.Status.MemSize() - int(unsafe.Sizeof(r.Status))
	s += r.Type.MemSize() - int(unsafe.Sizeof(r.Type))
	if r.SubType != nil {
		s += r.SubType.MemSize()
	}
	s += r.Use.MemSize() - int(unsafe.Sizeof(r.Use))
	s += r.Patient.MemSize() - int(unsafe.Sizeof(r.Patient))
	s += r.Created.MemSize() - int(unsafe.Sizeof(r.Created))
	if r.Insurer != nil {
		s += r.Insurer.MemSize()
	}
	if r.Requestor != nil {
		s += r.Requestor.MemSize()
	}
	if r.Request != nil {
		s += r.Request.MemSize()
	}
	s += r.Outcome.MemSize() - int(unsafe.Sizeof(r.Outcome))
	if r.Decision != nil {
		s += r.Decision.MemSize()
	}
	if r.Disposition != nil {
		s += r.Disposition.MemSize()
	}
	if r.PreAuthRef != nil {
		s += r.PreAuthRef.MemSize()
	}
	if r.PreAuthPeriod != nil {
		s += r.PreAuthPeriod.MemSize()
	}
	for _, i := range r.Event {
		s += i.MemSize()
	}
	s += (cap(r.Event) - len(r.Event)) * int(unsafe.Sizeof(ClaimResponseEvent{}))
	if r.PayeeType != nil {
		s += r.PayeeType.MemSize()
	}
	for _, i := range r.Encounter {
		s += i.MemSize()
	}
	s += (cap(r.Encounter) - len(r.Encounter)) * int(unsafe.Sizeof(Reference{}))
	if r.DiagnosisRelatedGroup != nil {
		s += r.DiagnosisRelatedGroup.MemSize()
	}
	for _, i := range r.Item {
		s += i.MemSize()
	}
	s += (cap(r.Item) - len(r.Item)) * int(unsafe.Sizeof(ClaimResponseItem{}))
	for _, i := range r.AddItem {
		s += i.MemSize()
	}
	s += (cap(r.AddItem) - len(r.AddItem)) * int(unsafe.Sizeof(ClaimResponseAddItem{}))
	for _, i := range r.Adjudication {
		s += i.MemSize()
	}
	s += (cap(r.Adjudication) - len(r.Adjudication)) * int(unsafe.Sizeof(ClaimResponseItemAdjudication{}))
	for _, i := range r.Total {
		s += i.MemSize()
	}
	s += (cap(r.Total) - len(r.Total)) * int(unsafe.Sizeof(ClaimResponseTotal{}))
	if r.Payment != nil {
		s += r.Payment.MemSize()
	}
	if r.FundsReserve != nil {
		s += r.FundsReserve.MemSize()
	}
	if r.FormCode != nil {
		s += r.FormCode.MemSize()
	}
	if r.Form != nil {
		s += r.Form.MemSize()
	}
	for _, i := range r.ProcessNote {
		s += i.MemSize()
	}
	s += (cap(r.ProcessNote) - len(r.ProcessNote)) * int(unsafe.Sizeof(ClaimResponseProcessNote{}))
	for _, i := range r.CommunicationRequest {
		s += i.MemSize()
	}
	s += (cap(r.CommunicationRequest) - len(r.CommunicationRequest)) * int(unsafe.Sizeof(Reference{}))
	for _, i := range r.Insurance {
		s += i.MemSize()
	}
	s += (cap(r.Insurance) - len(r.Insurance)) * int(unsafe.Sizeof(ClaimResponseInsurance{}))
	for _, i := range r.Error {
		s += i.MemSize()
	}
	s += (cap(r.Error) - len(r.Error)) * int(unsafe.Sizeof(ClaimResponseError{}))
	return s
}
func (r ClaimResponseEvent) MemSize() int {
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
func (r ClaimResponseItem) MemSize() int {
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
	s += r.ItemSequence.MemSize() - int(unsafe.Sizeof(r.ItemSequence))
	for _, i := range r.TraceNumber {
		s += i.MemSize()
	}
	s += (cap(r.TraceNumber) - len(r.TraceNumber)) * int(unsafe.Sizeof(Identifier{}))
	for _, i := range r.NoteNumber {
		s += i.MemSize()
	}
	s += (cap(r.NoteNumber) - len(r.NoteNumber)) * int(unsafe.Sizeof(PositiveInt{}))
	if r.ReviewOutcome != nil {
		s += r.ReviewOutcome.MemSize()
	}
	for _, i := range r.Adjudication {
		s += i.MemSize()
	}
	s += (cap(r.Adjudication) - len(r.Adjudication)) * int(unsafe.Sizeof(ClaimResponseItemAdjudication{}))
	for _, i := range r.Detail {
		s += i.MemSize()
	}
	s += (cap(r.Detail) - len(r.Detail)) * int(unsafe.Sizeof(ClaimResponseItemDetail{}))
	return s
}
func (r ClaimResponseItemReviewOutcome) MemSize() int {
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
	if r.Decision != nil {
		s += r.Decision.MemSize()
	}
	for _, i := range r.Reason {
		s += i.MemSize()
	}
	s += (cap(r.Reason) - len(r.Reason)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.PreAuthRef != nil {
		s += r.PreAuthRef.MemSize()
	}
	if r.PreAuthPeriod != nil {
		s += r.PreAuthPeriod.MemSize()
	}
	return s
}
func (r ClaimResponseItemAdjudication) MemSize() int {
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
	s += r.Category.MemSize() - int(unsafe.Sizeof(r.Category))
	if r.Reason != nil {
		s += r.Reason.MemSize()
	}
	if r.Amount != nil {
		s += r.Amount.MemSize()
	}
	if r.Quantity != nil {
		s += r.Quantity.MemSize()
	}
	return s
}
func (r ClaimResponseItemDetail) MemSize() int {
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
	s += r.DetailSequence.MemSize() - int(unsafe.Sizeof(r.DetailSequence))
	for _, i := range r.TraceNumber {
		s += i.MemSize()
	}
	s += (cap(r.TraceNumber) - len(r.TraceNumber)) * int(unsafe.Sizeof(Identifier{}))
	for _, i := range r.NoteNumber {
		s += i.MemSize()
	}
	s += (cap(r.NoteNumber) - len(r.NoteNumber)) * int(unsafe.Sizeof(PositiveInt{}))
	if r.ReviewOutcome != nil {
		s += r.ReviewOutcome.MemSize()
	}
	for _, i := range r.Adjudication {
		s += i.MemSize()
	}
	s += (cap(r.Adjudication) - len(r.Adjudication)) * int(unsafe.Sizeof(ClaimResponseItemAdjudication{}))
	for _, i := range r.SubDetail {
		s += i.MemSize()
	}
	s += (cap(r.SubDetail) - len(r.SubDetail)) * int(unsafe.Sizeof(ClaimResponseItemDetailSubDetail{}))
	return s
}
func (r ClaimResponseItemDetailSubDetail) MemSize() int {
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
	s += r.SubDetailSequence.MemSize() - int(unsafe.Sizeof(r.SubDetailSequence))
	for _, i := range r.TraceNumber {
		s += i.MemSize()
	}
	s += (cap(r.TraceNumber) - len(r.TraceNumber)) * int(unsafe.Sizeof(Identifier{}))
	for _, i := range r.NoteNumber {
		s += i.MemSize()
	}
	s += (cap(r.NoteNumber) - len(r.NoteNumber)) * int(unsafe.Sizeof(PositiveInt{}))
	if r.ReviewOutcome != nil {
		s += r.ReviewOutcome.MemSize()
	}
	for _, i := range r.Adjudication {
		s += i.MemSize()
	}
	s += (cap(r.Adjudication) - len(r.Adjudication)) * int(unsafe.Sizeof(ClaimResponseItemAdjudication{}))
	return s
}
func (r ClaimResponseAddItem) MemSize() int {
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
	for _, i := range r.ItemSequence {
		s += i.MemSize()
	}
	s += (cap(r.ItemSequence) - len(r.ItemSequence)) * int(unsafe.Sizeof(PositiveInt{}))
	for _, i := range r.DetailSequence {
		s += i.MemSize()
	}
	s += (cap(r.DetailSequence) - len(r.DetailSequence)) * int(unsafe.Sizeof(PositiveInt{}))
	for _, i := range r.SubdetailSequence {
		s += i.MemSize()
	}
	s += (cap(r.SubdetailSequence) - len(r.SubdetailSequence)) * int(unsafe.Sizeof(PositiveInt{}))
	for _, i := range r.TraceNumber {
		s += i.MemSize()
	}
	s += (cap(r.TraceNumber) - len(r.TraceNumber)) * int(unsafe.Sizeof(Identifier{}))
	for _, i := range r.Provider {
		s += i.MemSize()
	}
	s += (cap(r.Provider) - len(r.Provider)) * int(unsafe.Sizeof(Reference{}))
	if r.Revenue != nil {
		s += r.Revenue.MemSize()
	}
	if r.ProductOrService != nil {
		s += r.ProductOrService.MemSize()
	}
	if r.ProductOrServiceEnd != nil {
		s += r.ProductOrServiceEnd.MemSize()
	}
	for _, i := range r.Request {
		s += i.MemSize()
	}
	s += (cap(r.Request) - len(r.Request)) * int(unsafe.Sizeof(Reference{}))
	for _, i := range r.Modifier {
		s += i.MemSize()
	}
	s += (cap(r.Modifier) - len(r.Modifier)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.ProgramCode {
		s += i.MemSize()
	}
	s += (cap(r.ProgramCode) - len(r.ProgramCode)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Serviced != nil {
		s += r.Serviced.MemSize()
	}
	if r.Location != nil {
		s += r.Location.MemSize()
	}
	if r.Quantity != nil {
		s += r.Quantity.MemSize()
	}
	if r.UnitPrice != nil {
		s += r.UnitPrice.MemSize()
	}
	if r.Factor != nil {
		s += r.Factor.MemSize()
	}
	if r.Tax != nil {
		s += r.Tax.MemSize()
	}
	if r.Net != nil {
		s += r.Net.MemSize()
	}
	for _, i := range r.BodySite {
		s += i.MemSize()
	}
	s += (cap(r.BodySite) - len(r.BodySite)) * int(unsafe.Sizeof(ClaimResponseAddItemBodySite{}))
	for _, i := range r.NoteNumber {
		s += i.MemSize()
	}
	s += (cap(r.NoteNumber) - len(r.NoteNumber)) * int(unsafe.Sizeof(PositiveInt{}))
	if r.ReviewOutcome != nil {
		s += r.ReviewOutcome.MemSize()
	}
	for _, i := range r.Adjudication {
		s += i.MemSize()
	}
	s += (cap(r.Adjudication) - len(r.Adjudication)) * int(unsafe.Sizeof(ClaimResponseItemAdjudication{}))
	for _, i := range r.Detail {
		s += i.MemSize()
	}
	s += (cap(r.Detail) - len(r.Detail)) * int(unsafe.Sizeof(ClaimResponseAddItemDetail{}))
	return s
}
func (r ClaimResponseAddItemBodySite) MemSize() int {
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
	for _, i := range r.Site {
		s += i.MemSize()
	}
	s += (cap(r.Site) - len(r.Site)) * int(unsafe.Sizeof(CodeableReference{}))
	for _, i := range r.SubSite {
		s += i.MemSize()
	}
	s += (cap(r.SubSite) - len(r.SubSite)) * int(unsafe.Sizeof(CodeableConcept{}))
	return s
}
func (r ClaimResponseAddItemDetail) MemSize() int {
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
	for _, i := range r.TraceNumber {
		s += i.MemSize()
	}
	s += (cap(r.TraceNumber) - len(r.TraceNumber)) * int(unsafe.Sizeof(Identifier{}))
	if r.Revenue != nil {
		s += r.Revenue.MemSize()
	}
	if r.ProductOrService != nil {
		s += r.ProductOrService.MemSize()
	}
	if r.ProductOrServiceEnd != nil {
		s += r.ProductOrServiceEnd.MemSize()
	}
	for _, i := range r.Modifier {
		s += i.MemSize()
	}
	s += (cap(r.Modifier) - len(r.Modifier)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Quantity != nil {
		s += r.Quantity.MemSize()
	}
	if r.UnitPrice != nil {
		s += r.UnitPrice.MemSize()
	}
	if r.Factor != nil {
		s += r.Factor.MemSize()
	}
	if r.Tax != nil {
		s += r.Tax.MemSize()
	}
	if r.Net != nil {
		s += r.Net.MemSize()
	}
	for _, i := range r.NoteNumber {
		s += i.MemSize()
	}
	s += (cap(r.NoteNumber) - len(r.NoteNumber)) * int(unsafe.Sizeof(PositiveInt{}))
	if r.ReviewOutcome != nil {
		s += r.ReviewOutcome.MemSize()
	}
	for _, i := range r.Adjudication {
		s += i.MemSize()
	}
	s += (cap(r.Adjudication) - len(r.Adjudication)) * int(unsafe.Sizeof(ClaimResponseItemAdjudication{}))
	for _, i := range r.SubDetail {
		s += i.MemSize()
	}
	s += (cap(r.SubDetail) - len(r.SubDetail)) * int(unsafe.Sizeof(ClaimResponseAddItemDetailSubDetail{}))
	return s
}
func (r ClaimResponseAddItemDetailSubDetail) MemSize() int {
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
	for _, i := range r.TraceNumber {
		s += i.MemSize()
	}
	s += (cap(r.TraceNumber) - len(r.TraceNumber)) * int(unsafe.Sizeof(Identifier{}))
	if r.Revenue != nil {
		s += r.Revenue.MemSize()
	}
	if r.ProductOrService != nil {
		s += r.ProductOrService.MemSize()
	}
	if r.ProductOrServiceEnd != nil {
		s += r.ProductOrServiceEnd.MemSize()
	}
	for _, i := range r.Modifier {
		s += i.MemSize()
	}
	s += (cap(r.Modifier) - len(r.Modifier)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Quantity != nil {
		s += r.Quantity.MemSize()
	}
	if r.UnitPrice != nil {
		s += r.UnitPrice.MemSize()
	}
	if r.Factor != nil {
		s += r.Factor.MemSize()
	}
	if r.Tax != nil {
		s += r.Tax.MemSize()
	}
	if r.Net != nil {
		s += r.Net.MemSize()
	}
	for _, i := range r.NoteNumber {
		s += i.MemSize()
	}
	s += (cap(r.NoteNumber) - len(r.NoteNumber)) * int(unsafe.Sizeof(PositiveInt{}))
	if r.ReviewOutcome != nil {
		s += r.ReviewOutcome.MemSize()
	}
	for _, i := range r.Adjudication {
		s += i.MemSize()
	}
	s += (cap(r.Adjudication) - len(r.Adjudication)) * int(unsafe.Sizeof(ClaimResponseItemAdjudication{}))
	return s
}
func (r ClaimResponseTotal) MemSize() int {
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
	s += r.Category.MemSize() - int(unsafe.Sizeof(r.Category))
	s += r.Amount.MemSize() - int(unsafe.Sizeof(r.Amount))
	return s
}
func (r ClaimResponsePayment) MemSize() int {
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
	if r.Adjustment != nil {
		s += r.Adjustment.MemSize()
	}
	if r.AdjustmentReason != nil {
		s += r.AdjustmentReason.MemSize()
	}
	if r.Date != nil {
		s += r.Date.MemSize()
	}
	s += r.Amount.MemSize() - int(unsafe.Sizeof(r.Amount))
	if r.Identifier != nil {
		s += r.Identifier.MemSize()
	}
	return s
}
func (r ClaimResponseProcessNote) MemSize() int {
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
	if r.Number != nil {
		s += r.Number.MemSize()
	}
	if r.Type != nil {
		s += r.Type.MemSize()
	}
	s += r.Text.MemSize() - int(unsafe.Sizeof(r.Text))
	if r.Language != nil {
		s += r.Language.MemSize()
	}
	return s
}
func (r ClaimResponseInsurance) MemSize() int {
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
	s += r.Sequence.MemSize() - int(unsafe.Sizeof(r.Sequence))
	s += r.Focal.MemSize() - int(unsafe.Sizeof(r.Focal))
	s += r.Coverage.MemSize() - int(unsafe.Sizeof(r.Coverage))
	if r.BusinessArrangement != nil {
		s += r.BusinessArrangement.MemSize()
	}
	if r.ClaimResponse != nil {
		s += r.ClaimResponse.MemSize()
	}
	return s
}
func (r ClaimResponseError) MemSize() int {
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
	if r.ItemSequence != nil {
		s += r.ItemSequence.MemSize()
	}
	if r.DetailSequence != nil {
		s += r.DetailSequence.MemSize()
	}
	if r.SubDetailSequence != nil {
		s += r.SubDetailSequence.MemSize()
	}
	s += r.Code.MemSize() - int(unsafe.Sizeof(r.Code))
	for _, i := range r.Expression {
		s += i.MemSize()
	}
	s += (cap(r.Expression) - len(r.Expression)) * int(unsafe.Sizeof(String{}))
	return s
}
func (r ClaimResponse) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponseEvent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponseItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponseItemReviewOutcome) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponseItemAdjudication) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponseItemDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponseItemDetailSubDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponseAddItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponseAddItemBodySite) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponseAddItemDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponseAddItemDetailSubDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponseTotal) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponsePayment) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponseProcessNote) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponseInsurance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponseError) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ClaimResponse) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponse) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"ClaimResponse\""))
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
	if len(r.TraceNumber) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"traceNumber\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.TraceNumber {
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
	if r.SubType != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"subType\":"))
		if err != nil {
			return err
		}
		err = r.SubType.marshalJSON(w)
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
		_, err = w.Write([]byte("\"use\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Use)
		if err != nil {
			return err
		}
	}
	if r.Use.Id != nil || r.Use.Extension != nil {
		p := primitiveElement{Id: r.Use.Id, Extension: r.Use.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_use\":"))
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
	_, err = w.Write([]byte("\"patient\":"))
	if err != nil {
		return err
	}
	err = r.Patient.marshalJSON(w)
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
	if r.Insurer != nil {
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
	if r.Request != nil {
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
	if r.Decision != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"decision\":"))
		if err != nil {
			return err
		}
		err = r.Decision.marshalJSON(w)
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
	if r.PreAuthPeriod != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"preAuthPeriod\":"))
		if err != nil {
			return err
		}
		err = r.PreAuthPeriod.marshalJSON(w)
		if err != nil {
			return err
		}
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
	if r.PayeeType != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"payeeType\":"))
		if err != nil {
			return err
		}
		err = r.PayeeType.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Encounter) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"encounter\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Encounter {
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
	if r.DiagnosisRelatedGroup != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"diagnosisRelatedGroup\":"))
		if err != nil {
			return err
		}
		err = r.DiagnosisRelatedGroup.marshalJSON(w)
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
	if len(r.AddItem) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"addItem\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.AddItem {
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
	if len(r.Adjudication) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"adjudication\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Adjudication {
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
	if len(r.Total) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"total\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Total {
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
	if r.Payment != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"payment\":"))
		if err != nil {
			return err
		}
		err = r.Payment.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.FundsReserve != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"fundsReserve\":"))
		if err != nil {
			return err
		}
		err = r.FundsReserve.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.FormCode != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"formCode\":"))
		if err != nil {
			return err
		}
		err = r.FormCode.marshalJSON(w)
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
	if len(r.ProcessNote) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"processNote\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ProcessNote {
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
	if len(r.CommunicationRequest) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"communicationRequest\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.CommunicationRequest {
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
func (r ClaimResponseEvent) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponseEvent) marshalJSON(w io.Writer) error {
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
func (r ClaimResponseItem) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponseItem) marshalJSON(w io.Writer) error {
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
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"itemSequence\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ItemSequence)
		if err != nil {
			return err
		}
	}
	if r.ItemSequence.Id != nil || r.ItemSequence.Extension != nil {
		p := primitiveElement{Id: r.ItemSequence.Id, Extension: r.ItemSequence.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_itemSequence\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.TraceNumber) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"traceNumber\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.TraceNumber {
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
		anyValue := false
		for _, e := range r.NoteNumber {
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
			_, err = w.Write([]byte("\"noteNumber\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.NoteNumber)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.NoteNumber {
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
			_, err = w.Write([]byte("\"_noteNumber\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.NoteNumber {
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
	if r.ReviewOutcome != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reviewOutcome\":"))
		if err != nil {
			return err
		}
		err = r.ReviewOutcome.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Adjudication) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"adjudication\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Adjudication {
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
	if len(r.Detail) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"detail\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Detail {
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
func (r ClaimResponseItemReviewOutcome) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponseItemReviewOutcome) marshalJSON(w io.Writer) error {
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
	if r.Decision != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"decision\":"))
		if err != nil {
			return err
		}
		err = r.Decision.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Reason) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reason\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Reason {
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
	if r.PreAuthPeriod != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"preAuthPeriod\":"))
		if err != nil {
			return err
		}
		err = r.PreAuthPeriod.marshalJSON(w)
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
func (r ClaimResponseItemAdjudication) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponseItemAdjudication) marshalJSON(w io.Writer) error {
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
	_, err = w.Write([]byte("\"category\":"))
	if err != nil {
		return err
	}
	err = r.Category.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.Reason != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reason\":"))
		if err != nil {
			return err
		}
		err = r.Reason.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Amount != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"amount\":"))
		if err != nil {
			return err
		}
		err = r.Amount.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Quantity != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"quantity\":"))
		if err != nil {
			return err
		}
		err = r.Quantity.marshalJSON(w)
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
func (r ClaimResponseItemDetail) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponseItemDetail) marshalJSON(w io.Writer) error {
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
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"detailSequence\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.DetailSequence)
		if err != nil {
			return err
		}
	}
	if r.DetailSequence.Id != nil || r.DetailSequence.Extension != nil {
		p := primitiveElement{Id: r.DetailSequence.Id, Extension: r.DetailSequence.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_detailSequence\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.TraceNumber) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"traceNumber\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.TraceNumber {
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
		anyValue := false
		for _, e := range r.NoteNumber {
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
			_, err = w.Write([]byte("\"noteNumber\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.NoteNumber)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.NoteNumber {
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
			_, err = w.Write([]byte("\"_noteNumber\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.NoteNumber {
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
	if r.ReviewOutcome != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reviewOutcome\":"))
		if err != nil {
			return err
		}
		err = r.ReviewOutcome.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Adjudication) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"adjudication\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Adjudication {
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
	if len(r.SubDetail) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"subDetail\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.SubDetail {
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
func (r ClaimResponseItemDetailSubDetail) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponseItemDetailSubDetail) marshalJSON(w io.Writer) error {
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
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"subDetailSequence\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.SubDetailSequence)
		if err != nil {
			return err
		}
	}
	if r.SubDetailSequence.Id != nil || r.SubDetailSequence.Extension != nil {
		p := primitiveElement{Id: r.SubDetailSequence.Id, Extension: r.SubDetailSequence.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_subDetailSequence\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.TraceNumber) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"traceNumber\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.TraceNumber {
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
		anyValue := false
		for _, e := range r.NoteNumber {
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
			_, err = w.Write([]byte("\"noteNumber\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.NoteNumber)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.NoteNumber {
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
			_, err = w.Write([]byte("\"_noteNumber\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.NoteNumber {
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
	if r.ReviewOutcome != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reviewOutcome\":"))
		if err != nil {
			return err
		}
		err = r.ReviewOutcome.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Adjudication) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"adjudication\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Adjudication {
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
func (r ClaimResponseAddItem) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponseAddItem) marshalJSON(w io.Writer) error {
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
	{
		anyValue := false
		for _, e := range r.ItemSequence {
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
			_, err = w.Write([]byte("\"itemSequence\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.ItemSequence)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.ItemSequence {
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
			_, err = w.Write([]byte("\"_itemSequence\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.ItemSequence {
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
	{
		anyValue := false
		for _, e := range r.DetailSequence {
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
			_, err = w.Write([]byte("\"detailSequence\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.DetailSequence)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.DetailSequence {
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
			_, err = w.Write([]byte("\"_detailSequence\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.DetailSequence {
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
	{
		anyValue := false
		for _, e := range r.SubdetailSequence {
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
			_, err = w.Write([]byte("\"subdetailSequence\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.SubdetailSequence)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.SubdetailSequence {
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
			_, err = w.Write([]byte("\"_subdetailSequence\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.SubdetailSequence {
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
	if len(r.TraceNumber) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"traceNumber\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.TraceNumber {
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
	if len(r.Provider) > 0 {
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
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Provider {
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
	if r.Revenue != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"revenue\":"))
		if err != nil {
			return err
		}
		err = r.Revenue.marshalJSON(w)
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
	if r.ProductOrServiceEnd != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"productOrServiceEnd\":"))
		if err != nil {
			return err
		}
		err = r.ProductOrServiceEnd.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Request) > 0 {
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
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Request {
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
	if len(r.ProgramCode) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"programCode\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ProgramCode {
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
	switch v := r.Location.(type) {
	case CodeableConcept:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"locationCodeableConcept\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *CodeableConcept:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"locationCodeableConcept\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Address:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"locationAddress\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Address:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"locationAddress\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Reference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"locationReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Reference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"locationReference\":"))
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
	if r.Quantity != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"quantity\":"))
		if err != nil {
			return err
		}
		err = r.Quantity.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.UnitPrice != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"unitPrice\":"))
		if err != nil {
			return err
		}
		err = r.UnitPrice.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Factor != nil && r.Factor.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"factor\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Factor)
		if err != nil {
			return err
		}
	}
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		p := primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_factor\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Tax != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"tax\":"))
		if err != nil {
			return err
		}
		err = r.Tax.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Net != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"net\":"))
		if err != nil {
			return err
		}
		err = r.Net.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.BodySite) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"bodySite\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.BodySite {
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
		anyValue := false
		for _, e := range r.NoteNumber {
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
			_, err = w.Write([]byte("\"noteNumber\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.NoteNumber)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.NoteNumber {
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
			_, err = w.Write([]byte("\"_noteNumber\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.NoteNumber {
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
	if r.ReviewOutcome != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reviewOutcome\":"))
		if err != nil {
			return err
		}
		err = r.ReviewOutcome.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Adjudication) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"adjudication\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Adjudication {
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
	if len(r.Detail) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"detail\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Detail {
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
func (r ClaimResponseAddItemBodySite) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponseAddItemBodySite) marshalJSON(w io.Writer) error {
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
	if len(r.Site) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"site\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Site {
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
	if len(r.SubSite) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"subSite\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.SubSite {
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
func (r ClaimResponseAddItemDetail) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponseAddItemDetail) marshalJSON(w io.Writer) error {
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
	if len(r.TraceNumber) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"traceNumber\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.TraceNumber {
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
	if r.Revenue != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"revenue\":"))
		if err != nil {
			return err
		}
		err = r.Revenue.marshalJSON(w)
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
	if r.ProductOrServiceEnd != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"productOrServiceEnd\":"))
		if err != nil {
			return err
		}
		err = r.ProductOrServiceEnd.marshalJSON(w)
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
	if r.Quantity != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"quantity\":"))
		if err != nil {
			return err
		}
		err = r.Quantity.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.UnitPrice != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"unitPrice\":"))
		if err != nil {
			return err
		}
		err = r.UnitPrice.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Factor != nil && r.Factor.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"factor\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Factor)
		if err != nil {
			return err
		}
	}
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		p := primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_factor\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Tax != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"tax\":"))
		if err != nil {
			return err
		}
		err = r.Tax.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Net != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"net\":"))
		if err != nil {
			return err
		}
		err = r.Net.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		anyValue := false
		for _, e := range r.NoteNumber {
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
			_, err = w.Write([]byte("\"noteNumber\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.NoteNumber)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.NoteNumber {
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
			_, err = w.Write([]byte("\"_noteNumber\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.NoteNumber {
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
	if r.ReviewOutcome != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reviewOutcome\":"))
		if err != nil {
			return err
		}
		err = r.ReviewOutcome.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Adjudication) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"adjudication\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Adjudication {
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
	if len(r.SubDetail) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"subDetail\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.SubDetail {
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
func (r ClaimResponseAddItemDetailSubDetail) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponseAddItemDetailSubDetail) marshalJSON(w io.Writer) error {
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
	if len(r.TraceNumber) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"traceNumber\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.TraceNumber {
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
	if r.Revenue != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"revenue\":"))
		if err != nil {
			return err
		}
		err = r.Revenue.marshalJSON(w)
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
	if r.ProductOrServiceEnd != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"productOrServiceEnd\":"))
		if err != nil {
			return err
		}
		err = r.ProductOrServiceEnd.marshalJSON(w)
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
	if r.Quantity != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"quantity\":"))
		if err != nil {
			return err
		}
		err = r.Quantity.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.UnitPrice != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"unitPrice\":"))
		if err != nil {
			return err
		}
		err = r.UnitPrice.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Factor != nil && r.Factor.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"factor\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Factor)
		if err != nil {
			return err
		}
	}
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		p := primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_factor\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Tax != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"tax\":"))
		if err != nil {
			return err
		}
		err = r.Tax.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Net != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"net\":"))
		if err != nil {
			return err
		}
		err = r.Net.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		anyValue := false
		for _, e := range r.NoteNumber {
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
			_, err = w.Write([]byte("\"noteNumber\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.NoteNumber)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.NoteNumber {
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
			_, err = w.Write([]byte("\"_noteNumber\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.NoteNumber {
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
	if r.ReviewOutcome != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reviewOutcome\":"))
		if err != nil {
			return err
		}
		err = r.ReviewOutcome.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Adjudication) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"adjudication\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Adjudication {
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
func (r ClaimResponseTotal) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponseTotal) marshalJSON(w io.Writer) error {
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
	_, err = w.Write([]byte("\"category\":"))
	if err != nil {
		return err
	}
	err = r.Category.marshalJSON(w)
	if err != nil {
		return err
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"amount\":"))
	if err != nil {
		return err
	}
	err = r.Amount.marshalJSON(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r ClaimResponsePayment) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponsePayment) marshalJSON(w io.Writer) error {
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
	if r.Adjustment != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"adjustment\":"))
		if err != nil {
			return err
		}
		err = r.Adjustment.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.AdjustmentReason != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"adjustmentReason\":"))
		if err != nil {
			return err
		}
		err = r.AdjustmentReason.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Date != nil && r.Date.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"date\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Date)
		if err != nil {
			return err
		}
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		p := primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_date\":"))
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
	_, err = w.Write([]byte("\"amount\":"))
	if err != nil {
		return err
	}
	err = r.Amount.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.Identifier != nil {
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
		err = r.Identifier.marshalJSON(w)
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
func (r ClaimResponseProcessNote) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponseProcessNote) marshalJSON(w io.Writer) error {
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
	if r.Number != nil && r.Number.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"number\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Number)
		if err != nil {
			return err
		}
	}
	if r.Number != nil && (r.Number.Id != nil || r.Number.Extension != nil) {
		p := primitiveElement{Id: r.Number.Id, Extension: r.Number.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_number\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Type != nil {
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
	}
	{
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
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Text)
		if err != nil {
			return err
		}
	}
	if r.Text.Id != nil || r.Text.Extension != nil {
		p := primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_text\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Language != nil {
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
		err = r.Language.marshalJSON(w)
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
func (r ClaimResponseInsurance) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponseInsurance) marshalJSON(w io.Writer) error {
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
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sequence\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Sequence)
		if err != nil {
			return err
		}
	}
	if r.Sequence.Id != nil || r.Sequence.Extension != nil {
		p := primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_sequence\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
		_, err = w.Write([]byte("\"focal\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Focal)
		if err != nil {
			return err
		}
	}
	if r.Focal.Id != nil || r.Focal.Extension != nil {
		p := primitiveElement{Id: r.Focal.Id, Extension: r.Focal.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_focal\":"))
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
	_, err = w.Write([]byte("\"coverage\":"))
	if err != nil {
		return err
	}
	err = r.Coverage.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.BusinessArrangement != nil && r.BusinessArrangement.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"businessArrangement\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.BusinessArrangement)
		if err != nil {
			return err
		}
	}
	if r.BusinessArrangement != nil && (r.BusinessArrangement.Id != nil || r.BusinessArrangement.Extension != nil) {
		p := primitiveElement{Id: r.BusinessArrangement.Id, Extension: r.BusinessArrangement.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_businessArrangement\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ClaimResponse != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"claimResponse\":"))
		if err != nil {
			return err
		}
		err = r.ClaimResponse.marshalJSON(w)
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
func (r ClaimResponseError) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ClaimResponseError) marshalJSON(w io.Writer) error {
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
	if r.ItemSequence != nil && r.ItemSequence.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"itemSequence\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ItemSequence)
		if err != nil {
			return err
		}
	}
	if r.ItemSequence != nil && (r.ItemSequence.Id != nil || r.ItemSequence.Extension != nil) {
		p := primitiveElement{Id: r.ItemSequence.Id, Extension: r.ItemSequence.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_itemSequence\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.DetailSequence != nil && r.DetailSequence.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"detailSequence\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.DetailSequence)
		if err != nil {
			return err
		}
	}
	if r.DetailSequence != nil && (r.DetailSequence.Id != nil || r.DetailSequence.Extension != nil) {
		p := primitiveElement{Id: r.DetailSequence.Id, Extension: r.DetailSequence.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_detailSequence\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.SubDetailSequence != nil && r.SubDetailSequence.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"subDetailSequence\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.SubDetailSequence)
		if err != nil {
			return err
		}
	}
	if r.SubDetailSequence != nil && (r.SubDetailSequence.Id != nil || r.SubDetailSequence.Extension != nil) {
		p := primitiveElement{Id: r.SubDetailSequence.Id, Extension: r.SubDetailSequence.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_subDetailSequence\":"))
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
func (r *ClaimResponse) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *ClaimResponse) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponse element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponse element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
			}
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
			}
		case "traceNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.TraceNumber = append(r.TraceNumber, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
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
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = v
		case "subType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.SubType = &v
		case "use":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Use.Value = v.Value
		case "_use":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Use.Id = v.Id
			r.Use.Extension = v.Extension
		case "patient":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Patient = v
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
		case "insurer":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Insurer = &v
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
			r.Request = &v
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
		case "decision":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Decision = &v
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
		case "preAuthPeriod":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.PreAuthPeriod = &v
		case "event":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
			}
			for d.More() {
				var v ClaimResponseEvent
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
			}
		case "payeeType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.PayeeType = &v
		case "encounter":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Encounter = append(r.Encounter, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
			}
		case "diagnosisRelatedGroup":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.DiagnosisRelatedGroup = &v
		case "item":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
			}
			for d.More() {
				var v ClaimResponseItem
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
			}
		case "addItem":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
			}
			for d.More() {
				var v ClaimResponseAddItem
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.AddItem = append(r.AddItem, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
			}
		case "adjudication":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
			}
			for d.More() {
				var v ClaimResponseItemAdjudication
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
			}
		case "total":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
			}
			for d.More() {
				var v ClaimResponseTotal
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Total = append(r.Total, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
			}
		case "payment":
			var v ClaimResponsePayment
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Payment = &v
		case "fundsReserve":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.FundsReserve = &v
		case "formCode":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.FormCode = &v
		case "form":
			var v Attachment
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Form = &v
		case "processNote":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
			}
			for d.More() {
				var v ClaimResponseProcessNote
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ProcessNote = append(r.ProcessNote, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
			}
		case "communicationRequest":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.CommunicationRequest = append(r.CommunicationRequest, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
			}
		case "insurance":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
			}
			for d.More() {
				var v ClaimResponseInsurance
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
			}
		case "error":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponse element", t)
			}
			for d.More() {
				var v ClaimResponseError
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponse element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponse", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponse element", t)
	}
	return nil
}
func (r *ClaimResponseEvent) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponseEvent element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponseEvent element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseEvent element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseEvent element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseEvent element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseEvent element", t)
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
			return fmt.Errorf("invalid field: %s in ClaimResponseEvent", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponseEvent element", t)
	}
	return nil
}
func (r *ClaimResponseItem) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponseItem element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponseItem element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItem element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItem element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItem element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItem element", t)
			}
		case "itemSequence":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.ItemSequence.Value = v.Value
		case "_itemSequence":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ItemSequence.Id = v.Id
			r.ItemSequence.Extension = v.Extension
		case "traceNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItem element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.TraceNumber = append(r.TraceNumber, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItem element", t)
			}
		case "noteNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItem element", t)
			}
			for i := 0; d.More(); i++ {
				var v PositiveInt
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.NoteNumber) <= i {
					r.NoteNumber = append(r.NoteNumber, PositiveInt{})
				}
				r.NoteNumber[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItem element", t)
			}
		case "_noteNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItem element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.NoteNumber) <= i {
					r.NoteNumber = append(r.NoteNumber, PositiveInt{})
				}
				r.NoteNumber[i].Id = v.Id
				r.NoteNumber[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItem element", t)
			}
		case "reviewOutcome":
			var v ClaimResponseItemReviewOutcome
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ReviewOutcome = &v
		case "adjudication":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItem element", t)
			}
			for d.More() {
				var v ClaimResponseItemAdjudication
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItem element", t)
			}
		case "detail":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItem element", t)
			}
			for d.More() {
				var v ClaimResponseItemDetail
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Detail = append(r.Detail, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItem element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponseItem", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponseItem element", t)
	}
	return nil
}
func (r *ClaimResponseItemReviewOutcome) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponseItemReviewOutcome element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponseItemReviewOutcome element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemReviewOutcome element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemReviewOutcome element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemReviewOutcome element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemReviewOutcome element", t)
			}
		case "decision":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Decision = &v
		case "reason":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemReviewOutcome element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Reason = append(r.Reason, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemReviewOutcome element", t)
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
		case "preAuthPeriod":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.PreAuthPeriod = &v
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponseItemReviewOutcome", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponseItemReviewOutcome element", t)
	}
	return nil
}
func (r *ClaimResponseItemAdjudication) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponseItemAdjudication element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponseItemAdjudication element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemAdjudication element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemAdjudication element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemAdjudication element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemAdjudication element", t)
			}
		case "category":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Category = v
		case "reason":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Reason = &v
		case "amount":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Amount = &v
		case "quantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Quantity = &v
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponseItemAdjudication", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponseItemAdjudication element", t)
	}
	return nil
}
func (r *ClaimResponseItemDetail) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponseItemDetail element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponseItemDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemDetail element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemDetail element", t)
			}
		case "detailSequence":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.DetailSequence.Value = v.Value
		case "_detailSequence":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.DetailSequence.Id = v.Id
			r.DetailSequence.Extension = v.Extension
		case "traceNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemDetail element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.TraceNumber = append(r.TraceNumber, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemDetail element", t)
			}
		case "noteNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemDetail element", t)
			}
			for i := 0; d.More(); i++ {
				var v PositiveInt
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.NoteNumber) <= i {
					r.NoteNumber = append(r.NoteNumber, PositiveInt{})
				}
				r.NoteNumber[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemDetail element", t)
			}
		case "_noteNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemDetail element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.NoteNumber) <= i {
					r.NoteNumber = append(r.NoteNumber, PositiveInt{})
				}
				r.NoteNumber[i].Id = v.Id
				r.NoteNumber[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemDetail element", t)
			}
		case "reviewOutcome":
			var v ClaimResponseItemReviewOutcome
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ReviewOutcome = &v
		case "adjudication":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemDetail element", t)
			}
			for d.More() {
				var v ClaimResponseItemAdjudication
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemDetail element", t)
			}
		case "subDetail":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemDetail element", t)
			}
			for d.More() {
				var v ClaimResponseItemDetailSubDetail
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.SubDetail = append(r.SubDetail, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemDetail element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponseItemDetail", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponseItemDetail element", t)
	}
	return nil
}
func (r *ClaimResponseItemDetailSubDetail) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponseItemDetailSubDetail element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponseItemDetailSubDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemDetailSubDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemDetailSubDetail element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemDetailSubDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemDetailSubDetail element", t)
			}
		case "subDetailSequence":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.SubDetailSequence.Value = v.Value
		case "_subDetailSequence":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.SubDetailSequence.Id = v.Id
			r.SubDetailSequence.Extension = v.Extension
		case "traceNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemDetailSubDetail element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.TraceNumber = append(r.TraceNumber, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemDetailSubDetail element", t)
			}
		case "noteNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemDetailSubDetail element", t)
			}
			for i := 0; d.More(); i++ {
				var v PositiveInt
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.NoteNumber) <= i {
					r.NoteNumber = append(r.NoteNumber, PositiveInt{})
				}
				r.NoteNumber[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemDetailSubDetail element", t)
			}
		case "_noteNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemDetailSubDetail element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.NoteNumber) <= i {
					r.NoteNumber = append(r.NoteNumber, PositiveInt{})
				}
				r.NoteNumber[i].Id = v.Id
				r.NoteNumber[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemDetailSubDetail element", t)
			}
		case "reviewOutcome":
			var v ClaimResponseItemReviewOutcome
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ReviewOutcome = &v
		case "adjudication":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseItemDetailSubDetail element", t)
			}
			for d.More() {
				var v ClaimResponseItemAdjudication
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseItemDetailSubDetail element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponseItemDetailSubDetail", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponseItemDetailSubDetail element", t)
	}
	return nil
}
func (r *ClaimResponseAddItem) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponseAddItem element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponseAddItem element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "itemSequence":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for i := 0; d.More(); i++ {
				var v PositiveInt
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.ItemSequence) <= i {
					r.ItemSequence = append(r.ItemSequence, PositiveInt{})
				}
				r.ItemSequence[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "_itemSequence":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.ItemSequence) <= i {
					r.ItemSequence = append(r.ItemSequence, PositiveInt{})
				}
				r.ItemSequence[i].Id = v.Id
				r.ItemSequence[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "detailSequence":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for i := 0; d.More(); i++ {
				var v PositiveInt
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.DetailSequence) <= i {
					r.DetailSequence = append(r.DetailSequence, PositiveInt{})
				}
				r.DetailSequence[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "_detailSequence":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.DetailSequence) <= i {
					r.DetailSequence = append(r.DetailSequence, PositiveInt{})
				}
				r.DetailSequence[i].Id = v.Id
				r.DetailSequence[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "subdetailSequence":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for i := 0; d.More(); i++ {
				var v PositiveInt
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.SubdetailSequence) <= i {
					r.SubdetailSequence = append(r.SubdetailSequence, PositiveInt{})
				}
				r.SubdetailSequence[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "_subdetailSequence":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.SubdetailSequence) <= i {
					r.SubdetailSequence = append(r.SubdetailSequence, PositiveInt{})
				}
				r.SubdetailSequence[i].Id = v.Id
				r.SubdetailSequence[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "traceNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.TraceNumber = append(r.TraceNumber, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "provider":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Provider = append(r.Provider, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "revenue":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Revenue = &v
		case "productOrService":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ProductOrService = &v
		case "productOrServiceEnd":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ProductOrServiceEnd = &v
		case "request":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Request = append(r.Request, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "modifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "programCode":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ProgramCode = append(r.ProgramCode, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
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
		case "locationCodeableConcept":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Location = v
		case "locationAddress":
			var v Address
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Location = v
		case "locationReference":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Location = v
		case "quantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Quantity = &v
		case "unitPrice":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.UnitPrice = &v
		case "factor":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Factor == nil {
				r.Factor = &Decimal{}
			}
			r.Factor.Value = v.Value
		case "_factor":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Factor == nil {
				r.Factor = &Decimal{}
			}
			r.Factor.Id = v.Id
			r.Factor.Extension = v.Extension
		case "tax":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Tax = &v
		case "net":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Net = &v
		case "bodySite":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for d.More() {
				var v ClaimResponseAddItemBodySite
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.BodySite = append(r.BodySite, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "noteNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for i := 0; d.More(); i++ {
				var v PositiveInt
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.NoteNumber) <= i {
					r.NoteNumber = append(r.NoteNumber, PositiveInt{})
				}
				r.NoteNumber[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "_noteNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.NoteNumber) <= i {
					r.NoteNumber = append(r.NoteNumber, PositiveInt{})
				}
				r.NoteNumber[i].Id = v.Id
				r.NoteNumber[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "reviewOutcome":
			var v ClaimResponseItemReviewOutcome
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ReviewOutcome = &v
		case "adjudication":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for d.More() {
				var v ClaimResponseItemAdjudication
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		case "detail":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItem element", t)
			}
			for d.More() {
				var v ClaimResponseAddItemDetail
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Detail = append(r.Detail, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItem element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponseAddItem", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponseAddItem element", t)
	}
	return nil
}
func (r *ClaimResponseAddItemBodySite) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponseAddItemBodySite element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponseAddItemBodySite element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemBodySite element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemBodySite element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemBodySite element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemBodySite element", t)
			}
		case "site":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemBodySite element", t)
			}
			for d.More() {
				var v CodeableReference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Site = append(r.Site, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemBodySite element", t)
			}
		case "subSite":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemBodySite element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.SubSite = append(r.SubSite, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemBodySite element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponseAddItemBodySite", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponseAddItemBodySite element", t)
	}
	return nil
}
func (r *ClaimResponseAddItemDetail) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponseAddItemDetail element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponseAddItemDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetail element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetail element", t)
			}
		case "traceNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetail element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.TraceNumber = append(r.TraceNumber, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetail element", t)
			}
		case "revenue":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Revenue = &v
		case "productOrService":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ProductOrService = &v
		case "productOrServiceEnd":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ProductOrServiceEnd = &v
		case "modifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetail element", t)
			}
		case "quantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Quantity = &v
		case "unitPrice":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.UnitPrice = &v
		case "factor":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Factor == nil {
				r.Factor = &Decimal{}
			}
			r.Factor.Value = v.Value
		case "_factor":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Factor == nil {
				r.Factor = &Decimal{}
			}
			r.Factor.Id = v.Id
			r.Factor.Extension = v.Extension
		case "tax":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Tax = &v
		case "net":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Net = &v
		case "noteNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetail element", t)
			}
			for i := 0; d.More(); i++ {
				var v PositiveInt
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.NoteNumber) <= i {
					r.NoteNumber = append(r.NoteNumber, PositiveInt{})
				}
				r.NoteNumber[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetail element", t)
			}
		case "_noteNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetail element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.NoteNumber) <= i {
					r.NoteNumber = append(r.NoteNumber, PositiveInt{})
				}
				r.NoteNumber[i].Id = v.Id
				r.NoteNumber[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetail element", t)
			}
		case "reviewOutcome":
			var v ClaimResponseItemReviewOutcome
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ReviewOutcome = &v
		case "adjudication":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetail element", t)
			}
			for d.More() {
				var v ClaimResponseItemAdjudication
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetail element", t)
			}
		case "subDetail":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetail element", t)
			}
			for d.More() {
				var v ClaimResponseAddItemDetailSubDetail
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.SubDetail = append(r.SubDetail, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetail element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponseAddItemDetail", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponseAddItemDetail element", t)
	}
	return nil
}
func (r *ClaimResponseAddItemDetailSubDetail) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponseAddItemDetailSubDetail element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponseAddItemDetailSubDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetailSubDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetailSubDetail element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetailSubDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetailSubDetail element", t)
			}
		case "traceNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetailSubDetail element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.TraceNumber = append(r.TraceNumber, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetailSubDetail element", t)
			}
		case "revenue":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Revenue = &v
		case "productOrService":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ProductOrService = &v
		case "productOrServiceEnd":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ProductOrServiceEnd = &v
		case "modifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetailSubDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetailSubDetail element", t)
			}
		case "quantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Quantity = &v
		case "unitPrice":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.UnitPrice = &v
		case "factor":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Factor == nil {
				r.Factor = &Decimal{}
			}
			r.Factor.Value = v.Value
		case "_factor":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Factor == nil {
				r.Factor = &Decimal{}
			}
			r.Factor.Id = v.Id
			r.Factor.Extension = v.Extension
		case "tax":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Tax = &v
		case "net":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Net = &v
		case "noteNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetailSubDetail element", t)
			}
			for i := 0; d.More(); i++ {
				var v PositiveInt
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.NoteNumber) <= i {
					r.NoteNumber = append(r.NoteNumber, PositiveInt{})
				}
				r.NoteNumber[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetailSubDetail element", t)
			}
		case "_noteNumber":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetailSubDetail element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.NoteNumber) <= i {
					r.NoteNumber = append(r.NoteNumber, PositiveInt{})
				}
				r.NoteNumber[i].Id = v.Id
				r.NoteNumber[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetailSubDetail element", t)
			}
		case "reviewOutcome":
			var v ClaimResponseItemReviewOutcome
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ReviewOutcome = &v
		case "adjudication":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseAddItemDetailSubDetail element", t)
			}
			for d.More() {
				var v ClaimResponseItemAdjudication
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseAddItemDetailSubDetail element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponseAddItemDetailSubDetail", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponseAddItemDetailSubDetail element", t)
	}
	return nil
}
func (r *ClaimResponseTotal) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponseTotal element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponseTotal element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseTotal element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseTotal element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseTotal element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseTotal element", t)
			}
		case "category":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Category = v
		case "amount":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Amount = v
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponseTotal", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponseTotal element", t)
	}
	return nil
}
func (r *ClaimResponsePayment) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponsePayment element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponsePayment element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponsePayment element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponsePayment element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponsePayment element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponsePayment element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = v
		case "adjustment":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Adjustment = &v
		case "adjustmentReason":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.AdjustmentReason = &v
		case "date":
			var v Date
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Date == nil {
				r.Date = &Date{}
			}
			r.Date.Value = v.Value
		case "_date":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Date == nil {
				r.Date = &Date{}
			}
			r.Date.Id = v.Id
			r.Date.Extension = v.Extension
		case "amount":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Amount = v
		case "identifier":
			var v Identifier
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Identifier = &v
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponsePayment", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponsePayment element", t)
	}
	return nil
}
func (r *ClaimResponseProcessNote) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponseProcessNote element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponseProcessNote element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseProcessNote element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseProcessNote element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseProcessNote element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseProcessNote element", t)
			}
		case "number":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Number == nil {
				r.Number = &PositiveInt{}
			}
			r.Number.Value = v.Value
		case "_number":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Number == nil {
				r.Number = &PositiveInt{}
			}
			r.Number.Id = v.Id
			r.Number.Extension = v.Extension
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = &v
		case "text":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Text.Value = v.Value
		case "_text":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Text.Id = v.Id
			r.Text.Extension = v.Extension
		case "language":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Language = &v
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponseProcessNote", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponseProcessNote element", t)
	}
	return nil
}
func (r *ClaimResponseInsurance) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponseInsurance element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponseInsurance element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseInsurance element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseInsurance element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseInsurance element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseInsurance element", t)
			}
		case "sequence":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Sequence.Value = v.Value
		case "_sequence":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Sequence.Id = v.Id
			r.Sequence.Extension = v.Extension
		case "focal":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Focal.Value = v.Value
		case "_focal":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Focal.Id = v.Id
			r.Focal.Extension = v.Extension
		case "coverage":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Coverage = v
		case "businessArrangement":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.BusinessArrangement == nil {
				r.BusinessArrangement = &String{}
			}
			r.BusinessArrangement.Value = v.Value
		case "_businessArrangement":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.BusinessArrangement == nil {
				r.BusinessArrangement = &String{}
			}
			r.BusinessArrangement.Id = v.Id
			r.BusinessArrangement.Extension = v.Extension
		case "claimResponse":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ClaimResponse = &v
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponseInsurance", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponseInsurance element", t)
	}
	return nil
}
func (r *ClaimResponseError) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ClaimResponseError element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ClaimResponseError element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseError element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseError element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseError element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseError element", t)
			}
		case "itemSequence":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ItemSequence == nil {
				r.ItemSequence = &PositiveInt{}
			}
			r.ItemSequence.Value = v.Value
		case "_itemSequence":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ItemSequence == nil {
				r.ItemSequence = &PositiveInt{}
			}
			r.ItemSequence.Id = v.Id
			r.ItemSequence.Extension = v.Extension
		case "detailSequence":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.DetailSequence == nil {
				r.DetailSequence = &PositiveInt{}
			}
			r.DetailSequence.Value = v.Value
		case "_detailSequence":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.DetailSequence == nil {
				r.DetailSequence = &PositiveInt{}
			}
			r.DetailSequence.Id = v.Id
			r.DetailSequence.Extension = v.Extension
		case "subDetailSequence":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.SubDetailSequence == nil {
				r.SubDetailSequence = &PositiveInt{}
			}
			r.SubDetailSequence.Value = v.Value
		case "_subDetailSequence":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.SubDetailSequence == nil {
				r.SubDetailSequence = &PositiveInt{}
			}
			r.SubDetailSequence.Id = v.Id
			r.SubDetailSequence.Extension = v.Extension
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseError element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseError element", t)
			}
		case "_expression":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ClaimResponseError element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ClaimResponseError element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ClaimResponseError", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ClaimResponseError element", t)
	}
	return nil
}
func (r ClaimResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "ClaimResponse"
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
	err = e.EncodeElement(r.TraceNumber, xml.StartElement{Name: xml.Name{Local: "traceNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubType, xml.StartElement{Name: xml.Name{Local: "subType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Use, xml.StartElement{Name: xml.Name{Local: "use"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Patient, xml.StartElement{Name: xml.Name{Local: "patient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Created, xml.StartElement{Name: xml.Name{Local: "created"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Insurer, xml.StartElement{Name: xml.Name{Local: "insurer"}})
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
	err = e.EncodeElement(r.Decision, xml.StartElement{Name: xml.Name{Local: "decision"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Disposition, xml.StartElement{Name: xml.Name{Local: "disposition"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PreAuthRef, xml.StartElement{Name: xml.Name{Local: "preAuthRef"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PreAuthPeriod, xml.StartElement{Name: xml.Name{Local: "preAuthPeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Event, xml.StartElement{Name: xml.Name{Local: "event"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PayeeType, xml.StartElement{Name: xml.Name{Local: "payeeType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Encounter, xml.StartElement{Name: xml.Name{Local: "encounter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DiagnosisRelatedGroup, xml.StartElement{Name: xml.Name{Local: "diagnosisRelatedGroup"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Item, xml.StartElement{Name: xml.Name{Local: "item"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AddItem, xml.StartElement{Name: xml.Name{Local: "addItem"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Adjudication, xml.StartElement{Name: xml.Name{Local: "adjudication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Total, xml.StartElement{Name: xml.Name{Local: "total"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Payment, xml.StartElement{Name: xml.Name{Local: "payment"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FundsReserve, xml.StartElement{Name: xml.Name{Local: "fundsReserve"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FormCode, xml.StartElement{Name: xml.Name{Local: "formCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Form, xml.StartElement{Name: xml.Name{Local: "form"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProcessNote, xml.StartElement{Name: xml.Name{Local: "processNote"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CommunicationRequest, xml.StartElement{Name: xml.Name{Local: "communicationRequest"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Insurance, xml.StartElement{Name: xml.Name{Local: "insurance"}})
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
func (r ClaimResponseEvent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
func (r ClaimResponseItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ItemSequence, xml.StartElement{Name: xml.Name{Local: "itemSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TraceNumber, xml.StartElement{Name: xml.Name{Local: "traceNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NoteNumber, xml.StartElement{Name: xml.Name{Local: "noteNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReviewOutcome, xml.StartElement{Name: xml.Name{Local: "reviewOutcome"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Adjudication, xml.StartElement{Name: xml.Name{Local: "adjudication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Detail, xml.StartElement{Name: xml.Name{Local: "detail"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ClaimResponseItemReviewOutcome) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Decision, xml.StartElement{Name: xml.Name{Local: "decision"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Reason, xml.StartElement{Name: xml.Name{Local: "reason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PreAuthRef, xml.StartElement{Name: xml.Name{Local: "preAuthRef"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PreAuthPeriod, xml.StartElement{Name: xml.Name{Local: "preAuthPeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ClaimResponseItemAdjudication) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Reason, xml.StartElement{Name: xml.Name{Local: "reason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Amount, xml.StartElement{Name: xml.Name{Local: "amount"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ClaimResponseItemDetail) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.DetailSequence, xml.StartElement{Name: xml.Name{Local: "detailSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TraceNumber, xml.StartElement{Name: xml.Name{Local: "traceNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NoteNumber, xml.StartElement{Name: xml.Name{Local: "noteNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReviewOutcome, xml.StartElement{Name: xml.Name{Local: "reviewOutcome"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Adjudication, xml.StartElement{Name: xml.Name{Local: "adjudication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubDetail, xml.StartElement{Name: xml.Name{Local: "subDetail"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ClaimResponseItemDetailSubDetail) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.SubDetailSequence, xml.StartElement{Name: xml.Name{Local: "subDetailSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TraceNumber, xml.StartElement{Name: xml.Name{Local: "traceNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NoteNumber, xml.StartElement{Name: xml.Name{Local: "noteNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReviewOutcome, xml.StartElement{Name: xml.Name{Local: "reviewOutcome"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Adjudication, xml.StartElement{Name: xml.Name{Local: "adjudication"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ClaimResponseAddItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ItemSequence, xml.StartElement{Name: xml.Name{Local: "itemSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DetailSequence, xml.StartElement{Name: xml.Name{Local: "detailSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubdetailSequence, xml.StartElement{Name: xml.Name{Local: "subdetailSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TraceNumber, xml.StartElement{Name: xml.Name{Local: "traceNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Provider, xml.StartElement{Name: xml.Name{Local: "provider"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Revenue, xml.StartElement{Name: xml.Name{Local: "revenue"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProductOrService, xml.StartElement{Name: xml.Name{Local: "productOrService"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProductOrServiceEnd, xml.StartElement{Name: xml.Name{Local: "productOrServiceEnd"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Request, xml.StartElement{Name: xml.Name{Local: "request"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Modifier, xml.StartElement{Name: xml.Name{Local: "modifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProgramCode, xml.StartElement{Name: xml.Name{Local: "programCode"}})
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
	switch v := r.Location.(type) {
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "locationCodeableConcept"}})
		if err != nil {
			return err
		}
	case Address, *Address:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "locationAddress"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "locationReference"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UnitPrice, xml.StartElement{Name: xml.Name{Local: "unitPrice"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Factor, xml.StartElement{Name: xml.Name{Local: "factor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Tax, xml.StartElement{Name: xml.Name{Local: "tax"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Net, xml.StartElement{Name: xml.Name{Local: "net"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BodySite, xml.StartElement{Name: xml.Name{Local: "bodySite"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NoteNumber, xml.StartElement{Name: xml.Name{Local: "noteNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReviewOutcome, xml.StartElement{Name: xml.Name{Local: "reviewOutcome"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Adjudication, xml.StartElement{Name: xml.Name{Local: "adjudication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Detail, xml.StartElement{Name: xml.Name{Local: "detail"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ClaimResponseAddItemBodySite) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Site, xml.StartElement{Name: xml.Name{Local: "site"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubSite, xml.StartElement{Name: xml.Name{Local: "subSite"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ClaimResponseAddItemDetail) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.TraceNumber, xml.StartElement{Name: xml.Name{Local: "traceNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Revenue, xml.StartElement{Name: xml.Name{Local: "revenue"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProductOrService, xml.StartElement{Name: xml.Name{Local: "productOrService"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProductOrServiceEnd, xml.StartElement{Name: xml.Name{Local: "productOrServiceEnd"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Modifier, xml.StartElement{Name: xml.Name{Local: "modifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UnitPrice, xml.StartElement{Name: xml.Name{Local: "unitPrice"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Factor, xml.StartElement{Name: xml.Name{Local: "factor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Tax, xml.StartElement{Name: xml.Name{Local: "tax"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Net, xml.StartElement{Name: xml.Name{Local: "net"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NoteNumber, xml.StartElement{Name: xml.Name{Local: "noteNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReviewOutcome, xml.StartElement{Name: xml.Name{Local: "reviewOutcome"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Adjudication, xml.StartElement{Name: xml.Name{Local: "adjudication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubDetail, xml.StartElement{Name: xml.Name{Local: "subDetail"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ClaimResponseAddItemDetailSubDetail) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.TraceNumber, xml.StartElement{Name: xml.Name{Local: "traceNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Revenue, xml.StartElement{Name: xml.Name{Local: "revenue"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProductOrService, xml.StartElement{Name: xml.Name{Local: "productOrService"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProductOrServiceEnd, xml.StartElement{Name: xml.Name{Local: "productOrServiceEnd"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Modifier, xml.StartElement{Name: xml.Name{Local: "modifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UnitPrice, xml.StartElement{Name: xml.Name{Local: "unitPrice"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Factor, xml.StartElement{Name: xml.Name{Local: "factor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Tax, xml.StartElement{Name: xml.Name{Local: "tax"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Net, xml.StartElement{Name: xml.Name{Local: "net"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NoteNumber, xml.StartElement{Name: xml.Name{Local: "noteNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReviewOutcome, xml.StartElement{Name: xml.Name{Local: "reviewOutcome"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Adjudication, xml.StartElement{Name: xml.Name{Local: "adjudication"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ClaimResponseTotal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Amount, xml.StartElement{Name: xml.Name{Local: "amount"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ClaimResponsePayment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Adjustment, xml.StartElement{Name: xml.Name{Local: "adjustment"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AdjustmentReason, xml.StartElement{Name: xml.Name{Local: "adjustmentReason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Amount, xml.StartElement{Name: xml.Name{Local: "amount"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ClaimResponseProcessNote) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Number, xml.StartElement{Name: xml.Name{Local: "number"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ClaimResponseInsurance) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Sequence, xml.StartElement{Name: xml.Name{Local: "sequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Focal, xml.StartElement{Name: xml.Name{Local: "focal"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Coverage, xml.StartElement{Name: xml.Name{Local: "coverage"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BusinessArrangement, xml.StartElement{Name: xml.Name{Local: "businessArrangement"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ClaimResponse, xml.StartElement{Name: xml.Name{Local: "claimResponse"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ClaimResponseError) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ItemSequence, xml.StartElement{Name: xml.Name{Local: "itemSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DetailSequence, xml.StartElement{Name: xml.Name{Local: "detailSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubDetailSequence, xml.StartElement{Name: xml.Name{Local: "subDetailSequence"}})
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
func (r *ClaimResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "traceNumber":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TraceNumber = append(r.TraceNumber, v)
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "subType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubType = &v
			case "use":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Use = v
			case "patient":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Patient = v
			case "created":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Created = v
			case "insurer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Insurer = &v
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
				r.Request = &v
			case "outcome":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Outcome = v
			case "decision":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Decision = &v
			case "disposition":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Disposition = &v
			case "preAuthRef":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PreAuthRef = &v
			case "preAuthPeriod":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PreAuthPeriod = &v
			case "event":
				var v ClaimResponseEvent
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Event = append(r.Event, v)
			case "payeeType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PayeeType = &v
			case "encounter":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Encounter = append(r.Encounter, v)
			case "diagnosisRelatedGroup":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DiagnosisRelatedGroup = &v
			case "item":
				var v ClaimResponseItem
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Item = append(r.Item, v)
			case "addItem":
				var v ClaimResponseAddItem
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AddItem = append(r.AddItem, v)
			case "adjudication":
				var v ClaimResponseItemAdjudication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			case "total":
				var v ClaimResponseTotal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Total = append(r.Total, v)
			case "payment":
				var v ClaimResponsePayment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Payment = &v
			case "fundsReserve":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FundsReserve = &v
			case "formCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FormCode = &v
			case "form":
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Form = &v
			case "processNote":
				var v ClaimResponseProcessNote
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProcessNote = append(r.ProcessNote, v)
			case "communicationRequest":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CommunicationRequest = append(r.CommunicationRequest, v)
			case "insurance":
				var v ClaimResponseInsurance
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Insurance = append(r.Insurance, v)
			case "error":
				var v ClaimResponseError
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
func (r *ClaimResponseEvent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
func (r *ClaimResponseItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "itemSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ItemSequence = v
			case "traceNumber":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TraceNumber = append(r.TraceNumber, v)
			case "noteNumber":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NoteNumber = append(r.NoteNumber, v)
			case "reviewOutcome":
				var v ClaimResponseItemReviewOutcome
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReviewOutcome = &v
			case "adjudication":
				var v ClaimResponseItemAdjudication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			case "detail":
				var v ClaimResponseItemDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = append(r.Detail, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ClaimResponseItemReviewOutcome) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "decision":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Decision = &v
			case "reason":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reason = append(r.Reason, v)
			case "preAuthRef":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PreAuthRef = &v
			case "preAuthPeriod":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PreAuthPeriod = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ClaimResponseItemAdjudication) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.Category = v
			case "reason":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reason = &v
			case "amount":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = &v
			case "quantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ClaimResponseItemDetail) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "detailSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DetailSequence = v
			case "traceNumber":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TraceNumber = append(r.TraceNumber, v)
			case "noteNumber":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NoteNumber = append(r.NoteNumber, v)
			case "reviewOutcome":
				var v ClaimResponseItemReviewOutcome
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReviewOutcome = &v
			case "adjudication":
				var v ClaimResponseItemAdjudication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			case "subDetail":
				var v ClaimResponseItemDetailSubDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubDetail = append(r.SubDetail, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ClaimResponseItemDetailSubDetail) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "subDetailSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubDetailSequence = v
			case "traceNumber":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TraceNumber = append(r.TraceNumber, v)
			case "noteNumber":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NoteNumber = append(r.NoteNumber, v)
			case "reviewOutcome":
				var v ClaimResponseItemReviewOutcome
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReviewOutcome = &v
			case "adjudication":
				var v ClaimResponseItemAdjudication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ClaimResponseAddItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "itemSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ItemSequence = append(r.ItemSequence, v)
			case "detailSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DetailSequence = append(r.DetailSequence, v)
			case "subdetailSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubdetailSequence = append(r.SubdetailSequence, v)
			case "traceNumber":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TraceNumber = append(r.TraceNumber, v)
			case "provider":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Provider = append(r.Provider, v)
			case "revenue":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Revenue = &v
			case "productOrService":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductOrService = &v
			case "productOrServiceEnd":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductOrServiceEnd = &v
			case "request":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Request = append(r.Request, v)
			case "modifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Modifier = append(r.Modifier, v)
			case "programCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProgramCode = append(r.ProgramCode, v)
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
			case "locationCodeableConcept":
				if r.Location != nil {
					return fmt.Errorf("multiple values for field \"Location\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = v
			case "locationAddress":
				if r.Location != nil {
					return fmt.Errorf("multiple values for field \"Location\"")
				}
				var v Address
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = v
			case "locationReference":
				if r.Location != nil {
					return fmt.Errorf("multiple values for field \"Location\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = v
			case "quantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = &v
			case "unitPrice":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UnitPrice = &v
			case "factor":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Factor = &v
			case "tax":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Tax = &v
			case "net":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Net = &v
			case "bodySite":
				var v ClaimResponseAddItemBodySite
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BodySite = append(r.BodySite, v)
			case "noteNumber":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NoteNumber = append(r.NoteNumber, v)
			case "reviewOutcome":
				var v ClaimResponseItemReviewOutcome
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReviewOutcome = &v
			case "adjudication":
				var v ClaimResponseItemAdjudication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			case "detail":
				var v ClaimResponseAddItemDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = append(r.Detail, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ClaimResponseAddItemBodySite) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "site":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Site = append(r.Site, v)
			case "subSite":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubSite = append(r.SubSite, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ClaimResponseAddItemDetail) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "traceNumber":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TraceNumber = append(r.TraceNumber, v)
			case "revenue":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Revenue = &v
			case "productOrService":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductOrService = &v
			case "productOrServiceEnd":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductOrServiceEnd = &v
			case "modifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Modifier = append(r.Modifier, v)
			case "quantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = &v
			case "unitPrice":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UnitPrice = &v
			case "factor":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Factor = &v
			case "tax":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Tax = &v
			case "net":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Net = &v
			case "noteNumber":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NoteNumber = append(r.NoteNumber, v)
			case "reviewOutcome":
				var v ClaimResponseItemReviewOutcome
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReviewOutcome = &v
			case "adjudication":
				var v ClaimResponseItemAdjudication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			case "subDetail":
				var v ClaimResponseAddItemDetailSubDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubDetail = append(r.SubDetail, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ClaimResponseAddItemDetailSubDetail) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "traceNumber":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TraceNumber = append(r.TraceNumber, v)
			case "revenue":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Revenue = &v
			case "productOrService":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductOrService = &v
			case "productOrServiceEnd":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductOrServiceEnd = &v
			case "modifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Modifier = append(r.Modifier, v)
			case "quantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = &v
			case "unitPrice":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UnitPrice = &v
			case "factor":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Factor = &v
			case "tax":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Tax = &v
			case "net":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Net = &v
			case "noteNumber":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NoteNumber = append(r.NoteNumber, v)
			case "reviewOutcome":
				var v ClaimResponseItemReviewOutcome
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReviewOutcome = &v
			case "adjudication":
				var v ClaimResponseItemAdjudication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ClaimResponseTotal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.Category = v
			case "amount":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ClaimResponsePayment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "adjustment":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjustment = &v
			case "adjustmentReason":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AdjustmentReason = &v
			case "date":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = &v
			case "amount":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ClaimResponseProcessNote) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "number":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Number = &v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "text":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = v
			case "language":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ClaimResponseInsurance) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "sequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Sequence = v
			case "focal":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Focal = v
			case "coverage":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Coverage = v
			case "businessArrangement":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BusinessArrangement = &v
			case "claimResponse":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ClaimResponse = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ClaimResponseError) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "itemSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ItemSequence = &v
			case "detailSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DetailSequence = &v
			case "subDetailSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubDetailSequence = &v
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
func (r ClaimResponse) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "traceNumber") {
		for _, v := range r.TraceNumber {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "status") {
		children = append(children, r.Status)
	}
	if len(name) == 0 || slices.Contains(name, "type") {
		children = append(children, r.Type)
	}
	if len(name) == 0 || slices.Contains(name, "subType") {
		if r.SubType != nil {
			children = append(children, *r.SubType)
		}
	}
	if len(name) == 0 || slices.Contains(name, "use") {
		children = append(children, r.Use)
	}
	if len(name) == 0 || slices.Contains(name, "patient") {
		children = append(children, r.Patient)
	}
	if len(name) == 0 || slices.Contains(name, "created") {
		children = append(children, r.Created)
	}
	if len(name) == 0 || slices.Contains(name, "insurer") {
		if r.Insurer != nil {
			children = append(children, *r.Insurer)
		}
	}
	if len(name) == 0 || slices.Contains(name, "requestor") {
		if r.Requestor != nil {
			children = append(children, *r.Requestor)
		}
	}
	if len(name) == 0 || slices.Contains(name, "request") {
		if r.Request != nil {
			children = append(children, *r.Request)
		}
	}
	if len(name) == 0 || slices.Contains(name, "outcome") {
		children = append(children, r.Outcome)
	}
	if len(name) == 0 || slices.Contains(name, "decision") {
		if r.Decision != nil {
			children = append(children, *r.Decision)
		}
	}
	if len(name) == 0 || slices.Contains(name, "disposition") {
		if r.Disposition != nil {
			children = append(children, *r.Disposition)
		}
	}
	if len(name) == 0 || slices.Contains(name, "preAuthRef") {
		if r.PreAuthRef != nil {
			children = append(children, *r.PreAuthRef)
		}
	}
	if len(name) == 0 || slices.Contains(name, "preAuthPeriod") {
		if r.PreAuthPeriod != nil {
			children = append(children, *r.PreAuthPeriod)
		}
	}
	if len(name) == 0 || slices.Contains(name, "event") {
		for _, v := range r.Event {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "payeeType") {
		if r.PayeeType != nil {
			children = append(children, *r.PayeeType)
		}
	}
	if len(name) == 0 || slices.Contains(name, "encounter") {
		for _, v := range r.Encounter {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "diagnosisRelatedGroup") {
		if r.DiagnosisRelatedGroup != nil {
			children = append(children, *r.DiagnosisRelatedGroup)
		}
	}
	if len(name) == 0 || slices.Contains(name, "item") {
		for _, v := range r.Item {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "addItem") {
		for _, v := range r.AddItem {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "adjudication") {
		for _, v := range r.Adjudication {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "total") {
		for _, v := range r.Total {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "payment") {
		if r.Payment != nil {
			children = append(children, *r.Payment)
		}
	}
	if len(name) == 0 || slices.Contains(name, "fundsReserve") {
		if r.FundsReserve != nil {
			children = append(children, *r.FundsReserve)
		}
	}
	if len(name) == 0 || slices.Contains(name, "formCode") {
		if r.FormCode != nil {
			children = append(children, *r.FormCode)
		}
	}
	if len(name) == 0 || slices.Contains(name, "form") {
		if r.Form != nil {
			children = append(children, *r.Form)
		}
	}
	if len(name) == 0 || slices.Contains(name, "processNote") {
		for _, v := range r.ProcessNote {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "communicationRequest") {
		for _, v := range r.CommunicationRequest {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "insurance") {
		for _, v := range r.Insurance {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "error") {
		for _, v := range r.Error {
			children = append(children, v)
		}
	}
	return children
}
func (r ClaimResponse) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponse to Boolean")
}
func (r ClaimResponse) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponse to String")
}
func (r ClaimResponse) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponse to Integer")
}
func (r ClaimResponse) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponse to Decimal")
}
func (r ClaimResponse) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponse to Date")
}
func (r ClaimResponse) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponse to Time")
}
func (r ClaimResponse) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponse to DateTime")
}
func (r ClaimResponse) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponse to Quantity")
}
func (r ClaimResponse) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponse
	switch other := other.(type) {
	case ClaimResponse:
		o = &other
	case *ClaimResponse:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponse) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponse) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Id",
				Namespace: "FHIR",
			},
		}, {
			Name: "Meta",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Meta",
				Namespace: "FHIR",
			},
		}, {
			Name: "ImplicitRules",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Uri",
				Namespace: "FHIR",
			},
		}, {
			Name: "Language",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Text",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Narrative",
				Namespace: "FHIR",
			},
		}, {
			Name: "Contained",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Identifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "TraceNumber",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "Status",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "SubType",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Use",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Patient",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Created",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "DateTime",
				Namespace: "FHIR",
			},
		}, {
			Name: "Insurer",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Requestor",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Request",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Outcome",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Decision",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Disposition",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "PreAuthRef",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "PreAuthPeriod",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Period",
				Namespace: "FHIR",
			},
		}, {
			Name: "Event",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseEvent",
				Namespace: "FHIR",
			},
		}, {
			Name: "PayeeType",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Encounter",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "DiagnosisRelatedGroup",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Item",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseItem",
				Namespace: "FHIR",
			},
		}, {
			Name: "AddItem",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseAddItem",
				Namespace: "FHIR",
			},
		}, {
			Name: "Adjudication",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseItemAdjudication",
				Namespace: "FHIR",
			},
		}, {
			Name: "Total",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseTotal",
				Namespace: "FHIR",
			},
		}, {
			Name: "Payment",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "ClaimResponsePayment",
				Namespace: "FHIR",
			},
		}, {
			Name: "FundsReserve",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "FormCode",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Form",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Attachment",
				Namespace: "FHIR",
			},
		}, {
			Name: "ProcessNote",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseProcessNote",
				Namespace: "FHIR",
			},
		}, {
			Name: "CommunicationRequest",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Insurance",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseInsurance",
				Namespace: "FHIR",
			},
		}, {
			Name: "Error",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseError",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponse",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponseEvent) Children(name ...string) fhirpath.Collection {
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
func (r ClaimResponseEvent) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponseEvent to Boolean")
}
func (r ClaimResponseEvent) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponseEvent to String")
}
func (r ClaimResponseEvent) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponseEvent to Integer")
}
func (r ClaimResponseEvent) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponseEvent to Decimal")
}
func (r ClaimResponseEvent) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponseEvent to Date")
}
func (r ClaimResponseEvent) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponseEvent to Time")
}
func (r ClaimResponseEvent) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponseEvent to DateTime")
}
func (r ClaimResponseEvent) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponseEvent to Quantity")
}
func (r ClaimResponseEvent) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponseEvent
	switch other := other.(type) {
	case ClaimResponseEvent:
		o = &other
	case *ClaimResponseEvent:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponseEvent) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponseEvent) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "When",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PrimitiveElement",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponseEvent",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponseItem) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "itemSequence") {
		children = append(children, r.ItemSequence)
	}
	if len(name) == 0 || slices.Contains(name, "traceNumber") {
		for _, v := range r.TraceNumber {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "noteNumber") {
		for _, v := range r.NoteNumber {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "reviewOutcome") {
		if r.ReviewOutcome != nil {
			children = append(children, *r.ReviewOutcome)
		}
	}
	if len(name) == 0 || slices.Contains(name, "adjudication") {
		for _, v := range r.Adjudication {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "detail") {
		for _, v := range r.Detail {
			children = append(children, v)
		}
	}
	return children
}
func (r ClaimResponseItem) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponseItem to Boolean")
}
func (r ClaimResponseItem) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponseItem to String")
}
func (r ClaimResponseItem) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponseItem to Integer")
}
func (r ClaimResponseItem) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponseItem to Decimal")
}
func (r ClaimResponseItem) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponseItem to Date")
}
func (r ClaimResponseItem) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponseItem to Time")
}
func (r ClaimResponseItem) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponseItem to DateTime")
}
func (r ClaimResponseItem) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponseItem to Quantity")
}
func (r ClaimResponseItem) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponseItem
	switch other := other.(type) {
	case ClaimResponseItem:
		o = &other
	case *ClaimResponseItem:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponseItem) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponseItem) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ItemSequence",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "TraceNumber",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "NoteNumber",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "ReviewOutcome",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "ClaimResponseItemReviewOutcome",
				Namespace: "FHIR",
			},
		}, {
			Name: "Adjudication",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseItemAdjudication",
				Namespace: "FHIR",
			},
		}, {
			Name: "Detail",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseItemDetail",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponseItem",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponseItemReviewOutcome) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "decision") {
		if r.Decision != nil {
			children = append(children, *r.Decision)
		}
	}
	if len(name) == 0 || slices.Contains(name, "reason") {
		for _, v := range r.Reason {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "preAuthRef") {
		if r.PreAuthRef != nil {
			children = append(children, *r.PreAuthRef)
		}
	}
	if len(name) == 0 || slices.Contains(name, "preAuthPeriod") {
		if r.PreAuthPeriod != nil {
			children = append(children, *r.PreAuthPeriod)
		}
	}
	return children
}
func (r ClaimResponseItemReviewOutcome) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponseItemReviewOutcome to Boolean")
}
func (r ClaimResponseItemReviewOutcome) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponseItemReviewOutcome to String")
}
func (r ClaimResponseItemReviewOutcome) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponseItemReviewOutcome to Integer")
}
func (r ClaimResponseItemReviewOutcome) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponseItemReviewOutcome to Decimal")
}
func (r ClaimResponseItemReviewOutcome) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponseItemReviewOutcome to Date")
}
func (r ClaimResponseItemReviewOutcome) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponseItemReviewOutcome to Time")
}
func (r ClaimResponseItemReviewOutcome) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponseItemReviewOutcome to DateTime")
}
func (r ClaimResponseItemReviewOutcome) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponseItemReviewOutcome to Quantity")
}
func (r ClaimResponseItemReviewOutcome) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponseItemReviewOutcome
	switch other := other.(type) {
	case ClaimResponseItemReviewOutcome:
		o = &other
	case *ClaimResponseItemReviewOutcome:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponseItemReviewOutcome) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponseItemReviewOutcome) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Decision",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Reason",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "PreAuthRef",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "PreAuthPeriod",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Period",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponseItemReviewOutcome",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponseItemAdjudication) Children(name ...string) fhirpath.Collection {
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
		children = append(children, r.Category)
	}
	if len(name) == 0 || slices.Contains(name, "reason") {
		if r.Reason != nil {
			children = append(children, *r.Reason)
		}
	}
	if len(name) == 0 || slices.Contains(name, "amount") {
		if r.Amount != nil {
			children = append(children, *r.Amount)
		}
	}
	if len(name) == 0 || slices.Contains(name, "quantity") {
		if r.Quantity != nil {
			children = append(children, *r.Quantity)
		}
	}
	return children
}
func (r ClaimResponseItemAdjudication) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponseItemAdjudication to Boolean")
}
func (r ClaimResponseItemAdjudication) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponseItemAdjudication to String")
}
func (r ClaimResponseItemAdjudication) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponseItemAdjudication to Integer")
}
func (r ClaimResponseItemAdjudication) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponseItemAdjudication to Decimal")
}
func (r ClaimResponseItemAdjudication) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponseItemAdjudication to Date")
}
func (r ClaimResponseItemAdjudication) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponseItemAdjudication to Time")
}
func (r ClaimResponseItemAdjudication) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponseItemAdjudication to DateTime")
}
func (r ClaimResponseItemAdjudication) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponseItemAdjudication to Quantity")
}
func (r ClaimResponseItemAdjudication) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponseItemAdjudication
	switch other := other.(type) {
	case ClaimResponseItemAdjudication:
		o = &other
	case *ClaimResponseItemAdjudication:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponseItemAdjudication) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponseItemAdjudication) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Category",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Reason",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Amount",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Money",
				Namespace: "FHIR",
			},
		}, {
			Name: "Quantity",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Quantity",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponseItemAdjudication",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponseItemDetail) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "detailSequence") {
		children = append(children, r.DetailSequence)
	}
	if len(name) == 0 || slices.Contains(name, "traceNumber") {
		for _, v := range r.TraceNumber {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "noteNumber") {
		for _, v := range r.NoteNumber {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "reviewOutcome") {
		if r.ReviewOutcome != nil {
			children = append(children, *r.ReviewOutcome)
		}
	}
	if len(name) == 0 || slices.Contains(name, "adjudication") {
		for _, v := range r.Adjudication {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "subDetail") {
		for _, v := range r.SubDetail {
			children = append(children, v)
		}
	}
	return children
}
func (r ClaimResponseItemDetail) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponseItemDetail to Boolean")
}
func (r ClaimResponseItemDetail) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponseItemDetail to String")
}
func (r ClaimResponseItemDetail) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponseItemDetail to Integer")
}
func (r ClaimResponseItemDetail) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponseItemDetail to Decimal")
}
func (r ClaimResponseItemDetail) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponseItemDetail to Date")
}
func (r ClaimResponseItemDetail) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponseItemDetail to Time")
}
func (r ClaimResponseItemDetail) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponseItemDetail to DateTime")
}
func (r ClaimResponseItemDetail) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponseItemDetail to Quantity")
}
func (r ClaimResponseItemDetail) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponseItemDetail
	switch other := other.(type) {
	case ClaimResponseItemDetail:
		o = &other
	case *ClaimResponseItemDetail:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponseItemDetail) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponseItemDetail) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "DetailSequence",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "TraceNumber",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "NoteNumber",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "ReviewOutcome",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "ClaimResponseItemReviewOutcome",
				Namespace: "FHIR",
			},
		}, {
			Name: "Adjudication",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseItemAdjudication",
				Namespace: "FHIR",
			},
		}, {
			Name: "SubDetail",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseItemDetailSubDetail",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponseItemDetail",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponseItemDetailSubDetail) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "subDetailSequence") {
		children = append(children, r.SubDetailSequence)
	}
	if len(name) == 0 || slices.Contains(name, "traceNumber") {
		for _, v := range r.TraceNumber {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "noteNumber") {
		for _, v := range r.NoteNumber {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "reviewOutcome") {
		if r.ReviewOutcome != nil {
			children = append(children, *r.ReviewOutcome)
		}
	}
	if len(name) == 0 || slices.Contains(name, "adjudication") {
		for _, v := range r.Adjudication {
			children = append(children, v)
		}
	}
	return children
}
func (r ClaimResponseItemDetailSubDetail) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponseItemDetailSubDetail to Boolean")
}
func (r ClaimResponseItemDetailSubDetail) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponseItemDetailSubDetail to String")
}
func (r ClaimResponseItemDetailSubDetail) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponseItemDetailSubDetail to Integer")
}
func (r ClaimResponseItemDetailSubDetail) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponseItemDetailSubDetail to Decimal")
}
func (r ClaimResponseItemDetailSubDetail) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponseItemDetailSubDetail to Date")
}
func (r ClaimResponseItemDetailSubDetail) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponseItemDetailSubDetail to Time")
}
func (r ClaimResponseItemDetailSubDetail) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponseItemDetailSubDetail to DateTime")
}
func (r ClaimResponseItemDetailSubDetail) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponseItemDetailSubDetail to Quantity")
}
func (r ClaimResponseItemDetailSubDetail) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponseItemDetailSubDetail
	switch other := other.(type) {
	case ClaimResponseItemDetailSubDetail:
		o = &other
	case *ClaimResponseItemDetailSubDetail:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponseItemDetailSubDetail) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponseItemDetailSubDetail) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "SubDetailSequence",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "TraceNumber",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "NoteNumber",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "ReviewOutcome",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "ClaimResponseItemReviewOutcome",
				Namespace: "FHIR",
			},
		}, {
			Name: "Adjudication",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseItemAdjudication",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponseItemDetailSubDetail",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponseAddItem) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "itemSequence") {
		for _, v := range r.ItemSequence {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "detailSequence") {
		for _, v := range r.DetailSequence {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "subdetailSequence") {
		for _, v := range r.SubdetailSequence {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "traceNumber") {
		for _, v := range r.TraceNumber {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "provider") {
		for _, v := range r.Provider {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "revenue") {
		if r.Revenue != nil {
			children = append(children, *r.Revenue)
		}
	}
	if len(name) == 0 || slices.Contains(name, "productOrService") {
		if r.ProductOrService != nil {
			children = append(children, *r.ProductOrService)
		}
	}
	if len(name) == 0 || slices.Contains(name, "productOrServiceEnd") {
		if r.ProductOrServiceEnd != nil {
			children = append(children, *r.ProductOrServiceEnd)
		}
	}
	if len(name) == 0 || slices.Contains(name, "request") {
		for _, v := range r.Request {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifier") {
		for _, v := range r.Modifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "programCode") {
		for _, v := range r.ProgramCode {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "serviced") {
		if r.Serviced != nil {
			children = append(children, r.Serviced)
		}
	}
	if len(name) == 0 || slices.Contains(name, "location") {
		if r.Location != nil {
			children = append(children, r.Location)
		}
	}
	if len(name) == 0 || slices.Contains(name, "quantity") {
		if r.Quantity != nil {
			children = append(children, *r.Quantity)
		}
	}
	if len(name) == 0 || slices.Contains(name, "unitPrice") {
		if r.UnitPrice != nil {
			children = append(children, *r.UnitPrice)
		}
	}
	if len(name) == 0 || slices.Contains(name, "factor") {
		if r.Factor != nil {
			children = append(children, *r.Factor)
		}
	}
	if len(name) == 0 || slices.Contains(name, "tax") {
		if r.Tax != nil {
			children = append(children, *r.Tax)
		}
	}
	if len(name) == 0 || slices.Contains(name, "net") {
		if r.Net != nil {
			children = append(children, *r.Net)
		}
	}
	if len(name) == 0 || slices.Contains(name, "bodySite") {
		for _, v := range r.BodySite {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "noteNumber") {
		for _, v := range r.NoteNumber {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "reviewOutcome") {
		if r.ReviewOutcome != nil {
			children = append(children, *r.ReviewOutcome)
		}
	}
	if len(name) == 0 || slices.Contains(name, "adjudication") {
		for _, v := range r.Adjudication {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "detail") {
		for _, v := range r.Detail {
			children = append(children, v)
		}
	}
	return children
}
func (r ClaimResponseAddItem) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponseAddItem to Boolean")
}
func (r ClaimResponseAddItem) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponseAddItem to String")
}
func (r ClaimResponseAddItem) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponseAddItem to Integer")
}
func (r ClaimResponseAddItem) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponseAddItem to Decimal")
}
func (r ClaimResponseAddItem) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponseAddItem to Date")
}
func (r ClaimResponseAddItem) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponseAddItem to Time")
}
func (r ClaimResponseAddItem) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponseAddItem to DateTime")
}
func (r ClaimResponseAddItem) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponseAddItem to Quantity")
}
func (r ClaimResponseAddItem) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponseAddItem
	switch other := other.(type) {
	case ClaimResponseAddItem:
		o = &other
	case *ClaimResponseAddItem:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponseAddItem) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponseAddItem) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ItemSequence",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "DetailSequence",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "SubdetailSequence",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "TraceNumber",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "Provider",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Revenue",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ProductOrService",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ProductOrServiceEnd",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Request",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Modifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ProgramCode",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Serviced",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PrimitiveElement",
				Namespace: "FHIR",
			},
		}, {
			Name: "Location",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PrimitiveElement",
				Namespace: "FHIR",
			},
		}, {
			Name: "Quantity",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Quantity",
				Namespace: "FHIR",
			},
		}, {
			Name: "UnitPrice",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Money",
				Namespace: "FHIR",
			},
		}, {
			Name: "Factor",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "Tax",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Money",
				Namespace: "FHIR",
			},
		}, {
			Name: "Net",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Money",
				Namespace: "FHIR",
			},
		}, {
			Name: "BodySite",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseAddItemBodySite",
				Namespace: "FHIR",
			},
		}, {
			Name: "NoteNumber",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "ReviewOutcome",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "ClaimResponseItemReviewOutcome",
				Namespace: "FHIR",
			},
		}, {
			Name: "Adjudication",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseItemAdjudication",
				Namespace: "FHIR",
			},
		}, {
			Name: "Detail",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseAddItemDetail",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponseAddItem",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponseAddItemBodySite) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "site") {
		for _, v := range r.Site {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "subSite") {
		for _, v := range r.SubSite {
			children = append(children, v)
		}
	}
	return children
}
func (r ClaimResponseAddItemBodySite) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponseAddItemBodySite to Boolean")
}
func (r ClaimResponseAddItemBodySite) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponseAddItemBodySite to String")
}
func (r ClaimResponseAddItemBodySite) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponseAddItemBodySite to Integer")
}
func (r ClaimResponseAddItemBodySite) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponseAddItemBodySite to Decimal")
}
func (r ClaimResponseAddItemBodySite) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponseAddItemBodySite to Date")
}
func (r ClaimResponseAddItemBodySite) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponseAddItemBodySite to Time")
}
func (r ClaimResponseAddItemBodySite) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponseAddItemBodySite to DateTime")
}
func (r ClaimResponseAddItemBodySite) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponseAddItemBodySite to Quantity")
}
func (r ClaimResponseAddItemBodySite) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponseAddItemBodySite
	switch other := other.(type) {
	case ClaimResponseAddItemBodySite:
		o = &other
	case *ClaimResponseAddItemBodySite:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponseAddItemBodySite) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponseAddItemBodySite) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Site",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableReference",
				Namespace: "FHIR",
			},
		}, {
			Name: "SubSite",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponseAddItemBodySite",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponseAddItemDetail) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "traceNumber") {
		for _, v := range r.TraceNumber {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "revenue") {
		if r.Revenue != nil {
			children = append(children, *r.Revenue)
		}
	}
	if len(name) == 0 || slices.Contains(name, "productOrService") {
		if r.ProductOrService != nil {
			children = append(children, *r.ProductOrService)
		}
	}
	if len(name) == 0 || slices.Contains(name, "productOrServiceEnd") {
		if r.ProductOrServiceEnd != nil {
			children = append(children, *r.ProductOrServiceEnd)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifier") {
		for _, v := range r.Modifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "quantity") {
		if r.Quantity != nil {
			children = append(children, *r.Quantity)
		}
	}
	if len(name) == 0 || slices.Contains(name, "unitPrice") {
		if r.UnitPrice != nil {
			children = append(children, *r.UnitPrice)
		}
	}
	if len(name) == 0 || slices.Contains(name, "factor") {
		if r.Factor != nil {
			children = append(children, *r.Factor)
		}
	}
	if len(name) == 0 || slices.Contains(name, "tax") {
		if r.Tax != nil {
			children = append(children, *r.Tax)
		}
	}
	if len(name) == 0 || slices.Contains(name, "net") {
		if r.Net != nil {
			children = append(children, *r.Net)
		}
	}
	if len(name) == 0 || slices.Contains(name, "noteNumber") {
		for _, v := range r.NoteNumber {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "reviewOutcome") {
		if r.ReviewOutcome != nil {
			children = append(children, *r.ReviewOutcome)
		}
	}
	if len(name) == 0 || slices.Contains(name, "adjudication") {
		for _, v := range r.Adjudication {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "subDetail") {
		for _, v := range r.SubDetail {
			children = append(children, v)
		}
	}
	return children
}
func (r ClaimResponseAddItemDetail) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponseAddItemDetail to Boolean")
}
func (r ClaimResponseAddItemDetail) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponseAddItemDetail to String")
}
func (r ClaimResponseAddItemDetail) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponseAddItemDetail to Integer")
}
func (r ClaimResponseAddItemDetail) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponseAddItemDetail to Decimal")
}
func (r ClaimResponseAddItemDetail) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponseAddItemDetail to Date")
}
func (r ClaimResponseAddItemDetail) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponseAddItemDetail to Time")
}
func (r ClaimResponseAddItemDetail) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponseAddItemDetail to DateTime")
}
func (r ClaimResponseAddItemDetail) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponseAddItemDetail to Quantity")
}
func (r ClaimResponseAddItemDetail) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponseAddItemDetail
	switch other := other.(type) {
	case ClaimResponseAddItemDetail:
		o = &other
	case *ClaimResponseAddItemDetail:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponseAddItemDetail) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponseAddItemDetail) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "TraceNumber",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "Revenue",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ProductOrService",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ProductOrServiceEnd",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Modifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Quantity",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Quantity",
				Namespace: "FHIR",
			},
		}, {
			Name: "UnitPrice",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Money",
				Namespace: "FHIR",
			},
		}, {
			Name: "Factor",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "Tax",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Money",
				Namespace: "FHIR",
			},
		}, {
			Name: "Net",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Money",
				Namespace: "FHIR",
			},
		}, {
			Name: "NoteNumber",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "ReviewOutcome",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "ClaimResponseItemReviewOutcome",
				Namespace: "FHIR",
			},
		}, {
			Name: "Adjudication",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseItemAdjudication",
				Namespace: "FHIR",
			},
		}, {
			Name: "SubDetail",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseAddItemDetailSubDetail",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponseAddItemDetail",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponseAddItemDetailSubDetail) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "traceNumber") {
		for _, v := range r.TraceNumber {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "revenue") {
		if r.Revenue != nil {
			children = append(children, *r.Revenue)
		}
	}
	if len(name) == 0 || slices.Contains(name, "productOrService") {
		if r.ProductOrService != nil {
			children = append(children, *r.ProductOrService)
		}
	}
	if len(name) == 0 || slices.Contains(name, "productOrServiceEnd") {
		if r.ProductOrServiceEnd != nil {
			children = append(children, *r.ProductOrServiceEnd)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifier") {
		for _, v := range r.Modifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "quantity") {
		if r.Quantity != nil {
			children = append(children, *r.Quantity)
		}
	}
	if len(name) == 0 || slices.Contains(name, "unitPrice") {
		if r.UnitPrice != nil {
			children = append(children, *r.UnitPrice)
		}
	}
	if len(name) == 0 || slices.Contains(name, "factor") {
		if r.Factor != nil {
			children = append(children, *r.Factor)
		}
	}
	if len(name) == 0 || slices.Contains(name, "tax") {
		if r.Tax != nil {
			children = append(children, *r.Tax)
		}
	}
	if len(name) == 0 || slices.Contains(name, "net") {
		if r.Net != nil {
			children = append(children, *r.Net)
		}
	}
	if len(name) == 0 || slices.Contains(name, "noteNumber") {
		for _, v := range r.NoteNumber {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "reviewOutcome") {
		if r.ReviewOutcome != nil {
			children = append(children, *r.ReviewOutcome)
		}
	}
	if len(name) == 0 || slices.Contains(name, "adjudication") {
		for _, v := range r.Adjudication {
			children = append(children, v)
		}
	}
	return children
}
func (r ClaimResponseAddItemDetailSubDetail) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponseAddItemDetailSubDetail to Boolean")
}
func (r ClaimResponseAddItemDetailSubDetail) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponseAddItemDetailSubDetail to String")
}
func (r ClaimResponseAddItemDetailSubDetail) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponseAddItemDetailSubDetail to Integer")
}
func (r ClaimResponseAddItemDetailSubDetail) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponseAddItemDetailSubDetail to Decimal")
}
func (r ClaimResponseAddItemDetailSubDetail) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponseAddItemDetailSubDetail to Date")
}
func (r ClaimResponseAddItemDetailSubDetail) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponseAddItemDetailSubDetail to Time")
}
func (r ClaimResponseAddItemDetailSubDetail) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponseAddItemDetailSubDetail to DateTime")
}
func (r ClaimResponseAddItemDetailSubDetail) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponseAddItemDetailSubDetail to Quantity")
}
func (r ClaimResponseAddItemDetailSubDetail) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponseAddItemDetailSubDetail
	switch other := other.(type) {
	case ClaimResponseAddItemDetailSubDetail:
		o = &other
	case *ClaimResponseAddItemDetailSubDetail:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponseAddItemDetailSubDetail) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponseAddItemDetailSubDetail) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "TraceNumber",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "Revenue",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ProductOrService",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ProductOrServiceEnd",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Modifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Quantity",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Quantity",
				Namespace: "FHIR",
			},
		}, {
			Name: "UnitPrice",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Money",
				Namespace: "FHIR",
			},
		}, {
			Name: "Factor",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "Tax",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Money",
				Namespace: "FHIR",
			},
		}, {
			Name: "Net",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Money",
				Namespace: "FHIR",
			},
		}, {
			Name: "NoteNumber",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "ReviewOutcome",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "ClaimResponseItemReviewOutcome",
				Namespace: "FHIR",
			},
		}, {
			Name: "Adjudication",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ClaimResponseItemAdjudication",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponseAddItemDetailSubDetail",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponseTotal) Children(name ...string) fhirpath.Collection {
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
		children = append(children, r.Category)
	}
	if len(name) == 0 || slices.Contains(name, "amount") {
		children = append(children, r.Amount)
	}
	return children
}
func (r ClaimResponseTotal) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponseTotal to Boolean")
}
func (r ClaimResponseTotal) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponseTotal to String")
}
func (r ClaimResponseTotal) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponseTotal to Integer")
}
func (r ClaimResponseTotal) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponseTotal to Decimal")
}
func (r ClaimResponseTotal) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponseTotal to Date")
}
func (r ClaimResponseTotal) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponseTotal to Time")
}
func (r ClaimResponseTotal) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponseTotal to DateTime")
}
func (r ClaimResponseTotal) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponseTotal to Quantity")
}
func (r ClaimResponseTotal) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponseTotal
	switch other := other.(type) {
	case ClaimResponseTotal:
		o = &other
	case *ClaimResponseTotal:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponseTotal) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponseTotal) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Category",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Amount",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Money",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponseTotal",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponsePayment) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "adjustment") {
		if r.Adjustment != nil {
			children = append(children, *r.Adjustment)
		}
	}
	if len(name) == 0 || slices.Contains(name, "adjustmentReason") {
		if r.AdjustmentReason != nil {
			children = append(children, *r.AdjustmentReason)
		}
	}
	if len(name) == 0 || slices.Contains(name, "date") {
		if r.Date != nil {
			children = append(children, *r.Date)
		}
	}
	if len(name) == 0 || slices.Contains(name, "amount") {
		children = append(children, r.Amount)
	}
	if len(name) == 0 || slices.Contains(name, "identifier") {
		if r.Identifier != nil {
			children = append(children, *r.Identifier)
		}
	}
	return children
}
func (r ClaimResponsePayment) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponsePayment to Boolean")
}
func (r ClaimResponsePayment) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponsePayment to String")
}
func (r ClaimResponsePayment) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponsePayment to Integer")
}
func (r ClaimResponsePayment) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponsePayment to Decimal")
}
func (r ClaimResponsePayment) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponsePayment to Date")
}
func (r ClaimResponsePayment) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponsePayment to Time")
}
func (r ClaimResponsePayment) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponsePayment to DateTime")
}
func (r ClaimResponsePayment) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponsePayment to Quantity")
}
func (r ClaimResponsePayment) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponsePayment
	switch other := other.(type) {
	case ClaimResponsePayment:
		o = &other
	case *ClaimResponsePayment:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponsePayment) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponsePayment) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Adjustment",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Money",
				Namespace: "FHIR",
			},
		}, {
			Name: "AdjustmentReason",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Date",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Date",
				Namespace: "FHIR",
			},
		}, {
			Name: "Amount",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Money",
				Namespace: "FHIR",
			},
		}, {
			Name: "Identifier",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponsePayment",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponseProcessNote) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "number") {
		if r.Number != nil {
			children = append(children, *r.Number)
		}
	}
	if len(name) == 0 || slices.Contains(name, "type") {
		if r.Type != nil {
			children = append(children, *r.Type)
		}
	}
	if len(name) == 0 || slices.Contains(name, "text") {
		children = append(children, r.Text)
	}
	if len(name) == 0 || slices.Contains(name, "language") {
		if r.Language != nil {
			children = append(children, *r.Language)
		}
	}
	return children
}
func (r ClaimResponseProcessNote) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponseProcessNote to Boolean")
}
func (r ClaimResponseProcessNote) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponseProcessNote to String")
}
func (r ClaimResponseProcessNote) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponseProcessNote to Integer")
}
func (r ClaimResponseProcessNote) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponseProcessNote to Decimal")
}
func (r ClaimResponseProcessNote) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponseProcessNote to Date")
}
func (r ClaimResponseProcessNote) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponseProcessNote to Time")
}
func (r ClaimResponseProcessNote) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponseProcessNote to DateTime")
}
func (r ClaimResponseProcessNote) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponseProcessNote to Quantity")
}
func (r ClaimResponseProcessNote) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponseProcessNote
	switch other := other.(type) {
	case ClaimResponseProcessNote:
		o = &other
	case *ClaimResponseProcessNote:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponseProcessNote) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponseProcessNote) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Number",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Text",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Language",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponseProcessNote",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponseInsurance) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "sequence") {
		children = append(children, r.Sequence)
	}
	if len(name) == 0 || slices.Contains(name, "focal") {
		children = append(children, r.Focal)
	}
	if len(name) == 0 || slices.Contains(name, "coverage") {
		children = append(children, r.Coverage)
	}
	if len(name) == 0 || slices.Contains(name, "businessArrangement") {
		if r.BusinessArrangement != nil {
			children = append(children, *r.BusinessArrangement)
		}
	}
	if len(name) == 0 || slices.Contains(name, "claimResponse") {
		if r.ClaimResponse != nil {
			children = append(children, *r.ClaimResponse)
		}
	}
	return children
}
func (r ClaimResponseInsurance) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponseInsurance to Boolean")
}
func (r ClaimResponseInsurance) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponseInsurance to String")
}
func (r ClaimResponseInsurance) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponseInsurance to Integer")
}
func (r ClaimResponseInsurance) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponseInsurance to Decimal")
}
func (r ClaimResponseInsurance) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponseInsurance to Date")
}
func (r ClaimResponseInsurance) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponseInsurance to Time")
}
func (r ClaimResponseInsurance) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponseInsurance to DateTime")
}
func (r ClaimResponseInsurance) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponseInsurance to Quantity")
}
func (r ClaimResponseInsurance) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponseInsurance
	switch other := other.(type) {
	case ClaimResponseInsurance:
		o = &other
	case *ClaimResponseInsurance:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponseInsurance) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponseInsurance) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Sequence",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "Focal",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Coverage",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "BusinessArrangement",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "ClaimResponse",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponseInsurance",
			Namespace: "FHIR",
		},
	}
}
func (r ClaimResponseError) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "itemSequence") {
		if r.ItemSequence != nil {
			children = append(children, *r.ItemSequence)
		}
	}
	if len(name) == 0 || slices.Contains(name, "detailSequence") {
		if r.DetailSequence != nil {
			children = append(children, *r.DetailSequence)
		}
	}
	if len(name) == 0 || slices.Contains(name, "subDetailSequence") {
		if r.SubDetailSequence != nil {
			children = append(children, *r.SubDetailSequence)
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
func (r ClaimResponseError) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ClaimResponseError to Boolean")
}
func (r ClaimResponseError) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ClaimResponseError to String")
}
func (r ClaimResponseError) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ClaimResponseError to Integer")
}
func (r ClaimResponseError) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ClaimResponseError to Decimal")
}
func (r ClaimResponseError) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ClaimResponseError to Date")
}
func (r ClaimResponseError) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ClaimResponseError to Time")
}
func (r ClaimResponseError) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ClaimResponseError to DateTime")
}
func (r ClaimResponseError) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ClaimResponseError to Quantity")
}
func (r ClaimResponseError) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ClaimResponseError
	switch other := other.(type) {
	case ClaimResponseError:
		o = &other
	case *ClaimResponseError:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r ClaimResponseError) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r ClaimResponseError) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ItemSequence",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "DetailSequence",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "SubDetailSequence",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "Code",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Expression",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "String",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ClaimResponseError",
			Namespace: "FHIR",
		},
	}
}
