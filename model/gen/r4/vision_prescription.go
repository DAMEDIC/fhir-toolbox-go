package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// An authorization for the provision of glasses and/or contact lenses to a patient.
type VisionPrescription struct {
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
	// A unique identifier assigned to this vision prescription.
	Identifier []Identifier
	// The status of the resource instance.
	Status Code
	// The date this resource was created.
	Created DateTime
	// A resource reference to the person to whom the vision prescription applies.
	Patient Reference
	// A reference to a resource that identifies the particular occurrence of contact between patient and health care provider during which the prescription was issued.
	Encounter *Reference
	// The date (and perhaps time) when the prescription was written.
	DateWritten DateTime
	// The healthcare professional responsible for authorizing the prescription.
	Prescriber Reference
	// Contain the details of  the individual lens specifications and serves as the authorization for the fullfillment by certified professionals.
	LensSpecification []VisionPrescriptionLensSpecification
}

func (r VisionPrescription) ResourceType() string {
	return "VisionPrescription"
}
func (r VisionPrescription) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonVisionPrescription struct {
	ResourceType                  string                                `json:"resourceType"`
	Id                            *Id                                   `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement                     `json:"_id,omitempty"`
	Meta                          *Meta                                 `json:"meta,omitempty"`
	ImplicitRules                 *Uri                                  `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement                     `json:"_implicitRules,omitempty"`
	Language                      *Code                                 `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement                     `json:"_language,omitempty"`
	Text                          *Narrative                            `json:"text,omitempty"`
	Contained                     []ContainedResource                   `json:"contained,omitempty"`
	Extension                     []Extension                           `json:"extension,omitempty"`
	ModifierExtension             []Extension                           `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier                          `json:"identifier,omitempty"`
	Status                        Code                                  `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement                     `json:"_status,omitempty"`
	Created                       DateTime                              `json:"created,omitempty"`
	CreatedPrimitiveElement       *primitiveElement                     `json:"_created,omitempty"`
	Patient                       Reference                             `json:"patient,omitempty"`
	Encounter                     *Reference                            `json:"encounter,omitempty"`
	DateWritten                   DateTime                              `json:"dateWritten,omitempty"`
	DateWrittenPrimitiveElement   *primitiveElement                     `json:"_dateWritten,omitempty"`
	Prescriber                    Reference                             `json:"prescriber,omitempty"`
	LensSpecification             []VisionPrescriptionLensSpecification `json:"lensSpecification,omitempty"`
}

func (r VisionPrescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r VisionPrescription) marshalJSON() jsonVisionPrescription {
	m := jsonVisionPrescription{}
	m.ResourceType = "VisionPrescription"
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
	if r.Created.Value != nil {
		m.Created = r.Created
	}
	if r.Created.Id != nil || r.Created.Extension != nil {
		m.CreatedPrimitiveElement = &primitiveElement{Id: r.Created.Id, Extension: r.Created.Extension}
	}
	m.Patient = r.Patient
	m.Encounter = r.Encounter
	if r.DateWritten.Value != nil {
		m.DateWritten = r.DateWritten
	}
	if r.DateWritten.Id != nil || r.DateWritten.Extension != nil {
		m.DateWrittenPrimitiveElement = &primitiveElement{Id: r.DateWritten.Id, Extension: r.DateWritten.Extension}
	}
	m.Prescriber = r.Prescriber
	m.LensSpecification = r.LensSpecification
	return m
}
func (r *VisionPrescription) UnmarshalJSON(b []byte) error {
	var m jsonVisionPrescription
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *VisionPrescription) unmarshalJSON(m jsonVisionPrescription) error {
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
	r.Created = m.Created
	if m.CreatedPrimitiveElement != nil {
		r.Created.Id = m.CreatedPrimitiveElement.Id
		r.Created.Extension = m.CreatedPrimitiveElement.Extension
	}
	r.Patient = m.Patient
	r.Encounter = m.Encounter
	r.DateWritten = m.DateWritten
	if m.DateWrittenPrimitiveElement != nil {
		r.DateWritten.Id = m.DateWrittenPrimitiveElement.Id
		r.DateWritten.Extension = m.DateWrittenPrimitiveElement.Extension
	}
	r.Prescriber = m.Prescriber
	r.LensSpecification = m.LensSpecification
	return nil
}
func (r VisionPrescription) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "VisionPrescription"
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
	err = e.EncodeElement(r.Created, xml.StartElement{Name: xml.Name{Local: "created"}})
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
	err = e.EncodeElement(r.DateWritten, xml.StartElement{Name: xml.Name{Local: "dateWritten"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Prescriber, xml.StartElement{Name: xml.Name{Local: "prescriber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LensSpecification, xml.StartElement{Name: xml.Name{Local: "lensSpecification"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *VisionPrescription) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "created":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Created = v
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
			case "dateWritten":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DateWritten = v
			case "prescriber":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Prescriber = v
			case "lensSpecification":
				var v VisionPrescriptionLensSpecification
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LensSpecification = append(r.LensSpecification, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r VisionPrescription) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Contain the details of  the individual lens specifications and serves as the authorization for the fullfillment by certified professionals.
type VisionPrescriptionLensSpecification struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identifies the type of vision correction product which is required for the patient.
	Product CodeableConcept
	// The eye for which the lens specification applies.
	Eye Code
	// Lens power measured in dioptres (0.25 units).
	Sphere *Decimal
	// Power adjustment for astigmatism measured in dioptres (0.25 units).
	Cylinder *Decimal
	// Adjustment for astigmatism measured in integer degrees.
	Axis *Integer
	// Allows for adjustment on two axis.
	Prism []VisionPrescriptionLensSpecificationPrism
	// Power adjustment for multifocal lenses measured in dioptres (0.25 units).
	Add *Decimal
	// Contact lens power measured in dioptres (0.25 units).
	Power *Decimal
	// Back curvature measured in millimetres.
	BackCurve *Decimal
	// Contact lens diameter measured in millimetres.
	Diameter *Decimal
	// The recommended maximum wear period for the lens.
	Duration *Quantity
	// Special color or pattern.
	Color *String
	// Brand recommendations or restrictions.
	Brand *String
	// Notes for special requirements such as coatings and lens materials.
	Note []Annotation
}
type jsonVisionPrescriptionLensSpecification struct {
	Id                        *string                                    `json:"id,omitempty"`
	Extension                 []Extension                                `json:"extension,omitempty"`
	ModifierExtension         []Extension                                `json:"modifierExtension,omitempty"`
	Product                   CodeableConcept                            `json:"product,omitempty"`
	Eye                       Code                                       `json:"eye,omitempty"`
	EyePrimitiveElement       *primitiveElement                          `json:"_eye,omitempty"`
	Sphere                    *Decimal                                   `json:"sphere,omitempty"`
	SpherePrimitiveElement    *primitiveElement                          `json:"_sphere,omitempty"`
	Cylinder                  *Decimal                                   `json:"cylinder,omitempty"`
	CylinderPrimitiveElement  *primitiveElement                          `json:"_cylinder,omitempty"`
	Axis                      *Integer                                   `json:"axis,omitempty"`
	AxisPrimitiveElement      *primitiveElement                          `json:"_axis,omitempty"`
	Prism                     []VisionPrescriptionLensSpecificationPrism `json:"prism,omitempty"`
	Add                       *Decimal                                   `json:"add,omitempty"`
	AddPrimitiveElement       *primitiveElement                          `json:"_add,omitempty"`
	Power                     *Decimal                                   `json:"power,omitempty"`
	PowerPrimitiveElement     *primitiveElement                          `json:"_power,omitempty"`
	BackCurve                 *Decimal                                   `json:"backCurve,omitempty"`
	BackCurvePrimitiveElement *primitiveElement                          `json:"_backCurve,omitempty"`
	Diameter                  *Decimal                                   `json:"diameter,omitempty"`
	DiameterPrimitiveElement  *primitiveElement                          `json:"_diameter,omitempty"`
	Duration                  *Quantity                                  `json:"duration,omitempty"`
	Color                     *String                                    `json:"color,omitempty"`
	ColorPrimitiveElement     *primitiveElement                          `json:"_color,omitempty"`
	Brand                     *String                                    `json:"brand,omitempty"`
	BrandPrimitiveElement     *primitiveElement                          `json:"_brand,omitempty"`
	Note                      []Annotation                               `json:"note,omitempty"`
}

func (r VisionPrescriptionLensSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r VisionPrescriptionLensSpecification) marshalJSON() jsonVisionPrescriptionLensSpecification {
	m := jsonVisionPrescriptionLensSpecification{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Product = r.Product
	if r.Eye.Value != nil {
		m.Eye = r.Eye
	}
	if r.Eye.Id != nil || r.Eye.Extension != nil {
		m.EyePrimitiveElement = &primitiveElement{Id: r.Eye.Id, Extension: r.Eye.Extension}
	}
	if r.Sphere != nil && r.Sphere.Value != nil {
		m.Sphere = r.Sphere
	}
	if r.Sphere != nil && (r.Sphere.Id != nil || r.Sphere.Extension != nil) {
		m.SpherePrimitiveElement = &primitiveElement{Id: r.Sphere.Id, Extension: r.Sphere.Extension}
	}
	if r.Cylinder != nil && r.Cylinder.Value != nil {
		m.Cylinder = r.Cylinder
	}
	if r.Cylinder != nil && (r.Cylinder.Id != nil || r.Cylinder.Extension != nil) {
		m.CylinderPrimitiveElement = &primitiveElement{Id: r.Cylinder.Id, Extension: r.Cylinder.Extension}
	}
	if r.Axis != nil && r.Axis.Value != nil {
		m.Axis = r.Axis
	}
	if r.Axis != nil && (r.Axis.Id != nil || r.Axis.Extension != nil) {
		m.AxisPrimitiveElement = &primitiveElement{Id: r.Axis.Id, Extension: r.Axis.Extension}
	}
	m.Prism = r.Prism
	if r.Add != nil && r.Add.Value != nil {
		m.Add = r.Add
	}
	if r.Add != nil && (r.Add.Id != nil || r.Add.Extension != nil) {
		m.AddPrimitiveElement = &primitiveElement{Id: r.Add.Id, Extension: r.Add.Extension}
	}
	if r.Power != nil && r.Power.Value != nil {
		m.Power = r.Power
	}
	if r.Power != nil && (r.Power.Id != nil || r.Power.Extension != nil) {
		m.PowerPrimitiveElement = &primitiveElement{Id: r.Power.Id, Extension: r.Power.Extension}
	}
	if r.BackCurve != nil && r.BackCurve.Value != nil {
		m.BackCurve = r.BackCurve
	}
	if r.BackCurve != nil && (r.BackCurve.Id != nil || r.BackCurve.Extension != nil) {
		m.BackCurvePrimitiveElement = &primitiveElement{Id: r.BackCurve.Id, Extension: r.BackCurve.Extension}
	}
	if r.Diameter != nil && r.Diameter.Value != nil {
		m.Diameter = r.Diameter
	}
	if r.Diameter != nil && (r.Diameter.Id != nil || r.Diameter.Extension != nil) {
		m.DiameterPrimitiveElement = &primitiveElement{Id: r.Diameter.Id, Extension: r.Diameter.Extension}
	}
	m.Duration = r.Duration
	if r.Color != nil && r.Color.Value != nil {
		m.Color = r.Color
	}
	if r.Color != nil && (r.Color.Id != nil || r.Color.Extension != nil) {
		m.ColorPrimitiveElement = &primitiveElement{Id: r.Color.Id, Extension: r.Color.Extension}
	}
	if r.Brand != nil && r.Brand.Value != nil {
		m.Brand = r.Brand
	}
	if r.Brand != nil && (r.Brand.Id != nil || r.Brand.Extension != nil) {
		m.BrandPrimitiveElement = &primitiveElement{Id: r.Brand.Id, Extension: r.Brand.Extension}
	}
	m.Note = r.Note
	return m
}
func (r *VisionPrescriptionLensSpecification) UnmarshalJSON(b []byte) error {
	var m jsonVisionPrescriptionLensSpecification
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *VisionPrescriptionLensSpecification) unmarshalJSON(m jsonVisionPrescriptionLensSpecification) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Product = m.Product
	r.Eye = m.Eye
	if m.EyePrimitiveElement != nil {
		r.Eye.Id = m.EyePrimitiveElement.Id
		r.Eye.Extension = m.EyePrimitiveElement.Extension
	}
	r.Sphere = m.Sphere
	if m.SpherePrimitiveElement != nil {
		if r.Sphere == nil {
			r.Sphere = &Decimal{}
		}
		r.Sphere.Id = m.SpherePrimitiveElement.Id
		r.Sphere.Extension = m.SpherePrimitiveElement.Extension
	}
	r.Cylinder = m.Cylinder
	if m.CylinderPrimitiveElement != nil {
		if r.Cylinder == nil {
			r.Cylinder = &Decimal{}
		}
		r.Cylinder.Id = m.CylinderPrimitiveElement.Id
		r.Cylinder.Extension = m.CylinderPrimitiveElement.Extension
	}
	r.Axis = m.Axis
	if m.AxisPrimitiveElement != nil {
		if r.Axis == nil {
			r.Axis = &Integer{}
		}
		r.Axis.Id = m.AxisPrimitiveElement.Id
		r.Axis.Extension = m.AxisPrimitiveElement.Extension
	}
	r.Prism = m.Prism
	r.Add = m.Add
	if m.AddPrimitiveElement != nil {
		if r.Add == nil {
			r.Add = &Decimal{}
		}
		r.Add.Id = m.AddPrimitiveElement.Id
		r.Add.Extension = m.AddPrimitiveElement.Extension
	}
	r.Power = m.Power
	if m.PowerPrimitiveElement != nil {
		if r.Power == nil {
			r.Power = &Decimal{}
		}
		r.Power.Id = m.PowerPrimitiveElement.Id
		r.Power.Extension = m.PowerPrimitiveElement.Extension
	}
	r.BackCurve = m.BackCurve
	if m.BackCurvePrimitiveElement != nil {
		if r.BackCurve == nil {
			r.BackCurve = &Decimal{}
		}
		r.BackCurve.Id = m.BackCurvePrimitiveElement.Id
		r.BackCurve.Extension = m.BackCurvePrimitiveElement.Extension
	}
	r.Diameter = m.Diameter
	if m.DiameterPrimitiveElement != nil {
		if r.Diameter == nil {
			r.Diameter = &Decimal{}
		}
		r.Diameter.Id = m.DiameterPrimitiveElement.Id
		r.Diameter.Extension = m.DiameterPrimitiveElement.Extension
	}
	r.Duration = m.Duration
	r.Color = m.Color
	if m.ColorPrimitiveElement != nil {
		if r.Color == nil {
			r.Color = &String{}
		}
		r.Color.Id = m.ColorPrimitiveElement.Id
		r.Color.Extension = m.ColorPrimitiveElement.Extension
	}
	r.Brand = m.Brand
	if m.BrandPrimitiveElement != nil {
		if r.Brand == nil {
			r.Brand = &String{}
		}
		r.Brand.Id = m.BrandPrimitiveElement.Id
		r.Brand.Extension = m.BrandPrimitiveElement.Extension
	}
	r.Note = m.Note
	return nil
}
func (r VisionPrescriptionLensSpecification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Product, xml.StartElement{Name: xml.Name{Local: "product"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Eye, xml.StartElement{Name: xml.Name{Local: "eye"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Sphere, xml.StartElement{Name: xml.Name{Local: "sphere"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Cylinder, xml.StartElement{Name: xml.Name{Local: "cylinder"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Axis, xml.StartElement{Name: xml.Name{Local: "axis"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Prism, xml.StartElement{Name: xml.Name{Local: "prism"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Add, xml.StartElement{Name: xml.Name{Local: "add"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Power, xml.StartElement{Name: xml.Name{Local: "power"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BackCurve, xml.StartElement{Name: xml.Name{Local: "backCurve"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Diameter, xml.StartElement{Name: xml.Name{Local: "diameter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Duration, xml.StartElement{Name: xml.Name{Local: "duration"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Color, xml.StartElement{Name: xml.Name{Local: "color"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Brand, xml.StartElement{Name: xml.Name{Local: "brand"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *VisionPrescriptionLensSpecification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "product":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Product = v
			case "eye":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Eye = v
			case "sphere":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Sphere = &v
			case "cylinder":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Cylinder = &v
			case "axis":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Axis = &v
			case "prism":
				var v VisionPrescriptionLensSpecificationPrism
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Prism = append(r.Prism, v)
			case "add":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Add = &v
			case "power":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Power = &v
			case "backCurve":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BackCurve = &v
			case "diameter":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Diameter = &v
			case "duration":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Duration = &v
			case "color":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Color = &v
			case "brand":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Brand = &v
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r VisionPrescriptionLensSpecification) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Allows for adjustment on two axis.
type VisionPrescriptionLensSpecificationPrism struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Amount of prism to compensate for eye alignment in fractional units.
	Amount Decimal
	// The relative base, or reference lens edge, for the prism.
	Base Code
}
type jsonVisionPrescriptionLensSpecificationPrism struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Amount                 Decimal           `json:"amount,omitempty"`
	AmountPrimitiveElement *primitiveElement `json:"_amount,omitempty"`
	Base                   Code              `json:"base,omitempty"`
	BasePrimitiveElement   *primitiveElement `json:"_base,omitempty"`
}

func (r VisionPrescriptionLensSpecificationPrism) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r VisionPrescriptionLensSpecificationPrism) marshalJSON() jsonVisionPrescriptionLensSpecificationPrism {
	m := jsonVisionPrescriptionLensSpecificationPrism{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Amount.Value != nil {
		m.Amount = r.Amount
	}
	if r.Amount.Id != nil || r.Amount.Extension != nil {
		m.AmountPrimitiveElement = &primitiveElement{Id: r.Amount.Id, Extension: r.Amount.Extension}
	}
	if r.Base.Value != nil {
		m.Base = r.Base
	}
	if r.Base.Id != nil || r.Base.Extension != nil {
		m.BasePrimitiveElement = &primitiveElement{Id: r.Base.Id, Extension: r.Base.Extension}
	}
	return m
}
func (r *VisionPrescriptionLensSpecificationPrism) UnmarshalJSON(b []byte) error {
	var m jsonVisionPrescriptionLensSpecificationPrism
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *VisionPrescriptionLensSpecificationPrism) unmarshalJSON(m jsonVisionPrescriptionLensSpecificationPrism) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Amount = m.Amount
	if m.AmountPrimitiveElement != nil {
		r.Amount.Id = m.AmountPrimitiveElement.Id
		r.Amount.Extension = m.AmountPrimitiveElement.Extension
	}
	r.Base = m.Base
	if m.BasePrimitiveElement != nil {
		r.Base.Id = m.BasePrimitiveElement.Id
		r.Base.Extension = m.BasePrimitiveElement.Extension
	}
	return nil
}
func (r VisionPrescriptionLensSpecificationPrism) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Amount, xml.StartElement{Name: xml.Name{Local: "amount"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Base, xml.StartElement{Name: xml.Name{Local: "base"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *VisionPrescriptionLensSpecificationPrism) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "amount":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			case "base":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Base = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r VisionPrescriptionLensSpecificationPrism) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
