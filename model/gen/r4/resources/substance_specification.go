package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Moiety, for structural modifications.
type SubstanceSpecificationMoiety struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Role that the moiety is playing.
	Role *types.CodeableConcept
	// Identifier by which this moiety substance is known.
	Identifier *types.Identifier
	// Textual name for this moiety substance.
	Name *types.String
	// Optical activity type.
	OpticalActivity *types.CodeableConcept
	// Quantitative value for this moiety.
	Amount r4.Element
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Stereochemistry type.
	Stereochemistry *types.CodeableConcept
	// Molecular formula.
	MolecularFormula *types.String
}

func (s SubstanceSpecificationMoiety) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Codes associated with the substance.
type SubstanceSpecificationCode struct {
	// Supporting literature.
	Source []types.Reference
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The specific code.
	Code *types.CodeableConcept
	// Status of the code assignment.
	Status *types.CodeableConcept
	// The date at which the code status is changed as part of the terminology maintenance.
	StatusDate *types.DateTime
	// Any comment can be provided in this field, if necessary.
	Comment *types.String
}

func (s SubstanceSpecificationCode) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Details of the official nature of this name.
type SubstanceSpecificationNameOfficial struct {
	// The status of the official name.
	Status *types.CodeableConcept
	// Date of official name change.
	Date *types.DateTime
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Which authority uses this official name.
	Authority *types.CodeableConcept
}

func (s SubstanceSpecificationNameOfficial) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Names applicable to this substance.
type SubstanceSpecificationName struct {
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Language of the name.
	Language []types.CodeableConcept
	// A synonym of this name.
	Synonym []SubstanceSpecificationName
	// The status of the name.
	Status *types.CodeableConcept
	// The actual name.
	Name types.String
	// The jurisdiction where this name applies.
	Jurisdiction []types.CodeableConcept
	// A translation for this name.
	Translation []SubstanceSpecificationName
	// Details of the official nature of this name.
	Official []SubstanceSpecificationNameOfficial
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Name type.
	Type *types.CodeableConcept
	// If this is the preferred name for this substance.
	Preferred *types.Boolean
	// The use context of this name for example if there is a different name a drug active ingredient as opposed to a food colour additive.
	Domain []types.CodeableConcept
	// Supporting literature.
	Source []types.Reference
}

func (s SubstanceSpecificationName) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A link between this substance and another, with details of the relationship.
type SubstanceSpecificationRelationship struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A numeric factor for the relationship, for instance to express that the salt of a substance has some percentage of the active substance in relation to some other.
	Amount r4.Element
	// Supporting literature.
	Source []types.Reference
	// An operator for the amount, for example "average", "approximately", "less than".
	AmountType *types.CodeableConcept
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A pointer to another substance, as a resource or just a representational code.
	Substance r4.Element
	// For example "salt to parent", "active moiety", "starting material".
	Relationship *types.CodeableConcept
	// For example where an enzyme strongly bonds with a particular substance, this is a defining relationship for that enzyme, out of several possible substance relationships.
	IsDefining *types.Boolean
	// For use when the numeric.
	AmountRatioLowLimit *types.Ratio
}

func (s SubstanceSpecificationRelationship) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// General specifications for this substance, including how it is related to other substances.
type SubstanceSpecificationProperty struct {
	// A substance upon which a defining property depends (e.g. for solubility: in water, in alcohol).
	DefiningSubstance r4.Element
	// Quantitative value for this property.
	Amount r4.Element
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A category for this property, e.g. Physical, Chemical, Enzymatic.
	Category *types.CodeableConcept
	// Property type e.g. viscosity, pH, isoelectric point.
	Code *types.CodeableConcept
	// Parameters that were used in the measurement of a property (e.g. for viscosity: measured at 20C with a pH of 7.1).
	Parameters *types.String
}

func (s SubstanceSpecificationProperty) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Molecular structural representation.
type SubstanceSpecificationStructureRepresentation struct {
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The type of structure (e.g. Full, Partial, Representative).
	Type *types.CodeableConcept
	// The structural representation as text string in a format e.g. InChI, SMILES, MOLFILE, CDX.
	Representation *types.String
	// An attached file with the structural representation.
	Attachment *types.Attachment
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
}

