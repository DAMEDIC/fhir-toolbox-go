package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// Describes validation requirements, source(s), status and dates for one or more elements.
type VerificationResult struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *Id
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *Uri
	// The base language in which the resource is written.
	Language *Code
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *Narrative
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A resource that was validated.
	Target []Reference
	// The fhirpath location(s) within the resource that was validated.
	TargetLocation []String
	// The frequency with which the target must be validated (none; initial; periodic).
	Need *CodeableConcept
	// The validation status of the target (attested; validated; in process; requires revalidation; validation failed; revalidation failed).
	Status Code
	// When the validation status was updated.
	StatusDate *DateTime
	// What the target is validated against (nothing; primary source; multiple sources).
	ValidationType *CodeableConcept
	// The primary process by which the target is validated (edit check; value set; primary source; multiple sources; standalone; in context).
	ValidationProcess []CodeableConcept
	// Frequency of revalidation.
	Frequency *Timing
	// The date/time validation was last completed (including failed validations).
	LastPerformed *DateTime
	// The date when target is next validated, if appropriate.
	NextScheduled *Date
	// The result if validation fails (fatal; warning; record only; none).
	FailureAction *CodeableConcept
	// Information about the primary source(s) involved in validation.
	PrimarySource []VerificationResultPrimarySource
	// Information about the entity attesting to information.
	Attestation *VerificationResultAttestation
	// Information about the entity validating information.
	Validator []VerificationResultValidator
}
type jsonVerificationResult struct {
	ResourceType                   string                            `json:"resourceType"`
	Id                             *Id                               `json:"id,omitempty"`
	IdPrimitiveElement             *primitiveElement                 `json:"_id,omitempty"`
	Meta                           *Meta                             `json:"meta,omitempty"`
	ImplicitRules                  *Uri                              `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement  *primitiveElement                 `json:"_implicitRules,omitempty"`
	Language                       *Code                             `json:"language,omitempty"`
	LanguagePrimitiveElement       *primitiveElement                 `json:"_language,omitempty"`
	Text                           *Narrative                        `json:"text,omitempty"`
	Contained                      []containedResource               `json:"contained,omitempty"`
	Extension                      []Extension                       `json:"extension,omitempty"`
	ModifierExtension              []Extension                       `json:"modifierExtension,omitempty"`
	Target                         []Reference                       `json:"target,omitempty"`
	TargetLocation                 []String                          `json:"targetLocation,omitempty"`
	TargetLocationPrimitiveElement []*primitiveElement               `json:"_targetLocation,omitempty"`
	Need                           *CodeableConcept                  `json:"need,omitempty"`
	Status                         Code                              `json:"status,omitempty"`
	StatusPrimitiveElement         *primitiveElement                 `json:"_status,omitempty"`
	StatusDate                     *DateTime                         `json:"statusDate,omitempty"`
	StatusDatePrimitiveElement     *primitiveElement                 `json:"_statusDate,omitempty"`
	ValidationType                 *CodeableConcept                  `json:"validationType,omitempty"`
	ValidationProcess              []CodeableConcept                 `json:"validationProcess,omitempty"`
	Frequency                      *Timing                           `json:"frequency,omitempty"`
	LastPerformed                  *DateTime                         `json:"lastPerformed,omitempty"`
	LastPerformedPrimitiveElement  *primitiveElement                 `json:"_lastPerformed,omitempty"`
	NextScheduled                  *Date                             `json:"nextScheduled,omitempty"`
	NextScheduledPrimitiveElement  *primitiveElement                 `json:"_nextScheduled,omitempty"`
	FailureAction                  *CodeableConcept                  `json:"failureAction,omitempty"`
	PrimarySource                  []VerificationResultPrimarySource `json:"primarySource,omitempty"`
	Attestation                    *VerificationResultAttestation    `json:"attestation,omitempty"`
	Validator                      []VerificationResultValidator     `json:"validator,omitempty"`
}

func (r VerificationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r VerificationResult) marshalJSON() jsonVerificationResult {
	m := jsonVerificationResult{}
	m.ResourceType = "VerificationResult"
	m.Id = r.Id
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	m.ImplicitRules = r.ImplicitRules
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	m.Language = r.Language
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.Text = r.Text
	m.Contained = make([]containedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, containedResource{resource: c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Target = r.Target
	m.TargetLocation = r.TargetLocation
	anyTargetLocationIdOrExtension := false
	for _, e := range r.TargetLocation {
		if e.Id != nil || e.Extension != nil {
			anyTargetLocationIdOrExtension = true
			break
		}
	}
	if anyTargetLocationIdOrExtension {
		m.TargetLocationPrimitiveElement = make([]*primitiveElement, 0, len(r.TargetLocation))
		for _, e := range r.TargetLocation {
			if e.Id != nil || e.Extension != nil {
				m.TargetLocationPrimitiveElement = append(m.TargetLocationPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.TargetLocationPrimitiveElement = append(m.TargetLocationPrimitiveElement, nil)
			}
		}
	}
	m.Need = r.Need
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.StatusDate = r.StatusDate
	if r.StatusDate != nil && (r.StatusDate.Id != nil || r.StatusDate.Extension != nil) {
		m.StatusDatePrimitiveElement = &primitiveElement{Id: r.StatusDate.Id, Extension: r.StatusDate.Extension}
	}
	m.ValidationType = r.ValidationType
	m.ValidationProcess = r.ValidationProcess
	m.Frequency = r.Frequency
	m.LastPerformed = r.LastPerformed
	if r.LastPerformed != nil && (r.LastPerformed.Id != nil || r.LastPerformed.Extension != nil) {
		m.LastPerformedPrimitiveElement = &primitiveElement{Id: r.LastPerformed.Id, Extension: r.LastPerformed.Extension}
	}
	m.NextScheduled = r.NextScheduled
	if r.NextScheduled != nil && (r.NextScheduled.Id != nil || r.NextScheduled.Extension != nil) {
		m.NextScheduledPrimitiveElement = &primitiveElement{Id: r.NextScheduled.Id, Extension: r.NextScheduled.Extension}
	}
	m.FailureAction = r.FailureAction
	m.PrimarySource = r.PrimarySource
	m.Attestation = r.Attestation
	m.Validator = r.Validator
	return m
}
func (r *VerificationResult) UnmarshalJSON(b []byte) error {
	var m jsonVerificationResult
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *VerificationResult) unmarshalJSON(m jsonVerificationResult) error {
	r.Id = m.Id
	if m.IdPrimitiveElement != nil {
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Text = m.Text
	r.Contained = make([]model.Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Target = m.Target
	r.TargetLocation = m.TargetLocation
	for i, e := range m.TargetLocationPrimitiveElement {
		if len(r.TargetLocation) > i {
			r.TargetLocation[i].Id = e.Id
			r.TargetLocation[i].Extension = e.Extension
		} else {
			r.TargetLocation = append(r.TargetLocation, String{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Need = m.Need
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.StatusDate = m.StatusDate
	if m.StatusDatePrimitiveElement != nil {
		r.StatusDate.Id = m.StatusDatePrimitiveElement.Id
		r.StatusDate.Extension = m.StatusDatePrimitiveElement.Extension
	}
	r.ValidationType = m.ValidationType
	r.ValidationProcess = m.ValidationProcess
	r.Frequency = m.Frequency
	r.LastPerformed = m.LastPerformed
	if m.LastPerformedPrimitiveElement != nil {
		r.LastPerformed.Id = m.LastPerformedPrimitiveElement.Id
		r.LastPerformed.Extension = m.LastPerformedPrimitiveElement.Extension
	}
	r.NextScheduled = m.NextScheduled
	if m.NextScheduledPrimitiveElement != nil {
		r.NextScheduled.Id = m.NextScheduledPrimitiveElement.Id
		r.NextScheduled.Extension = m.NextScheduledPrimitiveElement.Extension
	}
	r.FailureAction = m.FailureAction
	r.PrimarySource = m.PrimarySource
	r.Attestation = m.Attestation
	r.Validator = m.Validator
	return nil
}
func (r VerificationResult) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about the primary source(s) involved in validation.
type VerificationResultPrimarySource struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Reference to the primary source.
	Who *Reference
	// Type of primary source (License Board; Primary Education; Continuing Education; Postal Service; Relationship owner; Registration Authority; legal source; issuing source; authoritative source).
	Type []CodeableConcept
	// Method for communicating with the primary source (manual; API; Push).
	CommunicationMethod []CodeableConcept
	// Status of the validation of the target against the primary source (successful; failed; unknown).
	ValidationStatus *CodeableConcept
	// When the target was validated against the primary source.
	ValidationDate *DateTime
	// Ability of the primary source to push updates/alerts (yes; no; undetermined).
	CanPushUpdates *CodeableConcept
	// Type of alerts/updates the primary source can send (specific requested changes; any changes; as defined by source).
	PushTypeAvailable []CodeableConcept
}
type jsonVerificationResultPrimarySource struct {
	Id                             *string           `json:"id,omitempty"`
	Extension                      []Extension       `json:"extension,omitempty"`
	ModifierExtension              []Extension       `json:"modifierExtension,omitempty"`
	Who                            *Reference        `json:"who,omitempty"`
	Type                           []CodeableConcept `json:"type,omitempty"`
	CommunicationMethod            []CodeableConcept `json:"communicationMethod,omitempty"`
	ValidationStatus               *CodeableConcept  `json:"validationStatus,omitempty"`
	ValidationDate                 *DateTime         `json:"validationDate,omitempty"`
	ValidationDatePrimitiveElement *primitiveElement `json:"_validationDate,omitempty"`
	CanPushUpdates                 *CodeableConcept  `json:"canPushUpdates,omitempty"`
	PushTypeAvailable              []CodeableConcept `json:"pushTypeAvailable,omitempty"`
}

func (r VerificationResultPrimarySource) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r VerificationResultPrimarySource) marshalJSON() jsonVerificationResultPrimarySource {
	m := jsonVerificationResultPrimarySource{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Who = r.Who
	m.Type = r.Type
	m.CommunicationMethod = r.CommunicationMethod
	m.ValidationStatus = r.ValidationStatus
	m.ValidationDate = r.ValidationDate
	if r.ValidationDate != nil && (r.ValidationDate.Id != nil || r.ValidationDate.Extension != nil) {
		m.ValidationDatePrimitiveElement = &primitiveElement{Id: r.ValidationDate.Id, Extension: r.ValidationDate.Extension}
	}
	m.CanPushUpdates = r.CanPushUpdates
	m.PushTypeAvailable = r.PushTypeAvailable
	return m
}
func (r *VerificationResultPrimarySource) UnmarshalJSON(b []byte) error {
	var m jsonVerificationResultPrimarySource
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *VerificationResultPrimarySource) unmarshalJSON(m jsonVerificationResultPrimarySource) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Who = m.Who
	r.Type = m.Type
	r.CommunicationMethod = m.CommunicationMethod
	r.ValidationStatus = m.ValidationStatus
	r.ValidationDate = m.ValidationDate
	if m.ValidationDatePrimitiveElement != nil {
		r.ValidationDate.Id = m.ValidationDatePrimitiveElement.Id
		r.ValidationDate.Extension = m.ValidationDatePrimitiveElement.Extension
	}
	r.CanPushUpdates = m.CanPushUpdates
	r.PushTypeAvailable = m.PushTypeAvailable
	return nil
}
func (r VerificationResultPrimarySource) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about the entity attesting to information.
type VerificationResultAttestation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The individual or organization attesting to information.
	Who *Reference
	// When the who is asserting on behalf of another (organization or individual).
	OnBehalfOf *Reference
	// The method by which attested information was submitted/retrieved (manual; API; Push).
	CommunicationMethod *CodeableConcept
	// The date the information was attested to.
	Date *Date
	// A digital identity certificate associated with the attestation source.
	SourceIdentityCertificate *String
	// A digital identity certificate associated with the proxy entity submitting attested information on behalf of the attestation source.
	ProxyIdentityCertificate *String
	// Signed assertion by the proxy entity indicating that they have the right to submit attested information on behalf of the attestation source.
	ProxySignature *Signature
	// Signed assertion by the attestation source that they have attested to the information.
	SourceSignature *Signature
}
type jsonVerificationResultAttestation struct {
	Id                                        *string           `json:"id,omitempty"`
	Extension                                 []Extension       `json:"extension,omitempty"`
	ModifierExtension                         []Extension       `json:"modifierExtension,omitempty"`
	Who                                       *Reference        `json:"who,omitempty"`
	OnBehalfOf                                *Reference        `json:"onBehalfOf,omitempty"`
	CommunicationMethod                       *CodeableConcept  `json:"communicationMethod,omitempty"`
	Date                                      *Date             `json:"date,omitempty"`
	DatePrimitiveElement                      *primitiveElement `json:"_date,omitempty"`
	SourceIdentityCertificate                 *String           `json:"sourceIdentityCertificate,omitempty"`
	SourceIdentityCertificatePrimitiveElement *primitiveElement `json:"_sourceIdentityCertificate,omitempty"`
	ProxyIdentityCertificate                  *String           `json:"proxyIdentityCertificate,omitempty"`
	ProxyIdentityCertificatePrimitiveElement  *primitiveElement `json:"_proxyIdentityCertificate,omitempty"`
	ProxySignature                            *Signature        `json:"proxySignature,omitempty"`
	SourceSignature                           *Signature        `json:"sourceSignature,omitempty"`
}

func (r VerificationResultAttestation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r VerificationResultAttestation) marshalJSON() jsonVerificationResultAttestation {
	m := jsonVerificationResultAttestation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Who = r.Who
	m.OnBehalfOf = r.OnBehalfOf
	m.CommunicationMethod = r.CommunicationMethod
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.SourceIdentityCertificate = r.SourceIdentityCertificate
	if r.SourceIdentityCertificate != nil && (r.SourceIdentityCertificate.Id != nil || r.SourceIdentityCertificate.Extension != nil) {
		m.SourceIdentityCertificatePrimitiveElement = &primitiveElement{Id: r.SourceIdentityCertificate.Id, Extension: r.SourceIdentityCertificate.Extension}
	}
	m.ProxyIdentityCertificate = r.ProxyIdentityCertificate
	if r.ProxyIdentityCertificate != nil && (r.ProxyIdentityCertificate.Id != nil || r.ProxyIdentityCertificate.Extension != nil) {
		m.ProxyIdentityCertificatePrimitiveElement = &primitiveElement{Id: r.ProxyIdentityCertificate.Id, Extension: r.ProxyIdentityCertificate.Extension}
	}
	m.ProxySignature = r.ProxySignature
	m.SourceSignature = r.SourceSignature
	return m
}
func (r *VerificationResultAttestation) UnmarshalJSON(b []byte) error {
	var m jsonVerificationResultAttestation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *VerificationResultAttestation) unmarshalJSON(m jsonVerificationResultAttestation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Who = m.Who
	r.OnBehalfOf = m.OnBehalfOf
	r.CommunicationMethod = m.CommunicationMethod
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.SourceIdentityCertificate = m.SourceIdentityCertificate
	if m.SourceIdentityCertificatePrimitiveElement != nil {
		r.SourceIdentityCertificate.Id = m.SourceIdentityCertificatePrimitiveElement.Id
		r.SourceIdentityCertificate.Extension = m.SourceIdentityCertificatePrimitiveElement.Extension
	}
	r.ProxyIdentityCertificate = m.ProxyIdentityCertificate
	if m.ProxyIdentityCertificatePrimitiveElement != nil {
		r.ProxyIdentityCertificate.Id = m.ProxyIdentityCertificatePrimitiveElement.Id
		r.ProxyIdentityCertificate.Extension = m.ProxyIdentityCertificatePrimitiveElement.Extension
	}
	r.ProxySignature = m.ProxySignature
	r.SourceSignature = m.SourceSignature
	return nil
}
func (r VerificationResultAttestation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
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
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Reference to the organization validating information.
	Organization Reference
	// A digital identity certificate associated with the validator.
	IdentityCertificate *String
	// Signed assertion by the validator that they have validated the information.
	AttestationSignature *Signature
}
type jsonVerificationResultValidator struct {
	Id                                  *string           `json:"id,omitempty"`
	Extension                           []Extension       `json:"extension,omitempty"`
	ModifierExtension                   []Extension       `json:"modifierExtension,omitempty"`
	Organization                        Reference         `json:"organization,omitempty"`
	IdentityCertificate                 *String           `json:"identityCertificate,omitempty"`
	IdentityCertificatePrimitiveElement *primitiveElement `json:"_identityCertificate,omitempty"`
	AttestationSignature                *Signature        `json:"attestationSignature,omitempty"`
}

func (r VerificationResultValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r VerificationResultValidator) marshalJSON() jsonVerificationResultValidator {
	m := jsonVerificationResultValidator{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Organization = r.Organization
	m.IdentityCertificate = r.IdentityCertificate
	if r.IdentityCertificate != nil && (r.IdentityCertificate.Id != nil || r.IdentityCertificate.Extension != nil) {
		m.IdentityCertificatePrimitiveElement = &primitiveElement{Id: r.IdentityCertificate.Id, Extension: r.IdentityCertificate.Extension}
	}
	m.AttestationSignature = r.AttestationSignature
	return m
}
func (r *VerificationResultValidator) UnmarshalJSON(b []byte) error {
	var m jsonVerificationResultValidator
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *VerificationResultValidator) unmarshalJSON(m jsonVerificationResultValidator) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Organization = m.Organization
	r.IdentityCertificate = m.IdentityCertificate
	if m.IdentityCertificatePrimitiveElement != nil {
		r.IdentityCertificate.Id = m.IdentityCertificatePrimitiveElement.Id
		r.IdentityCertificate.Extension = m.IdentityCertificatePrimitiveElement.Extension
	}
	r.AttestationSignature = m.AttestationSignature
	return nil
}
func (r VerificationResultValidator) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
