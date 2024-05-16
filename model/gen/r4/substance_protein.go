package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A SubstanceProtein is defined as a single unit of a linear amino acid sequence, or a combination of subunits that are either covalently linked or have a defined invariant stoichiometric relationship. This includes all synthetic, recombinant and purified SubstanceProteins of defined sequence, whether the use is therapeutic or prophylactic. This set of elements will be used to describe albumins, coagulation factors, cytokines, growth factors, peptide/SubstanceProtein hormones, enzymes, toxins, toxoids, recombinant vaccines, and immunomodulators.
type SubstanceProtein struct {
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
	// The SubstanceProtein descriptive elements will only be used when a complete or partial amino acid sequence is available or derivable from a nucleic acid sequence.
	SequenceType *CodeableConcept
	// Number of linear sequences of amino acids linked through peptide bonds. The number of subunits constituting the SubstanceProtein shall be described. It is possible that the number of subunits can be variable.
	NumberOfSubunits *Integer
	// The disulphide bond between two cysteine residues either on the same subunit or on two different subunits shall be described. The position of the disulfide bonds in the SubstanceProtein shall be listed in increasing order of subunit number and position within subunit followed by the abbreviation of the amino acids involved. The disulfide linkage positions shall actually contain the amino acid Cysteine at the respective positions.
	DisulfideLinkage []String
	// This subclause refers to the description of each subunit constituting the SubstanceProtein. A subunit is a linear sequence of amino acids linked through peptide bonds. The Subunit information shall be provided when the finished SubstanceProtein is a complex of multiple sequences; subunits are not used to delineate domains within a single sequence. Subunits are listed in order of decreasing length; sequences of the same length will be ordered by decreasing molecular weight; subunits that have identical sequences will be repeated multiple times.
	Subunit []SubstanceProteinSubunit
}

