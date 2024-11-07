package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// The regulatory authorization of a medicinal product.
type MedicinalProductAuthorization struct {
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
	// Business identifier for the marketing authorization, as assigned by a regulator.
	Identifier []Identifier
	// The medicinal product that is being authorized.
	Subject *Reference
	// The country in which the marketing authorization has been granted.
	Country []CodeableConcept
	// Jurisdiction within a country.
	Jurisdiction []CodeableConcept
	// The status of the marketing authorization.
	Status *CodeableConcept
	// The date at which the given status has become applicable.
	StatusDate *DateTime
	// The date when a suspended the marketing or the marketing authorization of the product is anticipated to be restored.
	RestoreDate *DateTime
	// The beginning of the time period in which the marketing authorization is in the specific status shall be specified A complete date consisting of day, month and year shall be specified using the ISO 8601 date format.
	ValidityPeriod *Period
	// A period of time after authorization before generic product applicatiosn can be submitted.
	DataExclusivityPeriod *Period
	// The date when the first authorization was granted by a Medicines Regulatory Agency.
	DateOfFirstAuthorization *DateTime
	// Date of first marketing authorization for a company's new medicinal product in any country in the World.
	InternationalBirthDate *DateTime
	// The legal framework against which this authorization is granted.
	LegalBasis *CodeableConcept
	// Authorization in areas within a country.
	JurisdictionalAuthorization []MedicinalProductAuthorizationJurisdictionalAuthorization
	// Marketing Authorization Holder.
	Holder *Reference
	// Medicines Regulatory Agency.
	Regulator *Reference
	// The regulatory procedure for granting or amending a marketing authorization.
	Procedure *MedicinalProductAuthorizationProcedure
}

