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

// Set of definitional characteristics for a kind of observation or measurement produced or consumed by an orderable health care service.
//
// In a catalog of health-related services that use or produce observations and measurements, this resource describes the expected characteristics of these observation / measurements.
type ObservationDefinition struct {
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
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code that classifies the general type of observation.
	Category []CodeableConcept
	// Describes what will be observed. Sometimes this is called the observation "name".
	Code CodeableConcept
	// A unique identifier assigned to this ObservationDefinition artifact.
	Identifier []Identifier
	// The data types allowed for the value element of the instance observations conforming to this ObservationDefinition.
	PermittedDataType []Code
	// Multiple results allowed for observations conforming to this ObservationDefinition.
	MultipleResultsAllowed *Boolean
	// The method or technique used to perform the observation.
	Method *CodeableConcept
	// The preferred name to be used when reporting the results of observations conforming to this ObservationDefinition.
	PreferredReportName *String
	// Characteristics for quantitative results of this observation.
	QuantitativeDetails *ObservationDefinitionQuantitativeDetails
	// Multiple  ranges of results qualified by different contexts for ordinal or continuous observations conforming to this ObservationDefinition.
	QualifiedInterval []ObservationDefinitionQualifiedInterval
	// The set of valid coded results for the observations  conforming to this ObservationDefinition.
	ValidCodedValueSet *Reference
	// The set of normal coded results for the observations conforming to this ObservationDefinition.
	NormalCodedValueSet *Reference
	// The set of abnormal coded results for the observation conforming to this ObservationDefinition.
	AbnormalCodedValueSet *Reference
	// The set of critical coded results for the observation conforming to this ObservationDefinition.
	CriticalCodedValueSet *Reference
}

// Characteristics for quantitative results of this observation.
type ObservationDefinitionQuantitativeDetails struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Customary unit used to report quantitative results of observations conforming to this ObservationDefinition.
	CustomaryUnit *CodeableConcept
	// SI unit used to report quantitative results of observations conforming to this ObservationDefinition.
	Unit *CodeableConcept
	// Factor for converting value expressed with SI unit to value expressed with customary unit.
	ConversionFactor *Decimal
	// Number of digits after decimal separator when the results of such observations are of type Quantity.
	DecimalPrecision *Integer
}

// Multiple  ranges of results qualified by different contexts for ordinal or continuous observations conforming to this ObservationDefinition.
type ObservationDefinitionQualifiedInterval struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The category of interval of values for continuous or ordinal observations conforming to this ObservationDefinition.
	Category *Code
	// The low and high values determining the interval. There may be only one of the two.
	Range *Range
	// Codes to indicate the health context the range applies to. For example, the normal or therapeutic range.
	Context *CodeableConcept
	// Codes to indicate the target population this reference range applies to.
	AppliesTo []CodeableConcept
	// Sex of the population the range applies to.
	Gender *Code
	// The age at which this reference range is applicable. This is a neonatal age (e.g. number of weeks at term) if the meaning says so.
	Age *Range
	// The gestational age to which this reference range is applicable, in the context of pregnancy.
	GestationalAge *Range
	// Text based condition for which the reference range is valid.
	Condition *String
}

