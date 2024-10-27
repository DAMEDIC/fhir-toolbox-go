package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// Raw data describing a biological sequence.
type MolecularSequence struct {
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
	// A unique identifier for this particular sequence instance. This is a FHIR-defined id.
	Identifier []Identifier
	// Amino Acid Sequence/ DNA Sequence / RNA Sequence.
	Type *Code
	// Whether the sequence is numbered starting at 0 (0-based numbering or coordinates, inclusive start, exclusive end) or starting at 1 (1-based numbering, inclusive start and inclusive end).
	CoordinateSystem Integer
	// The patient whose sequencing results are described by this resource.
	Patient *Reference
	// Specimen used for sequencing.
	Specimen *Reference
	// The method for sequencing, for example, chip information.
	Device *Reference
	// The organization or lab that should be responsible for this result.
	Performer *Reference
	// The number of copies of the sequence of interest. (RNASeq).
	Quantity *Quantity
	// A sequence that is used as a reference to describe variants that are present in a sequence analyzed.
	ReferenceSeq *MolecularSequenceReferenceSeq
	// The definition of variant here originates from Sequence ontology ([variant_of](http://www.sequenceontology.org/browser/current_svn/term/variant_of)). This element can represent amino acid or nucleic sequence change(including insertion,deletion,SNP,etc.)  It can represent some complex mutation or segment variation with the assist of CIGAR string.
	Variant []MolecularSequenceVariant
	// Sequence that was observed. It is the result marked by referenceSeq along with variant records on referenceSeq. This shall start from referenceSeq.windowStart and end by referenceSeq.windowEnd.
	ObservedSeq *String
	// An experimental feature attribute that defines the quality of the feature in a quantitative way, such as a phred quality score ([SO:0001686](http://www.sequenceontology.org/browser/current_svn/term/SO:0001686)).
	Quality []MolecularSequenceQuality
	// Coverage (read depth or depth) is the average number of reads representing a given nucleotide in the reconstructed sequence.
	ReadCoverage *Integer
	// Configurations of the external repository. The repository shall store target's observedSeq or records related with target's observedSeq.
	Repository []MolecularSequenceRepository
	// Pointer to next atomic sequence which at most contains one variant.
	Pointer []Reference
	// Information about chromosome structure variation.
	StructureVariant []MolecularSequenceStructureVariant
}

func (r MolecularSequence) ResourceType() string {
	return "MolecularSequence"
}
func (r MolecularSequence) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonMolecularSequence struct {
	ResourceType                     string                              `json:"resourceType"`
	Id                               *Id                                 `json:"id,omitempty"`
	IdPrimitiveElement               *primitiveElement                   `json:"_id,omitempty"`
	Meta                             *Meta                               `json:"meta,omitempty"`
	ImplicitRules                    *Uri                                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement    *primitiveElement                   `json:"_implicitRules,omitempty"`
	Language                         *Code                               `json:"language,omitempty"`
	LanguagePrimitiveElement         *primitiveElement                   `json:"_language,omitempty"`
	Text                             *Narrative                          `json:"text,omitempty"`
	Contained                        []ContainedResource                 `json:"contained,omitempty"`
	Extension                        []Extension                         `json:"extension,omitempty"`
	ModifierExtension                []Extension                         `json:"modifierExtension,omitempty"`
	Identifier                       []Identifier                        `json:"identifier,omitempty"`
	Type                             *Code                               `json:"type,omitempty"`
	TypePrimitiveElement             *primitiveElement                   `json:"_type,omitempty"`
	CoordinateSystem                 Integer                             `json:"coordinateSystem,omitempty"`
	CoordinateSystemPrimitiveElement *primitiveElement                   `json:"_coordinateSystem,omitempty"`
	Patient                          *Reference                          `json:"patient,omitempty"`
	Specimen                         *Reference                          `json:"specimen,omitempty"`
	Device                           *Reference                          `json:"device,omitempty"`
	Performer                        *Reference                          `json:"performer,omitempty"`
	Quantity                         *Quantity                           `json:"quantity,omitempty"`
	ReferenceSeq                     *MolecularSequenceReferenceSeq      `json:"referenceSeq,omitempty"`
	Variant                          []MolecularSequenceVariant          `json:"variant,omitempty"`
	ObservedSeq                      *String                             `json:"observedSeq,omitempty"`
	ObservedSeqPrimitiveElement      *primitiveElement                   `json:"_observedSeq,omitempty"`
	Quality                          []MolecularSequenceQuality          `json:"quality,omitempty"`
	ReadCoverage                     *Integer                            `json:"readCoverage,omitempty"`
	ReadCoveragePrimitiveElement     *primitiveElement                   `json:"_readCoverage,omitempty"`
	Repository                       []MolecularSequenceRepository       `json:"repository,omitempty"`
	Pointer                          []Reference                         `json:"pointer,omitempty"`
	StructureVariant                 []MolecularSequenceStructureVariant `json:"structureVariant,omitempty"`
}

