package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// An order or request for both supply of the medication and the instructions for administration of the medication to a patient. The resource is called "MedicationRequest" rather than "MedicationPrescription" or "MedicationOrder" to generalize the use across inpatient and outpatient settings, including care plans, etc., and to harmonize with workflow patterns.
type MedicationRequest struct {
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
	// Identifiers associated with this medication request that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate. They are business identifiers assigned to this resource by the performer or other systems and remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// A code specifying the current state of the order.  Generally, this will be active or completed state.
	Status Code
	// Captures the reason for the current state of the MedicationRequest.
	StatusReason *CodeableConcept
	// Whether the request is a proposal, plan, or an original order.
	Intent Code
	// Indicates the type of medication request (for example, where the medication is expected to be consumed or administered (i.e. inpatient or outpatient)).
	Category []CodeableConcept
	// Indicates how quickly the Medication Request should be addressed with respect to other requests.
	Priority *Code
	// If true indicates that the provider is asking for the medication request not to occur.
	DoNotPerform *Boolean
	// Indicates if this record was captured as a secondary 'reported' record rather than as an original primary source-of-truth record.  It may also indicate the source of the report.
	Reported isMedicationRequestReported
	// Identifies the medication being requested. This is a link to a resource that represents the medication which may be the details of the medication or simply an attribute carrying a code that identifies the medication from a known list of medications.
	Medication isMedicationRequestMedication
	// A link to a resource representing the person or set of individuals to whom the medication will be given.
	Subject Reference
	// The Encounter during which this [x] was created or to which the creation of this record is tightly associated.
	Encounter *Reference
	// Include additional information (for example, patient height and weight) that supports the ordering of the medication.
	SupportingInformation []Reference
	// The date (and perhaps time) when the prescription was initially written or authored on.
	AuthoredOn *DateTime
	// The individual, organization, or device that initiated the request and has responsibility for its activation.
	Requester *Reference
	// The specified desired performer of the medication treatment (e.g. the performer of the medication administration).
	Performer *Reference
	// Indicates the type of performer of the administration of the medication.
	PerformerType *CodeableConcept
	// The person who entered the order on behalf of another individual for example in the case of a verbal or a telephone order.
	Recorder *Reference
	// The reason or the indication for ordering or not ordering the medication.
	ReasonCode []CodeableConcept
	// Condition or observation that supports why the medication was ordered.
	ReasonReference []Reference
	// The URL pointing to a protocol, guideline, orderset, or other definition that is adhered to in whole or in part by this MedicationRequest.
	InstantiatesCanonical []Canonical
	// The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this MedicationRequest.
	InstantiatesUri []Uri
	// A plan or request that is fulfilled in whole or in part by this medication request.
	BasedOn []Reference
	// A shared identifier common to all requests that were authorized more or less simultaneously by a single author, representing the identifier of the requisition or prescription.
	GroupIdentifier *Identifier
	// The description of the overall patte3rn of the administration of the medication to the patient.
	CourseOfTherapyType *CodeableConcept
	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be required for delivering the requested service.
	Insurance []Reference
	// Extra information about the prescription that could not be conveyed by the other attributes.
	Note []Annotation
	// Indicates how the medication is to be used by the patient.
	DosageInstruction []Dosage
	// Indicates the specific details for the dispense or medication supply part of a medication request (also known as a Medication Prescription or Medication Order).  Note that this information is not always sent with the order.  There may be in some settings (e.g. hospitals) institutional or system support for completing the dispense details in the pharmacy department.
	DispenseRequest *MedicationRequestDispenseRequest
	// Indicates whether or not substitution can or should be part of the dispense. In some cases, substitution must happen, in other cases substitution must not happen. This block explains the prescriber's intent. If nothing is specified substitution may be done.
	Substitution *MedicationRequestSubstitution
	// A link to a resource representing an earlier order related order or prescription.
	PriorPrescription *Reference
	// Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions for a patient; e.g. Drug-drug interaction, duplicate therapy, dosage alert etc.
	DetectedIssue []Reference
	// Links to Provenance records for past versions of this resource or fulfilling request or event resources that identify key state transitions or updates that are likely to be relevant to a user looking at the current version of the resource.
	EventHistory []Reference
}

