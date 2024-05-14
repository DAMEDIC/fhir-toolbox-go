package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// This resource provides enrollment and plan details from the processing of an EnrollmentRequest resource.
type EnrollmentResponse struct {
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
	// The Response business identifier.
	Identifier []Identifier
	// The status of the resource instance.
	Status *Code
	// Original request resource reference.
	Request *Reference
	// Processing status: error, complete.
	Outcome *Code
	// A description of the status of the adjudication.
	Disposition *String
	// The date when the enclosed suite of services were performed or completed.
	Created *DateTime
	// The Insurer who produced this adjudicated response.
	Organization *Reference
	// The practitioner who is responsible for the services rendered to the patient.
	RequestProvider *Reference
}

func (r EnrollmentResponse) ResourceType() string {
	return "EnrollmentResponse"
}

type jsonEnrollmentResponse struct {
	ResourceType                  string              `json:"resourceType"`
	Id                            *Id                 `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement   `json:"_id,omitempty"`
	Meta                          *Meta               `json:"meta,omitempty"`
	ImplicitRules                 *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                      *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement   `json:"_language,omitempty"`
	Text                          *Narrative          `json:"text,omitempty"`
	Contained                     []containedResource `json:"contained,omitempty"`
	Extension                     []Extension         `json:"extension,omitempty"`
	ModifierExtension             []Extension         `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier        `json:"identifier,omitempty"`
	Status                        *Code               `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement   `json:"_status,omitempty"`
	Request                       *Reference          `json:"request,omitempty"`
	Outcome                       *Code               `json:"outcome,omitempty"`
	OutcomePrimitiveElement       *primitiveElement   `json:"_outcome,omitempty"`
	Disposition                   *String             `json:"disposition,omitempty"`
	DispositionPrimitiveElement   *primitiveElement   `json:"_disposition,omitempty"`
	Created                       *DateTime           `json:"created,omitempty"`
	CreatedPrimitiveElement       *primitiveElement   `json:"_created,omitempty"`
	Organization                  *Reference          `json:"organization,omitempty"`
	RequestProvider               *Reference          `json:"requestProvider,omitempty"`
}

func (r EnrollmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r EnrollmentResponse) marshalJSON() jsonEnrollmentResponse {
	m := jsonEnrollmentResponse{}
	m.ResourceType = "EnrollmentResponse"
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
	if r.Status != nil && (r.Status.Id != nil || r.Status.Extension != nil) {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Request = r.Request
	m.Outcome = r.Outcome
	if r.Outcome != nil && (r.Outcome.Id != nil || r.Outcome.Extension != nil) {
		m.OutcomePrimitiveElement = &primitiveElement{Id: r.Outcome.Id, Extension: r.Outcome.Extension}
	}
	m.Disposition = r.Disposition
	if r.Disposition != nil && (r.Disposition.Id != nil || r.Disposition.Extension != nil) {
		m.DispositionPrimitiveElement = &primitiveElement{Id: r.Disposition.Id, Extension: r.Disposition.Extension}
	}
	m.Created = r.Created
	if r.Created != nil && (r.Created.Id != nil || r.Created.Extension != nil) {
		m.CreatedPrimitiveElement = &primitiveElement{Id: r.Created.Id, Extension: r.Created.Extension}
	}
	m.Organization = r.Organization
	m.RequestProvider = r.RequestProvider
	return m
}
func (r *EnrollmentResponse) UnmarshalJSON(b []byte) error {
	var m jsonEnrollmentResponse
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *EnrollmentResponse) unmarshalJSON(m jsonEnrollmentResponse) error {
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
	r.Request = m.Request
	r.Outcome = m.Outcome
	if m.OutcomePrimitiveElement != nil {
		r.Outcome.Id = m.OutcomePrimitiveElement.Id
		r.Outcome.Extension = m.OutcomePrimitiveElement.Extension
	}
	r.Disposition = m.Disposition
	if m.DispositionPrimitiveElement != nil {
		r.Disposition.Id = m.DispositionPrimitiveElement.Id
		r.Disposition.Extension = m.DispositionPrimitiveElement.Extension
	}
	r.Created = m.Created
	if m.CreatedPrimitiveElement != nil {
		r.Created.Id = m.CreatedPrimitiveElement.Id
		r.Created.Extension = m.CreatedPrimitiveElement.Extension
	}
	r.Organization = m.Organization
	r.RequestProvider = m.RequestProvider
	return nil
}
func (r EnrollmentResponse) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
