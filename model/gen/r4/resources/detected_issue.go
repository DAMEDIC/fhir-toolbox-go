package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Supporting evidence or manifestations that provide the basis for identifying the detected issue such as a GuidanceResponse or MeasureReport.
type DetectedIssueEvidence struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A manifestation that led to the recording of this detected issue.
	Code []types.CodeableConcept
	// Links to resources that constitute evidence for the detected issue such as a GuidanceResponse or MeasureReport.
	Detail []types.Reference
}

func (s DetectedIssueEvidence) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates an action that has been taken or is committed to reduce or eliminate the likelihood of the risk identified by the detected issue from manifesting.  Can also reflect an observation of known mitigating factors that may reduce/eliminate the need for any action.
type DetectedIssueMitigation struct {
	// Describes the action that was taken or the observation that was made that reduces/eliminates the risk associated with the identified issue.
	Action types.CodeableConcept
	// Indicates when the mitigating action was documented.
	Date *types.DateTime
	// Identifies the practitioner who determined the mitigation and takes responsibility for the mitigation step occurring.
	Author *types.Reference
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
}

func (s DetectedIssueMitigation) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions for a patient; e.g. Drug-drug interaction, Ineffective treatment frequency, Procedure-condition conflict, etc.
type DetectedIssue struct {
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Identifies the general type of issue identified.
	Code *types.CodeableConcept
	// Indicates the resource representing the current activity or proposed activity that is potentially problematic.
	Implicated []types.Reference
	// Supporting evidence or manifestations that provide the basis for identifying the detected issue such as a GuidanceResponse or MeasureReport.
	Evidence []DetectedIssueEvidence
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// Business identifier associated with the detected issue record.
	Identifier []types.Identifier
	// Indicates the status of the detected issue.
	Status types.Code
	// Individual or device responsible for the issue being raised.  For example, a decision support application or a pharmacist conducting a medication review.
	Author *types.Reference
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Indicates the degree of importance associated with the identified issue based on the potential impact on the patient.
	Severity *types.Code
	// A textual explanation of the detected issue.
	Detail *types.String
	// Indicates the patient whose record the detected issue is associated with.
	Patient *types.Reference
	// The date or period when the detected issue was initially identified.
	Identified r4.Element
	// The literature, knowledge-base or similar reference that describes the propensity for the detected issue identified.
	Reference *types.Uri
	// Indicates an action that has been taken or is committed to reduce or eliminate the likelihood of the risk identified by the detected issue from manifesting.  Can also reflect an observation of known mitigating factors that may reduce/eliminate the need for any action.
	Mitigation []DetectedIssueMitigation
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// The base language in which the resource is written.
	Language *types.Code
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
}

func (s DetectedIssue) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
