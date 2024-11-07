package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
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
func (r SubstanceSourceMaterial) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
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
	Contained                            []ContainedResource                          `json:"contained,omitempty"`
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
	m.SourceMaterialClass = r.SourceMaterialClass
	m.SourceMaterialType = r.SourceMaterialType
	m.SourceMaterialState = r.SourceMaterialState
	m.OrganismId = r.OrganismId
	if r.OrganismName != nil && r.OrganismName.Value != nil {
		m.OrganismName = r.OrganismName
	}
	if r.OrganismName != nil && (r.OrganismName.Id != nil || r.OrganismName.Extension != nil) {
		m.OrganismNamePrimitiveElement = &primitiveElement{Id: r.OrganismName.Id, Extension: r.OrganismName.Extension}
	}
	m.ParentSubstanceId = r.ParentSubstanceId
	anyParentSubstanceNameValue := false
	for _, e := range r.ParentSubstanceName {
		if e.Value != nil {
			anyParentSubstanceNameValue = true
			break
		}
	}
	if anyParentSubstanceNameValue {
		m.ParentSubstanceName = r.ParentSubstanceName
	}
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
	anyGeographicalLocationValue := false
	for _, e := range r.GeographicalLocation {
		if e.Value != nil {
			anyGeographicalLocationValue = true
			break
		}
	}
	if anyGeographicalLocationValue {
		m.GeographicalLocation = r.GeographicalLocation
	}
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
	r.SourceMaterialClass = m.SourceMaterialClass
	r.SourceMaterialType = m.SourceMaterialType
	r.SourceMaterialState = m.SourceMaterialState
	r.OrganismId = m.OrganismId
	r.OrganismName = m.OrganismName
	if m.OrganismNamePrimitiveElement != nil {
		if r.OrganismName == nil {
			r.OrganismName = &String{}
		}
		r.OrganismName.Id = m.OrganismNamePrimitiveElement.Id
		r.OrganismName.Extension = m.OrganismNamePrimitiveElement.Extension
	}
	r.ParentSubstanceId = m.ParentSubstanceId
	r.ParentSubstanceName = m.ParentSubstanceName
	for i, e := range m.ParentSubstanceNamePrimitiveElement {
		if len(r.ParentSubstanceName) <= i {
			r.ParentSubstanceName = append(r.ParentSubstanceName, String{})
		}
		if e != nil {
			r.ParentSubstanceName[i].Id = e.Id
			r.ParentSubstanceName[i].Extension = e.Extension
		}
	}
	r.CountryOfOrigin = m.CountryOfOrigin
	r.GeographicalLocation = m.GeographicalLocation
	for i, e := range m.GeographicalLocationPrimitiveElement {
		if len(r.GeographicalLocation) <= i {
			r.GeographicalLocation = append(r.GeographicalLocation, String{})
		}
		if e != nil {
			r.GeographicalLocation[i].Id = e.Id
			r.GeographicalLocation[i].Extension = e.Extension
		}
	}
	r.DevelopmentStage = m.DevelopmentStage
	r.FractionDescription = m.FractionDescription
	r.Organism = m.Organism
	r.PartDescription = m.PartDescription
	return nil
}
func (r SubstanceSourceMaterial) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
func (r SubstanceSourceMaterial) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
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
	if r.Fraction != nil && r.Fraction.Value != nil {
		m.Fraction = r.Fraction
	}
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
		if r.Fraction == nil {
			r.Fraction = &String{}
		}
		r.Fraction.Id = m.FractionPrimitiveElement.Id
		r.Fraction.Extension = m.FractionPrimitiveElement.Extension
	}
	r.MaterialType = m.MaterialType
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
func (r SubstanceSourceMaterialFractionDescription) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
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
	if r.IntraspecificDescription != nil && r.IntraspecificDescription.Value != nil {
		m.IntraspecificDescription = r.IntraspecificDescription
	}
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
		if r.IntraspecificDescription == nil {
			r.IntraspecificDescription = &String{}
		}
		r.IntraspecificDescription.Id = m.IntraspecificDescriptionPrimitiveElement.Id
		r.IntraspecificDescription.Extension = m.IntraspecificDescriptionPrimitiveElement.Extension
	}
	r.Author = m.Author
	r.Hybrid = m.Hybrid
	r.OrganismGeneral = m.OrganismGeneral
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
func (r SubstanceSourceMaterialOrganism) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
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
	if r.AuthorDescription != nil && r.AuthorDescription.Value != nil {
		m.AuthorDescription = r.AuthorDescription
	}
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
		if r.AuthorDescription == nil {
			r.AuthorDescription = &String{}
		}
		r.AuthorDescription.Id = m.AuthorDescriptionPrimitiveElement.Id
		r.AuthorDescription.Extension = m.AuthorDescriptionPrimitiveElement.Extension
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
func (r SubstanceSourceMaterialOrganismAuthor) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
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
	if r.MaternalOrganismId != nil && r.MaternalOrganismId.Value != nil {
		m.MaternalOrganismId = r.MaternalOrganismId
	}
	if r.MaternalOrganismId != nil && (r.MaternalOrganismId.Id != nil || r.MaternalOrganismId.Extension != nil) {
		m.MaternalOrganismIdPrimitiveElement = &primitiveElement{Id: r.MaternalOrganismId.Id, Extension: r.MaternalOrganismId.Extension}
	}
	if r.MaternalOrganismName != nil && r.MaternalOrganismName.Value != nil {
		m.MaternalOrganismName = r.MaternalOrganismName
	}
	if r.MaternalOrganismName != nil && (r.MaternalOrganismName.Id != nil || r.MaternalOrganismName.Extension != nil) {
		m.MaternalOrganismNamePrimitiveElement = &primitiveElement{Id: r.MaternalOrganismName.Id, Extension: r.MaternalOrganismName.Extension}
	}
	if r.PaternalOrganismId != nil && r.PaternalOrganismId.Value != nil {
		m.PaternalOrganismId = r.PaternalOrganismId
	}
	if r.PaternalOrganismId != nil && (r.PaternalOrganismId.Id != nil || r.PaternalOrganismId.Extension != nil) {
		m.PaternalOrganismIdPrimitiveElement = &primitiveElement{Id: r.PaternalOrganismId.Id, Extension: r.PaternalOrganismId.Extension}
	}
	if r.PaternalOrganismName != nil && r.PaternalOrganismName.Value != nil {
		m.PaternalOrganismName = r.PaternalOrganismName
	}
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
		if r.MaternalOrganismId == nil {
			r.MaternalOrganismId = &String{}
		}
		r.MaternalOrganismId.Id = m.MaternalOrganismIdPrimitiveElement.Id
		r.MaternalOrganismId.Extension = m.MaternalOrganismIdPrimitiveElement.Extension
	}
	r.MaternalOrganismName = m.MaternalOrganismName
	if m.MaternalOrganismNamePrimitiveElement != nil {
		if r.MaternalOrganismName == nil {
			r.MaternalOrganismName = &String{}
		}
		r.MaternalOrganismName.Id = m.MaternalOrganismNamePrimitiveElement.Id
		r.MaternalOrganismName.Extension = m.MaternalOrganismNamePrimitiveElement.Extension
	}
	r.PaternalOrganismId = m.PaternalOrganismId
	if m.PaternalOrganismIdPrimitiveElement != nil {
		if r.PaternalOrganismId == nil {
			r.PaternalOrganismId = &String{}
		}
		r.PaternalOrganismId.Id = m.PaternalOrganismIdPrimitiveElement.Id
		r.PaternalOrganismId.Extension = m.PaternalOrganismIdPrimitiveElement.Extension
	}
	r.PaternalOrganismName = m.PaternalOrganismName
	if m.PaternalOrganismNamePrimitiveElement != nil {
		if r.PaternalOrganismName == nil {
			r.PaternalOrganismName = &String{}
		}
		r.PaternalOrganismName.Id = m.PaternalOrganismNamePrimitiveElement.Id
		r.PaternalOrganismName.Extension = m.PaternalOrganismNamePrimitiveElement.Extension
	}
	r.HybridType = m.HybridType
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
func (r SubstanceSourceMaterialOrganismHybrid) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
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
func (r SubstanceSourceMaterialOrganismOrganismGeneral) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
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
func (r SubstanceSourceMaterialPartDescription) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
