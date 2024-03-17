package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
type ExplanationOfBenefitSupportingInfo struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A number to uniquely identify supporting information entries.
	Sequence types.PositiveInt
	// The date when or period to which this information refers.
	Timing r4.Element
	// Provides the reason in the situation where a reason code is required in addition to the content.
	Reason *types.Coding
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// The general class of the information supplied: information; exception; accident, employment; onset, etc.
	Category types.CodeableConcept
	// System and code pertaining to the specific information regarding special conditions relating to the setting, treatment or patient  for which care is sought.
	Code *types.CodeableConcept
	// Additional data or information such as resources, documents, images etc. including references to the data or the actual inclusion of the data.
	Value r4.Element
}

func (s ExplanationOfBenefitSupportingInfo) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Categorized monetary totals for the adjudication.
type ExplanationOfBenefitTotal struct {
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

func (s ExplanationOfBenefitTotal) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Third-tier of goods and services.
type ExplanationOfBenefitItemDetailSubDetail struct {
	// The type of revenue or cost center providing the product and/or service.
	Revenue *types.CodeableConcept
	// The quantity times the unit price for an additional service or product or charge.
	Net *types.Money
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Identifies the program under which this may be recovered.
	ProgramCode []types.CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *types.Quantity
	// Unique Device Identifiers associated with this line item.
	Udi []types.Reference
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []types.PositiveInt
	// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
	Sequence types.PositiveInt
	// Code to identify the general type of benefits under which products and services are provided.
	Category *types.CodeableConcept
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *types.Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *types.Decimal
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService types.CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []types.CodeableConcept
	// The adjudication results.
	Adjudication []ExplanationOfBenefitItemAdjudication
}

func (s ExplanationOfBenefitItemDetailSubDetail) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Second-tier of goods and services.
type ExplanationOfBenefitItemDetail struct {
	// Code to identify the general type of benefits under which products and services are provided.
	Category *types.CodeableConcept
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService types.CodeableConcept
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *types.Money
	// The quantity times the unit price for an additional service or product or charge.
	Net *types.Money
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []types.PositiveInt
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// A claim detail line. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
	Sequence types.PositiveInt
	// The type of revenue or cost center providing the product and/or service.
	Revenue *types.CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *types.Quantity
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *types.Decimal
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []types.CodeableConcept
	// Third-tier of goods and services.
	SubDetail []ExplanationOfBenefitItemDetailSubDetail
	// The adjudication results.
	Adjudication []ExplanationOfBenefitItemAdjudication
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Identifies the program under which this may be recovered.
	ProgramCode []types.CodeableConcept
	// Unique Device Identifiers associated with this line item.
	Udi []types.Reference
}

func (s ExplanationOfBenefitItemDetail) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
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
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A code to indicate the information type of this adjudication record. Information types may include: the value submitted, maximum values or percentages allowed or payable under the plan, amounts that the patient is responsible for in-aggregate or pertaining to this item, amounts paid by other coverages, and the benefit payable for this item.
	Category types.CodeableConcept
	// A code supporting the understanding of the adjudication result and explaining variance from expected amount.
	Reason *types.CodeableConcept
	// Monetary amount associated with the category.
	Amount *types.Money
	// A non-monetary value associated with the category. Mutually exclusive to the amount element above.
	Value *types.Decimal
}

