package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// The detailed description of a substance, typically at a level beyond what is used for prescribing.
type SubstanceSpecification struct {
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
	// Identifier by which this substance is known.
	Identifier *Identifier
	// High level categorization, e.g. polymer or nucleic acid.
	Type *CodeableConcept
	// Status of substance within the catalogue e.g. approved.
	Status *CodeableConcept
	// If the substance applies to only human or veterinary use.
	Domain *CodeableConcept
	// Textual description of the substance.
	Description *String
	// Supporting literature.
	Source []Reference
	// Textual comment about this record of a substance.
	Comment *String
	// Moiety, for structural modifications.
	Moiety []SubstanceSpecificationMoiety
	// General specifications for this substance, including how it is related to other substances.
	Property []SubstanceSpecificationProperty
	// General information detailing this substance.
	ReferenceInformation *Reference
	// Structural information.
	Structure *SubstanceSpecificationStructure
	// Codes associated with the substance.
	Code []SubstanceSpecificationCode
	// Names applicable to this substance.
	Name []SubstanceSpecificationName
	// The molecular weight or weight range (for proteins, polymers or nucleic acids).
	MolecularWeight []SubstanceSpecificationStructureIsotopeMolecularWeight
	// A link between this substance and another, with details of the relationship.
	Relationship []SubstanceSpecificationRelationship
	// Data items specific to nucleic acids.
	NucleicAcid *Reference
	// Data items specific to polymers.
	Polymer *Reference
	// Data items specific to proteins.
	Protein *Reference
	// Material or taxonomic/anatomical source for the substance.
	SourceMaterial *Reference
}

func (r SubstanceSpecification) ResourceType() string {
	return "SubstanceSpecification"
}
func (r SubstanceSpecification) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonSubstanceSpecification struct {
	ResourceType                  string                                                  `json:"resourceType"`
	Id                            *Id                                                     `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement                                       `json:"_id,omitempty"`
	Meta                          *Meta                                                   `json:"meta,omitempty"`
	ImplicitRules                 *Uri                                                    `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement                                       `json:"_implicitRules,omitempty"`
	Language                      *Code                                                   `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement                                       `json:"_language,omitempty"`
	Text                          *Narrative                                              `json:"text,omitempty"`
	Contained                     []ContainedResource                                     `json:"contained,omitempty"`
	Extension                     []Extension                                             `json:"extension,omitempty"`
	ModifierExtension             []Extension                                             `json:"modifierExtension,omitempty"`
	Identifier                    *Identifier                                             `json:"identifier,omitempty"`
	Type                          *CodeableConcept                                        `json:"type,omitempty"`
	Status                        *CodeableConcept                                        `json:"status,omitempty"`
	Domain                        *CodeableConcept                                        `json:"domain,omitempty"`
	Description                   *String                                                 `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement                                       `json:"_description,omitempty"`
	Source                        []Reference                                             `json:"source,omitempty"`
	Comment                       *String                                                 `json:"comment,omitempty"`
	CommentPrimitiveElement       *primitiveElement                                       `json:"_comment,omitempty"`
	Moiety                        []SubstanceSpecificationMoiety                          `json:"moiety,omitempty"`
	Property                      []SubstanceSpecificationProperty                        `json:"property,omitempty"`
	ReferenceInformation          *Reference                                              `json:"referenceInformation,omitempty"`
	Structure                     *SubstanceSpecificationStructure                        `json:"structure,omitempty"`
	Code                          []SubstanceSpecificationCode                            `json:"code,omitempty"`
	Name                          []SubstanceSpecificationName                            `json:"name,omitempty"`
	MolecularWeight               []SubstanceSpecificationStructureIsotopeMolecularWeight `json:"molecularWeight,omitempty"`
	Relationship                  []SubstanceSpecificationRelationship                    `json:"relationship,omitempty"`
	NucleicAcid                   *Reference                                              `json:"nucleicAcid,omitempty"`
	Polymer                       *Reference                                              `json:"polymer,omitempty"`
	Protein                       *Reference                                              `json:"protein,omitempty"`
	SourceMaterial                *Reference                                              `json:"sourceMaterial,omitempty"`
}

