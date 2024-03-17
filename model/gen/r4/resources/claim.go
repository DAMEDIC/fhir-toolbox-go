package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Other claims which are related to this claim such as prior submissions or claims for related services or for the same event.
type ClaimRelated struct {
	// Reference to a related claim.
	Claim *types.Reference
	// A code to convey how the claims are related.
	Relationship *types.CodeableConcept
	// An alternate organizational reference to the case or file to which this particular claim pertains.
	Reference *types.Identifier
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
}

func (s ClaimRelated) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
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
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A number to uniquely identify care team entries.
	Sequence types.PositiveInt
	// Member of the team who provided the product or service.
	Provider types.Reference
	// The party who is billing and/or responsible for the claimed products or services.
	Responsible *types.Boolean
	// The lead, assisting or supervising practitioner and their discipline if a multidisciplinary team.
	Role *types.CodeableConcept
	// The qualification of the practitioner which is applicable for this service.
	Qualification *types.CodeableConcept
}

func (s ClaimCareTeam) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Financial instruments for reimbursement for the health care products and services specified on the claim.
type ClaimInsurance struct {
	// A flag to indicate that this Coverage is to be used for adjudication of this claim when set to true.
	Focal types.Boolean
	// Reference numbers previously provided by the insurer to the provider to be quoted on subsequent claims containing services or products related to the prior authorization.
	PreAuthRef []types.String
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A number to uniquely identify insurance entries and provide a sequence of coverages to convey coordination of benefit order.
	Sequence types.PositiveInt
	// The business identifier to be used when the claim is sent for adjudication against this insurance policy.
	Identifier *types.Identifier
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
}

func (s ClaimInsurance) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
type ClaimSupportingInfo struct {
	// System and code pertaining to the specific information regarding special conditions relating to the setting, treatment or patient  for which care is sought.
	Code *types.CodeableConcept
	// The date when or period to which this information refers.
	Timing r4.Element
	// Provides the reason in the situation where a reason code is required in addition to the content.
	Reason *types.CodeableConcept
	// A number to uniquely identify supporting information entries.
	Sequence types.PositiveInt
	// The general class of the information supplied: information; exception; accident, employment; onset, etc.
	Category types.CodeableConcept
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
	Value r4.Element
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
}

func (s ClaimSupportingInfo) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
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
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A number to uniquely identify procedure entries.
	Sequence types.PositiveInt
	// When the condition was observed or the relative ranking.
	Type []types.CodeableConcept
	// Date and optionally time the procedure was performed.
	Date *types.DateTime
	// The code or reference to a Procedure resource which identifies the clinical intervention performed.
	Procedure r4.Element
	// Unique Device Identifiers associated with this line item.
	Udi []types.Reference
}

func (s ClaimProcedure) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about diagnoses relevant to the claim items.
type ClaimDiagnosis struct {
	// Indication of whether the diagnosis was present on admission to a facility.
	OnAdmission *types.CodeableConcept
	// A package billing code or bundle code used to group products and services to a particular health condition (such as heart attack) which is based on a predetermined grouping code system.
	PackageCode *types.CodeableConcept
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A number to uniquely identify diagnosis entries.
	Sequence types.PositiveInt
	// The nature of illness or problem in a coded form or as a reference to an external defined Condition.
	Diagnosis r4.Element
	// When the condition was observed or the relative ranking.
	Type []types.CodeableConcept
}

func (s ClaimDiagnosis) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
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
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Type of Party to be reimbursed: subscriber, provider, other.
	Type types.CodeableConcept
	// Reference to the individual or organization to whom any payment will be made.
	Party *types.Reference
}

func (s ClaimPayee) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Details of an accident which resulted in injuries which required the products and services listed in the claim.
type ClaimAccident struct {
	// The type or context of the accident event for the purposes of selection of potential insurance coverages and determination of coordination between insurers.
	Type *types.CodeableConcept
	// The physical location of the accident event.
	Location r4.Element
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Date of an accident event  related to the products and services contained in the claim.
	Date types.Date
}

func (s ClaimAccident) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
type ClaimItemDetailSubDetail struct {
	// The type of revenue or cost center providing the product and/or service.
	Revenue *types.CodeableConcept
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService types.CodeableConcept
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *types.Money
	// Unique Device Identifiers associated with this line item.
	Udi []types.Reference
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Identifies the program under which this may be recovered.
	ProgramCode []types.CodeableConcept
	// The quantity times the unit price for an additional service or product or charge.
	Net *types.Money
	// A number to uniquely identify item entries.
	Sequence types.PositiveInt
	// Code to identify the general type of benefits under which products and services are provided.
	Category *types.CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []types.CodeableConcept
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The number of repetitions of a service or product.
	Quantity *types.Quantity
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *types.Decimal
}

func (s ClaimItemDetailSubDetail) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
type ClaimItemDetail struct {
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *types.Decimal
	// The quantity times the unit price for an additional service or product or charge.
	Net *types.Money
	// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
	SubDetail []ClaimItemDetailSubDetail
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The type of revenue or cost center providing the product and/or service.
	Revenue *types.CodeableConcept
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A number to uniquely identify item entries.
	Sequence types.PositiveInt
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []types.CodeableConcept
	// Identifies the program under which this may be recovered.
	ProgramCode []types.CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *types.Quantity
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Code to identify the general type of benefits under which products and services are provided.
	Category *types.CodeableConcept
	// Unique Device Identifiers associated with this line item.
	Udi []types.Reference
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService types.CodeableConcept
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *types.Money
}

