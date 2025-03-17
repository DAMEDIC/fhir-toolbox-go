package r4

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	model "github.com/DAMEDIC/fhir-toolbox-go/model"
	"io"
	"slices"
	"unsafe"
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

func (r VerificationResult) ResourceType() string {
	return "VerificationResult"
}
func (r VerificationResult) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r VerificationResult) MemSize() int {
	var emptyIface any
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += r.Id.MemSize()
	}
	if r.Meta != nil {
		s += r.Meta.MemSize()
	}
	if r.ImplicitRules != nil {
		s += r.ImplicitRules.MemSize()
	}
	if r.Language != nil {
		s += r.Language.MemSize()
	}
	if r.Text != nil {
		s += r.Text.MemSize()
	}
	for _, i := range r.Contained {
		s += i.MemSize()
	}
	s += (cap(r.Contained) - len(r.Contained)) * int(unsafe.Sizeof(emptyIface))
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Target {
		s += i.MemSize()
	}
	s += (cap(r.Target) - len(r.Target)) * int(unsafe.Sizeof(Reference{}))
	for _, i := range r.TargetLocation {
		s += i.MemSize()
	}
	s += (cap(r.TargetLocation) - len(r.TargetLocation)) * int(unsafe.Sizeof(String{}))
	if r.Need != nil {
		s += r.Need.MemSize()
	}
	s += r.Status.MemSize() - int(unsafe.Sizeof(r.Status))
	if r.StatusDate != nil {
		s += r.StatusDate.MemSize()
	}
	if r.ValidationType != nil {
		s += r.ValidationType.MemSize()
	}
	for _, i := range r.ValidationProcess {
		s += i.MemSize()
	}
	s += (cap(r.ValidationProcess) - len(r.ValidationProcess)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Frequency != nil {
		s += r.Frequency.MemSize()
	}
	if r.LastPerformed != nil {
		s += r.LastPerformed.MemSize()
	}
	if r.NextScheduled != nil {
		s += r.NextScheduled.MemSize()
	}
	if r.FailureAction != nil {
		s += r.FailureAction.MemSize()
	}
	for _, i := range r.PrimarySource {
		s += i.MemSize()
	}
	s += (cap(r.PrimarySource) - len(r.PrimarySource)) * int(unsafe.Sizeof(VerificationResultPrimarySource{}))
	if r.Attestation != nil {
		s += r.Attestation.MemSize()
	}
	for _, i := range r.Validator {
		s += i.MemSize()
	}
	s += (cap(r.Validator) - len(r.Validator)) * int(unsafe.Sizeof(VerificationResultValidator{}))
	return s
}
func (r VerificationResultPrimarySource) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Who != nil {
		s += r.Who.MemSize()
	}
	for _, i := range r.Type {
		s += i.MemSize()
	}
	s += (cap(r.Type) - len(r.Type)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.CommunicationMethod {
		s += i.MemSize()
	}
	s += (cap(r.CommunicationMethod) - len(r.CommunicationMethod)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.ValidationStatus != nil {
		s += r.ValidationStatus.MemSize()
	}
	if r.ValidationDate != nil {
		s += r.ValidationDate.MemSize()
	}
	if r.CanPushUpdates != nil {
		s += r.CanPushUpdates.MemSize()
	}
	for _, i := range r.PushTypeAvailable {
		s += i.MemSize()
	}
	s += (cap(r.PushTypeAvailable) - len(r.PushTypeAvailable)) * int(unsafe.Sizeof(CodeableConcept{}))
	return s
}
func (r VerificationResultAttestation) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Who != nil {
		s += r.Who.MemSize()
	}
	if r.OnBehalfOf != nil {
		s += r.OnBehalfOf.MemSize()
	}
	if r.CommunicationMethod != nil {
		s += r.CommunicationMethod.MemSize()
	}
	if r.Date != nil {
		s += r.Date.MemSize()
	}
	if r.SourceIdentityCertificate != nil {
		s += r.SourceIdentityCertificate.MemSize()
	}
	if r.ProxyIdentityCertificate != nil {
		s += r.ProxyIdentityCertificate.MemSize()
	}
	if r.ProxySignature != nil {
		s += r.ProxySignature.MemSize()
	}
	if r.SourceSignature != nil {
		s += r.SourceSignature.MemSize()
	}
	return s
}
func (r VerificationResultValidator) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	s += r.Organization.MemSize() - int(unsafe.Sizeof(r.Organization))
	if r.IdentityCertificate != nil {
		s += r.IdentityCertificate.MemSize()
	}
	if r.AttestationSignature != nil {
		s += r.AttestationSignature.MemSize()
	}
	return s
}
func (r VerificationResult) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r VerificationResultPrimarySource) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r VerificationResultAttestation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r VerificationResultValidator) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r VerificationResult) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r VerificationResult) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"VerificationResult\""))
	if err != nil {
		return err
	}
	setComma := true
	if r.Id != nil && r.Id.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		p := primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_id\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Meta != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"meta\":"))
		if err != nil {
			return err
		}
		err = r.Meta.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"implicitRules\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ImplicitRules)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		p := primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_implicitRules\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && r.Language.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"language\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Language)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		p := primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_language\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Text != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"text\":"))
		if err != nil {
			return err
		}
		err = r.Text.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Contained) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"contained\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, c := range r.Contained {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = ContainedResource{c}.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Target) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"target\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Target {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	{
		anyValue := false
		for _, e := range r.TargetLocation {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"targetLocation\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.TargetLocation)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.TargetLocation {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_targetLocation\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.TargetLocation {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	if r.Need != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"need\":"))
		if err != nil {
			return err
		}
		err = r.Need.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"status\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Status)
		if err != nil {
			return err
		}
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		p := primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_status\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.StatusDate != nil && r.StatusDate.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"statusDate\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.StatusDate)
		if err != nil {
			return err
		}
	}
	if r.StatusDate != nil && (r.StatusDate.Id != nil || r.StatusDate.Extension != nil) {
		p := primitiveElement{Id: r.StatusDate.Id, Extension: r.StatusDate.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_statusDate\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ValidationType != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"validationType\":"))
		if err != nil {
			return err
		}
		err = r.ValidationType.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.ValidationProcess) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"validationProcess\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ValidationProcess {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Frequency != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"frequency\":"))
		if err != nil {
			return err
		}
		err = r.Frequency.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.LastPerformed != nil && r.LastPerformed.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"lastPerformed\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.LastPerformed)
		if err != nil {
			return err
		}
	}
	if r.LastPerformed != nil && (r.LastPerformed.Id != nil || r.LastPerformed.Extension != nil) {
		p := primitiveElement{Id: r.LastPerformed.Id, Extension: r.LastPerformed.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_lastPerformed\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.NextScheduled != nil && r.NextScheduled.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"nextScheduled\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.NextScheduled)
		if err != nil {
			return err
		}
	}
	if r.NextScheduled != nil && (r.NextScheduled.Id != nil || r.NextScheduled.Extension != nil) {
		p := primitiveElement{Id: r.NextScheduled.Id, Extension: r.NextScheduled.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_nextScheduled\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.FailureAction != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"failureAction\":"))
		if err != nil {
			return err
		}
		err = r.FailureAction.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.PrimarySource) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"primarySource\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.PrimarySource {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Attestation != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"attestation\":"))
		if err != nil {
			return err
		}
		err = r.Attestation.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Validator) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"validator\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Validator {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r VerificationResultPrimarySource) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r VerificationResultPrimarySource) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Who != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"who\":"))
		if err != nil {
			return err
		}
		err = r.Who.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Type) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"type\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Type {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.CommunicationMethod) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"communicationMethod\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.CommunicationMethod {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.ValidationStatus != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"validationStatus\":"))
		if err != nil {
			return err
		}
		err = r.ValidationStatus.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ValidationDate != nil && r.ValidationDate.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"validationDate\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ValidationDate)
		if err != nil {
			return err
		}
	}
	if r.ValidationDate != nil && (r.ValidationDate.Id != nil || r.ValidationDate.Extension != nil) {
		p := primitiveElement{Id: r.ValidationDate.Id, Extension: r.ValidationDate.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_validationDate\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.CanPushUpdates != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"canPushUpdates\":"))
		if err != nil {
			return err
		}
		err = r.CanPushUpdates.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.PushTypeAvailable) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"pushTypeAvailable\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.PushTypeAvailable {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r VerificationResultAttestation) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r VerificationResultAttestation) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Who != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"who\":"))
		if err != nil {
			return err
		}
		err = r.Who.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.OnBehalfOf != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"onBehalfOf\":"))
		if err != nil {
			return err
		}
		err = r.OnBehalfOf.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.CommunicationMethod != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"communicationMethod\":"))
		if err != nil {
			return err
		}
		err = r.CommunicationMethod.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Date != nil && r.Date.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"date\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Date)
		if err != nil {
			return err
		}
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		p := primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_date\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.SourceIdentityCertificate != nil && r.SourceIdentityCertificate.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sourceIdentityCertificate\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.SourceIdentityCertificate)
		if err != nil {
			return err
		}
	}
	if r.SourceIdentityCertificate != nil && (r.SourceIdentityCertificate.Id != nil || r.SourceIdentityCertificate.Extension != nil) {
		p := primitiveElement{Id: r.SourceIdentityCertificate.Id, Extension: r.SourceIdentityCertificate.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_sourceIdentityCertificate\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ProxyIdentityCertificate != nil && r.ProxyIdentityCertificate.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"proxyIdentityCertificate\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ProxyIdentityCertificate)
		if err != nil {
			return err
		}
	}
	if r.ProxyIdentityCertificate != nil && (r.ProxyIdentityCertificate.Id != nil || r.ProxyIdentityCertificate.Extension != nil) {
		p := primitiveElement{Id: r.ProxyIdentityCertificate.Id, Extension: r.ProxyIdentityCertificate.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_proxyIdentityCertificate\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ProxySignature != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"proxySignature\":"))
		if err != nil {
			return err
		}
		err = r.ProxySignature.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.SourceSignature != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sourceSignature\":"))
		if err != nil {
			return err
		}
		err = r.SourceSignature.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r VerificationResultValidator) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r VerificationResultValidator) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"organization\":"))
	if err != nil {
		return err
	}
	err = r.Organization.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.IdentityCertificate != nil && r.IdentityCertificate.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"identityCertificate\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.IdentityCertificate)
		if err != nil {
			return err
		}
	}
	if r.IdentityCertificate != nil && (r.IdentityCertificate.Id != nil || r.IdentityCertificate.Extension != nil) {
		p := primitiveElement{Id: r.IdentityCertificate.Id, Extension: r.IdentityCertificate.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_identityCertificate\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.AttestationSignature != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"attestationSignature\":"))
		if err != nil {
			return err
		}
		err = r.AttestationSignature.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r *VerificationResult) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *VerificationResult) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in VerificationResult element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in VerificationResult element", t)
		}
		switch f {
		case "resourceType":
			_, err := d.Token()
			if err != nil {
				return err
			}
		case "id":
			var v Id
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Value = v.Value
		case "_id":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Id = v.Id
			r.Id.Extension = v.Extension
		case "meta":
			var v Meta
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Meta = &v
		case "implicitRules":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Value = v.Value
		case "_implicitRules":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Id = v.Id
			r.ImplicitRules.Extension = v.Extension
		case "language":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Value = v.Value
		case "_language":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Id = v.Id
			r.Language.Extension = v.Extension
		case "text":
			var v Narrative
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Text = &v
		case "contained":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResult element", t)
			}
			for d.More() {
				var v ContainedResource
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Contained = append(r.Contained, v.Resource)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResult element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResult element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResult element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResult element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResult element", t)
			}
		case "target":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResult element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Target = append(r.Target, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResult element", t)
			}
		case "targetLocation":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResult element", t)
			}
			for i := 0; d.More(); i++ {
				var v String
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.TargetLocation) <= i {
					r.TargetLocation = append(r.TargetLocation, String{})
				}
				r.TargetLocation[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResult element", t)
			}
		case "_targetLocation":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResult element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.TargetLocation) <= i {
					r.TargetLocation = append(r.TargetLocation, String{})
				}
				r.TargetLocation[i].Id = v.Id
				r.TargetLocation[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResult element", t)
			}
		case "need":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Need = &v
		case "status":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Status.Value = v.Value
		case "_status":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Status.Id = v.Id
			r.Status.Extension = v.Extension
		case "statusDate":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.StatusDate == nil {
				r.StatusDate = &DateTime{}
			}
			r.StatusDate.Value = v.Value
		case "_statusDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.StatusDate == nil {
				r.StatusDate = &DateTime{}
			}
			r.StatusDate.Id = v.Id
			r.StatusDate.Extension = v.Extension
		case "validationType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ValidationType = &v
		case "validationProcess":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResult element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ValidationProcess = append(r.ValidationProcess, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResult element", t)
			}
		case "frequency":
			var v Timing
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Frequency = &v
		case "lastPerformed":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.LastPerformed == nil {
				r.LastPerformed = &DateTime{}
			}
			r.LastPerformed.Value = v.Value
		case "_lastPerformed":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.LastPerformed == nil {
				r.LastPerformed = &DateTime{}
			}
			r.LastPerformed.Id = v.Id
			r.LastPerformed.Extension = v.Extension
		case "nextScheduled":
			var v Date
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.NextScheduled == nil {
				r.NextScheduled = &Date{}
			}
			r.NextScheduled.Value = v.Value
		case "_nextScheduled":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.NextScheduled == nil {
				r.NextScheduled = &Date{}
			}
			r.NextScheduled.Id = v.Id
			r.NextScheduled.Extension = v.Extension
		case "failureAction":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.FailureAction = &v
		case "primarySource":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResult element", t)
			}
			for d.More() {
				var v VerificationResultPrimarySource
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.PrimarySource = append(r.PrimarySource, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResult element", t)
			}
		case "attestation":
			var v VerificationResultAttestation
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Attestation = &v
		case "validator":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResult element", t)
			}
			for d.More() {
				var v VerificationResultValidator
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Validator = append(r.Validator, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResult element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in VerificationResult", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in VerificationResult element", t)
	}
	return nil
}
func (r *VerificationResultPrimarySource) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in VerificationResultPrimarySource element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in VerificationResultPrimarySource element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResultPrimarySource element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResultPrimarySource element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResultPrimarySource element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResultPrimarySource element", t)
			}
		case "who":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Who = &v
		case "type":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResultPrimarySource element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Type = append(r.Type, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResultPrimarySource element", t)
			}
		case "communicationMethod":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResultPrimarySource element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.CommunicationMethod = append(r.CommunicationMethod, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResultPrimarySource element", t)
			}
		case "validationStatus":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ValidationStatus = &v
		case "validationDate":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ValidationDate == nil {
				r.ValidationDate = &DateTime{}
			}
			r.ValidationDate.Value = v.Value
		case "_validationDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ValidationDate == nil {
				r.ValidationDate = &DateTime{}
			}
			r.ValidationDate.Id = v.Id
			r.ValidationDate.Extension = v.Extension
		case "canPushUpdates":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.CanPushUpdates = &v
		case "pushTypeAvailable":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResultPrimarySource element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.PushTypeAvailable = append(r.PushTypeAvailable, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResultPrimarySource element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in VerificationResultPrimarySource", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in VerificationResultPrimarySource element", t)
	}
	return nil
}
func (r *VerificationResultAttestation) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in VerificationResultAttestation element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in VerificationResultAttestation element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResultAttestation element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResultAttestation element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResultAttestation element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResultAttestation element", t)
			}
		case "who":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Who = &v
		case "onBehalfOf":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.OnBehalfOf = &v
		case "communicationMethod":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.CommunicationMethod = &v
		case "date":
			var v Date
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Date == nil {
				r.Date = &Date{}
			}
			r.Date.Value = v.Value
		case "_date":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Date == nil {
				r.Date = &Date{}
			}
			r.Date.Id = v.Id
			r.Date.Extension = v.Extension
		case "sourceIdentityCertificate":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.SourceIdentityCertificate == nil {
				r.SourceIdentityCertificate = &String{}
			}
			r.SourceIdentityCertificate.Value = v.Value
		case "_sourceIdentityCertificate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.SourceIdentityCertificate == nil {
				r.SourceIdentityCertificate = &String{}
			}
			r.SourceIdentityCertificate.Id = v.Id
			r.SourceIdentityCertificate.Extension = v.Extension
		case "proxyIdentityCertificate":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ProxyIdentityCertificate == nil {
				r.ProxyIdentityCertificate = &String{}
			}
			r.ProxyIdentityCertificate.Value = v.Value
		case "_proxyIdentityCertificate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ProxyIdentityCertificate == nil {
				r.ProxyIdentityCertificate = &String{}
			}
			r.ProxyIdentityCertificate.Id = v.Id
			r.ProxyIdentityCertificate.Extension = v.Extension
		case "proxySignature":
			var v Signature
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ProxySignature = &v
		case "sourceSignature":
			var v Signature
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.SourceSignature = &v
		default:
			return fmt.Errorf("invalid field: %s in VerificationResultAttestation", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in VerificationResultAttestation element", t)
	}
	return nil
}
func (r *VerificationResultValidator) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in VerificationResultValidator element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in VerificationResultValidator element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResultValidator element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResultValidator element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VerificationResultValidator element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VerificationResultValidator element", t)
			}
		case "organization":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Organization = v
		case "identityCertificate":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.IdentityCertificate == nil {
				r.IdentityCertificate = &String{}
			}
			r.IdentityCertificate.Value = v.Value
		case "_identityCertificate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.IdentityCertificate == nil {
				r.IdentityCertificate = &String{}
			}
			r.IdentityCertificate.Id = v.Id
			r.IdentityCertificate.Extension = v.Extension
		case "attestationSignature":
			var v Signature
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.AttestationSignature = &v
		default:
			return fmt.Errorf("invalid field: %s in VerificationResultValidator", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in VerificationResultValidator element", t)
	}
	return nil
}
func (r VerificationResult) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "VerificationResult"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Id, xml.StartElement{Name: xml.Name{Local: "id"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Meta, xml.StartElement{Name: xml.Name{Local: "meta"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ImplicitRules, xml.StartElement{Name: xml.Name{Local: "implicitRules"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	for _, c := range r.Contained {
		err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "contained"}})
		if err != nil {
			return err
		}
		err = e.EncodeElement(c, xml.StartElement{Name: xml.Name{Local: "__contained__"}})
		if err != nil {
			return err
		}
		err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "contained"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Target, xml.StartElement{Name: xml.Name{Local: "target"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TargetLocation, xml.StartElement{Name: xml.Name{Local: "targetLocation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Need, xml.StartElement{Name: xml.Name{Local: "need"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.StatusDate, xml.StartElement{Name: xml.Name{Local: "statusDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ValidationType, xml.StartElement{Name: xml.Name{Local: "validationType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ValidationProcess, xml.StartElement{Name: xml.Name{Local: "validationProcess"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Frequency, xml.StartElement{Name: xml.Name{Local: "frequency"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LastPerformed, xml.StartElement{Name: xml.Name{Local: "lastPerformed"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NextScheduled, xml.StartElement{Name: xml.Name{Local: "nextScheduled"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FailureAction, xml.StartElement{Name: xml.Name{Local: "failureAction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PrimarySource, xml.StartElement{Name: xml.Name{Local: "primarySource"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Attestation, xml.StartElement{Name: xml.Name{Local: "attestation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Validator, xml.StartElement{Name: xml.Name{Local: "validator"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r VerificationResultPrimarySource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Who, xml.StartElement{Name: xml.Name{Local: "who"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CommunicationMethod, xml.StartElement{Name: xml.Name{Local: "communicationMethod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ValidationStatus, xml.StartElement{Name: xml.Name{Local: "validationStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ValidationDate, xml.StartElement{Name: xml.Name{Local: "validationDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CanPushUpdates, xml.StartElement{Name: xml.Name{Local: "canPushUpdates"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PushTypeAvailable, xml.StartElement{Name: xml.Name{Local: "pushTypeAvailable"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r VerificationResultAttestation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Who, xml.StartElement{Name: xml.Name{Local: "who"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OnBehalfOf, xml.StartElement{Name: xml.Name{Local: "onBehalfOf"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CommunicationMethod, xml.StartElement{Name: xml.Name{Local: "communicationMethod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SourceIdentityCertificate, xml.StartElement{Name: xml.Name{Local: "sourceIdentityCertificate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProxyIdentityCertificate, xml.StartElement{Name: xml.Name{Local: "proxyIdentityCertificate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProxySignature, xml.StartElement{Name: xml.Name{Local: "proxySignature"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SourceSignature, xml.StartElement{Name: xml.Name{Local: "sourceSignature"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r VerificationResultValidator) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Id != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "id"},
			Value: *r.Id,
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Extension, xml.StartElement{Name: xml.Name{Local: "extension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ModifierExtension, xml.StartElement{Name: xml.Name{Local: "modifierExtension"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Organization, xml.StartElement{Name: xml.Name{Local: "organization"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IdentityCertificate, xml.StartElement{Name: xml.Name{Local: "identityCertificate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AttestationSignature, xml.StartElement{Name: xml.Name{Local: "attestationSignature"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *VerificationResult) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "id":
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Id = &v
			case "meta":
				var v Meta
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Meta = &v
			case "implicitRules":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ImplicitRules = &v
			case "language":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = &v
			case "text":
				var v Narrative
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = &v
			case "contained":
				var c ContainedResource
				err := d.DecodeElement(&c, &t)
				if err != nil {
					return err
				}
				r.Contained = append(r.Contained, c.Resource)
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "target":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Target = append(r.Target, v)
			case "targetLocation":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TargetLocation = append(r.TargetLocation, v)
			case "need":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Need = &v
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "statusDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StatusDate = &v
			case "validationType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ValidationType = &v
			case "validationProcess":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ValidationProcess = append(r.ValidationProcess, v)
			case "frequency":
				var v Timing
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Frequency = &v
			case "lastPerformed":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LastPerformed = &v
			case "nextScheduled":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NextScheduled = &v
			case "failureAction":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FailureAction = &v
			case "primarySource":
				var v VerificationResultPrimarySource
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PrimarySource = append(r.PrimarySource, v)
			case "attestation":
				var v VerificationResultAttestation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Attestation = &v
			case "validator":
				var v VerificationResultValidator
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Validator = append(r.Validator, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *VerificationResultPrimarySource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "who":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Who = &v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = append(r.Type, v)
			case "communicationMethod":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CommunicationMethod = append(r.CommunicationMethod, v)
			case "validationStatus":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ValidationStatus = &v
			case "validationDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ValidationDate = &v
			case "canPushUpdates":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CanPushUpdates = &v
			case "pushTypeAvailable":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PushTypeAvailable = append(r.PushTypeAvailable, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *VerificationResultAttestation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "who":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Who = &v
			case "onBehalfOf":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OnBehalfOf = &v
			case "communicationMethod":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CommunicationMethod = &v
			case "date":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = &v
			case "sourceIdentityCertificate":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SourceIdentityCertificate = &v
			case "proxyIdentityCertificate":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProxyIdentityCertificate = &v
			case "proxySignature":
				var v Signature
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProxySignature = &v
			case "sourceSignature":
				var v Signature
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SourceSignature = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *VerificationResultValidator) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "http://hl7.org/fhir" {
		return fmt.Errorf("invalid namespace: \"%s\", expected: \"http://hl7.org/fhir\"", start.Name.Space)
	}
	for _, a := range start.Attr {
		if a.Name.Space != "" {
			return fmt.Errorf("invalid attribute namespace: \"%s\", expected default namespace", start.Name.Space)
		}
		switch a.Name.Local {
		case "xmlns":
			continue
		case "id":
			r.Id = &a.Value
		default:
			return fmt.Errorf("invalid attribute: \"%s\"", a.Name.Local)
		}
	}
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "extension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			case "modifierExtension":
				var v Extension
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			case "organization":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Organization = v
			case "identityCertificate":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IdentityCertificate = &v
			case "attestationSignature":
				var v Signature
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AttestationSignature = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r VerificationResult) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, *r.Id)
		}
	}
	if len(name) == 0 || slices.Contains(name, "meta") {
		if r.Meta != nil {
			children = append(children, *r.Meta)
		}
	}
	if len(name) == 0 || slices.Contains(name, "implicitRules") {
		if r.ImplicitRules != nil {
			children = append(children, *r.ImplicitRules)
		}
	}
	if len(name) == 0 || slices.Contains(name, "language") {
		if r.Language != nil {
			children = append(children, *r.Language)
		}
	}
	if len(name) == 0 || slices.Contains(name, "text") {
		if r.Text != nil {
			children = append(children, *r.Text)
		}
	}
	if len(name) == 0 || slices.Contains(name, "contained") {
		for _, v := range r.Contained {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "target") {
		for _, v := range r.Target {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "targetLocation") {
		for _, v := range r.TargetLocation {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "need") {
		if r.Need != nil {
			children = append(children, *r.Need)
		}
	}
	if len(name) == 0 || slices.Contains(name, "status") {
		children = append(children, r.Status)
	}
	if len(name) == 0 || slices.Contains(name, "statusDate") {
		if r.StatusDate != nil {
			children = append(children, *r.StatusDate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "validationType") {
		if r.ValidationType != nil {
			children = append(children, *r.ValidationType)
		}
	}
	if len(name) == 0 || slices.Contains(name, "validationProcess") {
		for _, v := range r.ValidationProcess {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "frequency") {
		if r.Frequency != nil {
			children = append(children, *r.Frequency)
		}
	}
	if len(name) == 0 || slices.Contains(name, "lastPerformed") {
		if r.LastPerformed != nil {
			children = append(children, *r.LastPerformed)
		}
	}
	if len(name) == 0 || slices.Contains(name, "nextScheduled") {
		if r.NextScheduled != nil {
			children = append(children, *r.NextScheduled)
		}
	}
	if len(name) == 0 || slices.Contains(name, "failureAction") {
		if r.FailureAction != nil {
			children = append(children, *r.FailureAction)
		}
	}
	if len(name) == 0 || slices.Contains(name, "primarySource") {
		for _, v := range r.PrimarySource {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "attestation") {
		if r.Attestation != nil {
			children = append(children, *r.Attestation)
		}
	}
	if len(name) == 0 || slices.Contains(name, "validator") {
		for _, v := range r.Validator {
			children = append(children, v)
		}
	}
	return children
}
func (r VerificationResult) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert VerificationResult to Boolean")
}
func (r VerificationResult) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert VerificationResult to String")
}
func (r VerificationResult) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert VerificationResult to Integer")
}
func (r VerificationResult) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert VerificationResult to Decimal")
}
func (r VerificationResult) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert VerificationResult to Date")
}
func (r VerificationResult) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert VerificationResult to Time")
}
func (r VerificationResult) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert VerificationResult to DateTime")
}
func (r VerificationResult) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert VerificationResult to Quantity")
}
func (r VerificationResult) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *VerificationResult
	switch other := other.(type) {
	case VerificationResult:
		o = &other
	case *VerificationResult:
		o = other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r VerificationResult) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *VerificationResult
	switch other := other.(type) {
	case VerificationResult:
		o = &other
	case *VerificationResult:
		o = other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r VerificationResult) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Id",
				Namespace: "FHIR",
			},
		}, {
			Name: "Meta",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Meta",
				Namespace: "FHIR",
			},
		}, {
			Name: "ImplicitRules",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Uri",
				Namespace: "FHIR",
			},
		}, {
			Name: "Language",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Text",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Narrative",
				Namespace: "FHIR",
			},
		}, {
			Name: "Contained",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Target",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "TargetLocation",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Need",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Status",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "StatusDate",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "DateTime",
				Namespace: "FHIR",
			},
		}, {
			Name: "ValidationType",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ValidationProcess",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Frequency",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Timing",
				Namespace: "FHIR",
			},
		}, {
			Name: "LastPerformed",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "DateTime",
				Namespace: "FHIR",
			},
		}, {
			Name: "NextScheduled",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Date",
				Namespace: "FHIR",
			},
		}, {
			Name: "FailureAction",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "PrimarySource",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "VerificationResultPrimarySource",
				Namespace: "FHIR",
			},
		}, {
			Name: "Attestation",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "VerificationResultAttestation",
				Namespace: "FHIR",
			},
		}, {
			Name: "Validator",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "VerificationResultValidator",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "VerificationResult",
			Namespace: "FHIR",
		},
	}
}
func (r VerificationResultPrimarySource) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "who") {
		if r.Who != nil {
			children = append(children, *r.Who)
		}
	}
	if len(name) == 0 || slices.Contains(name, "type") {
		for _, v := range r.Type {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "communicationMethod") {
		for _, v := range r.CommunicationMethod {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "validationStatus") {
		if r.ValidationStatus != nil {
			children = append(children, *r.ValidationStatus)
		}
	}
	if len(name) == 0 || slices.Contains(name, "validationDate") {
		if r.ValidationDate != nil {
			children = append(children, *r.ValidationDate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "canPushUpdates") {
		if r.CanPushUpdates != nil {
			children = append(children, *r.CanPushUpdates)
		}
	}
	if len(name) == 0 || slices.Contains(name, "pushTypeAvailable") {
		for _, v := range r.PushTypeAvailable {
			children = append(children, v)
		}
	}
	return children
}
func (r VerificationResultPrimarySource) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert VerificationResultPrimarySource to Boolean")
}
func (r VerificationResultPrimarySource) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert VerificationResultPrimarySource to String")
}
func (r VerificationResultPrimarySource) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert VerificationResultPrimarySource to Integer")
}
func (r VerificationResultPrimarySource) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert VerificationResultPrimarySource to Decimal")
}
func (r VerificationResultPrimarySource) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert VerificationResultPrimarySource to Date")
}
func (r VerificationResultPrimarySource) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert VerificationResultPrimarySource to Time")
}
func (r VerificationResultPrimarySource) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert VerificationResultPrimarySource to DateTime")
}
func (r VerificationResultPrimarySource) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert VerificationResultPrimarySource to Quantity")
}
func (r VerificationResultPrimarySource) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *VerificationResultPrimarySource
	switch other := other.(type) {
	case VerificationResultPrimarySource:
		o = &other
	case *VerificationResultPrimarySource:
		o = other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r VerificationResultPrimarySource) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *VerificationResultPrimarySource
	switch other := other.(type) {
	case VerificationResultPrimarySource:
		o = &other
	case *VerificationResultPrimarySource:
		o = other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r VerificationResultPrimarySource) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Who",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "CommunicationMethod",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ValidationStatus",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ValidationDate",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "DateTime",
				Namespace: "FHIR",
			},
		}, {
			Name: "CanPushUpdates",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "PushTypeAvailable",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "VerificationResultPrimarySource",
			Namespace: "FHIR",
		},
	}
}
func (r VerificationResultAttestation) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "who") {
		if r.Who != nil {
			children = append(children, *r.Who)
		}
	}
	if len(name) == 0 || slices.Contains(name, "onBehalfOf") {
		if r.OnBehalfOf != nil {
			children = append(children, *r.OnBehalfOf)
		}
	}
	if len(name) == 0 || slices.Contains(name, "communicationMethod") {
		if r.CommunicationMethod != nil {
			children = append(children, *r.CommunicationMethod)
		}
	}
	if len(name) == 0 || slices.Contains(name, "date") {
		if r.Date != nil {
			children = append(children, *r.Date)
		}
	}
	if len(name) == 0 || slices.Contains(name, "sourceIdentityCertificate") {
		if r.SourceIdentityCertificate != nil {
			children = append(children, *r.SourceIdentityCertificate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "proxyIdentityCertificate") {
		if r.ProxyIdentityCertificate != nil {
			children = append(children, *r.ProxyIdentityCertificate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "proxySignature") {
		if r.ProxySignature != nil {
			children = append(children, *r.ProxySignature)
		}
	}
	if len(name) == 0 || slices.Contains(name, "sourceSignature") {
		if r.SourceSignature != nil {
			children = append(children, *r.SourceSignature)
		}
	}
	return children
}
func (r VerificationResultAttestation) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert VerificationResultAttestation to Boolean")
}
func (r VerificationResultAttestation) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert VerificationResultAttestation to String")
}
func (r VerificationResultAttestation) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert VerificationResultAttestation to Integer")
}
func (r VerificationResultAttestation) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert VerificationResultAttestation to Decimal")
}
func (r VerificationResultAttestation) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert VerificationResultAttestation to Date")
}
func (r VerificationResultAttestation) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert VerificationResultAttestation to Time")
}
func (r VerificationResultAttestation) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert VerificationResultAttestation to DateTime")
}
func (r VerificationResultAttestation) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert VerificationResultAttestation to Quantity")
}
func (r VerificationResultAttestation) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *VerificationResultAttestation
	switch other := other.(type) {
	case VerificationResultAttestation:
		o = &other
	case *VerificationResultAttestation:
		o = other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r VerificationResultAttestation) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *VerificationResultAttestation
	switch other := other.(type) {
	case VerificationResultAttestation:
		o = &other
	case *VerificationResultAttestation:
		o = other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r VerificationResultAttestation) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Who",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "OnBehalfOf",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "CommunicationMethod",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Date",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Date",
				Namespace: "FHIR",
			},
		}, {
			Name: "SourceIdentityCertificate",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "ProxyIdentityCertificate",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "ProxySignature",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Signature",
				Namespace: "FHIR",
			},
		}, {
			Name: "SourceSignature",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Signature",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "VerificationResultAttestation",
			Namespace: "FHIR",
		},
	}
}
func (r VerificationResultValidator) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "organization") {
		children = append(children, r.Organization)
	}
	if len(name) == 0 || slices.Contains(name, "identityCertificate") {
		if r.IdentityCertificate != nil {
			children = append(children, *r.IdentityCertificate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "attestationSignature") {
		if r.AttestationSignature != nil {
			children = append(children, *r.AttestationSignature)
		}
	}
	return children
}
func (r VerificationResultValidator) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert VerificationResultValidator to Boolean")
}
func (r VerificationResultValidator) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert VerificationResultValidator to String")
}
func (r VerificationResultValidator) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert VerificationResultValidator to Integer")
}
func (r VerificationResultValidator) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert VerificationResultValidator to Decimal")
}
func (r VerificationResultValidator) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert VerificationResultValidator to Date")
}
func (r VerificationResultValidator) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert VerificationResultValidator to Time")
}
func (r VerificationResultValidator) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert VerificationResultValidator to DateTime")
}
func (r VerificationResultValidator) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert VerificationResultValidator to Quantity")
}
func (r VerificationResultValidator) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *VerificationResultValidator
	switch other := other.(type) {
	case VerificationResultValidator:
		o = &other
	case *VerificationResultValidator:
		o = other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r VerificationResultValidator) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o *VerificationResultValidator
	switch other := other.(type) {
	case VerificationResultValidator:
		o = &other
	case *VerificationResultValidator:
		o = other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r VerificationResultValidator) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Organization",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "IdentityCertificate",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "AttestationSignature",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Signature",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "VerificationResultValidator",
			Namespace: "FHIR",
		},
	}
}
