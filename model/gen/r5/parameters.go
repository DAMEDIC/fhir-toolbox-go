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

// This resource is used to pass information into and back from an operation (whether invoked directly from REST or within a messaging environment).  It is not persisted or allowed to be referenced by other resources except as described in the definition of the Parameters resource.
type Parameters struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *Id
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *Uri
	// The base language in which the resource is written.
	Language *Code
	// A parameter passed to or received from the operation.
	Parameter []ParametersParameter
}

// A parameter passed to or received from the operation.
type ParametersParameter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The name of the parameter (reference to the operation definition).
	Name String
	// Conveys the content if the parameter is a data type.
	Value isParametersParameterValue
	// Conveys the content if the parameter is a whole resource.
	Resource model.Resource
	// A named part of a multi-part parameter.
	Part []ParametersParameter
}
type isParametersParameterValue interface {
	model.Element
	isParametersParameterValue()
}

func (r Base64Binary) isParametersParameterValue()          {}
func (r Boolean) isParametersParameterValue()               {}
func (r Canonical) isParametersParameterValue()             {}
func (r Code) isParametersParameterValue()                  {}
func (r Date) isParametersParameterValue()                  {}
func (r DateTime) isParametersParameterValue()              {}
func (r Decimal) isParametersParameterValue()               {}
func (r Id) isParametersParameterValue()                    {}
func (r Instant) isParametersParameterValue()               {}
func (r Integer) isParametersParameterValue()               {}
func (r Integer64) isParametersParameterValue()             {}
func (r Markdown) isParametersParameterValue()              {}
func (r Oid) isParametersParameterValue()                   {}
func (r PositiveInt) isParametersParameterValue()           {}
func (r String) isParametersParameterValue()                {}
func (r Time) isParametersParameterValue()                  {}
func (r UnsignedInt) isParametersParameterValue()           {}
func (r Uri) isParametersParameterValue()                   {}
func (r Url) isParametersParameterValue()                   {}
func (r Uuid) isParametersParameterValue()                  {}
func (r Address) isParametersParameterValue()               {}
func (r Age) isParametersParameterValue()                   {}
func (r Annotation) isParametersParameterValue()            {}
func (r Attachment) isParametersParameterValue()            {}
func (r CodeableConcept) isParametersParameterValue()       {}
func (r CodeableReference) isParametersParameterValue()     {}
func (r Coding) isParametersParameterValue()                {}
func (r ContactPoint) isParametersParameterValue()          {}
func (r Count) isParametersParameterValue()                 {}
func (r Distance) isParametersParameterValue()              {}
func (r Duration) isParametersParameterValue()              {}
func (r HumanName) isParametersParameterValue()             {}
func (r Identifier) isParametersParameterValue()            {}
func (r Money) isParametersParameterValue()                 {}
func (r Period) isParametersParameterValue()                {}
func (r Quantity) isParametersParameterValue()              {}
func (r Range) isParametersParameterValue()                 {}
func (r Ratio) isParametersParameterValue()                 {}
func (r RatioRange) isParametersParameterValue()            {}
func (r Reference) isParametersParameterValue()             {}
func (r SampledData) isParametersParameterValue()           {}
func (r Signature) isParametersParameterValue()             {}
func (r Timing) isParametersParameterValue()                {}
func (r ContactDetail) isParametersParameterValue()         {}
func (r DataRequirement) isParametersParameterValue()       {}
func (r Expression) isParametersParameterValue()            {}
func (r ParameterDefinition) isParametersParameterValue()   {}
func (r RelatedArtifact) isParametersParameterValue()       {}
func (r TriggerDefinition) isParametersParameterValue()     {}
func (r UsageContext) isParametersParameterValue()          {}
func (r Availability) isParametersParameterValue()          {}
func (r ExtendedContactDetail) isParametersParameterValue() {}
func (r Dosage) isParametersParameterValue()                {}
func (r Meta) isParametersParameterValue()                  {}
func (r Parameters) ResourceType() string {
	return "Parameters"
}
func (r Parameters) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r Parameters) MemSize() int {
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
	for _, i := range r.Parameter {
		s += i.MemSize()
	}
	s += (cap(r.Parameter) - len(r.Parameter)) * int(unsafe.Sizeof(ParametersParameter{}))
	return s
}
func (r ParametersParameter) MemSize() int {
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
	s += r.Name.MemSize() - int(unsafe.Sizeof(r.Name))
	if r.Value != nil {
		s += r.Value.MemSize()
	}
	if r.Resource != nil {
		s += (r.Resource).MemSize()
	}
	for _, i := range r.Part {
		s += i.MemSize()
	}
	s += (cap(r.Part) - len(r.Part)) * int(unsafe.Sizeof(ParametersParameter{}))
	return s
}
func (r Parameters) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Parameters) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Parameters) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"Parameters\""))
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
	if len(r.Parameter) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"parameter\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Parameter {
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
func (r ParametersParameter) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ParametersParameter) marshalJSON(w io.Writer) error {
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
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"name\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Name)
		if err != nil {
			return err
		}
	}
	if r.Name.Id != nil || r.Name.Extension != nil {
		p := primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_name\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	switch v := r.Value.(type) {
	case Base64Binary:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueBase64Binary\":"))
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
			_, err = w.Write([]byte("\"_valueBase64Binary\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Base64Binary:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueBase64Binary\":"))
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
			_, err = w.Write([]byte("\"_valueBase64Binary\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Boolean:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueBoolean\":"))
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
			_, err = w.Write([]byte("\"_valueBoolean\":"))
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
			_, err = w.Write([]byte("\"valueBoolean\":"))
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
			_, err = w.Write([]byte("\"_valueBoolean\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Canonical:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueCanonical\":"))
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
			_, err = w.Write([]byte("\"_valueCanonical\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Canonical:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueCanonical\":"))
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
			_, err = w.Write([]byte("\"_valueCanonical\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Code:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueCode\":"))
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
			_, err = w.Write([]byte("\"_valueCode\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Code:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueCode\":"))
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
			_, err = w.Write([]byte("\"_valueCode\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Date:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDate\":"))
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
			_, err = w.Write([]byte("\"_valueDate\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Date:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDate\":"))
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
			_, err = w.Write([]byte("\"_valueDate\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDateTime\":"))
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
			_, err = w.Write([]byte("\"_valueDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *DateTime:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDateTime\":"))
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
			_, err = w.Write([]byte("\"_valueDateTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Decimal:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDecimal\":"))
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
			_, err = w.Write([]byte("\"_valueDecimal\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Decimal:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueDecimal\":"))
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
			_, err = w.Write([]byte("\"_valueDecimal\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Id:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueId\":"))
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
			_, err = w.Write([]byte("\"_valueId\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Id:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueId\":"))
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
			_, err = w.Write([]byte("\"_valueId\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Instant:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInstant\":"))
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
			_, err = w.Write([]byte("\"_valueInstant\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Instant:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInstant\":"))
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
			_, err = w.Write([]byte("\"_valueInstant\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Integer:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInteger\":"))
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
			_, err = w.Write([]byte("\"_valueInteger\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Integer:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInteger\":"))
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
			_, err = w.Write([]byte("\"_valueInteger\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Integer64:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInteger64\":"))
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
			_, err = w.Write([]byte("\"_valueInteger64\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Integer64:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueInteger64\":"))
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
			_, err = w.Write([]byte("\"_valueInteger64\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Markdown:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueMarkdown\":"))
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
			_, err = w.Write([]byte("\"_valueMarkdown\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Markdown:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueMarkdown\":"))
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
			_, err = w.Write([]byte("\"_valueMarkdown\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Oid:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueOid\":"))
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
			_, err = w.Write([]byte("\"_valueOid\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Oid:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueOid\":"))
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
			_, err = w.Write([]byte("\"_valueOid\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case PositiveInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valuePositiveInt\":"))
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
			_, err = w.Write([]byte("\"_valuePositiveInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *PositiveInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valuePositiveInt\":"))
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
			_, err = w.Write([]byte("\"_valuePositiveInt\":"))
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
			_, err = w.Write([]byte("\"valueString\":"))
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
			_, err = w.Write([]byte("\"_valueString\":"))
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
			_, err = w.Write([]byte("\"valueString\":"))
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
			_, err = w.Write([]byte("\"_valueString\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Time:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueTime\":"))
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
			_, err = w.Write([]byte("\"_valueTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Time:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueTime\":"))
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
			_, err = w.Write([]byte("\"_valueTime\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case UnsignedInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUnsignedInt\":"))
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
			_, err = w.Write([]byte("\"_valueUnsignedInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *UnsignedInt:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUnsignedInt\":"))
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
			_, err = w.Write([]byte("\"_valueUnsignedInt\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Uri:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUri\":"))
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
			_, err = w.Write([]byte("\"_valueUri\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Uri:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUri\":"))
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
			_, err = w.Write([]byte("\"_valueUri\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Url:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUrl\":"))
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
			_, err = w.Write([]byte("\"_valueUrl\":"))
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
			_, err = w.Write([]byte("\"valueUrl\":"))
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
			_, err = w.Write([]byte("\"_valueUrl\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Uuid:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUuid\":"))
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
			_, err = w.Write([]byte("\"_valueUuid\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case *Uuid:
		if v.Value != nil {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"valueUuid\":"))
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
			_, err = w.Write([]byte("\"_valueUuid\":"))
			if err != nil {
				return err
			}
			err = p.marshalJSON(w)
			if err != nil {
				return err
			}
		}
	case Address:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAddress\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Address:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAddress\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Age:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAge\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Age:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAge\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Annotation:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAnnotation\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Annotation:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAnnotation\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Attachment:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAttachment\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Attachment:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAttachment\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case CodeableConcept:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCodeableConcept\":"))
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
		_, err = w.Write([]byte("\"valueCodeableConcept\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case CodeableReference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCodeableReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *CodeableReference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCodeableReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Coding:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCoding\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Coding:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCoding\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case ContactPoint:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueContactPoint\":"))
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
		_, err = w.Write([]byte("\"valueContactPoint\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Count:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCount\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Count:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueCount\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Distance:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDistance\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Distance:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDistance\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Duration:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDuration\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Duration:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDuration\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case HumanName:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueHumanName\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *HumanName:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueHumanName\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Identifier:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueIdentifier\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Identifier:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueIdentifier\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Money:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueMoney\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Money:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueMoney\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valuePeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valuePeriod\":"))
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
		_, err = w.Write([]byte("\"valueQuantity\":"))
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
		_, err = w.Write([]byte("\"valueQuantity\":"))
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
		_, err = w.Write([]byte("\"valueRange\":"))
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
		_, err = w.Write([]byte("\"valueRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Ratio:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRatio\":"))
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
		_, err = w.Write([]byte("\"valueRatio\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case RatioRange:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRatioRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *RatioRange:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRatioRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Reference:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueReference\":"))
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
		_, err = w.Write([]byte("\"valueReference\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case SampledData:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueSampledData\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *SampledData:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueSampledData\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Signature:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueSignature\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Signature:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueSignature\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Timing:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueTiming\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Timing:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueTiming\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case ContactDetail:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueContactDetail\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *ContactDetail:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueContactDetail\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case DataRequirement:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDataRequirement\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *DataRequirement:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDataRequirement\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Expression:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueExpression\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Expression:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueExpression\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case ParameterDefinition:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueParameterDefinition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *ParameterDefinition:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueParameterDefinition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case RelatedArtifact:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRelatedArtifact\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *RelatedArtifact:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueRelatedArtifact\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case TriggerDefinition:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueTriggerDefinition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *TriggerDefinition:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueTriggerDefinition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case UsageContext:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueUsageContext\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *UsageContext:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueUsageContext\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Availability:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAvailability\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Availability:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueAvailability\":"))
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
		_, err = w.Write([]byte("\"valueExtendedContactDetail\":"))
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
		_, err = w.Write([]byte("\"valueExtendedContactDetail\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Dosage:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDosage\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Dosage:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueDosage\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Meta:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueMeta\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Meta:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"valueMeta\":"))
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
	if r.Resource != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"resource\":"))
		if err != nil {
			return err
		}
		err = ContainedResource{r.Resource}.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Part) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"part\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Part {
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
func (r *Parameters) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *Parameters) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in Parameters element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in Parameters element", t)
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
		case "parameter":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in Parameters element", t)
			}
			for d.More() {
				var v ParametersParameter
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Parameter = append(r.Parameter, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in Parameters element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in Parameters", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in Parameters element", t)
	}
	return nil
}
func (r *ParametersParameter) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ParametersParameter element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ParametersParameter element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ParametersParameter element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ParametersParameter element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ParametersParameter element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ParametersParameter element", t)
			}
		case "name":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Name.Value = v.Value
		case "_name":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Name.Id = v.Id
			r.Name.Extension = v.Extension
		case "valueBase64Binary":
			var v Base64Binary
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Base64Binary{
					Extension: r.Value.(Base64Binary).Extension,
					Id:        r.Value.(Base64Binary).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueBase64Binary":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Base64Binary{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Base64Binary).Value,
				}
			} else {
				r.Value = Base64Binary{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueBoolean":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Boolean{
					Extension: r.Value.(Boolean).Extension,
					Id:        r.Value.(Boolean).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueBoolean":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Boolean{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Boolean).Value,
				}
			} else {
				r.Value = Boolean{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueCanonical":
			var v Canonical
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Canonical{
					Extension: r.Value.(Canonical).Extension,
					Id:        r.Value.(Canonical).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueCanonical":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Canonical{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Canonical).Value,
				}
			} else {
				r.Value = Canonical{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueCode":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Code{
					Extension: r.Value.(Code).Extension,
					Id:        r.Value.(Code).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueCode":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Code{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Code).Value,
				}
			} else {
				r.Value = Code{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueDate":
			var v Date
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Date{
					Extension: r.Value.(Date).Extension,
					Id:        r.Value.(Date).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Date{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Date).Value,
				}
			} else {
				r.Value = Date{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueDateTime":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = DateTime{
					Extension: r.Value.(DateTime).Extension,
					Id:        r.Value.(DateTime).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueDateTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(DateTime).Value,
				}
			} else {
				r.Value = DateTime{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueDecimal":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Decimal{
					Extension: r.Value.(Decimal).Extension,
					Id:        r.Value.(Decimal).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueDecimal":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Decimal{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Decimal).Value,
				}
			} else {
				r.Value = Decimal{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueId":
			var v Id
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Id{
					Extension: r.Value.(Id).Extension,
					Id:        r.Value.(Id).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueId":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Id{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Id).Value,
				}
			} else {
				r.Value = Id{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueInstant":
			var v Instant
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Instant{
					Extension: r.Value.(Instant).Extension,
					Id:        r.Value.(Instant).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueInstant":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Instant{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Instant).Value,
				}
			} else {
				r.Value = Instant{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueInteger":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Integer{
					Extension: r.Value.(Integer).Extension,
					Id:        r.Value.(Integer).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueInteger":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Integer{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Integer).Value,
				}
			} else {
				r.Value = Integer{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueInteger64":
			var v Integer64
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Integer64{
					Extension: r.Value.(Integer64).Extension,
					Id:        r.Value.(Integer64).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueInteger64":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Integer64{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Integer64).Value,
				}
			} else {
				r.Value = Integer64{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueMarkdown":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Markdown{
					Extension: r.Value.(Markdown).Extension,
					Id:        r.Value.(Markdown).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueMarkdown":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Markdown{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Markdown).Value,
				}
			} else {
				r.Value = Markdown{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueOid":
			var v Oid
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Oid{
					Extension: r.Value.(Oid).Extension,
					Id:        r.Value.(Oid).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueOid":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Oid{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Oid).Value,
				}
			} else {
				r.Value = Oid{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valuePositiveInt":
			var v PositiveInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = PositiveInt{
					Extension: r.Value.(PositiveInt).Extension,
					Id:        r.Value.(PositiveInt).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valuePositiveInt":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = PositiveInt{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(PositiveInt).Value,
				}
			} else {
				r.Value = PositiveInt{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueString":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = String{
					Extension: r.Value.(String).Extension,
					Id:        r.Value.(String).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueString":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = String{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(String).Value,
				}
			} else {
				r.Value = String{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueTime":
			var v Time
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Time{
					Extension: r.Value.(Time).Extension,
					Id:        r.Value.(Time).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Time{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Time).Value,
				}
			} else {
				r.Value = Time{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueUnsignedInt":
			var v UnsignedInt
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = UnsignedInt{
					Extension: r.Value.(UnsignedInt).Extension,
					Id:        r.Value.(UnsignedInt).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueUnsignedInt":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = UnsignedInt{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(UnsignedInt).Value,
				}
			} else {
				r.Value = UnsignedInt{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueUri":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Uri{
					Extension: r.Value.(Uri).Extension,
					Id:        r.Value.(Uri).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueUri":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Uri{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Uri).Value,
				}
			} else {
				r.Value = Uri{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueUrl":
			var v Url
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Url{
					Extension: r.Value.(Url).Extension,
					Id:        r.Value.(Url).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueUrl":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Url{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Url).Value,
				}
			} else {
				r.Value = Url{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueUuid":
			var v Uuid
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Uuid{
					Extension: r.Value.(Uuid).Extension,
					Id:        r.Value.(Uuid).Id,
					Value:     v.Value,
				}
			} else {
				r.Value = v
			}
		case "_valueUuid":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Value != nil {
				r.Value = Uuid{
					Extension: v.Extension,
					Id:        v.Id,
					Value:     r.Value.(Uuid).Value,
				}
			} else {
				r.Value = Uuid{
					Extension: v.Extension,
					Id:        v.Id,
				}
			}
		case "valueAddress":
			var v Address
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueAge":
			var v Age
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueAnnotation":
			var v Annotation
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueAttachment":
			var v Attachment
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueCodeableConcept":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueCodeableReference":
			var v CodeableReference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueCoding":
			var v Coding
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueContactPoint":
			var v ContactPoint
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueCount":
			var v Count
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueDistance":
			var v Distance
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueDuration":
			var v Duration
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueHumanName":
			var v HumanName
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueIdentifier":
			var v Identifier
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueMoney":
			var v Money
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valuePeriod":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueQuantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueRange":
			var v Range
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueRatio":
			var v Ratio
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueRatioRange":
			var v RatioRange
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueReference":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueSampledData":
			var v SampledData
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueSignature":
			var v Signature
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueTiming":
			var v Timing
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueContactDetail":
			var v ContactDetail
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueDataRequirement":
			var v DataRequirement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueExpression":
			var v Expression
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueParameterDefinition":
			var v ParameterDefinition
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueRelatedArtifact":
			var v RelatedArtifact
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueTriggerDefinition":
			var v TriggerDefinition
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueUsageContext":
			var v UsageContext
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueAvailability":
			var v Availability
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueExtendedContactDetail":
			var v ExtendedContactDetail
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueDosage":
			var v Dosage
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "valueMeta":
			var v Meta
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Value = v
		case "resource":
			var v ContainedResource
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Resource = v.Resource
		case "part":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ParametersParameter element", t)
			}
			for d.More() {
				var v ParametersParameter
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Part = append(r.Part, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ParametersParameter element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in ParametersParameter", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ParametersParameter element", t)
	}
	return nil
}
func (r Parameters) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "Parameters"
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
	err = e.EncodeElement(r.Parameter, xml.StartElement{Name: xml.Name{Local: "parameter"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ParametersParameter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	switch v := r.Value.(type) {
	case Base64Binary, *Base64Binary:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueBase64Binary"}})
		if err != nil {
			return err
		}
	case Boolean, *Boolean:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueBoolean"}})
		if err != nil {
			return err
		}
	case Canonical, *Canonical:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCanonical"}})
		if err != nil {
			return err
		}
	case Code, *Code:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCode"}})
		if err != nil {
			return err
		}
	case Date, *Date:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDate"}})
		if err != nil {
			return err
		}
	case DateTime, *DateTime:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDateTime"}})
		if err != nil {
			return err
		}
	case Decimal, *Decimal:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDecimal"}})
		if err != nil {
			return err
		}
	case Id, *Id:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueId"}})
		if err != nil {
			return err
		}
	case Instant, *Instant:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueInstant"}})
		if err != nil {
			return err
		}
	case Integer, *Integer:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueInteger"}})
		if err != nil {
			return err
		}
	case Integer64, *Integer64:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueInteger64"}})
		if err != nil {
			return err
		}
	case Markdown, *Markdown:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueMarkdown"}})
		if err != nil {
			return err
		}
	case Oid, *Oid:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueOid"}})
		if err != nil {
			return err
		}
	case PositiveInt, *PositiveInt:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valuePositiveInt"}})
		if err != nil {
			return err
		}
	case String, *String:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueString"}})
		if err != nil {
			return err
		}
	case Time, *Time:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueTime"}})
		if err != nil {
			return err
		}
	case UnsignedInt, *UnsignedInt:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUnsignedInt"}})
		if err != nil {
			return err
		}
	case Uri, *Uri:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUri"}})
		if err != nil {
			return err
		}
	case Url, *Url:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUrl"}})
		if err != nil {
			return err
		}
	case Uuid, *Uuid:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUuid"}})
		if err != nil {
			return err
		}
	case Address, *Address:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAddress"}})
		if err != nil {
			return err
		}
	case Age, *Age:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAge"}})
		if err != nil {
			return err
		}
	case Annotation, *Annotation:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAnnotation"}})
		if err != nil {
			return err
		}
	case Attachment, *Attachment:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAttachment"}})
		if err != nil {
			return err
		}
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCodeableConcept"}})
		if err != nil {
			return err
		}
	case CodeableReference, *CodeableReference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCodeableReference"}})
		if err != nil {
			return err
		}
	case Coding, *Coding:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCoding"}})
		if err != nil {
			return err
		}
	case ContactPoint, *ContactPoint:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueContactPoint"}})
		if err != nil {
			return err
		}
	case Count, *Count:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueCount"}})
		if err != nil {
			return err
		}
	case Distance, *Distance:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDistance"}})
		if err != nil {
			return err
		}
	case Duration, *Duration:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDuration"}})
		if err != nil {
			return err
		}
	case HumanName, *HumanName:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueHumanName"}})
		if err != nil {
			return err
		}
	case Identifier, *Identifier:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueIdentifier"}})
		if err != nil {
			return err
		}
	case Money, *Money:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueMoney"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valuePeriod"}})
		if err != nil {
			return err
		}
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueQuantity"}})
		if err != nil {
			return err
		}
	case Range, *Range:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueRange"}})
		if err != nil {
			return err
		}
	case Ratio, *Ratio:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueRatio"}})
		if err != nil {
			return err
		}
	case RatioRange, *RatioRange:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueRatioRange"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueReference"}})
		if err != nil {
			return err
		}
	case SampledData, *SampledData:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueSampledData"}})
		if err != nil {
			return err
		}
	case Signature, *Signature:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueSignature"}})
		if err != nil {
			return err
		}
	case Timing, *Timing:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueTiming"}})
		if err != nil {
			return err
		}
	case ContactDetail, *ContactDetail:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueContactDetail"}})
		if err != nil {
			return err
		}
	case DataRequirement, *DataRequirement:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDataRequirement"}})
		if err != nil {
			return err
		}
	case Expression, *Expression:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueExpression"}})
		if err != nil {
			return err
		}
	case ParameterDefinition, *ParameterDefinition:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueParameterDefinition"}})
		if err != nil {
			return err
		}
	case RelatedArtifact, *RelatedArtifact:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueRelatedArtifact"}})
		if err != nil {
			return err
		}
	case TriggerDefinition, *TriggerDefinition:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueTriggerDefinition"}})
		if err != nil {
			return err
		}
	case UsageContext, *UsageContext:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueUsageContext"}})
		if err != nil {
			return err
		}
	case Availability, *Availability:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueAvailability"}})
		if err != nil {
			return err
		}
	case ExtendedContactDetail, *ExtendedContactDetail:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueExtendedContactDetail"}})
		if err != nil {
			return err
		}
	case Dosage, *Dosage:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueDosage"}})
		if err != nil {
			return err
		}
	case Meta, *Meta:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "valueMeta"}})
		if err != nil {
			return err
		}
	}
	if r.Resource != nil {
		err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "resource"}})
		if err != nil {
			return err
		}
		err = e.EncodeElement(r.Resource, xml.StartElement{Name: xml.Name{Local: "__contained__"}})
		if err != nil {
			return err
		}
		err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "resource"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Part, xml.StartElement{Name: xml.Name{Local: "part"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Parameters) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "parameter":
				var v ParametersParameter
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Parameter = append(r.Parameter, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ParametersParameter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = v
			case "valueBase64Binary":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Base64Binary
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueBoolean":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCanonical":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCode":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDate":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
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
			case "valueDecimal":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueId":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Id
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueInstant":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Instant
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueInteger":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueInteger64":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Integer64
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueMarkdown":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueOid":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Oid
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valuePositiveInt":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueString":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueTime":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Time
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUnsignedInt":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v UnsignedInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUri":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUrl":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Url
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUuid":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Uuid
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAddress":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Address
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAge":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Age
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAnnotation":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAttachment":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCodeableConcept":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCodeableReference":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCoding":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Coding
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueContactPoint":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v ContactPoint
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueCount":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Count
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDistance":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Distance
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
			case "valueHumanName":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v HumanName
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueIdentifier":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueMoney":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Money
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
			case "valueQuantity":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueRange":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueRatio":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueRatioRange":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v RatioRange
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueReference":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueSampledData":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v SampledData
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueSignature":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Signature
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueTiming":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Timing
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueContactDetail":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDataRequirement":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v DataRequirement
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueExpression":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Expression
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueParameterDefinition":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v ParameterDefinition
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueRelatedArtifact":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v RelatedArtifact
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueTriggerDefinition":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v TriggerDefinition
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueUsageContext":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v UsageContext
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueAvailability":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Availability
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueExtendedContactDetail":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v ExtendedContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueDosage":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Dosage
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "valueMeta":
				if r.Value != nil {
					return fmt.Errorf("multiple values for field \"Value\"")
				}
				var v Meta
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = v
			case "resource":
				var c ContainedResource
				err := d.DecodeElement(&c, &t)
				if err != nil {
					return err
				}
				r.Resource = c.Resource
			case "part":
				var v ParametersParameter
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Part = append(r.Part, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Parameters) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "parameter") {
		for _, v := range r.Parameter {
			children = append(children, v)
		}
	}
	return children
}
func (r Parameters) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert Parameters to Boolean")
}
func (r Parameters) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert Parameters to String")
}
func (r Parameters) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert Parameters to Integer")
}
func (r Parameters) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert Parameters to Decimal")
}
func (r Parameters) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert Parameters to Date")
}
func (r Parameters) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert Parameters to Time")
}
func (r Parameters) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert Parameters to DateTime")
}
func (r Parameters) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert Parameters to Quantity")
}
func (r Parameters) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.Id",
		}, {
			Name: "Meta",
			Type: "FHIR.Meta",
		}, {
			Name: "ImplicitRules",
			Type: "FHIR.Uri",
		}, {
			Name: "Language",
			Type: "FHIR.Code",
		}, {
			Name: "Parameter",
			Type: "List<FHIR.ParametersParameter>",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "Parameters",
			Namespace: "FHIR",
		},
	}
}
func (r ParametersParameter) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "name") {
		children = append(children, r.Name)
	}
	if len(name) == 0 || slices.Contains(name, "value") {
		if r.Value != nil {
			children = append(children, r.Value)
		}
	}
	if len(name) == 0 || slices.Contains(name, "resource") {
		if r.Resource != nil {
			children = append(children, r.Resource)
		}
	}
	if len(name) == 0 || slices.Contains(name, "part") {
		for _, v := range r.Part {
			children = append(children, v)
		}
	}
	return children
}
func (r ParametersParameter) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert ParametersParameter to Boolean")
}
func (r ParametersParameter) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert ParametersParameter to String")
}
func (r ParametersParameter) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert ParametersParameter to Integer")
}
func (r ParametersParameter) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert ParametersParameter to Decimal")
}
func (r ParametersParameter) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert ParametersParameter to Date")
}
func (r ParametersParameter) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert ParametersParameter to Time")
}
func (r ParametersParameter) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert ParametersParameter to DateTime")
}
func (r ParametersParameter) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert ParametersParameter to Quantity")
}
func (r ParametersParameter) TypeInfo() fhirpath.TypeInfo {
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
			Name: "Name",
			Type: "FHIR.String",
		}, {
			Name: "Value",
			Type: "FHIR.PrimitiveElement",
		}, {
			Name: "Resource",
			Type: "FHIR.",
		}, {
			Name: "Part",
			Type: "List<FHIR.ParametersParameter>",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "ParametersParameter",
			Namespace: "FHIR",
		},
	}
}