func (r MedicationRequest) ResourceType() string {
	return "MedicationRequest"
}
func (r MedicationRequest) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type isMedicationRequestReported interface {
	isMedicationRequestReported()
}

func (r Boolean) isMedicationRequestReported()   {}
func (r Reference) isMedicationRequestReported() {}

type isMedicationRequestMedication interface {
	isMedicationRequestMedication()
}

func (r CodeableConcept) isMedicationRequestMedication() {}
func (r Reference) isMedicationRequestMedication()       {}

type jsonMedicationRequest struct {
	ResourceType                          string                            `json:"resourceType"`
	Id                                    *Id                               `json:"id,omitempty"`
	IdPrimitiveElement                    *primitiveElement                 `json:"_id,omitempty"`
	Meta                                  *Meta                             `json:"meta,omitempty"`
	ImplicitRules                         *Uri                              `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement         *primitiveElement                 `json:"_implicitRules,omitempty"`
	Language                              *Code                             `json:"language,omitempty"`
	LanguagePrimitiveElement              *primitiveElement                 `json:"_language,omitempty"`
	Text                                  *Narrative                        `json:"text,omitempty"`
	Contained                             []containedResource               `json:"contained,omitempty"`
	Extension                             []Extension                       `json:"extension,omitempty"`
	ModifierExtension                     []Extension                       `json:"modifierExtension,omitempty"`
	Identifier                            []Identifier                      `json:"identifier,omitempty"`
	Status                                Code                              `json:"status,omitempty"`
	StatusPrimitiveElement                *primitiveElement                 `json:"_status,omitempty"`
	StatusReason                          *CodeableConcept                  `json:"statusReason,omitempty"`
	Intent                                Code                              `json:"intent,omitempty"`
	IntentPrimitiveElement                *primitiveElement                 `json:"_intent,omitempty"`
	Category                              []CodeableConcept                 `json:"category,omitempty"`
	Priority                              *Code                             `json:"priority,omitempty"`
	PriorityPrimitiveElement              *primitiveElement                 `json:"_priority,omitempty"`
	DoNotPerform                          *Boolean                          `json:"doNotPerform,omitempty"`
	DoNotPerformPrimitiveElement          *primitiveElement                 `json:"_doNotPerform,omitempty"`
	ReportedBoolean                       *Boolean                          `json:"reportedBoolean,omitempty"`
	ReportedBooleanPrimitiveElement       *primitiveElement                 `json:"_reportedBoolean,omitempty"`
	ReportedReference                     *Reference                        `json:"reportedReference,omitempty"`
	MedicationCodeableConcept             *CodeableConcept                  `json:"medicationCodeableConcept,omitempty"`
	MedicationReference                   *Reference                        `json:"medicationReference,omitempty"`
	Subject                               Reference                         `json:"subject,omitempty"`
	Encounter                             *Reference                        `json:"encounter,omitempty"`
	SupportingInformation                 []Reference                       `json:"supportingInformation,omitempty"`
	AuthoredOn                            *DateTime                         `json:"authoredOn,omitempty"`
	AuthoredOnPrimitiveElement            *primitiveElement                 `json:"_authoredOn,omitempty"`
	Requester                             *Reference                        `json:"requester,omitempty"`
	Performer                             *Reference                        `json:"performer,omitempty"`
	PerformerType                         *CodeableConcept                  `json:"performerType,omitempty"`
	Recorder                              *Reference                        `json:"recorder,omitempty"`
	ReasonCode                            []CodeableConcept                 `json:"reasonCode,omitempty"`
	ReasonReference                       []Reference                       `json:"reasonReference,omitempty"`
	InstantiatesCanonical                 []Canonical                       `json:"instantiatesCanonical,omitempty"`
	InstantiatesCanonicalPrimitiveElement []*primitiveElement               `json:"_instantiatesCanonical,omitempty"`
	InstantiatesUri                       []Uri                             `json:"instantiatesUri,omitempty"`
	InstantiatesUriPrimitiveElement       []*primitiveElement               `json:"_instantiatesUri,omitempty"`
	BasedOn                               []Reference                       `json:"basedOn,omitempty"`
	GroupIdentifier                       *Identifier                       `json:"groupIdentifier,omitempty"`
	CourseOfTherapyType                   *CodeableConcept                  `json:"courseOfTherapyType,omitempty"`
	Insurance                             []Reference                       `json:"insurance,omitempty"`
	Note                                  []Annotation                      `json:"note,omitempty"`
	DosageInstruction                     []Dosage                          `json:"dosageInstruction,omitempty"`
	DispenseRequest                       *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty"`
	Substitution                          *MedicationRequestSubstitution    `json:"substitution,omitempty"`
	PriorPrescription                     *Reference                        `json:"priorPrescription,omitempty"`
	DetectedIssue                         []Reference                       `json:"detectedIssue,omitempty"`
	EventHistory                          []Reference                       `json:"eventHistory,omitempty"`
}

