package r5

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"io"
	"reflect"
	"slices"
)

// Availability Type: Availability data for an {item}.
type Availability struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Times the {item} is available.
	AvailableTime []AvailabilityAvailableTime
	// Not available during this time due to provided reason.
	NotAvailableTime []AvailabilityNotAvailableTime
}

// Times the {item} is available.
type AvailabilityAvailableTime struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// mon | tue | wed | thu | fri | sat | sun.
	DaysOfWeek []Code
	// Always available? i.e. 24 hour service.
	AllDay *Boolean
	// Opening time of day (ignored if allDay = true).
	AvailableStartTime *Time
	// Closing time of day (ignored if allDay = true).
	AvailableEndTime *Time
}

// Not available during this time due to provided reason.
type AvailabilityNotAvailableTime struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Reason presented to the user explaining why time not available.
	Description *String
	// Service not available during this period.
	During *Period
}

func (r Availability) MemSize() int {
	s := int(reflect.TypeOf(r).Size())
	if r.Id != nil {
		s += len(*r.Id) + int(reflect.TypeOf(*r.Id).Size())
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(reflect.TypeOf(Extension{}).Size())
	for _, i := range r.AvailableTime {
		s += i.MemSize()
	}
	s += (cap(r.AvailableTime) - len(r.AvailableTime)) * int(reflect.TypeOf(AvailabilityAvailableTime{}).Size())
	for _, i := range r.NotAvailableTime {
		s += i.MemSize()
	}
	s += (cap(r.NotAvailableTime) - len(r.NotAvailableTime)) * int(reflect.TypeOf(AvailabilityNotAvailableTime{}).Size())
	return s
}
func (r AvailabilityAvailableTime) MemSize() int {
	s := int(reflect.TypeOf(r).Size())
	if r.Id != nil {
		s += len(*r.Id) + int(reflect.TypeOf(*r.Id).Size())
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(reflect.TypeOf(Extension{}).Size())
	for _, i := range r.DaysOfWeek {
		s += i.MemSize()
	}
	s += (cap(r.DaysOfWeek) - len(r.DaysOfWeek)) * int(reflect.TypeOf(Code{}).Size())
	if r.AllDay != nil {
		s += r.AllDay.MemSize()
	}
	if r.AvailableStartTime != nil {
		s += r.AvailableStartTime.MemSize()
	}
	if r.AvailableEndTime != nil {
		s += r.AvailableEndTime.MemSize()
	}
	return s
}
func (r AvailabilityNotAvailableTime) MemSize() int {
	s := int(reflect.TypeOf(r).Size())
	if r.Id != nil {
		s += len(*r.Id) + int(reflect.TypeOf(*r.Id).Size())
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(reflect.TypeOf(Extension{}).Size())
	if r.Description != nil {
		s += r.Description.MemSize()
	}
	if r.During != nil {
		s += r.During.MemSize()
	}
	return s
}
func (r Availability) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r AvailabilityAvailableTime) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r AvailabilityNotAvailableTime) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Availability) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Availability) marshalJSON(w io.Writer) error {
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
	if len(r.AvailableTime) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"availableTime\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.AvailableTime {
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
	if len(r.NotAvailableTime) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"notAvailableTime\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.NotAvailableTime {
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
func (r AvailabilityAvailableTime) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r AvailabilityAvailableTime) marshalJSON(w io.Writer) error {
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
	{
		anyValue := false
		for _, e := range r.DaysOfWeek {
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
			_, err = w.Write([]byte("\"daysOfWeek\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.DaysOfWeek)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.DaysOfWeek {
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
			_, err = w.Write([]byte("\"_daysOfWeek\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.DaysOfWeek {
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
	if r.AllDay != nil && r.AllDay.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"allDay\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.AllDay)
		if err != nil {
			return err
		}
	}
	if r.AllDay != nil && (r.AllDay.Id != nil || r.AllDay.Extension != nil) {
		p := primitiveElement{Id: r.AllDay.Id, Extension: r.AllDay.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_allDay\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.AvailableStartTime != nil && r.AvailableStartTime.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"availableStartTime\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.AvailableStartTime)
		if err != nil {
			return err
		}
	}
	if r.AvailableStartTime != nil && (r.AvailableStartTime.Id != nil || r.AvailableStartTime.Extension != nil) {
		p := primitiveElement{Id: r.AvailableStartTime.Id, Extension: r.AvailableStartTime.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_availableStartTime\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.AvailableEndTime != nil && r.AvailableEndTime.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"availableEndTime\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.AvailableEndTime)
		if err != nil {
			return err
		}
	}
	if r.AvailableEndTime != nil && (r.AvailableEndTime.Id != nil || r.AvailableEndTime.Extension != nil) {
		p := primitiveElement{Id: r.AvailableEndTime.Id, Extension: r.AvailableEndTime.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_availableEndTime\":"))
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
func (r AvailabilityNotAvailableTime) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r AvailabilityNotAvailableTime) marshalJSON(w io.Writer) error {
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
	if r.During != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"during\":"))
		if err != nil {
			return err
		}
		err = r.During.marshalJSON(w)
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
func (r *Availability) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in Availability element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in Availability element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in Availability element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in Availability element", t)
			}
		case "availableTime":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Availability element", t)
			}
			for d.More() {
				var v AvailabilityAvailableTime
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.AvailableTime = append(r.AvailableTime, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Availability element", t)
			}
		case "notAvailableTime":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Availability element", t)
			}
			for d.More() {
				var v AvailabilityNotAvailableTime
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.NotAvailableTime = append(r.NotAvailableTime, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Availability element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in Availability", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in Availability element", t)
	}
	return nil
}
func (r *AvailabilityAvailableTime) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in AvailabilityAvailableTime element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in AvailabilityAvailableTime element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in AvailabilityAvailableTime element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in AvailabilityAvailableTime element", t)
			}
		case "daysOfWeek":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AvailabilityAvailableTime element", t)
			}
			for i := 0; d.More(); i++ {
				var v Code
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.DaysOfWeek) <= i {
					r.DaysOfWeek = append(r.DaysOfWeek, Code{})
				}
				r.DaysOfWeek[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AvailabilityAvailableTime element", t)
			}
		case "_daysOfWeek":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in AvailabilityAvailableTime element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.DaysOfWeek) <= i {
					r.DaysOfWeek = append(r.DaysOfWeek, Code{})
				}
				r.DaysOfWeek[i].Id = v.Id
				r.DaysOfWeek[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in AvailabilityAvailableTime element", t)
			}
		case "allDay":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.AllDay == nil {
				r.AllDay = &Boolean{}
			}
			r.AllDay.Value = v.Value
		case "_allDay":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.AllDay == nil {
				r.AllDay = &Boolean{}
			}
			r.AllDay.Id = v.Id
			r.AllDay.Extension = v.Extension
		case "availableStartTime":
			var v Time
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.AvailableStartTime == nil {
				r.AvailableStartTime = &Time{}
			}
			r.AvailableStartTime.Value = v.Value
		case "_availableStartTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.AvailableStartTime == nil {
				r.AvailableStartTime = &Time{}
			}
			r.AvailableStartTime.Id = v.Id
			r.AvailableStartTime.Extension = v.Extension
		case "availableEndTime":
			var v Time
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.AvailableEndTime == nil {
				r.AvailableEndTime = &Time{}
			}
			r.AvailableEndTime.Value = v.Value
		case "_availableEndTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.AvailableEndTime == nil {
				r.AvailableEndTime = &Time{}
			}
			r.AvailableEndTime.Id = v.Id
			r.AvailableEndTime.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in AvailabilityAvailableTime", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in AvailabilityAvailableTime element", t)
	}
	return nil
}
func (r *AvailabilityNotAvailableTime) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in AvailabilityNotAvailableTime element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in AvailabilityNotAvailableTime element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in AvailabilityNotAvailableTime element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in AvailabilityNotAvailableTime element", t)
			}
		case "description":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Value = v.Value
		case "_description":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Description == nil {
				r.Description = &String{}
			}
			r.Description.Id = v.Id
			r.Description.Extension = v.Extension
		case "during":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.During = &v
		default:
			return fmt.Errorf("invalid field: %s in AvailabilityNotAvailableTime", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in AvailabilityNotAvailableTime element", t)
	}
	return nil
}
func (r Availability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.AvailableTime, xml.StartElement{Name: xml.Name{Local: "availableTime"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NotAvailableTime, xml.StartElement{Name: xml.Name{Local: "notAvailableTime"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r AvailabilityAvailableTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.DaysOfWeek, xml.StartElement{Name: xml.Name{Local: "daysOfWeek"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AllDay, xml.StartElement{Name: xml.Name{Local: "allDay"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AvailableStartTime, xml.StartElement{Name: xml.Name{Local: "availableStartTime"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AvailableEndTime, xml.StartElement{Name: xml.Name{Local: "availableEndTime"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r AvailabilityNotAvailableTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.During, xml.StartElement{Name: xml.Name{Local: "during"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Availability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "availableTime":
				var v AvailabilityAvailableTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AvailableTime = append(r.AvailableTime, v)
			case "notAvailableTime":
				var v AvailabilityNotAvailableTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NotAvailableTime = append(r.NotAvailableTime, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *AvailabilityAvailableTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "daysOfWeek":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DaysOfWeek = append(r.DaysOfWeek, v)
			case "allDay":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AllDay = &v
			case "availableStartTime":
				var v Time
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AvailableStartTime = &v
			case "availableEndTime":
				var v Time
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AvailableEndTime = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *AvailabilityNotAvailableTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "during":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.During = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Availability) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "availableTime") {
		for _, v := range r.AvailableTime {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "notAvailableTime") {
		for _, v := range r.NotAvailableTime {
			children = append(children, v)
		}
	}
	return children
}
func (r Availability) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert Availability to Boolean")
}
func (r Availability) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert Availability to String")
}
func (r Availability) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert Availability to Integer")
}
func (r Availability) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert Availability to Decimal")
}
func (r Availability) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert Availability to Date")
}
func (r Availability) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert Availability to Time")
}
func (r Availability) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert Availability to DateTime")
}
func (r Availability) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert Availability to Quantity")
}
func (r Availability) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *Availability
	switch other := other.(type) {
	case Availability:
		o = &other
	case *Availability:
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
func (r Availability) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(Availability)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r Availability) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
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
			Name: "AvailableTime",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "AvailabilityAvailableTime",
				Namespace: "FHIR",
			},
		}, {
			Name: "NotAvailableTime",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "AvailabilityNotAvailableTime",
				Namespace: "FHIR",
			},
		}},
		Name:      "Availability",
		Namespace: "FHIR",
	}
}
func (r AvailabilityAvailableTime) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "daysOfWeek") {
		for _, v := range r.DaysOfWeek {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "allDay") {
		if r.AllDay != nil {
			children = append(children, *r.AllDay)
		}
	}
	if len(name) == 0 || slices.Contains(name, "availableStartTime") {
		if r.AvailableStartTime != nil {
			children = append(children, *r.AvailableStartTime)
		}
	}
	if len(name) == 0 || slices.Contains(name, "availableEndTime") {
		if r.AvailableEndTime != nil {
			children = append(children, *r.AvailableEndTime)
		}
	}
	return children
}
func (r AvailabilityAvailableTime) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert AvailabilityAvailableTime to Boolean")
}
func (r AvailabilityAvailableTime) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert AvailabilityAvailableTime to String")
}
func (r AvailabilityAvailableTime) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert AvailabilityAvailableTime to Integer")
}
func (r AvailabilityAvailableTime) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert AvailabilityAvailableTime to Decimal")
}
func (r AvailabilityAvailableTime) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert AvailabilityAvailableTime to Date")
}
func (r AvailabilityAvailableTime) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert AvailabilityAvailableTime to Time")
}
func (r AvailabilityAvailableTime) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert AvailabilityAvailableTime to DateTime")
}
func (r AvailabilityAvailableTime) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert AvailabilityAvailableTime to Quantity")
}
func (r AvailabilityAvailableTime) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *AvailabilityAvailableTime
	switch other := other.(type) {
	case AvailabilityAvailableTime:
		o = &other
	case *AvailabilityAvailableTime:
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
func (r AvailabilityAvailableTime) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(AvailabilityAvailableTime)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r AvailabilityAvailableTime) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
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
			Name: "DaysOfWeek",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "AllDay",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "AvailableStartTime",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Time",
				Namespace: "FHIR",
			},
		}, {
			Name: "AvailableEndTime",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Time",
				Namespace: "FHIR",
			},
		}},
		Name:      "AvailabilityAvailableTime",
		Namespace: "FHIR",
	}
}
func (r AvailabilityNotAvailableTime) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "description") {
		if r.Description != nil {
			children = append(children, *r.Description)
		}
	}
	if len(name) == 0 || slices.Contains(name, "during") {
		if r.During != nil {
			children = append(children, *r.During)
		}
	}
	return children
}
func (r AvailabilityNotAvailableTime) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert AvailabilityNotAvailableTime to Boolean")
}
func (r AvailabilityNotAvailableTime) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert AvailabilityNotAvailableTime to String")
}
func (r AvailabilityNotAvailableTime) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert AvailabilityNotAvailableTime to Integer")
}
func (r AvailabilityNotAvailableTime) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert AvailabilityNotAvailableTime to Decimal")
}
func (r AvailabilityNotAvailableTime) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert AvailabilityNotAvailableTime to Date")
}
func (r AvailabilityNotAvailableTime) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert AvailabilityNotAvailableTime to Time")
}
func (r AvailabilityNotAvailableTime) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert AvailabilityNotAvailableTime to DateTime")
}
func (r AvailabilityNotAvailableTime) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert AvailabilityNotAvailableTime to Quantity")
}
func (r AvailabilityNotAvailableTime) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *AvailabilityNotAvailableTime
	switch other := other.(type) {
	case AvailabilityNotAvailableTime:
		o = &other
	case *AvailabilityNotAvailableTime:
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
func (r AvailabilityNotAvailableTime) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(AvailabilityNotAvailableTime)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r AvailabilityNotAvailableTime) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DataType",
			Namespace: "FHIR",
		},
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
			Name: "Description",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "During",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Period",
				Namespace: "FHIR",
			},
		}},
		Name:      "AvailabilityNotAvailableTime",
		Namespace: "FHIR",
	}
}
