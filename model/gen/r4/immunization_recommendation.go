package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// A patient's point-in-time set of recommendations (i.e. forecasting) according to a published schedule with optional supporting justification.
type ImmunizationRecommendation struct {
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
	// A unique identifier assigned to this particular recommendation record.
	Identifier []Identifier
	// The patient the recommendation(s) are for.
	Patient Reference
	// The date the immunization recommendation(s) were created.
	Date DateTime
	// Indicates the authority who published the protocol (e.g. ACIP).
	Authority *Reference
	// Vaccine administration recommendations.
	Recommendation []ImmunizationRecommendationRecommendation
}

func (r ImmunizationRecommendation) ResourceType() string {
	return "ImmunizationRecommendation"
}
func (r ImmunizationRecommendation) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonImmunizationRecommendation struct {
	ResourceType                  string                                     `json:"resourceType"`
	Id                            *Id                                        `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement                          `json:"_id,omitempty"`
	Meta                          *Meta                                      `json:"meta,omitempty"`
	ImplicitRules                 *Uri                                       `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement                          `json:"_implicitRules,omitempty"`
	Language                      *Code                                      `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement                          `json:"_language,omitempty"`
	Text                          *Narrative                                 `json:"text,omitempty"`
	Contained                     []ContainedResource                        `json:"contained,omitempty"`
	Extension                     []Extension                                `json:"extension,omitempty"`
	ModifierExtension             []Extension                                `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier                               `json:"identifier,omitempty"`
	Patient                       Reference                                  `json:"patient,omitempty"`
	Date                          DateTime                                   `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement                          `json:"_date,omitempty"`
	Authority                     *Reference                                 `json:"authority,omitempty"`
	Recommendation                []ImmunizationRecommendationRecommendation `json:"recommendation,omitempty"`
}

func (r ImmunizationRecommendation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImmunizationRecommendation) marshalJSON() jsonImmunizationRecommendation {
	m := jsonImmunizationRecommendation{}
	m.ResourceType = "ImmunizationRecommendation"
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
	m.Patient = r.Patient
	if r.Date.Value != nil {
		m.Date = r.Date
	}
	if r.Date.Id != nil || r.Date.Extension != nil {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Authority = r.Authority
	m.Recommendation = r.Recommendation
	return m
}
func (r *ImmunizationRecommendation) UnmarshalJSON(b []byte) error {
	var m jsonImmunizationRecommendation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImmunizationRecommendation) unmarshalJSON(m jsonImmunizationRecommendation) error {
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
	r.Patient = m.Patient
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Authority = m.Authority
	r.Recommendation = m.Recommendation
	return nil
}
func (r ImmunizationRecommendation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "ImmunizationRecommendation"
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
	err = e.EncodeElement(r.Patient, xml.StartElement{Name: xml.Name{Local: "patient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Authority, xml.StartElement{Name: xml.Name{Local: "authority"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Recommendation, xml.StartElement{Name: xml.Name{Local: "recommendation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImmunizationRecommendation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "patient":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Patient = v
			case "date":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = v
			case "authority":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Authority = &v
			case "recommendation":
				var v ImmunizationRecommendationRecommendation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Recommendation = append(r.Recommendation, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImmunizationRecommendation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Vaccine administration recommendations.
type ImmunizationRecommendationRecommendation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Vaccine(s) or vaccine group that pertain to the recommendation.
	VaccineCode []CodeableConcept
	// The targeted disease for the recommendation.
	TargetDisease *CodeableConcept
	// Vaccine(s) which should not be used to fulfill the recommendation.
	ContraindicatedVaccineCode []CodeableConcept
	// Indicates the patient status with respect to the path to immunity for the target disease.
	ForecastStatus CodeableConcept
	// The reason for the assigned forecast status.
	ForecastReason []CodeableConcept
	// Vaccine date recommendations.  For example, earliest date to administer, latest date to administer, etc.
	DateCriterion []ImmunizationRecommendationRecommendationDateCriterion
	// Contains the description about the protocol under which the vaccine was administered.
	Description *String
	// One possible path to achieve presumed immunity against a disease - within the context of an authority.
	Series *String
	// Nominal position of the recommended dose in a series (e.g. dose 2 is the next recommended dose).
	DoseNumber isImmunizationRecommendationRecommendationDoseNumber
	// The recommended number of doses to achieve immunity.
	SeriesDoses isImmunizationRecommendationRecommendationSeriesDoses
	// Immunization event history and/or evaluation that supports the status and recommendation.
	SupportingImmunization []Reference
	// Patient Information that supports the status and recommendation.  This includes patient observations, adverse reactions and allergy/intolerance information.
	SupportingPatientInformation []Reference
}
type isImmunizationRecommendationRecommendationDoseNumber interface {
	isImmunizationRecommendationRecommendationDoseNumber()
}

func (r PositiveInt) isImmunizationRecommendationRecommendationDoseNumber() {}
func (r String) isImmunizationRecommendationRecommendationDoseNumber()      {}

type isImmunizationRecommendationRecommendationSeriesDoses interface {
	isImmunizationRecommendationRecommendationSeriesDoses()
}

func (r PositiveInt) isImmunizationRecommendationRecommendationSeriesDoses() {}
func (r String) isImmunizationRecommendationRecommendationSeriesDoses()      {}

type jsonImmunizationRecommendationRecommendation struct {
	Id                                     *string                                                 `json:"id,omitempty"`
	Extension                              []Extension                                             `json:"extension,omitempty"`
	ModifierExtension                      []Extension                                             `json:"modifierExtension,omitempty"`
	VaccineCode                            []CodeableConcept                                       `json:"vaccineCode,omitempty"`
	TargetDisease                          *CodeableConcept                                        `json:"targetDisease,omitempty"`
	ContraindicatedVaccineCode             []CodeableConcept                                       `json:"contraindicatedVaccineCode,omitempty"`
	ForecastStatus                         CodeableConcept                                         `json:"forecastStatus,omitempty"`
	ForecastReason                         []CodeableConcept                                       `json:"forecastReason,omitempty"`
	DateCriterion                          []ImmunizationRecommendationRecommendationDateCriterion `json:"dateCriterion,omitempty"`
	Description                            *String                                                 `json:"description,omitempty"`
	DescriptionPrimitiveElement            *primitiveElement                                       `json:"_description,omitempty"`
	Series                                 *String                                                 `json:"series,omitempty"`
	SeriesPrimitiveElement                 *primitiveElement                                       `json:"_series,omitempty"`
	DoseNumberPositiveInt                  *PositiveInt                                            `json:"doseNumberPositiveInt,omitempty"`
	DoseNumberPositiveIntPrimitiveElement  *primitiveElement                                       `json:"_doseNumberPositiveInt,omitempty"`
	DoseNumberString                       *String                                                 `json:"doseNumberString,omitempty"`
	DoseNumberStringPrimitiveElement       *primitiveElement                                       `json:"_doseNumberString,omitempty"`
	SeriesDosesPositiveInt                 *PositiveInt                                            `json:"seriesDosesPositiveInt,omitempty"`
	SeriesDosesPositiveIntPrimitiveElement *primitiveElement                                       `json:"_seriesDosesPositiveInt,omitempty"`
	SeriesDosesString                      *String                                                 `json:"seriesDosesString,omitempty"`
	SeriesDosesStringPrimitiveElement      *primitiveElement                                       `json:"_seriesDosesString,omitempty"`
	SupportingImmunization                 []Reference                                             `json:"supportingImmunization,omitempty"`
	SupportingPatientInformation           []Reference                                             `json:"supportingPatientInformation,omitempty"`
}

func (r ImmunizationRecommendationRecommendation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImmunizationRecommendationRecommendation) marshalJSON() jsonImmunizationRecommendationRecommendation {
	m := jsonImmunizationRecommendationRecommendation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.VaccineCode = r.VaccineCode
	m.TargetDisease = r.TargetDisease
	m.ContraindicatedVaccineCode = r.ContraindicatedVaccineCode
	m.ForecastStatus = r.ForecastStatus
	m.ForecastReason = r.ForecastReason
	m.DateCriterion = r.DateCriterion
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	if r.Series != nil && r.Series.Value != nil {
		m.Series = r.Series
	}
	if r.Series != nil && (r.Series.Id != nil || r.Series.Extension != nil) {
		m.SeriesPrimitiveElement = &primitiveElement{Id: r.Series.Id, Extension: r.Series.Extension}
	}
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
	m.SupportingImmunization = r.SupportingImmunization
	m.SupportingPatientInformation = r.SupportingPatientInformation
	return m
}
func (r *ImmunizationRecommendationRecommendation) UnmarshalJSON(b []byte) error {
	var m jsonImmunizationRecommendationRecommendation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImmunizationRecommendationRecommendation) unmarshalJSON(m jsonImmunizationRecommendationRecommendation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.VaccineCode = m.VaccineCode
	r.TargetDisease = m.TargetDisease
	r.ContraindicatedVaccineCode = m.ContraindicatedVaccineCode
	r.ForecastStatus = m.ForecastStatus
	r.ForecastReason = m.ForecastReason
	r.DateCriterion = m.DateCriterion
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Series = m.Series
	if m.SeriesPrimitiveElement != nil {
		if r.Series == nil {
			r.Series = &String{}
		}
		r.Series.Id = m.SeriesPrimitiveElement.Id
		r.Series.Extension = m.SeriesPrimitiveElement.Extension
	}
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
	r.SupportingImmunization = m.SupportingImmunization
	r.SupportingPatientInformation = m.SupportingPatientInformation
	return nil
}
func (r ImmunizationRecommendationRecommendation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.VaccineCode, xml.StartElement{Name: xml.Name{Local: "vaccineCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TargetDisease, xml.StartElement{Name: xml.Name{Local: "targetDisease"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ContraindicatedVaccineCode, xml.StartElement{Name: xml.Name{Local: "contraindicatedVaccineCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ForecastStatus, xml.StartElement{Name: xml.Name{Local: "forecastStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ForecastReason, xml.StartElement{Name: xml.Name{Local: "forecastReason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DateCriterion, xml.StartElement{Name: xml.Name{Local: "dateCriterion"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Series, xml.StartElement{Name: xml.Name{Local: "series"}})
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
	err = e.EncodeElement(r.SupportingImmunization, xml.StartElement{Name: xml.Name{Local: "supportingImmunization"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SupportingPatientInformation, xml.StartElement{Name: xml.Name{Local: "supportingPatientInformation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImmunizationRecommendationRecommendation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "vaccineCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.VaccineCode = append(r.VaccineCode, v)
			case "targetDisease":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TargetDisease = &v
			case "contraindicatedVaccineCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ContraindicatedVaccineCode = append(r.ContraindicatedVaccineCode, v)
			case "forecastStatus":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ForecastStatus = v
			case "forecastReason":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ForecastReason = append(r.ForecastReason, v)
			case "dateCriterion":
				var v ImmunizationRecommendationRecommendationDateCriterion
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DateCriterion = append(r.DateCriterion, v)
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "series":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Series = &v
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
			case "supportingImmunization":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SupportingImmunization = append(r.SupportingImmunization, v)
			case "supportingPatientInformation":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SupportingPatientInformation = append(r.SupportingPatientInformation, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImmunizationRecommendationRecommendation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Vaccine date recommendations.  For example, earliest date to administer, latest date to administer, etc.
type ImmunizationRecommendationRecommendationDateCriterion struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Date classification of recommendation.  For example, earliest date to give, latest date to give, etc.
	Code CodeableConcept
	// The date whose meaning is specified by dateCriterion.code.
	Value DateTime
}
type jsonImmunizationRecommendationRecommendationDateCriterion struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Code                  CodeableConcept   `json:"code,omitempty"`
	Value                 DateTime          `json:"value,omitempty"`
	ValuePrimitiveElement *primitiveElement `json:"_value,omitempty"`
}

func (r ImmunizationRecommendationRecommendationDateCriterion) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImmunizationRecommendationRecommendationDateCriterion) marshalJSON() jsonImmunizationRecommendationRecommendationDateCriterion {
	m := jsonImmunizationRecommendationRecommendationDateCriterion{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	if r.Value.Value != nil {
		m.Value = r.Value
	}
	if r.Value.Id != nil || r.Value.Extension != nil {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	return m
}
func (r *ImmunizationRecommendationRecommendationDateCriterion) UnmarshalJSON(b []byte) error {
	var m jsonImmunizationRecommendationRecommendationDateCriterion
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImmunizationRecommendationRecommendationDateCriterion) unmarshalJSON(m jsonImmunizationRecommendationRecommendationDateCriterion) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	return nil
}
func (r ImmunizationRecommendationRecommendationDateCriterion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Value, xml.StartElement{Name: xml.Name{Local: "value"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ImmunizationRecommendationRecommendationDateCriterion) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = v
			case "value":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ImmunizationRecommendationRecommendationDateCriterion) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