func (r MedicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationRequest) marshalJSON() jsonMedicationRequest {
	m := jsonMedicationRequest{}
	m.ResourceType = "MedicationRequest"
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
	m.StatusReason = r.StatusReason
	m.Intent = r.Intent
	if r.Intent.Id != nil || r.Intent.Extension != nil {
		m.IntentPrimitiveElement = &primitiveElement{Id: r.Intent.Id, Extension: r.Intent.Extension}
	}
	m.Category = r.Category
	m.Priority = r.Priority
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		m.PriorityPrimitiveElement = &primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
	}
	m.DoNotPerform = r.DoNotPerform
	if r.DoNotPerform != nil && (r.DoNotPerform.Id != nil || r.DoNotPerform.Extension != nil) {
		m.DoNotPerformPrimitiveElement = &primitiveElement{Id: r.DoNotPerform.Id, Extension: r.DoNotPerform.Extension}
	}
	switch v := r.Reported.(type) {
	case Boolean:
		m.ReportedBoolean = &v
		if v.Id != nil || v.Extension != nil {
			m.ReportedBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		m.ReportedBoolean = v
		if v.Id != nil || v.Extension != nil {
			m.ReportedBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Reference:
		m.ReportedReference = &v
	case *Reference:
		m.ReportedReference = v
	}
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
	m.Encounter = r.Encounter
	m.SupportingInformation = r.SupportingInformation
	m.AuthoredOn = r.AuthoredOn
	if r.AuthoredOn != nil && (r.AuthoredOn.Id != nil || r.AuthoredOn.Extension != nil) {
		m.AuthoredOnPrimitiveElement = &primitiveElement{Id: r.AuthoredOn.Id, Extension: r.AuthoredOn.Extension}
	}
	m.Requester = r.Requester
	m.Performer = r.Performer
	m.PerformerType = r.PerformerType
	m.Recorder = r.Recorder
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.InstantiatesCanonical = r.InstantiatesCanonical
	anyInstantiatesCanonicalIdOrExtension := false
	for _, e := range r.InstantiatesCanonical {
		if e.Id != nil || e.Extension != nil {
			anyInstantiatesCanonicalIdOrExtension = true
			break
		}
	}
	if anyInstantiatesCanonicalIdOrExtension {
		m.InstantiatesCanonicalPrimitiveElement = make([]*primitiveElement, 0, len(r.InstantiatesCanonical))
		for _, e := range r.InstantiatesCanonical {
			if e.Id != nil || e.Extension != nil {
				m.InstantiatesCanonicalPrimitiveElement = append(m.InstantiatesCanonicalPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.InstantiatesCanonicalPrimitiveElement = append(m.InstantiatesCanonicalPrimitiveElement, nil)
			}
		}
	}
	m.InstantiatesUri = r.InstantiatesUri
	anyInstantiatesUriIdOrExtension := false
	for _, e := range r.InstantiatesUri {
		if e.Id != nil || e.Extension != nil {
			anyInstantiatesUriIdOrExtension = true
			break
		}
	}
	if anyInstantiatesUriIdOrExtension {
		m.InstantiatesUriPrimitiveElement = make([]*primitiveElement, 0, len(r.InstantiatesUri))
		for _, e := range r.InstantiatesUri {
			if e.Id != nil || e.Extension != nil {
				m.InstantiatesUriPrimitiveElement = append(m.InstantiatesUriPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.InstantiatesUriPrimitiveElement = append(m.InstantiatesUriPrimitiveElement, nil)
			}
		}
	}
	m.BasedOn = r.BasedOn
	m.GroupIdentifier = r.GroupIdentifier
	m.CourseOfTherapyType = r.CourseOfTherapyType
	m.Insurance = r.Insurance
	m.Note = r.Note
	m.DosageInstruction = r.DosageInstruction
	m.DispenseRequest = r.DispenseRequest
	m.Substitution = r.Substitution
	m.PriorPrescription = r.PriorPrescription
	m.DetectedIssue = r.DetectedIssue
	m.EventHistory = r.EventHistory
	return m
}
func (r *MedicationRequest) UnmarshalJSON(b []byte) error {
	var m jsonMedicationRequest
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationRequest) unmarshalJSON(m jsonMedicationRequest) error {
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
	r.StatusReason = m.StatusReason
	r.Intent = m.Intent
	if m.IntentPrimitiveElement != nil {
		r.Intent.Id = m.IntentPrimitiveElement.Id
		r.Intent.Extension = m.IntentPrimitiveElement.Extension
	}
	r.Category = m.Category
	r.Priority = m.Priority
	if m.PriorityPrimitiveElement != nil {
		r.Priority.Id = m.PriorityPrimitiveElement.Id
		r.Priority.Extension = m.PriorityPrimitiveElement.Extension
	}
	r.DoNotPerform = m.DoNotPerform
	if m.DoNotPerformPrimitiveElement != nil {
		r.DoNotPerform.Id = m.DoNotPerformPrimitiveElement.Id
		r.DoNotPerform.Extension = m.DoNotPerformPrimitiveElement.Extension
	}
	if m.ReportedBoolean != nil || m.ReportedBooleanPrimitiveElement != nil {
		if r.Reported != nil {
			return fmt.Errorf("multiple values for field \"Reported\"")
		}
		v := m.ReportedBoolean
		if m.ReportedBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.ReportedBooleanPrimitiveElement.Id
			v.Extension = m.ReportedBooleanPrimitiveElement.Extension
		}
		r.Reported = v
	}
	if m.ReportedReference != nil {
		if r.Reported != nil {
			return fmt.Errorf("multiple values for field \"Reported\"")
		}
		v := m.ReportedReference
		r.Reported = v
	}
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
	r.Encounter = m.Encounter
	r.SupportingInformation = m.SupportingInformation
	r.AuthoredOn = m.AuthoredOn
	if m.AuthoredOnPrimitiveElement != nil {
		r.AuthoredOn.Id = m.AuthoredOnPrimitiveElement.Id
		r.AuthoredOn.Extension = m.AuthoredOnPrimitiveElement.Extension
	}
	r.Requester = m.Requester
	r.Performer = m.Performer
	r.PerformerType = m.PerformerType
	r.Recorder = m.Recorder
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.InstantiatesCanonical = m.InstantiatesCanonical
	for i, e := range m.InstantiatesCanonicalPrimitiveElement {
		if len(r.InstantiatesCanonical) > i {
			r.InstantiatesCanonical[i].Id = e.Id
			r.InstantiatesCanonical[i].Extension = e.Extension
		} else {
			r.InstantiatesCanonical = append(r.InstantiatesCanonical, Canonical{Id: e.Id, Extension: e.Extension})
		}
	}
	r.InstantiatesUri = m.InstantiatesUri
	for i, e := range m.InstantiatesUriPrimitiveElement {
		if len(r.InstantiatesUri) > i {
			r.InstantiatesUri[i].Id = e.Id
			r.InstantiatesUri[i].Extension = e.Extension
		} else {
			r.InstantiatesUri = append(r.InstantiatesUri, Uri{Id: e.Id, Extension: e.Extension})
		}
	}
	r.BasedOn = m.BasedOn
	r.GroupIdentifier = m.GroupIdentifier
	r.CourseOfTherapyType = m.CourseOfTherapyType
	r.Insurance = m.Insurance
	r.Note = m.Note
	r.DosageInstruction = m.DosageInstruction
	r.DispenseRequest = m.DispenseRequest
	r.Substitution = m.Substitution
	r.PriorPrescription = m.PriorPrescription
	r.DetectedIssue = m.DetectedIssue
	r.EventHistory = m.EventHistory
	return nil
}
func (r MedicationRequest) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates the specific details for the dispense or medication supply part of a medication request (also known as a Medication Prescription or Medication Order).  Note that this information is not always sent with the order.  There may be in some settings (e.g. hospitals) institutional or system support for completing the dispense details in the pharmacy department.
type MedicationRequestDispenseRequest struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates the quantity or duration for the first dispense of the medication.
	InitialFill *MedicationRequestDispenseRequestInitialFill
	// The minimum period of time that must occur between dispenses of the medication.
	DispenseInterval *Duration
	// This indicates the validity period of a prescription (stale dating the Prescription).
	ValidityPeriod *Period
	// An integer indicating the number of times, in addition to the original dispense, (aka refills or repeats) that the patient can receive the prescribed medication. Usage Notes: This integer does not include the original order dispense. This means that if an order indicates dispense 30 tablets plus "3 repeats", then the order can be dispensed a total of 4 times and the patient can receive a total of 120 tablets.  A prescriber may explicitly say that zero refills are permitted after the initial dispense.
	NumberOfRepeatsAllowed *UnsignedInt
	// The amount that is to be dispensed for one fill.
	Quantity *Quantity
	// Identifies the period time over which the supplied product is expected to be used, or the length of time the dispense is expected to last.
	ExpectedSupplyDuration *Duration
	// Indicates the intended dispensing Organization specified by the prescriber.
	Performer *Reference
}
type jsonMedicationRequestDispenseRequest struct {
	Id                                     *string                                      `json:"id,omitempty"`
	Extension                              []Extension                                  `json:"extension,omitempty"`
	ModifierExtension                      []Extension                                  `json:"modifierExtension,omitempty"`
	InitialFill                            *MedicationRequestDispenseRequestInitialFill `json:"initialFill,omitempty"`
	DispenseInterval                       *Duration                                    `json:"dispenseInterval,omitempty"`
	ValidityPeriod                         *Period                                      `json:"validityPeriod,omitempty"`
	NumberOfRepeatsAllowed                 *UnsignedInt                                 `json:"numberOfRepeatsAllowed,omitempty"`
	NumberOfRepeatsAllowedPrimitiveElement *primitiveElement                            `json:"_numberOfRepeatsAllowed,omitempty"`
	Quantity                               *Quantity                                    `json:"quantity,omitempty"`
	ExpectedSupplyDuration                 *Duration                                    `json:"expectedSupplyDuration,omitempty"`
	Performer                              *Reference                                   `json:"performer,omitempty"`
}

func (r MedicationRequestDispenseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationRequestDispenseRequest) marshalJSON() jsonMedicationRequestDispenseRequest {
	m := jsonMedicationRequestDispenseRequest{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.InitialFill = r.InitialFill
	m.DispenseInterval = r.DispenseInterval
	m.ValidityPeriod = r.ValidityPeriod
	m.NumberOfRepeatsAllowed = r.NumberOfRepeatsAllowed
	if r.NumberOfRepeatsAllowed != nil && (r.NumberOfRepeatsAllowed.Id != nil || r.NumberOfRepeatsAllowed.Extension != nil) {
		m.NumberOfRepeatsAllowedPrimitiveElement = &primitiveElement{Id: r.NumberOfRepeatsAllowed.Id, Extension: r.NumberOfRepeatsAllowed.Extension}
	}
	m.Quantity = r.Quantity
	m.ExpectedSupplyDuration = r.ExpectedSupplyDuration
	m.Performer = r.Performer
	return m
}
func (r *MedicationRequestDispenseRequest) UnmarshalJSON(b []byte) error {
	var m jsonMedicationRequestDispenseRequest
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationRequestDispenseRequest) unmarshalJSON(m jsonMedicationRequestDispenseRequest) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.InitialFill = m.InitialFill
	r.DispenseInterval = m.DispenseInterval
	r.ValidityPeriod = m.ValidityPeriod
	r.NumberOfRepeatsAllowed = m.NumberOfRepeatsAllowed
	if m.NumberOfRepeatsAllowedPrimitiveElement != nil {
		r.NumberOfRepeatsAllowed.Id = m.NumberOfRepeatsAllowedPrimitiveElement.Id
		r.NumberOfRepeatsAllowed.Extension = m.NumberOfRepeatsAllowedPrimitiveElement.Extension
	}
	r.Quantity = m.Quantity
	r.ExpectedSupplyDuration = m.ExpectedSupplyDuration
	r.Performer = m.Performer
	return nil
}
func (r MedicationRequestDispenseRequest) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates the quantity or duration for the first dispense of the medication.
type MedicationRequestDispenseRequestInitialFill struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The amount or quantity to provide as part of the first dispense.
	Quantity *Quantity
	// The length of time that the first dispense is expected to last.
	Duration *Duration
}
type jsonMedicationRequestDispenseRequestInitialFill struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Quantity          *Quantity   `json:"quantity,omitempty"`
	Duration          *Duration   `json:"duration,omitempty"`
}

