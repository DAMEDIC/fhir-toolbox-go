package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// Indicates that a medication product is to be or has been dispensed for a named person/patient.  This includes a description of the medication product (supply) provided and the instructions for administering the medication.  The medication dispense is the result of a pharmacy system responding to a medication order.
type MedicationDispense struct {
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
	// Identifiers associated with this Medication Dispense that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate. They are business identifiers assigned to this resource by the performer or other systems and remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// The procedure that trigger the dispense.
	PartOf []Reference
	// A code specifying the state of the set of dispense events.
	Status Code
	// Indicates the reason why a dispense was not performed.
	StatusReason isMedicationDispenseStatusReason
	// Indicates the type of medication dispense (for example, where the medication is expected to be consumed or administered (i.e. inpatient or outpatient)).
	Category *CodeableConcept
	// Identifies the medication being administered. This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
	Medication isMedicationDispenseMedication
	// A link to a resource representing the person or the group to whom the medication will be given.
	Subject *Reference
	// The encounter or episode of care that establishes the context for this event.
	Context *Reference
	// Additional information that supports the medication being dispensed.
	SupportingInformation []Reference
	// Indicates who or what performed the event.
	Performer []MedicationDispensePerformer
	// The principal physical location where the dispense was performed.
	Location *Reference
	// Indicates the medication order that is being dispensed against.
	AuthorizingPrescription []Reference
	// Indicates the type of dispensing event that is performed. For example, Trial Fill, Completion of Trial, Partial Fill, Emergency Fill, Samples, etc.
	Type *CodeableConcept
	// The amount of medication that has been dispensed. Includes unit of measure.
	Quantity *Quantity
	// The amount of medication expressed as a timing amount.
	DaysSupply *Quantity
	// The time when the dispensed product was packaged and reviewed.
	WhenPrepared *DateTime
	// The time the dispensed product was provided to the patient or their representative.
	WhenHandedOver *DateTime
	// Identification of the facility/location where the medication was shipped to, as part of the dispense event.
	Destination *Reference
	// Identifies the person who picked up the medication.  This will usually be a patient or their caregiver, but some cases exist where it can be a healthcare professional.
	Receiver []Reference
	// Extra information about the dispense that could not be conveyed in the other attributes.
	Note []Annotation
	// Indicates how the medication is to be used by the patient.
	DosageInstruction []Dosage
	// Indicates whether or not substitution was made as part of the dispense.  In some cases, substitution will be expected but does not happen, in other cases substitution is not expected but does happen.  This block explains what substitution did or did not happen and why.  If nothing is specified, substitution was not done.
	Substitution *MedicationDispenseSubstitution
	// Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions for a patient; e.g. drug-drug interaction, duplicate therapy, dosage alert etc.
	DetectedIssue []Reference
	// A summary of the events of interest that have occurred, such as when the dispense was verified.
	EventHistory []Reference
}

func (r MedicationDispense) ResourceType() string {
	return "MedicationDispense"
}
func (r MedicationDispense) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isMedicationDispenseStatusReason interface {
	isMedicationDispenseStatusReason()
}

func (r CodeableConcept) isMedicationDispenseStatusReason() {}
func (r Reference) isMedicationDispenseStatusReason()       {}

type isMedicationDispenseMedication interface {
	isMedicationDispenseMedication()
}

func (r CodeableConcept) isMedicationDispenseMedication() {}
func (r Reference) isMedicationDispenseMedication()       {}

