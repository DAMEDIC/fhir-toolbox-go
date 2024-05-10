package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// The Measure resource provides the definition of a quality measure.
type Measure struct {
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
	// An absolute URI that is used to identify this measure when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this measure is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the measure is stored on different servers.
	Url *Uri
	// A formal identifier that is used to identify this measure when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the measure when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the measure author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence. To provide a version consistent with the Decision Support Service specification, use the format Major.Minor.Revision (e.g. 1.0.0). For more information on versioning knowledge assets, refer to the Decision Support Service specification. Note that a version is required for non-experimental active artifacts.
	Version *String
	// A natural language name identifying the measure. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the measure.
	Title *String
	// An explanatory or alternate title for the measure giving additional information about its content.
	Subtitle *String
	// The status of this measure. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this measure is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The intended subjects for the measure. If this element is not provided, a Patient subject is assumed, but the subject of the measure can be anything.
	Subject isMeasureSubject
	// The date  (and optionally time) when the measure was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the measure changes.
	Date *DateTime
	// The name of the organization or individual that published the measure.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the measure from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate measure instances.
	UseContext []UsageContext
	// A legal or geographic region in which the measure is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this measure is needed and why it has been designed as it has.
	Purpose *Markdown
	// A detailed description, from a clinical perspective, of how the measure is used.
	Usage *String
	// A copyright statement relating to the measure and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the measure.
	Copyright *Markdown
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *Date
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *Date
	// The period during which the measure content was or is planned to be in active use.
	EffectivePeriod *Period
	// Descriptive topics related to the content of the measure. Topics provide a high-level categorization grouping types of measures that can be useful for filtering and searching.
	Topic []CodeableConcept
	// An individiual or organization primarily involved in the creation and maintenance of the content.
	Author []ContactDetail
	// An individual or organization primarily responsible for internal coherence of the content.
	Editor []ContactDetail
	// An individual or organization primarily responsible for review of some aspect of the content.
	Reviewer []ContactDetail
	// An individual or organization responsible for officially endorsing the content for use in some setting.
	Endorser []ContactDetail
	// Related artifacts such as additional documentation, justification, or bibliographic references.
	RelatedArtifact []RelatedArtifact
	// A reference to a Library resource containing the formal logic used by the measure.
	Library []Canonical
	// Notices and disclaimers regarding the use of the measure or related to intellectual property (such as code systems) referenced by the measure.
	Disclaimer *Markdown
	// Indicates how the calculation is performed for the measure, including proportion, ratio, continuous-variable, and cohort. The value set is extensible, allowing additional measure scoring types to be represented.
	Scoring *CodeableConcept
	// If this is a composite measure, the scoring method used to combine the component measures to determine the composite score.
	CompositeScoring *CodeableConcept
	// Indicates whether the measure is used to examine a process, an outcome over time, a patient-reported outcome, or a structure measure such as utilization.
	Type []CodeableConcept
	// A description of the risk adjustment factors that may impact the resulting score for the measure and how they may be accounted for when computing and reporting measure results.
	RiskAdjustment *String
	// Describes how to combine the information calculated, based on logic in each of several populations, into one summarized result.
	RateAggregation *String
	// Provides a succinct statement of the need for the measure. Usually includes statements pertaining to importance criterion: impact, gap in care, and evidence.
	Rationale *Markdown
	// Provides a summary of relevant clinical guidelines or other clinical recommendations supporting the measure.
	ClinicalRecommendationStatement *Markdown
	// Information on whether an increase or decrease in score is the preferred result (e.g., a higher score indicates better quality OR a lower score indicates better quality OR quality is within a range).
	ImprovementNotation *CodeableConcept
	// Provides a description of an individual term used within the measure.
	Definition []Markdown
	// Additional guidance for the measure including how it can be used in a clinical context, and the intent of the measure.
	Guidance *Markdown
	// A group of population criteria for the measure.
	Group []MeasureGroup
	// The supplemental data criteria for the measure report, specified as either the name of a valid CQL expression within a referenced library, or a valid FHIR Resource Path.
	SupplementalData []MeasureSupplementalData
}
type isMeasureSubject interface {
	isMeasureSubject()
}

func (r CodeableConcept) isMeasureSubject() {}
func (r Reference) isMeasureSubject()       {}

