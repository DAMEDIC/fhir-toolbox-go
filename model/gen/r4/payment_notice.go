package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// This resource provides the status of the payment for goods and services rendered, and the request and response resource references.
type PaymentNotice struct {
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
	// A unique identifier assigned to this payment notice.
	Identifier []Identifier
	// The status of the resource instance.
	Status Code
	// Reference of resource for which payment is being made.
	Request *Reference
	// Reference of response to resource for which payment is being made.
	Response *Reference
	// The date when this resource was created.
	Created DateTime
	// The practitioner who is responsible for the services rendered to the patient.
	Provider *Reference
	// A reference to the payment which is the subject of this notice.
	Payment Reference
	// The date when the above payment action occurred.
	PaymentDate *Date
	// The party who will receive or has received payment that is the subject of this notification.
	Payee *Reference
	// The party who is notified of the payment status.
	Recipient Reference
	// The amount sent to the payee.
	Amount Money
	// A code indicating whether payment has been sent or cleared.
	PaymentStatus *CodeableConcept
}

func (r PaymentNotice) ResourceType() string {
	return "PaymentNotice"
}
func (r PaymentNotice) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonPaymentNotice struct {
	ResourceType                  string              `json:"resourceType"`
	Id                            *Id                 `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement   `json:"_id,omitempty"`
	Meta                          *Meta               `json:"meta,omitempty"`
	ImplicitRules                 *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                      *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement   `json:"_language,omitempty"`
	Text                          *Narrative          `json:"text,omitempty"`
	Contained                     []ContainedResource `json:"contained,omitempty"`
	Extension                     []Extension         `json:"extension,omitempty"`
	ModifierExtension             []Extension         `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier        `json:"identifier,omitempty"`
	Status                        Code                `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement   `json:"_status,omitempty"`
	Request                       *Reference          `json:"request,omitempty"`
	Response                      *Reference          `json:"response,omitempty"`
	Created                       DateTime            `json:"created,omitempty"`
	CreatedPrimitiveElement       *primitiveElement   `json:"_created,omitempty"`
	Provider                      *Reference          `json:"provider,omitempty"`
	Payment                       Reference           `json:"payment,omitempty"`
	PaymentDate                   *Date               `json:"paymentDate,omitempty"`
	PaymentDatePrimitiveElement   *primitiveElement   `json:"_paymentDate,omitempty"`
	Payee                         *Reference          `json:"payee,omitempty"`
	Recipient                     Reference           `json:"recipient,omitempty"`
	Amount                        Money               `json:"amount,omitempty"`
	PaymentStatus                 *CodeableConcept    `json:"paymentStatus,omitempty"`
}

func (r PaymentNotice) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PaymentNotice) marshalJSON() jsonPaymentNotice {
	m := jsonPaymentNotice{}
	m.ResourceType = "PaymentNotice"
	if r.Id != nil && r.Id.Value != nil {
		m.Id = r.Id
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		m.ImplicitRules = r.ImplicitRules
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	if r.Language != nil && r.Language.Value != nil {
		m.Language = r.Language
	}
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
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Request = r.Request
	m.Response = r.Response
	if r.Created.Value != nil {
		m.Created = r.Created
	}
	if r.Created.Id != nil || r.Created.Extension != nil {
		m.CreatedPrimitiveElement = &primitiveElement{Id: r.Created.Id, Extension: r.Created.Extension}
	}
	m.Provider = r.Provider
	m.Payment = r.Payment
	if r.PaymentDate != nil && r.PaymentDate.Value != nil {
		m.PaymentDate = r.PaymentDate
	}
	if r.PaymentDate != nil && (r.PaymentDate.Id != nil || r.PaymentDate.Extension != nil) {
		m.PaymentDatePrimitiveElement = &primitiveElement{Id: r.PaymentDate.Id, Extension: r.PaymentDate.Extension}
	}
	m.Payee = r.Payee
	m.Recipient = r.Recipient
	m.Amount = r.Amount
	m.PaymentStatus = r.PaymentStatus
	return m
}
func (r *PaymentNotice) UnmarshalJSON(b []byte) error {
	var m jsonPaymentNotice
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PaymentNotice) unmarshalJSON(m jsonPaymentNotice) error {
	r.Id = m.Id
	if m.IdPrimitiveElement != nil {
		if r.Id == nil {
			r.Id = &Id{}
		}
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		if r.ImplicitRules == nil {
			r.ImplicitRules = &Uri{}
		}
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		if r.Language == nil {
			r.Language = &Code{}
		}
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
	r.Request = m.Request
	r.Response = m.Response
	r.Created = m.Created
	if m.CreatedPrimitiveElement != nil {
		r.Created.Id = m.CreatedPrimitiveElement.Id
		r.Created.Extension = m.CreatedPrimitiveElement.Extension
	}
	r.Provider = m.Provider
	r.Payment = m.Payment
	r.PaymentDate = m.PaymentDate
	if m.PaymentDatePrimitiveElement != nil {
		if r.PaymentDate == nil {
			r.PaymentDate = &Date{}
		}
		r.PaymentDate.Id = m.PaymentDatePrimitiveElement.Id
		r.PaymentDate.Extension = m.PaymentDatePrimitiveElement.Extension
	}
	r.Payee = m.Payee
	r.Recipient = m.Recipient
	r.Amount = m.Amount
	r.PaymentStatus = m.PaymentStatus
	return nil
}
func (r PaymentNotice) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
