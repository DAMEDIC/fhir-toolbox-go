package r4

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Base StructureDefinition for Timing Type: Specifies an event that may occur multiple times. Timing schedules are used to record when things are planned, expected or requested to occur. The most common usage is in dosage instructions for medications. They are also used when planning care of various kinds, and may be used for reporting the schedule to which past regular activities were carried out.
//
// Need to able to track proposed timing schedules. There are several different ways to do this: one or more specified times, a simple rules like three times a day, or  before/after meals.
type Timing struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identifies specific times when the event occurs.
	Event []DateTime
	// A set of rules that describe when the event is scheduled.
	Repeat *TimingRepeat
	// A code for the timing schedule (or just text in code.text). Some codes such as BID are ubiquitous, but many institutions define their own additional codes. If a code is provided, the code is understood to be a complete statement of whatever is specified in the structured timing data, and either the code or the data may be used to interpret the Timing, with the exception that .repeat.bounds still applies over the code (and is not contained in the code).
	Code *CodeableConcept
}
type jsonTiming struct {
	Id                    *string             `json:"id,omitempty"`
	Extension             []Extension         `json:"extension,omitempty"`
	ModifierExtension     []Extension         `json:"modifierExtension,omitempty"`
	Event                 []DateTime          `json:"event,omitempty"`
	EventPrimitiveElement []*primitiveElement `json:"_event,omitempty"`
	Repeat                *TimingRepeat       `json:"repeat,omitempty"`
	Code                  *CodeableConcept    `json:"code,omitempty"`
}

