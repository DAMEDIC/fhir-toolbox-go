package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Indicates who or what performed the event.
type MedicationDispensePerformer struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Distinguishes the type of performer in the dispense.  For example, date enterer, packager, final checker.
	Function *types.CodeableConcept
	// The device, practitioner, etc. who performed the action.  It should be assumed that the actor is the dispenser of the medication.
	Actor types.Reference
}

func (s MedicationDispensePerformer) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates whether or not substitution was made as part of the dispense.  In some cases, substitution will be expected but does not happen, in other cases substitution is not expected but does happen.  This block explains what substitution did or did not happen and why.  If nothing is specified, substitution was not done.
type MedicationDispenseSubstitution struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// True if the dispenser dispensed a different drug or product from what was prescribed.
	WasSubstituted types.Boolean
	// A code signifying whether a different drug was dispensed from what was prescribed.
	Type *types.CodeableConcept
	// Indicates the reason for the substitution (or lack of substitution) from what was prescribed.
	Reason []types.CodeableConcept
	// The person or organization that has primary responsibility for the substitution.
	ResponsibleParty []types.Reference
}

func (s MedicationDispenseSubstitution) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates that a medication product is to be or has been dispensed for a named person/patient.  This includes a description of the medication product (supply) provided and the instructions for administering the medication.  The medication dispense is the result of a pharmacy system responding to a medication order.
type MedicationDispense struct {
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// The amount of medication expressed as a timing amount.
	DaysSupply *types.Quantity
	// The time the dispensed product was provided to the patient or their representative.
	WhenHandedOver *types.DateTime
	// Identification of the facility/location where the medication was shipped to, as part of the dispense event.
	Destination *types.Reference
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A code specifying the state of the set of dispense events.
	Status types.Code
	// The encounter or episode of care that establishes the context for this event.
	Context *types.Reference
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// The procedure that trigger the dispense.
	PartOf []types.Reference
	// Identifies the medication being administered. This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
	Medication r4.Element
	// Indicates who or what performed the event.
	Performer []MedicationDispensePerformer
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Indicates the reason why a dispense was not performed.
	StatusReason r4.Element
	// Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions for a patient; e.g. drug-drug interaction, duplicate therapy, dosage alert etc.
	DetectedIssue []types.Reference
	// Identifiers associated with this Medication Dispense that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate. They are business identifiers assigned to this resource by the performer or other systems and remain constant as the resource is updated and propagates from server to server.
	Identifier []types.Identifier
	// Additional information that supports the medication being dispensed.
	SupportingInformation []types.Reference
	// The principal physical location where the dispense was performed.
	Location *types.Reference
	// Identifies the person who picked up the medication.  This will usually be a patient or their caregiver, but some cases exist where it can be a healthcare professional.
	Receiver []types.Reference
	// Indicates how the medication is to be used by the patient.
	DosageInstruction []types.Dosage
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// Indicates the type of medication dispense (for example, where the medication is expected to be consumed or administered (i.e. inpatient or outpatient)).
	Category *types.CodeableConcept
	// The time when the dispensed product was packaged and reviewed.
	WhenPrepared *types.DateTime
	// A link to a resource representing the person or the group to whom the medication will be given.
	Subject *types.Reference
	// Indicates the type of dispensing event that is performed. For example, Trial Fill, Completion of Trial, Partial Fill, Emergency Fill, Samples, etc.
	Type *types.CodeableConcept
	// The amount of medication that has been dispensed. Includes unit of measure.
	Quantity *types.Quantity
	// A summary of the events of interest that have occurred, such as when the dispense was verified.
	EventHistory []types.Reference
	// Indicates whether or not substitution was made as part of the dispense.  In some cases, substitution will be expected but does not happen, in other cases substitution is not expected but does happen.  This block explains what substitution did or did not happen and why.  If nothing is specified, substitution was not done.
	Substitution *MedicationDispenseSubstitution
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// The base language in which the resource is written.
	Language *types.Code
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// Indicates the medication order that is being dispensed against.
	AuthorizingPrescription []types.Reference
	// Extra information about the dispense that could not be conveyed in the other attributes.
	Note []types.Annotation
}

func (s MedicationDispense) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