func (r ObservationDefinition) ResourceType() string {
	return "ObservationDefinition"
}
func (r ObservationDefinition) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r ObservationDefinition) MemSize() int {
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
	for _, i := range r.Category {
		s += i.MemSize()
	}
	s += (cap(r.Category) - len(r.Category)) * int(unsafe.Sizeof(CodeableConcept{}))
	s += r.Code.MemSize() - int(unsafe.Sizeof(r.Code))
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	for _, i := range r.PermittedDataType {
		s += i.MemSize()
	}
	s += (cap(r.PermittedDataType) - len(r.PermittedDataType)) * int(unsafe.Sizeof(Code{}))
	if r.MultipleResultsAllowed != nil {
		s += r.MultipleResultsAllowed.MemSize()
	}
	if r.Method != nil {
		s += r.Method.MemSize()
	}
	if r.PreferredReportName != nil {
		s += r.PreferredReportName.MemSize()
	}
	if r.QuantitativeDetails != nil {
		s += r.QuantitativeDetails.MemSize()
	}
	for _, i := range r.QualifiedInterval {
		s += i.MemSize()
	}
	s += (cap(r.QualifiedInterval) - len(r.QualifiedInterval)) * int(unsafe.Sizeof(ObservationDefinitionQualifiedInterval{}))
	if r.ValidCodedValueSet != nil {
		s += r.ValidCodedValueSet.MemSize()
	}
	if r.NormalCodedValueSet != nil {
		s += r.NormalCodedValueSet.MemSize()
	}
	if r.AbnormalCodedValueSet != nil {
		s += r.AbnormalCodedValueSet.MemSize()
	}
	if r.CriticalCodedValueSet != nil {
		s += r.CriticalCodedValueSet.MemSize()
	}
	return s
}
func (r ObservationDefinitionQuantitativeDetails) MemSize() int {
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
	if r.CustomaryUnit != nil {
		s += r.CustomaryUnit.MemSize()
	}
	if r.Unit != nil {
		s += r.Unit.MemSize()
	}
	if r.ConversionFactor != nil {
		s += r.ConversionFactor.MemSize()
	}
	if r.DecimalPrecision != nil {
		s += r.DecimalPrecision.MemSize()
	}
	return s
}
func (r ObservationDefinitionQualifiedInterval) MemSize() int {
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
	if r.Category != nil {
		s += r.Category.MemSize()
	}
	if r.Range != nil {
		s += r.Range.MemSize()
	}
	if r.Context != nil {
		s += r.Context.MemSize()
	}
	for _, i := range r.AppliesTo {
		s += i.MemSize()
	}
	s += (cap(r.AppliesTo) - len(r.AppliesTo)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Gender != nil {
		s += r.Gender.MemSize()
	}
	if r.Age != nil {
		s += r.Age.MemSize()
	}
	if r.GestationalAge != nil {
		s += r.GestationalAge.MemSize()
	}
	if r.Condition != nil {
		s += r.Condition.MemSize()
	}
	return s
}
func (r ObservationDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ObservationDefinitionQuantitativeDetails) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ObservationDefinitionQualifiedInterval) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r ObservationDefinition) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ObservationDefinition) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"ObservationDefinition\""))
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
	if len(r.Category) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"category\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Category {
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
	_, err = w.Write([]byte("\"code\":"))
	if err != nil {
		return err
	}
	err = r.Code.marshalJSON(w)
	if err != nil {
		return err
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
	{
		anyValue := false
		for _, e := range r.PermittedDataType {
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
			_, err = w.Write([]byte("\"permittedDataType\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.PermittedDataType)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.PermittedDataType {
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
			_, err = w.Write([]byte("\"_permittedDataType\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.PermittedDataType {
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
	if r.MultipleResultsAllowed != nil && r.MultipleResultsAllowed.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"multipleResultsAllowed\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.MultipleResultsAllowed)
		if err != nil {
			return err
		}
	}
	if r.MultipleResultsAllowed != nil && (r.MultipleResultsAllowed.Id != nil || r.MultipleResultsAllowed.Extension != nil) {
		p := primitiveElement{Id: r.MultipleResultsAllowed.Id, Extension: r.MultipleResultsAllowed.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_multipleResultsAllowed\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
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
	if r.PreferredReportName != nil && r.PreferredReportName.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"preferredReportName\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.PreferredReportName)
		if err != nil {
			return err
		}
	}
	if r.PreferredReportName != nil && (r.PreferredReportName.Id != nil || r.PreferredReportName.Extension != nil) {
		p := primitiveElement{Id: r.PreferredReportName.Id, Extension: r.PreferredReportName.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_preferredReportName\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.QuantitativeDetails != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"quantitativeDetails\":"))
		if err != nil {
			return err
		}
		err = r.QuantitativeDetails.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.QualifiedInterval) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"qualifiedInterval\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.QualifiedInterval {
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
	if r.ValidCodedValueSet != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"validCodedValueSet\":"))
		if err != nil {
			return err
		}
		err = r.ValidCodedValueSet.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.NormalCodedValueSet != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"normalCodedValueSet\":"))
		if err != nil {
			return err
		}
		err = r.NormalCodedValueSet.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.AbnormalCodedValueSet != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"abnormalCodedValueSet\":"))
		if err != nil {
			return err
		}
		err = r.AbnormalCodedValueSet.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.CriticalCodedValueSet != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"criticalCodedValueSet\":"))
		if err != nil {
			return err
		}
		err = r.CriticalCodedValueSet.marshalJSON(w)
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
func (r ObservationDefinitionQuantitativeDetails) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ObservationDefinitionQuantitativeDetails) marshalJSON(w io.Writer) error {
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
	if r.CustomaryUnit != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"customaryUnit\":"))
		if err != nil {
			return err
		}
		err = r.CustomaryUnit.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Unit != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"unit\":"))
		if err != nil {
			return err
		}
		err = r.Unit.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ConversionFactor != nil && r.ConversionFactor.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"conversionFactor\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ConversionFactor)
		if err != nil {
			return err
		}
	}
	if r.ConversionFactor != nil && (r.ConversionFactor.Id != nil || r.ConversionFactor.Extension != nil) {
		p := primitiveElement{Id: r.ConversionFactor.Id, Extension: r.ConversionFactor.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_conversionFactor\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.DecimalPrecision != nil && r.DecimalPrecision.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"decimalPrecision\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.DecimalPrecision)
		if err != nil {
			return err
		}
	}
	if r.DecimalPrecision != nil && (r.DecimalPrecision.Id != nil || r.DecimalPrecision.Extension != nil) {
		p := primitiveElement{Id: r.DecimalPrecision.Id, Extension: r.DecimalPrecision.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_decimalPrecision\":"))
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
func (r ObservationDefinitionQualifiedInterval) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r ObservationDefinitionQualifiedInterval) marshalJSON(w io.Writer) error {
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
	if r.Category != nil && r.Category.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"category\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Category)
		if err != nil {
			return err
		}
	}
	if r.Category != nil && (r.Category.Id != nil || r.Category.Extension != nil) {
		p := primitiveElement{Id: r.Category.Id, Extension: r.Category.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_category\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Range != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"range\":"))
		if err != nil {
			return err
		}
		err = r.Range.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Context != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"context\":"))
		if err != nil {
			return err
		}
		err = r.Context.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.AppliesTo) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"appliesTo\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.AppliesTo {
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
	if r.Gender != nil && r.Gender.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"gender\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Gender)
		if err != nil {
			return err
		}
	}
	if r.Gender != nil && (r.Gender.Id != nil || r.Gender.Extension != nil) {
		p := primitiveElement{Id: r.Gender.Id, Extension: r.Gender.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_gender\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Age != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"age\":"))
		if err != nil {
			return err
		}
		err = r.Age.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.GestationalAge != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"gestationalAge\":"))
		if err != nil {
			return err
		}
		err = r.GestationalAge.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Condition != nil && r.Condition.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"condition\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Condition)
		if err != nil {
			return err
		}
	}
	if r.Condition != nil && (r.Condition.Id != nil || r.Condition.Extension != nil) {
		p := primitiveElement{Id: r.Condition.Id, Extension: r.Condition.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_condition\":"))
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
func (r *ObservationDefinition) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *ObservationDefinition) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ObservationDefinition element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ObservationDefinition element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ObservationDefinition element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ObservationDefinition element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ObservationDefinition element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ObservationDefinition element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ObservationDefinition element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ObservationDefinition element", t)
			}
		case "category":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ObservationDefinition element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Category = append(r.Category, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ObservationDefinition element", t)
			}
		case "code":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Code = v
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ObservationDefinition element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ObservationDefinition element", t)
			}
		case "permittedDataType":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ObservationDefinition element", t)
			}
			for i := 0; d.More(); i++ {
				var v Code
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.PermittedDataType) <= i {
					r.PermittedDataType = append(r.PermittedDataType, Code{})
				}
				r.PermittedDataType[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ObservationDefinition element", t)
			}
		case "_permittedDataType":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ObservationDefinition element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.PermittedDataType) <= i {
					r.PermittedDataType = append(r.PermittedDataType, Code{})
				}
				r.PermittedDataType[i].Id = v.Id
				r.PermittedDataType[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ObservationDefinition element", t)
			}
		case "multipleResultsAllowed":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.MultipleResultsAllowed == nil {
				r.MultipleResultsAllowed = &Boolean{}
			}
			r.MultipleResultsAllowed.Value = v.Value
		case "_multipleResultsAllowed":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.MultipleResultsAllowed == nil {
				r.MultipleResultsAllowed = &Boolean{}
			}
			r.MultipleResultsAllowed.Id = v.Id
			r.MultipleResultsAllowed.Extension = v.Extension
		case "method":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Method = &v
		case "preferredReportName":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.PreferredReportName == nil {
				r.PreferredReportName = &String{}
			}
			r.PreferredReportName.Value = v.Value
		case "_preferredReportName":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.PreferredReportName == nil {
				r.PreferredReportName = &String{}
			}
			r.PreferredReportName.Id = v.Id
			r.PreferredReportName.Extension = v.Extension
		case "quantitativeDetails":
			var v ObservationDefinitionQuantitativeDetails
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.QuantitativeDetails = &v
		case "qualifiedInterval":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ObservationDefinition element", t)
			}
			for d.More() {
				var v ObservationDefinitionQualifiedInterval
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.QualifiedInterval = append(r.QualifiedInterval, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ObservationDefinition element", t)
			}
		case "validCodedValueSet":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ValidCodedValueSet = &v
		case "normalCodedValueSet":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.NormalCodedValueSet = &v
		case "abnormalCodedValueSet":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.AbnormalCodedValueSet = &v
		case "criticalCodedValueSet":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.CriticalCodedValueSet = &v
		default:
			return fmt.Errorf("invalid field: %s in ObservationDefinition", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ObservationDefinition element", t)
	}
	return nil
}
func (r *ObservationDefinitionQuantitativeDetails) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ObservationDefinitionQuantitativeDetails element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ObservationDefinitionQuantitativeDetails element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ObservationDefinitionQuantitativeDetails element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ObservationDefinitionQuantitativeDetails element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ObservationDefinitionQuantitativeDetails element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ObservationDefinitionQuantitativeDetails element", t)
			}
		case "customaryUnit":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.CustomaryUnit = &v
		case "unit":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Unit = &v
		case "conversionFactor":
			var v Decimal
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ConversionFactor == nil {
				r.ConversionFactor = &Decimal{}
			}
			r.ConversionFactor.Value = v.Value
		case "_conversionFactor":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ConversionFactor == nil {
				r.ConversionFactor = &Decimal{}
			}
			r.ConversionFactor.Id = v.Id
			r.ConversionFactor.Extension = v.Extension
		case "decimalPrecision":
			var v Integer
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.DecimalPrecision == nil {
				r.DecimalPrecision = &Integer{}
			}
			r.DecimalPrecision.Value = v.Value
		case "_decimalPrecision":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.DecimalPrecision == nil {
				r.DecimalPrecision = &Integer{}
			}
			r.DecimalPrecision.Id = v.Id
			r.DecimalPrecision.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in ObservationDefinitionQuantitativeDetails", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ObservationDefinitionQuantitativeDetails element", t)
	}
	return nil
}
func (r *ObservationDefinitionQualifiedInterval) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in ObservationDefinitionQualifiedInterval element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in ObservationDefinitionQualifiedInterval element", t)
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
				return fmt.Errorf("invalid token: %v, expected: '[' in ObservationDefinitionQualifiedInterval element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ObservationDefinitionQualifiedInterval element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ObservationDefinitionQualifiedInterval element", t)
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
				return fmt.Errorf("invalid token: %v, expected: ']' in ObservationDefinitionQualifiedInterval element", t)
			}
		case "category":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Category == nil {
				r.Category = &Code{}
			}
			r.Category.Value = v.Value
		case "_category":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Category == nil {
				r.Category = &Code{}
			}
			r.Category.Id = v.Id
			r.Category.Extension = v.Extension
		case "range":
			var v Range
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Range = &v
		case "context":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Context = &v
		case "appliesTo":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in ObservationDefinitionQualifiedInterval element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.AppliesTo = append(r.AppliesTo, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in ObservationDefinitionQualifiedInterval element", t)
			}
		case "gender":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Gender == nil {
				r.Gender = &Code{}
			}
			r.Gender.Value = v.Value
		case "_gender":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Gender == nil {
				r.Gender = &Code{}
			}
			r.Gender.Id = v.Id
			r.Gender.Extension = v.Extension
		case "age":
			var v Range
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Age = &v
		case "gestationalAge":
			var v Range
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.GestationalAge = &v
		case "condition":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Condition == nil {
				r.Condition = &String{}
			}
			r.Condition.Value = v.Value
		case "_condition":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Condition == nil {
				r.Condition = &String{}
			}
			r.Condition.Id = v.Id
			r.Condition.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in ObservationDefinitionQualifiedInterval", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in ObservationDefinitionQualifiedInterval element", t)
	}
	return nil
}
func (r ObservationDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "ObservationDefinition"
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
	err = e.EncodeElement(r.Category, xml.StartElement{Name: xml.Name{Local: "category"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PermittedDataType, xml.StartElement{Name: xml.Name{Local: "permittedDataType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MultipleResultsAllowed, xml.StartElement{Name: xml.Name{Local: "multipleResultsAllowed"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Method, xml.StartElement{Name: xml.Name{Local: "method"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PreferredReportName, xml.StartElement{Name: xml.Name{Local: "preferredReportName"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.QuantitativeDetails, xml.StartElement{Name: xml.Name{Local: "quantitativeDetails"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.QualifiedInterval, xml.StartElement{Name: xml.Name{Local: "qualifiedInterval"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ValidCodedValueSet, xml.StartElement{Name: xml.Name{Local: "validCodedValueSet"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NormalCodedValueSet, xml.StartElement{Name: xml.Name{Local: "normalCodedValueSet"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AbnormalCodedValueSet, xml.StartElement{Name: xml.Name{Local: "abnormalCodedValueSet"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CriticalCodedValueSet, xml.StartElement{Name: xml.Name{Local: "criticalCodedValueSet"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ObservationDefinitionQuantitativeDetails) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.CustomaryUnit, xml.StartElement{Name: xml.Name{Local: "customaryUnit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Unit, xml.StartElement{Name: xml.Name{Local: "unit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ConversionFactor, xml.StartElement{Name: xml.Name{Local: "conversionFactor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DecimalPrecision, xml.StartElement{Name: xml.Name{Local: "decimalPrecision"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r ObservationDefinitionQualifiedInterval) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Category, xml.StartElement{Name: xml.Name{Local: "category"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Range, xml.StartElement{Name: xml.Name{Local: "range"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Context, xml.StartElement{Name: xml.Name{Local: "context"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AppliesTo, xml.StartElement{Name: xml.Name{Local: "appliesTo"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Gender, xml.StartElement{Name: xml.Name{Local: "gender"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Age, xml.StartElement{Name: xml.Name{Local: "age"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.GestationalAge, xml.StartElement{Name: xml.Name{Local: "gestationalAge"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Condition, xml.StartElement{Name: xml.Name{Local: "condition"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ObservationDefinition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "category":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = append(r.Category, v)
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = v
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "permittedDataType":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PermittedDataType = append(r.PermittedDataType, v)
			case "multipleResultsAllowed":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MultipleResultsAllowed = &v
			case "method":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Method = &v
			case "preferredReportName":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PreferredReportName = &v
			case "quantitativeDetails":
				var v ObservationDefinitionQuantitativeDetails
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.QuantitativeDetails = &v
			case "qualifiedInterval":
				var v ObservationDefinitionQualifiedInterval
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.QualifiedInterval = append(r.QualifiedInterval, v)
			case "validCodedValueSet":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ValidCodedValueSet = &v
			case "normalCodedValueSet":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NormalCodedValueSet = &v
			case "abnormalCodedValueSet":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AbnormalCodedValueSet = &v
			case "criticalCodedValueSet":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CriticalCodedValueSet = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ObservationDefinitionQuantitativeDetails) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "customaryUnit":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CustomaryUnit = &v
			case "unit":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Unit = &v
			case "conversionFactor":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ConversionFactor = &v
			case "decimalPrecision":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DecimalPrecision = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *ObservationDefinitionQualifiedInterval) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "category":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = &v
			case "range":
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Range = &v
			case "context":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Context = &v
			case "appliesTo":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AppliesTo = append(r.AppliesTo, v)
			case "gender":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Gender = &v
			case "age":
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Age = &v
			case "gestationalAge":
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.GestationalAge = &v
			case "condition":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Condition = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ObservationDefinition) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "category") {
		for _, v := range r.Category {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "code") {
		children = append(children, r.Code)
	}
	if len(name) == 0 || slices.Contains(name, "identifier") {
		for _, v := range r.Identifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "permittedDataType") {
		for _, v := range r.PermittedDataType {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "multipleResultsAllowed") {
		if r.MultipleResultsAllowed != nil {
			children = append(children, *r.MultipleResultsAllowed)
		}
	}
	if len(name) == 0 || slices.Contains(name, "method") {
		if r.Method != nil {
			children = append(children, *r.Method)
		}
	}
	if len(name) == 0 || slices.Contains(name, "preferredReportName") {
		if r.PreferredReportName != nil {
			children = append(children, *r.PreferredReportName)
		}
	}
	if len(name) == 0 || slices.Contains(name, "quantitativeDetails") {
		if r.QuantitativeDetails != nil {
			children = append(children, *r.QuantitativeDetails)
		}
	}
	if len(name) == 0 || slices.Contains(name, "qualifiedInterval") {
		for _, v := range r.QualifiedInterval {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "validCodedValueSet") {
		if r.ValidCodedValueSet != nil {
			children = append(children, *r.ValidCodedValueSet)
		}
	}
	if len(name) == 0 || slices.Contains(name, "normalCodedValueSet") {
		if r.NormalCodedValueSet != nil {
			children = append(children, *r.NormalCodedValueSet)
		}
	}
	if len(name) == 0 || slices.Contains(name, "abnormalCodedValueSet") {
		if r.AbnormalCodedValueSet != nil {
			children = append(children, *r.AbnormalCodedValueSet)
		}
	}
	if len(name) == 0 || slices.Contains(name, "criticalCodedValueSet") {
		if r.CriticalCodedValueSet != nil {
			children = append(children, *r.CriticalCodedValueSet)
		}
	}
	return children
}
func (r ObservationDefinition) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ObservationDefinition to Boolean")
}
func (r ObservationDefinition) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ObservationDefinition to String")
}
func (r ObservationDefinition) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ObservationDefinition to Integer")
}
func (r ObservationDefinition) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ObservationDefinition to Decimal")
}
func (r ObservationDefinition) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ObservationDefinition to Date")
}
func (r ObservationDefinition) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ObservationDefinition to Time")
}
func (r ObservationDefinition) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ObservationDefinition to DateTime")
}
func (r ObservationDefinition) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ObservationDefinition to Quantity")
}
func (r ObservationDefinition) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ObservationDefinition
	switch other := other.(type) {
	case ObservationDefinition:
		o = &other
	case *ObservationDefinition:
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
func (r ObservationDefinition) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(ObservationDefinition)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r ObservationDefinition) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		BaseType: fhirpath.TypeSpecifier{
			Name:      "DomainResource",
			Namespace: "FHIR",
		},
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
			Name: "Category",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Code",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
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
			Name: "PermittedDataType",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "MultipleResultsAllowed",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Boolean",
				Namespace: "FHIR",
			},
		}, {
			Name: "Method",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "PreferredReportName",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}, {
			Name: "QuantitativeDetails",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "ObservationDefinitionQuantitativeDetails",
				Namespace: "FHIR",
			},
		}, {
			Name: "QualifiedInterval",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "ObservationDefinitionQualifiedInterval",
				Namespace: "FHIR",
			},
		}, {
			Name: "ValidCodedValueSet",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "NormalCodedValueSet",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "AbnormalCodedValueSet",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}, {
			Name: "CriticalCodedValueSet",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Reference",
				Namespace: "FHIR",
			},
		}},
		Name:      "ObservationDefinition",
		Namespace: "FHIR",
	}
}
func (r ObservationDefinitionQuantitativeDetails) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "customaryUnit") {
		if r.CustomaryUnit != nil {
			children = append(children, *r.CustomaryUnit)
		}
	}
	if len(name) == 0 || slices.Contains(name, "unit") {
		if r.Unit != nil {
			children = append(children, *r.Unit)
		}
	}
	if len(name) == 0 || slices.Contains(name, "conversionFactor") {
		if r.ConversionFactor != nil {
			children = append(children, *r.ConversionFactor)
		}
	}
	if len(name) == 0 || slices.Contains(name, "decimalPrecision") {
		if r.DecimalPrecision != nil {
			children = append(children, *r.DecimalPrecision)
		}
	}
	return children
}
func (r ObservationDefinitionQuantitativeDetails) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ObservationDefinitionQuantitativeDetails to Boolean")
}
func (r ObservationDefinitionQuantitativeDetails) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ObservationDefinitionQuantitativeDetails to String")
}
func (r ObservationDefinitionQuantitativeDetails) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ObservationDefinitionQuantitativeDetails to Integer")
}
func (r ObservationDefinitionQuantitativeDetails) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ObservationDefinitionQuantitativeDetails to Decimal")
}
func (r ObservationDefinitionQuantitativeDetails) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ObservationDefinitionQuantitativeDetails to Date")
}
func (r ObservationDefinitionQuantitativeDetails) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ObservationDefinitionQuantitativeDetails to Time")
}
func (r ObservationDefinitionQuantitativeDetails) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ObservationDefinitionQuantitativeDetails to DateTime")
}
func (r ObservationDefinitionQuantitativeDetails) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ObservationDefinitionQuantitativeDetails to Quantity")
}
func (r ObservationDefinitionQuantitativeDetails) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ObservationDefinitionQuantitativeDetails
	switch other := other.(type) {
	case ObservationDefinitionQuantitativeDetails:
		o = &other
	case *ObservationDefinitionQuantitativeDetails:
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
func (r ObservationDefinitionQuantitativeDetails) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(ObservationDefinitionQuantitativeDetails)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r ObservationDefinitionQuantitativeDetails) TypeInfo() fhirpath.TypeInfo {
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
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "CustomaryUnit",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Unit",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "ConversionFactor",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "DecimalPrecision",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Integer",
				Namespace: "FHIR",
			},
		}},
		Name:      "ObservationDefinitionQuantitativeDetails",
		Namespace: "FHIR",
	}
}
func (r ObservationDefinitionQualifiedInterval) Children(name ...string) fhirpath.Collection {
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
	if len(name) == 0 || slices.Contains(name, "category") {
		if r.Category != nil {
			children = append(children, *r.Category)
		}
	}
	if len(name) == 0 || slices.Contains(name, "range") {
		if r.Range != nil {
			children = append(children, *r.Range)
		}
	}
	if len(name) == 0 || slices.Contains(name, "context") {
		if r.Context != nil {
			children = append(children, *r.Context)
		}
	}
	if len(name) == 0 || slices.Contains(name, "appliesTo") {
		for _, v := range r.AppliesTo {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "gender") {
		if r.Gender != nil {
			children = append(children, *r.Gender)
		}
	}
	if len(name) == 0 || slices.Contains(name, "age") {
		if r.Age != nil {
			children = append(children, *r.Age)
		}
	}
	if len(name) == 0 || slices.Contains(name, "gestationalAge") {
		if r.GestationalAge != nil {
			children = append(children, *r.GestationalAge)
		}
	}
	if len(name) == 0 || slices.Contains(name, "condition") {
		if r.Condition != nil {
			children = append(children, *r.Condition)
		}
	}
	return children
}
func (r ObservationDefinitionQualifiedInterval) ToBoolean(explicit bool) (fhirpath.Boolean, bool, error) {
	return false, false, errors.New("can not convert ObservationDefinitionQualifiedInterval to Boolean")
}
func (r ObservationDefinitionQualifiedInterval) ToString(explicit bool) (fhirpath.String, bool, error) {
	return "", false, errors.New("can not convert ObservationDefinitionQualifiedInterval to String")
}
func (r ObservationDefinitionQualifiedInterval) ToInteger(explicit bool) (fhirpath.Integer, bool, error) {
	return 0, false, errors.New("can not convert ObservationDefinitionQualifiedInterval to Integer")
}
func (r ObservationDefinitionQualifiedInterval) ToDecimal(explicit bool) (fhirpath.Decimal, bool, error) {
	return fhirpath.Decimal{}, false, errors.New("can not convert ObservationDefinitionQualifiedInterval to Decimal")
}
func (r ObservationDefinitionQualifiedInterval) ToDate(explicit bool) (fhirpath.Date, bool, error) {
	return fhirpath.Date{}, false, errors.New("can not convert ObservationDefinitionQualifiedInterval to Date")
}
func (r ObservationDefinitionQualifiedInterval) ToTime(explicit bool) (fhirpath.Time, bool, error) {
	return fhirpath.Time{}, false, errors.New("can not convert ObservationDefinitionQualifiedInterval to Time")
}
func (r ObservationDefinitionQualifiedInterval) ToDateTime(explicit bool) (fhirpath.DateTime, bool, error) {
	return fhirpath.DateTime{}, false, errors.New("can not convert ObservationDefinitionQualifiedInterval to DateTime")
}
func (r ObservationDefinitionQualifiedInterval) ToQuantity(explicit bool) (fhirpath.Quantity, bool, error) {
	return fhirpath.Quantity{}, false, errors.New("can not convert ObservationDefinitionQualifiedInterval to Quantity")
}
func (r ObservationDefinitionQualifiedInterval) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (bool, bool) {
	var o *ObservationDefinitionQualifiedInterval
	switch other := other.(type) {
	case ObservationDefinitionQualifiedInterval:
		o = &other
	case *ObservationDefinitionQualifiedInterval:
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
func (r ObservationDefinitionQualifiedInterval) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	o, ok := other.(ObservationDefinitionQualifiedInterval)
	if !ok {
		return false
	}
	r.Id = nil
	o.Id = nil
	eq, ok := r.Equal(o)
	return eq && ok
}
func (r ObservationDefinitionQualifiedInterval) TypeInfo() fhirpath.TypeInfo {
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
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Category",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Range",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Range",
				Namespace: "FHIR",
			},
		}, {
			Name: "Context",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "AppliesTo",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}, {
			Name: "Gender",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Age",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Range",
				Namespace: "FHIR",
			},
		}, {
			Name: "GestationalAge",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Range",
				Namespace: "FHIR",
			},
		}, {
			Name: "Condition",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "String",
				Namespace: "FHIR",
			},
		}},
		Name:      "ObservationDefinitionQualifiedInterval",
		Namespace: "FHIR",
	}
}
