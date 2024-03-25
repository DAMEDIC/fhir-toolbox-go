package r4

import "encoding/json"

// Describes a measurement, calculation or setting capability of a medical device.
type DeviceMetric struct {
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
	Contained []Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Unique instance identifiers assigned to a device by the device or gateway software, manufacturers, other organizations or owners. For example: handle ID.
	Identifier []Identifier
	// Describes the type of the metric. For example: Heart Rate, PEEP Setting, etc.
	Type CodeableConcept
	// Describes the unit that an observed value determined for this metric will have. For example: Percent, Seconds, etc.
	Unit *CodeableConcept
	// Describes the link to the  Device that this DeviceMetric belongs to and that contains administrative device information such as manufacturer, serial number, etc.
	Source *Reference
	// Describes the link to the  Device that this DeviceMetric belongs to and that provide information about the location of this DeviceMetric in the containment structure of the parent Device. An example would be a Device that represents a Channel. This reference can be used by a client application to distinguish DeviceMetrics that have the same type, but should be interpreted based on their containment location.
	Parent *Reference
	// Indicates current operational state of the device. For example: On, Off, Standby, etc.
	OperationalStatus *Code
	// Describes the color representation for the metric. This is often used to aid clinicians to track and identify parameter types by color. In practice, consider a Patient Monitor that has ECG/HR and Pleth for example; the parameters are displayed in different characteristic colors, such as HR-blue, BP-green, and PR and SpO2- magenta.
	Color *Code
	// Indicates the category of the observation generation process. A DeviceMetric can be for example a setting, measurement, or calculation.
	Category Code
	// Describes the measurement repetition time. This is not necessarily the same as the update period. The measurement repetition time can range from milliseconds up to hours. An example for a measurement repetition time in the range of milliseconds is the sampling rate of an ECG. An example for a measurement repetition time in the range of hours is a NIBP that is triggered automatically every hour. The update period may be different than the measurement repetition time, if the device does not update the published observed value with the same frequency as it was measured.
	MeasurementPeriod *Timing
	// Describes the calibrations that have been performed or that are required to be performed.
	Calibration []DeviceMetricCalibration
}
type jsonDeviceMetric struct {
	ResourceType                      string                    `json:"resourceType"`
	Id                                *Id                       `json:"id,omitempty"`
	IdPrimitiveElement                *primitiveElement         `json:"_id,omitempty"`
	Meta                              *Meta                     `json:"meta,omitempty"`
	ImplicitRules                     *Uri                      `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement     *primitiveElement         `json:"_implicitRules,omitempty"`
	Language                          *Code                     `json:"language,omitempty"`
	LanguagePrimitiveElement          *primitiveElement         `json:"_language,omitempty"`
	Text                              *Narrative                `json:"text,omitempty"`
	Contained                         []containedResource       `json:"contained,omitempty"`
	Extension                         []Extension               `json:"extension,omitempty"`
	ModifierExtension                 []Extension               `json:"modifierExtension,omitempty"`
	Identifier                        []Identifier              `json:"identifier,omitempty"`
	Type                              CodeableConcept           `json:"type,omitempty"`
	Unit                              *CodeableConcept          `json:"unit,omitempty"`
	Source                            *Reference                `json:"source,omitempty"`
	Parent                            *Reference                `json:"parent,omitempty"`
	OperationalStatus                 *Code                     `json:"operationalStatus,omitempty"`
	OperationalStatusPrimitiveElement *primitiveElement         `json:"_operationalStatus,omitempty"`
	Color                             *Code                     `json:"color,omitempty"`
	ColorPrimitiveElement             *primitiveElement         `json:"_color,omitempty"`
	Category                          Code                      `json:"category,omitempty"`
	CategoryPrimitiveElement          *primitiveElement         `json:"_category,omitempty"`
	MeasurementPeriod                 *Timing                   `json:"measurementPeriod,omitempty"`
	Calibration                       []DeviceMetricCalibration `json:"calibration,omitempty"`
}

func (r DeviceMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceMetric) marshalJSON() jsonDeviceMetric {
	m := jsonDeviceMetric{}
	m.ResourceType = "DeviceMetric"
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
	m.Type = r.Type
	m.Unit = r.Unit
	m.Source = r.Source
	m.Parent = r.Parent
	m.OperationalStatus = r.OperationalStatus
	if r.OperationalStatus != nil && (r.OperationalStatus.Id != nil || r.OperationalStatus.Extension != nil) {
		m.OperationalStatusPrimitiveElement = &primitiveElement{Id: r.OperationalStatus.Id, Extension: r.OperationalStatus.Extension}
	}
	m.Color = r.Color
	if r.Color != nil && (r.Color.Id != nil || r.Color.Extension != nil) {
		m.ColorPrimitiveElement = &primitiveElement{Id: r.Color.Id, Extension: r.Color.Extension}
	}
	m.Category = r.Category
	if r.Category.Id != nil || r.Category.Extension != nil {
		m.CategoryPrimitiveElement = &primitiveElement{Id: r.Category.Id, Extension: r.Category.Extension}
	}
	m.MeasurementPeriod = r.MeasurementPeriod
	m.Calibration = r.Calibration
	return m
}
func (r *DeviceMetric) UnmarshalJSON(b []byte) error {
	var m jsonDeviceMetric
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceMetric) unmarshalJSON(m jsonDeviceMetric) error {
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
	r.Contained = make([]Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Type = m.Type
	r.Unit = m.Unit
	r.Source = m.Source
	r.Parent = m.Parent
	r.OperationalStatus = m.OperationalStatus
	if m.OperationalStatusPrimitiveElement != nil {
		r.OperationalStatus.Id = m.OperationalStatusPrimitiveElement.Id
		r.OperationalStatus.Extension = m.OperationalStatusPrimitiveElement.Extension
	}
	r.Color = m.Color
	if m.ColorPrimitiveElement != nil {
		r.Color.Id = m.ColorPrimitiveElement.Id
		r.Color.Extension = m.ColorPrimitiveElement.Extension
	}
	r.Category = m.Category
	if m.CategoryPrimitiveElement != nil {
		r.Category.Id = m.CategoryPrimitiveElement.Id
		r.Category.Extension = m.CategoryPrimitiveElement.Extension
	}
	r.MeasurementPeriod = m.MeasurementPeriod
	r.Calibration = m.Calibration
	return nil
}
func (r DeviceMetric) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Describes the calibrations that have been performed or that are required to be performed.
type DeviceMetricCalibration struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Describes the type of the calibration method.
	Type *Code
	// Describes the state of the calibration.
	State *Code
	// Describes the time last calibration has been performed.
	Time *Instant
}
type jsonDeviceMetricCalibration struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Type                  *Code             `json:"type,omitempty"`
	TypePrimitiveElement  *primitiveElement `json:"_type,omitempty"`
	State                 *Code             `json:"state,omitempty"`
	StatePrimitiveElement *primitiveElement `json:"_state,omitempty"`
	Time                  *Instant          `json:"time,omitempty"`
	TimePrimitiveElement  *primitiveElement `json:"_time,omitempty"`
}

func (r DeviceMetricCalibration) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceMetricCalibration) marshalJSON() jsonDeviceMetricCalibration {
	m := jsonDeviceMetricCalibration{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.State = r.State
	if r.State != nil && (r.State.Id != nil || r.State.Extension != nil) {
		m.StatePrimitiveElement = &primitiveElement{Id: r.State.Id, Extension: r.State.Extension}
	}
	m.Time = r.Time
	if r.Time != nil && (r.Time.Id != nil || r.Time.Extension != nil) {
		m.TimePrimitiveElement = &primitiveElement{Id: r.Time.Id, Extension: r.Time.Extension}
	}
	return m
}
func (r *DeviceMetricCalibration) UnmarshalJSON(b []byte) error {
	var m jsonDeviceMetricCalibration
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceMetricCalibration) unmarshalJSON(m jsonDeviceMetricCalibration) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.State = m.State
	if m.StatePrimitiveElement != nil {
		r.State.Id = m.StatePrimitiveElement.Id
		r.State.Extension = m.StatePrimitiveElement.Extension
	}
	r.Time = m.Time
	if m.TimePrimitiveElement != nil {
		r.Time.Id = m.TimePrimitiveElement.Id
		r.Time.Extension = m.TimePrimitiveElement.Extension
	}
	return nil
}
func (r DeviceMetricCalibration) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
