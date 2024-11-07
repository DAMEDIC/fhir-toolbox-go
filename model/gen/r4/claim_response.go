package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// This resource provides the adjudication details from the processing of a Claim resource.
type ClaimResponse struct {
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
	// A unique identifier assigned to this claim response.
	Identifier []Identifier
	// The status of the resource instance.
	Status Code
	// A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
	Type CodeableConcept
	// A finer grained suite of claim type codes which may convey additional information such as Inpatient vs Outpatient and/or a specialty service.
	SubType *CodeableConcept
	// A code to indicate whether the nature of the request is: to request adjudication of products and services previously rendered; or requesting authorization and adjudication for provision in the future; or requesting the non-binding adjudication of the listed products and services which could be provided in the future.
	Use Code
	// The party to whom the professional services and/or products have been supplied or are being considered and for whom actual for facast reimbursement is sought.
	Patient Reference
	// The date this resource was created.
	Created DateTime
	// The party responsible for authorization, adjudication and reimbursement.
	Insurer Reference
	// The provider which is responsible for the claim, predetermination or preauthorization.
	Requestor *Reference
	// Original request resource reference.
	Request *Reference
	// The outcome of the claim, predetermination, or preauthorization processing.
	Outcome Code
	// A human readable description of the status of the adjudication.
	Disposition *String
	// Reference from the Insurer which is used in later communications which refers to this adjudication.
	PreAuthRef *String
	// The time frame during which this authorization is effective.
	PreAuthPeriod *Period
	// Type of Party to be reimbursed: subscriber, provider, other.
	PayeeType *CodeableConcept
	// A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
	Item []ClaimResponseItem
	// The first-tier service adjudications for payor added product or service lines.
	AddItem []ClaimResponseAddItem
	// The adjudication results which are presented at the header level rather than at the line-item or add-item levels.
	Adjudication []ClaimResponseItemAdjudication
	// Categorized monetary totals for the adjudication.
	Total []ClaimResponseTotal
	// Payment details for the adjudication of the claim.
	Payment *ClaimResponsePayment
	// A code, used only on a response to a preauthorization, to indicate whether the benefits payable have been reserved and for whom.
	FundsReserve *CodeableConcept
	// A code for the form to be used for printing the content.
	FormCode *CodeableConcept
	// The actual form, by reference or inclusion, for printing the content or an EOB.
	Form *Attachment
	// A note that describes or explains adjudication results in a human readable form.
	ProcessNote []ClaimResponseProcessNote
	// Request for additional supporting or authorizing information.
	CommunicationRequest []Reference
	// Financial instruments for reimbursement for the health care products and services specified on the claim.
	Insurance []ClaimResponseInsurance
	// Errors encountered during the processing of the adjudication.
	Error []ClaimResponseError
}

