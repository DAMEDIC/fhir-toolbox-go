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

// Base StructureDefinition for Dosage Type: Indicates how the medication is/was taken or should be taken by the patient.
type Dosage struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates the order in which the dosage instructions should be applied or interpreted.
	Sequence *Integer
	// Free text dosage instructions e.g. SIG.
	Text *String
	// Supplemental instructions to the patient on how to take the medication  (e.g. "with meals" or"take half to one hour before food") or warnings for the patient about the medication (e.g. "may cause drowsiness" or "avoid exposure of skin to direct sunlight or sunlamps").
	AdditionalInstruction []CodeableConcept
	// Instructions in terms that are understood by the patient or consumer.
	PatientInstruction *String
	// When medication should be administered.
	Timing *Timing
	// Indicates whether the Medication is only taken when needed within a specific dosing schedule (Boolean option), or it indicates the precondition for taking the Medication (CodeableConcept).
	AsNeeded isDosageAsNeeded
	// Body site to administer to.
	Site *CodeableConcept
	// How drug should enter body.
	Route *CodeableConcept
	// Technique for administering medication.
	Method *CodeableConcept
	// The amount of medication administered.
	DoseAndRate []DosageDoseAndRate
	// Upper limit on medication per unit of time.
	MaxDosePerPeriod *Ratio
	// Upper limit on medication per administration.
	MaxDosePerAdministration *Quantity
	// Upper limit on medication per lifetime of the patient.
	MaxDosePerLifetime *Quantity
}
type isDosageAsNeeded interface {
	model.Element
	isDosageAsNeeded()
}

func (r Boolean) isDosageAsNeeded()         {}
func (r CodeableConcept) isDosageAsNeeded() {}

// The amount of medication administered.
type DosageDoseAndRate struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The kind of dose or rate specified, for example, ordered or calculated.
	Type *CodeableConcept
	// Amount of medication per dose.
	Dose isDosageDoseAndRateDose
	// Amount of medication per unit of time.
	Rate isDosageDoseAndRateRate
}
type isDosageDoseAndRateDose interface {
	model.Element
	isDosageDoseAndRateDose()
}

func (r Range) isDosageDoseAndRateDose()    {}
func (r Quantity) isDosageDoseAndRateDose() {}

type isDosageDoseAndRateRate interface {
	model.Element
	isDosageDoseAndRateRate()
}