type jsonMedicationDispense struct {
	ResourceType                   string                          `json:"resourceType"`
	Id                             *Id                             `json:"id,omitempty"`
	IdPrimitiveElement             *primitiveElement               `json:"_id,omitempty"`
	Meta                           *Meta                           `json:"meta,omitempty"`
	ImplicitRules                  *Uri                            `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement  *primitiveElement               `json:"_implicitRules,omitempty"`
	Language                       *Code                           `json:"language,omitempty"`
	LanguagePrimitiveElement       *primitiveElement               `json:"_language,omitempty"`
	Text                           *Narrative                      `json:"text,omitempty"`
	Contained                      []ContainedResource             `json:"contained,omitempty"`
	Extension                      []Extension                     `json:"extension,omitempty"`
	ModifierExtension              []Extension                     `json:"modifierExtension,omitempty"`
	Identifier                     []Identifier                    `json:"identifier,omitempty"`
	PartOf                         []Reference                     `json:"partOf,omitempty"`
	Status                         Code                            `json:"status,omitempty"`
	StatusPrimitiveElement         *primitiveElement               `json:"_status,omitempty"`
	StatusReasonCodeableConcept    *CodeableConcept                `json:"statusReasonCodeableConcept,omitempty"`
	StatusReasonReference          *Reference                      `json:"statusReasonReference,omitempty"`
	Category                       *CodeableConcept                `json:"category,omitempty"`
	MedicationCodeableConcept      *CodeableConcept                `json:"medicationCodeableConcept,omitempty"`
	MedicationReference            *Reference                      `json:"medicationReference,omitempty"`
	Subject                        *Reference                      `json:"subject,omitempty"`
	Context                        *Reference                      `json:"context,omitempty"`
	SupportingInformation          []Reference                     `json:"supportingInformation,omitempty"`
	Performer                      []MedicationDispensePerformer   `json:"performer,omitempty"`
	Location                       *Reference                      `json:"location,omitempty"`
	AuthorizingPrescription        []Reference                     `json:"authorizingPrescription,omitempty"`
	Type                           *CodeableConcept                `json:"type,omitempty"`
	Quantity                       *Quantity                       `json:"quantity,omitempty"`
	DaysSupply                     *Quantity                       `json:"daysSupply,omitempty"`
	WhenPrepared                   *DateTime                       `json:"whenPrepared,omitempty"`
	WhenPreparedPrimitiveElement   *primitiveElement               `json:"_whenPrepared,omitempty"`
	WhenHandedOver                 *DateTime                       `json:"whenHandedOver,omitempty"`
	WhenHandedOverPrimitiveElement *primitiveElement               `json:"_whenHandedOver,omitempty"`
	Destination                    *Reference                      `json:"destination,omitempty"`
	Receiver                       []Reference                     `json:"receiver,omitempty"`
	Note                           []Annotation                    `json:"note,omitempty"`
	DosageInstruction              []Dosage                        `json:"dosageInstruction,omitempty"`
	Substitution                   *MedicationDispenseSubstitution `json:"substitution,omitempty"`
	DetectedIssue                  []Reference                     `json:"detectedIssue,omitempty"`
	EventHistory                   []Reference                     `json:"eventHistory,omitempty"`
}

func (r MedicationDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationDispense) marshalJSON() jsonMedicationDispense {
	m := jsonMedicationDispense{}
	m.ResourceType = "MedicationDispense"
	if r.Id != nil && r.Id.Value != nil {
		m.Id = r.Id
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		m.ImplicitRules = r.ImplicitRules
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	if r.Language != nil && r.Language.Value != nil {
		m.Language = r.Language
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.Text = r.Text
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.PartOf = r.PartOf
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	switch v := r.StatusReason.(type) {
	case CodeableConcept:
		m.StatusReasonCodeableConcept = &v
	case *CodeableConcept:
		m.StatusReasonCodeableConcept = v
	case Reference:
		m.StatusReasonReference = &v
	case *Reference:
		m.StatusReasonReference = v
	}
	m.Category = r.Category
	switch v := r.Medication.(type) {
	case CodeableConcept:
		m.MedicationCodeableConcept = &v
	case *CodeableConcept:
		m.MedicationCodeableConcept = v
	case Reference:
		m.MedicationReference = &v
	case *Reference:
		m.MedicationReference = v
	}
	m.Subject = r.Subject
	m.Context = r.Context
	m.SupportingInformation = r.SupportingInformation
	m.Performer = r.Performer
	m.Location = r.Location
	m.AuthorizingPrescription = r.AuthorizingPrescription
	m.Type = r.Type
	m.Quantity = r.Quantity
	m.DaysSupply = r.DaysSupply
	if r.WhenPrepared != nil && r.WhenPrepared.Value != nil {
		m.WhenPrepared = r.WhenPrepared
	}
	if r.WhenPrepared != nil && (r.WhenPrepared.Id != nil || r.WhenPrepared.Extension != nil) {
		m.WhenPreparedPrimitiveElement = &primitiveElement{Id: r.WhenPrepared.Id, Extension: r.WhenPrepared.Extension}
	}
	if r.WhenHandedOver != nil && r.WhenHandedOver.Value != nil {
		m.WhenHandedOver = r.WhenHandedOver
	}
	if r.WhenHandedOver != nil && (r.WhenHandedOver.Id != nil || r.WhenHandedOver.Extension != nil) {
		m.WhenHandedOverPrimitiveElement = &primitiveElement{Id: r.WhenHandedOver.Id, Extension: r.WhenHandedOver.Extension}
	}
	m.Destination = r.Destination
	m.Receiver = r.Receiver
	m.Note = r.Note
	m.DosageInstruction = r.DosageInstruction
	m.Substitution = r.Substitution
	m.DetectedIssue = r.DetectedIssue
	m.EventHistory = r.EventHistory
	return m
}
func (r *MedicationDispense) UnmarshalJSON(b []byte) error {
	var m jsonMedicationDispense
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationDispense) unmarshalJSON(m jsonMedicationDispense) error {
	r.Id = m.Id
	if m.IdPrimitiveElement != nil {
		if r.Id == nil {
			r.Id = &Id{}
		}
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		if r.ImplicitRules == nil {
			r.ImplicitRules = &Uri{}
		}
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		if r.Language == nil {
			r.Language = &Code{}
		}
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Text = m.Text
	r.Contained = make([]model.Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.Resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.PartOf = m.PartOf
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	if m.StatusReasonCodeableConcept != nil {
		if r.StatusReason != nil {
			return fmt.Errorf("multiple values for field \"StatusReason\"")
		}
		v := m.StatusReasonCodeableConcept
		r.StatusReason = v
	}
	if m.StatusReasonReference != nil {
		if r.StatusReason != nil {
			return fmt.Errorf("multiple values for field \"StatusReason\"")
		}
		v := m.StatusReasonReference
		r.StatusReason = v
	}
	r.Category = m.Category
	if m.MedicationCodeableConcept != nil {
		if r.Medication != nil {
			return fmt.Errorf("multiple values for field \"Medication\"")
		}
		v := m.MedicationCodeableConcept
		r.Medication = v
	}
	if m.MedicationReference != nil {
		if r.Medication != nil {
			return fmt.Errorf("multiple values for field \"Medication\"")
		}
		v := m.MedicationReference
		r.Medication = v
	}
	r.Subject = m.Subject
	r.Context = m.Context
	r.SupportingInformation = m.SupportingInformation
	r.Performer = m.Performer
	r.Location = m.Location
	r.AuthorizingPrescription = m.AuthorizingPrescription
	r.Type = m.Type
	r.Quantity = m.Quantity
	r.DaysSupply = m.DaysSupply
	r.WhenPrepared = m.WhenPrepared
	if m.WhenPreparedPrimitiveElement != nil {
		if r.WhenPrepared == nil {
			r.WhenPrepared = &DateTime{}
		}
		r.WhenPrepared.Id = m.WhenPreparedPrimitiveElement.Id
		r.WhenPrepared.Extension = m.WhenPreparedPrimitiveElement.Extension
	}
	r.WhenHandedOver = m.WhenHandedOver
	if m.WhenHandedOverPrimitiveElement != nil {
		if r.WhenHandedOver == nil {
			r.WhenHandedOver = &DateTime{}
		}
		r.WhenHandedOver.Id = m.WhenHandedOverPrimitiveElement.Id
		r.WhenHandedOver.Extension = m.WhenHandedOverPrimitiveElement.Extension
	}
	r.Destination = m.Destination
	r.Receiver = m.Receiver
	r.Note = m.Note
	r.DosageInstruction = m.DosageInstruction
	r.Substitution = m.Substitution
	r.DetectedIssue = m.DetectedIssue
	r.EventHistory = m.EventHistory
	return nil
}
func (r MedicationDispense) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates who or what performed the event.
type MedicationDispensePerformer struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Distinguishes the type of performer in the dispense.  For example, date enterer, packager, final checker.
	Function *CodeableConcept
	// The device, practitioner, etc. who performed the action.  It should be assumed that the actor is the dispenser of the medication.
	Actor Reference
}
type jsonMedicationDispensePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor,omitempty"`
}

