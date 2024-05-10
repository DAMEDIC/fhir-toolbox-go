package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A process where a researcher or organization plans and then executes a series of steps intended to increase the field of healthcare-related knowledge.  This includes studies of safety, efficacy, comparative effectiveness and other information about medications, devices, therapies and other interventional and investigative techniques.  A ResearchStudy involves the gathering of information about human or animal subjects.
type ResearchStudy struct {
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
	// Identifiers assigned to this research study by the sponsor or other systems.
	Identifier []Identifier
	// A short, descriptive user-friendly label for the study.
	Title *String
	// The set of steps expected to be performed as part of the execution of the study.
	Protocol []Reference
	// A larger research study of which this particular study is a component or step.
	PartOf []Reference
	// The current state of the study.
	Status Code
	// The type of study based upon the intent of the study's activities. A classification of the intent of the study.
	PrimaryPurposeType *CodeableConcept
	// The stage in the progression of a therapy from initial experimental use in humans in clinical trials to post-market evaluation.
	Phase *CodeableConcept
	// Codes categorizing the type of study such as investigational vs. observational, type of blinding, type of randomization, safety vs. efficacy, etc.
	Category []CodeableConcept
	// The medication(s), food(s), therapy(ies), device(s) or other concerns or interventions that the study is seeking to gain more information about.
	Focus []CodeableConcept
	// The condition that is the focus of the study.  For example, In a study to examine risk factors for Lupus, might have as an inclusion criterion "healthy volunteer", but the target condition code would be a Lupus SNOMED code.
	Condition []CodeableConcept
	// Contact details to assist a user in learning more about or engaging with the study.
	Contact []ContactDetail
	// Citations, references and other related documents.
	RelatedArtifact []RelatedArtifact
	// Key terms to aid in searching for or filtering the study.
	Keyword []CodeableConcept
	// Indicates a country, state or other region where the study is taking place.
	Location []CodeableConcept
	// A full description of how the study is being conducted.
	Description *Markdown
	// Reference to a Group that defines the criteria for and quantity of subjects participating in the study.  E.g. " 200 female Europeans between the ages of 20 and 45 with early onset diabetes".
	Enrollment []Reference
	// Identifies the start date and the expected (or actual, depending on status) end date for the study.
	Period *Period
	// An organization that initiates the investigation and is legally responsible for the study.
	Sponsor *Reference
	// A researcher in a study who oversees multiple aspects of the study, such as concept development, protocol writing, protocol submission for IRB approval, participant recruitment, informed consent, data collection, analysis, interpretation and presentation.
	PrincipalInvestigator *Reference
	// A facility in which study activities are conducted.
	Site []Reference
	// A description and/or code explaining the premature termination of the study.
	ReasonStopped *CodeableConcept
	// Comments made about the study by the performer, subject or other participants.
	Note []Annotation
	// Describes an expected sequence of events for one of the participants of a study.  E.g. Exposure to drug A, wash-out, exposure to drug B, wash-out, follow-up.
	Arm []ResearchStudyArm
	// A goal that the study is aiming to achieve in terms of a scientific question to be answered by the analysis of data collected during the study.
	Objective []ResearchStudyObjective
}
type jsonResearchStudy struct {
	ResourceType                  string                   `json:"resourceType"`
	Id                            *Id                      `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement        `json:"_id,omitempty"`
	Meta                          *Meta                    `json:"meta,omitempty"`
	ImplicitRules                 *Uri                     `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement        `json:"_implicitRules,omitempty"`
	Language                      *Code                    `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement        `json:"_language,omitempty"`
	Text                          *Narrative               `json:"text,omitempty"`
	Contained                     []containedResource      `json:"contained,omitempty"`
	Extension                     []Extension              `json:"extension,omitempty"`
	ModifierExtension             []Extension              `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier             `json:"identifier,omitempty"`
	Title                         *String                  `json:"title,omitempty"`
	TitlePrimitiveElement         *primitiveElement        `json:"_title,omitempty"`
	Protocol                      []Reference              `json:"protocol,omitempty"`
	PartOf                        []Reference              `json:"partOf,omitempty"`
	Status                        Code                     `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement        `json:"_status,omitempty"`
	PrimaryPurposeType            *CodeableConcept         `json:"primaryPurposeType,omitempty"`
	Phase                         *CodeableConcept         `json:"phase,omitempty"`
	Category                      []CodeableConcept        `json:"category,omitempty"`
	Focus                         []CodeableConcept        `json:"focus,omitempty"`
	Condition                     []CodeableConcept        `json:"condition,omitempty"`
	Contact                       []ContactDetail          `json:"contact,omitempty"`
	RelatedArtifact               []RelatedArtifact        `json:"relatedArtifact,omitempty"`
	Keyword                       []CodeableConcept        `json:"keyword,omitempty"`
	Location                      []CodeableConcept        `json:"location,omitempty"`
	Description                   *Markdown                `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement        `json:"_description,omitempty"`
	Enrollment                    []Reference              `json:"enrollment,omitempty"`
	Period                        *Period                  `json:"period,omitempty"`
	Sponsor                       *Reference               `json:"sponsor,omitempty"`
	PrincipalInvestigator         *Reference               `json:"principalInvestigator,omitempty"`
	Site                          []Reference              `json:"site,omitempty"`
	ReasonStopped                 *CodeableConcept         `json:"reasonStopped,omitempty"`
	Note                          []Annotation             `json:"note,omitempty"`
	Arm                           []ResearchStudyArm       `json:"arm,omitempty"`
	Objective                     []ResearchStudyObjective `json:"objective,omitempty"`
}

