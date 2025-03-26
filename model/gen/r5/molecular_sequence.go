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

// Representation of a molecular sequence.
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
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, nor can they have their own independent transaction scope. This is allowed to be a Parameters resource if and only if it is referenced by a resource that provides context/meaning.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A unique identifier for this particular sequence instance.
	Identifier []Identifier
	// Amino Acid Sequence/ DNA Sequence / RNA Sequence.
	Type *Code
	// Indicates the subject this sequence is associated too.
	Subject *Reference
	// The actual focus of a molecular sequence when it is not the patient of record representing something or someone associated with the patient such as a spouse, parent, child, or sibling. For example, in trio testing, the subject would be the child (proband) and the focus would be the parent.
	Focus []Reference
	// Specimen used for sequencing.
	Specimen *Reference
	// The method for sequencing, for example, chip information.
	Device *Reference
	// The organization or lab that should be responsible for this result.
	Performer *Reference
	// Sequence that was observed.
	Literal *String
	// Sequence that was observed as file content. Can be an actual file contents, or referenced by a URL to an external system.
	Formatted []Attachment
	// A sequence defined relative to another sequence.
	Relative []MolecularSequenceRelative
}

// A sequence defined relative to another sequence.
type MolecularSequenceRelative struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// These are different ways of identifying nucleotides or amino acids within a sequence. Different databases and file types may use different systems. For detail definitions, see https://loinc.org/92822-6/ for more detail.
	CoordinateSystem CodeableConcept
	// Indicates the order in which the sequence should be considered when putting multiple 'relative' elements together.
	OrdinalPosition *Integer
	// Indicates the nucleotide range in the composed sequence when multiple 'relative' elements are used together.
	SequenceRange *Range
	// A sequence that is used as a starting sequence to describe variants that are present in a sequence analyzed.
	StartingSequence *MolecularSequenceRelativeStartingSequence
	// Changes in sequence from the starting sequence.
	Edit []MolecularSequenceRelativeEdit
}

// A sequence that is used as a starting sequence to describe variants that are present in a sequence analyzed.
type MolecularSequenceRelativeStartingSequence struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The genome assembly used for starting sequence, e.g. GRCh38.
	GenomeAssembly *CodeableConcept
	// Structural unit composed of a nucleic acid molecule which controls its own replication through the interaction of specific proteins at one or more origins of replication ([SO:0000340](http://www.sequenceontology.org/browser/current_svn/term/SO:0000340)).
	Chromosome *CodeableConcept
	// The reference sequence that represents the starting sequence.
	Sequence isMolecularSequenceRelativeStartingSequenceSequence
	// Start position of the window on the starting sequence. This value should honor the rules of the coordinateSystem.
	WindowStart *Integer
	// End position of the window on the starting sequence. This value should honor the rules of the  coordinateSystem.
	WindowEnd *Integer
	// A relative reference to a DNA strand based on gene orientation. The strand that contains the open reading frame of the gene is the "sense" strand, and the opposite complementary strand is the "antisense" strand.
	Orientation *Code
	// An absolute reference to a strand. The Watson strand is the strand whose 5'-end is on the short arm of the chromosome, and the Crick strand as the one whose 5'-end is on the long arm.
	Strand *Code
}
type isMolecularSequenceRelativeStartingSequenceSequence interface {
	model.Element
	isMolecularSequenceRelativeStartingSequenceSequence()
}

func (r CodeableConcept) isMolecularSequenceRelativeStartingSequenceSequence() {}
func (r String) isMolecularSequenceRelativeStartingSequenceSequence()          {}
func (r Reference) isMolecularSequenceRelativeStartingSequenceSequence()       {}

