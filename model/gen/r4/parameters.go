package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// This resource is a non-persisted resource used to pass information into and back from an [operation](operations.html). It has no other use, and there is no RESTful endpoint associated with it.
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

type jsonParameters struct {
	ResourceType                  string                `json:"resourceType"`
	Id                            *Id                   `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement     `json:"_id,omitempty"`
	Meta                          *Meta                 `json:"meta,omitempty"`
	ImplicitRules                 *Uri                  `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement     `json:"_implicitRules,omitempty"`
	Language                      *Code                 `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement     `json:"_language,omitempty"`
	Parameter                     []ParametersParameter `json:"parameter,omitempty"`
}

func (r Parameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Parameters) marshalJSON() jsonParameters {
	m := jsonParameters{}
	m.ResourceType = "Parameters"
	if r.Id != nil && r.Id.Value != nil {
		m.Id = r.Id
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		m.ImplicitRules = r.ImplicitRules
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	if r.Language != nil && r.Language.Value != nil {
		m.Language = r.Language
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.Parameter = r.Parameter
	return m
}
func (r *Parameters) UnmarshalJSON(b []byte) error {
	var m jsonParameters
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Parameters) unmarshalJSON(m jsonParameters) error {
	r.Id = m.Id
	if m.IdPrimitiveElement != nil {
		if r.Id == nil {
			r.Id = &Id{}
		}
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		if r.ImplicitRules == nil {
			r.ImplicitRules = &Uri{}
		}
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		if r.Language == nil {
			r.Language = &Code{}
		}
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Parameter = m.Parameter
	return nil
}
func (r Parameters) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A parameter passed to or received from the operation.
type ParametersParameter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The name of the parameter (reference to the operation definition).
	Name String
	// If the parameter is a data type.
	Value isParametersParameterValue
	// If the parameter is a whole resource.
	Resource *model.Resource
	// A named part of a multi-part parameter.
	Part []ParametersParameter
}
type isParametersParameterValue interface {
	isParametersParameterValue()
}

func (r Base64Binary) isParametersParameterValue()        {}
func (r Boolean) isParametersParameterValue()             {}
func (r Canonical) isParametersParameterValue()           {}
func (r Code) isParametersParameterValue()                {}
func (r Date) isParametersParameterValue()                {}
func (r DateTime) isParametersParameterValue()            {}
func (r Decimal) isParametersParameterValue()             {}
func (r Id) isParametersParameterValue()                  {}
func (r Instant) isParametersParameterValue()             {}
func (r Integer) isParametersParameterValue()             {}
func (r Markdown) isParametersParameterValue()            {}
func (r Oid) isParametersParameterValue()                 {}
func (r PositiveInt) isParametersParameterValue()         {}
func (r String) isParametersParameterValue()              {}
func (r Time) isParametersParameterValue()                {}
func (r UnsignedInt) isParametersParameterValue()         {}
func (r Uri) isParametersParameterValue()                 {}
func (r Url) isParametersParameterValue()                 {}
func (r Uuid) isParametersParameterValue()                {}
func (r Address) isParametersParameterValue()             {}
func (r Age) isParametersParameterValue()                 {}
func (r Annotation) isParametersParameterValue()          {}
func (r Attachment) isParametersParameterValue()          {}
func (r CodeableConcept) isParametersParameterValue()     {}
func (r Coding) isParametersParameterValue()              {}
func (r ContactPoint) isParametersParameterValue()        {}
func (r Count) isParametersParameterValue()               {}
func (r Distance) isParametersParameterValue()            {}
func (r Duration) isParametersParameterValue()            {}
func (r HumanName) isParametersParameterValue()           {}
func (r Identifier) isParametersParameterValue()          {}
func (r Money) isParametersParameterValue()               {}
func (r Period) isParametersParameterValue()              {}
func (r Quantity) isParametersParameterValue()            {}
func (r Range) isParametersParameterValue()               {}
func (r Ratio) isParametersParameterValue()               {}
func (r Reference) isParametersParameterValue()           {}
func (r SampledData) isParametersParameterValue()         {}
func (r Signature) isParametersParameterValue()           {}
func (r Timing) isParametersParameterValue()              {}
func (r ContactDetail) isParametersParameterValue()       {}
func (r Contributor) isParametersParameterValue()         {}
func (r DataRequirement) isParametersParameterValue()     {}
func (r Expression) isParametersParameterValue()          {}
func (r ParameterDefinition) isParametersParameterValue() {}
func (r RelatedArtifact) isParametersParameterValue()     {}
func (r TriggerDefinition) isParametersParameterValue()   {}
func (r UsageContext) isParametersParameterValue()        {}
func (r Dosage) isParametersParameterValue()              {}
func (r Meta) isParametersParameterValue()                {}

type jsonParametersParameter struct {
	Id                                *string               `json:"id,omitempty"`
	Extension                         []Extension           `json:"extension,omitempty"`
	ModifierExtension                 []Extension           `json:"modifierExtension,omitempty"`
	Name                              String                `json:"name,omitempty"`
	NamePrimitiveElement              *primitiveElement     `json:"_name,omitempty"`
	ValueBase64Binary                 *Base64Binary         `json:"valueBase64Binary,omitempty"`
	ValueBase64BinaryPrimitiveElement *primitiveElement     `json:"_valueBase64Binary,omitempty"`
	ValueBoolean                      *Boolean              `json:"valueBoolean,omitempty"`
	ValueBooleanPrimitiveElement      *primitiveElement     `json:"_valueBoolean,omitempty"`
	ValueCanonical                    *Canonical            `json:"valueCanonical,omitempty"`
	ValueCanonicalPrimitiveElement    *primitiveElement     `json:"_valueCanonical,omitempty"`
	ValueCode                         *Code                 `json:"valueCode,omitempty"`
	ValueCodePrimitiveElement         *primitiveElement     `json:"_valueCode,omitempty"`
	ValueDate                         *Date                 `json:"valueDate,omitempty"`
	ValueDatePrimitiveElement         *primitiveElement     `json:"_valueDate,omitempty"`
	ValueDateTime                     *DateTime             `json:"valueDateTime,omitempty"`
	ValueDateTimePrimitiveElement     *primitiveElement     `json:"_valueDateTime,omitempty"`
	ValueDecimal                      *Decimal              `json:"valueDecimal,omitempty"`
	ValueDecimalPrimitiveElement      *primitiveElement     `json:"_valueDecimal,omitempty"`
	ValueId                           *Id                   `json:"valueId,omitempty"`
	ValueIdPrimitiveElement           *primitiveElement     `json:"_valueId,omitempty"`
	ValueInstant                      *Instant              `json:"valueInstant,omitempty"`
	ValueInstantPrimitiveElement      *primitiveElement     `json:"_valueInstant,omitempty"`
	ValueInteger                      *Integer              `json:"valueInteger,omitempty"`
	ValueIntegerPrimitiveElement      *primitiveElement     `json:"_valueInteger,omitempty"`
	ValueMarkdown                     *Markdown             `json:"valueMarkdown,omitempty"`
	ValueMarkdownPrimitiveElement     *primitiveElement     `json:"_valueMarkdown,omitempty"`
	ValueOid                          *Oid                  `json:"valueOid,omitempty"`
	ValueOidPrimitiveElement          *primitiveElement     `json:"_valueOid,omitempty"`
	ValuePositiveInt                  *PositiveInt          `json:"valuePositiveInt,omitempty"`
	ValuePositiveIntPrimitiveElement  *primitiveElement     `json:"_valuePositiveInt,omitempty"`
	ValueString                       *String               `json:"valueString,omitempty"`
	ValueStringPrimitiveElement       *primitiveElement     `json:"_valueString,omitempty"`
	ValueTime                         *Time                 `json:"valueTime,omitempty"`
	ValueTimePrimitiveElement         *primitiveElement     `json:"_valueTime,omitempty"`
	ValueUnsignedInt                  *UnsignedInt          `json:"valueUnsignedInt,omitempty"`
	ValueUnsignedIntPrimitiveElement  *primitiveElement     `json:"_valueUnsignedInt,omitempty"`
	ValueUri                          *Uri                  `json:"valueUri,omitempty"`
	ValueUriPrimitiveElement          *primitiveElement     `json:"_valueUri,omitempty"`
	ValueUrl                          *Url                  `json:"valueUrl,omitempty"`
	ValueUrlPrimitiveElement          *primitiveElement     `json:"_valueUrl,omitempty"`
	ValueUuid                         *Uuid                 `json:"valueUuid,omitempty"`
	ValueUuidPrimitiveElement         *primitiveElement     `json:"_valueUuid,omitempty"`
	ValueAddress                      *Address              `json:"valueAddress,omitempty"`
	ValueAge                          *Age                  `json:"valueAge,omitempty"`
	ValueAnnotation                   *Annotation           `json:"valueAnnotation,omitempty"`
	ValueAttachment                   *Attachment           `json:"valueAttachment,omitempty"`
	ValueCodeableConcept              *CodeableConcept      `json:"valueCodeableConcept,omitempty"`
	ValueCoding                       *Coding               `json:"valueCoding,omitempty"`
	ValueContactPoint                 *ContactPoint         `json:"valueContactPoint,omitempty"`
	ValueCount                        *Count                `json:"valueCount,omitempty"`
	ValueDistance                     *Distance             `json:"valueDistance,omitempty"`
	ValueDuration                     *Duration             `json:"valueDuration,omitempty"`
	ValueHumanName                    *HumanName            `json:"valueHumanName,omitempty"`
	ValueIdentifier                   *Identifier           `json:"valueIdentifier,omitempty"`
	ValueMoney                        *Money                `json:"valueMoney,omitempty"`
	ValuePeriod                       *Period               `json:"valuePeriod,omitempty"`
	ValueQuantity                     *Quantity             `json:"valueQuantity,omitempty"`
	ValueRange                        *Range                `json:"valueRange,omitempty"`
	ValueRatio                        *Ratio                `json:"valueRatio,omitempty"`
	ValueReference                    *Reference            `json:"valueReference,omitempty"`
	ValueSampledData                  *SampledData          `json:"valueSampledData,omitempty"`
	ValueSignature                    *Signature            `json:"valueSignature,omitempty"`
	ValueTiming                       *Timing               `json:"valueTiming,omitempty"`
	ValueContactDetail                *ContactDetail        `json:"valueContactDetail,omitempty"`
	ValueContributor                  *Contributor          `json:"valueContributor,omitempty"`
	ValueDataRequirement              *DataRequirement      `json:"valueDataRequirement,omitempty"`
	ValueExpression                   *Expression           `json:"valueExpression,omitempty"`
	ValueParameterDefinition          *ParameterDefinition  `json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact              *RelatedArtifact      `json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition            *TriggerDefinition    `json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext                 *UsageContext         `json:"valueUsageContext,omitempty"`
	ValueDosage                       *Dosage               `json:"valueDosage,omitempty"`
	ValueMeta                         *Meta                 `json:"valueMeta,omitempty"`
	Resource                          *ContainedResource    `json:"resource,omitempty"`
	Part                              []ParametersParameter `json:"part,omitempty"`
}

