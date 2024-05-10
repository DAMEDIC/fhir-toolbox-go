package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// Representation of the content produced in a DICOM imaging study. A study comprises a set of series, each of which includes a set of Service-Object Pair Instances (SOP Instances - images or other data) acquired or produced in a common context.  A series is of only one modality (e.g. X-ray, CT, MR, ultrasound), but a study may have multiple series of different modalities.
type ImagingStudy struct {
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
	// Identifiers for the ImagingStudy such as DICOM Study Instance UID, and Accession Number.
	Identifier []Identifier
	// The current state of the ImagingStudy.
	Status Code
	// A list of all the series.modality values that are actual acquisition modalities, i.e. those in the DICOM Context Group 29 (value set OID 1.2.840.10008.6.1.19).
	Modality []Coding
	// The subject, typically a patient, of the imaging study.
	Subject Reference
	// The healthcare event (e.g. a patient and healthcare provider interaction) during which this ImagingStudy is made.
	Encounter *Reference
	// Date and time the study started.
	Started *DateTime
	// A list of the diagnostic requests that resulted in this imaging study being performed.
	BasedOn []Reference
	// The requesting/referring physician.
	Referrer *Reference
	// Who read the study and interpreted the images or other content.
	Interpreter []Reference
	// The network service providing access (e.g., query, view, or retrieval) for the study. See implementation notes for information about using DICOM endpoints. A study-level endpoint applies to each series in the study, unless overridden by a series-level endpoint with the same Endpoint.connectionType.
	Endpoint []Reference
	// Number of Series in the Study. This value given may be larger than the number of series elements this Resource contains due to resource availability, security, or other factors. This element should be present if any series elements are present.
	NumberOfSeries *UnsignedInt
	// Number of SOP Instances in Study. This value given may be larger than the number of instance elements this resource contains due to resource availability, security, or other factors. This element should be present if any instance elements are present.
	NumberOfInstances *UnsignedInt
	// The procedure which this ImagingStudy was part of.
	ProcedureReference *Reference
	// The code for the performed procedure type.
	ProcedureCode []CodeableConcept
	// The principal physical location where the ImagingStudy was performed.
	Location *Reference
	// Description of clinical condition indicating why the ImagingStudy was requested.
	ReasonCode []CodeableConcept
	// Indicates another resource whose existence justifies this Study.
	ReasonReference []Reference
	// Per the recommended DICOM mapping, this element is derived from the Study Description attribute (0008,1030). Observations or findings about the imaging study should be recorded in another resource, e.g. Observation, and not in this element.
	Note []Annotation
	// The Imaging Manager description of the study. Institution-generated description or classification of the Study (component) performed.
	Description *String
	// Each study has one or more series of images or other content.
	Series []ImagingStudySeries
}
type jsonImagingStudy struct {
	ResourceType                      string               `json:"resourceType"`
	Id                                *Id                  `json:"id,omitempty"`
	IdPrimitiveElement                *primitiveElement    `json:"_id,omitempty"`
	Meta                              *Meta                `json:"meta,omitempty"`
	ImplicitRules                     *Uri                 `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement     *primitiveElement    `json:"_implicitRules,omitempty"`
	Language                          *Code                `json:"language,omitempty"`
	LanguagePrimitiveElement          *primitiveElement    `json:"_language,omitempty"`
	Text                              *Narrative           `json:"text,omitempty"`
	Contained                         []containedResource  `json:"contained,omitempty"`
	Extension                         []Extension          `json:"extension,omitempty"`
	ModifierExtension                 []Extension          `json:"modifierExtension,omitempty"`
	Identifier                        []Identifier         `json:"identifier,omitempty"`
	Status                            Code                 `json:"status,omitempty"`
	StatusPrimitiveElement            *primitiveElement    `json:"_status,omitempty"`
	Modality                          []Coding             `json:"modality,omitempty"`
	Subject                           Reference            `json:"subject,omitempty"`
	Encounter                         *Reference           `json:"encounter,omitempty"`
	Started                           *DateTime            `json:"started,omitempty"`
	StartedPrimitiveElement           *primitiveElement    `json:"_started,omitempty"`
	BasedOn                           []Reference          `json:"basedOn,omitempty"`
	Referrer                          *Reference           `json:"referrer,omitempty"`
	Interpreter                       []Reference          `json:"interpreter,omitempty"`
	Endpoint                          []Reference          `json:"endpoint,omitempty"`
	NumberOfSeries                    *UnsignedInt         `json:"numberOfSeries,omitempty"`
	NumberOfSeriesPrimitiveElement    *primitiveElement    `json:"_numberOfSeries,omitempty"`
	NumberOfInstances                 *UnsignedInt         `json:"numberOfInstances,omitempty"`
	NumberOfInstancesPrimitiveElement *primitiveElement    `json:"_numberOfInstances,omitempty"`
	ProcedureReference                *Reference           `json:"procedureReference,omitempty"`
	ProcedureCode                     []CodeableConcept    `json:"procedureCode,omitempty"`
	Location                          *Reference           `json:"location,omitempty"`
	ReasonCode                        []CodeableConcept    `json:"reasonCode,omitempty"`
	ReasonReference                   []Reference          `json:"reasonReference,omitempty"`
	Note                              []Annotation         `json:"note,omitempty"`
	Description                       *String              `json:"description,omitempty"`
	DescriptionPrimitiveElement       *primitiveElement    `json:"_description,omitempty"`
	Series                            []ImagingStudySeries `json:"series,omitempty"`
}

func (r ImagingStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImagingStudy) marshalJSON() jsonImagingStudy {
	m := jsonImagingStudy{}
	m.ResourceType = "ImagingStudy"
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
	m.Modality = r.Modality
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	m.Started = r.Started
	if r.Started != nil && (r.Started.Id != nil || r.Started.Extension != nil) {
		m.StartedPrimitiveElement = &primitiveElement{Id: r.Started.Id, Extension: r.Started.Extension}
	}
	m.BasedOn = r.BasedOn
	m.Referrer = r.Referrer
	m.Interpreter = r.Interpreter
	m.Endpoint = r.Endpoint
	m.NumberOfSeries = r.NumberOfSeries
	if r.NumberOfSeries != nil && (r.NumberOfSeries.Id != nil || r.NumberOfSeries.Extension != nil) {
		m.NumberOfSeriesPrimitiveElement = &primitiveElement{Id: r.NumberOfSeries.Id, Extension: r.NumberOfSeries.Extension}
	}
	m.NumberOfInstances = r.NumberOfInstances
	if r.NumberOfInstances != nil && (r.NumberOfInstances.Id != nil || r.NumberOfInstances.Extension != nil) {
		m.NumberOfInstancesPrimitiveElement = &primitiveElement{Id: r.NumberOfInstances.Id, Extension: r.NumberOfInstances.Extension}
	}
	m.ProcedureReference = r.ProcedureReference
	m.ProcedureCode = r.ProcedureCode
	m.Location = r.Location
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Note = r.Note
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Series = r.Series
	return m
}
func (r *ImagingStudy) UnmarshalJSON(b []byte) error {
	var m jsonImagingStudy
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImagingStudy) unmarshalJSON(m jsonImagingStudy) error {
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
	r.Modality = m.Modality
	r.Subject = m.Subject
	r.Encounter = m.Encounter
	r.Started = m.Started
	if m.StartedPrimitiveElement != nil {
		r.Started.Id = m.StartedPrimitiveElement.Id
		r.Started.Extension = m.StartedPrimitiveElement.Extension
	}
	r.BasedOn = m.BasedOn
	r.Referrer = m.Referrer
	r.Interpreter = m.Interpreter
	r.Endpoint = m.Endpoint
	r.NumberOfSeries = m.NumberOfSeries
	if m.NumberOfSeriesPrimitiveElement != nil {
		r.NumberOfSeries.Id = m.NumberOfSeriesPrimitiveElement.Id
		r.NumberOfSeries.Extension = m.NumberOfSeriesPrimitiveElement.Extension
	}
	r.NumberOfInstances = m.NumberOfInstances
	if m.NumberOfInstancesPrimitiveElement != nil {
		r.NumberOfInstances.Id = m.NumberOfInstancesPrimitiveElement.Id
		r.NumberOfInstances.Extension = m.NumberOfInstancesPrimitiveElement.Extension
	}
	r.ProcedureReference = m.ProcedureReference
	r.ProcedureCode = m.ProcedureCode
	r.Location = m.Location
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Note = m.Note
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Series = m.Series
	return nil
}
func (r ImagingStudy) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Each study has one or more series of images or other content.
type ImagingStudySeries struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The DICOM Series Instance UID for the series.
	Uid Id
	// The numeric identifier of this series in the study.
	Number *UnsignedInt
	// The modality of this series sequence.
	Modality Coding
	// A description of the series.
	Description *String
	// Number of SOP Instances in the Study. The value given may be larger than the number of instance elements this resource contains due to resource availability, security, or other factors. This element should be present if any instance elements are present.
	NumberOfInstances *UnsignedInt
	// The network service providing access (e.g., query, view, or retrieval) for this series. See implementation notes for information about using DICOM endpoints. A series-level endpoint, if present, has precedence over a study-level endpoint with the same Endpoint.connectionType.
	Endpoint []Reference
	// The anatomic structures examined. See DICOM Part 16 Annex L (http://dicom.nema.org/medical/dicom/current/output/chtml/part16/chapter_L.html) for DICOM to SNOMED-CT mappings. The bodySite may indicate the laterality of body part imaged; if so, it shall be consistent with any content of ImagingStudy.series.laterality.
	BodySite *Coding
	// The laterality of the (possibly paired) anatomic structures examined. E.g., the left knee, both lungs, or unpaired abdomen. If present, shall be consistent with any laterality information indicated in ImagingStudy.series.bodySite.
	Laterality *Coding
	// The specimen imaged, e.g., for whole slide imaging of a biopsy.
	Specimen []Reference
	// The date and time the series was started.
	Started *DateTime
	// Indicates who or what performed the series and how they were involved.
	Performer []ImagingStudySeriesPerformer
	// A single SOP instance within the series, e.g. an image, or presentation state.
	Instance []ImagingStudySeriesInstance
}
type jsonImagingStudySeries struct {
	Id                                *string                       `json:"id,omitempty"`
	Extension                         []Extension                   `json:"extension,omitempty"`
	ModifierExtension                 []Extension                   `json:"modifierExtension,omitempty"`
	Uid                               Id                            `json:"uid,omitempty"`
	UidPrimitiveElement               *primitiveElement             `json:"_uid,omitempty"`
	Number                            *UnsignedInt                  `json:"number,omitempty"`
	NumberPrimitiveElement            *primitiveElement             `json:"_number,omitempty"`
	Modality                          Coding                        `json:"modality,omitempty"`
	Description                       *String                       `json:"description,omitempty"`
	DescriptionPrimitiveElement       *primitiveElement             `json:"_description,omitempty"`
	NumberOfInstances                 *UnsignedInt                  `json:"numberOfInstances,omitempty"`
	NumberOfInstancesPrimitiveElement *primitiveElement             `json:"_numberOfInstances,omitempty"`
	Endpoint                          []Reference                   `json:"endpoint,omitempty"`
	BodySite                          *Coding                       `json:"bodySite,omitempty"`
	Laterality                        *Coding                       `json:"laterality,omitempty"`
	Specimen                          []Reference                   `json:"specimen,omitempty"`
	Started                           *DateTime                     `json:"started,omitempty"`
	StartedPrimitiveElement           *primitiveElement             `json:"_started,omitempty"`
	Performer                         []ImagingStudySeriesPerformer `json:"performer,omitempty"`
	Instance                          []ImagingStudySeriesInstance  `json:"instance,omitempty"`
}

func (r ImagingStudySeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImagingStudySeries) marshalJSON() jsonImagingStudySeries {
	m := jsonImagingStudySeries{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Uid = r.Uid
	if r.Uid.Id != nil || r.Uid.Extension != nil {
		m.UidPrimitiveElement = &primitiveElement{Id: r.Uid.Id, Extension: r.Uid.Extension}
	}
	m.Number = r.Number
	if r.Number != nil && (r.Number.Id != nil || r.Number.Extension != nil) {
		m.NumberPrimitiveElement = &primitiveElement{Id: r.Number.Id, Extension: r.Number.Extension}
	}
	m.Modality = r.Modality
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.NumberOfInstances = r.NumberOfInstances
	if r.NumberOfInstances != nil && (r.NumberOfInstances.Id != nil || r.NumberOfInstances.Extension != nil) {
		m.NumberOfInstancesPrimitiveElement = &primitiveElement{Id: r.NumberOfInstances.Id, Extension: r.NumberOfInstances.Extension}
	}
	m.Endpoint = r.Endpoint
	m.BodySite = r.BodySite
	m.Laterality = r.Laterality
	m.Specimen = r.Specimen
	m.Started = r.Started
	if r.Started != nil && (r.Started.Id != nil || r.Started.Extension != nil) {
		m.StartedPrimitiveElement = &primitiveElement{Id: r.Started.Id, Extension: r.Started.Extension}
	}
	m.Performer = r.Performer
	m.Instance = r.Instance
	return m
}
func (r *ImagingStudySeries) UnmarshalJSON(b []byte) error {
	var m jsonImagingStudySeries
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImagingStudySeries) unmarshalJSON(m jsonImagingStudySeries) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Uid = m.Uid
	if m.UidPrimitiveElement != nil {
		r.Uid.Id = m.UidPrimitiveElement.Id
		r.Uid.Extension = m.UidPrimitiveElement.Extension
	}
	r.Number = m.Number
	if m.NumberPrimitiveElement != nil {
		r.Number.Id = m.NumberPrimitiveElement.Id
		r.Number.Extension = m.NumberPrimitiveElement.Extension
	}
	r.Modality = m.Modality
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.NumberOfInstances = m.NumberOfInstances
	if m.NumberOfInstancesPrimitiveElement != nil {
		r.NumberOfInstances.Id = m.NumberOfInstancesPrimitiveElement.Id
		r.NumberOfInstances.Extension = m.NumberOfInstancesPrimitiveElement.Extension
	}
	r.Endpoint = m.Endpoint
	r.BodySite = m.BodySite
	r.Laterality = m.Laterality
	r.Specimen = m.Specimen
	r.Started = m.Started
	if m.StartedPrimitiveElement != nil {
		r.Started.Id = m.StartedPrimitiveElement.Id
		r.Started.Extension = m.StartedPrimitiveElement.Extension
	}
	r.Performer = m.Performer
	r.Instance = m.Instance
	return nil
}
func (r ImagingStudySeries) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates who or what performed the series and how they were involved.
type ImagingStudySeriesPerformer struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Distinguishes the type of involvement of the performer in the series.
	Function *CodeableConcept
	// Indicates who or what performed the series.
	Actor Reference
}
type jsonImagingStudySeriesPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor,omitempty"`
}

