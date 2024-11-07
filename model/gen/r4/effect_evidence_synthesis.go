package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// The EffectEvidenceSynthesis resource describes the difference in an outcome between exposures states in a population where the effect estimate is derived from a combination of research studies.
type EffectEvidenceSynthesis struct {
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
	// An absolute URI that is used to identify this effect evidence synthesis when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this effect evidence synthesis is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the effect evidence synthesis is stored on different servers.
	Url *Uri
	// A formal identifier that is used to identify this effect evidence synthesis when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the effect evidence synthesis when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the effect evidence synthesis author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the effect evidence synthesis. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the effect evidence synthesis.
	Title *String
	// The status of this effect evidence synthesis. Enables tracking the life-cycle of the content.
	Status Code
	// The date  (and optionally time) when the effect evidence synthesis was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the effect evidence synthesis changes.
	Date *DateTime
	// The name of the organization or individual that published the effect evidence synthesis.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the effect evidence synthesis from a consumer's perspective.
	Description *Markdown
	// A human-readable string to clarify or explain concepts about the resource.
	Note []Annotation
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate effect evidence synthesis instances.
	UseContext []UsageContext
	// A legal or geographic region in which the effect evidence synthesis is intended to be used.
	Jurisdiction []CodeableConcept
	// A copyright statement relating to the effect evidence synthesis and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the effect evidence synthesis.
	Copyright *Markdown
	// The date on which the resource content was approved by the publisher. Approval happens once when the content is officially approved for usage.
	ApprovalDate *Date
	// The date on which the resource content was last reviewed. Review happens periodically after approval but does not change the original approval date.
	LastReviewDate *Date
	// The period during which the effect evidence synthesis content was or is planned to be in active use.
	EffectivePeriod *Period
	// Descriptive topics related to the content of the EffectEvidenceSynthesis. Topics provide a high-level categorization grouping types of EffectEvidenceSynthesiss that can be useful for filtering and searching.
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
	Exposure Reference
	// A reference to a EvidenceVariable resource that defines the comparison exposure for the research.
	ExposureAlternative Reference
	// A reference to a EvidenceVariable resomece that defines the outcome for the research.
	Outcome Reference
	// A description of the size of the sample involved in the synthesis.
	SampleSize *EffectEvidenceSynthesisSampleSize
	// A description of the results for each exposure considered in the effect estimate.
	ResultsByExposure []EffectEvidenceSynthesisResultsByExposure
	// The estimated effect of the exposure variant.
	EffectEstimate []EffectEvidenceSynthesisEffectEstimate
	// A description of the certainty of the effect estimate.
	Certainty []EffectEvidenceSynthesisCertainty
}

