package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// Represents a defined collection of entities that may be discussed or acted upon collectively but which are not expected to act collectively, and are not formally or legally recognized; i.e. a collection of entities that isn't an Organization.
type Group struct {
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
	// A unique business identifier for this group.
	Identifier []Identifier
	// Indicates whether the record for the group is available for use or is merely being retained for historical purposes.
	Active *Boolean
	// Identifies the broad classification of the kind of resources the group includes.
	Type Code
	// If true, indicates that the resource refers to a specific group of real individuals.  If false, the group defines a set of intended individuals.
	Actual Boolean
	// Provides a specific type of resource the group includes; e.g. "cow", "syringe", etc.
	Code *CodeableConcept
	// A label assigned to the group for human identification and communication.
	Name *String
	// A count of the number of resource instances that are part of the group.
	Quantity *UnsignedInt
	// Entity responsible for defining and maintaining Group characteristics and/or registered members.
	ManagingEntity *Reference
	// Identifies traits whose presence r absence is shared by members of the group.
	Characteristic []GroupCharacteristic
	// Identifies the resource instances that are members of the group.
	Member []GroupMember
}

func (r Group) ResourceType() string {
	return "Group"
}
func (r Group) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonGroup struct {
	ResourceType                  string                `json:"resourceType"`
	Id                            *Id                   `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement     `json:"_id,omitempty"`
	Meta                          *Meta                 `json:"meta,omitempty"`
	ImplicitRules                 *Uri                  `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement     `json:"_implicitRules,omitempty"`
	Language                      *Code                 `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement     `json:"_language,omitempty"`
	Text                          *Narrative            `json:"text,omitempty"`
	Contained                     []ContainedResource   `json:"contained,omitempty"`
	Extension                     []Extension           `json:"extension,omitempty"`
	ModifierExtension             []Extension           `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier          `json:"identifier,omitempty"`
	Active                        *Boolean              `json:"active,omitempty"`
	ActivePrimitiveElement        *primitiveElement     `json:"_active,omitempty"`
	Type                          Code                  `json:"type,omitempty"`
	TypePrimitiveElement          *primitiveElement     `json:"_type,omitempty"`
	Actual                        Boolean               `json:"actual,omitempty"`
	ActualPrimitiveElement        *primitiveElement     `json:"_actual,omitempty"`
	Code                          *CodeableConcept      `json:"code,omitempty"`
	Name                          *String               `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement     `json:"_name,omitempty"`
	Quantity                      *UnsignedInt          `json:"quantity,omitempty"`
	QuantityPrimitiveElement      *primitiveElement     `json:"_quantity,omitempty"`
	ManagingEntity                *Reference            `json:"managingEntity,omitempty"`
	Characteristic                []GroupCharacteristic `json:"characteristic,omitempty"`
	Member                        []GroupMember         `json:"member,omitempty"`
}

func (r Group) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Group) marshalJSON() jsonGroup {
	m := jsonGroup{}
	m.ResourceType = "Group"
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
	m.Active = r.Active
	if r.Active != nil && (r.Active.Id != nil || r.Active.Extension != nil) {
		m.ActivePrimitiveElement = &primitiveElement{Id: r.Active.Id, Extension: r.Active.Extension}
	}
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Actual = r.Actual
	if r.Actual.Id != nil || r.Actual.Extension != nil {
		m.ActualPrimitiveElement = &primitiveElement{Id: r.Actual.Id, Extension: r.Actual.Extension}
	}
	m.Code = r.Code
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Quantity = r.Quantity
	if r.Quantity != nil && (r.Quantity.Id != nil || r.Quantity.Extension != nil) {
		m.QuantityPrimitiveElement = &primitiveElement{Id: r.Quantity.Id, Extension: r.Quantity.Extension}
	}
	m.ManagingEntity = r.ManagingEntity
	m.Characteristic = r.Characteristic
	m.Member = r.Member
	return m
}
func (r *Group) UnmarshalJSON(b []byte) error {
	var m jsonGroup
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Group) unmarshalJSON(m jsonGroup) error {
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
	r.Active = m.Active
	if m.ActivePrimitiveElement != nil {
		r.Active.Id = m.ActivePrimitiveElement.Id
		r.Active.Extension = m.ActivePrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Actual = m.Actual
	if m.ActualPrimitiveElement != nil {
		r.Actual.Id = m.ActualPrimitiveElement.Id
		r.Actual.Extension = m.ActualPrimitiveElement.Extension
	}
	r.Code = m.Code
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Quantity = m.Quantity
	if m.QuantityPrimitiveElement != nil {
		r.Quantity.Id = m.QuantityPrimitiveElement.Id
		r.Quantity.Extension = m.QuantityPrimitiveElement.Extension
	}
	r.ManagingEntity = m.ManagingEntity
	r.Characteristic = m.Characteristic
	r.Member = m.Member
	return nil
}
func (r Group) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Identifies traits whose presence r absence is shared by members of the group.
type GroupCharacteristic struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code that identifies the kind of trait being asserted.
	Code CodeableConcept
	// The value of the trait that holds (or does not hold - see 'exclude') for members of the group.
	Value isGroupCharacteristicValue
	// If true, indicates the characteristic is one that is NOT held by members of the group.
	Exclude Boolean
	// The period over which the characteristic is tested; e.g. the patient had an operation during the month of June.
	Period *Period
}
type isGroupCharacteristicValue interface {
	isGroupCharacteristicValue()
}

func (r CodeableConcept) isGroupCharacteristicValue() {}
func (r Boolean) isGroupCharacteristicValue()         {}
func (r Quantity) isGroupCharacteristicValue()        {}
func (r Range) isGroupCharacteristicValue()           {}
func (r Reference) isGroupCharacteristicValue()       {}

type jsonGroupCharacteristic struct {
	Id                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Code                         CodeableConcept   `json:"code,omitempty"`
	ValueCodeableConcept         *CodeableConcept  `json:"valueCodeableConcept,omitempty"`
	ValueBoolean                 *Boolean          `json:"valueBoolean,omitempty"`
	ValueBooleanPrimitiveElement *primitiveElement `json:"_valueBoolean,omitempty"`
	ValueQuantity                *Quantity         `json:"valueQuantity,omitempty"`
	ValueRange                   *Range            `json:"valueRange,omitempty"`
	ValueReference               *Reference        `json:"valueReference,omitempty"`
	Exclude                      Boolean           `json:"exclude,omitempty"`
	ExcludePrimitiveElement      *primitiveElement `json:"_exclude,omitempty"`
	Period                       *Period           `json:"period,omitempty"`
}

func (r GroupCharacteristic) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r GroupCharacteristic) marshalJSON() jsonGroupCharacteristic {
	m := jsonGroupCharacteristic{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	switch v := r.Value.(type) {
	case CodeableConcept:
		m.ValueCodeableConcept = &v
	case *CodeableConcept:
		m.ValueCodeableConcept = v
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
	case Quantity:
		m.ValueQuantity = &v
	case *Quantity:
		m.ValueQuantity = v
	case Range:
		m.ValueRange = &v
	case *Range:
		m.ValueRange = v
	case Reference:
		m.ValueReference = &v
	case *Reference:
		m.ValueReference = v
	}
	m.Exclude = r.Exclude
	if r.Exclude.Id != nil || r.Exclude.Extension != nil {
		m.ExcludePrimitiveElement = &primitiveElement{Id: r.Exclude.Id, Extension: r.Exclude.Extension}
	}
	m.Period = r.Period
	return m
}
func (r *GroupCharacteristic) UnmarshalJSON(b []byte) error {
	var m jsonGroupCharacteristic
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *GroupCharacteristic) unmarshalJSON(m jsonGroupCharacteristic) error {
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
	if m.ValueReference != nil {
		if r.Value != nil {
			return fmt.Errorf("multiple values for field \"Value\"")
		}
		v := m.ValueReference
		r.Value = v
	}
	r.Exclude = m.Exclude
	if m.ExcludePrimitiveElement != nil {
		r.Exclude.Id = m.ExcludePrimitiveElement.Id
		r.Exclude.Extension = m.ExcludePrimitiveElement.Extension
	}
	r.Period = m.Period
	return nil
}
func (r GroupCharacteristic) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Identifies the resource instances that are members of the group.
type GroupMember struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A reference to the entity that is a member of the group. Must be consistent with Group.type. If the entity is another group, then the type must be the same.
	Entity Reference
	// The period that the member was in the group, if known.
	Period *Period
	// A flag to indicate that the member is no longer in the group, but previously may have been a member.
	Inactive *Boolean
}
type jsonGroupMember struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Entity                   Reference         `json:"entity,omitempty"`
	Period                   *Period           `json:"period,omitempty"`
	Inactive                 *Boolean          `json:"inactive,omitempty"`
	InactivePrimitiveElement *primitiveElement `json:"_inactive,omitempty"`
}

func (r GroupMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r GroupMember) marshalJSON() jsonGroupMember {
	m := jsonGroupMember{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Entity = r.Entity
	m.Period = r.Period
	m.Inactive = r.Inactive
	if r.Inactive != nil && (r.Inactive.Id != nil || r.Inactive.Extension != nil) {
		m.InactivePrimitiveElement = &primitiveElement{Id: r.Inactive.Id, Extension: r.Inactive.Extension}
	}
	return m
}
func (r *GroupMember) UnmarshalJSON(b []byte) error {
	var m jsonGroupMember
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *GroupMember) unmarshalJSON(m jsonGroupMember) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Entity = m.Entity
	r.Period = m.Period
	r.Inactive = m.Inactive
	if m.InactivePrimitiveElement != nil {
		r.Inactive.Id = m.InactivePrimitiveElement.Id
		r.Inactive.Extension = m.InactivePrimitiveElement.Extension
	}
	return nil
}
func (r GroupMember) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
