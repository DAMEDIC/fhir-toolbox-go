package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Country where the name applies.
type MedicinalProductNameCountryLanguage struct {
	// Country code for where this name applies.
	Country types.CodeableConcept
	// Jurisdiction code for where this name applies.
	Jurisdiction *types.CodeableConcept
	// Language code for this name.
	Language types.CodeableConcept
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
}

func (s MedicinalProductNameCountryLanguage) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Coding words or phrases of the name.
type MedicinalProductNameNamePart struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A fragment of a product name.
	Part types.String
	// Idenifying type for this part of the name (e.g. strength part).
	Type types.Coding
}

func (s MedicinalProductNameNamePart) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The product's name, including full name and possibly coded parts.
type MedicinalProductName struct {
	// Country where the name applies.
	CountryLanguage []MedicinalProductNameCountryLanguage
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The full product name.
	ProductName types.String
	// Coding words or phrases of the name.
	NamePart []MedicinalProductNameNamePart
}

func (s MedicinalProductName) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// An operation applied to the product, for manufacturing or adminsitrative purpose.
type MedicinalProductManufacturingBusinessOperation struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Regulatory authorization reference number.
	AuthorisationReferenceNumber *types.Identifier
	// Regulatory authorization date.
	EffectiveDate *types.DateTime
	// To indicate if this proces is commercially confidential.
	ConfidentialityIndicator *types.CodeableConcept
	// The manufacturer or establishment associated with the process.
	Manufacturer []types.Reference
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// The type of manufacturing operation.
	OperationType *types.CodeableConcept
	// A regulator which oversees the operation.
	Regulator *types.Reference
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
}

func (s MedicinalProductManufacturingBusinessOperation) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates if the medicinal product has an orphan designation for the treatment of a rare disease.
type MedicinalProductSpecialDesignation struct {
	// Identifier for the designation, or procedure number.
	Identifier []types.Identifier
	// The intended use of the product, e.g. prevention, treatment.
	IntendedUse *types.CodeableConcept
	// Condition for which the medicinal use applies.
	Indication r4.Element
	// Date when the designation was granted.
	Date *types.DateTime
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// For example granted, pending, expired or withdrawn.
	Status *types.CodeableConcept
	// Animal species for which this applies.
	Species *types.CodeableConcept
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The type of special designation, e.g. orphan drug, minor use.
	Type *types.CodeableConcept
}

func (s MedicinalProductSpecialDesignation) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Detailed definition of a medicinal product, typically for uses other than direct patient care (e.g. regulatory use).
type MedicinalProduct struct {
	// Package representation for the product.
	PackagedMedicinalProduct []types.Reference
	// The product's name, including full name and possibly coded parts.
	Name []MedicinalProductName
	// Reference to another product, e.g. for linking authorised to investigational product.
	CrossReference []types.Identifier
	// An operation applied to the product, for manufacturing or adminsitrative purpose.
	ManufacturingBusinessOperation []MedicinalProductManufacturingBusinessOperation
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// Marketing status of the medicinal product, in contrast to marketing authorizaton.
	MarketingStatus []types.MarketingStatus
	// Business identifier for this product. Could be an MPID.
	Identifier []types.Identifier
	// Pharmaceutical aspects of product.
	PharmaceuticalProduct []types.Reference
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// Whether the Medicinal Product is subject to additional monitoring for regulatory reasons.
	AdditionalMonitoringIndicator *types.CodeableConcept
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Whether the Medicinal Product is subject to special measures for regulatory reasons.
	SpecialMeasures []types.String
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Regulatory type, e.g. Investigational or Authorized.
	Type *types.CodeableConcept
	// If this medicine applies to human or veterinary uses.
	Domain *types.Coding
	// The dose form for a single part product, or combined form of a multiple part product.
	CombinedPharmaceuticalDoseForm *types.CodeableConcept
	// A product specific contact, person (in a role), or an organization.
	Contact []types.Reference
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// The base language in which the resource is written.
	Language *types.Code
	// Clinical trials or studies that this product is involved in.
	ClinicalTrial []types.Reference
	// Indicates if the medicinal product has an orphan designation for the treatment of a rare disease.
	SpecialDesignation []MedicinalProductSpecialDesignation
	// If authorised for use in children.
	PaediatricUseIndicator *types.CodeableConcept
	// Allows the product to be classified by various systems.
	ProductClassification []types.CodeableConcept
	// Supporting documentation, typically for regulatory submission.
	AttachedDocument []types.Reference
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// The legal status of supply of the medicinal product as classified by the regulator.
	LegalStatusOfSupply *types.CodeableConcept
	// A master file for to the medicinal product (e.g. Pharmacovigilance System Master File).
	MasterFile []types.Reference
}

func (s MedicinalProduct) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
