package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// The findings and interpretation of diagnostic  tests performed on patients, groups of patients, devices, and locations, and/or specimens derived from these. The report includes clinical context such as requesting and provider information, and some mix of atomic results, images, textual and coded interpretations, and formatted representation of diagnostic reports.
//
// To support reporting for any diagnostic report into a clinical data repository.
type DiagnosticReport struct {
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
	// Identifiers assigned to this report by the performer or other systems.
	Identifier []Identifier
	// Details concerning a service requested.
	BasedOn []Reference
	// The status of the diagnostic report.
	Status Code
	// A code that classifies the clinical discipline, department or diagnostic service that created the report (e.g. cardiology, biochemistry, hematology, MRI). This is used for searching, sorting and display purposes.
	Category []CodeableConcept
	// A code or name that describes this diagnostic report.
	Code CodeableConcept
	// The subject of the report. Usually, but not always, this is a patient. However, diagnostic services also perform analyses on specimens collected from a variety of other sources.
	Subject *Reference
	// The healthcare event  (e.g. a patient and healthcare provider interaction) which this DiagnosticReport is about.
	Encounter *Reference
	// The time or time-period the observed values are related to. When the subject of the report is a patient, this is usually either the time of the procedure or of specimen collection(s), but very often the source of the date/time is not known, only the date/time itself.
	Effective isDiagnosticReportEffective
	// The date and time that this version of the report was made available to providers, typically after the report was reviewed and verified.
	Issued *Instant
	// The diagnostic service that is responsible for issuing the report.
	Performer []Reference
	// The practitioner or organization that is responsible for the report's conclusions and interpretations.
	ResultsInterpreter []Reference
	// Details about the specimens on which this diagnostic report is based.
	Specimen []Reference
	// [Observations](observation.html)  that are part of this diagnostic report.
	Result []Reference
	// One or more links to full details of any imaging performed during the diagnostic investigation. Typically, this is imaging performed by DICOM enabled modalities, but this is not required. A fully enabled PACS viewer can use this information to provide views of the source images.
	ImagingStudy []Reference
	// A list of key images associated with this report. The images are generally created during the diagnostic process, and may be directly of the patient, or of treated specimens (i.e. slides of interest).
	Media []DiagnosticReportMedia
	// Concise and clinically contextualized summary conclusion (interpretation/impression) of the diagnostic report.
	Conclusion *String
	// One or more codes that represent the summary conclusion (interpretation/impression) of the diagnostic report.
	ConclusionCode []CodeableConcept
	// Rich text representation of the entire result as issued by the diagnostic service. Multiple formats are allowed but they SHALL be semantically equivalent.
	PresentedForm []Attachment
}

func (r DiagnosticReport) ResourceType() string {
	return "DiagnosticReport"
}
func (r DiagnosticReport) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isDiagnosticReportEffective interface {
	isDiagnosticReportEffective()
}

func (r DateTime) isDiagnosticReportEffective() {}
func (r Period) isDiagnosticReportEffective()   {}

