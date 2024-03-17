package types

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
)

// The amount of medication administered.
type DosageDoseAndRate struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The kind of dose or rate specified, for example, ordered or calculated.
	Type *CodeableConcept
	// Amount of medication per dose.
	Dose r4.Element
	// Amount of medication per unit of time.
	Rate r4.Element
}

func (s DosageDoseAndRate) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Base StructureDefinition for Dosage Type: Indicates how the medication is/was taken or should be taken by the patient.
type Dosage struct {
	// Supplemental instructions to the patient on how to take the medication  (e.g. "with meals" or"take half to one hour before food") or warnings for the patient about the medication (e.g. "may cause drowsiness" or "avoid exposure of skin to direct sunlight or sunlamps").
	AdditionalInstruction []CodeableConcept
	// When medication should be administered.
	Timing *Timing
	// The amount of medication administered.
	DoseAndRate []DosageDoseAndRate
	// Body site to administer to.
	Site *CodeableConcept
	// Technique for administering medication.
	Method *CodeableConcept
	// Upper limit on medication per unit of time.
	MaxDosePerPeriod *Ratio
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Free text dosage instructions e.g. SIG.
	Text *String
	// Instructions in terms that are understood by the patient or consumer.
	PatientInstruction *String
	// How drug should enter body.
	Route *CodeableConcept
	// Upper limit on medication per administration.
	MaxDosePerAdministration *Quantity
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Indicates the order in which the dosage instructions should be applied or interpreted.
	Sequence *Integer
	// Indicates whether the Medication is only taken when needed within a specific dosing schedule (Boolean option), or it indicates the precondition for taking the Medication (CodeableConcept).
	AsNeeded r4.Element
	// Upper limit on medication per lifetime of the patient.
	MaxDosePerLifetime *Quantity
}

func (s Dosage) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
