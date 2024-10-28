package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
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
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A unique identifier assigned to this immunization record.
	Identifier []Identifier
	// Indicates the current status of the immunization event.
	Status Code
	// Indicates the reason the immunization event was not performed.
	StatusReason *CodeableConcept
	// Vaccine that was administered or was to be administered.
	VaccineCode CodeableConcept
	// The patient who either received or did not receive the immunization.
	Patient Reference
	// The visit or admission or other contact between patient and health care provider the immunization was performed as part of.
	Encounter *Reference
	// Date vaccine administered or was to be administered.
	Occurrence isImmunizationOccurrence
	// The date the occurrence of the immunization was first captured in the record - potentially significantly after the occurrence of the event.
	Recorded *DateTime
	// An indication that the content of the record is based on information from the person who administered the vaccine. This reflects the context under which the data was originally recorded.
	PrimarySource *Boolean
	// The source of the data when the report of the immunization event is not based on information from the person who administered the vaccine.
	ReportOrigin *CodeableConcept
	// The service delivery location where the vaccine administration occurred.
	Location *Reference
	// Name of vaccine manufacturer.
	Manufacturer *Reference
	// Lot number of the  vaccine product.
	LotNumber *String
	// Date vaccine batch expires.
	ExpirationDate *Date
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
	// Reasons why the vaccine was administered.
	ReasonCode []CodeableConcept
	// Condition, Observation or DiagnosticReport that supports why the immunization was administered.
	ReasonReference []Reference
	// Indication if a dose is considered to be subpotent. By default, a dose should be considered to be potent.
	IsSubpotent *Boolean
	// Reason why a dose is considered to be subpotent.
	SubpotentReason []CodeableConcept
	// Educational material presented to the patient (or guardian) at the time of vaccine administration.
	Education []ImmunizationEducation
	// Indicates a patient's eligibility for a funding program.
	ProgramEligibility []CodeableConcept
	// Indicates the source of the vaccine actually administered. This may be different than the patient eligibility (e.g. the patient may be eligible for a publically purchased vaccine but due to inventory issues, vaccine purchased with private funds was actually administered).
	FundingSource *CodeableConcept
	// Categorical data indicating that an adverse event is associated in time to an immunization.
	Reaction []ImmunizationReaction
	// The protocol (set of recommendations) being followed by the provider who administered the dose.
	ProtocolApplied []ImmunizationProtocolApplied
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

type isImmunizationOccurrence interface {
	isImmunizationOccurrence()
}

func (r DateTime) isImmunizationOccurrence() {}
func (r String) isImmunizationOccurrence()   {}