func (r Timing) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Timing) marshalJSON() jsonTiming {
	m := jsonTiming{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	anyEventValue := false
	for _, e := range r.Event {
		if e.Value != nil {
			anyEventValue = true
			break
		}
	}
	if anyEventValue {
		m.Event = r.Event
	}
	anyEventIdOrExtension := false
	for _, e := range r.Event {
		if e.Id != nil || e.Extension != nil {
			anyEventIdOrExtension = true
			break
		}
	}
	if anyEventIdOrExtension {
		m.EventPrimitiveElement = make([]*primitiveElement, 0, len(r.Event))
		for _, e := range r.Event {
			if e.Id != nil || e.Extension != nil {
				m.EventPrimitiveElement = append(m.EventPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.EventPrimitiveElement = append(m.EventPrimitiveElement, nil)
			}
		}
	}
	m.Repeat = r.Repeat
	m.Code = r.Code
	return m
}
func (r *Timing) UnmarshalJSON(b []byte) error {
	var m jsonTiming
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Timing) unmarshalJSON(m jsonTiming) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Event = m.Event
	for i, e := range m.EventPrimitiveElement {
		if len(r.Event) <= i {
			r.Event = append(r.Event, DateTime{})
		}
		if e != nil {
			r.Event[i].Id = e.Id
			r.Event[i].Extension = e.Extension
		}
	}
	r.Repeat = m.Repeat
	r.Code = m.Code
	return nil
}
func (r Timing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Event, xml.StartElement{Name: xml.Name{Local: "event"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Repeat, xml.StartElement{Name: xml.Name{Local: "repeat"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *Timing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "event":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Event = append(r.Event, v)
			case "repeat":
				var v TimingRepeat
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Repeat = &v
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r Timing) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A set of rules that describe when the event is scheduled.
type TimingRepeat struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Either a duration for the length of the timing schedule, a range of possible length, or outer bounds for start and/or end limits of the timing schedule.
	Bounds isTimingRepeatBounds
	// A total count of the desired number of repetitions across the duration of the entire timing specification. If countMax is present, this element indicates the lower bound of the allowed range of count values.
	Count *PositiveInt
	// If present, indicates that the count is a range - so to perform the action between [count] and [countMax] times.
	CountMax *PositiveInt
	// How long this thing happens for when it happens. If durationMax is present, this element indicates the lower bound of the allowed range of the duration.
	Duration *Decimal
	// If present, indicates that the duration is a range - so to perform the action between [duration] and [durationMax] time length.
	DurationMax *Decimal
	// The units of time for the duration, in UCUM units.
	DurationUnit *Code
	// The number of times to repeat the action within the specified period. If frequencyMax is present, this element indicates the lower bound of the allowed range of the frequency.
	Frequency *PositiveInt
	// If present, indicates that the frequency is a range - so to repeat between [frequency] and [frequencyMax] times within the period or period range.
	FrequencyMax *PositiveInt
	// Indicates the duration of time over which repetitions are to occur; e.g. to express "3 times per day", 3 would be the frequency and "1 day" would be the period. If periodMax is present, this element indicates the lower bound of the allowed range of the period length.
	Period *Decimal
	// If present, indicates that the period is a range from [period] to [periodMax], allowing expressing concepts such as "do this once every 3-5 days.
	PeriodMax *Decimal
	// The units of time for the period in UCUM units.
	PeriodUnit *Code
	// If one or more days of week is provided, then the action happens only on the specified day(s).
	DayOfWeek []Code
	// Specified time of day for action to take place.
	TimeOfDay []Time
	// An approximate time period during the day, potentially linked to an event of daily living that indicates when the action should occur.
	When []Code
	// The number of minutes from the event. If the event code does not indicate whether the minutes is before or after the event, then the offset is assumed to be after the event.
	Offset *UnsignedInt
}
type isTimingRepeatBounds interface {
	isTimingRepeatBounds()
}

func (r Duration) isTimingRepeatBounds() {}
func (r Range) isTimingRepeatBounds()    {}
func (r Period) isTimingRepeatBounds()   {}

type jsonTimingRepeat struct {
	Id                           *string             `json:"id,omitempty"`
	Extension                    []Extension         `json:"extension,omitempty"`
	BoundsDuration               *Duration           `json:"boundsDuration,omitempty"`
	BoundsRange                  *Range              `json:"boundsRange,omitempty"`
	BoundsPeriod                 *Period             `json:"boundsPeriod,omitempty"`
	Count                        *PositiveInt        `json:"count,omitempty"`
	CountPrimitiveElement        *primitiveElement   `json:"_count,omitempty"`
	CountMax                     *PositiveInt        `json:"countMax,omitempty"`
	CountMaxPrimitiveElement     *primitiveElement   `json:"_countMax,omitempty"`
	Duration                     *Decimal            `json:"duration,omitempty"`
	DurationPrimitiveElement     *primitiveElement   `json:"_duration,omitempty"`
	DurationMax                  *Decimal            `json:"durationMax,omitempty"`
	DurationMaxPrimitiveElement  *primitiveElement   `json:"_durationMax,omitempty"`
	DurationUnit                 *Code               `json:"durationUnit,omitempty"`
	DurationUnitPrimitiveElement *primitiveElement   `json:"_durationUnit,omitempty"`
	Frequency                    *PositiveInt        `json:"frequency,omitempty"`
	FrequencyPrimitiveElement    *primitiveElement   `json:"_frequency,omitempty"`
	FrequencyMax                 *PositiveInt        `json:"frequencyMax,omitempty"`
	FrequencyMaxPrimitiveElement *primitiveElement   `json:"_frequencyMax,omitempty"`
	Period                       *Decimal            `json:"period,omitempty"`
	PeriodPrimitiveElement       *primitiveElement   `json:"_period,omitempty"`
	PeriodMax                    *Decimal            `json:"periodMax,omitempty"`
	PeriodMaxPrimitiveElement    *primitiveElement   `json:"_periodMax,omitempty"`
	PeriodUnit                   *Code               `json:"periodUnit,omitempty"`
	PeriodUnitPrimitiveElement   *primitiveElement   `json:"_periodUnit,omitempty"`
	DayOfWeek                    []Code              `json:"dayOfWeek,omitempty"`
	DayOfWeekPrimitiveElement    []*primitiveElement `json:"_dayOfWeek,omitempty"`
	TimeOfDay                    []Time              `json:"timeOfDay,omitempty"`
	TimeOfDayPrimitiveElement    []*primitiveElement `json:"_timeOfDay,omitempty"`
	When                         []Code              `json:"when,omitempty"`
	WhenPrimitiveElement         []*primitiveElement `json:"_when,omitempty"`
	Offset                       *UnsignedInt        `json:"offset,omitempty"`
	OffsetPrimitiveElement       *primitiveElement   `json:"_offset,omitempty"`
}

func (r TimingRepeat) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r TimingRepeat) marshalJSON() jsonTimingRepeat {
	m := jsonTimingRepeat{}
	m.Id = r.Id
	m.Extension = r.Extension
	switch v := r.Bounds.(type) {
	case Duration:
		m.BoundsDuration = &v
	case *Duration:
		m.BoundsDuration = v
	case Range:
		m.BoundsRange = &v
	case *Range:
		m.BoundsRange = v
	case Period:
		m.BoundsPeriod = &v
	case *Period:
		m.BoundsPeriod = v
	}
	if r.Count != nil && r.Count.Value != nil {
		m.Count = r.Count
	}
	if r.Count != nil && (r.Count.Id != nil || r.Count.Extension != nil) {
		m.CountPrimitiveElement = &primitiveElement{Id: r.Count.Id, Extension: r.Count.Extension}
	}
	if r.CountMax != nil && r.CountMax.Value != nil {
		m.CountMax = r.CountMax
	}
	if r.CountMax != nil && (r.CountMax.Id != nil || r.CountMax.Extension != nil) {
		m.CountMaxPrimitiveElement = &primitiveElement{Id: r.CountMax.Id, Extension: r.CountMax.Extension}
	}
	if r.Duration != nil && r.Duration.Value != nil {
		m.Duration = r.Duration
	}
	if r.Duration != nil && (r.Duration.Id != nil || r.Duration.Extension != nil) {
		m.DurationPrimitiveElement = &primitiveElement{Id: r.Duration.Id, Extension: r.Duration.Extension}
	}
	if r.DurationMax != nil && r.DurationMax.Value != nil {
		m.DurationMax = r.DurationMax
	}
	if r.DurationMax != nil && (r.DurationMax.Id != nil || r.DurationMax.Extension != nil) {
		m.DurationMaxPrimitiveElement = &primitiveElement{Id: r.DurationMax.Id, Extension: r.DurationMax.Extension}
	}
	if r.DurationUnit != nil && r.DurationUnit.Value != nil {
		m.DurationUnit = r.DurationUnit
	}
	if r.DurationUnit != nil && (r.DurationUnit.Id != nil || r.DurationUnit.Extension != nil) {
		m.DurationUnitPrimitiveElement = &primitiveElement{Id: r.DurationUnit.Id, Extension: r.DurationUnit.Extension}
	}
	if r.Frequency != nil && r.Frequency.Value != nil {
		m.Frequency = r.Frequency
	}
	if r.Frequency != nil && (r.Frequency.Id != nil || r.Frequency.Extension != nil) {
		m.FrequencyPrimitiveElement = &primitiveElement{Id: r.Frequency.Id, Extension: r.Frequency.Extension}
	}
	if r.FrequencyMax != nil && r.FrequencyMax.Value != nil {
		m.FrequencyMax = r.FrequencyMax
	}
	if r.FrequencyMax != nil && (r.FrequencyMax.Id != nil || r.FrequencyMax.Extension != nil) {
		m.FrequencyMaxPrimitiveElement = &primitiveElement{Id: r.FrequencyMax.Id, Extension: r.FrequencyMax.Extension}
	}
	if r.Period != nil && r.Period.Value != nil {
		m.Period = r.Period
	}
	if r.Period != nil && (r.Period.Id != nil || r.Period.Extension != nil) {
		m.PeriodPrimitiveElement = &primitiveElement{Id: r.Period.Id, Extension: r.Period.Extension}
	}
	if r.PeriodMax != nil && r.PeriodMax.Value != nil {
		m.PeriodMax = r.PeriodMax
	}
	if r.PeriodMax != nil && (r.PeriodMax.Id != nil || r.PeriodMax.Extension != nil) {
		m.PeriodMaxPrimitiveElement = &primitiveElement{Id: r.PeriodMax.Id, Extension: r.PeriodMax.Extension}
	}
	if r.PeriodUnit != nil && r.PeriodUnit.Value != nil {
		m.PeriodUnit = r.PeriodUnit
	}
	if r.PeriodUnit != nil && (r.PeriodUnit.Id != nil || r.PeriodUnit.Extension != nil) {
		m.PeriodUnitPrimitiveElement = &primitiveElement{Id: r.PeriodUnit.Id, Extension: r.PeriodUnit.Extension}
	}
	anyDayOfWeekValue := false
	for _, e := range r.DayOfWeek {
		if e.Value != nil {
			anyDayOfWeekValue = true
			break
		}
	}
	if anyDayOfWeekValue {
		m.DayOfWeek = r.DayOfWeek
	}
	anyDayOfWeekIdOrExtension := false
	for _, e := range r.DayOfWeek {
		if e.Id != nil || e.Extension != nil {
			anyDayOfWeekIdOrExtension = true
			break
		}
	}
	if anyDayOfWeekIdOrExtension {
		m.DayOfWeekPrimitiveElement = make([]*primitiveElement, 0, len(r.DayOfWeek))
		for _, e := range r.DayOfWeek {
			if e.Id != nil || e.Extension != nil {
				m.DayOfWeekPrimitiveElement = append(m.DayOfWeekPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.DayOfWeekPrimitiveElement = append(m.DayOfWeekPrimitiveElement, nil)
			}
		}
	}
	anyTimeOfDayValue := false
	for _, e := range r.TimeOfDay {
		if e.Value != nil {
			anyTimeOfDayValue = true
			break
		}
	}
	if anyTimeOfDayValue {
		m.TimeOfDay = r.TimeOfDay
	}
	anyTimeOfDayIdOrExtension := false
	for _, e := range r.TimeOfDay {
		if e.Id != nil || e.Extension != nil {
			anyTimeOfDayIdOrExtension = true
			break
		}
	}
	if anyTimeOfDayIdOrExtension {
		m.TimeOfDayPrimitiveElement = make([]*primitiveElement, 0, len(r.TimeOfDay))
		for _, e := range r.TimeOfDay {
			if e.Id != nil || e.Extension != nil {
				m.TimeOfDayPrimitiveElement = append(m.TimeOfDayPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.TimeOfDayPrimitiveElement = append(m.TimeOfDayPrimitiveElement, nil)
			}
		}
	}
	anyWhenValue := false
	for _, e := range r.When {
		if e.Value != nil {
			anyWhenValue = true
			break
		}
	}
	if anyWhenValue {
		m.When = r.When
	}
	anyWhenIdOrExtension := false
	for _, e := range r.When {
		if e.Id != nil || e.Extension != nil {
			anyWhenIdOrExtension = true
			break
		}
	}
	if anyWhenIdOrExtension {
		m.WhenPrimitiveElement = make([]*primitiveElement, 0, len(r.When))
		for _, e := range r.When {
			if e.Id != nil || e.Extension != nil {
				m.WhenPrimitiveElement = append(m.WhenPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.WhenPrimitiveElement = append(m.WhenPrimitiveElement, nil)
			}
		}
	}
	if r.Offset != nil && r.Offset.Value != nil {
		m.Offset = r.Offset
	}
	if r.Offset != nil && (r.Offset.Id != nil || r.Offset.Extension != nil) {
		m.OffsetPrimitiveElement = &primitiveElement{Id: r.Offset.Id, Extension: r.Offset.Extension}
	}
	return m
}
func (r *TimingRepeat) UnmarshalJSON(b []byte) error {
	var m jsonTimingRepeat
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *TimingRepeat) unmarshalJSON(m jsonTimingRepeat) error {
	r.Id = m.Id
	r.Extension = m.Extension
	if m.BoundsDuration != nil {
		if r.Bounds != nil {
			return fmt.Errorf("multiple values for field \"Bounds\"")
		}
		v := m.BoundsDuration
		r.Bounds = v
	}
	if m.BoundsRange != nil {
		if r.Bounds != nil {
			return fmt.Errorf("multiple values for field \"Bounds\"")
		}
		v := m.BoundsRange
		r.Bounds = v
	}
	if m.BoundsPeriod != nil {
		if r.Bounds != nil {
			return fmt.Errorf("multiple values for field \"Bounds\"")
		}
		v := m.BoundsPeriod
		r.Bounds = v
	}
	r.Count = m.Count
	if m.CountPrimitiveElement != nil {
		if r.Count == nil {
			r.Count = &PositiveInt{}
		}
		r.Count.Id = m.CountPrimitiveElement.Id
		r.Count.Extension = m.CountPrimitiveElement.Extension
	}
	r.CountMax = m.CountMax
	if m.CountMaxPrimitiveElement != nil {
		if r.CountMax == nil {
			r.CountMax = &PositiveInt{}
		}
		r.CountMax.Id = m.CountMaxPrimitiveElement.Id
		r.CountMax.Extension = m.CountMaxPrimitiveElement.Extension
	}
	r.Duration = m.Duration
	if m.DurationPrimitiveElement != nil {
		if r.Duration == nil {
			r.Duration = &Decimal{}
		}
		r.Duration.Id = m.DurationPrimitiveElement.Id
		r.Duration.Extension = m.DurationPrimitiveElement.Extension
	}
	r.DurationMax = m.DurationMax
	if m.DurationMaxPrimitiveElement != nil {
		if r.DurationMax == nil {
			r.DurationMax = &Decimal{}
		}
		r.DurationMax.Id = m.DurationMaxPrimitiveElement.Id
		r.DurationMax.Extension = m.DurationMaxPrimitiveElement.Extension
	}
	r.DurationUnit = m.DurationUnit
	if m.DurationUnitPrimitiveElement != nil {
		if r.DurationUnit == nil {
			r.DurationUnit = &Code{}
		}
		r.DurationUnit.Id = m.DurationUnitPrimitiveElement.Id
		r.DurationUnit.Extension = m.DurationUnitPrimitiveElement.Extension
	}
	r.Frequency = m.Frequency
	if m.FrequencyPrimitiveElement != nil {
		if r.Frequency == nil {
			r.Frequency = &PositiveInt{}
		}
		r.Frequency.Id = m.FrequencyPrimitiveElement.Id
		r.Frequency.Extension = m.FrequencyPrimitiveElement.Extension
	}
	r.FrequencyMax = m.FrequencyMax
	if m.FrequencyMaxPrimitiveElement != nil {
		if r.FrequencyMax == nil {
			r.FrequencyMax = &PositiveInt{}
		}
		r.FrequencyMax.Id = m.FrequencyMaxPrimitiveElement.Id
		r.FrequencyMax.Extension = m.FrequencyMaxPrimitiveElement.Extension
	}
	r.Period = m.Period
	if m.PeriodPrimitiveElement != nil {
		if r.Period == nil {
			r.Period = &Decimal{}
		}
		r.Period.Id = m.PeriodPrimitiveElement.Id
		r.Period.Extension = m.PeriodPrimitiveElement.Extension
	}
	r.PeriodMax = m.PeriodMax
	if m.PeriodMaxPrimitiveElement != nil {
		if r.PeriodMax == nil {
			r.PeriodMax = &Decimal{}
		}
		r.PeriodMax.Id = m.PeriodMaxPrimitiveElement.Id
		r.PeriodMax.Extension = m.PeriodMaxPrimitiveElement.Extension
	}
	r.PeriodUnit = m.PeriodUnit
	if m.PeriodUnitPrimitiveElement != nil {
		if r.PeriodUnit == nil {
			r.PeriodUnit = &Code{}
		}
		r.PeriodUnit.Id = m.PeriodUnitPrimitiveElement.Id
		r.PeriodUnit.Extension = m.PeriodUnitPrimitiveElement.Extension
	}
	r.DayOfWeek = m.DayOfWeek
	for i, e := range m.DayOfWeekPrimitiveElement {
		if len(r.DayOfWeek) <= i {
			r.DayOfWeek = append(r.DayOfWeek, Code{})
		}
		if e != nil {
			r.DayOfWeek[i].Id = e.Id
			r.DayOfWeek[i].Extension = e.Extension
		}
	}
	r.TimeOfDay = m.TimeOfDay
	for i, e := range m.TimeOfDayPrimitiveElement {
		if len(r.TimeOfDay) <= i {
			r.TimeOfDay = append(r.TimeOfDay, Time{})
		}
		if e != nil {
			r.TimeOfDay[i].Id = e.Id
			r.TimeOfDay[i].Extension = e.Extension
		}
	}
	r.When = m.When
	for i, e := range m.WhenPrimitiveElement {
		if len(r.When) <= i {
			r.When = append(r.When, Code{})
		}
		if e != nil {
			r.When[i].Id = e.Id
			r.When[i].Extension = e.Extension
		}
	}
	r.Offset = m.Offset
	if m.OffsetPrimitiveElement != nil {
		if r.Offset == nil {
			r.Offset = &UnsignedInt{}
		}
		r.Offset.Id = m.OffsetPrimitiveElement.Id
		r.Offset.Extension = m.OffsetPrimitiveElement.Extension
	}
	return nil
}
func (r TimingRepeat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	switch v := r.Bounds.(type) {
	case Duration, *Duration:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "boundsDuration"}})
		if err != nil {
			return err
		}
	case Range, *Range:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "boundsRange"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "boundsPeriod"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Count, xml.StartElement{Name: xml.Name{Local: "count"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CountMax, xml.StartElement{Name: xml.Name{Local: "countMax"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Duration, xml.StartElement{Name: xml.Name{Local: "duration"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DurationMax, xml.StartElement{Name: xml.Name{Local: "durationMax"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DurationUnit, xml.StartElement{Name: xml.Name{Local: "durationUnit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Frequency, xml.StartElement{Name: xml.Name{Local: "frequency"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FrequencyMax, xml.StartElement{Name: xml.Name{Local: "frequencyMax"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Period, xml.StartElement{Name: xml.Name{Local: "period"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PeriodMax, xml.StartElement{Name: xml.Name{Local: "periodMax"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PeriodUnit, xml.StartElement{Name: xml.Name{Local: "periodUnit"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DayOfWeek, xml.StartElement{Name: xml.Name{Local: "dayOfWeek"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.TimeOfDay, xml.StartElement{Name: xml.Name{Local: "timeOfDay"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.When, xml.StartElement{Name: xml.Name{Local: "when"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Offset, xml.StartElement{Name: xml.Name{Local: "offset"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *TimingRepeat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "boundsDuration":
				if r.Bounds != nil {
					return fmt.Errorf("multiple values for field \"Bounds\"")
				}
				var v Duration
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Bounds = v
			case "boundsRange":
				if r.Bounds != nil {
					return fmt.Errorf("multiple values for field \"Bounds\"")
				}
				var v Range
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Bounds = v
			case "boundsPeriod":
				if r.Bounds != nil {
					return fmt.Errorf("multiple values for field \"Bounds\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Bounds = v
			case "count":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Count = &v
			case "countMax":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CountMax = &v
			case "duration":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Duration = &v
			case "durationMax":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DurationMax = &v
			case "durationUnit":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DurationUnit = &v
			case "frequency":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Frequency = &v
			case "frequencyMax":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FrequencyMax = &v
			case "period":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Period = &v
			case "periodMax":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PeriodMax = &v
			case "periodUnit":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PeriodUnit = &v
			case "dayOfWeek":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DayOfWeek = append(r.DayOfWeek, v)
			case "timeOfDay":
				var v Time
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.TimeOfDay = append(r.TimeOfDay, v)
			case "when":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.When = append(r.When, v)
			case "offset":
				var v UnsignedInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Offset = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r TimingRepeat) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
