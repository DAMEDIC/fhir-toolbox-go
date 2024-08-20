package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// Record of delivery of what is supplied.
type SupplyDelivery struct {
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
	// Identifier for the supply delivery event that is used to identify it across multiple disparate systems.
	Identifier []Identifier
	// A plan, proposal or order that is fulfilled in whole or in part by this event.
	BasedOn []Reference
	// A larger event of which this particular event is a component or step.
	PartOf []Reference
	// A code specifying the state of the dispense event.
	Status *Code
	// A link to a resource representing the person whom the delivered item is for.
	Patient *Reference
	// Indicates the type of dispensing event that is performed. Examples include: Trial Fill, Completion of Trial, Partial Fill, Emergency Fill, Samples, etc.
	Type *CodeableConcept
	// The item that is being delivered or has been supplied.
	SuppliedItem *SupplyDeliverySuppliedItem
	// The date or time(s) the activity occurred.
	Occurrence isSupplyDeliveryOccurrence
	// The individual responsible for dispensing the medication, supplier or device.
	Supplier *Reference
	// Identification of the facility/location where the Supply was shipped to, as part of the dispense event.
	Destination *Reference
	// Identifies the person who picked up the Supply.
	Receiver []Reference
}

func (r SupplyDelivery) ResourceType() string {
	return "SupplyDelivery"
}
func (r SupplyDelivery) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isSupplyDeliveryOccurrence interface {
	isSupplyDeliveryOccurrence()
}

func (r DateTime) isSupplyDeliveryOccurrence() {}
func (r Period) isSupplyDeliveryOccurrence()   {}
func (r Timing) isSupplyDeliveryOccurrence()   {}

type jsonSupplyDelivery struct {
	ResourceType                       string                      `json:"resourceType"`
	Id                                 *Id                         `json:"id,omitempty"`
	IdPrimitiveElement                 *primitiveElement           `json:"_id,omitempty"`
	Meta                               *Meta                       `json:"meta,omitempty"`
	ImplicitRules                      *Uri                        `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement      *primitiveElement           `json:"_implicitRules,omitempty"`
	Language                           *Code                       `json:"language,omitempty"`
	LanguagePrimitiveElement           *primitiveElement           `json:"_language,omitempty"`
	Text                               *Narrative                  `json:"text,omitempty"`
	Contained                          []ContainedResource         `json:"contained,omitempty"`
	Extension                          []Extension                 `json:"extension,omitempty"`
	ModifierExtension                  []Extension                 `json:"modifierExtension,omitempty"`
	Identifier                         []Identifier                `json:"identifier,omitempty"`
	BasedOn                            []Reference                 `json:"basedOn,omitempty"`
	PartOf                             []Reference                 `json:"partOf,omitempty"`
	Status                             *Code                       `json:"status,omitempty"`
	StatusPrimitiveElement             *primitiveElement           `json:"_status,omitempty"`
	Patient                            *Reference                  `json:"patient,omitempty"`
	Type                               *CodeableConcept            `json:"type,omitempty"`
	SuppliedItem                       *SupplyDeliverySuppliedItem `json:"suppliedItem,omitempty"`
	OccurrenceDateTime                 *DateTime                   `json:"occurrenceDateTime,omitempty"`
	OccurrenceDateTimePrimitiveElement *primitiveElement           `json:"_occurrenceDateTime,omitempty"`
	OccurrencePeriod                   *Period                     `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming                   *Timing                     `json:"occurrenceTiming,omitempty"`
	Supplier                           *Reference                  `json:"supplier,omitempty"`
	Destination                        *Reference                  `json:"destination,omitempty"`
	Receiver                           []Reference                 `json:"receiver,omitempty"`
}

func (r SupplyDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SupplyDelivery) marshalJSON() jsonSupplyDelivery {
	m := jsonSupplyDelivery{}
	m.ResourceType = "SupplyDelivery"
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
	m.BasedOn = r.BasedOn
	m.PartOf = r.PartOf
	m.Status = r.Status
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Patient = r.Patient
	m.Type = r.Type
	m.SuppliedItem = r.SuppliedItem
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
	m.Supplier = r.Supplier
	m.Destination = r.Destination
	m.Receiver = r.Receiver
	return m
}
func (r *SupplyDelivery) UnmarshalJSON(b []byte) error {
	var m jsonSupplyDelivery
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SupplyDelivery) unmarshalJSON(m jsonSupplyDelivery) error {
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
	r.BasedOn = m.BasedOn
	r.PartOf = m.PartOf
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Patient = m.Patient
	r.Type = m.Type
	r.SuppliedItem = m.SuppliedItem
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
	r.Supplier = m.Supplier
	r.Destination = m.Destination
	r.Receiver = m.Receiver
	return nil
}
func (r SupplyDelivery) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The item that is being delivered or has been supplied.
type SupplyDeliverySuppliedItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The amount of supply that has been dispensed. Includes unit of measure.
	Quantity *Quantity
	// Identifies the medication, substance or device being dispensed. This is either a link to a resource representing the details of the item or a code that identifies the item from a known list.
	Item isSupplyDeliverySuppliedItemItem
}
type isSupplyDeliverySuppliedItemItem interface {
	isSupplyDeliverySuppliedItemItem()
}

func (r CodeableConcept) isSupplyDeliverySuppliedItemItem() {}
func (r Reference) isSupplyDeliverySuppliedItemItem()       {}

type jsonSupplyDeliverySuppliedItem struct {
	Id                  *string          `json:"id,omitempty"`
	Extension           []Extension      `json:"extension,omitempty"`
	ModifierExtension   []Extension      `json:"modifierExtension,omitempty"`
	Quantity            *Quantity        `json:"quantity,omitempty"`
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `json:"itemReference,omitempty"`
}

func (r SupplyDeliverySuppliedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SupplyDeliverySuppliedItem) marshalJSON() jsonSupplyDeliverySuppliedItem {
	m := jsonSupplyDeliverySuppliedItem{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Quantity = r.Quantity
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
	return m
}
func (r *SupplyDeliverySuppliedItem) UnmarshalJSON(b []byte) error {
	var m jsonSupplyDeliverySuppliedItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SupplyDeliverySuppliedItem) unmarshalJSON(m jsonSupplyDeliverySuppliedItem) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Quantity = m.Quantity
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
	return nil
}
func (r SupplyDeliverySuppliedItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
