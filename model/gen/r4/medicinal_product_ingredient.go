package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// An ingredient of a manufactured item or pharmaceutical product.
type MedicinalProductIngredient struct {
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
	// The identifier(s) of this Ingredient that are assigned by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate.
	Identifier *Identifier
	// Ingredient role e.g. Active ingredient, excipient.
	Role CodeableConcept
	// If the ingredient is a known or suspected allergen.
	AllergenicIndicator *Boolean
	// Manufacturer of this Ingredient.
	Manufacturer []Reference
	// A specified substance that comprises this ingredient.
	SpecifiedSubstance []MedicinalProductIngredientSpecifiedSubstance
	// The ingredient substance.
	Substance *MedicinalProductIngredientSubstance
}

func (r MedicinalProductIngredient) ResourceType() string {
	return "MedicinalProductIngredient"
}
func (r MedicinalProductIngredient) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonMedicinalProductIngredient struct {
	ResourceType                        string                                         `json:"resourceType"`
	Id                                  *Id                                            `json:"id,omitempty"`
	IdPrimitiveElement                  *primitiveElement                              `json:"_id,omitempty"`
	Meta                                *Meta                                          `json:"meta,omitempty"`
	ImplicitRules                       *Uri                                           `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement       *primitiveElement                              `json:"_implicitRules,omitempty"`
	Language                            *Code                                          `json:"language,omitempty"`
	LanguagePrimitiveElement            *primitiveElement                              `json:"_language,omitempty"`
	Text                                *Narrative                                     `json:"text,omitempty"`
	Contained                           []ContainedResource                            `json:"contained,omitempty"`
	Extension                           []Extension                                    `json:"extension,omitempty"`
	ModifierExtension                   []Extension                                    `json:"modifierExtension,omitempty"`
	Identifier                          *Identifier                                    `json:"identifier,omitempty"`
	Role                                CodeableConcept                                `json:"role,omitempty"`
	AllergenicIndicator                 *Boolean                                       `json:"allergenicIndicator,omitempty"`
	AllergenicIndicatorPrimitiveElement *primitiveElement                              `json:"_allergenicIndicator,omitempty"`
	Manufacturer                        []Reference                                    `json:"manufacturer,omitempty"`
	SpecifiedSubstance                  []MedicinalProductIngredientSpecifiedSubstance `json:"specifiedSubstance,omitempty"`
	Substance                           *MedicinalProductIngredientSubstance           `json:"substance,omitempty"`
}

func (r MedicinalProductIngredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductIngredient) marshalJSON() jsonMedicinalProductIngredient {
	m := jsonMedicinalProductIngredient{}
	m.ResourceType = "MedicinalProductIngredient"
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
	m.Role = r.Role
	if r.AllergenicIndicator != nil && r.AllergenicIndicator.Value != nil {
		m.AllergenicIndicator = r.AllergenicIndicator
	}
	if r.AllergenicIndicator != nil && (r.AllergenicIndicator.Id != nil || r.AllergenicIndicator.Extension != nil) {
		m.AllergenicIndicatorPrimitiveElement = &primitiveElement{Id: r.AllergenicIndicator.Id, Extension: r.AllergenicIndicator.Extension}
	}
	m.Manufacturer = r.Manufacturer
	m.SpecifiedSubstance = r.SpecifiedSubstance
	m.Substance = r.Substance
	return m
}
func (r *MedicinalProductIngredient) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductIngredient
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductIngredient) unmarshalJSON(m jsonMedicinalProductIngredient) error {
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
	r.Role = m.Role
	r.AllergenicIndicator = m.AllergenicIndicator
	if m.AllergenicIndicatorPrimitiveElement != nil {
		if r.AllergenicIndicator == nil {
			r.AllergenicIndicator = &Boolean{}
		}
		r.AllergenicIndicator.Id = m.AllergenicIndicatorPrimitiveElement.Id
		r.AllergenicIndicator.Extension = m.AllergenicIndicatorPrimitiveElement.Extension
	}
	r.Manufacturer = m.Manufacturer
	r.SpecifiedSubstance = m.SpecifiedSubstance
	r.Substance = m.Substance
	return nil
}
func (r MedicinalProductIngredient) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "MedicinalProductIngredient"
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
	err = e.EncodeElement(r.Role, xml.StartElement{Name: xml.Name{Local: "role"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AllergenicIndicator, xml.StartElement{Name: xml.Name{Local: "allergenicIndicator"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Manufacturer, xml.StartElement{Name: xml.Name{Local: "manufacturer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SpecifiedSubstance, xml.StartElement{Name: xml.Name{Local: "specifiedSubstance"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Substance, xml.StartElement{Name: xml.Name{Local: "substance"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductIngredient) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.Identifier = &v
			case "role":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Role = v
			case "allergenicIndicator":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AllergenicIndicator = &v
			case "manufacturer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Manufacturer = append(r.Manufacturer, v)
			case "specifiedSubstance":
				var v MedicinalProductIngredientSpecifiedSubstance
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SpecifiedSubstance = append(r.SpecifiedSubstance, v)
			case "substance":
				var v MedicinalProductIngredientSubstance
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Substance = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductIngredient) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A specified substance that comprises this ingredient.
type MedicinalProductIngredientSpecifiedSubstance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The specified substance.
	Code CodeableConcept
	// The group of specified substance, e.g. group 1 to 4.
	Group CodeableConcept
	// Confidentiality level of the specified substance as the ingredient.
	Confidentiality *CodeableConcept
	// Quantity of the substance or specified substance present in the manufactured item or pharmaceutical product.
	Strength []MedicinalProductIngredientSpecifiedSubstanceStrength
}
type jsonMedicinalProductIngredientSpecifiedSubstance struct {
	Id                *string                                                `json:"id,omitempty"`
	Extension         []Extension                                            `json:"extension,omitempty"`
	ModifierExtension []Extension                                            `json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                        `json:"code,omitempty"`
	Group             CodeableConcept                                        `json:"group,omitempty"`
	Confidentiality   *CodeableConcept                                       `json:"confidentiality,omitempty"`
	Strength          []MedicinalProductIngredientSpecifiedSubstanceStrength `json:"strength,omitempty"`
}

