package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
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
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
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

func (r SubstanceSourceMaterial) ResourceType() string {
	return "SubstanceSourceMaterial"
}

type jsonSubstanceSourceMaterial struct {
	ResourceType                         string                                       `json:"resourceType"`
	Id                                   *Id                                          `json:"id,omitempty"`
	IdPrimitiveElement                   *primitiveElement                            `json:"_id,omitempty"`
	Meta                                 *Meta                                        `json:"meta,omitempty"`
	ImplicitRules                        *Uri                                         `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement        *primitiveElement                            `json:"_implicitRules,omitempty"`
	Language                             *Code                                        `json:"language,omitempty"`
	LanguagePrimitiveElement             *primitiveElement                            `json:"_language,omitempty"`
	Text                                 *Narrative                                   `json:"text,omitempty"`
	Contained                            []containedResource                          `json:"contained,omitempty"`
	Extension                            []Extension                                  `json:"extension,omitempty"`
	ModifierExtension                    []Extension                                  `json:"modifierExtension,omitempty"`
	SourceMaterialClass                  *CodeableConcept                             `json:"sourceMaterialClass,omitempty"`
	SourceMaterialType                   *CodeableConcept                             `json:"sourceMaterialType,omitempty"`
	SourceMaterialState                  *CodeableConcept                             `json:"sourceMaterialState,omitempty"`
	OrganismId                           *Identifier                                  `json:"organismId,omitempty"`
	OrganismName                         *String                                      `json:"organismName,omitempty"`
	OrganismNamePrimitiveElement         *primitiveElement                            `json:"_organismName,omitempty"`
	ParentSubstanceId                    []Identifier                                 `json:"parentSubstanceId,omitempty"`
	ParentSubstanceName                  []String                                     `json:"parentSubstanceName,omitempty"`
	ParentSubstanceNamePrimitiveElement  []*primitiveElement                          `json:"_parentSubstanceName,omitempty"`
	CountryOfOrigin                      []CodeableConcept                            `json:"countryOfOrigin,omitempty"`
	GeographicalLocation                 []String                                     `json:"geographicalLocation,omitempty"`
	GeographicalLocationPrimitiveElement []*primitiveElement                          `json:"_geographicalLocation,omitempty"`
	DevelopmentStage                     *CodeableConcept                             `json:"developmentStage,omitempty"`
	FractionDescription                  []SubstanceSourceMaterialFractionDescription `json:"fractionDescription,omitempty"`
	Organism                             *SubstanceSourceMaterialOrganism             `json:"organism,omitempty"`
	PartDescription                      []SubstanceSourceMaterialPartDescription     `json:"partDescription,omitempty"`
}

func (r SubstanceSourceMaterial) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSourceMaterial) marshalJSON() jsonSubstanceSourceMaterial {
	m := jsonSubstanceSourceMaterial{}
	m.ResourceType = "SubstanceSourceMaterial"
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
	m.SourceMaterialClass = r.SourceMaterialClass
	m.SourceMaterialType = r.SourceMaterialType
	m.SourceMaterialState = r.SourceMaterialState
	m.OrganismId = r.OrganismId
	m.OrganismName = r.OrganismName
	if r.OrganismName != nil && (r.OrganismName.Id != nil || r.OrganismName.Extension != nil) {
		m.OrganismNamePrimitiveElement = &primitiveElement{Id: r.OrganismName.Id, Extension: r.OrganismName.Extension}
	}
	m.ParentSubstanceId = r.ParentSubstanceId
	m.ParentSubstanceName = r.ParentSubstanceName
	anyParentSubstanceNameIdOrExtension := false
	for _, e := range r.ParentSubstanceName {
		if e.Id != nil || e.Extension != nil {
			anyParentSubstanceNameIdOrExtension = true
			break
		}
	}
	if anyParentSubstanceNameIdOrExtension {
		m.ParentSubstanceNamePrimitiveElement = make([]*primitiveElement, 0, len(r.ParentSubstanceName))
		for _, e := range r.ParentSubstanceName {
			if e.Id != nil || e.Extension != nil {
				m.ParentSubstanceNamePrimitiveElement = append(m.ParentSubstanceNamePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ParentSubstanceNamePrimitiveElement = append(m.ParentSubstanceNamePrimitiveElement, nil)
			}
		}
	}
	m.CountryOfOrigin = r.CountryOfOrigin
	m.GeographicalLocation = r.GeographicalLocation
	anyGeographicalLocationIdOrExtension := false
	for _, e := range r.GeographicalLocation {
		if e.Id != nil || e.Extension != nil {
			anyGeographicalLocationIdOrExtension = true
			break
		}
	}
	if anyGeographicalLocationIdOrExtension {
		m.GeographicalLocationPrimitiveElement = make([]*primitiveElement, 0, len(r.GeographicalLocation))
		for _, e := range r.GeographicalLocation {
			if e.Id != nil || e.Extension != nil {
				m.GeographicalLocationPrimitiveElement = append(m.GeographicalLocationPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.GeographicalLocationPrimitiveElement = append(m.GeographicalLocationPrimitiveElement, nil)
			}
		}
	}
	m.DevelopmentStage = r.DevelopmentStage
	m.FractionDescription = r.FractionDescription
	m.Organism = r.Organism
	m.PartDescription = r.PartDescription
	return m
}
func (r *SubstanceSourceMaterial) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSourceMaterial
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSourceMaterial) unmarshalJSON(m jsonSubstanceSourceMaterial) error {
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
	r.SourceMaterialClass = m.SourceMaterialClass
	r.SourceMaterialType = m.SourceMaterialType
	r.SourceMaterialState = m.SourceMaterialState
	r.OrganismId = m.OrganismId
	r.OrganismName = m.OrganismName
	if m.OrganismNamePrimitiveElement != nil {
		r.OrganismName.Id = m.OrganismNamePrimitiveElement.Id
		r.OrganismName.Extension = m.OrganismNamePrimitiveElement.Extension
	}
	r.ParentSubstanceId = m.ParentSubstanceId
	r.ParentSubstanceName = m.ParentSubstanceName
	for i, e := range m.ParentSubstanceNamePrimitiveElement {
		if len(r.ParentSubstanceName) > i {
			r.ParentSubstanceName[i].Id = e.Id
			r.ParentSubstanceName[i].Extension = e.Extension
		} else {
			r.ParentSubstanceName = append(r.ParentSubstanceName, String{Id: e.Id, Extension: e.Extension})
		}
	}
	r.CountryOfOrigin = m.CountryOfOrigin
	r.GeographicalLocation = m.GeographicalLocation
	for i, e := range m.GeographicalLocationPrimitiveElement {
		if len(r.GeographicalLocation) > i {
			r.GeographicalLocation[i].Id = e.Id
			r.GeographicalLocation[i].Extension = e.Extension
		} else {
			r.GeographicalLocation = append(r.GeographicalLocation, String{Id: e.Id, Extension: e.Extension})
		}
	}
	r.DevelopmentStage = m.DevelopmentStage
	r.FractionDescription = m.FractionDescription
	r.Organism = m.Organism
	r.PartDescription = m.PartDescription
	return nil
}
func (r SubstanceSourceMaterial) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Many complex materials are fractions of parts of plants, animals, or minerals. Fraction elements are often necessary to define both Substances and Specified Group 1 Substances. For substances derived from Plants, fraction information will be captured at the Substance information level ( . Oils, Juices and Exudates). Additional information for Extracts, such as extraction solvent composition, will be captured at the Specified Substance Group 1 information level. For plasma-derived products fraction information will be captured at the Substance and the Specified Substance Group 1 levels.
type SubstanceSourceMaterialFractionDescription struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// This element is capturing information about the fraction of a plant part, or human plasma for fractionation.
	Fraction *String
	// The specific type of the material constituting the component. For Herbal preparations the particulars of the extracts (liquid/dry) is described in Specified Substance Group 1.
	MaterialType *CodeableConcept
}
type jsonSubstanceSourceMaterialFractionDescription struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Fraction                 *String           `json:"fraction,omitempty"`
	FractionPrimitiveElement *primitiveElement `json:"_fraction,omitempty"`
	MaterialType             *CodeableConcept  `json:"materialType,omitempty"`
}

func (r SubstanceSourceMaterialFractionDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSourceMaterialFractionDescription) marshalJSON() jsonSubstanceSourceMaterialFractionDescription {
	m := jsonSubstanceSourceMaterialFractionDescription{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Fraction = r.Fraction
	if r.Fraction != nil && (r.Fraction.Id != nil || r.Fraction.Extension != nil) {
		m.FractionPrimitiveElement = &primitiveElement{Id: r.Fraction.Id, Extension: r.Fraction.Extension}
	}
	m.MaterialType = r.MaterialType
	return m
}
func (r *SubstanceSourceMaterialFractionDescription) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSourceMaterialFractionDescription
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSourceMaterialFractionDescription) unmarshalJSON(m jsonSubstanceSourceMaterialFractionDescription) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Fraction = m.Fraction
	if m.FractionPrimitiveElement != nil {
		r.Fraction.Id = m.FractionPrimitiveElement.Id
		r.Fraction.Extension = m.FractionPrimitiveElement.Extension
	}
	r.MaterialType = m.MaterialType
	return nil
}
func (r SubstanceSourceMaterialFractionDescription) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// This subclause describes the organism which the substance is derived from. For vaccines, the parent organism shall be specified based on these subclause elements. As an example, full taxonomy will be described for the Substance Name: ., Leaf.
type SubstanceSourceMaterialOrganism struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
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
type jsonSubstanceSourceMaterialOrganism struct {
	Id                                       *string                                         `json:"id,omitempty"`
	Extension                                []Extension                                     `json:"extension,omitempty"`
	ModifierExtension                        []Extension                                     `json:"modifierExtension,omitempty"`
	Family                                   *CodeableConcept                                `json:"family,omitempty"`
	Genus                                    *CodeableConcept                                `json:"genus,omitempty"`
	Species                                  *CodeableConcept                                `json:"species,omitempty"`
	IntraspecificType                        *CodeableConcept                                `json:"intraspecificType,omitempty"`
	IntraspecificDescription                 *String                                         `json:"intraspecificDescription,omitempty"`
	IntraspecificDescriptionPrimitiveElement *primitiveElement                               `json:"_intraspecificDescription,omitempty"`
	Author                                   []SubstanceSourceMaterialOrganismAuthor         `json:"author,omitempty"`
	Hybrid                                   *SubstanceSourceMaterialOrganismHybrid          `json:"hybrid,omitempty"`
	OrganismGeneral                          *SubstanceSourceMaterialOrganismOrganismGeneral `json:"organismGeneral,omitempty"`
}

func (r SubstanceSourceMaterialOrganism) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSourceMaterialOrganism) marshalJSON() jsonSubstanceSourceMaterialOrganism {
	m := jsonSubstanceSourceMaterialOrganism{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Family = r.Family
	m.Genus = r.Genus
	m.Species = r.Species
	m.IntraspecificType = r.IntraspecificType
	m.IntraspecificDescription = r.IntraspecificDescription
	if r.IntraspecificDescription != nil && (r.IntraspecificDescription.Id != nil || r.IntraspecificDescription.Extension != nil) {
		m.IntraspecificDescriptionPrimitiveElement = &primitiveElement{Id: r.IntraspecificDescription.Id, Extension: r.IntraspecificDescription.Extension}
	}
	m.Author = r.Author
	m.Hybrid = r.Hybrid
	m.OrganismGeneral = r.OrganismGeneral
	return m
}
func (r *SubstanceSourceMaterialOrganism) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSourceMaterialOrganism
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSourceMaterialOrganism) unmarshalJSON(m jsonSubstanceSourceMaterialOrganism) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Family = m.Family
	r.Genus = m.Genus
	r.Species = m.Species
	r.IntraspecificType = m.IntraspecificType
	r.IntraspecificDescription = m.IntraspecificDescription
	if m.IntraspecificDescriptionPrimitiveElement != nil {
		r.IntraspecificDescription.Id = m.IntraspecificDescriptionPrimitiveElement.Id
		r.IntraspecificDescription.Extension = m.IntraspecificDescriptionPrimitiveElement.Extension
	}
	r.Author = m.Author
	r.Hybrid = m.Hybrid
	r.OrganismGeneral = m.OrganismGeneral
	return nil
}
func (r SubstanceSourceMaterialOrganism) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// 4.9.13.6.1 Author type (Conditional).
type SubstanceSourceMaterialOrganismAuthor struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of author of an organism species shall be specified. The parenthetical author of an organism species refers to the first author who published the plant/animal name (of any rank). The primary author of an organism species refers to the first author(s), who validly published the plant/animal name.
	AuthorType *CodeableConcept
	// The author of an organism species shall be specified. The author year of an organism shall also be specified when applicable; refers to the year in which the first author(s) published the infraspecific plant/animal name (of any rank).
	AuthorDescription *String
}
type jsonSubstanceSourceMaterialOrganismAuthor struct {
	Id                                *string           `json:"id,omitempty"`
	Extension                         []Extension       `json:"extension,omitempty"`
	ModifierExtension                 []Extension       `json:"modifierExtension,omitempty"`
	AuthorType                        *CodeableConcept  `json:"authorType,omitempty"`
	AuthorDescription                 *String           `json:"authorDescription,omitempty"`
	AuthorDescriptionPrimitiveElement *primitiveElement `json:"_authorDescription,omitempty"`
}

func (r SubstanceSourceMaterialOrganismAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSourceMaterialOrganismAuthor) marshalJSON() jsonSubstanceSourceMaterialOrganismAuthor {
	m := jsonSubstanceSourceMaterialOrganismAuthor{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.AuthorType = r.AuthorType
	m.AuthorDescription = r.AuthorDescription
	if r.AuthorDescription != nil && (r.AuthorDescription.Id != nil || r.AuthorDescription.Extension != nil) {
		m.AuthorDescriptionPrimitiveElement = &primitiveElement{Id: r.AuthorDescription.Id, Extension: r.AuthorDescription.Extension}
	}
	return m
}
func (r *SubstanceSourceMaterialOrganismAuthor) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSourceMaterialOrganismAuthor
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSourceMaterialOrganismAuthor) unmarshalJSON(m jsonSubstanceSourceMaterialOrganismAuthor) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.AuthorType = m.AuthorType
	r.AuthorDescription = m.AuthorDescription
	if m.AuthorDescriptionPrimitiveElement != nil {
		r.AuthorDescription.Id = m.AuthorDescriptionPrimitiveElement.Id
		r.AuthorDescription.Extension = m.AuthorDescriptionPrimitiveElement.Extension
	}
	return nil
}
func (r SubstanceSourceMaterialOrganismAuthor) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// 4.9.13.8.1 Hybrid species maternal organism ID (Optional).
type SubstanceSourceMaterialOrganismHybrid struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
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
type jsonSubstanceSourceMaterialOrganismHybrid struct {
	Id                                   *string           `json:"id,omitempty"`
	Extension                            []Extension       `json:"extension,omitempty"`
	ModifierExtension                    []Extension       `json:"modifierExtension,omitempty"`
	MaternalOrganismId                   *String           `json:"maternalOrganismId,omitempty"`
	MaternalOrganismIdPrimitiveElement   *primitiveElement `json:"_maternalOrganismId,omitempty"`
	MaternalOrganismName                 *String           `json:"maternalOrganismName,omitempty"`
	MaternalOrganismNamePrimitiveElement *primitiveElement `json:"_maternalOrganismName,omitempty"`
	PaternalOrganismId                   *String           `json:"paternalOrganismId,omitempty"`
	PaternalOrganismIdPrimitiveElement   *primitiveElement `json:"_paternalOrganismId,omitempty"`
	PaternalOrganismName                 *String           `json:"paternalOrganismName,omitempty"`
	PaternalOrganismNamePrimitiveElement *primitiveElement `json:"_paternalOrganismName,omitempty"`
	HybridType                           *CodeableConcept  `json:"hybridType,omitempty"`
}

func (r SubstanceSourceMaterialOrganismHybrid) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSourceMaterialOrganismHybrid) marshalJSON() jsonSubstanceSourceMaterialOrganismHybrid {
	m := jsonSubstanceSourceMaterialOrganismHybrid{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.MaternalOrganismId = r.MaternalOrganismId
	if r.MaternalOrganismId != nil && (r.MaternalOrganismId.Id != nil || r.MaternalOrganismId.Extension != nil) {
		m.MaternalOrganismIdPrimitiveElement = &primitiveElement{Id: r.MaternalOrganismId.Id, Extension: r.MaternalOrganismId.Extension}
	}
	m.MaternalOrganismName = r.MaternalOrganismName
	if r.MaternalOrganismName != nil && (r.MaternalOrganismName.Id != nil || r.MaternalOrganismName.Extension != nil) {
		m.MaternalOrganismNamePrimitiveElement = &primitiveElement{Id: r.MaternalOrganismName.Id, Extension: r.MaternalOrganismName.Extension}
	}
	m.PaternalOrganismId = r.PaternalOrganismId
	if r.PaternalOrganismId != nil && (r.PaternalOrganismId.Id != nil || r.PaternalOrganismId.Extension != nil) {
		m.PaternalOrganismIdPrimitiveElement = &primitiveElement{Id: r.PaternalOrganismId.Id, Extension: r.PaternalOrganismId.Extension}
	}
	m.PaternalOrganismName = r.PaternalOrganismName
	if r.PaternalOrganismName != nil && (r.PaternalOrganismName.Id != nil || r.PaternalOrganismName.Extension != nil) {
		m.PaternalOrganismNamePrimitiveElement = &primitiveElement{Id: r.PaternalOrganismName.Id, Extension: r.PaternalOrganismName.Extension}
	}
	m.HybridType = r.HybridType
	return m
}
func (r *SubstanceSourceMaterialOrganismHybrid) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSourceMaterialOrganismHybrid
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSourceMaterialOrganismHybrid) unmarshalJSON(m jsonSubstanceSourceMaterialOrganismHybrid) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.MaternalOrganismId = m.MaternalOrganismId
	if m.MaternalOrganismIdPrimitiveElement != nil {
		r.MaternalOrganismId.Id = m.MaternalOrganismIdPrimitiveElement.Id
		r.MaternalOrganismId.Extension = m.MaternalOrganismIdPrimitiveElement.Extension
	}
	r.MaternalOrganismName = m.MaternalOrganismName
	if m.MaternalOrganismNamePrimitiveElement != nil {
		r.MaternalOrganismName.Id = m.MaternalOrganismNamePrimitiveElement.Id
		r.MaternalOrganismName.Extension = m.MaternalOrganismNamePrimitiveElement.Extension
	}
	r.PaternalOrganismId = m.PaternalOrganismId
	if m.PaternalOrganismIdPrimitiveElement != nil {
		r.PaternalOrganismId.Id = m.PaternalOrganismIdPrimitiveElement.Id
		r.PaternalOrganismId.Extension = m.PaternalOrganismIdPrimitiveElement.Extension
	}
	r.PaternalOrganismName = m.PaternalOrganismName
	if m.PaternalOrganismNamePrimitiveElement != nil {
		r.PaternalOrganismName.Id = m.PaternalOrganismNamePrimitiveElement.Id
		r.PaternalOrganismName.Extension = m.PaternalOrganismNamePrimitiveElement.Extension
	}
	r.HybridType = m.HybridType
	return nil
}
func (r SubstanceSourceMaterialOrganismHybrid) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// 4.9.13.7.1 Kingdom (Conditional).
type SubstanceSourceMaterialOrganismOrganismGeneral struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
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
type jsonSubstanceSourceMaterialOrganismOrganismGeneral struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Kingdom           *CodeableConcept `json:"kingdom,omitempty"`
	Phylum            *CodeableConcept `json:"phylum,omitempty"`
	Class             *CodeableConcept `json:"class,omitempty"`
	Order             *CodeableConcept `json:"order,omitempty"`
}

func (r SubstanceSourceMaterialOrganismOrganismGeneral) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) marshalJSON() jsonSubstanceSourceMaterialOrganismOrganismGeneral {
	m := jsonSubstanceSourceMaterialOrganismOrganismGeneral{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Kingdom = r.Kingdom
	m.Phylum = r.Phylum
	m.Class = r.Class
	m.Order = r.Order
	return m
}
func (r *SubstanceSourceMaterialOrganismOrganismGeneral) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSourceMaterialOrganismOrganismGeneral
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSourceMaterialOrganismOrganismGeneral) unmarshalJSON(m jsonSubstanceSourceMaterialOrganismOrganismGeneral) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Kingdom = m.Kingdom
	r.Phylum = m.Phylum
	r.Class = m.Class
	r.Order = m.Order
	return nil
}
func (r SubstanceSourceMaterialOrganismOrganismGeneral) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// To do.
type SubstanceSourceMaterialPartDescription struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Entity of anatomical origin of source material within an organism.
	Part *CodeableConcept
	// The detailed anatomic location when the part can be extracted from different anatomical locations of the organism. Multiple alternative locations may apply.
	PartLocation *CodeableConcept
}
type jsonSubstanceSourceMaterialPartDescription struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Part              *CodeableConcept `json:"part,omitempty"`
	PartLocation      *CodeableConcept `json:"partLocation,omitempty"`
}

func (r SubstanceSourceMaterialPartDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSourceMaterialPartDescription) marshalJSON() jsonSubstanceSourceMaterialPartDescription {
	m := jsonSubstanceSourceMaterialPartDescription{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Part = r.Part
	m.PartLocation = r.PartLocation
	return m
}
func (r *SubstanceSourceMaterialPartDescription) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSourceMaterialPartDescription
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSourceMaterialPartDescription) unmarshalJSON(m jsonSubstanceSourceMaterialPartDescription) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Part = m.Part
	r.PartLocation = m.PartLocation
	return nil
}
func (r SubstanceSourceMaterialPartDescription) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
