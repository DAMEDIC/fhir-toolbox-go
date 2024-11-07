package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// A pharmaceutical product described in terms of its composition and dose form.
type MedicinalProductPharmaceutical struct {
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
	// An identifier for the pharmaceutical medicinal product.
	Identifier []Identifier
	// The administrable dose form, after necessary reconstitution.
	AdministrableDoseForm CodeableConcept
	// Todo.
	UnitOfPresentation *CodeableConcept
	// Ingredient.
	Ingredient []Reference
	// Accompanying device.
	Device []Reference
	// Characteristics e.g. a products onset of action.
	Characteristics []MedicinalProductPharmaceuticalCharacteristics
	// The path by which the pharmaceutical product is taken into or makes contact with the body.
	RouteOfAdministration []MedicinalProductPharmaceuticalRouteOfAdministration
}

func (r MedicinalProductPharmaceutical) ResourceType() string {
	return "MedicinalProductPharmaceutical"
}
func (r MedicinalProductPharmaceutical) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonMedicinalProductPharmaceutical struct {
	ResourceType                  string                                                `json:"resourceType"`
	Id                            *Id                                                   `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement                                     `json:"_id,omitempty"`
	Meta                          *Meta                                                 `json:"meta,omitempty"`
	ImplicitRules                 *Uri                                                  `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement                                     `json:"_implicitRules,omitempty"`
	Language                      *Code                                                 `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement                                     `json:"_language,omitempty"`
	Text                          *Narrative                                            `json:"text,omitempty"`
	Contained                     []ContainedResource                                   `json:"contained,omitempty"`
	Extension                     []Extension                                           `json:"extension,omitempty"`
	ModifierExtension             []Extension                                           `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier                                          `json:"identifier,omitempty"`
	AdministrableDoseForm         CodeableConcept                                       `json:"administrableDoseForm,omitempty"`
	UnitOfPresentation            *CodeableConcept                                      `json:"unitOfPresentation,omitempty"`
	Ingredient                    []Reference                                           `json:"ingredient,omitempty"`
	Device                        []Reference                                           `json:"device,omitempty"`
	Characteristics               []MedicinalProductPharmaceuticalCharacteristics       `json:"characteristics,omitempty"`
	RouteOfAdministration         []MedicinalProductPharmaceuticalRouteOfAdministration `json:"routeOfAdministration,omitempty"`
}

