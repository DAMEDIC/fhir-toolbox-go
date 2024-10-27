package r4

import "encoding/json"

// Base StructureDefinition for HumanName Type: A human's name with the ability to identify parts and usage.
//
// Need to be able to record names, along with notes about their use.
type HumanName struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Identifies the purpose for this name.
	Use *Code
	// Specifies the entire name as it should be displayed e.g. on an application UI. This may be provided instead of or as well as the specific parts.
	Text *String
	// The part of a name that links to the genealogy. In some cultures (e.g. Eritrea) the family name of a son is the first name of his father.
	Family *String
	// Given name.
	Given []String
	// Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the start of the name.
	Prefix []String
	// Part of the name that is acquired as a title due to academic, legal, employment or nobility status, etc. and that appears at the end of the name.
	Suffix []String
	// Indicates the period of time when this name was valid for the named person.
	Period *Period
}
type jsonHumanName struct {
	Id                     *string             `json:"id,omitempty"`
	Extension              []Extension         `json:"extension,omitempty"`
	Use                    *Code               `json:"use,omitempty"`
	UsePrimitiveElement    *primitiveElement   `json:"_use,omitempty"`
	Text                   *String             `json:"text,omitempty"`
	TextPrimitiveElement   *primitiveElement   `json:"_text,omitempty"`
	Family                 *String             `json:"family,omitempty"`
	FamilyPrimitiveElement *primitiveElement   `json:"_family,omitempty"`
	Given                  []String            `json:"given,omitempty"`
	GivenPrimitiveElement  []*primitiveElement `json:"_given,omitempty"`
	Prefix                 []String            `json:"prefix,omitempty"`
	PrefixPrimitiveElement []*primitiveElement `json:"_prefix,omitempty"`
	Suffix                 []String            `json:"suffix,omitempty"`
	SuffixPrimitiveElement []*primitiveElement `json:"_suffix,omitempty"`
	Period                 *Period             `json:"period,omitempty"`
}

func (r HumanName) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r HumanName) marshalJSON() jsonHumanName {
	m := jsonHumanName{}
	m.Id = r.Id
	m.Extension = r.Extension
	if r.Use != nil && r.Use.Value != nil {
		m.Use = r.Use
	}
	if r.Use != nil && (r.Use.Id != nil || r.Use.Extension != nil) {
		m.UsePrimitiveElement = &primitiveElement{Id: r.Use.Id, Extension: r.Use.Extension}
	}
	if r.Text != nil && r.Text.Value != nil {
		m.Text = r.Text
	}
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	if r.Family != nil && r.Family.Value != nil {
		m.Family = r.Family
	}
	if r.Family != nil && (r.Family.Id != nil || r.Family.Extension != nil) {
		m.FamilyPrimitiveElement = &primitiveElement{Id: r.Family.Id, Extension: r.Family.Extension}
	}
	anyGivenValue := false
	for _, e := range r.Given {
		if e.Value != nil {
			anyGivenValue = true
			break
		}
	}
	if anyGivenValue {
		m.Given = r.Given
	}
	anyGivenIdOrExtension := false
	for _, e := range r.Given {
		if e.Id != nil || e.Extension != nil {
			anyGivenIdOrExtension = true
			break
		}
	}
	if anyGivenIdOrExtension {
		m.GivenPrimitiveElement = make([]*primitiveElement, 0, len(r.Given))
		for _, e := range r.Given {
			if e.Id != nil || e.Extension != nil {
				m.GivenPrimitiveElement = append(m.GivenPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.GivenPrimitiveElement = append(m.GivenPrimitiveElement, nil)
			}
		}
	}
	anyPrefixValue := false
	for _, e := range r.Prefix {
		if e.Value != nil {
			anyPrefixValue = true
			break
		}
	}
	if anyPrefixValue {
		m.Prefix = r.Prefix
	}
	anyPrefixIdOrExtension := false
	for _, e := range r.Prefix {
		if e.Id != nil || e.Extension != nil {
			anyPrefixIdOrExtension = true
			break
		}
	}
	if anyPrefixIdOrExtension {
		m.PrefixPrimitiveElement = make([]*primitiveElement, 0, len(r.Prefix))
		for _, e := range r.Prefix {
			if e.Id != nil || e.Extension != nil {
				m.PrefixPrimitiveElement = append(m.PrefixPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.PrefixPrimitiveElement = append(m.PrefixPrimitiveElement, nil)
			}
		}
	}
	anySuffixValue := false
	for _, e := range r.Suffix {
		if e.Value != nil {
			anySuffixValue = true
			break
		}
	}
	if anySuffixValue {
		m.Suffix = r.Suffix
	}
	anySuffixIdOrExtension := false
	for _, e := range r.Suffix {
		if e.Id != nil || e.Extension != nil {
			anySuffixIdOrExtension = true
			break
		}
	}
	if anySuffixIdOrExtension {
		m.SuffixPrimitiveElement = make([]*primitiveElement, 0, len(r.Suffix))
		for _, e := range r.Suffix {
			if e.Id != nil || e.Extension != nil {
				m.SuffixPrimitiveElement = append(m.SuffixPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SuffixPrimitiveElement = append(m.SuffixPrimitiveElement, nil)
			}
		}
	}
	m.Period = r.Period
	return m
}
func (r *HumanName) UnmarshalJSON(b []byte) error {
	var m jsonHumanName
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *HumanName) unmarshalJSON(m jsonHumanName) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Use = m.Use
	if m.UsePrimitiveElement != nil {
		if r.Use == nil {
			r.Use = &Code{}
		}
		r.Use.Id = m.UsePrimitiveElement.Id
		r.Use.Extension = m.UsePrimitiveElement.Extension
	}
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		if r.Text == nil {
			r.Text = &String{}
		}
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	r.Family = m.Family
	if m.FamilyPrimitiveElement != nil {
		if r.Family == nil {
			r.Family = &String{}
		}
		r.Family.Id = m.FamilyPrimitiveElement.Id
		r.Family.Extension = m.FamilyPrimitiveElement.Extension
	}
	r.Given = m.Given
	for i, e := range m.GivenPrimitiveElement {
		if len(r.Given) <= i {
			r.Given = append(r.Given, String{})
		}
		if e != nil {
			r.Given[i].Id = e.Id
			r.Given[i].Extension = e.Extension
		}
	}
	r.Prefix = m.Prefix
	for i, e := range m.PrefixPrimitiveElement {
		if len(r.Prefix) <= i {
			r.Prefix = append(r.Prefix, String{})
		}
		if e != nil {
			r.Prefix[i].Id = e.Id
			r.Prefix[i].Extension = e.Extension
		}
	}
	r.Suffix = m.Suffix
	for i, e := range m.SuffixPrimitiveElement {
		if len(r.Suffix) <= i {
			r.Suffix = append(r.Suffix, String{})
		}
		if e != nil {
			r.Suffix[i].Id = e.Id
			r.Suffix[i].Extension = e.Extension
		}
	}
	r.Period = m.Period
	return nil
}
func (r HumanName) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
