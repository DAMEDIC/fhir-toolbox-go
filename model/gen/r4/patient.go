package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// Demographics and other administrative information about an individual or animal receiving care or other health-related services.
//
// Tracking patient is the center of the healthcare process.
type Patient struct {
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
	// An identifier for this patient.
	Identifier []Identifier
	// Whether this patient record is in active use.
	// Many systems use this property to mark as non-current patients, such as those that have not been seen for a period of time based on an organization's business rules.
	//
	// # It is often used to filter patient lists to exclude inactive patients
	//
	// Deceased patients may also be marked as inactive for the same reasons, but may be active for some time after death.
	Active *Boolean
	// A name associated with the individual.
	Name []HumanName
	// A contact detail (e.g. a telephone number or an email address) by which the individual may be contacted.
	Telecom []ContactPoint
	// Administrative Gender - the gender that the patient is considered to have for administration and record keeping purposes.
	Gender *Code
	// The date of birth for the individual.
	BirthDate *Date
	// Indicates if the individual is deceased or not.
	Deceased isPatientDeceased
	// An address for the individual.
	Address []Address
	// This field contains a patient's most recent marital (civil) status.
	MaritalStatus *CodeableConcept
	// Indicates whether the patient is part of a multiple (boolean) or indicates the actual birth order (integer).
	MultipleBirth isPatientMultipleBirth
	// Image of the patient.
	Photo []Attachment
	// A contact party (e.g. guardian, partner, friend) for the patient.
	Contact []PatientContact
	// A language which may be used to communicate with the patient about his or her health.
	Communication []PatientCommunication
	// Patient's nominated care provider.
	GeneralPractitioner []Reference
	// Organization that is the custodian of the patient record.
	ManagingOrganization *Reference
	// Link to another patient resource that concerns the same actual patient.
	Link []PatientLink
}

func (r Patient) ResourceType() string {
	return "Patient"
}
func (r Patient) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isPatientDeceased interface {
	isPatientDeceased()
}

func (r Boolean) isPatientDeceased()  {}
func (r DateTime) isPatientDeceased() {}

type isPatientMultipleBirth interface {
	isPatientMultipleBirth()
}

func (r Boolean) isPatientMultipleBirth() {}
func (r Integer) isPatientMultipleBirth() {}