type jsonMeasure struct {
	ResourceType                                    string                    `json:"resourceType"`
	Id                                              *Id                       `json:"id,omitempty"`
	IdPrimitiveElement                              *primitiveElement         `json:"_id,omitempty"`
	Meta                                            *Meta                     `json:"meta,omitempty"`
	ImplicitRules                                   *Uri                      `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement                   *primitiveElement         `json:"_implicitRules,omitempty"`
	Language                                        *Code                     `json:"language,omitempty"`
	LanguagePrimitiveElement                        *primitiveElement         `json:"_language,omitempty"`
	Text                                            *Narrative                `json:"text,omitempty"`
	Contained                                       []containedResource       `json:"contained,omitempty"`
	Extension                                       []Extension               `json:"extension,omitempty"`
	ModifierExtension                               []Extension               `json:"modifierExtension,omitempty"`
	Url                                             *Uri                      `json:"url,omitempty"`
	UrlPrimitiveElement                             *primitiveElement         `json:"_url,omitempty"`
	Identifier                                      []Identifier              `json:"identifier,omitempty"`
	Version                                         *String                   `json:"version,omitempty"`
	VersionPrimitiveElement                         *primitiveElement         `json:"_version,omitempty"`
	Name                                            *String                   `json:"name,omitempty"`
	NamePrimitiveElement                            *primitiveElement         `json:"_name,omitempty"`
	Title                                           *String                   `json:"title,omitempty"`
	TitlePrimitiveElement                           *primitiveElement         `json:"_title,omitempty"`
	Subtitle                                        *String                   `json:"subtitle,omitempty"`
	SubtitlePrimitiveElement                        *primitiveElement         `json:"_subtitle,omitempty"`
	Status                                          Code                      `json:"status,omitempty"`
	StatusPrimitiveElement                          *primitiveElement         `json:"_status,omitempty"`
	Experimental                                    *Boolean                  `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement                    *primitiveElement         `json:"_experimental,omitempty"`
	SubjectCodeableConcept                          *CodeableConcept          `json:"subjectCodeableConcept,omitempty"`
	SubjectReference                                *Reference                `json:"subjectReference,omitempty"`
	Date                                            *DateTime                 `json:"date,omitempty"`
	DatePrimitiveElement                            *primitiveElement         `json:"_date,omitempty"`
	Publisher                                       *String                   `json:"publisher,omitempty"`
	PublisherPrimitiveElement                       *primitiveElement         `json:"_publisher,omitempty"`
	Contact                                         []ContactDetail           `json:"contact,omitempty"`
	Description                                     *Markdown                 `json:"description,omitempty"`
	DescriptionPrimitiveElement                     *primitiveElement         `json:"_description,omitempty"`
	UseContext                                      []UsageContext            `json:"useContext,omitempty"`
	Jurisdiction                                    []CodeableConcept         `json:"jurisdiction,omitempty"`
	Purpose                                         *Markdown                 `json:"purpose,omitempty"`
	PurposePrimitiveElement                         *primitiveElement         `json:"_purpose,omitempty"`
	Usage                                           *String                   `json:"usage,omitempty"`
	UsagePrimitiveElement                           *primitiveElement         `json:"_usage,omitempty"`
	Copyright                                       *Markdown                 `json:"copyright,omitempty"`
	CopyrightPrimitiveElement                       *primitiveElement         `json:"_copyright,omitempty"`
	ApprovalDate                                    *Date                     `json:"approvalDate,omitempty"`
	ApprovalDatePrimitiveElement                    *primitiveElement         `json:"_approvalDate,omitempty"`
	LastReviewDate                                  *Date                     `json:"lastReviewDate,omitempty"`
	LastReviewDatePrimitiveElement                  *primitiveElement         `json:"_lastReviewDate,omitempty"`
	EffectivePeriod                                 *Period                   `json:"effectivePeriod,omitempty"`
	Topic                                           []CodeableConcept         `json:"topic,omitempty"`
	Author                                          []ContactDetail           `json:"author,omitempty"`
	Editor                                          []ContactDetail           `json:"editor,omitempty"`
	Reviewer                                        []ContactDetail           `json:"reviewer,omitempty"`
	Endorser                                        []ContactDetail           `json:"endorser,omitempty"`
	RelatedArtifact                                 []RelatedArtifact         `json:"relatedArtifact,omitempty"`
	Library                                         []Canonical               `json:"library,omitempty"`
	LibraryPrimitiveElement                         []*primitiveElement       `json:"_library,omitempty"`
	Disclaimer                                      *Markdown                 `json:"disclaimer,omitempty"`
	DisclaimerPrimitiveElement                      *primitiveElement         `json:"_disclaimer,omitempty"`
	Scoring                                         *CodeableConcept          `json:"scoring,omitempty"`
	CompositeScoring                                *CodeableConcept          `json:"compositeScoring,omitempty"`
	Type                                            []CodeableConcept         `json:"type,omitempty"`
	RiskAdjustment                                  *String                   `json:"riskAdjustment,omitempty"`
	RiskAdjustmentPrimitiveElement                  *primitiveElement         `json:"_riskAdjustment,omitempty"`
	RateAggregation                                 *String                   `json:"rateAggregation,omitempty"`
	RateAggregationPrimitiveElement                 *primitiveElement         `json:"_rateAggregation,omitempty"`
	Rationale                                       *Markdown                 `json:"rationale,omitempty"`
	RationalePrimitiveElement                       *primitiveElement         `json:"_rationale,omitempty"`
	ClinicalRecommendationStatement                 *Markdown                 `json:"clinicalRecommendationStatement,omitempty"`
	ClinicalRecommendationStatementPrimitiveElement *primitiveElement         `json:"_clinicalRecommendationStatement,omitempty"`
	ImprovementNotation                             *CodeableConcept          `json:"improvementNotation,omitempty"`
	Definition                                      []Markdown                `json:"definition,omitempty"`
	DefinitionPrimitiveElement                      []*primitiveElement       `json:"_definition,omitempty"`
	Guidance                                        *Markdown                 `json:"guidance,omitempty"`
	GuidancePrimitiveElement                        *primitiveElement         `json:"_guidance,omitempty"`
	Group                                           []MeasureGroup            `json:"group,omitempty"`
	SupplementalData                                []MeasureSupplementalData `json:"supplementalData,omitempty"`
}

