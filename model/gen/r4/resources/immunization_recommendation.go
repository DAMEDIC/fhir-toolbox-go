package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Vaccine date recommendations.  For example, earliest date to administer, latest date to administer, etc.
type ImmunizationRecommendationRecommendationDateCriterion struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Date classification of recommendation.  For example, earliest date to give, latest date to give, etc.
	Code types.CodeableConcept
	// The date whose meaning is specified by dateCriterion.code.
	Value types.DateTime
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
}

func (s ImmunizationRecommendationRecommendationDateCriterion) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Vaccine administration recommendations.
type ImmunizationRecommendationRecommendation struct {
	// The reason for the assigned forecast status.
	ForecastReason []types.CodeableConcept
	// One possible path to achieve presumed immunity against a disease - within the context of an authority.
	Series *types.String
	// Patient Information that supports the status and recommendation.  This includes patient observations, adverse reactions and allergy/intolerance information.
	SupportingPatientInformation []types.Reference
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Vaccine(s) or vaccine group that pertain to the recommendation.
	VaccineCode []types.CodeableConcept
	// Indicates the patient status with respect to the path to immunity for the target disease.
	ForecastStatus types.CodeableConcept
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Vaccine(s) which should not be used to fulfill the recommendation.
	ContraindicatedVaccineCode []types.CodeableConcept
	// Immunization event history and/or evaluation that supports the status and recommendation.
	SupportingImmunization []types.Reference
	// The targeted disease for the recommendation.
	TargetDisease *types.CodeableConcept
	// Contains the description about the protocol under which the vaccine was administered.
	Description *types.String
	// The recommended number of doses to achieve immunity.
	SeriesDoses r4.Element
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Vaccine date recommendations.  For example, earliest date to administer, latest date to administer, etc.
	DateCriterion []ImmunizationRecommendationRecommendationDateCriterion
	// Nominal position of the recommended dose in a series (e.g. dose 2 is the next recommended dose).
	DoseNumber r4.Element
}

func (s ImmunizationRecommendationRecommendation) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A patient's point-in-time set of recommendations (i.e. forecasting) according to a published schedule with optional supporting justification.
type ImmunizationRecommendation struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// The base language in which the resource is written.
	Language *types.Code
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The patient the recommendation(s) are for.
	Patient types.Reference
	// Indicates the authority who published the protocol (e.g. ACIP).
	Authority *types.Reference
	// Vaccine administration recommendations.
	Recommendation []ImmunizationRecommendationRecommendation
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A unique identifier assigned to this particular recommendation record.
	Identifier []types.Identifier
	// The date the immunization recommendation(s) were created.
	Date types.DateTime
}

func (s ImmunizationRecommendation) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
