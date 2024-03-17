package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Batch numbering.
type MedicinalProductPackagedBatchIdentifier struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A number appearing on the outer packaging of a specific batch.
	OuterPackaging types.Identifier
	// A number appearing on the immediate packaging (and not the outer packaging).
	ImmediatePackaging *types.Identifier
}

func (s MedicinalProductPackagedBatchIdentifier) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A packaging item, as a contained for medicine, possibly with other packaging items within.
type MedicinalProductPackagedPackageItem struct {
	// The physical type of the container of the medicine.
	Type types.CodeableConcept
	// The manufactured item as contained in the packaged medicinal product.
	ManufacturedItem []types.Reference
	// Other codeable characteristics.
	OtherCharacteristics []types.CodeableConcept
	// The quantity of this package in the medicinal product, at the current level of packaging. The outermost is always 1.
	Quantity types.Quantity
	// A possible alternate material for the packaging.
	AlternateMaterial []types.CodeableConcept
	// Dimensions, color etc.
	PhysicalCharacteristics *types.ProdCharacteristic
	// Shelf Life and storage information.
	ShelfLifeStorage []types.ProductShelfLife
	// Manufacturer of this Package Item.
	Manufacturer []types.Reference
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Including possibly Data Carrier Identifier.
	Identifier []types.Identifier
	// Allows containers within containers.
	PackageItem []MedicinalProductPackagedPackageItem
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Material type of the package item.
	Material []types.CodeableConcept
	// A device accompanying a medicinal product.
	Device []types.Reference
}

func (s MedicinalProductPackagedPackageItem) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A medicinal product in a container or package.
type MedicinalProductPackaged struct {
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Unique identifier.
	Identifier []types.Identifier
	// Textual description.
	Description *types.String
	// The legal status of supply of the medicinal product as classified by the regulator.
	LegalStatusOfSupply *types.CodeableConcept
	// Marketing information.
	MarketingStatus []types.MarketingStatus
	// Manufacturer of this Package Item.
	Manufacturer []types.Reference
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// The base language in which the resource is written.
	Language *types.Code
	// The product with this is a pack for.
	Subject []types.Reference
	// Manufacturer of this Package Item.
	MarketingAuthorization *types.Reference
	// Batch numbering.
	BatchIdentifier []MedicinalProductPackagedBatchIdentifier
	// A packaging item, as a contained for medicine, possibly with other packaging items within.
	PackageItem []MedicinalProductPackagedPackageItem
}

func (s MedicinalProductPackaged) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