func (r ImagingStudySeriesPerformer) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImagingStudySeriesPerformer) marshalJSON() jsonImagingStudySeriesPerformer {
	m := jsonImagingStudySeriesPerformer{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Function = r.Function
	m.Actor = r.Actor
	return m
}
func (r *ImagingStudySeriesPerformer) UnmarshalJSON(b []byte) error {
	var m jsonImagingStudySeriesPerformer
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImagingStudySeriesPerformer) unmarshalJSON(m jsonImagingStudySeriesPerformer) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Function = m.Function
	r.Actor = m.Actor
	return nil
}
func (r ImagingStudySeriesPerformer) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A single SOP instance within the series, e.g. an image, or presentation state.
type ImagingStudySeriesInstance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The DICOM SOP Instance UID for this image or other DICOM content.
	Uid Id
	// DICOM instance  type.
	SopClass Coding
	// The number of instance in the series.
	Number *UnsignedInt
	// The description of the instance.
	Title *String
}
type jsonImagingStudySeriesInstance struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Uid                    Id                `json:"uid,omitempty"`
	UidPrimitiveElement    *primitiveElement `json:"_uid,omitempty"`
	SopClass               Coding            `json:"sopClass,omitempty"`
	Number                 *UnsignedInt      `json:"number,omitempty"`
	NumberPrimitiveElement *primitiveElement `json:"_number,omitempty"`
	Title                  *String           `json:"title,omitempty"`
	TitlePrimitiveElement  *primitiveElement `json:"_title,omitempty"`
}