func (r MolecularSequence) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MolecularSequence) marshalJSON() jsonMolecularSequence {
	m := jsonMolecularSequence{}
	m.ResourceType = "MolecularSequence"
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
	if r.Type != nil && r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.CoordinateSystem.Value != nil {
		m.CoordinateSystem = r.CoordinateSystem
	}
	if r.CoordinateSystem.Id != nil || r.CoordinateSystem.Extension != nil {
		m.CoordinateSystemPrimitiveElement = &primitiveElement{Id: r.CoordinateSystem.Id, Extension: r.CoordinateSystem.Extension}
	}
	m.Patient = r.Patient
	m.Specimen = r.Specimen
	m.Device = r.Device
	m.Performer = r.Performer
	m.Quantity = r.Quantity
	m.ReferenceSeq = r.ReferenceSeq
	m.Variant = r.Variant
	if r.ObservedSeq != nil && r.ObservedSeq.Value != nil {
		m.ObservedSeq = r.ObservedSeq
	}
	if r.ObservedSeq != nil && (r.ObservedSeq.Id != nil || r.ObservedSeq.Extension != nil) {
		m.ObservedSeqPrimitiveElement = &primitiveElement{Id: r.ObservedSeq.Id, Extension: r.ObservedSeq.Extension}
	}
	m.Quality = r.Quality
	if r.ReadCoverage != nil && r.ReadCoverage.Value != nil {
		m.ReadCoverage = r.ReadCoverage
	}
	if r.ReadCoverage != nil && (r.ReadCoverage.Id != nil || r.ReadCoverage.Extension != nil) {
		m.ReadCoveragePrimitiveElement = &primitiveElement{Id: r.ReadCoverage.Id, Extension: r.ReadCoverage.Extension}
	}
	m.Repository = r.Repository
	m.Pointer = r.Pointer
	m.StructureVariant = r.StructureVariant
	return m
}
func (r *MolecularSequence) UnmarshalJSON(b []byte) error {
	var m jsonMolecularSequence
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MolecularSequence) unmarshalJSON(m jsonMolecularSequence) error {
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
	if m.TypePrimitiveElement != nil {
		if r.Type == nil {
			r.Type = &Code{}
		}
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.CoordinateSystem = m.CoordinateSystem
	if m.CoordinateSystemPrimitiveElement != nil {
		r.CoordinateSystem.Id = m.CoordinateSystemPrimitiveElement.Id
		r.CoordinateSystem.Extension = m.CoordinateSystemPrimitiveElement.Extension
	}
	r.Patient = m.Patient
	r.Specimen = m.Specimen
	r.Device = m.Device
	r.Performer = m.Performer
	r.Quantity = m.Quantity
	r.ReferenceSeq = m.ReferenceSeq
	r.Variant = m.Variant
	r.ObservedSeq = m.ObservedSeq
	if m.ObservedSeqPrimitiveElement != nil {
		if r.ObservedSeq == nil {
			r.ObservedSeq = &String{}
		}
		r.ObservedSeq.Id = m.ObservedSeqPrimitiveElement.Id
		r.ObservedSeq.Extension = m.ObservedSeqPrimitiveElement.Extension
	}
	r.Quality = m.Quality
	r.ReadCoverage = m.ReadCoverage
	if m.ReadCoveragePrimitiveElement != nil {
		if r.ReadCoverage == nil {
			r.ReadCoverage = &Integer{}
		}
		r.ReadCoverage.Id = m.ReadCoveragePrimitiveElement.Id
		r.ReadCoverage.Extension = m.ReadCoveragePrimitiveElement.Extension
	}
	r.Repository = m.Repository
	r.Pointer = m.Pointer
	r.StructureVariant = m.StructureVariant
	return nil
}
func (r MolecularSequence) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A sequence that is used as a reference to describe variants that are present in a sequence analyzed.
type MolecularSequenceReferenceSeq struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Structural unit composed of a nucleic acid molecule which controls its own replication through the interaction of specific proteins at one or more origins of replication ([SO:0000340](http://www.sequenceontology.org/browser/current_svn/term/SO:0000340)).
	Chromosome *CodeableConcept
	// The Genome Build used for reference, following GRCh build versions e.g. 'GRCh 37'.  Version number must be included if a versioned release of a primary build was used.
	GenomeBuild *String
	// A relative reference to a DNA strand based on gene orientation. The strand that contains the open reading frame of the gene is the "sense" strand, and the opposite complementary strand is the "antisense" strand.
	Orientation *Code
	// Reference identifier of reference sequence submitted to NCBI. It must match the type in the MolecularSequence.type field. For example, the prefix, “NG_” identifies reference sequence for genes, “NM_” for messenger RNA transcripts, and “NP_” for amino acid sequences.
	ReferenceSeqId *CodeableConcept
	// A pointer to another MolecularSequence entity as reference sequence.
	ReferenceSeqPointer *Reference
	// A string like "ACGT".
	ReferenceSeqString *String
	// An absolute reference to a strand. The Watson strand is the strand whose 5'-end is on the short arm of the chromosome, and the Crick strand as the one whose 5'-end is on the long arm.
	Strand *Code
	// Start position of the window on the reference sequence. If the coordinate system is either 0-based or 1-based, then start position is inclusive.
	WindowStart *Integer
	// End position of the window on the reference sequence. If the coordinate system is 0-based then end is exclusive and does not include the last position. If the coordinate system is 1-base, then end is inclusive and includes the last position.
	WindowEnd *Integer
}
type jsonMolecularSequenceReferenceSeq struct {
	Id                                 *string           `json:"id,omitempty"`
	Extension                          []Extension       `json:"extension,omitempty"`
	ModifierExtension                  []Extension       `json:"modifierExtension,omitempty"`
	Chromosome                         *CodeableConcept  `json:"chromosome,omitempty"`
	GenomeBuild                        *String           `json:"genomeBuild,omitempty"`
	GenomeBuildPrimitiveElement        *primitiveElement `json:"_genomeBuild,omitempty"`
	Orientation                        *Code             `json:"orientation,omitempty"`
	OrientationPrimitiveElement        *primitiveElement `json:"_orientation,omitempty"`
	ReferenceSeqId                     *CodeableConcept  `json:"referenceSeqId,omitempty"`
	ReferenceSeqPointer                *Reference        `json:"referenceSeqPointer,omitempty"`
	ReferenceSeqString                 *String           `json:"referenceSeqString,omitempty"`
	ReferenceSeqStringPrimitiveElement *primitiveElement `json:"_referenceSeqString,omitempty"`
	Strand                             *Code             `json:"strand,omitempty"`
	StrandPrimitiveElement             *primitiveElement `json:"_strand,omitempty"`
	WindowStart                        *Integer          `json:"windowStart,omitempty"`
	WindowStartPrimitiveElement        *primitiveElement `json:"_windowStart,omitempty"`
	WindowEnd                          *Integer          `json:"windowEnd,omitempty"`
	WindowEndPrimitiveElement          *primitiveElement `json:"_windowEnd,omitempty"`
}

func (r MolecularSequenceReferenceSeq) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MolecularSequenceReferenceSeq) marshalJSON() jsonMolecularSequenceReferenceSeq {
	m := jsonMolecularSequenceReferenceSeq{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Chromosome = r.Chromosome
	if r.GenomeBuild != nil && r.GenomeBuild.Value != nil {
		m.GenomeBuild = r.GenomeBuild
	}
	if r.GenomeBuild != nil && (r.GenomeBuild.Id != nil || r.GenomeBuild.Extension != nil) {
		m.GenomeBuildPrimitiveElement = &primitiveElement{Id: r.GenomeBuild.Id, Extension: r.GenomeBuild.Extension}
	}
	if r.Orientation != nil && r.Orientation.Value != nil {
		m.Orientation = r.Orientation
	}
	if r.Orientation != nil && (r.Orientation.Id != nil || r.Orientation.Extension != nil) {
		m.OrientationPrimitiveElement = &primitiveElement{Id: r.Orientation.Id, Extension: r.Orientation.Extension}
	}
	m.ReferenceSeqId = r.ReferenceSeqId
	m.ReferenceSeqPointer = r.ReferenceSeqPointer
	if r.ReferenceSeqString != nil && r.ReferenceSeqString.Value != nil {
		m.ReferenceSeqString = r.ReferenceSeqString
	}
	if r.ReferenceSeqString != nil && (r.ReferenceSeqString.Id != nil || r.ReferenceSeqString.Extension != nil) {
		m.ReferenceSeqStringPrimitiveElement = &primitiveElement{Id: r.ReferenceSeqString.Id, Extension: r.ReferenceSeqString.Extension}
	}
	if r.Strand != nil && r.Strand.Value != nil {
		m.Strand = r.Strand
	}
	if r.Strand != nil && (r.Strand.Id != nil || r.Strand.Extension != nil) {
		m.StrandPrimitiveElement = &primitiveElement{Id: r.Strand.Id, Extension: r.Strand.Extension}
	}
	if r.WindowStart != nil && r.WindowStart.Value != nil {
		m.WindowStart = r.WindowStart
	}
	if r.WindowStart != nil && (r.WindowStart.Id != nil || r.WindowStart.Extension != nil) {
		m.WindowStartPrimitiveElement = &primitiveElement{Id: r.WindowStart.Id, Extension: r.WindowStart.Extension}
	}
	if r.WindowEnd != nil && r.WindowEnd.Value != nil {
		m.WindowEnd = r.WindowEnd
	}
	if r.WindowEnd != nil && (r.WindowEnd.Id != nil || r.WindowEnd.Extension != nil) {
		m.WindowEndPrimitiveElement = &primitiveElement{Id: r.WindowEnd.Id, Extension: r.WindowEnd.Extension}
	}
	return m
}
func (r *MolecularSequenceReferenceSeq) UnmarshalJSON(b []byte) error {
	var m jsonMolecularSequenceReferenceSeq
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MolecularSequenceReferenceSeq) unmarshalJSON(m jsonMolecularSequenceReferenceSeq) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Chromosome = m.Chromosome
	r.GenomeBuild = m.GenomeBuild
	if m.GenomeBuildPrimitiveElement != nil {
		if r.GenomeBuild == nil {
			r.GenomeBuild = &String{}
		}
		r.GenomeBuild.Id = m.GenomeBuildPrimitiveElement.Id
		r.GenomeBuild.Extension = m.GenomeBuildPrimitiveElement.Extension
	}
	r.Orientation = m.Orientation
	if m.OrientationPrimitiveElement != nil {
		if r.Orientation == nil {
			r.Orientation = &Code{}
		}
		r.Orientation.Id = m.OrientationPrimitiveElement.Id
		r.Orientation.Extension = m.OrientationPrimitiveElement.Extension
	}
	r.ReferenceSeqId = m.ReferenceSeqId
	r.ReferenceSeqPointer = m.ReferenceSeqPointer
	r.ReferenceSeqString = m.ReferenceSeqString
	if m.ReferenceSeqStringPrimitiveElement != nil {
		if r.ReferenceSeqString == nil {
			r.ReferenceSeqString = &String{}
		}
		r.ReferenceSeqString.Id = m.ReferenceSeqStringPrimitiveElement.Id
		r.ReferenceSeqString.Extension = m.ReferenceSeqStringPrimitiveElement.Extension
	}
	r.Strand = m.Strand
	if m.StrandPrimitiveElement != nil {
		if r.Strand == nil {
			r.Strand = &Code{}
		}
		r.Strand.Id = m.StrandPrimitiveElement.Id
		r.Strand.Extension = m.StrandPrimitiveElement.Extension
	}
	r.WindowStart = m.WindowStart
	if m.WindowStartPrimitiveElement != nil {
		if r.WindowStart == nil {
			r.WindowStart = &Integer{}
		}
		r.WindowStart.Id = m.WindowStartPrimitiveElement.Id
		r.WindowStart.Extension = m.WindowStartPrimitiveElement.Extension
	}
	r.WindowEnd = m.WindowEnd
	if m.WindowEndPrimitiveElement != nil {
		if r.WindowEnd == nil {
			r.WindowEnd = &Integer{}
		}
		r.WindowEnd.Id = m.WindowEndPrimitiveElement.Id
		r.WindowEnd.Extension = m.WindowEndPrimitiveElement.Extension
	}
	return nil
}
func (r MolecularSequenceReferenceSeq) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The definition of variant here originates from Sequence ontology ([variant_of](http://www.sequenceontology.org/browser/current_svn/term/variant_of)). This element can represent amino acid or nucleic sequence change(including insertion,deletion,SNP,etc.)  It can represent some complex mutation or segment variation with the assist of CIGAR string.
type MolecularSequenceVariant struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Start position of the variant on the  reference sequence. If the coordinate system is either 0-based or 1-based, then start position is inclusive.
	Start *Integer
	// End position of the variant on the reference sequence. If the coordinate system is 0-based then end is exclusive and does not include the last position. If the coordinate system is 1-base, then end is inclusive and includes the last position.
	End *Integer
	// An allele is one of a set of coexisting sequence variants of a gene ([SO:0001023](http://www.sequenceontology.org/browser/current_svn/term/SO:0001023)).  Nucleotide(s)/amino acids from start position of sequence to stop position of sequence on the positive (+) strand of the observed  sequence. When the sequence  type is DNA, it should be the sequence on the positive (+) strand. This will lay in the range between variant.start and variant.end.
	ObservedAllele *String
	// An allele is one of a set of coexisting sequence variants of a gene ([SO:0001023](http://www.sequenceontology.org/browser/current_svn/term/SO:0001023)). Nucleotide(s)/amino acids from start position of sequence to stop position of sequence on the positive (+) strand of the reference sequence. When the sequence  type is DNA, it should be the sequence on the positive (+) strand. This will lay in the range between variant.start and variant.end.
	ReferenceAllele *String
	// Extended CIGAR string for aligning the sequence with reference bases. See detailed documentation [here](http://support.illumina.com/help/SequencingAnalysisWorkflow/Content/Vault/Informatics/Sequencing_Analysis/CASAVA/swSEQ_mCA_ExtendedCIGARFormat.htm).
	Cigar *String
	// A pointer to an Observation containing variant information.
	VariantPointer *Reference
}
type jsonMolecularSequenceVariant struct {
	Id                              *string           `json:"id,omitempty"`
	Extension                       []Extension       `json:"extension,omitempty"`
	ModifierExtension               []Extension       `json:"modifierExtension,omitempty"`
	Start                           *Integer          `json:"start,omitempty"`
	StartPrimitiveElement           *primitiveElement `json:"_start,omitempty"`
	End                             *Integer          `json:"end,omitempty"`
	EndPrimitiveElement             *primitiveElement `json:"_end,omitempty"`
	ObservedAllele                  *String           `json:"observedAllele,omitempty"`
	ObservedAllelePrimitiveElement  *primitiveElement `json:"_observedAllele,omitempty"`
	ReferenceAllele                 *String           `json:"referenceAllele,omitempty"`
	ReferenceAllelePrimitiveElement *primitiveElement `json:"_referenceAllele,omitempty"`
	Cigar                           *String           `json:"cigar,omitempty"`
	CigarPrimitiveElement           *primitiveElement `json:"_cigar,omitempty"`
	VariantPointer                  *Reference        `json:"variantPointer,omitempty"`
}

func (r MolecularSequenceVariant) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MolecularSequenceVariant) marshalJSON() jsonMolecularSequenceVariant {
	m := jsonMolecularSequenceVariant{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Start != nil && r.Start.Value != nil {
		m.Start = r.Start
	}
	if r.Start != nil && (r.Start.Id != nil || r.Start.Extension != nil) {
		m.StartPrimitiveElement = &primitiveElement{Id: r.Start.Id, Extension: r.Start.Extension}
	}
	if r.End != nil && r.End.Value != nil {
		m.End = r.End
	}
	if r.End != nil && (r.End.Id != nil || r.End.Extension != nil) {
		m.EndPrimitiveElement = &primitiveElement{Id: r.End.Id, Extension: r.End.Extension}
	}
	if r.ObservedAllele != nil && r.ObservedAllele.Value != nil {
		m.ObservedAllele = r.ObservedAllele
	}
	if r.ObservedAllele != nil && (r.ObservedAllele.Id != nil || r.ObservedAllele.Extension != nil) {
		m.ObservedAllelePrimitiveElement = &primitiveElement{Id: r.ObservedAllele.Id, Extension: r.ObservedAllele.Extension}
	}
	if r.ReferenceAllele != nil && r.ReferenceAllele.Value != nil {
		m.ReferenceAllele = r.ReferenceAllele
	}
	if r.ReferenceAllele != nil && (r.ReferenceAllele.Id != nil || r.ReferenceAllele.Extension != nil) {
		m.ReferenceAllelePrimitiveElement = &primitiveElement{Id: r.ReferenceAllele.Id, Extension: r.ReferenceAllele.Extension}
	}
	if r.Cigar != nil && r.Cigar.Value != nil {
		m.Cigar = r.Cigar
	}
	if r.Cigar != nil && (r.Cigar.Id != nil || r.Cigar.Extension != nil) {
		m.CigarPrimitiveElement = &primitiveElement{Id: r.Cigar.Id, Extension: r.Cigar.Extension}
	}
	m.VariantPointer = r.VariantPointer
	return m
}
func (r *MolecularSequenceVariant) UnmarshalJSON(b []byte) error {
	var m jsonMolecularSequenceVariant
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MolecularSequenceVariant) unmarshalJSON(m jsonMolecularSequenceVariant) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Start = m.Start
	if m.StartPrimitiveElement != nil {
		if r.Start == nil {
			r.Start = &Integer{}
		}
		r.Start.Id = m.StartPrimitiveElement.Id
		r.Start.Extension = m.StartPrimitiveElement.Extension
	}
	r.End = m.End
	if m.EndPrimitiveElement != nil {
		if r.End == nil {
			r.End = &Integer{}
		}
		r.End.Id = m.EndPrimitiveElement.Id
		r.End.Extension = m.EndPrimitiveElement.Extension
	}
	r.ObservedAllele = m.ObservedAllele
	if m.ObservedAllelePrimitiveElement != nil {
		if r.ObservedAllele == nil {
			r.ObservedAllele = &String{}
		}
		r.ObservedAllele.Id = m.ObservedAllelePrimitiveElement.Id
		r.ObservedAllele.Extension = m.ObservedAllelePrimitiveElement.Extension
	}
	r.ReferenceAllele = m.ReferenceAllele
	if m.ReferenceAllelePrimitiveElement != nil {
		if r.ReferenceAllele == nil {
			r.ReferenceAllele = &String{}
		}
		r.ReferenceAllele.Id = m.ReferenceAllelePrimitiveElement.Id
		r.ReferenceAllele.Extension = m.ReferenceAllelePrimitiveElement.Extension
	}
	r.Cigar = m.Cigar
	if m.CigarPrimitiveElement != nil {
		if r.Cigar == nil {
			r.Cigar = &String{}
		}
		r.Cigar.Id = m.CigarPrimitiveElement.Id
		r.Cigar.Extension = m.CigarPrimitiveElement.Extension
	}
	r.VariantPointer = m.VariantPointer
	return nil
}
func (r MolecularSequenceVariant) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// An experimental feature attribute that defines the quality of the feature in a quantitative way, such as a phred quality score ([SO:0001686](http://www.sequenceontology.org/browser/current_svn/term/SO:0001686)).
type MolecularSequenceQuality struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// INDEL / SNP / Undefined variant.
	Type Code
	// Gold standard sequence used for comparing against.
	StandardSequence *CodeableConcept
	// Start position of the sequence. If the coordinate system is either 0-based or 1-based, then start position is inclusive.
	Start *Integer
	// End position of the sequence. If the coordinate system is 0-based then end is exclusive and does not include the last position. If the coordinate system is 1-base, then end is inclusive and includes the last position.
	End *Integer
	// The score of an experimentally derived feature such as a p-value ([SO:0001685](http://www.sequenceontology.org/browser/current_svn/term/SO:0001685)).
	Score *Quantity
	// Which method is used to get sequence quality.
	Method *CodeableConcept
	// True positives, from the perspective of the truth data, i.e. the number of sites in the Truth Call Set for which there are paths through the Query Call Set that are consistent with all of the alleles at this site, and for which there is an accurate genotype call for the event.
	TruthTp *Decimal
	// True positives, from the perspective of the query data, i.e. the number of sites in the Query Call Set for which there are paths through the Truth Call Set that are consistent with all of the alleles at this site, and for which there is an accurate genotype call for the event.
	QueryTp *Decimal
	// False negatives, i.e. the number of sites in the Truth Call Set for which there is no path through the Query Call Set that is consistent with all of the alleles at this site, or sites for which there is an inaccurate genotype call for the event. Sites with correct variant but incorrect genotype are counted here.
	TruthFn *Decimal
	// False positives, i.e. the number of sites in the Query Call Set for which there is no path through the Truth Call Set that is consistent with this site. Sites with correct variant but incorrect genotype are counted here.
	QueryFp *Decimal
	// The number of false positives where the non-REF alleles in the Truth and Query Call Sets match (i.e. cases where the truth is 1/1 and the query is 0/1 or similar).
	GtFp *Decimal
	// QUERY.TP / (QUERY.TP + QUERY.FP).
	Precision *Decimal
	// TRUTH.TP / (TRUTH.TP + TRUTH.FN).
	Recall *Decimal
	// Harmonic mean of Recall and Precision, computed as: 2 * precision * recall / (precision + recall).
	FScore *Decimal
	// Receiver Operator Characteristic (ROC) Curve  to give sensitivity/specificity tradeoff.
	Roc *MolecularSequenceQualityRoc
}
type jsonMolecularSequenceQuality struct {
	Id                        *string                      `json:"id,omitempty"`
	Extension                 []Extension                  `json:"extension,omitempty"`
	ModifierExtension         []Extension                  `json:"modifierExtension,omitempty"`
	Type                      Code                         `json:"type,omitempty"`
	TypePrimitiveElement      *primitiveElement            `json:"_type,omitempty"`
	StandardSequence          *CodeableConcept             `json:"standardSequence,omitempty"`
	Start                     *Integer                     `json:"start,omitempty"`
	StartPrimitiveElement     *primitiveElement            `json:"_start,omitempty"`
	End                       *Integer                     `json:"end,omitempty"`
	EndPrimitiveElement       *primitiveElement            `json:"_end,omitempty"`
	Score                     *Quantity                    `json:"score,omitempty"`
	Method                    *CodeableConcept             `json:"method,omitempty"`
	TruthTp                   *Decimal                     `json:"truthTP,omitempty"`
	TruthTpPrimitiveElement   *primitiveElement            `json:"_truthTP,omitempty"`
	QueryTp                   *Decimal                     `json:"queryTP,omitempty"`
	QueryTpPrimitiveElement   *primitiveElement            `json:"_queryTP,omitempty"`
	TruthFn                   *Decimal                     `json:"truthFN,omitempty"`
	TruthFnPrimitiveElement   *primitiveElement            `json:"_truthFN,omitempty"`
	QueryFp                   *Decimal                     `json:"queryFP,omitempty"`
	QueryFpPrimitiveElement   *primitiveElement            `json:"_queryFP,omitempty"`
	GtFp                      *Decimal                     `json:"gtFP,omitempty"`
	GtFpPrimitiveElement      *primitiveElement            `json:"_gtFP,omitempty"`
	Precision                 *Decimal                     `json:"precision,omitempty"`
	PrecisionPrimitiveElement *primitiveElement            `json:"_precision,omitempty"`
	Recall                    *Decimal                     `json:"recall,omitempty"`
	RecallPrimitiveElement    *primitiveElement            `json:"_recall,omitempty"`
	FScore                    *Decimal                     `json:"fScore,omitempty"`
	FScorePrimitiveElement    *primitiveElement            `json:"_fScore,omitempty"`
	Roc                       *MolecularSequenceQualityRoc `json:"roc,omitempty"`
}

