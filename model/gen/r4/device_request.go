package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// Represents a request for a patient to employ a medical device. The device may be an implantable device, or an external assistive device, such as a walker.
type DeviceRequest struct {
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
	// Identifiers assigned to this order by the orderer or by the receiver.
	Identifier []Identifier
	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this DeviceRequest.
	InstantiatesCanonical []Canonical
	// The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this DeviceRequest.
	InstantiatesUri []Uri
	// Plan/proposal/order fulfilled by this request.
	BasedOn []Reference
	// The request takes the place of the referenced completed or terminated request(s).
	PriorRequest []Reference
	// Composite request this is part of.
	GroupIdentifier *Identifier
	// The status of the request.
	Status *Code
	// Whether the request is a proposal, plan, an original order or a reflex order.
	Intent Code
	// Indicates how quickly the {{title}} should be addressed with respect to other requests.
	Priority *Code
	// The details of the device to be used.
	Code isDeviceRequestCode
	// Specific parameters for the ordered item.  For example, the prism value for lenses.
	Parameter []DeviceRequestParameter
	// The patient who will use the device.
	Subject Reference
	// An encounter that provides additional context in which this request is made.
	Encounter *Reference
	// The timing schedule for the use of the device. The Schedule data type allows many different expressions, for example. "Every 8 hours"; "Three times a day"; "1/2 an hour before breakfast for 10 days from 23-Dec 2011:"; "15 Oct 2013, 17 Oct 2013 and 1 Nov 2013".
	Occurrence isDeviceRequestOccurrence
	// When the request transitioned to being actionable.
	AuthoredOn *DateTime
	// The individual who initiated the request and has responsibility for its activation.
	Requester *Reference
	// Desired type of performer for doing the diagnostic testing.
	PerformerType *CodeableConcept
	// The desired performer for doing the diagnostic testing.
	Performer *Reference
	// Reason or justification for the use of this device.
	ReasonCode []CodeableConcept
	// Reason or justification for the use of this device.
	ReasonReference []Reference
	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be required for delivering the requested service.
	Insurance []Reference
	// Additional clinical information about the patient that may influence the request fulfilment.  For example, this may include where on the subject's body the device will be used (i.e. the target site).
	SupportingInfo []Reference
	// Details about this request that were not represented at all or sufficiently in one of the attributes provided in a class. These may include for example a comment, an instruction, or a note associated with the statement.
	Note []Annotation
	// Key events in the history of the request.
	RelevantHistory []Reference
}

func (r DeviceRequest) ResourceType() string {
	return "DeviceRequest"
}
func (r DeviceRequest) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isDeviceRequestCode interface {
	isDeviceRequestCode()
}

func (r Reference) isDeviceRequestCode()       {}
func (r CodeableConcept) isDeviceRequestCode() {}

type isDeviceRequestOccurrence interface {
	isDeviceRequestOccurrence()
}

func (r DateTime) isDeviceRequestOccurrence() {}
func (r Period) isDeviceRequestOccurrence()   {}
func (r Timing) isDeviceRequestOccurrence()   {}

