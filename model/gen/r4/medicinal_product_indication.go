package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// Indication for the Medicinal Product.
type MedicinalProductIndication struct {
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
	// The medication for which this is an indication.
	Subject []Reference
	// The disease, symptom or procedure that is the indication for treatment.
	DiseaseSymptomProcedure *CodeableConcept
	// The status of the disease or symptom for which the indication applies.
	DiseaseStatus *CodeableConcept
	// Comorbidity (concurrent condition) or co-infection as part of the indication.
	Comorbidity []CodeableConcept
	// The intended effect, aim or strategy to be achieved by the indication.
	IntendedEffect *CodeableConcept
	// Timing or duration information as part of the indication.
	Duration *Quantity
	// Information about the use of the medicinal product in relation to other therapies described as part of the indication.
	OtherTherapy []MedicinalProductIndicationOtherTherapy
	// Describe the undesirable effects of the medicinal product.
	UndesirableEffect []Reference
	// The population group to which this applies.
	Population []Population
}

func (r MedicinalProductIndication) ResourceType() string {
	return "MedicinalProductIndication"
}
func (r MedicinalProductIndication) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonMedicinalProductIndication struct {
	ResourceType                  string                                   `json:"resourceType"`
	Id                            *Id                                      `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement                        `json:"_id,omitempty"`
	Meta                          *Meta                                    `json:"meta,omitempty"`
	ImplicitRules                 *Uri                                     `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement                        `json:"_implicitRules,omitempty"`
	Language                      *Code                                    `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement                        `json:"_language,omitempty"`
	Text                          *Narrative                               `json:"text,omitempty"`
	Contained                     []ContainedResource                      `json:"contained,omitempty"`
	Extension                     []Extension                              `json:"extension,omitempty"`
	ModifierExtension             []Extension                              `json:"modifierExtension,omitempty"`
	Subject                       []Reference                              `json:"subject,omitempty"`
	DiseaseSymptomProcedure       *CodeableConcept                         `json:"diseaseSymptomProcedure,omitempty"`
	DiseaseStatus                 *CodeableConcept                         `json:"diseaseStatus,omitempty"`
	Comorbidity                   []CodeableConcept                        `json:"comorbidity,omitempty"`
	IntendedEffect                *CodeableConcept                         `json:"intendedEffect,omitempty"`
	Duration                      *Quantity                                `json:"duration,omitempty"`
	OtherTherapy                  []MedicinalProductIndicationOtherTherapy `json:"otherTherapy,omitempty"`
	UndesirableEffect             []Reference                              `json:"undesirableEffect,omitempty"`
	Population                    []Population                             `json:"population,omitempty"`
}

func (r MedicinalProductIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductIndication) marshalJSON() jsonMedicinalProductIndication {
	m := jsonMedicinalProductIndication{}
	m.ResourceType = "MedicinalProductIndication"
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
	m.Subject = r.Subject
	m.DiseaseSymptomProcedure = r.DiseaseSymptomProcedure
	m.DiseaseStatus = r.DiseaseStatus
	m.Comorbidity = r.Comorbidity
	m.IntendedEffect = r.IntendedEffect
	m.Duration = r.Duration
	m.OtherTherapy = r.OtherTherapy
	m.UndesirableEffect = r.UndesirableEffect
	m.Population = r.Population
	return m
}
func (r *MedicinalProductIndication) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductIndication
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductIndication) unmarshalJSON(m jsonMedicinalProductIndication) error {
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
	r.Subject = m.Subject
	r.DiseaseSymptomProcedure = m.DiseaseSymptomProcedure
	r.DiseaseStatus = m.DiseaseStatus
	r.Comorbidity = m.Comorbidity
	r.IntendedEffect = m.IntendedEffect
	r.Duration = m.Duration
	r.OtherTherapy = m.OtherTherapy
	r.UndesirableEffect = m.UndesirableEffect
	r.Population = m.Population
	return nil
}
func (r MedicinalProductIndication) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "MedicinalProductIndication"
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
	err = e.EncodeElement(r.Subject, xml.StartElement{Name: xml.Name{Local: "subject"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DiseaseSymptomProcedure, xml.StartElement{Name: xml.Name{Local: "diseaseSymptomProcedure"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DiseaseStatus, xml.StartElement{Name: xml.Name{Local: "diseaseStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Comorbidity, xml.StartElement{Name: xml.Name{Local: "comorbidity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IntendedEffect, xml.StartElement{Name: xml.Name{Local: "intendedEffect"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Duration, xml.StartElement{Name: xml.Name{Local: "duration"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OtherTherapy, xml.StartElement{Name: xml.Name{Local: "otherTherapy"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UndesirableEffect, xml.StartElement{Name: xml.Name{Local: "undesirableEffect"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Population, xml.StartElement{Name: xml.Name{Local: "population"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductIndication) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "subject":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = append(r.Subject, v)
			case "diseaseSymptomProcedure":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DiseaseSymptomProcedure = &v
			case "diseaseStatus":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DiseaseStatus = &v
			case "comorbidity":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Comorbidity = append(r.Comorbidity, v)
			case "intendedEffect":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IntendedEffect = &v
			case "duration":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Duration = &v
			case "otherTherapy":
				var v MedicinalProductIndicationOtherTherapy
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OtherTherapy = append(r.OtherTherapy, v)
			case "undesirableEffect":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UndesirableEffect = append(r.UndesirableEffect, v)
			case "population":
				var v Population
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Population = append(r.Population, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductIndication) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Information about the use of the medicinal product in relation to other therapies described as part of the indication.
type MedicinalProductIndicationOtherTherapy struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of relationship between the medicinal product indication or contraindication and another therapy.
	TherapyRelationshipType CodeableConcept
	// Reference to a specific medication (active substance, medicinal product or class of products) as part of an indication or contraindication.
	Medication isMedicinalProductIndicationOtherTherapyMedication
}
type isMedicinalProductIndicationOtherTherapyMedication interface {
	isMedicinalProductIndicationOtherTherapyMedication()
}

func (r CodeableConcept) isMedicinalProductIndicationOtherTherapyMedication() {}
func (r Reference) isMedicinalProductIndicationOtherTherapyMedication()       {}

type jsonMedicinalProductIndicationOtherTherapy struct {
	Id                        *string          `json:"id,omitempty"`
	Extension                 []Extension      `json:"extension,omitempty"`
	ModifierExtension         []Extension      `json:"modifierExtension,omitempty"`
	TherapyRelationshipType   CodeableConcept  `json:"therapyRelationshipType,omitempty"`
	MedicationCodeableConcept *CodeableConcept `json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference       `json:"medicationReference,omitempty"`
}