func (r MedicationDispensePerformer) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationDispensePerformer) marshalJSON() jsonMedicationDispensePerformer {
	m := jsonMedicationDispensePerformer{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Function = r.Function
	m.Actor = r.Actor
	return m
}
func (r *MedicationDispensePerformer) UnmarshalJSON(b []byte) error {
	var m jsonMedicationDispensePerformer
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationDispensePerformer) unmarshalJSON(m jsonMedicationDispensePerformer) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Function = m.Function
	r.Actor = m.Actor
	return nil
}
func (r MedicationDispensePerformer) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates whether or not substitution was made as part of the dispense.  In some cases, substitution will be expected but does not happen, in other cases substitution is not expected but does happen.  This block explains what substitution did or did not happen and why.  If nothing is specified, substitution was not done.
type MedicationDispenseSubstitution struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// True if the dispenser dispensed a different drug or product from what was prescribed.
	WasSubstituted Boolean
	// A code signifying whether a different drug was dispensed from what was prescribed.
	Type *CodeableConcept
	// Indicates the reason for the substitution (or lack of substitution) from what was prescribed.
	Reason []CodeableConcept
	// The person or organization that has primary responsibility for the substitution.
	ResponsibleParty []Reference
}
type jsonMedicationDispenseSubstitution struct {
	Id                             *string           `json:"id,omitempty"`
	Extension                      []Extension       `json:"extension,omitempty"`
	ModifierExtension              []Extension       `json:"modifierExtension,omitempty"`
	WasSubstituted                 Boolean           `json:"wasSubstituted,omitempty"`
	WasSubstitutedPrimitiveElement *primitiveElement `json:"_wasSubstituted,omitempty"`
	Type                           *CodeableConcept  `json:"type,omitempty"`
	Reason                         []CodeableConcept `json:"reason,omitempty"`
	ResponsibleParty               []Reference       `json:"responsibleParty,omitempty"`
}