func (r ClaimResponse) ResourceType() string {
	return "ClaimResponse"
}
func (r ClaimResponse) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonClaimResponse struct {
	ResourceType                  string                          `json:"resourceType"`
	Id                            *Id                             `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement               `json:"_id,omitempty"`
	Meta                          *Meta                           `json:"meta,omitempty"`
	ImplicitRules                 *Uri                            `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement               `json:"_implicitRules,omitempty"`
	Language                      *Code                           `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement               `json:"_language,omitempty"`
	Text                          *Narrative                      `json:"text,omitempty"`
	Contained                     []ContainedResource             `json:"contained,omitempty"`
	Extension                     []Extension                     `json:"extension,omitempty"`
	ModifierExtension             []Extension                     `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier                    `json:"identifier,omitempty"`
	Status                        Code                            `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement               `json:"_status,omitempty"`
	Type                          CodeableConcept                 `json:"type,omitempty"`
	SubType                       *CodeableConcept                `json:"subType,omitempty"`
	Use                           Code                            `json:"use,omitempty"`
	UsePrimitiveElement           *primitiveElement               `json:"_use,omitempty"`
	Patient                       Reference                       `json:"patient,omitempty"`
	Created                       DateTime                        `json:"created,omitempty"`
	CreatedPrimitiveElement       *primitiveElement               `json:"_created,omitempty"`
	Insurer                       Reference                       `json:"insurer,omitempty"`
	Requestor                     *Reference                      `json:"requestor,omitempty"`
	Request                       *Reference                      `json:"request,omitempty"`
	Outcome                       Code                            `json:"outcome,omitempty"`
	OutcomePrimitiveElement       *primitiveElement               `json:"_outcome,omitempty"`
	Disposition                   *String                         `json:"disposition,omitempty"`
	DispositionPrimitiveElement   *primitiveElement               `json:"_disposition,omitempty"`
	PreAuthRef                    *String                         `json:"preAuthRef,omitempty"`
	PreAuthRefPrimitiveElement    *primitiveElement               `json:"_preAuthRef,omitempty"`
	PreAuthPeriod                 *Period                         `json:"preAuthPeriod,omitempty"`
	PayeeType                     *CodeableConcept                `json:"payeeType,omitempty"`
	Item                          []ClaimResponseItem             `json:"item,omitempty"`
	AddItem                       []ClaimResponseAddItem          `json:"addItem,omitempty"`
	Adjudication                  []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
	Total                         []ClaimResponseTotal            `json:"total,omitempty"`
	Payment                       *ClaimResponsePayment           `json:"payment,omitempty"`
	FundsReserve                  *CodeableConcept                `json:"fundsReserve,omitempty"`
	FormCode                      *CodeableConcept                `json:"formCode,omitempty"`
	Form                          *Attachment                     `json:"form,omitempty"`
	ProcessNote                   []ClaimResponseProcessNote      `json:"processNote,omitempty"`
	CommunicationRequest          []Reference                     `json:"communicationRequest,omitempty"`
	Insurance                     []ClaimResponseInsurance        `json:"insurance,omitempty"`
	Error                         []ClaimResponseError            `json:"error,omitempty"`
}

func (r ClaimResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimResponse) marshalJSON() jsonClaimResponse {
	m := jsonClaimResponse{}
	m.ResourceType = "ClaimResponse"
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
	m.Type = r.Type
	m.SubType = r.SubType
	if r.Use.Value != nil {
		m.Use = r.Use
	}
	if r.Use.Id != nil || r.Use.Extension != nil {
		m.UsePrimitiveElement = &primitiveElement{Id: r.Use.Id, Extension: r.Use.Extension}
	}
	m.Patient = r.Patient
	if r.Created.Value != nil {
		m.Created = r.Created
	}
	if r.Created.Id != nil || r.Created.Extension != nil {
		m.CreatedPrimitiveElement = &primitiveElement{Id: r.Created.Id, Extension: r.Created.Extension}
	}
	m.Insurer = r.Insurer
	m.Requestor = r.Requestor
	m.Request = r.Request
	if r.Outcome.Value != nil {
		m.Outcome = r.Outcome
	}
	if r.Outcome.Id != nil || r.Outcome.Extension != nil {
		m.OutcomePrimitiveElement = &primitiveElement{Id: r.Outcome.Id, Extension: r.Outcome.Extension}
	}
	if r.Disposition != nil && r.Disposition.Value != nil {
		m.Disposition = r.Disposition
	}
	if r.Disposition != nil && (r.Disposition.Id != nil || r.Disposition.Extension != nil) {
		m.DispositionPrimitiveElement = &primitiveElement{Id: r.Disposition.Id, Extension: r.Disposition.Extension}
	}
	if r.PreAuthRef != nil && r.PreAuthRef.Value != nil {
		m.PreAuthRef = r.PreAuthRef
	}
	if r.PreAuthRef != nil && (r.PreAuthRef.Id != nil || r.PreAuthRef.Extension != nil) {
		m.PreAuthRefPrimitiveElement = &primitiveElement{Id: r.PreAuthRef.Id, Extension: r.PreAuthRef.Extension}
	}
	m.PreAuthPeriod = r.PreAuthPeriod
	m.PayeeType = r.PayeeType
	m.Item = r.Item
	m.AddItem = r.AddItem
	m.Adjudication = r.Adjudication
	m.Total = r.Total
	m.Payment = r.Payment
	m.FundsReserve = r.FundsReserve
	m.FormCode = r.FormCode
	m.Form = r.Form
	m.ProcessNote = r.ProcessNote
	m.CommunicationRequest = r.CommunicationRequest
	m.Insurance = r.Insurance
	m.Error = r.Error
	return m
}
func (r *ClaimResponse) UnmarshalJSON(b []byte) error {
	var m jsonClaimResponse
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimResponse) unmarshalJSON(m jsonClaimResponse) error {
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
	r.Type = m.Type
	r.SubType = m.SubType
	r.Use = m.Use
	if m.UsePrimitiveElement != nil {
		r.Use.Id = m.UsePrimitiveElement.Id
		r.Use.Extension = m.UsePrimitiveElement.Extension
	}
	r.Patient = m.Patient
	r.Created = m.Created
	if m.CreatedPrimitiveElement != nil {
		r.Created.Id = m.CreatedPrimitiveElement.Id
		r.Created.Extension = m.CreatedPrimitiveElement.Extension
	}
	r.Insurer = m.Insurer
	r.Requestor = m.Requestor
	r.Request = m.Request
	r.Outcome = m.Outcome
	if m.OutcomePrimitiveElement != nil {
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
	r.PreAuthRef = m.PreAuthRef
	if m.PreAuthRefPrimitiveElement != nil {
		if r.PreAuthRef == nil {
			r.PreAuthRef = &String{}
		}
		r.PreAuthRef.Id = m.PreAuthRefPrimitiveElement.Id
		r.PreAuthRef.Extension = m.PreAuthRefPrimitiveElement.Extension
	}
	r.PreAuthPeriod = m.PreAuthPeriod
	r.PayeeType = m.PayeeType
	r.Item = m.Item
	r.AddItem = m.AddItem
	r.Adjudication = m.Adjudication
	r.Total = m.Total
	r.Payment = m.Payment
	r.FundsReserve = m.FundsReserve
	r.FormCode = m.FormCode
	r.Form = m.Form
	r.ProcessNote = m.ProcessNote
	r.CommunicationRequest = m.CommunicationRequest
	r.Insurance = m.Insurance
	r.Error = m.Error
	return nil
}
func (r ClaimResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "ClaimResponse"
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
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubType, xml.StartElement{Name: xml.Name{Local: "subType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Use, xml.StartElement{Name: xml.Name{Local: "use"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Patient, xml.StartElement{Name: xml.Name{Local: "patient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Created, xml.StartElement{Name: xml.Name{Local: "created"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Insurer, xml.StartElement{Name: xml.Name{Local: "insurer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Requestor, xml.StartElement{Name: xml.Name{Local: "requestor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Request, xml.StartElement{Name: xml.Name{Local: "request"}})
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
	err = e.EncodeElement(r.PreAuthRef, xml.StartElement{Name: xml.Name{Local: "preAuthRef"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PreAuthPeriod, xml.StartElement{Name: xml.Name{Local: "preAuthPeriod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.PayeeType, xml.StartElement{Name: xml.Name{Local: "payeeType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Item, xml.StartElement{Name: xml.Name{Local: "item"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AddItem, xml.StartElement{Name: xml.Name{Local: "addItem"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Adjudication, xml.StartElement{Name: xml.Name{Local: "adjudication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Total, xml.StartElement{Name: xml.Name{Local: "total"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Payment, xml.StartElement{Name: xml.Name{Local: "payment"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FundsReserve, xml.StartElement{Name: xml.Name{Local: "fundsReserve"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FormCode, xml.StartElement{Name: xml.Name{Local: "formCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Form, xml.StartElement{Name: xml.Name{Local: "form"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProcessNote, xml.StartElement{Name: xml.Name{Local: "processNote"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CommunicationRequest, xml.StartElement{Name: xml.Name{Local: "communicationRequest"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Insurance, xml.StartElement{Name: xml.Name{Local: "insurance"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Error, xml.StartElement{Name: xml.Name{Local: "error"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ClaimResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "subType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubType = &v
			case "use":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Use = v
			case "patient":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Patient = v
			case "created":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Created = v
			case "insurer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Insurer = v
			case "requestor":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Requestor = &v
			case "request":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Request = &v
			case "outcome":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Outcome = v
			case "disposition":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Disposition = &v
			case "preAuthRef":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PreAuthRef = &v
			case "preAuthPeriod":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PreAuthPeriod = &v
			case "payeeType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.PayeeType = &v
			case "item":
				var v ClaimResponseItem
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Item = append(r.Item, v)
			case "addItem":
				var v ClaimResponseAddItem
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AddItem = append(r.AddItem, v)
			case "adjudication":
				var v ClaimResponseItemAdjudication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			case "total":
				var v ClaimResponseTotal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Total = append(r.Total, v)
			case "payment":
				var v ClaimResponsePayment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Payment = &v
			case "fundsReserve":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FundsReserve = &v
			case "formCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FormCode = &v
			case "form":
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Form = &v
			case "processNote":
				var v ClaimResponseProcessNote
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProcessNote = append(r.ProcessNote, v)
			case "communicationRequest":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CommunicationRequest = append(r.CommunicationRequest, v)
			case "insurance":
				var v ClaimResponseInsurance
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Insurance = append(r.Insurance, v)
			case "error":
				var v ClaimResponseError
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Error = append(r.Error, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ClaimResponse) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A claim line. Either a simple (a product or service) or a 'group' of details which can also be a simple items or groups of sub-details.
type ClaimResponseItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely reference the claim item entries.
	ItemSequence PositiveInt
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// If this item is a group then the values here are a summary of the adjudication of the detail items. If this item is a simple product or service then this is the result of the adjudication of this item.
	Adjudication []ClaimResponseItemAdjudication
	// A claim detail. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
	Detail []ClaimResponseItemDetail
}
type jsonClaimResponseItem struct {
	Id                           *string                         `json:"id,omitempty"`
	Extension                    []Extension                     `json:"extension,omitempty"`
	ModifierExtension            []Extension                     `json:"modifierExtension,omitempty"`
	ItemSequence                 PositiveInt                     `json:"itemSequence,omitempty"`
	ItemSequencePrimitiveElement *primitiveElement               `json:"_itemSequence,omitempty"`
	NoteNumber                   []PositiveInt                   `json:"noteNumber,omitempty"`
	NoteNumberPrimitiveElement   []*primitiveElement             `json:"_noteNumber,omitempty"`
	Adjudication                 []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
	Detail                       []ClaimResponseItemDetail       `json:"detail,omitempty"`
}

func (r ClaimResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimResponseItem) marshalJSON() jsonClaimResponseItem {
	m := jsonClaimResponseItem{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.ItemSequence.Value != nil {
		m.ItemSequence = r.ItemSequence
	}
	if r.ItemSequence.Id != nil || r.ItemSequence.Extension != nil {
		m.ItemSequencePrimitiveElement = &primitiveElement{Id: r.ItemSequence.Id, Extension: r.ItemSequence.Extension}
	}
	anyNoteNumberValue := false
	for _, e := range r.NoteNumber {
		if e.Value != nil {
			anyNoteNumberValue = true
			break
		}
	}
	if anyNoteNumberValue {
		m.NoteNumber = r.NoteNumber
	}
	anyNoteNumberIdOrExtension := false
	for _, e := range r.NoteNumber {
		if e.Id != nil || e.Extension != nil {
			anyNoteNumberIdOrExtension = true
			break
		}
	}
	if anyNoteNumberIdOrExtension {
		m.NoteNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.NoteNumber))
		for _, e := range r.NoteNumber {
			if e.Id != nil || e.Extension != nil {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, nil)
			}
		}
	}
	m.Adjudication = r.Adjudication
	m.Detail = r.Detail
	return m
}
func (r *ClaimResponseItem) UnmarshalJSON(b []byte) error {
	var m jsonClaimResponseItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimResponseItem) unmarshalJSON(m jsonClaimResponseItem) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ItemSequence = m.ItemSequence
	if m.ItemSequencePrimitiveElement != nil {
		r.ItemSequence.Id = m.ItemSequencePrimitiveElement.Id
		r.ItemSequence.Extension = m.ItemSequencePrimitiveElement.Extension
	}
	r.NoteNumber = m.NoteNumber
	for i, e := range m.NoteNumberPrimitiveElement {
		if len(r.NoteNumber) <= i {
			r.NoteNumber = append(r.NoteNumber, PositiveInt{})
		}
		if e != nil {
			r.NoteNumber[i].Id = e.Id
			r.NoteNumber[i].Extension = e.Extension
		}
	}
	r.Adjudication = m.Adjudication
	r.Detail = m.Detail
	return nil
}
func (r ClaimResponseItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ItemSequence, xml.StartElement{Name: xml.Name{Local: "itemSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NoteNumber, xml.StartElement{Name: xml.Name{Local: "noteNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Adjudication, xml.StartElement{Name: xml.Name{Local: "adjudication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Detail, xml.StartElement{Name: xml.Name{Local: "detail"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ClaimResponseItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "itemSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ItemSequence = v
			case "noteNumber":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NoteNumber = append(r.NoteNumber, v)
			case "adjudication":
				var v ClaimResponseItemAdjudication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			case "detail":
				var v ClaimResponseItemDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = append(r.Detail, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ClaimResponseItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// If this item is a group then the values here are a summary of the adjudication of the detail items. If this item is a simple product or service then this is the result of the adjudication of this item.
type ClaimResponseItemAdjudication struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code to indicate the information type of this adjudication record. Information types may include the value submitted, maximum values or percentages allowed or payable under the plan, amounts that: the patient is responsible for in aggregate or pertaining to this item; amounts paid by other coverages; and, the benefit payable for this item.
	Category CodeableConcept
	// A code supporting the understanding of the adjudication result and explaining variance from expected amount.
	Reason *CodeableConcept
	// Monetary amount associated with the category.
	Amount *Money
	// A non-monetary value associated with the category. Mutually exclusive to the amount element above.
	Value *Decimal
}
type jsonClaimResponseItemAdjudication struct {
	Id                    *string           `json:"id,omitempty"`
	Extension             []Extension       `json:"extension,omitempty"`
	ModifierExtension     []Extension       `json:"modifierExtension,omitempty"`
	Category              CodeableConcept   `json:"category,omitempty"`
	Reason                *CodeableConcept  `json:"reason,omitempty"`
	Amount                *Money            `json:"amount,omitempty"`
	Value                 *Decimal          `json:"value,omitempty"`
	ValuePrimitiveElement *primitiveElement `json:"_value,omitempty"`
}

func (r ClaimResponseItemAdjudication) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimResponseItemAdjudication) marshalJSON() jsonClaimResponseItemAdjudication {
	m := jsonClaimResponseItemAdjudication{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Category = r.Category
	m.Reason = r.Reason
	m.Amount = r.Amount
	if r.Value != nil && r.Value.Value != nil {
		m.Value = r.Value
	}
	if r.Value != nil && (r.Value.Id != nil || r.Value.Extension != nil) {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	return m
}
func (r *ClaimResponseItemAdjudication) UnmarshalJSON(b []byte) error {
	var m jsonClaimResponseItemAdjudication
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimResponseItemAdjudication) unmarshalJSON(m jsonClaimResponseItemAdjudication) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Category = m.Category
	r.Reason = m.Reason
	r.Amount = m.Amount
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		if r.Value == nil {
			r.Value = &Decimal{}
		}
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	return nil
}
func (r ClaimResponseItemAdjudication) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Category, xml.StartElement{Name: xml.Name{Local: "category"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Reason, xml.StartElement{Name: xml.Name{Local: "reason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Amount, xml.StartElement{Name: xml.Name{Local: "amount"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Value, xml.StartElement{Name: xml.Name{Local: "value"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ClaimResponseItemAdjudication) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "category":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = v
			case "reason":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Reason = &v
			case "amount":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = &v
			case "value":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Value = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ClaimResponseItemAdjudication) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A claim detail. Either a simple (a product or service) or a 'group' of sub-details which are simple items.
type ClaimResponseItemDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely reference the claim detail entry.
	DetailSequence PositiveInt
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
	// A sub-detail adjudication of a simple product or service.
	SubDetail []ClaimResponseItemDetailSubDetail
}
type jsonClaimResponseItemDetail struct {
	Id                             *string                            `json:"id,omitempty"`
	Extension                      []Extension                        `json:"extension,omitempty"`
	ModifierExtension              []Extension                        `json:"modifierExtension,omitempty"`
	DetailSequence                 PositiveInt                        `json:"detailSequence,omitempty"`
	DetailSequencePrimitiveElement *primitiveElement                  `json:"_detailSequence,omitempty"`
	NoteNumber                     []PositiveInt                      `json:"noteNumber,omitempty"`
	NoteNumberPrimitiveElement     []*primitiveElement                `json:"_noteNumber,omitempty"`
	Adjudication                   []ClaimResponseItemAdjudication    `json:"adjudication,omitempty"`
	SubDetail                      []ClaimResponseItemDetailSubDetail `json:"subDetail,omitempty"`
}

func (r ClaimResponseItemDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimResponseItemDetail) marshalJSON() jsonClaimResponseItemDetail {
	m := jsonClaimResponseItemDetail{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.DetailSequence.Value != nil {
		m.DetailSequence = r.DetailSequence
	}
	if r.DetailSequence.Id != nil || r.DetailSequence.Extension != nil {
		m.DetailSequencePrimitiveElement = &primitiveElement{Id: r.DetailSequence.Id, Extension: r.DetailSequence.Extension}
	}
	anyNoteNumberValue := false
	for _, e := range r.NoteNumber {
		if e.Value != nil {
			anyNoteNumberValue = true
			break
		}
	}
	if anyNoteNumberValue {
		m.NoteNumber = r.NoteNumber
	}
	anyNoteNumberIdOrExtension := false
	for _, e := range r.NoteNumber {
		if e.Id != nil || e.Extension != nil {
			anyNoteNumberIdOrExtension = true
			break
		}
	}
	if anyNoteNumberIdOrExtension {
		m.NoteNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.NoteNumber))
		for _, e := range r.NoteNumber {
			if e.Id != nil || e.Extension != nil {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, nil)
			}
		}
	}
	m.Adjudication = r.Adjudication
	m.SubDetail = r.SubDetail
	return m
}
func (r *ClaimResponseItemDetail) UnmarshalJSON(b []byte) error {
	var m jsonClaimResponseItemDetail
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimResponseItemDetail) unmarshalJSON(m jsonClaimResponseItemDetail) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.DetailSequence = m.DetailSequence
	if m.DetailSequencePrimitiveElement != nil {
		r.DetailSequence.Id = m.DetailSequencePrimitiveElement.Id
		r.DetailSequence.Extension = m.DetailSequencePrimitiveElement.Extension
	}
	r.NoteNumber = m.NoteNumber
	for i, e := range m.NoteNumberPrimitiveElement {
		if len(r.NoteNumber) <= i {
			r.NoteNumber = append(r.NoteNumber, PositiveInt{})
		}
		if e != nil {
			r.NoteNumber[i].Id = e.Id
			r.NoteNumber[i].Extension = e.Extension
		}
	}
	r.Adjudication = m.Adjudication
	r.SubDetail = m.SubDetail
	return nil
}
func (r ClaimResponseItemDetail) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.DetailSequence, xml.StartElement{Name: xml.Name{Local: "detailSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NoteNumber, xml.StartElement{Name: xml.Name{Local: "noteNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Adjudication, xml.StartElement{Name: xml.Name{Local: "adjudication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubDetail, xml.StartElement{Name: xml.Name{Local: "subDetail"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ClaimResponseItemDetail) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "detailSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DetailSequence = v
			case "noteNumber":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NoteNumber = append(r.NoteNumber, v)
			case "adjudication":
				var v ClaimResponseItemAdjudication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			case "subDetail":
				var v ClaimResponseItemDetailSubDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubDetail = append(r.SubDetail, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ClaimResponseItemDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A sub-detail adjudication of a simple product or service.
type ClaimResponseItemDetailSubDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely reference the claim sub-detail entry.
	SubDetailSequence PositiveInt
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
}
type jsonClaimResponseItemDetailSubDetail struct {
	Id                                *string                         `json:"id,omitempty"`
	Extension                         []Extension                     `json:"extension,omitempty"`
	ModifierExtension                 []Extension                     `json:"modifierExtension,omitempty"`
	SubDetailSequence                 PositiveInt                     `json:"subDetailSequence,omitempty"`
	SubDetailSequencePrimitiveElement *primitiveElement               `json:"_subDetailSequence,omitempty"`
	NoteNumber                        []PositiveInt                   `json:"noteNumber,omitempty"`
	NoteNumberPrimitiveElement        []*primitiveElement             `json:"_noteNumber,omitempty"`
	Adjudication                      []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
}

func (r ClaimResponseItemDetailSubDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimResponseItemDetailSubDetail) marshalJSON() jsonClaimResponseItemDetailSubDetail {
	m := jsonClaimResponseItemDetailSubDetail{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.SubDetailSequence.Value != nil {
		m.SubDetailSequence = r.SubDetailSequence
	}
	if r.SubDetailSequence.Id != nil || r.SubDetailSequence.Extension != nil {
		m.SubDetailSequencePrimitiveElement = &primitiveElement{Id: r.SubDetailSequence.Id, Extension: r.SubDetailSequence.Extension}
	}
	anyNoteNumberValue := false
	for _, e := range r.NoteNumber {
		if e.Value != nil {
			anyNoteNumberValue = true
			break
		}
	}
	if anyNoteNumberValue {
		m.NoteNumber = r.NoteNumber
	}
	anyNoteNumberIdOrExtension := false
	for _, e := range r.NoteNumber {
		if e.Id != nil || e.Extension != nil {
			anyNoteNumberIdOrExtension = true
			break
		}
	}
	if anyNoteNumberIdOrExtension {
		m.NoteNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.NoteNumber))
		for _, e := range r.NoteNumber {
			if e.Id != nil || e.Extension != nil {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, nil)
			}
		}
	}
	m.Adjudication = r.Adjudication
	return m
}
func (r *ClaimResponseItemDetailSubDetail) UnmarshalJSON(b []byte) error {
	var m jsonClaimResponseItemDetailSubDetail
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimResponseItemDetailSubDetail) unmarshalJSON(m jsonClaimResponseItemDetailSubDetail) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.SubDetailSequence = m.SubDetailSequence
	if m.SubDetailSequencePrimitiveElement != nil {
		r.SubDetailSequence.Id = m.SubDetailSequencePrimitiveElement.Id
		r.SubDetailSequence.Extension = m.SubDetailSequencePrimitiveElement.Extension
	}
	r.NoteNumber = m.NoteNumber
	for i, e := range m.NoteNumberPrimitiveElement {
		if len(r.NoteNumber) <= i {
			r.NoteNumber = append(r.NoteNumber, PositiveInt{})
		}
		if e != nil {
			r.NoteNumber[i].Id = e.Id
			r.NoteNumber[i].Extension = e.Extension
		}
	}
	r.Adjudication = m.Adjudication
	return nil
}
func (r ClaimResponseItemDetailSubDetail) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.SubDetailSequence, xml.StartElement{Name: xml.Name{Local: "subDetailSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NoteNumber, xml.StartElement{Name: xml.Name{Local: "noteNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Adjudication, xml.StartElement{Name: xml.Name{Local: "adjudication"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ClaimResponseItemDetailSubDetail) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "subDetailSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubDetailSequence = v
			case "noteNumber":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NoteNumber = append(r.NoteNumber, v)
			case "adjudication":
				var v ClaimResponseItemAdjudication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ClaimResponseItemDetailSubDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The first-tier service adjudications for payor added product or service lines.
type ClaimResponseAddItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Claim items which this service line is intended to replace.
	ItemSequence []PositiveInt
	// The sequence number of the details within the claim item which this line is intended to replace.
	DetailSequence []PositiveInt
	// The sequence number of the sub-details within the details within the claim item which this line is intended to replace.
	SubdetailSequence []PositiveInt
	// The providers who are authorized for the services rendered to the patient.
	Provider []Reference
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []CodeableConcept
	// Identifies the program under which this may be recovered.
	ProgramCode []CodeableConcept
	// The date or dates when the service or product was supplied, performed or completed.
	Serviced isClaimResponseAddItemServiced
	// Where the product or service was provided.
	Location isClaimResponseAddItemLocation
	// The number of repetitions of a service or product.
	Quantity *Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *Decimal
	// The quantity times the unit price for an additional service or product or charge.
	Net *Money
	// Physical service site on the patient (limb, tooth, etc.).
	BodySite *CodeableConcept
	// A region or surface of the bodySite, e.g. limb region or tooth surface(s).
	SubSite []CodeableConcept
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
	// The second-tier service adjudications for payor added services.
	Detail []ClaimResponseAddItemDetail
}
type isClaimResponseAddItemServiced interface {
	isClaimResponseAddItemServiced()
}

func (r Date) isClaimResponseAddItemServiced()   {}
func (r Period) isClaimResponseAddItemServiced() {}

type isClaimResponseAddItemLocation interface {
	isClaimResponseAddItemLocation()
}

func (r CodeableConcept) isClaimResponseAddItemLocation() {}
func (r Address) isClaimResponseAddItemLocation()         {}
func (r Reference) isClaimResponseAddItemLocation()       {}

type jsonClaimResponseAddItem struct {
	Id                                *string                         `json:"id,omitempty"`
	Extension                         []Extension                     `json:"extension,omitempty"`
	ModifierExtension                 []Extension                     `json:"modifierExtension,omitempty"`
	ItemSequence                      []PositiveInt                   `json:"itemSequence,omitempty"`
	ItemSequencePrimitiveElement      []*primitiveElement             `json:"_itemSequence,omitempty"`
	DetailSequence                    []PositiveInt                   `json:"detailSequence,omitempty"`
	DetailSequencePrimitiveElement    []*primitiveElement             `json:"_detailSequence,omitempty"`
	SubdetailSequence                 []PositiveInt                   `json:"subdetailSequence,omitempty"`
	SubdetailSequencePrimitiveElement []*primitiveElement             `json:"_subdetailSequence,omitempty"`
	Provider                          []Reference                     `json:"provider,omitempty"`
	ProductOrService                  CodeableConcept                 `json:"productOrService,omitempty"`
	Modifier                          []CodeableConcept               `json:"modifier,omitempty"`
	ProgramCode                       []CodeableConcept               `json:"programCode,omitempty"`
	ServicedDate                      *Date                           `json:"servicedDate,omitempty"`
	ServicedDatePrimitiveElement      *primitiveElement               `json:"_servicedDate,omitempty"`
	ServicedPeriod                    *Period                         `json:"servicedPeriod,omitempty"`
	LocationCodeableConcept           *CodeableConcept                `json:"locationCodeableConcept,omitempty"`
	LocationAddress                   *Address                        `json:"locationAddress,omitempty"`
	LocationReference                 *Reference                      `json:"locationReference,omitempty"`
	Quantity                          *Quantity                       `json:"quantity,omitempty"`
	UnitPrice                         *Money                          `json:"unitPrice,omitempty"`
	Factor                            *Decimal                        `json:"factor,omitempty"`
	FactorPrimitiveElement            *primitiveElement               `json:"_factor,omitempty"`
	Net                               *Money                          `json:"net,omitempty"`
	BodySite                          *CodeableConcept                `json:"bodySite,omitempty"`
	SubSite                           []CodeableConcept               `json:"subSite,omitempty"`
	NoteNumber                        []PositiveInt                   `json:"noteNumber,omitempty"`
	NoteNumberPrimitiveElement        []*primitiveElement             `json:"_noteNumber,omitempty"`
	Adjudication                      []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
	Detail                            []ClaimResponseAddItemDetail    `json:"detail,omitempty"`
}

func (r ClaimResponseAddItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimResponseAddItem) marshalJSON() jsonClaimResponseAddItem {
	m := jsonClaimResponseAddItem{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	anyItemSequenceValue := false
	for _, e := range r.ItemSequence {
		if e.Value != nil {
			anyItemSequenceValue = true
			break
		}
	}
	if anyItemSequenceValue {
		m.ItemSequence = r.ItemSequence
	}
	anyItemSequenceIdOrExtension := false
	for _, e := range r.ItemSequence {
		if e.Id != nil || e.Extension != nil {
			anyItemSequenceIdOrExtension = true
			break
		}
	}
	if anyItemSequenceIdOrExtension {
		m.ItemSequencePrimitiveElement = make([]*primitiveElement, 0, len(r.ItemSequence))
		for _, e := range r.ItemSequence {
			if e.Id != nil || e.Extension != nil {
				m.ItemSequencePrimitiveElement = append(m.ItemSequencePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ItemSequencePrimitiveElement = append(m.ItemSequencePrimitiveElement, nil)
			}
		}
	}
	anyDetailSequenceValue := false
	for _, e := range r.DetailSequence {
		if e.Value != nil {
			anyDetailSequenceValue = true
			break
		}
	}
	if anyDetailSequenceValue {
		m.DetailSequence = r.DetailSequence
	}
	anyDetailSequenceIdOrExtension := false
	for _, e := range r.DetailSequence {
		if e.Id != nil || e.Extension != nil {
			anyDetailSequenceIdOrExtension = true
			break
		}
	}
	if anyDetailSequenceIdOrExtension {
		m.DetailSequencePrimitiveElement = make([]*primitiveElement, 0, len(r.DetailSequence))
		for _, e := range r.DetailSequence {
			if e.Id != nil || e.Extension != nil {
				m.DetailSequencePrimitiveElement = append(m.DetailSequencePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.DetailSequencePrimitiveElement = append(m.DetailSequencePrimitiveElement, nil)
			}
		}
	}
	anySubdetailSequenceValue := false
	for _, e := range r.SubdetailSequence {
		if e.Value != nil {
			anySubdetailSequenceValue = true
			break
		}
	}
	if anySubdetailSequenceValue {
		m.SubdetailSequence = r.SubdetailSequence
	}
	anySubdetailSequenceIdOrExtension := false
	for _, e := range r.SubdetailSequence {
		if e.Id != nil || e.Extension != nil {
			anySubdetailSequenceIdOrExtension = true
			break
		}
	}
	if anySubdetailSequenceIdOrExtension {
		m.SubdetailSequencePrimitiveElement = make([]*primitiveElement, 0, len(r.SubdetailSequence))
		for _, e := range r.SubdetailSequence {
			if e.Id != nil || e.Extension != nil {
				m.SubdetailSequencePrimitiveElement = append(m.SubdetailSequencePrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.SubdetailSequencePrimitiveElement = append(m.SubdetailSequencePrimitiveElement, nil)
			}
		}
	}
	m.Provider = r.Provider
	m.ProductOrService = r.ProductOrService
	m.Modifier = r.Modifier
	m.ProgramCode = r.ProgramCode
	switch v := r.Serviced.(type) {
	case Date:
		if v.Value != nil {
			m.ServicedDate = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.ServicedDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Date:
		if v.Value != nil {
			m.ServicedDate = v
		}
		if v.Id != nil || v.Extension != nil {
			m.ServicedDatePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Period:
		m.ServicedPeriod = &v
	case *Period:
		m.ServicedPeriod = v
	}
	switch v := r.Location.(type) {
	case CodeableConcept:
		m.LocationCodeableConcept = &v
	case *CodeableConcept:
		m.LocationCodeableConcept = v
	case Address:
		m.LocationAddress = &v
	case *Address:
		m.LocationAddress = v
	case Reference:
		m.LocationReference = &v
	case *Reference:
		m.LocationReference = v
	}
	m.Quantity = r.Quantity
	m.UnitPrice = r.UnitPrice
	if r.Factor != nil && r.Factor.Value != nil {
		m.Factor = r.Factor
	}
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		m.FactorPrimitiveElement = &primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
	}
	m.Net = r.Net
	m.BodySite = r.BodySite
	m.SubSite = r.SubSite
	anyNoteNumberValue := false
	for _, e := range r.NoteNumber {
		if e.Value != nil {
			anyNoteNumberValue = true
			break
		}
	}
	if anyNoteNumberValue {
		m.NoteNumber = r.NoteNumber
	}
	anyNoteNumberIdOrExtension := false
	for _, e := range r.NoteNumber {
		if e.Id != nil || e.Extension != nil {
			anyNoteNumberIdOrExtension = true
			break
		}
	}
	if anyNoteNumberIdOrExtension {
		m.NoteNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.NoteNumber))
		for _, e := range r.NoteNumber {
			if e.Id != nil || e.Extension != nil {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, nil)
			}
		}
	}
	m.Adjudication = r.Adjudication
	m.Detail = r.Detail
	return m
}
func (r *ClaimResponseAddItem) UnmarshalJSON(b []byte) error {
	var m jsonClaimResponseAddItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimResponseAddItem) unmarshalJSON(m jsonClaimResponseAddItem) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ItemSequence = m.ItemSequence
	for i, e := range m.ItemSequencePrimitiveElement {
		if len(r.ItemSequence) <= i {
			r.ItemSequence = append(r.ItemSequence, PositiveInt{})
		}
		if e != nil {
			r.ItemSequence[i].Id = e.Id
			r.ItemSequence[i].Extension = e.Extension
		}
	}
	r.DetailSequence = m.DetailSequence
	for i, e := range m.DetailSequencePrimitiveElement {
		if len(r.DetailSequence) <= i {
			r.DetailSequence = append(r.DetailSequence, PositiveInt{})
		}
		if e != nil {
			r.DetailSequence[i].Id = e.Id
			r.DetailSequence[i].Extension = e.Extension
		}
	}
	r.SubdetailSequence = m.SubdetailSequence
	for i, e := range m.SubdetailSequencePrimitiveElement {
		if len(r.SubdetailSequence) <= i {
			r.SubdetailSequence = append(r.SubdetailSequence, PositiveInt{})
		}
		if e != nil {
			r.SubdetailSequence[i].Id = e.Id
			r.SubdetailSequence[i].Extension = e.Extension
		}
	}
	r.Provider = m.Provider
	r.ProductOrService = m.ProductOrService
	r.Modifier = m.Modifier
	r.ProgramCode = m.ProgramCode
	if m.ServicedDate != nil || m.ServicedDatePrimitiveElement != nil {
		if r.Serviced != nil {
			return fmt.Errorf("multiple values for field \"Serviced\"")
		}
		v := m.ServicedDate
		if m.ServicedDatePrimitiveElement != nil {
			if v == nil {
				v = &Date{}
			}
			v.Id = m.ServicedDatePrimitiveElement.Id
			v.Extension = m.ServicedDatePrimitiveElement.Extension
		}
		r.Serviced = v
	}
	if m.ServicedPeriod != nil {
		if r.Serviced != nil {
			return fmt.Errorf("multiple values for field \"Serviced\"")
		}
		v := m.ServicedPeriod
		r.Serviced = v
	}
	if m.LocationCodeableConcept != nil {
		if r.Location != nil {
			return fmt.Errorf("multiple values for field \"Location\"")
		}
		v := m.LocationCodeableConcept
		r.Location = v
	}
	if m.LocationAddress != nil {
		if r.Location != nil {
			return fmt.Errorf("multiple values for field \"Location\"")
		}
		v := m.LocationAddress
		r.Location = v
	}
	if m.LocationReference != nil {
		if r.Location != nil {
			return fmt.Errorf("multiple values for field \"Location\"")
		}
		v := m.LocationReference
		r.Location = v
	}
	r.Quantity = m.Quantity
	r.UnitPrice = m.UnitPrice
	r.Factor = m.Factor
	if m.FactorPrimitiveElement != nil {
		if r.Factor == nil {
			r.Factor = &Decimal{}
		}
		r.Factor.Id = m.FactorPrimitiveElement.Id
		r.Factor.Extension = m.FactorPrimitiveElement.Extension
	}
	r.Net = m.Net
	r.BodySite = m.BodySite
	r.SubSite = m.SubSite
	r.NoteNumber = m.NoteNumber
	for i, e := range m.NoteNumberPrimitiveElement {
		if len(r.NoteNumber) <= i {
			r.NoteNumber = append(r.NoteNumber, PositiveInt{})
		}
		if e != nil {
			r.NoteNumber[i].Id = e.Id
			r.NoteNumber[i].Extension = e.Extension
		}
	}
	r.Adjudication = m.Adjudication
	r.Detail = m.Detail
	return nil
}
func (r ClaimResponseAddItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ItemSequence, xml.StartElement{Name: xml.Name{Local: "itemSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DetailSequence, xml.StartElement{Name: xml.Name{Local: "detailSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubdetailSequence, xml.StartElement{Name: xml.Name{Local: "subdetailSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Provider, xml.StartElement{Name: xml.Name{Local: "provider"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProductOrService, xml.StartElement{Name: xml.Name{Local: "productOrService"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Modifier, xml.StartElement{Name: xml.Name{Local: "modifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProgramCode, xml.StartElement{Name: xml.Name{Local: "programCode"}})
	if err != nil {
		return err
	}
	switch v := r.Serviced.(type) {
	case Date, *Date:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "servicedDate"}})
		if err != nil {
			return err
		}
	case Period, *Period:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "servicedPeriod"}})
		if err != nil {
			return err
		}
	}
	switch v := r.Location.(type) {
	case CodeableConcept, *CodeableConcept:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "locationCodeableConcept"}})
		if err != nil {
			return err
		}
	case Address, *Address:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "locationAddress"}})
		if err != nil {
			return err
		}
	case Reference, *Reference:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "locationReference"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UnitPrice, xml.StartElement{Name: xml.Name{Local: "unitPrice"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Factor, xml.StartElement{Name: xml.Name{Local: "factor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Net, xml.StartElement{Name: xml.Name{Local: "net"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BodySite, xml.StartElement{Name: xml.Name{Local: "bodySite"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubSite, xml.StartElement{Name: xml.Name{Local: "subSite"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NoteNumber, xml.StartElement{Name: xml.Name{Local: "noteNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Adjudication, xml.StartElement{Name: xml.Name{Local: "adjudication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Detail, xml.StartElement{Name: xml.Name{Local: "detail"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ClaimResponseAddItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "itemSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ItemSequence = append(r.ItemSequence, v)
			case "detailSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DetailSequence = append(r.DetailSequence, v)
			case "subdetailSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubdetailSequence = append(r.SubdetailSequence, v)
			case "provider":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Provider = append(r.Provider, v)
			case "productOrService":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductOrService = v
			case "modifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Modifier = append(r.Modifier, v)
			case "programCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProgramCode = append(r.ProgramCode, v)
			case "servicedDate":
				if r.Serviced != nil {
					return fmt.Errorf("multiple values for field \"Serviced\"")
				}
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Serviced = v
			case "servicedPeriod":
				if r.Serviced != nil {
					return fmt.Errorf("multiple values for field \"Serviced\"")
				}
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Serviced = v
			case "locationCodeableConcept":
				if r.Location != nil {
					return fmt.Errorf("multiple values for field \"Location\"")
				}
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = v
			case "locationAddress":
				if r.Location != nil {
					return fmt.Errorf("multiple values for field \"Location\"")
				}
				var v Address
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = v
			case "locationReference":
				if r.Location != nil {
					return fmt.Errorf("multiple values for field \"Location\"")
				}
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = v
			case "quantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = &v
			case "unitPrice":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UnitPrice = &v
			case "factor":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Factor = &v
			case "net":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Net = &v
			case "bodySite":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BodySite = &v
			case "subSite":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubSite = append(r.SubSite, v)
			case "noteNumber":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NoteNumber = append(r.NoteNumber, v)
			case "adjudication":
				var v ClaimResponseItemAdjudication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			case "detail":
				var v ClaimResponseAddItemDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Detail = append(r.Detail, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ClaimResponseAddItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The second-tier service adjudications for payor added services.
type ClaimResponseAddItemDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *Decimal
	// The quantity times the unit price for an additional service or product or charge.
	Net *Money
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
	// The third-tier service adjudications for payor added services.
	SubDetail []ClaimResponseAddItemDetailSubDetail
}
type jsonClaimResponseAddItemDetail struct {
	Id                         *string                               `json:"id,omitempty"`
	Extension                  []Extension                           `json:"extension,omitempty"`
	ModifierExtension          []Extension                           `json:"modifierExtension,omitempty"`
	ProductOrService           CodeableConcept                       `json:"productOrService,omitempty"`
	Modifier                   []CodeableConcept                     `json:"modifier,omitempty"`
	Quantity                   *Quantity                             `json:"quantity,omitempty"`
	UnitPrice                  *Money                                `json:"unitPrice,omitempty"`
	Factor                     *Decimal                              `json:"factor,omitempty"`
	FactorPrimitiveElement     *primitiveElement                     `json:"_factor,omitempty"`
	Net                        *Money                                `json:"net,omitempty"`
	NoteNumber                 []PositiveInt                         `json:"noteNumber,omitempty"`
	NoteNumberPrimitiveElement []*primitiveElement                   `json:"_noteNumber,omitempty"`
	Adjudication               []ClaimResponseItemAdjudication       `json:"adjudication,omitempty"`
	SubDetail                  []ClaimResponseAddItemDetailSubDetail `json:"subDetail,omitempty"`
}

func (r ClaimResponseAddItemDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimResponseAddItemDetail) marshalJSON() jsonClaimResponseAddItemDetail {
	m := jsonClaimResponseAddItemDetail{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.ProductOrService = r.ProductOrService
	m.Modifier = r.Modifier
	m.Quantity = r.Quantity
	m.UnitPrice = r.UnitPrice
	if r.Factor != nil && r.Factor.Value != nil {
		m.Factor = r.Factor
	}
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		m.FactorPrimitiveElement = &primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
	}
	m.Net = r.Net
	anyNoteNumberValue := false
	for _, e := range r.NoteNumber {
		if e.Value != nil {
			anyNoteNumberValue = true
			break
		}
	}
	if anyNoteNumberValue {
		m.NoteNumber = r.NoteNumber
	}
	anyNoteNumberIdOrExtension := false
	for _, e := range r.NoteNumber {
		if e.Id != nil || e.Extension != nil {
			anyNoteNumberIdOrExtension = true
			break
		}
	}
	if anyNoteNumberIdOrExtension {
		m.NoteNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.NoteNumber))
		for _, e := range r.NoteNumber {
			if e.Id != nil || e.Extension != nil {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, nil)
			}
		}
	}
	m.Adjudication = r.Adjudication
	m.SubDetail = r.SubDetail
	return m
}
func (r *ClaimResponseAddItemDetail) UnmarshalJSON(b []byte) error {
	var m jsonClaimResponseAddItemDetail
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimResponseAddItemDetail) unmarshalJSON(m jsonClaimResponseAddItemDetail) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ProductOrService = m.ProductOrService
	r.Modifier = m.Modifier
	r.Quantity = m.Quantity
	r.UnitPrice = m.UnitPrice
	r.Factor = m.Factor
	if m.FactorPrimitiveElement != nil {
		if r.Factor == nil {
			r.Factor = &Decimal{}
		}
		r.Factor.Id = m.FactorPrimitiveElement.Id
		r.Factor.Extension = m.FactorPrimitiveElement.Extension
	}
	r.Net = m.Net
	r.NoteNumber = m.NoteNumber
	for i, e := range m.NoteNumberPrimitiveElement {
		if len(r.NoteNumber) <= i {
			r.NoteNumber = append(r.NoteNumber, PositiveInt{})
		}
		if e != nil {
			r.NoteNumber[i].Id = e.Id
			r.NoteNumber[i].Extension = e.Extension
		}
	}
	r.Adjudication = m.Adjudication
	r.SubDetail = m.SubDetail
	return nil
}
func (r ClaimResponseAddItemDetail) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ProductOrService, xml.StartElement{Name: xml.Name{Local: "productOrService"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Modifier, xml.StartElement{Name: xml.Name{Local: "modifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UnitPrice, xml.StartElement{Name: xml.Name{Local: "unitPrice"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Factor, xml.StartElement{Name: xml.Name{Local: "factor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Net, xml.StartElement{Name: xml.Name{Local: "net"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NoteNumber, xml.StartElement{Name: xml.Name{Local: "noteNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Adjudication, xml.StartElement{Name: xml.Name{Local: "adjudication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubDetail, xml.StartElement{Name: xml.Name{Local: "subDetail"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ClaimResponseAddItemDetail) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "productOrService":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductOrService = v
			case "modifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Modifier = append(r.Modifier, v)
			case "quantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = &v
			case "unitPrice":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UnitPrice = &v
			case "factor":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Factor = &v
			case "net":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Net = &v
			case "noteNumber":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NoteNumber = append(r.NoteNumber, v)
			case "adjudication":
				var v ClaimResponseItemAdjudication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			case "subDetail":
				var v ClaimResponseAddItemDetailSubDetail
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubDetail = append(r.SubDetail, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ClaimResponseAddItemDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The third-tier service adjudications for payor added services.
type ClaimResponseAddItemDetailSubDetail struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// When the value is a group code then this item collects a set of related claim details, otherwise this contains the product, service, drug or other billing code for the item.
	ProductOrService CodeableConcept
	// Item typification or modifiers codes to convey additional context for the product or service.
	Modifier []CodeableConcept
	// The number of repetitions of a service or product.
	Quantity *Quantity
	// If the item is not a group then this is the fee for the product or service, otherwise this is the total of the fees for the details of the group.
	UnitPrice *Money
	// A real number that represents a multiplier used in determining the overall value of services delivered and/or goods received. The concept of a Factor allows for a discount or surcharge multiplier to be applied to a monetary amount.
	Factor *Decimal
	// The quantity times the unit price for an additional service or product or charge.
	Net *Money
	// The numbers associated with notes below which apply to the adjudication of this item.
	NoteNumber []PositiveInt
	// The adjudication results.
	Adjudication []ClaimResponseItemAdjudication
}
type jsonClaimResponseAddItemDetailSubDetail struct {
	Id                         *string                         `json:"id,omitempty"`
	Extension                  []Extension                     `json:"extension,omitempty"`
	ModifierExtension          []Extension                     `json:"modifierExtension,omitempty"`
	ProductOrService           CodeableConcept                 `json:"productOrService,omitempty"`
	Modifier                   []CodeableConcept               `json:"modifier,omitempty"`
	Quantity                   *Quantity                       `json:"quantity,omitempty"`
	UnitPrice                  *Money                          `json:"unitPrice,omitempty"`
	Factor                     *Decimal                        `json:"factor,omitempty"`
	FactorPrimitiveElement     *primitiveElement               `json:"_factor,omitempty"`
	Net                        *Money                          `json:"net,omitempty"`
	NoteNumber                 []PositiveInt                   `json:"noteNumber,omitempty"`
	NoteNumberPrimitiveElement []*primitiveElement             `json:"_noteNumber,omitempty"`
	Adjudication               []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
}

func (r ClaimResponseAddItemDetailSubDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimResponseAddItemDetailSubDetail) marshalJSON() jsonClaimResponseAddItemDetailSubDetail {
	m := jsonClaimResponseAddItemDetailSubDetail{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.ProductOrService = r.ProductOrService
	m.Modifier = r.Modifier
	m.Quantity = r.Quantity
	m.UnitPrice = r.UnitPrice
	if r.Factor != nil && r.Factor.Value != nil {
		m.Factor = r.Factor
	}
	if r.Factor != nil && (r.Factor.Id != nil || r.Factor.Extension != nil) {
		m.FactorPrimitiveElement = &primitiveElement{Id: r.Factor.Id, Extension: r.Factor.Extension}
	}
	m.Net = r.Net
	anyNoteNumberValue := false
	for _, e := range r.NoteNumber {
		if e.Value != nil {
			anyNoteNumberValue = true
			break
		}
	}
	if anyNoteNumberValue {
		m.NoteNumber = r.NoteNumber
	}
	anyNoteNumberIdOrExtension := false
	for _, e := range r.NoteNumber {
		if e.Id != nil || e.Extension != nil {
			anyNoteNumberIdOrExtension = true
			break
		}
	}
	if anyNoteNumberIdOrExtension {
		m.NoteNumberPrimitiveElement = make([]*primitiveElement, 0, len(r.NoteNumber))
		for _, e := range r.NoteNumber {
			if e.Id != nil || e.Extension != nil {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.NoteNumberPrimitiveElement = append(m.NoteNumberPrimitiveElement, nil)
			}
		}
	}
	m.Adjudication = r.Adjudication
	return m
}
func (r *ClaimResponseAddItemDetailSubDetail) UnmarshalJSON(b []byte) error {
	var m jsonClaimResponseAddItemDetailSubDetail
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimResponseAddItemDetailSubDetail) unmarshalJSON(m jsonClaimResponseAddItemDetailSubDetail) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ProductOrService = m.ProductOrService
	r.Modifier = m.Modifier
	r.Quantity = m.Quantity
	r.UnitPrice = m.UnitPrice
	r.Factor = m.Factor
	if m.FactorPrimitiveElement != nil {
		if r.Factor == nil {
			r.Factor = &Decimal{}
		}
		r.Factor.Id = m.FactorPrimitiveElement.Id
		r.Factor.Extension = m.FactorPrimitiveElement.Extension
	}
	r.Net = m.Net
	r.NoteNumber = m.NoteNumber
	for i, e := range m.NoteNumberPrimitiveElement {
		if len(r.NoteNumber) <= i {
			r.NoteNumber = append(r.NoteNumber, PositiveInt{})
		}
		if e != nil {
			r.NoteNumber[i].Id = e.Id
			r.NoteNumber[i].Extension = e.Extension
		}
	}
	r.Adjudication = m.Adjudication
	return nil
}
func (r ClaimResponseAddItemDetailSubDetail) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ProductOrService, xml.StartElement{Name: xml.Name{Local: "productOrService"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Modifier, xml.StartElement{Name: xml.Name{Local: "modifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.UnitPrice, xml.StartElement{Name: xml.Name{Local: "unitPrice"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Factor, xml.StartElement{Name: xml.Name{Local: "factor"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Net, xml.StartElement{Name: xml.Name{Local: "net"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NoteNumber, xml.StartElement{Name: xml.Name{Local: "noteNumber"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Adjudication, xml.StartElement{Name: xml.Name{Local: "adjudication"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ClaimResponseAddItemDetailSubDetail) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "productOrService":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductOrService = v
			case "modifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Modifier = append(r.Modifier, v)
			case "quantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = &v
			case "unitPrice":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.UnitPrice = &v
			case "factor":
				var v Decimal
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Factor = &v
			case "net":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Net = &v
			case "noteNumber":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NoteNumber = append(r.NoteNumber, v)
			case "adjudication":
				var v ClaimResponseItemAdjudication
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjudication = append(r.Adjudication, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ClaimResponseAddItemDetailSubDetail) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Categorized monetary totals for the adjudication.
type ClaimResponseTotal struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A code to indicate the information type of this adjudication record. Information types may include: the value submitted, maximum values or percentages allowed or payable under the plan, amounts that the patient is responsible for in aggregate or pertaining to this item, amounts paid by other coverages, and the benefit payable for this item.
	Category CodeableConcept
	// Monetary total amount associated with the category.
	Amount Money
}
type jsonClaimResponseTotal struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Category          CodeableConcept `json:"category,omitempty"`
	Amount            Money           `json:"amount,omitempty"`
}

func (r ClaimResponseTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimResponseTotal) marshalJSON() jsonClaimResponseTotal {
	m := jsonClaimResponseTotal{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Category = r.Category
	m.Amount = r.Amount
	return m
}
func (r *ClaimResponseTotal) UnmarshalJSON(b []byte) error {
	var m jsonClaimResponseTotal
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimResponseTotal) unmarshalJSON(m jsonClaimResponseTotal) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Category = m.Category
	r.Amount = m.Amount
	return nil
}
func (r ClaimResponseTotal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Category, xml.StartElement{Name: xml.Name{Local: "category"}})
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
func (r *ClaimResponseTotal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "category":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = v
			case "amount":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ClaimResponseTotal) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Payment details for the adjudication of the claim.
type ClaimResponsePayment struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Whether this represents partial or complete payment of the benefits payable.
	Type CodeableConcept
	// Total amount of all adjustments to this payment included in this transaction which are not related to this claim's adjudication.
	Adjustment *Money
	// Reason for the payment adjustment.
	AdjustmentReason *CodeableConcept
	// Estimated date the payment will be issued or the actual issue date of payment.
	Date *Date
	// Benefits payable less any payment adjustment.
	Amount Money
	// Issuer's unique identifier for the payment instrument.
	Identifier *Identifier
}
type jsonClaimResponsePayment struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Type                 CodeableConcept   `json:"type,omitempty"`
	Adjustment           *Money            `json:"adjustment,omitempty"`
	AdjustmentReason     *CodeableConcept  `json:"adjustmentReason,omitempty"`
	Date                 *Date             `json:"date,omitempty"`
	DatePrimitiveElement *primitiveElement `json:"_date,omitempty"`
	Amount               Money             `json:"amount,omitempty"`
	Identifier           *Identifier       `json:"identifier,omitempty"`
}

func (r ClaimResponsePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimResponsePayment) marshalJSON() jsonClaimResponsePayment {
	m := jsonClaimResponsePayment{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Adjustment = r.Adjustment
	m.AdjustmentReason = r.AdjustmentReason
	if r.Date != nil && r.Date.Value != nil {
		m.Date = r.Date
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Amount = r.Amount
	m.Identifier = r.Identifier
	return m
}
func (r *ClaimResponsePayment) UnmarshalJSON(b []byte) error {
	var m jsonClaimResponsePayment
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimResponsePayment) unmarshalJSON(m jsonClaimResponsePayment) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Adjustment = m.Adjustment
	r.AdjustmentReason = m.AdjustmentReason
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		if r.Date == nil {
			r.Date = &Date{}
		}
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Amount = m.Amount
	r.Identifier = m.Identifier
	return nil
}
func (r ClaimResponsePayment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Adjustment, xml.StartElement{Name: xml.Name{Local: "adjustment"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AdjustmentReason, xml.StartElement{Name: xml.Name{Local: "adjustmentReason"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Date, xml.StartElement{Name: xml.Name{Local: "date"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Amount, xml.StartElement{Name: xml.Name{Local: "amount"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Identifier, xml.StartElement{Name: xml.Name{Local: "identifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ClaimResponsePayment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = v
			case "adjustment":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Adjustment = &v
			case "adjustmentReason":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AdjustmentReason = &v
			case "date":
				var v Date
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Date = &v
			case "amount":
				var v Money
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = v
			case "identifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Identifier = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ClaimResponsePayment) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A note that describes or explains adjudication results in a human readable form.
type ClaimResponseProcessNote struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely identify a note entry.
	Number *PositiveInt
	// The business purpose of the note text.
	Type *Code
	// The explanation or description associated with the processing.
	Text String
	// A code to define the language used in the text of the note.
	Language *CodeableConcept
}
type jsonClaimResponseProcessNote struct {
	Id                     *string           `json:"id,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Number                 *PositiveInt      `json:"number,omitempty"`
	NumberPrimitiveElement *primitiveElement `json:"_number,omitempty"`
	Type                   *Code             `json:"type,omitempty"`
	TypePrimitiveElement   *primitiveElement `json:"_type,omitempty"`
	Text                   String            `json:"text,omitempty"`
	TextPrimitiveElement   *primitiveElement `json:"_text,omitempty"`
	Language               *CodeableConcept  `json:"language,omitempty"`
}

