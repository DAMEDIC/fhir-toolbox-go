package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// An action that is or was performed on or for a patient. This can be a physical intervention like an operation, or less invasive like long term services, counseling, or hypnotherapy.
type Procedure struct {
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
	// Business identifiers assigned to this procedure by the performer or other systems which remain constant as the resource is updated and is propagated from server to server.
	Identifier []Identifier
	// The URL pointing to a FHIR-defined protocol, guideline, order set or other definition that is adhered to in whole or in part by this Procedure.
	InstantiatesCanonical []Canonical
	// The URL pointing to an externally maintained protocol, guideline, order set or other definition that is adhered to in whole or in part by this Procedure.
	InstantiatesUri []Uri
	// A reference to a resource that contains details of the request for this procedure.
	BasedOn []Reference
	// A larger event of which this particular procedure is a component or step.
	PartOf []Reference
	// A code specifying the state of the procedure. Generally, this will be the in-progress or completed state.
	Status Code
	// Captures the reason for the current state of the procedure.
	StatusReason *CodeableConcept
	// A code that classifies the procedure for searching, sorting and display purposes (e.g. "Surgical Procedure").
	Category *CodeableConcept
	// The specific procedure that is performed. Use text if the exact nature of the procedure cannot be coded (e.g. "Laparoscopic Appendectomy").
	Code *CodeableConcept
	// The person, animal or group on which the procedure was performed.
	Subject Reference
	// The Encounter during which this Procedure was created or performed or to which the creation of this record is tightly associated.
	Encounter *Reference
	// Estimated or actual date, date-time, period, or age when the procedure was performed.  Allows a period to support complex procedures that span more than one date, and also allows for the length of the procedure to be captured.
	Performed isProcedurePerformed
	// Individual who recorded the record and takes responsibility for its content.
	Recorder *Reference
	// Individual who is making the procedure statement.
	Asserter *Reference
	// Limited to "real" people rather than equipment.
	Performer []ProcedurePerformer
	// The location where the procedure actually happened.  E.g. a newborn at home, a tracheostomy at a restaurant.
	Location *Reference
	// The coded reason why the procedure was performed. This may be a coded entity of some type, or may simply be present as text.
	ReasonCode []CodeableConcept
	// The justification of why the procedure was performed.
	ReasonReference []Reference
	// Detailed and structured anatomical location information. Multiple locations are allowed - e.g. multiple punch biopsies of a lesion.
	BodySite []CodeableConcept
	// The outcome of the procedure - did it resolve the reasons for the procedure being performed?
	Outcome *CodeableConcept
	// This could be a histology result, pathology report, surgical report, etc.
	Report []Reference
	// Any complications that occurred during the procedure, or in the immediate post-performance period. These are generally tracked separately from the notes, which will typically describe the procedure itself rather than any 'post procedure' issues.
	Complication []CodeableConcept
	// Any complications that occurred during the procedure, or in the immediate post-performance period.
	ComplicationDetail []Reference
	// If the procedure required specific follow up - e.g. removal of sutures. The follow up may be represented as a simple note or could potentially be more complex, in which case the CarePlan resource can be used.
	FollowUp []CodeableConcept
	// Any other notes and comments about the procedure.
	Note []Annotation
	// A device that is implanted, removed or otherwise manipulated (calibration, battery replacement, fitting a prosthesis, attaching a wound-vac, etc.) as a focal portion of the Procedure.
	FocalDevice []ProcedureFocalDevice
	// Identifies medications, devices and any other substance used as part of the procedure.
	UsedReference []Reference
	// Identifies coded items that were used as part of the procedure.
	UsedCode []CodeableConcept
}

func (r Procedure) ResourceType() string {
	return "Procedure"
}
func (r Procedure) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isProcedurePerformed interface {
	isProcedurePerformed()
}

func (r DateTime) isProcedurePerformed() {}
func (r Period) isProcedurePerformed()   {}
func (r String) isProcedurePerformed()   {}
func (r Age) isProcedurePerformed()      {}
func (r Range) isProcedurePerformed()    {}

