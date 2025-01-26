package r5

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

// Source material shall capture information on the taxonomic and anatomical origins as well as the fraction of a material that can result in or can be modified to form a substance. This set of data elements shall be used to define polymer substances isolated from biological matrices. Taxonomic and anatomical origins shall be described using a controlled vocabulary as required. This information is captured for naturally derived polymers ( . starch) and structurally diverse substances. For Organisms belonging to the Kingdom Plantae the Substance level defines the fresh material of a single species or infraspecies, the Herbal Drug and the Herbal preparation. For Herbal preparations, the fraction information will be captured at the Substance information level and additional information for herbal extracts will be captured at the Specified Substance Group 1 information level. See for further explanation the Substance Class: Structurally Diverse and the herbal annex.
type SubstanceSourceMaterial struct {
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
	// General high level classification of the source material specific to the origin of the material.
	SourceMaterialClass *CodeableConcept
	// The type of the source material shall be specified based on a controlled vocabulary. For vaccines, this subclause refers to the class of infectious agent.
	SourceMaterialType *CodeableConcept
	// The state of the source material when extracted.
	SourceMaterialState *CodeableConcept
	// The unique identifier associated with the source material parent organism shall be specified.
	OrganismId *Identifier
	// The organism accepted Scientific name shall be provided based on the organism taxonomy.
	OrganismName *String
	// The parent of the herbal drug Ginkgo biloba, Leaf is the substance ID of the substance (fresh) of Ginkgo biloba L. or Ginkgo biloba L. (Whole plant).
	ParentSubstanceId []Identifier
	// The parent substance of the Herbal Drug, or Herbal preparation.
	ParentSubstanceName []String
	// The country where the plant material is harvested or the countries where the plasma is sourced from as laid down in accordance with the Plasma Master File. For “Plasma-derived substances” the attribute country of origin provides information about the countries used for the manufacturing of the Cryopoor plama or Crioprecipitate.
	CountryOfOrigin []CodeableConcept
	// The place/region where the plant is harvested or the places/regions where the animal source material has its habitat.
	GeographicalLocation []String
	// Stage of life for animals, plants, insects and microorganisms. This information shall be provided only when the substance is significantly different in these stages (e.g. foetal bovine serum).
	DevelopmentStage *CodeableConcept
	// Many complex materials are fractions of parts of plants, animals, or minerals. Fraction elements are often necessary to define both Substances and Specified Group 1 Substances. For substances derived from Plants, fraction information will be captured at the Substance information level ( . Oils, Juices and Exudates). Additional information for Extracts, such as extraction solvent composition, will be captured at the Specified Substance Group 1 information level. For plasma-derived products fraction information will be captured at the Substance and the Specified Substance Group 1 levels.
	FractionDescription []SubstanceSourceMaterialFractionDescription
	// This subclause describes the organism which the substance is derived from. For vaccines, the parent organism shall be specified based on these subclause elements. As an example, full taxonomy will be described for the Substance Name: ., Leaf.
	Organism *SubstanceSourceMaterialOrganism
	// To do.
	PartDescription []SubstanceSourceMaterialPartDescription
}

// Many complex materials are fractions of parts of plants, animals, or minerals. Fraction elements are often necessary to define both Substances and Specified Group 1 Substances. For substances derived from Plants, fraction information will be captured at the Substance information level ( . Oils, Juices and Exudates). Additional information for Extracts, such as extraction solvent composition, will be captured at the Specified Substance Group 1 information level. For plasma-derived products fraction information will be captured at the Substance and the Specified Substance Group 1 levels.
type SubstanceSourceMaterialFractionDescription struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// This element is capturing information about the fraction of a plant part, or human plasma for fractionation.
	Fraction *String
	// The specific type of the material constituting the component. For Herbal preparations the particulars of the extracts (liquid/dry) is described in Specified Substance Group 1.
	MaterialType *CodeableConcept
}

// This subclause describes the organism which the substance is derived from. For vaccines, the parent organism shall be specified based on these subclause elements. As an example, full taxonomy will be described for the Substance Name: ., Leaf.
type SubstanceSourceMaterialOrganism struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The family of an organism shall be specified.
	Family *CodeableConcept
	// The genus of an organism shall be specified; refers to the Latin epithet of the genus element of the plant/animal scientific name; it is present in names for genera, species and infraspecies.
	Genus *CodeableConcept
	// The species of an organism shall be specified; refers to the Latin epithet of the species of the plant/animal; it is present in names for species and infraspecies.
	Species *CodeableConcept
	// The Intraspecific type of an organism shall be specified.
	IntraspecificType *CodeableConcept
	// The intraspecific description of an organism shall be specified based on a controlled vocabulary. For Influenza Vaccine, the intraspecific description shall contain the syntax of the antigen in line with the WHO convention.
	IntraspecificDescription *String
	// 4.9.13.6.1 Author type (Conditional).
	Author []SubstanceSourceMaterialOrganismAuthor
	// 4.9.13.8.1 Hybrid species maternal organism ID (Optional).
	Hybrid *SubstanceSourceMaterialOrganismHybrid
	// 4.9.13.7.1 Kingdom (Conditional).
	OrganismGeneral *SubstanceSourceMaterialOrganismOrganismGeneral
}

// 4.9.13.6.1 Author type (Conditional).
type SubstanceSourceMaterialOrganismAuthor struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of author of an organism species shall be specified. The parenthetical author of an organism species refers to the first author who published the plant/animal name (of any rank). The primary author of an organism species refers to the first author(s), who validly published the plant/animal name.
	AuthorType *CodeableConcept
	// The author of an organism species shall be specified. The author year of an organism shall also be specified when applicable; refers to the year in which the first author(s) published the infraspecific plant/animal name (of any rank).
	AuthorDescription *String
}