func (r Measure) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Measure) marshalJSON() jsonMeasure {
	m := jsonMeasure{}
	m.ResourceType = "Measure"
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
	m.Url = r.Url
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	m.Identifier = r.Identifier
	m.Version = r.Version
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Title = r.Title
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	m.Subtitle = r.Subtitle
	if r.Subtitle != nil && (r.Subtitle.Id != nil || r.Subtitle.Extension != nil) {
		m.SubtitlePrimitiveElement = &primitiveElement{Id: r.Subtitle.Id, Extension: r.Subtitle.Extension}
	}
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Experimental = r.Experimental
	if r.Experimental != nil && (r.Experimental.Id != nil || r.Experimental.Extension != nil) {
		m.ExperimentalPrimitiveElement = &primitiveElement{Id: r.Experimental.Id, Extension: r.Experimental.Extension}
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
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Publisher = r.Publisher
	if r.Publisher != nil && (r.Publisher.Id != nil || r.Publisher.Extension != nil) {
		m.PublisherPrimitiveElement = &primitiveElement{Id: r.Publisher.Id, Extension: r.Publisher.Extension}
	}
	m.Contact = r.Contact
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.UseContext = r.UseContext
	m.Jurisdiction = r.Jurisdiction
	m.Purpose = r.Purpose
	if r.Purpose != nil && (r.Purpose.Id != nil || r.Purpose.Extension != nil) {
		m.PurposePrimitiveElement = &primitiveElement{Id: r.Purpose.Id, Extension: r.Purpose.Extension}
	}
	m.Usage = r.Usage
	if r.Usage != nil && (r.Usage.Id != nil || r.Usage.Extension != nil) {
		m.UsagePrimitiveElement = &primitiveElement{Id: r.Usage.Id, Extension: r.Usage.Extension}
	}
	m.Copyright = r.Copyright
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		m.CopyrightPrimitiveElement = &primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
	}
	m.ApprovalDate = r.ApprovalDate
	if r.ApprovalDate != nil && (r.ApprovalDate.Id != nil || r.ApprovalDate.Extension != nil) {
		m.ApprovalDatePrimitiveElement = &primitiveElement{Id: r.ApprovalDate.Id, Extension: r.ApprovalDate.Extension}
	}
	m.LastReviewDate = r.LastReviewDate
	if r.LastReviewDate != nil && (r.LastReviewDate.Id != nil || r.LastReviewDate.Extension != nil) {
		m.LastReviewDatePrimitiveElement = &primitiveElement{Id: r.LastReviewDate.Id, Extension: r.LastReviewDate.Extension}
	}
	m.EffectivePeriod = r.EffectivePeriod
	m.Topic = r.Topic
	m.Author = r.Author
	m.Editor = r.Editor
	m.Reviewer = r.Reviewer
	m.Endorser = r.Endorser
	m.RelatedArtifact = r.RelatedArtifact
	m.Library = r.Library
	anyLibraryIdOrExtension := false
	for _, e := range r.Library {
		if e.Id != nil || e.Extension != nil {
			anyLibraryIdOrExtension = true
			break
		}
	}
	if anyLibraryIdOrExtension {
		m.LibraryPrimitiveElement = make([]*primitiveElement, 0, len(r.Library))
		for _, e := range r.Library {
			if e.Id != nil || e.Extension != nil {
				m.LibraryPrimitiveElement = append(m.LibraryPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.LibraryPrimitiveElement = append(m.LibraryPrimitiveElement, nil)
			}
		}
	}
	m.Disclaimer = r.Disclaimer
	if r.Disclaimer != nil && (r.Disclaimer.Id != nil || r.Disclaimer.Extension != nil) {
		m.DisclaimerPrimitiveElement = &primitiveElement{Id: r.Disclaimer.Id, Extension: r.Disclaimer.Extension}
	}
	m.Scoring = r.Scoring
	m.CompositeScoring = r.CompositeScoring
	m.Type = r.Type
	m.RiskAdjustment = r.RiskAdjustment
	if r.RiskAdjustment != nil && (r.RiskAdjustment.Id != nil || r.RiskAdjustment.Extension != nil) {
		m.RiskAdjustmentPrimitiveElement = &primitiveElement{Id: r.RiskAdjustment.Id, Extension: r.RiskAdjustment.Extension}
	}
	m.RateAggregation = r.RateAggregation
	if r.RateAggregation != nil && (r.RateAggregation.Id != nil || r.RateAggregation.Extension != nil) {
		m.RateAggregationPrimitiveElement = &primitiveElement{Id: r.RateAggregation.Id, Extension: r.RateAggregation.Extension}
	}
	m.Rationale = r.Rationale
	if r.Rationale != nil && (r.Rationale.Id != nil || r.Rationale.Extension != nil) {
		m.RationalePrimitiveElement = &primitiveElement{Id: r.Rationale.Id, Extension: r.Rationale.Extension}
	}
	m.ClinicalRecommendationStatement = r.ClinicalRecommendationStatement
	if r.ClinicalRecommendationStatement != nil && (r.ClinicalRecommendationStatement.Id != nil || r.ClinicalRecommendationStatement.Extension != nil) {
		m.ClinicalRecommendationStatementPrimitiveElement = &primitiveElement{Id: r.ClinicalRecommendationStatement.Id, Extension: r.ClinicalRecommendationStatement.Extension}
	}
	m.ImprovementNotation = r.ImprovementNotation
	m.Definition = r.Definition
	anyDefinitionIdOrExtension := false
	for _, e := range r.Definition {
		if e.Id != nil || e.Extension != nil {
			anyDefinitionIdOrExtension = true
			break
		}
	}
	if anyDefinitionIdOrExtension {
		m.DefinitionPrimitiveElement = make([]*primitiveElement, 0, len(r.Definition))
		for _, e := range r.Definition {
			if e.Id != nil || e.Extension != nil {
				m.DefinitionPrimitiveElement = append(m.DefinitionPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.DefinitionPrimitiveElement = append(m.DefinitionPrimitiveElement, nil)
			}
		}
	}
	m.Guidance = r.Guidance
	if r.Guidance != nil && (r.Guidance.Id != nil || r.Guidance.Extension != nil) {
		m.GuidancePrimitiveElement = &primitiveElement{Id: r.Guidance.Id, Extension: r.Guidance.Extension}
	}
	m.Group = r.Group
	m.SupplementalData = r.SupplementalData
	return m
}
func (r *Measure) UnmarshalJSON(b []byte) error {
	var m jsonMeasure
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Measure) unmarshalJSON(m jsonMeasure) error {
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
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Identifier = m.Identifier
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Subtitle = m.Subtitle
	if m.SubtitlePrimitiveElement != nil {
		r.Subtitle.Id = m.SubtitlePrimitiveElement.Id
		r.Subtitle.Extension = m.SubtitlePrimitiveElement.Extension
	}
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Experimental = m.Experimental
	if m.ExperimentalPrimitiveElement != nil {
		r.Experimental.Id = m.ExperimentalPrimitiveElement.Id
		r.Experimental.Extension = m.ExperimentalPrimitiveElement.Extension
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
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Publisher = m.Publisher
	if m.PublisherPrimitiveElement != nil {
		r.Publisher.Id = m.PublisherPrimitiveElement.Id
		r.Publisher.Extension = m.PublisherPrimitiveElement.Extension
	}
	r.Contact = m.Contact
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.UseContext = m.UseContext
	r.Jurisdiction = m.Jurisdiction
	r.Purpose = m.Purpose
	if m.PurposePrimitiveElement != nil {
		r.Purpose.Id = m.PurposePrimitiveElement.Id
		r.Purpose.Extension = m.PurposePrimitiveElement.Extension
	}
	r.Usage = m.Usage
	if m.UsagePrimitiveElement != nil {
		r.Usage.Id = m.UsagePrimitiveElement.Id
		r.Usage.Extension = m.UsagePrimitiveElement.Extension
	}
	r.Copyright = m.Copyright
	if m.CopyrightPrimitiveElement != nil {
		r.Copyright.Id = m.CopyrightPrimitiveElement.Id
		r.Copyright.Extension = m.CopyrightPrimitiveElement.Extension
	}
	r.ApprovalDate = m.ApprovalDate
	if m.ApprovalDatePrimitiveElement != nil {
		r.ApprovalDate.Id = m.ApprovalDatePrimitiveElement.Id
		r.ApprovalDate.Extension = m.ApprovalDatePrimitiveElement.Extension
	}
	r.LastReviewDate = m.LastReviewDate
	if m.LastReviewDatePrimitiveElement != nil {
		r.LastReviewDate.Id = m.LastReviewDatePrimitiveElement.Id
		r.LastReviewDate.Extension = m.LastReviewDatePrimitiveElement.Extension
	}
	r.EffectivePeriod = m.EffectivePeriod
	r.Topic = m.Topic
	r.Author = m.Author
	r.Editor = m.Editor
	r.Reviewer = m.Reviewer
	r.Endorser = m.Endorser
	r.RelatedArtifact = m.RelatedArtifact
	r.Library = m.Library
	for i, e := range m.LibraryPrimitiveElement {
		if len(r.Library) > i {
			r.Library[i].Id = e.Id
			r.Library[i].Extension = e.Extension
		} else {
			r.Library = append(r.Library, Canonical{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Disclaimer = m.Disclaimer
	if m.DisclaimerPrimitiveElement != nil {
		r.Disclaimer.Id = m.DisclaimerPrimitiveElement.Id
		r.Disclaimer.Extension = m.DisclaimerPrimitiveElement.Extension
	}
	r.Scoring = m.Scoring
	r.CompositeScoring = m.CompositeScoring
	r.Type = m.Type
	r.RiskAdjustment = m.RiskAdjustment
	if m.RiskAdjustmentPrimitiveElement != nil {
		r.RiskAdjustment.Id = m.RiskAdjustmentPrimitiveElement.Id
		r.RiskAdjustment.Extension = m.RiskAdjustmentPrimitiveElement.Extension
	}
	r.RateAggregation = m.RateAggregation
	if m.RateAggregationPrimitiveElement != nil {
		r.RateAggregation.Id = m.RateAggregationPrimitiveElement.Id
		r.RateAggregation.Extension = m.RateAggregationPrimitiveElement.Extension
	}
	r.Rationale = m.Rationale
	if m.RationalePrimitiveElement != nil {
		r.Rationale.Id = m.RationalePrimitiveElement.Id
		r.Rationale.Extension = m.RationalePrimitiveElement.Extension
	}
	r.ClinicalRecommendationStatement = m.ClinicalRecommendationStatement
	if m.ClinicalRecommendationStatementPrimitiveElement != nil {
		r.ClinicalRecommendationStatement.Id = m.ClinicalRecommendationStatementPrimitiveElement.Id
		r.ClinicalRecommendationStatement.Extension = m.ClinicalRecommendationStatementPrimitiveElement.Extension
	}
	r.ImprovementNotation = m.ImprovementNotation
	r.Definition = m.Definition
	for i, e := range m.DefinitionPrimitiveElement {
		if len(r.Definition) > i {
			r.Definition[i].Id = e.Id
			r.Definition[i].Extension = e.Extension
		} else {
			r.Definition = append(r.Definition, Markdown{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Guidance = m.Guidance
	if m.GuidancePrimitiveElement != nil {
		r.Guidance.Id = m.GuidancePrimitiveElement.Id
		r.Guidance.Extension = m.GuidancePrimitiveElement.Extension
	}
	r.Group = m.Group
	r.SupplementalData = m.SupplementalData
	return nil
}
func (r Measure) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A group of population criteria for the measure.
type MeasureGroup struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates a meaning for the group. This can be as simple as a unique identifier, or it can establish meaning in a broader context by drawing from a terminology, allowing groups to be correlated across measures.
	Code *CodeableConcept
	// The human readable description of this population group.
	Description *String
	// A population criteria for the measure.
	Population []MeasureGroupPopulation
	// The stratifier criteria for the measure report, specified as either the name of a valid CQL expression defined within a referenced library or a valid FHIR Resource Path.
	Stratifier []MeasureGroupStratifier
}
type jsonMeasureGroup struct {
	Id                          *string                  `json:"id,omitempty"`
	Extension                   []Extension              `json:"extension,omitempty"`
	ModifierExtension           []Extension              `json:"modifierExtension,omitempty"`
	Code                        *CodeableConcept         `json:"code,omitempty"`
	Description                 *String                  `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement        `json:"_description,omitempty"`
	Population                  []MeasureGroupPopulation `json:"population,omitempty"`
	Stratifier                  []MeasureGroupStratifier `json:"stratifier,omitempty"`
}

func (r MeasureGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MeasureGroup) marshalJSON() jsonMeasureGroup {
	m := jsonMeasureGroup{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Population = r.Population
	m.Stratifier = r.Stratifier
	return m
}
func (r *MeasureGroup) UnmarshalJSON(b []byte) error {
	var m jsonMeasureGroup
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MeasureGroup) unmarshalJSON(m jsonMeasureGroup) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Population = m.Population
	r.Stratifier = m.Stratifier
	return nil
}
func (r MeasureGroup) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A population criteria for the measure.
type MeasureGroupPopulation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of population criteria.
	Code *CodeableConcept
	// The human readable description of this population criteria.
	Description *String
	// An expression that specifies the criteria for the population, typically the name of an expression in a library.
	Criteria Expression
}
type jsonMeasureGroupPopulation struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Code                        *CodeableConcept  `json:"code,omitempty"`
	Description                 *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
	Criteria                    Expression        `json:"criteria,omitempty"`
}

func (r MeasureGroupPopulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MeasureGroupPopulation) marshalJSON() jsonMeasureGroupPopulation {
	m := jsonMeasureGroupPopulation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Criteria = r.Criteria
	return m
}
func (r *MeasureGroupPopulation) UnmarshalJSON(b []byte) error {
	var m jsonMeasureGroupPopulation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MeasureGroupPopulation) unmarshalJSON(m jsonMeasureGroupPopulation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Criteria = m.Criteria
	return nil
}
func (r MeasureGroupPopulation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The stratifier criteria for the measure report, specified as either the name of a valid CQL expression defined within a referenced library or a valid FHIR Resource Path.
type MeasureGroupStratifier struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates a meaning for the stratifier. This can be as simple as a unique identifier, or it can establish meaning in a broader context by drawing from a terminology, allowing stratifiers to be correlated across measures.
	Code *CodeableConcept
	// The human readable description of this stratifier criteria.
	Description *String
	// An expression that specifies the criteria for the stratifier. This is typically the name of an expression defined within a referenced library, but it may also be a path to a stratifier element.
	Criteria *Expression
	// A component of the stratifier criteria for the measure report, specified as either the name of a valid CQL expression defined within a referenced library or a valid FHIR Resource Path.
	Component []MeasureGroupStratifierComponent
}
type jsonMeasureGroupStratifier struct {
	Id                          *string                           `json:"id,omitempty"`
	Extension                   []Extension                       `json:"extension,omitempty"`
	ModifierExtension           []Extension                       `json:"modifierExtension,omitempty"`
	Code                        *CodeableConcept                  `json:"code,omitempty"`
	Description                 *String                           `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement                 `json:"_description,omitempty"`
	Criteria                    *Expression                       `json:"criteria,omitempty"`
	Component                   []MeasureGroupStratifierComponent `json:"component,omitempty"`
}

func (r MeasureGroupStratifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MeasureGroupStratifier) marshalJSON() jsonMeasureGroupStratifier {
	m := jsonMeasureGroupStratifier{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Criteria = r.Criteria
	m.Component = r.Component
	return m
}
func (r *MeasureGroupStratifier) UnmarshalJSON(b []byte) error {
	var m jsonMeasureGroupStratifier
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MeasureGroupStratifier) unmarshalJSON(m jsonMeasureGroupStratifier) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Criteria = m.Criteria
	r.Component = m.Component
	return nil
}
func (r MeasureGroupStratifier) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A component of the stratifier criteria for the measure report, specified as either the name of a valid CQL expression defined within a referenced library or a valid FHIR Resource Path.
type MeasureGroupStratifierComponent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates a meaning for the stratifier component. This can be as simple as a unique identifier, or it can establish meaning in a broader context by drawing from a terminology, allowing stratifiers to be correlated across measures.
	Code *CodeableConcept
	// The human readable description of this stratifier criteria component.
	Description *String
	// An expression that specifies the criteria for this component of the stratifier. This is typically the name of an expression defined within a referenced library, but it may also be a path to a stratifier element.
	Criteria Expression
}
type jsonMeasureGroupStratifierComponent struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Code                        *CodeableConcept  `json:"code,omitempty"`
	Description                 *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
	Criteria                    Expression        `json:"criteria,omitempty"`
}

func (r MeasureGroupStratifierComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MeasureGroupStratifierComponent) marshalJSON() jsonMeasureGroupStratifierComponent {
	m := jsonMeasureGroupStratifierComponent{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Criteria = r.Criteria
	return m
}
func (r *MeasureGroupStratifierComponent) UnmarshalJSON(b []byte) error {
	var m jsonMeasureGroupStratifierComponent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MeasureGroupStratifierComponent) unmarshalJSON(m jsonMeasureGroupStratifierComponent) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Criteria = m.Criteria
	return nil
}
func (r MeasureGroupStratifierComponent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The supplemental data criteria for the measure report, specified as either the name of a valid CQL expression within a referenced library, or a valid FHIR Resource Path.
type MeasureSupplementalData struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates a meaning for the supplemental data. This can be as simple as a unique identifier, or it can establish meaning in a broader context by drawing from a terminology, allowing supplemental data to be correlated across measures.
	Code *CodeableConcept
	// An indicator of the intended usage for the supplemental data element. Supplemental data indicates the data is additional information requested to augment the measure information. Risk adjustment factor indicates the data is additional information used to calculate risk adjustment factors when applying a risk model to the measure calculation.
	Usage []CodeableConcept
	// The human readable description of this supplemental data.
	Description *String
	// The criteria for the supplemental data. This is typically the name of a valid expression defined within a referenced library, but it may also be a path to a specific data element. The criteria defines the data to be returned for this element.
	Criteria Expression
}
type jsonMeasureSupplementalData struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Code                        *CodeableConcept  `json:"code,omitempty"`
	Usage                       []CodeableConcept `json:"usage,omitempty"`
	Description                 *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
	Criteria                    Expression        `json:"criteria,omitempty"`
}

func (r MeasureSupplementalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r MeasureSupplementalData) marshalJSON() jsonMeasureSupplementalData {
	m := jsonMeasureSupplementalData{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	m.Usage = r.Usage
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Criteria = r.Criteria
	return m
}
func (r *MeasureSupplementalData) UnmarshalJSON(b []byte) error {
	var m jsonMeasureSupplementalData
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *MeasureSupplementalData) unmarshalJSON(m jsonMeasureSupplementalData) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Usage = m.Usage
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Criteria = m.Criteria
	return nil
}
func (r MeasureSupplementalData) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
