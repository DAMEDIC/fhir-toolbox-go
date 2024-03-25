package r4

import (
	"encoding/json"
	"fmt"
)

// Significant health conditions for a person related to the patient relevant in the context of care for the patient.
type FamilyMemberHistory struct {
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
	Contained []Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Business identifiers assigned to this family member history by the performer or other systems which remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this FamilyMemberHistory.
	InstantiatesCanonical []Canonical
	// The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this FamilyMemberHistory.
	InstantiatesUri []Uri
	// A code specifying the status of the record of the family history of a specific family member.
	Status Code
	// Describes why the family member's history is not available.
	DataAbsentReason *CodeableConcept
	// The person who this history concerns.
	Patient Reference
	// The date (and possibly time) when the family member history was recorded or last updated.
	Date *DateTime
	// This will either be a name or a description; e.g. "Aunt Susan", "my cousin with the red hair".
	Name *String
	// The type of relationship this person has to the patient (father, mother, brother etc.).
	Relationship CodeableConcept
	// The birth sex of the family member.
	Sex *CodeableConcept
	// The actual or approximate date of birth of the relative.
	Born isFamilyMemberHistoryBorn
	// The age of the relative at the time the family member history is recorded.
	Age isFamilyMemberHistoryAge
	// If true, indicates that the age value specified is an estimated value.
	EstimatedAge *Boolean
	// Deceased flag or the actual or approximate age of the relative at the time of death for the family member history record.
	Deceased isFamilyMemberHistoryDeceased
	// Describes why the family member history occurred in coded or textual form.
	ReasonCode []CodeableConcept
	// Indicates a Condition, Observation, AllergyIntolerance, or QuestionnaireResponse that justifies this family member history event.
	ReasonReference []Reference
	// This property allows a non condition-specific note to the made about the related person. Ideally, the note would be in the condition property, but this is not always possible.
	Note []Annotation
	// The significant Conditions (or condition) that the family member had. This is a repeating section to allow a system to represent more than one condition per resource, though there is nothing stopping multiple resources - one per condition.
	Condition []FamilyMemberHistoryCondition
}
type isFamilyMemberHistoryBorn interface {
	isFamilyMemberHistoryBorn()
}

func (r Period) isFamilyMemberHistoryBorn() {}
func (r Date) isFamilyMemberHistoryBorn()   {}
func (r String) isFamilyMemberHistoryBorn() {}

type isFamilyMemberHistoryAge interface {
	isFamilyMemberHistoryAge()
}

func (r Age) isFamilyMemberHistoryAge()    {}
func (r Range) isFamilyMemberHistoryAge()  {}
func (r String) isFamilyMemberHistoryAge() {}

type isFamilyMemberHistoryDeceased interface {
	isFamilyMemberHistoryDeceased()
}

func (r Boolean) isFamilyMemberHistoryDeceased() {}
func (r Age) isFamilyMemberHistoryDeceased()     {}
func (r Range) isFamilyMemberHistoryDeceased()   {}
func (r Date) isFamilyMemberHistoryDeceased()    {}
func (r String) isFamilyMemberHistoryDeceased()  {}

