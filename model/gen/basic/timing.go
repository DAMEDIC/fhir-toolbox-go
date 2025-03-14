package basic

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	model "github.com/DAMEDIC/fhir-toolbox-go/model"
	"io"
	"slices"
	"unsafe"
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
	model.Element
	isTimingRepeatBounds()
}

func (r Duration) isTimingRepeatBounds() {}
func (r Range) isTimingRepeatBounds()    {}
func (r Period) isTimingRepeatBounds()   {}
func (r Timing) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Event {
		s += i.MemSize()
	}
	s += (cap(r.Event) - len(r.Event)) * int(unsafe.Sizeof(DateTime{}))
	if r.Repeat != nil {
		s += r.Repeat.MemSize()
	}
	if r.Code != nil {
		s += r.Code.MemSize()
	}
	return s
}
func (r TimingRepeat) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	if r.Bounds != nil {
		s += r.Bounds.MemSize()
	}
	if r.Count != nil {
		s += r.Count.MemSize()
	}
	if r.CountMax != nil {
		s += r.CountMax.MemSize()
	}
	if r.Duration != nil {
		s += r.Duration.MemSize()
	}
	if r.DurationMax != nil {
		s += r.DurationMax.MemSize()
	}
	if r.DurationUnit != nil {
		s += r.DurationUnit.MemSize()
	}
	if r.Frequency != nil {
		s += r.Frequency.MemSize()
	}
	if r.FrequencyMax != nil {
		s += r.FrequencyMax.MemSize()
	}
	if r.Period != nil {
		s += r.Period.MemSize()
	}
	if r.PeriodMax != nil {
		s += r.PeriodMax.MemSize()
	}
	if r.PeriodUnit != nil {
		s += r.PeriodUnit.MemSize()
	}
	for _, i := range r.DayOfWeek {
		s += i.MemSize()
	}
	s += (cap(r.DayOfWeek) - len(r.DayOfWeek)) * int(unsafe.Sizeof(Code{}))
	for _, i := range r.TimeOfDay {
		s += i.MemSize()
	}
	s += (cap(r.TimeOfDay) - len(r.TimeOfDay)) * int(unsafe.Sizeof(Time{}))
	for _, i := range r.When {
		s += i.MemSize()
	}
	s += (cap(r.When) - len(r.When)) * int(unsafe.Sizeof(Code{}))
	if r.Offset != nil {
		s += r.Offset.MemSize()
	}
	return s
}
func (r Timing) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r TimingRepeat) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r Timing) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r Timing) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	{
		anyValue := false
		for _, e := range r.Event {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"event\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.Event)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.Event {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_event\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.Event {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	if r.Repeat != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"repeat\":"))
		if err != nil {
			return err
		}
		err = r.Repeat.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Code != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"code\":"))
		if err != nil {
			return err
		}
		err = r.Code.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r TimingRepeat) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r TimingRepeat) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	switch v := r.Bounds.(type) {
	case Duration:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"boundsDuration\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Duration:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"boundsDuration\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Range:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"boundsRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Range:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"boundsRange\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"boundsPeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Period:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"boundsPeriod\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	}
	if r.Count != nil && r.Count.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"count\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Count)
		if err != nil {
			return err
		}
	}
	if r.Count != nil && (r.Count.Id != nil || r.Count.Extension != nil) {
		p := primitiveElement{Id: r.Count.Id, Extension: r.Count.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_count\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.CountMax != nil && r.CountMax.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"countMax\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.CountMax)
		if err != nil {
			return err
		}
	}
	if r.CountMax != nil && (r.CountMax.Id != nil || r.CountMax.Extension != nil) {
		p := primitiveElement{Id: r.CountMax.Id, Extension: r.CountMax.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_countMax\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Duration != nil && r.Duration.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"duration\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Duration)
		if err != nil {
			return err
		}
	}
	if r.Duration != nil && (r.Duration.Id != nil || r.Duration.Extension != nil) {
		p := primitiveElement{Id: r.Duration.Id, Extension: r.Duration.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_duration\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.DurationMax != nil && r.DurationMax.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"durationMax\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.DurationMax)
		if err != nil {
			return err
		}
	}
	if r.DurationMax != nil && (r.DurationMax.Id != nil || r.DurationMax.Extension != nil) {
		p := primitiveElement{Id: r.DurationMax.Id, Extension: r.DurationMax.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_durationMax\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.DurationUnit != nil && r.DurationUnit.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"durationUnit\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.DurationUnit)
		if err != nil {
			return err
		}
	}
	if r.DurationUnit != nil && (r.DurationUnit.Id != nil || r.DurationUnit.Extension != nil) {
		p := primitiveElement{Id: r.DurationUnit.Id, Extension: r.DurationUnit.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_durationUnit\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Frequency != nil && r.Frequency.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"frequency\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Frequency)
		if err != nil {
			return err
		}
	}
	if r.Frequency != nil && (r.Frequency.Id != nil || r.Frequency.Extension != nil) {
		p := primitiveElement{Id: r.Frequency.Id, Extension: r.Frequency.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_frequency\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.FrequencyMax != nil && r.FrequencyMax.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"frequencyMax\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.FrequencyMax)
		if err != nil {
			return err
		}
	}
	if r.FrequencyMax != nil && (r.FrequencyMax.Id != nil || r.FrequencyMax.Extension != nil) {
		p := primitiveElement{Id: r.FrequencyMax.Id, Extension: r.FrequencyMax.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_frequencyMax\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Period != nil && r.Period.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"period\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Period)
		if err != nil {
			return err
		}
	}
	if r.Period != nil && (r.Period.Id != nil || r.Period.Extension != nil) {
		p := primitiveElement{Id: r.Period.Id, Extension: r.Period.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_period\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.PeriodMax != nil && r.PeriodMax.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"periodMax\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.PeriodMax)
		if err != nil {
			return err
		}
	}
	if r.PeriodMax != nil && (r.PeriodMax.Id != nil || r.PeriodMax.Extension != nil) {
		p := primitiveElement{Id: r.PeriodMax.Id, Extension: r.PeriodMax.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_periodMax\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.PeriodUnit != nil && r.PeriodUnit.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"periodUnit\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.PeriodUnit)
		if err != nil {
			return err
		}
	}
	if r.PeriodUnit != nil && (r.PeriodUnit.Id != nil || r.PeriodUnit.Extension != nil) {
		p := primitiveElement{Id: r.PeriodUnit.Id, Extension: r.PeriodUnit.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_periodUnit\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		anyValue := false
		for _, e := range r.DayOfWeek {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"dayOfWeek\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.DayOfWeek)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.DayOfWeek {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_dayOfWeek\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.DayOfWeek {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	{
		anyValue := false
		for _, e := range r.TimeOfDay {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"timeOfDay\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.TimeOfDay)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.TimeOfDay {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_timeOfDay\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.TimeOfDay {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	{
		anyValue := false
		for _, e := range r.When {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"when\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.When)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.When {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_when\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.When {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	if r.Offset != nil && r.Offset.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"offset\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Offset)
		if err != nil {
			return err
		}
	}
	if r.Offset != nil && (r.Offset.Id != nil || r.Offset.Extension != nil) {
		p := primitiveElement{Id: r.Offset.Id, Extension: r.Offset.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_offset\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
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
func (r Timing) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "event") {
		for _, v := range r.Event {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "repeat") {
		if r.Repeat != nil {
			children = append(children, *r.Repeat)
		}
	}
	if len(name) == 0 || slices.Contains(name, "code") {
		if r.Code != nil {
			children = append(children, *r.Code)
		}
	}
	return children
}
func (r Timing) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert Timing to Boolean")
}
func (r Timing) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert Timing to String")
}
func (r Timing) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert Timing to Integer")
}
func (r Timing) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert Timing to Decimal")
}
func (r Timing) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert Timing to Date")
}
func (r Timing) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert Timing to Time")
}
func (r Timing) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert Timing to DateTime")
}
func (r Timing) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert Timing to Quantity")
}
func (r Timing) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o Timing
	switch other := other.(type) {
	case Timing:
		o = other
	case *Timing:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r Timing) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o Timing
	switch other := other.(type) {
	case Timing:
		o = other
	case *Timing:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r Timing) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "ModifierExtension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Event",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "DateTime",
				Namespace: "FHIR",
			},
		}, {
			Name: "Repeat",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "TimingRepeat",
				Namespace: "FHIR",
			},
		}, {
			Name: "Code",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "CodeableConcept",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "Timing",
			Namespace: "FHIR",
		},
	}
}
func (r TimingRepeat) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "bounds") {
		if r.Bounds != nil {
			children = append(children, r.Bounds)
		}
	}
	if len(name) == 0 || slices.Contains(name, "count") {
		if r.Count != nil {
			children = append(children, *r.Count)
		}
	}
	if len(name) == 0 || slices.Contains(name, "countMax") {
		if r.CountMax != nil {
			children = append(children, *r.CountMax)
		}
	}
	if len(name) == 0 || slices.Contains(name, "duration") {
		if r.Duration != nil {
			children = append(children, *r.Duration)
		}
	}
	if len(name) == 0 || slices.Contains(name, "durationMax") {
		if r.DurationMax != nil {
			children = append(children, *r.DurationMax)
		}
	}
	if len(name) == 0 || slices.Contains(name, "durationUnit") {
		if r.DurationUnit != nil {
			children = append(children, *r.DurationUnit)
		}
	}
	if len(name) == 0 || slices.Contains(name, "frequency") {
		if r.Frequency != nil {
			children = append(children, *r.Frequency)
		}
	}
	if len(name) == 0 || slices.Contains(name, "frequencyMax") {
		if r.FrequencyMax != nil {
			children = append(children, *r.FrequencyMax)
		}
	}
	if len(name) == 0 || slices.Contains(name, "period") {
		if r.Period != nil {
			children = append(children, *r.Period)
		}
	}
	if len(name) == 0 || slices.Contains(name, "periodMax") {
		if r.PeriodMax != nil {
			children = append(children, *r.PeriodMax)
		}
	}
	if len(name) == 0 || slices.Contains(name, "periodUnit") {
		if r.PeriodUnit != nil {
			children = append(children, *r.PeriodUnit)
		}
	}
	if len(name) == 0 || slices.Contains(name, "dayOfWeek") {
		for _, v := range r.DayOfWeek {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "timeOfDay") {
		for _, v := range r.TimeOfDay {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "when") {
		for _, v := range r.When {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "offset") {
		if r.Offset != nil {
			children = append(children, *r.Offset)
		}
	}
	return children
}
func (r TimingRepeat) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert TimingRepeat to Boolean")
}
func (r TimingRepeat) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert TimingRepeat to String")
}
func (r TimingRepeat) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert TimingRepeat to Integer")
}
func (r TimingRepeat) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert TimingRepeat to Decimal")
}
func (r TimingRepeat) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert TimingRepeat to Date")
}
func (r TimingRepeat) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert TimingRepeat to Time")
}
func (r TimingRepeat) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert TimingRepeat to DateTime")
}
func (r TimingRepeat) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert TimingRepeat to Quantity")
}
func (r TimingRepeat) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o TimingRepeat
	switch other := other.(type) {
	case TimingRepeat:
		o = other
	case *TimingRepeat:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equal(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r TimingRepeat) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	var o TimingRepeat
	switch other := other.(type) {
	case TimingRepeat:
		o = other
	case *TimingRepeat:
		o = *other
	default:
		return false
	}
	eq := r.Children().Equivalent(o.Children())
	if eq == nil {
		return true
	}
	return *eq
}
func (r TimingRepeat) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "string",
				Namespace: "FHIR",
			},
		}, {
			Name: "Extension",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Extension",
				Namespace: "FHIR",
			},
		}, {
			Name: "Bounds",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PrimitiveElement",
				Namespace: "FHIR",
			},
		}, {
			Name: "Count",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "CountMax",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "Duration",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "DurationMax",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "DurationUnit",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Frequency",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "FrequencyMax",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "PositiveInt",
				Namespace: "FHIR",
			},
		}, {
			Name: "Period",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "PeriodMax",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Decimal",
				Namespace: "FHIR",
			},
		}, {
			Name: "PeriodUnit",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "DayOfWeek",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "TimeOfDay",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Time",
				Namespace: "FHIR",
			},
		}, {
			Name: "When",
			Type: fhirpath.TypeSpecifier{
				List:      true,
				Name:      "Code",
				Namespace: "FHIR",
			},
		}, {
			Name: "Offset",
			Type: fhirpath.TypeSpecifier{
				List:      false,
				Name:      "UnsignedInt",
				Namespace: "FHIR",
			},
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "TimingRepeat",
			Namespace: "FHIR",
		},
	}
}
