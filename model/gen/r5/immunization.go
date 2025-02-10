package r5

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

// Describes the event of a patient being administered a vaccine or a record of an immunization as reported by a patient, a clinician or another party.
type Immunization struct {
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
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, nor can they have their own independent transaction scope. This is allowed to be a Parameters resource if and only if it is referenced by a resource that provides context/meaning.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A unique identifier assigned to this immunization record.
	Identifier []Identifier
	// A plan, order or recommendation fulfilled in whole or in part by this immunization.
	BasedOn []Reference
	// Indicates the current status of the immunization event.
	Status Code
	// Indicates the reason the immunization event was not performed.
	StatusReason *CodeableConcept
	// Vaccine that was administered or was to be administered.
	VaccineCode CodeableConcept
	// An indication of which product was administered to the patient. This is typically a more detailed representation of the concept conveyed by the vaccineCode data element. If a Medication resource is referenced, it may be to a stand-alone resource or a contained resource within the Immunization resource.
	AdministeredProduct *CodeableReference
	// Name of vaccine manufacturer.
	Manufacturer *CodeableReference
	// Lot number of the  vaccine product.
	LotNumber *String
	// Date vaccine batch expires.
	ExpirationDate *Date
	// The patient who either received or did not receive the immunization.
	Patient Reference
	// The visit or admission or other contact between patient and health care provider the immunization was performed as part of.
	Encounter *Reference
	// Additional information that is relevant to the immunization (e.g. for a vaccine recipient who is pregnant, the gestational age of the fetus). The reason why a vaccine was given (e.g. occupation, underlying medical condition) should be conveyed in Immunization.reason, not as supporting information. The reason why a vaccine was not given (e.g. contraindication) should be conveyed in Immunization.statusReason, not as supporting information.
	SupportingInformation []Reference
	// Date vaccine administered or was to be administered.
	Occurrence isImmunizationOccurrence
	// Indicates whether the data contained in the resource was captured by the individual/organization which was responsible for the administration of the vaccine rather than as 'secondary reported' data documented by a third party. A value of 'true' means this data originated with the individual/organization which was responsible for the administration of the vaccine.
	PrimarySource *Boolean
	// Typically the source of the data when the report of the immunization event is not based on information from the person who administered the vaccine.
	InformationSource *CodeableReference
	// The service delivery location where the vaccine administration occurred.
	Location *Reference
	// Body site where vaccine was administered.
	Site *CodeableConcept
	// The path by which the vaccine product is taken into the body.
	Route *CodeableConcept
	// The quantity of vaccine product that was administered.
	DoseQuantity *Quantity
	// Indicates who performed the immunization event.
	Performer []ImmunizationPerformer
	// Extra information about the immunization that is not conveyed by the other attributes.
	Note []Annotation
	// Describes why the immunization occurred in coded or textual form, or Indicates another resource (Condition, Observation or DiagnosticReport) whose existence justifies this immunization.
	Reason []CodeableReference
	// Indication if a dose is considered to be subpotent. By default, a dose should be considered to be potent.
	IsSubpotent *Boolean
	// Reason why a dose is considered to be subpotent.
	SubpotentReason []CodeableConcept
	// Indicates a patient's eligibility for a funding program.
	ProgramEligibility []ImmunizationProgramEligibility
	// Indicates the source of the vaccine actually administered. This may be different than the patient eligibility (e.g. the patient may be eligible for a publically purchased vaccine but due to inventory issues, vaccine purchased with private funds was actually administered).
	FundingSource *CodeableConcept
	// Categorical data indicating that an adverse event is associated in time to an immunization.
	Reaction []ImmunizationReaction
	// The protocol (set of recommendations) being followed by the provider who administered the dose.
	ProtocolApplied []ImmunizationProtocolApplied
}
type isImmunizationOccurrence interface {
	model.Element
	isImmunizationOccurrence()
}

func (r DateTime) isImmunizationOccurrence() {}
func (r String) isImmunizationOccurrence()   {}

// Indicates who performed the immunization event.
type ImmunizationPerformer struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Describes the type of performance (e.g. ordering provider, administering provider, etc.).
	Function *CodeableConcept
	// The practitioner or organization who performed the action.
	Actor Reference
}

// Indicates a patient's eligibility for a funding program.
type ImmunizationProgramEligibility struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates which program the patient had their eligility evaluated for.
	Program CodeableConcept
	// Indicates the patient's eligility status for for a specific payment program.
	ProgramStatus CodeableConcept
}

// Categorical data indicating that an adverse event is associated in time to an immunization.
type ImmunizationReaction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Date of reaction to the immunization.
	Date *DateTime
	// Details of the reaction.
	Manifestation *CodeableReference
	// Self-reported indicator.
	Reported *Boolean
}

// The protocol (set of recommendations) being followed by the provider who administered the dose.
type ImmunizationProtocolApplied struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// One possible path to achieve presumed immunity against a disease - within the context of an authority.
	Series *String
	// Indicates the authority who published the protocol (e.g. ACIP) that is being followed.
	Authority *Reference
	// The vaccine preventable disease the dose is being administered against.
	TargetDisease []CodeableConcept
	// Nominal position in a series as intended by the practitioner administering the dose.
	DoseNumber String
	// The recommended number of doses to achieve immunity as intended by the practitioner administering the dose.
	SeriesDoses *String
}

