package r4

import "encoding/json"

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
	Contained []Resource
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
	Contained                     []containedResource        `json:"contained,omitempty"`
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
	m.Type = r.Type
	m.Orderable = r.Orderable
	if r.Orderable.Id != nil || r.Orderable.Extension != nil {
		m.OrderablePrimitiveElement = &primitiveElement{Id: r.Orderable.Id, Extension: r.Orderable.Extension}
	}
	m.ReferencedItem = r.ReferencedItem
	m.AdditionalIdentifier = r.AdditionalIdentifier
	m.Classification = r.Classification
	m.Status = r.Status
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.ValidityPeriod = r.ValidityPeriod
	m.ValidTo = r.ValidTo
	if r.ValidTo != nil && (r.ValidTo.Id != nil || r.ValidTo.Extension != nil) {
		m.ValidToPrimitiveElement = &primitiveElement{Id: r.ValidTo.Id, Extension: r.ValidTo.Extension}
	}
	m.LastUpdated = r.LastUpdated
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
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.ValidityPeriod = m.ValidityPeriod
	r.ValidTo = m.ValidTo
	if m.ValidToPrimitiveElement != nil {
		r.ValidTo.Id = m.ValidToPrimitiveElement.Id
		r.ValidTo.Extension = m.ValidToPrimitiveElement.Extension
	}
	r.LastUpdated = m.LastUpdated
	if m.LastUpdatedPrimitiveElement != nil {
		r.LastUpdated.Id = m.LastUpdatedPrimitiveElement.Id
		r.LastUpdated.Extension = m.LastUpdatedPrimitiveElement.Extension
	}
	r.AdditionalCharacteristic = m.AdditionalCharacteristic
	r.AdditionalClassification = m.AdditionalClassification
	r.RelatedEntry = m.RelatedEntry
	return nil
}
func (r CatalogEntry) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
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
	m.Relationtype = r.Relationtype
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
func (r CatalogEntryRelatedEntry) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
