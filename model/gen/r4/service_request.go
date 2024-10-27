package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A record of a request for service such as diagnostic investigations, treatments, or operations to be performed.
type ServiceRequest struct {
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
	// Identifiers assigned to this order instance by the orderer and/or the receiver and/or order fulfiller.
	Identifier []Identifier
	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this ServiceRequest.
	InstantiatesCanonical []Canonical
	// The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this ServiceRequest.
	InstantiatesUri []Uri
	// Plan/proposal/order fulfilled by this request.
	BasedOn []Reference
	// The request takes the place of the referenced completed or terminated request(s).
	Replaces []Reference
	// A shared identifier common to all service requests that were authorized more or less simultaneously by a single author, representing the composite or group identifier.
	Requisition *Identifier
	// The status of the order.
	Status Code
	// Whether the request is a proposal, plan, an original order or a reflex order.
	Intent Code
	// A code that classifies the service for searching, sorting and display purposes (e.g. "Surgical Procedure").
	Category []CodeableConcept
	// Indicates how quickly the ServiceRequest should be addressed with respect to other requests.
	Priority *Code
	// Set this to true if the record is saying that the service/procedure should NOT be performed.
	DoNotPerform *Boolean
	// A code that identifies a particular service (i.e., procedure, diagnostic investigation, or panel of investigations) that have been requested.
	Code *CodeableConcept
	// Additional details and instructions about the how the services are to be delivered.   For example, and order for a urinary catheter may have an order detail for an external or indwelling catheter, or an order for a bandage may require additional instructions specifying how the bandage should be applied.
	OrderDetail []CodeableConcept
	// An amount of service being requested which can be a quantity ( for example $1,500 home modification), a ratio ( for example, 20 half day visits per month), or a range (2.0 to 1.8 Gy per fraction).
	Quantity isServiceRequestQuantity
	// On whom or what the service is to be performed. This is usually a human patient, but can also be requested on animals, groups of humans or animals, devices such as dialysis machines, or even locations (typically for environmental scans).
	Subject Reference
	// An encounter that provides additional information about the healthcare context in which this request is made.
	Encounter *Reference
	// The date/time at which the requested service should occur.
	Occurrence isServiceRequestOccurrence
	// If a CodeableConcept is present, it indicates the pre-condition for performing the service.  For example "pain", "on flare-up", etc.
	AsNeeded isServiceRequestAsNeeded
	// When the request transitioned to being actionable.
	AuthoredOn *DateTime
	// The individual who initiated the request and has responsibility for its activation.
	Requester *Reference
	// Desired type of performer for doing the requested service.
	PerformerType *CodeableConcept
	// The desired performer for doing the requested service.  For example, the surgeon, dermatopathologist, endoscopist, etc.
	Performer []Reference
	// The preferred location(s) where the procedure should actually happen in coded or free text form. E.g. at home or nursing day care center.
	LocationCode []CodeableConcept
	// A reference to the the preferred location(s) where the procedure should actually happen. E.g. at home or nursing day care center.
	LocationReference []Reference
	// An explanation or justification for why this service is being requested in coded or textual form.   This is often for billing purposes.  May relate to the resources referred to in `supportingInfo`.
	ReasonCode []CodeableConcept
	// Indicates another resource that provides a justification for why this service is being requested.   May relate to the resources referred to in `supportingInfo`.
	ReasonReference []Reference
	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be needed for delivering the requested service.
	Insurance []Reference
	// Additional clinical information about the patient or specimen that may influence the services or their interpretations.     This information includes diagnosis, clinical findings and other observations.  In laboratory ordering these are typically referred to as "ask at order entry questions (AOEs)".  This includes observations explicitly requested by the producer (filler) to provide context or supporting information needed to complete the order. For example,  reporting the amount of inspired oxygen for blood gas measurements.
	SupportingInfo []Reference
	// One or more specimens that the laboratory procedure will use.
	Specimen []Reference
	// Anatomic location where the procedure should be performed. This is the target site.
	BodySite []CodeableConcept
	// Any other notes and comments made about the service request. For example, internal billing notes.
	Note []Annotation
	// Instructions in terms that are understood by the patient or consumer.
	PatientInstruction *String
	// Key events in the history of the request.
	RelevantHistory []Reference
}

