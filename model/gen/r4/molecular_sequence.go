package r4

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
func (r MolecularSequence) MemSize() int {
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
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	if r.Type != nil {
		s += r.Type.MemSize()
	}
	s += r.CoordinateSystem.MemSize() - int(unsafe.Sizeof(r.CoordinateSystem))
	if r.Patient != nil {
		s += r.Patient.MemSize()
	}
	if r.Specimen != nil {
		s += r.Specimen.MemSize()
	}
	if r.Device != nil {
		s += r.Device.MemSize()
	}
	if r.Performer != nil {
		s += r.Performer.MemSize()
	}
	if r.Quantity != nil {
		s += r.Quantity.MemSize()
	}
	if r.ReferenceSeq != nil {
		s += r.ReferenceSeq.MemSize()
	}
	for _, i := range r.Variant {
		s += i.MemSize()
	}
	s += (cap(r.Variant) - len(r.Variant)) * int(unsafe.Sizeof(MolecularSequenceVariant{}))
	if r.ObservedSeq != nil {
		s += r.ObservedSeq.MemSize()
	}
	for _, i := range r.Quality {
		s += i.MemSize()
	}
	s += (cap(r.Quality) - len(r.Quality)) * int(unsafe.Sizeof(MolecularSequenceQuality{}))
	if r.ReadCoverage != nil {
		s += r.ReadCoverage.MemSize()
	}
	for _, i := range r.Repository {
		s += i.MemSize()
	}
	s += (cap(r.Repository) - len(r.Repository)) * int(unsafe.Sizeof(MolecularSequenceRepository{}))
	for _, i := range r.Pointer {
		s += i.MemSize()
	}
	s += (cap(r.Pointer) - len(r.Pointer)) * int(unsafe.Sizeof(Reference{}))
	for _, i := range r.StructureVariant {
		s += i.MemSize()
	}
	s += (cap(r.StructureVariant) - len(r.StructureVariant)) * int(unsafe.Sizeof(MolecularSequenceStructureVariant{}))
	return s
}
func (r MolecularSequenceReferenceSeq) MemSize() int {
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
	if r.Chromosome != nil {
		s += r.Chromosome.MemSize()
	}
	if r.GenomeBuild != nil {
		s += r.GenomeBuild.MemSize()
	}
	if r.Orientation != nil {
		s += r.Orientation.MemSize()
	}
	if r.ReferenceSeqId != nil {
		s += r.ReferenceSeqId.MemSize()
	}
	if r.ReferenceSeqPointer != nil {
		s += r.ReferenceSeqPointer.MemSize()
	}
	if r.ReferenceSeqString != nil {
		s += r.ReferenceSeqString.MemSize()
	}
	if r.Strand != nil {
		s += r.Strand.MemSize()
	}
	if r.WindowStart != nil {
		s += r.WindowStart.MemSize()
	}
	if r.WindowEnd != nil {
		s += r.WindowEnd.MemSize()
	}
	return s
}
func (r MolecularSequenceVariant) MemSize() int {
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
	if r.Start != nil {
		s += r.Start.MemSize()
	}
	if r.End != nil {
		s += r.End.MemSize()
	}
	if r.ObservedAllele != nil {
		s += r.ObservedAllele.MemSize()
	}
	if r.ReferenceAllele != nil {
		s += r.ReferenceAllele.MemSize()
	}
	if r.Cigar != nil {
		s += r.Cigar.MemSize()
	}
	if r.VariantPointer != nil {
		s += r.VariantPointer.MemSize()
	}
	return s
}
func (r MolecularSequenceQuality) MemSize() int {
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
	s += r.Type.MemSize() - int(unsafe.Sizeof(r.Type))
	if r.StandardSequence != nil {
		s += r.StandardSequence.MemSize()
	}
	if r.Start != nil {
		s += r.Start.MemSize()
	}
	if r.End != nil {
		s += r.End.MemSize()
	}
	if r.Score != nil {
		s += r.Score.MemSize()
	}
	if r.Method != nil {
		s += r.Method.MemSize()
	}
	if r.TruthTp != nil {
		s += r.TruthTp.MemSize()
	}
	if r.QueryTp != nil {
		s += r.QueryTp.MemSize()
	}
	if r.TruthFn != nil {
		s += r.TruthFn.MemSize()
	}
	if r.QueryFp != nil {
		s += r.QueryFp.MemSize()
	}
	if r.GtFp != nil {
		s += r.GtFp.MemSize()
	}
	if r.Precision != nil {
		s += r.Precision.MemSize()
	}
	if r.Recall != nil {
		s += r.Recall.MemSize()
	}
	if r.FScore != nil {
		s += r.FScore.MemSize()
	}
	if r.Roc != nil {
		s += r.Roc.MemSize()
	}
	return s
}
func (r MolecularSequenceQualityRoc) MemSize() int {
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
	for _, i := range r.Score {
		s += i.MemSize()
	}
	s += (cap(r.Score) - len(r.Score)) * int(unsafe.Sizeof(Integer{}))
	for _, i := range r.NumTp {
		s += i.MemSize()
	}
	s += (cap(r.NumTp) - len(r.NumTp)) * int(unsafe.Sizeof(Integer{}))
	for _, i := range r.NumFp {
		s += i.MemSize()
	}
	s += (cap(r.NumFp) - len(r.NumFp)) * int(unsafe.Sizeof(Integer{}))
	for _, i := range r.NumFn {
		s += i.MemSize()
	}
	s += (cap(r.NumFn) - len(r.NumFn)) * int(unsafe.Sizeof(Integer{}))
	for _, i := range r.Precision {
		s += i.MemSize()
	}
	s += (cap(r.Precision) - len(r.Precision)) * int(unsafe.Sizeof(Decimal{}))
	for _, i := range r.Sensitivity {
		s += i.MemSize()
	}
	s += (cap(r.Sensitivity) - len(r.Sensitivity)) * int(unsafe.Sizeof(Decimal{}))
	for _, i := range r.FMeasure {
		s += i.MemSize()
	}
	s += (cap(r.FMeasure) - len(r.FMeasure)) * int(unsafe.Sizeof(Decimal{}))
	return s
}
func (r MolecularSequenceRepository) MemSize() int {
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
	s += r.Type.MemSize() - int(unsafe.Sizeof(r.Type))
	if r.Url != nil {
		s += r.Url.MemSize()
	}
	if r.Name != nil {
		s += r.Name.MemSize()
	}
	if r.DatasetId != nil {
		s += r.DatasetId.MemSize()
	}
	if r.VariantsetId != nil {
		s += r.VariantsetId.MemSize()
	}
	if r.ReadsetId != nil {
		s += r.ReadsetId.MemSize()
	}
	return s
}
func (r MolecularSequenceStructureVariant) MemSize() int {
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
	if r.VariantType != nil {
		s += r.VariantType.MemSize()
	}
	if r.Exact != nil {
		s += r.Exact.MemSize()
	}
	if r.Length != nil {
		s += r.Length.MemSize()
	}
	if r.Outer != nil {
		s += r.Outer.MemSize()
	}
	if r.Inner != nil {
		s += r.Inner.MemSize()
	}
	return s
}
func (r MolecularSequenceStructureVariantOuter) MemSize() int {
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
	if r.Start != nil {
		s += r.Start.MemSize()
	}
	if r.End != nil {
		s += r.End.MemSize()
	}
	return s
}
func (r MolecularSequenceStructureVariantInner) MemSize() int {
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
	if r.Start != nil {
		s += r.Start.MemSize()
	}
	if r.End != nil {
		s += r.End.MemSize()
	}
	return s
}
func (r MolecularSequence) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r MolecularSequenceReferenceSeq) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r MolecularSequenceVariant) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r MolecularSequenceQuality) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r MolecularSequenceQualityRoc) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r MolecularSequenceRepository) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r MolecularSequenceStructureVariant) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r MolecularSequenceStructureVariantOuter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r MolecularSequenceStructureVariantInner) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r MolecularSequence) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r MolecularSequence) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"MolecularSequence\""))
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
	if len(r.Identifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"identifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Identifier {
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
	if r.Type != nil && r.Type.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"type\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Type)
		if err != nil {
			return err
		}
	}
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		p := primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_type\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"coordinateSystem\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.CoordinateSystem)
		if err != nil {
			return err
		}
	}
	if r.CoordinateSystem.Id != nil || r.CoordinateSystem.Extension != nil {
		p := primitiveElement{Id: r.CoordinateSystem.Id, Extension: r.CoordinateSystem.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_coordinateSystem\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Patient != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"patient\":"))
		if err != nil {
			return err
		}
		err = r.Patient.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Specimen != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"specimen\":"))
		if err != nil {
			return err
		}
		err = r.Specimen.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Device != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"device\":"))
		if err != nil {
			return err
		}
		err = r.Device.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Performer != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"performer\":"))
		if err != nil {
			return err
		}
		err = r.Performer.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Quantity != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"quantity\":"))
		if err != nil {
			return err
		}
		err = r.Quantity.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ReferenceSeq != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"referenceSeq\":"))
		if err != nil {
			return err
		}
		err = r.ReferenceSeq.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Variant) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"variant\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Variant {
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
	if r.ObservedSeq != nil && r.ObservedSeq.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"observedSeq\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ObservedSeq)
		if err != nil {
			return err
		}
	}
	if r.ObservedSeq != nil && (r.ObservedSeq.Id != nil || r.ObservedSeq.Extension != nil) {
		p := primitiveElement{Id: r.ObservedSeq.Id, Extension: r.ObservedSeq.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_observedSeq\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Quality) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"quality\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Quality {
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
	if r.ReadCoverage != nil && r.ReadCoverage.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"readCoverage\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ReadCoverage)
		if err != nil {
			return err
		}
	}
	if r.ReadCoverage != nil && (r.ReadCoverage.Id != nil || r.ReadCoverage.Extension != nil) {
		p := primitiveElement{Id: r.ReadCoverage.Id, Extension: r.ReadCoverage.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_readCoverage\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Repository) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"repository\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Repository {
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
	if len(r.Pointer) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"pointer\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Pointer {
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
	if len(r.StructureVariant) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"structureVariant\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.StructureVariant {
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
func (r MolecularSequenceReferenceSeq) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r MolecularSequenceReferenceSeq) marshalJSON(w io.Writer) error {
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
	if r.Chromosome != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"chromosome\":"))
		if err != nil {
			return err
		}
		err = r.Chromosome.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.GenomeBuild != nil && r.GenomeBuild.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"genomeBuild\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.GenomeBuild)
		if err != nil {
			return err
		}
	}
	if r.GenomeBuild != nil && (r.GenomeBuild.Id != nil || r.GenomeBuild.Extension != nil) {
		p := primitiveElement{Id: r.GenomeBuild.Id, Extension: r.GenomeBuild.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_genomeBuild\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Orientation != nil && r.Orientation.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"orientation\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Orientation)
		if err != nil {
			return err
		}
	}
	if r.Orientation != nil && (r.Orientation.Id != nil || r.Orientation.Extension != nil) {
		p := primitiveElement{Id: r.Orientation.Id, Extension: r.Orientation.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_orientation\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ReferenceSeqId != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"referenceSeqId\":"))
		if err != nil {
			return err
		}
		err = r.ReferenceSeqId.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ReferenceSeqPointer != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"referenceSeqPointer\":"))
		if err != nil {
			return err
		}
		err = r.ReferenceSeqPointer.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ReferenceSeqString != nil && r.ReferenceSeqString.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"referenceSeqString\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ReferenceSeqString)
		if err != nil {
			return err
		}
	}
	if r.ReferenceSeqString != nil && (r.ReferenceSeqString.Id != nil || r.ReferenceSeqString.Extension != nil) {
		p := primitiveElement{Id: r.ReferenceSeqString.Id, Extension: r.ReferenceSeqString.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_referenceSeqString\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Strand != nil && r.Strand.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"strand\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Strand)
		if err != nil {
			return err
		}
	}
	if r.Strand != nil && (r.Strand.Id != nil || r.Strand.Extension != nil) {
		p := primitiveElement{Id: r.Strand.Id, Extension: r.Strand.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_strand\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.WindowStart != nil && r.WindowStart.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"windowStart\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.WindowStart)
		if err != nil {
			return err
		}
	}
	if r.WindowStart != nil && (r.WindowStart.Id != nil || r.WindowStart.Extension != nil) {
		p := primitiveElement{Id: r.WindowStart.Id, Extension: r.WindowStart.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_windowStart\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.WindowEnd != nil && r.WindowEnd.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"windowEnd\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.WindowEnd)
		if err != nil {
			return err
		}
	}
	if r.WindowEnd != nil && (r.WindowEnd.Id != nil || r.WindowEnd.Extension != nil) {
		p := primitiveElement{Id: r.WindowEnd.Id, Extension: r.WindowEnd.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_windowEnd\":"))
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
func (r MolecularSequenceVariant) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r MolecularSequenceVariant) marshalJSON(w io.Writer) error {
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
	if r.Start != nil && r.Start.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"start\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Start)
		if err != nil {
			return err
		}
	}
	if r.Start != nil && (r.Start.Id != nil || r.Start.Extension != nil) {
		p := primitiveElement{Id: r.Start.Id, Extension: r.Start.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_start\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.End != nil && r.End.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"end\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.End)
		if err != nil {
			return err
		}
	}
	if r.End != nil && (r.End.Id != nil || r.End.Extension != nil) {
		p := primitiveElement{Id: r.End.Id, Extension: r.End.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_end\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ObservedAllele != nil && r.ObservedAllele.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"observedAllele\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ObservedAllele)
		if err != nil {
			return err
		}
	}
	if r.ObservedAllele != nil && (r.ObservedAllele.Id != nil || r.ObservedAllele.Extension != nil) {
		p := primitiveElement{Id: r.ObservedAllele.Id, Extension: r.ObservedAllele.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_observedAllele\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ReferenceAllele != nil && r.ReferenceAllele.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"referenceAllele\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ReferenceAllele)
		if err != nil {
			return err
		}
	}
	if r.ReferenceAllele != nil && (r.ReferenceAllele.Id != nil || r.ReferenceAllele.Extension != nil) {
		p := primitiveElement{Id: r.ReferenceAllele.Id, Extension: r.ReferenceAllele.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_referenceAllele\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Cigar != nil && r.Cigar.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"cigar\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Cigar)
		if err != nil {
			return err
		}
	}
	if r.Cigar != nil && (r.Cigar.Id != nil || r.Cigar.Extension != nil) {
		p := primitiveElement{Id: r.Cigar.Id, Extension: r.Cigar.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_cigar\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.VariantPointer != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"variantPointer\":"))
		if err != nil {
			return err
		}
		err = r.VariantPointer.marshalJSON(w)
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
func (r MolecularSequenceQuality) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r MolecularSequenceQuality) marshalJSON(w io.Writer) error {
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
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"type\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Type)
		if err != nil {
			return err
		}
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		p := primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_type\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.StandardSequence != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"standardSequence\":"))
		if err != nil {
			return err
		}
		err = r.StandardSequence.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Start != nil && r.Start.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"start\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Start)
		if err != nil {
			return err
		}
	}
	if r.Start != nil && (r.Start.Id != nil || r.Start.Extension != nil) {
		p := primitiveElement{Id: r.Start.Id, Extension: r.Start.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_start\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.End != nil && r.End.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"end\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.End)
		if err != nil {
			return err
		}
	}
	if r.End != nil && (r.End.Id != nil || r.End.Extension != nil) {
		p := primitiveElement{Id: r.End.Id, Extension: r.End.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_end\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Score != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"score\":"))
		if err != nil {
			return err
		}
		err = r.Score.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Method != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"method\":"))
		if err != nil {
			return err
		}
		err = r.Method.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.TruthTp != nil && r.TruthTp.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"truthTP\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.TruthTp)
		if err != nil {
			return err
		}
	}
	if r.TruthTp != nil && (r.TruthTp.Id != nil || r.TruthTp.Extension != nil) {
		p := primitiveElement{Id: r.TruthTp.Id, Extension: r.TruthTp.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_truthTP\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.QueryTp != nil && r.QueryTp.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"queryTP\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.QueryTp)
		if err != nil {
			return err
		}
	}
	if r.QueryTp != nil && (r.QueryTp.Id != nil || r.QueryTp.Extension != nil) {
		p := primitiveElement{Id: r.QueryTp.Id, Extension: r.QueryTp.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_queryTP\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.TruthFn != nil && r.TruthFn.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"truthFN\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.TruthFn)
		if err != nil {
			return err
		}
	}
	if r.TruthFn != nil && (r.TruthFn.Id != nil || r.TruthFn.Extension != nil) {
		p := primitiveElement{Id: r.TruthFn.Id, Extension: r.TruthFn.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_truthFN\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.QueryFp != nil && r.QueryFp.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"queryFP\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.QueryFp)
		if err != nil {
			return err
		}
	}
	if r.QueryFp != nil && (r.QueryFp.Id != nil || r.QueryFp.Extension != nil) {
		p := primitiveElement{Id: r.QueryFp.Id, Extension: r.QueryFp.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_queryFP\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.GtFp != nil && r.GtFp.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"gtFP\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.GtFp)
		if err != nil {
			return err
		}
	}
	if r.GtFp != nil && (r.GtFp.Id != nil || r.GtFp.Extension != nil) {
		p := primitiveElement{Id: r.GtFp.Id, Extension: r.GtFp.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_gtFP\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Precision != nil && r.Precision.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"precision\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Precision)
		if err != nil {
			return err
		}
	}
	if r.Precision != nil && (r.Precision.Id != nil || r.Precision.Extension != nil) {
		p := primitiveElement{Id: r.Precision.Id, Extension: r.Precision.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_precision\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Recall != nil && r.Recall.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"recall\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Recall)
		if err != nil {
			return err
		}
	}
	if r.Recall != nil && (r.Recall.Id != nil || r.Recall.Extension != nil) {
		p := primitiveElement{Id: r.Recall.Id, Extension: r.Recall.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_recall\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.FScore != nil && r.FScore.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"fScore\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.FScore)
		if err != nil {
			return err
		}
	}
	if r.FScore != nil && (r.FScore.Id != nil || r.FScore.Extension != nil) {
		p := primitiveElement{Id: r.FScore.Id, Extension: r.FScore.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_fScore\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Roc != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"roc\":"))
		if err != nil {
			return err
		}
		err = r.Roc.marshalJSON(w)
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
func (r MolecularSequenceQualityRoc) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r MolecularSequenceQualityRoc) marshalJSON(w io.Writer) error {
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
	{
		anyValue := false
		for _, e := range r.Score {
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
			_, err = w.Write([]byte("\"score\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.Score)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.Score {
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
			_, err = w.Write([]byte("\"_score\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.Score {
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
	{
		anyValue := false
		for _, e := range r.NumTp {
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
			_, err = w.Write([]byte("\"numTP\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.NumTp)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.NumTp {
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
			_, err = w.Write([]byte("\"_numTP\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.NumTp {
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
	{
		anyValue := false
		for _, e := range r.NumFp {
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
			_, err = w.Write([]byte("\"numFP\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.NumFp)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.NumFp {
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
			_, err = w.Write([]byte("\"_numFP\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.NumFp {
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
	{
		anyValue := false
		for _, e := range r.NumFn {
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
			_, err = w.Write([]byte("\"numFN\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.NumFn)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.NumFn {
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
			_, err = w.Write([]byte("\"_numFN\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.NumFn {
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
	{
		anyValue := false
		for _, e := range r.Precision {
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
			_, err = w.Write([]byte("\"precision\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.Precision)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.Precision {
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
			_, err = w.Write([]byte("\"_precision\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.Precision {
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
	{
		anyValue := false
		for _, e := range r.Sensitivity {
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
			_, err = w.Write([]byte("\"sensitivity\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.Sensitivity)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.Sensitivity {
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
			_, err = w.Write([]byte("\"_sensitivity\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.Sensitivity {
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
	{
		anyValue := false
		for _, e := range r.FMeasure {
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
			_, err = w.Write([]byte("\"fMeasure\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.FMeasure)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.FMeasure {
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
			_, err = w.Write([]byte("\"_fMeasure\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.FMeasure {
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
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r MolecularSequenceRepository) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r MolecularSequenceRepository) marshalJSON(w io.Writer) error {
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
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"type\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Type)
		if err != nil {
			return err
		}
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		p := primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_type\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Url != nil && r.Url.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"url\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Url)
		if err != nil {
			return err
		}
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		p := primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_url\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Name != nil && r.Name.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"name\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Name)
		if err != nil {
			return err
		}
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		p := primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_name\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.DatasetId != nil && r.DatasetId.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"datasetId\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.DatasetId)
		if err != nil {
			return err
		}
	}
	if r.DatasetId != nil && (r.DatasetId.Id != nil || r.DatasetId.Extension != nil) {
		p := primitiveElement{Id: r.DatasetId.Id, Extension: r.DatasetId.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_datasetId\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.VariantsetId != nil && r.VariantsetId.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"variantsetId\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.VariantsetId)
		if err != nil {
			return err
		}
	}
	if r.VariantsetId != nil && (r.VariantsetId.Id != nil || r.VariantsetId.Extension != nil) {
		p := primitiveElement{Id: r.VariantsetId.Id, Extension: r.VariantsetId.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_variantsetId\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ReadsetId != nil && r.ReadsetId.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"readsetId\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ReadsetId)
		if err != nil {
			return err
		}
	}
	if r.ReadsetId != nil && (r.ReadsetId.Id != nil || r.ReadsetId.Extension != nil) {
		p := primitiveElement{Id: r.ReadsetId.Id, Extension: r.ReadsetId.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_readsetId\":"))
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
func (r MolecularSequenceStructureVariant) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r MolecularSequenceStructureVariant) marshalJSON(w io.Writer) error {
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
	if r.VariantType != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"variantType\":"))
		if err != nil {
			return err
		}
		err = r.VariantType.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Exact != nil && r.Exact.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"exact\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Exact)
		if err != nil {
			return err
		}
	}
	if r.Exact != nil && (r.Exact.Id != nil || r.Exact.Extension != nil) {
		p := primitiveElement{Id: r.Exact.Id, Extension: r.Exact.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_exact\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Length != nil && r.Length.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"length\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Length)
		if err != nil {
			return err
		}
	}
	if r.Length != nil && (r.Length.Id != nil || r.Length.Extension != nil) {
		p := primitiveElement{Id: r.Length.Id, Extension: r.Length.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_length\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Outer != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"outer\":"))
		if err != nil {
			return err
		}
		err = r.Outer.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Inner != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"inner\":"))
		if err != nil {
			return err
		}
		err = r.Inner.marshalJSON(w)
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
func (r MolecularSequenceStructureVariantOuter) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r MolecularSequenceStructureVariantOuter) marshalJSON(w io.Writer) error {
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
	if r.Start != nil && r.Start.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"start\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Start)
		if err != nil {
			return err
		}
	}
	if r.Start != nil && (r.Start.Id != nil || r.Start.Extension != nil) {
		p := primitiveElement{Id: r.Start.Id, Extension: r.Start.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_start\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.End != nil && r.End.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"end\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.End)
		if err != nil {
			return err
		}
	}
	if r.End != nil && (r.End.Id != nil || r.End.Extension != nil) {
		p := primitiveElement{Id: r.End.Id, Extension: r.End.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_end\":"))
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
func (r MolecularSequenceStructureVariantInner) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r MolecularSequenceStructureVariantInner) marshalJSON(w io.Writer) error {
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
	if r.Start != nil && r.Start.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"start\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Start)
		if err != nil {
			return err
		}
	}
	if r.Start != nil && (r.Start.Id != nil || r.Start.Extension != nil) {
		p := primitiveElement{Id: r.Start.Id, Extension: r.Start.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_start\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.End != nil && r.End.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"end\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.End)
		if err != nil {
			return err
		}
	}
	if r.End != nil && (r.End.Id != nil || r.End.Extension != nil) {
		p := primitiveElement{Id: r.End.Id, Extension: r.End.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_end\":"))
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
func (r *MolecularSequence) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *MolecularSequence) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in MolecularSequence element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in MolecularSequence element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequence element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequence element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequence element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequence element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequence element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequence element", t)
			}
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequence element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequence element", t)
			}
		case "type":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Type == nil {
				r.Type = &Code{}
			}
			r.Type.Value = v.Value
		case "_type":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Type == nil {
				r.Type = &Code{}
			}
			r.Type.Id = v.Id
			r.Type.Extension = v.Extension
		case "coordinateSystem":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.CoordinateSystem.Value = v.Value
		case "_coordinateSystem":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.CoordinateSystem.Id = v.Id
			r.CoordinateSystem.Extension = v.Extension
		case "patient":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Patient = &v
		case "specimen":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Specimen = &v
		case "device":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Device = &v
		case "performer":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Performer = &v
		case "quantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Quantity = &v
		case "referenceSeq":
			var v MolecularSequenceReferenceSeq
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ReferenceSeq = &v
		case "variant":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequence element", t)
			}
			for d.More() {
				var v MolecularSequenceVariant
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Variant = append(r.Variant, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequence element", t)
			}
		case "observedSeq":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ObservedSeq == nil {
				r.ObservedSeq = &String{}
			}
			r.ObservedSeq.Value = v.Value
		case "_observedSeq":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ObservedSeq == nil {
				r.ObservedSeq = &String{}
			}
			r.ObservedSeq.Id = v.Id
			r.ObservedSeq.Extension = v.Extension
		case "quality":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequence element", t)
			}
			for d.More() {
				var v MolecularSequenceQuality
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Quality = append(r.Quality, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequence element", t)
			}
		case "readCoverage":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ReadCoverage == nil {
				r.ReadCoverage = &Integer{}
			}
			r.ReadCoverage.Value = v.Value
		case "_readCoverage":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ReadCoverage == nil {
				r.ReadCoverage = &Integer{}
			}
			r.ReadCoverage.Id = v.Id
			r.ReadCoverage.Extension = v.Extension
		case "repository":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequence element", t)
			}
			for d.More() {
				var v MolecularSequenceRepository
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Repository = append(r.Repository, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequence element", t)
			}
		case "pointer":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequence element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Pointer = append(r.Pointer, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequence element", t)
			}
		case "structureVariant":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequence element", t)
			}
			for d.More() {
				var v MolecularSequenceStructureVariant
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.StructureVariant = append(r.StructureVariant, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequence element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in MolecularSequence", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in MolecularSequence element", t)
	}
	return nil
}
func (r *MolecularSequenceReferenceSeq) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in MolecularSequenceReferenceSeq element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in MolecularSequenceReferenceSeq element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceReferenceSeq element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceReferenceSeq element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceReferenceSeq element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceReferenceSeq element", t)
			}
		case "chromosome":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Chromosome = &v
		case "genomeBuild":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.GenomeBuild == nil {
				r.GenomeBuild = &String{}
			}
			r.GenomeBuild.Value = v.Value
		case "_genomeBuild":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.GenomeBuild == nil {
				r.GenomeBuild = &String{}
			}
			r.GenomeBuild.Id = v.Id
			r.GenomeBuild.Extension = v.Extension
		case "orientation":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Orientation == nil {
				r.Orientation = &Code{}
			}
			r.Orientation.Value = v.Value
		case "_orientation":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Orientation == nil {
				r.Orientation = &Code{}
			}
			r.Orientation.Id = v.Id
			r.Orientation.Extension = v.Extension
		case "referenceSeqId":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ReferenceSeqId = &v
		case "referenceSeqPointer":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ReferenceSeqPointer = &v
		case "referenceSeqString":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ReferenceSeqString == nil {
				r.ReferenceSeqString = &String{}
			}
			r.ReferenceSeqString.Value = v.Value
		case "_referenceSeqString":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ReferenceSeqString == nil {
				r.ReferenceSeqString = &String{}
			}
			r.ReferenceSeqString.Id = v.Id
			r.ReferenceSeqString.Extension = v.Extension
		case "strand":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Strand == nil {
				r.Strand = &Code{}
			}
			r.Strand.Value = v.Value
		case "_strand":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Strand == nil {
				r.Strand = &Code{}
			}
			r.Strand.Id = v.Id
			r.Strand.Extension = v.Extension
		case "windowStart":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.WindowStart == nil {
				r.WindowStart = &Integer{}
			}
			r.WindowStart.Value = v.Value
		case "_windowStart":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.WindowStart == nil {
				r.WindowStart = &Integer{}
			}
			r.WindowStart.Id = v.Id
			r.WindowStart.Extension = v.Extension
		case "windowEnd":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.WindowEnd == nil {
				r.WindowEnd = &Integer{}
			}
			r.WindowEnd.Value = v.Value
		case "_windowEnd":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.WindowEnd == nil {
				r.WindowEnd = &Integer{}
			}
			r.WindowEnd.Id = v.Id
			r.WindowEnd.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in MolecularSequenceReferenceSeq", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in MolecularSequenceReferenceSeq element", t)
	}
	return nil
}
func (r *MolecularSequenceVariant) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in MolecularSequenceVariant element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in MolecularSequenceVariant element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceVariant element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceVariant element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceVariant element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceVariant element", t)
			}
		case "start":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Start == nil {
				r.Start = &Integer{}
			}
			r.Start.Value = v.Value
		case "_start":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Start == nil {
				r.Start = &Integer{}
			}
			r.Start.Id = v.Id
			r.Start.Extension = v.Extension
		case "end":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.End == nil {
				r.End = &Integer{}
			}
			r.End.Value = v.Value
		case "_end":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.End == nil {
				r.End = &Integer{}
			}
			r.End.Id = v.Id
			r.End.Extension = v.Extension
		case "observedAllele":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ObservedAllele == nil {
				r.ObservedAllele = &String{}
			}
			r.ObservedAllele.Value = v.Value
		case "_observedAllele":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ObservedAllele == nil {
				r.ObservedAllele = &String{}
			}
			r.ObservedAllele.Id = v.Id
			r.ObservedAllele.Extension = v.Extension
		case "referenceAllele":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ReferenceAllele == nil {
				r.ReferenceAllele = &String{}
			}
			r.ReferenceAllele.Value = v.Value
		case "_referenceAllele":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ReferenceAllele == nil {
				r.ReferenceAllele = &String{}
			}
			r.ReferenceAllele.Id = v.Id
			r.ReferenceAllele.Extension = v.Extension
		case "cigar":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Cigar == nil {
				r.Cigar = &String{}
			}
			r.Cigar.Value = v.Value
		case "_cigar":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Cigar == nil {
				r.Cigar = &String{}
			}
			r.Cigar.Id = v.Id
			r.Cigar.Extension = v.Extension
		case "variantPointer":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.VariantPointer = &v
		default:
			return fmt.Errorf("invalid field: %s in MolecularSequenceVariant", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in MolecularSequenceVariant element", t)
	}
	return nil
}
func (r *MolecularSequenceQuality) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in MolecularSequenceQuality element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in MolecularSequenceQuality element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQuality element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQuality element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQuality element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQuality element", t)
			}
		case "type":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Type.Value = v.Value
		case "_type":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type.Id = v.Id
			r.Type.Extension = v.Extension
		case "standardSequence":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.StandardSequence = &v
		case "start":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Start == nil {
				r.Start = &Integer{}
			}
			r.Start.Value = v.Value
		case "_start":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Start == nil {
				r.Start = &Integer{}
			}
			r.Start.Id = v.Id
			r.Start.Extension = v.Extension
		case "end":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.End == nil {
				r.End = &Integer{}
			}
			r.End.Value = v.Value
		case "_end":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.End == nil {
				r.End = &Integer{}
			}
			r.End.Id = v.Id
			r.End.Extension = v.Extension
		case "score":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Score = &v
		case "method":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Method = &v
		case "truthTP":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.TruthTp == nil {
				r.TruthTp = &Decimal{}
			}
			r.TruthTp.Value = v.Value
		case "_truthTP":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.TruthTp == nil {
				r.TruthTp = &Decimal{}
			}
			r.TruthTp.Id = v.Id
			r.TruthTp.Extension = v.Extension
		case "queryTP":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.QueryTp == nil {
				r.QueryTp = &Decimal{}
			}
			r.QueryTp.Value = v.Value
		case "_queryTP":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.QueryTp == nil {
				r.QueryTp = &Decimal{}
			}
			r.QueryTp.Id = v.Id
			r.QueryTp.Extension = v.Extension
		case "truthFN":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.TruthFn == nil {
				r.TruthFn = &Decimal{}
			}
			r.TruthFn.Value = v.Value
		case "_truthFN":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.TruthFn == nil {
				r.TruthFn = &Decimal{}
			}
			r.TruthFn.Id = v.Id
			r.TruthFn.Extension = v.Extension
		case "queryFP":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.QueryFp == nil {
				r.QueryFp = &Decimal{}
			}
			r.QueryFp.Value = v.Value
		case "_queryFP":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.QueryFp == nil {
				r.QueryFp = &Decimal{}
			}
			r.QueryFp.Id = v.Id
			r.QueryFp.Extension = v.Extension
		case "gtFP":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.GtFp == nil {
				r.GtFp = &Decimal{}
			}
			r.GtFp.Value = v.Value
		case "_gtFP":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.GtFp == nil {
				r.GtFp = &Decimal{}
			}
			r.GtFp.Id = v.Id
			r.GtFp.Extension = v.Extension
		case "precision":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Precision == nil {
				r.Precision = &Decimal{}
			}
			r.Precision.Value = v.Value
		case "_precision":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Precision == nil {
				r.Precision = &Decimal{}
			}
			r.Precision.Id = v.Id
			r.Precision.Extension = v.Extension
		case "recall":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Recall == nil {
				r.Recall = &Decimal{}
			}
			r.Recall.Value = v.Value
		case "_recall":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Recall == nil {
				r.Recall = &Decimal{}
			}
			r.Recall.Id = v.Id
			r.Recall.Extension = v.Extension
		case "fScore":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.FScore == nil {
				r.FScore = &Decimal{}
			}
			r.FScore.Value = v.Value
		case "_fScore":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.FScore == nil {
				r.FScore = &Decimal{}
			}
			r.FScore.Id = v.Id
			r.FScore.Extension = v.Extension
		case "roc":
			var v MolecularSequenceQualityRoc
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Roc = &v
		default:
			return fmt.Errorf("invalid field: %s in MolecularSequenceQuality", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in MolecularSequenceQuality element", t)
	}
	return nil
}
func (r *MolecularSequenceQualityRoc) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in MolecularSequenceQualityRoc element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in MolecularSequenceQualityRoc element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "score":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
			}
			for i := 0; d.More(); i++ {
				var v Integer
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.Score) <= i {
					r.Score = append(r.Score, Integer{})
				}
				r.Score[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "_score":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.Score) <= i {
					r.Score = append(r.Score, Integer{})
				}
				r.Score[i].Id = v.Id
				r.Score[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "numTP":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
			}
			for i := 0; d.More(); i++ {
				var v Integer
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.NumTp) <= i {
					r.NumTp = append(r.NumTp, Integer{})
				}
				r.NumTp[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "_numTP":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.NumTp) <= i {
					r.NumTp = append(r.NumTp, Integer{})
				}
				r.NumTp[i].Id = v.Id
				r.NumTp[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "numFP":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
			}
			for i := 0; d.More(); i++ {
				var v Integer
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.NumFp) <= i {
					r.NumFp = append(r.NumFp, Integer{})
				}
				r.NumFp[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "_numFP":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.NumFp) <= i {
					r.NumFp = append(r.NumFp, Integer{})
				}
				r.NumFp[i].Id = v.Id
				r.NumFp[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "numFN":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
			}
			for i := 0; d.More(); i++ {
				var v Integer
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.NumFn) <= i {
					r.NumFn = append(r.NumFn, Integer{})
				}
				r.NumFn[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "_numFN":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.NumFn) <= i {
					r.NumFn = append(r.NumFn, Integer{})
				}
				r.NumFn[i].Id = v.Id
				r.NumFn[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "precision":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
			}
			for i := 0; d.More(); i++ {
				var v Decimal
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.Precision) <= i {
					r.Precision = append(r.Precision, Decimal{})
				}
				r.Precision[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "_precision":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.Precision) <= i {
					r.Precision = append(r.Precision, Decimal{})
				}
				r.Precision[i].Id = v.Id
				r.Precision[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "sensitivity":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
			}
			for i := 0; d.More(); i++ {
				var v Decimal
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.Sensitivity) <= i {
					r.Sensitivity = append(r.Sensitivity, Decimal{})
				}
				r.Sensitivity[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "_sensitivity":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.Sensitivity) <= i {
					r.Sensitivity = append(r.Sensitivity, Decimal{})
				}
				r.Sensitivity[i].Id = v.Id
				r.Sensitivity[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "fMeasure":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
			}
			for i := 0; d.More(); i++ {
				var v Decimal
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.FMeasure) <= i {
					r.FMeasure = append(r.FMeasure, Decimal{})
				}
				r.FMeasure[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		case "_fMeasure":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceQualityRoc element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.FMeasure) <= i {
					r.FMeasure = append(r.FMeasure, Decimal{})
				}
				r.FMeasure[i].Id = v.Id
				r.FMeasure[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceQualityRoc element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in MolecularSequenceQualityRoc", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in MolecularSequenceQualityRoc element", t)
	}
	return nil
}
func (r *MolecularSequenceRepository) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in MolecularSequenceRepository element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in MolecularSequenceRepository element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceRepository element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceRepository element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceRepository element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceRepository element", t)
			}
		case "type":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Type.Value = v.Value
		case "_type":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type.Id = v.Id
			r.Type.Extension = v.Extension
		case "url":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Url == nil {
				r.Url = &Uri{}
			}
			r.Url.Value = v.Value
		case "_url":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Url == nil {
				r.Url = &Uri{}
			}
			r.Url.Id = v.Id
			r.Url.Extension = v.Extension
		case "name":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Name == nil {
				r.Name = &String{}
			}
			r.Name.Value = v.Value
		case "_name":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Name == nil {
				r.Name = &String{}
			}
			r.Name.Id = v.Id
			r.Name.Extension = v.Extension
		case "datasetId":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.DatasetId == nil {
				r.DatasetId = &String{}
			}
			r.DatasetId.Value = v.Value
		case "_datasetId":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.DatasetId == nil {
				r.DatasetId = &String{}
			}
			r.DatasetId.Id = v.Id
			r.DatasetId.Extension = v.Extension
		case "variantsetId":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.VariantsetId == nil {
				r.VariantsetId = &String{}
			}
			r.VariantsetId.Value = v.Value
		case "_variantsetId":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.VariantsetId == nil {
				r.VariantsetId = &String{}
			}
			r.VariantsetId.Id = v.Id
			r.VariantsetId.Extension = v.Extension
		case "readsetId":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ReadsetId == nil {
				r.ReadsetId = &String{}
			}
			r.ReadsetId.Value = v.Value
		case "_readsetId":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ReadsetId == nil {
				r.ReadsetId = &String{}
			}
			r.ReadsetId.Id = v.Id
			r.ReadsetId.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in MolecularSequenceRepository", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in MolecularSequenceRepository element", t)
	}
	return nil
}
func (r *MolecularSequenceStructureVariant) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in MolecularSequenceStructureVariant element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in MolecularSequenceStructureVariant element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceStructureVariant element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceStructureVariant element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceStructureVariant element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceStructureVariant element", t)
			}
		case "variantType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.VariantType = &v
		case "exact":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Exact == nil {
				r.Exact = &Boolean{}
			}
			r.Exact.Value = v.Value
		case "_exact":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Exact == nil {
				r.Exact = &Boolean{}
			}
			r.Exact.Id = v.Id
			r.Exact.Extension = v.Extension
		case "length":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Length == nil {
				r.Length = &Integer{}
			}
			r.Length.Value = v.Value
		case "_length":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Length == nil {
				r.Length = &Integer{}
			}
			r.Length.Id = v.Id
			r.Length.Extension = v.Extension
		case "outer":
			var v MolecularSequenceStructureVariantOuter
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Outer = &v
		case "inner":
			var v MolecularSequenceStructureVariantInner
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Inner = &v
		default:
			return fmt.Errorf("invalid field: %s in MolecularSequenceStructureVariant", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in MolecularSequenceStructureVariant element", t)
	}
	return nil
}
func (r *MolecularSequenceStructureVariantOuter) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in MolecularSequenceStructureVariantOuter element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in MolecularSequenceStructureVariantOuter element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceStructureVariantOuter element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceStructureVariantOuter element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceStructureVariantOuter element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceStructureVariantOuter element", t)
			}
		case "start":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Start == nil {
				r.Start = &Integer{}
			}
			r.Start.Value = v.Value
		case "_start":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Start == nil {
				r.Start = &Integer{}
			}
			r.Start.Id = v.Id
			r.Start.Extension = v.Extension
		case "end":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.End == nil {
				r.End = &Integer{}
			}
			r.End.Value = v.Value
		case "_end":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.End == nil {
				r.End = &Integer{}
			}
			r.End.Id = v.Id
			r.End.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in MolecularSequenceStructureVariantOuter", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in MolecularSequenceStructureVariantOuter element", t)
	}
	return nil
}
func (r *MolecularSequenceStructureVariantInner) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in MolecularSequenceStructureVariantInner element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in MolecularSequenceStructureVariantInner element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceStructureVariantInner element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceStructureVariantInner element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceStructureVariantInner element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceStructureVariantInner element", t)
			}
		case "start":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Start == nil {
				r.Start = &Integer{}
			}
			r.Start.Value = v.Value
		case "_start":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Start == nil {
				r.Start = &Integer{}
			}
			r.Start.Id = v.Id
			r.Start.Extension = v.Extension
		case "end":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.End == nil {
				r.End = &Integer{}
			}
			r.End.Value = v.Value
		case "_end":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.End == nil {
				r.End = &Integer{}
			}
			r.End.Id = v.Id
			r.End.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in MolecularSequenceStructureVariantInner", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in MolecularSequenceStructureVariantInner element", t)
	}
	return nil
}
func (r MolecularSequence) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "MolecularSequence"
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CoordinateSystem, xml.StartElement{Name: xml.Name{Local: "coordinateSystem"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Patient, xml.StartElement{Name: xml.Name{Local: "patient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Specimen, xml.StartElement{Name: xml.Name{Local: "specimen"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Device, xml.StartElement{Name: xml.Name{Local: "device"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Performer, xml.StartElement{Name: xml.Name{Local: "performer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReferenceSeq, xml.StartElement{Name: xml.Name{Local: "referenceSeq"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Variant, xml.StartElement{Name: xml.Name{Local: "variant"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ObservedSeq, xml.StartElement{Name: xml.Name{Local: "observedSeq"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Quality, xml.StartElement{Name: xml.Name{Local: "quality"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReadCoverage, xml.StartElement{Name: xml.Name{Local: "readCoverage"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Repository, xml.StartElement{Name: xml.Name{Local: "repository"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Pointer, xml.StartElement{Name: xml.Name{Local: "pointer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.StructureVariant, xml.StartElement{Name: xml.Name{Local: "structureVariant"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r MolecularSequenceReferenceSeq) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Chromosome, xml.StartElement{Name: xml.Name{Local: "chromosome"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.GenomeBuild, xml.StartElement{Name: xml.Name{Local: "genomeBuild"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Orientation, xml.StartElement{Name: xml.Name{Local: "orientation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReferenceSeqId, xml.StartElement{Name: xml.Name{Local: "referenceSeqId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReferenceSeqPointer, xml.StartElement{Name: xml.Name{Local: "referenceSeqPointer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReferenceSeqString, xml.StartElement{Name: xml.Name{Local: "referenceSeqString"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Strand, xml.StartElement{Name: xml.Name{Local: "strand"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.WindowStart, xml.StartElement{Name: xml.Name{Local: "windowStart"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.WindowEnd, xml.StartElement{Name: xml.Name{Local: "windowEnd"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r MolecularSequenceVariant) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Start, xml.StartElement{Name: xml.Name{Local: "start"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.End, xml.StartElement{Name: xml.Name{Local: "end"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ObservedAllele, xml.StartElement{Name: xml.Name{Local: "observedAllele"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReferenceAllele, xml.StartElement{Name: xml.Name{Local: "referenceAllele"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Cigar, xml.StartElement{Name: xml.Name{Local: "cigar"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.VariantPointer, xml.StartElement{Name: xml.Name{Local: "variantPointer"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r MolecularSequenceQuality) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.StandardSequence, xml.StartElement{Name: xml.Name{Local: "standardSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Start, xml.StartElement{Name: xml.Name{Local: "start"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.End, xml.StartElement{Name: xml.Name{Local: "end"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Score, xml.StartElement{Name: xml.Name{Local: "score"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Method, xml.StartElement{Name: xml.Name{Local: "method"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TruthTp, xml.StartElement{Name: xml.Name{Local: "truthTP"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.QueryTp, xml.StartElement{Name: xml.Name{Local: "queryTP"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TruthFn, xml.StartElement{Name: xml.Name{Local: "truthFN"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.QueryFp, xml.StartElement{Name: xml.Name{Local: "queryFP"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.GtFp, xml.StartElement{Name: xml.Name{Local: "gtFP"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Precision, xml.StartElement{Name: xml.Name{Local: "precision"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Recall, xml.StartElement{Name: xml.Name{Local: "recall"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FScore, xml.StartElement{Name: xml.Name{Local: "fScore"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Roc, xml.StartElement{Name: xml.Name{Local: "roc"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r MolecularSequenceQualityRoc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Score, xml.StartElement{Name: xml.Name{Local: "score"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NumTp, xml.StartElement{Name: xml.Name{Local: "numTP"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NumFp, xml.StartElement{Name: xml.Name{Local: "numFP"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NumFn, xml.StartElement{Name: xml.Name{Local: "numFN"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Precision, xml.StartElement{Name: xml.Name{Local: "precision"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Sensitivity, xml.StartElement{Name: xml.Name{Local: "sensitivity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FMeasure, xml.StartElement{Name: xml.Name{Local: "fMeasure"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r MolecularSequenceRepository) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DatasetId, xml.StartElement{Name: xml.Name{Local: "datasetId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.VariantsetId, xml.StartElement{Name: xml.Name{Local: "variantsetId"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReadsetId, xml.StartElement{Name: xml.Name{Local: "readsetId"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r MolecularSequenceStructureVariant) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.VariantType, xml.StartElement{Name: xml.Name{Local: "variantType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Exact, xml.StartElement{Name: xml.Name{Local: "exact"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Length, xml.StartElement{Name: xml.Name{Local: "length"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Outer, xml.StartElement{Name: xml.Name{Local: "outer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Inner, xml.StartElement{Name: xml.Name{Local: "inner"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r MolecularSequenceStructureVariantOuter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Start, xml.StartElement{Name: xml.Name{Local: "start"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.End, xml.StartElement{Name: xml.Name{Local: "end"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r MolecularSequenceStructureVariantInner) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Start, xml.StartElement{Name: xml.Name{Local: "start"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.End, xml.StartElement{Name: xml.Name{Local: "end"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *MolecularSequence) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.Identifier = append(r.Identifier, v)
			case "type":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "coordinateSystem":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CoordinateSystem = v
			case "patient":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Patient = &v
			case "specimen":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Specimen = &v
			case "device":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Device = &v
			case "performer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Performer = &v
			case "quantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = &v
			case "referenceSeq":
				var v MolecularSequenceReferenceSeq
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReferenceSeq = &v
			case "variant":
				var v MolecularSequenceVariant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Variant = append(r.Variant, v)
			case "observedSeq":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ObservedSeq = &v
			case "quality":
				var v MolecularSequenceQuality
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quality = append(r.Quality, v)
			case "readCoverage":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReadCoverage = &v
			case "repository":
				var v MolecularSequenceRepository
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Repository = append(r.Repository, v)
			case "pointer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Pointer = append(r.Pointer, v)
			case "structureVariant":
				var v MolecularSequenceStructureVariant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StructureVariant = append(r.StructureVariant, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *MolecularSequenceReferenceSeq) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "chromosome":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Chromosome = &v
			case "genomeBuild":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.GenomeBuild = &v
			case "orientation":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Orientation = &v
			case "referenceSeqId":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReferenceSeqId = &v
			case "referenceSeqPointer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReferenceSeqPointer = &v
			case "referenceSeqString":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReferenceSeqString = &v
			case "strand":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Strand = &v
			case "windowStart":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.WindowStart = &v
			case "windowEnd":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.WindowEnd = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *MolecularSequenceVariant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "start":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Start = &v
			case "end":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.End = &v
			case "observedAllele":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ObservedAllele = &v
			case "referenceAllele":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReferenceAllele = &v
			case "cigar":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Cigar = &v
			case "variantPointer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.VariantPointer = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *MolecularSequenceQuality) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "standardSequence":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StandardSequence = &v
			case "start":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Start = &v
			case "end":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.End = &v
			case "score":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Score = &v
			case "method":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Method = &v
			case "truthTP":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TruthTp = &v
			case "queryTP":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.QueryTp = &v
			case "truthFN":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TruthFn = &v
			case "queryFP":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.QueryFp = &v
			case "gtFP":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.GtFp = &v
			case "precision":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Precision = &v
			case "recall":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Recall = &v
			case "fScore":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FScore = &v
			case "roc":
				var v MolecularSequenceQualityRoc
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Roc = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *MolecularSequenceQualityRoc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "score":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Score = append(r.Score, v)
			case "numTP":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NumTp = append(r.NumTp, v)
			case "numFP":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NumFp = append(r.NumFp, v)
			case "numFN":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NumFn = append(r.NumFn, v)
			case "precision":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Precision = append(r.Precision, v)
			case "sensitivity":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Sensitivity = append(r.Sensitivity, v)
			case "fMeasure":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FMeasure = append(r.FMeasure, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *MolecularSequenceRepository) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "url":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Url = &v
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "datasetId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DatasetId = &v
			case "variantsetId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.VariantsetId = &v
			case "readsetId":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReadsetId = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *MolecularSequenceStructureVariant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "variantType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.VariantType = &v
			case "exact":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Exact = &v
			case "length":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Length = &v
			case "outer":
				var v MolecularSequenceStructureVariantOuter
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Outer = &v
			case "inner":
				var v MolecularSequenceStructureVariantInner
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Inner = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *MolecularSequenceStructureVariantOuter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "start":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Start = &v
			case "end":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.End = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *MolecularSequenceStructureVariantInner) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "start":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Start = &v
			case "end":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.End = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r MolecularSequence) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "identifier") {
		for _, v := range r.Identifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "type") {
		if r.Type != nil {
			children = append(children, *r.Type)
		}
	}
	if len(name) == 0 || slices.Contains(name, "coordinateSystem") {
		children = append(children, r.CoordinateSystem)
	}
	if len(name) == 0 || slices.Contains(name, "patient") {
		if r.Patient != nil {
			children = append(children, *r.Patient)
		}
	}
	if len(name) == 0 || slices.Contains(name, "specimen") {
		if r.Specimen != nil {
			children = append(children, *r.Specimen)
		}
	}
	if len(name) == 0 || slices.Contains(name, "device") {
		if r.Device != nil {
			children = append(children, *r.Device)
		}
	}
	if len(name) == 0 || slices.Contains(name, "performer") {
		if r.Performer != nil {
			children = append(children, *r.Performer)
		}
	}
	if len(name) == 0 || slices.Contains(name, "quantity") {
		if r.Quantity != nil {
			children = append(children, *r.Quantity)
		}
	}
	if len(name) == 0 || slices.Contains(name, "referenceSeq") {
		if r.ReferenceSeq != nil {
			children = append(children, *r.ReferenceSeq)
		}
	}
	if len(name) == 0 || slices.Contains(name, "variant") {
		for _, v := range r.Variant {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "observedSeq") {
		if r.ObservedSeq != nil {
			children = append(children, *r.ObservedSeq)
		}
	}
	if len(name) == 0 || slices.Contains(name, "quality") {
		for _, v := range r.Quality {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "readCoverage") {
		if r.ReadCoverage != nil {
			children = append(children, *r.ReadCoverage)
		}
	}
	if len(name) == 0 || slices.Contains(name, "repository") {
		for _, v := range r.Repository {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "pointer") {
		for _, v := range r.Pointer {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "structureVariant") {
		for _, v := range r.StructureVariant {
			children = append(children, v)
		}
	}
	return children
}
func (r MolecularSequence) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert MolecularSequence to Boolean")
}
func (r MolecularSequence) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert MolecularSequence to String")
}
func (r MolecularSequence) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert MolecularSequence to Integer")
}
func (r MolecularSequence) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert MolecularSequence to Decimal")
}
func (r MolecularSequence) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert MolecularSequence to Date")
}
func (r MolecularSequence) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert MolecularSequence to Time")
}
func (r MolecularSequence) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert MolecularSequence to DateTime")
}
func (r MolecularSequence) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert MolecularSequence to Quantity")
}
func (r MolecularSequence) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequence
	switch other := other.(type) {
	case MolecularSequence:
		o = other
	case *MolecularSequence:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequence) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequence
	switch other := other.(type) {
	case MolecularSequence:
		o = other
	case *MolecularSequence:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequence) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Id",
				Namespace: "FHIR",
			},
		}, {
			Name: "Meta",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Meta",
				Namespace: "FHIR",
			},
		}, {
			Name: "ImplicitRules",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Uri",
				Namespace: "FHIR",
			},
		}, {
			Name: "Language",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Text",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Narrative",
				Namespace: "FHIR",
			},
		}, {
			Name: "Contained",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Identifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Identifier",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "CoordinateSystem",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "Patient",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Specimen",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Device",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Performer",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Quantity",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Quantity",
				Namespace: "FHIR",
			},
		}, {
			Name: "ReferenceSeq",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "MolecularSequenceReferenceSeq",
				Namespace: "FHIR",
			},
		}, {
			Name: "Variant",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "MolecularSequenceVariant",
				Namespace: "FHIR",
			},
		}, {
			Name: "ObservedSeq",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Quality",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "MolecularSequenceQuality",
				Namespace: "FHIR",
			},
		}, {
			Name: "ReadCoverage",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "Repository",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "MolecularSequenceRepository",
				Namespace: "FHIR",
			},
		}, {
			Name: "Pointer",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "StructureVariant",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "MolecularSequenceStructureVariant",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "MolecularSequence",
			Namespace: "FHIR",
		},
	}
}
func (r MolecularSequenceReferenceSeq) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "chromosome") {
		if r.Chromosome != nil {
			children = append(children, *r.Chromosome)
		}
	}
	if len(name) == 0 || slices.Contains(name, "genomeBuild") {
		if r.GenomeBuild != nil {
			children = append(children, *r.GenomeBuild)
		}
	}
	if len(name) == 0 || slices.Contains(name, "orientation") {
		if r.Orientation != nil {
			children = append(children, *r.Orientation)
		}
	}
	if len(name) == 0 || slices.Contains(name, "referenceSeqId") {
		if r.ReferenceSeqId != nil {
			children = append(children, *r.ReferenceSeqId)
		}
	}
	if len(name) == 0 || slices.Contains(name, "referenceSeqPointer") {
		if r.ReferenceSeqPointer != nil {
			children = append(children, *r.ReferenceSeqPointer)
		}
	}
	if len(name) == 0 || slices.Contains(name, "referenceSeqString") {
		if r.ReferenceSeqString != nil {
			children = append(children, *r.ReferenceSeqString)
		}
	}
	if len(name) == 0 || slices.Contains(name, "strand") {
		if r.Strand != nil {
			children = append(children, *r.Strand)
		}
	}
	if len(name) == 0 || slices.Contains(name, "windowStart") {
		if r.WindowStart != nil {
			children = append(children, *r.WindowStart)
		}
	}
	if len(name) == 0 || slices.Contains(name, "windowEnd") {
		if r.WindowEnd != nil {
			children = append(children, *r.WindowEnd)
		}
	}
	return children
}
func (r MolecularSequenceReferenceSeq) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert MolecularSequenceReferenceSeq to Boolean")
}
func (r MolecularSequenceReferenceSeq) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert MolecularSequenceReferenceSeq to String")
}
func (r MolecularSequenceReferenceSeq) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert MolecularSequenceReferenceSeq to Integer")
}
func (r MolecularSequenceReferenceSeq) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert MolecularSequenceReferenceSeq to Decimal")
}
func (r MolecularSequenceReferenceSeq) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert MolecularSequenceReferenceSeq to Date")
}
func (r MolecularSequenceReferenceSeq) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert MolecularSequenceReferenceSeq to Time")
}
func (r MolecularSequenceReferenceSeq) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert MolecularSequenceReferenceSeq to DateTime")
}
func (r MolecularSequenceReferenceSeq) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert MolecularSequenceReferenceSeq to Quantity")
}
func (r MolecularSequenceReferenceSeq) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceReferenceSeq
	switch other := other.(type) {
	case MolecularSequenceReferenceSeq:
		o = other
	case *MolecularSequenceReferenceSeq:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceReferenceSeq) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceReferenceSeq
	switch other := other.(type) {
	case MolecularSequenceReferenceSeq:
		o = other
	case *MolecularSequenceReferenceSeq:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceReferenceSeq) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Chromosome",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "GenomeBuild",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Orientation",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "ReferenceSeqId",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ReferenceSeqPointer",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "ReferenceSeqString",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Strand",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "WindowStart",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "WindowEnd",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "MolecularSequenceReferenceSeq",
			Namespace: "FHIR",
		},
	}
}
func (r MolecularSequenceVariant) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "start") {
		if r.Start != nil {
			children = append(children, *r.Start)
		}
	}
	if len(name) == 0 || slices.Contains(name, "end") {
		if r.End != nil {
			children = append(children, *r.End)
		}
	}
	if len(name) == 0 || slices.Contains(name, "observedAllele") {
		if r.ObservedAllele != nil {
			children = append(children, *r.ObservedAllele)
		}
	}
	if len(name) == 0 || slices.Contains(name, "referenceAllele") {
		if r.ReferenceAllele != nil {
			children = append(children, *r.ReferenceAllele)
		}
	}
	if len(name) == 0 || slices.Contains(name, "cigar") {
		if r.Cigar != nil {
			children = append(children, *r.Cigar)
		}
	}
	if len(name) == 0 || slices.Contains(name, "variantPointer") {
		if r.VariantPointer != nil {
			children = append(children, *r.VariantPointer)
		}
	}
	return children
}
func (r MolecularSequenceVariant) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert MolecularSequenceVariant to Boolean")
}
func (r MolecularSequenceVariant) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert MolecularSequenceVariant to String")
}
func (r MolecularSequenceVariant) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert MolecularSequenceVariant to Integer")
}
func (r MolecularSequenceVariant) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert MolecularSequenceVariant to Decimal")
}
func (r MolecularSequenceVariant) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert MolecularSequenceVariant to Date")
}
func (r MolecularSequenceVariant) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert MolecularSequenceVariant to Time")
}
func (r MolecularSequenceVariant) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert MolecularSequenceVariant to DateTime")
}
func (r MolecularSequenceVariant) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert MolecularSequenceVariant to Quantity")
}
func (r MolecularSequenceVariant) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceVariant
	switch other := other.(type) {
	case MolecularSequenceVariant:
		o = other
	case *MolecularSequenceVariant:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceVariant) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceVariant
	switch other := other.(type) {
	case MolecularSequenceVariant:
		o = other
	case *MolecularSequenceVariant:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceVariant) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Start",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "End",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "ObservedAllele",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "ReferenceAllele",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Cigar",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "VariantPointer",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "MolecularSequenceVariant",
			Namespace: "FHIR",
		},
	}
}
func (r MolecularSequenceQuality) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "type") {
		children = append(children, r.Type)
	}
	if len(name) == 0 || slices.Contains(name, "standardSequence") {
		if r.StandardSequence != nil {
			children = append(children, *r.StandardSequence)
		}
	}
	if len(name) == 0 || slices.Contains(name, "start") {
		if r.Start != nil {
			children = append(children, *r.Start)
		}
	}
	if len(name) == 0 || slices.Contains(name, "end") {
		if r.End != nil {
			children = append(children, *r.End)
		}
	}
	if len(name) == 0 || slices.Contains(name, "score") {
		if r.Score != nil {
			children = append(children, *r.Score)
		}
	}
	if len(name) == 0 || slices.Contains(name, "method") {
		if r.Method != nil {
			children = append(children, *r.Method)
		}
	}
	if len(name) == 0 || slices.Contains(name, "truthTP") {
		if r.TruthTp != nil {
			children = append(children, *r.TruthTp)
		}
	}
	if len(name) == 0 || slices.Contains(name, "queryTP") {
		if r.QueryTp != nil {
			children = append(children, *r.QueryTp)
		}
	}
	if len(name) == 0 || slices.Contains(name, "truthFN") {
		if r.TruthFn != nil {
			children = append(children, *r.TruthFn)
		}
	}
	if len(name) == 0 || slices.Contains(name, "queryFP") {
		if r.QueryFp != nil {
			children = append(children, *r.QueryFp)
		}
	}
	if len(name) == 0 || slices.Contains(name, "gtFP") {
		if r.GtFp != nil {
			children = append(children, *r.GtFp)
		}
	}
	if len(name) == 0 || slices.Contains(name, "precision") {
		if r.Precision != nil {
			children = append(children, *r.Precision)
		}
	}
	if len(name) == 0 || slices.Contains(name, "recall") {
		if r.Recall != nil {
			children = append(children, *r.Recall)
		}
	}
	if len(name) == 0 || slices.Contains(name, "fScore") {
		if r.FScore != nil {
			children = append(children, *r.FScore)
		}
	}
	if len(name) == 0 || slices.Contains(name, "roc") {
		if r.Roc != nil {
			children = append(children, *r.Roc)
		}
	}
	return children
}
func (r MolecularSequenceQuality) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert MolecularSequenceQuality to Boolean")
}
func (r MolecularSequenceQuality) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert MolecularSequenceQuality to String")
}
func (r MolecularSequenceQuality) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert MolecularSequenceQuality to Integer")
}
func (r MolecularSequenceQuality) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert MolecularSequenceQuality to Decimal")
}
func (r MolecularSequenceQuality) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert MolecularSequenceQuality to Date")
}
func (r MolecularSequenceQuality) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert MolecularSequenceQuality to Time")
}
func (r MolecularSequenceQuality) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert MolecularSequenceQuality to DateTime")
}
func (r MolecularSequenceQuality) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert MolecularSequenceQuality to Quantity")
}
func (r MolecularSequenceQuality) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceQuality
	switch other := other.(type) {
	case MolecularSequenceQuality:
		o = other
	case *MolecularSequenceQuality:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceQuality) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceQuality
	switch other := other.(type) {
	case MolecularSequenceQuality:
		o = other
	case *MolecularSequenceQuality:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceQuality) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "StandardSequence",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Start",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "End",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "Score",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Quantity",
				Namespace: "FHIR",
			},
		}, {
			Name: "Method",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "TruthTp",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "QueryTp",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "TruthFn",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "QueryFp",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "GtFp",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "Precision",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "Recall",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "FScore",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "Roc",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "MolecularSequenceQualityRoc",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "MolecularSequenceQuality",
			Namespace: "FHIR",
		},
	}
}
func (r MolecularSequenceQualityRoc) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "score") {
		for _, v := range r.Score {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "numTP") {
		for _, v := range r.NumTp {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "numFP") {
		for _, v := range r.NumFp {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "numFN") {
		for _, v := range r.NumFn {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "precision") {
		for _, v := range r.Precision {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "sensitivity") {
		for _, v := range r.Sensitivity {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "fMeasure") {
		for _, v := range r.FMeasure {
			children = append(children, v)
		}
	}
	return children
}
func (r MolecularSequenceQualityRoc) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert MolecularSequenceQualityRoc to Boolean")
}
func (r MolecularSequenceQualityRoc) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert MolecularSequenceQualityRoc to String")
}
func (r MolecularSequenceQualityRoc) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert MolecularSequenceQualityRoc to Integer")
}
func (r MolecularSequenceQualityRoc) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert MolecularSequenceQualityRoc to Decimal")
}
func (r MolecularSequenceQualityRoc) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert MolecularSequenceQualityRoc to Date")
}
func (r MolecularSequenceQualityRoc) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert MolecularSequenceQualityRoc to Time")
}
func (r MolecularSequenceQualityRoc) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert MolecularSequenceQualityRoc to DateTime")
}
func (r MolecularSequenceQualityRoc) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert MolecularSequenceQualityRoc to Quantity")
}
func (r MolecularSequenceQualityRoc) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceQualityRoc
	switch other := other.(type) {
	case MolecularSequenceQualityRoc:
		o = other
	case *MolecularSequenceQualityRoc:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceQualityRoc) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceQualityRoc
	switch other := other.(type) {
	case MolecularSequenceQualityRoc:
		o = other
	case *MolecularSequenceQualityRoc:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceQualityRoc) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Score",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "NumTp",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "NumFp",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "NumFn",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "Precision",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "Sensitivity",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "FMeasure",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "MolecularSequenceQualityRoc",
			Namespace: "FHIR",
		},
	}
}
func (r MolecularSequenceRepository) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "type") {
		children = append(children, r.Type)
	}
	if len(name) == 0 || slices.Contains(name, "url") {
		if r.Url != nil {
			children = append(children, *r.Url)
		}
	}
	if len(name) == 0 || slices.Contains(name, "name") {
		if r.Name != nil {
			children = append(children, *r.Name)
		}
	}
	if len(name) == 0 || slices.Contains(name, "datasetId") {
		if r.DatasetId != nil {
			children = append(children, *r.DatasetId)
		}
	}
	if len(name) == 0 || slices.Contains(name, "variantsetId") {
		if r.VariantsetId != nil {
			children = append(children, *r.VariantsetId)
		}
	}
	if len(name) == 0 || slices.Contains(name, "readsetId") {
		if r.ReadsetId != nil {
			children = append(children, *r.ReadsetId)
		}
	}
	return children
}
func (r MolecularSequenceRepository) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert MolecularSequenceRepository to Boolean")
}
func (r MolecularSequenceRepository) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert MolecularSequenceRepository to String")
}
func (r MolecularSequenceRepository) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert MolecularSequenceRepository to Integer")
}
func (r MolecularSequenceRepository) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert MolecularSequenceRepository to Decimal")
}
func (r MolecularSequenceRepository) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert MolecularSequenceRepository to Date")
}
func (r MolecularSequenceRepository) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert MolecularSequenceRepository to Time")
}
func (r MolecularSequenceRepository) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert MolecularSequenceRepository to DateTime")
}
func (r MolecularSequenceRepository) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert MolecularSequenceRepository to Quantity")
}
func (r MolecularSequenceRepository) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceRepository
	switch other := other.(type) {
	case MolecularSequenceRepository:
		o = other
	case *MolecularSequenceRepository:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceRepository) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceRepository
	switch other := other.(type) {
	case MolecularSequenceRepository:
		o = other
	case *MolecularSequenceRepository:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceRepository) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Type",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Url",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Uri",
				Namespace: "FHIR",
			},
		}, {
			Name: "Name",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "DatasetId",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "VariantsetId",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "ReadsetId",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "MolecularSequenceRepository",
			Namespace: "FHIR",
		},
	}
}
func (r MolecularSequenceStructureVariant) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "variantType") {
		if r.VariantType != nil {
			children = append(children, *r.VariantType)
		}
	}
	if len(name) == 0 || slices.Contains(name, "exact") {
		if r.Exact != nil {
			children = append(children, *r.Exact)
		}
	}
	if len(name) == 0 || slices.Contains(name, "length") {
		if r.Length != nil {
			children = append(children, *r.Length)
		}
	}
	if len(name) == 0 || slices.Contains(name, "outer") {
		if r.Outer != nil {
			children = append(children, *r.Outer)
		}
	}
	if len(name) == 0 || slices.Contains(name, "inner") {
		if r.Inner != nil {
			children = append(children, *r.Inner)
		}
	}
	return children
}
func (r MolecularSequenceStructureVariant) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariant to Boolean")
}
func (r MolecularSequenceStructureVariant) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariant to String")
}
func (r MolecularSequenceStructureVariant) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariant to Integer")
}
func (r MolecularSequenceStructureVariant) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariant to Decimal")
}
func (r MolecularSequenceStructureVariant) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariant to Date")
}
func (r MolecularSequenceStructureVariant) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariant to Time")
}
func (r MolecularSequenceStructureVariant) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariant to DateTime")
}
func (r MolecularSequenceStructureVariant) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariant to Quantity")
}
func (r MolecularSequenceStructureVariant) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceStructureVariant
	switch other := other.(type) {
	case MolecularSequenceStructureVariant:
		o = other
	case *MolecularSequenceStructureVariant:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceStructureVariant) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceStructureVariant
	switch other := other.(type) {
	case MolecularSequenceStructureVariant:
		o = other
	case *MolecularSequenceStructureVariant:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceStructureVariant) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "VariantType",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Exact",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Length",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "Outer",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "MolecularSequenceStructureVariantOuter",
				Namespace: "FHIR",
			},
		}, {
			Name: "Inner",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "MolecularSequenceStructureVariantInner",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "MolecularSequenceStructureVariant",
			Namespace: "FHIR",
		},
	}
}
func (r MolecularSequenceStructureVariantOuter) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "start") {
		if r.Start != nil {
			children = append(children, *r.Start)
		}
	}
	if len(name) == 0 || slices.Contains(name, "end") {
		if r.End != nil {
			children = append(children, *r.End)
		}
	}
	return children
}
func (r MolecularSequenceStructureVariantOuter) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantOuter to Boolean")
}
func (r MolecularSequenceStructureVariantOuter) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantOuter to String")
}
func (r MolecularSequenceStructureVariantOuter) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantOuter to Integer")
}
func (r MolecularSequenceStructureVariantOuter) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantOuter to Decimal")
}
func (r MolecularSequenceStructureVariantOuter) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantOuter to Date")
}
func (r MolecularSequenceStructureVariantOuter) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantOuter to Time")
}
func (r MolecularSequenceStructureVariantOuter) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantOuter to DateTime")
}
func (r MolecularSequenceStructureVariantOuter) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantOuter to Quantity")
}
func (r MolecularSequenceStructureVariantOuter) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceStructureVariantOuter
	switch other := other.(type) {
	case MolecularSequenceStructureVariantOuter:
		o = other
	case *MolecularSequenceStructureVariantOuter:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceStructureVariantOuter) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceStructureVariantOuter
	switch other := other.(type) {
	case MolecularSequenceStructureVariantOuter:
		o = other
	case *MolecularSequenceStructureVariantOuter:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceStructureVariantOuter) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Start",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "End",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "MolecularSequenceStructureVariantOuter",
			Namespace: "FHIR",
		},
	}
}
func (r MolecularSequenceStructureVariantInner) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "start") {
		if r.Start != nil {
			children = append(children, *r.Start)
		}
	}
	if len(name) == 0 || slices.Contains(name, "end") {
		if r.End != nil {
			children = append(children, *r.End)
		}
	}
	return children
}
func (r MolecularSequenceStructureVariantInner) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantInner to Boolean")
}
func (r MolecularSequenceStructureVariantInner) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantInner to String")
}
func (r MolecularSequenceStructureVariantInner) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantInner to Integer")
}
func (r MolecularSequenceStructureVariantInner) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantInner to Decimal")
}
func (r MolecularSequenceStructureVariantInner) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantInner to Date")
}
func (r MolecularSequenceStructureVariantInner) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantInner to Time")
}
func (r MolecularSequenceStructureVariantInner) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantInner to DateTime")
}
func (r MolecularSequenceStructureVariantInner) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert MolecularSequenceStructureVariantInner to Quantity")
}
func (r MolecularSequenceStructureVariantInner) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceStructureVariantInner
	switch other := other.(type) {
	case MolecularSequenceStructureVariantInner:
		o = other
	case *MolecularSequenceStructureVariantInner:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceStructureVariantInner) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o MolecularSequenceStructureVariantInner
	switch other := other.(type) {
	case MolecularSequenceStructureVariantInner:
		o = other
	case *MolecularSequenceStructureVariantInner:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r MolecularSequenceStructureVariantInner) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Start",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "End",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "MolecularSequenceStructureVariantInner",
			Namespace: "FHIR",
		},
	}
}
