package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A record of a healthcare consumer’s  choices, which permits or denies identified recipient(s) or recipient role(s) to perform one or more actions within a given policy context, for specific purposes and periods of time.
type Consent struct {
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
	// Unique identifier for this copy of the Consent Statement.
	Identifier []Identifier
	// Indicates the current state of this consent.
	Status Code
	// A selector of the type of consent being presented: ADR, Privacy, Treatment, Research.  This list is now extensible.
	Scope CodeableConcept
	// A classification of the type of consents found in the statement. This element supports indexing and retrieval of consent statements.
	Category []CodeableConcept
	// The patient/healthcare consumer to whom this consent applies.
	Patient *Reference
	// When this  Consent was issued / created / indexed.
	DateTime *DateTime
	// Either the Grantor, which is the entity responsible for granting the rights listed in a Consent Directive or the Grantee, which is the entity responsible for complying with the Consent Directive, including any obligations or limitations on authorizations and enforcement of prohibitions.
	Performer []Reference
	// The organization that manages the consent, and the framework within which it is executed.
	Organization []Reference
	// The source on which this consent statement is based. The source might be a scanned original paper form, or a reference to a consent that links back to such a source, a reference to a document repository (e.g. XDS) that stores the original consent document.
	Source isConsentSource
	// The references to the policies that are included in this consent scope. Policies may be organizational, but are often defined jurisdictionally, or in law.
	Policy []ConsentPolicy
	// A reference to the specific base computable regulation or policy.
	PolicyRule *CodeableConcept
	// Whether a treatment instruction (e.g. artificial respiration yes or no) was verified with the patient, his/her family or another authorized person.
	Verification []ConsentVerification
	// An exception to the base policy of this consent. An exception can be an addition or removal of access permissions.
	Provision *ConsentProvision
}

func (r Consent) ResourceType() string {
	return "Consent"
}

type isConsentSource interface {
	isConsentSource()
}

func (r Attachment) isConsentSource() {}
func (r Reference) isConsentSource()  {}

