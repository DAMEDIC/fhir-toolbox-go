package r4

import (
	"encoding/json"
	"encoding/xml"
	model "fhir-toolbox/model"
	"fmt"
)

// The details of a healthcare service available at a location.
type HealthcareService struct {
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
	// External identifiers for this item.
	Identifier []Identifier
	// This flag is used to mark the record to not be used. This is not used when a center is closed for maintenance, or for holidays, the notAvailable period is to be used for this.
	Active *Boolean
	// The organization that provides this healthcare service.
	ProvidedBy *Reference
	// Identifies the broad category of service being performed or delivered.
	Category []CodeableConcept
	// The specific type of service that may be delivered or performed.
	Type []CodeableConcept
	// Collection of specialties handled by the service site. This is more of a medical term.
	Specialty []CodeableConcept
	// The location(s) where this healthcare service may be provided.
	Location []Reference
	// Further description of the service as it would be presented to a consumer while searching.
	Name *String
	// Any additional description of the service and/or any specific issues not covered by the other attributes, which can be displayed as further detail under the serviceName.
	Comment *String
	// Extra details about the service that can't be placed in the other fields.
	ExtraDetails *Markdown
	// If there is a photo/symbol associated with this HealthcareService, it may be included here to facilitate quick identification of the service in a list.
	Photo *Attachment
	// List of contacts related to this specific healthcare service.
	Telecom []ContactPoint
	// The location(s) that this service is available to (not where the service is provided).
	CoverageArea []Reference
	// The code(s) that detail the conditions under which the healthcare service is available/offered.
	ServiceProvisionCode []CodeableConcept
	// Does this service have specific eligibility requirements that need to be met in order to use the service?
	Eligibility []HealthcareServiceEligibility
	// Programs that this service is applicable to.
	Program []CodeableConcept
	// Collection of characteristics (attributes).
	Characteristic []CodeableConcept
	// Some services are specifically made available in multiple languages, this property permits a directory to declare the languages this is offered in. Typically this is only provided where a service operates in communities with mixed languages used.
	Communication []CodeableConcept
	// Ways that the service accepts referrals, if this is not provided then it is implied that no referral is required.
	ReferralMethod []CodeableConcept
	// Indicates whether or not a prospective consumer will require an appointment for a particular service at a site to be provided by the Organization. Indicates if an appointment is required for access to this service.
	AppointmentRequired *Boolean
	// A collection of times that the Service Site is available.
	AvailableTime []HealthcareServiceAvailableTime
	// The HealthcareService is not available during this period of time due to the provided reason.
	NotAvailable []HealthcareServiceNotAvailable
	// A description of site availability exceptions, e.g. public holiday availability. Succinctly describing all possible exceptions to normal site availability as details in the available Times and not available Times.
	AvailabilityExceptions *String
	// Technical endpoints providing access to services operated for the specific healthcare services defined at this resource.
	Endpoint []Reference
}