func (s ExplanationOfBenefitItemAdjudication) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
type ExplanationOfBenefitItem struct {
	// The date or dates when the service or product was supplied, performed or completed.
	Serviced r4.Element
	// A region or surface of the bodySite, e.g. limb region or tooth surface(s).
	SubSite []types.CodeableConcept
	// A billed item may include goods or services provided in multiple encounters.
	Encounter []types.Reference
	// Second-tier of goods and services.
	Detail []ExplanationOfBenefitItemDetail
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Exceptions, special conditions and supporting information applicable for this service or product.
	InformationSequence []types.PositiveInt
	// Code to identify the general type of benefits under which products and services are provided.
	Category *types.CodeableConcept
	// Identifies the program under which this may be recovered.
	ProgramCode []types.CodeableConcept
	// If this item is a group then the values here are a summary of the adjudication of the detail items. If this item is a simple product or service then this is the result of the adjudication of this item.
	Adjudication []ExplanationOfBenefitItemAdjudication
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A number to uniquely identify item entries.
	Sequence types.PositiveInt
	// Procedures applicable for this service or product.
	ProcedureSequence []types.PositiveInt
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []types.CodeableConcept
	// Unique Device Identifiers associated with this line item.
	Udi []types.Reference
	// Physical service site on the patient (limb, tooth, etc.).
	BodySite *types.CodeableConcept
	// Care team members related to this service or product.
	CareTeamSequence []types.PositiveInt
	// The type of revenue or cost center providing the product and/or service.
	Revenue *types.CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *types.Quantity
	// The quantity times the unit price for an additional service or product or charge.
	Net *types.Money
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *types.Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *types.Decimal
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []types.PositiveInt
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Diagnoses applicable for this service or product.
	DiagnosisSequence []types.PositiveInt
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService types.CodeableConcept
	// Where the product or service was provided.
	Location r4.Element
}

func (s ExplanationOfBenefitItem) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Benefits Used to date.
type ExplanationOfBenefitBenefitBalanceFinancial struct {
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
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
}

func (s ExplanationOfBenefitBenefitBalanceFinancial) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Balance by Benefit Category.
type ExplanationOfBenefitBenefitBalance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Code to identify the general type of benefits under which products and services are provided.
	Category types.CodeableConcept
	// True if the indicated class of service is excluded from the plan, missing or False indicates the product or service is included in the coverage.
	Excluded *types.Boolean
	// Is a flag to indicate whether the benefits refer to in-network providers or out-of-network providers.
	Network *types.CodeableConcept
	// The term or period of the values such as 'maximum lifetime benefit' or 'maximum annual visits'.
	Term *types.CodeableConcept
	// Benefits Used to date.
	Financial []ExplanationOfBenefitBenefitBalanceFinancial
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A short name or tag for the benefit.
	Name *types.String
	// A richer description of the benefit or services covered.
	Description *types.String
	// Indicates if the benefits apply to an individual or to the family.
	Unit *types.CodeableConcept
}

func (s ExplanationOfBenefitBenefitBalance) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
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

func (s ExplanationOfBenefitProcedure) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
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
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Type of Party to be reimbursed: Subscriber, provider, other.
	Type *types.CodeableConcept
	// Reference to the individual or organization to whom any payment will be made.
	Party *types.Reference
}

func (s ExplanationOfBenefitPayee) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Details of a accident which resulted in injuries which required the products and services listed in the claim.
type ExplanationOfBenefitAccident struct {
	// Date of an accident event  related to the products and services contained in the claim.
	Date *types.Date
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
}

func (s ExplanationOfBenefitAccident) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The third-tier service adjudications for payor added services.
type ExplanationOfBenefitAddItemDetailSubDetail struct {
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService types.CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []types.CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *types.Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *types.Money
	// The quantity times the unit price for an additional service or product or charge.
	Net *types.Money
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []types.PositiveInt
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *types.Decimal
	// The adjudication results.
	Adjudication []ExplanationOfBenefitItemAdjudication
}

func (s ExplanationOfBenefitAddItemDetailSubDetail) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The second-tier service adjudications for payor added services.
type ExplanationOfBenefitAddItemDetail struct {
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *types.Money
	// The third-tier service adjudications for payor added services.
	SubDetail []ExplanationOfBenefitAddItemDetailSubDetail
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []types.CodeableConcept
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService types.CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *types.Quantity
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *types.Decimal
	// The quantity times the unit price for an additional service or product or charge.
	Net *types.Money
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []types.PositiveInt
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// The adjudication results.
	Adjudication []ExplanationOfBenefitItemAdjudication
}

