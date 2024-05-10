package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A guidance response is the formal response to a guidance request, including any output parameters returned by the evaluation, as well as the description of any proposed actions to be taken.
//
// The GuidanceResponse resource supports recording the results of decision support interactions, reportability determination for public health, as well as the communication of additional data requirements for subsequent interactions.
type GuidanceResponse struct {
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
	// The identifier of the request associated with this response. If an identifier was given as part of the request, it will be reproduced here to enable the requester to more easily identify the response in a multi-request scenario.
	RequestIdentifier *Identifier
	// Allows a service to provide  unique, business identifiers for the response.
	Identifier []Identifier
	// An identifier, CodeableConcept or canonical reference to the guidance that was requested.
	Module isGuidanceResponseModule
	// The status of the response. If the evaluation is completed successfully, the status will indicate success. However, in order to complete the evaluation, the engine may require more information. In this case, the status will be data-required, and the response will contain a description of the additional required information. If the evaluation completed successfully, but the engine determines that a potentially more accurate response could be provided if more data was available, the status will be data-requested, and the response will contain a description of the additional requested information.
	Status Code
	// The patient for which the request was processed.
	Subject *Reference
	// The encounter during which this response was created or to which the creation of this record is tightly associated.
	Encounter *Reference
	// Indicates when the guidance response was processed.
	OccurrenceDateTime *DateTime
	// Provides a reference to the device that performed the guidance.
	Performer *Reference
	// Describes the reason for the guidance response in coded or textual form.
	ReasonCode []CodeableConcept
	// Indicates the reason the request was initiated. This is typically provided as a parameter to the evaluation and echoed by the service, although for some use cases, such as subscription- or event-based scenarios, it may provide an indication of the cause for the response.
	ReasonReference []Reference
	// Provides a mechanism to communicate additional information about the response.
	Note []Annotation
	// Messages resulting from the evaluation of the artifact or artifacts. As part of evaluating the request, the engine may produce informational or warning messages. These messages will be provided by this element.
	EvaluationMessage []Reference
	// The output parameters of the evaluation, if any. Many modules will result in the return of specific resources such as procedure or communication requests that are returned as part of the operation result. However, modules may define specific outputs that would be returned as the result of the evaluation, and these would be returned in this element.
	OutputParameters *Reference
	// The actions, if any, produced by the evaluation of the artifact.
	Result *Reference
	// If the evaluation could not be completed due to lack of information, or additional information would potentially result in a more accurate response, this element will a description of the data required in order to proceed with the evaluation. A subsequent request to the service should include this data.
	DataRequirement []DataRequirement
}
type isGuidanceResponseModule interface {
	isGuidanceResponseModule()
}

func (r Uri) isGuidanceResponseModule()             {}
func (r Canonical) isGuidanceResponseModule()       {}
func (r CodeableConcept) isGuidanceResponseModule() {}

type jsonGuidanceResponse struct {
	ResourceType                       string              `json:"resourceType"`
	Id                                 *Id                 `json:"id,omitempty"`
	IdPrimitiveElement                 *primitiveElement   `json:"_id,omitempty"`
	Meta                               *Meta               `json:"meta,omitempty"`
	ImplicitRules                      *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement      *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                           *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement           *primitiveElement   `json:"_language,omitempty"`
	Text                               *Narrative          `json:"text,omitempty"`
	Contained                          []containedResource `json:"contained,omitempty"`
	Extension                          []Extension         `json:"extension,omitempty"`
	ModifierExtension                  []Extension         `json:"modifierExtension,omitempty"`
	RequestIdentifier                  *Identifier         `json:"requestIdentifier,omitempty"`
	Identifier                         []Identifier        `json:"identifier,omitempty"`
	ModuleUri                          *Uri                `json:"moduleUri,omitempty"`
	ModuleUriPrimitiveElement          *primitiveElement   `json:"_moduleUri,omitempty"`
	ModuleCanonical                    *Canonical          `json:"moduleCanonical,omitempty"`
	ModuleCanonicalPrimitiveElement    *primitiveElement   `json:"_moduleCanonical,omitempty"`
	ModuleCodeableConcept              *CodeableConcept    `json:"moduleCodeableConcept,omitempty"`
	Status                             Code                `json:"status,omitempty"`
	StatusPrimitiveElement             *primitiveElement   `json:"_status,omitempty"`
	Subject                            *Reference          `json:"subject,omitempty"`
	Encounter                          *Reference          `json:"encounter,omitempty"`
	OccurrenceDateTime                 *DateTime           `json:"occurrenceDateTime,omitempty"`
	OccurrenceDateTimePrimitiveElement *primitiveElement   `json:"_occurrenceDateTime,omitempty"`
	Performer                          *Reference          `json:"performer,omitempty"`
	ReasonCode                         []CodeableConcept   `json:"reasonCode,omitempty"`
	ReasonReference                    []Reference         `json:"reasonReference,omitempty"`
	Note                               []Annotation        `json:"note,omitempty"`
	EvaluationMessage                  []Reference         `json:"evaluationMessage,omitempty"`
	OutputParameters                   *Reference          `json:"outputParameters,omitempty"`
	Result                             *Reference          `json:"result,omitempty"`
	DataRequirement                    []DataRequirement   `json:"dataRequirement,omitempty"`
}

