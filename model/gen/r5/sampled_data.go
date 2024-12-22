package r5

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"unsafe"
)

// SampledData Type: A series of measurements taken by a device, with upper and lower limits. There may be more than one dimension in the data.
//
// There is a need for a concise way to handle the data produced by devices that sample a physical state at a high frequency.
type SampledData struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The base quantity that a measured value of zero represents. In addition, this provides the units of the entire measurement series.
	Origin Quantity
	// Amount of intervalUnits between samples, e.g. milliseconds for time-based sampling.
	Interval *Decimal
	// The measurement unit in which the sample interval is expressed.
	IntervalUnit Code
	// A correction factor that is applied to the sampled data points before they are added to the origin.
	Factor *Decimal
	// The lower limit of detection of the measured points. This is needed if any of the data points have the value "L" (lower than detection limit).
	LowerLimit *Decimal
	// The upper limit of detection of the measured points. This is needed if any of the data points have the value "U" (higher than detection limit).
	UpperLimit *Decimal
	// The number of sample points at each time point. If this value is greater than one, then the dimensions will be interlaced - all the sample points for a point in time will be recorded at once.
	Dimensions PositiveInt
	// Reference to ConceptMap that defines the codes used in the data.
	CodeMap *Canonical
	// A series of data points which are decimal values separated by a single space (character u20).  The units in which the offsets are expressed are found in intervalUnit.  The absolute point at which the measurements begin SHALL be conveyed outside the scope of this datatype, e.g. Observation.effectiveDateTime for a timing offset.
	Offsets *String
	// A series of data points which are decimal values or codes separated by a single space (character u20). The special codes "E" (error), "L" (below detection limit) and "U" (above detection limit) are also defined for used in place of decimal values.
	Data *String
}

