package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// Invoice containing collected ChargeItems from an Account with calculated individual and total price for Billing purpose.
type Invoice struct {
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
	// Identifier of this Invoice, often used for reference in correspondence about this invoice or for tracking of payments.
	Identifier []Identifier
	// The current state of the Invoice.
	Status Code
	// In case of Invoice cancellation a reason must be given (entered in error, superseded by corrected invoice etc.).
	CancelledReason *String
	// Type of Invoice depending on domain, realm an usage (e.g. internal/external, dental, preliminary).
	Type *CodeableConcept
	// The individual or set of individuals receiving the goods and services billed in this invoice.
	Subject *Reference
	// The individual or Organization responsible for balancing of this invoice.
	Recipient *Reference
	// Date/time(s) of when this Invoice was posted.
	Date *DateTime
	// Indicates who or what performed or participated in the charged service.
	Participant []InvoiceParticipant
	// The organizationissuing the Invoice.
	Issuer *Reference
	// Account which is supposed to be balanced with this Invoice.
	Account *Reference
	// Each line item represents one charge for goods and services rendered. Details such as date, code and amount are found in the referenced ChargeItem resource.
	LineItem []InvoiceLineItem
	// The total amount for the Invoice may be calculated as the sum of the line items with surcharges/deductions that apply in certain conditions.  The priceComponent element can be used to offer transparency to the recipient of the Invoice of how the total price was calculated.
	TotalPriceComponent []InvoiceLineItemPriceComponent
	// Invoice total , taxes excluded.
	TotalNet *Money
	// Invoice total, tax included.
	TotalGross *Money
	// Payment details such as banking details, period of payment, deductibles, methods of payment.
	PaymentTerms *Markdown
	// Comments made about the invoice by the issuer, subject, or other participants.
	Note []Annotation
}
type jsonInvoice struct {
	ResourceType                    string                          `json:"resourceType"`
	Id                              *Id                             `json:"id,omitempty"`
	IdPrimitiveElement              *primitiveElement               `json:"_id,omitempty"`
	Meta                            *Meta                           `json:"meta,omitempty"`
	ImplicitRules                   *Uri                            `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement   *primitiveElement               `json:"_implicitRules,omitempty"`
	Language                        *Code                           `json:"language,omitempty"`
	LanguagePrimitiveElement        *primitiveElement               `json:"_language,omitempty"`
	Text                            *Narrative                      `json:"text,omitempty"`
	Contained                       []containedResource             `json:"contained,omitempty"`
	Extension                       []Extension                     `json:"extension,omitempty"`
	ModifierExtension               []Extension                     `json:"modifierExtension,omitempty"`
	Identifier                      []Identifier                    `json:"identifier,omitempty"`
	Status                          Code                            `json:"status,omitempty"`
	StatusPrimitiveElement          *primitiveElement               `json:"_status,omitempty"`
	CancelledReason                 *String                         `json:"cancelledReason,omitempty"`
	CancelledReasonPrimitiveElement *primitiveElement               `json:"_cancelledReason,omitempty"`
	Type                            *CodeableConcept                `json:"type,omitempty"`
	Subject                         *Reference                      `json:"subject,omitempty"`
	Recipient                       *Reference                      `json:"recipient,omitempty"`
	Date                            *DateTime                       `json:"date,omitempty"`
	DatePrimitiveElement            *primitiveElement               `json:"_date,omitempty"`
	Participant                     []InvoiceParticipant            `json:"participant,omitempty"`
	Issuer                          *Reference                      `json:"issuer,omitempty"`
	Account                         *Reference                      `json:"account,omitempty"`
	LineItem                        []InvoiceLineItem               `json:"lineItem,omitempty"`
	TotalPriceComponent             []InvoiceLineItemPriceComponent `json:"totalPriceComponent,omitempty"`
	TotalNet                        *Money                          `json:"totalNet,omitempty"`
	TotalGross                      *Money                          `json:"totalGross,omitempty"`
	PaymentTerms                    *Markdown                       `json:"paymentTerms,omitempty"`
	PaymentTermsPrimitiveElement    *primitiveElement               `json:"_paymentTerms,omitempty"`
	Note                            []Annotation                    `json:"note,omitempty"`
}

func (r Invoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Invoice) marshalJSON() jsonInvoice {
	m := jsonInvoice{}
	m.ResourceType = "Invoice"
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
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.CancelledReason = r.CancelledReason
	if r.CancelledReason != nil && (r.CancelledReason.Id != nil || r.CancelledReason.Extension != nil) {
		m.CancelledReasonPrimitiveElement = &primitiveElement{Id: r.CancelledReason.Id, Extension: r.CancelledReason.Extension}
	}
	m.Type = r.Type
	m.Subject = r.Subject
	m.Recipient = r.Recipient
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Participant = r.Participant
	m.Issuer = r.Issuer
	m.Account = r.Account
	m.LineItem = r.LineItem
	m.TotalPriceComponent = r.TotalPriceComponent
	m.TotalNet = r.TotalNet
	m.TotalGross = r.TotalGross
	m.PaymentTerms = r.PaymentTerms
	if r.PaymentTerms != nil && (r.PaymentTerms.Id != nil || r.PaymentTerms.Extension != nil) {
		m.PaymentTermsPrimitiveElement = &primitiveElement{Id: r.PaymentTerms.Id, Extension: r.PaymentTerms.Extension}
	}
	m.Note = r.Note
	return m
}
func (r *Invoice) UnmarshalJSON(b []byte) error {
	var m jsonInvoice
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Invoice) unmarshalJSON(m jsonInvoice) error {
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
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.CancelledReason = m.CancelledReason
	if m.CancelledReasonPrimitiveElement != nil {
		r.CancelledReason.Id = m.CancelledReasonPrimitiveElement.Id
		r.CancelledReason.Extension = m.CancelledReasonPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Subject = m.Subject
	r.Recipient = m.Recipient
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Participant = m.Participant
	r.Issuer = m.Issuer
	r.Account = m.Account
	r.LineItem = m.LineItem
	r.TotalPriceComponent = m.TotalPriceComponent
	r.TotalNet = m.TotalNet
	r.TotalGross = m.TotalGross
	r.PaymentTerms = m.PaymentTerms
	if m.PaymentTermsPrimitiveElement != nil {
		r.PaymentTerms.Id = m.PaymentTermsPrimitiveElement.Id
		r.PaymentTerms.Extension = m.PaymentTermsPrimitiveElement.Extension
	}
	r.Note = m.Note
	return nil
}
func (r Invoice) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates who or what performed or participated in the charged service.
type InvoiceParticipant struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Describes the type of involvement (e.g. transcriptionist, creator etc.). If the invoice has been created automatically, the Participant may be a billing engine or another kind of device.
	Role *CodeableConcept
	// The device, practitioner, etc. who performed or participated in the service.
	Actor Reference
}
type jsonInvoiceParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Actor             Reference        `json:"actor,omitempty"`
}

func (r InvoiceParticipant) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r InvoiceParticipant) marshalJSON() jsonInvoiceParticipant {
	m := jsonInvoiceParticipant{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Role = r.Role
	m.Actor = r.Actor
	return m
}
func (r *InvoiceParticipant) UnmarshalJSON(b []byte) error {
	var m jsonInvoiceParticipant
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *InvoiceParticipant) unmarshalJSON(m jsonInvoiceParticipant) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Role = m.Role
	r.Actor = m.Actor
	return nil
}
func (r InvoiceParticipant) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Each line item represents one charge for goods and services rendered. Details such as date, code and amount are found in the referenced ChargeItem resource.
type InvoiceLineItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Sequence in which the items appear on the invoice.
	Sequence *PositiveInt
	// The ChargeItem contains information such as the billing code, date, amount etc. If no further details are required for the lineItem, inline billing codes can be added using the CodeableConcept data type instead of the Reference.
	ChargeItem isInvoiceLineItemChargeItem
	// The price for a ChargeItem may be calculated as a base price with surcharges/deductions that apply in certain conditions. A ChargeItemDefinition resource that defines the prices, factors and conditions that apply to a billing code is currently under development. The priceComponent element can be used to offer transparency to the recipient of the Invoice as to how the prices have been calculated.
	PriceComponent []InvoiceLineItemPriceComponent
}
type isInvoiceLineItemChargeItem interface {
	isInvoiceLineItemChargeItem()
}

func (r Reference) isInvoiceLineItemChargeItem()       {}
func (r CodeableConcept) isInvoiceLineItemChargeItem() {}

type jsonInvoiceLineItem struct {
	Id                        *string                         `json:"id,omitempty"`
	Extension                 []Extension                     `json:"extension,omitempty"`
	ModifierExtension         []Extension                     `json:"modifierExtension,omitempty"`
	Sequence                  *PositiveInt                    `json:"sequence,omitempty"`
	SequencePrimitiveElement  *primitiveElement               `json:"_sequence,omitempty"`
	ChargeItemReference       *Reference                      `json:"chargeItemReference,omitempty"`
	ChargeItemCodeableConcept *CodeableConcept                `json:"chargeItemCodeableConcept,omitempty"`
	PriceComponent            []InvoiceLineItemPriceComponent `json:"priceComponent,omitempty"`
}

func (r InvoiceLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r InvoiceLineItem) marshalJSON() jsonInvoiceLineItem {
	m := jsonInvoiceLineItem{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Sequence = r.Sequence
	if r.Sequence != nil && (r.Sequence.Id != nil || r.Sequence.Extension != nil) {
		m.SequencePrimitiveElement = &primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
	}
	switch v := r.ChargeItem.(type) {
	case Reference:
		m.ChargeItemReference = &v
	case *Reference:
		m.ChargeItemReference = v
	case CodeableConcept:
		m.ChargeItemCodeableConcept = &v
	case *CodeableConcept:
		m.ChargeItemCodeableConcept = v
	}
	m.PriceComponent = r.PriceComponent
	return m
}
func (r *InvoiceLineItem) UnmarshalJSON(b []byte) error {
	var m jsonInvoiceLineItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *InvoiceLineItem) unmarshalJSON(m jsonInvoiceLineItem) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Sequence = m.Sequence
	if m.SequencePrimitiveElement != nil {
		r.Sequence.Id = m.SequencePrimitiveElement.Id
		r.Sequence.Extension = m.SequencePrimitiveElement.Extension
	}
	if m.ChargeItemReference != nil {
		if r.ChargeItem != nil {
			return fmt.Errorf("multiple values for field \"ChargeItem\"")
		}
		v := m.ChargeItemReference
		r.ChargeItem = v
	}
	if m.ChargeItemCodeableConcept != nil {
		if r.ChargeItem != nil {
			return fmt.Errorf("multiple values for field \"ChargeItem\"")
		}
		v := m.ChargeItemCodeableConcept
		r.ChargeItem = v
	}
	r.PriceComponent = m.PriceComponent
	return nil
}
func (r InvoiceLineItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The price for a ChargeItem may be calculated as a base price with surcharges/deductions that apply in certain conditions. A ChargeItemDefinition resource that defines the prices, factors and conditions that apply to a billing code is currently under development. The priceComponent element can be used to offer transparency to the recipient of the Invoice as to how the prices have been calculated.
type InvoiceLineItemPriceComponent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// This code identifies the type of the component.
	Type Code
	// A code that identifies the component. Codes may be used to differentiate between kinds of taxes, surcharges, discounts etc.
	Code *CodeableConcept
	// The factor that has been applied on the base price for calculating this component.
	Factor *Decimal
	// The amount calculated for this component.
	Amount *Money
}
type jsonInvoiceLineItemPriceComponent struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Type                   Code              `json:"type,omitempty"`
	TypePrimitiveElement   *primitiveElement `json:"_type,omitempty"`
	Code                   *CodeableConcept  `json:"code,omitempty"`
	Factor                 *Decimal          `json:"factor,omitempty"`
	FactorPrimitiveElement *primitiveElement `json:"_factor,omitempty"`
	Amount                 *Money            `json:"amount,omitempty"`
}

func (r InvoiceLineItemPriceComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r InvoiceLineItemPriceComponent) marshalJSON() jsonInvoiceLineItemPriceComponent {
	m := jsonInvoiceLineItemPriceComponent{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Code = r.Code
	m.Factor = r.Factor
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		m.FactorPrimitiveElement = &primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
	}
	m.Amount = r.Amount
	return m
}
func (r *InvoiceLineItemPriceComponent) UnmarshalJSON(b []byte) error {
	var m jsonInvoiceLineItemPriceComponent
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *InvoiceLineItemPriceComponent) unmarshalJSON(m jsonInvoiceLineItemPriceComponent) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Code = m.Code
	r.Factor = m.Factor
	if m.FactorPrimitiveElement != nil {
		r.Factor.Id = m.FactorPrimitiveElement.Id
		r.Factor.Extension = m.FactorPrimitiveElement.Extension
	}
	r.Amount = m.Amount
	return nil
}
func (r InvoiceLineItemPriceComponent) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
