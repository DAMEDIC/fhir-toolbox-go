package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
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

type jsonBiologicallyDerivedProduct struct {
	ResourceType                    string                                  `json:"resourceType"`
	Id                              *Id                                     `json:"id,omitempty"`
	IdPrimitiveElement              *primitiveElement                       `json:"_id,omitempty"`
	Meta                            *Meta                                   `json:"meta,omitempty"`
	ImplicitRules                   *Uri                                    `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement   *primitiveElement                       `json:"_implicitRules,omitempty"`
	Language                        *Code                                   `json:"language,omitempty"`
	LanguagePrimitiveElement        *primitiveElement                       `json:"_language,omitempty"`
	Text                            *Narrative                              `json:"text,omitempty"`
	Contained                       []ContainedResource                     `json:"contained,omitempty"`
	Extension                       []Extension                             `json:"extension,omitempty"`
	ModifierExtension               []Extension                             `json:"modifierExtension,omitempty"`
	Identifier                      []Identifier                            `json:"identifier,omitempty"`
	ProductCategory                 *Code                                   `json:"productCategory,omitempty"`
	ProductCategoryPrimitiveElement *primitiveElement                       `json:"_productCategory,omitempty"`
	ProductCode                     *CodeableConcept                        `json:"productCode,omitempty"`
	Status                          *Code                                   `json:"status,omitempty"`
	StatusPrimitiveElement          *primitiveElement                       `json:"_status,omitempty"`
	Request                         []Reference                             `json:"request,omitempty"`
	Quantity                        *Integer                                `json:"quantity,omitempty"`
	QuantityPrimitiveElement        *primitiveElement                       `json:"_quantity,omitempty"`
	Parent                          []Reference                             `json:"parent,omitempty"`
	Collection                      *BiologicallyDerivedProductCollection   `json:"collection,omitempty"`
	Processing                      []BiologicallyDerivedProductProcessing  `json:"processing,omitempty"`
	Manipulation                    *BiologicallyDerivedProductManipulation `json:"manipulation,omitempty"`
	Storage                         []BiologicallyDerivedProductStorage     `json:"storage,omitempty"`
}

func (r BiologicallyDerivedProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r BiologicallyDerivedProduct) marshalJSON() jsonBiologicallyDerivedProduct {
	m := jsonBiologicallyDerivedProduct{}
	m.ResourceType = "BiologicallyDerivedProduct"
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
	if r.ProductCategory != nil && r.ProductCategory.Value != nil {
		m.ProductCategory = r.ProductCategory
	}
	if r.ProductCategory != nil && (r.ProductCategory.Id != nil || r.ProductCategory.Extension != nil) {
		m.ProductCategoryPrimitiveElement = &primitiveElement{Id: r.ProductCategory.Id, Extension: r.ProductCategory.Extension}
	}
	m.ProductCode = r.ProductCode
	if r.Status != nil && r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Request = r.Request
	if r.Quantity != nil && r.Quantity.Value != nil {
		m.Quantity = r.Quantity
	}
	if r.Quantity != nil && (r.Quantity.Id != nil || r.Quantity.Extension != nil) {
		m.QuantityPrimitiveElement = &primitiveElement{Id: r.Quantity.Id, Extension: r.Quantity.Extension}
	}
	m.Parent = r.Parent
	m.Collection = r.Collection
	m.Processing = r.Processing
	m.Manipulation = r.Manipulation
	m.Storage = r.Storage
	return m
}
func (r *BiologicallyDerivedProduct) UnmarshalJSON(b []byte) error {
	var m jsonBiologicallyDerivedProduct
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *BiologicallyDerivedProduct) unmarshalJSON(m jsonBiologicallyDerivedProduct) error {
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
	r.ProductCategory = m.ProductCategory
	if m.ProductCategoryPrimitiveElement != nil {
		if r.ProductCategory == nil {
			r.ProductCategory = &Code{}
		}
		r.ProductCategory.Id = m.ProductCategoryPrimitiveElement.Id
		r.ProductCategory.Extension = m.ProductCategoryPrimitiveElement.Extension
	}
	r.ProductCode = m.ProductCode
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		if r.Status == nil {
			r.Status = &Code{}
		}
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Request = m.Request
	r.Quantity = m.Quantity
	if m.QuantityPrimitiveElement != nil {
		if r.Quantity == nil {
			r.Quantity = &Integer{}
		}
		r.Quantity.Id = m.QuantityPrimitiveElement.Id
		r.Quantity.Extension = m.QuantityPrimitiveElement.Extension
	}
	r.Parent = m.Parent
	r.Collection = m.Collection
	r.Processing = m.Processing
	r.Manipulation = m.Manipulation
	r.Storage = m.Storage
	return nil
}
func (r BiologicallyDerivedProduct) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
func (r BiologicallyDerivedProduct) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
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
	isBiologicallyDerivedProductCollectionCollected()
}

func (r DateTime) isBiologicallyDerivedProductCollectionCollected() {}
func (r Period) isBiologicallyDerivedProductCollectionCollected()   {}

type jsonBiologicallyDerivedProductCollection struct {
	Id                                *string           `json:"id,omitempty"`
	Extension                         []Extension       `json:"extension,omitempty"`
	ModifierExtension                 []Extension       `json:"modifierExtension,omitempty"`
	Collector                         *Reference        `json:"collector,omitempty"`
	Source                            *Reference        `json:"source,omitempty"`
	CollectedDateTime                 *DateTime         `json:"collectedDateTime,omitempty"`
	CollectedDateTimePrimitiveElement *primitiveElement `json:"_collectedDateTime,omitempty"`
	CollectedPeriod                   *Period           `json:"collectedPeriod,omitempty"`
}

func (r BiologicallyDerivedProductCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r BiologicallyDerivedProductCollection) marshalJSON() jsonBiologicallyDerivedProductCollection {
	m := jsonBiologicallyDerivedProductCollection{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Collector = r.Collector
	m.Source = r.Source
	switch v := r.Collected.(type) {
	case DateTime:
		if v.Value != nil {
			m.CollectedDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.CollectedDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.CollectedDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.CollectedDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.CollectedPeriod = &v
	case *Period:
		m.CollectedPeriod = v
	}
	return m
}
func (r *BiologicallyDerivedProductCollection) UnmarshalJSON(b []byte) error {
	var m jsonBiologicallyDerivedProductCollection
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *BiologicallyDerivedProductCollection) unmarshalJSON(m jsonBiologicallyDerivedProductCollection) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Collector = m.Collector
	r.Source = m.Source
	if m.CollectedDateTime != nil || m.CollectedDateTimePrimitiveElement != nil {
		if r.Collected != nil {
			return fmt.Errorf("multiple values for field \"Collected\"")
		}
		v := m.CollectedDateTime
		if m.CollectedDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.CollectedDateTimePrimitiveElement.Id
			v.Extension = m.CollectedDateTimePrimitiveElement.Extension
		}
		r.Collected = v
	}
	if m.CollectedPeriod != nil {
		if r.Collected != nil {
			return fmt.Errorf("multiple values for field \"Collected\"")
		}
		v := m.CollectedPeriod
		r.Collected = v
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
func (r BiologicallyDerivedProductCollection) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

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
	isBiologicallyDerivedProductProcessingTime()
}

func (r DateTime) isBiologicallyDerivedProductProcessingTime() {}
func (r Period) isBiologicallyDerivedProductProcessingTime()   {}

type jsonBiologicallyDerivedProductProcessing struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Description                  *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement  *primitiveElement `json:"_description,omitempty"`
	Procedure                    *CodeableConcept  `json:"procedure,omitempty"`
	Additive                     *Reference        `json:"additive,omitempty"`
	TimeDateTime                 *DateTime         `json:"timeDateTime,omitempty"`
	TimeDateTimePrimitiveElement *primitiveElement `json:"_timeDateTime,omitempty"`
	TimePeriod                   *Period           `json:"timePeriod,omitempty"`
}

func (r BiologicallyDerivedProductProcessing) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r BiologicallyDerivedProductProcessing) marshalJSON() jsonBiologicallyDerivedProductProcessing {
	m := jsonBiologicallyDerivedProductProcessing{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Procedure = r.Procedure
	m.Additive = r.Additive
	switch v := r.Time.(type) {
	case DateTime:
		if v.Value != nil {
			m.TimeDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.TimeDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.TimeDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.TimeDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.TimePeriod = &v
	case *Period:
		m.TimePeriod = v
	}
	return m
}
func (r *BiologicallyDerivedProductProcessing) UnmarshalJSON(b []byte) error {
	var m jsonBiologicallyDerivedProductProcessing
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *BiologicallyDerivedProductProcessing) unmarshalJSON(m jsonBiologicallyDerivedProductProcessing) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Procedure = m.Procedure
	r.Additive = m.Additive
	if m.TimeDateTime != nil || m.TimeDateTimePrimitiveElement != nil {
		if r.Time != nil {
			return fmt.Errorf("multiple values for field \"Time\"")
		}
		v := m.TimeDateTime
		if m.TimeDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.TimeDateTimePrimitiveElement.Id
			v.Extension = m.TimeDateTimePrimitiveElement.Extension
		}
		r.Time = v
	}
	if m.TimePeriod != nil {
		if r.Time != nil {
			return fmt.Errorf("multiple values for field \"Time\"")
		}
		v := m.TimePeriod
		r.Time = v
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
func (r BiologicallyDerivedProductProcessing) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

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
	isBiologicallyDerivedProductManipulationTime()
}

func (r DateTime) isBiologicallyDerivedProductManipulationTime() {}
func (r Period) isBiologicallyDerivedProductManipulationTime()   {}

type jsonBiologicallyDerivedProductManipulation struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Description                  *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement  *primitiveElement `json:"_description,omitempty"`
	TimeDateTime                 *DateTime         `json:"timeDateTime,omitempty"`
	TimeDateTimePrimitiveElement *primitiveElement `json:"_timeDateTime,omitempty"`
	TimePeriod                   *Period           `json:"timePeriod,omitempty"`
}

func (r BiologicallyDerivedProductManipulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r BiologicallyDerivedProductManipulation) marshalJSON() jsonBiologicallyDerivedProductManipulation {
	m := jsonBiologicallyDerivedProductManipulation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	switch v := r.Time.(type) {
	case DateTime:
		if v.Value != nil {
			m.TimeDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.TimeDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.TimeDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.TimeDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.TimePeriod = &v
	case *Period:
		m.TimePeriod = v
	}
	return m
}
func (r *BiologicallyDerivedProductManipulation) UnmarshalJSON(b []byte) error {
	var m jsonBiologicallyDerivedProductManipulation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *BiologicallyDerivedProductManipulation) unmarshalJSON(m jsonBiologicallyDerivedProductManipulation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	if m.TimeDateTime != nil || m.TimeDateTimePrimitiveElement != nil {
		if r.Time != nil {
			return fmt.Errorf("multiple values for field \"Time\"")
		}
		v := m.TimeDateTime
		if m.TimeDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.TimeDateTimePrimitiveElement.Id
			v.Extension = m.TimeDateTimePrimitiveElement.Extension
		}
		r.Time = v
	}
	if m.TimePeriod != nil {
		if r.Time != nil {
			return fmt.Errorf("multiple values for field \"Time\"")
		}
		v := m.TimePeriod
		r.Time = v
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
func (r BiologicallyDerivedProductManipulation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

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
type jsonBiologicallyDerivedProductStorage struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Description                 *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
	Temperature                 *Decimal          `json:"temperature,omitempty"`
	TemperaturePrimitiveElement *primitiveElement `json:"_temperature,omitempty"`
	Scale                       *Code             `json:"scale,omitempty"`
	ScalePrimitiveElement       *primitiveElement `json:"_scale,omitempty"`
	Duration                    *Period           `json:"duration,omitempty"`
}

func (r BiologicallyDerivedProductStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r BiologicallyDerivedProductStorage) marshalJSON() jsonBiologicallyDerivedProductStorage {
	m := jsonBiologicallyDerivedProductStorage{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	if r.Temperature != nil && r.Temperature.Value != nil {
		m.Temperature = r.Temperature
	}
	if r.Temperature != nil && (r.Temperature.Id != nil || r.Temperature.Extension != nil) {
		m.TemperaturePrimitiveElement = &primitiveElement{Id: r.Temperature.Id, Extension: r.Temperature.Extension}
	}
	if r.Scale != nil && r.Scale.Value != nil {
		m.Scale = r.Scale
	}
	if r.Scale != nil && (r.Scale.Id != nil || r.Scale.Extension != nil) {
		m.ScalePrimitiveElement = &primitiveElement{Id: r.Scale.Id, Extension: r.Scale.Extension}
	}
	m.Duration = r.Duration
	return m
}
func (r *BiologicallyDerivedProductStorage) UnmarshalJSON(b []byte) error {
	var m jsonBiologicallyDerivedProductStorage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *BiologicallyDerivedProductStorage) unmarshalJSON(m jsonBiologicallyDerivedProductStorage) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Temperature = m.Temperature
	if m.TemperaturePrimitiveElement != nil {
		if r.Temperature == nil {
			r.Temperature = &Decimal{}
		}
		r.Temperature.Id = m.TemperaturePrimitiveElement.Id
		r.Temperature.Extension = m.TemperaturePrimitiveElement.Extension
	}
	r.Scale = m.Scale
	if m.ScalePrimitiveElement != nil {
		if r.Scale == nil {
			r.Scale = &Code{}
		}
		r.Scale.Id = m.ScalePrimitiveElement.Id
		r.Scale.Extension = m.ScalePrimitiveElement.Extension
	}
	r.Duration = m.Duration
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
func (r BiologicallyDerivedProductStorage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