func (s ClaimItemDetail) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A claim line. Either a simple  product or service or a 'group' of details which can each be a simple items or groups of sub-details.
type ClaimItem struct {
	// The type of revenue or cost center providing the product and/or service.
	Revenue *types.CodeableConcept
	// The date or dates when the service or product was supplied, performed or completed.
	Serviced r4.Element
	// The number of repetitions of a service or product.
	Quantity *types.Quantity
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *types.Decimal
	// A region or surface of the bodySite, e.g. limb region or tooth surface(s).
	SubSite []types.CodeableConcept
	// The Encounters during which this Claim was created or to which the creation of this record is tightly associated.
	Encounter []types.Reference
	// Exceptions, special conditions and supporting information applicable for this service or product.
	InformationSequence []types.PositiveInt
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService types.CodeableConcept
	// Where the product or service was provided.
	Location r4.Element
	// Unique Device Identifiers associated with this line item.
	Udi []types.Reference
	// Physical service site on the patient (limb, tooth, etc.).
	BodySite *types.CodeableConcept
	// A number to uniquely identify item entries.
	Sequence types.PositiveInt
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Diagnosis applicable for this service or product.
	DiagnosisSequence []types.PositiveInt
	// Identifies the program under which this may be recovered.
	ProgramCode []types.CodeableConcept
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *types.Money
	// The quantity times the unit price for an additional service or product or charge.
	Net *types.Money
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// CareTeam members related to this service or product.
	CareTeamSequence []types.PositiveInt
	// Procedures applicable for this service or product.
	ProcedureSequence []types.PositiveInt
	// Code to identify the general type of benefits under which products and services are provided.
	Category *types.CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []types.CodeableConcept
	// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
	Detail []ClaimItemDetail
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
}

func (s ClaimItem) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A provider issued list of professional services and products which have been provided, or are to be provided, to a patient which is sent to an insurer for reimbursement.
//
// The Claim resource is used by providers to exchange services and products rendered to patients or planned to be rendered with insurers for reimbuserment. It is also used by insurers to exchange claims information with statutory reporting and data analytics firms.
type Claim struct {
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A unique identifier assigned to this claim.
	Identifier []types.Identifier
	// Other claims which are related to this claim such as prior submissions or claims for related services or for the same event.
	Related []ClaimRelated
	// The members of the team who provided the products and services.
	CareTeam []ClaimCareTeam
	// Financial instruments for reimbursement for the health care products and services specified on the claim.
	Insurance []ClaimInsurance
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// The period for which charges are being submitted.
	BillablePeriod *types.Period
	// Facility where the services were provided.
	Facility *types.Reference
	// The total value of the all the items in the claim.
	Total *types.Money
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// A code to indicate whether the nature of the request is: to request adjudication of products and services previously rendered; or requesting authorization and adjudication for provision in the future; or requesting the non-binding adjudication of the listed products and services which could be provided in the future.
	Use types.Code
	// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
	SupportingInfo []ClaimSupportingInfo
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
	SubType *types.CodeableConcept
	// The date this resource was created.
	Created types.DateTime
	// Procedures performed on the patient relevant to the billing items with the claim.
	Procedure []ClaimProcedure
	// Prescription to support the dispensing of pharmacy, device or vision products.
	Prescription *types.Reference
	// The category of claim, e.g. oral, pharmacy, vision, institutional, professional.
	Type types.CodeableConcept
	// Original prescription which has been superseded by this prescription to support the dispensing of pharmacy services, medications or products.
	OriginalPrescription *types.Reference
	// Information about diagnoses relevant to the claim items.
	Diagnosis []ClaimDiagnosis
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// The provider-required urgency of processing the request. Typical values include: stat, routine deferred.
	Priority types.CodeableConcept
	// The party to be reimbursed for cost of the products and services according to the terms of the policy.
	Payee *ClaimPayee
	// Details of an accident which resulted in injuries which required the products and services listed in the claim.
	Accident *ClaimAccident
	// The base language in which the resource is written.
	Language *types.Code
	// The status of the resource instance.
	Status types.Code
	// The party to whom the professional services and/or products have been supplied or are being considered and for whom actual or forecast reimbursement is sought.
	Patient types.Reference
	// Individual who created the claim, predetermination or preauthorization.
	Enterer *types.Reference
	// The Insurer who is target of the request.
	Insurer *types.Reference
	// The provider which is responsible for the claim, predetermination or preauthorization.
	Provider types.Reference
	// A code to indicate whether and for whom funds are to be reserved for future claims.
	FundsReserve *types.CodeableConcept
	// A reference to a referral resource.
	Referral *types.Reference
	// A claim line. Either a simple  product or service or a 'group' of details which can each be a simple items or groups of sub-details.
	Item []ClaimItem
}

func (s Claim) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
