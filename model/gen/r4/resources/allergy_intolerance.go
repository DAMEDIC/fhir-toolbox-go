package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Details about each adverse reaction event linked to exposure to the identified substance.
type AllergyIntoleranceReaction struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Identification of the specific substance (or pharmaceutical product) considered to be responsible for the Adverse Reaction event. Note: the substance for a specific reaction may be different from the substance identified as the cause of the risk, but it must be consistent with it. For instance, it may be a more specific substance (e.g. a brand medication) or a composite product that includes the identified substance. It must be clinically safe to only process the 'code' and ignore the 'reaction.substance'.  If a receiving system is unable to confirm that AllergyIntolerance.reaction.substance falls within the semantic scope of AllergyIntolerance.code, then the receiving system should ignore AllergyIntolerance.reaction.substance.
	Substance *types.CodeableConcept
	// Clinical symptoms and/or signs that are observed or associated with the adverse reaction event.
	Manifestation []types.CodeableConcept
	// Text description about the reaction as a whole, including details of the manifestation if required.
	Description *types.String
	// Record of the date and/or time of the onset of the Reaction.
	Onset *types.DateTime
	// Clinical assessment of the severity of the reaction event as a whole, potentially considering multiple different manifestations.
	Severity *types.Code
	// Additional text about the adverse reaction event not captured in other fields.
	Note []types.Annotation
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Identification of the route by which the subject was exposed to the substance.
	ExposureRoute *types.CodeableConcept
}

func (s AllergyIntoleranceReaction) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Risk of harmful or undesirable, physiological response which is unique to an individual and associated with exposure to a substance.
//
// To record a clinical assessment of a propensity, or potential risk to an individual, of an adverse reaction upon future exposure to the specified substance, or class of substance.
type AllergyIntolerance struct {
	// Identification of the underlying physiological mechanism for the reaction risk.
	Type *types.Code
	// The source of the information about the allergy that is recorded.
	Asserter *types.Reference
	// Represents the date and/or time of the last known occurrence of a reaction event.
	LastOccurrence *types.DateTime
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// The clinical status of the allergy or intolerance.
	ClinicalStatus *types.CodeableConcept
	// Code for an allergy or intolerance statement (either a positive or a negated/excluded statement).  This may be a code for a substance or pharmaceutical product that is considered to be responsible for the adverse reaction risk (e.g., "Latex"), an allergy or intolerance condition (e.g., "Latex allergy"), or a negated/excluded code for a specific substance or class (e.g., "No latex allergy") or a general or categorical negated statement (e.g.,  "No known allergy", "No known drug allergies").  Note: the substance for a specific reaction may be different from the substance identified as the cause of the risk, but it must be consistent with it. For instance, it may be a more specific substance (e.g. a brand medication) or a composite product that includes the identified substance. It must be clinically safe to only process the 'code' and ignore the 'reaction.substance'.  If a receiving system is unable to confirm that AllergyIntolerance.reaction.substance falls within the semantic scope of AllergyIntolerance.code, then the receiving system should ignore AllergyIntolerance.reaction.substance.
	Code *types.CodeableConcept
	// Additional narrative about the propensity for the Adverse Reaction, not captured in other fields.
	Note []types.Annotation
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Business identifiers assigned to this AllergyIntolerance by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier []types.Identifier
	// Estimate of the potential clinical harm, or seriousness, of the reaction to the identified substance.
	Criticality *types.Code
	// The encounter when the allergy or intolerance was asserted.
	Encounter *types.Reference
	// Estimated or actual date,  date-time, or age when allergy or intolerance was identified.
	Onset r4.Element
	// Individual who recorded the record and takes responsibility for its content.
	Recorder *types.Reference
	// The base language in which the resource is written.
	Language *types.Code
	// Assertion about certainty associated with the propensity, or potential risk, of a reaction to the identified substance (including pharmaceutical product).
	VerificationStatus *types.CodeableConcept
	// Category of the identified substance.
	Category []types.Code
	// The patient who has the allergy or intolerance.
	Patient types.Reference
	// The recordedDate represents when this particular AllergyIntolerance record was created in the system, which is often a system-generated date.
	RecordedDate *types.DateTime
	// Details about each adverse reaction event linked to exposure to the identified substance.
	Reaction []AllergyIntoleranceReaction
}

func (s AllergyIntolerance) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
