package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// This resource provides the details including amount of a payment and allocates the payment items being paid.
type PaymentReconciliation struct {
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
	// A unique identifier assigned to this payment reconciliation.
	Identifier []Identifier
	// The status of the resource instance.
	Status Code
	// The period of time for which payments have been gathered into this bulk payment for settlement.
	Period *Period
	// The date when the resource was created.
	Created DateTime
	// The party who generated the payment.
	PaymentIssuer *Reference
	// Original request resource reference.
	Request *Reference
	// The practitioner who is responsible for the services rendered to the patient.
	Requestor *Reference
	// The outcome of a request for a reconciliation.
	Outcome *Code
	// A human readable description of the status of the request for the reconciliation.
	Disposition *String
	// The date of payment as indicated on the financial instrument.
	PaymentDate Date
	// Total payment amount as indicated on the financial instrument.
	PaymentAmount Money
	// Issuer's unique identifier for the payment instrument.
	PaymentIdentifier *Identifier
	// Distribution of the payment amount for a previously acknowledged payable.
	Detail []PaymentReconciliationDetail
	// A code for the form to be used for printing the content.
	FormCode *CodeableConcept
	// A note that describes or explains the processing in a human readable form.
	ProcessNote []PaymentReconciliationProcessNote
}

func (r PaymentReconciliation) ResourceType() string {
	return "PaymentReconciliation"
}
func (r PaymentReconciliation) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonPaymentReconciliation struct {
	ResourceType                  string                             `json:"resourceType"`
	Id                            *Id                                `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement                  `json:"_id,omitempty"`
	Meta                          *Meta                              `json:"meta,omitempty"`
	ImplicitRules                 *Uri                               `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement                  `json:"_implicitRules,omitempty"`
	Language                      *Code                              `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement                  `json:"_language,omitempty"`
	Text                          *Narrative                         `json:"text,omitempty"`
	Contained                     []ContainedResource                `json:"contained,omitempty"`
	Extension                     []Extension                        `json:"extension,omitempty"`
	ModifierExtension             []Extension                        `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier                       `json:"identifier,omitempty"`
	Status                        Code                               `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement                  `json:"_status,omitempty"`
	Period                        *Period                            `json:"period,omitempty"`
	Created                       DateTime                           `json:"created,omitempty"`
	CreatedPrimitiveElement       *primitiveElement                  `json:"_created,omitempty"`
	PaymentIssuer                 *Reference                         `json:"paymentIssuer,omitempty"`
	Request                       *Reference                         `json:"request,omitempty"`
	Requestor                     *Reference                         `json:"requestor,omitempty"`
	Outcome                       *Code                              `json:"outcome,omitempty"`
	OutcomePrimitiveElement       *primitiveElement                  `json:"_outcome,omitempty"`
	Disposition                   *String                            `json:"disposition,omitempty"`
	DispositionPrimitiveElement   *primitiveElement                  `json:"_disposition,omitempty"`
	PaymentDate                   Date                               `json:"paymentDate,omitempty"`
	PaymentDatePrimitiveElement   *primitiveElement                  `json:"_paymentDate,omitempty"`
	PaymentAmount                 Money                              `json:"paymentAmount,omitempty"`
	PaymentIdentifier             *Identifier                        `json:"paymentIdentifier,omitempty"`
	Detail                        []PaymentReconciliationDetail      `json:"detail,omitempty"`
	FormCode                      *CodeableConcept                   `json:"formCode,omitempty"`
	ProcessNote                   []PaymentReconciliationProcessNote `json:"processNote,omitempty"`
}

func (r PaymentReconciliation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PaymentReconciliation) marshalJSON() jsonPaymentReconciliation {
	m := jsonPaymentReconciliation{}
	m.ResourceType = "PaymentReconciliation"
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
	m.Period = r.Period
	if r.Created.Value != nil {
		m.Created = r.Created
	}
	if r.Created.Id != nil || r.Created.Extension != nil {
		m.CreatedPrimitiveElement = &primitiveElement{Id: r.Created.Id, Extension: r.Created.Extension}
	}
	m.PaymentIssuer = r.PaymentIssuer
	m.Request = r.Request
	m.Requestor = r.Requestor
	if r.Outcome != nil && r.Outcome.Value != nil {
		m.Outcome = r.Outcome
	}
	if r.Outcome != nil && (r.Outcome.Id != nil || r.Outcome.Extension != nil) {
		m.OutcomePrimitiveElement = &primitiveElement{Id: r.Outcome.Id, Extension: r.Outcome.Extension}
	}
	if r.Disposition != nil && r.Disposition.Value != nil {
		m.Disposition = r.Disposition
	}
	if r.Disposition != nil && (r.Disposition.Id != nil || r.Disposition.Extension != nil) {
		m.DispositionPrimitiveElement = &primitiveElement{Id: r.Disposition.Id, Extension: r.Disposition.Extension}
	}
	if r.PaymentDate.Value != nil {
		m.PaymentDate = r.PaymentDate
	}
	if r.PaymentDate.Id != nil || r.PaymentDate.Extension != nil {
		m.PaymentDatePrimitiveElement = &primitiveElement{Id: r.PaymentDate.Id, Extension: r.PaymentDate.Extension}
	}
	m.PaymentAmount = r.PaymentAmount
	m.PaymentIdentifier = r.PaymentIdentifier
	m.Detail = r.Detail
	m.FormCode = r.FormCode
	m.ProcessNote = r.ProcessNote
	return m
}
func (r *PaymentReconciliation) UnmarshalJSON(b []byte) error {
	var m jsonPaymentReconciliation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PaymentReconciliation) unmarshalJSON(m jsonPaymentReconciliation) error {
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
	r.Period = m.Period
	r.Created = m.Created
	if m.CreatedPrimitiveElement != nil {
		r.Created.Id = m.CreatedPrimitiveElement.Id
		r.Created.Extension = m.CreatedPrimitiveElement.Extension
	}
	r.PaymentIssuer = m.PaymentIssuer
	r.Request = m.Request
	r.Requestor = m.Requestor
	r.Outcome = m.Outcome
	if m.OutcomePrimitiveElement != nil {
		if r.Outcome == nil {
			r.Outcome = &Code{}
		}
		r.Outcome.Id = m.OutcomePrimitiveElement.Id
		r.Outcome.Extension = m.OutcomePrimitiveElement.Extension
	}
	r.Disposition = m.Disposition
	if m.DispositionPrimitiveElement != nil {
		if r.Disposition == nil {
			r.Disposition = &String{}
		}
		r.Disposition.Id = m.DispositionPrimitiveElement.Id
		r.Disposition.Extension = m.DispositionPrimitiveElement.Extension
	}
	r.PaymentDate = m.PaymentDate
	if m.PaymentDatePrimitiveElement != nil {
		r.PaymentDate.Id = m.PaymentDatePrimitiveElement.Id
		r.PaymentDate.Extension = m.PaymentDatePrimitiveElement.Extension
	}
	r.PaymentAmount = m.PaymentAmount
	r.PaymentIdentifier = m.PaymentIdentifier
	r.Detail = m.Detail
	r.FormCode = m.FormCode
	r.ProcessNote = m.ProcessNote
	return nil
}
func (r PaymentReconciliation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "PaymentReconciliation"
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Period, xml.StartElement{Name: xml.Name{Local: "period"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Created, xml.StartElement{Name: xml.Name{Local: "created"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PaymentIssuer, xml.StartElement{Name: xml.Name{Local: "paymentIssuer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Request, xml.StartElement{Name: xml.Name{Local: "request"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Requestor, xml.StartElement{Name: xml.Name{Local: "requestor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Outcome, xml.StartElement{Name: xml.Name{Local: "outcome"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Disposition, xml.StartElement{Name: xml.Name{Local: "disposition"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PaymentDate, xml.StartElement{Name: xml.Name{Local: "paymentDate"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PaymentAmount, xml.StartElement{Name: xml.Name{Local: "paymentAmount"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PaymentIdentifier, xml.StartElement{Name: xml.Name{Local: "paymentIdentifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Detail, xml.StartElement{Name: xml.Name{Local: "detail"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FormCode, xml.StartElement{Name: xml.Name{Local: "formCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProcessNote, xml.StartElement{Name: xml.Name{Local: "processNote"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *PaymentReconciliation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "period":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Period = &v
			case "created":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Created = v
			case "paymentIssuer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PaymentIssuer = &v
			case "request":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Request = &v
			case "requestor":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Requestor = &v
			case "outcome":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Outcome = &v
			case "disposition":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Disposition = &v
			case "paymentDate":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PaymentDate = v
			case "paymentAmount":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PaymentAmount = v
			case "paymentIdentifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PaymentIdentifier = &v
			case "detail":
				var v PaymentReconciliationDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = append(r.Detail, v)
			case "formCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FormCode = &v
			case "processNote":
				var v PaymentReconciliationProcessNote
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProcessNote = append(r.ProcessNote, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r PaymentReconciliation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Distribution of the payment amount for a previously acknowledged payable.
type PaymentReconciliationDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Unique identifier for the current payment item for the referenced payable.
	Identifier *Identifier
	// Unique identifier for the prior payment item for the referenced payable.
	Predecessor *Identifier
	// Code to indicate the nature of the payment.
	Type CodeableConcept
	// A resource, such as a Claim, the evaluation of which could lead to payment.
	Request *Reference
	// The party which submitted the claim or financial transaction.
	Submitter *Reference
	// A resource, such as a ClaimResponse, which contains a commitment to payment.
	Response *Reference
	// The date from the response resource containing a commitment to pay.
	Date *Date
	// A reference to the individual who is responsible for inquiries regarding the response and its payment.
	Responsible *Reference
	// The party which is receiving the payment.
	Payee *Reference
	// The monetary amount allocated from the total payment to the payable.
	Amount *Money
}
type jsonPaymentReconciliationDetail struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Identifier           *Identifier       `json:"identifier,omitempty"`
	Predecessor          *Identifier       `json:"predecessor,omitempty"`
	Type                 CodeableConcept   `json:"type,omitempty"`
	Request              *Reference        `json:"request,omitempty"`
	Submitter            *Reference        `json:"submitter,omitempty"`
	Response             *Reference        `json:"response,omitempty"`
	Date                 *Date             `json:"date,omitempty"`
	DatePrimitiveElement *primitiveElement `json:"_date,omitempty"`
	Responsible          *Reference        `json:"responsible,omitempty"`
	Payee                *Reference        `json:"payee,omitempty"`
	Amount               *Money            `json:"amount,omitempty"`
}

func (r PaymentReconciliationDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PaymentReconciliationDetail) marshalJSON() jsonPaymentReconciliationDetail {
	m := jsonPaymentReconciliationDetail{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Identifier = r.Identifier
	m.Predecessor = r.Predecessor
	m.Type = r.Type
	m.Request = r.Request
	m.Submitter = r.Submitter
	m.Response = r.Response
	if r.Date != nil && r.Date.Value != nil {
		m.Date = r.Date
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Responsible = r.Responsible
	m.Payee = r.Payee
	m.Amount = r.Amount
	return m
}
func (r *PaymentReconciliationDetail) UnmarshalJSON(b []byte) error {
	var m jsonPaymentReconciliationDetail
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PaymentReconciliationDetail) unmarshalJSON(m jsonPaymentReconciliationDetail) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Identifier = m.Identifier
	r.Predecessor = m.Predecessor
	r.Type = m.Type
	r.Request = m.Request
	r.Submitter = m.Submitter
	r.Response = m.Response
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		if r.Date == nil {
			r.Date = &Date{}
		}
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Responsible = m.Responsible
	r.Payee = m.Payee
	r.Amount = m.Amount
	return nil
}
func (r PaymentReconciliationDetail) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Predecessor, xml.StartElement{Name: xml.Name{Local: "predecessor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Request, xml.StartElement{Name: xml.Name{Local: "request"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Submitter, xml.StartElement{Name: xml.Name{Local: "submitter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Response, xml.StartElement{Name: xml.Name{Local: "response"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Responsible, xml.StartElement{Name: xml.Name{Local: "responsible"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Payee, xml.StartElement{Name: xml.Name{Local: "payee"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Amount, xml.StartElement{Name: xml.Name{Local: "amount"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *PaymentReconciliationDetail) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = &v
			case "predecessor":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Predecessor = &v
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "request":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Request = &v
			case "submitter":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Submitter = &v
			case "response":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Response = &v
			case "date":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = &v
			case "responsible":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Responsible = &v
			case "payee":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Payee = &v
			case "amount":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r PaymentReconciliationDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A note that describes or explains the processing in a human readable form.
type PaymentReconciliationProcessNote struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The business purpose of the note text.
	Type *Code
	// The explanation or description associated with the processing.
	Text *String
}
type jsonPaymentReconciliationProcessNote struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Type                 *Code             `json:"type,omitempty"`
	TypePrimitiveElement *primitiveElement `json:"_type,omitempty"`
	Text                 *String           `json:"text,omitempty"`
	TextPrimitiveElement *primitiveElement `json:"_text,omitempty"`
}

func (r PaymentReconciliationProcessNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r PaymentReconciliationProcessNote) marshalJSON() jsonPaymentReconciliationProcessNote {
	m := jsonPaymentReconciliationProcessNote{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Type != nil && r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.Text != nil && r.Text.Value != nil {
		m.Text = r.Text
	}
	if r.Text != nil && (r.Text.Id != nil || r.Text.Extension != nil) {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	return m
}
func (r *PaymentReconciliationProcessNote) UnmarshalJSON(b []byte) error {
	var m jsonPaymentReconciliationProcessNote
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *PaymentReconciliationProcessNote) unmarshalJSON(m jsonPaymentReconciliationProcessNote) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		if r.Type == nil {
			r.Type = &Code{}
		}
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Text = m.Text
	if m.TextPrimitiveElement != nil {
		if r.Text == nil {
			r.Text = &String{}
		}
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	return nil
}
func (r PaymentReconciliationProcessNote) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Text, xml.StartElement{Name: xml.Name{Local: "text"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *PaymentReconciliationProcessNote) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "text":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Text = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r PaymentReconciliationProcessNote) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
