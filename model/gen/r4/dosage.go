package r4

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
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
	isDosageAsNeeded()
}

func (r Boolean) isDosageAsNeeded()         {}
func (r CodeableConcept) isDosageAsNeeded() {}

type jsonDosage struct {
	Id                                 *string             `json:"id,omitempty"`
	Extension                          []Extension         `json:"extension,omitempty"`
	ModifierExtension                  []Extension         `json:"modifierExtension,omitempty"`
	Sequence                           *Integer            `json:"sequence,omitempty"`
	SequencePrimitiveElement           *primitiveElement   `json:"_sequence,omitempty"`
	Text                               *String             `json:"text,omitempty"`
	TextPrimitiveElement               *primitiveElement   `json:"_text,omitempty"`
	AdditionalInstruction              []CodeableConcept   `json:"additionalInstruction,omitempty"`
	PatientInstruction                 *String             `json:"patientInstruction,omitempty"`
	PatientInstructionPrimitiveElement *primitiveElement   `json:"_patientInstruction,omitempty"`
	Timing                             *Timing             `json:"timing,omitempty"`
	AsNeededBoolean                    *Boolean            `json:"asNeededBoolean,omitempty"`
	AsNeededBooleanPrimitiveElement    *primitiveElement   `json:"_asNeededBoolean,omitempty"`
	AsNeededCodeableConcept            *CodeableConcept    `json:"asNeededCodeableConcept,omitempty"`
	Site                               *CodeableConcept    `json:"site,omitempty"`
	Route                              *CodeableConcept    `json:"route,omitempty"`
	Method                             *CodeableConcept    `json:"method,omitempty"`
	DoseAndRate                        []DosageDoseAndRate `json:"doseAndRate,omitempty"`
	MaxDosePerPeriod                   *Ratio              `json:"maxDosePerPeriod,omitempty"`
	MaxDosePerAdministration           *Quantity           `json:"maxDosePerAdministration,omitempty"`
	MaxDosePerLifetime                 *Quantity           `json:"maxDosePerLifetime,omitempty"`
}