func (r SampledData) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	s += r.Origin.MemSize() - int(unsafe.Sizeof(r.Origin))
	if r.Interval != nil {
		s += r.Interval.MemSize()
	}
	s += r.IntervalUnit.MemSize() - int(unsafe.Sizeof(r.IntervalUnit))
	if r.Factor != nil {
		s += r.Factor.MemSize()
	}
	if r.LowerLimit != nil {
		s += r.LowerLimit.MemSize()
	}
	if r.UpperLimit != nil {
		s += r.UpperLimit.MemSize()
	}
	s += r.Dimensions.MemSize() - int(unsafe.Sizeof(r.Dimensions))
	if r.CodeMap != nil {
		s += r.CodeMap.MemSize()
	}
	if r.Offsets != nil {
		s += r.Offsets.MemSize()
	}
	if r.Data != nil {
		s += r.Data.MemSize()
	}
	return s
}
func (r SampledData) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r SampledData) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r SampledData) marshalJSON(w io.Writer) error {
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
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"origin\":"))
	if err != nil {
		return err
	}
	err = r.Origin.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.Interval != nil && r.Interval.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"interval\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Interval)
		if err != nil {
			return err
		}
	}
	if r.Interval != nil && (r.Interval.Id != nil || r.Interval.Extension != nil) {
		p := primitiveElement{Id: r.Interval.Id, Extension: r.Interval.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_interval\":"))
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
		_, err = w.Write([]byte("\"intervalUnit\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.IntervalUnit)
		if err != nil {
			return err
		}
	}
	if r.IntervalUnit.Id != nil || r.IntervalUnit.Extension != nil {
		p := primitiveElement{Id: r.IntervalUnit.Id, Extension: r.IntervalUnit.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_intervalUnit\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Factor != nil && r.Factor.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"factor\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Factor)
		if err != nil {
			return err
		}
	}
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		p := primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_factor\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.LowerLimit != nil && r.LowerLimit.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"lowerLimit\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.LowerLimit)
		if err != nil {
			return err
		}
	}
	if r.LowerLimit != nil && (r.LowerLimit.Id != nil || r.LowerLimit.Extension != nil) {
		p := primitiveElement{Id: r.LowerLimit.Id, Extension: r.LowerLimit.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_lowerLimit\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.UpperLimit != nil && r.UpperLimit.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"upperLimit\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.UpperLimit)
		if err != nil {
			return err
		}
	}
	if r.UpperLimit != nil && (r.UpperLimit.Id != nil || r.UpperLimit.Extension != nil) {
		p := primitiveElement{Id: r.UpperLimit.Id, Extension: r.UpperLimit.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_upperLimit\":"))
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
		_, err = w.Write([]byte("\"dimensions\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Dimensions)
		if err != nil {
			return err
		}
	}
	if r.Dimensions.Id != nil || r.Dimensions.Extension != nil {
		p := primitiveElement{Id: r.Dimensions.Id, Extension: r.Dimensions.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_dimensions\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.CodeMap != nil && r.CodeMap.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"codeMap\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.CodeMap)
		if err != nil {
			return err
		}
	}
	if r.CodeMap != nil && (r.CodeMap.Id != nil || r.CodeMap.Extension != nil) {
		p := primitiveElement{Id: r.CodeMap.Id, Extension: r.CodeMap.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_codeMap\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Offsets != nil && r.Offsets.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"offsets\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Offsets)
		if err != nil {
			return err
		}
	}
	if r.Offsets != nil && (r.Offsets.Id != nil || r.Offsets.Extension != nil) {
		p := primitiveElement{Id: r.Offsets.Id, Extension: r.Offsets.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_offsets\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Data != nil && r.Data.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"data\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Data)
		if err != nil {
			return err
		}
	}
	if r.Data != nil && (r.Data.Id != nil || r.Data.Extension != nil) {
		p := primitiveElement{Id: r.Data.Id, Extension: r.Data.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_data\":"))
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
func (r *SampledData) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in SampledData element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in SampledData element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in SampledData element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in SampledData element", t)
			}
		case "origin":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Origin = v
		case "interval":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Interval == nil {
				r.Interval = &Decimal{}
			}
			r.Interval.Value = v.Value
		case "_interval":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Interval == nil {
				r.Interval = &Decimal{}
			}
			r.Interval.Id = v.Id
			r.Interval.Extension = v.Extension
		case "intervalUnit":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.IntervalUnit.Value = v.Value
		case "_intervalUnit":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.IntervalUnit.Id = v.Id
			r.IntervalUnit.Extension = v.Extension
		case "factor":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Factor == nil {
				r.Factor = &Decimal{}
			}
			r.Factor.Value = v.Value
		case "_factor":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Factor == nil {
				r.Factor = &Decimal{}
			}
			r.Factor.Id = v.Id
			r.Factor.Extension = v.Extension
		case "lowerLimit":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.LowerLimit == nil {
				r.LowerLimit = &Decimal{}
			}
			r.LowerLimit.Value = v.Value
		case "_lowerLimit":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.LowerLimit == nil {
				r.LowerLimit = &Decimal{}
			}
			r.LowerLimit.Id = v.Id
			r.LowerLimit.Extension = v.Extension
		case "upperLimit":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.UpperLimit == nil {
				r.UpperLimit = &Decimal{}
			}
			r.UpperLimit.Value = v.Value
		case "_upperLimit":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.UpperLimit == nil {
				r.UpperLimit = &Decimal{}
			}
			r.UpperLimit.Id = v.Id
			r.UpperLimit.Extension = v.Extension
		case "dimensions":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Dimensions.Value = v.Value
		case "_dimensions":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Dimensions.Id = v.Id
			r.Dimensions.Extension = v.Extension
		case "codeMap":
			var v Canonical
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.CodeMap == nil {
				r.CodeMap = &Canonical{}
			}
			r.CodeMap.Value = v.Value
		case "_codeMap":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.CodeMap == nil {
				r.CodeMap = &Canonical{}
			}
			r.CodeMap.Id = v.Id
			r.CodeMap.Extension = v.Extension
		case "offsets":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Offsets == nil {
				r.Offsets = &String{}
			}
			r.Offsets.Value = v.Value
		case "_offsets":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Offsets == nil {
				r.Offsets = &String{}
			}
			r.Offsets.Id = v.Id
			r.Offsets.Extension = v.Extension
		case "data":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Data == nil {
				r.Data = &String{}
			}
			r.Data.Value = v.Value
		case "_data":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Data == nil {
				r.Data = &String{}
			}
			r.Data.Id = v.Id
			r.Data.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in SampledData", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in SampledData element", t)
	}
	return nil
}
func (r SampledData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Origin, xml.StartElement{Name: xml.Name{Local: "origin"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Interval, xml.StartElement{Name: xml.Name{Local: "interval"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.IntervalUnit, xml.StartElement{Name: xml.Name{Local: "intervalUnit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Factor, xml.StartElement{Name: xml.Name{Local: "factor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LowerLimit, xml.StartElement{Name: xml.Name{Local: "lowerLimit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UpperLimit, xml.StartElement{Name: xml.Name{Local: "upperLimit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Dimensions, xml.StartElement{Name: xml.Name{Local: "dimensions"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CodeMap, xml.StartElement{Name: xml.Name{Local: "codeMap"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Offsets, xml.StartElement{Name: xml.Name{Local: "offsets"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Data, xml.StartElement{Name: xml.Name{Local: "data"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *SampledData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "origin":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Origin = v
			case "interval":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Interval = &v
			case "intervalUnit":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.IntervalUnit = v
			case "factor":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Factor = &v
			case "lowerLimit":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LowerLimit = &v
			case "upperLimit":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UpperLimit = &v
			case "dimensions":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Dimensions = v
			case "codeMap":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CodeMap = &v
			case "offsets":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Offsets = &v
			case "data":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Data = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