// Changes in sequence from the starting sequence.
type MolecularSequenceRelativeEdit struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Start position of the edit on the starting sequence. If the coordinate system is either 0-based or 1-based, then start position is inclusive.
	Start *Integer
	// End position of the edit on the starting sequence. If the coordinate system is 0-based then end is exclusive and does not include the last position. If the coordinate system is 1-base, then end is inclusive and includes the last position.
	End *Integer
	// Allele that was observed. Nucleotide(s)/amino acids from start position of sequence to stop position of sequence on the positive (+) strand of the observed sequence. When the sequence type is DNA, it should be the sequence on the positive (+) strand. This will lay in the range between variant.start and variant.end.
	ReplacementSequence *String
	// Allele in the starting sequence. Nucleotide(s)/amino acids from start position of sequence to stop position of sequence on the positive (+) strand of the starting sequence. When the sequence  type is DNA, it should be the sequence on the positive (+) strand. This will lay in the range between variant.start and variant.end.
	ReplacedSequence *String
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
	if r.Subject != nil {
		s += r.Subject.MemSize()
	}
	for _, i := range r.Focus {
		s += i.MemSize()
	}
	s += (cap(r.Focus) - len(r.Focus)) * int(unsafe.Sizeof(Reference{}))
	if r.Specimen != nil {
		s += r.Specimen.MemSize()
	}
	if r.Device != nil {
		s += r.Device.MemSize()
	}
	if r.Performer != nil {
		s += r.Performer.MemSize()
	}
	if r.Literal != nil {
		s += r.Literal.MemSize()
	}
	for _, i := range r.Formatted {
		s += i.MemSize()
	}
	s += (cap(r.Formatted) - len(r.Formatted)) * int(unsafe.Sizeof(Attachment{}))
	for _, i := range r.Relative {
		s += i.MemSize()
	}
	s += (cap(r.Relative) - len(r.Relative)) * int(unsafe.Sizeof(MolecularSequenceRelative{}))
	return s
}
func (r MolecularSequenceRelative) MemSize() int {
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
	s += r.CoordinateSystem.MemSize() - int(unsafe.Sizeof(r.CoordinateSystem))
	if r.OrdinalPosition != nil {
		s += r.OrdinalPosition.MemSize()
	}
	if r.SequenceRange != nil {
		s += r.SequenceRange.MemSize()
	}
	if r.StartingSequence != nil {
		s += r.StartingSequence.MemSize()
	}
	for _, i := range r.Edit {
		s += i.MemSize()
	}
	s += (cap(r.Edit) - len(r.Edit)) * int(unsafe.Sizeof(MolecularSequenceRelativeEdit{}))
	return s
}
func (r MolecularSequenceRelativeStartingSequence) MemSize() int {
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
	if r.GenomeAssembly != nil {
		s += r.GenomeAssembly.MemSize()
	}
	if r.Chromosome != nil {
		s += r.Chromosome.MemSize()
	}
	if r.Sequence != nil {
		s += r.Sequence.MemSize()
	}
	if r.WindowStart != nil {
		s += r.WindowStart.MemSize()
	}
	if r.WindowEnd != nil {
		s += r.WindowEnd.MemSize()
	}
	if r.Orientation != nil {
		s += r.Orientation.MemSize()
	}
	if r.Strand != nil {
		s += r.Strand.MemSize()
	}
	return s
}
func (r MolecularSequenceRelativeEdit) MemSize() int {
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
	if r.ReplacementSequence != nil {
		s += r.ReplacementSequence.MemSize()
	}
	if r.ReplacedSequence != nil {
		s += r.ReplacedSequence.MemSize()
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
func (r MolecularSequenceRelative) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r MolecularSequenceRelativeStartingSequence) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r MolecularSequenceRelativeEdit) String() string {
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
	if r.Subject != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"subject\":"))
		if err != nil {
			return err
		}
		err = r.Subject.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Focus) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"focus\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Focus {
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
	if r.Literal != nil && r.Literal.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"literal\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Literal)
		if err != nil {
			return err
		}
	}
	if r.Literal != nil && (r.Literal.Id != nil || r.Literal.Extension != nil) {
		p := primitiveElement{Id: r.Literal.Id, Extension: r.Literal.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_literal\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Formatted) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"formatted\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Formatted {
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
	if len(r.Relative) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"relative\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Relative {
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
func (r MolecularSequenceRelative) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r MolecularSequenceRelative) marshalJSON(w io.Writer) error {
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
	err = r.CoordinateSystem.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.OrdinalPosition != nil && r.OrdinalPosition.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"ordinalPosition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.OrdinalPosition)
		if err != nil {
			return err
		}
	}
	if r.OrdinalPosition != nil && (r.OrdinalPosition.Id != nil || r.OrdinalPosition.Extension != nil) {
		p := primitiveElement{Id: r.OrdinalPosition.Id, Extension: r.OrdinalPosition.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_ordinalPosition\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.SequenceRange != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sequenceRange\":"))
		if err != nil {
			return err
		}
		err = r.SequenceRange.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.StartingSequence != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"startingSequence\":"))
		if err != nil {
			return err
		}
		err = r.StartingSequence.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Edit) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"edit\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Edit {
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
func (r MolecularSequenceRelativeStartingSequence) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r MolecularSequenceRelativeStartingSequence) marshalJSON(w io.Writer) error {
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
	if r.GenomeAssembly != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"genomeAssembly\":"))
		if err != nil {
			return err
		}
		err = r.GenomeAssembly.marshalJSON(w)
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
	switch v := r.Sequence.(type) {
	case CodeableConcept:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sequenceCodeableConcept\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *CodeableConcept:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sequenceCodeableConcept\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"sequenceString\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_sequenceString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *String:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"sequenceString\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(v)
			if err != nil {
				return err
			}
		}
		if v.Id != nil || v.Extension != nil {
			p := primitiveElement{Id: v.Id, Extension: v.Extension}
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_sequenceString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Reference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sequenceReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Reference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sequenceReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
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
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r MolecularSequenceRelativeEdit) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r MolecularSequenceRelativeEdit) marshalJSON(w io.Writer) error {
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
	if r.ReplacementSequence != nil && r.ReplacementSequence.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"replacementSequence\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ReplacementSequence)
		if err != nil {
			return err
		}
	}
	if r.ReplacementSequence != nil && (r.ReplacementSequence.Id != nil || r.ReplacementSequence.Extension != nil) {
		p := primitiveElement{Id: r.ReplacementSequence.Id, Extension: r.ReplacementSequence.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_replacementSequence\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ReplacedSequence != nil && r.ReplacedSequence.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"replacedSequence\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ReplacedSequence)
		if err != nil {
			return err
		}
	}
	if r.ReplacedSequence != nil && (r.ReplacedSequence.Id != nil || r.ReplacedSequence.Extension != nil) {
		p := primitiveElement{Id: r.ReplacedSequence.Id, Extension: r.ReplacedSequence.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_replacedSequence\":"))
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
		case "subject":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Subject = &v
		case "focus":
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
				r.Focus = append(r.Focus, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequence element", t)
			}
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
		case "literal":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Literal == nil {
				r.Literal = &String{}
			}
			r.Literal.Value = v.Value
		case "_literal":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Literal == nil {
				r.Literal = &String{}
			}
			r.Literal.Id = v.Id
			r.Literal.Extension = v.Extension
		case "formatted":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequence element", t)
			}
			for d.More() {
				var v Attachment
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Formatted = append(r.Formatted, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequence element", t)
			}
		case "relative":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequence element", t)
			}
			for d.More() {
				var v MolecularSequenceRelative
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Relative = append(r.Relative, v)
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
func (r *MolecularSequenceRelative) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in MolecularSequenceRelative element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in MolecularSequenceRelative element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceRelative element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceRelative element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceRelative element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceRelative element", t)
			}
		case "coordinateSystem":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.CoordinateSystem = v
		case "ordinalPosition":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.OrdinalPosition == nil {
				r.OrdinalPosition = &Integer{}
			}
			r.OrdinalPosition.Value = v.Value
		case "_ordinalPosition":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.OrdinalPosition == nil {
				r.OrdinalPosition = &Integer{}
			}
			r.OrdinalPosition.Id = v.Id
			r.OrdinalPosition.Extension = v.Extension
		case "sequenceRange":
			var v Range
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.SequenceRange = &v
		case "startingSequence":
			var v MolecularSequenceRelativeStartingSequence
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.StartingSequence = &v
		case "edit":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceRelative element", t)
			}
			for d.More() {
				var v MolecularSequenceRelativeEdit
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Edit = append(r.Edit, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceRelative element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in MolecularSequenceRelative", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in MolecularSequenceRelative element", t)
	}
	return nil
}
func (r *MolecularSequenceRelativeStartingSequence) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in MolecularSequenceRelativeStartingSequence element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in MolecularSequenceRelativeStartingSequence element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceRelativeStartingSequence element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceRelativeStartingSequence element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceRelativeStartingSequence element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceRelativeStartingSequence element", t)
			}
		case "genomeAssembly":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.GenomeAssembly = &v
		case "chromosome":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Chromosome = &v
		case "sequenceCodeableConcept":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Sequence = v
		case "sequenceString":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Sequence != nil {
				r.Sequence = String{
					Extension: r.Sequence.(String).Extension,
					Id:        r.Sequence.(String).Id,
					Value:     v.Value,
				}
			} else {
				r.Sequence = v
			}
		case "_sequenceString":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Sequence != nil {
				r.Sequence = String{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Sequence.(String).Value,
				}
			} else {
				r.Sequence = String{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "sequenceReference":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Sequence = v
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
		default:
			return fmt.Errorf("invalid field: %s in MolecularSequenceRelativeStartingSequence", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in MolecularSequenceRelativeStartingSequence element", t)
	}
	return nil
}
func (r *MolecularSequenceRelativeEdit) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in MolecularSequenceRelativeEdit element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in MolecularSequenceRelativeEdit element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceRelativeEdit element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceRelativeEdit element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in MolecularSequenceRelativeEdit element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in MolecularSequenceRelativeEdit element", t)
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
		case "replacementSequence":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ReplacementSequence == nil {
				r.ReplacementSequence = &String{}
			}
			r.ReplacementSequence.Value = v.Value
		case "_replacementSequence":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ReplacementSequence == nil {
				r.ReplacementSequence = &String{}
			}
			r.ReplacementSequence.Id = v.Id
			r.ReplacementSequence.Extension = v.Extension
		case "replacedSequence":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ReplacedSequence == nil {
				r.ReplacedSequence = &String{}
			}
			r.ReplacedSequence.Value = v.Value
		case "_replacedSequence":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ReplacedSequence == nil {
				r.ReplacedSequence = &String{}
			}
			r.ReplacedSequence.Id = v.Id
			r.ReplacedSequence.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in MolecularSequenceRelativeEdit", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in MolecularSequenceRelativeEdit element", t)
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
	err = e.EncodeElement(r.Subject, xml.StartElement{Name: xml.Name{Local: "subject"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Focus, xml.StartElement{Name: xml.Name{Local: "focus"}})
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
	err = e.EncodeElement(r.Literal, xml.StartElement{Name: xml.Name{Local: "literal"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Formatted, xml.StartElement{Name: xml.Name{Local: "formatted"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Relative, xml.StartElement{Name: xml.Name{Local: "relative"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r MolecularSequenceRelative) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.CoordinateSystem, xml.StartElement{Name: xml.Name{Local: "coordinateSystem"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OrdinalPosition, xml.StartElement{Name: xml.Name{Local: "ordinalPosition"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SequenceRange, xml.StartElement{Name: xml.Name{Local: "sequenceRange"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.StartingSequence, xml.StartElement{Name: xml.Name{Local: "startingSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Edit, xml.StartElement{Name: xml.Name{Local: "edit"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r MolecularSequenceRelativeStartingSequence) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.GenomeAssembly, xml.StartElement{Name: xml.Name{Local: "genomeAssembly"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Chromosome, xml.StartElement{Name: xml.Name{Local: "chromosome"}})
	if err != nil {
		return err
	}
	switch v := r.Sequence.(type) {
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "sequenceCodeableConcept"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "sequenceString"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "sequenceReference"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.WindowStart, xml.StartElement{Name: xml.Name{Local: "windowStart"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.WindowEnd, xml.StartElement{Name: xml.Name{Local: "windowEnd"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Orientation, xml.StartElement{Name: xml.Name{Local: "orientation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Strand, xml.StartElement{Name: xml.Name{Local: "strand"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r MolecularSequenceRelativeEdit) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ReplacementSequence, xml.StartElement{Name: xml.Name{Local: "replacementSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReplacedSequence, xml.StartElement{Name: xml.Name{Local: "replacedSequence"}})
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
			case "subject":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = &v
			case "focus":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Focus = append(r.Focus, v)
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
			case "literal":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Literal = &v
			case "formatted":
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Formatted = append(r.Formatted, v)
			case "relative":
				var v MolecularSequenceRelative
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Relative = append(r.Relative, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *MolecularSequenceRelative) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "coordinateSystem":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CoordinateSystem = v
			case "ordinalPosition":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OrdinalPosition = &v
			case "sequenceRange":
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SequenceRange = &v
			case "startingSequence":
				var v MolecularSequenceRelativeStartingSequence
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StartingSequence = &v
			case "edit":
				var v MolecularSequenceRelativeEdit
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Edit = append(r.Edit, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *MolecularSequenceRelativeStartingSequence) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "genomeAssembly":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.GenomeAssembly = &v
			case "chromosome":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Chromosome = &v
			case "sequenceCodeableConcept":
				if r.Sequence != nil {
					return fmt.Errorf("multiple values for field \"Sequence\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Sequence = v
			case "sequenceString":
				if r.Sequence != nil {
					return fmt.Errorf("multiple values for field \"Sequence\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Sequence = v
			case "sequenceReference":
				if r.Sequence != nil {
					return fmt.Errorf("multiple values for field \"Sequence\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Sequence = v
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
			case "orientation":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Orientation = &v
			case "strand":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Strand = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *MolecularSequenceRelativeEdit) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "replacementSequence":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReplacementSequence = &v
			case "replacedSequence":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReplacedSequence = &v
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
	if len(name) == 0 || slices.Contains(name, "subject") {
		if r.Subject != nil {
			children = append(children, *r.Subject)
		}
	}
	if len(name) == 0 || slices.Contains(name, "focus") {
		for _, v := range r.Focus {
			children = append(children, v)
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
	if len(name) == 0 || slices.Contains(name, "literal") {
		if r.Literal != nil {
			children = append(children, *r.Literal)
		}
	}
	if len(name) == 0 || slices.Contains(name, "formatted") {
		for _, v := range r.Formatted {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "relative") {
		for _, v := range r.Relative {
			children = append(children, v)
		}
	}
	return children
}
func (r MolecularSequence) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert MolecularSequence to Boolean")
}
func (r MolecularSequence) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert MolecularSequence to String")
}
func (r MolecularSequence) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert MolecularSequence to Integer")
}
func (r MolecularSequence) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert MolecularSequence to Decimal")
}
func (r MolecularSequence) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert MolecularSequence to Date")
}
func (r MolecularSequence) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert MolecularSequence to Time")
}
func (r MolecularSequence) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert MolecularSequence to DateTime")
}
func (r MolecularSequence) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert MolecularSequence to Quantity")
}
func (r MolecularSequence) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *MolecularSequence
	switch other := other.(type) {
	case MolecularSequence:
		o = &other
	case *MolecularSequence:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r MolecularSequence) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(MolecularSequence)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
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
			Name: "Subject",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Focus",
			Type: fhirpath.TypeSpecifier{
				List:      true,
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
			Name: "Literal",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "Formatted",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Attachment",
				Namespace: "FHIR",
			},
		}, {
			Name: "Relative",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "MolecularSequenceRelative",
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
func (r MolecularSequenceRelative) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "coordinateSystem") {
		children = append(children, r.CoordinateSystem)
	}
	if len(name) == 0 || slices.Contains(name, "ordinalPosition") {
		if r.OrdinalPosition != nil {
			children = append(children, *r.OrdinalPosition)
		}
	}
	if len(name) == 0 || slices.Contains(name, "sequenceRange") {
		if r.SequenceRange != nil {
			children = append(children, *r.SequenceRange)
		}
	}
	if len(name) == 0 || slices.Contains(name, "startingSequence") {
		if r.StartingSequence != nil {
			children = append(children, *r.StartingSequence)
		}
	}
	if len(name) == 0 || slices.Contains(name, "edit") {
		for _, v := range r.Edit {
			children = append(children, v)
		}
	}
	return children
}
func (r MolecularSequenceRelative) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert MolecularSequenceRelative to Boolean")
}
func (r MolecularSequenceRelative) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert MolecularSequenceRelative to String")
}
func (r MolecularSequenceRelative) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert MolecularSequenceRelative to Integer")
}
func (r MolecularSequenceRelative) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert MolecularSequenceRelative to Decimal")
}
func (r MolecularSequenceRelative) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert MolecularSequenceRelative to Date")
}
func (r MolecularSequenceRelative) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert MolecularSequenceRelative to Time")
}
func (r MolecularSequenceRelative) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert MolecularSequenceRelative to DateTime")
}
func (r MolecularSequenceRelative) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert MolecularSequenceRelative to Quantity")
}
func (r MolecularSequenceRelative) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *MolecularSequenceRelative
	switch other := other.(type) {
	case MolecularSequenceRelative:
		o = &other
	case *MolecularSequenceRelative:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r MolecularSequenceRelative) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(MolecularSequenceRelative)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r MolecularSequenceRelative) TypeInfo() fhirpath.TypeInfo {
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
			Name: "CoordinateSystem",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "OrdinalPosition",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}, {
			Name: "SequenceRange",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Range",
				Namespace: "FHIR",
			},
		}, {
			Name: "StartingSequence",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "MolecularSequenceRelativeStartingSequence",
				Namespace: "FHIR",
			},
		}, {
			Name: "Edit",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "MolecularSequenceRelativeEdit",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "MolecularSequenceRelative",
			Namespace: "FHIR",
		},
	}
}
func (r MolecularSequenceRelativeStartingSequence) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "genomeAssembly") {
		if r.GenomeAssembly != nil {
			children = append(children, *r.GenomeAssembly)
		}
	}
	if len(name) == 0 || slices.Contains(name, "chromosome") {
		if r.Chromosome != nil {
			children = append(children, *r.Chromosome)
		}
	}
	if len(name) == 0 || slices.Contains(name, "sequence") {
		if r.Sequence != nil {
			children = append(children, r.Sequence)
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
	if len(name) == 0 || slices.Contains(name, "orientation") {
		if r.Orientation != nil {
			children = append(children, *r.Orientation)
		}
	}
	if len(name) == 0 || slices.Contains(name, "strand") {
		if r.Strand != nil {
			children = append(children, *r.Strand)
		}
	}
	return children
}
func (r MolecularSequenceRelativeStartingSequence) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert MolecularSequenceRelativeStartingSequence to Boolean")
}
func (r MolecularSequenceRelativeStartingSequence) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert MolecularSequenceRelativeStartingSequence to String")
}
func (r MolecularSequenceRelativeStartingSequence) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert MolecularSequenceRelativeStartingSequence to Integer")
}
func (r MolecularSequenceRelativeStartingSequence) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert MolecularSequenceRelativeStartingSequence to Decimal")
}
func (r MolecularSequenceRelativeStartingSequence) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert MolecularSequenceRelativeStartingSequence to Date")
}
func (r MolecularSequenceRelativeStartingSequence) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert MolecularSequenceRelativeStartingSequence to Time")
}
func (r MolecularSequenceRelativeStartingSequence) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert MolecularSequenceRelativeStartingSequence to DateTime")
}
func (r MolecularSequenceRelativeStartingSequence) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert MolecularSequenceRelativeStartingSequence to Quantity")
}
func (r MolecularSequenceRelativeStartingSequence) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *MolecularSequenceRelativeStartingSequence
	switch other := other.(type) {
	case MolecularSequenceRelativeStartingSequence:
		o = &other
	case *MolecularSequenceRelativeStartingSequence:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r MolecularSequenceRelativeStartingSequence) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(MolecularSequenceRelativeStartingSequence)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r MolecularSequenceRelativeStartingSequence) TypeInfo() fhirpath.TypeInfo {
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
			Name: "GenomeAssembly",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
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
			Name: "Sequence",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PrimitiveElement",
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
		}, {
			Name: "Orientation",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Strand",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "MolecularSequenceRelativeStartingSequence",
			Namespace: "FHIR",
		},
	}
}
func (r MolecularSequenceRelativeEdit) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "replacementSequence") {
		if r.ReplacementSequence != nil {
			children = append(children, *r.ReplacementSequence)
		}
	}
	if len(name) == 0 || slices.Contains(name, "replacedSequence") {
		if r.ReplacedSequence != nil {
			children = append(children, *r.ReplacedSequence)
		}
	}
	return children
}
func (r MolecularSequenceRelativeEdit) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert MolecularSequenceRelativeEdit to Boolean")
}
func (r MolecularSequenceRelativeEdit) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert MolecularSequenceRelativeEdit to String")
}
func (r MolecularSequenceRelativeEdit) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert MolecularSequenceRelativeEdit to Integer")
}
func (r MolecularSequenceRelativeEdit) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert MolecularSequenceRelativeEdit to Decimal")
}
func (r MolecularSequenceRelativeEdit) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert MolecularSequenceRelativeEdit to Date")
}
func (r MolecularSequenceRelativeEdit) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert MolecularSequenceRelativeEdit to Time")
}
func (r MolecularSequenceRelativeEdit) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert MolecularSequenceRelativeEdit to DateTime")
}
func (r MolecularSequenceRelativeEdit) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert MolecularSequenceRelativeEdit to Quantity")
}
func (r MolecularSequenceRelativeEdit) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *MolecularSequenceRelativeEdit
	switch other := other.(type) {
	case MolecularSequenceRelativeEdit:
		o = &other
	case *MolecularSequenceRelativeEdit:
		o = other
	default:
		return false, true
	}
	if o == nil {
		return false, true
	}
	eq, ok := r.Children().Equal(o.Children())
	return eq && ok, true
}
func (r MolecularSequenceRelativeEdit) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(MolecularSequenceRelativeEdit)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r MolecularSequenceRelativeEdit) TypeInfo() fhirpath.TypeInfo {
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
			Name: "ReplacementSequence",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "ReplacedSequence",
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
			Name:      "MolecularSequenceRelativeEdit",
			Namespace: "FHIR",
		},
	}
}
