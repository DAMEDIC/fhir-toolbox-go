package r4

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	model "github.com/DAMEDIC/fhir-toolbox-go/model"
	"io"
	"slices"
	"unsafe"
)

// A material substance originating from a biological entity intended to be transplanted or infused
// into another (possibly the same) biological entity.
type BiologicallyDerivedProduct struct {
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
	// This records identifiers associated with this biologically derived product instance that are defined by business processes and/or used to refer to it when a direct URL reference to the resource itself is not appropriate (e.g. in CDA documents, or in written / printed documentation).
	Identifier []Identifier
	// Broad category of this product.
	ProductCategory *Code
	// A code that identifies the kind of this biologically derived product (SNOMED Ctcode).
	ProductCode *CodeableConcept
	// Whether the product is currently available.
	Status *Code
	// Procedure request to obtain this biologically derived product.
	Request []Reference
	// Number of discrete units within this product.
	Quantity *Integer
	// Parent product (if any).
	Parent []Reference
	// How this product was collected.
	Collection *BiologicallyDerivedProductCollection
	// Any processing of the product during collection that does not change the fundamental nature of the product. For example adding anti-coagulants during the collection of Peripheral Blood Stem Cells.
	Processing []BiologicallyDerivedProductProcessing
	// Any manipulation of product post-collection that is intended to alter the product.  For example a buffy-coat enrichment or CD8 reduction of Peripheral Blood Stem Cells to make it more suitable for infusion.
	Manipulation *BiologicallyDerivedProductManipulation
	// Product storage.
	Storage []BiologicallyDerivedProductStorage
}

// How this product was collected.
type BiologicallyDerivedProductCollection struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Healthcare professional who is performing the collection.
	Collector *Reference
	// The patient or entity, such as a hospital or vendor in the case of a processed/manipulated/manufactured product, providing the product.
	Source *Reference
	// Time of product collection.
	Collected isBiologicallyDerivedProductCollectionCollected
}
type isBiologicallyDerivedProductCollectionCollected interface {
	model.Element
	isBiologicallyDerivedProductCollectionCollected()
}

func (r DateTime) isBiologicallyDerivedProductCollectionCollected() {}
func (r Period) isBiologicallyDerivedProductCollectionCollected()   {}

// Any processing of the product during collection that does not change the fundamental nature of the product. For example adding anti-coagulants during the collection of Peripheral Blood Stem Cells.
type BiologicallyDerivedProductProcessing struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Description of of processing.
	Description *String
	// Procesing code.
	Procedure *CodeableConcept
	// Substance added during processing.
	Additive *Reference
	// Time of processing.
	Time isBiologicallyDerivedProductProcessingTime
}
type isBiologicallyDerivedProductProcessingTime interface {
	model.Element
	isBiologicallyDerivedProductProcessingTime()
}

func (r DateTime) isBiologicallyDerivedProductProcessingTime() {}
func (r Period) isBiologicallyDerivedProductProcessingTime()   {}

// Any manipulation of product post-collection that is intended to alter the product.  For example a buffy-coat enrichment or CD8 reduction of Peripheral Blood Stem Cells to make it more suitable for infusion.
type BiologicallyDerivedProductManipulation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Description of manipulation.
	Description *String
	// Time of manipulation.
	Time isBiologicallyDerivedProductManipulationTime
}
type isBiologicallyDerivedProductManipulationTime interface {
	model.Element
	isBiologicallyDerivedProductManipulationTime()
}

func (r DateTime) isBiologicallyDerivedProductManipulationTime() {}
func (r Period) isBiologicallyDerivedProductManipulationTime()   {}

// Product storage.
type BiologicallyDerivedProductStorage struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Description of storage.
	Description *String
	// Storage temperature.
	Temperature *Decimal
	// Temperature scale used.
	Scale *Code
	// Storage timeperiod.
	Duration *Period
}