func (r ServiceRequest) ResourceType() string {
	return "ServiceRequest"
}
func (r ServiceRequest) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isServiceRequestQuantity interface {
	isServiceRequestQuantity()
}

func (r Quantity) isServiceRequestQuantity() {}
func (r Ratio) isServiceRequestQuantity()    {}
func (r Range) isServiceRequestQuantity()    {}

type isServiceRequestOccurrence interface {
	isServiceRequestOccurrence()
}

func (r DateTime) isServiceRequestOccurrence() {}
func (r Period) isServiceRequestOccurrence()   {}
func (r Timing) isServiceRequestOccurrence()   {}

type isServiceRequestAsNeeded interface {
	isServiceRequestAsNeeded()
}

func (r Boolean) isServiceRequestAsNeeded()         {}
func (r CodeableConcept) isServiceRequestAsNeeded() {}

type jsonServiceRequest struct {
	ResourceType                          string              `json:"resourceType"`
	Id                                    *Id                 `json:"id,omitempty"`
	IdPrimitiveElement                    *primitiveElement   `json:"_id,omitempty"`
	Meta                                  *Meta               `json:"meta,omitempty"`
	ImplicitRules                         *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement         *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                              *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement              *primitiveElement   `json:"_language,omitempty"`
	Text                                  *Narrative          `json:"text,omitempty"`
	Contained                             []ContainedResource `json:"contained,omitempty"`
	Extension                             []Extension         `json:"extension,omitempty"`
	ModifierExtension                     []Extension         `json:"modifierExtension,omitempty"`
	Identifier                            []Identifier        `json:"identifier,omitempty"`
	InstantiatesCanonical                 []Canonical         `json:"instantiatesCanonical,omitempty"`
	InstantiatesCanonicalPrimitiveElement []*primitiveElement `json:"_instantiatesCanonical,omitempty"`
	InstantiatesUri                       []Uri               `json:"instantiatesUri,omitempty"`
	InstantiatesUriPrimitiveElement       []*primitiveElement `json:"_instantiatesUri,omitempty"`
	BasedOn                               []Reference         `json:"basedOn,omitempty"`
	Replaces                              []Reference         `json:"replaces,omitempty"`
	Requisition                           *Identifier         `json:"requisition,omitempty"`
	Status                                Code                `json:"status,omitempty"`
	StatusPrimitiveElement                *primitiveElement   `json:"_status,omitempty"`
	Intent                                Code                `json:"intent,omitempty"`
	IntentPrimitiveElement                *primitiveElement   `json:"_intent,omitempty"`
	Category                              []CodeableConcept   `json:"category,omitempty"`
	Priority                              *Code               `json:"priority,omitempty"`
	PriorityPrimitiveElement              *primitiveElement   `json:"_priority,omitempty"`
	DoNotPerform                          *Boolean            `json:"doNotPerform,omitempty"`
	DoNotPerformPrimitiveElement          *primitiveElement   `json:"_doNotPerform,omitempty"`
	Code                                  *CodeableConcept    `json:"code,omitempty"`
	OrderDetail                           []CodeableConcept   `json:"orderDetail,omitempty"`
	QuantityQuantity                      *Quantity           `json:"quantityQuantity,omitempty"`
	QuantityRatio                         *Ratio              `json:"quantityRatio,omitempty"`
	QuantityRange                         *Range              `json:"quantityRange,omitempty"`
	Subject                               Reference           `json:"subject,omitempty"`
	Encounter                             *Reference          `json:"encounter,omitempty"`
	OccurrenceDateTime                    *DateTime           `json:"occurrenceDateTime,omitempty"`
	OccurrenceDateTimePrimitiveElement    *primitiveElement   `json:"_occurrenceDateTime,omitempty"`
	OccurrencePeriod                      *Period             `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming                      *Timing             `json:"occurrenceTiming,omitempty"`
	AsNeededBoolean                       *Boolean            `json:"asNeededBoolean,omitempty"`
	AsNeededBooleanPrimitiveElement       *primitiveElement   `json:"_asNeededBoolean,omitempty"`
	AsNeededCodeableConcept               *CodeableConcept    `json:"asNeededCodeableConcept,omitempty"`
	AuthoredOn                            *DateTime           `json:"authoredOn,omitempty"`
	AuthoredOnPrimitiveElement            *primitiveElement   `json:"_authoredOn,omitempty"`
	Requester                             *Reference          `json:"requester,omitempty"`
	PerformerType                         *CodeableConcept    `json:"performerType,omitempty"`
	Performer                             []Reference         `json:"performer,omitempty"`
	LocationCode                          []CodeableConcept   `json:"locationCode,omitempty"`
	LocationReference                     []Reference         `json:"locationReference,omitempty"`
	ReasonCode                            []CodeableConcept   `json:"reasonCode,omitempty"`
	ReasonReference                       []Reference         `json:"reasonReference,omitempty"`
	Insurance                             []Reference         `json:"insurance,omitempty"`
	SupportingInfo                        []Reference         `json:"supportingInfo,omitempty"`
	Specimen                              []Reference         `json:"specimen,omitempty"`
	BodySite                              []CodeableConcept   `json:"bodySite,omitempty"`
	Note                                  []Annotation        `json:"note,omitempty"`
	PatientInstruction                    *String             `json:"patientInstruction,omitempty"`
	PatientInstructionPrimitiveElement    *primitiveElement   `json:"_patientInstruction,omitempty"`
	RelevantHistory                       []Reference         `json:"relevantHistory,omitempty"`
}

func (r ServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ServiceRequest) marshalJSON() jsonServiceRequest {
	m := jsonServiceRequest{}
	m.ResourceType = "ServiceRequest"
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
	anyInstantiatesCanonicalValue := false
	for _, e := range r.InstantiatesCanonical {
		if e.Value != nil {
			anyInstantiatesCanonicalValue = true
			break
		}
	}
	if anyInstantiatesCanonicalValue {
		m.InstantiatesCanonical = r.InstantiatesCanonical
	}
	anyInstantiatesCanonicalIdOrExtension := false
	for _, e := range r.InstantiatesCanonical {
		if e.Id != nil || e.Extension != nil {
			anyInstantiatesCanonicalIdOrExtension = true
			break
		}
	}
	if anyInstantiatesCanonicalIdOrExtension {
		m.InstantiatesCanonicalPrimitiveElement = make([]*primitiveElement, 0, len(r.InstantiatesCanonical))
		for _, e := range r.InstantiatesCanonical {
			if e.Id != nil || e.Extension != nil {
				m.InstantiatesCanonicalPrimitiveElement = append(m.InstantiatesCanonicalPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.InstantiatesCanonicalPrimitiveElement = append(m.InstantiatesCanonicalPrimitiveElement, nil)
			}
		}
	}
	anyInstantiatesUriValue := false
	for _, e := range r.InstantiatesUri {
		if e.Value != nil {
			anyInstantiatesUriValue = true
			break
		}
	}
	if anyInstantiatesUriValue {
		m.InstantiatesUri = r.InstantiatesUri
	}
	anyInstantiatesUriIdOrExtension := false
	for _, e := range r.InstantiatesUri {
		if e.Id != nil || e.Extension != nil {
			anyInstantiatesUriIdOrExtension = true
			break
		}
	}
	if anyInstantiatesUriIdOrExtension {
		m.InstantiatesUriPrimitiveElement = make([]*primitiveElement, 0, len(r.InstantiatesUri))
		for _, e := range r.InstantiatesUri {
			if e.Id != nil || e.Extension != nil {
				m.InstantiatesUriPrimitiveElement = append(m.InstantiatesUriPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.InstantiatesUriPrimitiveElement = append(m.InstantiatesUriPrimitiveElement, nil)
			}
		}
	}
	m.BasedOn = r.BasedOn
	m.Replaces = r.Replaces
	m.Requisition = r.Requisition
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	if r.Intent.Value != nil {
		m.Intent = r.Intent
	}
	if r.Intent.Id != nil || r.Intent.Extension != nil {
		m.IntentPrimitiveElement = &primitiveElement{Id: r.Intent.Id, Extension: r.Intent.Extension}
	}
	m.Category = r.Category
	if r.Priority != nil && r.Priority.Value != nil {
		m.Priority = r.Priority
	}
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		m.PriorityPrimitiveElement = &primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
	}
	if r.DoNotPerform != nil && r.DoNotPerform.Value != nil {
		m.DoNotPerform = r.DoNotPerform
	}
	if r.DoNotPerform != nil && (r.DoNotPerform.Id != nil || r.DoNotPerform.Extension != nil) {
		m.DoNotPerformPrimitiveElement = &primitiveElement{Id: r.DoNotPerform.Id, Extension: r.DoNotPerform.Extension}
	}
	m.Code = r.Code
	m.OrderDetail = r.OrderDetail
	switch v := r.Quantity.(type) {
	case Quantity:
		m.QuantityQuantity = &v
	case *Quantity:
		m.QuantityQuantity = v
	case Ratio:
		m.QuantityRatio = &v
	case *Ratio:
		m.QuantityRatio = v
	case Range:
		m.QuantityRange = &v
	case *Range:
		m.QuantityRange = v
	}
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	switch v := r.Occurrence.(type) {
	case DateTime:
		if v.Value != nil {
			m.OccurrenceDateTime = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.OccurrenceDateTimePrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *DateTime:
		if v.Value != nil {
			m.OccurrenceDateTime = v
		}
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
	switch v := r.AsNeeded.(type) {
	case Boolean:
		if v.Value != nil {
			m.AsNeededBoolean = &v
		}
		if v.Id != nil || v.Extension != nil {
			m.AsNeededBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Boolean:
		if v.Value != nil {
			m.AsNeededBoolean = v
		}
		if v.Id != nil || v.Extension != nil {
			m.AsNeededBooleanPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case CodeableConcept:
		m.AsNeededCodeableConcept = &v
	case *CodeableConcept:
		m.AsNeededCodeableConcept = v
	}
	if r.AuthoredOn != nil && r.AuthoredOn.Value != nil {
		m.AuthoredOn = r.AuthoredOn
	}
	if r.AuthoredOn != nil && (r.AuthoredOn.Id != nil || r.AuthoredOn.Extension != nil) {
		m.AuthoredOnPrimitiveElement = &primitiveElement{Id: r.AuthoredOn.Id, Extension: r.AuthoredOn.Extension}
	}
	m.Requester = r.Requester
	m.PerformerType = r.PerformerType
	m.Performer = r.Performer
	m.LocationCode = r.LocationCode
	m.LocationReference = r.LocationReference
	m.ReasonCode = r.ReasonCode
	m.ReasonReference = r.ReasonReference
	m.Insurance = r.Insurance
	m.SupportingInfo = r.SupportingInfo
	m.Specimen = r.Specimen
	m.BodySite = r.BodySite
	m.Note = r.Note
	if r.PatientInstruction != nil && r.PatientInstruction.Value != nil {
		m.PatientInstruction = r.PatientInstruction
	}
	if r.PatientInstruction != nil && (r.PatientInstruction.Id != nil || r.PatientInstruction.Extension != nil) {
		m.PatientInstructionPrimitiveElement = &primitiveElement{Id: r.PatientInstruction.Id, Extension: r.PatientInstruction.Extension}
	}
	m.RelevantHistory = r.RelevantHistory
	return m
}
func (r *ServiceRequest) UnmarshalJSON(b []byte) error {
	var m jsonServiceRequest
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ServiceRequest) unmarshalJSON(m jsonServiceRequest) error {
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
	r.InstantiatesCanonical = m.InstantiatesCanonical
	for i, e := range m.InstantiatesCanonicalPrimitiveElement {
		if len(r.InstantiatesCanonical) <= i {
			r.InstantiatesCanonical = append(r.InstantiatesCanonical, Canonical{})
		}
		if e != nil {
			r.InstantiatesCanonical[i].Id = e.Id
			r.InstantiatesCanonical[i].Extension = e.Extension
		}
	}
	r.InstantiatesUri = m.InstantiatesUri
	for i, e := range m.InstantiatesUriPrimitiveElement {
		if len(r.InstantiatesUri) <= i {
			r.InstantiatesUri = append(r.InstantiatesUri, Uri{})
		}
		if e != nil {
			r.InstantiatesUri[i].Id = e.Id
			r.InstantiatesUri[i].Extension = e.Extension
		}
	}
	r.BasedOn = m.BasedOn
	r.Replaces = m.Replaces
	r.Requisition = m.Requisition
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Intent = m.Intent
	if m.IntentPrimitiveElement != nil {
		r.Intent.Id = m.IntentPrimitiveElement.Id
		r.Intent.Extension = m.IntentPrimitiveElement.Extension
	}
	r.Category = m.Category
	r.Priority = m.Priority
	if m.PriorityPrimitiveElement != nil {
		if r.Priority == nil {
			r.Priority = &Code{}
		}
		r.Priority.Id = m.PriorityPrimitiveElement.Id
		r.Priority.Extension = m.PriorityPrimitiveElement.Extension
	}
	r.DoNotPerform = m.DoNotPerform
	if m.DoNotPerformPrimitiveElement != nil {
		if r.DoNotPerform == nil {
			r.DoNotPerform = &Boolean{}
		}
		r.DoNotPerform.Id = m.DoNotPerformPrimitiveElement.Id
		r.DoNotPerform.Extension = m.DoNotPerformPrimitiveElement.Extension
	}
	r.Code = m.Code
	r.OrderDetail = m.OrderDetail
	if m.QuantityQuantity != nil {
		if r.Quantity != nil {
			return fmt.Errorf("multiple values for field \"Quantity\"")
		}
		v := m.QuantityQuantity
		r.Quantity = v
	}
	if m.QuantityRatio != nil {
		if r.Quantity != nil {
			return fmt.Errorf("multiple values for field \"Quantity\"")
		}
		v := m.QuantityRatio
		r.Quantity = v
	}
	if m.QuantityRange != nil {
		if r.Quantity != nil {
			return fmt.Errorf("multiple values for field \"Quantity\"")
		}
		v := m.QuantityRange
		r.Quantity = v
	}
	r.Subject = m.Subject
	r.Encounter = m.Encounter
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
	if m.AsNeededBoolean != nil || m.AsNeededBooleanPrimitiveElement != nil {
		if r.AsNeeded != nil {
			return fmt.Errorf("multiple values for field \"AsNeeded\"")
		}
		v := m.AsNeededBoolean
		if m.AsNeededBooleanPrimitiveElement != nil {
			if v == nil {
				v = &Boolean{}
			}
			v.Id = m.AsNeededBooleanPrimitiveElement.Id
			v.Extension = m.AsNeededBooleanPrimitiveElement.Extension
		}
		r.AsNeeded = v
	}
	if m.AsNeededCodeableConcept != nil {
		if r.AsNeeded != nil {
			return fmt.Errorf("multiple values for field \"AsNeeded\"")
		}
		v := m.AsNeededCodeableConcept
		r.AsNeeded = v
	}
	r.AuthoredOn = m.AuthoredOn
	if m.AuthoredOnPrimitiveElement != nil {
		if r.AuthoredOn == nil {
			r.AuthoredOn = &DateTime{}
		}
		r.AuthoredOn.Id = m.AuthoredOnPrimitiveElement.Id
		r.AuthoredOn.Extension = m.AuthoredOnPrimitiveElement.Extension
	}
	r.Requester = m.Requester
	r.PerformerType = m.PerformerType
	r.Performer = m.Performer
	r.LocationCode = m.LocationCode
	r.LocationReference = m.LocationReference
	r.ReasonCode = m.ReasonCode
	r.ReasonReference = m.ReasonReference
	r.Insurance = m.Insurance
	r.SupportingInfo = m.SupportingInfo
	r.Specimen = m.Specimen
	r.BodySite = m.BodySite
	r.Note = m.Note
	r.PatientInstruction = m.PatientInstruction
	if m.PatientInstructionPrimitiveElement != nil {
		if r.PatientInstruction == nil {
			r.PatientInstruction = &String{}
		}
		r.PatientInstruction.Id = m.PatientInstructionPrimitiveElement.Id
		r.PatientInstruction.Extension = m.PatientInstructionPrimitiveElement.Extension
	}
	r.RelevantHistory = m.RelevantHistory
	return nil
}
func (r ServiceRequest) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