func (s ExplanationOfBenefitAddItemDetail) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The first-tier service adjudications for payor added product or service lines.
type ExplanationOfBenefitAddItem struct {
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *types.Decimal
	// The quantity times the unit price for an additional service or product or charge.
	Net *types.Money
	// Claim items which this service line is intended to replace.
	ItemSequence []types.PositiveInt
	// The sequence number of the details within the claim item which this line is intended to replace.
	DetailSequence []types.PositiveInt
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService types.CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []types.CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *types.Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *types.Money
	// The sequence number of the sub-details woithin the details within the claim item which this line is intended to replace.
	SubDetailSequence []types.PositiveInt
	// The date or dates when the service or product was supplied, performed or completed.
	Serviced r4.Element
	// Physical service site on the patient (limb, tooth, etc.).
	BodySite *types.CodeableConcept
	// A region or surface of the bodySite, e.g. limb region or tooth surface(s).
	SubSite []types.CodeableConcept
	// The adjudication results.
	Adjudication []ExplanationOfBenefitItemAdjudication
	// The second-tier service adjudications for payor added services.
	Detail []ExplanationOfBenefitAddItemDetail
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// The providers who are authorized for the services rendered to the patient.
	Provider []types.Reference
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []types.PositiveInt
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Identifies the program under which this may be recovered.
	ProgramCode []types.CodeableConcept
	// Where the product or service was provided.
	Location r4.Element
}

func (s ExplanationOfBenefitAddItem) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Other claims which are related to this claim such as prior submissions or claims for related services or for the same event.
type ExplanationOfBenefitRelated struct {
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
	// Reference to a related claim.
	Claim *types.Reference
	// A code to convey how the claims are related.
	Relationship *types.CodeableConcept
}

func (s ExplanationOfBenefitRelated) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
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

func (s ExplanationOfBenefitCareTeam) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A note that describes or explains adjudication results in a human readable form.
type ExplanationOfBenefitProcessNote struct {
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
	// The explanation or description associated with the processing.
	Text *types.String
	// A code to define the language used in the text of the note.
	Language *types.CodeableConcept
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
}

func (s ExplanationOfBenefitProcessNote) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about diagnoses relevant to the claim items.
type ExplanationOfBenefitDiagnosis struct {
	// A number to uniquely identify diagnosis entries.
	Sequence types.PositiveInt
	// The nature of illness or problem in a coded form or as a reference to an external defined Condition.
	Diagnosis r4.Element
	// When the condition was observed or the relative ranking.
	Type []types.CodeableConcept
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
}

func (s ExplanationOfBenefitDiagnosis) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
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
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A flag to indicate that this Coverage is to be used for adjudication of this claim when set to true.
	Focal types.Boolean
	// Reference to the insurance card level information contained in the Coverage resource. The coverage issuing insurer will use these details to locate the patient's actual coverage within the insurer's information system.
	Coverage types.Reference
	// Reference numbers previously provided by the insurer to the provider to be quoted on subsequent claims containing services or products related to the prior authorization.
	PreAuthRef []types.String
}

func (s ExplanationOfBenefitInsurance) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Payment details for the adjudication of the claim.
type ExplanationOfBenefitPayment struct {
	// Reason for the payment adjustment.
	AdjustmentReason *types.CodeableConcept
	// Issuer's unique identifier for the payment instrument.
	Identifier *types.Identifier
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Total amount of all adjustments to this payment included in this transaction which are not related to this claim's adjudication.
	Adjustment *types.Money
	// Estimated date the payment will be issued or the actual issue date of payment.
	Date *types.Date
	// Benefits payable less any payment adjustment.
	Amount *types.Money
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Whether this represents partial or complete payment of the benefits payable.
	Type *types.CodeableConcept
}

