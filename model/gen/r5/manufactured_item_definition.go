package r5

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
	"io"
)

// The definition and characteristics of a medicinal manufactured item, such as a tablet or capsule, as contained in a packaged medicinal product.
type ManufacturedItemDefinition struct {
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
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, nor can they have their own independent transaction scope. This is allowed to be a Parameters resource if and only if it is referenced by a resource that provides context/meaning.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Unique identifier.
	Identifier []Identifier
	// The status of this item. Enables tracking the life-cycle of the content.
	Status Code
	// A descriptive name applied to this item.
	Name *String
	// Dose form as manufactured and before any transformation into the pharmaceutical product.
	ManufacturedDoseForm CodeableConcept
	// The “real-world” units in which the quantity of the manufactured item is described.
	UnitOfPresentation *CodeableConcept
	// Manufacturer of the item, one of several possible.
	Manufacturer []Reference
	// Allows specifying that an item is on the market for sale, or that it is not available, and the dates and locations associated.
	MarketingStatus []MarketingStatus
	// The ingredients of this manufactured item. This is only needed if the ingredients are not specified by incoming references from the Ingredient resource.
	Ingredient []CodeableConcept
	// General characteristics of this item.
	Property []ManufacturedItemDefinitionProperty
	// Physical parts of the manufactured item, that it is intrisically made from. This is distinct from the ingredients that are part of its chemical makeup.
	Component []ManufacturedItemDefinitionComponent
}

// General characteristics of this item.
type ManufacturedItemDefinitionProperty struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code expressing the type of characteristic.
	Type CodeableConcept
	// A value for the characteristic.
	Value isManufacturedItemDefinitionPropertyValue
}
type isManufacturedItemDefinitionPropertyValue interface {
	isManufacturedItemDefinitionPropertyValue()
}

func (r CodeableConcept) isManufacturedItemDefinitionPropertyValue() {}
func (r Quantity) isManufacturedItemDefinitionPropertyValue()        {}
func (r Date) isManufacturedItemDefinitionPropertyValue()            {}
func (r Boolean) isManufacturedItemDefinitionPropertyValue()         {}
func (r Markdown) isManufacturedItemDefinitionPropertyValue()        {}
func (r Attachment) isManufacturedItemDefinitionPropertyValue()      {}
func (r Reference) isManufacturedItemDefinitionPropertyValue()       {}

// Physical parts of the manufactured item, that it is intrisically made from. This is distinct from the ingredients that are part of its chemical makeup.
type ManufacturedItemDefinitionComponent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Defining type of the component e.g. shell, layer, ink.
	Type CodeableConcept
	// The function of this component within the item e.g. delivers active ingredient, masks taste.
	Function []CodeableConcept
	// The measurable amount of total quantity of all substances in the component, expressable in different ways (e.g. by mass or volume).
	Amount []Quantity
	// A reference to a constituent of the manufactured item as a whole, linked here so that its component location within the item can be indicated. This not where the item's ingredient are primarily stated (for which see Ingredient.for or ManufacturedItemDefinition.ingredient).
	Constituent []ManufacturedItemDefinitionComponentConstituent
	// General characteristics of this component.
	Property []ManufacturedItemDefinitionProperty
	// A component that this component contains or is made from.
	Component []ManufacturedItemDefinitionComponent
}

// A reference to a constituent of the manufactured item as a whole, linked here so that its component location within the item can be indicated. This not where the item's ingredient are primarily stated (for which see Ingredient.for or ManufacturedItemDefinition.ingredient).
type ManufacturedItemDefinitionComponentConstituent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The measurable amount of the substance, expressable in different ways (e.g. by mass or volume).
	Amount []Quantity
	// The physical location of the constituent/ingredient within the component. Example – if the component is the bead in the capsule, then the location would be where the ingredient resides within the product part – intragranular, extra-granular, etc.
	Location []CodeableConcept
	// The function of this constituent within the component e.g. binder.
	Function []CodeableConcept
	// The ingredient that is the constituent of the given component.
	HasIngredient []CodeableReference
}

