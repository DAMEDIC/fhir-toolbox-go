package r5

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
	"io"
	"unsafe"
)

// A record of significant events/milestones key data throughout the history of an Encounter
type EncounterHistory struct {
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
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, nor can they have their own independent transaction scope. This is allowed to be a Parameters resource if and only if it is referenced by a resource that provides context/meaning.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The Encounter associated with this set of historic values.
	Encounter *Reference
	// Identifier(s) by which this encounter is known.
	Identifier []Identifier
	// planned | in-progress | on-hold | discharged | completed | cancelled | discontinued | entered-in-error | unknown.
	Status Code
	// Concepts representing classification of patient encounter such as ambulatory (outpatient), inpatient, emergency, home health or others due to local variations.
	Class CodeableConcept
	// Specific type of encounter (e.g. e-mail consultation, surgical day-care, skilled nursing, rehabilitation).
	Type []CodeableConcept
	// Broad categorization of the service that is to be provided (e.g. cardiology).
	ServiceType []CodeableReference
	// The patient or group related to this encounter. In some use-cases the patient MAY not be present, such as a case meeting about a patient between several practitioners or a careteam.
	Subject *Reference
	// The subjectStatus value can be used to track the patient's status within the encounter. It details whether the patient has arrived or departed, has been triaged or is currently in a waiting status.
	SubjectStatus *CodeableConcept
	// The start and end time associated with this set of values associated with the encounter, may be different to the planned times for various reasons.
	ActualPeriod *Period
	// The planned start date/time (or admission date) of the encounter.
	PlannedStartDate *DateTime
	// The planned end date/time (or discharge date) of the encounter.
	PlannedEndDate *DateTime
	// Actual quantity of time the encounter lasted. This excludes the time during leaves of absence.When missing it is the time in between the start and end values.
	Length *Duration
	// The location of the patient at this point in the encounter, the multiple cardinality permits de-normalizing the levels of the location hierarchy, such as site/ward/room/bed.
	Location []EncounterHistoryLocation
}

// The location of the patient at this point in the encounter, the multiple cardinality permits de-normalizing the levels of the location hierarchy, such as site/ward/room/bed.
type EncounterHistoryLocation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The location where the encounter takes place.
	Location Reference
	// This will be used to specify the required levels (bed/ward/room/etc.) desired to be recorded to simplify either messaging or query.
	Form *CodeableConcept
}