type jsonDiagnosticReport struct {
	ResourceType                      string                  `json:"resourceType"`
	Id                                *Id                     `json:"id,omitempty"`
	IdPrimitiveElement                *primitiveElement       `json:"_id,omitempty"`
	Meta                              *Meta                   `json:"meta,omitempty"`
	ImplicitRules                     *Uri                    `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement     *primitiveElement       `json:"_implicitRules,omitempty"`
	Language                          *Code                   `json:"language,omitempty"`
	LanguagePrimitiveElement          *primitiveElement       `json:"_language,omitempty"`
	Text                              *Narrative              `json:"text,omitempty"`
	Contained                         []ContainedResource     `json:"contained,omitempty"`
	Extension                         []Extension             `json:"extension,omitempty"`
	ModifierExtension                 []Extension             `json:"modifierExtension,omitempty"`
	Identifier                        []Identifier            `json:"identifier,omitempty"`
	BasedOn                           []Reference             `json:"basedOn,omitempty"`
	Status                            Code                    `json:"status,omitempty"`
	StatusPrimitiveElement            *primitiveElement       `json:"_status,omitempty"`
	Category                          []CodeableConcept       `json:"category,omitempty"`
	Code                              CodeableConcept         `json:"code,omitempty"`
	Subject                           *Reference              `json:"subject,omitempty"`
	Encounter                         *Reference              `json:"encounter,omitempty"`
	EffectiveDateTime                 *DateTime               `json:"effectiveDateTime,omitempty"`
	EffectiveDateTimePrimitiveElement *primitiveElement       `json:"_effectiveDateTime,omitempty"`
	EffectivePeriod                   *Period                 `json:"effectivePeriod,omitempty"`
	Issued                            *Instant                `json:"issued,omitempty"`
	IssuedPrimitiveElement            *primitiveElement       `json:"_issued,omitempty"`
	Performer                         []Reference             `json:"performer,omitempty"`
	ResultsInterpreter                []Reference             `json:"resultsInterpreter,omitempty"`
	Specimen                          []Reference             `json:"specimen,omitempty"`
	Result                            []Reference             `json:"result,omitempty"`
	ImagingStudy                      []Reference             `json:"imagingStudy,omitempty"`
	Media                             []DiagnosticReportMedia `json:"media,omitempty"`
	Conclusion                        *String                 `json:"conclusion,omitempty"`
	ConclusionPrimitiveElement        *primitiveElement       `json:"_conclusion,omitempty"`
	ConclusionCode                    []CodeableConcept       `json:"conclusionCode,omitempty"`
	PresentedForm                     []Attachment            `json:"presentedForm,omitempty"`
}

func (r DiagnosticReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DiagnosticReport) marshalJSON() jsonDiagnosticReport {
	m := jsonDiagnosticReport{}
	m.ResourceType = "DiagnosticReport"
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
	m.Identifier = r.Identifier
	m.BasedOn = r.BasedOn
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Category = r.Category
	m.Code = r.Code
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	switch v := r.Effective.(type) {
	case DateTime:
		if v.Value != nil {
			m.EffectiveDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.EffectiveDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.EffectiveDateTime = v
		}
		if v.Id != nil || v.Extension != nil {
			m.EffectiveDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.EffectivePeriod = &v
	case *Period:
		m.EffectivePeriod = v
	}
	if r.Issued != nil && r.Issued.Value != nil {
		m.Issued = r.Issued
	}
	if r.Issued != nil && (r.Issued.Id != nil || r.Issued.Extension != nil) {
		m.IssuedPrimitiveElement = &primitiveElement{Id: r.Issued.Id, Extension: r.Issued.Extension}
	}
	m.Performer = r.Performer
	m.ResultsInterpreter = r.ResultsInterpreter
	m.Specimen = r.Specimen
	m.Result = r.Result
	m.ImagingStudy = r.ImagingStudy
	m.Media = r.Media
	if r.Conclusion != nil && r.Conclusion.Value != nil {
		m.Conclusion = r.Conclusion
	}
	if r.Conclusion != nil && (r.Conclusion.Id != nil || r.Conclusion.Extension != nil) {
		m.ConclusionPrimitiveElement = &primitiveElement{Id: r.Conclusion.Id, Extension: r.Conclusion.Extension}
	}
	m.ConclusionCode = r.ConclusionCode
	m.PresentedForm = r.PresentedForm
	return m
}
func (r *DiagnosticReport) UnmarshalJSON(b []byte) error {
	var m jsonDiagnosticReport
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DiagnosticReport) unmarshalJSON(m jsonDiagnosticReport) error {
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
	r.Identifier = m.Identifier
	r.BasedOn = m.BasedOn
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Category = m.Category
	r.Code = m.Code
	r.Subject = m.Subject
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
	r.Issued = m.Issued
	if m.IssuedPrimitiveElement != nil {
		if r.Issued == nil {
			r.Issued = &Instant{}
		}
		r.Issued.Id = m.IssuedPrimitiveElement.Id
		r.Issued.Extension = m.IssuedPrimitiveElement.Extension
	}
	r.Performer = m.Performer
	r.ResultsInterpreter = m.ResultsInterpreter
	r.Specimen = m.Specimen
	r.Result = m.Result
	r.ImagingStudy = m.ImagingStudy
	r.Media = m.Media
	r.Conclusion = m.Conclusion
	if m.ConclusionPrimitiveElement != nil {
		if r.Conclusion == nil {
			r.Conclusion = &String{}
		}
		r.Conclusion.Id = m.ConclusionPrimitiveElement.Id
		r.Conclusion.Extension = m.ConclusionPrimitiveElement.Extension
	}
	r.ConclusionCode = m.ConclusionCode
	r.PresentedForm = m.PresentedForm
	return nil
}
func (r DiagnosticReport) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A list of key images associated with this report. The images are generally created during the diagnostic process, and may be directly of the patient, or of treated specimens (i.e. slides of interest).
type DiagnosticReportMedia struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A comment about the image. Typically, this is used to provide an explanation for why the image is included, or to draw the viewer's attention to important features.
	Comment *String
	// Reference to the image source.
	Link Reference
}
type jsonDiagnosticReportMedia struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Comment                 *String           `json:"comment,omitempty"`
	CommentPrimitiveElement *primitiveElement `json:"_comment,omitempty"`
	Link                    Reference         `json:"link,omitempty"`
}

func (r DiagnosticReportMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DiagnosticReportMedia) marshalJSON() jsonDiagnosticReportMedia {
	m := jsonDiagnosticReportMedia{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Comment != nil && r.Comment.Value != nil {
		m.Comment = r.Comment
	}
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	m.Link = r.Link
	return m
}
func (r *DiagnosticReportMedia) UnmarshalJSON(b []byte) error {
	var m jsonDiagnosticReportMedia
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DiagnosticReportMedia) unmarshalJSON(m jsonDiagnosticReportMedia) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		if r.Comment == nil {
			r.Comment = &String{}
		}
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	r.Link = m.Link
	return nil
}
func (r DiagnosticReportMedia) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