func (r ManufacturedItemDefinition) ResourceType() string {
	return "ManufacturedItemDefinition"
}
func (r ManufacturedItemDefinition) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r ManufacturedItemDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ManufacturedItemDefinition) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ManufacturedItemDefinition) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"ManufacturedItemDefinition\""))
	if err != nil {
		return err
	}
	setComma := true
	if r.Id != nil && r.Id.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		p := primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_id\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Meta != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"meta\":"))
		if err != nil {
			return err
		}
		err = r.Meta.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"implicitRules\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ImplicitRules)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		p := primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_implicitRules\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && r.Language.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"language\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Language)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		p := primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_language\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Text != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"text\":"))
		if err != nil {
			return err
		}
		err = r.Text.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Contained) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"contained\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, c := range r.Contained {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = ContainedResource{c}.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Identifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"identifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Identifier {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"status\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Status)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		p := primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_status\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Name != nil && r.Name.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"name\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Name)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		p := primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_name\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"manufacturedDoseForm\":"))
	if err != nil {
		return err
	}
	err = r.ManufacturedDoseForm.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.UnitOfPresentation != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"unitOfPresentation\":"))
		if err != nil {
			return err
		}
		err = r.UnitOfPresentation.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Manufacturer) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"manufacturer\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Manufacturer {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.MarketingStatus) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"marketingStatus\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.MarketingStatus {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Ingredient) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"ingredient\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Ingredient {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Property) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"property\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Property {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Component) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"component\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Component {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r ManufacturedItemDefinitionProperty) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ManufacturedItemDefinitionProperty) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"type\":"))
	if err != nil {
		return err
	}
	err = r.Type.marshalJSON(w)
	if err != nil {
		return err
	}
	switch v := r.Value.(type) {
	case CodeableConcept:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCodeableConcept\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	case *CodeableConcept:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCodeableConcept\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	case Quantity:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueQuantity\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	case *Quantity:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueQuantity\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	case Date:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDate\":"))
			if err != nil {
				return err
			}
			var b bytes.Buffer
			enc := json.NewEncoder(&b)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
			_, err = w.Write(b.Bytes())
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_valueDate\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Date:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDate\":"))
			if err != nil {
				return err
			}
			var b bytes.Buffer
			enc := json.NewEncoder(&b)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
			_, err = w.Write(b.Bytes())
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_valueDate\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Boolean:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueBoolean\":"))
			if err != nil {
				return err
			}
			var b bytes.Buffer
			enc := json.NewEncoder(&b)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
			_, err = w.Write(b.Bytes())
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_valueBoolean\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Boolean:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueBoolean\":"))
			if err != nil {
				return err
			}
			var b bytes.Buffer
			enc := json.NewEncoder(&b)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
			_, err = w.Write(b.Bytes())
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_valueBoolean\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Markdown:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueMarkdown\":"))
			if err != nil {
				return err
			}
			var b bytes.Buffer
			enc := json.NewEncoder(&b)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
			_, err = w.Write(b.Bytes())
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_valueMarkdown\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Markdown:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueMarkdown\":"))
			if err != nil {
				return err
			}
			var b bytes.Buffer
			enc := json.NewEncoder(&b)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
			_, err = w.Write(b.Bytes())
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_valueMarkdown\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Attachment:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAttachment\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	case *Attachment:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAttachment\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	case Reference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueReference\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	case *Reference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueReference\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r ManufacturedItemDefinitionComponent) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ManufacturedItemDefinitionComponent) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"type\":"))
	if err != nil {
		return err
	}
	err = r.Type.marshalJSON(w)
	if err != nil {
		return err
	}
	if len(r.Function) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"function\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Function {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Amount) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"amount\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Amount {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Constituent) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"constituent\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Constituent {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Property) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"property\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Property {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Component) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"component\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Component {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r ManufacturedItemDefinitionComponentConstituent) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ManufacturedItemDefinitionComponentConstituent) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Amount) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"amount\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Amount {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Location) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"location\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Location {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Function) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"function\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Function {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.HasIngredient) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"hasIngredient\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.HasIngredient {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r *ManufacturedItemDefinition) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *ManufacturedItemDefinition) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ManufacturedItemDefinition element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ManufacturedItemDefinition element", t)
		}
		switch f {
		case "resourceType":
			_, err := d.Token()
			if err != nil {
				return err
			}
		case "id":
			var v Id
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Value = v.Value
		case "_id":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Id = v.Id
			r.Id.Extension = v.Extension
		case "meta":
			var v Meta
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Meta = &v
		case "implicitRules":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Value = v.Value
		case "_implicitRules":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Id = v.Id
			r.ImplicitRules.Extension = v.Extension
		case "language":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Value = v.Value
		case "_language":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Id = v.Id
			r.Language.Extension = v.Extension
		case "text":
			var v Narrative
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Text = &v
		case "contained":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinition element", t)
			}
			for d.More() {
				var v ContainedResource
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Contained = append(r.Contained, v.Resource)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinition element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinition element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinition element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinition element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinition element", t)
			}
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinition element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinition element", t)
			}
		case "status":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Status.Value = v.Value
		case "_status":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Status.Id = v.Id
			r.Status.Extension = v.Extension
		case "name":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Name == nil {
				r.Name = &String{}
			}
			r.Name.Value = v.Value
		case "_name":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Name == nil {
				r.Name = &String{}
			}
			r.Name.Id = v.Id
			r.Name.Extension = v.Extension
		case "manufacturedDoseForm":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ManufacturedDoseForm = v
		case "unitOfPresentation":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.UnitOfPresentation = &v
		case "manufacturer":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinition element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Manufacturer = append(r.Manufacturer, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinition element", t)
			}
		case "marketingStatus":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinition element", t)
			}
			for d.More() {
				var v MarketingStatus
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.MarketingStatus = append(r.MarketingStatus, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinition element", t)
			}
		case "ingredient":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinition element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Ingredient = append(r.Ingredient, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinition element", t)
			}
		case "property":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinition element", t)
			}
			for d.More() {
				var v ManufacturedItemDefinitionProperty
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Property = append(r.Property, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinition element", t)
			}
		case "component":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinition element", t)
			}
			for d.More() {
				var v ManufacturedItemDefinitionComponent
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Component = append(r.Component, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinition element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ManufacturedItemDefinition", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ManufacturedItemDefinition element", t)
	}
	return nil
}
func (r *ManufacturedItemDefinitionProperty) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ManufacturedItemDefinitionProperty element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ManufacturedItemDefinitionProperty element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionProperty element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionProperty element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionProperty element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionProperty element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = v
		case "valueCodeableConcept":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueQuantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueDate":
			var v Date
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Date{
					Extension: r.Value.(Date).Extension,
					Id:        r.Value.(Date).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Date{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Date).Value,
				}
			} else {
				r.Value = Date{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueBoolean":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Boolean{
					Extension: r.Value.(Boolean).Extension,
					Id:        r.Value.(Boolean).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueBoolean":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Boolean{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Boolean).Value,
				}
			} else {
				r.Value = Boolean{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueMarkdown":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Markdown{
					Extension: r.Value.(Markdown).Extension,
					Id:        r.Value.(Markdown).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueMarkdown":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Markdown{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Markdown).Value,
				}
			} else {
				r.Value = Markdown{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueAttachment":
			var v Attachment
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueReference":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		default:
			return fmt.Errorf("invalid field: %s in ManufacturedItemDefinitionProperty", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ManufacturedItemDefinitionProperty element", t)
	}
	return nil
}
func (r *ManufacturedItemDefinitionComponent) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ManufacturedItemDefinitionComponent element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ManufacturedItemDefinitionComponent element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionComponent element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionComponent element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionComponent element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionComponent element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = v
		case "function":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionComponent element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Function = append(r.Function, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionComponent element", t)
			}
		case "amount":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionComponent element", t)
			}
			for d.More() {
				var v Quantity
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Amount = append(r.Amount, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionComponent element", t)
			}
		case "constituent":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionComponent element", t)
			}
			for d.More() {
				var v ManufacturedItemDefinitionComponentConstituent
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Constituent = append(r.Constituent, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionComponent element", t)
			}
		case "property":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionComponent element", t)
			}
			for d.More() {
				var v ManufacturedItemDefinitionProperty
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Property = append(r.Property, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionComponent element", t)
			}
		case "component":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionComponent element", t)
			}
			for d.More() {
				var v ManufacturedItemDefinitionComponent
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Component = append(r.Component, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionComponent element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ManufacturedItemDefinitionComponent", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ManufacturedItemDefinitionComponent element", t)
	}
	return nil
}
func (r *ManufacturedItemDefinitionComponentConstituent) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ManufacturedItemDefinitionComponentConstituent element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ManufacturedItemDefinitionComponentConstituent element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionComponentConstituent element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionComponentConstituent element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionComponentConstituent element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionComponentConstituent element", t)
			}
		case "amount":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionComponentConstituent element", t)
			}
			for d.More() {
				var v Quantity
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Amount = append(r.Amount, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionComponentConstituent element", t)
			}
		case "location":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionComponentConstituent element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Location = append(r.Location, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionComponentConstituent element", t)
			}
		case "function":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionComponentConstituent element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Function = append(r.Function, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionComponentConstituent element", t)
			}
		case "hasIngredient":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ManufacturedItemDefinitionComponentConstituent element", t)
			}
			for d.More() {
				var v CodeableReference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.HasIngredient = append(r.HasIngredient, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ManufacturedItemDefinitionComponentConstituent element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ManufacturedItemDefinitionComponentConstituent", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ManufacturedItemDefinitionComponentConstituent element", t)
	}
	return nil
}
func (r ManufacturedItemDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "ManufacturedItemDefinition"
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
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ManufacturedDoseForm, xml.StartElement{Name: xml.Name{Local: "manufacturedDoseForm"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UnitOfPresentation, xml.StartElement{Name: xml.Name{Local: "unitOfPresentation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Manufacturer, xml.StartElement{Name: xml.Name{Local: "manufacturer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MarketingStatus, xml.StartElement{Name: xml.Name{Local: "marketingStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Ingredient, xml.StartElement{Name: xml.Name{Local: "ingredient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Property, xml.StartElement{Name: xml.Name{Local: "property"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Component, xml.StartElement{Name: xml.Name{Local: "component"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ManufacturedItemDefinitionProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	switch v := r.Value.(type) {
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCodeableConcept"}})
		if err != nil {
			return err
		}
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueQuantity"}})
		if err != nil {
			return err
		}
	case Date, *Date:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDate"}})
		if err != nil {
			return err
		}
	case Boolean, *Boolean:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueBoolean"}})
		if err != nil {
			return err
		}
	case Markdown, *Markdown:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueMarkdown"}})
		if err != nil {
			return err
		}
	case Attachment, *Attachment:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAttachment"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueReference"}})
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
func (r ManufacturedItemDefinitionComponent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Function, xml.StartElement{Name: xml.Name{Local: "function"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Amount, xml.StartElement{Name: xml.Name{Local: "amount"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Constituent, xml.StartElement{Name: xml.Name{Local: "constituent"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Property, xml.StartElement{Name: xml.Name{Local: "property"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Component, xml.StartElement{Name: xml.Name{Local: "component"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ManufacturedItemDefinitionComponentConstituent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Location, xml.StartElement{Name: xml.Name{Local: "location"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Function, xml.StartElement{Name: xml.Name{Local: "function"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.HasIngredient, xml.StartElement{Name: xml.Name{Local: "hasIngredient"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ManufacturedItemDefinition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "manufacturedDoseForm":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ManufacturedDoseForm = v
			case "unitOfPresentation":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UnitOfPresentation = &v
			case "manufacturer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Manufacturer = append(r.Manufacturer, v)
			case "marketingStatus":
				var v MarketingStatus
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MarketingStatus = append(r.MarketingStatus, v)
			case "ingredient":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Ingredient = append(r.Ingredient, v)
			case "property":
				var v ManufacturedItemDefinitionProperty
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Property = append(r.Property, v)
			case "component":
				var v ManufacturedItemDefinitionComponent
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Component = append(r.Component, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ManufacturedItemDefinitionProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "valueCodeableConcept":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueQuantity":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDate":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueBoolean":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueMarkdown":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAttachment":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueReference":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Reference
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
func (r *ManufacturedItemDefinitionComponent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "function":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Function = append(r.Function, v)
			case "amount":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = append(r.Amount, v)
			case "constituent":
				var v ManufacturedItemDefinitionComponentConstituent
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Constituent = append(r.Constituent, v)
			case "property":
				var v ManufacturedItemDefinitionProperty
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Property = append(r.Property, v)
			case "component":
				var v ManufacturedItemDefinitionComponent
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Component = append(r.Component, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ManufacturedItemDefinitionComponentConstituent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = append(r.Amount, v)
			case "location":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = append(r.Location, v)
			case "function":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Function = append(r.Function, v)
			case "hasIngredient":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.HasIngredient = append(r.HasIngredient, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