func (r EncounterHistory) ResourceType() string {
	return "EncounterHistory"
}
func (r EncounterHistory) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r EncounterHistory) MemSize() int {
	var emptyIface any
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += r.Id.MemSize()
	}
	if r.Meta != nil {
		s += r.Meta.MemSize()
	}
	if r.ImplicitRules != nil {
		s += r.ImplicitRules.MemSize()
	}
	if r.Language != nil {
		s += r.Language.MemSize()
	}
	if r.Text != nil {
		s += r.Text.MemSize()
	}
	for _, i := range r.Contained {
		s += i.MemSize()
	}
	s += (cap(r.Contained) - len(r.Contained)) * int(unsafe.Sizeof(emptyIface))
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Encounter != nil {
		s += r.Encounter.MemSize()
	}
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	s += r.Status.MemSize() - int(unsafe.Sizeof(r.Status))
	s += r.Class.MemSize() - int(unsafe.Sizeof(r.Class))
	for _, i := range r.Type {
		s += i.MemSize()
	}
	s += (cap(r.Type) - len(r.Type)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.ServiceType {
		s += i.MemSize()
	}
	s += (cap(r.ServiceType) - len(r.ServiceType)) * int(unsafe.Sizeof(CodeableReference{}))
	if r.Subject != nil {
		s += r.Subject.MemSize()
	}
	if r.SubjectStatus != nil {
		s += r.SubjectStatus.MemSize()
	}
	if r.ActualPeriod != nil {
		s += r.ActualPeriod.MemSize()
	}
	if r.PlannedStartDate != nil {
		s += r.PlannedStartDate.MemSize()
	}
	if r.PlannedEndDate != nil {
		s += r.PlannedEndDate.MemSize()
	}
	if r.Length != nil {
		s += r.Length.MemSize()
	}
	for _, i := range r.Location {
		s += i.MemSize()
	}
	s += (cap(r.Location) - len(r.Location)) * int(unsafe.Sizeof(EncounterHistoryLocation{}))
	return s
}
func (r EncounterHistoryLocation) MemSize() int {
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
	s += r.Location.MemSize() - int(unsafe.Sizeof(r.Location))
	if r.Form != nil {
		s += r.Form.MemSize()
	}
	return s
}
func (r EncounterHistory) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r EncounterHistory) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r EncounterHistory) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"EncounterHistory\""))
	if err != nil {
		return err
	}
	setComma := true
	if r.Id != nil && r.Id.Value != nil {
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
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		p := primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_id\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Meta != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"meta\":"))
		if err != nil {
			return err
		}
		err = r.Meta.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"implicitRules\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ImplicitRules)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		p := primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_implicitRules\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && r.Language.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"language\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Language)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		p := primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_language\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Text != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"text\":"))
		if err != nil {
			return err
		}
		err = r.Text.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Contained) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"contained\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, c := range r.Contained {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = ContainedResource{c}.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
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
	if r.Encounter != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"encounter\":"))
		if err != nil {
			return err
		}
		err = r.Encounter.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Identifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"identifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Identifier {
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
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"status\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Status)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		p := primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_status\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"class\":"))
	if err != nil {
		return err
	}
	err = r.Class.marshalJSON(w)
	if err != nil {
		return err
	}
	if len(r.Type) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"type\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Type {
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
	if len(r.ServiceType) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"serviceType\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ServiceType {
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
	if r.Subject != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"subject\":"))
		if err != nil {
			return err
		}
		err = r.Subject.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.SubjectStatus != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"subjectStatus\":"))
		if err != nil {
			return err
		}
		err = r.SubjectStatus.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ActualPeriod != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"actualPeriod\":"))
		if err != nil {
			return err
		}
		err = r.ActualPeriod.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.PlannedStartDate != nil && r.PlannedStartDate.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"plannedStartDate\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.PlannedStartDate)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if r.PlannedStartDate != nil && (r.PlannedStartDate.Id != nil || r.PlannedStartDate.Extension != nil) {
		p := primitiveElement{Id: r.PlannedStartDate.Id, Extension: r.PlannedStartDate.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_plannedStartDate\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.PlannedEndDate != nil && r.PlannedEndDate.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"plannedEndDate\":"))
		if err != nil {
			return err
		}
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.PlannedEndDate)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
		if err != nil {
			return err
		}
	}
	if r.PlannedEndDate != nil && (r.PlannedEndDate.Id != nil || r.PlannedEndDate.Extension != nil) {
		p := primitiveElement{Id: r.PlannedEndDate.Id, Extension: r.PlannedEndDate.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_plannedEndDate\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Length != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"length\":"))
		if err != nil {
			return err
		}
		err = r.Length.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Location) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"location\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Location {
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
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r EncounterHistoryLocation) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r EncounterHistoryLocation) marshalJSON(w io.Writer) error {
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
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
		_, err = w.Write(b.Bytes())
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
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"location\":"))
	if err != nil {
		return err
	}
	err = r.Location.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.Form != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"form\":"))
		if err != nil {
			return err
		}
		err = r.Form.marshalJSON(w)
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
func (r *EncounterHistory) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *EncounterHistory) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in EncounterHistory element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in EncounterHistory element", t)
		}
		switch f {
		case "resourceType":
			_, err := d.Token()
			if err != nil {
				return err
			}
		case "id":
			var v Id
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Value = v.Value
		case "_id":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Id = v.Id
			r.Id.Extension = v.Extension
		case "meta":
			var v Meta
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Meta = &v
		case "implicitRules":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Value = v.Value
		case "_implicitRules":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Id = v.Id
			r.ImplicitRules.Extension = v.Extension
		case "language":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Value = v.Value
		case "_language":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Id = v.Id
			r.Language.Extension = v.Extension
		case "text":
			var v Narrative
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Text = &v
		case "contained":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in EncounterHistory element", t)
			}
			for d.More() {
				var v ContainedResource
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Contained = append(r.Contained, v.Resource)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in EncounterHistory element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in EncounterHistory element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in EncounterHistory element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in EncounterHistory element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in EncounterHistory element", t)
			}
		case "encounter":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Encounter = &v
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in EncounterHistory element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in EncounterHistory element", t)
			}
		case "status":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Status.Value = v.Value
		case "_status":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Status.Id = v.Id
			r.Status.Extension = v.Extension
		case "class":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Class = v
		case "type":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in EncounterHistory element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Type = append(r.Type, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in EncounterHistory element", t)
			}
		case "serviceType":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in EncounterHistory element", t)
			}
			for d.More() {
				var v CodeableReference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ServiceType = append(r.ServiceType, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in EncounterHistory element", t)
			}
		case "subject":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Subject = &v
		case "subjectStatus":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.SubjectStatus = &v
		case "actualPeriod":
			var v Period
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.ActualPeriod = &v
		case "plannedStartDate":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.PlannedStartDate == nil {
				r.PlannedStartDate = &DateTime{}
			}
			r.PlannedStartDate.Value = v.Value
		case "_plannedStartDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.PlannedStartDate == nil {
				r.PlannedStartDate = &DateTime{}
			}
			r.PlannedStartDate.Id = v.Id
			r.PlannedStartDate.Extension = v.Extension
		case "plannedEndDate":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.PlannedEndDate == nil {
				r.PlannedEndDate = &DateTime{}
			}
			r.PlannedEndDate.Value = v.Value
		case "_plannedEndDate":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.PlannedEndDate == nil {
				r.PlannedEndDate = &DateTime{}
			}
			r.PlannedEndDate.Id = v.Id
			r.PlannedEndDate.Extension = v.Extension
		case "length":
			var v Duration
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Length = &v
		case "location":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in EncounterHistory element", t)
			}
			for d.More() {
				var v EncounterHistoryLocation
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Location = append(r.Location, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in EncounterHistory element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in EncounterHistory", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in EncounterHistory element", t)
	}
	return nil
}
func (r *EncounterHistoryLocation) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in EncounterHistoryLocation element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in EncounterHistoryLocation element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in EncounterHistoryLocation element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in EncounterHistoryLocation element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in EncounterHistoryLocation element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in EncounterHistoryLocation element", t)
			}
		case "location":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Location = v
		case "form":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Form = &v
		default:
			return fmt.Errorf("invalid field: %s in EncounterHistoryLocation", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in EncounterHistoryLocation element", t)
	}
	return nil
}
func (r EncounterHistory) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "EncounterHistory"
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
	err = e.EncodeElement(r.Encounter, xml.StartElement{Name: xml.Name{Local: "encounter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Class, xml.StartElement{Name: xml.Name{Local: "class"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ServiceType, xml.StartElement{Name: xml.Name{Local: "serviceType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Subject, xml.StartElement{Name: xml.Name{Local: "subject"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubjectStatus, xml.StartElement{Name: xml.Name{Local: "subjectStatus"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ActualPeriod, xml.StartElement{Name: xml.Name{Local: "actualPeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PlannedStartDate, xml.StartElement{Name: xml.Name{Local: "plannedStartDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PlannedEndDate, xml.StartElement{Name: xml.Name{Local: "plannedEndDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Length, xml.StartElement{Name: xml.Name{Local: "length"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Location, xml.StartElement{Name: xml.Name{Local: "location"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r EncounterHistoryLocation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Location, xml.StartElement{Name: xml.Name{Local: "location"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Form, xml.StartElement{Name: xml.Name{Local: "form"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *EncounterHistory) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "encounter":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Encounter = &v
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "class":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Class = v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = append(r.Type, v)
			case "serviceType":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ServiceType = append(r.ServiceType, v)
			case "subject":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = &v
			case "subjectStatus":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubjectStatus = &v
			case "actualPeriod":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ActualPeriod = &v
			case "plannedStartDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PlannedStartDate = &v
			case "plannedEndDate":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PlannedEndDate = &v
			case "length":
				var v Duration
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Length = &v
			case "location":
				var v EncounterHistoryLocation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = append(r.Location, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *EncounterHistoryLocation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "location":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = v
			case "form":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Form = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