func (r MolecularSequenceQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MolecularSequenceQuality) marshalJSON() jsonMolecularSequenceQuality {
	m := jsonMolecularSequenceQuality{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.StandardSequence = r.StandardSequence
	if r.Start != nil && r.Start.Value != nil {
		m.Start = r.Start
	}
	if r.Start != nil && (r.Start.Id != nil || r.Start.Extension != nil) {
		m.StartPrimitiveElement = &primitiveElement{Id: r.Start.Id, Extension: r.Start.Extension}
	}
	if r.End != nil && r.End.Value != nil {
		m.End = r.End
	}
	if r.End != nil && (r.End.Id != nil || r.End.Extension != nil) {
		m.EndPrimitiveElement = &primitiveElement{Id: r.End.Id, Extension: r.End.Extension}
	}
	m.Score = r.Score
	m.Method = r.Method
	if r.TruthTp != nil && r.TruthTp.Value != nil {
		m.TruthTp = r.TruthTp
	}
	if r.TruthTp != nil && (r.TruthTp.Id != nil || r.TruthTp.Extension != nil) {
		m.TruthTpPrimitiveElement = &primitiveElement{Id: r.TruthTp.Id, Extension: r.TruthTp.Extension}
	}
	if r.QueryTp != nil && r.QueryTp.Value != nil {
		m.QueryTp = r.QueryTp
	}
	if r.QueryTp != nil && (r.QueryTp.Id != nil || r.QueryTp.Extension != nil) {
		m.QueryTpPrimitiveElement = &primitiveElement{Id: r.QueryTp.Id, Extension: r.QueryTp.Extension}
	}
	if r.TruthFn != nil && r.TruthFn.Value != nil {
		m.TruthFn = r.TruthFn
	}
	if r.TruthFn != nil && (r.TruthFn.Id != nil || r.TruthFn.Extension != nil) {
		m.TruthFnPrimitiveElement = &primitiveElement{Id: r.TruthFn.Id, Extension: r.TruthFn.Extension}
	}
	if r.QueryFp != nil && r.QueryFp.Value != nil {
		m.QueryFp = r.QueryFp
	}
	if r.QueryFp != nil && (r.QueryFp.Id != nil || r.QueryFp.Extension != nil) {
		m.QueryFpPrimitiveElement = &primitiveElement{Id: r.QueryFp.Id, Extension: r.QueryFp.Extension}
	}
	if r.GtFp != nil && r.GtFp.Value != nil {
		m.GtFp = r.GtFp
	}
	if r.GtFp != nil && (r.GtFp.Id != nil || r.GtFp.Extension != nil) {
		m.GtFpPrimitiveElement = &primitiveElement{Id: r.GtFp.Id, Extension: r.GtFp.Extension}
	}
	if r.Precision != nil && r.Precision.Value != nil {
		m.Precision = r.Precision
	}
	if r.Precision != nil && (r.Precision.Id != nil || r.Precision.Extension != nil) {
		m.PrecisionPrimitiveElement = &primitiveElement{Id: r.Precision.Id, Extension: r.Precision.Extension}
	}
	if r.Recall != nil && r.Recall.Value != nil {
		m.Recall = r.Recall
	}
	if r.Recall != nil && (r.Recall.Id != nil || r.Recall.Extension != nil) {
		m.RecallPrimitiveElement = &primitiveElement{Id: r.Recall.Id, Extension: r.Recall.Extension}
	}
	if r.FScore != nil && r.FScore.Value != nil {
		m.FScore = r.FScore
	}
	if r.FScore != nil && (r.FScore.Id != nil || r.FScore.Extension != nil) {
		m.FScorePrimitiveElement = &primitiveElement{Id: r.FScore.Id, Extension: r.FScore.Extension}
	}
	m.Roc = r.Roc
	return m
}
func (r *MolecularSequenceQuality) UnmarshalJSON(b []byte) error {
	var m jsonMolecularSequenceQuality
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MolecularSequenceQuality) unmarshalJSON(m jsonMolecularSequenceQuality) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.StandardSequence = m.StandardSequence
	r.Start = m.Start
	if m.StartPrimitiveElement != nil {
		if r.Start == nil {
			r.Start = &Integer{}
		}
		r.Start.Id = m.StartPrimitiveElement.Id
		r.Start.Extension = m.StartPrimitiveElement.Extension
	}
	r.End = m.End
	if m.EndPrimitiveElement != nil {
		if r.End == nil {
			r.End = &Integer{}
		}
		r.End.Id = m.EndPrimitiveElement.Id
		r.End.Extension = m.EndPrimitiveElement.Extension
	}
	r.Score = m.Score
	r.Method = m.Method
	r.TruthTp = m.TruthTp
	if m.TruthTpPrimitiveElement != nil {
		if r.TruthTp == nil {
			r.TruthTp = &Decimal{}
		}
		r.TruthTp.Id = m.TruthTpPrimitiveElement.Id
		r.TruthTp.Extension = m.TruthTpPrimitiveElement.Extension
	}
	r.QueryTp = m.QueryTp
	if m.QueryTpPrimitiveElement != nil {
		if r.QueryTp == nil {
			r.QueryTp = &Decimal{}
		}
		r.QueryTp.Id = m.QueryTpPrimitiveElement.Id
		r.QueryTp.Extension = m.QueryTpPrimitiveElement.Extension
	}
	r.TruthFn = m.TruthFn
	if m.TruthFnPrimitiveElement != nil {
		if r.TruthFn == nil {
			r.TruthFn = &Decimal{}
		}
		r.TruthFn.Id = m.TruthFnPrimitiveElement.Id
		r.TruthFn.Extension = m.TruthFnPrimitiveElement.Extension
	}
	r.QueryFp = m.QueryFp
	if m.QueryFpPrimitiveElement != nil {
		if r.QueryFp == nil {
			r.QueryFp = &Decimal{}
		}
		r.QueryFp.Id = m.QueryFpPrimitiveElement.Id
		r.QueryFp.Extension = m.QueryFpPrimitiveElement.Extension
	}
	r.GtFp = m.GtFp
	if m.GtFpPrimitiveElement != nil {
		if r.GtFp == nil {
			r.GtFp = &Decimal{}
		}
		r.GtFp.Id = m.GtFpPrimitiveElement.Id
		r.GtFp.Extension = m.GtFpPrimitiveElement.Extension
	}
	r.Precision = m.Precision
	if m.PrecisionPrimitiveElement != nil {
		if r.Precision == nil {
			r.Precision = &Decimal{}
		}
		r.Precision.Id = m.PrecisionPrimitiveElement.Id
		r.Precision.Extension = m.PrecisionPrimitiveElement.Extension
	}
	r.Recall = m.Recall
	if m.RecallPrimitiveElement != nil {
		if r.Recall == nil {
			r.Recall = &Decimal{}
		}
		r.Recall.Id = m.RecallPrimitiveElement.Id
		r.Recall.Extension = m.RecallPrimitiveElement.Extension
	}
	r.FScore = m.FScore
	if m.FScorePrimitiveElement != nil {
		if r.FScore == nil {
			r.FScore = &Decimal{}
		}
		r.FScore.Id = m.FScorePrimitiveElement.Id
		r.FScore.Extension = m.FScorePrimitiveElement.Extension
	}
	r.Roc = m.Roc
	return nil
}
func (r MolecularSequenceQuality) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Receiver Operator Characteristic (ROC) Curve  to give sensitivity/specificity tradeoff.
type MolecularSequenceQualityRoc struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Invidual data point representing the GQ (genotype quality) score threshold.
	Score []Integer
	// The number of true positives if the GQ score threshold was set to "score" field value.
	NumTp []Integer
	// The number of false positives if the GQ score threshold was set to "score" field value.
	NumFp []Integer
	// The number of false negatives if the GQ score threshold was set to "score" field value.
	NumFn []Integer
	// Calculated precision if the GQ score threshold was set to "score" field value.
	Precision []Decimal
	// Calculated sensitivity if the GQ score threshold was set to "score" field value.
	Sensitivity []Decimal
	// Calculated fScore if the GQ score threshold was set to "score" field value.
	FMeasure []Decimal
}
type jsonMolecularSequenceQualityRoc struct {
	Id                          *string             `json:"id,omitempty"`
	Extension                   []Extension         `json:"extension,omitempty"`
	ModifierExtension           []Extension         `json:"modifierExtension,omitempty"`
	Score                       []Integer           `json:"score,omitempty"`
	ScorePrimitiveElement       []*primitiveElement `json:"_score,omitempty"`
	NumTp                       []Integer           `json:"numTP,omitempty"`
	NumTpPrimitiveElement       []*primitiveElement `json:"_numTP,omitempty"`
	NumFp                       []Integer           `json:"numFP,omitempty"`
	NumFpPrimitiveElement       []*primitiveElement `json:"_numFP,omitempty"`
	NumFn                       []Integer           `json:"numFN,omitempty"`
	NumFnPrimitiveElement       []*primitiveElement `json:"_numFN,omitempty"`
	Precision                   []Decimal           `json:"precision,omitempty"`
	PrecisionPrimitiveElement   []*primitiveElement `json:"_precision,omitempty"`
	Sensitivity                 []Decimal           `json:"sensitivity,omitempty"`
	SensitivityPrimitiveElement []*primitiveElement `json:"_sensitivity,omitempty"`
	FMeasure                    []Decimal           `json:"fMeasure,omitempty"`
	FMeasurePrimitiveElement    []*primitiveElement `json:"_fMeasure,omitempty"`
}

func (r MolecularSequenceQualityRoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MolecularSequenceQualityRoc) marshalJSON() jsonMolecularSequenceQualityRoc {
	m := jsonMolecularSequenceQualityRoc{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	anyScoreValue := false
	for _, e := range r.Score {
		if e.Value != nil {
			anyScoreValue = true
			break
		}
	}
	if anyScoreValue {
		m.Score = r.Score
	}
	anyScoreIdOrExtension := false
	for _, e := range r.Score {
		if e.Id != nil || e.Extension != nil {
			anyScoreIdOrExtension = true
			break
		}
	}
	if anyScoreIdOrExtension {
		m.ScorePrimitiveElement = make([]*primitiveElement, 0, len(r.Score))
		for _, e := range r.Score {
			if e.Id != nil || e.Extension != nil {
				m.ScorePrimitiveElement = append(m.ScorePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ScorePrimitiveElement = append(m.ScorePrimitiveElement, nil)
			}
		}
	}
	anyNumTpValue := false
	for _, e := range r.NumTp {
		if e.Value != nil {
			anyNumTpValue = true
			break
		}
	}
	if anyNumTpValue {
		m.NumTp = r.NumTp
	}
	anyNumTpIdOrExtension := false
	for _, e := range r.NumTp {
		if e.Id != nil || e.Extension != nil {
			anyNumTpIdOrExtension = true
			break
		}
	}
	if anyNumTpIdOrExtension {
		m.NumTpPrimitiveElement = make([]*primitiveElement, 0, len(r.NumTp))
		for _, e := range r.NumTp {
			if e.Id != nil || e.Extension != nil {
				m.NumTpPrimitiveElement = append(m.NumTpPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NumTpPrimitiveElement = append(m.NumTpPrimitiveElement, nil)
			}
		}
	}
	anyNumFpValue := false
	for _, e := range r.NumFp {
		if e.Value != nil {
			anyNumFpValue = true
			break
		}
	}
	if anyNumFpValue {
		m.NumFp = r.NumFp
	}
	anyNumFpIdOrExtension := false
	for _, e := range r.NumFp {
		if e.Id != nil || e.Extension != nil {
			anyNumFpIdOrExtension = true
			break
		}
	}
	if anyNumFpIdOrExtension {
		m.NumFpPrimitiveElement = make([]*primitiveElement, 0, len(r.NumFp))
		for _, e := range r.NumFp {
			if e.Id != nil || e.Extension != nil {
				m.NumFpPrimitiveElement = append(m.NumFpPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NumFpPrimitiveElement = append(m.NumFpPrimitiveElement, nil)
			}
		}
	}
	anyNumFnValue := false
	for _, e := range r.NumFn {
		if e.Value != nil {
			anyNumFnValue = true
			break
		}
	}
	if anyNumFnValue {
		m.NumFn = r.NumFn
	}
	anyNumFnIdOrExtension := false
	for _, e := range r.NumFn {
		if e.Id != nil || e.Extension != nil {
			anyNumFnIdOrExtension = true
			break
		}
	}
	if anyNumFnIdOrExtension {
		m.NumFnPrimitiveElement = make([]*primitiveElement, 0, len(r.NumFn))
		for _, e := range r.NumFn {
			if e.Id != nil || e.Extension != nil {
				m.NumFnPrimitiveElement = append(m.NumFnPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NumFnPrimitiveElement = append(m.NumFnPrimitiveElement, nil)
			}
		}
	}
	anyPrecisionValue := false
	for _, e := range r.Precision {
		if e.Value != nil {
			anyPrecisionValue = true
			break
		}
	}
	if anyPrecisionValue {
		m.Precision = r.Precision
	}
	anyPrecisionIdOrExtension := false
	for _, e := range r.Precision {
		if e.Id != nil || e.Extension != nil {
			anyPrecisionIdOrExtension = true
			break
		}
	}
	if anyPrecisionIdOrExtension {
		m.PrecisionPrimitiveElement = make([]*primitiveElement, 0, len(r.Precision))
		for _, e := range r.Precision {
			if e.Id != nil || e.Extension != nil {
				m.PrecisionPrimitiveElement = append(m.PrecisionPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.PrecisionPrimitiveElement = append(m.PrecisionPrimitiveElement, nil)
			}
		}
	}
	anySensitivityValue := false
	for _, e := range r.Sensitivity {
		if e.Value != nil {
			anySensitivityValue = true
			break
		}
	}
	if anySensitivityValue {
		m.Sensitivity = r.Sensitivity
	}
	anySensitivityIdOrExtension := false
	for _, e := range r.Sensitivity {
		if e.Id != nil || e.Extension != nil {
			anySensitivityIdOrExtension = true
			break
		}
	}
	if anySensitivityIdOrExtension {
		m.SensitivityPrimitiveElement = make([]*primitiveElement, 0, len(r.Sensitivity))
		for _, e := range r.Sensitivity {
			if e.Id != nil || e.Extension != nil {
				m.SensitivityPrimitiveElement = append(m.SensitivityPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SensitivityPrimitiveElement = append(m.SensitivityPrimitiveElement, nil)
			}
		}
	}
	anyFMeasureValue := false
	for _, e := range r.FMeasure {
		if e.Value != nil {
			anyFMeasureValue = true
			break
		}
	}
	if anyFMeasureValue {
		m.FMeasure = r.FMeasure
	}
	anyFMeasureIdOrExtension := false
	for _, e := range r.FMeasure {
		if e.Id != nil || e.Extension != nil {
			anyFMeasureIdOrExtension = true
			break
		}
	}
	if anyFMeasureIdOrExtension {
		m.FMeasurePrimitiveElement = make([]*primitiveElement, 0, len(r.FMeasure))
		for _, e := range r.FMeasure {
			if e.Id != nil || e.Extension != nil {
				m.FMeasurePrimitiveElement = append(m.FMeasurePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.FMeasurePrimitiveElement = append(m.FMeasurePrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *MolecularSequenceQualityRoc) UnmarshalJSON(b []byte) error {
	var m jsonMolecularSequenceQualityRoc
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MolecularSequenceQualityRoc) unmarshalJSON(m jsonMolecularSequenceQualityRoc) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Score = m.Score
	for i, e := range m.ScorePrimitiveElement {
		if len(r.Score) <= i {
			r.Score = append(r.Score, Integer{})
		}
		if e != nil {
			r.Score[i].Id = e.Id
			r.Score[i].Extension = e.Extension
		}
	}
	r.NumTp = m.NumTp
	for i, e := range m.NumTpPrimitiveElement {
		if len(r.NumTp) <= i {
			r.NumTp = append(r.NumTp, Integer{})
		}
		if e != nil {
			r.NumTp[i].Id = e.Id
			r.NumTp[i].Extension = e.Extension
		}
	}
	r.NumFp = m.NumFp
	for i, e := range m.NumFpPrimitiveElement {
		if len(r.NumFp) <= i {
			r.NumFp = append(r.NumFp, Integer{})
		}
		if e != nil {
			r.NumFp[i].Id = e.Id
			r.NumFp[i].Extension = e.Extension
		}
	}
	r.NumFn = m.NumFn
	for i, e := range m.NumFnPrimitiveElement {
		if len(r.NumFn) <= i {
			r.NumFn = append(r.NumFn, Integer{})
		}
		if e != nil {
			r.NumFn[i].Id = e.Id
			r.NumFn[i].Extension = e.Extension
		}
	}
	r.Precision = m.Precision
	for i, e := range m.PrecisionPrimitiveElement {
		if len(r.Precision) <= i {
			r.Precision = append(r.Precision, Decimal{})
		}
		if e != nil {
			r.Precision[i].Id = e.Id
			r.Precision[i].Extension = e.Extension
		}
	}
	r.Sensitivity = m.Sensitivity
	for i, e := range m.SensitivityPrimitiveElement {
		if len(r.Sensitivity) <= i {
			r.Sensitivity = append(r.Sensitivity, Decimal{})
		}
		if e != nil {
			r.Sensitivity[i].Id = e.Id
			r.Sensitivity[i].Extension = e.Extension
		}
	}
	r.FMeasure = m.FMeasure
	for i, e := range m.FMeasurePrimitiveElement {
		if len(r.FMeasure) <= i {
			r.FMeasure = append(r.FMeasure, Decimal{})
		}
		if e != nil {
			r.FMeasure[i].Id = e.Id
			r.FMeasure[i].Extension = e.Extension
		}
	}
	return nil
}
func (r MolecularSequenceQualityRoc) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Configurations of the external repository. The repository shall store target's observedSeq or records related with target's observedSeq.
type MolecularSequenceRepository struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Click and see / RESTful API / Need login to see / RESTful API with authentication / Other ways to see resource.
	Type Code
	// URI of an external repository which contains further details about the genetics data.
	Url *Uri
	// URI of an external repository which contains further details about the genetics data.
	Name *String
	// Id of the variant in this external repository. The server will understand how to use this id to call for more info about datasets in external repository.
	DatasetId *String
	// Id of the variantset in this external repository. The server will understand how to use this id to call for more info about variantsets in external repository.
	VariantsetId *String
	// Id of the read in this external repository.
	ReadsetId *String
}
type jsonMolecularSequenceRepository struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Type                         Code              `json:"type,omitempty"`
	TypePrimitiveElement         *primitiveElement `json:"_type,omitempty"`
	Url                          *Uri              `json:"url,omitempty"`
	UrlPrimitiveElement          *primitiveElement `json:"_url,omitempty"`
	Name                         *String           `json:"name,omitempty"`
	NamePrimitiveElement         *primitiveElement `json:"_name,omitempty"`
	DatasetId                    *String           `json:"datasetId,omitempty"`
	DatasetIdPrimitiveElement    *primitiveElement `json:"_datasetId,omitempty"`
	VariantsetId                 *String           `json:"variantsetId,omitempty"`
	VariantsetIdPrimitiveElement *primitiveElement `json:"_variantsetId,omitempty"`
	ReadsetId                    *String           `json:"readsetId,omitempty"`
	ReadsetIdPrimitiveElement    *primitiveElement `json:"_readsetId,omitempty"`
}

func (r MolecularSequenceRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MolecularSequenceRepository) marshalJSON() jsonMolecularSequenceRepository {
	m := jsonMolecularSequenceRepository{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.Url != nil && r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.DatasetId != nil && r.DatasetId.Value != nil {
		m.DatasetId = r.DatasetId
	}
	if r.DatasetId != nil && (r.DatasetId.Id != nil || r.DatasetId.Extension != nil) {
		m.DatasetIdPrimitiveElement = &primitiveElement{Id: r.DatasetId.Id, Extension: r.DatasetId.Extension}
	}
	if r.VariantsetId != nil && r.VariantsetId.Value != nil {
		m.VariantsetId = r.VariantsetId
	}
	if r.VariantsetId != nil && (r.VariantsetId.Id != nil || r.VariantsetId.Extension != nil) {
		m.VariantsetIdPrimitiveElement = &primitiveElement{Id: r.VariantsetId.Id, Extension: r.VariantsetId.Extension}
	}
	if r.ReadsetId != nil && r.ReadsetId.Value != nil {
		m.ReadsetId = r.ReadsetId
	}
	if r.ReadsetId != nil && (r.ReadsetId.Id != nil || r.ReadsetId.Extension != nil) {
		m.ReadsetIdPrimitiveElement = &primitiveElement{Id: r.ReadsetId.Id, Extension: r.ReadsetId.Extension}
	}
	return m
}
func (r *MolecularSequenceRepository) UnmarshalJSON(b []byte) error {
	var m jsonMolecularSequenceRepository
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MolecularSequenceRepository) unmarshalJSON(m jsonMolecularSequenceRepository) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		if r.Url == nil {
			r.Url = &Uri{}
		}
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.DatasetId = m.DatasetId
	if m.DatasetIdPrimitiveElement != nil {
		if r.DatasetId == nil {
			r.DatasetId = &String{}
		}
		r.DatasetId.Id = m.DatasetIdPrimitiveElement.Id
		r.DatasetId.Extension = m.DatasetIdPrimitiveElement.Extension
	}
	r.VariantsetId = m.VariantsetId
	if m.VariantsetIdPrimitiveElement != nil {
		if r.VariantsetId == nil {
			r.VariantsetId = &String{}
		}
		r.VariantsetId.Id = m.VariantsetIdPrimitiveElement.Id
		r.VariantsetId.Extension = m.VariantsetIdPrimitiveElement.Extension
	}
	r.ReadsetId = m.ReadsetId
	if m.ReadsetIdPrimitiveElement != nil {
		if r.ReadsetId == nil {
			r.ReadsetId = &String{}
		}
		r.ReadsetId.Id = m.ReadsetIdPrimitiveElement.Id
		r.ReadsetId.Extension = m.ReadsetIdPrimitiveElement.Extension
	}
	return nil
}
func (r MolecularSequenceRepository) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Information about chromosome structure variation.
type MolecularSequenceStructureVariant struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Information about chromosome structure variation DNA change type.
	VariantType *CodeableConcept
	// Used to indicate if the outer and inner start-end values have the same meaning.
	Exact *Boolean
	// Length of the variant chromosome.
	Length *Integer
	// Structural variant outer.
	Outer *MolecularSequenceStructureVariantOuter
	// Structural variant inner.
	Inner *MolecularSequenceStructureVariantInner
}
type jsonMolecularSequenceStructureVariant struct {
	Id                     *string                                 `json:"id,omitempty"`
	Extension              []Extension                             `json:"extension,omitempty"`
	ModifierExtension      []Extension                             `json:"modifierExtension,omitempty"`
	VariantType            *CodeableConcept                        `json:"variantType,omitempty"`
	Exact                  *Boolean                                `json:"exact,omitempty"`
	ExactPrimitiveElement  *primitiveElement                       `json:"_exact,omitempty"`
	Length                 *Integer                                `json:"length,omitempty"`
	LengthPrimitiveElement *primitiveElement                       `json:"_length,omitempty"`
	Outer                  *MolecularSequenceStructureVariantOuter `json:"outer,omitempty"`
	Inner                  *MolecularSequenceStructureVariantInner `json:"inner,omitempty"`
}

func (r MolecularSequenceStructureVariant) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MolecularSequenceStructureVariant) marshalJSON() jsonMolecularSequenceStructureVariant {
	m := jsonMolecularSequenceStructureVariant{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.VariantType = r.VariantType
	if r.Exact != nil && r.Exact.Value != nil {
		m.Exact = r.Exact
	}
	if r.Exact != nil && (r.Exact.Id != nil || r.Exact.Extension != nil) {
		m.ExactPrimitiveElement = &primitiveElement{Id: r.Exact.Id, Extension: r.Exact.Extension}
	}
	if r.Length != nil && r.Length.Value != nil {
		m.Length = r.Length
	}
	if r.Length != nil && (r.Length.Id != nil || r.Length.Extension != nil) {
		m.LengthPrimitiveElement = &primitiveElement{Id: r.Length.Id, Extension: r.Length.Extension}
	}
	m.Outer = r.Outer
	m.Inner = r.Inner
	return m
}
func (r *MolecularSequenceStructureVariant) UnmarshalJSON(b []byte) error {
	var m jsonMolecularSequenceStructureVariant
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MolecularSequenceStructureVariant) unmarshalJSON(m jsonMolecularSequenceStructureVariant) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.VariantType = m.VariantType
	r.Exact = m.Exact
	if m.ExactPrimitiveElement != nil {
		if r.Exact == nil {
			r.Exact = &Boolean{}
		}
		r.Exact.Id = m.ExactPrimitiveElement.Id
		r.Exact.Extension = m.ExactPrimitiveElement.Extension
	}
	r.Length = m.Length
	if m.LengthPrimitiveElement != nil {
		if r.Length == nil {
			r.Length = &Integer{}
		}
		r.Length.Id = m.LengthPrimitiveElement.Id
		r.Length.Extension = m.LengthPrimitiveElement.Extension
	}
	r.Outer = m.Outer
	r.Inner = m.Inner
	return nil
}
func (r MolecularSequenceStructureVariant) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Structural variant outer.
type MolecularSequenceStructureVariantOuter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Structural variant outer start. If the coordinate system is either 0-based or 1-based, then start position is inclusive.
	Start *Integer
	// Structural variant outer end. If the coordinate system is 0-based then end is exclusive and does not include the last position. If the coordinate system is 1-base, then end is inclusive and includes the last position.
	End *Integer
}
type jsonMolecularSequenceStructureVariantOuter struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Start                 *Integer          `json:"start,omitempty"`
	StartPrimitiveElement *primitiveElement `json:"_start,omitempty"`
	End                   *Integer          `json:"end,omitempty"`
	EndPrimitiveElement   *primitiveElement `json:"_end,omitempty"`
}

func (r MolecularSequenceStructureVariantOuter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MolecularSequenceStructureVariantOuter) marshalJSON() jsonMolecularSequenceStructureVariantOuter {
	m := jsonMolecularSequenceStructureVariantOuter{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Start != nil && r.Start.Value != nil {
		m.Start = r.Start
	}
	if r.Start != nil && (r.Start.Id != nil || r.Start.Extension != nil) {
		m.StartPrimitiveElement = &primitiveElement{Id: r.Start.Id, Extension: r.Start.Extension}
	}
	if r.End != nil && r.End.Value != nil {
		m.End = r.End
	}
	if r.End != nil && (r.End.Id != nil || r.End.Extension != nil) {
		m.EndPrimitiveElement = &primitiveElement{Id: r.End.Id, Extension: r.End.Extension}
	}
	return m
}
func (r *MolecularSequenceStructureVariantOuter) UnmarshalJSON(b []byte) error {
	var m jsonMolecularSequenceStructureVariantOuter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MolecularSequenceStructureVariantOuter) unmarshalJSON(m jsonMolecularSequenceStructureVariantOuter) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Start = m.Start
	if m.StartPrimitiveElement != nil {
		if r.Start == nil {
			r.Start = &Integer{}
		}
		r.Start.Id = m.StartPrimitiveElement.Id
		r.Start.Extension = m.StartPrimitiveElement.Extension
	}
	r.End = m.End
	if m.EndPrimitiveElement != nil {
		if r.End == nil {
			r.End = &Integer{}
		}
		r.End.Id = m.EndPrimitiveElement.Id
		r.End.Extension = m.EndPrimitiveElement.Extension
	}
	return nil
}
func (r MolecularSequenceStructureVariantOuter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Structural variant inner.
type MolecularSequenceStructureVariantInner struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Structural variant inner start. If the coordinate system is either 0-based or 1-based, then start position is inclusive.
	Start *Integer
	// Structural variant inner end. If the coordinate system is 0-based then end is exclusive and does not include the last position. If the coordinate system is 1-base, then end is inclusive and includes the last position.
	End *Integer
}
type jsonMolecularSequenceStructureVariantInner struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Start                 *Integer          `json:"start,omitempty"`
	StartPrimitiveElement *primitiveElement `json:"_start,omitempty"`
	End                   *Integer          `json:"end,omitempty"`
	EndPrimitiveElement   *primitiveElement `json:"_end,omitempty"`
}

func (r MolecularSequenceStructureVariantInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MolecularSequenceStructureVariantInner) marshalJSON() jsonMolecularSequenceStructureVariantInner {
	m := jsonMolecularSequenceStructureVariantInner{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Start != nil && r.Start.Value != nil {
		m.Start = r.Start
	}
	if r.Start != nil && (r.Start.Id != nil || r.Start.Extension != nil) {
		m.StartPrimitiveElement = &primitiveElement{Id: r.Start.Id, Extension: r.Start.Extension}
	}
	if r.End != nil && r.End.Value != nil {
		m.End = r.End
	}
	if r.End != nil && (r.End.Id != nil || r.End.Extension != nil) {
		m.EndPrimitiveElement = &primitiveElement{Id: r.End.Id, Extension: r.End.Extension}
	}
	return m
}
func (r *MolecularSequenceStructureVariantInner) UnmarshalJSON(b []byte) error {
	var m jsonMolecularSequenceStructureVariantInner
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MolecularSequenceStructureVariantInner) unmarshalJSON(m jsonMolecularSequenceStructureVariantInner) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Start = m.Start
	if m.StartPrimitiveElement != nil {
		if r.Start == nil {
			r.Start = &Integer{}
		}
		r.Start.Id = m.StartPrimitiveElement.Id
		r.Start.Extension = m.StartPrimitiveElement.Extension
	}
	r.End = m.End
	if m.EndPrimitiveElement != nil {
		if r.End == nil {
			r.End = &Integer{}
		}
		r.End.Id = m.EndPrimitiveElement.Id
		r.End.Extension = m.EndPrimitiveElement.Extension
	}
	return nil
}
func (r MolecularSequenceStructureVariantInner) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
