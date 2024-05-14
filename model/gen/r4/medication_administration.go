package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// Describes the event of a patient consuming or otherwise being administered a medication.  This may be as simple as swallowing a tablet or it may be a long running infusion.  Related resources tie this event to the authorizing prescription, and the specific encounter between patient and health care practitioner.
type MedicationAdministration struct {
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
	// Identifiers associated with this Medication Administration that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate. They are business identifiers assigned to this resource by the performer or other systems and remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// A protocol, guideline, orderset, or other definition that was adhered to in whole or in part by this event.
	Instantiates []Uri
	// A larger event of which this particular event is a component or step.
	PartOf []Reference
	// Will generally be set to show that the administration has been completed.  For some long running administrations such as infusions, it is possible for an administration to be started but not completed or it may be paused while some other process is under way.
	Status Code
	// A code indicating why the administration was not performed.
	StatusReason []CodeableConcept
	// Indicates where the medication is expected to be consumed or administered.
	Category *CodeableConcept
	// Identifies the medication that was administered. This is either a link to a resource representing the details of the medication or a simple attribute carrying a code that identifies the medication from a known list of medications.
	Medication isMedicationAdministrationMedication
	// The person or animal or group receiving the medication.
	Subject Reference
	// The visit, admission, or other contact between patient and health care provider during which the medication administration was performed.
	Context *Reference
	// Additional information (for example, patient height and weight) that supports the administration of the medication.
	SupportingInformation []Reference
	// A specific date/time or interval of time during which the administration took place (or did not take place, when the 'notGiven' attribute is true). For many administrations, such as swallowing a tablet the use of dateTime is more appropriate.
	Effective isMedicationAdministrationEffective
	// Indicates who or what performed the medication administration and how they were involved.
	Performer []MedicationAdministrationPerformer
	// A code indicating why the medication was given.
	ReasonCode []CodeableConcept
	// Condition or observation that supports why the medication was administered.
	ReasonReference []Reference
	// The original request, instruction or authority to perform the administration.
	Request *Reference
	// The device used in administering the medication to the patient.  For example, a particular infusion pump.
	Device []Reference
	// Extra information about the medication administration that is not conveyed by the other attributes.
	Note []Annotation
	// Describes the medication dosage information details e.g. dose, rate, site, route, etc.
	Dosage *MedicationAdministrationDosage
	// A summary of the events of interest that have occurred, such as when the administration was verified.
	EventHistory []Reference
}

func (r MedicationAdministration) ResourceType() string {
	return "MedicationAdministration"
}

type isMedicationAdministrationMedication interface {
	isMedicationAdministrationMedication()
}

func (r CodeableConcept) isMedicationAdministrationMedication() {}
func (r Reference) isMedicationAdministrationMedication()       {}

type isMedicationAdministrationEffective interface {
	isMedicationAdministrationEffective()
}

func (r DateTime) isMedicationAdministrationEffective() {}
func (r Period) isMedicationAdministrationEffective()   {}

type jsonMedicationAdministration struct {
	ResourceType                      string                              `json:"resourceType"`
	Id                                *Id                                 `json:"id,omitempty"`
	IdPrimitiveElement                *primitiveElement                   `json:"_id,omitempty"`
	Meta                              *Meta                               `json:"meta,omitempty"`
	ImplicitRules                     *Uri                                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement     *primitiveElement                   `json:"_implicitRules,omitempty"`
	Language                          *Code                               `json:"language,omitempty"`
	LanguagePrimitiveElement          *primitiveElement                   `json:"_language,omitempty"`
	Text                              *Narrative                          `json:"text,omitempty"`
	Contained                         []containedResource                 `json:"contained,omitempty"`
	Extension                         []Extension                         `json:"extension,omitempty"`
	ModifierExtension                 []Extension                         `json:"modifierExtension,omitempty"`
	Identifier                        []Identifier                        `json:"identifier,omitempty"`
	Instantiates                      []Uri                               `json:"instantiates,omitempty"`
	InstantiatesPrimitiveElement      []*primitiveElement                 `json:"_instantiates,omitempty"`
	PartOf                            []Reference                         `json:"partOf,omitempty"`
	Status                            Code                                `json:"status,omitempty"`
	StatusPrimitiveElement            *primitiveElement                   `json:"_status,omitempty"`
	StatusReason                      []CodeableConcept                   `json:"statusReason,omitempty"`
	Category                          *CodeableConcept                    `json:"category,omitempty"`
	MedicationCodeableConcept         *CodeableConcept                    `json:"medicationCodeableConcept,omitempty"`
	MedicationReference               *Reference                          `json:"medicationReference,omitempty"`
	Subject                           Reference                           `json:"subject,omitempty"`
	Context                           *Reference                          `json:"context,omitempty"`
	SupportingInformation             []Reference                         `json:"supportingInformation,omitempty"`
	EffectiveDateTime                 *DateTime                           `json:"effectiveDateTime,omitempty"`
	EffectiveDateTimePrimitiveElement *primitiveElement                   `json:"_effectiveDateTime,omitempty"`
	EffectivePeriod                   *Period                             `json:"effectivePeriod,omitempty"`
	Performer                         []MedicationAdministrationPerformer `json:"performer,omitempty"`
	ReasonCode                        []CodeableConcept                   `json:"reasonCode,omitempty"`
	ReasonReference                   []Reference                         `json:"reasonReference,omitempty"`
	Request                           *Reference                          `json:"request,omitempty"`
	Device                            []Reference                         `json:"device,omitempty"`
	Note                              []Annotation                        `json:"note,omitempty"`
	Dosage                            *MedicationAdministrationDosage     `json:"dosage,omitempty"`
	EventHistory                      []Reference                         `json:"eventHistory,omitempty"`
}