func (r MedicationRequestDispenseRequestInitialFill) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationRequestDispenseRequestInitialFill) marshalJSON() jsonMedicationRequestDispenseRequestInitialFill {
	m := jsonMedicationRequestDispenseRequestInitialFill{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Quantity = r.Quantity
	m.Duration = r.Duration
	return m
}
func (r *MedicationRequestDispenseRequestInitialFill) UnmarshalJSON(b []byte) error {
	var m jsonMedicationRequestDispenseRequestInitialFill
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationRequestDispenseRequestInitialFill) unmarshalJSON(m jsonMedicationRequestDispenseRequestInitialFill) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Quantity = m.Quantity
	r.Duration = m.Duration
	return nil
}
func (r MedicationRequestDispenseRequestInitialFill) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates whether or not substitution can or should be part of the dispense. In some cases, substitution must happen, in other cases substitution must not happen. This block explains the prescriber's intent. If nothing is specified substitution may be done.
type MedicationRequestSubstitution struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// True if the prescriber allows a different drug to be dispensed from what was prescribed.
	Allowed isMedicationRequestSubstitutionAllowed
	// Indicates the reason for the substitution, or why substitution must or must not be performed.
	Reason *CodeableConcept
}
type isMedicationRequestSubstitutionAllowed interface {
	isMedicationRequestSubstitutionAllowed()
}