type jsonProcedure struct {
	ResourceType                          string                 `json:"resourceType"`
	Id                                    *Id                    `json:"id,omitempty"`
	IdPrimitiveElement                    *primitiveElement      `json:"_id,omitempty"`
	Meta                                  *Meta                  `json:"meta,omitempty"`
	ImplicitRules                         *Uri                   `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement         *primitiveElement      `json:"_implicitRules,omitempty"`
	Language                              *Code                  `json:"language,omitempty"`
	LanguagePrimitiveElement              *primitiveElement      `json:"_language,omitempty"`
	Text                                  *Narrative             `json:"text,omitempty"`
	Contained                             []containedResource    `json:"contained,omitempty"`
	Extension                             []Extension            `json:"extension,omitempty"`
	ModifierExtension                     []Extension            `json:"modifierExtension,omitempty"`
	Identifier                            []Identifier           `json:"identifier,omitempty"`
	InstantiatesCanonical                 []Canonical            `json:"instantiatesCanonical,omitempty"`
	InstantiatesCanonicalPrimitiveElement []*primitiveElement    `json:"_instantiatesCanonical,omitempty"`
	InstantiatesUri                       []Uri                  `json:"instantiatesUri,omitempty"`
	InstantiatesUriPrimitiveElement       []*primitiveElement    `json:"_instantiatesUri,omitempty"`
	BasedOn                               []Reference            `json:"basedOn,omitempty"`
	PartOf                                []Reference            `json:"partOf,omitempty"`
	Status                                Code                   `json:"status,omitempty"`
	StatusPrimitiveElement                *primitiveElement      `json:"_status,omitempty"`
	StatusReason                          *CodeableConcept       `json:"statusReason,omitempty"`
	Category                              *CodeableConcept       `json:"category,omitempty"`
	Code                                  *CodeableConcept       `json:"code,omitempty"`
	Subject                               Reference              `json:"subject,omitempty"`
	Encounter                             *Reference             `json:"encounter,omitempty"`
	PerformedDateTime                     *DateTime              `json:"performedDateTime,omitempty"`
	PerformedDateTimePrimitiveElement     *primitiveElement      `json:"_performedDateTime,omitempty"`
	PerformedPeriod                       *Period                `json:"performedPeriod,omitempty"`
	PerformedString                       *String                `json:"performedString,omitempty"`
	PerformedStringPrimitiveElement       *primitiveElement      `json:"_performedString,omitempty"`
	PerformedAge                          *Age                   `json:"performedAge,omitempty"`
	PerformedRange                        *Range                 `json:"performedRange,omitempty"`
	Recorder                              *Reference             `json:"recorder,omitempty"`
	Asserter                              *Reference             `json:"asserter,omitempty"`
	Performer                             []ProcedurePerformer   `json:"performer,omitempty"`
	Location                              *Reference             `json:"location,omitempty"`
	ReasonCode                            []CodeableConcept      `json:"reasonCode,omitempty"`
	ReasonReference                       []Reference            `json:"reasonReference,omitempty"`
	BodySite                              []CodeableConcept      `json:"bodySite,omitempty"`
	Outcome                               *CodeableConcept       `json:"outcome,omitempty"`
	Report                                []Reference            `json:"report,omitempty"`
	Complication                          []CodeableConcept      `json:"complication,omitempty"`
	ComplicationDetail                    []Reference            `json:"complicationDetail,omitempty"`
	FollowUp                              []CodeableConcept      `json:"followUp,omitempty"`
	Note                                  []Annotation           `json:"note,omitempty"`
	FocalDevice                           []ProcedureFocalDevice `json:"focalDevice,omitempty"`
	UsedReference                         []Reference            `json:"usedReference,omitempty"`
	UsedCode                              []CodeableConcept      `json:"usedCode,omitempty"`
}

func (r Procedure) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Procedure) marshalJSON() jsonProcedure {
	m := jsonProcedure{}
	m.ResourceType = "Procedure"
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
	m.InstantiatesCanonical = r.InstantiatesCanonical
	anyInstantiatesCanonicalIdOrExtension := false
	for _, e := range r.InstantiatesCanonical {
		if e.Id != nil || e.Extension != nil {
			anyInstantiatesCanonicalIdOrExtension = true
			break
		}
	}
	if anyInstantiatesCanonicalIdOrExtension {
		m.InstantiatesCanonicalPrimitiveElement = make([]*primitiveElement, 0, len(r.InstantiatesCanonical))
		for _, e := range r.InstantiatesCanonical {
			if e.Id != nil || e.Extension != nil {
				m.InstantiatesCanonicalPrimitiveElement = append(m.InstantiatesCanonicalPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.InstantiatesCanonicalPrimitiveElement = append(m.InstantiatesCanonicalPrimitiveElement, nil)
			}
		}
	}
	m.InstantiatesUri = r.InstantiatesUri
	anyInstantiatesUriIdOrExtension := false
	for _, e := range r.InstantiatesUri {
		if e.Id != nil || e.Extension != nil {
			anyInstantiatesUriIdOrExtension = true
			break
		}
	}
	if anyInstantiatesUriIdOrExtension {
		m.InstantiatesUriPrimitiveElement = make([]*primitiveElement, 0, len(r.InstantiatesUri))
		for _, e := range r.InstantiatesUri {
			if e.Id != nil || e.Extension != nil {
				m.InstantiatesUriPrimitiveElement = append(m.InstantiatesUriPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.InstantiatesUriPrimitiveElement = append(m.InstantiatesUriPrimitiveElement, nil)
			}
		}
	}
	m.BasedOn = r.BasedOn
	m.PartOf = r.PartOf
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.StatusReason = r.StatusReason
	m.Category = r.Category
	m.Code = r.Code
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	switch v := r.Performed.(type) {
	case DateTime:
		m.PerformedDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.PerformedDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.PerformedDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.PerformedDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.PerformedPeriod = &v
	case *Period:
		m.PerformedPeriod = v
	case String:
		m.PerformedString = &v
		if v.Id != nil || v.Extension != nil {
			m.PerformedStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.PerformedString = v
		if v.Id != nil || v.Extension != nil {
			m.PerformedStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Age:
		m.PerformedAge = &v
	case *Age:
		m.PerformedAge = v
	case Range:
		m.PerformedRange = &v
	case *Range:
		m.PerformedRange = v
	}
	m.Recorder = r.Recorder
	m.Asserter = r.Asserter
	m.Performer = r.Performer
	m.Location = r.Location
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.BodySite = r.BodySite
	m.Outcome = r.Outcome
	m.Report = r.Report
	m.Complication = r.Complication
	m.ComplicationDetail = r.ComplicationDetail
	m.FollowUp = r.FollowUp
	m.Note = r.Note
	m.FocalDevice = r.FocalDevice
	m.UsedReference = r.UsedReference
	m.UsedCode = r.UsedCode
	return m
}
func (r *Procedure) UnmarshalJSON(b []byte) error {
	var m jsonProcedure
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Procedure) unmarshalJSON(m jsonProcedure) error {
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
	r.InstantiatesCanonical = m.InstantiatesCanonical
	for i, e := range m.InstantiatesCanonicalPrimitiveElement {
		if len(r.InstantiatesCanonical) > i {
			r.InstantiatesCanonical[i].Id = e.Id
			r.InstantiatesCanonical[i].Extension = e.Extension
		} else {
			r.InstantiatesCanonical = append(r.InstantiatesCanonical, Canonical{Id: e.Id, Extension: e.Extension})
		}
	}
	r.InstantiatesUri = m.InstantiatesUri
	for i, e := range m.InstantiatesUriPrimitiveElement {
		if len(r.InstantiatesUri) > i {
			r.InstantiatesUri[i].Id = e.Id
			r.InstantiatesUri[i].Extension = e.Extension
		} else {
			r.InstantiatesUri = append(r.InstantiatesUri, Uri{Id: e.Id, Extension: e.Extension})
		}
	}
	r.BasedOn = m.BasedOn
	r.PartOf = m.PartOf
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.StatusReason = m.StatusReason
	r.Category = m.Category
	r.Code = m.Code
	r.Subject = m.Subject
	r.Encounter = m.Encounter
	if m.PerformedDateTime != nil || m.PerformedDateTimePrimitiveElement != nil {
		if r.Performed != nil {
			return fmt.Errorf("multiple values for field \"Performed\"")
		}
		v := m.PerformedDateTime
		if m.PerformedDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.PerformedDateTimePrimitiveElement.Id
			v.Extension = m.PerformedDateTimePrimitiveElement.Extension
		}
		r.Performed = v
	}
	if m.PerformedPeriod != nil {
		if r.Performed != nil {
			return fmt.Errorf("multiple values for field \"Performed\"")
		}
		v := m.PerformedPeriod
		r.Performed = v
	}
	if m.PerformedString != nil || m.PerformedStringPrimitiveElement != nil {
		if r.Performed != nil {
			return fmt.Errorf("multiple values for field \"Performed\"")
		}
		v := m.PerformedString
		if m.PerformedStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.PerformedStringPrimitiveElement.Id
			v.Extension = m.PerformedStringPrimitiveElement.Extension
		}
		r.Performed = v
	}
	if m.PerformedAge != nil {
		if r.Performed != nil {
			return fmt.Errorf("multiple values for field \"Performed\"")
		}
		v := m.PerformedAge
		r.Performed = v
	}
	if m.PerformedRange != nil {
		if r.Performed != nil {
			return fmt.Errorf("multiple values for field \"Performed\"")
		}
		v := m.PerformedRange
		r.Performed = v
	}
	r.Recorder = m.Recorder
	r.Asserter = m.Asserter
	r.Performer = m.Performer
	r.Location = m.Location
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.BodySite = m.BodySite
	r.Outcome = m.Outcome
	r.Report = m.Report
	r.Complication = m.Complication
	r.ComplicationDetail = m.ComplicationDetail
	r.FollowUp = m.FollowUp
	r.Note = m.Note
	r.FocalDevice = m.FocalDevice
	r.UsedReference = m.UsedReference
	r.UsedCode = m.UsedCode
	return nil
}
func (r Procedure) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Limited to "real" people rather than equipment.
type ProcedurePerformer struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Distinguishes the type of involvement of the performer in the procedure. For example, surgeon, anaesthetist, endoscopist.
	Function *CodeableConcept
	// The practitioner who was involved in the procedure.
	Actor Reference
	// The organization the device or practitioner was acting on behalf of.
	OnBehalfOf *Reference
}
type jsonProcedurePerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor,omitempty"`
	OnBehalfOf        *Reference       `json:"onBehalfOf,omitempty"`
}