type jsonDeviceRequest struct {
	ResourceType                          string                   `json:"resourceType"`
	Id                                    *Id                      `json:"id,omitempty"`
	IdPrimitiveElement                    *primitiveElement        `json:"_id,omitempty"`
	Meta                                  *Meta                    `json:"meta,omitempty"`
	ImplicitRules                         *Uri                     `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement         *primitiveElement        `json:"_implicitRules,omitempty"`
	Language                              *Code                    `json:"language,omitempty"`
	LanguagePrimitiveElement              *primitiveElement        `json:"_language,omitempty"`
	Text                                  *Narrative               `json:"text,omitempty"`
	Contained                             []ContainedResource      `json:"contained,omitempty"`
	Extension                             []Extension              `json:"extension,omitempty"`
	ModifierExtension                     []Extension              `json:"modifierExtension,omitempty"`
	Identifier                            []Identifier             `json:"identifier,omitempty"`
	InstantiatesCanonical                 []Canonical              `json:"instantiatesCanonical,omitempty"`
	InstantiatesCanonicalPrimitiveElement []*primitiveElement      `json:"_instantiatesCanonical,omitempty"`
	InstantiatesUri                       []Uri                    `json:"instantiatesUri,omitempty"`
	InstantiatesUriPrimitiveElement       []*primitiveElement      `json:"_instantiatesUri,omitempty"`
	BasedOn                               []Reference              `json:"basedOn,omitempty"`
	PriorRequest                          []Reference              `json:"priorRequest,omitempty"`
	GroupIdentifier                       *Identifier              `json:"groupIdentifier,omitempty"`
	Status                                *Code                    `json:"status,omitempty"`
	StatusPrimitiveElement                *primitiveElement        `json:"_status,omitempty"`
	Intent                                Code                     `json:"intent,omitempty"`
	IntentPrimitiveElement                *primitiveElement        `json:"_intent,omitempty"`
	Priority                              *Code                    `json:"priority,omitempty"`
	PriorityPrimitiveElement              *primitiveElement        `json:"_priority,omitempty"`
	CodeReference                         *Reference               `json:"codeReference,omitempty"`
	CodeCodeableConcept                   *CodeableConcept         `json:"codeCodeableConcept,omitempty"`
	Parameter                             []DeviceRequestParameter `json:"parameter,omitempty"`
	Subject                               Reference                `json:"subject,omitempty"`
	Encounter                             *Reference               `json:"encounter,omitempty"`
	OccurrenceDateTime                    *DateTime                `json:"occurrenceDateTime,omitempty"`
	OccurrenceDateTimePrimitiveElement    *primitiveElement        `json:"_occurrenceDateTime,omitempty"`
	OccurrencePeriod                      *Period                  `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming                      *Timing                  `json:"occurrenceTiming,omitempty"`
	AuthoredOn                            *DateTime                `json:"authoredOn,omitempty"`
	AuthoredOnPrimitiveElement            *primitiveElement        `json:"_authoredOn,omitempty"`
	Requester                             *Reference               `json:"requester,omitempty"`
	PerformerType                         *CodeableConcept         `json:"performerType,omitempty"`
	Performer                             *Reference               `json:"performer,omitempty"`
	ReasonCode                            []CodeableConcept        `json:"reasonCode,omitempty"`
	ReasonReference                       []Reference              `json:"reasonReference,omitempty"`
	Insurance                             []Reference              `json:"insurance,omitempty"`
	SupportingInfo                        []Reference              `json:"supportingInfo,omitempty"`
	Note                                  []Annotation             `json:"note,omitempty"`
	RelevantHistory                       []Reference              `json:"relevantHistory,omitempty"`
}

func (r DeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceRequest) marshalJSON() jsonDeviceRequest {
	m := jsonDeviceRequest{}
	m.ResourceType = "DeviceRequest"
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
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
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
	m.PriorRequest = r.PriorRequest
	m.GroupIdentifier = r.GroupIdentifier
	m.Status = r.Status
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Intent = r.Intent
	if r.Intent.Id != nil || r.Intent.Extension != nil {
		m.IntentPrimitiveElement = &primitiveElement{Id: r.Intent.Id, Extension: r.Intent.Extension}
	}
	m.Priority = r.Priority
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		m.PriorityPrimitiveElement = &primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
	}
	switch v := r.Code.(type) {
	case Reference:
		m.CodeReference = &v
	case *Reference:
		m.CodeReference = v
	case CodeableConcept:
		m.CodeCodeableConcept = &v
	case *CodeableConcept:
		m.CodeCodeableConcept = v
	}
	m.Parameter = r.Parameter
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	switch v := r.Occurrence.(type) {
	case DateTime:
		m.OccurrenceDateTime = &v
		if v.Id != nil || v.Extension != nil {
			m.OccurrenceDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		m.OccurrenceDateTime = v
		if v.Id != nil || v.Extension != nil {
			m.OccurrenceDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.OccurrencePeriod = &v
	case *Period:
		m.OccurrencePeriod = v
	case Timing:
		m.OccurrenceTiming = &v
	case *Timing:
		m.OccurrenceTiming = v
	}
	m.AuthoredOn = r.AuthoredOn
	if r.AuthoredOn != nil && (r.AuthoredOn.Id != nil || r.AuthoredOn.Extension != nil) {
		m.AuthoredOnPrimitiveElement = &primitiveElement{Id: r.AuthoredOn.Id, Extension: r.AuthoredOn.Extension}
	}
	m.Requester = r.Requester
	m.PerformerType = r.PerformerType
	m.Performer = r.Performer
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Insurance = r.Insurance
	m.SupportingInfo = r.SupportingInfo
	m.Note = r.Note
	m.RelevantHistory = r.RelevantHistory
	return m
}
func (r *DeviceRequest) UnmarshalJSON(b []byte) error {
	var m jsonDeviceRequest
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceRequest) unmarshalJSON(m jsonDeviceRequest) error {
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
		r.Contained = append(r.Contained, v.Resource)
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
	r.PriorRequest = m.PriorRequest
	r.GroupIdentifier = m.GroupIdentifier
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Intent = m.Intent
	if m.IntentPrimitiveElement != nil {
		r.Intent.Id = m.IntentPrimitiveElement.Id
		r.Intent.Extension = m.IntentPrimitiveElement.Extension
	}
	r.Priority = m.Priority
	if m.PriorityPrimitiveElement != nil {
		r.Priority.Id = m.PriorityPrimitiveElement.Id
		r.Priority.Extension = m.PriorityPrimitiveElement.Extension
	}
	if m.CodeReference != nil {
		if r.Code != nil {
			return fmt.Errorf("multiple values for field \"Code\"")
		}
		v := m.CodeReference
		r.Code = v
	}
	if m.CodeCodeableConcept != nil {
		if r.Code != nil {
			return fmt.Errorf("multiple values for field \"Code\"")
		}
		v := m.CodeCodeableConcept
		r.Code = v
	}
	r.Parameter = m.Parameter
	r.Subject = m.Subject
	r.Encounter = m.Encounter
	if m.OccurrenceDateTime != nil || m.OccurrenceDateTimePrimitiveElement != nil {
		if r.Occurrence != nil {
			return fmt.Errorf("multiple values for field \"Occurrence\"")
		}
		v := m.OccurrenceDateTime
		if m.OccurrenceDateTimePrimitiveElement != nil {
			if v == nil {
				v = &DateTime{}
			}
			v.Id = m.OccurrenceDateTimePrimitiveElement.Id
			v.Extension = m.OccurrenceDateTimePrimitiveElement.Extension
		}
		r.Occurrence = v
	}
	if m.OccurrencePeriod != nil {
		if r.Occurrence != nil {
			return fmt.Errorf("multiple values for field \"Occurrence\"")
		}
		v := m.OccurrencePeriod
		r.Occurrence = v
	}
	if m.OccurrenceTiming != nil {
		if r.Occurrence != nil {
			return fmt.Errorf("multiple values for field \"Occurrence\"")
		}
		v := m.OccurrenceTiming
		r.Occurrence = v
	}
	r.AuthoredOn = m.AuthoredOn
	if m.AuthoredOnPrimitiveElement != nil {
		r.AuthoredOn.Id = m.AuthoredOnPrimitiveElement.Id
		r.AuthoredOn.Extension = m.AuthoredOnPrimitiveElement.Extension
	}
	r.Requester = m.Requester
	r.PerformerType = m.PerformerType
	r.Performer = m.Performer
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Insurance = m.Insurance
	r.SupportingInfo = m.SupportingInfo
	r.Note = m.Note
	r.RelevantHistory = m.RelevantHistory
	return nil
}
func (r DeviceRequest) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Specific parameters for the ordered item.  For example, the prism value for lenses.
type DeviceRequestParameter struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code or string that identifies the device detail being asserted.
	Code *CodeableConcept
	// The value of the device detail.
	Value isDeviceRequestParameterValue
}
type isDeviceRequestParameterValue interface {
	isDeviceRequestParameterValue()
}

