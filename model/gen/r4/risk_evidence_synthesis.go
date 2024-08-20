package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// The RiskEvidenceSynthesis resource describes the likelihood of an outcome in a population plus exposure state where the risk estimate is derived from a combination of research studies.
type RiskEvidenceSynthesis struct {
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
	// An absolute URI that is used to identify this risk evidence synthesis when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this risk evidence synthesis is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the risk evidence synthesis is stored on different servers.
	Url *Uri
	// A formal identifier that is used to identify this risk evidence synthesis when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the risk evidence synthesis when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the risk evidence synthesis author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the risk evidence synthesis. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the risk evidence synthesis.
	Title *String
	// The status of this risk evidence synthesis. Enables tracking the life-cycle of the content.
	Status Code
	// The date  (and optionally time) when the risk evidence synthesis was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the risk evidence synthesis changes.
	Date *DateTime
	// The name of the organization or individual that published the risk evidence synthesis.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the risk evidence synthesis from a consumer's perspective.
	Description *Markdown
	// A human-readable string to clarify or explain concepts about the resource.
	Note []Annotation
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate risk evidence synthesis instances.
	UseContext []UsageContext
	// A legal or geographic region in which the risk evidence synthesis is intended to be used.
	Jurisdiction []CodeableConcept
	// A copyright statement relating to the risk evidence synthesis and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the risk evidence synthesis.
	Copyright *Markdown
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *Date
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *Date
	// The period during which the risk evidence synthesis content was or is planned to be in active use.
	EffectivePeriod *Period
	// Descriptive topics related to the content of the RiskEvidenceSynthesis. Topics provide a high-level categorization grouping types of EffectEvidenceSynthesiss that can be useful for filtering and searching.
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
	// Type of synthesis eg meta-analysis.
	SynthesisType *CodeableConcept
	// Type of study eg randomized trial.
	StudyType *CodeableConcept
	// A reference to a EvidenceVariable resource that defines the population for the research.
	Population Reference
	// A reference to a EvidenceVariable resource that defines the exposure for the research.
	Exposure *Reference
	// A reference to a EvidenceVariable resomece that defines the outcome for the research.
	Outcome Reference
	// A description of the size of the sample involved in the synthesis.
	SampleSize *RiskEvidenceSynthesisSampleSize
	// The estimated risk of the outcome.
	RiskEstimate *RiskEvidenceSynthesisRiskEstimate
	// A description of the certainty of the risk estimate.
	Certainty []RiskEvidenceSynthesisCertainty
}