func (r ClaimResponseProcessNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimResponseProcessNote) marshalJSON() jsonClaimResponseProcessNote {
	m := jsonClaimResponseProcessNote{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Number != nil && r.Number.Value != nil {
		m.Number = r.Number
	}
	if r.Number != nil && (r.Number.Id != nil || r.Number.Extension != nil) {
		m.NumberPrimitiveElement = &primitiveElement{Id: r.Number.Id, Extension: r.Number.Extension}
	}
	if r.Type != nil && r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.Text.Value != nil {
		m.Text = r.Text
	}
	if r.Text.Id != nil || r.Text.Extension != nil {
		m.TextPrimitiveElement = &primitiveElement{Id: r.Text.Id, Extension: r.Text.Extension}
	}
	m.Language = r.Language
	return m
}
func (r *ClaimResponseProcessNote) UnmarshalJSON(b []byte) error {
	var m jsonClaimResponseProcessNote
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimResponseProcessNote) unmarshalJSON(m jsonClaimResponseProcessNote) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Number = m.Number
	if m.NumberPrimitiveElement != nil {
		if r.Number == nil {
			r.Number = &PositiveInt{}
		}
		r.Number.Id = m.NumberPrimitiveElement.Id
		r.Number.Extension = m.NumberPrimitiveElement.Extension
	}
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
		r.Text.Id = m.TextPrimitiveElement.Id
		r.Text.Extension = m.TextPrimitiveElement.Extension
	}
	r.Language = m.Language
	return nil
}
func (r ClaimResponseProcessNote) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Number, xml.StartElement{Name: xml.Name{Local: "number"}})
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
	err = e.EncodeElement(r.Language, xml.StartElement{Name: xml.Name{Local: "language"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ClaimResponseProcessNote) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "number":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Number = &v
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
				r.Text = v
			case "language":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Language = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ClaimResponseProcessNote) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Financial instruments for reimbursement for the health care products and services specified on the claim.
type ClaimResponseInsurance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A number to uniquely identify insurance entries and provide a sequence of coverages to convey coordination of benefit order.
	Sequence PositiveInt
	// A flag to indicate that this Coverage is to be used for adjudication of this claim when set to true.
	Focal Boolean
	// Reference to the insurance card level information contained in the Coverage resource. The coverage issuing insurer will use these details to locate the patient's actual coverage within the insurer's information system.
	Coverage Reference
	// A business agreement number established between the provider and the insurer for special business processing purposes.
	BusinessArrangement *String
	// The result of the adjudication of the line items for the Coverage specified in this insurance.
	ClaimResponse *Reference
}
type jsonClaimResponseInsurance struct {
	Id                                  *string           `json:"id,omitempty"`
	Extension                           []Extension       `json:"extension,omitempty"`
	ModifierExtension                   []Extension       `json:"modifierExtension,omitempty"`
	Sequence                            PositiveInt       `json:"sequence,omitempty"`
	SequencePrimitiveElement            *primitiveElement `json:"_sequence,omitempty"`
	Focal                               Boolean           `json:"focal,omitempty"`
	FocalPrimitiveElement               *primitiveElement `json:"_focal,omitempty"`
	Coverage                            Reference         `json:"coverage,omitempty"`
	BusinessArrangement                 *String           `json:"businessArrangement,omitempty"`
	BusinessArrangementPrimitiveElement *primitiveElement `json:"_businessArrangement,omitempty"`
	ClaimResponse                       *Reference        `json:"claimResponse,omitempty"`
}