// 4.9.13.8.1 Hybrid species maternal organism ID (Optional).
type SubstanceSourceMaterialOrganismHybrid struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The identifier of the maternal species constituting the hybrid organism shall be specified based on a controlled vocabulary. For plants, the parents aren’t always known, and it is unlikely that it will be known which is maternal and which is paternal.
	MaternalOrganismId *String
	// The name of the maternal species constituting the hybrid organism shall be specified. For plants, the parents aren’t always known, and it is unlikely that it will be known which is maternal and which is paternal.
	MaternalOrganismName *String
	// The identifier of the paternal species constituting the hybrid organism shall be specified based on a controlled vocabulary.
	PaternalOrganismId *String
	// The name of the paternal species constituting the hybrid organism shall be specified.
	PaternalOrganismName *String
	// The hybrid type of an organism shall be specified.
	HybridType *CodeableConcept
}

// 4.9.13.7.1 Kingdom (Conditional).
type SubstanceSourceMaterialOrganismOrganismGeneral struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The kingdom of an organism shall be specified.
	Kingdom *CodeableConcept
	// The phylum of an organism shall be specified.
	Phylum *CodeableConcept
	// The class of an organism shall be specified.
	Class *CodeableConcept
	// The order of an organism shall be specified,.
	Order *CodeableConcept
}

// To do.
type SubstanceSourceMaterialPartDescription struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Entity of anatomical origin of source material within an organism.
	Part *CodeableConcept
	// The detailed anatomic location when the part can be extracted from different anatomical locations of the organism. Multiple alternative locations may apply.
	PartLocation *CodeableConcept
}