func (r MedicinalProductPharmaceutical) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductPharmaceutical) marshalJSON() jsonMedicinalProductPharmaceutical {
	m := jsonMedicinalProductPharmaceutical{}
	m.ResourceType = "MedicinalProductPharmaceutical"
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
	m.AdministrableDoseForm = r.AdministrableDoseForm
	m.UnitOfPresentation = r.UnitOfPresentation
	m.Ingredient = r.Ingredient
	m.Device = r.Device
	m.Characteristics = r.Characteristics
	m.RouteOfAdministration = r.RouteOfAdministration
	return m
}
func (r *MedicinalProductPharmaceutical) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductPharmaceutical
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductPharmaceutical) unmarshalJSON(m jsonMedicinalProductPharmaceutical) error {
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
	r.AdministrableDoseForm = m.AdministrableDoseForm
	r.UnitOfPresentation = m.UnitOfPresentation
	r.Ingredient = m.Ingredient
	r.Device = m.Device
	r.Characteristics = m.Characteristics
	r.RouteOfAdministration = m.RouteOfAdministration
	return nil
}
func (r MedicinalProductPharmaceutical) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "MedicinalProductPharmaceutical"
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
	err = e.EncodeElement(r.AdministrableDoseForm, xml.StartElement{Name: xml.Name{Local: "administrableDoseForm"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UnitOfPresentation, xml.StartElement{Name: xml.Name{Local: "unitOfPresentation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Ingredient, xml.StartElement{Name: xml.Name{Local: "ingredient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Device, xml.StartElement{Name: xml.Name{Local: "device"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Characteristics, xml.StartElement{Name: xml.Name{Local: "characteristics"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RouteOfAdministration, xml.StartElement{Name: xml.Name{Local: "routeOfAdministration"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductPharmaceutical) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "administrableDoseForm":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AdministrableDoseForm = v
			case "unitOfPresentation":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UnitOfPresentation = &v
			case "ingredient":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Ingredient = append(r.Ingredient, v)
			case "device":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Device = append(r.Device, v)
			case "characteristics":
				var v MedicinalProductPharmaceuticalCharacteristics
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Characteristics = append(r.Characteristics, v)
			case "routeOfAdministration":
				var v MedicinalProductPharmaceuticalRouteOfAdministration
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RouteOfAdministration = append(r.RouteOfAdministration, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductPharmaceutical) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Characteristics e.g. a products onset of action.
type MedicinalProductPharmaceuticalCharacteristics struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A coded characteristic.
	Code CodeableConcept
	// The status of characteristic e.g. assigned or pending.
	Status *CodeableConcept
}
type jsonMedicinalProductPharmaceuticalCharacteristics struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              CodeableConcept  `json:"code,omitempty"`
	Status            *CodeableConcept `json:"status,omitempty"`
}

func (r MedicinalProductPharmaceuticalCharacteristics) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductPharmaceuticalCharacteristics) marshalJSON() jsonMedicinalProductPharmaceuticalCharacteristics {
	m := jsonMedicinalProductPharmaceuticalCharacteristics{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Status = r.Status
	return m
}
func (r *MedicinalProductPharmaceuticalCharacteristics) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductPharmaceuticalCharacteristics
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductPharmaceuticalCharacteristics) unmarshalJSON(m jsonMedicinalProductPharmaceuticalCharacteristics) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Status = m.Status
	return nil
}
func (r MedicinalProductPharmaceuticalCharacteristics) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductPharmaceuticalCharacteristics) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "status":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductPharmaceuticalCharacteristics) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The path by which the pharmaceutical product is taken into or makes contact with the body.
type MedicinalProductPharmaceuticalRouteOfAdministration struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Coded expression for the route.
	Code CodeableConcept
	// The first dose (dose quantity) administered in humans can be specified, for a product under investigation, using a numerical value and its unit of measurement.
	FirstDose *Quantity
	// The maximum single dose that can be administered as per the protocol of a clinical trial can be specified using a numerical value and its unit of measurement.
	MaxSingleDose *Quantity
	// The maximum dose per day (maximum dose quantity to be administered in any one 24-h period) that can be administered as per the protocol referenced in the clinical trial authorisation.
	MaxDosePerDay *Quantity
	// The maximum dose per treatment period that can be administered as per the protocol referenced in the clinical trial authorisation.
	MaxDosePerTreatmentPeriod *Ratio
	// The maximum treatment period during which an Investigational Medicinal Product can be administered as per the protocol referenced in the clinical trial authorisation.
	MaxTreatmentPeriod *Duration
	// A species for which this route applies.
	TargetSpecies []MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies
}
type jsonMedicinalProductPharmaceuticalRouteOfAdministration struct {
	Id                        *string                                                            `json:"id,omitempty"`
	Extension                 []Extension                                                        `json:"extension,omitempty"`
	ModifierExtension         []Extension                                                        `json:"modifierExtension,omitempty"`
	Code                      CodeableConcept                                                    `json:"code,omitempty"`
	FirstDose                 *Quantity                                                          `json:"firstDose,omitempty"`
	MaxSingleDose             *Quantity                                                          `json:"maxSingleDose,omitempty"`
	MaxDosePerDay             *Quantity                                                          `json:"maxDosePerDay,omitempty"`
	MaxDosePerTreatmentPeriod *Ratio                                                             `json:"maxDosePerTreatmentPeriod,omitempty"`
	MaxTreatmentPeriod        *Duration                                                          `json:"maxTreatmentPeriod,omitempty"`
	TargetSpecies             []MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies `json:"targetSpecies,omitempty"`
}

func (r MedicinalProductPharmaceuticalRouteOfAdministration) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductPharmaceuticalRouteOfAdministration) marshalJSON() jsonMedicinalProductPharmaceuticalRouteOfAdministration {
	m := jsonMedicinalProductPharmaceuticalRouteOfAdministration{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.FirstDose = r.FirstDose
	m.MaxSingleDose = r.MaxSingleDose
	m.MaxDosePerDay = r.MaxDosePerDay
	m.MaxDosePerTreatmentPeriod = r.MaxDosePerTreatmentPeriod
	m.MaxTreatmentPeriod = r.MaxTreatmentPeriod
	m.TargetSpecies = r.TargetSpecies
	return m
}
func (r *MedicinalProductPharmaceuticalRouteOfAdministration) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductPharmaceuticalRouteOfAdministration
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductPharmaceuticalRouteOfAdministration) unmarshalJSON(m jsonMedicinalProductPharmaceuticalRouteOfAdministration) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.FirstDose = m.FirstDose
	r.MaxSingleDose = m.MaxSingleDose
	r.MaxDosePerDay = m.MaxDosePerDay
	r.MaxDosePerTreatmentPeriod = m.MaxDosePerTreatmentPeriod
	r.MaxTreatmentPeriod = m.MaxTreatmentPeriod
	r.TargetSpecies = m.TargetSpecies
	return nil
}
func (r MedicinalProductPharmaceuticalRouteOfAdministration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.FirstDose, xml.StartElement{Name: xml.Name{Local: "firstDose"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MaxSingleDose, xml.StartElement{Name: xml.Name{Local: "maxSingleDose"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MaxDosePerDay, xml.StartElement{Name: xml.Name{Local: "maxDosePerDay"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MaxDosePerTreatmentPeriod, xml.StartElement{Name: xml.Name{Local: "maxDosePerTreatmentPeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MaxTreatmentPeriod, xml.StartElement{Name: xml.Name{Local: "maxTreatmentPeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TargetSpecies, xml.StartElement{Name: xml.Name{Local: "targetSpecies"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductPharmaceuticalRouteOfAdministration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "firstDose":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FirstDose = &v
			case "maxSingleDose":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MaxSingleDose = &v
			case "maxDosePerDay":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MaxDosePerDay = &v
			case "maxDosePerTreatmentPeriod":
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MaxDosePerTreatmentPeriod = &v
			case "maxTreatmentPeriod":
				var v Duration
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MaxTreatmentPeriod = &v
			case "targetSpecies":
				var v MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TargetSpecies = append(r.TargetSpecies, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductPharmaceuticalRouteOfAdministration) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A species for which this route applies.
type MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Coded expression for the species.
	Code CodeableConcept
	// A species specific time during which consumption of animal product is not appropriate.
	WithdrawalPeriod []MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod
}
type jsonMedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies struct {
	Id                *string                                                                            `json:"id,omitempty"`
	Extension         []Extension                                                                        `json:"extension,omitempty"`
	ModifierExtension []Extension                                                                        `json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                                                    `json:"code,omitempty"`
	WithdrawalPeriod  []MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod `json:"withdrawalPeriod,omitempty"`
}

func (r MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies) marshalJSON() jsonMedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies {
	m := jsonMedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.WithdrawalPeriod = r.WithdrawalPeriod
	return m
}
func (r *MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies) unmarshalJSON(m jsonMedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.WithdrawalPeriod = m.WithdrawalPeriod
	return nil
}
func (r MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.WithdrawalPeriod, xml.StartElement{Name: xml.Name{Local: "withdrawalPeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "withdrawalPeriod":
				var v MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.WithdrawalPeriod = append(r.WithdrawalPeriod, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpecies) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A species specific time during which consumption of animal product is not appropriate.
type MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Coded expression for the type of tissue for which the withdrawal period applues, e.g. meat, milk.
	Tissue CodeableConcept
	// A value for the time.
	Value Quantity
	// Extra information about the withdrawal period.
	SupportingInformation *String
}
type jsonMedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
	Id                                    *string           `json:"id,omitempty"`
	Extension                             []Extension       `json:"extension,omitempty"`
	ModifierExtension                     []Extension       `json:"modifierExtension,omitempty"`
	Tissue                                CodeableConcept   `json:"tissue,omitempty"`
	Value                                 Quantity          `json:"value,omitempty"`
	SupportingInformation                 *String           `json:"supportingInformation,omitempty"`
	SupportingInformationPrimitiveElement *primitiveElement `json:"_supportingInformation,omitempty"`
}

func (r MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod) marshalJSON() jsonMedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod {
	m := jsonMedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Tissue = r.Tissue
	m.Value = r.Value
	if r.SupportingInformation != nil && r.SupportingInformation.Value != nil {
		m.SupportingInformation = r.SupportingInformation
	}
	if r.SupportingInformation != nil && (r.SupportingInformation.Id != nil || r.SupportingInformation.Extension != nil) {
		m.SupportingInformationPrimitiveElement = &primitiveElement{Id: r.SupportingInformation.Id, Extension: r.SupportingInformation.Extension}
	}
	return m
}
func (r *MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod) unmarshalJSON(m jsonMedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Tissue = m.Tissue
	r.Value = m.Value
	r.SupportingInformation = m.SupportingInformation
	if m.SupportingInformationPrimitiveElement != nil {
		if r.SupportingInformation == nil {
			r.SupportingInformation = &String{}
		}
		r.SupportingInformation.Id = m.SupportingInformationPrimitiveElement.Id
		r.SupportingInformation.Extension = m.SupportingInformationPrimitiveElement.Extension
	}
	return nil
}
func (r MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Tissue, xml.StartElement{Name: xml.Name{Local: "tissue"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Value, xml.StartElement{Name: xml.Name{Local: "value"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SupportingInformation, xml.StartElement{Name: xml.Name{Local: "supportingInformation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "tissue":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Tissue = v
			case "value":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "supportingInformation":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SupportingInformation = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MedicinalProductPharmaceuticalRouteOfAdministrationTargetSpeciesWithdrawalPeriod) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
