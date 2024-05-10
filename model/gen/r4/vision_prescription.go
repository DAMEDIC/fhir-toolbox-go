package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
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
	Contained                     []containedResource                   `json:"contained,omitempty"`
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
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Created = r.Created
	if r.Created.Id != nil || r.Created.Extension != nil {
		m.CreatedPrimitiveElement = &primitiveElement{Id: r.Created.Id, Extension: r.Created.Extension}
	}
	m.Patient = r.Patient
	m.Encounter = r.Encounter
	m.DateWritten = r.DateWritten
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
func (r VisionPrescription) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
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
	m.Eye = r.Eye
	if r.Eye.Id != nil || r.Eye.Extension != nil {
		m.EyePrimitiveElement = &primitiveElement{Id: r.Eye.Id, Extension: r.Eye.Extension}
	}
	m.Sphere = r.Sphere
	if r.Sphere != nil && (r.Sphere.Id != nil || r.Sphere.Extension != nil) {
		m.SpherePrimitiveElement = &primitiveElement{Id: r.Sphere.Id, Extension: r.Sphere.Extension}
	}
	m.Cylinder = r.Cylinder
	if r.Cylinder != nil && (r.Cylinder.Id != nil || r.Cylinder.Extension != nil) {
		m.CylinderPrimitiveElement = &primitiveElement{Id: r.Cylinder.Id, Extension: r.Cylinder.Extension}
	}
	m.Axis = r.Axis
	if r.Axis != nil && (r.Axis.Id != nil || r.Axis.Extension != nil) {
		m.AxisPrimitiveElement = &primitiveElement{Id: r.Axis.Id, Extension: r.Axis.Extension}
	}
	m.Prism = r.Prism
	m.Add = r.Add
	if r.Add != nil && (r.Add.Id != nil || r.Add.Extension != nil) {
		m.AddPrimitiveElement = &primitiveElement{Id: r.Add.Id, Extension: r.Add.Extension}
	}
	m.Power = r.Power
	if r.Power != nil && (r.Power.Id != nil || r.Power.Extension != nil) {
		m.PowerPrimitiveElement = &primitiveElement{Id: r.Power.Id, Extension: r.Power.Extension}
	}
	m.BackCurve = r.BackCurve
	if r.BackCurve != nil && (r.BackCurve.Id != nil || r.BackCurve.Extension != nil) {
		m.BackCurvePrimitiveElement = &primitiveElement{Id: r.BackCurve.Id, Extension: r.BackCurve.Extension}
	}
	m.Diameter = r.Diameter
	if r.Diameter != nil && (r.Diameter.Id != nil || r.Diameter.Extension != nil) {
		m.DiameterPrimitiveElement = &primitiveElement{Id: r.Diameter.Id, Extension: r.Diameter.Extension}
	}
	m.Duration = r.Duration
	m.Color = r.Color
	if r.Color != nil && (r.Color.Id != nil || r.Color.Extension != nil) {
		m.ColorPrimitiveElement = &primitiveElement{Id: r.Color.Id, Extension: r.Color.Extension}
	}
	m.Brand = r.Brand
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
		r.Sphere.Id = m.SpherePrimitiveElement.Id
		r.Sphere.Extension = m.SpherePrimitiveElement.Extension
	}
	r.Cylinder = m.Cylinder
	if m.CylinderPrimitiveElement != nil {
		r.Cylinder.Id = m.CylinderPrimitiveElement.Id
		r.Cylinder.Extension = m.CylinderPrimitiveElement.Extension
	}
	r.Axis = m.Axis
	if m.AxisPrimitiveElement != nil {
		r.Axis.Id = m.AxisPrimitiveElement.Id
		r.Axis.Extension = m.AxisPrimitiveElement.Extension
	}
	r.Prism = m.Prism
	r.Add = m.Add
	if m.AddPrimitiveElement != nil {
		r.Add.Id = m.AddPrimitiveElement.Id
		r.Add.Extension = m.AddPrimitiveElement.Extension
	}
	r.Power = m.Power
	if m.PowerPrimitiveElement != nil {
		r.Power.Id = m.PowerPrimitiveElement.Id
		r.Power.Extension = m.PowerPrimitiveElement.Extension
	}
	r.BackCurve = m.BackCurve
	if m.BackCurvePrimitiveElement != nil {
		r.BackCurve.Id = m.BackCurvePrimitiveElement.Id
		r.BackCurve.Extension = m.BackCurvePrimitiveElement.Extension
	}
	r.Diameter = m.Diameter
	if m.DiameterPrimitiveElement != nil {
		r.Diameter.Id = m.DiameterPrimitiveElement.Id
		r.Diameter.Extension = m.DiameterPrimitiveElement.Extension
	}
	r.Duration = m.Duration
	r.Color = m.Color
	if m.ColorPrimitiveElement != nil {
		r.Color.Id = m.ColorPrimitiveElement.Id
		r.Color.Extension = m.ColorPrimitiveElement.Extension
	}
	r.Brand = m.Brand
	if m.BrandPrimitiveElement != nil {
		r.Brand.Id = m.BrandPrimitiveElement.Id
		r.Brand.Extension = m.BrandPrimitiveElement.Extension
	}
	r.Note = m.Note
	return nil
}
func (r VisionPrescriptionLensSpecification) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
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
	m.Amount = r.Amount
	if r.Amount.Id != nil || r.Amount.Extension != nil {
		m.AmountPrimitiveElement = &primitiveElement{Id: r.Amount.Id, Extension: r.Amount.Extension}
	}
	m.Base = r.Base
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
func (r VisionPrescriptionLensSpecificationPrism) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
