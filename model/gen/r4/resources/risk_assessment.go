package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Describes the expected outcome for the subject.
type RiskAssessmentPrediction struct {
	// Indicates the period of time or age range of the subject to which the specified probability applies.
	When r4.Element
	// Additional information explaining the basis for the prediction.
	Rationale *types.String
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Indicates how likely the outcome is (in the specified timeframe).
	Probability r4.Element
	// Indicates how likely the outcome is (in the specified timeframe), expressed as a qualitative value (e.g. low, medium, or high).
	QualitativeRisk *types.CodeableConcept
	// Indicates the risk for this particular subject (with their specific characteristics) divided by the risk of the population in general.  (Numbers greater than 1 = higher risk than the population, numbers less than 1 = lower risk.).
	RelativeRisk *types.Decimal
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// One of the potential outcomes for the patient (e.g. remission, death,  a particular condition).
	Outcome *types.CodeableConcept
}

func (s RiskAssessmentPrediction) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// An assessment of the likely outcome(s) for a patient or other subject as well as the likelihood of each outcome.
type RiskAssessment struct {
	// A description of the steps that might be taken to reduce the identified risk(s).
	Mitigation *types.String
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// The base language in which the resource is written.
	Language *types.Code
	// The algorithm, process or mechanism used to evaluate the risk.
	Method *types.CodeableConcept
	// The reason the risk assessment was performed.
	ReasonCode []types.CodeableConcept
	// A reference to the request that is fulfilled by this risk assessment.
	BasedOn *types.Reference
	// The status of the RiskAssessment, using the same statuses as an Observation.
	Status types.Code
	// The patient or group the risk assessment applies to.
	Subject types.Reference
	// Resources supporting the reason the risk assessment was performed.
	ReasonReference []types.Reference
	// The date (and possibly time) the risk assessment was performed.
	Occurrence r4.Element
	// Indicates the source data considered as part of the assessment (for example, FamilyHistory, Observations, Procedures, Conditions, etc.).
	Basis []types.Reference
	// Describes the expected outcome for the subject.
	Prediction []RiskAssessmentPrediction
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// A reference to a resource that this risk assessment is part of, such as a Procedure.
	Parent *types.Reference
	// The type of the risk assessment performed.
	Code *types.CodeableConcept
	// The encounter where the assessment was performed.
	Encounter *types.Reference
	// For assessments or prognosis specific to a particular condition, indicates the condition being assessed.
	Condition *types.Reference
	// The provider or software application that performed the assessment.
	Performer *types.Reference
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Business identifier assigned to the risk assessment.
	Identifier []types.Identifier
	// Additional comments about the risk assessment.
	Note []types.Annotation
}

func (s RiskAssessment) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
