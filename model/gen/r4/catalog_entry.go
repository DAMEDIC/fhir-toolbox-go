package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// Catalog entries are wrappers that contextualize items included in a catalog.
type CatalogEntry struct {
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
	// Used in supporting different identifiers for the same product, e.g. manufacturer code and retailer code.
	Identifier []Identifier
	// The type of item - medication, device, service, protocol or other.
	Type *CodeableConcept
	// Whether the entry represents an orderable item.
	Orderable Boolean
	// The item in a catalog or definition.
	ReferencedItem Reference
	// Used in supporting related concepts, e.g. NDC to RxNorm.
	AdditionalIdentifier []Identifier
	// Classes of devices, or ATC for medication.
	Classification []CodeableConcept
	// Used to support catalog exchange even for unsupported products, e.g. getting list of medications even if not prescribable.
	Status *Code
	// The time period in which this catalog entry is expected to be active.
	ValidityPeriod *Period
	// The date until which this catalog entry is expected to be active.
	ValidTo *DateTime
	// Typically date of issue is different from the beginning of the validity. This can be used to see when an item was last updated.
	LastUpdated *DateTime
	// Used for examplefor Out of Formulary, or any specifics.
	AdditionalCharacteristic []CodeableConcept
	// User for example for ATC classification, or.
	AdditionalClassification []CodeableConcept
	// Used for example, to point to a substance, or to a device used to administer a medication.
	RelatedEntry []CatalogEntryRelatedEntry
}

