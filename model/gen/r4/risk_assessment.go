package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// An assessment of the likely outcome(s) for a patient or other subject as well as the likelihood of each outcome.
type RiskAssessment struct {
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
	// Business identifier assigned to the risk assessment.
	Identifier []Identifier
	// A reference to the request that is fulfilled by this risk assessment.
	BasedOn *Reference
	// A reference to a resource that this risk assessment is part of, such as a Procedure.
	Parent *Reference
	// The status of the RiskAssessment, using the same statuses as an Observation.
	Status Code
	// The algorithm, process or mechanism used to evaluate the risk.
	Method *CodeableConcept
	// The type of the risk assessment performed.
	Code *CodeableConcept
	// The patient or group the risk assessment applies to.
	Subject Reference
	// The encounter where the assessment was performed.
	Encounter *Reference
	// The date (and possibly time) the risk assessment was performed.
	Occurrence isRiskAssessmentOccurrence
	// For assessments or prognosis specific to a particular condition, indicates the condition being assessed.
	Condition *Reference
	// The provider or software application that performed the assessment.
	Performer *Reference
	// The reason the risk assessment was performed.
	ReasonCode []CodeableConcept
	// Resources supporting the reason the risk assessment was performed.
	ReasonReference []Reference
	// Indicates the source data considered as part of the assessment (for example, FamilyHistory, Observations, Procedures, Conditions, etc.).
	Basis []Reference
	// Describes the expected outcome for the subject.
	Prediction []RiskAssessmentPrediction
	// A description of the steps that might be taken to reduce the identified risk(s).
	Mitigation *String
	// Additional comments about the risk assessment.
	Note []Annotation
}

func (r RiskAssessment) ResourceType() string {
	return "RiskAssessment"
}

type isRiskAssessmentOccurrence interface {
	isRiskAssessmentOccurrence()
}

func (r DateTime) isRiskAssessmentOccurrence() {}
func (r Period) isRiskAssessmentOccurrence()   {}