func (r MedicationAdministration) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationAdministration) marshalJSON() jsonMedicationAdministration {
	m := jsonMedicationAdministration{}
	m.ResourceType = "MedicationAdministration"
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
	m.Instantiates = r.Instantiates
	anyInstantiatesIdOrExtension := false
	for _, e := range r.Instantiates {
		if e.Id != nil || e.Extension != nil {
			anyInstantiatesIdOrExtension = true
			break
		}
	}
	if anyInstantiatesIdOrExtension {
		m.InstantiatesPrimitiveElement = make([]*primitiveElement, 0, len(r.Instantiates))
		for _, e := range r.Instantiates {
			if e.Id != nil || e.Extension != nil {
				m.InstantiatesPrimitiveElement = append(m.InstantiatesPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.InstantiatesPrimitiveElement = append(m.InstantiatesPrimitiveElement, nil)
			}
		}
	}
	m.PartOf = r.PartOf
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.StatusReason = r.StatusReason
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
	switch v := r.Effective.(type) {
	case DateTime:
		m.EffectiveDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.EffectiveDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.EffectiveDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.EffectiveDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.EffectivePeriod = &v
	case *Period:
		m.EffectivePeriod = v
	}
	m.Performer = r.Performer
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Request = r.Request
	m.Device = r.Device
	m.Note = r.Note
	m.Dosage = r.Dosage
	m.EventHistory = r.EventHistory
	return m
}
func (r *MedicationAdministration) UnmarshalJSON(b []byte) error {
	var m jsonMedicationAdministration
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationAdministration) unmarshalJSON(m jsonMedicationAdministration) error {
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
	r.Instantiates = m.Instantiates
	for i, e := range m.InstantiatesPrimitiveElement {
		if len(r.Instantiates) > i {
			r.Instantiates[i].Id = e.Id
			r.Instantiates[i].Extension = e.Extension
		} else {
			r.Instantiates = append(r.Instantiates, Uri{Id: e.Id, Extension: e.Extension})
		}
	}
	r.PartOf = m.PartOf
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.StatusReason = m.StatusReason
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
	if m.EffectiveDateTime != nil || m.EffectiveDateTimePrimitiveElement != nil {
		if r.Effective != nil {
			return fmt.Errorf("multiple values for field \"Effective\"")
		}
		v := m.EffectiveDateTime
		if m.EffectiveDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.EffectiveDateTimePrimitiveElement.Id
			v.Extension = m.EffectiveDateTimePrimitiveElement.Extension
		}
		r.Effective = v
	}
	if m.EffectivePeriod != nil {
		if r.Effective != nil {
			return fmt.Errorf("multiple values for field \"Effective\"")
		}
		v := m.EffectivePeriod
		r.Effective = v
	}
	r.Performer = m.Performer
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Request = m.Request
	r.Device = m.Device
	r.Note = m.Note
	r.Dosage = m.Dosage
	r.EventHistory = m.EventHistory
	return nil
}
func (r MedicationAdministration) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates who or what performed the medication administration and how they were involved.
type MedicationAdministrationPerformer struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Distinguishes the type of involvement of the performer in the medication administration.
	Function *CodeableConcept
	// Indicates who or what performed the medication administration.
	Actor Reference
}
type jsonMedicationAdministrationPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor,omitempty"`
}

