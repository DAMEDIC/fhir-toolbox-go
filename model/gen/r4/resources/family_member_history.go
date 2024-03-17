package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// The significant Conditions (or condition) that the family member had. This is a repeating section to allow a system to represent more than one condition per resource, though there is nothing stopping multiple resources - one per condition.
type FamilyMemberHistoryCondition struct {
	// The actual condition specified. Could be a coded condition (like MI or Diabetes) or a less specific string like 'cancer' depending on how much is known about the condition and the capabilities of the creating system.
	Code types.CodeableConcept
	// Indicates what happened following the condition.  If the condition resulted in death, deceased date is captured on the relation.
	Outcome *types.CodeableConcept
	// This condition contributed to the cause of death of the related person. If contributedToDeath is not populated, then it is unknown.
	ContributedToDeath *types.Boolean
	// Either the age of onset, range of approximate age or descriptive string can be recorded.  For conditions with multiple occurrences, this describes the first known occurrence.
	Onset r4.Element
	// An area where general notes can be placed about this specific condition.
	Note []types.Annotation
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
}

func (s FamilyMemberHistoryCondition) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Significant health conditions for a person related to the patient relevant in the context of care for the patient.
type FamilyMemberHistory struct {
	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this FamilyMemberHistory.
	InstantiatesCanonical []types.Canonical
	// A code specifying the status of the record of the family history of a specific family member.
	Status types.Code
	// Describes why the family member's history is not available.
	DataAbsentReason *types.CodeableConcept
	// The person who this history concerns.
	Patient types.Reference
	// The date (and possibly time) when the family member history was recorded or last updated.
	Date *types.DateTime
	// If true, indicates that the age value specified is an estimated value.
	EstimatedAge *types.Boolean
	// Deceased flag or the actual or approximate age of the relative at the time of death for the family member history record.
	Deceased r4.Element
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// The significant Conditions (or condition) that the family member had. This is a repeating section to allow a system to represent more than one condition per resource, though there is nothing stopping multiple resources - one per condition.
	Condition []FamilyMemberHistoryCondition
	// The base language in which the resource is written.
	Language *types.Code
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Business identifiers assigned to this family member history by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier []types.Identifier
	// The birth sex of the family member.
	Sex *types.CodeableConcept
	// Describes why the family member history occurred in coded or textual form.
	ReasonCode []types.CodeableConcept
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// The type of relationship this person has to the patient (father, mother, brother etc.).
	Relationship types.CodeableConcept
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this FamilyMemberHistory.
	InstantiatesUri []types.Uri
	// This will either be a name or a description; e.g. "Aunt Susan", "my cousin with the red hair".
	Name *types.String
	// The actual or approximate date of birth of the relative.
	Born r4.Element
	// The age of the relative at the time the family member history is recorded.
	Age r4.Element
	// Indicates a Condition, Observation, AllergyIntolerance, or QuestionnaireResponse that justifies this family member history event.
	ReasonReference []types.Reference
	// This property allows a non condition-specific note to the made about the related person. Ideally, the note would be in the condition property, but this is not always possible.
	Note []types.Annotation
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
}

func (s FamilyMemberHistory) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