func (r CatalogEntry) ResourceType() string {
	return "CatalogEntry"
}
func (r CatalogEntry) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonCatalogEntry struct {
	ResourceType                  string                     `json:"resourceType"`
	Id                            *Id                        `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement          `json:"_id,omitempty"`
	Meta                          *Meta                      `json:"meta,omitempty"`
	ImplicitRules                 *Uri                       `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement          `json:"_implicitRules,omitempty"`
	Language                      *Code                      `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement          `json:"_language,omitempty"`
	Text                          *Narrative                 `json:"text,omitempty"`
	Contained                     []ContainedResource        `json:"contained,omitempty"`
	Extension                     []Extension                `json:"extension,omitempty"`
	ModifierExtension             []Extension                `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier               `json:"identifier,omitempty"`
	Type                          *CodeableConcept           `json:"type,omitempty"`
	Orderable                     Boolean                    `json:"orderable,omitempty"`
	OrderablePrimitiveElement     *primitiveElement          `json:"_orderable,omitempty"`
	ReferencedItem                Reference                  `json:"referencedItem,omitempty"`
	AdditionalIdentifier          []Identifier               `json:"additionalIdentifier,omitempty"`
	Classification                []CodeableConcept          `json:"classification,omitempty"`
	Status                        *Code                      `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement          `json:"_status,omitempty"`
	ValidityPeriod                *Period                    `json:"validityPeriod,omitempty"`
	ValidTo                       *DateTime                  `json:"validTo,omitempty"`
	ValidToPrimitiveElement       *primitiveElement          `json:"_validTo,omitempty"`
	LastUpdated                   *DateTime                  `json:"lastUpdated,omitempty"`
	LastUpdatedPrimitiveElement   *primitiveElement          `json:"_lastUpdated,omitempty"`
	AdditionalCharacteristic      []CodeableConcept          `json:"additionalCharacteristic,omitempty"`
	AdditionalClassification      []CodeableConcept          `json:"additionalClassification,omitempty"`
	RelatedEntry                  []CatalogEntryRelatedEntry `json:"relatedEntry,omitempty"`
}

func (r CatalogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CatalogEntry) marshalJSON() jsonCatalogEntry {
	m := jsonCatalogEntry{}
	m.ResourceType = "CatalogEntry"
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
	m.Type = r.Type
	if r.Orderable.Value != nil {
		m.Orderable = r.Orderable
	}
	if r.Orderable.Id != nil || r.Orderable.Extension != nil {
		m.OrderablePrimitiveElement = &primitiveElement{Id: r.Orderable.Id, Extension: r.Orderable.Extension}
	}
	m.ReferencedItem = r.ReferencedItem
	m.AdditionalIdentifier = r.AdditionalIdentifier
	m.Classification = r.Classification
	if r.Status != nil && r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.ValidityPeriod = r.ValidityPeriod
	if r.ValidTo != nil && r.ValidTo.Value != nil {
		m.ValidTo = r.ValidTo
	}
	if r.ValidTo != nil && (r.ValidTo.Id != nil || r.ValidTo.Extension != nil) {
		m.ValidToPrimitiveElement = &primitiveElement{Id: r.ValidTo.Id, Extension: r.ValidTo.Extension}
	}
	if r.LastUpdated != nil && r.LastUpdated.Value != nil {
		m.LastUpdated = r.LastUpdated
	}
	if r.LastUpdated != nil && (r.LastUpdated.Id != nil || r.LastUpdated.Extension != nil) {
		m.LastUpdatedPrimitiveElement = &primitiveElement{Id: r.LastUpdated.Id, Extension: r.LastUpdated.Extension}
	}
	m.AdditionalCharacteristic = r.AdditionalCharacteristic
	m.AdditionalClassification = r.AdditionalClassification
	m.RelatedEntry = r.RelatedEntry
	return m
}
func (r *CatalogEntry) UnmarshalJSON(b []byte) error {
	var m jsonCatalogEntry
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CatalogEntry) unmarshalJSON(m jsonCatalogEntry) error {
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
	r.Type = m.Type
	r.Orderable = m.Orderable
	if m.OrderablePrimitiveElement != nil {
		r.Orderable.Id = m.OrderablePrimitiveElement.Id
		r.Orderable.Extension = m.OrderablePrimitiveElement.Extension
	}
	r.ReferencedItem = m.ReferencedItem
	r.AdditionalIdentifier = m.AdditionalIdentifier
	r.Classification = m.Classification
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		if r.Status == nil {
			r.Status = &Code{}
		}
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.ValidityPeriod = m.ValidityPeriod
	r.ValidTo = m.ValidTo
	if m.ValidToPrimitiveElement != nil {
		if r.ValidTo == nil {
			r.ValidTo = &DateTime{}
		}
		r.ValidTo.Id = m.ValidToPrimitiveElement.Id
		r.ValidTo.Extension = m.ValidToPrimitiveElement.Extension
	}
	r.LastUpdated = m.LastUpdated
	if m.LastUpdatedPrimitiveElement != nil {
		if r.LastUpdated == nil {
			r.LastUpdated = &DateTime{}
		}
		r.LastUpdated.Id = m.LastUpdatedPrimitiveElement.Id
		r.LastUpdated.Extension = m.LastUpdatedPrimitiveElement.Extension
	}
	r.AdditionalCharacteristic = m.AdditionalCharacteristic
	r.AdditionalClassification = m.AdditionalClassification
	r.RelatedEntry = m.RelatedEntry
	return nil
}
func (r CatalogEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "CatalogEntry"
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Orderable, xml.StartElement{Name: xml.Name{Local: "orderable"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReferencedItem, xml.StartElement{Name: xml.Name{Local: "referencedItem"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AdditionalIdentifier, xml.StartElement{Name: xml.Name{Local: "additionalIdentifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Classification, xml.StartElement{Name: xml.Name{Local: "classification"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ValidityPeriod, xml.StartElement{Name: xml.Name{Local: "validityPeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ValidTo, xml.StartElement{Name: xml.Name{Local: "validTo"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LastUpdated, xml.StartElement{Name: xml.Name{Local: "lastUpdated"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AdditionalCharacteristic, xml.StartElement{Name: xml.Name{Local: "additionalCharacteristic"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AdditionalClassification, xml.StartElement{Name: xml.Name{Local: "additionalClassification"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RelatedEntry, xml.StartElement{Name: xml.Name{Local: "relatedEntry"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CatalogEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "orderable":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Orderable = v
			case "referencedItem":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReferencedItem = v
			case "additionalIdentifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AdditionalIdentifier = append(r.AdditionalIdentifier, v)
			case "classification":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Classification = append(r.Classification, v)
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = &v
			case "validityPeriod":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ValidityPeriod = &v
			case "validTo":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ValidTo = &v
			case "lastUpdated":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LastUpdated = &v
			case "additionalCharacteristic":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AdditionalCharacteristic = append(r.AdditionalCharacteristic, v)
			case "additionalClassification":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AdditionalClassification = append(r.AdditionalClassification, v)
			case "relatedEntry":
				var v CatalogEntryRelatedEntry
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RelatedEntry = append(r.RelatedEntry, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CatalogEntry) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Used for example, to point to a substance, or to a device used to administer a medication.
type CatalogEntryRelatedEntry struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of relation to the related item: child, parent, packageContent, containerPackage, usedIn, uses, requires, etc.
	Relationtype Code
	// The reference to the related item.
	Item Reference
}
type jsonCatalogEntryRelatedEntry struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Relationtype                 Code              `json:"relationtype,omitempty"`
	RelationtypePrimitiveElement *primitiveElement `json:"_relationtype,omitempty"`
	Item                         Reference         `json:"item,omitempty"`
}

func (r CatalogEntryRelatedEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r CatalogEntryRelatedEntry) marshalJSON() jsonCatalogEntryRelatedEntry {
	m := jsonCatalogEntryRelatedEntry{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Relationtype.Value != nil {
		m.Relationtype = r.Relationtype
	}
	if r.Relationtype.Id != nil || r.Relationtype.Extension != nil {
		m.RelationtypePrimitiveElement = &primitiveElement{Id: r.Relationtype.Id, Extension: r.Relationtype.Extension}
	}
	m.Item = r.Item
	return m
}
func (r *CatalogEntryRelatedEntry) UnmarshalJSON(b []byte) error {
	var m jsonCatalogEntryRelatedEntry
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *CatalogEntryRelatedEntry) unmarshalJSON(m jsonCatalogEntryRelatedEntry) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Relationtype = m.Relationtype
	if m.RelationtypePrimitiveElement != nil {
		r.Relationtype.Id = m.RelationtypePrimitiveElement.Id
		r.Relationtype.Extension = m.RelationtypePrimitiveElement.Extension
	}
	r.Item = m.Item
	return nil
}
func (r CatalogEntryRelatedEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Relationtype, xml.StartElement{Name: xml.Name{Local: "relationtype"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Item, xml.StartElement{Name: xml.Name{Local: "item"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *CatalogEntryRelatedEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "relationtype":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Relationtype = v
			case "item":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Item = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r CatalogEntryRelatedEntry) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