func (r MedicationAdministrationPerformer) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationAdministrationPerformer) marshalJSON() jsonMedicationAdministrationPerformer {
	m := jsonMedicationAdministrationPerformer{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Function = r.Function
	m.Actor = r.Actor
	return m
}
func (r *MedicationAdministrationPerformer) UnmarshalJSON(b []byte) error {
	var m jsonMedicationAdministrationPerformer
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationAdministrationPerformer) unmarshalJSON(m jsonMedicationAdministrationPerformer) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Function = m.Function
	r.Actor = m.Actor
	return nil
}
func (r MedicationAdministrationPerformer) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Describes the medication dosage information details e.g. dose, rate, site, route, etc.
type MedicationAdministrationDosage struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Free text dosage can be used for cases where the dosage administered is too complex to code. When coded dosage is present, the free text dosage may still be present for display to humans.The dosage instructions should reflect the dosage of the medication that was administered.
	Text *String
	// A coded specification of the anatomic site where the medication first entered the body.  For example, "left arm".
	Site *CodeableConcept
	// A code specifying the route or physiological path of administration of a therapeutic agent into or onto the patient.  For example, topical, intravenous, etc.
	Route *CodeableConcept
	// A coded value indicating the method by which the medication is intended to be or was introduced into or on the body.  This attribute will most often NOT be populated.  It is most commonly used for injections.  For example, Slow Push, Deep IV.
	Method *CodeableConcept
	// The amount of the medication given at one administration event.   Use this value when the administration is essentially an instantaneous event such as a swallowing a tablet or giving an injection.
	Dose *Quantity
	// Identifies the speed with which the medication was or will be introduced into the patient.  Typically, the rate for an infusion e.g. 100 ml per 1 hour or 100 ml/hr.  May also be expressed as a rate per unit of time, e.g. 500 ml per 2 hours.  Other examples:  200 mcg/min or 200 mcg/1 minute; 1 liter/8 hours.
	Rate isMedicationAdministrationDosageRate
}
type isMedicationAdministrationDosageRate interface {
	isMedicationAdministrationDosageRate()
}

func (r Ratio) isMedicationAdministrationDosageRate()    {}
func (r Quantity) isMedicationAdministrationDosageRate() {}

type jsonMedicationAdministrationDosage struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Text                 *String           `json:"text,omitempty"`
	TextPrimitiveElement *primitiveElement `json:"_text,omitempty"`
	Site                 *CodeableConcept  `json:"site,omitempty"`
	Route                *CodeableConcept  `json:"route,omitempty"`
	Method               *CodeableConcept  `json:"method,omitempty"`
	Dose                 *Quantity         `json:"dose,omitempty"`
	RateRatio            *Ratio            `json:"rateRatio,omitempty"`
	RateQuantity         *Quantity         `json:"rateQuantity,omitempty"`
}

func (r MedicationAdministrationDosage) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicationAdministrationDosage) marshalJSON() jsonMedicationAdministrationDosage {
	m := jsonMedicationAdministrationDosage{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Text = r.Text
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	m.Site = r.Site
	m.Route = r.Route
	m.Method = r.Method
	m.Dose = r.Dose
	switch v := r.Rate.(type) {
	case Ratio:
		m.RateRatio = &v
	case *Ratio:
		m.RateRatio = v
	case Quantity:
		m.RateQuantity = &v
	case *Quantity:
		m.RateQuantity = v
	}
	return m
}
func (r *MedicationAdministrationDosage) UnmarshalJSON(b []byte) error {
	var m jsonMedicationAdministrationDosage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicationAdministrationDosage) unmarshalJSON(m jsonMedicationAdministrationDosage) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	r.Site = m.Site
	r.Route = m.Route
	r.Method = m.Method
	r.Dose = m.Dose
	if m.RateRatio != nil {
		if r.Rate != nil {
			return fmt.Errorf("multiple values for field \"Rate\"")
		}
		v := m.RateRatio
		r.Rate = v
	}
	if m.RateQuantity != nil {
		if r.Rate != nil {
			return fmt.Errorf("multiple values for field \"Rate\"")
		}
		v := m.RateQuantity
		r.Rate = v
	}
	return nil
}
func (r MedicationAdministrationDosage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