func (r MedicinalProductIndicationOtherTherapy) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductIndicationOtherTherapy) marshalJSON() jsonMedicinalProductIndicationOtherTherapy {
	m := jsonMedicinalProductIndicationOtherTherapy{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.TherapyRelationshipType = r.TherapyRelationshipType
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
	return m
}
func (r *MedicinalProductIndicationOtherTherapy) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductIndicationOtherTherapy
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductIndicationOtherTherapy) unmarshalJSON(m jsonMedicinalProductIndicationOtherTherapy) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.TherapyRelationshipType = m.TherapyRelationshipType
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
	return nil
}
func (r MedicinalProductIndicationOtherTherapy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.TherapyRelationshipType, xml.StartElement{Name: xml.Name{Local: "therapyRelationshipType"}})
	if err != nil {
		return err
	}
	switch v := r.Medication.(type) {
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "medicationCodeableConcept"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "medicationReference"}})
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
func (r *MedicinalProductIndicationOtherTherapy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "therapyRelationshipType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TherapyRelationshipType = v
			case "medicationCodeableConcept":
				if r.Medication != nil {
					return fmt.Errorf("multiple values for field \"Medication\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Medication = v
			case "medicationReference":
				if r.Medication != nil {
					return fmt.Errorf("multiple values for field \"Medication\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Medication = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductIndicationOtherTherapy) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
