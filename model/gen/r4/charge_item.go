package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// The resource ChargeItem describes the provision of healthcare provider products for a certain patient, therefore referring not only to the product, but containing in addition details of the provision, like date, time, amounts and participating organizations and persons. Main Usage of the ChargeItem is to enable the billing process and internal cost allocation.
type ChargeItem struct {
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
	// Identifiers assigned to this event performer or other systems.
	Identifier []Identifier
	// References the (external) source of pricing information, rules of application for the code this ChargeItem uses.
	DefinitionUri []Uri
	// References the source of pricing information, rules of application for the code this ChargeItem uses.
	DefinitionCanonical []Canonical
	// The current state of the ChargeItem.
	Status Code
	// ChargeItems can be grouped to larger ChargeItems covering the whole set.
	PartOf []Reference
	// A code that identifies the charge, like a billing code.
	Code CodeableConcept
	// The individual or set of individuals the action is being or was performed on.
	Subject Reference
	// The encounter or episode of care that establishes the context for this event.
	Context *Reference
	// Date/time(s) or duration when the charged service was applied.
	Occurrence isChargeItemOccurrence
	// Indicates who or what performed or participated in the charged service.
	Performer []ChargeItemPerformer
	// The organization requesting the service.
	PerformingOrganization *Reference
	// The organization performing the service.
	RequestingOrganization *Reference
	// The financial cost center permits the tracking of charge attribution.
	CostCenter *Reference
	// Quantity of which the charge item has been serviced.
	Quantity *Quantity
	// The anatomical location where the related service has been applied.
	Bodysite []CodeableConcept
	// Factor overriding the factor determined by the rules associated with the code.
	FactorOverride *Decimal
	// Total price of the charge overriding the list price associated with the code.
	PriceOverride *Money
	// If the list price or the rule-based factor associated with the code is overridden, this attribute can capture a text to indicate the  reason for this action.
	OverrideReason *String
	// The device, practitioner, etc. who entered the charge item.
	Enterer *Reference
	// Date the charge item was entered.
	EnteredDate *DateTime
	// Describes why the event occurred in coded or textual form.
	Reason []CodeableConcept
	// Indicated the rendered service that caused this charge.
	Service []Reference
	// Identifies the device, food, drug or other product being charged either by type code or reference to an instance.
	Product isChargeItemProduct
	// Account into which this ChargeItems belongs.
	Account []Reference
	// Comments made about the event by the performer, subject or other participants.
	Note []Annotation
	// Further information supporting this charge.
	SupportingInformation []Reference
}

func (r ChargeItem) ResourceType() string {
	return "ChargeItem"
}

type isChargeItemOccurrence interface {
	isChargeItemOccurrence()
}

func (r DateTime) isChargeItemOccurrence() {}
func (r Period) isChargeItemOccurrence()   {}
func (r Timing) isChargeItemOccurrence()   {}

type isChargeItemProduct interface {
	isChargeItemProduct()
}

func (r Reference) isChargeItemProduct()       {}
func (r CodeableConcept) isChargeItemProduct() {}