func (r ProcedurePerformer) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ProcedurePerformer) marshalJSON() jsonProcedurePerformer {
	m := jsonProcedurePerformer{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Function = r.Function
	m.Actor = r.Actor
	m.OnBehalfOf = r.OnBehalfOf
	return m
}
func (r *ProcedurePerformer) UnmarshalJSON(b []byte) error {
	var m jsonProcedurePerformer
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ProcedurePerformer) unmarshalJSON(m jsonProcedurePerformer) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Function = m.Function
	r.Actor = m.Actor
	r.OnBehalfOf = m.OnBehalfOf
	return nil
}
func (r ProcedurePerformer) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A device that is implanted, removed or otherwise manipulated (calibration, battery replacement, fitting a prosthesis, attaching a wound-vac, etc.) as a focal portion of the Procedure.
type ProcedureFocalDevice struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The kind of change that happened to the device during the procedure.
	Action *CodeableConcept
	// The device that was manipulated (changed) during the procedure.
	Manipulated Reference
}
type jsonProcedureFocalDevice struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Action            *CodeableConcept `json:"action,omitempty"`
	Manipulated       Reference        `json:"manipulated,omitempty"`
}

func (r ProcedureFocalDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ProcedureFocalDevice) marshalJSON() jsonProcedureFocalDevice {
	m := jsonProcedureFocalDevice{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Action = r.Action
	m.Manipulated = r.Manipulated
	return m
}
func (r *ProcedureFocalDevice) UnmarshalJSON(b []byte) error {
	var m jsonProcedureFocalDevice
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ProcedureFocalDevice) unmarshalJSON(m jsonProcedureFocalDevice) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Action = m.Action
	r.Manipulated = m.Manipulated
	return nil
}
func (r ProcedureFocalDevice) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