type jsonConsent struct {
	ResourceType                  string                `json:"resourceType"`
	Id                            *Id                   `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement     `json:"_id,omitempty"`
	Meta                          *Meta                 `json:"meta,omitempty"`
	ImplicitRules                 *Uri                  `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement     `json:"_implicitRules,omitempty"`
	Language                      *Code                 `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement     `json:"_language,omitempty"`
	Text                          *Narrative            `json:"text,omitempty"`
	Contained                     []containedResource   `json:"contained,omitempty"`
	Extension                     []Extension           `json:"extension,omitempty"`
	ModifierExtension             []Extension           `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier          `json:"identifier,omitempty"`
	Status                        Code                  `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement     `json:"_status,omitempty"`
	Scope                         CodeableConcept       `json:"scope,omitempty"`
	Category                      []CodeableConcept     `json:"category,omitempty"`
	Patient                       *Reference            `json:"patient,omitempty"`
	DateTime                      *DateTime             `json:"dateTime,omitempty"`
	DateTimePrimitiveElement      *primitiveElement     `json:"_dateTime,omitempty"`
	Performer                     []Reference           `json:"performer,omitempty"`
	Organization                  []Reference           `json:"organization,omitempty"`
	SourceAttachment              *Attachment           `json:"sourceAttachment,omitempty"`
	SourceReference               *Reference            `json:"sourceReference,omitempty"`
	Policy                        []ConsentPolicy       `json:"policy,omitempty"`
	PolicyRule                    *CodeableConcept      `json:"policyRule,omitempty"`
	Verification                  []ConsentVerification `json:"verification,omitempty"`
	Provision                     *ConsentProvision     `json:"provision,omitempty"`
}

func (r Consent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Consent) marshalJSON() jsonConsent {
	m := jsonConsent{}
	m.ResourceType = "Consent"
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
	m.Identifier = r.Identifier
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Scope = r.Scope
	m.Category = r.Category
	m.Patient = r.Patient
	m.DateTime = r.DateTime
	if r.DateTime != nil && (r.DateTime.Id != nil || r.DateTime.Extension != nil) {
		m.DateTimePrimitiveElement = &primitiveElement{Id: r.DateTime.Id, Extension: r.DateTime.Extension}
	}
	m.Performer = r.Performer
	m.Organization = r.Organization
	switch v := r.Source.(type) {
	case Attachment:
		m.SourceAttachment = &v
	case *Attachment:
		m.SourceAttachment = v
	case Reference:
		m.SourceReference = &v
	case *Reference:
		m.SourceReference = v
	}
	m.Policy = r.Policy
	m.PolicyRule = r.PolicyRule
	m.Verification = r.Verification
	m.Provision = r.Provision
	return m
}
func (r *Consent) UnmarshalJSON(b []byte) error {
	var m jsonConsent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Consent) unmarshalJSON(m jsonConsent) error {
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
	r.Identifier = m.Identifier
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Scope = m.Scope
	r.Category = m.Category
	r.Patient = m.Patient
	r.DateTime = m.DateTime
	if m.DateTimePrimitiveElement != nil {
		r.DateTime.Id = m.DateTimePrimitiveElement.Id
		r.DateTime.Extension = m.DateTimePrimitiveElement.Extension
	}
	r.Performer = m.Performer
	r.Organization = m.Organization
	if m.SourceAttachment != nil {
		if r.Source != nil {
			return fmt.Errorf("multiple values for field \"Source\"")
		}
		v := m.SourceAttachment
		r.Source = v
	}
	if m.SourceReference != nil {
		if r.Source != nil {
			return fmt.Errorf("multiple values for field \"Source\"")
		}
		v := m.SourceReference
		r.Source = v
	}
	r.Policy = m.Policy
	r.PolicyRule = m.PolicyRule
	r.Verification = m.Verification
	r.Provision = m.Provision
	return nil
}
func (r Consent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The references to the policies that are included in this consent scope. Policies may be organizational, but are often defined jurisdictionally, or in law.
type ConsentPolicy struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Entity or Organization having regulatory jurisdiction or accountability for  enforcing policies pertaining to Consent Directives.
	Authority *Uri
	// The references to the policies that are included in this consent scope. Policies may be organizational, but are often defined jurisdictionally, or in law.
	Uri *Uri
}
type jsonConsentPolicy struct {
	Id                        *string           `json:"id,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Authority                 *Uri              `json:"authority,omitempty"`
	AuthorityPrimitiveElement *primitiveElement `json:"_authority,omitempty"`
	Uri                       *Uri              `json:"uri,omitempty"`
	UriPrimitiveElement       *primitiveElement `json:"_uri,omitempty"`
}

func (r ConsentPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ConsentPolicy) marshalJSON() jsonConsentPolicy {
	m := jsonConsentPolicy{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Authority = r.Authority
	if r.Authority != nil && (r.Authority.Id != nil || r.Authority.Extension != nil) {
		m.AuthorityPrimitiveElement = &primitiveElement{Id: r.Authority.Id, Extension: r.Authority.Extension}
	}
	m.Uri = r.Uri
	if r.Uri != nil && (r.Uri.Id != nil || r.Uri.Extension != nil) {
		m.UriPrimitiveElement = &primitiveElement{Id: r.Uri.Id, Extension: r.Uri.Extension}
	}
	return m
}
func (r *ConsentPolicy) UnmarshalJSON(b []byte) error {
	var m jsonConsentPolicy
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ConsentPolicy) unmarshalJSON(m jsonConsentPolicy) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Authority = m.Authority
	if m.AuthorityPrimitiveElement != nil {
		r.Authority.Id = m.AuthorityPrimitiveElement.Id
		r.Authority.Extension = m.AuthorityPrimitiveElement.Extension
	}
	r.Uri = m.Uri
	if m.UriPrimitiveElement != nil {
		r.Uri.Id = m.UriPrimitiveElement.Id
		r.Uri.Extension = m.UriPrimitiveElement.Extension
	}
	return nil
}
func (r ConsentPolicy) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Whether a treatment instruction (e.g. artificial respiration yes or no) was verified with the patient, his/her family or another authorized person.
type ConsentVerification struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Has the instruction been verified.
	Verified Boolean
	// Who verified the instruction (Patient, Relative or other Authorized Person).
	VerifiedWith *Reference
	// Date verification was collected.
	VerificationDate *DateTime
}
type jsonConsentVerification struct {
	Id                               *string           `json:"id,omitempty"`
	Extension                        []Extension       `json:"extension,omitempty"`
	ModifierExtension                []Extension       `json:"modifierExtension,omitempty"`
	Verified                         Boolean           `json:"verified,omitempty"`
	VerifiedPrimitiveElement         *primitiveElement `json:"_verified,omitempty"`
	VerifiedWith                     *Reference        `json:"verifiedWith,omitempty"`
	VerificationDate                 *DateTime         `json:"verificationDate,omitempty"`
	VerificationDatePrimitiveElement *primitiveElement `json:"_verificationDate,omitempty"`
}

