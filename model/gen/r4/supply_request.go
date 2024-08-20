package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A record of a request for a medication, substance or device used in the healthcare setting.
type SupplyRequest struct {
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
	// Business identifiers assigned to this SupplyRequest by the author and/or other systems. These identifiers remain constant as the resource is updated and propagates from server to server.
	Identifier []Identifier
	// Status of the supply request.
	Status *Code
	// Category of supply, e.g.  central, non-stock, etc. This is used to support work flows associated with the supply process.
	Category *CodeableConcept
	// Indicates how quickly this SupplyRequest should be addressed with respect to other requests.
	Priority *Code
	// The item that is requested to be supplied. This is either a link to a resource representing the details of the item or a code that identifies the item from a known list.
	Item isSupplyRequestItem
	// The amount that is being ordered of the indicated item.
	Quantity Quantity
	// Specific parameters for the ordered item.  For example, the size of the indicated item.
	Parameter []SupplyRequestParameter
	// When the request should be fulfilled.
	Occurrence isSupplyRequestOccurrence
	// When the request was made.
	AuthoredOn *DateTime
	// The device, practitioner, etc. who initiated the request.
	Requester *Reference
	// Who is intended to fulfill the request.
	Supplier []Reference
	// The reason why the supply item was requested.
	ReasonCode []CodeableConcept
	// The reason why the supply item was requested.
	ReasonReference []Reference
	// Where the supply is expected to come from.
	DeliverFrom *Reference
	// Where the supply is destined to go.
	DeliverTo *Reference
}

func (r SupplyRequest) ResourceType() string {
	return "SupplyRequest"
}
func (r SupplyRequest) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isSupplyRequestItem interface {
	isSupplyRequestItem()
}

func (r CodeableConcept) isSupplyRequestItem() {}
func (r Reference) isSupplyRequestItem()       {}

type isSupplyRequestOccurrence interface {
	isSupplyRequestOccurrence()
}

func (r DateTime) isSupplyRequestOccurrence() {}
func (r Period) isSupplyRequestOccurrence()   {}
func (r Timing) isSupplyRequestOccurrence()   {}

type jsonSupplyRequest struct {
	ResourceType                       string                   `json:"resourceType"`
	Id                                 *Id                      `json:"id,omitempty"`
	IdPrimitiveElement                 *primitiveElement        `json:"_id,omitempty"`
	Meta                               *Meta                    `json:"meta,omitempty"`
	ImplicitRules                      *Uri                     `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement      *primitiveElement        `json:"_implicitRules,omitempty"`
	Language                           *Code                    `json:"language,omitempty"`
	LanguagePrimitiveElement           *primitiveElement        `json:"_language,omitempty"`
	Text                               *Narrative               `json:"text,omitempty"`
	Contained                          []ContainedResource      `json:"contained,omitempty"`
	Extension                          []Extension              `json:"extension,omitempty"`
	ModifierExtension                  []Extension              `json:"modifierExtension,omitempty"`
	Identifier                         []Identifier             `json:"identifier,omitempty"`
	Status                             *Code                    `json:"status,omitempty"`
	StatusPrimitiveElement             *primitiveElement        `json:"_status,omitempty"`
	Category                           *CodeableConcept         `json:"category,omitempty"`
	Priority                           *Code                    `json:"priority,omitempty"`
	PriorityPrimitiveElement           *primitiveElement        `json:"_priority,omitempty"`
	ItemCodeableConcept                *CodeableConcept         `json:"itemCodeableConcept,omitempty"`
	ItemReference                      *Reference               `json:"itemReference,omitempty"`
	Quantity                           Quantity                 `json:"quantity,omitempty"`
	Parameter                          []SupplyRequestParameter `json:"parameter,omitempty"`
	OccurrenceDateTime                 *DateTime                `json:"occurrenceDateTime,omitempty"`
	OccurrenceDateTimePrimitiveElement *primitiveElement        `json:"_occurrenceDateTime,omitempty"`
	OccurrencePeriod                   *Period                  `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming                   *Timing                  `json:"occurrenceTiming,omitempty"`
	AuthoredOn                         *DateTime                `json:"authoredOn,omitempty"`
	AuthoredOnPrimitiveElement         *primitiveElement        `json:"_authoredOn,omitempty"`
	Requester                          *Reference               `json:"requester,omitempty"`
	Supplier                           []Reference              `json:"supplier,omitempty"`
	ReasonCode                         []CodeableConcept        `json:"reasonCode,omitempty"`
	ReasonReference                    []Reference              `json:"reasonReference,omitempty"`
	DeliverFrom                        *Reference               `json:"deliverFrom,omitempty"`
	DeliverTo                          *Reference               `json:"deliverTo,omitempty"`
}

func (r SupplyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SupplyRequest) marshalJSON() jsonSupplyRequest {
	m := jsonSupplyRequest{}
	m.ResourceType = "SupplyRequest"
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
	m.Status = r.Status
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Category = r.Category
	m.Priority = r.Priority
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		m.PriorityPrimitiveElement = &primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
	}
	switch v := r.Item.(type) {
	case CodeableConcept:
		m.ItemCodeableConcept = &v
	case *CodeableConcept:
		m.ItemCodeableConcept = v
	case Reference:
		m.ItemReference = &v
	case *Reference:
		m.ItemReference = v
	}
	m.Quantity = r.Quantity
	m.Parameter = r.Parameter
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
	m.Supplier = r.Supplier
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.DeliverFrom = r.DeliverFrom
	m.DeliverTo = r.DeliverTo
	return m
}
func (r *SupplyRequest) UnmarshalJSON(b []byte) error {
	var m jsonSupplyRequest
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SupplyRequest) unmarshalJSON(m jsonSupplyRequest) error {
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
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Category = m.Category
	r.Priority = m.Priority
	if m.PriorityPrimitiveElement != nil {
		r.Priority.Id = m.PriorityPrimitiveElement.Id
		r.Priority.Extension = m.PriorityPrimitiveElement.Extension
	}
	if m.ItemCodeableConcept != nil {
		if r.Item != nil {
			return fmt.Errorf("multiple values for field \"Item\"")
		}
		v := m.ItemCodeableConcept
		r.Item = v
	}
	if m.ItemReference != nil {
		if r.Item != nil {
			return fmt.Errorf("multiple values for field \"Item\"")
		}
		v := m.ItemReference
		r.Item = v
	}
	r.Quantity = m.Quantity
	r.Parameter = m.Parameter
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
	r.Supplier = m.Supplier
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.DeliverFrom = m.DeliverFrom
	r.DeliverTo = m.DeliverTo
	return nil
}
func (r SupplyRequest) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Specific parameters for the ordered item.  For example, the size of the indicated item.
type SupplyRequestParameter struct {
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
	Value isSupplyRequestParameterValue
}
type isSupplyRequestParameterValue interface {
	isSupplyRequestParameterValue()
}

func (r CodeableConcept) isSupplyRequestParameterValue() {}
func (r Quantity) isSupplyRequestParameterValue()        {}
func (r Range) isSupplyRequestParameterValue()           {}
func (r Boolean) isSupplyRequestParameterValue()         {}

type jsonSupplyRequestParameter struct {
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

func (r SupplyRequestParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SupplyRequestParameter) marshalJSON() jsonSupplyRequestParameter {
	m := jsonSupplyRequestParameter{}
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
func (r *SupplyRequestParameter) UnmarshalJSON(b []byte) error {
	var m jsonSupplyRequestParameter
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SupplyRequestParameter) unmarshalJSON(m jsonSupplyRequestParameter) error {
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
func (r SupplyRequestParameter) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