type jsonPatient struct {
	ResourceType                         string                 `json:"resourceType"`
	Id                                   *Id                    `json:"id,omitempty"`
	IdPrimitiveElement                   *primitiveElement      `json:"_id,omitempty"`
	Meta                                 *Meta                  `json:"meta,omitempty"`
	ImplicitRules                        *Uri                   `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement        *primitiveElement      `json:"_implicitRules,omitempty"`
	Language                             *Code                  `json:"language,omitempty"`
	LanguagePrimitiveElement             *primitiveElement      `json:"_language,omitempty"`
	Text                                 *Narrative             `json:"text,omitempty"`
	Contained                            []ContainedResource    `json:"contained,omitempty"`
	Extension                            []Extension            `json:"extension,omitempty"`
	ModifierExtension                    []Extension            `json:"modifierExtension,omitempty"`
	Identifier                           []Identifier           `json:"identifier,omitempty"`
	Active                               *Boolean               `json:"active,omitempty"`
	ActivePrimitiveElement               *primitiveElement      `json:"_active,omitempty"`
	Name                                 []HumanName            `json:"name,omitempty"`
	Telecom                              []ContactPoint         `json:"telecom,omitempty"`
	Gender                               *Code                  `json:"gender,omitempty"`
	GenderPrimitiveElement               *primitiveElement      `json:"_gender,omitempty"`
	BirthDate                            *Date                  `json:"birthDate,omitempty"`
	BirthDatePrimitiveElement            *primitiveElement      `json:"_birthDate,omitempty"`
	DeceasedBoolean                      *Boolean               `json:"deceasedBoolean,omitempty"`
	DeceasedBooleanPrimitiveElement      *primitiveElement      `json:"_deceasedBoolean,omitempty"`
	DeceasedDateTime                     *DateTime              `json:"deceasedDateTime,omitempty"`
	DeceasedDateTimePrimitiveElement     *primitiveElement      `json:"_deceasedDateTime,omitempty"`
	Address                              []Address              `json:"address,omitempty"`
	MaritalStatus                        *CodeableConcept       `json:"maritalStatus,omitempty"`
	MultipleBirthBoolean                 *Boolean               `json:"multipleBirthBoolean,omitempty"`
	MultipleBirthBooleanPrimitiveElement *primitiveElement      `json:"_multipleBirthBoolean,omitempty"`
	MultipleBirthInteger                 *Integer               `json:"multipleBirthInteger,omitempty"`
	MultipleBirthIntegerPrimitiveElement *primitiveElement      `json:"_multipleBirthInteger,omitempty"`
	Photo                                []Attachment           `json:"photo,omitempty"`
	Contact                              []PatientContact       `json:"contact,omitempty"`
	Communication                        []PatientCommunication `json:"communication,omitempty"`
	GeneralPractitioner                  []Reference            `json:"generalPractitioner,omitempty"`
	ManagingOrganization                 *Reference             `json:"managingOrganization,omitempty"`
	Link                                 []PatientLink          `json:"link,omitempty"`
}

func (r Patient) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Patient) marshalJSON() jsonPatient {
	m := jsonPatient{}
	m.ResourceType = "Patient"
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
	if r.Active != nil && r.Active.Value != nil {
		m.Active = r.Active
	}
	if r.Active != nil && (r.Active.Id != nil || r.Active.Extension != nil) {
		m.ActivePrimitiveElement = &primitiveElement{Id: r.Active.Id, Extension: r.Active.Extension}
	}
	m.Name = r.Name
	m.Telecom = r.Telecom
	if r.Gender != nil && r.Gender.Value != nil {
		m.Gender = r.Gender
	}
	if r.Gender != nil && (r.Gender.Id != nil || r.Gender.Extension != nil) {
		m.GenderPrimitiveElement = &primitiveElement{Id: r.Gender.Id, Extension: r.Gender.Extension}
	}
	if r.BirthDate != nil && r.BirthDate.Value != nil {
		m.BirthDate = r.BirthDate
	}
	if r.BirthDate != nil && (r.BirthDate.Id != nil || r.BirthDate.Extension != nil) {
		m.BirthDatePrimitiveElement = &primitiveElement{Id: r.BirthDate.Id, Extension: r.BirthDate.Extension}
	}
	switch v := r.Deceased.(type) {
	case Boolean:
		if v.Value != nil {
			m.DeceasedBoolean = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.DeceasedBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		if v.Value != nil {
			m.DeceasedBoolean = v
		}
		if v.Id != nil || v.Extension != nil {
			m.DeceasedBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case DateTime:
		if v.Value != nil {
			m.DeceasedDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.DeceasedDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.DeceasedDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.DeceasedDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.Address = r.Address
	m.MaritalStatus = r.MaritalStatus
	switch v := r.MultipleBirth.(type) {
	case Boolean:
		if v.Value != nil {
			m.MultipleBirthBoolean = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.MultipleBirthBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		if v.Value != nil {
			m.MultipleBirthBoolean = v
		}
		if v.Id != nil || v.Extension != nil {
			m.MultipleBirthBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Integer:
		if v.Value != nil {
			m.MultipleBirthInteger = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.MultipleBirthIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Integer:
		if v.Value != nil {
			m.MultipleBirthInteger = v
		}
		if v.Id != nil || v.Extension != nil {
			m.MultipleBirthIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.Photo = r.Photo
	m.Contact = r.Contact
	m.Communication = r.Communication
	m.GeneralPractitioner = r.GeneralPractitioner
	m.ManagingOrganization = r.ManagingOrganization
	m.Link = r.Link
	return m
}
func (r *Patient) UnmarshalJSON(b []byte) error {
	var m jsonPatient
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Patient) unmarshalJSON(m jsonPatient) error {
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
	r.Active = m.Active
	if m.ActivePrimitiveElement != nil {
		if r.Active == nil {
			r.Active = &Boolean{}
		}
		r.Active.Id = m.ActivePrimitiveElement.Id
		r.Active.Extension = m.ActivePrimitiveElement.Extension
	}
	r.Name = m.Name
	r.Telecom = m.Telecom
	r.Gender = m.Gender
	if m.GenderPrimitiveElement != nil {
		if r.Gender == nil {
			r.Gender = &Code{}
		}
		r.Gender.Id = m.GenderPrimitiveElement.Id
		r.Gender.Extension = m.GenderPrimitiveElement.Extension
	}
	r.BirthDate = m.BirthDate
	if m.BirthDatePrimitiveElement != nil {
		if r.BirthDate == nil {
			r.BirthDate = &Date{}
		}
		r.BirthDate.Id = m.BirthDatePrimitiveElement.Id
		r.BirthDate.Extension = m.BirthDatePrimitiveElement.Extension
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
	if m.DeceasedDateTime != nil || m.DeceasedDateTimePrimitiveElement != nil {
		if r.Deceased != nil {
			return fmt.Errorf("multiple values for field \"Deceased\"")
		}
		v := m.DeceasedDateTime
		if m.DeceasedDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.DeceasedDateTimePrimitiveElement.Id
			v.Extension = m.DeceasedDateTimePrimitiveElement.Extension
		}
		r.Deceased = v
	}
	r.Address = m.Address
	r.MaritalStatus = m.MaritalStatus
	if m.MultipleBirthBoolean != nil || m.MultipleBirthBooleanPrimitiveElement != nil {
		if r.MultipleBirth != nil {
			return fmt.Errorf("multiple values for field \"MultipleBirth\"")
		}
		v := m.MultipleBirthBoolean
		if m.MultipleBirthBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.MultipleBirthBooleanPrimitiveElement.Id
			v.Extension = m.MultipleBirthBooleanPrimitiveElement.Extension
		}
		r.MultipleBirth = v
	}
	if m.MultipleBirthInteger != nil || m.MultipleBirthIntegerPrimitiveElement != nil {
		if r.MultipleBirth != nil {
			return fmt.Errorf("multiple values for field \"MultipleBirth\"")
		}
		v := m.MultipleBirthInteger
		if m.MultipleBirthIntegerPrimitiveElement != nil {
			if v == nil {
				v = &Integer{}
			}
			v.Id = m.MultipleBirthIntegerPrimitiveElement.Id
			v.Extension = m.MultipleBirthIntegerPrimitiveElement.Extension
		}
		r.MultipleBirth = v
	}
	r.Photo = m.Photo
	r.Contact = m.Contact
	r.Communication = m.Communication
	r.GeneralPractitioner = m.GeneralPractitioner
	r.ManagingOrganization = m.ManagingOrganization
	r.Link = m.Link
	return nil
}
func (r Patient) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Patient"
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
	err = e.EncodeElement(r.Active, xml.StartElement{Name: xml.Name{Local: "active"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Telecom, xml.StartElement{Name: xml.Name{Local: "telecom"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Gender, xml.StartElement{Name: xml.Name{Local: "gender"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BirthDate, xml.StartElement{Name: xml.Name{Local: "birthDate"}})
	if err != nil {
		return err
	}
	switch v := r.Deceased.(type) {
	case Boolean, *Boolean:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "deceasedBoolean"}})
		if err != nil {
			return err
		}
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "deceasedDateTime"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Address, xml.StartElement{Name: xml.Name{Local: "address"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MaritalStatus, xml.StartElement{Name: xml.Name{Local: "maritalStatus"}})
	if err != nil {
		return err
	}
	switch v := r.MultipleBirth.(type) {
	case Boolean, *Boolean:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "multipleBirthBoolean"}})
		if err != nil {
			return err
		}
	case Integer, *Integer:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "multipleBirthInteger"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Photo, xml.StartElement{Name: xml.Name{Local: "photo"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Contact, xml.StartElement{Name: xml.Name{Local: "contact"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Communication, xml.StartElement{Name: xml.Name{Local: "communication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.GeneralPractitioner, xml.StartElement{Name: xml.Name{Local: "generalPractitioner"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ManagingOrganization, xml.StartElement{Name: xml.Name{Local: "managingOrganization"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Link, xml.StartElement{Name: xml.Name{Local: "link"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Patient) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "active":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Active = &v
			case "name":
				var v HumanName
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = append(r.Name, v)
			case "telecom":
				var v ContactPoint
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Telecom = append(r.Telecom, v)
			case "gender":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Gender = &v
			case "birthDate":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BirthDate = &v
			case "deceasedBoolean":
				if r.Deceased != nil {
					return fmt.Errorf("multiple values for field \"Deceased\"")
				}
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Deceased = v
			case "deceasedDateTime":
				if r.Deceased != nil {
					return fmt.Errorf("multiple values for field \"Deceased\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Deceased = v
			case "address":
				var v Address
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Address = append(r.Address, v)
			case "maritalStatus":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MaritalStatus = &v
			case "multipleBirthBoolean":
				if r.MultipleBirth != nil {
					return fmt.Errorf("multiple values for field \"MultipleBirth\"")
				}
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MultipleBirth = v
			case "multipleBirthInteger":
				if r.MultipleBirth != nil {
					return fmt.Errorf("multiple values for field \"MultipleBirth\"")
				}
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MultipleBirth = v
			case "photo":
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Photo = append(r.Photo, v)
			case "contact":
				var v PatientContact
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Contact = append(r.Contact, v)
			case "communication":
				var v PatientCommunication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Communication = append(r.Communication, v)
			case "generalPractitioner":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.GeneralPractitioner = append(r.GeneralPractitioner, v)
			case "managingOrganization":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ManagingOrganization = &v
			case "link":
				var v PatientLink
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Link = append(r.Link, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Patient) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A contact party (e.g. guardian, partner, friend) for the patient.
type PatientContact struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The nature of the relationship between the patient and the contact person.
	Relationship []CodeableConcept
	// A name associated with the contact person.
	Name *HumanName
	// A contact detail for the person, e.g. a telephone number or an email address.
	Telecom []ContactPoint
	// Address for the contact person.
	Address *Address
	// Administrative Gender - the gender that the contact person is considered to have for administration and record keeping purposes.
	Gender *Code
	// Organization on behalf of which the contact is acting or for which the contact is working.
	Organization *Reference
	// The period during which this contact person or organization is valid to be contacted relating to this patient.
	Period *Period
}
type jsonPatientContact struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Relationship           []CodeableConcept `json:"relationship,omitempty"`
	Name                   *HumanName        `json:"name,omitempty"`
	Telecom                []ContactPoint    `json:"telecom,omitempty"`
	Address                *Address          `json:"address,omitempty"`
	Gender                 *Code             `json:"gender,omitempty"`
	GenderPrimitiveElement *primitiveElement `json:"_gender,omitempty"`
	Organization           *Reference        `json:"organization,omitempty"`
	Period                 *Period           `json:"period,omitempty"`
}

func (r PatientContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PatientContact) marshalJSON() jsonPatientContact {
	m := jsonPatientContact{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Relationship = r.Relationship
	m.Name = r.Name
	m.Telecom = r.Telecom
	m.Address = r.Address
	if r.Gender != nil && r.Gender.Value != nil {
		m.Gender = r.Gender
	}
	if r.Gender != nil && (r.Gender.Id != nil || r.Gender.Extension != nil) {
		m.GenderPrimitiveElement = &primitiveElement{Id: r.Gender.Id, Extension: r.Gender.Extension}
	}
	m.Organization = r.Organization
	m.Period = r.Period
	return m
}
func (r *PatientContact) UnmarshalJSON(b []byte) error {
	var m jsonPatientContact
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PatientContact) unmarshalJSON(m jsonPatientContact) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Relationship = m.Relationship
	r.Name = m.Name
	r.Telecom = m.Telecom
	r.Address = m.Address
	r.Gender = m.Gender
	if m.GenderPrimitiveElement != nil {
		if r.Gender == nil {
			r.Gender = &Code{}
		}
		r.Gender.Id = m.GenderPrimitiveElement.Id
		r.Gender.Extension = m.GenderPrimitiveElement.Extension
	}
	r.Organization = m.Organization
	r.Period = m.Period
	return nil
}
func (r PatientContact) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Relationship, xml.StartElement{Name: xml.Name{Local: "relationship"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Telecom, xml.StartElement{Name: xml.Name{Local: "telecom"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Address, xml.StartElement{Name: xml.Name{Local: "address"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Gender, xml.StartElement{Name: xml.Name{Local: "gender"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Organization, xml.StartElement{Name: xml.Name{Local: "organization"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Period, xml.StartElement{Name: xml.Name{Local: "period"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *PatientContact) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "relationship":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Relationship = append(r.Relationship, v)
			case "name":
				var v HumanName
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "telecom":
				var v ContactPoint
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Telecom = append(r.Telecom, v)
			case "address":
				var v Address
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Address = &v
			case "gender":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Gender = &v
			case "organization":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Organization = &v
			case "period":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Period = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r PatientContact) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A language which may be used to communicate with the patient about his or her health.
type PatientCommunication struct {
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
type jsonPatientCommunication struct {
	Id                        *string           `json:"id,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	ModifierExtension         []Extension       `json:"modifierExtension,omitempty"`
	Language                  CodeableConcept   `json:"language,omitempty"`
	Preferred                 *Boolean          `json:"preferred,omitempty"`
	PreferredPrimitiveElement *primitiveElement `json:"_preferred,omitempty"`
}

