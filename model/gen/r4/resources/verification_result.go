package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Information about the primary source(s) involved in validation.
type VerificationResultPrimarySource struct {
	// Reference to the primary source.
	Who *types.Reference
	// Method for communicating with the primary source (manual; API; Push).
	CommunicationMethod []types.CodeableConcept
	// When the target was validated against the primary source.
	ValidationDate *types.DateTime
	// Ability of the primary source to push updates/alerts (yes; no; undetermined).
	CanPushUpdates *types.CodeableConcept
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Status of the validation of the target against the primary source (successful; failed; unknown).
	ValidationStatus *types.CodeableConcept
	// Type of alerts/updates the primary source can send (specific requested changes; any changes; as defined by source).
	PushTypeAvailable []types.CodeableConcept
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Type of primary source (License Board; Primary Education; Continuing Education; Postal Service; Relationship owner; Registration Authority; legal source; issuing source; authoritative source).
	Type []types.CodeableConcept
}

func (s VerificationResultPrimarySource) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about the entity attesting to information.
type VerificationResultAttestation struct {
	// When the who is asserting on behalf of another (organization or individual).
	OnBehalfOf *types.Reference
	// A digital identity certificate associated with the proxy entity submitting attested information on behalf of the attestation source.
	ProxyIdentityCertificate *types.String
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A digital identity certificate associated with the attestation source.
	SourceIdentityCertificate *types.String
	// Signed assertion by the proxy entity indicating that they have the right to submit attested information on behalf of the attestation source.
	ProxySignature *types.Signature
	// Signed assertion by the attestation source that they have attested to the information.
	SourceSignature *types.Signature
	// The individual or organization attesting to information.
	Who *types.Reference
	// The method by which attested information was submitted/retrieved (manual; API; Push).
	CommunicationMethod *types.CodeableConcept
	// The date the information was attested to.
	Date *types.Date
}

func (s VerificationResultAttestation) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about the entity validating information.
type VerificationResultValidator struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Reference to the organization validating information.
	Organization types.Reference
	// A digital identity certificate associated with the validator.
	IdentityCertificate *types.String
	// Signed assertion by the validator that they have validated the information.
	AttestationSignature *types.Signature
}

func (s VerificationResultValidator) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Describes validation requirements, source(s), status and dates for one or more elements.
type VerificationResult struct {
	// The date/time validation was last completed (including failed validations).
	LastPerformed *types.DateTime
	// Information about the primary source(s) involved in validation.
	PrimarySource []VerificationResultPrimarySource
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// The date when target is next validated, if appropriate.
	NextScheduled *types.Date
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A resource that was validated.
	Target []types.Reference
	// The fhirpath location(s) within the resource that was validated.
	TargetLocation []types.String
	// The frequency with which the target must be validated (none; initial; periodic).
	Need *types.CodeableConcept
	// When the validation status was updated.
	StatusDate *types.DateTime
	// What the target is validated against (nothing; primary source; multiple sources).
	ValidationType *types.CodeableConcept
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// The base language in which the resource is written.
	Language *types.Code
	// Information about the entity attesting to information.
	Attestation *VerificationResultAttestation
	// The validation status of the target (attested; validated; in process; requires revalidation; validation failed; revalidation failed).
	Status types.Code
	// The primary process by which the target is validated (edit check; value set; primary source; multiple sources; standalone; in context).
	ValidationProcess []types.CodeableConcept
	// Frequency of revalidation.
	Frequency *types.Timing
	// The result if validation fails (fatal; warning; record only; none).
	FailureAction *types.CodeableConcept
	// Information about the entity validating information.
	Validator []VerificationResultValidator
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
}

func (s VerificationResult) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
