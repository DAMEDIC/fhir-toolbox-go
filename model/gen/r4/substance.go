package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A homogeneous material with a definite composition.
type Substance struct {
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
	// Unique identifier for the substance.
	Identifier []Identifier
	// A code to indicate if the substance is actively used.
	Status *Code
	// A code that classifies the general type of substance.  This is used  for searching, sorting and display purposes.
	Category []CodeableConcept
	// A code (or set of codes) that identify this substance.
	Code CodeableConcept
	// A description of the substance - its appearance, handling requirements, and other usage notes.
	Description *String
	// Substance may be used to describe a kind of substance, or a specific package/container of the substance: an instance.
	Instance []SubstanceInstance
	// A substance can be composed of other substances.
	Ingredient []SubstanceIngredient
}

func (r Substance) ResourceType() string {
	return "Substance"
}
func (r Substance) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonSubstance struct {
	ResourceType                  string                `json:"resourceType"`
	Id                            *Id                   `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement     `json:"_id,omitempty"`
	Meta                          *Meta                 `json:"meta,omitempty"`
	ImplicitRules                 *Uri                  `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement     `json:"_implicitRules,omitempty"`
	Language                      *Code                 `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement     `json:"_language,omitempty"`
	Text                          *Narrative            `json:"text,omitempty"`
	Contained                     []ContainedResource   `json:"contained,omitempty"`
	Extension                     []Extension           `json:"extension,omitempty"`
	ModifierExtension             []Extension           `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier          `json:"identifier,omitempty"`
	Status                        *Code                 `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement     `json:"_status,omitempty"`
	Category                      []CodeableConcept     `json:"category,omitempty"`
	Code                          CodeableConcept       `json:"code,omitempty"`
	Description                   *String               `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement     `json:"_description,omitempty"`
	Instance                      []SubstanceInstance   `json:"instance,omitempty"`
	Ingredient                    []SubstanceIngredient `json:"ingredient,omitempty"`
}

func (r Substance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Substance) marshalJSON() jsonSubstance {
	m := jsonSubstance{}
	m.ResourceType = "Substance"
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
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Status = r.Status
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Category = r.Category
	m.Code = r.Code
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Instance = r.Instance
	m.Ingredient = r.Ingredient
	return m
}
func (r *Substance) UnmarshalJSON(b []byte) error {
	var m jsonSubstance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Substance) unmarshalJSON(m jsonSubstance) error {
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
	r.Category = m.Category
	r.Code = m.Code
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Instance = m.Instance
	r.Ingredient = m.Ingredient
	return nil
}
func (r Substance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Substance may be used to describe a kind of substance, or a specific package/container of the substance: an instance.
type SubstanceInstance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identifier associated with the package/container (usually a label affixed directly).
	Identifier *Identifier
	// When the substance is no longer valid to use. For some substances, a single arbitrary date is used for expiry.
	Expiry *DateTime
	// The amount of the substance.
	Quantity *Quantity
}
type jsonSubstanceInstance struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Identifier             *Identifier       `json:"identifier,omitempty"`
	Expiry                 *DateTime         `json:"expiry,omitempty"`
	ExpiryPrimitiveElement *primitiveElement `json:"_expiry,omitempty"`
	Quantity               *Quantity         `json:"quantity,omitempty"`
}

func (r SubstanceInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceInstance) marshalJSON() jsonSubstanceInstance {
	m := jsonSubstanceInstance{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Expiry = r.Expiry
	if r.Expiry != nil && (r.Expiry.Id != nil || r.Expiry.Extension != nil) {
		m.ExpiryPrimitiveElement = &primitiveElement{Id: r.Expiry.Id, Extension: r.Expiry.Extension}
	}
	m.Quantity = r.Quantity
	return m
}
func (r *SubstanceInstance) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceInstance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceInstance) unmarshalJSON(m jsonSubstanceInstance) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Expiry = m.Expiry
	if m.ExpiryPrimitiveElement != nil {
		r.Expiry.Id = m.ExpiryPrimitiveElement.Id
		r.Expiry.Extension = m.ExpiryPrimitiveElement.Extension
	}
	r.Quantity = m.Quantity
	return nil
}
func (r SubstanceInstance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A substance can be composed of other substances.
type SubstanceIngredient struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The amount of the ingredient in the substance - a concentration ratio.
	Quantity *Ratio
	// Another substance that is a component of this substance.
	Substance isSubstanceIngredientSubstance
}
type isSubstanceIngredientSubstance interface {
	isSubstanceIngredientSubstance()
}

func (r CodeableConcept) isSubstanceIngredientSubstance() {}
func (r Reference) isSubstanceIngredientSubstance()       {}

type jsonSubstanceIngredient struct {
	Id                       *string          `json:"id,omitempty"`
	Extension                []Extension      `json:"extension,omitempty"`
	ModifierExtension        []Extension      `json:"modifierExtension,omitempty"`
	Quantity                 *Ratio           `json:"quantity,omitempty"`
	SubstanceCodeableConcept *CodeableConcept `json:"substanceCodeableConcept,omitempty"`
	SubstanceReference       *Reference       `json:"substanceReference,omitempty"`
}

func (r SubstanceIngredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceIngredient) marshalJSON() jsonSubstanceIngredient {
	m := jsonSubstanceIngredient{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Quantity = r.Quantity
	switch v := r.Substance.(type) {
	case CodeableConcept:
		m.SubstanceCodeableConcept = &v
	case *CodeableConcept:
		m.SubstanceCodeableConcept = v
	case Reference:
		m.SubstanceReference = &v
	case *Reference:
		m.SubstanceReference = v
	}
	return m
}
func (r *SubstanceIngredient) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceIngredient
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceIngredient) unmarshalJSON(m jsonSubstanceIngredient) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Quantity = m.Quantity
	if m.SubstanceCodeableConcept != nil {
		if r.Substance != nil {
			return fmt.Errorf("multiple values for field \"Substance\"")
		}
		v := m.SubstanceCodeableConcept
		r.Substance = v
	}
	if m.SubstanceReference != nil {
		if r.Substance != nil {
			return fmt.Errorf("multiple values for field \"Substance\"")
		}
		v := m.SubstanceReference
		r.Substance = v
	}
	return nil
}
func (r SubstanceIngredient) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