func (r PatientCommunication) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PatientCommunication) marshalJSON() jsonPatientCommunication {
	m := jsonPatientCommunication{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Language = r.Language
	if r.Preferred != nil && r.Preferred.Value != nil {
		m.Preferred = r.Preferred
	}
	if r.Preferred != nil && (r.Preferred.Id != nil || r.Preferred.Extension != nil) {
		m.PreferredPrimitiveElement = &primitiveElement{Id: r.Preferred.Id, Extension: r.Preferred.Extension}
	}
	return m
}
func (r *PatientCommunication) UnmarshalJSON(b []byte) error {
	var m jsonPatientCommunication
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PatientCommunication) unmarshalJSON(m jsonPatientCommunication) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Language = m.Language
	r.Preferred = m.Preferred
	if m.PreferredPrimitiveElement != nil {
		if r.Preferred == nil {
			r.Preferred = &Boolean{}
		}
		r.Preferred.Id = m.PreferredPrimitiveElement.Id
		r.Preferred.Extension = m.PreferredPrimitiveElement.Extension
	}
	return nil
}
func (r PatientCommunication) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Preferred, xml.StartElement{Name: xml.Name{Local: "preferred"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *PatientCommunication) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "language":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = v
			case "preferred":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Preferred = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r PatientCommunication) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Link to another patient resource that concerns the same actual patient.
type PatientLink struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The other patient resource that the link refers to.
	Other Reference
	// The type of link between this patient resource and another patient resource.
	Type Code
}
type jsonPatientLink struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Other                Reference         `json:"other,omitempty"`
	Type                 Code              `json:"type,omitempty"`
	TypePrimitiveElement *primitiveElement `json:"_type,omitempty"`
}

func (r PatientLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PatientLink) marshalJSON() jsonPatientLink {
	m := jsonPatientLink{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Other = r.Other
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	return m
}
func (r *PatientLink) UnmarshalJSON(b []byte) error {
	var m jsonPatientLink
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PatientLink) unmarshalJSON(m jsonPatientLink) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Other = m.Other
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	return nil
}
func (r PatientLink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Other, xml.StartElement{Name: xml.Name{Local: "other"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *PatientLink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "other":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Other = v
			case "type":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r PatientLink) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
