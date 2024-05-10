package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A medicinal product in a container or package.
type MedicinalProductPackaged struct {
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
	// Unique identifier.
	Identifier []Identifier
	// The product with this is a pack for.
	Subject []Reference
	// Textual description.
	Description *String
	// The legal status of supply of the medicinal product as classified by the regulator.
	LegalStatusOfSupply *CodeableConcept
	// Marketing information.
	MarketingStatus []MarketingStatus
	// Manufacturer of this Package Item.
	MarketingAuthorization *Reference
	// Manufacturer of this Package Item.
	Manufacturer []Reference
	// Batch numbering.
	BatchIdentifier []MedicinalProductPackagedBatchIdentifier
	// A packaging item, as a contained for medicine, possibly with other packaging items within.
	PackageItem []MedicinalProductPackagedPackageItem
}
type jsonMedicinalProductPackaged struct {
	ResourceType                  string                                    `json:"resourceType"`
	Id                            *Id                                       `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement                         `json:"_id,omitempty"`
	Meta                          *Meta                                     `json:"meta,omitempty"`
	ImplicitRules                 *Uri                                      `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement                         `json:"_implicitRules,omitempty"`
	Language                      *Code                                     `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement                         `json:"_language,omitempty"`
	Text                          *Narrative                                `json:"text,omitempty"`
	Contained                     []containedResource                       `json:"contained,omitempty"`
	Extension                     []Extension                               `json:"extension,omitempty"`
	ModifierExtension             []Extension                               `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier                              `json:"identifier,omitempty"`
	Subject                       []Reference                               `json:"subject,omitempty"`
	Description                   *String                                   `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement                         `json:"_description,omitempty"`
	LegalStatusOfSupply           *CodeableConcept                          `json:"legalStatusOfSupply,omitempty"`
	MarketingStatus               []MarketingStatus                         `json:"marketingStatus,omitempty"`
	MarketingAuthorization        *Reference                                `json:"marketingAuthorization,omitempty"`
	Manufacturer                  []Reference                               `json:"manufacturer,omitempty"`
	BatchIdentifier               []MedicinalProductPackagedBatchIdentifier `json:"batchIdentifier,omitempty"`
	PackageItem                   []MedicinalProductPackagedPackageItem     `json:"packageItem,omitempty"`
}

func (r MedicinalProductPackaged) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductPackaged) marshalJSON() jsonMedicinalProductPackaged {
	m := jsonMedicinalProductPackaged{}
	m.ResourceType = "MedicinalProductPackaged"
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
	m.Subject = r.Subject
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.LegalStatusOfSupply = r.LegalStatusOfSupply
	m.MarketingStatus = r.MarketingStatus
	m.MarketingAuthorization = r.MarketingAuthorization
	m.Manufacturer = r.Manufacturer
	m.BatchIdentifier = r.BatchIdentifier
	m.PackageItem = r.PackageItem
	return m
}
func (r *MedicinalProductPackaged) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductPackaged
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductPackaged) unmarshalJSON(m jsonMedicinalProductPackaged) error {
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
	r.Subject = m.Subject
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.LegalStatusOfSupply = m.LegalStatusOfSupply
	r.MarketingStatus = m.MarketingStatus
	r.MarketingAuthorization = m.MarketingAuthorization
	r.Manufacturer = m.Manufacturer
	r.BatchIdentifier = m.BatchIdentifier
	r.PackageItem = m.PackageItem
	return nil
}
func (r MedicinalProductPackaged) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Batch numbering.
type MedicinalProductPackagedBatchIdentifier struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number appearing on the outer packaging of a specific batch.
	OuterPackaging Identifier
	// A number appearing on the immediate packaging (and not the outer packaging).
	ImmediatePackaging *Identifier
}
type jsonMedicinalProductPackagedBatchIdentifier struct {
	Id                 *string     `json:"id,omitempty"`
	Extension          []Extension `json:"extension,omitempty"`
	ModifierExtension  []Extension `json:"modifierExtension,omitempty"`
	OuterPackaging     Identifier  `json:"outerPackaging,omitempty"`
	ImmediatePackaging *Identifier `json:"immediatePackaging,omitempty"`
}

func (r MedicinalProductPackagedBatchIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductPackagedBatchIdentifier) marshalJSON() jsonMedicinalProductPackagedBatchIdentifier {
	m := jsonMedicinalProductPackagedBatchIdentifier{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.OuterPackaging = r.OuterPackaging
	m.ImmediatePackaging = r.ImmediatePackaging
	return m
}
func (r *MedicinalProductPackagedBatchIdentifier) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductPackagedBatchIdentifier
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductPackagedBatchIdentifier) unmarshalJSON(m jsonMedicinalProductPackagedBatchIdentifier) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.OuterPackaging = m.OuterPackaging
	r.ImmediatePackaging = m.ImmediatePackaging
	return nil
}
func (r MedicinalProductPackagedBatchIdentifier) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A packaging item, as a contained for medicine, possibly with other packaging items within.
type MedicinalProductPackagedPackageItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Including possibly Data Carrier Identifier.
	Identifier []Identifier
	// The physical type of the container of the medicine.
	Type CodeableConcept
	// The quantity of this package in the medicinal product, at the current level of packaging. The outermost is always 1.
	Quantity Quantity
	// Material type of the package item.
	Material []CodeableConcept
	// A possible alternate material for the packaging.
	AlternateMaterial []CodeableConcept
	// A device accompanying a medicinal product.
	Device []Reference
	// The manufactured item as contained in the packaged medicinal product.
	ManufacturedItem []Reference
	// Allows containers within containers.
	PackageItem []MedicinalProductPackagedPackageItem
	// Dimensions, color etc.
	PhysicalCharacteristics *ProdCharacteristic
	// Other codeable characteristics.
	OtherCharacteristics []CodeableConcept
	// Shelf Life and storage information.
	ShelfLifeStorage []ProductShelfLife
	// Manufacturer of this Package Item.
	Manufacturer []Reference
}
type jsonMedicinalProductPackagedPackageItem struct {
	Id                      *string                               `json:"id,omitempty"`
	Extension               []Extension                           `json:"extension,omitempty"`
	ModifierExtension       []Extension                           `json:"modifierExtension,omitempty"`
	Identifier              []Identifier                          `json:"identifier,omitempty"`
	Type                    CodeableConcept                       `json:"type,omitempty"`
	Quantity                Quantity                              `json:"quantity,omitempty"`
	Material                []CodeableConcept                     `json:"material,omitempty"`
	AlternateMaterial       []CodeableConcept                     `json:"alternateMaterial,omitempty"`
	Device                  []Reference                           `json:"device,omitempty"`
	ManufacturedItem        []Reference                           `json:"manufacturedItem,omitempty"`
	PackageItem             []MedicinalProductPackagedPackageItem `json:"packageItem,omitempty"`
	PhysicalCharacteristics *ProdCharacteristic                   `json:"physicalCharacteristics,omitempty"`
	OtherCharacteristics    []CodeableConcept                     `json:"otherCharacteristics,omitempty"`
	ShelfLifeStorage        []ProductShelfLife                    `json:"shelfLifeStorage,omitempty"`
	Manufacturer            []Reference                           `json:"manufacturer,omitempty"`
}

func (r MedicinalProductPackagedPackageItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MedicinalProductPackagedPackageItem) marshalJSON() jsonMedicinalProductPackagedPackageItem {
	m := jsonMedicinalProductPackagedPackageItem{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Type = r.Type
	m.Quantity = r.Quantity
	m.Material = r.Material
	m.AlternateMaterial = r.AlternateMaterial
	m.Device = r.Device
	m.ManufacturedItem = r.ManufacturedItem
	m.PackageItem = r.PackageItem
	m.PhysicalCharacteristics = r.PhysicalCharacteristics
	m.OtherCharacteristics = r.OtherCharacteristics
	m.ShelfLifeStorage = r.ShelfLifeStorage
	m.Manufacturer = r.Manufacturer
	return m
}
func (r *MedicinalProductPackagedPackageItem) UnmarshalJSON(b []byte) error {
	var m jsonMedicinalProductPackagedPackageItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MedicinalProductPackagedPackageItem) unmarshalJSON(m jsonMedicinalProductPackagedPackageItem) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Type = m.Type
	r.Quantity = m.Quantity
	r.Material = m.Material
	r.AlternateMaterial = m.AlternateMaterial
	r.Device = m.Device
	r.ManufacturedItem = m.ManufacturedItem
	r.PackageItem = m.PackageItem
	r.PhysicalCharacteristics = m.PhysicalCharacteristics
	r.OtherCharacteristics = m.OtherCharacteristics
	r.ShelfLifeStorage = m.ShelfLifeStorage
	r.Manufacturer = m.Manufacturer
	return nil
}
func (r MedicinalProductPackagedPackageItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