func (r ConsentVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ConsentVerification) marshalJSON() jsonConsentVerification {
	m := jsonConsentVerification{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Verified = r.Verified
	if r.Verified.Id != nil || r.Verified.Extension != nil {
		m.VerifiedPrimitiveElement = &primitiveElement{Id: r.Verified.Id, Extension: r.Verified.Extension}
	}
	m.VerifiedWith = r.VerifiedWith
	m.VerificationDate = r.VerificationDate
	if r.VerificationDate != nil && (r.VerificationDate.Id != nil || r.VerificationDate.Extension != nil) {
		m.VerificationDatePrimitiveElement = &primitiveElement{Id: r.VerificationDate.Id, Extension: r.VerificationDate.Extension}
	}
	return m
}
func (r *ConsentVerification) UnmarshalJSON(b []byte) error {
	var m jsonConsentVerification
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ConsentVerification) unmarshalJSON(m jsonConsentVerification) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Verified = m.Verified
	if m.VerifiedPrimitiveElement != nil {
		r.Verified.Id = m.VerifiedPrimitiveElement.Id
		r.Verified.Extension = m.VerifiedPrimitiveElement.Extension
	}
	r.VerifiedWith = m.VerifiedWith
	r.VerificationDate = m.VerificationDate
	if m.VerificationDatePrimitiveElement != nil {
		r.VerificationDate.Id = m.VerificationDatePrimitiveElement.Id
		r.VerificationDate.Extension = m.VerificationDatePrimitiveElement.Extension
	}
	return nil
}
func (r ConsentVerification) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// An exception to the base policy of this consent. An exception can be an addition or removal of access permissions.
type ConsentProvision struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Action  to take - permit or deny - when the rule conditions are met.  Not permitted in root rule, required in all nested rules.
	Type *Code
	// The timeframe in this rule is valid.
	Period *Period
	// Who or what is controlled by this rule. Use group to identify a set of actors by some property they share (e.g. 'admitting officers').
	Actor []ConsentProvisionActor
	// Actions controlled by this Rule.
	Action []CodeableConcept
	// A security label, comprised of 0..* security label fields (Privacy tags), which define which resources are controlled by this exception.
	SecurityLabel []Coding
	// The context of the activities a user is taking - why the user is accessing the data - that are controlled by this rule.
	Purpose []Coding
	// The class of information covered by this rule. The type can be a FHIR resource type, a profile on a type, or a CDA document, or some other type that indicates what sort of information the consent relates to.
	Class []Coding
	// If this code is found in an instance, then the rule applies.
	Code []CodeableConcept
	// Clinical or Operational Relevant period of time that bounds the data controlled by this rule.
	DataPeriod *Period
	// The resources controlled by this rule if specific resources are referenced.
	Data []ConsentProvisionData
	// Rules which provide exceptions to the base rule or subrules.
	Provision []ConsentProvision
}
type jsonConsentProvision struct {
	Id                   *string                 `json:"id,omitempty"`
	Extension            []Extension             `json:"extension,omitempty"`
	ModifierExtension    []Extension             `json:"modifierExtension,omitempty"`
	Type                 *Code                   `json:"type,omitempty"`
	TypePrimitiveElement *primitiveElement       `json:"_type,omitempty"`
	Period               *Period                 `json:"period,omitempty"`
	Actor                []ConsentProvisionActor `json:"actor,omitempty"`
	Action               []CodeableConcept       `json:"action,omitempty"`
	SecurityLabel        []Coding                `json:"securityLabel,omitempty"`
	Purpose              []Coding                `json:"purpose,omitempty"`
	Class                []Coding                `json:"class,omitempty"`
	Code                 []CodeableConcept       `json:"code,omitempty"`
	DataPeriod           *Period                 `json:"dataPeriod,omitempty"`
	Data                 []ConsentProvisionData  `json:"data,omitempty"`
	Provision            []ConsentProvision      `json:"provision,omitempty"`
}