func (s ExplanationOfBenefitPayment) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// This resource provides: the claim details; adjudication details from the processing of a Claim; and optionally account balance information, for informing the subscriber of the benefits provided.
type ExplanationOfBenefit struct {
	// A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
	SubType *types.CodeableConcept
	// A code to indicate whether the nature of the request is: to request adjudication of products and services previously rendered; or requesting authorization and adjudication for provision in the future; or requesting the non-binding adjudication of the listed products and services which could be provided in the future.
	Use types.Code
	// A code to indicate whether and for whom funds are to be reserved for future claims.
	FundsReserveRequested *types.CodeableConcept
	// Additional information codes regarding exceptions, special considerations, the condition, situation, prior or concurrent issues.
	SupportingInfo []ExplanationOfBenefitSupportingInfo
	// Categorized monetary totals for the adjudication.
	Total []ExplanationOfBenefitTotal
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// A unique identifier assigned to this explanation of benefit.
	Identifier []types.Identifier
	// The status of the resource instance.
	Status types.Code
	// Prescription to support the dispensing of pharmacy, device or vision products.
	Prescription *types.Reference
	// The timeframe during which the supplied preauthorization reference may be quoted on claims to obtain the adjudication as provided.
	PreAuthRefPeriod []types.Period
	// A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
	Item []ExplanationOfBenefitItem
	// Balance by Benefit Category.
	BenefitBalance []ExplanationOfBenefitBenefitBalance
	// Procedures performed on the patient relevant to the billing items with the claim.
	Procedure []ExplanationOfBenefitProcedure
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// The period for which charges are being submitted.
	BillablePeriod *types.Period
	// The party to be reimbursed for cost of the products and services according to the terms of the policy.
	Payee *ExplanationOfBenefitPayee
	// A reference to a referral resource.
	Referral *types.Reference
	// The business identifier for the instance of the adjudication request: claim predetermination or preauthorization.
	Claim *types.Reference
	// Reference from the Insurer which is used in later communications which refers to this adjudication.
	PreAuthRef []types.String
	// Details of a accident which resulted in injuries which required the products and services listed in the claim.
	Accident *ExplanationOfBenefitAccident
	// The first-tier service adjudications for payor added product or service lines.
	AddItem []ExplanationOfBenefitAddItem
	// The date this resource was created.
	Created types.DateTime
	// The provider which is responsible for the claim, predetermination or preauthorization.
	Provider types.Reference
	// This indicates the relative order of a series of EOBs related to different coverages for the same suite of services.
	Precedence *types.PositiveInt
	// The term of the benefits documented in this response.
	BenefitPeriod *types.Period
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The party to whom the professional services and/or products have been supplied or are being considered and for whom actual for forecast reimbursement is sought.
	Patient types.Reference
	// A code, used only on a response to a preauthorization, to indicate whether the benefits payable have been reserved and for whom.
	FundsReserve *types.CodeableConcept
	// Other claims which are related to this claim such as prior submissions or claims for related services or for the same event.
	Related []ExplanationOfBenefitRelated
	// The actual form, by reference or inclusion, for printing the content or an EOB.
	Form *types.Attachment
	// The base language in which the resource is written.
	Language *types.Code
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// Individual who created the claim, predetermination or preauthorization.
	Enterer *types.Reference
	// The outcome of the claim, predetermination, or preauthorization processing.
	Outcome types.Code
	// The members of the team who provided the products and services.
	CareTeam []ExplanationOfBenefitCareTeam
	// The adjudication results which are presented at the header level rather than at the line-item or add-item levels.
	Adjudication []ExplanationOfBenefitItemAdjudication
	// A note that describes or explains adjudication results in a human readable form.
	ProcessNote []ExplanationOfBenefitProcessNote
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// The category of claim, e.g. oral, pharmacy, vision, institutional, professional.
	Type types.CodeableConcept
	// The party responsible for authorization, adjudication and reimbursement.
	Insurer types.Reference
	// The business identifier for the instance of the adjudication response: claim, predetermination or preauthorization response.
	ClaimResponse *types.Reference
	// A human readable description of the status of the adjudication.
	Disposition *types.String
	// Information about diagnoses relevant to the claim items.
	Diagnosis []ExplanationOfBenefitDiagnosis
	// A code for the form to be used for printing the content.
	FormCode *types.CodeableConcept
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// The provider-required urgency of processing the request. Typical values include: stat, routine deferred.
	Priority *types.CodeableConcept
	// Original prescription which has been superseded by this prescription to support the dispensing of pharmacy services, medications or products.
	OriginalPrescription *types.Reference
	// Facility where the services were provided.
	Facility *types.Reference
	// Financial instruments for reimbursement for the health care products and services specified on the claim.
	Insurance []ExplanationOfBenefitInsurance
	// Payment details for the adjudication of the claim.
	Payment *ExplanationOfBenefitPayment
}

func (s ExplanationOfBenefit) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