func (r SubstanceSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSpecification) marshalJSON() jsonSubstanceSpecification {
	m := jsonSubstanceSpecification{}
	m.ResourceType = "SubstanceSpecification"
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
	m.Status = r.Status
	m.Domain = r.Domain
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Source = r.Source
	if r.Comment != nil && r.Comment.Value != nil {
		m.Comment = r.Comment
	}
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	m.Moiety = r.Moiety
	m.Property = r.Property
	m.ReferenceInformation = r.ReferenceInformation
	m.Structure = r.Structure
	m.Code = r.Code
	m.Name = r.Name
	m.MolecularWeight = r.MolecularWeight
	m.Relationship = r.Relationship
	m.NucleicAcid = r.NucleicAcid
	m.Polymer = r.Polymer
	m.Protein = r.Protein
	m.SourceMaterial = r.SourceMaterial
	return m
}
func (r *SubstanceSpecification) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSpecification
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSpecification) unmarshalJSON(m jsonSubstanceSpecification) error {
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
	r.Status = m.Status
	r.Domain = m.Domain
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Source = m.Source
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		if r.Comment == nil {
			r.Comment = &String{}
		}
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	r.Moiety = m.Moiety
	r.Property = m.Property
	r.ReferenceInformation = m.ReferenceInformation
	r.Structure = m.Structure
	r.Code = m.Code
	r.Name = m.Name
	r.MolecularWeight = m.MolecularWeight
	r.Relationship = m.Relationship
	r.NucleicAcid = m.NucleicAcid
	r.Polymer = m.Polymer
	r.Protein = m.Protein
	r.SourceMaterial = m.SourceMaterial
	return nil
}
func (r SubstanceSpecification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "SubstanceSpecification"
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
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Domain, xml.StartElement{Name: xml.Name{Local: "domain"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Source, xml.StartElement{Name: xml.Name{Local: "source"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Comment, xml.StartElement{Name: xml.Name{Local: "comment"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Moiety, xml.StartElement{Name: xml.Name{Local: "moiety"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Property, xml.StartElement{Name: xml.Name{Local: "property"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReferenceInformation, xml.StartElement{Name: xml.Name{Local: "referenceInformation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Structure, xml.StartElement{Name: xml.Name{Local: "structure"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MolecularWeight, xml.StartElement{Name: xml.Name{Local: "molecularWeight"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Relationship, xml.StartElement{Name: xml.Name{Local: "relationship"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NucleicAcid, xml.StartElement{Name: xml.Name{Local: "nucleicAcid"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Polymer, xml.StartElement{Name: xml.Name{Local: "polymer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Protein, xml.StartElement{Name: xml.Name{Local: "protein"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SourceMaterial, xml.StartElement{Name: xml.Name{Local: "sourceMaterial"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceSpecification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.Identifier = &v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "status":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = &v
			case "domain":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Domain = &v
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "source":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Source = append(r.Source, v)
			case "comment":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Comment = &v
			case "moiety":
				var v SubstanceSpecificationMoiety
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Moiety = append(r.Moiety, v)
			case "property":
				var v SubstanceSpecificationProperty
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Property = append(r.Property, v)
			case "referenceInformation":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReferenceInformation = &v
			case "structure":
				var v SubstanceSpecificationStructure
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Structure = &v
			case "code":
				var v SubstanceSpecificationCode
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = append(r.Code, v)
			case "name":
				var v SubstanceSpecificationName
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = append(r.Name, v)
			case "molecularWeight":
				var v SubstanceSpecificationStructureIsotopeMolecularWeight
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MolecularWeight = append(r.MolecularWeight, v)
			case "relationship":
				var v SubstanceSpecificationRelationship
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Relationship = append(r.Relationship, v)
			case "nucleicAcid":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NucleicAcid = &v
			case "polymer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Polymer = &v
			case "protein":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Protein = &v
			case "sourceMaterial":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SourceMaterial = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceSpecification) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Moiety, for structural modifications.
type SubstanceSpecificationMoiety struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Role that the moiety is playing.
	Role *CodeableConcept
	// Identifier by which this moiety substance is known.
	Identifier *Identifier
	// Textual name for this moiety substance.
	Name *String
	// Stereochemistry type.
	Stereochemistry *CodeableConcept
	// Optical activity type.
	OpticalActivity *CodeableConcept
	// Molecular formula.
	MolecularFormula *String
	// Quantitative value for this moiety.
	Amount isSubstanceSpecificationMoietyAmount
}
type isSubstanceSpecificationMoietyAmount interface {
	isSubstanceSpecificationMoietyAmount()
}

func (r Quantity) isSubstanceSpecificationMoietyAmount() {}
func (r String) isSubstanceSpecificationMoietyAmount()   {}

type jsonSubstanceSpecificationMoiety struct {
	Id                               *string           `json:"id,omitempty"`
	Extension                        []Extension       `json:"extension,omitempty"`
	ModifierExtension                []Extension       `json:"modifierExtension,omitempty"`
	Role                             *CodeableConcept  `json:"role,omitempty"`
	Identifier                       *Identifier       `json:"identifier,omitempty"`
	Name                             *String           `json:"name,omitempty"`
	NamePrimitiveElement             *primitiveElement `json:"_name,omitempty"`
	Stereochemistry                  *CodeableConcept  `json:"stereochemistry,omitempty"`
	OpticalActivity                  *CodeableConcept  `json:"opticalActivity,omitempty"`
	MolecularFormula                 *String           `json:"molecularFormula,omitempty"`
	MolecularFormulaPrimitiveElement *primitiveElement `json:"_molecularFormula,omitempty"`
	AmountQuantity                   *Quantity         `json:"amountQuantity,omitempty"`
	AmountString                     *String           `json:"amountString,omitempty"`
	AmountStringPrimitiveElement     *primitiveElement `json:"_amountString,omitempty"`
}

func (r SubstanceSpecificationMoiety) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSpecificationMoiety) marshalJSON() jsonSubstanceSpecificationMoiety {
	m := jsonSubstanceSpecificationMoiety{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Role = r.Role
	m.Identifier = r.Identifier
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Stereochemistry = r.Stereochemistry
	m.OpticalActivity = r.OpticalActivity
	if r.MolecularFormula != nil && r.MolecularFormula.Value != nil {
		m.MolecularFormula = r.MolecularFormula
	}
	if r.MolecularFormula != nil && (r.MolecularFormula.Id != nil || r.MolecularFormula.Extension != nil) {
		m.MolecularFormulaPrimitiveElement = &primitiveElement{Id: r.MolecularFormula.Id, Extension: r.MolecularFormula.Extension}
	}
	switch v := r.Amount.(type) {
	case Quantity:
		m.AmountQuantity = &v
	case *Quantity:
		m.AmountQuantity = v
	case String:
		if v.Value != nil {
			m.AmountString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.AmountStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.AmountString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.AmountStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	return m
}
func (r *SubstanceSpecificationMoiety) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSpecificationMoiety
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSpecificationMoiety) unmarshalJSON(m jsonSubstanceSpecificationMoiety) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Role = m.Role
	r.Identifier = m.Identifier
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Stereochemistry = m.Stereochemistry
	r.OpticalActivity = m.OpticalActivity
	r.MolecularFormula = m.MolecularFormula
	if m.MolecularFormulaPrimitiveElement != nil {
		if r.MolecularFormula == nil {
			r.MolecularFormula = &String{}
		}
		r.MolecularFormula.Id = m.MolecularFormulaPrimitiveElement.Id
		r.MolecularFormula.Extension = m.MolecularFormulaPrimitiveElement.Extension
	}
	if m.AmountQuantity != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountQuantity
		r.Amount = v
	}
	if m.AmountString != nil || m.AmountStringPrimitiveElement != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountString
		if m.AmountStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.AmountStringPrimitiveElement.Id
			v.Extension = m.AmountStringPrimitiveElement.Extension
		}
		r.Amount = v
	}
	return nil
}
func (r SubstanceSpecificationMoiety) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Role, xml.StartElement{Name: xml.Name{Local: "role"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Stereochemistry, xml.StartElement{Name: xml.Name{Local: "stereochemistry"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OpticalActivity, xml.StartElement{Name: xml.Name{Local: "opticalActivity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MolecularFormula, xml.StartElement{Name: xml.Name{Local: "molecularFormula"}})
	if err != nil {
		return err
	}
	switch v := r.Amount.(type) {
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "amountQuantity"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "amountString"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceSpecificationMoiety) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "role":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Role = &v
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = &v
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "stereochemistry":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Stereochemistry = &v
			case "opticalActivity":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OpticalActivity = &v
			case "molecularFormula":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MolecularFormula = &v
			case "amountQuantity":
				if r.Amount != nil {
					return fmt.Errorf("multiple values for field \"Amount\"")
				}
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			case "amountString":
				if r.Amount != nil {
					return fmt.Errorf("multiple values for field \"Amount\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceSpecificationMoiety) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// General specifications for this substance, including how it is related to other substances.
type SubstanceSpecificationProperty struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A category for this property, e.g. Physical, Chemical, Enzymatic.
	Category *CodeableConcept
	// Property type e.g. viscosity, pH, isoelectric point.
	Code *CodeableConcept
	// Parameters that were used in the measurement of a property (e.g. for viscosity: measured at 20C with a pH of 7.1).
	Parameters *String
	// A substance upon which a defining property depends (e.g. for solubility: in water, in alcohol).
	DefiningSubstance isSubstanceSpecificationPropertyDefiningSubstance
	// Quantitative value for this property.
	Amount isSubstanceSpecificationPropertyAmount
}
type isSubstanceSpecificationPropertyDefiningSubstance interface {
	isSubstanceSpecificationPropertyDefiningSubstance()
}

func (r Reference) isSubstanceSpecificationPropertyDefiningSubstance()       {}
func (r CodeableConcept) isSubstanceSpecificationPropertyDefiningSubstance() {}

type isSubstanceSpecificationPropertyAmount interface {
	isSubstanceSpecificationPropertyAmount()
}

func (r Quantity) isSubstanceSpecificationPropertyAmount() {}
func (r String) isSubstanceSpecificationPropertyAmount()   {}

type jsonSubstanceSpecificationProperty struct {
	Id                               *string           `json:"id,omitempty"`
	Extension                        []Extension       `json:"extension,omitempty"`
	ModifierExtension                []Extension       `json:"modifierExtension,omitempty"`
	Category                         *CodeableConcept  `json:"category,omitempty"`
	Code                             *CodeableConcept  `json:"code,omitempty"`
	Parameters                       *String           `json:"parameters,omitempty"`
	ParametersPrimitiveElement       *primitiveElement `json:"_parameters,omitempty"`
	DefiningSubstanceReference       *Reference        `json:"definingSubstanceReference,omitempty"`
	DefiningSubstanceCodeableConcept *CodeableConcept  `json:"definingSubstanceCodeableConcept,omitempty"`
	AmountQuantity                   *Quantity         `json:"amountQuantity,omitempty"`
	AmountString                     *String           `json:"amountString,omitempty"`
	AmountStringPrimitiveElement     *primitiveElement `json:"_amountString,omitempty"`
}

func (r SubstanceSpecificationProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSpecificationProperty) marshalJSON() jsonSubstanceSpecificationProperty {
	m := jsonSubstanceSpecificationProperty{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Category = r.Category
	m.Code = r.Code
	if r.Parameters != nil && r.Parameters.Value != nil {
		m.Parameters = r.Parameters
	}
	if r.Parameters != nil && (r.Parameters.Id != nil || r.Parameters.Extension != nil) {
		m.ParametersPrimitiveElement = &primitiveElement{Id: r.Parameters.Id, Extension: r.Parameters.Extension}
	}
	switch v := r.DefiningSubstance.(type) {
	case Reference:
		m.DefiningSubstanceReference = &v
	case *Reference:
		m.DefiningSubstanceReference = v
	case CodeableConcept:
		m.DefiningSubstanceCodeableConcept = &v
	case *CodeableConcept:
		m.DefiningSubstanceCodeableConcept = v
	}
	switch v := r.Amount.(type) {
	case Quantity:
		m.AmountQuantity = &v
	case *Quantity:
		m.AmountQuantity = v
	case String:
		if v.Value != nil {
			m.AmountString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.AmountStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.AmountString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.AmountStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	return m
}
func (r *SubstanceSpecificationProperty) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSpecificationProperty
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSpecificationProperty) unmarshalJSON(m jsonSubstanceSpecificationProperty) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Category = m.Category
	r.Code = m.Code
	r.Parameters = m.Parameters
	if m.ParametersPrimitiveElement != nil {
		if r.Parameters == nil {
			r.Parameters = &String{}
		}
		r.Parameters.Id = m.ParametersPrimitiveElement.Id
		r.Parameters.Extension = m.ParametersPrimitiveElement.Extension
	}
	if m.DefiningSubstanceReference != nil {
		if r.DefiningSubstance != nil {
			return fmt.Errorf("multiple values for field \"DefiningSubstance\"")
		}
		v := m.DefiningSubstanceReference
		r.DefiningSubstance = v
	}
	if m.DefiningSubstanceCodeableConcept != nil {
		if r.DefiningSubstance != nil {
			return fmt.Errorf("multiple values for field \"DefiningSubstance\"")
		}
		v := m.DefiningSubstanceCodeableConcept
		r.DefiningSubstance = v
	}
	if m.AmountQuantity != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountQuantity
		r.Amount = v
	}
	if m.AmountString != nil || m.AmountStringPrimitiveElement != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountString
		if m.AmountStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.AmountStringPrimitiveElement.Id
			v.Extension = m.AmountStringPrimitiveElement.Extension
		}
		r.Amount = v
	}
	return nil
}
func (r SubstanceSpecificationProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Category, xml.StartElement{Name: xml.Name{Local: "category"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Parameters, xml.StartElement{Name: xml.Name{Local: "parameters"}})
	if err != nil {
		return err
	}
	switch v := r.DefiningSubstance.(type) {
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "definingSubstanceReference"}})
		if err != nil {
			return err
		}
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "definingSubstanceCodeableConcept"}})
		if err != nil {
			return err
		}
	}
	switch v := r.Amount.(type) {
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "amountQuantity"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "amountString"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceSpecificationProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "category":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = &v
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = &v
			case "parameters":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Parameters = &v
			case "definingSubstanceReference":
				if r.DefiningSubstance != nil {
					return fmt.Errorf("multiple values for field \"DefiningSubstance\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DefiningSubstance = v
			case "definingSubstanceCodeableConcept":
				if r.DefiningSubstance != nil {
					return fmt.Errorf("multiple values for field \"DefiningSubstance\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DefiningSubstance = v
			case "amountQuantity":
				if r.Amount != nil {
					return fmt.Errorf("multiple values for field \"Amount\"")
				}
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			case "amountString":
				if r.Amount != nil {
					return fmt.Errorf("multiple values for field \"Amount\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceSpecificationProperty) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Structural information.
type SubstanceSpecificationStructure struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Stereochemistry type.
	Stereochemistry *CodeableConcept
	// Optical activity type.
	OpticalActivity *CodeableConcept
	// Molecular formula.
	MolecularFormula *String
	// Specified per moiety according to the Hill system, i.e. first C, then H, then alphabetical, each moiety separated by a dot.
	MolecularFormulaByMoiety *String
	// Applicable for single substances that contain a radionuclide or a non-natural isotopic ratio.
	Isotope []SubstanceSpecificationStructureIsotope
	// The molecular weight or weight range (for proteins, polymers or nucleic acids).
	MolecularWeight *SubstanceSpecificationStructureIsotopeMolecularWeight
	// Supporting literature.
	Source []Reference
	// Molecular structural representation.
	Representation []SubstanceSpecificationStructureRepresentation
}
type jsonSubstanceSpecificationStructure struct {
	Id                                       *string                                                `json:"id,omitempty"`
	Extension                                []Extension                                            `json:"extension,omitempty"`
	ModifierExtension                        []Extension                                            `json:"modifierExtension,omitempty"`
	Stereochemistry                          *CodeableConcept                                       `json:"stereochemistry,omitempty"`
	OpticalActivity                          *CodeableConcept                                       `json:"opticalActivity,omitempty"`
	MolecularFormula                         *String                                                `json:"molecularFormula,omitempty"`
	MolecularFormulaPrimitiveElement         *primitiveElement                                      `json:"_molecularFormula,omitempty"`
	MolecularFormulaByMoiety                 *String                                                `json:"molecularFormulaByMoiety,omitempty"`
	MolecularFormulaByMoietyPrimitiveElement *primitiveElement                                      `json:"_molecularFormulaByMoiety,omitempty"`
	Isotope                                  []SubstanceSpecificationStructureIsotope               `json:"isotope,omitempty"`
	MolecularWeight                          *SubstanceSpecificationStructureIsotopeMolecularWeight `json:"molecularWeight,omitempty"`
	Source                                   []Reference                                            `json:"source,omitempty"`
	Representation                           []SubstanceSpecificationStructureRepresentation        `json:"representation,omitempty"`
}

func (r SubstanceSpecificationStructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSpecificationStructure) marshalJSON() jsonSubstanceSpecificationStructure {
	m := jsonSubstanceSpecificationStructure{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Stereochemistry = r.Stereochemistry
	m.OpticalActivity = r.OpticalActivity
	if r.MolecularFormula != nil && r.MolecularFormula.Value != nil {
		m.MolecularFormula = r.MolecularFormula
	}
	if r.MolecularFormula != nil && (r.MolecularFormula.Id != nil || r.MolecularFormula.Extension != nil) {
		m.MolecularFormulaPrimitiveElement = &primitiveElement{Id: r.MolecularFormula.Id, Extension: r.MolecularFormula.Extension}
	}
	if r.MolecularFormulaByMoiety != nil && r.MolecularFormulaByMoiety.Value != nil {
		m.MolecularFormulaByMoiety = r.MolecularFormulaByMoiety
	}
	if r.MolecularFormulaByMoiety != nil && (r.MolecularFormulaByMoiety.Id != nil || r.MolecularFormulaByMoiety.Extension != nil) {
		m.MolecularFormulaByMoietyPrimitiveElement = &primitiveElement{Id: r.MolecularFormulaByMoiety.Id, Extension: r.MolecularFormulaByMoiety.Extension}
	}
	m.Isotope = r.Isotope
	m.MolecularWeight = r.MolecularWeight
	m.Source = r.Source
	m.Representation = r.Representation
	return m
}
func (r *SubstanceSpecificationStructure) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSpecificationStructure
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSpecificationStructure) unmarshalJSON(m jsonSubstanceSpecificationStructure) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Stereochemistry = m.Stereochemistry
	r.OpticalActivity = m.OpticalActivity
	r.MolecularFormula = m.MolecularFormula
	if m.MolecularFormulaPrimitiveElement != nil {
		if r.MolecularFormula == nil {
			r.MolecularFormula = &String{}
		}
		r.MolecularFormula.Id = m.MolecularFormulaPrimitiveElement.Id
		r.MolecularFormula.Extension = m.MolecularFormulaPrimitiveElement.Extension
	}
	r.MolecularFormulaByMoiety = m.MolecularFormulaByMoiety
	if m.MolecularFormulaByMoietyPrimitiveElement != nil {
		if r.MolecularFormulaByMoiety == nil {
			r.MolecularFormulaByMoiety = &String{}
		}
		r.MolecularFormulaByMoiety.Id = m.MolecularFormulaByMoietyPrimitiveElement.Id
		r.MolecularFormulaByMoiety.Extension = m.MolecularFormulaByMoietyPrimitiveElement.Extension
	}
	r.Isotope = m.Isotope
	r.MolecularWeight = m.MolecularWeight
	r.Source = m.Source
	r.Representation = m.Representation
	return nil
}
func (r SubstanceSpecificationStructure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Stereochemistry, xml.StartElement{Name: xml.Name{Local: "stereochemistry"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OpticalActivity, xml.StartElement{Name: xml.Name{Local: "opticalActivity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MolecularFormula, xml.StartElement{Name: xml.Name{Local: "molecularFormula"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MolecularFormulaByMoiety, xml.StartElement{Name: xml.Name{Local: "molecularFormulaByMoiety"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Isotope, xml.StartElement{Name: xml.Name{Local: "isotope"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MolecularWeight, xml.StartElement{Name: xml.Name{Local: "molecularWeight"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Source, xml.StartElement{Name: xml.Name{Local: "source"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Representation, xml.StartElement{Name: xml.Name{Local: "representation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceSpecificationStructure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "stereochemistry":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Stereochemistry = &v
			case "opticalActivity":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OpticalActivity = &v
			case "molecularFormula":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MolecularFormula = &v
			case "molecularFormulaByMoiety":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MolecularFormulaByMoiety = &v
			case "isotope":
				var v SubstanceSpecificationStructureIsotope
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Isotope = append(r.Isotope, v)
			case "molecularWeight":
				var v SubstanceSpecificationStructureIsotopeMolecularWeight
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MolecularWeight = &v
			case "source":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Source = append(r.Source, v)
			case "representation":
				var v SubstanceSpecificationStructureRepresentation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Representation = append(r.Representation, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceSpecificationStructure) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Applicable for single substances that contain a radionuclide or a non-natural isotopic ratio.
type SubstanceSpecificationStructureIsotope struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Substance identifier for each non-natural or radioisotope.
	Identifier *Identifier
	// Substance name for each non-natural or radioisotope.
	Name *CodeableConcept
	// The type of isotopic substitution present in a single substance.
	Substitution *CodeableConcept
	// Half life - for a non-natural nuclide.
	HalfLife *Quantity
	// The molecular weight or weight range (for proteins, polymers or nucleic acids).
	MolecularWeight *SubstanceSpecificationStructureIsotopeMolecularWeight
}
type jsonSubstanceSpecificationStructureIsotope struct {
	Id                *string                                                `json:"id,omitempty"`
	Extension         []Extension                                            `json:"extension,omitempty"`
	ModifierExtension []Extension                                            `json:"modifierExtension,omitempty"`
	Identifier        *Identifier                                            `json:"identifier,omitempty"`
	Name              *CodeableConcept                                       `json:"name,omitempty"`
	Substitution      *CodeableConcept                                       `json:"substitution,omitempty"`
	HalfLife          *Quantity                                              `json:"halfLife,omitempty"`
	MolecularWeight   *SubstanceSpecificationStructureIsotopeMolecularWeight `json:"molecularWeight,omitempty"`
}

func (r SubstanceSpecificationStructureIsotope) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSpecificationStructureIsotope) marshalJSON() jsonSubstanceSpecificationStructureIsotope {
	m := jsonSubstanceSpecificationStructureIsotope{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Name = r.Name
	m.Substitution = r.Substitution
	m.HalfLife = r.HalfLife
	m.MolecularWeight = r.MolecularWeight
	return m
}
func (r *SubstanceSpecificationStructureIsotope) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSpecificationStructureIsotope
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSpecificationStructureIsotope) unmarshalJSON(m jsonSubstanceSpecificationStructureIsotope) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Name = m.Name
	r.Substitution = m.Substitution
	r.HalfLife = m.HalfLife
	r.MolecularWeight = m.MolecularWeight
	return nil
}
func (r SubstanceSpecificationStructureIsotope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Substitution, xml.StartElement{Name: xml.Name{Local: "substitution"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.HalfLife, xml.StartElement{Name: xml.Name{Local: "halfLife"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MolecularWeight, xml.StartElement{Name: xml.Name{Local: "molecularWeight"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceSpecificationStructureIsotope) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = &v
			case "name":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "substitution":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Substitution = &v
			case "halfLife":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.HalfLife = &v
			case "molecularWeight":
				var v SubstanceSpecificationStructureIsotopeMolecularWeight
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MolecularWeight = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceSpecificationStructureIsotope) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The molecular weight or weight range (for proteins, polymers or nucleic acids).
type SubstanceSpecificationStructureIsotopeMolecularWeight struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The method by which the molecular weight was determined.
	Method *CodeableConcept
	// Type of molecular weight such as exact, average (also known as. number average), weight average.
	Type *CodeableConcept
	// Used to capture quantitative values for a variety of elements. If only limits are given, the arithmetic mean would be the average. If only a single definite value for a given element is given, it would be captured in this field.
	Amount *Quantity
}
type jsonSubstanceSpecificationStructureIsotopeMolecularWeight struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Method            *CodeableConcept `json:"method,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Amount            *Quantity        `json:"amount,omitempty"`
}

func (r SubstanceSpecificationStructureIsotopeMolecularWeight) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSpecificationStructureIsotopeMolecularWeight) marshalJSON() jsonSubstanceSpecificationStructureIsotopeMolecularWeight {
	m := jsonSubstanceSpecificationStructureIsotopeMolecularWeight{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Method = r.Method
	m.Type = r.Type
	m.Amount = r.Amount
	return m
}
func (r *SubstanceSpecificationStructureIsotopeMolecularWeight) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSpecificationStructureIsotopeMolecularWeight
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSpecificationStructureIsotopeMolecularWeight) unmarshalJSON(m jsonSubstanceSpecificationStructureIsotopeMolecularWeight) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Method = m.Method
	r.Type = m.Type
	r.Amount = m.Amount
	return nil
}
func (r SubstanceSpecificationStructureIsotopeMolecularWeight) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Method, xml.StartElement{Name: xml.Name{Local: "method"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Amount, xml.StartElement{Name: xml.Name{Local: "amount"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceSpecificationStructureIsotopeMolecularWeight) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "method":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Method = &v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "amount":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceSpecificationStructureIsotopeMolecularWeight) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Molecular structural representation.
type SubstanceSpecificationStructureRepresentation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of structure (e.g. Full, Partial, Representative).
	Type *CodeableConcept
	// The structural representation as text string in a format e.g. InChI, SMILES, MOLFILE, CDX.
	Representation *String
	// An attached file with the structural representation.
	Attachment *Attachment
}
type jsonSubstanceSpecificationStructureRepresentation struct {
	Id                             *string           `json:"id,omitempty"`
	Extension                      []Extension       `json:"extension,omitempty"`
	ModifierExtension              []Extension       `json:"modifierExtension,omitempty"`
	Type                           *CodeableConcept  `json:"type,omitempty"`
	Representation                 *String           `json:"representation,omitempty"`
	RepresentationPrimitiveElement *primitiveElement `json:"_representation,omitempty"`
	Attachment                     *Attachment       `json:"attachment,omitempty"`
}

func (r SubstanceSpecificationStructureRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSpecificationStructureRepresentation) marshalJSON() jsonSubstanceSpecificationStructureRepresentation {
	m := jsonSubstanceSpecificationStructureRepresentation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Representation != nil && r.Representation.Value != nil {
		m.Representation = r.Representation
	}
	if r.Representation != nil && (r.Representation.Id != nil || r.Representation.Extension != nil) {
		m.RepresentationPrimitiveElement = &primitiveElement{Id: r.Representation.Id, Extension: r.Representation.Extension}
	}
	m.Attachment = r.Attachment
	return m
}
func (r *SubstanceSpecificationStructureRepresentation) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSpecificationStructureRepresentation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSpecificationStructureRepresentation) unmarshalJSON(m jsonSubstanceSpecificationStructureRepresentation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Representation = m.Representation
	if m.RepresentationPrimitiveElement != nil {
		if r.Representation == nil {
			r.Representation = &String{}
		}
		r.Representation.Id = m.RepresentationPrimitiveElement.Id
		r.Representation.Extension = m.RepresentationPrimitiveElement.Extension
	}
	r.Attachment = m.Attachment
	return nil
}
func (r SubstanceSpecificationStructureRepresentation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Representation, xml.StartElement{Name: xml.Name{Local: "representation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Attachment, xml.StartElement{Name: xml.Name{Local: "attachment"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceSpecificationStructureRepresentation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "representation":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Representation = &v
			case "attachment":
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Attachment = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceSpecificationStructureRepresentation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Codes associated with the substance.
type SubstanceSpecificationCode struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The specific code.
	Code *CodeableConcept
	// Status of the code assignment.
	Status *CodeableConcept
	// The date at which the code status is changed as part of the terminology maintenance.
	StatusDate *DateTime
	// Any comment can be provided in this field, if necessary.
	Comment *String
	// Supporting literature.
	Source []Reference
}
type jsonSubstanceSpecificationCode struct {
	Id                         *string           `json:"id,omitempty"`
	Extension                  []Extension       `json:"extension,omitempty"`
	ModifierExtension          []Extension       `json:"modifierExtension,omitempty"`
	Code                       *CodeableConcept  `json:"code,omitempty"`
	Status                     *CodeableConcept  `json:"status,omitempty"`
	StatusDate                 *DateTime         `json:"statusDate,omitempty"`
	StatusDatePrimitiveElement *primitiveElement `json:"_statusDate,omitempty"`
	Comment                    *String           `json:"comment,omitempty"`
	CommentPrimitiveElement    *primitiveElement `json:"_comment,omitempty"`
	Source                     []Reference       `json:"source,omitempty"`
}

func (r SubstanceSpecificationCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSpecificationCode) marshalJSON() jsonSubstanceSpecificationCode {
	m := jsonSubstanceSpecificationCode{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Status = r.Status
	if r.StatusDate != nil && r.StatusDate.Value != nil {
		m.StatusDate = r.StatusDate
	}
	if r.StatusDate != nil && (r.StatusDate.Id != nil || r.StatusDate.Extension != nil) {
		m.StatusDatePrimitiveElement = &primitiveElement{Id: r.StatusDate.Id, Extension: r.StatusDate.Extension}
	}
	if r.Comment != nil && r.Comment.Value != nil {
		m.Comment = r.Comment
	}
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	m.Source = r.Source
	return m
}
func (r *SubstanceSpecificationCode) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSpecificationCode
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSpecificationCode) unmarshalJSON(m jsonSubstanceSpecificationCode) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Status = m.Status
	r.StatusDate = m.StatusDate
	if m.StatusDatePrimitiveElement != nil {
		if r.StatusDate == nil {
			r.StatusDate = &DateTime{}
		}
		r.StatusDate.Id = m.StatusDatePrimitiveElement.Id
		r.StatusDate.Extension = m.StatusDatePrimitiveElement.Extension
	}
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		if r.Comment == nil {
			r.Comment = &String{}
		}
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	r.Source = m.Source
	return nil
}
func (r SubstanceSpecificationCode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.StatusDate, xml.StartElement{Name: xml.Name{Local: "statusDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Comment, xml.StartElement{Name: xml.Name{Local: "comment"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Source, xml.StartElement{Name: xml.Name{Local: "source"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceSpecificationCode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = &v
			case "status":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = &v
			case "statusDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StatusDate = &v
			case "comment":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Comment = &v
			case "source":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Source = append(r.Source, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceSpecificationCode) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Names applicable to this substance.
type SubstanceSpecificationName struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The actual name.
	Name String
	// Name type.
	Type *CodeableConcept
	// The status of the name.
	Status *CodeableConcept
	// If this is the preferred name for this substance.
	Preferred *Boolean
	// Language of the name.
	Language []CodeableConcept
	// The use context of this name for example if there is a different name a drug active ingredient as opposed to a food colour additive.
	Domain []CodeableConcept
	// The jurisdiction where this name applies.
	Jurisdiction []CodeableConcept
	// A synonym of this name.
	Synonym []SubstanceSpecificationName
	// A translation for this name.
	Translation []SubstanceSpecificationName
	// Details of the official nature of this name.
	Official []SubstanceSpecificationNameOfficial
	// Supporting literature.
	Source []Reference
}
type jsonSubstanceSpecificationName struct {
	Id                        *string                              `json:"id,omitempty"`
	Extension                 []Extension                          `json:"extension,omitempty"`
	ModifierExtension         []Extension                          `json:"modifierExtension,omitempty"`
	Name                      String                               `json:"name,omitempty"`
	NamePrimitiveElement      *primitiveElement                    `json:"_name,omitempty"`
	Type                      *CodeableConcept                     `json:"type,omitempty"`
	Status                    *CodeableConcept                     `json:"status,omitempty"`
	Preferred                 *Boolean                             `json:"preferred,omitempty"`
	PreferredPrimitiveElement *primitiveElement                    `json:"_preferred,omitempty"`
	Language                  []CodeableConcept                    `json:"language,omitempty"`
	Domain                    []CodeableConcept                    `json:"domain,omitempty"`
	Jurisdiction              []CodeableConcept                    `json:"jurisdiction,omitempty"`
	Synonym                   []SubstanceSpecificationName         `json:"synonym,omitempty"`
	Translation               []SubstanceSpecificationName         `json:"translation,omitempty"`
	Official                  []SubstanceSpecificationNameOfficial `json:"official,omitempty"`
	Source                    []Reference                          `json:"source,omitempty"`
}

func (r SubstanceSpecificationName) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSpecificationName) marshalJSON() jsonSubstanceSpecificationName {
	m := jsonSubstanceSpecificationName{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Type = r.Type
	m.Status = r.Status
	if r.Preferred != nil && r.Preferred.Value != nil {
		m.Preferred = r.Preferred
	}
	if r.Preferred != nil && (r.Preferred.Id != nil || r.Preferred.Extension != nil) {
		m.PreferredPrimitiveElement = &primitiveElement{Id: r.Preferred.Id, Extension: r.Preferred.Extension}
	}
	m.Language = r.Language
	m.Domain = r.Domain
	m.Jurisdiction = r.Jurisdiction
	m.Synonym = r.Synonym
	m.Translation = r.Translation
	m.Official = r.Official
	m.Source = r.Source
	return m
}
func (r *SubstanceSpecificationName) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSpecificationName
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSpecificationName) unmarshalJSON(m jsonSubstanceSpecificationName) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Status = m.Status
	r.Preferred = m.Preferred
	if m.PreferredPrimitiveElement != nil {
		if r.Preferred == nil {
			r.Preferred = &Boolean{}
		}
		r.Preferred.Id = m.PreferredPrimitiveElement.Id
		r.Preferred.Extension = m.PreferredPrimitiveElement.Extension
	}
	r.Language = m.Language
	r.Domain = m.Domain
	r.Jurisdiction = m.Jurisdiction
	r.Synonym = m.Synonym
	r.Translation = m.Translation
	r.Official = m.Official
	r.Source = m.Source
	return nil
}
func (r SubstanceSpecificationName) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Preferred, xml.StartElement{Name: xml.Name{Local: "preferred"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Domain, xml.StartElement{Name: xml.Name{Local: "domain"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Jurisdiction, xml.StartElement{Name: xml.Name{Local: "jurisdiction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Synonym, xml.StartElement{Name: xml.Name{Local: "synonym"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Translation, xml.StartElement{Name: xml.Name{Local: "translation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Official, xml.StartElement{Name: xml.Name{Local: "official"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Source, xml.StartElement{Name: xml.Name{Local: "source"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceSpecificationName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "status":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = &v
			case "preferred":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Preferred = &v
			case "language":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = append(r.Language, v)
			case "domain":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Domain = append(r.Domain, v)
			case "jurisdiction":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Jurisdiction = append(r.Jurisdiction, v)
			case "synonym":
				var v SubstanceSpecificationName
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Synonym = append(r.Synonym, v)
			case "translation":
				var v SubstanceSpecificationName
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Translation = append(r.Translation, v)
			case "official":
				var v SubstanceSpecificationNameOfficial
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Official = append(r.Official, v)
			case "source":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Source = append(r.Source, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceSpecificationName) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Details of the official nature of this name.
type SubstanceSpecificationNameOfficial struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Which authority uses this official name.
	Authority *CodeableConcept
	// The status of the official name.
	Status *CodeableConcept
	// Date of official name change.
	Date *DateTime
}
type jsonSubstanceSpecificationNameOfficial struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Authority            *CodeableConcept  `json:"authority,omitempty"`
	Status               *CodeableConcept  `json:"status,omitempty"`
	Date                 *DateTime         `json:"date,omitempty"`
	DatePrimitiveElement *primitiveElement `json:"_date,omitempty"`
}

func (r SubstanceSpecificationNameOfficial) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSpecificationNameOfficial) marshalJSON() jsonSubstanceSpecificationNameOfficial {
	m := jsonSubstanceSpecificationNameOfficial{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Authority = r.Authority
	m.Status = r.Status
	if r.Date != nil && r.Date.Value != nil {
		m.Date = r.Date
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	return m
}
func (r *SubstanceSpecificationNameOfficial) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSpecificationNameOfficial
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSpecificationNameOfficial) unmarshalJSON(m jsonSubstanceSpecificationNameOfficial) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Authority = m.Authority
	r.Status = m.Status
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		if r.Date == nil {
			r.Date = &DateTime{}
		}
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	return nil
}
func (r SubstanceSpecificationNameOfficial) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Authority, xml.StartElement{Name: xml.Name{Local: "authority"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceSpecificationNameOfficial) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "authority":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Authority = &v
			case "status":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = &v
			case "date":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceSpecificationNameOfficial) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A link between this substance and another, with details of the relationship.
type SubstanceSpecificationRelationship struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A pointer to another substance, as a resource or just a representational code.
	Substance isSubstanceSpecificationRelationshipSubstance
	// For example "salt to parent", "active moiety", "starting material".
	Relationship *CodeableConcept
	// For example where an enzyme strongly bonds with a particular substance, this is a defining relationship for that enzyme, out of several possible substance relationships.
	IsDefining *Boolean
	// A numeric factor for the relationship, for instance to express that the salt of a substance has some percentage of the active substance in relation to some other.
	Amount isSubstanceSpecificationRelationshipAmount
	// For use when the numeric.
	AmountRatioLowLimit *Ratio
	// An operator for the amount, for example "average", "approximately", "less than".
	AmountType *CodeableConcept
	// Supporting literature.
	Source []Reference
}
type isSubstanceSpecificationRelationshipSubstance interface {
	isSubstanceSpecificationRelationshipSubstance()
}

func (r Reference) isSubstanceSpecificationRelationshipSubstance()       {}
func (r CodeableConcept) isSubstanceSpecificationRelationshipSubstance() {}

type isSubstanceSpecificationRelationshipAmount interface {
	isSubstanceSpecificationRelationshipAmount()
}

func (r Quantity) isSubstanceSpecificationRelationshipAmount() {}
func (r Range) isSubstanceSpecificationRelationshipAmount()    {}
func (r Ratio) isSubstanceSpecificationRelationshipAmount()    {}
func (r String) isSubstanceSpecificationRelationshipAmount()   {}

type jsonSubstanceSpecificationRelationship struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	SubstanceReference           *Reference        `json:"substanceReference,omitempty"`
	SubstanceCodeableConcept     *CodeableConcept  `json:"substanceCodeableConcept,omitempty"`
	Relationship                 *CodeableConcept  `json:"relationship,omitempty"`
	IsDefining                   *Boolean          `json:"isDefining,omitempty"`
	IsDefiningPrimitiveElement   *primitiveElement `json:"_isDefining,omitempty"`
	AmountQuantity               *Quantity         `json:"amountQuantity,omitempty"`
	AmountRange                  *Range            `json:"amountRange,omitempty"`
	AmountRatio                  *Ratio            `json:"amountRatio,omitempty"`
	AmountString                 *String           `json:"amountString,omitempty"`
	AmountStringPrimitiveElement *primitiveElement `json:"_amountString,omitempty"`
	AmountRatioLowLimit          *Ratio            `json:"amountRatioLowLimit,omitempty"`
	AmountType                   *CodeableConcept  `json:"amountType,omitempty"`
	Source                       []Reference       `json:"source,omitempty"`
}

func (r SubstanceSpecificationRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceSpecificationRelationship) marshalJSON() jsonSubstanceSpecificationRelationship {
	m := jsonSubstanceSpecificationRelationship{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Substance.(type) {
	case Reference:
		m.SubstanceReference = &v
	case *Reference:
		m.SubstanceReference = v
	case CodeableConcept:
		m.SubstanceCodeableConcept = &v
	case *CodeableConcept:
		m.SubstanceCodeableConcept = v
	}
	m.Relationship = r.Relationship
	if r.IsDefining != nil && r.IsDefining.Value != nil {
		m.IsDefining = r.IsDefining
	}
	if r.IsDefining != nil && (r.IsDefining.Id != nil || r.IsDefining.Extension != nil) {
		m.IsDefiningPrimitiveElement = &primitiveElement{Id: r.IsDefining.Id, Extension: r.IsDefining.Extension}
	}
	switch v := r.Amount.(type) {
	case Quantity:
		m.AmountQuantity = &v
	case *Quantity:
		m.AmountQuantity = v
	case Range:
		m.AmountRange = &v
	case *Range:
		m.AmountRange = v
	case Ratio:
		m.AmountRatio = &v
	case *Ratio:
		m.AmountRatio = v
	case String:
		if v.Value != nil {
			m.AmountString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.AmountStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.AmountString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.AmountStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.AmountRatioLowLimit = r.AmountRatioLowLimit
	m.AmountType = r.AmountType
	m.Source = r.Source
	return m
}
func (r *SubstanceSpecificationRelationship) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceSpecificationRelationship
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceSpecificationRelationship) unmarshalJSON(m jsonSubstanceSpecificationRelationship) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.SubstanceReference != nil {
		if r.Substance != nil {
			return fmt.Errorf("multiple values for field \"Substance\"")
		}
		v := m.SubstanceReference
		r.Substance = v
	}
	if m.SubstanceCodeableConcept != nil {
		if r.Substance != nil {
			return fmt.Errorf("multiple values for field \"Substance\"")
		}
		v := m.SubstanceCodeableConcept
		r.Substance = v
	}
	r.Relationship = m.Relationship
	r.IsDefining = m.IsDefining
	if m.IsDefiningPrimitiveElement != nil {
		if r.IsDefining == nil {
			r.IsDefining = &Boolean{}
		}
		r.IsDefining.Id = m.IsDefiningPrimitiveElement.Id
		r.IsDefining.Extension = m.IsDefiningPrimitiveElement.Extension
	}
	if m.AmountQuantity != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountQuantity
		r.Amount = v
	}
	if m.AmountRange != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountRange
		r.Amount = v
	}
	if m.AmountRatio != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountRatio
		r.Amount = v
	}
	if m.AmountString != nil || m.AmountStringPrimitiveElement != nil {
		if r.Amount != nil {
			return fmt.Errorf("multiple values for field \"Amount\"")
		}
		v := m.AmountString
		if m.AmountStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.AmountStringPrimitiveElement.Id
			v.Extension = m.AmountStringPrimitiveElement.Extension
		}
		r.Amount = v
	}
	r.AmountRatioLowLimit = m.AmountRatioLowLimit
	r.AmountType = m.AmountType
	r.Source = m.Source
	return nil
}
func (r SubstanceSpecificationRelationship) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	switch v := r.Substance.(type) {
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "substanceReference"}})
		if err != nil {
			return err
		}
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "substanceCodeableConcept"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Relationship, xml.StartElement{Name: xml.Name{Local: "relationship"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IsDefining, xml.StartElement{Name: xml.Name{Local: "isDefining"}})
	if err != nil {
		return err
	}
	switch v := r.Amount.(type) {
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "amountQuantity"}})
		if err != nil {
			return err
		}
	case Range, *Range:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "amountRange"}})
		if err != nil {
			return err
		}
	case Ratio, *Ratio:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "amountRatio"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "amountString"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.AmountRatioLowLimit, xml.StartElement{Name: xml.Name{Local: "amountRatioLowLimit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AmountType, xml.StartElement{Name: xml.Name{Local: "amountType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Source, xml.StartElement{Name: xml.Name{Local: "source"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceSpecificationRelationship) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "substanceReference":
				if r.Substance != nil {
					return fmt.Errorf("multiple values for field \"Substance\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Substance = v
			case "substanceCodeableConcept":
				if r.Substance != nil {
					return fmt.Errorf("multiple values for field \"Substance\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Substance = v
			case "relationship":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Relationship = &v
			case "isDefining":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IsDefining = &v
			case "amountQuantity":
				if r.Amount != nil {
					return fmt.Errorf("multiple values for field \"Amount\"")
				}
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			case "amountRange":
				if r.Amount != nil {
					return fmt.Errorf("multiple values for field \"Amount\"")
				}
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			case "amountRatio":
				if r.Amount != nil {
					return fmt.Errorf("multiple values for field \"Amount\"")
				}
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			case "amountString":
				if r.Amount != nil {
					return fmt.Errorf("multiple values for field \"Amount\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			case "amountRatioLowLimit":
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AmountRatioLowLimit = &v
			case "amountType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AmountType = &v
			case "source":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Source = append(r.Source, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceSpecificationRelationship) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
