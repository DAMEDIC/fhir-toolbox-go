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

// Record details about an anatomical structure.  This resource may be used when a coded concept does not provide the necessary detail needed for the use case.
type BodyStructure struct {
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
	// Identifier for this instance of the anatomical structure.
	Identifier []Identifier
	// Whether this body site is in active use.
	Active *Boolean
	// The kind of structure being represented by the body structure at `BodyStructure.location`.  This can define both normal and abnormal morphologies.
	Morphology *CodeableConcept
	// The anatomical location(s) or region(s) of the specimen, lesion, or body structure.
	IncludedStructure []BodyStructureIncludedStructure
	// The anatomical location(s) or region(s) not occupied or represented by the specimen, lesion, or body structure.
	ExcludedStructure []BodyStructureIncludedStructure
	// A summary, characterization or explanation of the body structure.
	Description *Markdown
	// Image or images used to identify a location.
	Image []Attachment
	// The person to which the body site belongs.
	Patient Reference
}

// The anatomical location(s) or region(s) of the specimen, lesion, or body structure.
type BodyStructureIncludedStructure struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Code that represents the included structure.
	Structure CodeableConcept
	// Code that represents the included structure laterality.
	Laterality *CodeableConcept
	// Body locations in relation to a specific body landmark (tatoo, scar, other body structure).
	BodyLandmarkOrientation []BodyStructureIncludedStructureBodyLandmarkOrientation
	// XY or XYZ-coordinate orientation for structure.
	SpatialReference []Reference
	// Code that represents the included structure qualifier.
	Qualifier []CodeableConcept
}

// Body locations in relation to a specific body landmark (tatoo, scar, other body structure).
type BodyStructureIncludedStructureBodyLandmarkOrientation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A description of a landmark on the body used as a reference to locate something else.
	LandmarkDescription []CodeableConcept
	// An description of the direction away from a landmark something is located based on a radial clock dial.
	ClockFacePosition []CodeableConcept
	// The distance in centimeters a certain observation is made from a body landmark.
	DistanceFromLandmark []BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark
	// The surface area a body location is in relation to a landmark.
	SurfaceOrientation []CodeableConcept
}

// The distance in centimeters a certain observation is made from a body landmark.
type BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An instrument, tool, analyzer, etc. used in the measurement.
	Device []CodeableReference
	// The measured distance (e.g., in cm) from a body landmark.
	Value []Quantity
}