func (r SubstanceProtein) ResourceType() string {
	return "SubstanceProtein"
}
func (r SubstanceProtein) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonSubstanceProtein struct {
	ResourceType                     string                    `json:"resourceType"`
	Id                               *Id                       `json:"id,omitempty"`
	IdPrimitiveElement               *primitiveElement         `json:"_id,omitempty"`
	Meta                             *Meta                     `json:"meta,omitempty"`
	ImplicitRules                    *Uri                      `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement    *primitiveElement         `json:"_implicitRules,omitempty"`
	Language                         *Code                     `json:"language,omitempty"`
	LanguagePrimitiveElement         *primitiveElement         `json:"_language,omitempty"`
	Text                             *Narrative                `json:"text,omitempty"`
	Contained                        []containedResource       `json:"contained,omitempty"`
	Extension                        []Extension               `json:"extension,omitempty"`
	ModifierExtension                []Extension               `json:"modifierExtension,omitempty"`
	SequenceType                     *CodeableConcept          `json:"sequenceType,omitempty"`
	NumberOfSubunits                 *Integer                  `json:"numberOfSubunits,omitempty"`
	NumberOfSubunitsPrimitiveElement *primitiveElement         `json:"_numberOfSubunits,omitempty"`
	DisulfideLinkage                 []String                  `json:"disulfideLinkage,omitempty"`
	DisulfideLinkagePrimitiveElement []*primitiveElement       `json:"_disulfideLinkage,omitempty"`
	Subunit                          []SubstanceProteinSubunit `json:"subunit,omitempty"`
}

func (r SubstanceProtein) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceProtein) marshalJSON() jsonSubstanceProtein {
	m := jsonSubstanceProtein{}
	m.ResourceType = "SubstanceProtein"
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
	m.SequenceType = r.SequenceType
	m.NumberOfSubunits = r.NumberOfSubunits
	if r.NumberOfSubunits != nil && (r.NumberOfSubunits.Id != nil || r.NumberOfSubunits.Extension != nil) {
		m.NumberOfSubunitsPrimitiveElement = &primitiveElement{Id: r.NumberOfSubunits.Id, Extension: r.NumberOfSubunits.Extension}
	}
	m.DisulfideLinkage = r.DisulfideLinkage
	anyDisulfideLinkageIdOrExtension := false
	for _, e := range r.DisulfideLinkage {
		if e.Id != nil || e.Extension != nil {
			anyDisulfideLinkageIdOrExtension = true
			break
		}
	}
	if anyDisulfideLinkageIdOrExtension {
		m.DisulfideLinkagePrimitiveElement = make([]*primitiveElement, 0, len(r.DisulfideLinkage))
		for _, e := range r.DisulfideLinkage {
			if e.Id != nil || e.Extension != nil {
				m.DisulfideLinkagePrimitiveElement = append(m.DisulfideLinkagePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.DisulfideLinkagePrimitiveElement = append(m.DisulfideLinkagePrimitiveElement, nil)
			}
		}
	}
	m.Subunit = r.Subunit
	return m
}
func (r *SubstanceProtein) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceProtein
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceProtein) unmarshalJSON(m jsonSubstanceProtein) error {
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
	r.SequenceType = m.SequenceType
	r.NumberOfSubunits = m.NumberOfSubunits
	if m.NumberOfSubunitsPrimitiveElement != nil {
		r.NumberOfSubunits.Id = m.NumberOfSubunitsPrimitiveElement.Id
		r.NumberOfSubunits.Extension = m.NumberOfSubunitsPrimitiveElement.Extension
	}
	r.DisulfideLinkage = m.DisulfideLinkage
	for i, e := range m.DisulfideLinkagePrimitiveElement {
		if len(r.DisulfideLinkage) > i {
			r.DisulfideLinkage[i].Id = e.Id
			r.DisulfideLinkage[i].Extension = e.Extension
		} else {
			r.DisulfideLinkage = append(r.DisulfideLinkage, String{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Subunit = m.Subunit
	return nil
}
func (r SubstanceProtein) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// This subclause refers to the description of each subunit constituting the SubstanceProtein. A subunit is a linear sequence of amino acids linked through peptide bonds. The Subunit information shall be provided when the finished SubstanceProtein is a complex of multiple sequences; subunits are not used to delineate domains within a single sequence. Subunits are listed in order of decreasing length; sequences of the same length will be ordered by decreasing molecular weight; subunits that have identical sequences will be repeated multiple times.
type SubstanceProteinSubunit struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Index of primary sequences of amino acids linked through peptide bonds in order of decreasing length. Sequences of the same length will be ordered by molecular weight. Subunits that have identical sequences will be repeated and have sequential subscripts.
	Subunit *Integer
	// The sequence information shall be provided enumerating the amino acids from N- to C-terminal end using standard single-letter amino acid codes. Uppercase shall be used for L-amino acids and lowercase for D-amino acids. Transcribed SubstanceProteins will always be described using the translated sequence; for synthetic peptide containing amino acids that are not represented with a single letter code an X should be used within the sequence. The modified amino acids will be distinguished by their position in the sequence.
	Sequence *String
	// Length of linear sequences of amino acids contained in the subunit.
	Length *Integer
	// The sequence information shall be provided enumerating the amino acids from N- to C-terminal end using standard single-letter amino acid codes. Uppercase shall be used for L-amino acids and lowercase for D-amino acids. Transcribed SubstanceProteins will always be described using the translated sequence; for synthetic peptide containing amino acids that are not represented with a single letter code an X should be used within the sequence. The modified amino acids will be distinguished by their position in the sequence.
	SequenceAttachment *Attachment
	// Unique identifier for molecular fragment modification based on the ISO 11238 Substance ID.
	NTerminalModificationId *Identifier
	// The name of the fragment modified at the N-terminal of the SubstanceProtein shall be specified.
	NTerminalModification *String
	// Unique identifier for molecular fragment modification based on the ISO 11238 Substance ID.
	CTerminalModificationId *Identifier
	// The modification at the C-terminal shall be specified.
	CTerminalModification *String
}
type jsonSubstanceProteinSubunit struct {
	Id                                    *string           `json:"id,omitempty"`
	Extension                             []Extension       `json:"extension,omitempty"`
	ModifierExtension                     []Extension       `json:"modifierExtension,omitempty"`
	Subunit                               *Integer          `json:"subunit,omitempty"`
	SubunitPrimitiveElement               *primitiveElement `json:"_subunit,omitempty"`
	Sequence                              *String           `json:"sequence,omitempty"`
	SequencePrimitiveElement              *primitiveElement `json:"_sequence,omitempty"`
	Length                                *Integer          `json:"length,omitempty"`
	LengthPrimitiveElement                *primitiveElement `json:"_length,omitempty"`
	SequenceAttachment                    *Attachment       `json:"sequenceAttachment,omitempty"`
	NTerminalModificationId               *Identifier       `json:"nTerminalModificationId,omitempty"`
	NTerminalModification                 *String           `json:"nTerminalModification,omitempty"`
	NTerminalModificationPrimitiveElement *primitiveElement `json:"_nTerminalModification,omitempty"`
	CTerminalModificationId               *Identifier       `json:"cTerminalModificationId,omitempty"`
	CTerminalModification                 *String           `json:"cTerminalModification,omitempty"`
	CTerminalModificationPrimitiveElement *primitiveElement `json:"_cTerminalModification,omitempty"`
}

func (r SubstanceProteinSubunit) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstanceProteinSubunit) marshalJSON() jsonSubstanceProteinSubunit {
	m := jsonSubstanceProteinSubunit{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Subunit = r.Subunit
	if r.Subunit != nil && (r.Subunit.Id != nil || r.Subunit.Extension != nil) {
		m.SubunitPrimitiveElement = &primitiveElement{Id: r.Subunit.Id, Extension: r.Subunit.Extension}
	}
	m.Sequence = r.Sequence
	if r.Sequence != nil && (r.Sequence.Id != nil || r.Sequence.Extension != nil) {
		m.SequencePrimitiveElement = &primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
	}
	m.Length = r.Length
	if r.Length != nil && (r.Length.Id != nil || r.Length.Extension != nil) {
		m.LengthPrimitiveElement = &primitiveElement{Id: r.Length.Id, Extension: r.Length.Extension}
	}
	m.SequenceAttachment = r.SequenceAttachment
	m.NTerminalModificationId = r.NTerminalModificationId
	m.NTerminalModification = r.NTerminalModification
	if r.NTerminalModification != nil && (r.NTerminalModification.Id != nil || r.NTerminalModification.Extension != nil) {
		m.NTerminalModificationPrimitiveElement = &primitiveElement{Id: r.NTerminalModification.Id, Extension: r.NTerminalModification.Extension}
	}
	m.CTerminalModificationId = r.CTerminalModificationId
	m.CTerminalModification = r.CTerminalModification
	if r.CTerminalModification != nil && (r.CTerminalModification.Id != nil || r.CTerminalModification.Extension != nil) {
		m.CTerminalModificationPrimitiveElement = &primitiveElement{Id: r.CTerminalModification.Id, Extension: r.CTerminalModification.Extension}
	}
	return m
}
func (r *SubstanceProteinSubunit) UnmarshalJSON(b []byte) error {
	var m jsonSubstanceProteinSubunit
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstanceProteinSubunit) unmarshalJSON(m jsonSubstanceProteinSubunit) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Subunit = m.Subunit
	if m.SubunitPrimitiveElement != nil {
		r.Subunit.Id = m.SubunitPrimitiveElement.Id
		r.Subunit.Extension = m.SubunitPrimitiveElement.Extension
	}
	r.Sequence = m.Sequence
	if m.SequencePrimitiveElement != nil {
		r.Sequence.Id = m.SequencePrimitiveElement.Id
		r.Sequence.Extension = m.SequencePrimitiveElement.Extension
	}
	r.Length = m.Length
	if m.LengthPrimitiveElement != nil {
		r.Length.Id = m.LengthPrimitiveElement.Id
		r.Length.Extension = m.LengthPrimitiveElement.Extension
	}
	r.SequenceAttachment = m.SequenceAttachment
	r.NTerminalModificationId = m.NTerminalModificationId
	r.NTerminalModification = m.NTerminalModification
	if m.NTerminalModificationPrimitiveElement != nil {
		r.NTerminalModification.Id = m.NTerminalModificationPrimitiveElement.Id
		r.NTerminalModification.Extension = m.NTerminalModificationPrimitiveElement.Extension
	}
	r.CTerminalModificationId = m.CTerminalModificationId
	r.CTerminalModification = m.CTerminalModification
	if m.CTerminalModificationPrimitiveElement != nil {
		r.CTerminalModification.Id = m.CTerminalModificationPrimitiveElement.Id
		r.CTerminalModification.Extension = m.CTerminalModificationPrimitiveElement.Extension
	}
	return nil
}
func (r SubstanceProteinSubunit) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
