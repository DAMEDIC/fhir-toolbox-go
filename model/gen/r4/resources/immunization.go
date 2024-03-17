package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// The protocol (set of recommendations) being followed by the provider who administered the dose.
type ImmunizationProtocolApplied struct {
	// The vaccine preventable disease the dose is being administered against.
	TargetDisease []types.CodeableConcept
	// Nominal position in a series.
	DoseNumber r4.Element
	// The recommended number of doses to achieve immunity.
	SeriesDoses r4.Element
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// One possible path to achieve presumed immunity against a disease - within the context of an authority.
	Series *types.String
	// Indicates the authority who published the protocol (e.g. ACIP) that is being followed.
	Authority *types.Reference
}

func (s ImmunizationProtocolApplied) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Categorical data indicating that an adverse event is associated in time to an immunization.
type ImmunizationReaction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Date of reaction to the immunization.
	Date *types.DateTime
	// Details of the reaction.
	Detail *types.Reference
	// Self-reported indicator.
	Reported *types.Boolean
}

func (s ImmunizationReaction) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates who performed the immunization event.
type ImmunizationPerformer struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Describes the type of performance (e.g. ordering provider, administering provider, etc.).
	Function *types.CodeableConcept
	// The practitioner or organization who performed the action.
	Actor types.Reference
}

func (s ImmunizationPerformer) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Educational material presented to the patient (or guardian) at the time of vaccine administration.
type ImmunizationEducation struct {
	// Date the educational material was published.
	PublicationDate *types.DateTime
	// Date the educational material was given to the patient.
	PresentationDate *types.DateTime
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Identifier of the material presented to the patient.
	DocumentType *types.String
	// Reference pointer to the educational material given to the patient if the information was on line.
	Reference *types.Uri
}

func (s ImmunizationEducation) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Describes the event of a patient being administered a vaccine or a record of an immunization as reported by a patient, a clinician or another party.
type Immunization struct {
	// The base language in which the resource is written.
	Language *types.Code
	// The date the occurrence of the immunization was first captured in the record - potentially significantly after the occurrence of the event.
	Recorded *types.DateTime
	// The path by which the vaccine product is taken into the body.
	Route *types.CodeableConcept
	// The protocol (set of recommendations) being followed by the provider who administered the dose.
	ProtocolApplied []ImmunizationProtocolApplied
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// Reasons why the vaccine was administered.
	ReasonCode []types.CodeableConcept
	// Indicates a patient's eligibility for a funding program.
	ProgramEligibility []types.CodeableConcept
	// Indicates the source of the vaccine actually administered. This may be different than the patient eligibility (e.g. the patient may be eligible for a publically purchased vaccine but due to inventory issues, vaccine purchased with private funds was actually administered).
	FundingSource *types.CodeableConcept
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// The service delivery location where the vaccine administration occurred.
	Location *types.Reference
	// Condition, Observation or DiagnosticReport that supports why the immunization was administered.
	ReasonReference []types.Reference
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// Indicates the current status of the immunization event.
	Status types.Code
	// The patient who either received or did not receive the immunization.
	Patient types.Reference
	// The visit or admission or other contact between patient and health care provider the immunization was performed as part of.
	Encounter *types.Reference
	// The quantity of vaccine product that was administered.
	DoseQuantity *types.Quantity
	// Categorical data indicating that an adverse event is associated in time to an immunization.
	Reaction []ImmunizationReaction
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Vaccine that was administered or was to be administered.
	VaccineCode types.CodeableConcept
	// The source of the data when the report of the immunization event is not based on information from the person who administered the vaccine.
	ReportOrigin *types.CodeableConcept
	// Indicates who performed the immunization event.
	Performer []ImmunizationPerformer
	// A unique identifier assigned to this immunization record.
	Identifier []types.Identifier
	// Date vaccine administered or was to be administered.
	Occurrence r4.Element
	// Lot number of the  vaccine product.
	LotNumber *types.String
	// Body site where vaccine was administered.
	Site *types.CodeableConcept
	// Extra information about the immunization that is not conveyed by the other attributes.
	Note []types.Annotation
	// Reason why a dose is considered to be subpotent.
	SubpotentReason []types.CodeableConcept
	// Educational material presented to the patient (or guardian) at the time of vaccine administration.
	Education []ImmunizationEducation
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// An indication that the content of the record is based on information from the person who administered the vaccine. This reflects the context under which the data was originally recorded.
	PrimarySource *types.Boolean
	// Indication if a dose is considered to be subpotent. By default, a dose should be considered to be potent.
	IsSubpotent *types.Boolean
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// Indicates the reason the immunization event was not performed.
	StatusReason *types.CodeableConcept
	// Name of vaccine manufacturer.
	Manufacturer *types.Reference
	// Date vaccine batch expires.
	ExpirationDate *types.Date
}

func (s Immunization) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
