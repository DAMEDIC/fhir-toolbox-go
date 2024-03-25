package r4

import (
	"encoding/json"
	"fmt"
)

// The interactions of the medicinal product with other medicinal products, or other forms of interactions.
type MedicinalProductInteraction struct {
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
	Contained []Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The medication for which this is a described interaction.
	Subject []Reference
	// The interaction described.
	Description *String
	// The specific medication, food or laboratory test that interacts.
	Interactant []MedicinalProductInteractionInteractant
	// The type of the interaction e.g. drug-drug interaction, drug-food interaction, drug-lab test interaction.
	Type *CodeableConcept
	// The effect of the interaction, for example "reduced gastric absorption of primary medication".
	Effect *CodeableConcept
	// The incidence of the interaction, e.g. theoretical, observed.
	Incidence *CodeableConcept
	// Actions for managing the interaction.
	Management *CodeableConcept
}
type jsonMedicinalProductInteraction struct {
	ResourceType                  string                                   `json:"resourceType"`
	Id                            *Id                                      `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement                        `json:"_id,omitempty"`
	Meta                          *Meta                                    `json:"meta,omitempty"`
	ImplicitRules                 *Uri                                     `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement                        `json:"_implicitRules,omitempty"`
	Language                      *Code                                    `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement                        `json:"_language,omitempty"`
	Text                          *Narrative                               `json:"text,omitempty"`
	Contained                     []containedResource                      `json:"contained,omitempty"`
	Extension                     []Extension                              `json:"extension,omitempty"`
	ModifierExtension             []Extension                              `json:"modifierExtension,omitempty"`
	Subject                       []Reference                              `json:"subject,omitempty"`
	Description                   *String                                  `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement                        `json:"_description,omitempty"`
	Interactant                   []MedicinalProductInteractionInteractant `json:"interactant,omitempty"`
	Type                          *CodeableConcept                         `json:"type,omitempty"`
	Effect                        *CodeableConcept                         `json:"effect,omitempty"`
	Incidence                     *CodeableConcept                         `json:"incidence,omitempty"`
	Management                    *CodeableConcept                         `json:"management,omitempty"`
}

func (r MedicinalProductInteraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductInteraction) marshalJSON() jsonMedicinalProductInteraction {
	m := jsonMedicinalProductInteraction{}
	m.ResourceType = "MedicinalProductInteraction"
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
	m.Subject = r.Subject
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Interactant = r.Interactant
	m.Type = r.Type
	m.Effect = r.Effect
	m.Incidence = r.Incidence
	m.Management = r.Management
	return m
}
func (r *MedicinalProductInteraction) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductInteraction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductInteraction) unmarshalJSON(m jsonMedicinalProductInteraction) error {
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
	r.Contained = make([]Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Subject = m.Subject
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Interactant = m.Interactant
	r.Type = m.Type
	r.Effect = m.Effect
	r.Incidence = m.Incidence
	r.Management = m.Management
	return nil
}
func (r MedicinalProductInteraction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The specific medication, food or laboratory test that interacts.
type MedicinalProductInteractionInteractant struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The specific medication, food or laboratory test that interacts.
	Item isMedicinalProductInteractionInteractantItem
}
type isMedicinalProductInteractionInteractantItem interface {
	isMedicinalProductInteractionInteractantItem()
}

func (r Reference) isMedicinalProductInteractionInteractantItem()       {}
func (r CodeableConcept) isMedicinalProductInteractionInteractantItem() {}

type jsonMedicinalProductInteractionInteractant struct {
	Id                  *string          `json:"id,omitempty"`
	Extension           []Extension      `json:"extension,omitempty"`
	ModifierExtension   []Extension      `json:"modifierExtension,omitempty"`
	ItemReference       *Reference       `json:"itemReference,omitempty"`
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

func (r MedicinalProductInteractionInteractant) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductInteractionInteractant) marshalJSON() jsonMedicinalProductInteractionInteractant {
	m := jsonMedicinalProductInteractionInteractant{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Item.(type) {
	case Reference:
		m.ItemReference = &v
	case *Reference:
		m.ItemReference = v
	case CodeableConcept:
		m.ItemCodeableConcept = &v
	case *CodeableConcept:
		m.ItemCodeableConcept = v
	}
	return m
}
func (r *MedicinalProductInteractionInteractant) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductInteractionInteractant
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductInteractionInteractant) unmarshalJSON(m jsonMedicinalProductInteractionInteractant) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.ItemReference != nil {
		if r.Item != nil {
			return fmt.Errorf("multiple values for field \"Item\"")
		}
		v := m.ItemReference
		r.Item = v
	}
	if m.ItemCodeableConcept != nil {
		if r.Item != nil {
			return fmt.Errorf("multiple values for field \"Item\"")
		}
		v := m.ItemCodeableConcept
		r.Item = v
	}
	return nil
}
func (r MedicinalProductInteractionInteractant) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
