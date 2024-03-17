package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Indicates who or what performed or participated in the charged service.
type ChargeItemPerformer struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Describes the type of performance or participation(e.g. primary surgeon, anesthesiologiest, etc.).
	Function *types.CodeableConcept
	// The device, practitioner, etc. who performed or participated in the service.
	Actor types.Reference
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
}

func (s ChargeItemPerformer) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The resource ChargeItem describes the provision of healthcare provider products for a certain patient, therefore referring not only to the product, but containing in addition details of the provision, like date, time, amounts and participating organizations and persons. Main Usage of the ChargeItem is to enable the billing process and internal cost allocation.
type ChargeItem struct {
	// The financial cost center permits the tracking of charge attribution.
	CostCenter *types.Reference
	// Quantity of which the charge item has been serviced.
	Quantity *types.Quantity
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// The base language in which the resource is written.
	Language *types.Code
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Identifiers assigned to this event performer or other systems.
	Identifier []types.Identifier
	// The encounter or episode of care that establishes the context for this event.
	Context *types.Reference
	// The organization requesting the service.
	PerformingOrganization *types.Reference
	// References the source of pricing information, rules of application for the code this ChargeItem uses.
	DefinitionCanonical []types.Canonical
	// The individual or set of individuals the action is being or was performed on.
	Subject types.Reference
	// Date/time(s) or duration when the charged service was applied.
	Occurrence r4.Element
	// The organization performing the service.
	RequestingOrganization *types.Reference
	// Factor overriding the factor determined by the rules associated with the code.
	FactorOverride *types.Decimal
	// Total price of the charge overriding the list price associated with the code.
	PriceOverride *types.Money
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// A code that identifies the charge, like a billing code.
	Code types.CodeableConcept
	// The device, practitioner, etc. who entered the charge item.
	Enterer *types.Reference
	// Identifies the device, food, drug or other product being charged either by type code or reference to an instance.
	Product r4.Element
	// Account into which this ChargeItems belongs.
	Account []types.Reference
	// References the (external) source of pricing information, rules of application for the code this ChargeItem uses.
	DefinitionUri []types.Uri
	// The anatomical location where the related service has been applied.
	Bodysite []types.CodeableConcept
	// Indicated the rendered service that caused this charge.
	Service []types.Reference
	// Comments made about the event by the performer, subject or other participants.
	Note []types.Annotation
	// Date the charge item was entered.
	EnteredDate *types.DateTime
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// If the list price or the rule-based factor associated with the code is overridden, this attribute can capture a text to indicate the  reason for this action.
	OverrideReason *types.String
	// Further information supporting this charge.
	SupportingInformation []types.Reference
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The current state of the ChargeItem.
	Status types.Code
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// ChargeItems can be grouped to larger ChargeItems covering the whole set.
	PartOf []types.Reference
	// Indicates who or what performed or participated in the charged service.
	Performer []ChargeItemPerformer
	// Describes why the event occurred in coded or textual form.
	Reason []types.CodeableConcept
}

func (s ChargeItem) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