type jsonImmunization struct {
	ResourceType                       string                        `json:"resourceType"`
	Id                                 *Id                           `json:"id,omitempty"`
	IdPrimitiveElement                 *primitiveElement             `json:"_id,omitempty"`
	Meta                               *Meta                         `json:"meta,omitempty"`
	ImplicitRules                      *Uri                          `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement      *primitiveElement             `json:"_implicitRules,omitempty"`
	Language                           *Code                         `json:"language,omitempty"`
	LanguagePrimitiveElement           *primitiveElement             `json:"_language,omitempty"`
	Text                               *Narrative                    `json:"text,omitempty"`
	Contained                          []ContainedResource           `json:"contained,omitempty"`
	Extension                          []Extension                   `json:"extension,omitempty"`
	ModifierExtension                  []Extension                   `json:"modifierExtension,omitempty"`
	Identifier                         []Identifier                  `json:"identifier,omitempty"`
	Status                             Code                          `json:"status,omitempty"`
	StatusPrimitiveElement             *primitiveElement             `json:"_status,omitempty"`
	StatusReason                       *CodeableConcept              `json:"statusReason,omitempty"`
	VaccineCode                        CodeableConcept               `json:"vaccineCode,omitempty"`
	Patient                            Reference                     `json:"patient,omitempty"`
	Encounter                          *Reference                    `json:"encounter,omitempty"`
	OccurrenceDateTime                 *DateTime                     `json:"occurrenceDateTime,omitempty"`
	OccurrenceDateTimePrimitiveElement *primitiveElement             `json:"_occurrenceDateTime,omitempty"`
	OccurrenceString                   *String                       `json:"occurrenceString,omitempty"`
	OccurrenceStringPrimitiveElement   *primitiveElement             `json:"_occurrenceString,omitempty"`
	Recorded                           *DateTime                     `json:"recorded,omitempty"`
	RecordedPrimitiveElement           *primitiveElement             `json:"_recorded,omitempty"`
	PrimarySource                      *Boolean                      `json:"primarySource,omitempty"`
	PrimarySourcePrimitiveElement      *primitiveElement             `json:"_primarySource,omitempty"`
	ReportOrigin                       *CodeableConcept              `json:"reportOrigin,omitempty"`
	Location                           *Reference                    `json:"location,omitempty"`
	Manufacturer                       *Reference                    `json:"manufacturer,omitempty"`
	LotNumber                          *String                       `json:"lotNumber,omitempty"`
	LotNumberPrimitiveElement          *primitiveElement             `json:"_lotNumber,omitempty"`
	ExpirationDate                     *Date                         `json:"expirationDate,omitempty"`
	ExpirationDatePrimitiveElement     *primitiveElement             `json:"_expirationDate,omitempty"`
	Site                               *CodeableConcept              `json:"site,omitempty"`
	Route                              *CodeableConcept              `json:"route,omitempty"`
	DoseQuantity                       *Quantity                     `json:"doseQuantity,omitempty"`
	Performer                          []ImmunizationPerformer       `json:"performer,omitempty"`
	Note                               []Annotation                  `json:"note,omitempty"`
	ReasonCode                         []CodeableConcept             `json:"reasonCode,omitempty"`
	ReasonReference                    []Reference                   `json:"reasonReference,omitempty"`
	IsSubpotent                        *Boolean                      `json:"isSubpotent,omitempty"`
	IsSubpotentPrimitiveElement        *primitiveElement             `json:"_isSubpotent,omitempty"`
	SubpotentReason                    []CodeableConcept             `json:"subpotentReason,omitempty"`
	Education                          []ImmunizationEducation       `json:"education,omitempty"`
	ProgramEligibility                 []CodeableConcept             `json:"programEligibility,omitempty"`
	FundingSource                      *CodeableConcept              `json:"fundingSource,omitempty"`
	Reaction                           []ImmunizationReaction        `json:"reaction,omitempty"`
	ProtocolApplied                    []ImmunizationProtocolApplied `json:"protocolApplied,omitempty"`
}

func (r Immunization) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Immunization) marshalJSON() jsonImmunization {
	m := jsonImmunization{}
	m.ResourceType = "Immunization"
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
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.StatusReason = r.StatusReason
	m.VaccineCode = r.VaccineCode
	m.Patient = r.Patient
	m.Encounter = r.Encounter
	switch v := r.Occurrence.(type) {
	case DateTime:
		if v.Value != nil {
			m.OccurrenceDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.OccurrenceDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.OccurrenceDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.OccurrenceDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case String:
		if v.Value != nil {
			m.OccurrenceString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.OccurrenceStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.OccurrenceString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.OccurrenceStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	if r.Recorded != nil && r.Recorded.Value != nil {
		m.Recorded = r.Recorded
	}
	if r.Recorded != nil && (r.Recorded.Id != nil || r.Recorded.Extension != nil) {
		m.RecordedPrimitiveElement = &primitiveElement{Id: r.Recorded.Id, Extension: r.Recorded.Extension}
	}
	if r.PrimarySource != nil && r.PrimarySource.Value != nil {
		m.PrimarySource = r.PrimarySource
	}
	if r.PrimarySource != nil && (r.PrimarySource.Id != nil || r.PrimarySource.Extension != nil) {
		m.PrimarySourcePrimitiveElement = &primitiveElement{Id: r.PrimarySource.Id, Extension: r.PrimarySource.Extension}
	}
	m.ReportOrigin = r.ReportOrigin
	m.Location = r.Location
	m.Manufacturer = r.Manufacturer
	if r.LotNumber != nil && r.LotNumber.Value != nil {
		m.LotNumber = r.LotNumber
	}
	if r.LotNumber != nil && (r.LotNumber.Id != nil || r.LotNumber.Extension != nil) {
		m.LotNumberPrimitiveElement = &primitiveElement{Id: r.LotNumber.Id, Extension: r.LotNumber.Extension}
	}
	if r.ExpirationDate != nil && r.ExpirationDate.Value != nil {
		m.ExpirationDate = r.ExpirationDate
	}
	if r.ExpirationDate != nil && (r.ExpirationDate.Id != nil || r.ExpirationDate.Extension != nil) {
		m.ExpirationDatePrimitiveElement = &primitiveElement{Id: r.ExpirationDate.Id, Extension: r.ExpirationDate.Extension}
	}
	m.Site = r.Site
	m.Route = r.Route
	m.DoseQuantity = r.DoseQuantity
	m.Performer = r.Performer
	m.Note = r.Note
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	if r.IsSubpotent != nil && r.IsSubpotent.Value != nil {
		m.IsSubpotent = r.IsSubpotent
	}
	if r.IsSubpotent != nil && (r.IsSubpotent.Id != nil || r.IsSubpotent.Extension != nil) {
		m.IsSubpotentPrimitiveElement = &primitiveElement{Id: r.IsSubpotent.Id, Extension: r.IsSubpotent.Extension}
	}
	m.SubpotentReason = r.SubpotentReason
	m.Education = r.Education
	m.ProgramEligibility = r.ProgramEligibility
	m.FundingSource = r.FundingSource
	m.Reaction = r.Reaction
	m.ProtocolApplied = r.ProtocolApplied
	return m
}
func (r *Immunization) UnmarshalJSON(b []byte) error {
	var m jsonImmunization
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Immunization) unmarshalJSON(m jsonImmunization) error {
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
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.StatusReason = m.StatusReason
	r.VaccineCode = m.VaccineCode
	r.Patient = m.Patient
	r.Encounter = m.Encounter
	if m.OccurrenceDateTime != nil || m.OccurrenceDateTimePrimitiveElement != nil {
		if r.Occurrence != nil {
			return fmt.Errorf("multiple values for field \"Occurrence\"")
		}
		v := m.OccurrenceDateTime
		if m.OccurrenceDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.OccurrenceDateTimePrimitiveElement.Id
			v.Extension = m.OccurrenceDateTimePrimitiveElement.Extension
		}
		r.Occurrence = v
	}
	if m.OccurrenceString != nil || m.OccurrenceStringPrimitiveElement != nil {
		if r.Occurrence != nil {
			return fmt.Errorf("multiple values for field \"Occurrence\"")
		}
		v := m.OccurrenceString
		if m.OccurrenceStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.OccurrenceStringPrimitiveElement.Id
			v.Extension = m.OccurrenceStringPrimitiveElement.Extension
		}
		r.Occurrence = v
	}
	r.Recorded = m.Recorded
	if m.RecordedPrimitiveElement != nil {
		if r.Recorded == nil {
			r.Recorded = &DateTime{}
		}
		r.Recorded.Id = m.RecordedPrimitiveElement.Id
		r.Recorded.Extension = m.RecordedPrimitiveElement.Extension
	}
	r.PrimarySource = m.PrimarySource
	if m.PrimarySourcePrimitiveElement != nil {
		if r.PrimarySource == nil {
			r.PrimarySource = &Boolean{}
		}
		r.PrimarySource.Id = m.PrimarySourcePrimitiveElement.Id
		r.PrimarySource.Extension = m.PrimarySourcePrimitiveElement.Extension
	}
	r.ReportOrigin = m.ReportOrigin
	r.Location = m.Location
	r.Manufacturer = m.Manufacturer
	r.LotNumber = m.LotNumber
	if m.LotNumberPrimitiveElement != nil {
		if r.LotNumber == nil {
			r.LotNumber = &String{}
		}
		r.LotNumber.Id = m.LotNumberPrimitiveElement.Id
		r.LotNumber.Extension = m.LotNumberPrimitiveElement.Extension
	}
	r.ExpirationDate = m.ExpirationDate
	if m.ExpirationDatePrimitiveElement != nil {
		if r.ExpirationDate == nil {
			r.ExpirationDate = &Date{}
		}
		r.ExpirationDate.Id = m.ExpirationDatePrimitiveElement.Id
		r.ExpirationDate.Extension = m.ExpirationDatePrimitiveElement.Extension
	}
	r.Site = m.Site
	r.Route = m.Route
	r.DoseQuantity = m.DoseQuantity
	r.Performer = m.Performer
	r.Note = m.Note
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.IsSubpotent = m.IsSubpotent
	if m.IsSubpotentPrimitiveElement != nil {
		if r.IsSubpotent == nil {
			r.IsSubpotent = &Boolean{}
		}
		r.IsSubpotent.Id = m.IsSubpotentPrimitiveElement.Id
		r.IsSubpotent.Extension = m.IsSubpotentPrimitiveElement.Extension
	}
	r.SubpotentReason = m.SubpotentReason
	r.Education = m.Education
	r.ProgramEligibility = m.ProgramEligibility
	r.FundingSource = m.FundingSource
	r.Reaction = m.Reaction
	r.ProtocolApplied = m.ProtocolApplied
	return nil
}
func (r Immunization) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	v := make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		v = append(v, ContainedResource{c})
	}
	err = e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contained"}})
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
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
	err = e.EncodeElement(r.Patient, xml.StartElement{Name: xml.Name{Local: "patient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Encounter, xml.StartElement{Name: xml.Name{Local: "encounter"}})
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
	err = e.EncodeElement(r.Recorded, xml.StartElement{Name: xml.Name{Local: "recorded"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PrimarySource, xml.StartElement{Name: xml.Name{Local: "primarySource"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReportOrigin, xml.StartElement{Name: xml.Name{Local: "reportOrigin"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Location, xml.StartElement{Name: xml.Name{Local: "location"}})
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
	err = e.EncodeElement(r.ReasonCode, xml.StartElement{Name: xml.Name{Local: "reasonCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReasonReference, xml.StartElement{Name: xml.Name{Local: "reasonReference"}})
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
	err = e.EncodeElement(r.Education, xml.StartElement{Name: xml.Name{Local: "education"}})
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
			case "recorded":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Recorded = &v
			case "primarySource":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PrimarySource = &v
			case "reportOrigin":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReportOrigin = &v
			case "location":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = &v
			case "manufacturer":
				var v Reference
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
			case "reasonCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReasonCode = append(r.ReasonCode, v)
			case "reasonReference":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReasonReference = append(r.ReasonReference, v)
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
			case "education":
				var v ImmunizationEducation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Education = append(r.Education, v)
			case "programEligibility":
				var v CodeableConcept
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
func (r Immunization) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Indicates who performed the immunization event.
type ImmunizationPerformer struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Describes the type of performance (e.g. ordering provider, administering provider, etc.).
	Function *CodeableConcept
	// The practitioner or organization who performed the action.
	Actor Reference
}
type jsonImmunizationPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor,omitempty"`
}

