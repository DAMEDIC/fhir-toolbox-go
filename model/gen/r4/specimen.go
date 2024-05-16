package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A sample to be used for analysis.
type Specimen struct {
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
	// Id for specimen.
	Identifier []Identifier
	// The identifier assigned by the lab when accessioning specimen(s). This is not necessarily the same as the specimen identifier, depending on local lab procedures.
	AccessionIdentifier *Identifier
	// The availability of the specimen.
	Status *Code
	// The kind of material that forms the specimen.
	Type *CodeableConcept
	// Where the specimen came from. This may be from patient(s), from a location (e.g., the source of an environmental sample), or a sampling of a substance or a device.
	Subject *Reference
	// Time when specimen was received for processing or testing.
	ReceivedTime *DateTime
	// Reference to the parent (source) specimen which is used when the specimen was either derived from or a component of another specimen.
	Parent []Reference
	// Details concerning a service request that required a specimen to be collected.
	Request []Reference
	// Details concerning the specimen collection.
	Collection *SpecimenCollection
	// Details concerning processing and processing steps for the specimen.
	Processing []SpecimenProcessing
	// The container holding the specimen.  The recursive nature of containers; i.e. blood in tube in tray in rack is not addressed here.
	Container []SpecimenContainer
	// A mode or state of being that describes the nature of the specimen.
	Condition []CodeableConcept
	// To communicate any details or issues about the specimen or during the specimen collection. (for example: broken vial, sent with patient, frozen).
	Note []Annotation
}

