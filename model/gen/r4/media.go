package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A photo, video, or audio recording acquired or used in healthcare. The actual content may be inline or provided by direct reference.
type Media struct {
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
	// Identifiers associated with the image - these may include identifiers for the image itself, identifiers for the context of its collection (e.g. series ids) and context ids such as accession numbers or other workflow identifiers.
	Identifier []Identifier
	// A procedure that is fulfilled in whole or in part by the creation of this media.
	BasedOn []Reference
	// A larger event of which this particular event is a component or step.
	PartOf []Reference
	// The current state of the {{title}}.
	Status Code
	// A code that classifies whether the media is an image, video or audio recording or some other media category.
	Type *CodeableConcept
	// Details of the type of the media - usually, how it was acquired (what type of device). If images sourced from a DICOM system, are wrapped in a Media resource, then this is the modality.
	Modality *CodeableConcept
	// The name of the imaging view e.g. Lateral or Antero-posterior (AP).
	View *CodeableConcept
	// Who/What this Media is a record of.
	Subject *Reference
	// The encounter that establishes the context for this media.
	Encounter *Reference
	// The date and time(s) at which the media was collected.
	Created isMediaCreated
	// The date and time this version of the media was made available to providers, typically after having been reviewed.
	Issued *Instant
	// The person who administered the collection of the image.
	Operator *Reference
	// Describes why the event occurred in coded or textual form.
	ReasonCode []CodeableConcept
	// Indicates the site on the subject's body where the observation was made (i.e. the target site).
	BodySite *CodeableConcept
	// The name of the device / manufacturer of the device  that was used to make the recording.
	DeviceName *String
	// The device used to collect the media.
	Device *Reference
	// Height of the image in pixels (photo/video).
	Height *PositiveInt
	// Width of the image in pixels (photo/video).
	Width *PositiveInt
	// The number of frames in a photo. This is used with a multi-page fax, or an imaging acquisition context that takes multiple slices in a single image, or an animated gif. If there is more than one frame, this SHALL have a value in order to alert interface software that a multi-frame capable rendering widget is required.
	Frames *PositiveInt
	// The duration of the recording in seconds - for audio and video.
	Duration *Decimal
	// The actual content of the media - inline or by direct reference to the media source file.
	Content Attachment
	// Comments made about the media by the performer, subject or other participants.
	Note []Annotation
}
type isMediaCreated interface {
	isMediaCreated()
}

func (r DateTime) isMediaCreated() {}
func (r Period) isMediaCreated()   {}

type jsonMedia struct {
	ResourceType                    string              `json:"resourceType"`
	Id                              *Id                 `json:"id,omitempty"`
	IdPrimitiveElement              *primitiveElement   `json:"_id,omitempty"`
	Meta                            *Meta               `json:"meta,omitempty"`
	ImplicitRules                   *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement   *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                        *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement        *primitiveElement   `json:"_language,omitempty"`
	Text                            *Narrative          `json:"text,omitempty"`
	Contained                       []containedResource `json:"contained,omitempty"`
	Extension                       []Extension         `json:"extension,omitempty"`
	ModifierExtension               []Extension         `json:"modifierExtension,omitempty"`
	Identifier                      []Identifier        `json:"identifier,omitempty"`
	BasedOn                         []Reference         `json:"basedOn,omitempty"`
	PartOf                          []Reference         `json:"partOf,omitempty"`
	Status                          Code                `json:"status,omitempty"`
	StatusPrimitiveElement          *primitiveElement   `json:"_status,omitempty"`
	Type                            *CodeableConcept    `json:"type,omitempty"`
	Modality                        *CodeableConcept    `json:"modality,omitempty"`
	View                            *CodeableConcept    `json:"view,omitempty"`
	Subject                         *Reference          `json:"subject,omitempty"`
	Encounter                       *Reference          `json:"encounter,omitempty"`
	CreatedDateTime                 *DateTime           `json:"createdDateTime,omitempty"`
	CreatedDateTimePrimitiveElement *primitiveElement   `json:"_createdDateTime,omitempty"`
	CreatedPeriod                   *Period             `json:"createdPeriod,omitempty"`
	Issued                          *Instant            `json:"issued,omitempty"`
	IssuedPrimitiveElement          *primitiveElement   `json:"_issued,omitempty"`
	Operator                        *Reference          `json:"operator,omitempty"`
	ReasonCode                      []CodeableConcept   `json:"reasonCode,omitempty"`
	BodySite                        *CodeableConcept    `json:"bodySite,omitempty"`
	DeviceName                      *String             `json:"deviceName,omitempty"`
	DeviceNamePrimitiveElement      *primitiveElement   `json:"_deviceName,omitempty"`
	Device                          *Reference          `json:"device,omitempty"`
	Height                          *PositiveInt        `json:"height,omitempty"`
	HeightPrimitiveElement          *primitiveElement   `json:"_height,omitempty"`
	Width                           *PositiveInt        `json:"width,omitempty"`
	WidthPrimitiveElement           *primitiveElement   `json:"_width,omitempty"`
	Frames                          *PositiveInt        `json:"frames,omitempty"`
	FramesPrimitiveElement          *primitiveElement   `json:"_frames,omitempty"`
	Duration                        *Decimal            `json:"duration,omitempty"`
	DurationPrimitiveElement        *primitiveElement   `json:"_duration,omitempty"`
	Content                         Attachment          `json:"content,omitempty"`
	Note                            []Annotation        `json:"note,omitempty"`
}