func (r ParametersParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ParametersParameter) marshalJSON() jsonParametersParameter {
	m := jsonParametersParameter{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	switch v := r.Value.(type) {
	case Base64Binary:
		if v.Value != nil {
			m.ValueBase64Binary = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueBase64BinaryPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Base64Binary:
		if v.Value != nil {
			m.ValueBase64Binary = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueBase64BinaryPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Boolean:
		if v.Value != nil {
			m.ValueBoolean = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		if v.Value != nil {
			m.ValueBoolean = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Canonical:
		if v.Value != nil {
			m.ValueCanonical = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Canonical:
		if v.Value != nil {
			m.ValueCanonical = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Code:
		if v.Value != nil {
			m.ValueCode = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueCodePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Code:
		if v.Value != nil {
			m.ValueCode = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueCodePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Date:
		if v.Value != nil {
			m.ValueDate = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		if v.Value != nil {
			m.ValueDate = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
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
	case Decimal:
		if v.Value != nil {
			m.ValueDecimal = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Decimal:
		if v.Value != nil {
			m.ValueDecimal = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Id:
		if v.Value != nil {
			m.ValueId = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueIdPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Id:
		if v.Value != nil {
			m.ValueId = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueIdPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Instant:
		if v.Value != nil {
			m.ValueInstant = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueInstantPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Instant:
		if v.Value != nil {
			m.ValueInstant = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueInstantPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Integer:
		if v.Value != nil {
			m.ValueInteger = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Integer:
		if v.Value != nil {
			m.ValueInteger = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Markdown:
		if v.Value != nil {
			m.ValueMarkdown = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueMarkdownPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Markdown:
		if v.Value != nil {
			m.ValueMarkdown = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueMarkdownPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Oid:
		if v.Value != nil {
			m.ValueOid = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueOidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Oid:
		if v.Value != nil {
			m.ValueOid = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueOidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case PositiveInt:
		if v.Value != nil {
			m.ValuePositiveInt = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValuePositiveIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *PositiveInt:
		if v.Value != nil {
			m.ValuePositiveInt = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValuePositiveIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case String:
		if v.Value != nil {
			m.ValueString = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		if v.Value != nil {
			m.ValueString = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Time:
		if v.Value != nil {
			m.ValueTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Time:
		if v.Value != nil {
			m.ValueTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case UnsignedInt:
		if v.Value != nil {
			m.ValueUnsignedInt = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueUnsignedIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *UnsignedInt:
		if v.Value != nil {
			m.ValueUnsignedInt = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueUnsignedIntPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Uri:
		if v.Value != nil {
			m.ValueUri = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uri:
		if v.Value != nil {
			m.ValueUri = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Url:
		if v.Value != nil {
			m.ValueUrl = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueUrlPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Url:
		if v.Value != nil {
			m.ValueUrl = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueUrlPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Uuid:
		if v.Value != nil {
			m.ValueUuid = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueUuidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uuid:
		if v.Value != nil {
			m.ValueUuid = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ValueUuidPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Address:
		m.ValueAddress = &v
	case *Address:
		m.ValueAddress = v
	case Age:
		m.ValueAge = &v
	case *Age:
		m.ValueAge = v
	case Annotation:
		m.ValueAnnotation = &v
	case *Annotation:
		m.ValueAnnotation = v
	case Attachment:
		m.ValueAttachment = &v
	case *Attachment:
		m.ValueAttachment = v
	case CodeableConcept:
		m.ValueCodeableConcept = &v
	case *CodeableConcept:
		m.ValueCodeableConcept = v
	case Coding:
		m.ValueCoding = &v
	case *Coding:
		m.ValueCoding = v
	case ContactPoint:
		m.ValueContactPoint = &v
	case *ContactPoint:
		m.ValueContactPoint = v
	case Count:
		m.ValueCount = &v
	case *Count:
		m.ValueCount = v
	case Distance:
		m.ValueDistance = &v
	case *Distance:
		m.ValueDistance = v
	case Duration:
		m.ValueDuration = &v
	case *Duration:
		m.ValueDuration = v
	case HumanName:
		m.ValueHumanName = &v
	case *HumanName:
		m.ValueHumanName = v
	case Identifier:
		m.ValueIdentifier = &v
	case *Identifier:
		m.ValueIdentifier = v
	case Money:
		m.ValueMoney = &v
	case *Money:
		m.ValueMoney = v
	case Period:
		m.ValuePeriod = &v
	case *Period:
		m.ValuePeriod = v
	case Quantity:
		m.ValueQuantity = &v
	case *Quantity:
		m.ValueQuantity = v
	case Range:
		m.ValueRange = &v
	case *Range:
		m.ValueRange = v
	case Ratio:
		m.ValueRatio = &v
	case *Ratio:
		m.ValueRatio = v
	case Reference:
		m.ValueReference = &v
	case *Reference:
		m.ValueReference = v
	case SampledData:
		m.ValueSampledData = &v
	case *SampledData:
		m.ValueSampledData = v
	case Signature:
		m.ValueSignature = &v
	case *Signature:
		m.ValueSignature = v
	case Timing:
		m.ValueTiming = &v
	case *Timing:
		m.ValueTiming = v
	case ContactDetail:
		m.ValueContactDetail = &v
	case *ContactDetail:
		m.ValueContactDetail = v
	case Contributor:
		m.ValueContributor = &v
	case *Contributor:
		m.ValueContributor = v
	case DataRequirement:
		m.ValueDataRequirement = &v
	case *DataRequirement:
		m.ValueDataRequirement = v
	case Expression:
		m.ValueExpression = &v
	case *Expression:
		m.ValueExpression = v
	case ParameterDefinition:
		m.ValueParameterDefinition = &v
	case *ParameterDefinition:
		m.ValueParameterDefinition = v
	case RelatedArtifact:
		m.ValueRelatedArtifact = &v
	case *RelatedArtifact:
		m.ValueRelatedArtifact = v
	case TriggerDefinition:
		m.ValueTriggerDefinition = &v
	case *TriggerDefinition:
		m.ValueTriggerDefinition = v
	case UsageContext:
		m.ValueUsageContext = &v
	case *UsageContext:
		m.ValueUsageContext = v
	case Dosage:
		m.ValueDosage = &v
	case *Dosage:
		m.ValueDosage = v
	case Meta:
		m.ValueMeta = &v
	case *Meta:
		m.ValueMeta = v
	}
	if r.Resource != nil {
		m.Resource = &ContainedResource{*r.Resource}
	}
	m.Part = r.Part
	return m
}
func (r *ParametersParameter) UnmarshalJSON(b []byte) error {
	var m jsonParametersParameter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ParametersParameter) unmarshalJSON(m jsonParametersParameter) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	if m.ValueBase64Binary != nil || m.ValueBase64BinaryPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueBase64Binary
		if m.ValueBase64BinaryPrimitiveElement != nil {
			if v == nil {
				v = &Base64Binary{}
			}
			v.Id = m.ValueBase64BinaryPrimitiveElement.Id
			v.Extension = m.ValueBase64BinaryPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueBoolean != nil || m.ValueBooleanPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueBoolean
		if m.ValueBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.ValueBooleanPrimitiveElement.Id
			v.Extension = m.ValueBooleanPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueCanonical != nil || m.ValueCanonicalPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCanonical
		if m.ValueCanonicalPrimitiveElement != nil {
			if v == nil {
				v = &Canonical{}
			}
			v.Id = m.ValueCanonicalPrimitiveElement.Id
			v.Extension = m.ValueCanonicalPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueCode != nil || m.ValueCodePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCode
		if m.ValueCodePrimitiveElement != nil {
			if v == nil {
				v = &Code{}
			}
			v.Id = m.ValueCodePrimitiveElement.Id
			v.Extension = m.ValueCodePrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueDate != nil || m.ValueDatePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDate
		if m.ValueDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.ValueDatePrimitiveElement.Id
			v.Extension = m.ValueDatePrimitiveElement.Extension
		}
		r.Value = v
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
	if m.ValueDecimal != nil || m.ValueDecimalPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDecimal
		if m.ValueDecimalPrimitiveElement != nil {
			if v == nil {
				v = &Decimal{}
			}
			v.Id = m.ValueDecimalPrimitiveElement.Id
			v.Extension = m.ValueDecimalPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueId != nil || m.ValueIdPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueId
		if m.ValueIdPrimitiveElement != nil {
			if v == nil {
				v = &Id{}
			}
			v.Id = m.ValueIdPrimitiveElement.Id
			v.Extension = m.ValueIdPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueInstant != nil || m.ValueInstantPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueInstant
		if m.ValueInstantPrimitiveElement != nil {
			if v == nil {
				v = &Instant{}
			}
			v.Id = m.ValueInstantPrimitiveElement.Id
			v.Extension = m.ValueInstantPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueInteger != nil || m.ValueIntegerPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueInteger
		if m.ValueIntegerPrimitiveElement != nil {
			if v == nil {
				v = &Integer{}
			}
			v.Id = m.ValueIntegerPrimitiveElement.Id
			v.Extension = m.ValueIntegerPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueMarkdown != nil || m.ValueMarkdownPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueMarkdown
		if m.ValueMarkdownPrimitiveElement != nil {
			if v == nil {
				v = &Markdown{}
			}
			v.Id = m.ValueMarkdownPrimitiveElement.Id
			v.Extension = m.ValueMarkdownPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueOid != nil || m.ValueOidPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueOid
		if m.ValueOidPrimitiveElement != nil {
			if v == nil {
				v = &Oid{}
			}
			v.Id = m.ValueOidPrimitiveElement.Id
			v.Extension = m.ValueOidPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValuePositiveInt != nil || m.ValuePositiveIntPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValuePositiveInt
		if m.ValuePositiveIntPrimitiveElement != nil {
			if v == nil {
				v = &PositiveInt{}
			}
			v.Id = m.ValuePositiveIntPrimitiveElement.Id
			v.Extension = m.ValuePositiveIntPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueString != nil || m.ValueStringPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueString
		if m.ValueStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.ValueStringPrimitiveElement.Id
			v.Extension = m.ValueStringPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueTime != nil || m.ValueTimePrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueTime
		if m.ValueTimePrimitiveElement != nil {
			if v == nil {
				v = &Time{}
			}
			v.Id = m.ValueTimePrimitiveElement.Id
			v.Extension = m.ValueTimePrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueUnsignedInt != nil || m.ValueUnsignedIntPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUnsignedInt
		if m.ValueUnsignedIntPrimitiveElement != nil {
			if v == nil {
				v = &UnsignedInt{}
			}
			v.Id = m.ValueUnsignedIntPrimitiveElement.Id
			v.Extension = m.ValueUnsignedIntPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueUri != nil || m.ValueUriPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUri
		if m.ValueUriPrimitiveElement != nil {
			if v == nil {
				v = &Uri{}
			}
			v.Id = m.ValueUriPrimitiveElement.Id
			v.Extension = m.ValueUriPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueUrl != nil || m.ValueUrlPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUrl
		if m.ValueUrlPrimitiveElement != nil {
			if v == nil {
				v = &Url{}
			}
			v.Id = m.ValueUrlPrimitiveElement.Id
			v.Extension = m.ValueUrlPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueUuid != nil || m.ValueUuidPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUuid
		if m.ValueUuidPrimitiveElement != nil {
			if v == nil {
				v = &Uuid{}
			}
			v.Id = m.ValueUuidPrimitiveElement.Id
			v.Extension = m.ValueUuidPrimitiveElement.Extension
		}
		r.Value = v
	}
	if m.ValueAddress != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAddress
		r.Value = v
	}
	if m.ValueAge != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAge
		r.Value = v
	}
	if m.ValueAnnotation != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAnnotation
		r.Value = v
	}
	if m.ValueAttachment != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueAttachment
		r.Value = v
	}
	if m.ValueCodeableConcept != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCodeableConcept
		r.Value = v
	}
	if m.ValueCoding != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCoding
		r.Value = v
	}
	if m.ValueContactPoint != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueContactPoint
		r.Value = v
	}
	if m.ValueCount != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCount
		r.Value = v
	}
	if m.ValueDistance != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDistance
		r.Value = v
	}
	if m.ValueDuration != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDuration
		r.Value = v
	}
	if m.ValueHumanName != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueHumanName
		r.Value = v
	}
	if m.ValueIdentifier != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueIdentifier
		r.Value = v
	}
	if m.ValueMoney != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueMoney
		r.Value = v
	}
	if m.ValuePeriod != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValuePeriod
		r.Value = v
	}
	if m.ValueQuantity != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueQuantity
		r.Value = v
	}
	if m.ValueRange != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueRange
		r.Value = v
	}
	if m.ValueRatio != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueRatio
		r.Value = v
	}
	if m.ValueReference != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueReference
		r.Value = v
	}
	if m.ValueSampledData != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueSampledData
		r.Value = v
	}
	if m.ValueSignature != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueSignature
		r.Value = v
	}
	if m.ValueTiming != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueTiming
		r.Value = v
	}
	if m.ValueContactDetail != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueContactDetail
		r.Value = v
	}
	if m.ValueContributor != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueContributor
		r.Value = v
	}
	if m.ValueDataRequirement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDataRequirement
		r.Value = v
	}
	if m.ValueExpression != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueExpression
		r.Value = v
	}
	if m.ValueParameterDefinition != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueParameterDefinition
		r.Value = v
	}
	if m.ValueRelatedArtifact != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueRelatedArtifact
		r.Value = v
	}
	if m.ValueTriggerDefinition != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueTriggerDefinition
		r.Value = v
	}
	if m.ValueUsageContext != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueUsageContext
		r.Value = v
	}
	if m.ValueDosage != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueDosage
		r.Value = v
	}
	if m.ValueMeta != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueMeta
		r.Value = v
	}
	if m.Resource != nil {
		r.Resource = &m.Resource.Resource
	}
	r.Part = m.Part
	return nil
}
func (r ParametersParameter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
