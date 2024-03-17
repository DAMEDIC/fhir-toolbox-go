package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Used for example, to point to a substance, or to a device used to administer a medication.
type CatalogEntryRelatedEntry struct {
	// The type of relation to the related item: child, parent, packageContent, containerPackage, usedIn, uses, requires, etc.
	Relationtype types.Code
	// The reference to the related item.
	Item types.Reference
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
}

func (s CatalogEntryRelatedEntry) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Catalog entries are wrappers that contextualize items included in a catalog.
type CatalogEntry struct {
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// The base language in which the resource is written.
	Language *types.Code
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Classes of devices, or ATC for medication.
	Classification []types.CodeableConcept
	// The date until which this catalog entry is expected to be active.
	ValidTo *types.DateTime
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// The item in a catalog or definition.
	ReferencedItem types.Reference
	// The time period in which this catalog entry is expected to be active.
	ValidityPeriod *types.Period
	// Typically date of issue is different from the beginning of the validity. This can be used to see when an item was last updated.
	LastUpdated *types.DateTime
	// Used for examplefor Out of Formulary, or any specifics.
	AdditionalCharacteristic []types.CodeableConcept
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// Used to support catalog exchange even for unsupported products, e.g. getting list of medications even if not prescribable.
	Status *types.Code
	// Used for example, to point to a substance, or to a device used to administer a medication.
	RelatedEntry []CatalogEntryRelatedEntry
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Used in supporting different identifiers for the same product, e.g. manufacturer code and retailer code.
	Identifier []types.Identifier
	// The type of item - medication, device, service, protocol or other.
	Type *types.CodeableConcept
	// Whether the entry represents an orderable item.
	Orderable types.Boolean
	// Used in supporting related concepts, e.g. NDC to RxNorm.
	AdditionalIdentifier []types.Identifier
	// User for example for ATC classification, or.
	AdditionalClassification []types.CodeableConcept
}

func (s CatalogEntry) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