func (r Media) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Media) marshalJSON() jsonMedia {
	m := jsonMedia{}
	m.ResourceType = "Media"
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
	m.Type = r.Type
	m.Modality = r.Modality
	m.View = r.View
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	switch v := r.Created.(type) {
	case DateTime:
		m.CreatedDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.CreatedDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.CreatedDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.CreatedDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.CreatedPeriod = &v
	case *Period:
		m.CreatedPeriod = v
	}
	m.Issued = r.Issued
	if r.Issued != nil && (r.Issued.Id != nil || r.Issued.Extension != nil) {
		m.IssuedPrimitiveElement = &primitiveElement{Id: r.Issued.Id, Extension: r.Issued.Extension}
	}
	m.Operator = r.Operator
	m.ReasonCode = r.ReasonCode
	m.BodySite = r.BodySite
	m.DeviceName = r.DeviceName
	if r.DeviceName != nil && (r.DeviceName.Id != nil || r.DeviceName.Extension != nil) {
		m.DeviceNamePrimitiveElement = &primitiveElement{Id: r.DeviceName.Id, Extension: r.DeviceName.Extension}
	}
	m.Device = r.Device
	m.Height = r.Height
	if r.Height != nil && (r.Height.Id != nil || r.Height.Extension != nil) {
		m.HeightPrimitiveElement = &primitiveElement{Id: r.Height.Id, Extension: r.Height.Extension}
	}
	m.Width = r.Width
	if r.Width != nil && (r.Width.Id != nil || r.Width.Extension != nil) {
		m.WidthPrimitiveElement = &primitiveElement{Id: r.Width.Id, Extension: r.Width.Extension}
	}
	m.Frames = r.Frames
	if r.Frames != nil && (r.Frames.Id != nil || r.Frames.Extension != nil) {
		m.FramesPrimitiveElement = &primitiveElement{Id: r.Frames.Id, Extension: r.Frames.Extension}
	}
	m.Duration = r.Duration
	if r.Duration != nil && (r.Duration.Id != nil || r.Duration.Extension != nil) {
		m.DurationPrimitiveElement = &primitiveElement{Id: r.Duration.Id, Extension: r.Duration.Extension}
	}
	m.Content = r.Content
	m.Note = r.Note
	return m
}
func (r *Media) UnmarshalJSON(b []byte) error {
	var m jsonMedia
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Media) unmarshalJSON(m jsonMedia) error {
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
	r.PartOf = m.PartOf
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Modality = m.Modality
	r.View = m.View
	r.Subject = m.Subject
	r.Encounter = m.Encounter
	if m.CreatedDateTime != nil || m.CreatedDateTimePrimitiveElement != nil {
		if r.Created != nil {
			return fmt.Errorf("multiple values for field \"Created\"")
		}
		v := m.CreatedDateTime
		if m.CreatedDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.CreatedDateTimePrimitiveElement.Id
			v.Extension = m.CreatedDateTimePrimitiveElement.Extension
		}
		r.Created = v
	}
	if m.CreatedPeriod != nil {
		if r.Created != nil {
			return fmt.Errorf("multiple values for field \"Created\"")
		}
		v := m.CreatedPeriod
		r.Created = v
	}
	r.Issued = m.Issued
	if m.IssuedPrimitiveElement != nil {
		r.Issued.Id = m.IssuedPrimitiveElement.Id
		r.Issued.Extension = m.IssuedPrimitiveElement.Extension
	}
	r.Operator = m.Operator
	r.ReasonCode = m.ReasonCode
	r.BodySite = m.BodySite
	r.DeviceName = m.DeviceName
	if m.DeviceNamePrimitiveElement != nil {
		r.DeviceName.Id = m.DeviceNamePrimitiveElement.Id
		r.DeviceName.Extension = m.DeviceNamePrimitiveElement.Extension
	}
	r.Device = m.Device
	r.Height = m.Height
	if m.HeightPrimitiveElement != nil {
		r.Height.Id = m.HeightPrimitiveElement.Id
		r.Height.Extension = m.HeightPrimitiveElement.Extension
	}
	r.Width = m.Width
	if m.WidthPrimitiveElement != nil {
		r.Width.Id = m.WidthPrimitiveElement.Id
		r.Width.Extension = m.WidthPrimitiveElement.Extension
	}
	r.Frames = m.Frames
	if m.FramesPrimitiveElement != nil {
		r.Frames.Id = m.FramesPrimitiveElement.Id
		r.Frames.Extension = m.FramesPrimitiveElement.Extension
	}
	r.Duration = m.Duration
	if m.DurationPrimitiveElement != nil {
		r.Duration.Id = m.DurationPrimitiveElement.Id
		r.Duration.Extension = m.DurationPrimitiveElement.Extension
	}
	r.Content = m.Content
	r.Note = m.Note
	return nil
}
func (r Media) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