func (r Immunization) ResourceType() string {
	return "Immunization"
}
func (r Immunization) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r Immunization) MemSize() int {
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
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	for _, i := range r.BasedOn {
		s += i.MemSize()
	}
	s += (cap(r.BasedOn) - len(r.BasedOn)) * int(unsafe.Sizeof(Reference{}))
	s += r.Status.MemSize() - int(unsafe.Sizeof(r.Status))
	if r.StatusReason != nil {
		s += r.StatusReason.MemSize()
	}
	s += r.VaccineCode.MemSize() - int(unsafe.Sizeof(r.VaccineCode))
	if r.AdministeredProduct != nil {
		s += r.AdministeredProduct.MemSize()
	}
	if r.Manufacturer != nil {
		s += r.Manufacturer.MemSize()
	}
	if r.LotNumber != nil {
		s += r.LotNumber.MemSize()
	}
	if r.ExpirationDate != nil {
		s += r.ExpirationDate.MemSize()
	}
	s += r.Patient.MemSize() - int(unsafe.Sizeof(r.Patient))
	if r.Encounter != nil {
		s += r.Encounter.MemSize()
	}
	for _, i := range r.SupportingInformation {
		s += i.MemSize()
	}
	s += (cap(r.SupportingInformation) - len(r.SupportingInformation)) * int(unsafe.Sizeof(Reference{}))
	if r.Occurrence != nil {
		s += r.Occurrence.MemSize()
	}
	if r.PrimarySource != nil {
		s += r.PrimarySource.MemSize()
	}
	if r.InformationSource != nil {
		s += r.InformationSource.MemSize()
	}
	if r.Location != nil {
		s += r.Location.MemSize()
	}
	if r.Site != nil {
		s += r.Site.MemSize()
	}
	if r.Route != nil {
		s += r.Route.MemSize()
	}
	if r.DoseQuantity != nil {
		s += r.DoseQuantity.MemSize()
	}
	for _, i := range r.Performer {
		s += i.MemSize()
	}
	s += (cap(r.Performer) - len(r.Performer)) * int(unsafe.Sizeof(ImmunizationPerformer{}))
	for _, i := range r.Note {
		s += i.MemSize()
	}
	s += (cap(r.Note) - len(r.Note)) * int(unsafe.Sizeof(Annotation{}))
	for _, i := range r.Reason {
		s += i.MemSize()
	}
	s += (cap(r.Reason) - len(r.Reason)) * int(unsafe.Sizeof(CodeableReference{}))
	if r.IsSubpotent != nil {
		s += r.IsSubpotent.MemSize()
	}
	for _, i := range r.SubpotentReason {
		s += i.MemSize()
	}
	s += (cap(r.SubpotentReason) - len(r.SubpotentReason)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.ProgramEligibility {
		s += i.MemSize()
	}
	s += (cap(r.ProgramEligibility) - len(r.ProgramEligibility)) * int(unsafe.Sizeof(ImmunizationProgramEligibility{}))
	if r.FundingSource != nil {
		s += r.FundingSource.MemSize()
	}
	for _, i := range r.Reaction {
		s += i.MemSize()
	}
	s += (cap(r.Reaction) - len(r.Reaction)) * int(unsafe.Sizeof(ImmunizationReaction{}))
	for _, i := range r.ProtocolApplied {
		s += i.MemSize()
	}
	s += (cap(r.ProtocolApplied) - len(r.ProtocolApplied)) * int(unsafe.Sizeof(ImmunizationProtocolApplied{}))
	return s
}
func (r ImmunizationPerformer) MemSize() int {
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
	if r.Function != nil {
		s += r.Function.MemSize()
	}
	s += r.Actor.MemSize() - int(unsafe.Sizeof(r.Actor))
	return s
}
func (r ImmunizationProgramEligibility) MemSize() int {
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
	s += r.Program.MemSize() - int(unsafe.Sizeof(r.Program))
	s += r.ProgramStatus.MemSize() - int(unsafe.Sizeof(r.ProgramStatus))
	return s
}
func (r ImmunizationReaction) MemSize() int {
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
	if r.Date != nil {
		s += r.Date.MemSize()
	}
	if r.Manifestation != nil {
		s += r.Manifestation.MemSize()
	}
	if r.Reported != nil {
		s += r.Reported.MemSize()
	}
	return s
}
func (r ImmunizationProtocolApplied) MemSize() int {
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
	if r.Series != nil {
		s += r.Series.MemSize()
	}
	if r.Authority != nil {
		s += r.Authority.MemSize()
	}
	for _, i := range r.TargetDisease {
		s += i.MemSize()
	}
	s += (cap(r.TargetDisease) - len(r.TargetDisease)) * int(unsafe.Sizeof(CodeableConcept{}))
	s += r.DoseNumber.MemSize() - int(unsafe.Sizeof(r.DoseNumber))
	if r.SeriesDoses != nil {
		s += r.SeriesDoses.MemSize()
	}
	return s
}
func (r Immunization) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ImmunizationPerformer) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ImmunizationProgramEligibility) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ImmunizationReaction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ImmunizationProtocolApplied) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Immunization) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Immunization) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"Immunization\""))
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
	if len(r.Identifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"identifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Identifier {
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
	if len(r.BasedOn) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"basedOn\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.BasedOn {
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
	if r.StatusReason != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"statusReason\":"))
		if err != nil {
			return err
		}
		err = r.StatusReason.marshalJSON(w)
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
	_, err = w.Write([]byte("\"vaccineCode\":"))
	if err != nil {
		return err
	}
	err = r.VaccineCode.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.AdministeredProduct != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"administeredProduct\":"))
		if err != nil {
			return err
		}
		err = r.AdministeredProduct.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Manufacturer != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"manufacturer\":"))
		if err != nil {
			return err
		}
		err = r.Manufacturer.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.LotNumber != nil && r.LotNumber.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"lotNumber\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.LotNumber)
		if err != nil {
			return err
		}
	}
	if r.LotNumber != nil && (r.LotNumber.Id != nil || r.LotNumber.Extension != nil) {
		p := primitiveElement{Id: r.LotNumber.Id, Extension: r.LotNumber.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_lotNumber\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ExpirationDate != nil && r.ExpirationDate.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"expirationDate\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ExpirationDate)
		if err != nil {
			return err
		}
	}
	if r.ExpirationDate != nil && (r.ExpirationDate.Id != nil || r.ExpirationDate.Extension != nil) {
		p := primitiveElement{Id: r.ExpirationDate.Id, Extension: r.ExpirationDate.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_expirationDate\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
	_, err = w.Write([]byte("\"patient\":"))
	if err != nil {
		return err
	}
	err = r.Patient.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.Encounter != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"encounter\":"))
		if err != nil {
			return err
		}
		err = r.Encounter.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.SupportingInformation) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"supportingInformation\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.SupportingInformation {
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
	switch v := r.Occurrence.(type) {
	case DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"occurrenceDateTime\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_occurrenceDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"occurrenceDateTime\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_occurrenceDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"occurrenceString\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_occurrenceString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"occurrenceString\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_occurrenceString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	}
	if r.PrimarySource != nil && r.PrimarySource.Value != nil {
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
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.PrimarySource)
		if err != nil {
			return err
		}
	}
	if r.PrimarySource != nil && (r.PrimarySource.Id != nil || r.PrimarySource.Extension != nil) {
		p := primitiveElement{Id: r.PrimarySource.Id, Extension: r.PrimarySource.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_primarySource\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.InformationSource != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"informationSource\":"))
		if err != nil {
			return err
		}
		err = r.InformationSource.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Location != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"location\":"))
		if err != nil {
			return err
		}
		err = r.Location.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Site != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"site\":"))
		if err != nil {
			return err
		}
		err = r.Site.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Route != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"route\":"))
		if err != nil {
			return err
		}
		err = r.Route.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.DoseQuantity != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"doseQuantity\":"))
		if err != nil {
			return err
		}
		err = r.DoseQuantity.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Performer) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"performer\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Performer {
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
	if len(r.Note) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"note\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Note {
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
	if len(r.Reason) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reason\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Reason {
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
	if r.IsSubpotent != nil && r.IsSubpotent.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"isSubpotent\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.IsSubpotent)
		if err != nil {
			return err
		}
	}
	if r.IsSubpotent != nil && (r.IsSubpotent.Id != nil || r.IsSubpotent.Extension != nil) {
		p := primitiveElement{Id: r.IsSubpotent.Id, Extension: r.IsSubpotent.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_isSubpotent\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.SubpotentReason) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"subpotentReason\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.SubpotentReason {
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
	if len(r.ProgramEligibility) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"programEligibility\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ProgramEligibility {
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
	if r.FundingSource != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"fundingSource\":"))
		if err != nil {
			return err
		}
		err = r.FundingSource.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Reaction) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reaction\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Reaction {
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
	if len(r.ProtocolApplied) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"protocolApplied\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ProtocolApplied {
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
func (r ImmunizationPerformer) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ImmunizationPerformer) marshalJSON(w io.Writer) error {
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
	if r.Function != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"function\":"))
		if err != nil {
			return err
		}
		err = r.Function.marshalJSON(w)
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
	_, err = w.Write([]byte("\"actor\":"))
	if err != nil {
		return err
	}
	err = r.Actor.marshalJSON(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r ImmunizationProgramEligibility) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ImmunizationProgramEligibility) marshalJSON(w io.Writer) error {
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
	_, err = w.Write([]byte("\"program\":"))
	if err != nil {
		return err
	}
	err = r.Program.marshalJSON(w)
	if err != nil {
		return err
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"programStatus\":"))
	if err != nil {
		return err
	}
	err = r.ProgramStatus.marshalJSON(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r ImmunizationReaction) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ImmunizationReaction) marshalJSON(w io.Writer) error {
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
	if r.Manifestation != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"manifestation\":"))
		if err != nil {
			return err
		}
		err = r.Manifestation.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Reported != nil && r.Reported.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"reported\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Reported)
		if err != nil {
			return err
		}
	}
	if r.Reported != nil && (r.Reported.Id != nil || r.Reported.Extension != nil) {
		p := primitiveElement{Id: r.Reported.Id, Extension: r.Reported.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_reported\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r ImmunizationProtocolApplied) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ImmunizationProtocolApplied) marshalJSON(w io.Writer) error {
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
	if r.Series != nil && r.Series.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"series\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Series)
		if err != nil {
			return err
		}
	}
	if r.Series != nil && (r.Series.Id != nil || r.Series.Extension != nil) {
		p := primitiveElement{Id: r.Series.Id, Extension: r.Series.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_series\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Authority != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"authority\":"))
		if err != nil {
			return err
		}
		err = r.Authority.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.TargetDisease) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"targetDisease\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.TargetDisease {
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
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"doseNumber\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.DoseNumber)
		if err != nil {
			return err
		}
	}
	if r.DoseNumber.Id != nil || r.DoseNumber.Extension != nil {
		p := primitiveElement{Id: r.DoseNumber.Id, Extension: r.DoseNumber.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_doseNumber\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.SeriesDoses != nil && r.SeriesDoses.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"seriesDoses\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.SeriesDoses)
		if err != nil {
			return err
		}
	}
	if r.SeriesDoses != nil && (r.SeriesDoses.Id != nil || r.SeriesDoses.Extension != nil) {
		p := primitiveElement{Id: r.SeriesDoses.Id, Extension: r.SeriesDoses.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_seriesDoses\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r *Immunization) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *Immunization) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in Immunization element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in Immunization element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in Immunization element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Immunization element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Immunization element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Immunization element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Immunization element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Immunization element", t)
			}
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Immunization element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Immunization element", t)
			}
		case "basedOn":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Immunization element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.BasedOn = append(r.BasedOn, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Immunization element", t)
			}
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
		case "statusReason":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.StatusReason = &v
		case "vaccineCode":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.VaccineCode = v
		case "administeredProduct":
			var v CodeableReference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.AdministeredProduct = &v
		case "manufacturer":
			var v CodeableReference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Manufacturer = &v
		case "lotNumber":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.LotNumber == nil {
				r.LotNumber = &String{}
			}
			r.LotNumber.Value = v.Value
		case "_lotNumber":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.LotNumber == nil {
				r.LotNumber = &String{}
			}
			r.LotNumber.Id = v.Id
			r.LotNumber.Extension = v.Extension
		case "expirationDate":
			var v Date
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ExpirationDate == nil {
				r.ExpirationDate = &Date{}
			}
			r.ExpirationDate.Value = v.Value
		case "_expirationDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ExpirationDate == nil {
				r.ExpirationDate = &Date{}
			}
			r.ExpirationDate.Id = v.Id
			r.ExpirationDate.Extension = v.Extension
		case "patient":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Patient = v
		case "encounter":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Encounter = &v
		case "supportingInformation":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Immunization element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.SupportingInformation = append(r.SupportingInformation, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Immunization element", t)
			}
		case "occurrenceDateTime":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Occurrence != nil {
				r.Occurrence = DateTime{
					Extension: r.Occurrence.(DateTime).Extension,
					Id:        r.Occurrence.(DateTime).Id,
					Value:     v.Value,
				}
			} else {
				r.Occurrence = v
			}
		case "_occurrenceDateTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Occurrence != nil {
				r.Occurrence = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Occurrence.(DateTime).Value,
				}
			} else {
				r.Occurrence = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "occurrenceString":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Occurrence != nil {
				r.Occurrence = String{
					Extension: r.Occurrence.(String).Extension,
					Id:        r.Occurrence.(String).Id,
					Value:     v.Value,
				}
			} else {
				r.Occurrence = v
			}
		case "_occurrenceString":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Occurrence != nil {
				r.Occurrence = String{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Occurrence.(String).Value,
				}
			} else {
				r.Occurrence = String{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "primarySource":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.PrimarySource == nil {
				r.PrimarySource = &Boolean{}
			}
			r.PrimarySource.Value = v.Value
		case "_primarySource":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.PrimarySource == nil {
				r.PrimarySource = &Boolean{}
			}
			r.PrimarySource.Id = v.Id
			r.PrimarySource.Extension = v.Extension
		case "informationSource":
			var v CodeableReference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.InformationSource = &v
		case "location":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Location = &v
		case "site":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Site = &v
		case "route":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Route = &v
		case "doseQuantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.DoseQuantity = &v
		case "performer":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Immunization element", t)
			}
			for d.More() {
				var v ImmunizationPerformer
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Performer = append(r.Performer, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Immunization element", t)
			}
		case "note":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Immunization element", t)
			}
			for d.More() {
				var v Annotation
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Immunization element", t)
			}
		case "reason":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Immunization element", t)
			}
			for d.More() {
				var v CodeableReference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Reason = append(r.Reason, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Immunization element", t)
			}
		case "isSubpotent":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.IsSubpotent == nil {
				r.IsSubpotent = &Boolean{}
			}
			r.IsSubpotent.Value = v.Value
		case "_isSubpotent":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.IsSubpotent == nil {
				r.IsSubpotent = &Boolean{}
			}
			r.IsSubpotent.Id = v.Id
			r.IsSubpotent.Extension = v.Extension
		case "subpotentReason":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Immunization element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.SubpotentReason = append(r.SubpotentReason, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Immunization element", t)
			}
		case "programEligibility":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Immunization element", t)
			}
			for d.More() {
				var v ImmunizationProgramEligibility
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ProgramEligibility = append(r.ProgramEligibility, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Immunization element", t)
			}
		case "fundingSource":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.FundingSource = &v
		case "reaction":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Immunization element", t)
			}
			for d.More() {
				var v ImmunizationReaction
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Reaction = append(r.Reaction, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Immunization element", t)
			}
		case "protocolApplied":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Immunization element", t)
			}
			for d.More() {
				var v ImmunizationProtocolApplied
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ProtocolApplied = append(r.ProtocolApplied, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Immunization element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in Immunization", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in Immunization element", t)
	}
	return nil
}
func (r *ImmunizationPerformer) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ImmunizationPerformer element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ImmunizationPerformer element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationPerformer element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationPerformer element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationPerformer element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationPerformer element", t)
			}
		case "function":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Function = &v
		case "actor":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Actor = v
		default:
			return fmt.Errorf("invalid field: %s in ImmunizationPerformer", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ImmunizationPerformer element", t)
	}
	return nil
}
func (r *ImmunizationProgramEligibility) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ImmunizationProgramEligibility element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ImmunizationProgramEligibility element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationProgramEligibility element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationProgramEligibility element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationProgramEligibility element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationProgramEligibility element", t)
			}
		case "program":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Program = v
		case "programStatus":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ProgramStatus = v
		default:
			return fmt.Errorf("invalid field: %s in ImmunizationProgramEligibility", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ImmunizationProgramEligibility element", t)
	}
	return nil
}
func (r *ImmunizationReaction) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ImmunizationReaction element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ImmunizationReaction element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationReaction element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationReaction element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationReaction element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationReaction element", t)
			}
		case "date":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Date == nil {
				r.Date = &DateTime{}
			}
			r.Date.Value = v.Value
		case "_date":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Date == nil {
				r.Date = &DateTime{}
			}
			r.Date.Id = v.Id
			r.Date.Extension = v.Extension
		case "manifestation":
			var v CodeableReference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Manifestation = &v
		case "reported":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Reported == nil {
				r.Reported = &Boolean{}
			}
			r.Reported.Value = v.Value
		case "_reported":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Reported == nil {
				r.Reported = &Boolean{}
			}
			r.Reported.Id = v.Id
			r.Reported.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in ImmunizationReaction", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ImmunizationReaction element", t)
	}
	return nil
}
func (r *ImmunizationProtocolApplied) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ImmunizationProtocolApplied element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ImmunizationProtocolApplied element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationProtocolApplied element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationProtocolApplied element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationProtocolApplied element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationProtocolApplied element", t)
			}
		case "series":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Series == nil {
				r.Series = &String{}
			}
			r.Series.Value = v.Value
		case "_series":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Series == nil {
				r.Series = &String{}
			}
			r.Series.Id = v.Id
			r.Series.Extension = v.Extension
		case "authority":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Authority = &v
		case "targetDisease":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ImmunizationProtocolApplied element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.TargetDisease = append(r.TargetDisease, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ImmunizationProtocolApplied element", t)
			}
		case "doseNumber":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.DoseNumber.Value = v.Value
		case "_doseNumber":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.DoseNumber.Id = v.Id
			r.DoseNumber.Extension = v.Extension
		case "seriesDoses":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.SeriesDoses == nil {
				r.SeriesDoses = &String{}
			}
			r.SeriesDoses.Value = v.Value
		case "_seriesDoses":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.SeriesDoses == nil {
				r.SeriesDoses = &String{}
			}
			r.SeriesDoses.Id = v.Id
			r.SeriesDoses.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in ImmunizationProtocolApplied", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ImmunizationProtocolApplied element", t)
	}
	return nil
}
func (r Immunization) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "Immunization"
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BasedOn, xml.StartElement{Name: xml.Name{Local: "basedOn"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.StatusReason, xml.StartElement{Name: xml.Name{Local: "statusReason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.VaccineCode, xml.StartElement{Name: xml.Name{Local: "vaccineCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AdministeredProduct, xml.StartElement{Name: xml.Name{Local: "administeredProduct"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Manufacturer, xml.StartElement{Name: xml.Name{Local: "manufacturer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LotNumber, xml.StartElement{Name: xml.Name{Local: "lotNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ExpirationDate, xml.StartElement{Name: xml.Name{Local: "expirationDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Patient, xml.StartElement{Name: xml.Name{Local: "patient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Encounter, xml.StartElement{Name: xml.Name{Local: "encounter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SupportingInformation, xml.StartElement{Name: xml.Name{Local: "supportingInformation"}})
	if err != nil {
		return err
	}
	switch v := r.Occurrence.(type) {
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "occurrenceDateTime"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "occurrenceString"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.PrimarySource, xml.StartElement{Name: xml.Name{Local: "primarySource"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.InformationSource, xml.StartElement{Name: xml.Name{Local: "informationSource"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Location, xml.StartElement{Name: xml.Name{Local: "location"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Site, xml.StartElement{Name: xml.Name{Local: "site"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Route, xml.StartElement{Name: xml.Name{Local: "route"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DoseQuantity, xml.StartElement{Name: xml.Name{Local: "doseQuantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Performer, xml.StartElement{Name: xml.Name{Local: "performer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Reason, xml.StartElement{Name: xml.Name{Local: "reason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IsSubpotent, xml.StartElement{Name: xml.Name{Local: "isSubpotent"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubpotentReason, xml.StartElement{Name: xml.Name{Local: "subpotentReason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProgramEligibility, xml.StartElement{Name: xml.Name{Local: "programEligibility"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FundingSource, xml.StartElement{Name: xml.Name{Local: "fundingSource"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Reaction, xml.StartElement{Name: xml.Name{Local: "reaction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProtocolApplied, xml.StartElement{Name: xml.Name{Local: "protocolApplied"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ImmunizationPerformer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Function, xml.StartElement{Name: xml.Name{Local: "function"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Actor, xml.StartElement{Name: xml.Name{Local: "actor"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ImmunizationProgramEligibility) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Program, xml.StartElement{Name: xml.Name{Local: "program"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProgramStatus, xml.StartElement{Name: xml.Name{Local: "programStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ImmunizationReaction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Manifestation, xml.StartElement{Name: xml.Name{Local: "manifestation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Reported, xml.StartElement{Name: xml.Name{Local: "reported"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ImmunizationProtocolApplied) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Series, xml.StartElement{Name: xml.Name{Local: "series"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Authority, xml.StartElement{Name: xml.Name{Local: "authority"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TargetDisease, xml.StartElement{Name: xml.Name{Local: "targetDisease"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DoseNumber, xml.StartElement{Name: xml.Name{Local: "doseNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SeriesDoses, xml.StartElement{Name: xml.Name{Local: "seriesDoses"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Immunization) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "basedOn":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BasedOn = append(r.BasedOn, v)
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "statusReason":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StatusReason = &v
			case "vaccineCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.VaccineCode = v
			case "administeredProduct":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AdministeredProduct = &v
			case "manufacturer":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Manufacturer = &v
			case "lotNumber":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LotNumber = &v
			case "expirationDate":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ExpirationDate = &v
			case "patient":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Patient = v
			case "encounter":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Encounter = &v
			case "supportingInformation":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SupportingInformation = append(r.SupportingInformation, v)
			case "occurrenceDateTime":
				if r.Occurrence != nil {
					return fmt.Errorf("multiple values for field \"Occurrence\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Occurrence = v
			case "occurrenceString":
				if r.Occurrence != nil {
					return fmt.Errorf("multiple values for field \"Occurrence\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Occurrence = v
			case "primarySource":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PrimarySource = &v
			case "informationSource":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.InformationSource = &v
			case "location":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = &v
			case "site":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Site = &v
			case "route":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Route = &v
			case "doseQuantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DoseQuantity = &v
			case "performer":
				var v ImmunizationPerformer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Performer = append(r.Performer, v)
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			case "reason":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reason = append(r.Reason, v)
			case "isSubpotent":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IsSubpotent = &v
			case "subpotentReason":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubpotentReason = append(r.SubpotentReason, v)
			case "programEligibility":
				var v ImmunizationProgramEligibility
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProgramEligibility = append(r.ProgramEligibility, v)
			case "fundingSource":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FundingSource = &v
			case "reaction":
				var v ImmunizationReaction
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reaction = append(r.Reaction, v)
			case "protocolApplied":
				var v ImmunizationProtocolApplied
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProtocolApplied = append(r.ProtocolApplied, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ImmunizationPerformer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "function":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Function = &v
			case "actor":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Actor = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ImmunizationProgramEligibility) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "program":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Program = v
			case "programStatus":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProgramStatus = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ImmunizationReaction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "date":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = &v
			case "manifestation":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Manifestation = &v
			case "reported":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reported = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ImmunizationProtocolApplied) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "series":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Series = &v
			case "authority":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Authority = &v
			case "targetDisease":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TargetDisease = append(r.TargetDisease, v)
			case "doseNumber":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DoseNumber = v
			case "seriesDoses":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SeriesDoses = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Immunization) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "identifier") {
		for _, v := range r.Identifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "basedOn") {
		for _, v := range r.BasedOn {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "status") {
		children = append(children, r.Status)
	}
	if len(name) == 0 || slices.Contains(name, "statusReason") {
		if r.StatusReason != nil {
			children = append(children, *r.StatusReason)
		}
	}
	if len(name) == 0 || slices.Contains(name, "vaccineCode") {
		children = append(children, r.VaccineCode)
	}
	if len(name) == 0 || slices.Contains(name, "administeredProduct") {
		if r.AdministeredProduct != nil {
			children = append(children, *r.AdministeredProduct)
		}
	}
	if len(name) == 0 || slices.Contains(name, "manufacturer") {
		if r.Manufacturer != nil {
			children = append(children, *r.Manufacturer)
		}
	}
	if len(name) == 0 || slices.Contains(name, "lotNumber") {
		if r.LotNumber != nil {
			children = append(children, *r.LotNumber)
		}
	}
	if len(name) == 0 || slices.Contains(name, "expirationDate") {
		if r.ExpirationDate != nil {
			children = append(children, *r.ExpirationDate)
		}
	}
	if len(name) == 0 || slices.Contains(name, "patient") {
		children = append(children, r.Patient)
	}
	if len(name) == 0 || slices.Contains(name, "encounter") {
		if r.Encounter != nil {
			children = append(children, *r.Encounter)
		}
	}
	if len(name) == 0 || slices.Contains(name, "supportingInformation") {
		for _, v := range r.SupportingInformation {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "occurrence") {
		children = append(children, r.Occurrence)
	}
	if len(name) == 0 || slices.Contains(name, "primarySource") {
		if r.PrimarySource != nil {
			children = append(children, *r.PrimarySource)
		}
	}
	if len(name) == 0 || slices.Contains(name, "informationSource") {
		if r.InformationSource != nil {
			children = append(children, *r.InformationSource)
		}
	}
	if len(name) == 0 || slices.Contains(name, "location") {
		if r.Location != nil {
			children = append(children, *r.Location)
		}
	}
	if len(name) == 0 || slices.Contains(name, "site") {
		if r.Site != nil {
			children = append(children, *r.Site)
		}
	}
	if len(name) == 0 || slices.Contains(name, "route") {
		if r.Route != nil {
			children = append(children, *r.Route)
		}
	}
	if len(name) == 0 || slices.Contains(name, "doseQuantity") {
		if r.DoseQuantity != nil {
			children = append(children, *r.DoseQuantity)
		}
	}
	if len(name) == 0 || slices.Contains(name, "performer") {
		for _, v := range r.Performer {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "note") {
		for _, v := range r.Note {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "reason") {
		for _, v := range r.Reason {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "isSubpotent") {
		if r.IsSubpotent != nil {
			children = append(children, *r.IsSubpotent)
		}
	}
	if len(name) == 0 || slices.Contains(name, "subpotentReason") {
		for _, v := range r.SubpotentReason {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "programEligibility") {
		for _, v := range r.ProgramEligibility {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "fundingSource") {
		if r.FundingSource != nil {
			children = append(children, *r.FundingSource)
		}
	}
	if len(name) == 0 || slices.Contains(name, "reaction") {
		for _, v := range r.Reaction {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "protocolApplied") {
		for _, v := range r.ProtocolApplied {
			children = append(children, v)
		}
	}
	return children
}
func (r Immunization) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert Immunization to Boolean")
}
func (r Immunization) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert Immunization to String")
}
func (r Immunization) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert Immunization to Integer")
}
func (r Immunization) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert Immunization to Decimal")
}
func (r Immunization) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert Immunization to Date")
}
func (r Immunization) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert Immunization to Time")
}
func (r Immunization) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert Immunization to DateTime")
}
func (r Immunization) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert Immunization to Quantity")
}
func (r Immunization) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.Id",
		}, {
			Name: "Meta",
			Type: "FHIR.Meta",
		}, {
			Name: "ImplicitRules",
			Type: "FHIR.Uri",
		}, {
			Name: "Language",
			Type: "FHIR.Code",
		}, {
			Name: "Text",
			Type: "FHIR.Narrative",
		}, {
			Name: "Contained",
			Type: "List<FHIR.>",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Identifier",
			Type: "List<FHIR.Identifier>",
		}, {
			Name: "BasedOn",
			Type: "List<FHIR.Reference>",
		}, {
			Name: "Status",
			Type: "FHIR.Code",
		}, {
			Name: "StatusReason",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "VaccineCode",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "AdministeredProduct",
			Type: "FHIR.CodeableReference",
		}, {
			Name: "Manufacturer",
			Type: "FHIR.CodeableReference",
		}, {
			Name: "LotNumber",
			Type: "FHIR.String",
		}, {
			Name: "ExpirationDate",
			Type: "FHIR.Date",
		}, {
			Name: "Patient",
			Type: "FHIR.Reference",
		}, {
			Name: "Encounter",
			Type: "FHIR.Reference",
		}, {
			Name: "SupportingInformation",
			Type: "List<FHIR.Reference>",
		}, {
			Name: "Occurrence",
			Type: "FHIR.PrimitiveElement",
		}, {
			Name: "PrimarySource",
			Type: "FHIR.Boolean",
		}, {
			Name: "InformationSource",
			Type: "FHIR.CodeableReference",
		}, {
			Name: "Location",
			Type: "FHIR.Reference",
		}, {
			Name: "Site",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Route",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "DoseQuantity",
			Type: "FHIR.Quantity",
		}, {
			Name: "Performer",
			Type: "List<FHIR.ImmunizationPerformer>",
		}, {
			Name: "Note",
			Type: "List<FHIR.Annotation>",
		}, {
			Name: "Reason",
			Type: "List<FHIR.CodeableReference>",
		}, {
			Name: "IsSubpotent",
			Type: "FHIR.Boolean",
		}, {
			Name: "SubpotentReason",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "ProgramEligibility",
			Type: "List<FHIR.ImmunizationProgramEligibility>",
		}, {
			Name: "FundingSource",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Reaction",
			Type: "List<FHIR.ImmunizationReaction>",
		}, {
			Name: "ProtocolApplied",
			Type: "List<FHIR.ImmunizationProtocolApplied>",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "Immunization",
			Namespace: "FHIR",
		},
	}
}
func (r ImmunizationPerformer) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "function") {
		if r.Function != nil {
			children = append(children, *r.Function)
		}
	}
	if len(name) == 0 || slices.Contains(name, "actor") {
		children = append(children, r.Actor)
	}
	return children
}
func (r ImmunizationPerformer) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert ImmunizationPerformer to Boolean")
}
func (r ImmunizationPerformer) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert ImmunizationPerformer to String")
}
func (r ImmunizationPerformer) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert ImmunizationPerformer to Integer")
}
func (r ImmunizationPerformer) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert ImmunizationPerformer to Decimal")
}
func (r ImmunizationPerformer) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert ImmunizationPerformer to Date")
}
func (r ImmunizationPerformer) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert ImmunizationPerformer to Time")
}
func (r ImmunizationPerformer) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert ImmunizationPerformer to DateTime")
}
func (r ImmunizationPerformer) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert ImmunizationPerformer to Quantity")
}
func (r ImmunizationPerformer) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Function",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Actor",
			Type: "FHIR.Reference",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ImmunizationPerformer",
			Namespace: "FHIR",
		},
	}
}
func (r ImmunizationProgramEligibility) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "program") {
		children = append(children, r.Program)
	}
	if len(name) == 0 || slices.Contains(name, "programStatus") {
		children = append(children, r.ProgramStatus)
	}
	return children
}
func (r ImmunizationProgramEligibility) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert ImmunizationProgramEligibility to Boolean")
}
func (r ImmunizationProgramEligibility) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert ImmunizationProgramEligibility to String")
}
func (r ImmunizationProgramEligibility) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert ImmunizationProgramEligibility to Integer")
}
func (r ImmunizationProgramEligibility) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert ImmunizationProgramEligibility to Decimal")
}
func (r ImmunizationProgramEligibility) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert ImmunizationProgramEligibility to Date")
}
func (r ImmunizationProgramEligibility) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert ImmunizationProgramEligibility to Time")
}
func (r ImmunizationProgramEligibility) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert ImmunizationProgramEligibility to DateTime")
}
func (r ImmunizationProgramEligibility) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert ImmunizationProgramEligibility to Quantity")
}
func (r ImmunizationProgramEligibility) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Program",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "ProgramStatus",
			Type: "FHIR.CodeableConcept",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ImmunizationProgramEligibility",
			Namespace: "FHIR",
		},
	}
}
func (r ImmunizationReaction) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "date") {
		if r.Date != nil {
			children = append(children, *r.Date)
		}
	}
	if len(name) == 0 || slices.Contains(name, "manifestation") {
		if r.Manifestation != nil {
			children = append(children, *r.Manifestation)
		}
	}
	if len(name) == 0 || slices.Contains(name, "reported") {
		if r.Reported != nil {
			children = append(children, *r.Reported)
		}
	}
	return children
}
func (r ImmunizationReaction) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert ImmunizationReaction to Boolean")
}
func (r ImmunizationReaction) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert ImmunizationReaction to String")
}
func (r ImmunizationReaction) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert ImmunizationReaction to Integer")
}
func (r ImmunizationReaction) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert ImmunizationReaction to Decimal")
}
func (r ImmunizationReaction) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert ImmunizationReaction to Date")
}
func (r ImmunizationReaction) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert ImmunizationReaction to Time")
}
func (r ImmunizationReaction) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert ImmunizationReaction to DateTime")
}
func (r ImmunizationReaction) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert ImmunizationReaction to Quantity")
}
func (r ImmunizationReaction) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Date",
			Type: "FHIR.DateTime",
		}, {
			Name: "Manifestation",
			Type: "FHIR.CodeableReference",
		}, {
			Name: "Reported",
			Type: "FHIR.Boolean",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ImmunizationReaction",
			Namespace: "FHIR",
		},
	}
}
func (r ImmunizationProtocolApplied) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "series") {
		if r.Series != nil {
			children = append(children, *r.Series)
		}
	}
	if len(name) == 0 || slices.Contains(name, "authority") {
		if r.Authority != nil {
			children = append(children, *r.Authority)
		}
	}
	if len(name) == 0 || slices.Contains(name, "targetDisease") {
		for _, v := range r.TargetDisease {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "doseNumber") {
		children = append(children, r.DoseNumber)
	}
	if len(name) == 0 || slices.Contains(name, "seriesDoses") {
		if r.SeriesDoses != nil {
			children = append(children, *r.SeriesDoses)
		}
	}
	return children
}
func (r ImmunizationProtocolApplied) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert ImmunizationProtocolApplied to Boolean")
}
func (r ImmunizationProtocolApplied) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert ImmunizationProtocolApplied to String")
}
func (r ImmunizationProtocolApplied) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert ImmunizationProtocolApplied to Integer")
}
func (r ImmunizationProtocolApplied) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert ImmunizationProtocolApplied to Decimal")
}
func (r ImmunizationProtocolApplied) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert ImmunizationProtocolApplied to Date")
}
func (r ImmunizationProtocolApplied) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert ImmunizationProtocolApplied to Time")
}
func (r ImmunizationProtocolApplied) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert ImmunizationProtocolApplied to DateTime")
}
func (r ImmunizationProtocolApplied) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert ImmunizationProtocolApplied to Quantity")
}
func (r ImmunizationProtocolApplied) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Series",
			Type: "FHIR.String",
		}, {
			Name: "Authority",
			Type: "FHIR.Reference",
		}, {
			Name: "TargetDisease",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "DoseNumber",
			Type: "FHIR.String",
		}, {
			Name: "SeriesDoses",
			Type: "FHIR.String",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ImmunizationProtocolApplied",
			Namespace: "FHIR",
		},
	}
}