func (r Specimen) ResourceType() string {
	return "Specimen"
}
func (r Specimen) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonSpecimen struct {
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
	AccessionIdentifier           *Identifier          `json:"accessionIdentifier,omitempty"`
	Status                        *Code                `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement    `json:"_status,omitempty"`
	Type                          *CodeableConcept     `json:"type,omitempty"`
	Subject                       *Reference           `json:"subject,omitempty"`
	ReceivedTime                  *DateTime            `json:"receivedTime,omitempty"`
	ReceivedTimePrimitiveElement  *primitiveElement    `json:"_receivedTime,omitempty"`
	Parent                        []Reference          `json:"parent,omitempty"`
	Request                       []Reference          `json:"request,omitempty"`
	Collection                    *SpecimenCollection  `json:"collection,omitempty"`
	Processing                    []SpecimenProcessing `json:"processing,omitempty"`
	Container                     []SpecimenContainer  `json:"container,omitempty"`
	Condition                     []CodeableConcept    `json:"condition,omitempty"`
	Note                          []Annotation         `json:"note,omitempty"`
}

func (r Specimen) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Specimen) marshalJSON() jsonSpecimen {
	m := jsonSpecimen{}
	m.ResourceType = "Specimen"
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
	m.AccessionIdentifier = r.AccessionIdentifier
	m.Status = r.Status
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Type = r.Type
	m.Subject = r.Subject
	m.ReceivedTime = r.ReceivedTime
	if r.ReceivedTime != nil && (r.ReceivedTime.Id != nil || r.ReceivedTime.Extension != nil) {
		m.ReceivedTimePrimitiveElement = &primitiveElement{Id: r.ReceivedTime.Id, Extension: r.ReceivedTime.Extension}
	}
	m.Parent = r.Parent
	m.Request = r.Request
	m.Collection = r.Collection
	m.Processing = r.Processing
	m.Container = r.Container
	m.Condition = r.Condition
	m.Note = r.Note
	return m
}
func (r *Specimen) UnmarshalJSON(b []byte) error {
	var m jsonSpecimen
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Specimen) unmarshalJSON(m jsonSpecimen) error {
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
	r.AccessionIdentifier = m.AccessionIdentifier
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Subject = m.Subject
	r.ReceivedTime = m.ReceivedTime
	if m.ReceivedTimePrimitiveElement != nil {
		r.ReceivedTime.Id = m.ReceivedTimePrimitiveElement.Id
		r.ReceivedTime.Extension = m.ReceivedTimePrimitiveElement.Extension
	}
	r.Parent = m.Parent
	r.Request = m.Request
	r.Collection = m.Collection
	r.Processing = m.Processing
	r.Container = m.Container
	r.Condition = m.Condition
	r.Note = m.Note
	return nil
}
func (r Specimen) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Details concerning the specimen collection.
type SpecimenCollection struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Person who collected the specimen.
	Collector *Reference
	// Time when specimen was collected from subject - the physiologically relevant time.
	Collected isSpecimenCollectionCollected
	// The span of time over which the collection of a specimen occurred.
	Duration *Duration
	// The quantity of specimen collected; for instance the volume of a blood sample, or the physical measurement of an anatomic pathology sample.
	Quantity *Quantity
	// A coded value specifying the technique that is used to perform the procedure.
	Method *CodeableConcept
	// Anatomical location from which the specimen was collected (if subject is a patient). This is the target site.  This element is not used for environmental specimens.
	BodySite *CodeableConcept
	// Abstinence or reduction from some or all food, drink, or both, for a period of time prior to sample collection.
	FastingStatus isSpecimenCollectionFastingStatus
}
type isSpecimenCollectionCollected interface {
	isSpecimenCollectionCollected()
}

func (r DateTime) isSpecimenCollectionCollected() {}
func (r Period) isSpecimenCollectionCollected()   {}

type isSpecimenCollectionFastingStatus interface {
	isSpecimenCollectionFastingStatus()
}

func (r CodeableConcept) isSpecimenCollectionFastingStatus() {}
func (r Duration) isSpecimenCollectionFastingStatus()        {}

type jsonSpecimenCollection struct {
	Id                                *string           `json:"id,omitempty"`
	Extension                         []Extension       `json:"extension,omitempty"`
	ModifierExtension                 []Extension       `json:"modifierExtension,omitempty"`
	Collector                         *Reference        `json:"collector,omitempty"`
	CollectedDateTime                 *DateTime         `json:"collectedDateTime,omitempty"`
	CollectedDateTimePrimitiveElement *primitiveElement `json:"_collectedDateTime,omitempty"`
	CollectedPeriod                   *Period           `json:"collectedPeriod,omitempty"`
	Duration                          *Duration         `json:"duration,omitempty"`
	Quantity                          *Quantity         `json:"quantity,omitempty"`
	Method                            *CodeableConcept  `json:"method,omitempty"`
	BodySite                          *CodeableConcept  `json:"bodySite,omitempty"`
	FastingStatusCodeableConcept      *CodeableConcept  `json:"fastingStatusCodeableConcept,omitempty"`
	FastingStatusDuration             *Duration         `json:"fastingStatusDuration,omitempty"`
}

func (r SpecimenCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SpecimenCollection) marshalJSON() jsonSpecimenCollection {
	m := jsonSpecimenCollection{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Collector = r.Collector
	switch v := r.Collected.(type) {
	case DateTime:
		m.CollectedDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.CollectedDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.CollectedDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.CollectedDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.CollectedPeriod = &v
	case *Period:
		m.CollectedPeriod = v
	}
	m.Duration = r.Duration
	m.Quantity = r.Quantity
	m.Method = r.Method
	m.BodySite = r.BodySite
	switch v := r.FastingStatus.(type) {
	case CodeableConcept:
		m.FastingStatusCodeableConcept = &v
	case *CodeableConcept:
		m.FastingStatusCodeableConcept = v
	case Duration:
		m.FastingStatusDuration = &v
	case *Duration:
		m.FastingStatusDuration = v
	}
	return m
}
func (r *SpecimenCollection) UnmarshalJSON(b []byte) error {
	var m jsonSpecimenCollection
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SpecimenCollection) unmarshalJSON(m jsonSpecimenCollection) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Collector = m.Collector
	if m.CollectedDateTime != nil || m.CollectedDateTimePrimitiveElement != nil {
		if r.Collected != nil {
			return fmt.Errorf("multiple values for field \"Collected\"")
		}
		v := m.CollectedDateTime
		if m.CollectedDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.CollectedDateTimePrimitiveElement.Id
			v.Extension = m.CollectedDateTimePrimitiveElement.Extension
		}
		r.Collected = v
	}
	if m.CollectedPeriod != nil {
		if r.Collected != nil {
			return fmt.Errorf("multiple values for field \"Collected\"")
		}
		v := m.CollectedPeriod
		r.Collected = v
	}
	r.Duration = m.Duration
	r.Quantity = m.Quantity
	r.Method = m.Method
	r.BodySite = m.BodySite
	if m.FastingStatusCodeableConcept != nil {
		if r.FastingStatus != nil {
			return fmt.Errorf("multiple values for field \"FastingStatus\"")
		}
		v := m.FastingStatusCodeableConcept
		r.FastingStatus = v
	}
	if m.FastingStatusDuration != nil {
		if r.FastingStatus != nil {
			return fmt.Errorf("multiple values for field \"FastingStatus\"")
		}
		v := m.FastingStatusDuration
		r.FastingStatus = v
	}
	return nil
}
func (r SpecimenCollection) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Details concerning processing and processing steps for the specimen.
type SpecimenProcessing struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Textual description of procedure.
	Description *String
	// A coded value specifying the procedure used to process the specimen.
	Procedure *CodeableConcept
	// Material used in the processing step.
	Additive []Reference
	// A record of the time or period when the specimen processing occurred.  For example the time of sample fixation or the period of time the sample was in formalin.
	Time isSpecimenProcessingTime
}
type isSpecimenProcessingTime interface {
	isSpecimenProcessingTime()
}

func (r DateTime) isSpecimenProcessingTime() {}
func (r Period) isSpecimenProcessingTime()   {}

type jsonSpecimenProcessing struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Description                  *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement  *primitiveElement `json:"_description,omitempty"`
	Procedure                    *CodeableConcept  `json:"procedure,omitempty"`
	Additive                     []Reference       `json:"additive,omitempty"`
	TimeDateTime                 *DateTime         `json:"timeDateTime,omitempty"`
	TimeDateTimePrimitiveElement *primitiveElement `json:"_timeDateTime,omitempty"`
	TimePeriod                   *Period           `json:"timePeriod,omitempty"`
}

func (r SpecimenProcessing) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SpecimenProcessing) marshalJSON() jsonSpecimenProcessing {
	m := jsonSpecimenProcessing{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Procedure = r.Procedure
	m.Additive = r.Additive
	switch v := r.Time.(type) {
	case DateTime:
		m.TimeDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.TimeDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.TimeDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.TimeDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.TimePeriod = &v
	case *Period:
		m.TimePeriod = v
	}
	return m
}
func (r *SpecimenProcessing) UnmarshalJSON(b []byte) error {
	var m jsonSpecimenProcessing
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SpecimenProcessing) unmarshalJSON(m jsonSpecimenProcessing) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Procedure = m.Procedure
	r.Additive = m.Additive
	if m.TimeDateTime != nil || m.TimeDateTimePrimitiveElement != nil {
		if r.Time != nil {
			return fmt.Errorf("multiple values for field \"Time\"")
		}
		v := m.TimeDateTime
		if m.TimeDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.TimeDateTimePrimitiveElement.Id
			v.Extension = m.TimeDateTimePrimitiveElement.Extension
		}
		r.Time = v
	}
	if m.TimePeriod != nil {
		if r.Time != nil {
			return fmt.Errorf("multiple values for field \"Time\"")
		}
		v := m.TimePeriod
		r.Time = v
	}
	return nil
}
func (r SpecimenProcessing) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The container holding the specimen.  The recursive nature of containers; i.e. blood in tube in tray in rack is not addressed here.
type SpecimenContainer struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Id for container. There may be multiple; a manufacturer's bar code, lab assigned identifier, etc. The container ID may differ from the specimen id in some circumstances.
	Identifier []Identifier
	// Textual description of the container.
	Description *String
	// The type of container associated with the specimen (e.g. slide, aliquot, etc.).
	Type *CodeableConcept
	// The capacity (volume or other measure) the container may contain.
	Capacity *Quantity
	// The quantity of specimen in the container; may be volume, dimensions, or other appropriate measurements, depending on the specimen type.
	SpecimenQuantity *Quantity
	// Introduced substance to preserve, maintain or enhance the specimen. Examples: Formalin, Citrate, EDTA.
	Additive isSpecimenContainerAdditive
}
type isSpecimenContainerAdditive interface {
	isSpecimenContainerAdditive()
}

func (r CodeableConcept) isSpecimenContainerAdditive() {}
func (r Reference) isSpecimenContainerAdditive()       {}

type jsonSpecimenContainer struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Identifier                  []Identifier      `json:"identifier,omitempty"`
	Description                 *String           `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
	Type                        *CodeableConcept  `json:"type,omitempty"`
	Capacity                    *Quantity         `json:"capacity,omitempty"`
	SpecimenQuantity            *Quantity         `json:"specimenQuantity,omitempty"`
	AdditiveCodeableConcept     *CodeableConcept  `json:"additiveCodeableConcept,omitempty"`
	AdditiveReference           *Reference        `json:"additiveReference,omitempty"`
}

