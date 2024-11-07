package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// Nucleic acids are defined by three distinct elements: the base, sugar and linkage. Individual substance/moiety IDs will be created for each of these elements. The nucleotide sequence will be always entered in the 5’-3’ direction.
type SubstanceNucleicAcid struct {
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
	// The type of the sequence shall be specified based on a controlled vocabulary.
	SequenceType *CodeableConcept
	// The number of linear sequences of nucleotides linked through phosphodiester bonds shall be described. Subunits would be strands of nucleic acids that are tightly associated typically through Watson-Crick base pairing. NOTE: If not specified in the reference source, the assumption is that there is 1 subunit.
	NumberOfSubunits *Integer
	// The area of hybridisation shall be described if applicable for double stranded RNA or DNA. The number associated with the subunit followed by the number associated to the residue shall be specified in increasing order. The underscore “” shall be used as separator as follows: “Subunitnumber Residue”.
	AreaOfHybridisation *String
	// (TBC).
	OligoNucleotideType *CodeableConcept
	// Subunits are listed in order of decreasing length; sequences of the same length will be ordered by molecular weight; subunits that have identical sequences will be repeated multiple times.
	Subunit []SubstanceNucleicAcidSubunit
}

func (r SubstanceNucleicAcid) ResourceType() string {
	return "SubstanceNucleicAcid"
}
func (r SubstanceNucleicAcid) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonSubstanceNucleicAcid struct {
	ResourceType                        string                        `json:"resourceType"`
	Id                                  *Id                           `json:"id,omitempty"`
	IdPrimitiveElement                  *primitiveElement             `json:"_id,omitempty"`
	Meta                                *Meta                         `json:"meta,omitempty"`
	ImplicitRules                       *Uri                          `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement       *primitiveElement             `json:"_implicitRules,omitempty"`
	Language                            *Code                         `json:"language,omitempty"`
	LanguagePrimitiveElement            *primitiveElement             `json:"_language,omitempty"`
	Text                                *Narrative                    `json:"text,omitempty"`
	Contained                           []ContainedResource           `json:"contained,omitempty"`
	Extension                           []Extension                   `json:"extension,omitempty"`
	ModifierExtension                   []Extension                   `json:"modifierExtension,omitempty"`
	SequenceType                        *CodeableConcept              `json:"sequenceType,omitempty"`
	NumberOfSubunits                    *Integer                      `json:"numberOfSubunits,omitempty"`
	NumberOfSubunitsPrimitiveElement    *primitiveElement             `json:"_numberOfSubunits,omitempty"`
	AreaOfHybridisation                 *String                       `json:"areaOfHybridisation,omitempty"`
	AreaOfHybridisationPrimitiveElement *primitiveElement             `json:"_areaOfHybridisation,omitempty"`
	OligoNucleotideType                 *CodeableConcept              `json:"oligoNucleotideType,omitempty"`
	Subunit                             []SubstanceNucleicAcidSubunit `json:"subunit,omitempty"`
}

func (r SubstanceNucleicAcid) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceNucleicAcid) marshalJSON() jsonSubstanceNucleicAcid {
	m := jsonSubstanceNucleicAcid{}
	m.ResourceType = "SubstanceNucleicAcid"
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
	m.SequenceType = r.SequenceType
	if r.NumberOfSubunits != nil && r.NumberOfSubunits.Value != nil {
		m.NumberOfSubunits = r.NumberOfSubunits
	}
	if r.NumberOfSubunits != nil && (r.NumberOfSubunits.Id != nil || r.NumberOfSubunits.Extension != nil) {
		m.NumberOfSubunitsPrimitiveElement = &primitiveElement{Id: r.NumberOfSubunits.Id, Extension: r.NumberOfSubunits.Extension}
	}
	if r.AreaOfHybridisation != nil && r.AreaOfHybridisation.Value != nil {
		m.AreaOfHybridisation = r.AreaOfHybridisation
	}
	if r.AreaOfHybridisation != nil && (r.AreaOfHybridisation.Id != nil || r.AreaOfHybridisation.Extension != nil) {
		m.AreaOfHybridisationPrimitiveElement = &primitiveElement{Id: r.AreaOfHybridisation.Id, Extension: r.AreaOfHybridisation.Extension}
	}
	m.OligoNucleotideType = r.OligoNucleotideType
	m.Subunit = r.Subunit
	return m
}
func (r *SubstanceNucleicAcid) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceNucleicAcid
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceNucleicAcid) unmarshalJSON(m jsonSubstanceNucleicAcid) error {
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
	r.SequenceType = m.SequenceType
	r.NumberOfSubunits = m.NumberOfSubunits
	if m.NumberOfSubunitsPrimitiveElement != nil {
		if r.NumberOfSubunits == nil {
			r.NumberOfSubunits = &Integer{}
		}
		r.NumberOfSubunits.Id = m.NumberOfSubunitsPrimitiveElement.Id
		r.NumberOfSubunits.Extension = m.NumberOfSubunitsPrimitiveElement.Extension
	}
	r.AreaOfHybridisation = m.AreaOfHybridisation
	if m.AreaOfHybridisationPrimitiveElement != nil {
		if r.AreaOfHybridisation == nil {
			r.AreaOfHybridisation = &String{}
		}
		r.AreaOfHybridisation.Id = m.AreaOfHybridisationPrimitiveElement.Id
		r.AreaOfHybridisation.Extension = m.AreaOfHybridisationPrimitiveElement.Extension
	}
	r.OligoNucleotideType = m.OligoNucleotideType
	r.Subunit = m.Subunit
	return nil
}
func (r SubstanceNucleicAcid) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "SubstanceNucleicAcid"
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
	err = e.EncodeElement(r.SequenceType, xml.StartElement{Name: xml.Name{Local: "sequenceType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NumberOfSubunits, xml.StartElement{Name: xml.Name{Local: "numberOfSubunits"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AreaOfHybridisation, xml.StartElement{Name: xml.Name{Local: "areaOfHybridisation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OligoNucleotideType, xml.StartElement{Name: xml.Name{Local: "oligoNucleotideType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Subunit, xml.StartElement{Name: xml.Name{Local: "subunit"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceNucleicAcid) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "sequenceType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SequenceType = &v
			case "numberOfSubunits":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NumberOfSubunits = &v
			case "areaOfHybridisation":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AreaOfHybridisation = &v
			case "oligoNucleotideType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OligoNucleotideType = &v
			case "subunit":
				var v SubstanceNucleicAcidSubunit
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subunit = append(r.Subunit, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceNucleicAcid) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Subunits are listed in order of decreasing length; sequences of the same length will be ordered by molecular weight; subunits that have identical sequences will be repeated multiple times.
type SubstanceNucleicAcidSubunit struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Index of linear sequences of nucleic acids in order of decreasing length. Sequences of the same length will be ordered by molecular weight. Subunits that have identical sequences will be repeated and have sequential subscripts.
	Subunit *Integer
	// Actual nucleotide sequence notation from 5' to 3' end using standard single letter codes. In addition to the base sequence, sugar and type of phosphate or non-phosphate linkage should also be captured.
	Sequence *String
	// The length of the sequence shall be captured.
	Length *Integer
	// (TBC).
	SequenceAttachment *Attachment
	// The nucleotide present at the 5’ terminal shall be specified based on a controlled vocabulary. Since the sequence is represented from the 5' to the 3' end, the 5’ prime nucleotide is the letter at the first position in the sequence. A separate representation would be redundant.
	FivePrime *CodeableConcept
	// The nucleotide present at the 3’ terminal shall be specified based on a controlled vocabulary. Since the sequence is represented from the 5' to the 3' end, the 5’ prime nucleotide is the letter at the last position in the sequence. A separate representation would be redundant.
	ThreePrime *CodeableConcept
	// The linkages between sugar residues will also be captured.
	Linkage []SubstanceNucleicAcidSubunitLinkage
	// 5.3.6.8.1 Sugar ID (Mandatory).
	Sugar []SubstanceNucleicAcidSubunitSugar
}
type jsonSubstanceNucleicAcidSubunit struct {
	Id                       *string                              `json:"id,omitempty"`
	Extension                []Extension                          `json:"extension,omitempty"`
	ModifierExtension        []Extension                          `json:"modifierExtension,omitempty"`
	Subunit                  *Integer                             `json:"subunit,omitempty"`
	SubunitPrimitiveElement  *primitiveElement                    `json:"_subunit,omitempty"`
	Sequence                 *String                              `json:"sequence,omitempty"`
	SequencePrimitiveElement *primitiveElement                    `json:"_sequence,omitempty"`
	Length                   *Integer                             `json:"length,omitempty"`
	LengthPrimitiveElement   *primitiveElement                    `json:"_length,omitempty"`
	SequenceAttachment       *Attachment                          `json:"sequenceAttachment,omitempty"`
	FivePrime                *CodeableConcept                     `json:"fivePrime,omitempty"`
	ThreePrime               *CodeableConcept                     `json:"threePrime,omitempty"`
	Linkage                  []SubstanceNucleicAcidSubunitLinkage `json:"linkage,omitempty"`
	Sugar                    []SubstanceNucleicAcidSubunitSugar   `json:"sugar,omitempty"`
}

func (r SubstanceNucleicAcidSubunit) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceNucleicAcidSubunit) marshalJSON() jsonSubstanceNucleicAcidSubunit {
	m := jsonSubstanceNucleicAcidSubunit{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Subunit != nil && r.Subunit.Value != nil {
		m.Subunit = r.Subunit
	}
	if r.Subunit != nil && (r.Subunit.Id != nil || r.Subunit.Extension != nil) {
		m.SubunitPrimitiveElement = &primitiveElement{Id: r.Subunit.Id, Extension: r.Subunit.Extension}
	}
	if r.Sequence != nil && r.Sequence.Value != nil {
		m.Sequence = r.Sequence
	}
	if r.Sequence != nil && (r.Sequence.Id != nil || r.Sequence.Extension != nil) {
		m.SequencePrimitiveElement = &primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
	}
	if r.Length != nil && r.Length.Value != nil {
		m.Length = r.Length
	}
	if r.Length != nil && (r.Length.Id != nil || r.Length.Extension != nil) {
		m.LengthPrimitiveElement = &primitiveElement{Id: r.Length.Id, Extension: r.Length.Extension}
	}
	m.SequenceAttachment = r.SequenceAttachment
	m.FivePrime = r.FivePrime
	m.ThreePrime = r.ThreePrime
	m.Linkage = r.Linkage
	m.Sugar = r.Sugar
	return m
}
func (r *SubstanceNucleicAcidSubunit) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceNucleicAcidSubunit
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceNucleicAcidSubunit) unmarshalJSON(m jsonSubstanceNucleicAcidSubunit) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Subunit = m.Subunit
	if m.SubunitPrimitiveElement != nil {
		if r.Subunit == nil {
			r.Subunit = &Integer{}
		}
		r.Subunit.Id = m.SubunitPrimitiveElement.Id
		r.Subunit.Extension = m.SubunitPrimitiveElement.Extension
	}
	r.Sequence = m.Sequence
	if m.SequencePrimitiveElement != nil {
		if r.Sequence == nil {
			r.Sequence = &String{}
		}
		r.Sequence.Id = m.SequencePrimitiveElement.Id
		r.Sequence.Extension = m.SequencePrimitiveElement.Extension
	}
	r.Length = m.Length
	if m.LengthPrimitiveElement != nil {
		if r.Length == nil {
			r.Length = &Integer{}
		}
		r.Length.Id = m.LengthPrimitiveElement.Id
		r.Length.Extension = m.LengthPrimitiveElement.Extension
	}
	r.SequenceAttachment = m.SequenceAttachment
	r.FivePrime = m.FivePrime
	r.ThreePrime = m.ThreePrime
	r.Linkage = m.Linkage
	r.Sugar = m.Sugar
	return nil
}
func (r SubstanceNucleicAcidSubunit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Subunit, xml.StartElement{Name: xml.Name{Local: "subunit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Sequence, xml.StartElement{Name: xml.Name{Local: "sequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Length, xml.StartElement{Name: xml.Name{Local: "length"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SequenceAttachment, xml.StartElement{Name: xml.Name{Local: "sequenceAttachment"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FivePrime, xml.StartElement{Name: xml.Name{Local: "fivePrime"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ThreePrime, xml.StartElement{Name: xml.Name{Local: "threePrime"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Linkage, xml.StartElement{Name: xml.Name{Local: "linkage"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Sugar, xml.StartElement{Name: xml.Name{Local: "sugar"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceNucleicAcidSubunit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "subunit":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subunit = &v
			case "sequence":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Sequence = &v
			case "length":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Length = &v
			case "sequenceAttachment":
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SequenceAttachment = &v
			case "fivePrime":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FivePrime = &v
			case "threePrime":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ThreePrime = &v
			case "linkage":
				var v SubstanceNucleicAcidSubunitLinkage
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Linkage = append(r.Linkage, v)
			case "sugar":
				var v SubstanceNucleicAcidSubunitSugar
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Sugar = append(r.Sugar, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceNucleicAcidSubunit) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The linkages between sugar residues will also be captured.
type SubstanceNucleicAcidSubunitLinkage struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The entity that links the sugar residues together should also be captured for nearly all naturally occurring nucleic acid the linkage is a phosphate group. For many synthetic oligonucleotides phosphorothioate linkages are often seen. Linkage connectivity is assumed to be 3’-5’. If the linkage is either 3’-3’ or 5’-5’ this should be specified.
	Connectivity *String
	// Each linkage will be registered as a fragment and have an ID.
	Identifier *Identifier
	// Each linkage will be registered as a fragment and have at least one name. A single name shall be assigned to each linkage.
	Name *String
	// Residues shall be captured as described in 5.3.6.8.3.
	ResidueSite *String
}
type jsonSubstanceNucleicAcidSubunitLinkage struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Connectivity                 *String           `json:"connectivity,omitempty"`
	ConnectivityPrimitiveElement *primitiveElement `json:"_connectivity,omitempty"`
	Identifier                   *Identifier       `json:"identifier,omitempty"`
	Name                         *String           `json:"name,omitempty"`
	NamePrimitiveElement         *primitiveElement `json:"_name,omitempty"`
	ResidueSite                  *String           `json:"residueSite,omitempty"`
	ResidueSitePrimitiveElement  *primitiveElement `json:"_residueSite,omitempty"`
}

func (r SubstanceNucleicAcidSubunitLinkage) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceNucleicAcidSubunitLinkage) marshalJSON() jsonSubstanceNucleicAcidSubunitLinkage {
	m := jsonSubstanceNucleicAcidSubunitLinkage{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Connectivity != nil && r.Connectivity.Value != nil {
		m.Connectivity = r.Connectivity
	}
	if r.Connectivity != nil && (r.Connectivity.Id != nil || r.Connectivity.Extension != nil) {
		m.ConnectivityPrimitiveElement = &primitiveElement{Id: r.Connectivity.Id, Extension: r.Connectivity.Extension}
	}
	m.Identifier = r.Identifier
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.ResidueSite != nil && r.ResidueSite.Value != nil {
		m.ResidueSite = r.ResidueSite
	}
	if r.ResidueSite != nil && (r.ResidueSite.Id != nil || r.ResidueSite.Extension != nil) {
		m.ResidueSitePrimitiveElement = &primitiveElement{Id: r.ResidueSite.Id, Extension: r.ResidueSite.Extension}
	}
	return m
}
func (r *SubstanceNucleicAcidSubunitLinkage) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceNucleicAcidSubunitLinkage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceNucleicAcidSubunitLinkage) unmarshalJSON(m jsonSubstanceNucleicAcidSubunitLinkage) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Connectivity = m.Connectivity
	if m.ConnectivityPrimitiveElement != nil {
		if r.Connectivity == nil {
			r.Connectivity = &String{}
		}
		r.Connectivity.Id = m.ConnectivityPrimitiveElement.Id
		r.Connectivity.Extension = m.ConnectivityPrimitiveElement.Extension
	}
	r.Identifier = m.Identifier
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.ResidueSite = m.ResidueSite
	if m.ResidueSitePrimitiveElement != nil {
		if r.ResidueSite == nil {
			r.ResidueSite = &String{}
		}
		r.ResidueSite.Id = m.ResidueSitePrimitiveElement.Id
		r.ResidueSite.Extension = m.ResidueSitePrimitiveElement.Extension
	}
	return nil
}
func (r SubstanceNucleicAcidSubunitLinkage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Connectivity, xml.StartElement{Name: xml.Name{Local: "connectivity"}})
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
	err = e.EncodeElement(r.ResidueSite, xml.StartElement{Name: xml.Name{Local: "residueSite"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceNucleicAcidSubunitLinkage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "connectivity":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Connectivity = &v
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
			case "residueSite":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ResidueSite = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceNucleicAcidSubunitLinkage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// 5.3.6.8.1 Sugar ID (Mandatory).
type SubstanceNucleicAcidSubunitSugar struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The Substance ID of the sugar or sugar-like component that make up the nucleotide.
	Identifier *Identifier
	// The name of the sugar or sugar-like component that make up the nucleotide.
	Name *String
	// The residues that contain a given sugar will be captured. The order of given residues will be captured in the 5‘-3‘direction consistent with the base sequences listed above.
	ResidueSite *String
}
type jsonSubstanceNucleicAcidSubunitSugar struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Identifier                  *Identifier       `json:"identifier,omitempty"`
	Name                        *String           `json:"name,omitempty"`
	NamePrimitiveElement        *primitiveElement `json:"_name,omitempty"`
	ResidueSite                 *String           `json:"residueSite,omitempty"`
	ResidueSitePrimitiveElement *primitiveElement `json:"_residueSite,omitempty"`
}

func (r SubstanceNucleicAcidSubunitSugar) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceNucleicAcidSubunitSugar) marshalJSON() jsonSubstanceNucleicAcidSubunitSugar {
	m := jsonSubstanceNucleicAcidSubunitSugar{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.ResidueSite != nil && r.ResidueSite.Value != nil {
		m.ResidueSite = r.ResidueSite
	}
	if r.ResidueSite != nil && (r.ResidueSite.Id != nil || r.ResidueSite.Extension != nil) {
		m.ResidueSitePrimitiveElement = &primitiveElement{Id: r.ResidueSite.Id, Extension: r.ResidueSite.Extension}
	}
	return m
}
func (r *SubstanceNucleicAcidSubunitSugar) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceNucleicAcidSubunitSugar
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceNucleicAcidSubunitSugar) unmarshalJSON(m jsonSubstanceNucleicAcidSubunitSugar) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.ResidueSite = m.ResidueSite
	if m.ResidueSitePrimitiveElement != nil {
		if r.ResidueSite == nil {
			r.ResidueSite = &String{}
		}
		r.ResidueSite.Id = m.ResidueSitePrimitiveElement.Id
		r.ResidueSite.Extension = m.ResidueSitePrimitiveElement.Extension
	}
	return nil
}
func (r SubstanceNucleicAcidSubunitSugar) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ResidueSite, xml.StartElement{Name: xml.Name{Local: "residueSite"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SubstanceNucleicAcidSubunitSugar) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "residueSite":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ResidueSite = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r SubstanceNucleicAcidSubunitSugar) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
