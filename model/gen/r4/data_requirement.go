package r4

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Base StructureDefinition for DataRequirement Type: Describes a required data item for evaluation in terms of the type of data, and optional code or date-based filters of the data.
type DataRequirement struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The type of the required data, specified as the type name of a resource. For profiles, this value is set to the type of the base resource of the profile.
	Type Code
	// The profile of the required data, specified as the uri of the profile definition.
	Profile []Canonical
	// The intended subjects of the data requirement. If this element is not provided, a Patient subject is assumed.
	Subject isDataRequirementSubject
	// Indicates that specific elements of the type are referenced by the knowledge module and must be supported by the consumer in order to obtain an effective evaluation. This does not mean that a value is required for this element, only that the consuming system must understand the element and be able to provide values for it if they are available.
	//
	// The value of mustSupport SHALL be a FHIRPath resolveable on the type of the DataRequirement. The path SHALL consist only of identifiers, constant indexers, and .resolve() (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details).
	MustSupport []String
	// Code filters specify additional constraints on the data, specifying the value set of interest for a particular element of the data. Each code filter defines an additional constraint on the data, i.e. code filters are AND'ed, not OR'ed.
	CodeFilter []DataRequirementCodeFilter
	// Date filters specify additional constraints on the data in terms of the applicable date range for specific elements. Each date filter specifies an additional constraint on the data, i.e. date filters are AND'ed, not OR'ed.
	DateFilter []DataRequirementDateFilter
	// Specifies a maximum number of results that are required (uses the _count search parameter).
	Limit *PositiveInt
	// Specifies the order of the results to be returned.
	Sort []DataRequirementSort
}
type isDataRequirementSubject interface {
	isDataRequirementSubject()
}

func (r CodeableConcept) isDataRequirementSubject() {}
func (r Reference) isDataRequirementSubject()       {}

type jsonDataRequirement struct {
	Id                          *string                     `json:"id,omitempty"`
	Extension                   []Extension                 `json:"extension,omitempty"`
	Type                        Code                        `json:"type,omitempty"`
	TypePrimitiveElement        *primitiveElement           `json:"_type,omitempty"`
	Profile                     []Canonical                 `json:"profile,omitempty"`
	ProfilePrimitiveElement     []*primitiveElement         `json:"_profile,omitempty"`
	SubjectCodeableConcept      *CodeableConcept            `json:"subjectCodeableConcept,omitempty"`
	SubjectReference            *Reference                  `json:"subjectReference,omitempty"`
	MustSupport                 []String                    `json:"mustSupport,omitempty"`
	MustSupportPrimitiveElement []*primitiveElement         `json:"_mustSupport,omitempty"`
	CodeFilter                  []DataRequirementCodeFilter `json:"codeFilter,omitempty"`
	DateFilter                  []DataRequirementDateFilter `json:"dateFilter,omitempty"`
	Limit                       *PositiveInt                `json:"limit,omitempty"`
	LimitPrimitiveElement       *primitiveElement           `json:"_limit,omitempty"`
	Sort                        []DataRequirementSort       `json:"sort,omitempty"`
}