func (r CodeableConcept) isDeviceRequestParameterValue() {}
func (r Quantity) isDeviceRequestParameterValue()        {}
func (r Range) isDeviceRequestParameterValue()           {}
func (r Boolean) isDeviceRequestParameterValue()         {}

type jsonDeviceRequestParameter struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Code                         *CodeableConcept  `json:"code,omitempty"`
	ValueCodeableConcept         *CodeableConcept  `json:"valueCodeableConcept,omitempty"`
	ValueQuantity                *Quantity         `json:"valueQuantity,omitempty"`
	ValueRange                   *Range            `json:"valueRange,omitempty"`
	ValueBoolean                 *Boolean          `json:"valueBoolean,omitempty"`
	ValueBooleanPrimitiveElement *primitiveElement `json:"_valueBoolean,omitempty"`
}

func (r DeviceRequestParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r DeviceRequestParameter) marshalJSON() jsonDeviceRequestParameter {
	m := jsonDeviceRequestParameter{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	switch v := r.Value.(type) {
	case CodeableConcept:
		m.ValueCodeableConcept = &v
	case *CodeableConcept:
		m.ValueCodeableConcept = v
	case Quantity:
		m.ValueQuantity = &v
	case *Quantity:
		m.ValueQuantity = v
	case Range:
		m.ValueRange = &v
	case *Range:
		m.ValueRange = v
	case Boolean:
		m.ValueBoolean = &v
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		m.ValueBoolean = v
		if v.Id != nil || v.Extension != nil {
			m.ValueBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	return m
}
func (r *DeviceRequestParameter) UnmarshalJSON(b []byte) error {
	var m jsonDeviceRequestParameter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *DeviceRequestParameter) unmarshalJSON(m jsonDeviceRequestParameter) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.ValueCodeableConcept != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueCodeableConcept
		r.Value = v
	}
	if m.ValueQuantity != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueQuantity
		r.Value = v
	}
	if m.ValueRange != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueRange
		r.Value = v
	}
	if m.ValueBoolean != nil || m.ValueBooleanPrimitiveElement != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueBoolean
		if m.ValueBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.ValueBooleanPrimitiveElement.Id
			v.Extension = m.ValueBooleanPrimitiveElement.Extension
		}
		r.Value = v
	}
	return nil
}
func (r DeviceRequestParameter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