func (r MedicationDispenseSubstitution) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationDispenseSubstitution) marshalJSON() jsonMedicationDispenseSubstitution {
	m := jsonMedicationDispenseSubstitution{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.WasSubstituted.Value != nil {
		m.WasSubstituted = r.WasSubstituted
	}
	if r.WasSubstituted.Id != nil || r.WasSubstituted.Extension != nil {
		m.WasSubstitutedPrimitiveElement = &primitiveElement{Id: r.WasSubstituted.Id, Extension: r.WasSubstituted.Extension}
	}
	m.Type = r.Type
	m.Reason = r.Reason
	m.ResponsibleParty = r.ResponsibleParty
	return m
}
func (r *MedicationDispenseSubstitution) UnmarshalJSON(b []byte) error {
	var m jsonMedicationDispenseSubstitution
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationDispenseSubstitution) unmarshalJSON(m jsonMedicationDispenseSubstitution) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.WasSubstituted = m.WasSubstituted
	if m.WasSubstitutedPrimitiveElement != nil {
		r.WasSubstituted.Id = m.WasSubstitutedPrimitiveElement.Id
		r.WasSubstituted.Extension = m.WasSubstitutedPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Reason = m.Reason
	r.ResponsibleParty = m.ResponsibleParty
	return nil
}
func (r MedicationDispenseSubstitution) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
