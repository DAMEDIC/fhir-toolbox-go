package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// A record of a request for service such as diagnostic investigations, treatments, or operations to be performed.
type ServiceRequest struct {
	// The preferred location(s) where the procedure should actually happen in coded or free text form. E.g. at home or nursing day care center.
	LocationCode []types.CodeableConcept
	// Indicates another resource that provides a justification for why this service is being requested.   May relate to the resources referred to in `supportingInfo`.
	ReasonReference []types.Reference
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// An encounter that provides additional information about the healthcare context in which this request is made.
	Encounter *types.Reference
	// The desired performer for doing the requested service.  For example, the surgeon, dermatopathologist, endoscopist, etc.
	Performer []types.Reference
	// Set this to true if the record is saying that the service/procedure should NOT be performed.
	DoNotPerform *types.Boolean
	// A reference to the the preferred location(s) where the procedure should actually happen. E.g. at home or nursing day care center.
	LocationReference []types.Reference
	// Any other notes and comments made about the service request. For example, internal billing notes.
	Note []types.Annotation
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// Identifiers assigned to this order instance by the orderer and/or the receiver and/or order fulfiller.
	Identifier []types.Identifier
	// Indicates how quickly the ServiceRequest should be addressed with respect to other requests.
	Priority *types.Code
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// An explanation or justification for why this service is being requested in coded or textual form.   This is often for billing purposes.  May relate to the resources referred to in `supportingInfo`.
	ReasonCode []types.CodeableConcept
	// The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this ServiceRequest.
	InstantiatesUri []types.Uri
	// The date/time at which the requested service should occur.
	Occurrence r4.Element
	// Key events in the history of the request.
	RelevantHistory []types.Reference
	// The base language in which the resource is written.
	Language *types.Code
	// Plan/proposal/order fulfilled by this request.
	BasedOn []types.Reference
	// Additional details and instructions about the how the services are to be delivered.   For example, and order for a urinary catheter may have an order detail for an external or indwelling catheter, or an order for a bandage may require additional instructions specifying how the bandage should be applied.
	OrderDetail []types.CodeableConcept
	// If a CodeableConcept is present, it indicates the pre-condition for performing the service.  For example "pain", "on flare-up", etc.
	AsNeeded r4.Element
	// When the request transitioned to being actionable.
	AuthoredOn *types.DateTime
	// The individual who initiated the request and has responsibility for its activation.
	Requester *types.Reference
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// The request takes the place of the referenced completed or terminated request(s).
	Replaces []types.Reference
	// A shared identifier common to all service requests that were authorized more or less simultaneously by a single author, representing the composite or group identifier.
	Requisition *types.Identifier
	// A code that identifies a particular service (i.e., procedure, diagnostic investigation, or panel of investigations) that have been requested.
	Code *types.CodeableConcept
	// An amount of service being requested which can be a quantity ( for example $1,500 home modification), a ratio ( for example, 20 half day visits per month), or a range (2.0 to 1.8 Gy per fraction).
	Quantity r4.Element
	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be needed for delivering the requested service.
	Insurance []types.Reference
	// Additional clinical information about the patient or specimen that may influence the services or their interpretations.     This information includes diagnosis, clinical findings and other observations.  In laboratory ordering these are typically referred to as "ask at order entry questions (AOEs)".  This includes observations explicitly requested by the producer (filler) to provide context or supporting information needed to complete the order. For example,  reporting the amount of inspired oxygen for blood gas measurements.
	SupportingInfo []types.Reference
	// Instructions in terms that are understood by the patient or consumer.
	PatientInstruction *types.String
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The status of the order.
	Status types.Code
	// A code that classifies the service for searching, sorting and display purposes (e.g. "Surgical Procedure").
	Category []types.CodeableConcept
	// Whether the request is a proposal, plan, an original order or a reflex order.
	Intent types.Code
	// On whom or what the service is to be performed. This is usually a human patient, but can also be requested on animals, groups of humans or animals, devices such as dialysis machines, or even locations (typically for environmental scans).
	Subject types.Reference
	// Desired type of performer for doing the requested service.
	PerformerType *types.CodeableConcept
	// One or more specimens that the laboratory procedure will use.
	Specimen []types.Reference
	// Anatomic location where the procedure should be performed. This is the target site.
	BodySite []types.CodeableConcept
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this ServiceRequest.
	InstantiatesCanonical []types.Canonical
}

func (s ServiceRequest) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