func (r RiskEvidenceSynthesis) ResourceType() string {
	return "RiskEvidenceSynthesis"
}
func (r RiskEvidenceSynthesis) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonRiskEvidenceSynthesis struct {
	ResourceType                   string                             `json:"resourceType"`
	Id                             *Id                                `json:"id,omitempty"`
	IdPrimitiveElement             *primitiveElement                  `json:"_id,omitempty"`
	Meta                           *Meta                              `json:"meta,omitempty"`
	ImplicitRules                  *Uri                               `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement  *primitiveElement                  `json:"_implicitRules,omitempty"`
	Language                       *Code                              `json:"language,omitempty"`
	LanguagePrimitiveElement       *primitiveElement                  `json:"_language,omitempty"`
	Text                           *Narrative                         `json:"text,omitempty"`
	Contained                      []ContainedResource                `json:"contained,omitempty"`
	Extension                      []Extension                        `json:"extension,omitempty"`
	ModifierExtension              []Extension                        `json:"modifierExtension,omitempty"`
	Url                            *Uri                               `json:"url,omitempty"`
	UrlPrimitiveElement            *primitiveElement                  `json:"_url,omitempty"`
	Identifier                     []Identifier                       `json:"identifier,omitempty"`
	Version                        *String                            `json:"version,omitempty"`
	VersionPrimitiveElement        *primitiveElement                  `json:"_version,omitempty"`
	Name                           *String                            `json:"name,omitempty"`
	NamePrimitiveElement           *primitiveElement                  `json:"_name,omitempty"`
	Title                          *String                            `json:"title,omitempty"`
	TitlePrimitiveElement          *primitiveElement                  `json:"_title,omitempty"`
	Status                         Code                               `json:"status,omitempty"`
	StatusPrimitiveElement         *primitiveElement                  `json:"_status,omitempty"`
	Date                           *DateTime                          `json:"date,omitempty"`
	DatePrimitiveElement           *primitiveElement                  `json:"_date,omitempty"`
	Publisher                      *String                            `json:"publisher,omitempty"`
	PublisherPrimitiveElement      *primitiveElement                  `json:"_publisher,omitempty"`
	Contact                        []ContactDetail                    `json:"contact,omitempty"`
	Description                    *Markdown                          `json:"description,omitempty"`
	DescriptionPrimitiveElement    *primitiveElement                  `json:"_description,omitempty"`
	Note                           []Annotation                       `json:"note,omitempty"`
	UseContext                     []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction                   []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Copyright                      *Markdown                          `json:"copyright,omitempty"`
	CopyrightPrimitiveElement      *primitiveElement                  `json:"_copyright,omitempty"`
	ApprovalDate                   *Date                              `json:"approvalDate,omitempty"`
	ApprovalDatePrimitiveElement   *primitiveElement                  `json:"_approvalDate,omitempty"`
	LastReviewDate                 *Date                              `json:"lastReviewDate,omitempty"`
	LastReviewDatePrimitiveElement *primitiveElement                  `json:"_lastReviewDate,omitempty"`
	EffectivePeriod                *Period                            `json:"effectivePeriod,omitempty"`
	Topic                          []CodeableConcept                  `json:"topic,omitempty"`
	Author                         []ContactDetail                    `json:"author,omitempty"`
	Editor                         []ContactDetail                    `json:"editor,omitempty"`
	Reviewer                       []ContactDetail                    `json:"reviewer,omitempty"`
	Endorser                       []ContactDetail                    `json:"endorser,omitempty"`
	RelatedArtifact                []RelatedArtifact                  `json:"relatedArtifact,omitempty"`
	SynthesisType                  *CodeableConcept                   `json:"synthesisType,omitempty"`
	StudyType                      *CodeableConcept                   `json:"studyType,omitempty"`
	Population                     Reference                          `json:"population,omitempty"`
	Exposure                       *Reference                         `json:"exposure,omitempty"`
	Outcome                        Reference                          `json:"outcome,omitempty"`
	SampleSize                     *RiskEvidenceSynthesisSampleSize   `json:"sampleSize,omitempty"`
	RiskEstimate                   *RiskEvidenceSynthesisRiskEstimate `json:"riskEstimate,omitempty"`
	Certainty                      []RiskEvidenceSynthesisCertainty   `json:"certainty,omitempty"`
}

func (r RiskEvidenceSynthesis) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RiskEvidenceSynthesis) marshalJSON() jsonRiskEvidenceSynthesis {
	m := jsonRiskEvidenceSynthesis{}
	m.ResourceType = "RiskEvidenceSynthesis"
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
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
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
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
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
	m.Note = r.Note
	m.UseContext = r.UseContext
	m.Jurisdiction = r.Jurisdiction
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
	m.SynthesisType = r.SynthesisType
	m.StudyType = r.StudyType
	m.Population = r.Population
	m.Exposure = r.Exposure
	m.Outcome = r.Outcome
	m.SampleSize = r.SampleSize
	m.RiskEstimate = r.RiskEstimate
	m.Certainty = r.Certainty
	return m
}
func (r *RiskEvidenceSynthesis) UnmarshalJSON(b []byte) error {
	var m jsonRiskEvidenceSynthesis
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RiskEvidenceSynthesis) unmarshalJSON(m jsonRiskEvidenceSynthesis) error {
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
		r.Contained = append(r.Contained, v.Resource)
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
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
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
	r.Note = m.Note
	r.UseContext = m.UseContext
	r.Jurisdiction = m.Jurisdiction
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
	r.SynthesisType = m.SynthesisType
	r.StudyType = m.StudyType
	r.Population = m.Population
	r.Exposure = m.Exposure
	r.Outcome = m.Outcome
	r.SampleSize = m.SampleSize
	r.RiskEstimate = m.RiskEstimate
	r.Certainty = m.Certainty
	return nil
}
func (r RiskEvidenceSynthesis) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A description of the size of the sample involved in the synthesis.
type RiskEvidenceSynthesisSampleSize struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Human-readable summary of sample size.
	Description *String
	// Number of studies included in this evidence synthesis.
	NumberOfStudies *Integer
	// Number of participants included in this evidence synthesis.
	NumberOfParticipants *Integer
}
type jsonRiskEvidenceSynthesisSampleSize struct {
	Id                                   *string           `json:"id,omitempty"`
	Extension                            []Extension       `json:"extension,omitempty"`
	ModifierExtension                    []Extension       `json:"modifierExtension,omitempty"`
	Description                          *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement          *primitiveElement `json:"_description,omitempty"`
	NumberOfStudies                      *Integer          `json:"numberOfStudies,omitempty"`
	NumberOfStudiesPrimitiveElement      *primitiveElement `json:"_numberOfStudies,omitempty"`
	NumberOfParticipants                 *Integer          `json:"numberOfParticipants,omitempty"`
	NumberOfParticipantsPrimitiveElement *primitiveElement `json:"_numberOfParticipants,omitempty"`
}

func (r RiskEvidenceSynthesisSampleSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RiskEvidenceSynthesisSampleSize) marshalJSON() jsonRiskEvidenceSynthesisSampleSize {
	m := jsonRiskEvidenceSynthesisSampleSize{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.NumberOfStudies = r.NumberOfStudies
	if r.NumberOfStudies != nil && (r.NumberOfStudies.Id != nil || r.NumberOfStudies.Extension != nil) {
		m.NumberOfStudiesPrimitiveElement = &primitiveElement{Id: r.NumberOfStudies.Id, Extension: r.NumberOfStudies.Extension}
	}
	m.NumberOfParticipants = r.NumberOfParticipants
	if r.NumberOfParticipants != nil && (r.NumberOfParticipants.Id != nil || r.NumberOfParticipants.Extension != nil) {
		m.NumberOfParticipantsPrimitiveElement = &primitiveElement{Id: r.NumberOfParticipants.Id, Extension: r.NumberOfParticipants.Extension}
	}
	return m
}
func (r *RiskEvidenceSynthesisSampleSize) UnmarshalJSON(b []byte) error {
	var m jsonRiskEvidenceSynthesisSampleSize
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RiskEvidenceSynthesisSampleSize) unmarshalJSON(m jsonRiskEvidenceSynthesisSampleSize) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.NumberOfStudies = m.NumberOfStudies
	if m.NumberOfStudiesPrimitiveElement != nil {
		r.NumberOfStudies.Id = m.NumberOfStudiesPrimitiveElement.Id
		r.NumberOfStudies.Extension = m.NumberOfStudiesPrimitiveElement.Extension
	}
	r.NumberOfParticipants = m.NumberOfParticipants
	if m.NumberOfParticipantsPrimitiveElement != nil {
		r.NumberOfParticipants.Id = m.NumberOfParticipantsPrimitiveElement.Id
		r.NumberOfParticipants.Extension = m.NumberOfParticipantsPrimitiveElement.Extension
	}
	return nil
}
func (r RiskEvidenceSynthesisSampleSize) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The estimated risk of the outcome.
type RiskEvidenceSynthesisRiskEstimate struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Human-readable summary of risk estimate.
	Description *String
	// Examples include proportion and mean.
	Type *CodeableConcept
	// The point estimate of the risk estimate.
	Value *Decimal
	// Specifies the UCUM unit for the outcome.
	UnitOfMeasure *CodeableConcept
	// The sample size for the group that was measured for this risk estimate.
	DenominatorCount *Integer
	// The number of group members with the outcome of interest.
	NumeratorCount *Integer
	// A description of the precision of the estimate for the effect.
	PrecisionEstimate []RiskEvidenceSynthesisRiskEstimatePrecisionEstimate
}
type jsonRiskEvidenceSynthesisRiskEstimate struct {
	Id                               *string                                              `json:"id,omitempty"`
	Extension                        []Extension                                          `json:"extension,omitempty"`
	ModifierExtension                []Extension                                          `json:"modifierExtension,omitempty"`
	Description                      *String                                              `json:"description,omitempty"`
	DescriptionPrimitiveElement      *primitiveElement                                    `json:"_description,omitempty"`
	Type                             *CodeableConcept                                     `json:"type,omitempty"`
	Value                            *Decimal                                             `json:"value,omitempty"`
	ValuePrimitiveElement            *primitiveElement                                    `json:"_value,omitempty"`
	UnitOfMeasure                    *CodeableConcept                                     `json:"unitOfMeasure,omitempty"`
	DenominatorCount                 *Integer                                             `json:"denominatorCount,omitempty"`
	DenominatorCountPrimitiveElement *primitiveElement                                    `json:"_denominatorCount,omitempty"`
	NumeratorCount                   *Integer                                             `json:"numeratorCount,omitempty"`
	NumeratorCountPrimitiveElement   *primitiveElement                                    `json:"_numeratorCount,omitempty"`
	PrecisionEstimate                []RiskEvidenceSynthesisRiskEstimatePrecisionEstimate `json:"precisionEstimate,omitempty"`
}

func (r RiskEvidenceSynthesisRiskEstimate) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RiskEvidenceSynthesisRiskEstimate) marshalJSON() jsonRiskEvidenceSynthesisRiskEstimate {
	m := jsonRiskEvidenceSynthesisRiskEstimate{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Type = r.Type
	m.Value = r.Value
	if r.Value != nil && (r.Value.Id != nil || r.Value.Extension != nil) {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	m.UnitOfMeasure = r.UnitOfMeasure
	m.DenominatorCount = r.DenominatorCount
	if r.DenominatorCount != nil && (r.DenominatorCount.Id != nil || r.DenominatorCount.Extension != nil) {
		m.DenominatorCountPrimitiveElement = &primitiveElement{Id: r.DenominatorCount.Id, Extension: r.DenominatorCount.Extension}
	}
	m.NumeratorCount = r.NumeratorCount
	if r.NumeratorCount != nil && (r.NumeratorCount.Id != nil || r.NumeratorCount.Extension != nil) {
		m.NumeratorCountPrimitiveElement = &primitiveElement{Id: r.NumeratorCount.Id, Extension: r.NumeratorCount.Extension}
	}
	m.PrecisionEstimate = r.PrecisionEstimate
	return m
}
func (r *RiskEvidenceSynthesisRiskEstimate) UnmarshalJSON(b []byte) error {
	var m jsonRiskEvidenceSynthesisRiskEstimate
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RiskEvidenceSynthesisRiskEstimate) unmarshalJSON(m jsonRiskEvidenceSynthesisRiskEstimate) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	r.UnitOfMeasure = m.UnitOfMeasure
	r.DenominatorCount = m.DenominatorCount
	if m.DenominatorCountPrimitiveElement != nil {
		r.DenominatorCount.Id = m.DenominatorCountPrimitiveElement.Id
		r.DenominatorCount.Extension = m.DenominatorCountPrimitiveElement.Extension
	}
	r.NumeratorCount = m.NumeratorCount
	if m.NumeratorCountPrimitiveElement != nil {
		r.NumeratorCount.Id = m.NumeratorCountPrimitiveElement.Id
		r.NumeratorCount.Extension = m.NumeratorCountPrimitiveElement.Extension
	}
	r.PrecisionEstimate = m.PrecisionEstimate
	return nil
}
func (r RiskEvidenceSynthesisRiskEstimate) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A description of the precision of the estimate for the effect.
type RiskEvidenceSynthesisRiskEstimatePrecisionEstimate struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Examples include confidence interval and interquartile range.
	Type *CodeableConcept
	// Use 95 for a 95% confidence interval.
	Level *Decimal
	// Lower bound of confidence interval.
	From *Decimal
	// Upper bound of confidence interval.
	To *Decimal
}
type jsonRiskEvidenceSynthesisRiskEstimatePrecisionEstimate struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Type                  *CodeableConcept  `json:"type,omitempty"`
	Level                 *Decimal          `json:"level,omitempty"`
	LevelPrimitiveElement *primitiveElement `json:"_level,omitempty"`
	From                  *Decimal          `json:"from,omitempty"`
	FromPrimitiveElement  *primitiveElement `json:"_from,omitempty"`
	To                    *Decimal          `json:"to,omitempty"`
	ToPrimitiveElement    *primitiveElement `json:"_to,omitempty"`
}

func (r RiskEvidenceSynthesisRiskEstimatePrecisionEstimate) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RiskEvidenceSynthesisRiskEstimatePrecisionEstimate) marshalJSON() jsonRiskEvidenceSynthesisRiskEstimatePrecisionEstimate {
	m := jsonRiskEvidenceSynthesisRiskEstimatePrecisionEstimate{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Level = r.Level
	if r.Level != nil && (r.Level.Id != nil || r.Level.Extension != nil) {
		m.LevelPrimitiveElement = &primitiveElement{Id: r.Level.Id, Extension: r.Level.Extension}
	}
	m.From = r.From
	if r.From != nil && (r.From.Id != nil || r.From.Extension != nil) {
		m.FromPrimitiveElement = &primitiveElement{Id: r.From.Id, Extension: r.From.Extension}
	}
	m.To = r.To
	if r.To != nil && (r.To.Id != nil || r.To.Extension != nil) {
		m.ToPrimitiveElement = &primitiveElement{Id: r.To.Id, Extension: r.To.Extension}
	}
	return m
}
func (r *RiskEvidenceSynthesisRiskEstimatePrecisionEstimate) UnmarshalJSON(b []byte) error {
	var m jsonRiskEvidenceSynthesisRiskEstimatePrecisionEstimate
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RiskEvidenceSynthesisRiskEstimatePrecisionEstimate) unmarshalJSON(m jsonRiskEvidenceSynthesisRiskEstimatePrecisionEstimate) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Level = m.Level
	if m.LevelPrimitiveElement != nil {
		r.Level.Id = m.LevelPrimitiveElement.Id
		r.Level.Extension = m.LevelPrimitiveElement.Extension
	}
	r.From = m.From
	if m.FromPrimitiveElement != nil {
		r.From.Id = m.FromPrimitiveElement.Id
		r.From.Extension = m.FromPrimitiveElement.Extension
	}
	r.To = m.To
	if m.ToPrimitiveElement != nil {
		r.To.Id = m.ToPrimitiveElement.Id
		r.To.Extension = m.ToPrimitiveElement.Extension
	}
	return nil
}
func (r RiskEvidenceSynthesisRiskEstimatePrecisionEstimate) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A description of the certainty of the risk estimate.
type RiskEvidenceSynthesisCertainty struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A rating of the certainty of the effect estimate.
	Rating []CodeableConcept
	// A human-readable string to clarify or explain concepts about the resource.
	Note []Annotation
	// A description of a component of the overall certainty.
	CertaintySubcomponent []RiskEvidenceSynthesisCertaintyCertaintySubcomponent
}
type jsonRiskEvidenceSynthesisCertainty struct {
	Id                    *string                                               `json:"id,omitempty"`
	Extension             []Extension                                           `json:"extension,omitempty"`
	ModifierExtension     []Extension                                           `json:"modifierExtension,omitempty"`
	Rating                []CodeableConcept                                     `json:"rating,omitempty"`
	Note                  []Annotation                                          `json:"note,omitempty"`
	CertaintySubcomponent []RiskEvidenceSynthesisCertaintyCertaintySubcomponent `json:"certaintySubcomponent,omitempty"`
}

func (r RiskEvidenceSynthesisCertainty) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RiskEvidenceSynthesisCertainty) marshalJSON() jsonRiskEvidenceSynthesisCertainty {
	m := jsonRiskEvidenceSynthesisCertainty{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Rating = r.Rating
	m.Note = r.Note
	m.CertaintySubcomponent = r.CertaintySubcomponent
	return m
}
func (r *RiskEvidenceSynthesisCertainty) UnmarshalJSON(b []byte) error {
	var m jsonRiskEvidenceSynthesisCertainty
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RiskEvidenceSynthesisCertainty) unmarshalJSON(m jsonRiskEvidenceSynthesisCertainty) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Rating = m.Rating
	r.Note = m.Note
	r.CertaintySubcomponent = m.CertaintySubcomponent
	return nil
}
func (r RiskEvidenceSynthesisCertainty) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A description of a component of the overall certainty.
type RiskEvidenceSynthesisCertaintyCertaintySubcomponent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Type of subcomponent of certainty rating.
	Type *CodeableConcept
	// A rating of a subcomponent of rating certainty.
	Rating []CodeableConcept
	// A human-readable string to clarify or explain concepts about the resource.
	Note []Annotation
}
type jsonRiskEvidenceSynthesisCertaintyCertaintySubcomponent struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Rating            []CodeableConcept `json:"rating,omitempty"`
	Note              []Annotation      `json:"note,omitempty"`
}

func (r RiskEvidenceSynthesisCertaintyCertaintySubcomponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r RiskEvidenceSynthesisCertaintyCertaintySubcomponent) marshalJSON() jsonRiskEvidenceSynthesisCertaintyCertaintySubcomponent {
	m := jsonRiskEvidenceSynthesisCertaintyCertaintySubcomponent{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Rating = r.Rating
	m.Note = r.Note
	return m
}
func (r *RiskEvidenceSynthesisCertaintyCertaintySubcomponent) UnmarshalJSON(b []byte) error {
	var m jsonRiskEvidenceSynthesisCertaintyCertaintySubcomponent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *RiskEvidenceSynthesisCertaintyCertaintySubcomponent) unmarshalJSON(m jsonRiskEvidenceSynthesisCertaintyCertaintySubcomponent) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Rating = m.Rating
	r.Note = m.Note
	return nil
}
func (r RiskEvidenceSynthesisCertaintyCertaintySubcomponent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