func (r ResearchStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ResearchStudy) marshalJSON() jsonResearchStudy {
	m := jsonResearchStudy{}
	m.ResourceType = "ResearchStudy"
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
	m.Title = r.Title
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	m.Protocol = r.Protocol
	m.PartOf = r.PartOf
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.PrimaryPurposeType = r.PrimaryPurposeType
	m.Phase = r.Phase
	m.Category = r.Category
	m.Focus = r.Focus
	m.Condition = r.Condition
	m.Contact = r.Contact
	m.RelatedArtifact = r.RelatedArtifact
	m.Keyword = r.Keyword
	m.Location = r.Location
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Enrollment = r.Enrollment
	m.Period = r.Period
	m.Sponsor = r.Sponsor
	m.PrincipalInvestigator = r.PrincipalInvestigator
	m.Site = r.Site
	m.ReasonStopped = r.ReasonStopped
	m.Note = r.Note
	m.Arm = r.Arm
	m.Objective = r.Objective
	return m
}
func (r *ResearchStudy) UnmarshalJSON(b []byte) error {
	var m jsonResearchStudy
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ResearchStudy) unmarshalJSON(m jsonResearchStudy) error {
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
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Protocol = m.Protocol
	r.PartOf = m.PartOf
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.PrimaryPurposeType = m.PrimaryPurposeType
	r.Phase = m.Phase
	r.Category = m.Category
	r.Focus = m.Focus
	r.Condition = m.Condition
	r.Contact = m.Contact
	r.RelatedArtifact = m.RelatedArtifact
	r.Keyword = m.Keyword
	r.Location = m.Location
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Enrollment = m.Enrollment
	r.Period = m.Period
	r.Sponsor = m.Sponsor
	r.PrincipalInvestigator = m.PrincipalInvestigator
	r.Site = m.Site
	r.ReasonStopped = m.ReasonStopped
	r.Note = m.Note
	r.Arm = m.Arm
	r.Objective = m.Objective
	return nil
}
func (r ResearchStudy) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Describes an expected sequence of events for one of the participants of a study.  E.g. Exposure to drug A, wash-out, exposure to drug B, wash-out, follow-up.
type ResearchStudyArm struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Unique, human-readable label for this arm of the study.
	Name String
	// Categorization of study arm, e.g. experimental, active comparator, placebo comparater.
	Type *CodeableConcept
	// A succinct description of the path through the study that would be followed by a subject adhering to this arm.
	Description *String
}
type jsonResearchStudyArm struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Name                        String            `json:"name,omitempty"`
	NamePrimitiveElement        *primitiveElement `json:"_name,omitempty"`
	Type                        *CodeableConcept  `json:"type,omitempty"`
	Description                 *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
}

func (r ResearchStudyArm) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ResearchStudyArm) marshalJSON() jsonResearchStudyArm {
	m := jsonResearchStudyArm{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Name = r.Name
	if r.Name.Id != nil || r.Name.Extension != nil {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Type = r.Type
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	return m
}
func (r *ResearchStudyArm) UnmarshalJSON(b []byte) error {
	var m jsonResearchStudyArm
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ResearchStudyArm) unmarshalJSON(m jsonResearchStudyArm) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	return nil
}
func (r ResearchStudyArm) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A goal that the study is aiming to achieve in terms of a scientific question to be answered by the analysis of data collected during the study.
type ResearchStudyObjective struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Unique, human-readable label for this objective of the study.
	Name *String
	// The kind of study objective.
	Type *CodeableConcept
}
type jsonResearchStudyObjective struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Name                 *String           `json:"name,omitempty"`
	NamePrimitiveElement *primitiveElement `json:"_name,omitempty"`
	Type                 *CodeableConcept  `json:"type,omitempty"`
}

func (r ResearchStudyObjective) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ResearchStudyObjective) marshalJSON() jsonResearchStudyObjective {
	m := jsonResearchStudyObjective{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Type = r.Type
	return m
}
func (r *ResearchStudyObjective) UnmarshalJSON(b []byte) error {
	var m jsonResearchStudyObjective
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ResearchStudyObjective) unmarshalJSON(m jsonResearchStudyObjective) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Type = m.Type
	return nil
}
func (r ResearchStudyObjective) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
