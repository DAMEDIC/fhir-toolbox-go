package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// The MeasureReport resource contains the results of the calculation of a measure; and optionally a reference to the resources involved in that calculation.
type MeasureReport struct {
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
	// A formal identifier that is used to identify this MeasureReport when it is represented in other formats or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The MeasureReport status. No data will be available until the MeasureReport status is complete.
	Status Code
	// The type of measure report. This may be an individual report, which provides the score for the measure for an individual member of the population; a subject-listing, which returns the list of members that meet the various criteria in the measure; a summary report, which returns a population count for each of the criteria in the measure; or a data-collection, which enables the MeasureReport to be used to exchange the data-of-interest for a quality measure.
	Type Code
	// A reference to the Measure that was calculated to produce this report.
	Measure Canonical
	// Optional subject identifying the individual or individuals the report is for.
	Subject *Reference
	// The date this measure report was generated.
	Date *DateTime
	// The individual, location, or organization that is reporting the data.
	Reporter *Reference
	// The reporting period for which the report was calculated.
	Period Period
	// Whether improvement in the measure is noted by an increase or decrease in the measure score.
	ImprovementNotation *CodeableConcept
	// The results of the calculation, one for each population group in the measure.
	Group []MeasureReportGroup
	// A reference to a Bundle containing the Resources that were used in the calculation of this measure.
	EvaluatedResource []Reference
}
type jsonMeasureReport struct {
	ResourceType                  string               `json:"resourceType"`
	Id                            *Id                  `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement    `json:"_id,omitempty"`
	Meta                          *Meta                `json:"meta,omitempty"`
	ImplicitRules                 *Uri                 `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement    `json:"_implicitRules,omitempty"`
	Language                      *Code                `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement    `json:"_language,omitempty"`
	Text                          *Narrative           `json:"text,omitempty"`
	Contained                     []containedResource  `json:"contained,omitempty"`
	Extension                     []Extension          `json:"extension,omitempty"`
	ModifierExtension             []Extension          `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier         `json:"identifier,omitempty"`
	Status                        Code                 `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement    `json:"_status,omitempty"`
	Type                          Code                 `json:"type,omitempty"`
	TypePrimitiveElement          *primitiveElement    `json:"_type,omitempty"`
	Measure                       Canonical            `json:"measure,omitempty"`
	MeasurePrimitiveElement       *primitiveElement    `json:"_measure,omitempty"`
	Subject                       *Reference           `json:"subject,omitempty"`
	Date                          *DateTime            `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement    `json:"_date,omitempty"`
	Reporter                      *Reference           `json:"reporter,omitempty"`
	Period                        Period               `json:"period,omitempty"`
	ImprovementNotation           *CodeableConcept     `json:"improvementNotation,omitempty"`
	Group                         []MeasureReportGroup `json:"group,omitempty"`
	EvaluatedResource             []Reference          `json:"evaluatedResource,omitempty"`
}

func (r MeasureReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MeasureReport) marshalJSON() jsonMeasureReport {
	m := jsonMeasureReport{}
	m.ResourceType = "MeasureReport"
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
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Measure = r.Measure
	if r.Measure.Id != nil || r.Measure.Extension != nil {
		m.MeasurePrimitiveElement = &primitiveElement{Id: r.Measure.Id, Extension: r.Measure.Extension}
	}
	m.Subject = r.Subject
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Reporter = r.Reporter
	m.Period = r.Period
	m.ImprovementNotation = r.ImprovementNotation
	m.Group = r.Group
	m.EvaluatedResource = r.EvaluatedResource
	return m
}
func (r *MeasureReport) UnmarshalJSON(b []byte) error {
	var m jsonMeasureReport
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MeasureReport) unmarshalJSON(m jsonMeasureReport) error {
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
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Measure = m.Measure
	if m.MeasurePrimitiveElement != nil {
		r.Measure.Id = m.MeasurePrimitiveElement.Id
		r.Measure.Extension = m.MeasurePrimitiveElement.Extension
	}
	r.Subject = m.Subject
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Reporter = m.Reporter
	r.Period = m.Period
	r.ImprovementNotation = m.ImprovementNotation
	r.Group = m.Group
	r.EvaluatedResource = m.EvaluatedResource
	return nil
}
func (r MeasureReport) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The results of the calculation, one for each population group in the measure.
type MeasureReportGroup struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The meaning of the population group as defined in the measure definition.
	Code *CodeableConcept
	// The populations that make up the population group, one for each type of population appropriate for the measure.
	Population []MeasureReportGroupPopulation
	// The measure score for this population group, calculated as appropriate for the measure type and scoring method, and based on the contents of the populations defined in the group.
	MeasureScore *Quantity
	// When a measure includes multiple stratifiers, there will be a stratifier group for each stratifier defined by the measure.
	Stratifier []MeasureReportGroupStratifier
}
type jsonMeasureReportGroup struct {
	Id                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept               `json:"code,omitempty"`
	Population        []MeasureReportGroupPopulation `json:"population,omitempty"`
	MeasureScore      *Quantity                      `json:"measureScore,omitempty"`
	Stratifier        []MeasureReportGroupStratifier `json:"stratifier,omitempty"`
}