func (r HealthcareService) ResourceType() string {
	return "HealthcareService"
}
func (r HealthcareService) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonHealthcareService struct {
	ResourceType                           string                           `json:"resourceType"`
	Id                                     *Id                              `json:"id,omitempty"`
	IdPrimitiveElement                     *primitiveElement                `json:"_id,omitempty"`
	Meta                                   *Meta                            `json:"meta,omitempty"`
	ImplicitRules                          *Uri                             `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement          *primitiveElement                `json:"_implicitRules,omitempty"`
	Language                               *Code                            `json:"language,omitempty"`
	LanguagePrimitiveElement               *primitiveElement                `json:"_language,omitempty"`
	Text                                   *Narrative                       `json:"text,omitempty"`
	Contained                              []ContainedResource              `json:"contained,omitempty"`
	Extension                              []Extension                      `json:"extension,omitempty"`
	ModifierExtension                      []Extension                      `json:"modifierExtension,omitempty"`
	Identifier                             []Identifier                     `json:"identifier,omitempty"`
	Active                                 *Boolean                         `json:"active,omitempty"`
	ActivePrimitiveElement                 *primitiveElement                `json:"_active,omitempty"`
	ProvidedBy                             *Reference                       `json:"providedBy,omitempty"`
	Category                               []CodeableConcept                `json:"category,omitempty"`
	Type                                   []CodeableConcept                `json:"type,omitempty"`
	Specialty                              []CodeableConcept                `json:"specialty,omitempty"`
	Location                               []Reference                      `json:"location,omitempty"`
	Name                                   *String                          `json:"name,omitempty"`
	NamePrimitiveElement                   *primitiveElement                `json:"_name,omitempty"`
	Comment                                *String                          `json:"comment,omitempty"`
	CommentPrimitiveElement                *primitiveElement                `json:"_comment,omitempty"`
	ExtraDetails                           *Markdown                        `json:"extraDetails,omitempty"`
	ExtraDetailsPrimitiveElement           *primitiveElement                `json:"_extraDetails,omitempty"`
	Photo                                  *Attachment                      `json:"photo,omitempty"`
	Telecom                                []ContactPoint                   `json:"telecom,omitempty"`
	CoverageArea                           []Reference                      `json:"coverageArea,omitempty"`
	ServiceProvisionCode                   []CodeableConcept                `json:"serviceProvisionCode,omitempty"`
	Eligibility                            []HealthcareServiceEligibility   `json:"eligibility,omitempty"`
	Program                                []CodeableConcept                `json:"program,omitempty"`
	Characteristic                         []CodeableConcept                `json:"characteristic,omitempty"`
	Communication                          []CodeableConcept                `json:"communication,omitempty"`
	ReferralMethod                         []CodeableConcept                `json:"referralMethod,omitempty"`
	AppointmentRequired                    *Boolean                         `json:"appointmentRequired,omitempty"`
	AppointmentRequiredPrimitiveElement    *primitiveElement                `json:"_appointmentRequired,omitempty"`
	AvailableTime                          []HealthcareServiceAvailableTime `json:"availableTime,omitempty"`
	NotAvailable                           []HealthcareServiceNotAvailable  `json:"notAvailable,omitempty"`
	AvailabilityExceptions                 *String                          `json:"availabilityExceptions,omitempty"`
	AvailabilityExceptionsPrimitiveElement *primitiveElement                `json:"_availabilityExceptions,omitempty"`
	Endpoint                               []Reference                      `json:"endpoint,omitempty"`
}

func (r HealthcareService) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r HealthcareService) marshalJSON() jsonHealthcareService {
	m := jsonHealthcareService{}
	m.ResourceType = "HealthcareService"
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
	if r.Active != nil && r.Active.Value != nil {
		m.Active = r.Active
	}
	if r.Active != nil && (r.Active.Id != nil || r.Active.Extension != nil) {
		m.ActivePrimitiveElement = &primitiveElement{Id: r.Active.Id, Extension: r.Active.Extension}
	}
	m.ProvidedBy = r.ProvidedBy
	m.Category = r.Category
	m.Type = r.Type
	m.Specialty = r.Specialty
	m.Location = r.Location
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Comment != nil && r.Comment.Value != nil {
		m.Comment = r.Comment
	}
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	if r.ExtraDetails != nil && r.ExtraDetails.Value != nil {
		m.ExtraDetails = r.ExtraDetails
	}
	if r.ExtraDetails != nil && (r.ExtraDetails.Id != nil || r.ExtraDetails.Extension != nil) {
		m.ExtraDetailsPrimitiveElement = &primitiveElement{Id: r.ExtraDetails.Id, Extension: r.ExtraDetails.Extension}
	}
	m.Photo = r.Photo
	m.Telecom = r.Telecom
	m.CoverageArea = r.CoverageArea
	m.ServiceProvisionCode = r.ServiceProvisionCode
	m.Eligibility = r.Eligibility
	m.Program = r.Program
	m.Characteristic = r.Characteristic
	m.Communication = r.Communication
	m.ReferralMethod = r.ReferralMethod
	if r.AppointmentRequired != nil && r.AppointmentRequired.Value != nil {
		m.AppointmentRequired = r.AppointmentRequired
	}
	if r.AppointmentRequired != nil && (r.AppointmentRequired.Id != nil || r.AppointmentRequired.Extension != nil) {
		m.AppointmentRequiredPrimitiveElement = &primitiveElement{Id: r.AppointmentRequired.Id, Extension: r.AppointmentRequired.Extension}
	}
	m.AvailableTime = r.AvailableTime
	m.NotAvailable = r.NotAvailable
	if r.AvailabilityExceptions != nil && r.AvailabilityExceptions.Value != nil {
		m.AvailabilityExceptions = r.AvailabilityExceptions
	}
	if r.AvailabilityExceptions != nil && (r.AvailabilityExceptions.Id != nil || r.AvailabilityExceptions.Extension != nil) {
		m.AvailabilityExceptionsPrimitiveElement = &primitiveElement{Id: r.AvailabilityExceptions.Id, Extension: r.AvailabilityExceptions.Extension}
	}
	m.Endpoint = r.Endpoint
	return m
}
func (r *HealthcareService) UnmarshalJSON(b []byte) error {
	var m jsonHealthcareService
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *HealthcareService) unmarshalJSON(m jsonHealthcareService) error {
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
	r.Active = m.Active
	if m.ActivePrimitiveElement != nil {
		if r.Active == nil {
			r.Active = &Boolean{}
		}
		r.Active.Id = m.ActivePrimitiveElement.Id
		r.Active.Extension = m.ActivePrimitiveElement.Extension
	}
	r.ProvidedBy = m.ProvidedBy
	r.Category = m.Category
	r.Type = m.Type
	r.Specialty = m.Specialty
	r.Location = m.Location
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		if r.Comment == nil {
			r.Comment = &String{}
		}
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	r.ExtraDetails = m.ExtraDetails
	if m.ExtraDetailsPrimitiveElement != nil {
		if r.ExtraDetails == nil {
			r.ExtraDetails = &Markdown{}
		}
		r.ExtraDetails.Id = m.ExtraDetailsPrimitiveElement.Id
		r.ExtraDetails.Extension = m.ExtraDetailsPrimitiveElement.Extension
	}
	r.Photo = m.Photo
	r.Telecom = m.Telecom
	r.CoverageArea = m.CoverageArea
	r.ServiceProvisionCode = m.ServiceProvisionCode
	r.Eligibility = m.Eligibility
	r.Program = m.Program
	r.Characteristic = m.Characteristic
	r.Communication = m.Communication
	r.ReferralMethod = m.ReferralMethod
	r.AppointmentRequired = m.AppointmentRequired
	if m.AppointmentRequiredPrimitiveElement != nil {
		if r.AppointmentRequired == nil {
			r.AppointmentRequired = &Boolean{}
		}
		r.AppointmentRequired.Id = m.AppointmentRequiredPrimitiveElement.Id
		r.AppointmentRequired.Extension = m.AppointmentRequiredPrimitiveElement.Extension
	}
	r.AvailableTime = m.AvailableTime
	r.NotAvailable = m.NotAvailable
	r.AvailabilityExceptions = m.AvailabilityExceptions
	if m.AvailabilityExceptionsPrimitiveElement != nil {
		if r.AvailabilityExceptions == nil {
			r.AvailabilityExceptions = &String{}
		}
		r.AvailabilityExceptions.Id = m.AvailabilityExceptionsPrimitiveElement.Id
		r.AvailabilityExceptions.Extension = m.AvailabilityExceptionsPrimitiveElement.Extension
	}
	r.Endpoint = m.Endpoint
	return nil
}
func (r HealthcareService) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "HealthcareService"
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
	err = e.EncodeElement(r.Active, xml.StartElement{Name: xml.Name{Local: "active"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ProvidedBy, xml.StartElement{Name: xml.Name{Local: "providedBy"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Category, xml.StartElement{Name: xml.Name{Local: "category"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Type, xml.StartElement{Name: xml.Name{Local: "type"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Specialty, xml.StartElement{Name: xml.Name{Local: "specialty"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Location, xml.StartElement{Name: xml.Name{Local: "location"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Name, xml.StartElement{Name: xml.Name{Local: "name"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Comment, xml.StartElement{Name: xml.Name{Local: "comment"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ExtraDetails, xml.StartElement{Name: xml.Name{Local: "extraDetails"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Photo, xml.StartElement{Name: xml.Name{Local: "photo"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Telecom, xml.StartElement{Name: xml.Name{Local: "telecom"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CoverageArea, xml.StartElement{Name: xml.Name{Local: "coverageArea"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ServiceProvisionCode, xml.StartElement{Name: xml.Name{Local: "serviceProvisionCode"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Eligibility, xml.StartElement{Name: xml.Name{Local: "eligibility"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Program, xml.StartElement{Name: xml.Name{Local: "program"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Characteristic, xml.StartElement{Name: xml.Name{Local: "characteristic"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Communication, xml.StartElement{Name: xml.Name{Local: "communication"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ReferralMethod, xml.StartElement{Name: xml.Name{Local: "referralMethod"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AppointmentRequired, xml.StartElement{Name: xml.Name{Local: "appointmentRequired"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AvailableTime, xml.StartElement{Name: xml.Name{Local: "availableTime"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.NotAvailable, xml.StartElement{Name: xml.Name{Local: "notAvailable"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AvailabilityExceptions, xml.StartElement{Name: xml.Name{Local: "availabilityExceptions"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Endpoint, xml.StartElement{Name: xml.Name{Local: "endpoint"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *HealthcareService) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "active":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Active = &v
			case "providedBy":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProvidedBy = &v
			case "category":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Category = append(r.Category, v)
			case "type":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = append(r.Type, v)
			case "specialty":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Specialty = append(r.Specialty, v)
			case "location":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Location = append(r.Location, v)
			case "name":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Name = &v
			case "comment":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Comment = &v
			case "extraDetails":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ExtraDetails = &v
			case "photo":
				var v Attachment
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Photo = &v
			case "telecom":
				var v ContactPoint
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Telecom = append(r.Telecom, v)
			case "coverageArea":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CoverageArea = append(r.CoverageArea, v)
			case "serviceProvisionCode":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ServiceProvisionCode = append(r.ServiceProvisionCode, v)
			case "eligibility":
				var v HealthcareServiceEligibility
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Eligibility = append(r.Eligibility, v)
			case "program":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Program = append(r.Program, v)
			case "characteristic":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Characteristic = append(r.Characteristic, v)
			case "communication":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Communication = append(r.Communication, v)
			case "referralMethod":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ReferralMethod = append(r.ReferralMethod, v)
			case "appointmentRequired":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AppointmentRequired = &v
			case "availableTime":
				var v HealthcareServiceAvailableTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AvailableTime = append(r.AvailableTime, v)
			case "notAvailable":
				var v HealthcareServiceNotAvailable
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.NotAvailable = append(r.NotAvailable, v)
			case "availabilityExceptions":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AvailabilityExceptions = &v
			case "endpoint":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Endpoint = append(r.Endpoint, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r HealthcareService) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// Does this service have specific eligibility requirements that need to be met in order to use the service?
type HealthcareServiceEligibility struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Coded value for the eligibility.
	Code *CodeableConcept
	// Describes the eligibility conditions for the service.
	Comment *Markdown
}
type jsonHealthcareServiceEligibility struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Code                    *CodeableConcept  `json:"code,omitempty"`
	Comment                 *Markdown         `json:"comment,omitempty"`
	CommentPrimitiveElement *primitiveElement `json:"_comment,omitempty"`
}

func (r HealthcareServiceEligibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r HealthcareServiceEligibility) marshalJSON() jsonHealthcareServiceEligibility {
	m := jsonHealthcareServiceEligibility{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	if r.Comment != nil && r.Comment.Value != nil {
		m.Comment = r.Comment
	}
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	return m
}
func (r *HealthcareServiceEligibility) UnmarshalJSON(b []byte) error {
	var m jsonHealthcareServiceEligibility
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *HealthcareServiceEligibility) unmarshalJSON(m jsonHealthcareServiceEligibility) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		if r.Comment == nil {
			r.Comment = &Markdown{}
		}
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	return nil
}
func (r HealthcareServiceEligibility) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Code, xml.StartElement{Name: xml.Name{Local: "code"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Comment, xml.StartElement{Name: xml.Name{Local: "comment"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *HealthcareServiceEligibility) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "code":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Code = &v
			case "comment":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Comment = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r HealthcareServiceEligibility) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// A collection of times that the Service Site is available.
type HealthcareServiceAvailableTime struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates which days of the week are available between the start and end Times.
	DaysOfWeek []Code
	// Is this always available? (hence times are irrelevant) e.g. 24 hour service.
	AllDay *Boolean
	// The opening time of day. Note: If the AllDay flag is set, then this time is ignored.
	AvailableStartTime *Time
	// The closing time of day. Note: If the AllDay flag is set, then this time is ignored.
	AvailableEndTime *Time
}
type jsonHealthcareServiceAvailableTime struct {
	Id                                 *string             `json:"id,omitempty"`
	Extension                          []Extension         `json:"extension,omitempty"`
	ModifierExtension                  []Extension         `json:"modifierExtension,omitempty"`
	DaysOfWeek                         []Code              `json:"daysOfWeek,omitempty"`
	DaysOfWeekPrimitiveElement         []*primitiveElement `json:"_daysOfWeek,omitempty"`
	AllDay                             *Boolean            `json:"allDay,omitempty"`
	AllDayPrimitiveElement             *primitiveElement   `json:"_allDay,omitempty"`
	AvailableStartTime                 *Time               `json:"availableStartTime,omitempty"`
	AvailableStartTimePrimitiveElement *primitiveElement   `json:"_availableStartTime,omitempty"`
	AvailableEndTime                   *Time               `json:"availableEndTime,omitempty"`
	AvailableEndTimePrimitiveElement   *primitiveElement   `json:"_availableEndTime,omitempty"`
}

func (r HealthcareServiceAvailableTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r HealthcareServiceAvailableTime) marshalJSON() jsonHealthcareServiceAvailableTime {
	m := jsonHealthcareServiceAvailableTime{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	anyDaysOfWeekValue := false
	for _, e := range r.DaysOfWeek {
		if e.Value != nil {
			anyDaysOfWeekValue = true
			break
		}
	}
	if anyDaysOfWeekValue {
		m.DaysOfWeek = r.DaysOfWeek
	}
	anyDaysOfWeekIdOrExtension := false
	for _, e := range r.DaysOfWeek {
		if e.Id != nil || e.Extension != nil {
			anyDaysOfWeekIdOrExtension = true
			break
		}
	}
	if anyDaysOfWeekIdOrExtension {
		m.DaysOfWeekPrimitiveElement = make([]*primitiveElement, 0, len(r.DaysOfWeek))
		for _, e := range r.DaysOfWeek {
			if e.Id != nil || e.Extension != nil {
				m.DaysOfWeekPrimitiveElement = append(m.DaysOfWeekPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.DaysOfWeekPrimitiveElement = append(m.DaysOfWeekPrimitiveElement, nil)
			}
		}
	}
	if r.AllDay != nil && r.AllDay.Value != nil {
		m.AllDay = r.AllDay
	}
	if r.AllDay != nil && (r.AllDay.Id != nil || r.AllDay.Extension != nil) {
		m.AllDayPrimitiveElement = &primitiveElement{Id: r.AllDay.Id, Extension: r.AllDay.Extension}
	}
	if r.AvailableStartTime != nil && r.AvailableStartTime.Value != nil {
		m.AvailableStartTime = r.AvailableStartTime
	}
	if r.AvailableStartTime != nil && (r.AvailableStartTime.Id != nil || r.AvailableStartTime.Extension != nil) {
		m.AvailableStartTimePrimitiveElement = &primitiveElement{Id: r.AvailableStartTime.Id, Extension: r.AvailableStartTime.Extension}
	}
	if r.AvailableEndTime != nil && r.AvailableEndTime.Value != nil {
		m.AvailableEndTime = r.AvailableEndTime
	}
	if r.AvailableEndTime != nil && (r.AvailableEndTime.Id != nil || r.AvailableEndTime.Extension != nil) {
		m.AvailableEndTimePrimitiveElement = &primitiveElement{Id: r.AvailableEndTime.Id, Extension: r.AvailableEndTime.Extension}
	}
	return m
}
func (r *HealthcareServiceAvailableTime) UnmarshalJSON(b []byte) error {
	var m jsonHealthcareServiceAvailableTime
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *HealthcareServiceAvailableTime) unmarshalJSON(m jsonHealthcareServiceAvailableTime) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.DaysOfWeek = m.DaysOfWeek
	for i, e := range m.DaysOfWeekPrimitiveElement {
		if len(r.DaysOfWeek) <= i {
			r.DaysOfWeek = append(r.DaysOfWeek, Code{})
		}
		if e != nil {
			r.DaysOfWeek[i].Id = e.Id
			r.DaysOfWeek[i].Extension = e.Extension
		}
	}
	r.AllDay = m.AllDay
	if m.AllDayPrimitiveElement != nil {
		if r.AllDay == nil {
			r.AllDay = &Boolean{}
		}
		r.AllDay.Id = m.AllDayPrimitiveElement.Id
		r.AllDay.Extension = m.AllDayPrimitiveElement.Extension
	}
	r.AvailableStartTime = m.AvailableStartTime
	if m.AvailableStartTimePrimitiveElement != nil {
		if r.AvailableStartTime == nil {
			r.AvailableStartTime = &Time{}
		}
		r.AvailableStartTime.Id = m.AvailableStartTimePrimitiveElement.Id
		r.AvailableStartTime.Extension = m.AvailableStartTimePrimitiveElement.Extension
	}
	r.AvailableEndTime = m.AvailableEndTime
	if m.AvailableEndTimePrimitiveElement != nil {
		if r.AvailableEndTime == nil {
			r.AvailableEndTime = &Time{}
		}
		r.AvailableEndTime.Id = m.AvailableEndTimePrimitiveElement.Id
		r.AvailableEndTime.Extension = m.AvailableEndTimePrimitiveElement.Extension
	}
	return nil
}
func (r HealthcareServiceAvailableTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.DaysOfWeek, xml.StartElement{Name: xml.Name{Local: "daysOfWeek"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AllDay, xml.StartElement{Name: xml.Name{Local: "allDay"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AvailableStartTime, xml.StartElement{Name: xml.Name{Local: "availableStartTime"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AvailableEndTime, xml.StartElement{Name: xml.Name{Local: "availableEndTime"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *HealthcareServiceAvailableTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "daysOfWeek":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DaysOfWeek = append(r.DaysOfWeek, v)
			case "allDay":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AllDay = &v
			case "availableStartTime":
				var v Time
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AvailableStartTime = &v
			case "availableEndTime":
				var v Time
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AvailableEndTime = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r HealthcareServiceAvailableTime) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}

// The HealthcareService is not available during this period of time due to the provided reason.
type HealthcareServiceNotAvailable struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The reason that can be presented to the user as to why this time is not available.
	Description String
	// Service is not available (seasonally or for a public holiday) from this date.
	During *Period
}
type jsonHealthcareServiceNotAvailable struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Description                 String            `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
	During                      *Period           `json:"during,omitempty"`
}

func (r HealthcareServiceNotAvailable) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r HealthcareServiceNotAvailable) marshalJSON() jsonHealthcareServiceNotAvailable {
	m := jsonHealthcareServiceNotAvailable{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description.Id != nil || r.Description.Extension != nil {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.During = r.During
	return m
}
func (r *HealthcareServiceNotAvailable) UnmarshalJSON(b []byte) error {
	var m jsonHealthcareServiceNotAvailable
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *HealthcareServiceNotAvailable) unmarshalJSON(m jsonHealthcareServiceNotAvailable) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.During = m.During
	return nil
}
func (r HealthcareServiceNotAvailable) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Description, xml.StartElement{Name: xml.Name{Local: "description"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.During, xml.StartElement{Name: xml.Name{Local: "during"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *HealthcareServiceNotAvailable) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "description":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Description = v
			case "during":
				var v Period
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.During = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r HealthcareServiceNotAvailable) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