func (r Boolean) isMedicationRequestSubstitutionAllowed()         {}
func (r CodeableConcept) isMedicationRequestSubstitutionAllowed() {}

type jsonMedicationRequestSubstitution struct {
	Id                             *string           `json:"id,omitempty"`
	Extension                      []Extension       `json:"extension,omitempty"`
	ModifierExtension              []Extension       `json:"modifierExtension,omitempty"`
	AllowedBoolean                 *Boolean          `json:"allowedBoolean,omitempty"`
	AllowedBooleanPrimitiveElement *primitiveElement `json:"_allowedBoolean,omitempty"`
	AllowedCodeableConcept         *CodeableConcept  `json:"allowedCodeableConcept,omitempty"`
	Reason                         *CodeableConcept  `json:"reason,omitempty"`
}

func (r MedicationRequestSubstitution) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationRequestSubstitution) marshalJSON() jsonMedicationRequestSubstitution {
	m := jsonMedicationRequestSubstitution{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Allowed.(type) {
	case Boolean:
		m.AllowedBoolean = &v
		if v.Id != nil || v.Extension != nil {
			m.AllowedBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		m.AllowedBoolean = v
		if v.Id != nil || v.Extension != nil {
			m.AllowedBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case CodeableConcept:
		m.AllowedCodeableConcept = &v
	case *CodeableConcept:
		m.AllowedCodeableConcept = v
	}
	m.Reason = r.Reason
	return m
}
func (r *MedicationRequestSubstitution) UnmarshalJSON(b []byte) error {
	var m jsonMedicationRequestSubstitution
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationRequestSubstitution) unmarshalJSON(m jsonMedicationRequestSubstitution) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.AllowedBoolean != nil || m.AllowedBooleanPrimitiveElement != nil {
		if r.Allowed != nil {
			return fmt.Errorf("multiple values for field \"Allowed\"")
		}
		v := m.AllowedBoolean
		if m.AllowedBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.AllowedBooleanPrimitiveElement.Id
			v.Extension = m.AllowedBooleanPrimitiveElement.Extension
		}
		r.Allowed = v
	}
	if m.AllowedCodeableConcept != nil {
		if r.Allowed != nil {
			return fmt.Errorf("multiple values for field \"Allowed\"")
		}
		v := m.AllowedCodeableConcept
		r.Allowed = v
	}
	r.Reason = m.Reason
	return nil
}
func (r MedicationRequestSubstitution) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