func (r BodyStructure) ResourceType() string {
	return "BodyStructure"
}
func (r BodyStructure) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r BodyStructure) MemSize() int {
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
	if r.Active != nil {
		s += r.Active.MemSize()
	}
	if r.Morphology != nil {
		s += r.Morphology.MemSize()
	}
	for _, i := range r.IncludedStructure {
		s += i.MemSize()
	}
	s += (cap(r.IncludedStructure) - len(r.IncludedStructure)) * int(unsafe.Sizeof(BodyStructureIncludedStructure{}))
	for _, i := range r.ExcludedStructure {
		s += i.MemSize()
	}
	s += (cap(r.ExcludedStructure) - len(r.ExcludedStructure)) * int(unsafe.Sizeof(BodyStructureIncludedStructure{}))
	if r.Description != nil {
		s += r.Description.MemSize()
	}
	for _, i := range r.Image {
		s += i.MemSize()
	}
	s += (cap(r.Image) - len(r.Image)) * int(unsafe.Sizeof(Attachment{}))
	s += r.Patient.MemSize() - int(unsafe.Sizeof(r.Patient))
	return s
}
func (r BodyStructureIncludedStructure) MemSize() int {
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
	s += r.Structure.MemSize() - int(unsafe.Sizeof(r.Structure))
	if r.Laterality != nil {
		s += r.Laterality.MemSize()
	}
	for _, i := range r.BodyLandmarkOrientation {
		s += i.MemSize()
	}
	s += (cap(r.BodyLandmarkOrientation) - len(r.BodyLandmarkOrientation)) * int(unsafe.Sizeof(BodyStructureIncludedStructureBodyLandmarkOrientation{}))
	for _, i := range r.SpatialReference {
		s += i.MemSize()
	}
	s += (cap(r.SpatialReference) - len(r.SpatialReference)) * int(unsafe.Sizeof(Reference{}))
	for _, i := range r.Qualifier {
		s += i.MemSize()
	}
	s += (cap(r.Qualifier) - len(r.Qualifier)) * int(unsafe.Sizeof(CodeableConcept{}))
	return s
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) MemSize() int {
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
	for _, i := range r.LandmarkDescription {
		s += i.MemSize()
	}
	s += (cap(r.LandmarkDescription) - len(r.LandmarkDescription)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.ClockFacePosition {
		s += i.MemSize()
	}
	s += (cap(r.ClockFacePosition) - len(r.ClockFacePosition)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.DistanceFromLandmark {
		s += i.MemSize()
	}
	s += (cap(r.DistanceFromLandmark) - len(r.DistanceFromLandmark)) * int(unsafe.Sizeof(BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark{}))
	for _, i := range r.SurfaceOrientation {
		s += i.MemSize()
	}
	s += (cap(r.SurfaceOrientation) - len(r.SurfaceOrientation)) * int(unsafe.Sizeof(CodeableConcept{}))
	return s
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) MemSize() int {
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
	for _, i := range r.Device {
		s += i.MemSize()
	}
	s += (cap(r.Device) - len(r.Device)) * int(unsafe.Sizeof(CodeableReference{}))
	for _, i := range r.Value {
		s += i.MemSize()
	}
	s += (cap(r.Value) - len(r.Value)) * int(unsafe.Sizeof(Quantity{}))
	return s
}
func (r BodyStructure) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r BodyStructureIncludedStructure) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r BodyStructure) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r BodyStructure) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"BodyStructure\""))
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
	if r.Active != nil && r.Active.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"active\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Active)
		if err != nil {
			return err
		}
	}
	if r.Active != nil && (r.Active.Id != nil || r.Active.Extension != nil) {
		p := primitiveElement{Id: r.Active.Id, Extension: r.Active.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_active\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Morphology != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"morphology\":"))
		if err != nil {
			return err
		}
		err = r.Morphology.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.IncludedStructure) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"includedStructure\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.IncludedStructure {
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
	if len(r.ExcludedStructure) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"excludedStructure\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ExcludedStructure {
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
	if r.Description != nil && r.Description.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"description\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Description)
		if err != nil {
			return err
		}
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		p := primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_description\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Image) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"image\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Image {
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
	_, err = w.Write([]byte("\"patient\":"))
	if err != nil {
		return err
	}
	err = r.Patient.marshalJSON(w)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r BodyStructureIncludedStructure) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r BodyStructureIncludedStructure) marshalJSON(w io.Writer) error {
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
	_, err = w.Write([]byte("\"structure\":"))
	if err != nil {
		return err
	}
	err = r.Structure.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.Laterality != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"laterality\":"))
		if err != nil {
			return err
		}
		err = r.Laterality.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.BodyLandmarkOrientation) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"bodyLandmarkOrientation\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.BodyLandmarkOrientation {
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
	if len(r.SpatialReference) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"spatialReference\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.SpatialReference {
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
	if len(r.Qualifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"qualifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Qualifier {
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
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) marshalJSON(w io.Writer) error {
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
	if len(r.LandmarkDescription) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"landmarkDescription\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.LandmarkDescription {
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
	if len(r.ClockFacePosition) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"clockFacePosition\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ClockFacePosition {
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
	if len(r.DistanceFromLandmark) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"distanceFromLandmark\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.DistanceFromLandmark {
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
	if len(r.SurfaceOrientation) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"surfaceOrientation\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.SurfaceOrientation {
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
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) marshalJSON(w io.Writer) error {
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
	if len(r.Device) > 0 {
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
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Device {
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
	if len(r.Value) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"value\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Value {
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
func (r *BodyStructure) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *BodyStructure) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in BodyStructure element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in BodyStructure element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructure element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructure element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructure element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructure element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructure element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructure element", t)
			}
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructure element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructure element", t)
			}
		case "active":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Active == nil {
				r.Active = &Boolean{}
			}
			r.Active.Value = v.Value
		case "_active":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Active == nil {
				r.Active = &Boolean{}
			}
			r.Active.Id = v.Id
			r.Active.Extension = v.Extension
		case "morphology":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Morphology = &v
		case "includedStructure":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructure element", t)
			}
			for d.More() {
				var v BodyStructureIncludedStructure
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.IncludedStructure = append(r.IncludedStructure, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructure element", t)
			}
		case "excludedStructure":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructure element", t)
			}
			for d.More() {
				var v BodyStructureIncludedStructure
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ExcludedStructure = append(r.ExcludedStructure, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructure element", t)
			}
		case "description":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &Markdown{}
			}
			r.Description.Value = v.Value
		case "_description":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &Markdown{}
			}
			r.Description.Id = v.Id
			r.Description.Extension = v.Extension
		case "image":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructure element", t)
			}
			for d.More() {
				var v Attachment
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Image = append(r.Image, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructure element", t)
			}
		case "patient":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Patient = v
		default:
			return fmt.Errorf("invalid field: %s in BodyStructure", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in BodyStructure element", t)
	}
	return nil
}
func (r *BodyStructureIncludedStructure) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in BodyStructureIncludedStructure element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in BodyStructureIncludedStructure element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructure element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructure element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructure element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructure element", t)
			}
		case "structure":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Structure = v
		case "laterality":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Laterality = &v
		case "bodyLandmarkOrientation":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructure element", t)
			}
			for d.More() {
				var v BodyStructureIncludedStructureBodyLandmarkOrientation
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.BodyLandmarkOrientation = append(r.BodyLandmarkOrientation, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructure element", t)
			}
		case "spatialReference":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructure element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.SpatialReference = append(r.SpatialReference, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructure element", t)
			}
		case "qualifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructure element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Qualifier = append(r.Qualifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructure element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in BodyStructureIncludedStructure", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in BodyStructureIncludedStructure element", t)
	}
	return nil
}
func (r *BodyStructureIncludedStructureBodyLandmarkOrientation) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
			}
		case "landmarkDescription":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.LandmarkDescription = append(r.LandmarkDescription, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
			}
		case "clockFacePosition":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ClockFacePosition = append(r.ClockFacePosition, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
			}
		case "distanceFromLandmark":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
			}
			for d.More() {
				var v BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.DistanceFromLandmark = append(r.DistanceFromLandmark, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
			}
		case "surfaceOrientation":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.SurfaceOrientation = append(r.SurfaceOrientation, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in BodyStructureIncludedStructureBodyLandmarkOrientation", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in BodyStructureIncludedStructureBodyLandmarkOrientation element", t)
	}
	return nil
}
func (r *BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark element", t)
			}
		case "device":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark element", t)
			}
			for d.More() {
				var v CodeableReference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Device = append(r.Device, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark element", t)
			}
		case "value":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark element", t)
			}
			for d.More() {
				var v Quantity
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Value = append(r.Value, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark element", t)
	}
	return nil
}
func (r BodyStructure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "BodyStructure"
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
	err = e.EncodeElement(r.Active, xml.StartElement{Name: xml.Name{Local: "active"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Morphology, xml.StartElement{Name: xml.Name{Local: "morphology"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IncludedStructure, xml.StartElement{Name: xml.Name{Local: "includedStructure"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ExcludedStructure, xml.StartElement{Name: xml.Name{Local: "excludedStructure"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Image, xml.StartElement{Name: xml.Name{Local: "image"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Patient, xml.StartElement{Name: xml.Name{Local: "patient"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r BodyStructureIncludedStructure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Structure, xml.StartElement{Name: xml.Name{Local: "structure"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Laterality, xml.StartElement{Name: xml.Name{Local: "laterality"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BodyLandmarkOrientation, xml.StartElement{Name: xml.Name{Local: "bodyLandmarkOrientation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SpatialReference, xml.StartElement{Name: xml.Name{Local: "spatialReference"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Qualifier, xml.StartElement{Name: xml.Name{Local: "qualifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.LandmarkDescription, xml.StartElement{Name: xml.Name{Local: "landmarkDescription"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ClockFacePosition, xml.StartElement{Name: xml.Name{Local: "clockFacePosition"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DistanceFromLandmark, xml.StartElement{Name: xml.Name{Local: "distanceFromLandmark"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SurfaceOrientation, xml.StartElement{Name: xml.Name{Local: "surfaceOrientation"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Device, xml.StartElement{Name: xml.Name{Local: "device"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Value, xml.StartElement{Name: xml.Name{Local: "value"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *BodyStructure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "active":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Active = &v
			case "morphology":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Morphology = &v
			case "includedStructure":
				var v BodyStructureIncludedStructure
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IncludedStructure = append(r.IncludedStructure, v)
			case "excludedStructure":
				var v BodyStructureIncludedStructure
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ExcludedStructure = append(r.ExcludedStructure, v)
			case "description":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "image":
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Image = append(r.Image, v)
			case "patient":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Patient = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *BodyStructureIncludedStructure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "structure":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Structure = v
			case "laterality":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Laterality = &v
			case "bodyLandmarkOrientation":
				var v BodyStructureIncludedStructureBodyLandmarkOrientation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BodyLandmarkOrientation = append(r.BodyLandmarkOrientation, v)
			case "spatialReference":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SpatialReference = append(r.SpatialReference, v)
			case "qualifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Qualifier = append(r.Qualifier, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *BodyStructureIncludedStructureBodyLandmarkOrientation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "landmarkDescription":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LandmarkDescription = append(r.LandmarkDescription, v)
			case "clockFacePosition":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ClockFacePosition = append(r.ClockFacePosition, v)
			case "distanceFromLandmark":
				var v BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DistanceFromLandmark = append(r.DistanceFromLandmark, v)
			case "surfaceOrientation":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SurfaceOrientation = append(r.SurfaceOrientation, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "device":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Device = append(r.Device, v)
			case "value":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = append(r.Value, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r BodyStructure) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "active") {
		if r.Active != nil {
			children = append(children, *r.Active)
		}
	}
	if len(name) == 0 || slices.Contains(name, "morphology") {
		if r.Morphology != nil {
			children = append(children, *r.Morphology)
		}
	}
	if len(name) == 0 || slices.Contains(name, "includedStructure") {
		for _, v := range r.IncludedStructure {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "excludedStructure") {
		for _, v := range r.ExcludedStructure {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "description") {
		if r.Description != nil {
			children = append(children, *r.Description)
		}
	}
	if len(name) == 0 || slices.Contains(name, "image") {
		for _, v := range r.Image {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "patient") {
		children = append(children, r.Patient)
	}
	return children
}
func (r BodyStructure) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert BodyStructure to Boolean")
}
func (r BodyStructure) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert BodyStructure to String")
}
func (r BodyStructure) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert BodyStructure to Integer")
}
func (r BodyStructure) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert BodyStructure to Decimal")
}
func (r BodyStructure) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert BodyStructure to Date")
}
func (r BodyStructure) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert BodyStructure to Time")
}
func (r BodyStructure) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert BodyStructure to DateTime")
}
func (r BodyStructure) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert BodyStructure to Quantity")
}
func (r BodyStructure) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *BodyStructure
	switch other := other.(type) {
	case BodyStructure:
		o = &other
	case *BodyStructure:
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
func (r BodyStructure) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r BodyStructure) TypeInfo() fhirpath.TypeInfo {
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
			Name: "Active",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Morphology",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "IncludedStructure",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "BodyStructureIncludedStructure",
				Namespace: "FHIR",
			},
		}, {
			Name: "ExcludedStructure",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "BodyStructureIncludedStructure",
				Namespace: "FHIR",
			},
		}, {
			Name: "Description",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Markdown",
				Namespace: "FHIR",
			},
		}, {
			Name: "Image",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Attachment",
				Namespace: "FHIR",
			},
		}, {
			Name: "Patient",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "BodyStructure",
			Namespace: "FHIR",
		},
	}
}
func (r BodyStructureIncludedStructure) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "structure") {
		children = append(children, r.Structure)
	}
	if len(name) == 0 || slices.Contains(name, "laterality") {
		if r.Laterality != nil {
			children = append(children, *r.Laterality)
		}
	}
	if len(name) == 0 || slices.Contains(name, "bodyLandmarkOrientation") {
		for _, v := range r.BodyLandmarkOrientation {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "spatialReference") {
		for _, v := range r.SpatialReference {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "qualifier") {
		for _, v := range r.Qualifier {
			children = append(children, v)
		}
	}
	return children
}
func (r BodyStructureIncludedStructure) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert BodyStructureIncludedStructure to Boolean")
}
func (r BodyStructureIncludedStructure) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert BodyStructureIncludedStructure to String")
}
func (r BodyStructureIncludedStructure) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert BodyStructureIncludedStructure to Integer")
}
func (r BodyStructureIncludedStructure) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert BodyStructureIncludedStructure to Decimal")
}
func (r BodyStructureIncludedStructure) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert BodyStructureIncludedStructure to Date")
}
func (r BodyStructureIncludedStructure) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert BodyStructureIncludedStructure to Time")
}
func (r BodyStructureIncludedStructure) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert BodyStructureIncludedStructure to DateTime")
}
func (r BodyStructureIncludedStructure) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert BodyStructureIncludedStructure to Quantity")
}
func (r BodyStructureIncludedStructure) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *BodyStructureIncludedStructure
	switch other := other.(type) {
	case BodyStructureIncludedStructure:
		o = &other
	case *BodyStructureIncludedStructure:
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
func (r BodyStructureIncludedStructure) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r BodyStructureIncludedStructure) TypeInfo() fhirpath.TypeInfo {
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
			Name: "Structure",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Laterality",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "BodyLandmarkOrientation",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "BodyStructureIncludedStructureBodyLandmarkOrientation",
				Namespace: "FHIR",
			},
		}, {
			Name: "SpatialReference",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Qualifier",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "BodyStructureIncludedStructure",
			Namespace: "FHIR",
		},
	}
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "landmarkDescription") {
		for _, v := range r.LandmarkDescription {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "clockFacePosition") {
		for _, v := range r.ClockFacePosition {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "distanceFromLandmark") {
		for _, v := range r.DistanceFromLandmark {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "surfaceOrientation") {
		for _, v := range r.SurfaceOrientation {
			children = append(children, v)
		}
	}
	return children
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientation to Boolean")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientation to String")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientation to Integer")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientation to Decimal")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientation to Date")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientation to Time")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientation to DateTime")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientation to Quantity")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *BodyStructureIncludedStructureBodyLandmarkOrientation
	switch other := other.(type) {
	case BodyStructureIncludedStructureBodyLandmarkOrientation:
		o = &other
	case *BodyStructureIncludedStructureBodyLandmarkOrientation:
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
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientation) TypeInfo() fhirpath.TypeInfo {
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
			Name: "LandmarkDescription",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ClockFacePosition",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "DistanceFromLandmark",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark",
				Namespace: "FHIR",
			},
		}, {
			Name: "SurfaceOrientation",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "BodyStructureIncludedStructureBodyLandmarkOrientation",
			Namespace: "FHIR",
		},
	}
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "device") {
		for _, v := range r.Device {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "value") {
		for _, v := range r.Value {
			children = append(children, v)
		}
	}
	return children
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark to Boolean")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark to String")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark to Integer")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark to Decimal")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark to Date")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark to Time")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark to DateTime")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark to Quantity")
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark
	switch other := other.(type) {
	case BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark:
		o = &other
	case *BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark:
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
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, ok := r.Equal(other)
	return eq && ok
}
func (r BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark) TypeInfo() fhirpath.TypeInfo {
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
			Name: "Device",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableReference",
				Namespace: "FHIR",
			},
		}, {
			Name: "Value",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Quantity",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark",
			Namespace: "FHIR",
		},
	}
}