func (r MedicinalProductIngredientSpecifiedSubstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductIngredientSpecifiedSubstance) marshalJSON() jsonMedicinalProductIngredientSpecifiedSubstance {
	m := jsonMedicinalProductIngredientSpecifiedSubstance{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Group = r.Group
	m.Confidentiality = r.Confidentiality
	m.Strength = r.Strength
	return m
}
func (r *MedicinalProductIngredientSpecifiedSubstance) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductIngredientSpecifiedSubstance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductIngredientSpecifiedSubstance) unmarshalJSON(m jsonMedicinalProductIngredientSpecifiedSubstance) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Group = m.Group
	r.Confidentiality = m.Confidentiality
	r.Strength = m.Strength
	return nil
}
func (r MedicinalProductIngredientSpecifiedSubstance) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Group, xml.StartElement{Name: xml.Name{Local: "group"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Confidentiality, xml.StartElement{Name: xml.Name{Local: "confidentiality"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Strength, xml.StartElement{Name: xml.Name{Local: "strength"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductIngredientSpecifiedSubstance) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "group":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Group = v
			case "confidentiality":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Confidentiality = &v
			case "strength":
				var v MedicinalProductIngredientSpecifiedSubstanceStrength
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Strength = append(r.Strength, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductIngredientSpecifiedSubstance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Quantity of the substance or specified substance present in the manufactured item or pharmaceutical product.
type MedicinalProductIngredientSpecifiedSubstanceStrength struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The quantity of substance in the unit of presentation, or in the volume (or mass) of the single pharmaceutical product or manufactured item.
	Presentation Ratio
	// A lower limit for the quantity of substance in the unit of presentation. For use when there is a range of strengths, this is the lower limit, with the presentation attribute becoming the upper limit.
	PresentationLowLimit *Ratio
	// The strength per unitary volume (or mass).
	Concentration *Ratio
	// A lower limit for the strength per unitary volume (or mass), for when there is a range. The concentration attribute then becomes the upper limit.
	ConcentrationLowLimit *Ratio
	// For when strength is measured at a particular point or distance.
	MeasurementPoint *String
	// The country or countries for which the strength range applies.
	Country []CodeableConcept
	// Strength expressed in terms of a reference substance.
	ReferenceStrength []MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength
}
type jsonMedicinalProductIngredientSpecifiedSubstanceStrength struct {
	Id                               *string                                                                 `json:"id,omitempty"`
	Extension                        []Extension                                                             `json:"extension,omitempty"`
	ModifierExtension                []Extension                                                             `json:"modifierExtension,omitempty"`
	Presentation                     Ratio                                                                   `json:"presentation,omitempty"`
	PresentationLowLimit             *Ratio                                                                  `json:"presentationLowLimit,omitempty"`
	Concentration                    *Ratio                                                                  `json:"concentration,omitempty"`
	ConcentrationLowLimit            *Ratio                                                                  `json:"concentrationLowLimit,omitempty"`
	MeasurementPoint                 *String                                                                 `json:"measurementPoint,omitempty"`
	MeasurementPointPrimitiveElement *primitiveElement                                                       `json:"_measurementPoint,omitempty"`
	Country                          []CodeableConcept                                                       `json:"country,omitempty"`
	ReferenceStrength                []MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength `json:"referenceStrength,omitempty"`
}

func (r MedicinalProductIngredientSpecifiedSubstanceStrength) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductIngredientSpecifiedSubstanceStrength) marshalJSON() jsonMedicinalProductIngredientSpecifiedSubstanceStrength {
	m := jsonMedicinalProductIngredientSpecifiedSubstanceStrength{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Presentation = r.Presentation
	m.PresentationLowLimit = r.PresentationLowLimit
	m.Concentration = r.Concentration
	m.ConcentrationLowLimit = r.ConcentrationLowLimit
	if r.MeasurementPoint != nil && r.MeasurementPoint.Value != nil {
		m.MeasurementPoint = r.MeasurementPoint
	}
	if r.MeasurementPoint != nil && (r.MeasurementPoint.Id != nil || r.MeasurementPoint.Extension != nil) {
		m.MeasurementPointPrimitiveElement = &primitiveElement{Id: r.MeasurementPoint.Id, Extension: r.MeasurementPoint.Extension}
	}
	m.Country = r.Country
	m.ReferenceStrength = r.ReferenceStrength
	return m
}
func (r *MedicinalProductIngredientSpecifiedSubstanceStrength) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductIngredientSpecifiedSubstanceStrength
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductIngredientSpecifiedSubstanceStrength) unmarshalJSON(m jsonMedicinalProductIngredientSpecifiedSubstanceStrength) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Presentation = m.Presentation
	r.PresentationLowLimit = m.PresentationLowLimit
	r.Concentration = m.Concentration
	r.ConcentrationLowLimit = m.ConcentrationLowLimit
	r.MeasurementPoint = m.MeasurementPoint
	if m.MeasurementPointPrimitiveElement != nil {
		if r.MeasurementPoint == nil {
			r.MeasurementPoint = &String{}
		}
		r.MeasurementPoint.Id = m.MeasurementPointPrimitiveElement.Id
		r.MeasurementPoint.Extension = m.MeasurementPointPrimitiveElement.Extension
	}
	r.Country = m.Country
	r.ReferenceStrength = m.ReferenceStrength
	return nil
}
func (r MedicinalProductIngredientSpecifiedSubstanceStrength) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Presentation, xml.StartElement{Name: xml.Name{Local: "presentation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PresentationLowLimit, xml.StartElement{Name: xml.Name{Local: "presentationLowLimit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Concentration, xml.StartElement{Name: xml.Name{Local: "concentration"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ConcentrationLowLimit, xml.StartElement{Name: xml.Name{Local: "concentrationLowLimit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MeasurementPoint, xml.StartElement{Name: xml.Name{Local: "measurementPoint"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Country, xml.StartElement{Name: xml.Name{Local: "country"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReferenceStrength, xml.StartElement{Name: xml.Name{Local: "referenceStrength"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductIngredientSpecifiedSubstanceStrength) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "presentation":
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Presentation = v
			case "presentationLowLimit":
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PresentationLowLimit = &v
			case "concentration":
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Concentration = &v
			case "concentrationLowLimit":
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ConcentrationLowLimit = &v
			case "measurementPoint":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MeasurementPoint = &v
			case "country":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Country = append(r.Country, v)
			case "referenceStrength":
				var v MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReferenceStrength = append(r.ReferenceStrength, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductIngredientSpecifiedSubstanceStrength) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Strength expressed in terms of a reference substance.
type MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Relevant reference substance.
	Substance *CodeableConcept
	// Strength expressed in terms of a reference substance.
	Strength Ratio
	// Strength expressed in terms of a reference substance.
	StrengthLowLimit *Ratio
	// For when strength is measured at a particular point or distance.
	MeasurementPoint *String
	// The country or countries for which the strength range applies.
	Country []CodeableConcept
}
type jsonMedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength struct {
	Id                               *string           `json:"id,omitempty"`
	Extension                        []Extension       `json:"extension,omitempty"`
	ModifierExtension                []Extension       `json:"modifierExtension,omitempty"`
	Substance                        *CodeableConcept  `json:"substance,omitempty"`
	Strength                         Ratio             `json:"strength,omitempty"`
	StrengthLowLimit                 *Ratio            `json:"strengthLowLimit,omitempty"`
	MeasurementPoint                 *String           `json:"measurementPoint,omitempty"`
	MeasurementPointPrimitiveElement *primitiveElement `json:"_measurementPoint,omitempty"`
	Country                          []CodeableConcept `json:"country,omitempty"`
}

func (r MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength) marshalJSON() jsonMedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength {
	m := jsonMedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Substance = r.Substance
	m.Strength = r.Strength
	m.StrengthLowLimit = r.StrengthLowLimit
	if r.MeasurementPoint != nil && r.MeasurementPoint.Value != nil {
		m.MeasurementPoint = r.MeasurementPoint
	}
	if r.MeasurementPoint != nil && (r.MeasurementPoint.Id != nil || r.MeasurementPoint.Extension != nil) {
		m.MeasurementPointPrimitiveElement = &primitiveElement{Id: r.MeasurementPoint.Id, Extension: r.MeasurementPoint.Extension}
	}
	m.Country = r.Country
	return m
}
func (r *MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength) unmarshalJSON(m jsonMedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Substance = m.Substance
	r.Strength = m.Strength
	r.StrengthLowLimit = m.StrengthLowLimit
	r.MeasurementPoint = m.MeasurementPoint
	if m.MeasurementPointPrimitiveElement != nil {
		if r.MeasurementPoint == nil {
			r.MeasurementPoint = &String{}
		}
		r.MeasurementPoint.Id = m.MeasurementPointPrimitiveElement.Id
		r.MeasurementPoint.Extension = m.MeasurementPointPrimitiveElement.Extension
	}
	r.Country = m.Country
	return nil
}
func (r MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Substance, xml.StartElement{Name: xml.Name{Local: "substance"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Strength, xml.StartElement{Name: xml.Name{Local: "strength"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.StrengthLowLimit, xml.StartElement{Name: xml.Name{Local: "strengthLowLimit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MeasurementPoint, xml.StartElement{Name: xml.Name{Local: "measurementPoint"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Country, xml.StartElement{Name: xml.Name{Local: "country"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "substance":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Substance = &v
			case "strength":
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Strength = v
			case "strengthLowLimit":
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StrengthLowLimit = &v
			case "measurementPoint":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MeasurementPoint = &v
			case "country":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Country = append(r.Country, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The ingredient substance.
type MedicinalProductIngredientSubstance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The ingredient substance.
	Code CodeableConcept
	// Quantity of the substance or specified substance present in the manufactured item or pharmaceutical product.
	Strength []MedicinalProductIngredientSpecifiedSubstanceStrength
}
type jsonMedicinalProductIngredientSubstance struct {
	Id                *string                                                `json:"id,omitempty"`
	Extension         []Extension                                            `json:"extension,omitempty"`
	ModifierExtension []Extension                                            `json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                        `json:"code,omitempty"`
	Strength          []MedicinalProductIngredientSpecifiedSubstanceStrength `json:"strength,omitempty"`
}

func (r MedicinalProductIngredientSubstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductIngredientSubstance) marshalJSON() jsonMedicinalProductIngredientSubstance {
	m := jsonMedicinalProductIngredientSubstance{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Strength = r.Strength
	return m
}
func (r *MedicinalProductIngredientSubstance) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductIngredientSubstance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductIngredientSubstance) unmarshalJSON(m jsonMedicinalProductIngredientSubstance) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Strength = m.Strength
	return nil
}
func (r MedicinalProductIngredientSubstance) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Strength, xml.StartElement{Name: xml.Name{Local: "strength"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductIngredientSubstance) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "strength":
				var v MedicinalProductIngredientSpecifiedSubstanceStrength
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Strength = append(r.Strength, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductIngredientSubstance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