func (r GuidanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r GuidanceResponse) marshalJSON() jsonGuidanceResponse {
	m := jsonGuidanceResponse{}
	m.ResourceType = "GuidanceResponse"
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
	m.RequestIdentifier = r.RequestIdentifier
	m.Identifier = r.Identifier
	switch v := r.Module.(type) {
	case Uri:
		m.ModuleUri = &v
		if v.Id != nil || v.Extension != nil {
			m.ModuleUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uri:
		m.ModuleUri = v
		if v.Id != nil || v.Extension != nil {
			m.ModuleUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Canonical:
		m.ModuleCanonical = &v
		if v.Id != nil || v.Extension != nil {
			m.ModuleCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Canonical:
		m.ModuleCanonical = v
		if v.Id != nil || v.Extension != nil {
			m.ModuleCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case CodeableConcept:
		m.ModuleCodeableConcept = &v
	case *CodeableConcept:
		m.ModuleCodeableConcept = v
	}
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	m.OccurrenceDateTime = r.OccurrenceDateTime
	if r.OccurrenceDateTime != nil && (r.OccurrenceDateTime.Id != nil || r.OccurrenceDateTime.Extension != nil) {
		m.OccurrenceDateTimePrimitiveElement = &primitiveElement{Id: r.OccurrenceDateTime.Id, Extension: r.OccurrenceDateTime.Extension}
	}
	m.Performer = r.Performer
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Note = r.Note
	m.EvaluationMessage = r.EvaluationMessage
	m.OutputParameters = r.OutputParameters
	m.Result = r.Result
	m.DataRequirement = r.DataRequirement
	return m
}
func (r *GuidanceResponse) UnmarshalJSON(b []byte) error {
	var m jsonGuidanceResponse
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *GuidanceResponse) unmarshalJSON(m jsonGuidanceResponse) error {
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
	r.RequestIdentifier = m.RequestIdentifier
	r.Identifier = m.Identifier
	if m.ModuleUri != nil || m.ModuleUriPrimitiveElement != nil {
		if r.Module != nil {
			return fmt.Errorf("multiple values for field \"Module\"")
		}
		v := m.ModuleUri
		if m.ModuleUriPrimitiveElement != nil {
			if v == nil {
				v = &Uri{}
			}
			v.Id = m.ModuleUriPrimitiveElement.Id
			v.Extension = m.ModuleUriPrimitiveElement.Extension
		}
		r.Module = v
	}
	if m.ModuleCanonical != nil || m.ModuleCanonicalPrimitiveElement != nil {
		if r.Module != nil {
			return fmt.Errorf("multiple values for field \"Module\"")
		}
		v := m.ModuleCanonical
		if m.ModuleCanonicalPrimitiveElement != nil {
			if v == nil {
				v = &Canonical{}
			}
			v.Id = m.ModuleCanonicalPrimitiveElement.Id
			v.Extension = m.ModuleCanonicalPrimitiveElement.Extension
		}
		r.Module = v
	}
	if m.ModuleCodeableConcept != nil {
		if r.Module != nil {
			return fmt.Errorf("multiple values for field \"Module\"")
		}
		v := m.ModuleCodeableConcept
		r.Module = v
	}
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Subject = m.Subject
	r.Encounter = m.Encounter
	r.OccurrenceDateTime = m.OccurrenceDateTime
	if m.OccurrenceDateTimePrimitiveElement != nil {
		r.OccurrenceDateTime.Id = m.OccurrenceDateTimePrimitiveElement.Id
		r.OccurrenceDateTime.Extension = m.OccurrenceDateTimePrimitiveElement.Extension
	}
	r.Performer = m.Performer
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Note = m.Note
	r.EvaluationMessage = m.EvaluationMessage
	r.OutputParameters = m.OutputParameters
	r.Result = m.Result
	r.DataRequirement = m.DataRequirement
	return nil
}
func (r GuidanceResponse) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
