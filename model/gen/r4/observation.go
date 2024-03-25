package r4

import (
	"encoding/json"
	"fmt"
)

// Measurements and simple assertions made about a patient, device or other subject.
//
// Observations are a key aspect of healthcare.  This resource is used to capture those that do not require more sophisticated mechanisms.
type Observation struct {
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
	Contained []Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A unique identifier assigned to this observation.
	Identifier []Identifier
	// A plan, proposal or order that is fulfilled in whole or in part by this event.  For example, a MedicationRequest may require a patient to have laboratory test performed before  it is dispensed.
	BasedOn []Reference
	// A larger event of which this particular Observation is a component or step.  For example,  an observation as part of a procedure.
	PartOf []Reference
	// The status of the result value.
	Status Code
	// A code that classifies the general type of observation being made.
	Category []CodeableConcept
	// Describes what was observed. Sometimes this is called the observation "name".
	Code CodeableConcept
	// The patient, or group of patients, location, or device this observation is about and into whose record the observation is placed. If the actual focus of the observation is different from the subject (or a sample of, part, or region of the subject), the `focus` element or the `code` itself specifies the actual focus of the observation.
	Subject *Reference
	// The actual focus of an observation when it is not the patient of record representing something or someone associated with the patient such as a spouse, parent, fetus, or donor. For example, fetus observations in a mother's record.  The focus of an observation could also be an existing condition,  an intervention, the subject's diet,  another observation of the subject,  or a body structure such as tumor or implanted device.   An example use case would be using the Observation resource to capture whether the mother is trained to change her child's tracheostomy tube. In this example, the child is the patient of record and the mother is the focus.
	Focus []Reference
	// The healthcare event  (e.g. a patient and healthcare provider interaction) during which this observation is made.
	Encounter *Reference
	// The time or time-period the observed value is asserted as being true. For biological subjects - e.g. human patients - this is usually called the "physiologically relevant time". This is usually either the time of the procedure or of specimen collection, but very often the source of the date/time is not known, only the date/time itself.
	Effective isObservationEffective
	// The date and time this version of the observation was made available to providers, typically after the results have been reviewed and verified.
	Issued *Instant
	// Who was responsible for asserting the observed value as "true".
	Performer []Reference
	// The information determined as a result of making the observation, if the information has a simple value.
	Value isObservationValue
	// Provides a reason why the expected value in the element Observation.value[x] is missing.
	DataAbsentReason *CodeableConcept
	// A categorical assessment of an observation value.  For example, high, low, normal.
	Interpretation []CodeableConcept
	// Comments about the observation or the results.
	Note []Annotation
	// Indicates the site on the subject's body where the observation was made (i.e. the target site).
	BodySite *CodeableConcept
	// Indicates the mechanism used to perform the observation.
	Method *CodeableConcept
	// The specimen that was used when this observation was made.
	Specimen *Reference
	// The device used to generate the observation data.
	Device *Reference
	// Guidance on how to interpret the value by comparison to a normal or recommended range.  Multiple reference ranges are interpreted as an "OR".   In other words, to represent two distinct target populations, two `referenceRange` elements would be used.
	ReferenceRange []ObservationReferenceRange
	// This observation is a group observation (e.g. a battery, a panel of tests, a set of vital sign measurements) that includes the target as a member of the group.
	HasMember []Reference
	// The target resource that represents a measurement from which this observation value is derived. For example, a calculated anion gap or a fetal measurement based on an ultrasound image.
	DerivedFrom []Reference
	// Some observations have multiple component observations.  These component observations are expressed as separate code value pairs that share the same attributes.  Examples include systolic and diastolic component observations for blood pressure measurement and multiple component observations for genetics observations.
	Component []ObservationComponent
}
type isObservationEffective interface {
	isObservationEffective()
}

func (r DateTime) isObservationEffective() {}
func (r Period) isObservationEffective()   {}
func (r Timing) isObservationEffective()   {}
func (r Instant) isObservationEffective()  {}