func (r MedicinalProductAuthorization) ResourceType() string {
	return "MedicinalProductAuthorization"
}
func (r MedicinalProductAuthorization) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonMedicinalProductAuthorization struct {
	ResourceType                             string                                                     `json:"resourceType"`
	Id                                       *Id                                                        `json:"id,omitempty"`
	IdPrimitiveElement                       *primitiveElement                                          `json:"_id,omitempty"`
	Meta                                     *Meta                                                      `json:"meta,omitempty"`
	ImplicitRules                            *Uri                                                       `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement            *primitiveElement                                          `json:"_implicitRules,omitempty"`
	Language                                 *Code                                                      `json:"language,omitempty"`
	LanguagePrimitiveElement                 *primitiveElement                                          `json:"_language,omitempty"`
	Text                                     *Narrative                                                 `json:"text,omitempty"`
	Contained                                []ContainedResource                                        `json:"contained,omitempty"`
	Extension                                []Extension                                                `json:"extension,omitempty"`
	ModifierExtension                        []Extension                                                `json:"modifierExtension,omitempty"`
	Identifier                               []Identifier                                               `json:"identifier,omitempty"`
	Subject                                  *Reference                                                 `json:"subject,omitempty"`
	Country                                  []CodeableConcept                                          `json:"country,omitempty"`
	Jurisdiction                             []CodeableConcept                                          `json:"jurisdiction,omitempty"`
	Status                                   *CodeableConcept                                           `json:"status,omitempty"`
	StatusDate                               *DateTime                                                  `json:"statusDate,omitempty"`
	StatusDatePrimitiveElement               *primitiveElement                                          `json:"_statusDate,omitempty"`
	RestoreDate                              *DateTime                                                  `json:"restoreDate,omitempty"`
	RestoreDatePrimitiveElement              *primitiveElement                                          `json:"_restoreDate,omitempty"`
	ValidityPeriod                           *Period                                                    `json:"validityPeriod,omitempty"`
	DataExclusivityPeriod                    *Period                                                    `json:"dataExclusivityPeriod,omitempty"`
	DateOfFirstAuthorization                 *DateTime                                                  `json:"dateOfFirstAuthorization,omitempty"`
	DateOfFirstAuthorizationPrimitiveElement *primitiveElement                                          `json:"_dateOfFirstAuthorization,omitempty"`
	InternationalBirthDate                   *DateTime                                                  `json:"internationalBirthDate,omitempty"`
	InternationalBirthDatePrimitiveElement   *primitiveElement                                          `json:"_internationalBirthDate,omitempty"`
	LegalBasis                               *CodeableConcept                                           `json:"legalBasis,omitempty"`
	JurisdictionalAuthorization              []MedicinalProductAuthorizationJurisdictionalAuthorization `json:"jurisdictionalAuthorization,omitempty"`
	Holder                                   *Reference                                                 `json:"holder,omitempty"`
	Regulator                                *Reference                                                 `json:"regulator,omitempty"`
	Procedure                                *MedicinalProductAuthorizationProcedure                    `json:"procedure,omitempty"`
}

func (r MedicinalProductAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductAuthorization) marshalJSON() jsonMedicinalProductAuthorization {
	m := jsonMedicinalProductAuthorization{}
	m.ResourceType = "MedicinalProductAuthorization"
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
	m.Subject = r.Subject
	m.Country = r.Country
	m.Jurisdiction = r.Jurisdiction
	m.Status = r.Status
	if r.StatusDate != nil && r.StatusDate.Value != nil {
		m.StatusDate = r.StatusDate
	}
	if r.StatusDate != nil && (r.StatusDate.Id != nil || r.StatusDate.Extension != nil) {
		m.StatusDatePrimitiveElement = &primitiveElement{Id: r.StatusDate.Id, Extension: r.StatusDate.Extension}
	}
	if r.RestoreDate != nil && r.RestoreDate.Value != nil {
		m.RestoreDate = r.RestoreDate
	}
	if r.RestoreDate != nil && (r.RestoreDate.Id != nil || r.RestoreDate.Extension != nil) {
		m.RestoreDatePrimitiveElement = &primitiveElement{Id: r.RestoreDate.Id, Extension: r.RestoreDate.Extension}
	}
	m.ValidityPeriod = r.ValidityPeriod
	m.DataExclusivityPeriod = r.DataExclusivityPeriod
	if r.DateOfFirstAuthorization != nil && r.DateOfFirstAuthorization.Value != nil {
		m.DateOfFirstAuthorization = r.DateOfFirstAuthorization
	}
	if r.DateOfFirstAuthorization != nil && (r.DateOfFirstAuthorization.Id != nil || r.DateOfFirstAuthorization.Extension != nil) {
		m.DateOfFirstAuthorizationPrimitiveElement = &primitiveElement{Id: r.DateOfFirstAuthorization.Id, Extension: r.DateOfFirstAuthorization.Extension}
	}
	if r.InternationalBirthDate != nil && r.InternationalBirthDate.Value != nil {
		m.InternationalBirthDate = r.InternationalBirthDate
	}
	if r.InternationalBirthDate != nil && (r.InternationalBirthDate.Id != nil || r.InternationalBirthDate.Extension != nil) {
		m.InternationalBirthDatePrimitiveElement = &primitiveElement{Id: r.InternationalBirthDate.Id, Extension: r.InternationalBirthDate.Extension}
	}
	m.LegalBasis = r.LegalBasis
	m.JurisdictionalAuthorization = r.JurisdictionalAuthorization
	m.Holder = r.Holder
	m.Regulator = r.Regulator
	m.Procedure = r.Procedure
	return m
}
func (r *MedicinalProductAuthorization) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductAuthorization
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductAuthorization) unmarshalJSON(m jsonMedicinalProductAuthorization) error {
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
	r.Subject = m.Subject
	r.Country = m.Country
	r.Jurisdiction = m.Jurisdiction
	r.Status = m.Status
	r.StatusDate = m.StatusDate
	if m.StatusDatePrimitiveElement != nil {
		if r.StatusDate == nil {
			r.StatusDate = &DateTime{}
		}
		r.StatusDate.Id = m.StatusDatePrimitiveElement.Id
		r.StatusDate.Extension = m.StatusDatePrimitiveElement.Extension
	}
	r.RestoreDate = m.RestoreDate
	if m.RestoreDatePrimitiveElement != nil {
		if r.RestoreDate == nil {
			r.RestoreDate = &DateTime{}
		}
		r.RestoreDate.Id = m.RestoreDatePrimitiveElement.Id
		r.RestoreDate.Extension = m.RestoreDatePrimitiveElement.Extension
	}
	r.ValidityPeriod = m.ValidityPeriod
	r.DataExclusivityPeriod = m.DataExclusivityPeriod
	r.DateOfFirstAuthorization = m.DateOfFirstAuthorization
	if m.DateOfFirstAuthorizationPrimitiveElement != nil {
		if r.DateOfFirstAuthorization == nil {
			r.DateOfFirstAuthorization = &DateTime{}
		}
		r.DateOfFirstAuthorization.Id = m.DateOfFirstAuthorizationPrimitiveElement.Id
		r.DateOfFirstAuthorization.Extension = m.DateOfFirstAuthorizationPrimitiveElement.Extension
	}
	r.InternationalBirthDate = m.InternationalBirthDate
	if m.InternationalBirthDatePrimitiveElement != nil {
		if r.InternationalBirthDate == nil {
			r.InternationalBirthDate = &DateTime{}
		}
		r.InternationalBirthDate.Id = m.InternationalBirthDatePrimitiveElement.Id
		r.InternationalBirthDate.Extension = m.InternationalBirthDatePrimitiveElement.Extension
	}
	r.LegalBasis = m.LegalBasis
	r.JurisdictionalAuthorization = m.JurisdictionalAuthorization
	r.Holder = m.Holder
	r.Regulator = m.Regulator
	r.Procedure = m.Procedure
	return nil
}
func (r MedicinalProductAuthorization) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "MedicinalProductAuthorization"
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
	err = e.EncodeElement(r.Subject, xml.StartElement{Name: xml.Name{Local: "subject"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Country, xml.StartElement{Name: xml.Name{Local: "country"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Jurisdiction, xml.StartElement{Name: xml.Name{Local: "jurisdiction"}})
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
	err = e.EncodeElement(r.RestoreDate, xml.StartElement{Name: xml.Name{Local: "restoreDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ValidityPeriod, xml.StartElement{Name: xml.Name{Local: "validityPeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DataExclusivityPeriod, xml.StartElement{Name: xml.Name{Local: "dataExclusivityPeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DateOfFirstAuthorization, xml.StartElement{Name: xml.Name{Local: "dateOfFirstAuthorization"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.InternationalBirthDate, xml.StartElement{Name: xml.Name{Local: "internationalBirthDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LegalBasis, xml.StartElement{Name: xml.Name{Local: "legalBasis"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.JurisdictionalAuthorization, xml.StartElement{Name: xml.Name{Local: "jurisdictionalAuthorization"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Holder, xml.StartElement{Name: xml.Name{Local: "holder"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Regulator, xml.StartElement{Name: xml.Name{Local: "regulator"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Procedure, xml.StartElement{Name: xml.Name{Local: "procedure"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductAuthorization) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "subject":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = &v
			case "country":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Country = append(r.Country, v)
			case "jurisdiction":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Jurisdiction = append(r.Jurisdiction, v)
			case "status":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = &v
			case "statusDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StatusDate = &v
			case "restoreDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RestoreDate = &v
			case "validityPeriod":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ValidityPeriod = &v
			case "dataExclusivityPeriod":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DataExclusivityPeriod = &v
			case "dateOfFirstAuthorization":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DateOfFirstAuthorization = &v
			case "internationalBirthDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.InternationalBirthDate = &v
			case "legalBasis":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LegalBasis = &v
			case "jurisdictionalAuthorization":
				var v MedicinalProductAuthorizationJurisdictionalAuthorization
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.JurisdictionalAuthorization = append(r.JurisdictionalAuthorization, v)
			case "holder":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Holder = &v
			case "regulator":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Regulator = &v
			case "procedure":
				var v MedicinalProductAuthorizationProcedure
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Procedure = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductAuthorization) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Authorization in areas within a country.
type MedicinalProductAuthorizationJurisdictionalAuthorization struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The assigned number for the marketing authorization.
	Identifier []Identifier
	// Country of authorization.
	Country *CodeableConcept
	// Jurisdiction within a country.
	Jurisdiction []CodeableConcept
	// The legal status of supply in a jurisdiction or region.
	LegalStatusOfSupply *CodeableConcept
	// The start and expected end date of the authorization.
	ValidityPeriod *Period
}
type jsonMedicinalProductAuthorizationJurisdictionalAuthorization struct {
	Id                  *string           `json:"id,omitempty"`
	Extension           []Extension       `json:"extension,omitempty"`
	ModifierExtension   []Extension       `json:"modifierExtension,omitempty"`
	Identifier          []Identifier      `json:"identifier,omitempty"`
	Country             *CodeableConcept  `json:"country,omitempty"`
	Jurisdiction        []CodeableConcept `json:"jurisdiction,omitempty"`
	LegalStatusOfSupply *CodeableConcept  `json:"legalStatusOfSupply,omitempty"`
	ValidityPeriod      *Period           `json:"validityPeriod,omitempty"`
}

func (r MedicinalProductAuthorizationJurisdictionalAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductAuthorizationJurisdictionalAuthorization) marshalJSON() jsonMedicinalProductAuthorizationJurisdictionalAuthorization {
	m := jsonMedicinalProductAuthorizationJurisdictionalAuthorization{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Country = r.Country
	m.Jurisdiction = r.Jurisdiction
	m.LegalStatusOfSupply = r.LegalStatusOfSupply
	m.ValidityPeriod = r.ValidityPeriod
	return m
}
func (r *MedicinalProductAuthorizationJurisdictionalAuthorization) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductAuthorizationJurisdictionalAuthorization
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductAuthorizationJurisdictionalAuthorization) unmarshalJSON(m jsonMedicinalProductAuthorizationJurisdictionalAuthorization) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Country = m.Country
	r.Jurisdiction = m.Jurisdiction
	r.LegalStatusOfSupply = m.LegalStatusOfSupply
	r.ValidityPeriod = m.ValidityPeriod
	return nil
}
func (r MedicinalProductAuthorizationJurisdictionalAuthorization) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Country, xml.StartElement{Name: xml.Name{Local: "country"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Jurisdiction, xml.StartElement{Name: xml.Name{Local: "jurisdiction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LegalStatusOfSupply, xml.StartElement{Name: xml.Name{Local: "legalStatusOfSupply"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ValidityPeriod, xml.StartElement{Name: xml.Name{Local: "validityPeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductAuthorizationJurisdictionalAuthorization) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "country":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Country = &v
			case "jurisdiction":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Jurisdiction = append(r.Jurisdiction, v)
			case "legalStatusOfSupply":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LegalStatusOfSupply = &v
			case "validityPeriod":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ValidityPeriod = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductAuthorizationJurisdictionalAuthorization) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The regulatory procedure for granting or amending a marketing authorization.
type MedicinalProductAuthorizationProcedure struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identifier for this procedure.
	Identifier *Identifier
	// Type of procedure.
	Type CodeableConcept
	// Date of procedure.
	Date isMedicinalProductAuthorizationProcedureDate
	// Applcations submitted to obtain a marketing authorization.
	Application []MedicinalProductAuthorizationProcedure
}
type isMedicinalProductAuthorizationProcedureDate interface {
	isMedicinalProductAuthorizationProcedureDate()
}

func (r Period) isMedicinalProductAuthorizationProcedureDate()   {}
func (r DateTime) isMedicinalProductAuthorizationProcedureDate() {}

type jsonMedicinalProductAuthorizationProcedure struct {
	Id                           *string                                  `json:"id,omitempty"`
	Extension                    []Extension                              `json:"extension,omitempty"`
	ModifierExtension            []Extension                              `json:"modifierExtension,omitempty"`
	Identifier                   *Identifier                              `json:"identifier,omitempty"`
	Type                         CodeableConcept                          `json:"type,omitempty"`
	DatePeriod                   *Period                                  `json:"datePeriod,omitempty"`
	DateDateTime                 *DateTime                                `json:"dateDateTime,omitempty"`
	DateDateTimePrimitiveElement *primitiveElement                        `json:"_dateDateTime,omitempty"`
	Application                  []MedicinalProductAuthorizationProcedure `json:"application,omitempty"`
}

func (r MedicinalProductAuthorizationProcedure) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductAuthorizationProcedure) marshalJSON() jsonMedicinalProductAuthorizationProcedure {
	m := jsonMedicinalProductAuthorizationProcedure{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Type = r.Type
	switch v := r.Date.(type) {
	case Period:
		m.DatePeriod = &v
	case *Period:
		m.DatePeriod = v
	case DateTime:
		if v.Value != nil {
			m.DateDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.DateDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.DateDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.DateDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.Application = r.Application
	return m
}
func (r *MedicinalProductAuthorizationProcedure) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductAuthorizationProcedure
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductAuthorizationProcedure) unmarshalJSON(m jsonMedicinalProductAuthorizationProcedure) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Type = m.Type
	if m.DatePeriod != nil {
		if r.Date != nil {
			return fmt.Errorf("multiple values for field \"Date\"")
		}
		v := m.DatePeriod
		r.Date = v
	}
	if m.DateDateTime != nil || m.DateDateTimePrimitiveElement != nil {
		if r.Date != nil {
			return fmt.Errorf("multiple values for field \"Date\"")
		}
		v := m.DateDateTime
		if m.DateDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.DateDateTimePrimitiveElement.Id
			v.Extension = m.DateDateTimePrimitiveElement.Extension
		}
		r.Date = v
	}
	r.Application = m.Application
	return nil
}
func (r MedicinalProductAuthorizationProcedure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	switch v := r.Date.(type) {
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "datePeriod"}})
		if err != nil {
			return err
		}
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "dateDateTime"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Application, xml.StartElement{Name: xml.Name{Local: "application"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductAuthorizationProcedure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = &v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "datePeriod":
				if r.Date != nil {
					return fmt.Errorf("multiple values for field \"Date\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = v
			case "dateDateTime":
				if r.Date != nil {
					return fmt.Errorf("multiple values for field \"Date\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = v
			case "application":
				var v MedicinalProductAuthorizationProcedure
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Application = append(r.Application, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductAuthorizationProcedure) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