func (r MeasureReportGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MeasureReportGroup) marshalJSON() jsonMeasureReportGroup {
	m := jsonMeasureReportGroup{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Population = r.Population
	m.MeasureScore = r.MeasureScore
	m.Stratifier = r.Stratifier
	return m
}
func (r *MeasureReportGroup) UnmarshalJSON(b []byte) error {
	var m jsonMeasureReportGroup
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MeasureReportGroup) unmarshalJSON(m jsonMeasureReportGroup) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Population = m.Population
	r.MeasureScore = m.MeasureScore
	r.Stratifier = m.Stratifier
	return nil
}
func (r MeasureReportGroup) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The populations that make up the population group, one for each type of population appropriate for the measure.
type MeasureReportGroupPopulation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of the population.
	Code *CodeableConcept
	// The number of members of the population.
	Count *Integer
	// This element refers to a List of subject level MeasureReport resources, one for each subject in this population.
	SubjectResults *Reference
}
type jsonMeasureReportGroupPopulation struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Code                  *CodeableConcept  `json:"code,omitempty"`
	Count                 *Integer          `json:"count,omitempty"`
	CountPrimitiveElement *primitiveElement `json:"_count,omitempty"`
	SubjectResults        *Reference        `json:"subjectResults,omitempty"`
}

func (r MeasureReportGroupPopulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MeasureReportGroupPopulation) marshalJSON() jsonMeasureReportGroupPopulation {
	m := jsonMeasureReportGroupPopulation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Count = r.Count
	if r.Count != nil && (r.Count.Id != nil || r.Count.Extension != nil) {
		m.CountPrimitiveElement = &primitiveElement{Id: r.Count.Id, Extension: r.Count.Extension}
	}
	m.SubjectResults = r.SubjectResults
	return m
}
func (r *MeasureReportGroupPopulation) UnmarshalJSON(b []byte) error {
	var m jsonMeasureReportGroupPopulation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MeasureReportGroupPopulation) unmarshalJSON(m jsonMeasureReportGroupPopulation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Count = m.Count
	if m.CountPrimitiveElement != nil {
		r.Count.Id = m.CountPrimitiveElement.Id
		r.Count.Extension = m.CountPrimitiveElement.Extension
	}
	r.SubjectResults = m.SubjectResults
	return nil
}
func (r MeasureReportGroupPopulation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// When a measure includes multiple stratifiers, there will be a stratifier group for each stratifier defined by the measure.
type MeasureReportGroupStratifier struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The meaning of this stratifier, as defined in the measure definition.
	Code []CodeableConcept
	// This element contains the results for a single stratum within the stratifier. For example, when stratifying on administrative gender, there will be four strata, one for each possible gender value.
	Stratum []MeasureReportGroupStratifierStratum
}
type jsonMeasureReportGroupStratifier struct {
	Id                *string                               `json:"id,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept                     `json:"code,omitempty"`
	Stratum           []MeasureReportGroupStratifierStratum `json:"stratum,omitempty"`
}

func (r MeasureReportGroupStratifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MeasureReportGroupStratifier) marshalJSON() jsonMeasureReportGroupStratifier {
	m := jsonMeasureReportGroupStratifier{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Stratum = r.Stratum
	return m
}
func (r *MeasureReportGroupStratifier) UnmarshalJSON(b []byte) error {
	var m jsonMeasureReportGroupStratifier
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MeasureReportGroupStratifier) unmarshalJSON(m jsonMeasureReportGroupStratifier) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Stratum = m.Stratum
	return nil
}
func (r MeasureReportGroupStratifier) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// This element contains the results for a single stratum within the stratifier. For example, when stratifying on administrative gender, there will be four strata, one for each possible gender value.
type MeasureReportGroupStratifierStratum struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The value for this stratum, expressed as a CodeableConcept. When defining stratifiers on complex values, the value must be rendered such that the value for each stratum within the stratifier is unique.
	Value *CodeableConcept
	// A stratifier component value.
	Component []MeasureReportGroupStratifierStratumComponent
	// The populations that make up the stratum, one for each type of population appropriate to the measure.
	Population []MeasureReportGroupStratifierStratumPopulation
	// The measure score for this stratum, calculated as appropriate for the measure type and scoring method, and based on only the members of this stratum.
	MeasureScore *Quantity
}
type jsonMeasureReportGroupStratifierStratum struct {
	Id                *string                                         `json:"id,omitempty"`
	Extension         []Extension                                     `json:"extension,omitempty"`
	ModifierExtension []Extension                                     `json:"modifierExtension,omitempty"`
	Value             *CodeableConcept                                `json:"value,omitempty"`
	Component         []MeasureReportGroupStratifierStratumComponent  `json:"component,omitempty"`
	Population        []MeasureReportGroupStratifierStratumPopulation `json:"population,omitempty"`
	MeasureScore      *Quantity                                       `json:"measureScore,omitempty"`
}

func (r MeasureReportGroupStratifierStratum) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MeasureReportGroupStratifierStratum) marshalJSON() jsonMeasureReportGroupStratifierStratum {
	m := jsonMeasureReportGroupStratifierStratum{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Value = r.Value
	m.Component = r.Component
	m.Population = r.Population
	m.MeasureScore = r.MeasureScore
	return m
}
func (r *MeasureReportGroupStratifierStratum) UnmarshalJSON(b []byte) error {
	var m jsonMeasureReportGroupStratifierStratum
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MeasureReportGroupStratifierStratum) unmarshalJSON(m jsonMeasureReportGroupStratifierStratum) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Value = m.Value
	r.Component = m.Component
	r.Population = m.Population
	r.MeasureScore = m.MeasureScore
	return nil
}
func (r MeasureReportGroupStratifierStratum) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A stratifier component value.
type MeasureReportGroupStratifierStratumComponent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The code for the stratum component value.
	Code CodeableConcept
	// The stratum component value.
	Value CodeableConcept
}
type jsonMeasureReportGroupStratifierStratumComponent struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code,omitempty"`
	Value             CodeableConcept `json:"value,omitempty"`
}

func (r MeasureReportGroupStratifierStratumComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MeasureReportGroupStratifierStratumComponent) marshalJSON() jsonMeasureReportGroupStratifierStratumComponent {
	m := jsonMeasureReportGroupStratifierStratumComponent{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Value = r.Value
	return m
}
func (r *MeasureReportGroupStratifierStratumComponent) UnmarshalJSON(b []byte) error {
	var m jsonMeasureReportGroupStratifierStratumComponent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MeasureReportGroupStratifierStratumComponent) unmarshalJSON(m jsonMeasureReportGroupStratifierStratumComponent) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Value = m.Value
	return nil
}
func (r MeasureReportGroupStratifierStratumComponent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The populations that make up the stratum, one for each type of population appropriate to the measure.
type MeasureReportGroupStratifierStratumPopulation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of the population.
	Code *CodeableConcept
	// The number of members of the population in this stratum.
	Count *Integer
	// This element refers to a List of subject level MeasureReport resources, one for each subject in this population in this stratum.
	SubjectResults *Reference
}
type jsonMeasureReportGroupStratifierStratumPopulation struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Code                  *CodeableConcept  `json:"code,omitempty"`
	Count                 *Integer          `json:"count,omitempty"`
	CountPrimitiveElement *primitiveElement `json:"_count,omitempty"`
	SubjectResults        *Reference        `json:"subjectResults,omitempty"`
}

func (r MeasureReportGroupStratifierStratumPopulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MeasureReportGroupStratifierStratumPopulation) marshalJSON() jsonMeasureReportGroupStratifierStratumPopulation {
	m := jsonMeasureReportGroupStratifierStratumPopulation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Count = r.Count
	if r.Count != nil && (r.Count.Id != nil || r.Count.Extension != nil) {
		m.CountPrimitiveElement = &primitiveElement{Id: r.Count.Id, Extension: r.Count.Extension}
	}
	m.SubjectResults = r.SubjectResults
	return m
}
func (r *MeasureReportGroupStratifierStratumPopulation) UnmarshalJSON(b []byte) error {
	var m jsonMeasureReportGroupStratifierStratumPopulation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MeasureReportGroupStratifierStratumPopulation) unmarshalJSON(m jsonMeasureReportGroupStratifierStratumPopulation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Count = m.Count
	if m.CountPrimitiveElement != nil {
		r.Count.Id = m.CountPrimitiveElement.Id
		r.Count.Extension = m.CountPrimitiveElement.Extension
	}
	r.SubjectResults = m.SubjectResults
	return nil
}
func (r MeasureReportGroupStratifierStratumPopulation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