func (r ImmunizationPerformer) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImmunizationPerformer) marshalJSON() jsonImmunizationPerformer {
	m := jsonImmunizationPerformer{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Function = r.Function
	m.Actor = r.Actor
	return m
}
func (r *ImmunizationPerformer) UnmarshalJSON(b []byte) error {
	var m jsonImmunizationPerformer
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImmunizationPerformer) unmarshalJSON(m jsonImmunizationPerformer) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Function = m.Function
	r.Actor = m.Actor
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
func (r ImmunizationPerformer) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Educational material presented to the patient (or guardian) at the time of vaccine administration.
type ImmunizationEducation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identifier of the material presented to the patient.
	DocumentType *String
	// Reference pointer to the educational material given to the patient if the information was on line.
	Reference *Uri
	// Date the educational material was published.
	PublicationDate *DateTime
	// Date the educational material was given to the patient.
	PresentationDate *DateTime
}
type jsonImmunizationEducation struct {
	Id                               *string           `json:"id,omitempty"`
	Extension                        []Extension       `json:"extension,omitempty"`
	ModifierExtension                []Extension       `json:"modifierExtension,omitempty"`
	DocumentType                     *String           `json:"documentType,omitempty"`
	DocumentTypePrimitiveElement     *primitiveElement `json:"_documentType,omitempty"`
	Reference                        *Uri              `json:"reference,omitempty"`
	ReferencePrimitiveElement        *primitiveElement `json:"_reference,omitempty"`
	PublicationDate                  *DateTime         `json:"publicationDate,omitempty"`
	PublicationDatePrimitiveElement  *primitiveElement `json:"_publicationDate,omitempty"`
	PresentationDate                 *DateTime         `json:"presentationDate,omitempty"`
	PresentationDatePrimitiveElement *primitiveElement `json:"_presentationDate,omitempty"`
}