type jsonFamilyMemberHistory struct {
	ResourceType                          string                         `json:"resourceType"`
	Id                                    *Id                            `json:"id,omitempty"`
	IdPrimitiveElement                    *primitiveElement              `json:"_id,omitempty"`
	Meta                                  *Meta                          `json:"meta,omitempty"`
	ImplicitRules                         *Uri                           `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement         *primitiveElement              `json:"_implicitRules,omitempty"`
	Language                              *Code                          `json:"language,omitempty"`
	LanguagePrimitiveElement              *primitiveElement              `json:"_language,omitempty"`
	Text                                  *Narrative                     `json:"text,omitempty"`
	Contained                             []containedResource            `json:"contained,omitempty"`
	Extension                             []Extension                    `json:"extension,omitempty"`
	ModifierExtension                     []Extension                    `json:"modifierExtension,omitempty"`
	Identifier                            []Identifier                   `json:"identifier,omitempty"`
	InstantiatesCanonical                 []Canonical                    `json:"instantiatesCanonical,omitempty"`
	InstantiatesCanonicalPrimitiveElement []*primitiveElement            `json:"_instantiatesCanonical,omitempty"`
	InstantiatesUri                       []Uri                          `json:"instantiatesUri,omitempty"`
	InstantiatesUriPrimitiveElement       []*primitiveElement            `json:"_instantiatesUri,omitempty"`
	Status                                Code                           `json:"status,omitempty"`
	StatusPrimitiveElement                *primitiveElement              `json:"_status,omitempty"`
	DataAbsentReason                      *CodeableConcept               `json:"dataAbsentReason,omitempty"`
	Patient                               Reference                      `json:"patient,omitempty"`
	Date                                  *DateTime                      `json:"date,omitempty"`
	DatePrimitiveElement                  *primitiveElement              `json:"_date,omitempty"`
	Name                                  *String                        `json:"name,omitempty"`
	NamePrimitiveElement                  *primitiveElement              `json:"_name,omitempty"`
	Relationship                          CodeableConcept                `json:"relationship,omitempty"`
	Sex                                   *CodeableConcept               `json:"sex,omitempty"`
	BornPeriod                            *Period                        `json:"bornPeriod,omitempty"`
	BornDate                              *Date                          `json:"bornDate,omitempty"`
	BornDatePrimitiveElement              *primitiveElement              `json:"_bornDate,omitempty"`
	BornString                            *String                        `json:"bornString,omitempty"`
	BornStringPrimitiveElement            *primitiveElement              `json:"_bornString,omitempty"`
	AgeAge                                *Age                           `json:"ageAge,omitempty"`
	AgeRange                              *Range                         `json:"ageRange,omitempty"`
	AgeString                             *String                        `json:"ageString,omitempty"`
	AgeStringPrimitiveElement             *primitiveElement              `json:"_ageString,omitempty"`
	EstimatedAge                          *Boolean                       `json:"estimatedAge,omitempty"`
	EstimatedAgePrimitiveElement          *primitiveElement              `json:"_estimatedAge,omitempty"`
	DeceasedBoolean                       *Boolean                       `json:"deceasedBoolean,omitempty"`
	DeceasedBooleanPrimitiveElement       *primitiveElement              `json:"_deceasedBoolean,omitempty"`
	DeceasedAge                           *Age                           `json:"deceasedAge,omitempty"`
	DeceasedRange                         *Range                         `json:"deceasedRange,omitempty"`
	DeceasedDate                          *Date                          `json:"deceasedDate,omitempty"`
	DeceasedDatePrimitiveElement          *primitiveElement              `json:"_deceasedDate,omitempty"`
	DeceasedString                        *String                        `json:"deceasedString,omitempty"`
	DeceasedStringPrimitiveElement        *primitiveElement              `json:"_deceasedString,omitempty"`
	ReasonCode                            []CodeableConcept              `json:"reasonCode,omitempty"`
	ReasonReference                       []Reference                    `json:"reasonReference,omitempty"`
	Note                                  []Annotation                   `json:"note,omitempty"`
	Condition                             []FamilyMemberHistoryCondition `json:"condition,omitempty"`
}

func (r FamilyMemberHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r FamilyMemberHistory) marshalJSON() jsonFamilyMemberHistory {
	m := jsonFamilyMemberHistory{}
	m.ResourceType = "FamilyMemberHistory"
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
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.DataAbsentReason = r.DataAbsentReason
	m.Patient = r.Patient
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Relationship = r.Relationship
	m.Sex = r.Sex
	switch v := r.Born.(type) {
	case Period:
		m.BornPeriod = &v
	case *Period:
		m.BornPeriod = v
	case Date:
		m.BornDate = &v
		if v.Id != nil || v.Extension != nil {
			m.BornDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		m.BornDate = v
		if v.Id != nil || v.Extension != nil {
			m.BornDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case String:
		m.BornString = &v
		if v.Id != nil || v.Extension != nil {
			m.BornStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.BornString = v
		if v.Id != nil || v.Extension != nil {
			m.BornStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	switch v := r.Age.(type) {
	case Age:
		m.AgeAge = &v
	case *Age:
		m.AgeAge = v
	case Range:
		m.AgeRange = &v
	case *Range:
		m.AgeRange = v
	case String:
		m.AgeString = &v
		if v.Id != nil || v.Extension != nil {
			m.AgeStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.AgeString = v
		if v.Id != nil || v.Extension != nil {
			m.AgeStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.EstimatedAge = r.EstimatedAge
	if r.EstimatedAge != nil && (r.EstimatedAge.Id != nil || r.EstimatedAge.Extension != nil) {
		m.EstimatedAgePrimitiveElement = &primitiveElement{Id: r.EstimatedAge.Id, Extension: r.EstimatedAge.Extension}
	}
	switch v := r.Deceased.(type) {
	case Boolean:
		m.DeceasedBoolean = &v
		if v.Id != nil || v.Extension != nil {
			m.DeceasedBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		m.DeceasedBoolean = v
		if v.Id != nil || v.Extension != nil {
			m.DeceasedBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Age:
		m.DeceasedAge = &v
	case *Age:
		m.DeceasedAge = v
	case Range:
		m.DeceasedRange = &v
	case *Range:
		m.DeceasedRange = v
	case Date:
		m.DeceasedDate = &v
		if v.Id != nil || v.Extension != nil {
			m.DeceasedDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		m.DeceasedDate = v
		if v.Id != nil || v.Extension != nil {
			m.DeceasedDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case String:
		m.DeceasedString = &v
		if v.Id != nil || v.Extension != nil {
			m.DeceasedStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.DeceasedString = v
		if v.Id != nil || v.Extension != nil {
			m.DeceasedStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Note = r.Note
	m.Condition = r.Condition
	return m
}
func (r *FamilyMemberHistory) UnmarshalJSON(b []byte) error {
	var m jsonFamilyMemberHistory
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *FamilyMemberHistory) unmarshalJSON(m jsonFamilyMemberHistory) error {
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
	r.Contained = make([]Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
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
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.DataAbsentReason = m.DataAbsentReason
	r.Patient = m.Patient
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Relationship = m.Relationship
	r.Sex = m.Sex
	if m.BornPeriod != nil {
		if r.Born != nil {
			return fmt.Errorf("multiple values for field \"Born\"")
		}
		v := m.BornPeriod
		r.Born = v
	}
	if m.BornDate != nil || m.BornDatePrimitiveElement != nil {
		if r.Born != nil {
			return fmt.Errorf("multiple values for field \"Born\"")
		}
		v := m.BornDate
		if m.BornDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.BornDatePrimitiveElement.Id
			v.Extension = m.BornDatePrimitiveElement.Extension
		}
		r.Born = v
	}
	if m.BornString != nil || m.BornStringPrimitiveElement != nil {
		if r.Born != nil {
			return fmt.Errorf("multiple values for field \"Born\"")
		}
		v := m.BornString
		if m.BornStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.BornStringPrimitiveElement.Id
			v.Extension = m.BornStringPrimitiveElement.Extension
		}
		r.Born = v
	}
	if m.AgeAge != nil {
		if r.Age != nil {
			return fmt.Errorf("multiple values for field \"Age\"")
		}
		v := m.AgeAge
		r.Age = v
	}
	if m.AgeRange != nil {
		if r.Age != nil {
			return fmt.Errorf("multiple values for field \"Age\"")
		}
		v := m.AgeRange
		r.Age = v
	}
	if m.AgeString != nil || m.AgeStringPrimitiveElement != nil {
		if r.Age != nil {
			return fmt.Errorf("multiple values for field \"Age\"")
		}
		v := m.AgeString
		if m.AgeStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.AgeStringPrimitiveElement.Id
			v.Extension = m.AgeStringPrimitiveElement.Extension
		}
		r.Age = v
	}
	r.EstimatedAge = m.EstimatedAge
	if m.EstimatedAgePrimitiveElement != nil {
		r.EstimatedAge.Id = m.EstimatedAgePrimitiveElement.Id
		r.EstimatedAge.Extension = m.EstimatedAgePrimitiveElement.Extension
	}
	if m.DeceasedBoolean != nil || m.DeceasedBooleanPrimitiveElement != nil {
		if r.Deceased != nil {
			return fmt.Errorf("multiple values for field \"Deceased\"")
		}
		v := m.DeceasedBoolean
		if m.DeceasedBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.DeceasedBooleanPrimitiveElement.Id
			v.Extension = m.DeceasedBooleanPrimitiveElement.Extension
		}
		r.Deceased = v
	}
	if m.DeceasedAge != nil {
		if r.Deceased != nil {
			return fmt.Errorf("multiple values for field \"Deceased\"")
		}
		v := m.DeceasedAge
		r.Deceased = v
	}
	if m.DeceasedRange != nil {
		if r.Deceased != nil {
			return fmt.Errorf("multiple values for field \"Deceased\"")
		}
		v := m.DeceasedRange
		r.Deceased = v
	}
	if m.DeceasedDate != nil || m.DeceasedDatePrimitiveElement != nil {
		if r.Deceased != nil {
			return fmt.Errorf("multiple values for field \"Deceased\"")
		}
		v := m.DeceasedDate
		if m.DeceasedDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.DeceasedDatePrimitiveElement.Id
			v.Extension = m.DeceasedDatePrimitiveElement.Extension
		}
		r.Deceased = v
	}
	if m.DeceasedString != nil || m.DeceasedStringPrimitiveElement != nil {
		if r.Deceased != nil {
			return fmt.Errorf("multiple values for field \"Deceased\"")
		}
		v := m.DeceasedString
		if m.DeceasedStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.DeceasedStringPrimitiveElement.Id
			v.Extension = m.DeceasedStringPrimitiveElement.Extension
		}
		r.Deceased = v
	}
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Note = m.Note
	r.Condition = m.Condition
	return nil
}
func (r FamilyMemberHistory) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The significant Conditions (or condition) that the family member had. This is a repeating section to allow a system to represent more than one condition per resource, though there is nothing stopping multiple resources - one per condition.
type FamilyMemberHistoryCondition struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The actual condition specified. Could be a coded condition (like MI or Diabetes) or a less specific string like 'cancer' depending on how much is known about the condition and the capabilities of the creating system.
	Code CodeableConcept
	// Indicates what happened following the condition.  If the condition resulted in death, deceased date is captured on the relation.
	Outcome *CodeableConcept
	// This condition contributed to the cause of death of the related person. If contributedToDeath is not populated, then it is unknown.
	ContributedToDeath *Boolean
	// Either the age of onset, range of approximate age or descriptive string can be recorded.  For conditions with multiple occurrences, this describes the first known occurrence.
	Onset isFamilyMemberHistoryConditionOnset
	// An area where general notes can be placed about this specific condition.
	Note []Annotation
}
type isFamilyMemberHistoryConditionOnset interface {
	isFamilyMemberHistoryConditionOnset()
}

func (r Age) isFamilyMemberHistoryConditionOnset()    {}
func (r Range) isFamilyMemberHistoryConditionOnset()  {}
func (r Period) isFamilyMemberHistoryConditionOnset() {}
func (r String) isFamilyMemberHistoryConditionOnset() {}

type jsonFamilyMemberHistoryCondition struct {
	Id                                 *string           `json:"id,omitempty"`
	Extension                          []Extension       `json:"extension,omitempty"`
	ModifierExtension                  []Extension       `json:"modifierExtension,omitempty"`
	Code                               CodeableConcept   `json:"code,omitempty"`
	Outcome                            *CodeableConcept  `json:"outcome,omitempty"`
	ContributedToDeath                 *Boolean          `json:"contributedToDeath,omitempty"`
	ContributedToDeathPrimitiveElement *primitiveElement `json:"_contributedToDeath,omitempty"`
	OnsetAge                           *Age              `json:"onsetAge,omitempty"`
	OnsetRange                         *Range            `json:"onsetRange,omitempty"`
	OnsetPeriod                        *Period           `json:"onsetPeriod,omitempty"`
	OnsetString                        *String           `json:"onsetString,omitempty"`
	OnsetStringPrimitiveElement        *primitiveElement `json:"_onsetString,omitempty"`
	Note                               []Annotation      `json:"note,omitempty"`
}

func (r FamilyMemberHistoryCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r FamilyMemberHistoryCondition) marshalJSON() jsonFamilyMemberHistoryCondition {
	m := jsonFamilyMemberHistoryCondition{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Outcome = r.Outcome
	m.ContributedToDeath = r.ContributedToDeath
	if r.ContributedToDeath != nil && (r.ContributedToDeath.Id != nil || r.ContributedToDeath.Extension != nil) {
		m.ContributedToDeathPrimitiveElement = &primitiveElement{Id: r.ContributedToDeath.Id, Extension: r.ContributedToDeath.Extension}
	}
	switch v := r.Onset.(type) {
	case Age:
		m.OnsetAge = &v
	case *Age:
		m.OnsetAge = v
	case Range:
		m.OnsetRange = &v
	case *Range:
		m.OnsetRange = v
	case Period:
		m.OnsetPeriod = &v
	case *Period:
		m.OnsetPeriod = v
	case String:
		m.OnsetString = &v
		if v.Id != nil || v.Extension != nil {
			m.OnsetStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.OnsetString = v
		if v.Id != nil || v.Extension != nil {
			m.OnsetStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.Note = r.Note
	return m
}
func (r *FamilyMemberHistoryCondition) UnmarshalJSON(b []byte) error {
	var m jsonFamilyMemberHistoryCondition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *FamilyMemberHistoryCondition) unmarshalJSON(m jsonFamilyMemberHistoryCondition) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Outcome = m.Outcome
	r.ContributedToDeath = m.ContributedToDeath
	if m.ContributedToDeathPrimitiveElement != nil {
		r.ContributedToDeath.Id = m.ContributedToDeathPrimitiveElement.Id
		r.ContributedToDeath.Extension = m.ContributedToDeathPrimitiveElement.Extension
	}
	if m.OnsetAge != nil {
		if r.Onset != nil {
			return fmt.Errorf("multiple values for field \"Onset\"")
		}
		v := m.OnsetAge
		r.Onset = v
	}
	if m.OnsetRange != nil {
		if r.Onset != nil {
			return fmt.Errorf("multiple values for field \"Onset\"")
		}
		v := m.OnsetRange
		r.Onset = v
	}
	if m.OnsetPeriod != nil {
		if r.Onset != nil {
			return fmt.Errorf("multiple values for field \"Onset\"")
		}
		v := m.OnsetPeriod
		r.Onset = v
	}
	if m.OnsetString != nil || m.OnsetStringPrimitiveElement != nil {
		if r.Onset != nil {
			return fmt.Errorf("multiple values for field \"Onset\"")
		}
		v := m.OnsetString
		if m.OnsetStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.OnsetStringPrimitiveElement.Id
			v.Extension = m.OnsetStringPrimitiveElement.Extension
		}
		r.Onset = v
	}
	r.Note = m.Note
	return nil
}
func (r FamilyMemberHistoryCondition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