type jsonRiskAssessment struct {
	ResourceType                       string                     `json:"resourceType"`
	Id                                 *Id                        `json:"id,omitempty"`
	IdPrimitiveElement                 *primitiveElement          `json:"_id,omitempty"`
	Meta                               *Meta                      `json:"meta,omitempty"`
	ImplicitRules                      *Uri                       `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement      *primitiveElement          `json:"_implicitRules,omitempty"`
	Language                           *Code                      `json:"language,omitempty"`
	LanguagePrimitiveElement           *primitiveElement          `json:"_language,omitempty"`
	Text                               *Narrative                 `json:"text,omitempty"`
	Contained                          []containedResource        `json:"contained,omitempty"`
	Extension                          []Extension                `json:"extension,omitempty"`
	ModifierExtension                  []Extension                `json:"modifierExtension,omitempty"`
	Identifier                         []Identifier               `json:"identifier,omitempty"`
	BasedOn                            *Reference                 `json:"basedOn,omitempty"`
	Parent                             *Reference                 `json:"parent,omitempty"`
	Status                             Code                       `json:"status,omitempty"`
	StatusPrimitiveElement             *primitiveElement          `json:"_status,omitempty"`
	Method                             *CodeableConcept           `json:"method,omitempty"`
	Code                               *CodeableConcept           `json:"code,omitempty"`
	Subject                            Reference                  `json:"subject,omitempty"`
	Encounter                          *Reference                 `json:"encounter,omitempty"`
	OccurrenceDateTime                 *DateTime                  `json:"occurrenceDateTime,omitempty"`
	OccurrenceDateTimePrimitiveElement *primitiveElement          `json:"_occurrenceDateTime,omitempty"`
	OccurrencePeriod                   *Period                    `json:"occurrencePeriod,omitempty"`
	Condition                          *Reference                 `json:"condition,omitempty"`
	Performer                          *Reference                 `json:"performer,omitempty"`
	ReasonCode                         []CodeableConcept          `json:"reasonCode,omitempty"`
	ReasonReference                    []Reference                `json:"reasonReference,omitempty"`
	Basis                              []Reference                `json:"basis,omitempty"`
	Prediction                         []RiskAssessmentPrediction `json:"prediction,omitempty"`
	Mitigation                         *String                    `json:"mitigation,omitempty"`
	MitigationPrimitiveElement         *primitiveElement          `json:"_mitigation,omitempty"`
	Note                               []Annotation               `json:"note,omitempty"`
}

func (r RiskAssessment) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RiskAssessment) marshalJSON() jsonRiskAssessment {
	m := jsonRiskAssessment{}
	m.ResourceType = "RiskAssessment"
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
	m.Parent = r.Parent
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Method = r.Method
	m.Code = r.Code
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	switch v := r.Occurrence.(type) {
	case DateTime:
		m.OccurrenceDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.OccurrenceDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.OccurrenceDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.OccurrenceDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.OccurrencePeriod = &v
	case *Period:
		m.OccurrencePeriod = v
	}
	m.Condition = r.Condition
	m.Performer = r.Performer
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Basis = r.Basis
	m.Prediction = r.Prediction
	m.Mitigation = r.Mitigation
	if r.Mitigation != nil && (r.Mitigation.Id != nil || r.Mitigation.Extension != nil) {
		m.MitigationPrimitiveElement = &primitiveElement{Id: r.Mitigation.Id, Extension: r.Mitigation.Extension}
	}
	m.Note = r.Note
	return m
}
func (r *RiskAssessment) UnmarshalJSON(b []byte) error {
	var m jsonRiskAssessment
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RiskAssessment) unmarshalJSON(m jsonRiskAssessment) error {
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
	r.Contained = make([]model.Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.BasedOn = m.BasedOn
	r.Parent = m.Parent
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Method = m.Method
	r.Code = m.Code
	r.Subject = m.Subject
	r.Encounter = m.Encounter
	if m.OccurrenceDateTime != nil || m.OccurrenceDateTimePrimitiveElement != nil {
		if r.Occurrence != nil {
			return fmt.Errorf("multiple values for field \"Occurrence\"")
		}
		v := m.OccurrenceDateTime
		if m.OccurrenceDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.OccurrenceDateTimePrimitiveElement.Id
			v.Extension = m.OccurrenceDateTimePrimitiveElement.Extension
		}
		r.Occurrence = v
	}
	if m.OccurrencePeriod != nil {
		if r.Occurrence != nil {
			return fmt.Errorf("multiple values for field \"Occurrence\"")
		}
		v := m.OccurrencePeriod
		r.Occurrence = v
	}
	r.Condition = m.Condition
	r.Performer = m.Performer
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Basis = m.Basis
	r.Prediction = m.Prediction
	r.Mitigation = m.Mitigation
	if m.MitigationPrimitiveElement != nil {
		r.Mitigation.Id = m.MitigationPrimitiveElement.Id
		r.Mitigation.Extension = m.MitigationPrimitiveElement.Extension
	}
	r.Note = m.Note
	return nil
}
func (r RiskAssessment) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Describes the expected outcome for the subject.
type RiskAssessmentPrediction struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// One of the potential outcomes for the patient (e.g. remission, death,  a particular condition).
	Outcome *CodeableConcept
	// Indicates how likely the outcome is (in the specified timeframe).
	Probability isRiskAssessmentPredictionProbability
	// Indicates how likely the outcome is (in the specified timeframe), expressed as a qualitative value (e.g. low, medium, or high).
	QualitativeRisk *CodeableConcept
	// Indicates the risk for this particular subject (with their specific characteristics) divided by the risk of the population in general.  (Numbers greater than 1 = higher risk than the population, numbers less than 1 = lower risk.).
	RelativeRisk *Decimal
	// Indicates the period of time or age range of the subject to which the specified probability applies.
	When isRiskAssessmentPredictionWhen
	// Additional information explaining the basis for the prediction.
	Rationale *String
}
type isRiskAssessmentPredictionProbability interface {
	isRiskAssessmentPredictionProbability()
}

func (r Decimal) isRiskAssessmentPredictionProbability() {}
func (r Range) isRiskAssessmentPredictionProbability()   {}

type isRiskAssessmentPredictionWhen interface {
	isRiskAssessmentPredictionWhen()
}

func (r Period) isRiskAssessmentPredictionWhen() {}
func (r Range) isRiskAssessmentPredictionWhen()  {}

type jsonRiskAssessmentPrediction struct {
	Id                                 *string           `json:"id,omitempty"`
	Extension                          []Extension       `json:"extension,omitempty"`
	ModifierExtension                  []Extension       `json:"modifierExtension,omitempty"`
	Outcome                            *CodeableConcept  `json:"outcome,omitempty"`
	ProbabilityDecimal                 *Decimal          `json:"probabilityDecimal,omitempty"`
	ProbabilityDecimalPrimitiveElement *primitiveElement `json:"_probabilityDecimal,omitempty"`
	ProbabilityRange                   *Range            `json:"probabilityRange,omitempty"`
	QualitativeRisk                    *CodeableConcept  `json:"qualitativeRisk,omitempty"`
	RelativeRisk                       *Decimal          `json:"relativeRisk,omitempty"`
	RelativeRiskPrimitiveElement       *primitiveElement `json:"_relativeRisk,omitempty"`
	WhenPeriod                         *Period           `json:"whenPeriod,omitempty"`
	WhenRange                          *Range            `json:"whenRange,omitempty"`
	Rationale                          *String           `json:"rationale,omitempty"`
	RationalePrimitiveElement          *primitiveElement `json:"_rationale,omitempty"`
}

func (r RiskAssessmentPrediction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RiskAssessmentPrediction) marshalJSON() jsonRiskAssessmentPrediction {
	m := jsonRiskAssessmentPrediction{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Outcome = r.Outcome
	switch v := r.Probability.(type) {
	case Decimal:
		m.ProbabilityDecimal = &v
		if v.Id != nil || v.Extension != nil {
			m.ProbabilityDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Decimal:
		m.ProbabilityDecimal = v
		if v.Id != nil || v.Extension != nil {
			m.ProbabilityDecimalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Range:
		m.ProbabilityRange = &v
	case *Range:
		m.ProbabilityRange = v
	}
	m.QualitativeRisk = r.QualitativeRisk
	m.RelativeRisk = r.RelativeRisk
	if r.RelativeRisk != nil && (r.RelativeRisk.Id != nil || r.RelativeRisk.Extension != nil) {
		m.RelativeRiskPrimitiveElement = &primitiveElement{Id: r.RelativeRisk.Id, Extension: r.RelativeRisk.Extension}
	}
	switch v := r.When.(type) {
	case Period:
		m.WhenPeriod = &v
	case *Period:
		m.WhenPeriod = v
	case Range:
		m.WhenRange = &v
	case *Range:
		m.WhenRange = v
	}
	m.Rationale = r.Rationale
	if r.Rationale != nil && (r.Rationale.Id != nil || r.Rationale.Extension != nil) {
		m.RationalePrimitiveElement = &primitiveElement{Id: r.Rationale.Id, Extension: r.Rationale.Extension}
	}
	return m
}
func (r *RiskAssessmentPrediction) UnmarshalJSON(b []byte) error {
	var m jsonRiskAssessmentPrediction
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RiskAssessmentPrediction) unmarshalJSON(m jsonRiskAssessmentPrediction) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Outcome = m.Outcome
	if m.ProbabilityDecimal != nil || m.ProbabilityDecimalPrimitiveElement != nil {
		if r.Probability != nil {
			return fmt.Errorf("multiple values for field \"Probability\"")
		}
		v := m.ProbabilityDecimal
		if m.ProbabilityDecimalPrimitiveElement != nil {
			if v == nil {
				v = &Decimal{}
			}
			v.Id = m.ProbabilityDecimalPrimitiveElement.Id
			v.Extension = m.ProbabilityDecimalPrimitiveElement.Extension
		}
		r.Probability = v
	}
	if m.ProbabilityRange != nil {
		if r.Probability != nil {
			return fmt.Errorf("multiple values for field \"Probability\"")
		}
		v := m.ProbabilityRange
		r.Probability = v
	}
	r.QualitativeRisk = m.QualitativeRisk
	r.RelativeRisk = m.RelativeRisk
	if m.RelativeRiskPrimitiveElement != nil {
		r.RelativeRisk.Id = m.RelativeRiskPrimitiveElement.Id
		r.RelativeRisk.Extension = m.RelativeRiskPrimitiveElement.Extension
	}
	if m.WhenPeriod != nil {
		if r.When != nil {
			return fmt.Errorf("multiple values for field \"When\"")
		}
		v := m.WhenPeriod
		r.When = v
	}
	if m.WhenRange != nil {
		if r.When != nil {
			return fmt.Errorf("multiple values for field \"When\"")
		}
		v := m.WhenRange
		r.When = v
	}
	r.Rationale = m.Rationale
	if m.RationalePrimitiveElement != nil {
		r.Rationale.Id = m.RationalePrimitiveElement.Id
		r.Rationale.Extension = m.RationalePrimitiveElement.Extension
	}
	return nil
}
func (r RiskAssessmentPrediction) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
