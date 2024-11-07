package r4

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Base StructureDefinition for SampledData Type: A series of measurements taken by a device, with upper and lower limits. There may be more than one dimension in the data.
//
// There is a need for a concise way to handle the data produced by devices that sample a physical state at a high frequency.
type SampledData struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The base quantity that a measured value of zero represents. In addition, this provides the units of the entire measurement series.
	Origin Quantity
	// The length of time between sampling times, measured in milliseconds.
	Period Decimal
	// A correction factor that is applied to the sampled data points before they are added to the origin.
	Factor *Decimal
	// The lower limit of detection of the measured points. This is needed if any of the data points have the value "L" (lower than detection limit).
	LowerLimit *Decimal
	// The upper limit of detection of the measured points. This is needed if any of the data points have the value "U" (higher than detection limit).
	UpperLimit *Decimal
	// The number of sample points at each time point. If this value is greater than one, then the dimensions will be interlaced - all the sample points for a point in time will be recorded at once.
	Dimensions PositiveInt
	// A series of data points which are decimal values separated by a single space (character u20). The special values "E" (error), "L" (below detection limit) and "U" (above detection limit) can also be used in place of a decimal value.
	Data *String
}
type jsonSampledData struct {
	Id                         *string           `json:"id,omitempty"`
	Extension                  []Extension       `json:"extension,omitempty"`
	Origin                     Quantity          `json:"origin,omitempty"`
	Period                     Decimal           `json:"period,omitempty"`
	PeriodPrimitiveElement     *primitiveElement `json:"_period,omitempty"`
	Factor                     *Decimal          `json:"factor,omitempty"`
	FactorPrimitiveElement     *primitiveElement `json:"_factor,omitempty"`
	LowerLimit                 *Decimal          `json:"lowerLimit,omitempty"`
	LowerLimitPrimitiveElement *primitiveElement `json:"_lowerLimit,omitempty"`
	UpperLimit                 *Decimal          `json:"upperLimit,omitempty"`
	UpperLimitPrimitiveElement *primitiveElement `json:"_upperLimit,omitempty"`
	Dimensions                 PositiveInt       `json:"dimensions,omitempty"`
	DimensionsPrimitiveElement *primitiveElement `json:"_dimensions,omitempty"`
	Data                       *String           `json:"data,omitempty"`
	DataPrimitiveElement       *primitiveElement `json:"_data,omitempty"`
}

func (r SampledData) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SampledData) marshalJSON() jsonSampledData {
	m := jsonSampledData{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.Origin = r.Origin
	if r.Period.Value != nil {
		m.Period = r.Period
	}
	if r.Period.Id != nil || r.Period.Extension != nil {
		m.PeriodPrimitiveElement = &primitiveElement{Id: r.Period.Id, Extension: r.Period.Extension}
	}
	if r.Factor != nil && r.Factor.Value != nil {
		m.Factor = r.Factor
	}
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		m.FactorPrimitiveElement = &primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
	}
	if r.LowerLimit != nil && r.LowerLimit.Value != nil {
		m.LowerLimit = r.LowerLimit
	}
	if r.LowerLimit != nil && (r.LowerLimit.Id != nil || r.LowerLimit.Extension != nil) {
		m.LowerLimitPrimitiveElement = &primitiveElement{Id: r.LowerLimit.Id, Extension: r.LowerLimit.Extension}
	}
	if r.UpperLimit != nil && r.UpperLimit.Value != nil {
		m.UpperLimit = r.UpperLimit
	}
	if r.UpperLimit != nil && (r.UpperLimit.Id != nil || r.UpperLimit.Extension != nil) {
		m.UpperLimitPrimitiveElement = &primitiveElement{Id: r.UpperLimit.Id, Extension: r.UpperLimit.Extension}
	}
	if r.Dimensions.Value != nil {
		m.Dimensions = r.Dimensions
	}
	if r.Dimensions.Id != nil || r.Dimensions.Extension != nil {
		m.DimensionsPrimitiveElement = &primitiveElement{Id: r.Dimensions.Id, Extension: r.Dimensions.Extension}
	}
	if r.Data != nil && r.Data.Value != nil {
		m.Data = r.Data
	}
	if r.Data != nil && (r.Data.Id != nil || r.Data.Extension != nil) {
		m.DataPrimitiveElement = &primitiveElement{Id: r.Data.Id, Extension: r.Data.Extension}
	}
	return m
}
func (r *SampledData) UnmarshalJSON(b []byte) error {
	var m jsonSampledData
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SampledData) unmarshalJSON(m jsonSampledData) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Origin = m.Origin
	r.Period = m.Period
	if m.PeriodPrimitiveElement != nil {
		r.Period.Id = m.PeriodPrimitiveElement.Id
		r.Period.Extension = m.PeriodPrimitiveElement.Extension
	}
	r.Factor = m.Factor
	if m.FactorPrimitiveElement != nil {
		if r.Factor == nil {
			r.Factor = &Decimal{}
		}
		r.Factor.Id = m.FactorPrimitiveElement.Id
		r.Factor.Extension = m.FactorPrimitiveElement.Extension
	}
	r.LowerLimit = m.LowerLimit
	if m.LowerLimitPrimitiveElement != nil {
		if r.LowerLimit == nil {
			r.LowerLimit = &Decimal{}
		}
		r.LowerLimit.Id = m.LowerLimitPrimitiveElement.Id
		r.LowerLimit.Extension = m.LowerLimitPrimitiveElement.Extension
	}
	r.UpperLimit = m.UpperLimit
	if m.UpperLimitPrimitiveElement != nil {
		if r.UpperLimit == nil {
			r.UpperLimit = &Decimal{}
		}
		r.UpperLimit.Id = m.UpperLimitPrimitiveElement.Id
		r.UpperLimit.Extension = m.UpperLimitPrimitiveElement.Extension
	}
	r.Dimensions = m.Dimensions
	if m.DimensionsPrimitiveElement != nil {
		r.Dimensions.Id = m.DimensionsPrimitiveElement.Id
		r.Dimensions.Extension = m.DimensionsPrimitiveElement.Extension
	}
	r.Data = m.Data
	if m.DataPrimitiveElement != nil {
		if r.Data == nil {
			r.Data = &String{}
		}
		r.Data.Id = m.DataPrimitiveElement.Id
		r.Data.Extension = m.DataPrimitiveElement.Extension
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
	err = e.EncodeElement(r.Period, xml.StartElement{Name: xml.Name{Local: "period"}})
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
			case "period":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Period = v
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
func (r SampledData) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