func (r Dosage) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Dosage) marshalJSON() jsonDosage {
	m := jsonDosage{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Sequence != nil && r.Sequence.Value != nil {
		m.Sequence = r.Sequence
	}
	if r.Sequence != nil && (r.Sequence.Id != nil || r.Sequence.Extension != nil) {
		m.SequencePrimitiveElement = &primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
	}
	if r.Text != nil && r.Text.Value != nil {
		m.Text = r.Text
	}
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	m.AdditionalInstruction = r.AdditionalInstruction
	if r.PatientInstruction != nil && r.PatientInstruction.Value != nil {
		m.PatientInstruction = r.PatientInstruction
	}
	if r.PatientInstruction != nil && (r.PatientInstruction.Id != nil || r.PatientInstruction.Extension != nil) {
		m.PatientInstructionPrimitiveElement = &primitiveElement{Id: r.PatientInstruction.Id, Extension: r.PatientInstruction.Extension}
	}
	m.Timing = r.Timing
	switch v := r.AsNeeded.(type) {
	case Boolean:
		if v.Value != nil {
			m.AsNeededBoolean = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.AsNeededBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		if v.Value != nil {
			m.AsNeededBoolean = v
		}
		if v.Id != nil || v.Extension != nil {
			m.AsNeededBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case CodeableConcept:
		m.AsNeededCodeableConcept = &v
	case *CodeableConcept:
		m.AsNeededCodeableConcept = v
	}
	m.Site = r.Site
	m.Route = r.Route
	m.Method = r.Method
	m.DoseAndRate = r.DoseAndRate
	m.MaxDosePerPeriod = r.MaxDosePerPeriod
	m.MaxDosePerAdministration = r.MaxDosePerAdministration
	m.MaxDosePerLifetime = r.MaxDosePerLifetime
	return m
}
func (r *Dosage) UnmarshalJSON(b []byte) error {
	var m jsonDosage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Dosage) unmarshalJSON(m jsonDosage) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Sequence = m.Sequence
	if m.SequencePrimitiveElement != nil {
		if r.Sequence == nil {
			r.Sequence = &Integer{}
		}
		r.Sequence.Id = m.SequencePrimitiveElement.Id
		r.Sequence.Extension = m.SequencePrimitiveElement.Extension
	}
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		if r.Text == nil {
			r.Text = &String{}
		}
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	r.AdditionalInstruction = m.AdditionalInstruction
	r.PatientInstruction = m.PatientInstruction
	if m.PatientInstructionPrimitiveElement != nil {
		if r.PatientInstruction == nil {
			r.PatientInstruction = &String{}
		}
		r.PatientInstruction.Id = m.PatientInstructionPrimitiveElement.Id
		r.PatientInstruction.Extension = m.PatientInstructionPrimitiveElement.Extension
	}
	r.Timing = m.Timing
	if m.AsNeededBoolean != nil || m.AsNeededBooleanPrimitiveElement != nil {
		if r.AsNeeded != nil {
			return fmt.Errorf("multiple values for field \"AsNeeded\"")
		}
		v := m.AsNeededBoolean
		if m.AsNeededBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.AsNeededBooleanPrimitiveElement.Id
			v.Extension = m.AsNeededBooleanPrimitiveElement.Extension
		}
		r.AsNeeded = v
	}
	if m.AsNeededCodeableConcept != nil {
		if r.AsNeeded != nil {
			return fmt.Errorf("multiple values for field \"AsNeeded\"")
		}
		v := m.AsNeededCodeableConcept
		r.AsNeeded = v
	}
	r.Site = m.Site
	r.Route = m.Route
	r.Method = m.Method
	r.DoseAndRate = m.DoseAndRate
	r.MaxDosePerPeriod = m.MaxDosePerPeriod
	r.MaxDosePerAdministration = m.MaxDosePerAdministration
	r.MaxDosePerLifetime = m.MaxDosePerLifetime
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
func (r Dosage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

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
	isDosageDoseAndRateDose()
}

func (r Range) isDosageDoseAndRateDose()    {}
func (r Quantity) isDosageDoseAndRateDose() {}

type isDosageDoseAndRateRate interface {
	isDosageDoseAndRateRate()
}

func (r Ratio) isDosageDoseAndRateRate()    {}
func (r Range) isDosageDoseAndRateRate()    {}
func (r Quantity) isDosageDoseAndRateRate() {}

type jsonDosageDoseAndRate struct {
	Id           *string          `json:"id,omitempty"`
	Extension    []Extension      `json:"extension,omitempty"`
	Type         *CodeableConcept `json:"type,omitempty"`
	DoseRange    *Range           `json:"doseRange,omitempty"`
	DoseQuantity *Quantity        `json:"doseQuantity,omitempty"`
	RateRatio    *Ratio           `json:"rateRatio,omitempty"`
	RateRange    *Range           `json:"rateRange,omitempty"`
	RateQuantity *Quantity        `json:"rateQuantity,omitempty"`
}

func (r DosageDoseAndRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DosageDoseAndRate) marshalJSON() jsonDosageDoseAndRate {
	m := jsonDosageDoseAndRate{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Type = r.Type
	switch v := r.Dose.(type) {
	case Range:
		m.DoseRange = &v
	case *Range:
		m.DoseRange = v
	case Quantity:
		m.DoseQuantity = &v
	case *Quantity:
		m.DoseQuantity = v
	}
	switch v := r.Rate.(type) {
	case Ratio:
		m.RateRatio = &v
	case *Ratio:
		m.RateRatio = v
	case Range:
		m.RateRange = &v
	case *Range:
		m.RateRange = v
	case Quantity:
		m.RateQuantity = &v
	case *Quantity:
		m.RateQuantity = v
	}
	return m
}
func (r *DosageDoseAndRate) UnmarshalJSON(b []byte) error {
	var m jsonDosageDoseAndRate
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DosageDoseAndRate) unmarshalJSON(m jsonDosageDoseAndRate) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Type = m.Type
	if m.DoseRange != nil {
		if r.Dose != nil {
			return fmt.Errorf("multiple values for field \"Dose\"")
		}
		v := m.DoseRange
		r.Dose = v
	}
	if m.DoseQuantity != nil {
		if r.Dose != nil {
			return fmt.Errorf("multiple values for field \"Dose\"")
		}
		v := m.DoseQuantity
		r.Dose = v
	}
	if m.RateRatio != nil {
		if r.Rate != nil {
			return fmt.Errorf("multiple values for field \"Rate\"")
		}
		v := m.RateRatio
		r.Rate = v
	}
	if m.RateRange != nil {
		if r.Rate != nil {
			return fmt.Errorf("multiple values for field \"Rate\"")
		}
		v := m.RateRange
		r.Rate = v
	}
	if m.RateQuantity != nil {
		if r.Rate != nil {
			return fmt.Errorf("multiple values for field \"Rate\"")
		}
		v := m.RateQuantity
		r.Rate = v
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
func (r DosageDoseAndRate) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