type jsonChargeItem struct {
	ResourceType                        string                `json:"resourceType"`
	Id                                  *Id                   `json:"id,omitempty"`
	IdPrimitiveElement                  *primitiveElement     `json:"_id,omitempty"`
	Meta                                *Meta                 `json:"meta,omitempty"`
	ImplicitRules                       *Uri                  `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement       *primitiveElement     `json:"_implicitRules,omitempty"`
	Language                            *Code                 `json:"language,omitempty"`
	LanguagePrimitiveElement            *primitiveElement     `json:"_language,omitempty"`
	Text                                *Narrative            `json:"text,omitempty"`
	Contained                           []containedResource   `json:"contained,omitempty"`
	Extension                           []Extension           `json:"extension,omitempty"`
	ModifierExtension                   []Extension           `json:"modifierExtension,omitempty"`
	Identifier                          []Identifier          `json:"identifier,omitempty"`
	DefinitionUri                       []Uri                 `json:"definitionUri,omitempty"`
	DefinitionUriPrimitiveElement       []*primitiveElement   `json:"_definitionUri,omitempty"`
	DefinitionCanonical                 []Canonical           `json:"definitionCanonical,omitempty"`
	DefinitionCanonicalPrimitiveElement []*primitiveElement   `json:"_definitionCanonical,omitempty"`
	Status                              Code                  `json:"status,omitempty"`
	StatusPrimitiveElement              *primitiveElement     `json:"_status,omitempty"`
	PartOf                              []Reference           `json:"partOf,omitempty"`
	Code                                CodeableConcept       `json:"code,omitempty"`
	Subject                             Reference             `json:"subject,omitempty"`
	Context                             *Reference            `json:"context,omitempty"`
	OccurrenceDateTime                  *DateTime             `json:"occurrenceDateTime,omitempty"`
	OccurrenceDateTimePrimitiveElement  *primitiveElement     `json:"_occurrenceDateTime,omitempty"`
	OccurrencePeriod                    *Period               `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming                    *Timing               `json:"occurrenceTiming,omitempty"`
	Performer                           []ChargeItemPerformer `json:"performer,omitempty"`
	PerformingOrganization              *Reference            `json:"performingOrganization,omitempty"`
	RequestingOrganization              *Reference            `json:"requestingOrganization,omitempty"`
	CostCenter                          *Reference            `json:"costCenter,omitempty"`
	Quantity                            *Quantity             `json:"quantity,omitempty"`
	Bodysite                            []CodeableConcept     `json:"bodysite,omitempty"`
	FactorOverride                      *Decimal              `json:"factorOverride,omitempty"`
	FactorOverridePrimitiveElement      *primitiveElement     `json:"_factorOverride,omitempty"`
	PriceOverride                       *Money                `json:"priceOverride,omitempty"`
	OverrideReason                      *String               `json:"overrideReason,omitempty"`
	OverrideReasonPrimitiveElement      *primitiveElement     `json:"_overrideReason,omitempty"`
	Enterer                             *Reference            `json:"enterer,omitempty"`
	EnteredDate                         *DateTime             `json:"enteredDate,omitempty"`
	EnteredDatePrimitiveElement         *primitiveElement     `json:"_enteredDate,omitempty"`
	Reason                              []CodeableConcept     `json:"reason,omitempty"`
	Service                             []Reference           `json:"service,omitempty"`
	ProductReference                    *Reference            `json:"productReference,omitempty"`
	ProductCodeableConcept              *CodeableConcept      `json:"productCodeableConcept,omitempty"`
	Account                             []Reference           `json:"account,omitempty"`
	Note                                []Annotation          `json:"note,omitempty"`
	SupportingInformation               []Reference           `json:"supportingInformation,omitempty"`
}