func (r SubstanceSourceMaterial) ResourceType() string {
	return "SubstanceSourceMaterial"
}
func (r SubstanceSourceMaterial) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r SubstanceSourceMaterial) MemSize() int {
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
	if r.SourceMaterialClass != nil {
		s += r.SourceMaterialClass.MemSize()
	}
	if r.SourceMaterialType != nil {
		s += r.SourceMaterialType.MemSize()
	}
	if r.SourceMaterialState != nil {
		s += r.SourceMaterialState.MemSize()
	}
	if r.OrganismId != nil {
		s += r.OrganismId.MemSize()
	}
	if r.OrganismName != nil {
		s += r.OrganismName.MemSize()
	}
	for _, i := range r.ParentSubstanceId {
		s += i.MemSize()
	}
	s += (cap(r.ParentSubstanceId) - len(r.ParentSubstanceId)) * int(unsafe.Sizeof(Identifier{}))
	for _, i := range r.ParentSubstanceName {
		s += i.MemSize()
	}
	s += (cap(r.ParentSubstanceName) - len(r.ParentSubstanceName)) * int(unsafe.Sizeof(String{}))
	for _, i := range r.CountryOfOrigin {
		s += i.MemSize()
	}
	s += (cap(r.CountryOfOrigin) - len(r.CountryOfOrigin)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.GeographicalLocation {
		s += i.MemSize()
	}
	s += (cap(r.GeographicalLocation) - len(r.GeographicalLocation)) * int(unsafe.Sizeof(String{}))
	if r.DevelopmentStage != nil {
		s += r.DevelopmentStage.MemSize()
	}
	for _, i := range r.FractionDescription {
		s += i.MemSize()
	}
	s += (cap(r.FractionDescription) - len(r.FractionDescription)) * int(unsafe.Sizeof(SubstanceSourceMaterialFractionDescription{}))
	if r.Organism != nil {
		s += r.Organism.MemSize()
	}
	for _, i := range r.PartDescription {
		s += i.MemSize()
	}
	s += (cap(r.PartDescription) - len(r.PartDescription)) * int(unsafe.Sizeof(SubstanceSourceMaterialPartDescription{}))
	return s
}
func (r SubstanceSourceMaterialFractionDescription) MemSize() int {
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
	if r.Fraction != nil {
		s += r.Fraction.MemSize()
	}
	if r.MaterialType != nil {
		s += r.MaterialType.MemSize()
	}
	return s
}
func (r SubstanceSourceMaterialOrganism) MemSize() int {
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
	if r.Family != nil {
		s += r.Family.MemSize()
	}
	if r.Genus != nil {
		s += r.Genus.MemSize()
	}
	if r.Species != nil {
		s += r.Species.MemSize()
	}
	if r.IntraspecificType != nil {
		s += r.IntraspecificType.MemSize()
	}
	if r.IntraspecificDescription != nil {
		s += r.IntraspecificDescription.MemSize()
	}
	for _, i := range r.Author {
		s += i.MemSize()
	}
	s += (cap(r.Author) - len(r.Author)) * int(unsafe.Sizeof(SubstanceSourceMaterialOrganismAuthor{}))
	if r.Hybrid != nil {
		s += r.Hybrid.MemSize()
	}
	if r.OrganismGeneral != nil {
		s += r.OrganismGeneral.MemSize()
	}
	return s
}
func (r SubstanceSourceMaterialOrganismAuthor) MemSize() int {
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
	if r.AuthorType != nil {
		s += r.AuthorType.MemSize()
	}
	if r.AuthorDescription != nil {
		s += r.AuthorDescription.MemSize()
	}
	return s
}
func (r SubstanceSourceMaterialOrganismHybrid) MemSize() int {
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
	if r.MaternalOrganismId != nil {
		s += r.MaternalOrganismId.MemSize()
	}
	if r.MaternalOrganismName != nil {
		s += r.MaternalOrganismName.MemSize()
	}
	if r.PaternalOrganismId != nil {
		s += r.PaternalOrganismId.MemSize()
	}
	if r.PaternalOrganismName != nil {
		s += r.PaternalOrganismName.MemSize()
	}
	if r.HybridType != nil {
		s += r.HybridType.MemSize()
	}
	return s
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) MemSize() int {
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
	if r.Kingdom != nil {
		s += r.Kingdom.MemSize()
	}
	if r.Phylum != nil {
		s += r.Phylum.MemSize()
	}
	if r.Class != nil {
		s += r.Class.MemSize()
	}
	if r.Order != nil {
		s += r.Order.MemSize()
	}
	return s
}
func (r SubstanceSourceMaterialPartDescription) MemSize() int {
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
	if r.Part != nil {
		s += r.Part.MemSize()
	}
	if r.PartLocation != nil {
		s += r.PartLocation.MemSize()
	}
	return s
}
func (r SubstanceSourceMaterial) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r SubstanceSourceMaterial) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r SubstanceSourceMaterial) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"SubstanceSourceMaterial\""))
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
	if r.SourceMaterialClass != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sourceMaterialClass\":"))
		if err != nil {
			return err
		}
		err = r.SourceMaterialClass.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.SourceMaterialType != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sourceMaterialType\":"))
		if err != nil {
			return err
		}
		err = r.SourceMaterialType.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.SourceMaterialState != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sourceMaterialState\":"))
		if err != nil {
			return err
		}
		err = r.SourceMaterialState.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.OrganismId != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"organismId\":"))
		if err != nil {
			return err
		}
		err = r.OrganismId.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.OrganismName != nil && r.OrganismName.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"organismName\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.OrganismName)
		if err != nil {
			return err
		}
	}
	if r.OrganismName != nil && (r.OrganismName.Id != nil || r.OrganismName.Extension != nil) {
		p := primitiveElement{Id: r.OrganismName.Id, Extension: r.OrganismName.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_organismName\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.ParentSubstanceId) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"parentSubstanceId\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ParentSubstanceId {
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
		anyValue := false
		for _, e := range r.ParentSubstanceName {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"parentSubstanceName\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.ParentSubstanceName)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.ParentSubstanceName {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_parentSubstanceName\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.ParentSubstanceName {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	if len(r.CountryOfOrigin) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"countryOfOrigin\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.CountryOfOrigin {
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
		anyValue := false
		for _, e := range r.GeographicalLocation {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"geographicalLocation\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.GeographicalLocation)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.GeographicalLocation {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_geographicalLocation\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.GeographicalLocation {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	if r.DevelopmentStage != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"developmentStage\":"))
		if err != nil {
			return err
		}
		err = r.DevelopmentStage.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.FractionDescription) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"fractionDescription\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.FractionDescription {
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
	if r.Organism != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"organism\":"))
		if err != nil {
			return err
		}
		err = r.Organism.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.PartDescription) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"partDescription\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.PartDescription {
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
func (r SubstanceSourceMaterialFractionDescription) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r SubstanceSourceMaterialFractionDescription) marshalJSON(w io.Writer) error {
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
	if r.Fraction != nil && r.Fraction.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"fraction\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Fraction)
		if err != nil {
			return err
		}
	}
	if r.Fraction != nil && (r.Fraction.Id != nil || r.Fraction.Extension != nil) {
		p := primitiveElement{Id: r.Fraction.Id, Extension: r.Fraction.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_fraction\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.MaterialType != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"materialType\":"))
		if err != nil {
			return err
		}
		err = r.MaterialType.marshalJSON(w)
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
func (r SubstanceSourceMaterialOrganism) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r SubstanceSourceMaterialOrganism) marshalJSON(w io.Writer) error {
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
	if r.Family != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"family\":"))
		if err != nil {
			return err
		}
		err = r.Family.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Genus != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"genus\":"))
		if err != nil {
			return err
		}
		err = r.Genus.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Species != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"species\":"))
		if err != nil {
			return err
		}
		err = r.Species.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.IntraspecificType != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"intraspecificType\":"))
		if err != nil {
			return err
		}
		err = r.IntraspecificType.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.IntraspecificDescription != nil && r.IntraspecificDescription.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"intraspecificDescription\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.IntraspecificDescription)
		if err != nil {
			return err
		}
	}
	if r.IntraspecificDescription != nil && (r.IntraspecificDescription.Id != nil || r.IntraspecificDescription.Extension != nil) {
		p := primitiveElement{Id: r.IntraspecificDescription.Id, Extension: r.IntraspecificDescription.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_intraspecificDescription\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Author) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"author\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Author {
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
	if r.Hybrid != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"hybrid\":"))
		if err != nil {
			return err
		}
		err = r.Hybrid.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.OrganismGeneral != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"organismGeneral\":"))
		if err != nil {
			return err
		}
		err = r.OrganismGeneral.marshalJSON(w)
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
func (r SubstanceSourceMaterialOrganismAuthor) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r SubstanceSourceMaterialOrganismAuthor) marshalJSON(w io.Writer) error {
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
	if r.AuthorType != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"authorType\":"))
		if err != nil {
			return err
		}
		err = r.AuthorType.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.AuthorDescription != nil && r.AuthorDescription.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"authorDescription\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.AuthorDescription)
		if err != nil {
			return err
		}
	}
	if r.AuthorDescription != nil && (r.AuthorDescription.Id != nil || r.AuthorDescription.Extension != nil) {
		p := primitiveElement{Id: r.AuthorDescription.Id, Extension: r.AuthorDescription.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_authorDescription\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
func (r SubstanceSourceMaterialOrganismHybrid) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r SubstanceSourceMaterialOrganismHybrid) marshalJSON(w io.Writer) error {
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
	if r.MaternalOrganismId != nil && r.MaternalOrganismId.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"maternalOrganismId\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.MaternalOrganismId)
		if err != nil {
			return err
		}
	}
	if r.MaternalOrganismId != nil && (r.MaternalOrganismId.Id != nil || r.MaternalOrganismId.Extension != nil) {
		p := primitiveElement{Id: r.MaternalOrganismId.Id, Extension: r.MaternalOrganismId.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_maternalOrganismId\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.MaternalOrganismName != nil && r.MaternalOrganismName.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"maternalOrganismName\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.MaternalOrganismName)
		if err != nil {
			return err
		}
	}
	if r.MaternalOrganismName != nil && (r.MaternalOrganismName.Id != nil || r.MaternalOrganismName.Extension != nil) {
		p := primitiveElement{Id: r.MaternalOrganismName.Id, Extension: r.MaternalOrganismName.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_maternalOrganismName\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.PaternalOrganismId != nil && r.PaternalOrganismId.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"paternalOrganismId\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.PaternalOrganismId)
		if err != nil {
			return err
		}
	}
	if r.PaternalOrganismId != nil && (r.PaternalOrganismId.Id != nil || r.PaternalOrganismId.Extension != nil) {
		p := primitiveElement{Id: r.PaternalOrganismId.Id, Extension: r.PaternalOrganismId.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_paternalOrganismId\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.PaternalOrganismName != nil && r.PaternalOrganismName.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"paternalOrganismName\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.PaternalOrganismName)
		if err != nil {
			return err
		}
	}
	if r.PaternalOrganismName != nil && (r.PaternalOrganismName.Id != nil || r.PaternalOrganismName.Extension != nil) {
		p := primitiveElement{Id: r.PaternalOrganismName.Id, Extension: r.PaternalOrganismName.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_paternalOrganismName\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.HybridType != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"hybridType\":"))
		if err != nil {
			return err
		}
		err = r.HybridType.marshalJSON(w)
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
func (r SubstanceSourceMaterialOrganismOrganismGeneral) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) marshalJSON(w io.Writer) error {
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
	if r.Kingdom != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"kingdom\":"))
		if err != nil {
			return err
		}
		err = r.Kingdom.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Phylum != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"phylum\":"))
		if err != nil {
			return err
		}
		err = r.Phylum.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Class != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"class\":"))
		if err != nil {
			return err
		}
		err = r.Class.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Order != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"order\":"))
		if err != nil {
			return err
		}
		err = r.Order.marshalJSON(w)
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
func (r SubstanceSourceMaterialPartDescription) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r SubstanceSourceMaterialPartDescription) marshalJSON(w io.Writer) error {
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
	if r.Part != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"part\":"))
		if err != nil {
			return err
		}
		err = r.Part.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.PartLocation != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"partLocation\":"))
		if err != nil {
			return err
		}
		err = r.PartLocation.marshalJSON(w)
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
func (r *SubstanceSourceMaterial) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *SubstanceSourceMaterial) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in SubstanceSourceMaterial element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in SubstanceSourceMaterial element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterial element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterial element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterial element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterial element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterial element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterial element", t)
			}
		case "sourceMaterialClass":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.SourceMaterialClass = &v
		case "sourceMaterialType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.SourceMaterialType = &v
		case "sourceMaterialState":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.SourceMaterialState = &v
		case "organismId":
			var v Identifier
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.OrganismId = &v
		case "organismName":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.OrganismName == nil {
				r.OrganismName = &String{}
			}
			r.OrganismName.Value = v.Value
		case "_organismName":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.OrganismName == nil {
				r.OrganismName = &String{}
			}
			r.OrganismName.Id = v.Id
			r.OrganismName.Extension = v.Extension
		case "parentSubstanceId":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterial element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ParentSubstanceId = append(r.ParentSubstanceId, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterial element", t)
			}
		case "parentSubstanceName":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterial element", t)
			}
			for i := 0; d.More(); i++ {
				var v String
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.ParentSubstanceName) <= i {
					r.ParentSubstanceName = append(r.ParentSubstanceName, String{})
				}
				r.ParentSubstanceName[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterial element", t)
			}
		case "_parentSubstanceName":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterial element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.ParentSubstanceName) <= i {
					r.ParentSubstanceName = append(r.ParentSubstanceName, String{})
				}
				r.ParentSubstanceName[i].Id = v.Id
				r.ParentSubstanceName[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterial element", t)
			}
		case "countryOfOrigin":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterial element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.CountryOfOrigin = append(r.CountryOfOrigin, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterial element", t)
			}
		case "geographicalLocation":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterial element", t)
			}
			for i := 0; d.More(); i++ {
				var v String
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.GeographicalLocation) <= i {
					r.GeographicalLocation = append(r.GeographicalLocation, String{})
				}
				r.GeographicalLocation[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterial element", t)
			}
		case "_geographicalLocation":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterial element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.GeographicalLocation) <= i {
					r.GeographicalLocation = append(r.GeographicalLocation, String{})
				}
				r.GeographicalLocation[i].Id = v.Id
				r.GeographicalLocation[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterial element", t)
			}
		case "developmentStage":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.DevelopmentStage = &v
		case "fractionDescription":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterial element", t)
			}
			for d.More() {
				var v SubstanceSourceMaterialFractionDescription
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.FractionDescription = append(r.FractionDescription, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterial element", t)
			}
		case "organism":
			var v SubstanceSourceMaterialOrganism
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Organism = &v
		case "partDescription":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterial element", t)
			}
			for d.More() {
				var v SubstanceSourceMaterialPartDescription
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.PartDescription = append(r.PartDescription, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterial element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in SubstanceSourceMaterial", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in SubstanceSourceMaterial element", t)
	}
	return nil
}
func (r *SubstanceSourceMaterialFractionDescription) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in SubstanceSourceMaterialFractionDescription element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in SubstanceSourceMaterialFractionDescription element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterialFractionDescription element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterialFractionDescription element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterialFractionDescription element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterialFractionDescription element", t)
			}
		case "fraction":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Fraction == nil {
				r.Fraction = &String{}
			}
			r.Fraction.Value = v.Value
		case "_fraction":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Fraction == nil {
				r.Fraction = &String{}
			}
			r.Fraction.Id = v.Id
			r.Fraction.Extension = v.Extension
		case "materialType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.MaterialType = &v
		default:
			return fmt.Errorf("invalid field: %s in SubstanceSourceMaterialFractionDescription", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in SubstanceSourceMaterialFractionDescription element", t)
	}
	return nil
}
func (r *SubstanceSourceMaterialOrganism) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in SubstanceSourceMaterialOrganism element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in SubstanceSourceMaterialOrganism element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterialOrganism element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterialOrganism element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterialOrganism element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterialOrganism element", t)
			}
		case "family":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Family = &v
		case "genus":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Genus = &v
		case "species":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Species = &v
		case "intraspecificType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.IntraspecificType = &v
		case "intraspecificDescription":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.IntraspecificDescription == nil {
				r.IntraspecificDescription = &String{}
			}
			r.IntraspecificDescription.Value = v.Value
		case "_intraspecificDescription":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.IntraspecificDescription == nil {
				r.IntraspecificDescription = &String{}
			}
			r.IntraspecificDescription.Id = v.Id
			r.IntraspecificDescription.Extension = v.Extension
		case "author":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterialOrganism element", t)
			}
			for d.More() {
				var v SubstanceSourceMaterialOrganismAuthor
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Author = append(r.Author, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterialOrganism element", t)
			}
		case "hybrid":
			var v SubstanceSourceMaterialOrganismHybrid
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Hybrid = &v
		case "organismGeneral":
			var v SubstanceSourceMaterialOrganismOrganismGeneral
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.OrganismGeneral = &v
		default:
			return fmt.Errorf("invalid field: %s in SubstanceSourceMaterialOrganism", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in SubstanceSourceMaterialOrganism element", t)
	}
	return nil
}
func (r *SubstanceSourceMaterialOrganismAuthor) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in SubstanceSourceMaterialOrganismAuthor element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in SubstanceSourceMaterialOrganismAuthor element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterialOrganismAuthor element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterialOrganismAuthor element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterialOrganismAuthor element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterialOrganismAuthor element", t)
			}
		case "authorType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.AuthorType = &v
		case "authorDescription":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.AuthorDescription == nil {
				r.AuthorDescription = &String{}
			}
			r.AuthorDescription.Value = v.Value
		case "_authorDescription":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.AuthorDescription == nil {
				r.AuthorDescription = &String{}
			}
			r.AuthorDescription.Id = v.Id
			r.AuthorDescription.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in SubstanceSourceMaterialOrganismAuthor", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in SubstanceSourceMaterialOrganismAuthor element", t)
	}
	return nil
}
func (r *SubstanceSourceMaterialOrganismHybrid) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in SubstanceSourceMaterialOrganismHybrid element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in SubstanceSourceMaterialOrganismHybrid element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterialOrganismHybrid element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterialOrganismHybrid element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterialOrganismHybrid element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterialOrganismHybrid element", t)
			}
		case "maternalOrganismId":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.MaternalOrganismId == nil {
				r.MaternalOrganismId = &String{}
			}
			r.MaternalOrganismId.Value = v.Value
		case "_maternalOrganismId":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.MaternalOrganismId == nil {
				r.MaternalOrganismId = &String{}
			}
			r.MaternalOrganismId.Id = v.Id
			r.MaternalOrganismId.Extension = v.Extension
		case "maternalOrganismName":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.MaternalOrganismName == nil {
				r.MaternalOrganismName = &String{}
			}
			r.MaternalOrganismName.Value = v.Value
		case "_maternalOrganismName":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.MaternalOrganismName == nil {
				r.MaternalOrganismName = &String{}
			}
			r.MaternalOrganismName.Id = v.Id
			r.MaternalOrganismName.Extension = v.Extension
		case "paternalOrganismId":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.PaternalOrganismId == nil {
				r.PaternalOrganismId = &String{}
			}
			r.PaternalOrganismId.Value = v.Value
		case "_paternalOrganismId":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.PaternalOrganismId == nil {
				r.PaternalOrganismId = &String{}
			}
			r.PaternalOrganismId.Id = v.Id
			r.PaternalOrganismId.Extension = v.Extension
		case "paternalOrganismName":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.PaternalOrganismName == nil {
				r.PaternalOrganismName = &String{}
			}
			r.PaternalOrganismName.Value = v.Value
		case "_paternalOrganismName":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.PaternalOrganismName == nil {
				r.PaternalOrganismName = &String{}
			}
			r.PaternalOrganismName.Id = v.Id
			r.PaternalOrganismName.Extension = v.Extension
		case "hybridType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.HybridType = &v
		default:
			return fmt.Errorf("invalid field: %s in SubstanceSourceMaterialOrganismHybrid", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in SubstanceSourceMaterialOrganismHybrid element", t)
	}
	return nil
}
func (r *SubstanceSourceMaterialOrganismOrganismGeneral) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in SubstanceSourceMaterialOrganismOrganismGeneral element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in SubstanceSourceMaterialOrganismOrganismGeneral element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterialOrganismOrganismGeneral element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterialOrganismOrganismGeneral element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterialOrganismOrganismGeneral element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterialOrganismOrganismGeneral element", t)
			}
		case "kingdom":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Kingdom = &v
		case "phylum":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Phylum = &v
		case "class":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Class = &v
		case "order":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Order = &v
		default:
			return fmt.Errorf("invalid field: %s in SubstanceSourceMaterialOrganismOrganismGeneral", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in SubstanceSourceMaterialOrganismOrganismGeneral element", t)
	}
	return nil
}
func (r *SubstanceSourceMaterialPartDescription) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in SubstanceSourceMaterialPartDescription element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in SubstanceSourceMaterialPartDescription element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterialPartDescription element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterialPartDescription element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in SubstanceSourceMaterialPartDescription element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SubstanceSourceMaterialPartDescription element", t)
			}
		case "part":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Part = &v
		case "partLocation":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.PartLocation = &v
		default:
			return fmt.Errorf("invalid field: %s in SubstanceSourceMaterialPartDescription", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in SubstanceSourceMaterialPartDescription element", t)
	}
	return nil
}
func (r SubstanceSourceMaterial) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "SubstanceSourceMaterial"
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
	err = e.EncodeElement(r.SourceMaterialClass, xml.StartElement{Name: xml.Name{Local: "sourceMaterialClass"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SourceMaterialType, xml.StartElement{Name: xml.Name{Local: "sourceMaterialType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SourceMaterialState, xml.StartElement{Name: xml.Name{Local: "sourceMaterialState"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OrganismId, xml.StartElement{Name: xml.Name{Local: "organismId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OrganismName, xml.StartElement{Name: xml.Name{Local: "organismName"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ParentSubstanceId, xml.StartElement{Name: xml.Name{Local: "parentSubstanceId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ParentSubstanceName, xml.StartElement{Name: xml.Name{Local: "parentSubstanceName"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CountryOfOrigin, xml.StartElement{Name: xml.Name{Local: "countryOfOrigin"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.GeographicalLocation, xml.StartElement{Name: xml.Name{Local: "geographicalLocation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DevelopmentStage, xml.StartElement{Name: xml.Name{Local: "developmentStage"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FractionDescription, xml.StartElement{Name: xml.Name{Local: "fractionDescription"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Organism, xml.StartElement{Name: xml.Name{Local: "organism"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PartDescription, xml.StartElement{Name: xml.Name{Local: "partDescription"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r SubstanceSourceMaterialFractionDescription) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Fraction, xml.StartElement{Name: xml.Name{Local: "fraction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MaterialType, xml.StartElement{Name: xml.Name{Local: "materialType"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r SubstanceSourceMaterialOrganism) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Family, xml.StartElement{Name: xml.Name{Local: "family"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Genus, xml.StartElement{Name: xml.Name{Local: "genus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Species, xml.StartElement{Name: xml.Name{Local: "species"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IntraspecificType, xml.StartElement{Name: xml.Name{Local: "intraspecificType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IntraspecificDescription, xml.StartElement{Name: xml.Name{Local: "intraspecificDescription"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Author, xml.StartElement{Name: xml.Name{Local: "author"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Hybrid, xml.StartElement{Name: xml.Name{Local: "hybrid"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OrganismGeneral, xml.StartElement{Name: xml.Name{Local: "organismGeneral"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r SubstanceSourceMaterialOrganismAuthor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.AuthorType, xml.StartElement{Name: xml.Name{Local: "authorType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AuthorDescription, xml.StartElement{Name: xml.Name{Local: "authorDescription"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r SubstanceSourceMaterialOrganismHybrid) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.MaternalOrganismId, xml.StartElement{Name: xml.Name{Local: "maternalOrganismId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MaternalOrganismName, xml.StartElement{Name: xml.Name{Local: "maternalOrganismName"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PaternalOrganismId, xml.StartElement{Name: xml.Name{Local: "paternalOrganismId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PaternalOrganismName, xml.StartElement{Name: xml.Name{Local: "paternalOrganismName"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.HybridType, xml.StartElement{Name: xml.Name{Local: "hybridType"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Kingdom, xml.StartElement{Name: xml.Name{Local: "kingdom"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Phylum, xml.StartElement{Name: xml.Name{Local: "phylum"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Class, xml.StartElement{Name: xml.Name{Local: "class"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Order, xml.StartElement{Name: xml.Name{Local: "order"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r SubstanceSourceMaterialPartDescription) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Part, xml.StartElement{Name: xml.Name{Local: "part"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PartLocation, xml.StartElement{Name: xml.Name{Local: "partLocation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceSourceMaterial) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "sourceMaterialClass":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SourceMaterialClass = &v
			case "sourceMaterialType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SourceMaterialType = &v
			case "sourceMaterialState":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SourceMaterialState = &v
			case "organismId":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OrganismId = &v
			case "organismName":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OrganismName = &v
			case "parentSubstanceId":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ParentSubstanceId = append(r.ParentSubstanceId, v)
			case "parentSubstanceName":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ParentSubstanceName = append(r.ParentSubstanceName, v)
			case "countryOfOrigin":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CountryOfOrigin = append(r.CountryOfOrigin, v)
			case "geographicalLocation":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.GeographicalLocation = append(r.GeographicalLocation, v)
			case "developmentStage":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DevelopmentStage = &v
			case "fractionDescription":
				var v SubstanceSourceMaterialFractionDescription
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FractionDescription = append(r.FractionDescription, v)
			case "organism":
				var v SubstanceSourceMaterialOrganism
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Organism = &v
			case "partDescription":
				var v SubstanceSourceMaterialPartDescription
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PartDescription = append(r.PartDescription, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *SubstanceSourceMaterialFractionDescription) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "fraction":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Fraction = &v
			case "materialType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MaterialType = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *SubstanceSourceMaterialOrganism) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "family":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Family = &v
			case "genus":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Genus = &v
			case "species":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Species = &v
			case "intraspecificType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IntraspecificType = &v
			case "intraspecificDescription":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IntraspecificDescription = &v
			case "author":
				var v SubstanceSourceMaterialOrganismAuthor
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Author = append(r.Author, v)
			case "hybrid":
				var v SubstanceSourceMaterialOrganismHybrid
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Hybrid = &v
			case "organismGeneral":
				var v SubstanceSourceMaterialOrganismOrganismGeneral
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OrganismGeneral = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *SubstanceSourceMaterialOrganismAuthor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "authorType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AuthorType = &v
			case "authorDescription":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AuthorDescription = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *SubstanceSourceMaterialOrganismHybrid) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "maternalOrganismId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MaternalOrganismId = &v
			case "maternalOrganismName":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MaternalOrganismName = &v
			case "paternalOrganismId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PaternalOrganismId = &v
			case "paternalOrganismName":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PaternalOrganismName = &v
			case "hybridType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.HybridType = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *SubstanceSourceMaterialOrganismOrganismGeneral) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "kingdom":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Kingdom = &v
			case "phylum":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Phylum = &v
			case "class":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Class = &v
			case "order":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Order = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *SubstanceSourceMaterialPartDescription) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "part":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Part = &v
			case "partLocation":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PartLocation = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceSourceMaterial) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "sourceMaterialClass") {
		if r.SourceMaterialClass != nil {
			children = append(children, *r.SourceMaterialClass)
		}
	}
	if len(name) == 0 || slices.Contains(name, "sourceMaterialType") {
		if r.SourceMaterialType != nil {
			children = append(children, *r.SourceMaterialType)
		}
	}
	if len(name) == 0 || slices.Contains(name, "sourceMaterialState") {
		if r.SourceMaterialState != nil {
			children = append(children, *r.SourceMaterialState)
		}
	}
	if len(name) == 0 || slices.Contains(name, "organismId") {
		if r.OrganismId != nil {
			children = append(children, *r.OrganismId)
		}
	}
	if len(name) == 0 || slices.Contains(name, "organismName") {
		if r.OrganismName != nil {
			children = append(children, *r.OrganismName)
		}
	}
	if len(name) == 0 || slices.Contains(name, "parentSubstanceId") {
		for _, v := range r.ParentSubstanceId {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "parentSubstanceName") {
		for _, v := range r.ParentSubstanceName {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "countryOfOrigin") {
		for _, v := range r.CountryOfOrigin {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "geographicalLocation") {
		for _, v := range r.GeographicalLocation {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "developmentStage") {
		if r.DevelopmentStage != nil {
			children = append(children, *r.DevelopmentStage)
		}
	}
	if len(name) == 0 || slices.Contains(name, "fractionDescription") {
		for _, v := range r.FractionDescription {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "organism") {
		if r.Organism != nil {
			children = append(children, *r.Organism)
		}
	}
	if len(name) == 0 || slices.Contains(name, "partDescription") {
		for _, v := range r.PartDescription {
			children = append(children, v)
		}
	}
	return children
}
func (r SubstanceSourceMaterial) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterial to Boolean")
}
func (r SubstanceSourceMaterial) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterial to String")
}
func (r SubstanceSourceMaterial) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterial to Integer")
}
func (r SubstanceSourceMaterial) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterial to Decimal")
}
func (r SubstanceSourceMaterial) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterial to Date")
}
func (r SubstanceSourceMaterial) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterial to Time")
}
func (r SubstanceSourceMaterial) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterial to DateTime")
}
func (r SubstanceSourceMaterial) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterial to Quantity")
}
func (r SubstanceSourceMaterial) TypeInfo() fhirpath.TypeInfo {
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
			Name: "SourceMaterialClass",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "SourceMaterialType",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "SourceMaterialState",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "OrganismId",
			Type: "FHIR.Identifier",
		}, {
			Name: "OrganismName",
			Type: "FHIR.String",
		}, {
			Name: "ParentSubstanceId",
			Type: "List<FHIR.Identifier>",
		}, {
			Name: "ParentSubstanceName",
			Type: "List<FHIR.String>",
		}, {
			Name: "CountryOfOrigin",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "GeographicalLocation",
			Type: "List<FHIR.String>",
		}, {
			Name: "DevelopmentStage",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "FractionDescription",
			Type: "List<FHIR.SubstanceSourceMaterialFractionDescription>",
		}, {
			Name: "Organism",
			Type: "FHIR.SubstanceSourceMaterialOrganism",
		}, {
			Name: "PartDescription",
			Type: "List<FHIR.SubstanceSourceMaterialPartDescription>",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "SubstanceSourceMaterial",
			Namespace: "FHIR",
		},
	}
}
func (r SubstanceSourceMaterialFractionDescription) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "fraction") {
		if r.Fraction != nil {
			children = append(children, *r.Fraction)
		}
	}
	if len(name) == 0 || slices.Contains(name, "materialType") {
		if r.MaterialType != nil {
			children = append(children, *r.MaterialType)
		}
	}
	return children
}
func (r SubstanceSourceMaterialFractionDescription) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialFractionDescription to Boolean")
}
func (r SubstanceSourceMaterialFractionDescription) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialFractionDescription to String")
}
func (r SubstanceSourceMaterialFractionDescription) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialFractionDescription to Integer")
}
func (r SubstanceSourceMaterialFractionDescription) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialFractionDescription to Decimal")
}
func (r SubstanceSourceMaterialFractionDescription) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialFractionDescription to Date")
}
func (r SubstanceSourceMaterialFractionDescription) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialFractionDescription to Time")
}
func (r SubstanceSourceMaterialFractionDescription) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialFractionDescription to DateTime")
}
func (r SubstanceSourceMaterialFractionDescription) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialFractionDescription to Quantity")
}
func (r SubstanceSourceMaterialFractionDescription) TypeInfo() fhirpath.TypeInfo {
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
			Name: "Fraction",
			Type: "FHIR.String",
		}, {
			Name: "MaterialType",
			Type: "FHIR.CodeableConcept",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "SubstanceSourceMaterialFractionDescription",
			Namespace: "FHIR",
		},
	}
}
func (r SubstanceSourceMaterialOrganism) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "family") {
		if r.Family != nil {
			children = append(children, *r.Family)
		}
	}
	if len(name) == 0 || slices.Contains(name, "genus") {
		if r.Genus != nil {
			children = append(children, *r.Genus)
		}
	}
	if len(name) == 0 || slices.Contains(name, "species") {
		if r.Species != nil {
			children = append(children, *r.Species)
		}
	}
	if len(name) == 0 || slices.Contains(name, "intraspecificType") {
		if r.IntraspecificType != nil {
			children = append(children, *r.IntraspecificType)
		}
	}
	if len(name) == 0 || slices.Contains(name, "intraspecificDescription") {
		if r.IntraspecificDescription != nil {
			children = append(children, *r.IntraspecificDescription)
		}
	}
	if len(name) == 0 || slices.Contains(name, "author") {
		for _, v := range r.Author {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "hybrid") {
		if r.Hybrid != nil {
			children = append(children, *r.Hybrid)
		}
	}
	if len(name) == 0 || slices.Contains(name, "organismGeneral") {
		if r.OrganismGeneral != nil {
			children = append(children, *r.OrganismGeneral)
		}
	}
	return children
}
func (r SubstanceSourceMaterialOrganism) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganism to Boolean")
}
func (r SubstanceSourceMaterialOrganism) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganism to String")
}
func (r SubstanceSourceMaterialOrganism) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganism to Integer")
}
func (r SubstanceSourceMaterialOrganism) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganism to Decimal")
}
func (r SubstanceSourceMaterialOrganism) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganism to Date")
}
func (r SubstanceSourceMaterialOrganism) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganism to Time")
}
func (r SubstanceSourceMaterialOrganism) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganism to DateTime")
}
func (r SubstanceSourceMaterialOrganism) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganism to Quantity")
}
func (r SubstanceSourceMaterialOrganism) TypeInfo() fhirpath.TypeInfo {
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
			Name: "Family",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Genus",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Species",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "IntraspecificType",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "IntraspecificDescription",
			Type: "FHIR.String",
		}, {
			Name: "Author",
			Type: "List<FHIR.SubstanceSourceMaterialOrganismAuthor>",
		}, {
			Name: "Hybrid",
			Type: "FHIR.SubstanceSourceMaterialOrganismHybrid",
		}, {
			Name: "OrganismGeneral",
			Type: "FHIR.SubstanceSourceMaterialOrganismOrganismGeneral",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "SubstanceSourceMaterialOrganism",
			Namespace: "FHIR",
		},
	}
}
func (r SubstanceSourceMaterialOrganismAuthor) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "authorType") {
		if r.AuthorType != nil {
			children = append(children, *r.AuthorType)
		}
	}
	if len(name) == 0 || slices.Contains(name, "authorDescription") {
		if r.AuthorDescription != nil {
			children = append(children, *r.AuthorDescription)
		}
	}
	return children
}
func (r SubstanceSourceMaterialOrganismAuthor) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismAuthor to Boolean")
}
func (r SubstanceSourceMaterialOrganismAuthor) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismAuthor to String")
}
func (r SubstanceSourceMaterialOrganismAuthor) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismAuthor to Integer")
}
func (r SubstanceSourceMaterialOrganismAuthor) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismAuthor to Decimal")
}
func (r SubstanceSourceMaterialOrganismAuthor) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismAuthor to Date")
}
func (r SubstanceSourceMaterialOrganismAuthor) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismAuthor to Time")
}
func (r SubstanceSourceMaterialOrganismAuthor) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismAuthor to DateTime")
}
func (r SubstanceSourceMaterialOrganismAuthor) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismAuthor to Quantity")
}
func (r SubstanceSourceMaterialOrganismAuthor) TypeInfo() fhirpath.TypeInfo {
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
			Name: "AuthorType",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "AuthorDescription",
			Type: "FHIR.String",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "SubstanceSourceMaterialOrganismAuthor",
			Namespace: "FHIR",
		},
	}
}
func (r SubstanceSourceMaterialOrganismHybrid) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "maternalOrganismId") {
		if r.MaternalOrganismId != nil {
			children = append(children, *r.MaternalOrganismId)
		}
	}
	if len(name) == 0 || slices.Contains(name, "maternalOrganismName") {
		if r.MaternalOrganismName != nil {
			children = append(children, *r.MaternalOrganismName)
		}
	}
	if len(name) == 0 || slices.Contains(name, "paternalOrganismId") {
		if r.PaternalOrganismId != nil {
			children = append(children, *r.PaternalOrganismId)
		}
	}
	if len(name) == 0 || slices.Contains(name, "paternalOrganismName") {
		if r.PaternalOrganismName != nil {
			children = append(children, *r.PaternalOrganismName)
		}
	}
	if len(name) == 0 || slices.Contains(name, "hybridType") {
		if r.HybridType != nil {
			children = append(children, *r.HybridType)
		}
	}
	return children
}
func (r SubstanceSourceMaterialOrganismHybrid) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismHybrid to Boolean")
}
func (r SubstanceSourceMaterialOrganismHybrid) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismHybrid to String")
}
func (r SubstanceSourceMaterialOrganismHybrid) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismHybrid to Integer")
}
func (r SubstanceSourceMaterialOrganismHybrid) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismHybrid to Decimal")
}
func (r SubstanceSourceMaterialOrganismHybrid) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismHybrid to Date")
}
func (r SubstanceSourceMaterialOrganismHybrid) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismHybrid to Time")
}
func (r SubstanceSourceMaterialOrganismHybrid) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismHybrid to DateTime")
}
func (r SubstanceSourceMaterialOrganismHybrid) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismHybrid to Quantity")
}
func (r SubstanceSourceMaterialOrganismHybrid) TypeInfo() fhirpath.TypeInfo {
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
			Name: "MaternalOrganismId",
			Type: "FHIR.String",
		}, {
			Name: "MaternalOrganismName",
			Type: "FHIR.String",
		}, {
			Name: "PaternalOrganismId",
			Type: "FHIR.String",
		}, {
			Name: "PaternalOrganismName",
			Type: "FHIR.String",
		}, {
			Name: "HybridType",
			Type: "FHIR.CodeableConcept",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "SubstanceSourceMaterialOrganismHybrid",
			Namespace: "FHIR",
		},
	}
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "kingdom") {
		if r.Kingdom != nil {
			children = append(children, *r.Kingdom)
		}
	}
	if len(name) == 0 || slices.Contains(name, "phylum") {
		if r.Phylum != nil {
			children = append(children, *r.Phylum)
		}
	}
	if len(name) == 0 || slices.Contains(name, "class") {
		if r.Class != nil {
			children = append(children, *r.Class)
		}
	}
	if len(name) == 0 || slices.Contains(name, "order") {
		if r.Order != nil {
			children = append(children, *r.Order)
		}
	}
	return children
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismOrganismGeneral to Boolean")
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismOrganismGeneral to String")
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismOrganismGeneral to Integer")
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismOrganismGeneral to Decimal")
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismOrganismGeneral to Date")
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismOrganismGeneral to Time")
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismOrganismGeneral to DateTime")
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialOrganismOrganismGeneral to Quantity")
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) TypeInfo() fhirpath.TypeInfo {
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
			Name: "Kingdom",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Phylum",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Class",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Order",
			Type: "FHIR.CodeableConcept",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "SubstanceSourceMaterialOrganismOrganismGeneral",
			Namespace: "FHIR",
		},
	}
}
func (r SubstanceSourceMaterialPartDescription) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "part") {
		if r.Part != nil {
			children = append(children, *r.Part)
		}
	}
	if len(name) == 0 || slices.Contains(name, "partLocation") {
		if r.PartLocation != nil {
			children = append(children, *r.PartLocation)
		}
	}
	return children
}
func (r SubstanceSourceMaterialPartDescription) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialPartDescription to Boolean")
}
func (r SubstanceSourceMaterialPartDescription) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialPartDescription to String")
}
func (r SubstanceSourceMaterialPartDescription) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialPartDescription to Integer")
}
func (r SubstanceSourceMaterialPartDescription) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialPartDescription to Decimal")
}
func (r SubstanceSourceMaterialPartDescription) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialPartDescription to Date")
}
func (r SubstanceSourceMaterialPartDescription) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialPartDescription to Time")
}
func (r SubstanceSourceMaterialPartDescription) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialPartDescription to DateTime")
}
func (r SubstanceSourceMaterialPartDescription) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert SubstanceSourceMaterialPartDescription to Quantity")
}
func (r SubstanceSourceMaterialPartDescription) TypeInfo() fhirpath.TypeInfo {
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
			Name: "Part",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "PartLocation",
			Type: "FHIR.CodeableConcept",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "SubstanceSourceMaterialPartDescription",
			Namespace: "FHIR",
		},
	}
}