func (r EffectEvidenceSynthesis) ResourceType() string {
	return "EffectEvidenceSynthesis"
}
func (r EffectEvidenceSynthesis) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonEffectEvidenceSynthesis struct {
	ResourceType                   string                                     `json:"resourceType"`
	Id                             *Id                                        `json:"id,omitempty"`
	IdPrimitiveElement             *primitiveElement                          `json:"_id,omitempty"`
	Meta                           *Meta                                      `json:"meta,omitempty"`
	ImplicitRules                  *Uri                                       `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement  *primitiveElement                          `json:"_implicitRules,omitempty"`
	Language                       *Code                                      `json:"language,omitempty"`
	LanguagePrimitiveElement       *primitiveElement                          `json:"_language,omitempty"`
	Text                           *Narrative                                 `json:"text,omitempty"`
	Contained                      []ContainedResource                        `json:"contained,omitempty"`
	Extension                      []Extension                                `json:"extension,omitempty"`
	ModifierExtension              []Extension                                `json:"modifierExtension,omitempty"`
	Url                            *Uri                                       `json:"url,omitempty"`
	UrlPrimitiveElement            *primitiveElement                          `json:"_url,omitempty"`
	Identifier                     []Identifier                               `json:"identifier,omitempty"`
	Version                        *String                                    `json:"version,omitempty"`
	VersionPrimitiveElement        *primitiveElement                          `json:"_version,omitempty"`
	Name                           *String                                    `json:"name,omitempty"`
	NamePrimitiveElement           *primitiveElement                          `json:"_name,omitempty"`
	Title                          *String                                    `json:"title,omitempty"`
	TitlePrimitiveElement          *primitiveElement                          `json:"_title,omitempty"`
	Status                         Code                                       `json:"status,omitempty"`
	StatusPrimitiveElement         *primitiveElement                          `json:"_status,omitempty"`
	Date                           *DateTime                                  `json:"date,omitempty"`
	DatePrimitiveElement           *primitiveElement                          `json:"_date,omitempty"`
	Publisher                      *String                                    `json:"publisher,omitempty"`
	PublisherPrimitiveElement      *primitiveElement                          `json:"_publisher,omitempty"`
	Contact                        []ContactDetail                            `json:"contact,omitempty"`
	Description                    *Markdown                                  `json:"description,omitempty"`
	DescriptionPrimitiveElement    *primitiveElement                          `json:"_description,omitempty"`
	Note                           []Annotation                               `json:"note,omitempty"`
	UseContext                     []UsageContext                             `json:"useContext,omitempty"`
	Jurisdiction                   []CodeableConcept                          `json:"jurisdiction,omitempty"`
	Copyright                      *Markdown                                  `json:"copyright,omitempty"`
	CopyrightPrimitiveElement      *primitiveElement                          `json:"_copyright,omitempty"`
	ApprovalDate                   *Date                                      `json:"approvalDate,omitempty"`
	ApprovalDatePrimitiveElement   *primitiveElement                          `json:"_approvalDate,omitempty"`
	LastReviewDate                 *Date                                      `json:"lastReviewDate,omitempty"`
	LastReviewDatePrimitiveElement *primitiveElement                          `json:"_lastReviewDate,omitempty"`
	EffectivePeriod                *Period                                    `json:"effectivePeriod,omitempty"`
	Topic                          []CodeableConcept                          `json:"topic,omitempty"`
	Author                         []ContactDetail                            `json:"author,omitempty"`
	Editor                         []ContactDetail                            `json:"editor,omitempty"`
	Reviewer                       []ContactDetail                            `json:"reviewer,omitempty"`
	Endorser                       []ContactDetail                            `json:"endorser,omitempty"`
	RelatedArtifact                []RelatedArtifact                          `json:"relatedArtifact,omitempty"`
	SynthesisType                  *CodeableConcept                           `json:"synthesisType,omitempty"`
	StudyType                      *CodeableConcept                           `json:"studyType,omitempty"`
	Population                     Reference                                  `json:"population,omitempty"`
	Exposure                       Reference                                  `json:"exposure,omitempty"`
	ExposureAlternative            Reference                                  `json:"exposureAlternative,omitempty"`
	Outcome                        Reference                                  `json:"outcome,omitempty"`
	SampleSize                     *EffectEvidenceSynthesisSampleSize         `json:"sampleSize,omitempty"`
	ResultsByExposure              []EffectEvidenceSynthesisResultsByExposure `json:"resultsByExposure,omitempty"`
	EffectEstimate                 []EffectEvidenceSynthesisEffectEstimate    `json:"effectEstimate,omitempty"`
	Certainty                      []EffectEvidenceSynthesisCertainty         `json:"certainty,omitempty"`
}

func (r EffectEvidenceSynthesis) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EffectEvidenceSynthesis) marshalJSON() jsonEffectEvidenceSynthesis {
	m := jsonEffectEvidenceSynthesis{}
	m.ResourceType = "EffectEvidenceSynthesis"
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
	m.Text = r.Text
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Url != nil && r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	m.Identifier = r.Identifier
	if r.Version != nil && r.Version.Value != nil {
		m.Version = r.Version
	}
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Title != nil && r.Title.Value != nil {
		m.Title = r.Title
	}
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	if r.Date != nil && r.Date.Value != nil {
		m.Date = r.Date
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	if r.Publisher != nil && r.Publisher.Value != nil {
		m.Publisher = r.Publisher
	}
	if r.Publisher != nil && (r.Publisher.Id != nil || r.Publisher.Extension != nil) {
		m.PublisherPrimitiveElement = &primitiveElement{Id: r.Publisher.Id, Extension: r.Publisher.Extension}
	}
	m.Contact = r.Contact
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Note = r.Note
	m.UseContext = r.UseContext
	m.Jurisdiction = r.Jurisdiction
	if r.Copyright != nil && r.Copyright.Value != nil {
		m.Copyright = r.Copyright
	}
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		m.CopyrightPrimitiveElement = &primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
	}
	if r.ApprovalDate != nil && r.ApprovalDate.Value != nil {
		m.ApprovalDate = r.ApprovalDate
	}
	if r.ApprovalDate != nil && (r.ApprovalDate.Id != nil || r.ApprovalDate.Extension != nil) {
		m.ApprovalDatePrimitiveElement = &primitiveElement{Id: r.ApprovalDate.Id, Extension: r.ApprovalDate.Extension}
	}
	if r.LastReviewDate != nil && r.LastReviewDate.Value != nil {
		m.LastReviewDate = r.LastReviewDate
	}
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
	m.ExposureAlternative = r.ExposureAlternative
	m.Outcome = r.Outcome
	m.SampleSize = r.SampleSize
	m.ResultsByExposure = r.ResultsByExposure
	m.EffectEstimate = r.EffectEstimate
	m.Certainty = r.Certainty
	return m
}
func (r *EffectEvidenceSynthesis) UnmarshalJSON(b []byte) error {
	var m jsonEffectEvidenceSynthesis
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EffectEvidenceSynthesis) unmarshalJSON(m jsonEffectEvidenceSynthesis) error {
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
	r.Text = m.Text
	r.Contained = make([]model.Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.Resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		if r.Url == nil {
			r.Url = &Uri{}
		}
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Identifier = m.Identifier
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		if r.Version == nil {
			r.Version = &String{}
		}
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		if r.Title == nil {
			r.Title = &String{}
		}
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
		if r.Date == nil {
			r.Date = &DateTime{}
		}
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Publisher = m.Publisher
	if m.PublisherPrimitiveElement != nil {
		if r.Publisher == nil {
			r.Publisher = &String{}
		}
		r.Publisher.Id = m.PublisherPrimitiveElement.Id
		r.Publisher.Extension = m.PublisherPrimitiveElement.Extension
	}
	r.Contact = m.Contact
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &Markdown{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Note = m.Note
	r.UseContext = m.UseContext
	r.Jurisdiction = m.Jurisdiction
	r.Copyright = m.Copyright
	if m.CopyrightPrimitiveElement != nil {
		if r.Copyright == nil {
			r.Copyright = &Markdown{}
		}
		r.Copyright.Id = m.CopyrightPrimitiveElement.Id
		r.Copyright.Extension = m.CopyrightPrimitiveElement.Extension
	}
	r.ApprovalDate = m.ApprovalDate
	if m.ApprovalDatePrimitiveElement != nil {
		if r.ApprovalDate == nil {
			r.ApprovalDate = &Date{}
		}
		r.ApprovalDate.Id = m.ApprovalDatePrimitiveElement.Id
		r.ApprovalDate.Extension = m.ApprovalDatePrimitiveElement.Extension
	}
	r.LastReviewDate = m.LastReviewDate
	if m.LastReviewDatePrimitiveElement != nil {
		if r.LastReviewDate == nil {
			r.LastReviewDate = &Date{}
		}
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
	r.ExposureAlternative = m.ExposureAlternative
	r.Outcome = m.Outcome
	r.SampleSize = m.SampleSize
	r.ResultsByExposure = m.ResultsByExposure
	r.EffectEstimate = m.EffectEstimate
	r.Certainty = m.Certainty
	return nil
}
func (r EffectEvidenceSynthesis) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "EffectEvidenceSynthesis"
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
	v := make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		v = append(v, ContainedResource{c})
	}
	err = e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "contained"}})
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
	err = e.EncodeElement(r.Url, xml.StartElement{Name: xml.Name{Local: "url"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Version, xml.StartElement{Name: xml.Name{Local: "version"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Title, xml.StartElement{Name: xml.Name{Local: "title"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Publisher, xml.StartElement{Name: xml.Name{Local: "publisher"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Contact, xml.StartElement{Name: xml.Name{Local: "contact"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UseContext, xml.StartElement{Name: xml.Name{Local: "useContext"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Jurisdiction, xml.StartElement{Name: xml.Name{Local: "jurisdiction"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Copyright, xml.StartElement{Name: xml.Name{Local: "copyright"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ApprovalDate, xml.StartElement{Name: xml.Name{Local: "approvalDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.LastReviewDate, xml.StartElement{Name: xml.Name{Local: "lastReviewDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.EffectivePeriod, xml.StartElement{Name: xml.Name{Local: "effectivePeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Topic, xml.StartElement{Name: xml.Name{Local: "topic"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Author, xml.StartElement{Name: xml.Name{Local: "author"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Editor, xml.StartElement{Name: xml.Name{Local: "editor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Reviewer, xml.StartElement{Name: xml.Name{Local: "reviewer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Endorser, xml.StartElement{Name: xml.Name{Local: "endorser"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RelatedArtifact, xml.StartElement{Name: xml.Name{Local: "relatedArtifact"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SynthesisType, xml.StartElement{Name: xml.Name{Local: "synthesisType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.StudyType, xml.StartElement{Name: xml.Name{Local: "studyType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Population, xml.StartElement{Name: xml.Name{Local: "population"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Exposure, xml.StartElement{Name: xml.Name{Local: "exposure"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ExposureAlternative, xml.StartElement{Name: xml.Name{Local: "exposureAlternative"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Outcome, xml.StartElement{Name: xml.Name{Local: "outcome"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SampleSize, xml.StartElement{Name: xml.Name{Local: "sampleSize"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ResultsByExposure, xml.StartElement{Name: xml.Name{Local: "resultsByExposure"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.EffectEstimate, xml.StartElement{Name: xml.Name{Local: "effectEstimate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Certainty, xml.StartElement{Name: xml.Name{Local: "certainty"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *EffectEvidenceSynthesis) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "url":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Url = &v
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "version":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Version = &v
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "title":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Title = &v
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "date":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = &v
			case "publisher":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Publisher = &v
			case "contact":
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Contact = append(r.Contact, v)
			case "description":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			case "useContext":
				var v UsageContext
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UseContext = append(r.UseContext, v)
			case "jurisdiction":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Jurisdiction = append(r.Jurisdiction, v)
			case "copyright":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Copyright = &v
			case "approvalDate":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ApprovalDate = &v
			case "lastReviewDate":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.LastReviewDate = &v
			case "effectivePeriod":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.EffectivePeriod = &v
			case "topic":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Topic = append(r.Topic, v)
			case "author":
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Author = append(r.Author, v)
			case "editor":
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Editor = append(r.Editor, v)
			case "reviewer":
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reviewer = append(r.Reviewer, v)
			case "endorser":
				var v ContactDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Endorser = append(r.Endorser, v)
			case "relatedArtifact":
				var v RelatedArtifact
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RelatedArtifact = append(r.RelatedArtifact, v)
			case "synthesisType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SynthesisType = &v
			case "studyType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.StudyType = &v
			case "population":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Population = v
			case "exposure":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Exposure = v
			case "exposureAlternative":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ExposureAlternative = v
			case "outcome":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Outcome = v
			case "sampleSize":
				var v EffectEvidenceSynthesisSampleSize
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SampleSize = &v
			case "resultsByExposure":
				var v EffectEvidenceSynthesisResultsByExposure
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ResultsByExposure = append(r.ResultsByExposure, v)
			case "effectEstimate":
				var v EffectEvidenceSynthesisEffectEstimate
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.EffectEstimate = append(r.EffectEstimate, v)
			case "certainty":
				var v EffectEvidenceSynthesisCertainty
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Certainty = append(r.Certainty, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r EffectEvidenceSynthesis) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A description of the size of the sample involved in the synthesis.
type EffectEvidenceSynthesisSampleSize struct {
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
type jsonEffectEvidenceSynthesisSampleSize struct {
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

func (r EffectEvidenceSynthesisSampleSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EffectEvidenceSynthesisSampleSize) marshalJSON() jsonEffectEvidenceSynthesisSampleSize {
	m := jsonEffectEvidenceSynthesisSampleSize{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	if r.NumberOfStudies != nil && r.NumberOfStudies.Value != nil {
		m.NumberOfStudies = r.NumberOfStudies
	}
	if r.NumberOfStudies != nil && (r.NumberOfStudies.Id != nil || r.NumberOfStudies.Extension != nil) {
		m.NumberOfStudiesPrimitiveElement = &primitiveElement{Id: r.NumberOfStudies.Id, Extension: r.NumberOfStudies.Extension}
	}
	if r.NumberOfParticipants != nil && r.NumberOfParticipants.Value != nil {
		m.NumberOfParticipants = r.NumberOfParticipants
	}
	if r.NumberOfParticipants != nil && (r.NumberOfParticipants.Id != nil || r.NumberOfParticipants.Extension != nil) {
		m.NumberOfParticipantsPrimitiveElement = &primitiveElement{Id: r.NumberOfParticipants.Id, Extension: r.NumberOfParticipants.Extension}
	}
	return m
}
func (r *EffectEvidenceSynthesisSampleSize) UnmarshalJSON(b []byte) error {
	var m jsonEffectEvidenceSynthesisSampleSize
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EffectEvidenceSynthesisSampleSize) unmarshalJSON(m jsonEffectEvidenceSynthesisSampleSize) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.NumberOfStudies = m.NumberOfStudies
	if m.NumberOfStudiesPrimitiveElement != nil {
		if r.NumberOfStudies == nil {
			r.NumberOfStudies = &Integer{}
		}
		r.NumberOfStudies.Id = m.NumberOfStudiesPrimitiveElement.Id
		r.NumberOfStudies.Extension = m.NumberOfStudiesPrimitiveElement.Extension
	}
	r.NumberOfParticipants = m.NumberOfParticipants
	if m.NumberOfParticipantsPrimitiveElement != nil {
		if r.NumberOfParticipants == nil {
			r.NumberOfParticipants = &Integer{}
		}
		r.NumberOfParticipants.Id = m.NumberOfParticipantsPrimitiveElement.Id
		r.NumberOfParticipants.Extension = m.NumberOfParticipantsPrimitiveElement.Extension
	}
	return nil
}
func (r EffectEvidenceSynthesisSampleSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NumberOfStudies, xml.StartElement{Name: xml.Name{Local: "numberOfStudies"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NumberOfParticipants, xml.StartElement{Name: xml.Name{Local: "numberOfParticipants"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *EffectEvidenceSynthesisSampleSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "numberOfStudies":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NumberOfStudies = &v
			case "numberOfParticipants":
				var v Integer
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NumberOfParticipants = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r EffectEvidenceSynthesisSampleSize) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A description of the results for each exposure considered in the effect estimate.
type EffectEvidenceSynthesisResultsByExposure struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Human-readable summary of results by exposure state.
	Description *String
	// Whether these results are for the exposure state or alternative exposure state.
	ExposureState *Code
	// Used to define variant exposure states such as low-risk state.
	VariantState *CodeableConcept
	// Reference to a RiskEvidenceSynthesis resource.
	RiskEvidenceSynthesis Reference
}
type jsonEffectEvidenceSynthesisResultsByExposure struct {
	Id                            *string           `json:"id,omitempty"`
	Extension                     []Extension       `json:"extension,omitempty"`
	ModifierExtension             []Extension       `json:"modifierExtension,omitempty"`
	Description                   *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement `json:"_description,omitempty"`
	ExposureState                 *Code             `json:"exposureState,omitempty"`
	ExposureStatePrimitiveElement *primitiveElement `json:"_exposureState,omitempty"`
	VariantState                  *CodeableConcept  `json:"variantState,omitempty"`
	RiskEvidenceSynthesis         Reference         `json:"riskEvidenceSynthesis,omitempty"`
}

func (r EffectEvidenceSynthesisResultsByExposure) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EffectEvidenceSynthesisResultsByExposure) marshalJSON() jsonEffectEvidenceSynthesisResultsByExposure {
	m := jsonEffectEvidenceSynthesisResultsByExposure{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	if r.ExposureState != nil && r.ExposureState.Value != nil {
		m.ExposureState = r.ExposureState
	}
	if r.ExposureState != nil && (r.ExposureState.Id != nil || r.ExposureState.Extension != nil) {
		m.ExposureStatePrimitiveElement = &primitiveElement{Id: r.ExposureState.Id, Extension: r.ExposureState.Extension}
	}
	m.VariantState = r.VariantState
	m.RiskEvidenceSynthesis = r.RiskEvidenceSynthesis
	return m
}
func (r *EffectEvidenceSynthesisResultsByExposure) UnmarshalJSON(b []byte) error {
	var m jsonEffectEvidenceSynthesisResultsByExposure
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EffectEvidenceSynthesisResultsByExposure) unmarshalJSON(m jsonEffectEvidenceSynthesisResultsByExposure) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.ExposureState = m.ExposureState
	if m.ExposureStatePrimitiveElement != nil {
		if r.ExposureState == nil {
			r.ExposureState = &Code{}
		}
		r.ExposureState.Id = m.ExposureStatePrimitiveElement.Id
		r.ExposureState.Extension = m.ExposureStatePrimitiveElement.Extension
	}
	r.VariantState = m.VariantState
	r.RiskEvidenceSynthesis = m.RiskEvidenceSynthesis
	return nil
}
func (r EffectEvidenceSynthesisResultsByExposure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ExposureState, xml.StartElement{Name: xml.Name{Local: "exposureState"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.VariantState, xml.StartElement{Name: xml.Name{Local: "variantState"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RiskEvidenceSynthesis, xml.StartElement{Name: xml.Name{Local: "riskEvidenceSynthesis"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *EffectEvidenceSynthesisResultsByExposure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "exposureState":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ExposureState = &v
			case "variantState":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.VariantState = &v
			case "riskEvidenceSynthesis":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RiskEvidenceSynthesis = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r EffectEvidenceSynthesisResultsByExposure) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The estimated effect of the exposure variant.
type EffectEvidenceSynthesisEffectEstimate struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Human-readable summary of effect estimate.
	Description *String
	// Examples include relative risk and mean difference.
	Type *CodeableConcept
	// Used to define variant exposure states such as low-risk state.
	VariantState *CodeableConcept
	// The point estimate of the effect estimate.
	Value *Decimal
	// Specifies the UCUM unit for the outcome.
	UnitOfMeasure *CodeableConcept
	// A description of the precision of the estimate for the effect.
	PrecisionEstimate []EffectEvidenceSynthesisEffectEstimatePrecisionEstimate
}
type jsonEffectEvidenceSynthesisEffectEstimate struct {
	Id                          *string                                                  `json:"id,omitempty"`
	Extension                   []Extension                                              `json:"extension,omitempty"`
	ModifierExtension           []Extension                                              `json:"modifierExtension,omitempty"`
	Description                 *String                                                  `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement                                        `json:"_description,omitempty"`
	Type                        *CodeableConcept                                         `json:"type,omitempty"`
	VariantState                *CodeableConcept                                         `json:"variantState,omitempty"`
	Value                       *Decimal                                                 `json:"value,omitempty"`
	ValuePrimitiveElement       *primitiveElement                                        `json:"_value,omitempty"`
	UnitOfMeasure               *CodeableConcept                                         `json:"unitOfMeasure,omitempty"`
	PrecisionEstimate           []EffectEvidenceSynthesisEffectEstimatePrecisionEstimate `json:"precisionEstimate,omitempty"`
}