func (r SpecimenContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SpecimenContainer) marshalJSON() jsonSpecimenContainer {
	m := jsonSpecimenContainer{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Type = r.Type
	m.Capacity = r.Capacity
	m.SpecimenQuantity = r.SpecimenQuantity
	switch v := r.Additive.(type) {
	case CodeableConcept:
		m.AdditiveCodeableConcept = &v
	case *CodeableConcept:
		m.AdditiveCodeableConcept = v
	case Reference:
		m.AdditiveReference = &v
	case *Reference:
		m.AdditiveReference = v
	}
	return m
}
func (r *SpecimenContainer) UnmarshalJSON(b []byte) error {
	var m jsonSpecimenContainer
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SpecimenContainer) unmarshalJSON(m jsonSpecimenContainer) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Capacity = m.Capacity
	r.SpecimenQuantity = m.SpecimenQuantity
	if m.AdditiveCodeableConcept != nil {
		if r.Additive != nil {
			return fmt.Errorf("multiple values for field \"Additive\"")
		}
		v := m.AdditiveCodeableConcept
		r.Additive = v
	}
	if m.AdditiveReference != nil {
		if r.Additive != nil {
			return fmt.Errorf("multiple values for field \"Additive\"")
		}
		v := m.AdditiveReference
		r.Additive = v
	}
	return nil
}
func (r SpecimenContainer) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