type isObservationValue interface {
	isObservationValue()
}

func (r Quantity) isObservationValue()        {}
func (r CodeableConcept) isObservationValue() {}
func (r String) isObservationValue()          {}
func (r Boolean) isObservationValue()         {}
func (r Integer) isObservationValue()         {}
func (r Range) isObservationValue()           {}
func (r Ratio) isObservationValue()           {}
func (r SampledData) isObservationValue()     {}
func (r Time) isObservationValue()            {}
func (r DateTime) isObservationValue()        {}
func (r Period) isObservationValue()          {}

type jsonObservation struct {
	ResourceType                      string                      `json:"resourceType"`
	Id                                *Id                         `json:"id,omitempty"`
	IdPrimitiveElement                *primitiveElement           `json:"_id,omitempty"`
	Meta                              *Meta                       `json:"meta,omitempty"`
	ImplicitRules                     *Uri                        `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement     *primitiveElement           `json:"_implicitRules,omitempty"`
	Language                          *Code                       `json:"language,omitempty"`
	LanguagePrimitiveElement          *primitiveElement           `json:"_language,omitempty"`
	Text                              *Narrative                  `json:"text,omitempty"`
	Contained                         []containedResource         `json:"contained,omitempty"`
	Extension                         []Extension                 `json:"extension,omitempty"`
	ModifierExtension                 []Extension                 `json:"modifierExtension,omitempty"`
	Identifier                        []Identifier                `json:"identifier,omitempty"`
	BasedOn                           []Reference                 `json:"basedOn,omitempty"`
	PartOf                            []Reference                 `json:"partOf,omitempty"`
	Status                            Code                        `json:"status,omitempty"`
	StatusPrimitiveElement            *primitiveElement           `json:"_status,omitempty"`
	Category                          []CodeableConcept           `json:"category,omitempty"`
	Code                              CodeableConcept             `json:"code,omitempty"`
	Subject                           *Reference                  `json:"subject,omitempty"`
	Focus                             []Reference                 `json:"focus,omitempty"`
	Encounter                         *Reference                  `json:"encounter,omitempty"`
	EffectiveDateTime                 *DateTime                   `json:"effectiveDateTime,omitempty"`
	EffectiveDateTimePrimitiveElement *primitiveElement           `json:"_effectiveDateTime,omitempty"`
	EffectivePeriod                   *Period                     `json:"effectivePeriod,omitempty"`
	EffectiveTiming                   *Timing                     `json:"effectiveTiming,omitempty"`
	EffectiveInstant                  *Instant                    `json:"effectiveInstant,omitempty"`
	EffectiveInstantPrimitiveElement  *primitiveElement           `json:"_effectiveInstant,omitempty"`
	Issued                            *Instant                    `json:"issued,omitempty"`
	IssuedPrimitiveElement            *primitiveElement           `json:"_issued,omitempty"`
	Performer                         []Reference                 `json:"performer,omitempty"`
	ValueQuantity                     *Quantity                   `json:"valueQuantity,omitempty"`
	ValueCodeableConcept              *CodeableConcept            `json:"valueCodeableConcept,omitempty"`
	ValueString                       *String                     `json:"valueString,omitempty"`
	ValueStringPrimitiveElement       *primitiveElement           `json:"_valueString,omitempty"`
	ValueBoolean                      *Boolean                    `json:"valueBoolean,omitempty"`
	ValueBooleanPrimitiveElement      *primitiveElement           `json:"_valueBoolean,omitempty"`
	ValueInteger                      *Integer                    `json:"valueInteger,omitempty"`
	ValueIntegerPrimitiveElement      *primitiveElement           `json:"_valueInteger,omitempty"`
	ValueRange                        *Range                      `json:"valueRange,omitempty"`
	ValueRatio                        *Ratio                      `json:"valueRatio,omitempty"`
	ValueSampledData                  *SampledData                `json:"valueSampledData,omitempty"`
	ValueTime                         *Time                       `json:"valueTime,omitempty"`
	ValueTimePrimitiveElement         *primitiveElement           `json:"_valueTime,omitempty"`
	ValueDateTime                     *DateTime                   `json:"valueDateTime,omitempty"`
	ValueDateTimePrimitiveElement     *primitiveElement           `json:"_valueDateTime,omitempty"`
	ValuePeriod                       *Period                     `json:"valuePeriod,omitempty"`
	DataAbsentReason                  *CodeableConcept            `json:"dataAbsentReason,omitempty"`
	Interpretation                    []CodeableConcept           `json:"interpretation,omitempty"`
	Note                              []Annotation                `json:"note,omitempty"`
	BodySite                          *CodeableConcept            `json:"bodySite,omitempty"`
	Method                            *CodeableConcept            `json:"method,omitempty"`
	Specimen                          *Reference                  `json:"specimen,omitempty"`
	Device                            *Reference                  `json:"device,omitempty"`
	ReferenceRange                    []ObservationReferenceRange `json:"referenceRange,omitempty"`
	HasMember                         []Reference                 `json:"hasMember,omitempty"`
	DerivedFrom                       []Reference                 `json:"derivedFrom,omitempty"`
	Component                         []ObservationComponent      `json:"component,omitempty"`
}

func (r Observation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Observation) marshalJSON() jsonObservation {
	m := jsonObservation{}
	m.ResourceType = "Observation"
	m.Id = r.Id
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	m.ImplicitRules = r.ImplicitRules
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	m.Language = r.Language
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.Text = r.Text
	m.Contained = make([]containedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, containedResource{resource: c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.BasedOn = r.BasedOn
	m.PartOf = r.PartOf
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Category = r.Category
	m.Code = r.Code
	m.Subject = r.Subject
	m.Focus = r.Focus
	m.Encounter = r.Encounter
	switch v := r.Effective.(type) {
	case DateTime:
		m.EffectiveDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.EffectiveDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.EffectiveDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.EffectiveDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.EffectivePeriod = &v
	case *Period:
		m.EffectivePeriod = v
	case Timing:
		m.EffectiveTiming = &v
	case *Timing:
		m.EffectiveTiming = v
	case Instant:
		m.EffectiveInstant = &v
		if v.Id != nil || v.Extension != nil {
			m.EffectiveInstantPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Instant:
		m.EffectiveInstant = v
		if v.Id != nil || v.Extension != nil {
			m.EffectiveInstantPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.Issued = r.Issued
	if r.Issued != nil && (r.Issued.Id != nil || r.Issued.Extension != nil) {
		m.IssuedPrimitiveElement = &primitiveElement{Id: r.Issued.Id, Extension: r.Issued.Extension}
	}
	m.Performer = r.Performer
	switch v := r.Value.(type) {
	case Quantity:
		m.ValueQuantity = &v
	case *Quantity:
		m.ValueQuantity = v
	case CodeableConcept:
		m.ValueCodeableConcept = &v
	case *CodeableConcept:
		m.ValueCodeableConcept = v
	case String:
		m.ValueString = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.ValueString = v
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Boolean:
		m.ValueBoolean = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		m.ValueBoolean = v
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Integer:
		m.ValueInteger = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Integer:
		m.ValueInteger = v
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Range:
		m.ValueRange = &v
	case *Range:
		m.ValueRange = v
	case Ratio:
		m.ValueRatio = &v
	case *Ratio:
		m.ValueRatio = v
	case SampledData:
		m.ValueSampledData = &v
	case *SampledData:
		m.ValueSampledData = v
	case Time:
		m.ValueTime = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Time:
		m.ValueTime = v
		if v.Id != nil || v.Extension != nil {
			m.ValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case DateTime:
		m.ValueDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.ValueDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.ValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.ValuePeriod = &v
	case *Period:
		m.ValuePeriod = v
	}
	m.DataAbsentReason = r.DataAbsentReason
	m.Interpretation = r.Interpretation
	m.Note = r.Note
	m.BodySite = r.BodySite
	m.Method = r.Method
	m.Specimen = r.Specimen
	m.Device = r.Device
	m.ReferenceRange = r.ReferenceRange
	m.HasMember = r.HasMember
	m.DerivedFrom = r.DerivedFrom
	m.Component = r.Component
	return m
}
func (r *Observation) UnmarshalJSON(b []byte) error {
	var m jsonObservation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Observation) unmarshalJSON(m jsonObservation) error {
	r.Id = m.Id
	if m.IdPrimitiveElement != nil {
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Text = m.Text
	r.Contained = make([]Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.BasedOn = m.BasedOn
	r.PartOf = m.PartOf
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Category = m.Category
	r.Code = m.Code
	r.Subject = m.Subject
	r.Focus = m.Focus
	r.Encounter = m.Encounter
	if m.EffectiveDateTime != nil || m.EffectiveDateTimePrimitiveElement != nil {
		if r.Effective != nil {
			return fmt.Errorf("multiple values for field \"Effective\"")
		}
		v := m.EffectiveDateTime
		if m.EffectiveDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.EffectiveDateTimePrimitiveElement.Id
			v.Extension = m.EffectiveDateTimePrimitiveElement.Extension
		}
		r.Effective = v
	}
	if m.EffectivePeriod != nil {
		if r.Effective != nil {
			return fmt.Errorf("multiple values for field \"Effective\"")
		}
		v := m.EffectivePeriod
		r.Effective = v
	}
	if m.EffectiveTiming != nil {
		if r.Effective != nil {
			return fmt.Errorf("multiple values for field \"Effective\"")
		}
		v := m.EffectiveTiming
		r.Effective = v
	}
	if m.EffectiveInstant != nil || m.EffectiveInstantPrimitiveElement != nil {
		if r.Effective != nil {
			return fmt.Errorf("multiple values for field \"Effective\"")
		}
		v := m.EffectiveInstant
		if m.EffectiveInstantPrimitiveElement != nil {
			if v == nil {
				v = &Instant{}
			}
			v.Id = m.EffectiveInstantPrimitiveElement.Id
			v.Extension = m.EffectiveInstantPrimitiveElement.Extension
		}
		r.Effective = v
	}
	r.Issued = m.Issued
	if m.IssuedPrimitiveElement != nil {
		r.Issued.Id = m.IssuedPrimitiveElement.Id
		r.Issued.Extension = m.IssuedPrimitiveElement.Extension
	}
	r.Performer = m.Performer
	if m.ValueQuantity != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueQuantity
		r.Value = v
	}
	if m.ValueCodeableConcept != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCodeableConcept
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
	if m.ValueSampledData != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueSampledData
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
	r.DataAbsentReason = m.DataAbsentReason
	r.Interpretation = m.Interpretation
	r.Note = m.Note
	r.BodySite = m.BodySite
	r.Method = m.Method
	r.Specimen = m.Specimen
	r.Device = m.Device
	r.ReferenceRange = m.ReferenceRange
	r.HasMember = m.HasMember
	r.DerivedFrom = m.DerivedFrom
	r.Component = m.Component
	return nil
}
func (r Observation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Guidance on how to interpret the value by comparison to a normal or recommended range.  Multiple reference ranges are interpreted as an "OR".   In other words, to represent two distinct target populations, two `referenceRange` elements would be used.
type ObservationReferenceRange struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The value of the low bound of the reference range.  The low bound of the reference range endpoint is inclusive of the value (e.g.  reference range is >=5 - <=9). If the low bound is omitted,  it is assumed to be meaningless (e.g. reference range is <=2.3).
	Low *Quantity
	// The value of the high bound of the reference range.  The high bound of the reference range endpoint is inclusive of the value (e.g.  reference range is >=5 - <=9). If the high bound is omitted,  it is assumed to be meaningless (e.g. reference range is >= 2.3).
	High *Quantity
	// Codes to indicate the what part of the targeted reference population it applies to. For example, the normal or therapeutic range.
	Type *CodeableConcept
	// Codes to indicate the target population this reference range applies to.  For example, a reference range may be based on the normal population or a particular sex or race.  Multiple `appliesTo`  are interpreted as an "AND" of the target populations.  For example, to represent a target population of African American females, both a code of female and a code for African American would be used.
	AppliesTo []CodeableConcept
	// The age at which this reference range is applicable. This is a neonatal age (e.g. number of weeks at term) if the meaning says so.
	Age *Range
	// Text based reference range in an observation which may be used when a quantitative range is not appropriate for an observation.  An example would be a reference value of "Negative" or a list or table of "normals".
	Text *String
}
type jsonObservationReferenceRange struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Low                  *Quantity         `json:"low,omitempty"`
	High                 *Quantity         `json:"high,omitempty"`
	Type                 *CodeableConcept  `json:"type,omitempty"`
	AppliesTo            []CodeableConcept `json:"appliesTo,omitempty"`
	Age                  *Range            `json:"age,omitempty"`
	Text                 *String           `json:"text,omitempty"`
	TextPrimitiveElement *primitiveElement `json:"_text,omitempty"`
}

func (r ObservationReferenceRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ObservationReferenceRange) marshalJSON() jsonObservationReferenceRange {
	m := jsonObservationReferenceRange{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Low = r.Low
	m.High = r.High
	m.Type = r.Type
	m.AppliesTo = r.AppliesTo
	m.Age = r.Age
	m.Text = r.Text
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	return m
}
func (r *ObservationReferenceRange) UnmarshalJSON(b []byte) error {
	var m jsonObservationReferenceRange
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ObservationReferenceRange) unmarshalJSON(m jsonObservationReferenceRange) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Low = m.Low
	r.High = m.High
	r.Type = m.Type
	r.AppliesTo = m.AppliesTo
	r.Age = m.Age
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	return nil
}
func (r ObservationReferenceRange) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Some observations have multiple component observations.  These component observations are expressed as separate code value pairs that share the same attributes.  Examples include systolic and diastolic component observations for blood pressure measurement and multiple component observations for genetics observations.
type ObservationComponent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Describes what was observed. Sometimes this is called the observation "code".
	Code CodeableConcept
	// The information determined as a result of making the observation, if the information has a simple value.
	Value isObservationComponentValue
	// Provides a reason why the expected value in the element Observation.component.value[x] is missing.
	DataAbsentReason *CodeableConcept
	// A categorical assessment of an observation value.  For example, high, low, normal.
	Interpretation []CodeableConcept
	// Guidance on how to interpret the value by comparison to a normal or recommended range.
	ReferenceRange []ObservationReferenceRange
}
type isObservationComponentValue interface {
	isObservationComponentValue()
}

func (r Quantity) isObservationComponentValue()        {}
func (r CodeableConcept) isObservationComponentValue() {}
func (r String) isObservationComponentValue()          {}
func (r Boolean) isObservationComponentValue()         {}
func (r Integer) isObservationComponentValue()         {}
func (r Range) isObservationComponentValue()           {}
func (r Ratio) isObservationComponentValue()           {}
func (r SampledData) isObservationComponentValue()     {}
func (r Time) isObservationComponentValue()            {}
func (r DateTime) isObservationComponentValue()        {}
func (r Period) isObservationComponentValue()          {}

type jsonObservationComponent struct {
	Id                            *string                     `json:"id,omitempty"`
	Extension                     []Extension                 `json:"extension,omitempty"`
	ModifierExtension             []Extension                 `json:"modifierExtension,omitempty"`
	Code                          CodeableConcept             `json:"code,omitempty"`
	ValueQuantity                 *Quantity                   `json:"valueQuantity,omitempty"`
	ValueCodeableConcept          *CodeableConcept            `json:"valueCodeableConcept,omitempty"`
	ValueString                   *String                     `json:"valueString,omitempty"`
	ValueStringPrimitiveElement   *primitiveElement           `json:"_valueString,omitempty"`
	ValueBoolean                  *Boolean                    `json:"valueBoolean,omitempty"`
	ValueBooleanPrimitiveElement  *primitiveElement           `json:"_valueBoolean,omitempty"`
	ValueInteger                  *Integer                    `json:"valueInteger,omitempty"`
	ValueIntegerPrimitiveElement  *primitiveElement           `json:"_valueInteger,omitempty"`
	ValueRange                    *Range                      `json:"valueRange,omitempty"`
	ValueRatio                    *Ratio                      `json:"valueRatio,omitempty"`
	ValueSampledData              *SampledData                `json:"valueSampledData,omitempty"`
	ValueTime                     *Time                       `json:"valueTime,omitempty"`
	ValueTimePrimitiveElement     *primitiveElement           `json:"_valueTime,omitempty"`
	ValueDateTime                 *DateTime                   `json:"valueDateTime,omitempty"`
	ValueDateTimePrimitiveElement *primitiveElement           `json:"_valueDateTime,omitempty"`
	ValuePeriod                   *Period                     `json:"valuePeriod,omitempty"`
	DataAbsentReason              *CodeableConcept            `json:"dataAbsentReason,omitempty"`
	Interpretation                []CodeableConcept           `json:"interpretation,omitempty"`
	ReferenceRange                []ObservationReferenceRange `json:"referenceRange,omitempty"`
}

func (r ObservationComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ObservationComponent) marshalJSON() jsonObservationComponent {
	m := jsonObservationComponent{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	switch v := r.Value.(type) {
	case Quantity:
		m.ValueQuantity = &v
	case *Quantity:
		m.ValueQuantity = v
	case CodeableConcept:
		m.ValueCodeableConcept = &v
	case *CodeableConcept:
		m.ValueCodeableConcept = v
	case String:
		m.ValueString = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.ValueString = v
		if v.Id != nil || v.Extension != nil {
			m.ValueStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Boolean:
		m.ValueBoolean = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		m.ValueBoolean = v
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Integer:
		m.ValueInteger = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Integer:
		m.ValueInteger = v
		if v.Id != nil || v.Extension != nil {
			m.ValueIntegerPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Range:
		m.ValueRange = &v
	case *Range:
		m.ValueRange = v
	case Ratio:
		m.ValueRatio = &v
	case *Ratio:
		m.ValueRatio = v
	case SampledData:
		m.ValueSampledData = &v
	case *SampledData:
		m.ValueSampledData = v
	case Time:
		m.ValueTime = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Time:
		m.ValueTime = v
		if v.Id != nil || v.Extension != nil {
			m.ValueTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case DateTime:
		m.ValueDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.ValueDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.ValueDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.ValuePeriod = &v
	case *Period:
		m.ValuePeriod = v
	}
	m.DataAbsentReason = r.DataAbsentReason
	m.Interpretation = r.Interpretation
	m.ReferenceRange = r.ReferenceRange
	return m
}
func (r *ObservationComponent) UnmarshalJSON(b []byte) error {
	var m jsonObservationComponent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ObservationComponent) unmarshalJSON(m jsonObservationComponent) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.ValueQuantity != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueQuantity
		r.Value = v
	}
	if m.ValueCodeableConcept != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCodeableConcept
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
	if m.ValueSampledData != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueSampledData
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
	r.DataAbsentReason = m.DataAbsentReason
	r.Interpretation = m.Interpretation
	r.ReferenceRange = m.ReferenceRange
	return nil
}
func (r ObservationComponent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
