package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// The manufactured item as contained in the packaged medicinal product.
type MedicinalProductManufactured struct {
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
	// Dose form as manufactured and before any transformation into the pharmaceutical product.
	ManufacturedDoseForm CodeableConcept
	// The “real world” units in which the quantity of the manufactured item is described.
	UnitOfPresentation *CodeableConcept
	// The quantity or "count number" of the manufactured item.
	Quantity Quantity
	// Manufacturer of the item (Note that this should be named "manufacturer" but it currently causes technical issues).
	Manufacturer []Reference
	// Ingredient.
	Ingredient []Reference
	// Dimensions, color etc.
	PhysicalCharacteristics *ProdCharacteristic
	// Other codeable characteristics.
	OtherCharacteristics []CodeableConcept
}

func (r MedicinalProductManufactured) ResourceType() string {
	return "MedicinalProductManufactured"
}

type jsonMedicinalProductManufactured struct {
	ResourceType                  string              `json:"resourceType"`
	Id                            *Id                 `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement   `json:"_id,omitempty"`
	Meta                          *Meta               `json:"meta,omitempty"`
	ImplicitRules                 *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                      *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement   `json:"_language,omitempty"`
	Text                          *Narrative          `json:"text,omitempty"`
	Contained                     []containedResource `json:"contained,omitempty"`
	Extension                     []Extension         `json:"extension,omitempty"`
	ModifierExtension             []Extension         `json:"modifierExtension,omitempty"`
	ManufacturedDoseForm          CodeableConcept     `json:"manufacturedDoseForm,omitempty"`
	UnitOfPresentation            *CodeableConcept    `json:"unitOfPresentation,omitempty"`
	Quantity                      Quantity            `json:"quantity,omitempty"`
	Manufacturer                  []Reference         `json:"manufacturer,omitempty"`
	Ingredient                    []Reference         `json:"ingredient,omitempty"`
	PhysicalCharacteristics       *ProdCharacteristic `json:"physicalCharacteristics,omitempty"`
	OtherCharacteristics          []CodeableConcept   `json:"otherCharacteristics,omitempty"`
}

func (r MedicinalProductManufactured) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductManufactured) marshalJSON() jsonMedicinalProductManufactured {
	m := jsonMedicinalProductManufactured{}
	m.ResourceType = "MedicinalProductManufactured"
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
	m.ManufacturedDoseForm = r.ManufacturedDoseForm
	m.UnitOfPresentation = r.UnitOfPresentation
	m.Quantity = r.Quantity
	m.Manufacturer = r.Manufacturer
	m.Ingredient = r.Ingredient
	m.PhysicalCharacteristics = r.PhysicalCharacteristics
	m.OtherCharacteristics = r.OtherCharacteristics
	return m
}
func (r *MedicinalProductManufactured) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductManufactured
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductManufactured) unmarshalJSON(m jsonMedicinalProductManufactured) error {
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
	r.ManufacturedDoseForm = m.ManufacturedDoseForm
	r.UnitOfPresentation = m.UnitOfPresentation
	r.Quantity = m.Quantity
	r.Manufacturer = m.Manufacturer
	r.Ingredient = m.Ingredient
	r.PhysicalCharacteristics = m.PhysicalCharacteristics
	r.OtherCharacteristics = m.OtherCharacteristics
	return nil
}
func (r MedicinalProductManufactured) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