func (r ConsentProvision) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ConsentProvision) marshalJSON() jsonConsentProvision {
	m := jsonConsentProvision{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Period = r.Period
	m.Actor = r.Actor
	m.Action = r.Action
	m.SecurityLabel = r.SecurityLabel
	m.Purpose = r.Purpose
	m.Class = r.Class
	m.Code = r.Code
	m.DataPeriod = r.DataPeriod
	m.Data = r.Data
	m.Provision = r.Provision
	return m
}
func (r *ConsentProvision) UnmarshalJSON(b []byte) error {
	var m jsonConsentProvision
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ConsentProvision) unmarshalJSON(m jsonConsentProvision) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Period = m.Period
	r.Actor = m.Actor
	r.Action = m.Action
	r.SecurityLabel = m.SecurityLabel
	r.Purpose = m.Purpose
	r.Class = m.Class
	r.Code = m.Code
	r.DataPeriod = m.DataPeriod
	r.Data = m.Data
	r.Provision = m.Provision
	return nil
}
func (r ConsentProvision) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Who or what is controlled by this rule. Use group to identify a set of actors by some property they share (e.g. 'admitting officers').
type ConsentProvisionActor struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// How the individual is involved in the resources content that is described in the exception.
	Role CodeableConcept
	// The resource that identifies the actor. To identify actors by type, use group to identify a set of actors by some property they share (e.g. 'admitting officers').
	Reference Reference
}
type jsonConsentProvisionActor struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Role              CodeableConcept `json:"role,omitempty"`
	Reference         Reference       `json:"reference,omitempty"`
}

func (r ConsentProvisionActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ConsentProvisionActor) marshalJSON() jsonConsentProvisionActor {
	m := jsonConsentProvisionActor{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Role = r.Role
	m.Reference = r.Reference
	return m
}
func (r *ConsentProvisionActor) UnmarshalJSON(b []byte) error {
	var m jsonConsentProvisionActor
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ConsentProvisionActor) unmarshalJSON(m jsonConsentProvisionActor) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Role = m.Role
	r.Reference = m.Reference
	return nil
}
func (r ConsentProvisionActor) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The resources controlled by this rule if specific resources are referenced.
type ConsentProvisionData struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// How the resource reference is interpreted when testing consent restrictions.
	Meaning Code
	// A reference to a specific resource that defines which resources are covered by this consent.
	Reference Reference
}
type jsonConsentProvisionData struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Meaning                 Code              `json:"meaning,omitempty"`
	MeaningPrimitiveElement *primitiveElement `json:"_meaning,omitempty"`
	Reference               Reference         `json:"reference,omitempty"`
}

func (r ConsentProvisionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ConsentProvisionData) marshalJSON() jsonConsentProvisionData {
	m := jsonConsentProvisionData{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Meaning = r.Meaning
	if r.Meaning.Id != nil || r.Meaning.Extension != nil {
		m.MeaningPrimitiveElement = &primitiveElement{Id: r.Meaning.Id, Extension: r.Meaning.Extension}
	}
	m.Reference = r.Reference
	return m
}
func (r *ConsentProvisionData) UnmarshalJSON(b []byte) error {
	var m jsonConsentProvisionData
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ConsentProvisionData) unmarshalJSON(m jsonConsentProvisionData) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Meaning = m.Meaning
	if m.MeaningPrimitiveElement != nil {
		r.Meaning.Id = m.MeaningPrimitiveElement.Id
		r.Meaning.Extension = m.MeaningPrimitiveElement.Extension
	}
	r.Reference = m.Reference
	return nil
}
func (r ConsentProvisionData) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