func (r ImmunizationEducation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImmunizationEducation) marshalJSON() jsonImmunizationEducation {
	m := jsonImmunizationEducation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.DocumentType != nil && r.DocumentType.Value != nil {
		m.DocumentType = r.DocumentType
	}
	if r.DocumentType != nil && (r.DocumentType.Id != nil || r.DocumentType.Extension != nil) {
		m.DocumentTypePrimitiveElement = &primitiveElement{Id: r.DocumentType.Id, Extension: r.DocumentType.Extension}
	}
	if r.Reference != nil && r.Reference.Value != nil {
		m.Reference = r.Reference
	}
	if r.Reference != nil && (r.Reference.Id != nil || r.Reference.Extension != nil) {
		m.ReferencePrimitiveElement = &primitiveElement{Id: r.Reference.Id, Extension: r.Reference.Extension}
	}
	if r.PublicationDate != nil && r.PublicationDate.Value != nil {
		m.PublicationDate = r.PublicationDate
	}
	if r.PublicationDate != nil && (r.PublicationDate.Id != nil || r.PublicationDate.Extension != nil) {
		m.PublicationDatePrimitiveElement = &primitiveElement{Id: r.PublicationDate.Id, Extension: r.PublicationDate.Extension}
	}
	if r.PresentationDate != nil && r.PresentationDate.Value != nil {
		m.PresentationDate = r.PresentationDate
	}
	if r.PresentationDate != nil && (r.PresentationDate.Id != nil || r.PresentationDate.Extension != nil) {
		m.PresentationDatePrimitiveElement = &primitiveElement{Id: r.PresentationDate.Id, Extension: r.PresentationDate.Extension}
	}
	return m
}
func (r *ImmunizationEducation) UnmarshalJSON(b []byte) error {
	var m jsonImmunizationEducation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImmunizationEducation) unmarshalJSON(m jsonImmunizationEducation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.DocumentType = m.DocumentType
	if m.DocumentTypePrimitiveElement != nil {
		if r.DocumentType == nil {
			r.DocumentType = &String{}
		}
		r.DocumentType.Id = m.DocumentTypePrimitiveElement.Id
		r.DocumentType.Extension = m.DocumentTypePrimitiveElement.Extension
	}
	r.Reference = m.Reference
	if m.ReferencePrimitiveElement != nil {
		if r.Reference == nil {
			r.Reference = &Uri{}
		}
		r.Reference.Id = m.ReferencePrimitiveElement.Id
		r.Reference.Extension = m.ReferencePrimitiveElement.Extension
	}
	r.PublicationDate = m.PublicationDate
	if m.PublicationDatePrimitiveElement != nil {
		if r.PublicationDate == nil {
			r.PublicationDate = &DateTime{}
		}
		r.PublicationDate.Id = m.PublicationDatePrimitiveElement.Id
		r.PublicationDate.Extension = m.PublicationDatePrimitiveElement.Extension
	}
	r.PresentationDate = m.PresentationDate
	if m.PresentationDatePrimitiveElement != nil {
		if r.PresentationDate == nil {
			r.PresentationDate = &DateTime{}
		}
		r.PresentationDate.Id = m.PresentationDatePrimitiveElement.Id
		r.PresentationDate.Extension = m.PresentationDatePrimitiveElement.Extension
	}
	return nil
}
func (r ImmunizationEducation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.DocumentType, xml.StartElement{Name: xml.Name{Local: "documentType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Reference, xml.StartElement{Name: xml.Name{Local: "reference"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PublicationDate, xml.StartElement{Name: xml.Name{Local: "publicationDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PresentationDate, xml.StartElement{Name: xml.Name{Local: "presentationDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImmunizationEducation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "documentType":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DocumentType = &v
			case "reference":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reference = &v
			case "publicationDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PublicationDate = &v
			case "presentationDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PresentationDate = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImmunizationEducation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Categorical data indicating that an adverse event is associated in time to an immunization.
type ImmunizationReaction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Date of reaction to the immunization.
	Date *DateTime
	// Details of the reaction.
	Detail *Reference
	// Self-reported indicator.
	Reported *Boolean
}
type jsonImmunizationReaction struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Date                     *DateTime         `json:"date,omitempty"`
	DatePrimitiveElement     *primitiveElement `json:"_date,omitempty"`
	Detail                   *Reference        `json:"detail,omitempty"`
	Reported                 *Boolean          `json:"reported,omitempty"`
	ReportedPrimitiveElement *primitiveElement `json:"_reported,omitempty"`
}

func (r ImmunizationReaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImmunizationReaction) marshalJSON() jsonImmunizationReaction {
	m := jsonImmunizationReaction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Date != nil && r.Date.Value != nil {
		m.Date = r.Date
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Detail = r.Detail
	if r.Reported != nil && r.Reported.Value != nil {
		m.Reported = r.Reported
	}
	if r.Reported != nil && (r.Reported.Id != nil || r.Reported.Extension != nil) {
		m.ReportedPrimitiveElement = &primitiveElement{Id: r.Reported.Id, Extension: r.Reported.Extension}
	}
	return m
}
func (r *ImmunizationReaction) UnmarshalJSON(b []byte) error {
	var m jsonImmunizationReaction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImmunizationReaction) unmarshalJSON(m jsonImmunizationReaction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		if r.Date == nil {
			r.Date = &DateTime{}
		}
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Detail = m.Detail
	r.Reported = m.Reported
	if m.ReportedPrimitiveElement != nil {
		if r.Reported == nil {
			r.Reported = &Boolean{}
		}
		r.Reported.Id = m.ReportedPrimitiveElement.Id
		r.Reported.Extension = m.ReportedPrimitiveElement.Extension
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
	err = e.EncodeElement(r.Detail, xml.StartElement{Name: xml.Name{Local: "detail"}})
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
			case "detail":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = &v
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
func (r ImmunizationReaction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The protocol (set of recommendations) being followed by the provider who administered the dose.
type ImmunizationProtocolApplied struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// One possible path to achieve presumed immunity against a disease - within the context of an authority.
	Series *String
	// Indicates the authority who published the protocol (e.g. ACIP) that is being followed.
	Authority *Reference
	// The vaccine preventable disease the dose is being administered against.
	TargetDisease []CodeableConcept
	// Nominal position in a series.
	DoseNumber isImmunizationProtocolAppliedDoseNumber
	// The recommended number of doses to achieve immunity.
	SeriesDoses isImmunizationProtocolAppliedSeriesDoses
}
type isImmunizationProtocolAppliedDoseNumber interface {
	isImmunizationProtocolAppliedDoseNumber()
}

func (r PositiveInt) isImmunizationProtocolAppliedDoseNumber() {}
func (r String) isImmunizationProtocolAppliedDoseNumber()      {}

type isImmunizationProtocolAppliedSeriesDoses interface {
	isImmunizationProtocolAppliedSeriesDoses()
}

func (r PositiveInt) isImmunizationProtocolAppliedSeriesDoses() {}
func (r String) isImmunizationProtocolAppliedSeriesDoses()      {}

type jsonImmunizationProtocolApplied struct {
	Id                                     *string           `json:"id,omitempty"`
	Extension                              []Extension       `json:"extension,omitempty"`
	ModifierExtension                      []Extension       `json:"modifierExtension,omitempty"`
	Series                                 *String           `json:"series,omitempty"`
	SeriesPrimitiveElement                 *primitiveElement `json:"_series,omitempty"`
	Authority                              *Reference        `json:"authority,omitempty"`
	TargetDisease                          []CodeableConcept `json:"targetDisease,omitempty"`
	DoseNumberPositiveInt                  *PositiveInt      `json:"doseNumberPositiveInt,omitempty"`
	DoseNumberPositiveIntPrimitiveElement  *primitiveElement `json:"_doseNumberPositiveInt,omitempty"`
	DoseNumberString                       *String           `json:"doseNumberString,omitempty"`
	DoseNumberStringPrimitiveElement       *primitiveElement `json:"_doseNumberString,omitempty"`
	SeriesDosesPositiveInt                 *PositiveInt      `json:"seriesDosesPositiveInt,omitempty"`
	SeriesDosesPositiveIntPrimitiveElement *primitiveElement `json:"_seriesDosesPositiveInt,omitempty"`
	SeriesDosesString                      *String           `json:"seriesDosesString,omitempty"`
	SeriesDosesStringPrimitiveElement      *primitiveElement `json:"_seriesDosesString,omitempty"`
}

func (r ImmunizationProtocolApplied) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImmunizationProtocolApplied) marshalJSON() jsonImmunizationProtocolApplied {
	m := jsonImmunizationProtocolApplied{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Series != nil && r.Series.Value != nil {
		m.Series = r.Series
	}
	if r.Series != nil && (r.Series.Id != nil || r.Series.Extension != nil) {
		m.SeriesPrimitiveElement = &primitiveElement{Id: r.Series.Id, Extension: r.Series.Extension}
	}
	m.Authority = r.Authority
	m.TargetDisease = r.TargetDisease
	switch v := r.DoseNumber.(type) {
	case PositiveInt:
		if v.Value != nil {
			m.DoseNumberPositiveInt = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.DoseNumberPositiveIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *PositiveInt:
		if v.Value != nil {
			m.DoseNumberPositiveInt = v
		}
		if v.Id != nil || v.Extension != nil {
			m.DoseNumberPositiveIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case String:
		if v.Value != nil {
			m.DoseNumberString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.DoseNumberStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.DoseNumberString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.DoseNumberStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	switch v := r.SeriesDoses.(type) {
	case PositiveInt:
		if v.Value != nil {
			m.SeriesDosesPositiveInt = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.SeriesDosesPositiveIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *PositiveInt:
		if v.Value != nil {
			m.SeriesDosesPositiveInt = v
		}
		if v.Id != nil || v.Extension != nil {
			m.SeriesDosesPositiveIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case String:
		if v.Value != nil {
			m.SeriesDosesString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.SeriesDosesStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.SeriesDosesString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.SeriesDosesStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	return m
}
func (r *ImmunizationProtocolApplied) UnmarshalJSON(b []byte) error {
	var m jsonImmunizationProtocolApplied
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImmunizationProtocolApplied) unmarshalJSON(m jsonImmunizationProtocolApplied) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Series = m.Series
	if m.SeriesPrimitiveElement != nil {
		if r.Series == nil {
			r.Series = &String{}
		}
		r.Series.Id = m.SeriesPrimitiveElement.Id
		r.Series.Extension = m.SeriesPrimitiveElement.Extension
	}
	r.Authority = m.Authority
	r.TargetDisease = m.TargetDisease
	if m.DoseNumberPositiveInt != nil || m.DoseNumberPositiveIntPrimitiveElement != nil {
		if r.DoseNumber != nil {
			return fmt.Errorf("multiple values for field \"DoseNumber\"")
		}
		v := m.DoseNumberPositiveInt
		if m.DoseNumberPositiveIntPrimitiveElement != nil {
			if v == nil {
				v = &PositiveInt{}
			}
			v.Id = m.DoseNumberPositiveIntPrimitiveElement.Id
			v.Extension = m.DoseNumberPositiveIntPrimitiveElement.Extension
		}
		r.DoseNumber = v
	}
	if m.DoseNumberString != nil || m.DoseNumberStringPrimitiveElement != nil {
		if r.DoseNumber != nil {
			return fmt.Errorf("multiple values for field \"DoseNumber\"")
		}
		v := m.DoseNumberString
		if m.DoseNumberStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.DoseNumberStringPrimitiveElement.Id
			v.Extension = m.DoseNumberStringPrimitiveElement.Extension
		}
		r.DoseNumber = v
	}
	if m.SeriesDosesPositiveInt != nil || m.SeriesDosesPositiveIntPrimitiveElement != nil {
		if r.SeriesDoses != nil {
			return fmt.Errorf("multiple values for field \"SeriesDoses\"")
		}
		v := m.SeriesDosesPositiveInt
		if m.SeriesDosesPositiveIntPrimitiveElement != nil {
			if v == nil {
				v = &PositiveInt{}
			}
			v.Id = m.SeriesDosesPositiveIntPrimitiveElement.Id
			v.Extension = m.SeriesDosesPositiveIntPrimitiveElement.Extension
		}
		r.SeriesDoses = v
	}
	if m.SeriesDosesString != nil || m.SeriesDosesStringPrimitiveElement != nil {
		if r.SeriesDoses != nil {
			return fmt.Errorf("multiple values for field \"SeriesDoses\"")
		}
		v := m.SeriesDosesString
		if m.SeriesDosesStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.SeriesDosesStringPrimitiveElement.Id
			v.Extension = m.SeriesDosesStringPrimitiveElement.Extension
		}
		r.SeriesDoses = v
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
	switch v := r.DoseNumber.(type) {
	case PositiveInt, *PositiveInt:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "doseNumberPositiveInt"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "doseNumberString"}})
		if err != nil {
			return err
		}
	}
	switch v := r.SeriesDoses.(type) {
	case PositiveInt, *PositiveInt:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "seriesDosesPositiveInt"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "seriesDosesString"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
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
			case "doseNumberPositiveInt":
				if r.DoseNumber != nil {
					return fmt.Errorf("multiple values for field \"DoseNumber\"")
				}
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DoseNumber = v
			case "doseNumberString":
				if r.DoseNumber != nil {
					return fmt.Errorf("multiple values for field \"DoseNumber\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DoseNumber = v
			case "seriesDosesPositiveInt":
				if r.SeriesDoses != nil {
					return fmt.Errorf("multiple values for field \"SeriesDoses\"")
				}
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SeriesDoses = v
			case "seriesDosesString":
				if r.SeriesDoses != nil {
					return fmt.Errorf("multiple values for field \"SeriesDoses\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SeriesDoses = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImmunizationProtocolApplied) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
