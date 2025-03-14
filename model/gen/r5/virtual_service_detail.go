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

// VirtualServiceDetail Type: Virtual Service Contact Details.
type VirtualServiceDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The type of virtual service to connect to (i.e. Teams, Zoom, Specific VMR technology, WhatsApp).
	ChannelType *Coding
	// What address or number needs to be used for a user to connect to the virtual service to join. The channelType informs as to which datatype is appropriate to use (requires knowledge of the specific type).
	Address isVirtualServiceDetailAddress
	// Address to see alternative connection details.
	AdditionalInfo []Url
	// Maximum number of participants supported by the virtual service.
	MaxParticipants *PositiveInt
	// Session Key required by the virtual service.
	SessionKey *String
}
type isVirtualServiceDetailAddress interface {
	model.Element
	isVirtualServiceDetailAddress()
}

func (r Url) isVirtualServiceDetailAddress()                   {}
func (r String) isVirtualServiceDetailAddress()                {}
func (r ContactPoint) isVirtualServiceDetailAddress()          {}
func (r ExtendedContactDetail) isVirtualServiceDetailAddress() {}
func (r VirtualServiceDetail) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	if r.ChannelType != nil {
		s += r.ChannelType.MemSize()
	}
	if r.Address != nil {
		s += r.Address.MemSize()
	}
	for _, i := range r.AdditionalInfo {
		s += i.MemSize()
	}
	s += (cap(r.AdditionalInfo) - len(r.AdditionalInfo)) * int(unsafe.Sizeof(Url{}))
	if r.MaxParticipants != nil {
		s += r.MaxParticipants.MemSize()
	}
	if r.SessionKey != nil {
		s += r.SessionKey.MemSize()
	}
	return s
}
func (r VirtualServiceDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r VirtualServiceDetail) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r VirtualServiceDetail) marshalJSON(w io.Writer) error {
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
	if r.ChannelType != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"channelType\":"))
		if err != nil {
			return err
		}
		err = r.ChannelType.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	switch v := r.Address.(type) {
	case Url:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"addressUrl\":"))
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
			_, err = w.Write([]byte("\"_addressUrl\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Url:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"addressUrl\":"))
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
			_, err = w.Write([]byte("\"_addressUrl\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
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
			_, err = w.Write([]byte("\"addressString\":"))
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
			_, err = w.Write([]byte("\"_addressString\":"))
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
			_, err = w.Write([]byte("\"addressString\":"))
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
			_, err = w.Write([]byte("\"_addressString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case ContactPoint:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"addressContactPoint\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *ContactPoint:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"addressContactPoint\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case ExtendedContactDetail:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"addressExtendedContactDetail\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *ExtendedContactDetail:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"addressExtendedContactDetail\":"))
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
	{
		anyValue := false
		for _, e := range r.AdditionalInfo {
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
			_, err = w.Write([]byte("\"additionalInfo\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.AdditionalInfo)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.AdditionalInfo {
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
			_, err = w.Write([]byte("\"_additionalInfo\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.AdditionalInfo {
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
	if r.MaxParticipants != nil && r.MaxParticipants.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"maxParticipants\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.MaxParticipants)
		if err != nil {
			return err
		}
	}
	if r.MaxParticipants != nil && (r.MaxParticipants.Id != nil || r.MaxParticipants.Extension != nil) {
		p := primitiveElement{Id: r.MaxParticipants.Id, Extension: r.MaxParticipants.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_maxParticipants\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.SessionKey != nil && r.SessionKey.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"sessionKey\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.SessionKey)
		if err != nil {
			return err
		}
	}
	if r.SessionKey != nil && (r.SessionKey.Id != nil || r.SessionKey.Extension != nil) {
		p := primitiveElement{Id: r.SessionKey.Id, Extension: r.SessionKey.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_sessionKey\":"))
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
func (r *VirtualServiceDetail) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in VirtualServiceDetail element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in VirtualServiceDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in VirtualServiceDetail element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in VirtualServiceDetail element", t)
			}
		case "channelType":
			var v Coding
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ChannelType = &v
		case "addressUrl":
			var v Url
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Address != nil {
				r.Address = Url{
					Extension: r.Address.(Url).Extension,
					Id:        r.Address.(Url).Id,
					Value:     v.Value,
				}
			} else {
				r.Address = v
			}
		case "_addressUrl":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Address != nil {
				r.Address = Url{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Address.(Url).Value,
				}
			} else {
				r.Address = Url{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "addressString":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Address != nil {
				r.Address = String{
					Extension: r.Address.(String).Extension,
					Id:        r.Address.(String).Id,
					Value:     v.Value,
				}
			} else {
				r.Address = v
			}
		case "_addressString":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Address != nil {
				r.Address = String{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Address.(String).Value,
				}
			} else {
				r.Address = String{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "addressContactPoint":
			var v ContactPoint
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Address = v
		case "addressExtendedContactDetail":
			var v ExtendedContactDetail
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Address = v
		case "additionalInfo":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VirtualServiceDetail element", t)
			}
			for i := 0; d.More(); i++ {
				var v Url
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.AdditionalInfo) <= i {
					r.AdditionalInfo = append(r.AdditionalInfo, Url{})
				}
				r.AdditionalInfo[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VirtualServiceDetail element", t)
			}
		case "_additionalInfo":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in VirtualServiceDetail element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.AdditionalInfo) <= i {
					r.AdditionalInfo = append(r.AdditionalInfo, Url{})
				}
				r.AdditionalInfo[i].Id = v.Id
				r.AdditionalInfo[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in VirtualServiceDetail element", t)
			}
		case "maxParticipants":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.MaxParticipants == nil {
				r.MaxParticipants = &PositiveInt{}
			}
			r.MaxParticipants.Value = v.Value
		case "_maxParticipants":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.MaxParticipants == nil {
				r.MaxParticipants = &PositiveInt{}
			}
			r.MaxParticipants.Id = v.Id
			r.MaxParticipants.Extension = v.Extension
		case "sessionKey":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.SessionKey == nil {
				r.SessionKey = &String{}
			}
			r.SessionKey.Value = v.Value
		case "_sessionKey":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.SessionKey == nil {
				r.SessionKey = &String{}
			}
			r.SessionKey.Id = v.Id
			r.SessionKey.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in VirtualServiceDetail", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in VirtualServiceDetail element", t)
	}
	return nil
}
func (r VirtualServiceDetail) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ChannelType, xml.StartElement{Name: xml.Name{Local: "channelType"}})
	if err != nil {
		return err
	}
	switch v := r.Address.(type) {
	case Url, *Url:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "addressUrl"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "addressString"}})
		if err != nil {
			return err
		}
	case ContactPoint, *ContactPoint:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "addressContactPoint"}})
		if err != nil {
			return err
		}
	case ExtendedContactDetail, *ExtendedContactDetail:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "addressExtendedContactDetail"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.AdditionalInfo, xml.StartElement{Name: xml.Name{Local: "additionalInfo"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MaxParticipants, xml.StartElement{Name: xml.Name{Local: "maxParticipants"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SessionKey, xml.StartElement{Name: xml.Name{Local: "sessionKey"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *VirtualServiceDetail) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "channelType":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ChannelType = &v
			case "addressUrl":
				if r.Address != nil {
					return fmt.Errorf("multiple values for field \"Address\"")
				}
				var v Url
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Address = v
			case "addressString":
				if r.Address != nil {
					return fmt.Errorf("multiple values for field \"Address\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Address = v
			case "addressContactPoint":
				if r.Address != nil {
					return fmt.Errorf("multiple values for field \"Address\"")
				}
				var v ContactPoint
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Address = v
			case "addressExtendedContactDetail":
				if r.Address != nil {
					return fmt.Errorf("multiple values for field \"Address\"")
				}
				var v ExtendedContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Address = v
			case "additionalInfo":
				var v Url
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AdditionalInfo = append(r.AdditionalInfo, v)
			case "maxParticipants":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MaxParticipants = &v
			case "sessionKey":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SessionKey = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r VirtualServiceDetail) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "channelType") {
		if r.ChannelType != nil {
			children = append(children, *r.ChannelType)
		}
	}
	if len(name) == 0 || slices.Contains(name, "address") {
		if r.Address != nil {
			children = append(children, r.Address)
		}
	}
	if len(name) == 0 || slices.Contains(name, "additionalInfo") {
		for _, v := range r.AdditionalInfo {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "maxParticipants") {
		if r.MaxParticipants != nil {
			children = append(children, *r.MaxParticipants)
		}
	}
	if len(name) == 0 || slices.Contains(name, "sessionKey") {
		if r.SessionKey != nil {
			children = append(children, *r.SessionKey)
		}
	}
	return children
}
func (r VirtualServiceDetail) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert VirtualServiceDetail to Boolean")
}
func (r VirtualServiceDetail) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert VirtualServiceDetail to String")
}
func (r VirtualServiceDetail) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert VirtualServiceDetail to Integer")
}
func (r VirtualServiceDetail) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert VirtualServiceDetail to Decimal")
}
func (r VirtualServiceDetail) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert VirtualServiceDetail to Date")
}
func (r VirtualServiceDetail) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert VirtualServiceDetail to Time")
}
func (r VirtualServiceDetail) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert VirtualServiceDetail to DateTime")
}
func (r VirtualServiceDetail) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert VirtualServiceDetail to Quantity")
}
func (r VirtualServiceDetail) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o VirtualServiceDetail
	switch other := other.(type) {
	case VirtualServiceDetail:
		o = other
	case *VirtualServiceDetail:
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
func (r VirtualServiceDetail) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o VirtualServiceDetail
	switch other := other.(type) {
	case VirtualServiceDetail:
		o = other
	case *VirtualServiceDetail:
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
func (r VirtualServiceDetail) TypeInfo() fhirpath.TypeInfo {
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
			Name: "ChannelType",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Coding",
				Namespace: "FHIR",
			},
		}, {
			Name: "Address",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PrimitiveElement",
				Namespace: "FHIR",
			},
		}, {
			Name: "AdditionalInfo",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Url",
				Namespace: "FHIR",
			},
		}, {
			Name: "MaxParticipants",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "SessionKey",
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
			Name:      "VirtualServiceDetail",
			Namespace: "FHIR",
		},
	}
}