func (r BiologicallyDerivedProduct) ResourceType() string {
	return "BiologicallyDerivedProduct"
}
func (r BiologicallyDerivedProduct) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r BiologicallyDerivedProduct) MemSize() int {
	var emptyIface any
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += r.Id.MemSize()
	}
	if r.Meta != nil {
		s += r.Meta.MemSize()
	}
	if r.ImplicitRules != nil {
		s += r.ImplicitRules.MemSize()
	}
	if r.Language != nil {
		s += r.Language.MemSize()
	}
	if r.Text != nil {
		s += r.Text.MemSize()
	}
	for _, i := range r.Contained {
		s += i.MemSize()
	}
	s += (cap(r.Contained) - len(r.Contained)) * int(unsafe.Sizeof(emptyIface))
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	if r.ProductCategory != nil {
		s += r.ProductCategory.MemSize()
	}
	if r.ProductCode != nil {
		s += r.ProductCode.MemSize()
	}
	if r.Status != nil {
		s += r.Status.MemSize()
	}
	for _, i := range r.Request {
		s += i.MemSize()
	}
	s += (cap(r.Request) - len(r.Request)) * int(unsafe.Sizeof(Reference{}))
	if r.Quantity != nil {
		s += r.Quantity.MemSize()
	}
	for _, i := range r.Parent {
		s += i.MemSize()
	}
	s += (cap(r.Parent) - len(r.Parent)) * int(unsafe.Sizeof(Reference{}))
	if r.Collection != nil {
		s += r.Collection.MemSize()
	}
	for _, i := range r.Processing {
		s += i.MemSize()
	}
	s += (cap(r.Processing) - len(r.Processing)) * int(unsafe.Sizeof(BiologicallyDerivedProductProcessing{}))
	if r.Manipulation != nil {
		s += r.Manipulation.MemSize()
	}
	for _, i := range r.Storage {
		s += i.MemSize()
	}
	s += (cap(r.Storage) - len(r.Storage)) * int(unsafe.Sizeof(BiologicallyDerivedProductStorage{}))
	return s
}
func (r BiologicallyDerivedProductCollection) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Collector != nil {
		s += r.Collector.MemSize()
	}
	if r.Source != nil {
		s += r.Source.MemSize()
	}
	if r.Collected != nil {
		s += r.Collected.MemSize()
	}
	return s
}
func (r BiologicallyDerivedProductProcessing) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Description != nil {
		s += r.Description.MemSize()
	}
	if r.Procedure != nil {
		s += r.Procedure.MemSize()
	}
	if r.Additive != nil {
		s += r.Additive.MemSize()
	}
	if r.Time != nil {
		s += r.Time.MemSize()
	}
	return s
}
func (r BiologicallyDerivedProductManipulation) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Description != nil {
		s += r.Description.MemSize()
	}
	if r.Time != nil {
		s += r.Time.MemSize()
	}
	return s
}
func (r BiologicallyDerivedProductStorage) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Description != nil {
		s += r.Description.MemSize()
	}
	if r.Temperature != nil {
		s += r.Temperature.MemSize()
	}
	if r.Scale != nil {
		s += r.Scale.MemSize()
	}
	if r.Duration != nil {
		s += r.Duration.MemSize()
	}
	return s
}
func (r BiologicallyDerivedProduct) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r BiologicallyDerivedProductCollection) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r BiologicallyDerivedProductProcessing) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r BiologicallyDerivedProductManipulation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r BiologicallyDerivedProductStorage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r BiologicallyDerivedProduct) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r BiologicallyDerivedProduct) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"BiologicallyDerivedProduct\""))
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
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
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
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ImplicitRules)
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
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Language)
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
	if r.ProductCategory != nil && r.ProductCategory.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"productCategory\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ProductCategory)
		if err != nil {
			return err
		}
	}
	if r.ProductCategory != nil && (r.ProductCategory.Id != nil || r.ProductCategory.Extension != nil) {
		p := primitiveElement{Id: r.ProductCategory.Id, Extension: r.ProductCategory.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_productCategory\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ProductCode != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"productCode\":"))
		if err != nil {
			return err
		}
		err = r.ProductCode.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Status != nil && r.Status.Value != nil {
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
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Status)
		if err != nil {
			return err
		}
	}
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
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
	if len(r.Request) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"request\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Request {
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
	if r.Quantity != nil && r.Quantity.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"quantity\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Quantity)
		if err != nil {
			return err
		}
	}
	if r.Quantity != nil && (r.Quantity.Id != nil || r.Quantity.Extension != nil) {
		p := primitiveElement{Id: r.Quantity.Id, Extension: r.Quantity.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_quantity\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Parent) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"parent\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Parent {
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
	if r.Collection != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"collection\":"))
		if err != nil {
			return err
		}
		err = r.Collection.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Processing) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"processing\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Processing {
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
	if r.Manipulation != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"manipulation\":"))
		if err != nil {
			return err
		}
		err = r.Manipulation.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Storage) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"storage\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Storage {
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
func (r BiologicallyDerivedProductCollection) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r BiologicallyDerivedProductCollection) marshalJSON(w io.Writer) error {
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
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
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
	if r.Collector != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"collector\":"))
		if err != nil {
			return err
		}
		err = r.Collector.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Source != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"source\":"))
		if err != nil {
			return err
		}
		err = r.Source.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	switch v := r.Collected.(type) {
	case DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"collectedDateTime\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
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
			_, err = w.Write([]byte("\"_collectedDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"collectedDateTime\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
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
			_, err = w.Write([]byte("\"_collectedDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"collectedPeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"collectedPeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
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
func (r BiologicallyDerivedProductProcessing) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r BiologicallyDerivedProductProcessing) marshalJSON(w io.Writer) error {
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
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
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
	if r.Description != nil && r.Description.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"description\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Description)
		if err != nil {
			return err
		}
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		p := primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_description\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Procedure != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"procedure\":"))
		if err != nil {
			return err
		}
		err = r.Procedure.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Additive != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"additive\":"))
		if err != nil {
			return err
		}
		err = r.Additive.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	switch v := r.Time.(type) {
	case DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"timeDateTime\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
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
			_, err = w.Write([]byte("\"_timeDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"timeDateTime\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
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
			_, err = w.Write([]byte("\"_timeDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timePeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timePeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
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
func (r BiologicallyDerivedProductManipulation) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r BiologicallyDerivedProductManipulation) marshalJSON(w io.Writer) error {
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
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
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
	if r.Description != nil && r.Description.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"description\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Description)
		if err != nil {
			return err
		}
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		p := primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_description\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	switch v := r.Time.(type) {
	case DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"timeDateTime\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
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
			_, err = w.Write([]byte("\"_timeDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"timeDateTime\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
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
			_, err = w.Write([]byte("\"_timeDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timePeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timePeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
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
func (r BiologicallyDerivedProductStorage) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r BiologicallyDerivedProductStorage) marshalJSON(w io.Writer) error {
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
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
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
	if r.Description != nil && r.Description.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"description\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Description)
		if err != nil {
			return err
		}
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		p := primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_description\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Temperature != nil && r.Temperature.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"temperature\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Temperature)
		if err != nil {
			return err
		}
	}
	if r.Temperature != nil && (r.Temperature.Id != nil || r.Temperature.Extension != nil) {
		p := primitiveElement{Id: r.Temperature.Id, Extension: r.Temperature.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_temperature\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Scale != nil && r.Scale.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"scale\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Scale)
		if err != nil {
			return err
		}
	}
	if r.Scale != nil && (r.Scale.Id != nil || r.Scale.Extension != nil) {
		p := primitiveElement{Id: r.Scale.Id, Extension: r.Scale.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_scale\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Duration != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"duration\":"))
		if err != nil {
			return err
		}
		err = r.Duration.marshalJSON(w)
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
func (r *BiologicallyDerivedProduct) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *BiologicallyDerivedProduct) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in BiologicallyDerivedProduct element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in BiologicallyDerivedProduct element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProduct element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProduct element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProduct element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProduct element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProduct element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProduct element", t)
			}
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProduct element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProduct element", t)
			}
		case "productCategory":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ProductCategory == nil {
				r.ProductCategory = &Code{}
			}
			r.ProductCategory.Value = v.Value
		case "_productCategory":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ProductCategory == nil {
				r.ProductCategory = &Code{}
			}
			r.ProductCategory.Id = v.Id
			r.ProductCategory.Extension = v.Extension
		case "productCode":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ProductCode = &v
		case "status":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Status == nil {
				r.Status = &Code{}
			}
			r.Status.Value = v.Value
		case "_status":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Status == nil {
				r.Status = &Code{}
			}
			r.Status.Id = v.Id
			r.Status.Extension = v.Extension
		case "request":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProduct element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Request = append(r.Request, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProduct element", t)
			}
		case "quantity":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Quantity == nil {
				r.Quantity = &Integer{}
			}
			r.Quantity.Value = v.Value
		case "_quantity":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Quantity == nil {
				r.Quantity = &Integer{}
			}
			r.Quantity.Id = v.Id
			r.Quantity.Extension = v.Extension
		case "parent":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProduct element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Parent = append(r.Parent, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProduct element", t)
			}
		case "collection":
			var v BiologicallyDerivedProductCollection
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Collection = &v
		case "processing":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProduct element", t)
			}
			for d.More() {
				var v BiologicallyDerivedProductProcessing
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Processing = append(r.Processing, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProduct element", t)
			}
		case "manipulation":
			var v BiologicallyDerivedProductManipulation
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Manipulation = &v
		case "storage":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProduct element", t)
			}
			for d.More() {
				var v BiologicallyDerivedProductStorage
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Storage = append(r.Storage, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProduct element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in BiologicallyDerivedProduct", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in BiologicallyDerivedProduct element", t)
	}
	return nil
}
func (r *BiologicallyDerivedProductCollection) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in BiologicallyDerivedProductCollection element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in BiologicallyDerivedProductCollection element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProductCollection element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProductCollection element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProductCollection element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProductCollection element", t)
			}
		case "collector":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Collector = &v
		case "source":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Source = &v
		case "collectedDateTime":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Collected != nil {
				r.Collected = DateTime{
					Extension: r.Collected.(DateTime).Extension,
					Id:        r.Collected.(DateTime).Id,
					Value:     v.Value,
				}
			} else {
				r.Collected = v
			}
		case "_collectedDateTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Collected != nil {
				r.Collected = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Collected.(DateTime).Value,
				}
			} else {
				r.Collected = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "collectedPeriod":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Collected = v
		default:
			return fmt.Errorf("invalid field: %s in BiologicallyDerivedProductCollection", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in BiologicallyDerivedProductCollection element", t)
	}
	return nil
}
func (r *BiologicallyDerivedProductProcessing) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in BiologicallyDerivedProductProcessing element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in BiologicallyDerivedProductProcessing element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProductProcessing element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProductProcessing element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProductProcessing element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProductProcessing element", t)
			}
		case "description":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Value = v.Value
		case "_description":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Id = v.Id
			r.Description.Extension = v.Extension
		case "procedure":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Procedure = &v
		case "additive":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Additive = &v
		case "timeDateTime":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Time != nil {
				r.Time = DateTime{
					Extension: r.Time.(DateTime).Extension,
					Id:        r.Time.(DateTime).Id,
					Value:     v.Value,
				}
			} else {
				r.Time = v
			}
		case "_timeDateTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Time != nil {
				r.Time = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Time.(DateTime).Value,
				}
			} else {
				r.Time = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "timePeriod":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Time = v
		default:
			return fmt.Errorf("invalid field: %s in BiologicallyDerivedProductProcessing", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in BiologicallyDerivedProductProcessing element", t)
	}
	return nil
}
func (r *BiologicallyDerivedProductManipulation) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in BiologicallyDerivedProductManipulation element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in BiologicallyDerivedProductManipulation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProductManipulation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProductManipulation element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProductManipulation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProductManipulation element", t)
			}
		case "description":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Value = v.Value
		case "_description":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Id = v.Id
			r.Description.Extension = v.Extension
		case "timeDateTime":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Time != nil {
				r.Time = DateTime{
					Extension: r.Time.(DateTime).Extension,
					Id:        r.Time.(DateTime).Id,
					Value:     v.Value,
				}
			} else {
				r.Time = v
			}
		case "_timeDateTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Time != nil {
				r.Time = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Time.(DateTime).Value,
				}
			} else {
				r.Time = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "timePeriod":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Time = v
		default:
			return fmt.Errorf("invalid field: %s in BiologicallyDerivedProductManipulation", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in BiologicallyDerivedProductManipulation element", t)
	}
	return nil
}
func (r *BiologicallyDerivedProductStorage) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in BiologicallyDerivedProductStorage element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in BiologicallyDerivedProductStorage element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProductStorage element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProductStorage element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BiologicallyDerivedProductStorage element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BiologicallyDerivedProductStorage element", t)
			}
		case "description":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Value = v.Value
		case "_description":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Id = v.Id
			r.Description.Extension = v.Extension
		case "temperature":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Temperature == nil {
				r.Temperature = &Decimal{}
			}
			r.Temperature.Value = v.Value
		case "_temperature":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Temperature == nil {
				r.Temperature = &Decimal{}
			}
			r.Temperature.Id = v.Id
			r.Temperature.Extension = v.Extension
		case "scale":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Scale == nil {
				r.Scale = &Code{}
			}
			r.Scale.Value = v.Value
		case "_scale":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Scale == nil {
				r.Scale = &Code{}
			}
			r.Scale.Id = v.Id
			r.Scale.Extension = v.Extension
		case "duration":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Duration = &v
		default:
			return fmt.Errorf("invalid field: %s in BiologicallyDerivedProductStorage", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in BiologicallyDerivedProductStorage element", t)
	}
	return nil
}
func (r BiologicallyDerivedProduct) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "BiologicallyDerivedProduct"
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
	for _, c := range r.Contained {
		err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "contained"}})
		if err != nil {
			return err
		}
		err = e.EncodeElement(c, xml.StartElement{Name: xml.Name{Local: "__contained__"}})
		if err != nil {
			return err
		}
		err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "contained"}})
		if err != nil {
			return err
		}
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
	err = e.EncodeElement(r.ProductCategory, xml.StartElement{Name: xml.Name{Local: "productCategory"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProductCode, xml.StartElement{Name: xml.Name{Local: "productCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Request, xml.StartElement{Name: xml.Name{Local: "request"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Parent, xml.StartElement{Name: xml.Name{Local: "parent"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Collection, xml.StartElement{Name: xml.Name{Local: "collection"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Processing, xml.StartElement{Name: xml.Name{Local: "processing"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Manipulation, xml.StartElement{Name: xml.Name{Local: "manipulation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Storage, xml.StartElement{Name: xml.Name{Local: "storage"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r BiologicallyDerivedProductCollection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Collector, xml.StartElement{Name: xml.Name{Local: "collector"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Source, xml.StartElement{Name: xml.Name{Local: "source"}})
	if err != nil {
		return err
	}
	switch v := r.Collected.(type) {
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "collectedDateTime"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "collectedPeriod"}})
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
func (r BiologicallyDerivedProductProcessing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Procedure, xml.StartElement{Name: xml.Name{Local: "procedure"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Additive, xml.StartElement{Name: xml.Name{Local: "additive"}})
	if err != nil {
		return err
	}
	switch v := r.Time.(type) {
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "timeDateTime"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "timePeriod"}})
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
func (r BiologicallyDerivedProductManipulation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	switch v := r.Time.(type) {
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "timeDateTime"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "timePeriod"}})
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
func (r BiologicallyDerivedProductStorage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Temperature, xml.StartElement{Name: xml.Name{Local: "temperature"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Scale, xml.StartElement{Name: xml.Name{Local: "scale"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Duration, xml.StartElement{Name: xml.Name{Local: "duration"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *BiologicallyDerivedProduct) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "productCategory":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductCategory = &v
			case "productCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductCode = &v
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = &v
			case "request":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Request = append(r.Request, v)
			case "quantity":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = &v
			case "parent":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Parent = append(r.Parent, v)
			case "collection":
				var v BiologicallyDerivedProductCollection
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Collection = &v
			case "processing":
				var v BiologicallyDerivedProductProcessing
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Processing = append(r.Processing, v)
			case "manipulation":
				var v BiologicallyDerivedProductManipulation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Manipulation = &v
			case "storage":
				var v BiologicallyDerivedProductStorage
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Storage = append(r.Storage, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *BiologicallyDerivedProductCollection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "collector":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Collector = &v
			case "source":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Source = &v
			case "collectedDateTime":
				if r.Collected != nil {
					return fmt.Errorf("multiple values for field \"Collected\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Collected = v
			case "collectedPeriod":
				if r.Collected != nil {
					return fmt.Errorf("multiple values for field \"Collected\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Collected = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *BiologicallyDerivedProductProcessing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "procedure":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Procedure = &v
			case "additive":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Additive = &v
			case "timeDateTime":
				if r.Time != nil {
					return fmt.Errorf("multiple values for field \"Time\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Time = v
			case "timePeriod":
				if r.Time != nil {
					return fmt.Errorf("multiple values for field \"Time\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Time = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *BiologicallyDerivedProductManipulation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "timeDateTime":
				if r.Time != nil {
					return fmt.Errorf("multiple values for field \"Time\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Time = v
			case "timePeriod":
				if r.Time != nil {
					return fmt.Errorf("multiple values for field \"Time\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Time = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *BiologicallyDerivedProductStorage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "temperature":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Temperature = &v
			case "scale":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Scale = &v
			case "duration":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Duration = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r BiologicallyDerivedProduct) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, *r.Id)
		}
	}
	if len(name) == 0 || slices.Contains(name, "meta") {
		if r.Meta != nil {
			children = append(children, *r.Meta)
		}
	}
	if len(name) == 0 || slices.Contains(name, "implicitRules") {
		if r.ImplicitRules != nil {
			children = append(children, *r.ImplicitRules)
		}
	}
	if len(name) == 0 || slices.Contains(name, "language") {
		if r.Language != nil {
			children = append(children, *r.Language)
		}
	}
	if len(name) == 0 || slices.Contains(name, "text") {
		if r.Text != nil {
			children = append(children, *r.Text)
		}
	}
	if len(name) == 0 || slices.Contains(name, "contained") {
		for _, v := range r.Contained {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "identifier") {
		for _, v := range r.Identifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "productCategory") {
		if r.ProductCategory != nil {
			children = append(children, *r.ProductCategory)
		}
	}
	if len(name) == 0 || slices.Contains(name, "productCode") {
		if r.ProductCode != nil {
			children = append(children, *r.ProductCode)
		}
	}
	if len(name) == 0 || slices.Contains(name, "status") {
		if r.Status != nil {
			children = append(children, *r.Status)
		}
	}
	if len(name) == 0 || slices.Contains(name, "request") {
		for _, v := range r.Request {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "quantity") {
		if r.Quantity != nil {
			children = append(children, *r.Quantity)
		}
	}
	if len(name) == 0 || slices.Contains(name, "parent") {
		for _, v := range r.Parent {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "collection") {
		if r.Collection != nil {
			children = append(children, *r.Collection)
		}
	}
	if len(name) == 0 || slices.Contains(name, "processing") {
		for _, v := range r.Processing {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "manipulation") {
		if r.Manipulation != nil {
			children = append(children, *r.Manipulation)
		}
	}
	if len(name) == 0 || slices.Contains(name, "storage") {
		for _, v := range r.Storage {
			children = append(children, v)
		}
	}
	return children
}
func (r BiologicallyDerivedProduct) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProduct to Boolean")
}
func (r BiologicallyDerivedProduct) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProduct to String")
}
func (r BiologicallyDerivedProduct) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProduct to Integer")
}
func (r BiologicallyDerivedProduct) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProduct to Decimal")
}
func (r BiologicallyDerivedProduct) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProduct to Date")
}
func (r BiologicallyDerivedProduct) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProduct to Time")
}
func (r BiologicallyDerivedProduct) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProduct to DateTime")
}
func (r BiologicallyDerivedProduct) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProduct to Quantity")
}
func (r BiologicallyDerivedProduct) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o BiologicallyDerivedProduct
	switch other := other.(type) {
	case BiologicallyDerivedProduct:
		o = other
	case *BiologicallyDerivedProduct:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r BiologicallyDerivedProduct) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o BiologicallyDerivedProduct
	switch other := other.(type) {
	case BiologicallyDerivedProduct:
		o = other
	case *BiologicallyDerivedProduct:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r BiologicallyDerivedProduct) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.Id",
		}, {
			Name: "Meta",
			Type: "FHIR.Meta",
		}, {
			Name: "ImplicitRules",
			Type: "FHIR.Uri",
		}, {
			Name: "Language",
			Type: "FHIR.Code",
		}, {
			Name: "Text",
			Type: "FHIR.Narrative",
		}, {
			Name: "Contained",
			Type: "List<FHIR.>",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Identifier",
			Type: "List<FHIR.Identifier>",
		}, {
			Name: "ProductCategory",
			Type: "FHIR.Code",
		}, {
			Name: "ProductCode",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Status",
			Type: "FHIR.Code",
		}, {
			Name: "Request",
			Type: "List<FHIR.Reference>",
		}, {
			Name: "Quantity",
			Type: "FHIR.Integer",
		}, {
			Name: "Parent",
			Type: "List<FHIR.Reference>",
		}, {
			Name: "Collection",
			Type: "FHIR.BiologicallyDerivedProductCollection",
		}, {
			Name: "Processing",
			Type: "List<FHIR.BiologicallyDerivedProductProcessing>",
		}, {
			Name: "Manipulation",
			Type: "FHIR.BiologicallyDerivedProductManipulation",
		}, {
			Name: "Storage",
			Type: "List<FHIR.BiologicallyDerivedProductStorage>",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "BiologicallyDerivedProduct",
			Namespace: "FHIR",
		},
	}
}
func (r BiologicallyDerivedProductCollection) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "collector") {
		if r.Collector != nil {
			children = append(children, *r.Collector)
		}
	}
	if len(name) == 0 || slices.Contains(name, "source") {
		if r.Source != nil {
			children = append(children, *r.Source)
		}
	}
	if len(name) == 0 || slices.Contains(name, "collected") {
		if r.Collected != nil {
			children = append(children, r.Collected)
		}
	}
	return children
}
func (r BiologicallyDerivedProductCollection) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductCollection to Boolean")
}
func (r BiologicallyDerivedProductCollection) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductCollection to String")
}
func (r BiologicallyDerivedProductCollection) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductCollection to Integer")
}
func (r BiologicallyDerivedProductCollection) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductCollection to Decimal")
}
func (r BiologicallyDerivedProductCollection) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductCollection to Date")
}
func (r BiologicallyDerivedProductCollection) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductCollection to Time")
}
func (r BiologicallyDerivedProductCollection) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductCollection to DateTime")
}
func (r BiologicallyDerivedProductCollection) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductCollection to Quantity")
}
func (r BiologicallyDerivedProductCollection) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o BiologicallyDerivedProductCollection
	switch other := other.(type) {
	case BiologicallyDerivedProductCollection:
		o = other
	case *BiologicallyDerivedProductCollection:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r BiologicallyDerivedProductCollection) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o BiologicallyDerivedProductCollection
	switch other := other.(type) {
	case BiologicallyDerivedProductCollection:
		o = other
	case *BiologicallyDerivedProductCollection:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r BiologicallyDerivedProductCollection) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Collector",
			Type: "FHIR.Reference",
		}, {
			Name: "Source",
			Type: "FHIR.Reference",
		}, {
			Name: "Collected",
			Type: "FHIR.PrimitiveElement",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "BiologicallyDerivedProductCollection",
			Namespace: "FHIR",
		},
	}
}
func (r BiologicallyDerivedProductProcessing) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "description") {
		if r.Description != nil {
			children = append(children, *r.Description)
		}
	}
	if len(name) == 0 || slices.Contains(name, "procedure") {
		if r.Procedure != nil {
			children = append(children, *r.Procedure)
		}
	}
	if len(name) == 0 || slices.Contains(name, "additive") {
		if r.Additive != nil {
			children = append(children, *r.Additive)
		}
	}
	if len(name) == 0 || slices.Contains(name, "time") {
		if r.Time != nil {
			children = append(children, r.Time)
		}
	}
	return children
}
func (r BiologicallyDerivedProductProcessing) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductProcessing to Boolean")
}
func (r BiologicallyDerivedProductProcessing) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductProcessing to String")
}
func (r BiologicallyDerivedProductProcessing) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductProcessing to Integer")
}
func (r BiologicallyDerivedProductProcessing) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductProcessing to Decimal")
}
func (r BiologicallyDerivedProductProcessing) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductProcessing to Date")
}
func (r BiologicallyDerivedProductProcessing) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductProcessing to Time")
}
func (r BiologicallyDerivedProductProcessing) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductProcessing to DateTime")
}
func (r BiologicallyDerivedProductProcessing) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductProcessing to Quantity")
}
func (r BiologicallyDerivedProductProcessing) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o BiologicallyDerivedProductProcessing
	switch other := other.(type) {
	case BiologicallyDerivedProductProcessing:
		o = other
	case *BiologicallyDerivedProductProcessing:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r BiologicallyDerivedProductProcessing) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o BiologicallyDerivedProductProcessing
	switch other := other.(type) {
	case BiologicallyDerivedProductProcessing:
		o = other
	case *BiologicallyDerivedProductProcessing:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r BiologicallyDerivedProductProcessing) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Description",
			Type: "FHIR.String",
		}, {
			Name: "Procedure",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Additive",
			Type: "FHIR.Reference",
		}, {
			Name: "Time",
			Type: "FHIR.PrimitiveElement",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "BiologicallyDerivedProductProcessing",
			Namespace: "FHIR",
		},
	}
}
func (r BiologicallyDerivedProductManipulation) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "description") {
		if r.Description != nil {
			children = append(children, *r.Description)
		}
	}
	if len(name) == 0 || slices.Contains(name, "time") {
		if r.Time != nil {
			children = append(children, r.Time)
		}
	}
	return children
}
func (r BiologicallyDerivedProductManipulation) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductManipulation to Boolean")
}
func (r BiologicallyDerivedProductManipulation) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductManipulation to String")
}
func (r BiologicallyDerivedProductManipulation) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductManipulation to Integer")
}
func (r BiologicallyDerivedProductManipulation) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductManipulation to Decimal")
}
func (r BiologicallyDerivedProductManipulation) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductManipulation to Date")
}
func (r BiologicallyDerivedProductManipulation) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductManipulation to Time")
}
func (r BiologicallyDerivedProductManipulation) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductManipulation to DateTime")
}
func (r BiologicallyDerivedProductManipulation) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductManipulation to Quantity")
}
func (r BiologicallyDerivedProductManipulation) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o BiologicallyDerivedProductManipulation
	switch other := other.(type) {
	case BiologicallyDerivedProductManipulation:
		o = other
	case *BiologicallyDerivedProductManipulation:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r BiologicallyDerivedProductManipulation) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o BiologicallyDerivedProductManipulation
	switch other := other.(type) {
	case BiologicallyDerivedProductManipulation:
		o = other
	case *BiologicallyDerivedProductManipulation:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r BiologicallyDerivedProductManipulation) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Description",
			Type: "FHIR.String",
		}, {
			Name: "Time",
			Type: "FHIR.PrimitiveElement",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "BiologicallyDerivedProductManipulation",
			Namespace: "FHIR",
		},
	}
}
func (r BiologicallyDerivedProductStorage) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "description") {
		if r.Description != nil {
			children = append(children, *r.Description)
		}
	}
	if len(name) == 0 || slices.Contains(name, "temperature") {
		if r.Temperature != nil {
			children = append(children, *r.Temperature)
		}
	}
	if len(name) == 0 || slices.Contains(name, "scale") {
		if r.Scale != nil {
			children = append(children, *r.Scale)
		}
	}
	if len(name) == 0 || slices.Contains(name, "duration") {
		if r.Duration != nil {
			children = append(children, *r.Duration)
		}
	}
	return children
}
func (r BiologicallyDerivedProductStorage) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductStorage to Boolean")
}
func (r BiologicallyDerivedProductStorage) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductStorage to String")
}
func (r BiologicallyDerivedProductStorage) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductStorage to Integer")
}
func (r BiologicallyDerivedProductStorage) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductStorage to Decimal")
}
func (r BiologicallyDerivedProductStorage) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductStorage to Date")
}
func (r BiologicallyDerivedProductStorage) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductStorage to Time")
}
func (r BiologicallyDerivedProductStorage) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductStorage to DateTime")
}
func (r BiologicallyDerivedProductStorage) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert BiologicallyDerivedProductStorage to Quantity")
}
func (r BiologicallyDerivedProductStorage) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o BiologicallyDerivedProductStorage
	switch other := other.(type) {
	case BiologicallyDerivedProductStorage:
		o = other
	case *BiologicallyDerivedProductStorage:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r BiologicallyDerivedProductStorage) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o BiologicallyDerivedProductStorage
	switch other := other.(type) {
	case BiologicallyDerivedProductStorage:
		o = other
	case *BiologicallyDerivedProductStorage:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r BiologicallyDerivedProductStorage) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Description",
			Type: "FHIR.String",
		}, {
			Name: "Temperature",
			Type: "FHIR.Decimal",
		}, {
			Name: "Scale",
			Type: "FHIR.Code",
		}, {
			Name: "Duration",
			Type: "FHIR.Period",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "BiologicallyDerivedProductStorage",
			Namespace: "FHIR",
		},
	}
}