func (r ImagingStudySeriesInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ImagingStudySeriesInstance) marshalJSON() jsonImagingStudySeriesInstance {
	m := jsonImagingStudySeriesInstance{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Uid = r.Uid
	if r.Uid.Id != nil || r.Uid.Extension != nil {
		m.UidPrimitiveElement = &primitiveElement{Id: r.Uid.Id, Extension: r.Uid.Extension}
	}
	m.SopClass = r.SopClass
	m.Number = r.Number
	if r.Number != nil && (r.Number.Id != nil || r.Number.Extension != nil) {
		m.NumberPrimitiveElement = &primitiveElement{Id: r.Number.Id, Extension: r.Number.Extension}
	}
	m.Title = r.Title
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	return m
}
func (r *ImagingStudySeriesInstance) UnmarshalJSON(b []byte) error {
	var m jsonImagingStudySeriesInstance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ImagingStudySeriesInstance) unmarshalJSON(m jsonImagingStudySeriesInstance) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Uid = m.Uid
	if m.UidPrimitiveElement != nil {
		r.Uid.Id = m.UidPrimitiveElement.Id
		r.Uid.Extension = m.UidPrimitiveElement.Extension
	}
	r.SopClass = m.SopClass
	r.Number = m.Number
	if m.NumberPrimitiveElement != nil {
		r.Number.Id = m.NumberPrimitiveElement.Id
		r.Number.Extension = m.NumberPrimitiveElement.Extension
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	return nil
}
func (r ImagingStudySeriesInstance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
