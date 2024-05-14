package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// Information about a person that is involved in the care for a patient, but who is not the target of healthcare, nor has a formal responsibility in the care process.
//
// Need to track persons related to the patient or the healthcare process.
type RelatedPerson struct {
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
	// Identifier for a person within a particular scope.
	Identifier []Identifier
	// Whether this related person record is in active use.
	Active *Boolean
	// The patient this person is related to.
	Patient Reference
	// The nature of the relationship between a patient and the related person.
	Relationship []CodeableConcept
	// A name associated with the person.
	Name []HumanName
	// A contact detail for the person, e.g. a telephone number or an email address.
	Telecom []ContactPoint
	// Administrative Gender - the gender that the person is considered to have for administration and record keeping purposes.
	Gender *Code
	// The date on which the related person was born.
	BirthDate *Date
	// Address where the related person can be contacted or visited.
	Address []Address
	// Image of the person.
	Photo []Attachment
	// The period of time during which this relationship is or was active. If there are no dates defined, then the interval is unknown.
	Period *Period
	// A language which may be used to communicate with about the patient's health.
	Communication []RelatedPersonCommunication
}

func (r RelatedPerson) ResourceType() string {
	return "RelatedPerson"
}

type jsonRelatedPerson struct {
	ResourceType                  string                       `json:"resourceType"`
	Id                            *Id                          `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement            `json:"_id,omitempty"`
	Meta                          *Meta                        `json:"meta,omitempty"`
	ImplicitRules                 *Uri                         `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement            `json:"_implicitRules,omitempty"`
	Language                      *Code                        `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement            `json:"_language,omitempty"`
	Text                          *Narrative                   `json:"text,omitempty"`
	Contained                     []containedResource          `json:"contained,omitempty"`
	Extension                     []Extension                  `json:"extension,omitempty"`
	ModifierExtension             []Extension                  `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier                 `json:"identifier,omitempty"`
	Active                        *Boolean                     `json:"active,omitempty"`
	ActivePrimitiveElement        *primitiveElement            `json:"_active,omitempty"`
	Patient                       Reference                    `json:"patient,omitempty"`
	Relationship                  []CodeableConcept            `json:"relationship,omitempty"`
	Name                          []HumanName                  `json:"name,omitempty"`
	Telecom                       []ContactPoint               `json:"telecom,omitempty"`
	Gender                        *Code                        `json:"gender,omitempty"`
	GenderPrimitiveElement        *primitiveElement            `json:"_gender,omitempty"`
	BirthDate                     *Date                        `json:"birthDate,omitempty"`
	BirthDatePrimitiveElement     *primitiveElement            `json:"_birthDate,omitempty"`
	Address                       []Address                    `json:"address,omitempty"`
	Photo                         []Attachment                 `json:"photo,omitempty"`
	Period                        *Period                      `json:"period,omitempty"`
	Communication                 []RelatedPersonCommunication `json:"communication,omitempty"`
}

func (r RelatedPerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RelatedPerson) marshalJSON() jsonRelatedPerson {
	m := jsonRelatedPerson{}
	m.ResourceType = "RelatedPerson"
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
	m.Active = r.Active
	if r.Active != nil && (r.Active.Id != nil || r.Active.Extension != nil) {
		m.ActivePrimitiveElement = &primitiveElement{Id: r.Active.Id, Extension: r.Active.Extension}
	}
	m.Patient = r.Patient
	m.Relationship = r.Relationship
	m.Name = r.Name
	m.Telecom = r.Telecom
	m.Gender = r.Gender
	if r.Gender != nil && (r.Gender.Id != nil || r.Gender.Extension != nil) {
		m.GenderPrimitiveElement = &primitiveElement{Id: r.Gender.Id, Extension: r.Gender.Extension}
	}
	m.BirthDate = r.BirthDate
	if r.BirthDate != nil && (r.BirthDate.Id != nil || r.BirthDate.Extension != nil) {
		m.BirthDatePrimitiveElement = &primitiveElement{Id: r.BirthDate.Id, Extension: r.BirthDate.Extension}
	}
	m.Address = r.Address
	m.Photo = r.Photo
	m.Period = r.Period
	m.Communication = r.Communication
	return m
}
func (r *RelatedPerson) UnmarshalJSON(b []byte) error {
	var m jsonRelatedPerson
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RelatedPerson) unmarshalJSON(m jsonRelatedPerson) error {
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
	r.Active = m.Active
	if m.ActivePrimitiveElement != nil {
		r.Active.Id = m.ActivePrimitiveElement.Id
		r.Active.Extension = m.ActivePrimitiveElement.Extension
	}
	r.Patient = m.Patient
	r.Relationship = m.Relationship
	r.Name = m.Name
	r.Telecom = m.Telecom
	r.Gender = m.Gender
	if m.GenderPrimitiveElement != nil {
		r.Gender.Id = m.GenderPrimitiveElement.Id
		r.Gender.Extension = m.GenderPrimitiveElement.Extension
	}
	r.BirthDate = m.BirthDate
	if m.BirthDatePrimitiveElement != nil {
		r.BirthDate.Id = m.BirthDatePrimitiveElement.Id
		r.BirthDate.Extension = m.BirthDatePrimitiveElement.Extension
	}
	r.Address = m.Address
	r.Photo = m.Photo
	r.Period = m.Period
	r.Communication = m.Communication
	return nil
}
func (r RelatedPerson) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A language which may be used to communicate with about the patient's health.
type RelatedPersonCommunication struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The ISO-639-1 alpha 2 code in lower case for the language, optionally followed by a hyphen and the ISO-3166-1 alpha 2 code for the region in upper case; e.g. "en" for English, or "en-US" for American English versus "en-EN" for England English.
	Language CodeableConcept
	// Indicates whether or not the patient prefers this language (over other languages he masters up a certain level).
	Preferred *Boolean
}
type jsonRelatedPersonCommunication struct {
	Id                        *string           `json:"id,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Language                  CodeableConcept   `json:"language,omitempty"`
	Preferred                 *Boolean          `json:"preferred,omitempty"`
	PreferredPrimitiveElement *primitiveElement `json:"_preferred,omitempty"`
}

func (r RelatedPersonCommunication) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RelatedPersonCommunication) marshalJSON() jsonRelatedPersonCommunication {
	m := jsonRelatedPersonCommunication{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Language = r.Language
	m.Preferred = r.Preferred
	if r.Preferred != nil && (r.Preferred.Id != nil || r.Preferred.Extension != nil) {
		m.PreferredPrimitiveElement = &primitiveElement{Id: r.Preferred.Id, Extension: r.Preferred.Extension}
	}
	return m
}
func (r *RelatedPersonCommunication) UnmarshalJSON(b []byte) error {
	var m jsonRelatedPersonCommunication
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RelatedPersonCommunication) unmarshalJSON(m jsonRelatedPersonCommunication) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Language = m.Language
	r.Preferred = m.Preferred
	if m.PreferredPrimitiveElement != nil {
		r.Preferred.Id = m.PreferredPrimitiveElement.Id
		r.Preferred.Extension = m.PreferredPrimitiveElement.Extension
	}
	return nil
}
func (r RelatedPersonCommunication) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