func (r DataRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DataRequirement) marshalJSON() jsonDataRequirement {
	m := jsonDataRequirement{}
	m.Id = r.Id
	m.Extension = r.Extension
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	anyProfileValue := false
	for _, e := range r.Profile {
		if e.Value != nil {
			anyProfileValue = true
			break
		}
	}
	if anyProfileValue {
		m.Profile = r.Profile
	}
	anyProfileIdOrExtension := false
	for _, e := range r.Profile {
		if e.Id != nil || e.Extension != nil {
			anyProfileIdOrExtension = true
			break
		}
	}
	if anyProfileIdOrExtension {
		m.ProfilePrimitiveElement = make([]*primitiveElement, 0, len(r.Profile))
		for _, e := range r.Profile {
			if e.Id != nil || e.Extension != nil {
				m.ProfilePrimitiveElement = append(m.ProfilePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ProfilePrimitiveElement = append(m.ProfilePrimitiveElement, nil)
			}
		}
	}
	switch v := r.Subject.(type) {
	case CodeableConcept:
		m.SubjectCodeableConcept = &v
	case *CodeableConcept:
		m.SubjectCodeableConcept = v
	case Reference:
		m.SubjectReference = &v
	case *Reference:
		m.SubjectReference = v
	}
	anyMustSupportValue := false
	for _, e := range r.MustSupport {
		if e.Value != nil {
			anyMustSupportValue = true
			break
		}
	}
	if anyMustSupportValue {
		m.MustSupport = r.MustSupport
	}
	anyMustSupportIdOrExtension := false
	for _, e := range r.MustSupport {
		if e.Id != nil || e.Extension != nil {
			anyMustSupportIdOrExtension = true
			break
		}
	}
	if anyMustSupportIdOrExtension {
		m.MustSupportPrimitiveElement = make([]*primitiveElement, 0, len(r.MustSupport))
		for _, e := range r.MustSupport {
			if e.Id != nil || e.Extension != nil {
				m.MustSupportPrimitiveElement = append(m.MustSupportPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.MustSupportPrimitiveElement = append(m.MustSupportPrimitiveElement, nil)
			}
		}
	}
	m.CodeFilter = r.CodeFilter
	m.DateFilter = r.DateFilter
	if r.Limit != nil && r.Limit.Value != nil {
		m.Limit = r.Limit
	}
	if r.Limit != nil && (r.Limit.Id != nil || r.Limit.Extension != nil) {
		m.LimitPrimitiveElement = &primitiveElement{Id: r.Limit.Id, Extension: r.Limit.Extension}
	}
	m.Sort = r.Sort
	return m
}
func (r *DataRequirement) UnmarshalJSON(b []byte) error {
	var m jsonDataRequirement
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DataRequirement) unmarshalJSON(m jsonDataRequirement) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Profile = m.Profile
	for i, e := range m.ProfilePrimitiveElement {
		if len(r.Profile) <= i {
			r.Profile = append(r.Profile, Canonical{})
		}
		if e != nil {
			r.Profile[i].Id = e.Id
			r.Profile[i].Extension = e.Extension
		}
	}
	if m.SubjectCodeableConcept != nil {
		if r.Subject != nil {
			return fmt.Errorf("multiple values for field \"Subject\"")
		}
		v := m.SubjectCodeableConcept
		r.Subject = v
	}
	if m.SubjectReference != nil {
		if r.Subject != nil {
			return fmt.Errorf("multiple values for field \"Subject\"")
		}
		v := m.SubjectReference
		r.Subject = v
	}
	r.MustSupport = m.MustSupport
	for i, e := range m.MustSupportPrimitiveElement {
		if len(r.MustSupport) <= i {
			r.MustSupport = append(r.MustSupport, String{})
		}
		if e != nil {
			r.MustSupport[i].Id = e.Id
			r.MustSupport[i].Extension = e.Extension
		}
	}
	r.CodeFilter = m.CodeFilter
	r.DateFilter = m.DateFilter
	r.Limit = m.Limit
	if m.LimitPrimitiveElement != nil {
		if r.Limit == nil {
			r.Limit = &PositiveInt{}
		}
		r.Limit.Id = m.LimitPrimitiveElement.Id
		r.Limit.Extension = m.LimitPrimitiveElement.Extension
	}
	r.Sort = m.Sort
	return nil
}
func (r DataRequirement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Profile, xml.StartElement{Name: xml.Name{Local: "profile"}})
	if err != nil {
		return err
	}
	switch v := r.Subject.(type) {
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "subjectCodeableConcept"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "subjectReference"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.MustSupport, xml.StartElement{Name: xml.Name{Local: "mustSupport"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CodeFilter, xml.StartElement{Name: xml.Name{Local: "codeFilter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DateFilter, xml.StartElement{Name: xml.Name{Local: "dateFilter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Limit, xml.StartElement{Name: xml.Name{Local: "limit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Sort, xml.StartElement{Name: xml.Name{Local: "sort"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *DataRequirement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "profile":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Profile = append(r.Profile, v)
			case "subjectCodeableConcept":
				if r.Subject != nil {
					return fmt.Errorf("multiple values for field \"Subject\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = v
			case "subjectReference":
				if r.Subject != nil {
					return fmt.Errorf("multiple values for field \"Subject\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = v
			case "mustSupport":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MustSupport = append(r.MustSupport, v)
			case "codeFilter":
				var v DataRequirementCodeFilter
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CodeFilter = append(r.CodeFilter, v)
			case "dateFilter":
				var v DataRequirementDateFilter
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DateFilter = append(r.DateFilter, v)
			case "limit":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Limit = &v
			case "sort":
				var v DataRequirementSort
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Sort = append(r.Sort, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r DataRequirement) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Code filters specify additional constraints on the data, specifying the value set of interest for a particular element of the data. Each code filter defines an additional constraint on the data, i.e. code filters are AND'ed, not OR'ed.
type DataRequirementCodeFilter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The code-valued attribute of the filter. The specified path SHALL be a FHIRPath resolveable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of type code, Coding, or CodeableConcept.
	Path *String
	// A token parameter that refers to a search parameter defined on the specified type of the DataRequirement, and which searches on elements of type code, Coding, or CodeableConcept.
	SearchParam *String
	// The valueset for the code filter. The valueSet and code elements are additive. If valueSet is specified, the filter will return only those data items for which the value of the code-valued element specified in the path is a member of the specified valueset.
	ValueSet *Canonical
	// The codes for the code filter. If values are given, the filter will return only those data items for which the code-valued attribute specified by the path has a value that is one of the specified codes. If codes are specified in addition to a value set, the filter returns items matching a code in the value set or one of the specified codes.
	Code []Coding
}
type jsonDataRequirementCodeFilter struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	Path                        *String           `json:"path,omitempty"`
	PathPrimitiveElement        *primitiveElement `json:"_path,omitempty"`
	SearchParam                 *String           `json:"searchParam,omitempty"`
	SearchParamPrimitiveElement *primitiveElement `json:"_searchParam,omitempty"`
	ValueSet                    *Canonical        `json:"valueSet,omitempty"`
	ValueSetPrimitiveElement    *primitiveElement `json:"_valueSet,omitempty"`
	Code                        []Coding          `json:"code,omitempty"`
}

func (r DataRequirementCodeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DataRequirementCodeFilter) marshalJSON() jsonDataRequirementCodeFilter {
	m := jsonDataRequirementCodeFilter{}
	m.Id = r.Id
	m.Extension = r.Extension
	if r.Path != nil && r.Path.Value != nil {
		m.Path = r.Path
	}
	if r.Path != nil && (r.Path.Id != nil || r.Path.Extension != nil) {
		m.PathPrimitiveElement = &primitiveElement{Id: r.Path.Id, Extension: r.Path.Extension}
	}
	if r.SearchParam != nil && r.SearchParam.Value != nil {
		m.SearchParam = r.SearchParam
	}
	if r.SearchParam != nil && (r.SearchParam.Id != nil || r.SearchParam.Extension != nil) {
		m.SearchParamPrimitiveElement = &primitiveElement{Id: r.SearchParam.Id, Extension: r.SearchParam.Extension}
	}
	if r.ValueSet != nil && r.ValueSet.Value != nil {
		m.ValueSet = r.ValueSet
	}
	if r.ValueSet != nil && (r.ValueSet.Id != nil || r.ValueSet.Extension != nil) {
		m.ValueSetPrimitiveElement = &primitiveElement{Id: r.ValueSet.Id, Extension: r.ValueSet.Extension}
	}
	m.Code = r.Code
	return m
}
func (r *DataRequirementCodeFilter) UnmarshalJSON(b []byte) error {
	var m jsonDataRequirementCodeFilter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DataRequirementCodeFilter) unmarshalJSON(m jsonDataRequirementCodeFilter) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Path = m.Path
	if m.PathPrimitiveElement != nil {
		if r.Path == nil {
			r.Path = &String{}
		}
		r.Path.Id = m.PathPrimitiveElement.Id
		r.Path.Extension = m.PathPrimitiveElement.Extension
	}
	r.SearchParam = m.SearchParam
	if m.SearchParamPrimitiveElement != nil {
		if r.SearchParam == nil {
			r.SearchParam = &String{}
		}
		r.SearchParam.Id = m.SearchParamPrimitiveElement.Id
		r.SearchParam.Extension = m.SearchParamPrimitiveElement.Extension
	}
	r.ValueSet = m.ValueSet
	if m.ValueSetPrimitiveElement != nil {
		if r.ValueSet == nil {
			r.ValueSet = &Canonical{}
		}
		r.ValueSet.Id = m.ValueSetPrimitiveElement.Id
		r.ValueSet.Extension = m.ValueSetPrimitiveElement.Extension
	}
	r.Code = m.Code
	return nil
}
func (r DataRequirementCodeFilter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Path, xml.StartElement{Name: xml.Name{Local: "path"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SearchParam, xml.StartElement{Name: xml.Name{Local: "searchParam"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ValueSet, xml.StartElement{Name: xml.Name{Local: "valueSet"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *DataRequirementCodeFilter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "path":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Path = &v
			case "searchParam":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SearchParam = &v
			case "valueSet":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ValueSet = &v
			case "code":
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = append(r.Code, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r DataRequirementCodeFilter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Date filters specify additional constraints on the data in terms of the applicable date range for specific elements. Each date filter specifies an additional constraint on the data, i.e. date filters are AND'ed, not OR'ed.
type DataRequirementDateFilter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The date-valued attribute of the filter. The specified path SHALL be a FHIRPath resolveable on the specified type of the DataRequirement, and SHALL consist only of identifiers, constant indexers, and .resolve(). The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements (see the [Simple FHIRPath Profile](fhirpath.html#simple) for full details). Note that the index must be an integer constant. The path must resolve to an element of type date, dateTime, Period, Schedule, or Timing.
	Path *String
	// A date parameter that refers to a search parameter defined on the specified type of the DataRequirement, and which searches on elements of type date, dateTime, Period, Schedule, or Timing.
	SearchParam *String
	// The value of the filter. If period is specified, the filter will return only those data items that fall within the bounds determined by the Period, inclusive of the period boundaries. If dateTime is specified, the filter will return only those data items that are equal to the specified dateTime. If a Duration is specified, the filter will return only those data items that fall within Duration before now.
	Value isDataRequirementDateFilterValue
}
type isDataRequirementDateFilterValue interface {
	isDataRequirementDateFilterValue()
}

func (r DateTime) isDataRequirementDateFilterValue() {}
func (r Period) isDataRequirementDateFilterValue()   {}
func (r Duration) isDataRequirementDateFilterValue() {}

type jsonDataRequirementDateFilter struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	Path                          *String           `json:"path,omitempty"`
	PathPrimitiveElement          *primitiveElement `json:"_path,omitempty"`
	SearchParam                   *String           `json:"searchParam,omitempty"`
	SearchParamPrimitiveElement   *primitiveElement `json:"_searchParam,omitempty"`
	ValueDateTime                 *DateTime         `json:"valueDateTime,omitempty"`
	ValueDateTimePrimitiveElement *primitiveElement `json:"_valueDateTime,omitempty"`
	ValuePeriod                   *Period           `json:"valuePeriod,omitempty"`
	ValueDuration                 *Duration         `json:"valueDuration,omitempty"`
}

func (r DataRequirementDateFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DataRequirementDateFilter) marshalJSON() jsonDataRequirementDateFilter {
	m := jsonDataRequirementDateFilter{}
	m.Id = r.Id
	m.Extension = r.Extension
	if r.Path != nil && r.Path.Value != nil {
		m.Path = r.Path
	}
	if r.Path != nil && (r.Path.Id != nil || r.Path.Extension != nil) {
		m.PathPrimitiveElement = &primitiveElement{Id: r.Path.Id, Extension: r.Path.Extension}
	}
	if r.SearchParam != nil && r.SearchParam.Value != nil {
		m.SearchParam = r.SearchParam
	}
	if r.SearchParam != nil && (r.SearchParam.Id != nil || r.SearchParam.Extension != nil) {
		m.SearchParamPrimitiveElement = &primitiveElement{Id: r.SearchParam.Id, Extension: r.SearchParam.Extension}
	}
	switch v := r.Value.(type) {
	case DateTime:
		if v.Value != nil {
			m.ValueDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.ValueDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.ValuePeriod = &v
	case *Period:
		m.ValuePeriod = v
	case Duration:
		m.ValueDuration = &v
	case *Duration:
		m.ValueDuration = v
	}
	return m
}
func (r *DataRequirementDateFilter) UnmarshalJSON(b []byte) error {
	var m jsonDataRequirementDateFilter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DataRequirementDateFilter) unmarshalJSON(m jsonDataRequirementDateFilter) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Path = m.Path
	if m.PathPrimitiveElement != nil {
		if r.Path == nil {
			r.Path = &String{}
		}
		r.Path.Id = m.PathPrimitiveElement.Id
		r.Path.Extension = m.PathPrimitiveElement.Extension
	}
	r.SearchParam = m.SearchParam
	if m.SearchParamPrimitiveElement != nil {
		if r.SearchParam == nil {
			r.SearchParam = &String{}
		}
		r.SearchParam.Id = m.SearchParamPrimitiveElement.Id
		r.SearchParam.Extension = m.SearchParamPrimitiveElement.Extension
	}
	if m.ValueDateTime != nil || m.ValueDateTimePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDateTime
		if m.ValueDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.ValueDateTimePrimitiveElement.Id
			v.Extension = m.ValueDateTimePrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValuePeriod != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValuePeriod
		r.Value = v
	}
	if m.ValueDuration != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDuration
		r.Value = v
	}
	return nil
}
func (r DataRequirementDateFilter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Path, xml.StartElement{Name: xml.Name{Local: "path"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SearchParam, xml.StartElement{Name: xml.Name{Local: "searchParam"}})
	if err != nil {
		return err
	}
	switch v := r.Value.(type) {
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDateTime"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valuePeriod"}})
		if err != nil {
			return err
		}
	case Duration, *Duration:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDuration"}})
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
func (r *DataRequirementDateFilter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "path":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Path = &v
			case "searchParam":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SearchParam = &v
			case "valueDateTime":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valuePeriod":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDuration":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Duration
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r DataRequirementDateFilter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Specifies the order of the results to be returned.
type DataRequirementSort struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The attribute of the sort. The specified path must be resolvable from the type of the required data. The path is allowed to contain qualifiers (.) to traverse sub-elements, as well as indexers ([x]) to traverse multiple-cardinality sub-elements. Note that the index must be an integer constant.
	Path String
	// The direction of the sort, ascending or descending.
	Direction Code
}
type jsonDataRequirementSort struct {
	Id                        *string           `json:"id,omitempty"`
	Extension                 []Extension       `json:"extension,omitempty"`
	Path                      String            `json:"path,omitempty"`
	PathPrimitiveElement      *primitiveElement `json:"_path,omitempty"`
	Direction                 Code              `json:"direction,omitempty"`
	DirectionPrimitiveElement *primitiveElement `json:"_direction,omitempty"`
}

func (r DataRequirementSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DataRequirementSort) marshalJSON() jsonDataRequirementSort {
	m := jsonDataRequirementSort{}
	m.Id = r.Id
	m.Extension = r.Extension
	if r.Path.Value != nil {
		m.Path = r.Path
	}
	if r.Path.Id != nil || r.Path.Extension != nil {
		m.PathPrimitiveElement = &primitiveElement{Id: r.Path.Id, Extension: r.Path.Extension}
	}
	if r.Direction.Value != nil {
		m.Direction = r.Direction
	}
	if r.Direction.Id != nil || r.Direction.Extension != nil {
		m.DirectionPrimitiveElement = &primitiveElement{Id: r.Direction.Id, Extension: r.Direction.Extension}
	}
	return m
}
func (r *DataRequirementSort) UnmarshalJSON(b []byte) error {
	var m jsonDataRequirementSort
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DataRequirementSort) unmarshalJSON(m jsonDataRequirementSort) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.Path = m.Path
	if m.PathPrimitiveElement != nil {
		r.Path.Id = m.PathPrimitiveElement.Id
		r.Path.Extension = m.PathPrimitiveElement.Extension
	}
	r.Direction = m.Direction
	if m.DirectionPrimitiveElement != nil {
		r.Direction.Id = m.DirectionPrimitiveElement.Id
		r.Direction.Extension = m.DirectionPrimitiveElement.Extension
	}
	return nil
}
func (r DataRequirementSort) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Path, xml.StartElement{Name: xml.Name{Local: "path"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Direction, xml.StartElement{Name: xml.Name{Local: "direction"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *DataRequirementSort) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "path":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Path = v
			case "direction":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Direction = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r DataRequirementSort) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