func (r ChargeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ChargeItem) marshalJSON() jsonChargeItem {
	m := jsonChargeItem{}
	m.ResourceType = "ChargeItem"
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
	m.DefinitionUri = r.DefinitionUri
	anyDefinitionUriIdOrExtension := false
	for _, e := range r.DefinitionUri {
		if e.Id != nil || e.Extension != nil {
			anyDefinitionUriIdOrExtension = true
			break
		}
	}
	if anyDefinitionUriIdOrExtension {
		m.DefinitionUriPrimitiveElement = make([]*primitiveElement, 0, len(r.DefinitionUri))
		for _, e := range r.DefinitionUri {
			if e.Id != nil || e.Extension != nil {
				m.DefinitionUriPrimitiveElement = append(m.DefinitionUriPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.DefinitionUriPrimitiveElement = append(m.DefinitionUriPrimitiveElement, nil)
			}
		}
	}
	m.DefinitionCanonical = r.DefinitionCanonical
	anyDefinitionCanonicalIdOrExtension := false
	for _, e := range r.DefinitionCanonical {
		if e.Id != nil || e.Extension != nil {
			anyDefinitionCanonicalIdOrExtension = true
			break
		}
	}
	if anyDefinitionCanonicalIdOrExtension {
		m.DefinitionCanonicalPrimitiveElement = make([]*primitiveElement, 0, len(r.DefinitionCanonical))
		for _, e := range r.DefinitionCanonical {
			if e.Id != nil || e.Extension != nil {
				m.DefinitionCanonicalPrimitiveElement = append(m.DefinitionCanonicalPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.DefinitionCanonicalPrimitiveElement = append(m.DefinitionCanonicalPrimitiveElement, nil)
			}
		}
	}
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.PartOf = r.PartOf
	m.Code = r.Code
	m.Subject = r.Subject
	m.Context = r.Context
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
	m.Performer = r.Performer
	m.PerformingOrganization = r.PerformingOrganization
	m.RequestingOrganization = r.RequestingOrganization
	m.CostCenter = r.CostCenter
	m.Quantity = r.Quantity
	m.Bodysite = r.Bodysite
	m.FactorOverride = r.FactorOverride
	if r.FactorOverride != nil && (r.FactorOverride.Id != nil || r.FactorOverride.Extension != nil) {
		m.FactorOverridePrimitiveElement = &primitiveElement{Id: r.FactorOverride.Id, Extension: r.FactorOverride.Extension}
	}
	m.PriceOverride = r.PriceOverride
	m.OverrideReason = r.OverrideReason
	if r.OverrideReason != nil && (r.OverrideReason.Id != nil || r.OverrideReason.Extension != nil) {
		m.OverrideReasonPrimitiveElement = &primitiveElement{Id: r.OverrideReason.Id, Extension: r.OverrideReason.Extension}
	}
	m.Enterer = r.Enterer
	m.EnteredDate = r.EnteredDate
	if r.EnteredDate != nil && (r.EnteredDate.Id != nil || r.EnteredDate.Extension != nil) {
		m.EnteredDatePrimitiveElement = &primitiveElement{Id: r.EnteredDate.Id, Extension: r.EnteredDate.Extension}
	}
	m.Reason = r.Reason
	m.Service = r.Service
	switch v := r.Product.(type) {
	case Reference:
		m.ProductReference = &v
	case *Reference:
		m.ProductReference = v
	case CodeableConcept:
		m.ProductCodeableConcept = &v
	case *CodeableConcept:
		m.ProductCodeableConcept = v
	}
	m.Account = r.Account
	m.Note = r.Note
	m.SupportingInformation = r.SupportingInformation
	return m
}
func (r *ChargeItem) UnmarshalJSON(b []byte) error {
	var m jsonChargeItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ChargeItem) unmarshalJSON(m jsonChargeItem) error {
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
	r.DefinitionUri = m.DefinitionUri
	for i, e := range m.DefinitionUriPrimitiveElement {
		if len(r.DefinitionUri) > i {
			r.DefinitionUri[i].Id = e.Id
			r.DefinitionUri[i].Extension = e.Extension
		} else {
			r.DefinitionUri = append(r.DefinitionUri, Uri{Id: e.Id, Extension: e.Extension})
		}
	}
	r.DefinitionCanonical = m.DefinitionCanonical
	for i, e := range m.DefinitionCanonicalPrimitiveElement {
		if len(r.DefinitionCanonical) > i {
			r.DefinitionCanonical[i].Id = e.Id
			r.DefinitionCanonical[i].Extension = e.Extension
		} else {
			r.DefinitionCanonical = append(r.DefinitionCanonical, Canonical{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.PartOf = m.PartOf
	r.Code = m.Code
	r.Subject = m.Subject
	r.Context = m.Context
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
	r.Performer = m.Performer
	r.PerformingOrganization = m.PerformingOrganization
	r.RequestingOrganization = m.RequestingOrganization
	r.CostCenter = m.CostCenter
	r.Quantity = m.Quantity
	r.Bodysite = m.Bodysite
	r.FactorOverride = m.FactorOverride
	if m.FactorOverridePrimitiveElement != nil {
		r.FactorOverride.Id = m.FactorOverridePrimitiveElement.Id
		r.FactorOverride.Extension = m.FactorOverridePrimitiveElement.Extension
	}
	r.PriceOverride = m.PriceOverride
	r.OverrideReason = m.OverrideReason
	if m.OverrideReasonPrimitiveElement != nil {
		r.OverrideReason.Id = m.OverrideReasonPrimitiveElement.Id
		r.OverrideReason.Extension = m.OverrideReasonPrimitiveElement.Extension
	}
	r.Enterer = m.Enterer
	r.EnteredDate = m.EnteredDate
	if m.EnteredDatePrimitiveElement != nil {
		r.EnteredDate.Id = m.EnteredDatePrimitiveElement.Id
		r.EnteredDate.Extension = m.EnteredDatePrimitiveElement.Extension
	}
	r.Reason = m.Reason
	r.Service = m.Service
	if m.ProductReference != nil {
		if r.Product != nil {
			return fmt.Errorf("multiple values for field \"Product\"")
		}
		v := m.ProductReference
		r.Product = v
	}
	if m.ProductCodeableConcept != nil {
		if r.Product != nil {
			return fmt.Errorf("multiple values for field \"Product\"")
		}
		v := m.ProductCodeableConcept
		r.Product = v
	}
	r.Account = m.Account
	r.Note = m.Note
	r.SupportingInformation = m.SupportingInformation
	return nil
}
func (r ChargeItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates who or what performed or participated in the charged service.
type ChargeItemPerformer struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Describes the type of performance or participation(e.g. primary surgeon, anesthesiologiest, etc.).
	Function *CodeableConcept
	// The device, practitioner, etc. who performed or participated in the service.
	Actor Reference
}
type jsonChargeItemPerformer struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor,omitempty"`
}

func (r ChargeItemPerformer) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ChargeItemPerformer) marshalJSON() jsonChargeItemPerformer {
	m := jsonChargeItemPerformer{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Function = r.Function
	m.Actor = r.Actor
	return m
}
func (r *ChargeItemPerformer) UnmarshalJSON(b []byte) error {
	var m jsonChargeItemPerformer
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ChargeItemPerformer) unmarshalJSON(m jsonChargeItemPerformer) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Function = m.Function
	r.Actor = m.Actor
	return nil
}
func (r ChargeItemPerformer) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