func (r ClaimResponseInsurance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimResponseInsurance) marshalJSON() jsonClaimResponseInsurance {
	m := jsonClaimResponseInsurance{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Sequence.Value != nil {
		m.Sequence = r.Sequence
	}
	if r.Sequence.Id != nil || r.Sequence.Extension != nil {
		m.SequencePrimitiveElement = &primitiveElement{Id: r.Sequence.Id, Extension: r.Sequence.Extension}
	}
	if r.Focal.Value != nil {
		m.Focal = r.Focal
	}
	if r.Focal.Id != nil || r.Focal.Extension != nil {
		m.FocalPrimitiveElement = &primitiveElement{Id: r.Focal.Id, Extension: r.Focal.Extension}
	}
	m.Coverage = r.Coverage
	if r.BusinessArrangement != nil && r.BusinessArrangement.Value != nil {
		m.BusinessArrangement = r.BusinessArrangement
	}
	if r.BusinessArrangement != nil && (r.BusinessArrangement.Id != nil || r.BusinessArrangement.Extension != nil) {
		m.BusinessArrangementPrimitiveElement = &primitiveElement{Id: r.BusinessArrangement.Id, Extension: r.BusinessArrangement.Extension}
	}
	m.ClaimResponse = r.ClaimResponse
	return m
}
func (r *ClaimResponseInsurance) UnmarshalJSON(b []byte) error {
	var m jsonClaimResponseInsurance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimResponseInsurance) unmarshalJSON(m jsonClaimResponseInsurance) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Sequence = m.Sequence
	if m.SequencePrimitiveElement != nil {
		r.Sequence.Id = m.SequencePrimitiveElement.Id
		r.Sequence.Extension = m.SequencePrimitiveElement.Extension
	}
	r.Focal = m.Focal
	if m.FocalPrimitiveElement != nil {
		r.Focal.Id = m.FocalPrimitiveElement.Id
		r.Focal.Extension = m.FocalPrimitiveElement.Extension
	}
	r.Coverage = m.Coverage
	r.BusinessArrangement = m.BusinessArrangement
	if m.BusinessArrangementPrimitiveElement != nil {
		if r.BusinessArrangement == nil {
			r.BusinessArrangement = &String{}
		}
		r.BusinessArrangement.Id = m.BusinessArrangementPrimitiveElement.Id
		r.BusinessArrangement.Extension = m.BusinessArrangementPrimitiveElement.Extension
	}
	r.ClaimResponse = m.ClaimResponse
	return nil
}
func (r ClaimResponseInsurance) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Sequence, xml.StartElement{Name: xml.Name{Local: "sequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Focal, xml.StartElement{Name: xml.Name{Local: "focal"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Coverage, xml.StartElement{Name: xml.Name{Local: "coverage"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BusinessArrangement, xml.StartElement{Name: xml.Name{Local: "businessArrangement"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ClaimResponse, xml.StartElement{Name: xml.Name{Local: "claimResponse"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ClaimResponseInsurance) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "sequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Sequence = v
			case "focal":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Focal = v
			case "coverage":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Coverage = v
			case "businessArrangement":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BusinessArrangement = &v
			case "claimResponse":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ClaimResponse = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ClaimResponseInsurance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Errors encountered during the processing of the adjudication.
type ClaimResponseError struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The sequence number of the line item submitted which contains the error. This value is omitted when the error occurs outside of the item structure.
	ItemSequence *PositiveInt
	// The sequence number of the detail within the line item submitted which contains the error. This value is omitted when the error occurs outside of the item structure.
	DetailSequence *PositiveInt
	// The sequence number of the sub-detail within the detail within the line item submitted which contains the error. This value is omitted when the error occurs outside of the item structure.
	SubDetailSequence *PositiveInt
	// An error code, from a specified code system, which details why the claim could not be adjudicated.
	Code CodeableConcept
}
type jsonClaimResponseError struct {
	Id                                *string           `json:"id,omitempty"`
	Extension                         []Extension       `json:"extension,omitempty"`
	ModifierExtension                 []Extension       `json:"modifierExtension,omitempty"`
	ItemSequence                      *PositiveInt      `json:"itemSequence,omitempty"`
	ItemSequencePrimitiveElement      *primitiveElement `json:"_itemSequence,omitempty"`
	DetailSequence                    *PositiveInt      `json:"detailSequence,omitempty"`
	DetailSequencePrimitiveElement    *primitiveElement `json:"_detailSequence,omitempty"`
	SubDetailSequence                 *PositiveInt      `json:"subDetailSequence,omitempty"`
	SubDetailSequencePrimitiveElement *primitiveElement `json:"_subDetailSequence,omitempty"`
	Code                              CodeableConcept   `json:"code,omitempty"`
}

func (r ClaimResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ClaimResponseError) marshalJSON() jsonClaimResponseError {
	m := jsonClaimResponseError{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.ItemSequence != nil && r.ItemSequence.Value != nil {
		m.ItemSequence = r.ItemSequence
	}
	if r.ItemSequence != nil && (r.ItemSequence.Id != nil || r.ItemSequence.Extension != nil) {
		m.ItemSequencePrimitiveElement = &primitiveElement{Id: r.ItemSequence.Id, Extension: r.ItemSequence.Extension}
	}
	if r.DetailSequence != nil && r.DetailSequence.Value != nil {
		m.DetailSequence = r.DetailSequence
	}
	if r.DetailSequence != nil && (r.DetailSequence.Id != nil || r.DetailSequence.Extension != nil) {
		m.DetailSequencePrimitiveElement = &primitiveElement{Id: r.DetailSequence.Id, Extension: r.DetailSequence.Extension}
	}
	if r.SubDetailSequence != nil && r.SubDetailSequence.Value != nil {
		m.SubDetailSequence = r.SubDetailSequence
	}
	if r.SubDetailSequence != nil && (r.SubDetailSequence.Id != nil || r.SubDetailSequence.Extension != nil) {
		m.SubDetailSequencePrimitiveElement = &primitiveElement{Id: r.SubDetailSequence.Id, Extension: r.SubDetailSequence.Extension}
	}
	m.Code = r.Code
	return m
}
func (r *ClaimResponseError) UnmarshalJSON(b []byte) error {
	var m jsonClaimResponseError
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ClaimResponseError) unmarshalJSON(m jsonClaimResponseError) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ItemSequence = m.ItemSequence
	if m.ItemSequencePrimitiveElement != nil {
		if r.ItemSequence == nil {
			r.ItemSequence = &PositiveInt{}
		}
		r.ItemSequence.Id = m.ItemSequencePrimitiveElement.Id
		r.ItemSequence.Extension = m.ItemSequencePrimitiveElement.Extension
	}
	r.DetailSequence = m.DetailSequence
	if m.DetailSequencePrimitiveElement != nil {
		if r.DetailSequence == nil {
			r.DetailSequence = &PositiveInt{}
		}
		r.DetailSequence.Id = m.DetailSequencePrimitiveElement.Id
		r.DetailSequence.Extension = m.DetailSequencePrimitiveElement.Extension
	}
	r.SubDetailSequence = m.SubDetailSequence
	if m.SubDetailSequencePrimitiveElement != nil {
		if r.SubDetailSequence == nil {
			r.SubDetailSequence = &PositiveInt{}
		}
		r.SubDetailSequence.Id = m.SubDetailSequencePrimitiveElement.Id
		r.SubDetailSequence.Extension = m.SubDetailSequencePrimitiveElement.Extension
	}
	r.Code = m.Code
	return nil
}
func (r ClaimResponseError) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ItemSequence, xml.StartElement{Name: xml.Name{Local: "itemSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DetailSequence, xml.StartElement{Name: xml.Name{Local: "detailSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SubDetailSequence, xml.StartElement{Name: xml.Name{Local: "subDetailSequence"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *ClaimResponseError) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "itemSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ItemSequence = &v
			case "detailSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DetailSequence = &v
			case "subDetailSequence":
				var v PositiveInt
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SubDetailSequence = &v
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r ClaimResponseError) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