func (r Ratio) isDosageDoseAndRateRate()    {}
func (r Range) isDosageDoseAndRateRate()    {}
func (r Quantity) isDosageDoseAndRateRate() {}
func (r Dosage) MemSize() int {
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
	if r.Sequence != nil {
		s += r.Sequence.MemSize()
	}
	if r.Text != nil {
		s += r.Text.MemSize()
	}
	for _, i := range r.AdditionalInstruction {
		s += i.MemSize()
	}
	s += (cap(r.AdditionalInstruction) - len(r.AdditionalInstruction)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.PatientInstruction != nil {
		s += r.PatientInstruction.MemSize()
	}
	if r.Timing != nil {
		s += r.Timing.MemSize()
	}
	if r.AsNeeded != nil {
		s += r.AsNeeded.MemSize()
	}
	if r.Site != nil {
		s += r.Site.MemSize()
	}
	if r.Route != nil {
		s += r.Route.MemSize()
	}
	if r.Method != nil {
		s += r.Method.MemSize()
	}
	for _, i := range r.DoseAndRate {
		s += i.MemSize()
	}
	s += (cap(r.DoseAndRate) - len(r.DoseAndRate)) * int(unsafe.Sizeof(DosageDoseAndRate{}))
	if r.MaxDosePerPeriod != nil {
		s += r.MaxDosePerPeriod.MemSize()
	}
	if r.MaxDosePerAdministration != nil {
		s += r.MaxDosePerAdministration.MemSize()
	}
	if r.MaxDosePerLifetime != nil {
		s += r.MaxDosePerLifetime.MemSize()
	}
	return s
}
func (r DosageDoseAndRate) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	if r.Type != nil {
		s += r.Type.MemSize()
	}
	if r.Dose != nil {
		s += r.Dose.MemSize()
	}
	if r.Rate != nil {
		s += r.Rate.MemSize()
	}
	return s
}
func (r Dosage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Dosage) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Dosage) marshalJSON(w io.Writer) error {
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
	if r.Sequence != nil && r.Sequence.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sequence\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Sequence)
		if err != nil {
			return err
		}
	}
	if r.Sequence != nil && (r.Sequence.Id != nil || r.Sequence.Extension != nil) {
		p := primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_sequence\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Text != nil && r.Text.Value != nil {
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
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Text)
		if err != nil {
			return err
		}
	}
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		p := primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_text\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.AdditionalInstruction) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"additionalInstruction\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.AdditionalInstruction {
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
	if r.PatientInstruction != nil && r.PatientInstruction.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"patientInstruction\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.PatientInstruction)
		if err != nil {
			return err
		}
	}
	if r.PatientInstruction != nil && (r.PatientInstruction.Id != nil || r.PatientInstruction.Extension != nil) {
		p := primitiveElement{Id: r.PatientInstruction.Id, Extension: r.PatientInstruction.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_patientInstruction\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Timing != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timing\":"))
		if err != nil {
			return err
		}
		err = r.Timing.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	switch v := r.AsNeeded.(type) {
	case Boolean:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"asNeededBoolean\":"))
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
			_, err = w.Write([]byte("\"_asNeededBoolean\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Boolean:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"asNeededBoolean\":"))
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
			_, err = w.Write([]byte("\"_asNeededBoolean\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case CodeableConcept:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"asNeededCodeableConcept\":"))
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
		_, err = w.Write([]byte("\"asNeededCodeableConcept\":"))
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
	if r.Site != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"site\":"))
		if err != nil {
			return err
		}
		err = r.Site.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Route != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"route\":"))
		if err != nil {
			return err
		}
		err = r.Route.marshalJSON(w)
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
	if len(r.DoseAndRate) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"doseAndRate\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.DoseAndRate {
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
	if r.MaxDosePerPeriod != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"maxDosePerPeriod\":"))
		if err != nil {
			return err
		}
		err = r.MaxDosePerPeriod.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.MaxDosePerAdministration != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"maxDosePerAdministration\":"))
		if err != nil {
			return err
		}
		err = r.MaxDosePerAdministration.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.MaxDosePerLifetime != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"maxDosePerLifetime\":"))
		if err != nil {
			return err
		}
		err = r.MaxDosePerLifetime.marshalJSON(w)
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
func (r DosageDoseAndRate) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r DosageDoseAndRate) marshalJSON(w io.Writer) error {
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
	if r.Type != nil {
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
		err = r.Type.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	switch v := r.Dose.(type) {
	case Range:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"doseRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Range:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"doseRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Quantity:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"doseQuantity\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Quantity:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"doseQuantity\":"))
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
	switch v := r.Rate.(type) {
	case Ratio:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"rateRatio\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Ratio:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"rateRatio\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Range:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"rateRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Range:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"rateRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Quantity:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"rateQuantity\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Quantity:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"rateQuantity\":"))
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
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r *Dosage) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in Dosage element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in Dosage element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in Dosage element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Dosage element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Dosage element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Dosage element", t)
			}
		case "sequence":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Sequence == nil {
				r.Sequence = &Integer{}
			}
			r.Sequence.Value = v.Value
		case "_sequence":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Sequence == nil {
				r.Sequence = &Integer{}
			}
			r.Sequence.Id = v.Id
			r.Sequence.Extension = v.Extension
		case "text":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Text == nil {
				r.Text = &String{}
			}
			r.Text.Value = v.Value
		case "_text":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Text == nil {
				r.Text = &String{}
			}
			r.Text.Id = v.Id
			r.Text.Extension = v.Extension
		case "additionalInstruction":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Dosage element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.AdditionalInstruction = append(r.AdditionalInstruction, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Dosage element", t)
			}
		case "patientInstruction":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.PatientInstruction == nil {
				r.PatientInstruction = &String{}
			}
			r.PatientInstruction.Value = v.Value
		case "_patientInstruction":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.PatientInstruction == nil {
				r.PatientInstruction = &String{}
			}
			r.PatientInstruction.Id = v.Id
			r.PatientInstruction.Extension = v.Extension
		case "timing":
			var v Timing
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Timing = &v
		case "asNeededBoolean":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.AsNeeded != nil {
				r.AsNeeded = Boolean{
					Extension: r.AsNeeded.(Boolean).Extension,
					Id:        r.AsNeeded.(Boolean).Id,
					Value:     v.Value,
				}
			} else {
				r.AsNeeded = v
			}
		case "_asNeededBoolean":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.AsNeeded != nil {
				r.AsNeeded = Boolean{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.AsNeeded.(Boolean).Value,
				}
			} else {
				r.AsNeeded = Boolean{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "asNeededCodeableConcept":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.AsNeeded = v
		case "site":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Site = &v
		case "route":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Route = &v
		case "method":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Method = &v
		case "doseAndRate":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Dosage element", t)
			}
			for d.More() {
				var v DosageDoseAndRate
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.DoseAndRate = append(r.DoseAndRate, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Dosage element", t)
			}
		case "maxDosePerPeriod":
			var v Ratio
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.MaxDosePerPeriod = &v
		case "maxDosePerAdministration":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.MaxDosePerAdministration = &v
		case "maxDosePerLifetime":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.MaxDosePerLifetime = &v
		default:
			return fmt.Errorf("invalid field: %s in Dosage", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in Dosage element", t)
	}
	return nil
}
func (r *DosageDoseAndRate) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in DosageDoseAndRate element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in DosageDoseAndRate element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in DosageDoseAndRate element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in DosageDoseAndRate element", t)
			}
		case "type":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = &v
		case "doseRange":
			var v Range
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Dose = v
		case "doseQuantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Dose = v
		case "rateRatio":
			var v Ratio
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Rate = v
		case "rateRange":
			var v Range
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Rate = v
		case "rateQuantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Rate = v
		default:
			return fmt.Errorf("invalid field: %s in DosageDoseAndRate", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in DosageDoseAndRate element", t)
	}
	return nil
}
func (r Dosage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Sequence, xml.StartElement{Name: xml.Name{Local: "sequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AdditionalInstruction, xml.StartElement{Name: xml.Name{Local: "additionalInstruction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PatientInstruction, xml.StartElement{Name: xml.Name{Local: "patientInstruction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Timing, xml.StartElement{Name: xml.Name{Local: "timing"}})
	if err != nil {
		return err
	}
	switch v := r.AsNeeded.(type) {
	case Boolean, *Boolean:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "asNeededBoolean"}})
		if err != nil {
			return err
		}
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "asNeededCodeableConcept"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Site, xml.StartElement{Name: xml.Name{Local: "site"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Route, xml.StartElement{Name: xml.Name{Local: "route"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Method, xml.StartElement{Name: xml.Name{Local: "method"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DoseAndRate, xml.StartElement{Name: xml.Name{Local: "doseAndRate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MaxDosePerPeriod, xml.StartElement{Name: xml.Name{Local: "maxDosePerPeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MaxDosePerAdministration, xml.StartElement{Name: xml.Name{Local: "maxDosePerAdministration"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MaxDosePerLifetime, xml.StartElement{Name: xml.Name{Local: "maxDosePerLifetime"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r DosageDoseAndRate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	switch v := r.Dose.(type) {
	case Range, *Range:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "doseRange"}})
		if err != nil {
			return err
		}
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "doseQuantity"}})
		if err != nil {
			return err
		}
	}
	switch v := r.Rate.(type) {
	case Ratio, *Ratio:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "rateRatio"}})
		if err != nil {
			return err
		}
	case Range, *Range:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "rateRange"}})
		if err != nil {
			return err
		}
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "rateQuantity"}})
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
func (r *Dosage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "sequence":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Sequence = &v
			case "text":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = &v
			case "additionalInstruction":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AdditionalInstruction = append(r.AdditionalInstruction, v)
			case "patientInstruction":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PatientInstruction = &v
			case "timing":
				var v Timing
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Timing = &v
			case "asNeededBoolean":
				if r.AsNeeded != nil {
					return fmt.Errorf("multiple values for field \"AsNeeded\"")
				}
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AsNeeded = v
			case "asNeededCodeableConcept":
				if r.AsNeeded != nil {
					return fmt.Errorf("multiple values for field \"AsNeeded\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AsNeeded = v
			case "site":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Site = &v
			case "route":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Route = &v
			case "method":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Method = &v
			case "doseAndRate":
				var v DosageDoseAndRate
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DoseAndRate = append(r.DoseAndRate, v)
			case "maxDosePerPeriod":
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MaxDosePerPeriod = &v
			case "maxDosePerAdministration":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MaxDosePerAdministration = &v
			case "maxDosePerLifetime":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MaxDosePerLifetime = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *DosageDoseAndRate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "doseRange":
				if r.Dose != nil {
					return fmt.Errorf("multiple values for field \"Dose\"")
				}
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Dose = v
			case "doseQuantity":
				if r.Dose != nil {
					return fmt.Errorf("multiple values for field \"Dose\"")
				}
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Dose = v
			case "rateRatio":
				if r.Rate != nil {
					return fmt.Errorf("multiple values for field \"Rate\"")
				}
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Rate = v
			case "rateRange":
				if r.Rate != nil {
					return fmt.Errorf("multiple values for field \"Rate\"")
				}
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Rate = v
			case "rateQuantity":
				if r.Rate != nil {
					return fmt.Errorf("multiple values for field \"Rate\"")
				}
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Rate = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Dosage) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "sequence") {
		if r.Sequence != nil {
			children = append(children, *r.Sequence)
		}
	}
	if len(name) == 0 || slices.Contains(name, "text") {
		if r.Text != nil {
			children = append(children, *r.Text)
		}
	}
	if len(name) == 0 || slices.Contains(name, "additionalInstruction") {
		for _, v := range r.AdditionalInstruction {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "patientInstruction") {
		if r.PatientInstruction != nil {
			children = append(children, *r.PatientInstruction)
		}
	}
	if len(name) == 0 || slices.Contains(name, "timing") {
		if r.Timing != nil {
			children = append(children, *r.Timing)
		}
	}
	if len(name) == 0 || slices.Contains(name, "asNeeded") {
		if r.AsNeeded != nil {
			children = append(children, r.AsNeeded)
		}
	}
	if len(name) == 0 || slices.Contains(name, "site") {
		if r.Site != nil {
			children = append(children, *r.Site)
		}
	}
	if len(name) == 0 || slices.Contains(name, "route") {
		if r.Route != nil {
			children = append(children, *r.Route)
		}
	}
	if len(name) == 0 || slices.Contains(name, "method") {
		if r.Method != nil {
			children = append(children, *r.Method)
		}
	}
	if len(name) == 0 || slices.Contains(name, "doseAndRate") {
		for _, v := range r.DoseAndRate {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "maxDosePerPeriod") {
		if r.MaxDosePerPeriod != nil {
			children = append(children, *r.MaxDosePerPeriod)
		}
	}
	if len(name) == 0 || slices.Contains(name, "maxDosePerAdministration") {
		if r.MaxDosePerAdministration != nil {
			children = append(children, *r.MaxDosePerAdministration)
		}
	}
	if len(name) == 0 || slices.Contains(name, "maxDosePerLifetime") {
		if r.MaxDosePerLifetime != nil {
			children = append(children, *r.MaxDosePerLifetime)
		}
	}
	return children
}
func (r Dosage) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert Dosage to Boolean")
}
func (r Dosage) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert Dosage to String")
}
func (r Dosage) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert Dosage to Integer")
}
func (r Dosage) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert Dosage to Decimal")
}
func (r Dosage) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert Dosage to Date")
}
func (r Dosage) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert Dosage to Time")
}
func (r Dosage) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert Dosage to DateTime")
}
func (r Dosage) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert Dosage to Quantity")
}
func (r Dosage) TypeInfo() fhirpath.TypeInfo {
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
			Name: "Sequence",
			Type: "FHIR.Integer",
		}, {
			Name: "Text",
			Type: "FHIR.String",
		}, {
			Name: "AdditionalInstruction",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "PatientInstruction",
			Type: "FHIR.String",
		}, {
			Name: "Timing",
			Type: "FHIR.Timing",
		}, {
			Name: "AsNeeded",
			Type: "FHIR.PrimitiveElement",
		}, {
			Name: "Site",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Route",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Method",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "DoseAndRate",
			Type: "List<FHIR.DosageDoseAndRate>",
		}, {
			Name: "MaxDosePerPeriod",
			Type: "FHIR.Ratio",
		}, {
			Name: "MaxDosePerAdministration",
			Type: "FHIR.Quantity",
		}, {
			Name: "MaxDosePerLifetime",
			Type: "FHIR.Quantity",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "Dosage",
			Namespace: "FHIR",
		},
	}
}
func (r DosageDoseAndRate) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "type") {
		if r.Type != nil {
			children = append(children, *r.Type)
		}
	}
	if len(name) == 0 || slices.Contains(name, "dose") {
		if r.Dose != nil {
			children = append(children, r.Dose)
		}
	}
	if len(name) == 0 || slices.Contains(name, "rate") {
		if r.Rate != nil {
			children = append(children, r.Rate)
		}
	}
	return children
}
func (r DosageDoseAndRate) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert DosageDoseAndRate to Boolean")
}
func (r DosageDoseAndRate) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert DosageDoseAndRate to String")
}
func (r DosageDoseAndRate) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert DosageDoseAndRate to Integer")
}
func (r DosageDoseAndRate) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert DosageDoseAndRate to Decimal")
}
func (r DosageDoseAndRate) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert DosageDoseAndRate to Date")
}
func (r DosageDoseAndRate) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert DosageDoseAndRate to Time")
}
func (r DosageDoseAndRate) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert DosageDoseAndRate to DateTime")
}
func (r DosageDoseAndRate) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert DosageDoseAndRate to Quantity")
}
func (r DosageDoseAndRate) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Type",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Dose",
			Type: "FHIR.PrimitiveElement",
		}, {
			Name: "Rate",
			Type: "FHIR.PrimitiveElement",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "DosageDoseAndRate",
			Namespace: "FHIR",
		},
	}
}