func (r EffectEvidenceSynthesisEffectEstimate) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EffectEvidenceSynthesisEffectEstimate) marshalJSON() jsonEffectEvidenceSynthesisEffectEstimate {
	m := jsonEffectEvidenceSynthesisEffectEstimate{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Type = r.Type
	m.VariantState = r.VariantState
	if r.Value != nil && r.Value.Value != nil {
		m.Value = r.Value
	}
	if r.Value != nil && (r.Value.Id != nil || r.Value.Extension != nil) {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	m.UnitOfMeasure = r.UnitOfMeasure
	m.PrecisionEstimate = r.PrecisionEstimate
	return m
}
func (r *EffectEvidenceSynthesisEffectEstimate) UnmarshalJSON(b []byte) error {
	var m jsonEffectEvidenceSynthesisEffectEstimate
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EffectEvidenceSynthesisEffectEstimate) unmarshalJSON(m jsonEffectEvidenceSynthesisEffectEstimate) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &String{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.VariantState = m.VariantState
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		if r.Value == nil {
			r.Value = &Decimal{}
		}
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	r.UnitOfMeasure = m.UnitOfMeasure
	r.PrecisionEstimate = m.PrecisionEstimate
	return nil
}
func (r EffectEvidenceSynthesisEffectEstimate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.VariantState, xml.StartElement{Name: xml.Name{Local: "variantState"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Value, xml.StartElement{Name: xml.Name{Local: "value"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UnitOfMeasure, xml.StartElement{Name: xml.Name{Local: "unitOfMeasure"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PrecisionEstimate, xml.StartElement{Name: xml.Name{Local: "precisionEstimate"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *EffectEvidenceSynthesisEffectEstimate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = &v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "variantState":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.VariantState = &v
			case "value":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = &v
			case "unitOfMeasure":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UnitOfMeasure = &v
			case "precisionEstimate":
				var v EffectEvidenceSynthesisEffectEstimatePrecisionEstimate
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PrecisionEstimate = append(r.PrecisionEstimate, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r EffectEvidenceSynthesisEffectEstimate) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A description of the precision of the estimate for the effect.
type EffectEvidenceSynthesisEffectEstimatePrecisionEstimate struct {
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
type jsonEffectEvidenceSynthesisEffectEstimatePrecisionEstimate struct {
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

func (r EffectEvidenceSynthesisEffectEstimatePrecisionEstimate) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EffectEvidenceSynthesisEffectEstimatePrecisionEstimate) marshalJSON() jsonEffectEvidenceSynthesisEffectEstimatePrecisionEstimate {
	m := jsonEffectEvidenceSynthesisEffectEstimatePrecisionEstimate{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Level != nil && r.Level.Value != nil {
		m.Level = r.Level
	}
	if r.Level != nil && (r.Level.Id != nil || r.Level.Extension != nil) {
		m.LevelPrimitiveElement = &primitiveElement{Id: r.Level.Id, Extension: r.Level.Extension}
	}
	if r.From != nil && r.From.Value != nil {
		m.From = r.From
	}
	if r.From != nil && (r.From.Id != nil || r.From.Extension != nil) {
		m.FromPrimitiveElement = &primitiveElement{Id: r.From.Id, Extension: r.From.Extension}
	}
	if r.To != nil && r.To.Value != nil {
		m.To = r.To
	}
	if r.To != nil && (r.To.Id != nil || r.To.Extension != nil) {
		m.ToPrimitiveElement = &primitiveElement{Id: r.To.Id, Extension: r.To.Extension}
	}
	return m
}
func (r *EffectEvidenceSynthesisEffectEstimatePrecisionEstimate) UnmarshalJSON(b []byte) error {
	var m jsonEffectEvidenceSynthesisEffectEstimatePrecisionEstimate
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EffectEvidenceSynthesisEffectEstimatePrecisionEstimate) unmarshalJSON(m jsonEffectEvidenceSynthesisEffectEstimatePrecisionEstimate) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Level = m.Level
	if m.LevelPrimitiveElement != nil {
		if r.Level == nil {
			r.Level = &Decimal{}
		}
		r.Level.Id = m.LevelPrimitiveElement.Id
		r.Level.Extension = m.LevelPrimitiveElement.Extension
	}
	r.From = m.From
	if m.FromPrimitiveElement != nil {
		if r.From == nil {
			r.From = &Decimal{}
		}
		r.From.Id = m.FromPrimitiveElement.Id
		r.From.Extension = m.FromPrimitiveElement.Extension
	}
	r.To = m.To
	if m.ToPrimitiveElement != nil {
		if r.To == nil {
			r.To = &Decimal{}
		}
		r.To.Id = m.ToPrimitiveElement.Id
		r.To.Extension = m.ToPrimitiveElement.Extension
	}
	return nil
}
func (r EffectEvidenceSynthesisEffectEstimatePrecisionEstimate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Level, xml.StartElement{Name: xml.Name{Local: "level"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.From, xml.StartElement{Name: xml.Name{Local: "from"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.To, xml.StartElement{Name: xml.Name{Local: "to"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *EffectEvidenceSynthesisEffectEstimatePrecisionEstimate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "level":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Level = &v
			case "from":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.From = &v
			case "to":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.To = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r EffectEvidenceSynthesisEffectEstimatePrecisionEstimate) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A description of the certainty of the effect estimate.
type EffectEvidenceSynthesisCertainty struct {
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
	CertaintySubcomponent []EffectEvidenceSynthesisCertaintyCertaintySubcomponent
}
type jsonEffectEvidenceSynthesisCertainty struct {
	Id                    *string                                                 `json:"id,omitempty"`
	Extension             []Extension                                             `json:"extension,omitempty"`
	ModifierExtension     []Extension                                             `json:"modifierExtension,omitempty"`
	Rating                []CodeableConcept                                       `json:"rating,omitempty"`
	Note                  []Annotation                                            `json:"note,omitempty"`
	CertaintySubcomponent []EffectEvidenceSynthesisCertaintyCertaintySubcomponent `json:"certaintySubcomponent,omitempty"`
}

func (r EffectEvidenceSynthesisCertainty) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EffectEvidenceSynthesisCertainty) marshalJSON() jsonEffectEvidenceSynthesisCertainty {
	m := jsonEffectEvidenceSynthesisCertainty{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Rating = r.Rating
	m.Note = r.Note
	m.CertaintySubcomponent = r.CertaintySubcomponent
	return m
}
func (r *EffectEvidenceSynthesisCertainty) UnmarshalJSON(b []byte) error {
	var m jsonEffectEvidenceSynthesisCertainty
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EffectEvidenceSynthesisCertainty) unmarshalJSON(m jsonEffectEvidenceSynthesisCertainty) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Rating = m.Rating
	r.Note = m.Note
	r.CertaintySubcomponent = m.CertaintySubcomponent
	return nil
}
func (r EffectEvidenceSynthesisCertainty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Rating, xml.StartElement{Name: xml.Name{Local: "rating"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CertaintySubcomponent, xml.StartElement{Name: xml.Name{Local: "certaintySubcomponent"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *EffectEvidenceSynthesisCertainty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "rating":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Rating = append(r.Rating, v)
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			case "certaintySubcomponent":
				var v EffectEvidenceSynthesisCertaintyCertaintySubcomponent
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CertaintySubcomponent = append(r.CertaintySubcomponent, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r EffectEvidenceSynthesisCertainty) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A description of a component of the overall certainty.
type EffectEvidenceSynthesisCertaintyCertaintySubcomponent struct {
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
type jsonEffectEvidenceSynthesisCertaintyCertaintySubcomponent struct {
	Id                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Rating            []CodeableConcept `json:"rating,omitempty"`
	Note              []Annotation      `json:"note,omitempty"`
}

func (r EffectEvidenceSynthesisCertaintyCertaintySubcomponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EffectEvidenceSynthesisCertaintyCertaintySubcomponent) marshalJSON() jsonEffectEvidenceSynthesisCertaintyCertaintySubcomponent {
	m := jsonEffectEvidenceSynthesisCertaintyCertaintySubcomponent{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Rating = r.Rating
	m.Note = r.Note
	return m
}
func (r *EffectEvidenceSynthesisCertaintyCertaintySubcomponent) UnmarshalJSON(b []byte) error {
	var m jsonEffectEvidenceSynthesisCertaintyCertaintySubcomponent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EffectEvidenceSynthesisCertaintyCertaintySubcomponent) unmarshalJSON(m jsonEffectEvidenceSynthesisCertaintyCertaintySubcomponent) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Rating = m.Rating
	r.Note = m.Note
	return nil
}
func (r EffectEvidenceSynthesisCertaintyCertaintySubcomponent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Rating, xml.StartElement{Name: xml.Name{Local: "rating"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *EffectEvidenceSynthesisCertaintyCertaintySubcomponent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "rating":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Rating = append(r.Rating, v)
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r EffectEvidenceSynthesisCertaintyCertaintySubcomponent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