func (s SubstanceSpecificationStructureRepresentation) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The molecular weight or weight range (for proteins, polymers or nucleic acids).
type SubstanceSpecificationStructureIsotopeMolecularWeight struct {
	// Used to capture quantitative values for a variety of elements. If only limits are given, the arithmetic mean would be the average. If only a single definite value for a given element is given, it would be captured in this field.
	Amount *types.Quantity
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The method by which the molecular weight was determined.
	Method *types.CodeableConcept
	// Type of molecular weight such as exact, average (also known as. number average), weight average.
	Type *types.CodeableConcept
}

func (s SubstanceSpecificationStructureIsotopeMolecularWeight) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Applicable for single substances that contain a radionuclide or a non-natural isotopic ratio.
type SubstanceSpecificationStructureIsotope struct {
	// The type of isotopic substitution present in a single substance.
	Substitution *types.CodeableConcept
	// Half life - for a non-natural nuclide.
	HalfLife *types.Quantity
	// The molecular weight or weight range (for proteins, polymers or nucleic acids).
	MolecularWeight *SubstanceSpecificationStructureIsotopeMolecularWeight
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Substance identifier for each non-natural or radioisotope.
	Identifier *types.Identifier
	// Substance name for each non-natural or radioisotope.
	Name *types.CodeableConcept
}

func (s SubstanceSpecificationStructureIsotope) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Structural information.
type SubstanceSpecificationStructure struct {
	// Stereochemistry type.
	Stereochemistry *types.CodeableConcept
	// Specified per moiety according to the Hill system, i.e. first C, then H, then alphabetical, each moiety separated by a dot.
	MolecularFormulaByMoiety *types.String
	// The molecular weight or weight range (for proteins, polymers or nucleic acids).
	MolecularWeight *SubstanceSpecificationStructureIsotopeMolecularWeight
	// Supporting literature.
	Source []types.Reference
	// Molecular structural representation.
	Representation []SubstanceSpecificationStructureRepresentation
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Optical activity type.
	OpticalActivity *types.CodeableConcept
	// Molecular formula.
	MolecularFormula *types.String
	// Applicable for single substances that contain a radionuclide or a non-natural isotopic ratio.
	Isotope []SubstanceSpecificationStructureIsotope
}

func (s SubstanceSpecificationStructure) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The detailed description of a substance, typically at a level beyond what is used for prescribing.
type SubstanceSpecification struct {
	// Supporting literature.
	Source []types.Reference
	// General information detailing this substance.
	ReferenceInformation *types.Reference
	// If the substance applies to only human or veterinary use.
	Domain *types.CodeableConcept
	// Material or taxonomic/anatomical source for the substance.
	SourceMaterial *types.Reference
	// Status of substance within the catalogue e.g. approved.
	Status *types.CodeableConcept
	// Moiety, for structural modifications.
	Moiety []SubstanceSpecificationMoiety
	// Data items specific to proteins.
	Protein *types.Reference
	// Codes associated with the substance.
	Code []SubstanceSpecificationCode
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// Names applicable to this substance.
	Name []SubstanceSpecificationName
	// Data items specific to nucleic acids.
	NucleicAcid *types.Reference
	// Textual comment about this record of a substance.
	Comment *types.String
	// Identifier by which this substance is known.
	Identifier *types.Identifier
	// Textual description of the substance.
	Description *types.String
	// A link between this substance and another, with details of the relationship.
	Relationship []SubstanceSpecificationRelationship
	// High level categorization, e.g. polymer or nucleic acid.
	Type *types.CodeableConcept
	// General specifications for this substance, including how it is related to other substances.
	Property []SubstanceSpecificationProperty
	// Structural information.
	Structure *SubstanceSpecificationStructure
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The molecular weight or weight range (for proteins, polymers or nucleic acids).
	MolecularWeight []SubstanceSpecificationStructureIsotopeMolecularWeight
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// The base language in which the resource is written.
	Language *types.Code
	// Data items specific to polymers.
	Polymer *types.Reference
}

func (s SubstanceSpecification) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
