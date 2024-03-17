package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Specific parameters for the ordered item.  For example, the prism value for lenses.
type DeviceRequestParameter struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A code or string that identifies the device detail being asserted.
	Code *types.CodeableConcept
	// The value of the device detail.
	Value r4.Element
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
}

func (s DeviceRequestParameter) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Represents a request for a patient to employ a medical device. The device may be an implantable device, or an external assistive device, such as a walker.
type DeviceRequest struct {
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The request takes the place of the referenced completed or terminated request(s).
	PriorRequest []types.Reference
	// Specific parameters for the ordered item.  For example, the prism value for lenses.
	Parameter []DeviceRequestParameter
	// The patient who will use the device.
	Subject types.Reference
	// An encounter that provides additional context in which this request is made.
	Encounter *types.Reference
	// Key events in the history of the request.
	RelevantHistory []types.Reference
	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this DeviceRequest.
	InstantiatesCanonical []types.Canonical
	// Composite request this is part of.
	GroupIdentifier *types.Identifier
	// Desired type of performer for doing the diagnostic testing.
	PerformerType *types.CodeableConcept
	// Insurance plans, coverage extensions, pre-authorizations and/or pre-determinations that may be required for delivering the requested service.
	Insurance []types.Reference
	// The status of the request.
	Status *types.Code
	// Indicates how quickly the {{title}} should be addressed with respect to other requests.
	Priority *types.Code
	// Reason or justification for the use of this device.
	ReasonCode []types.CodeableConcept
	// Details about this request that were not represented at all or sufficiently in one of the attributes provided in a class. These may include for example a comment, an instruction, or a note associated with the statement.
	Note []types.Annotation
	// Identifiers assigned to this order by the orderer or by the receiver.
	Identifier []types.Identifier
	// The desired performer for doing the diagnostic testing.
	Performer *types.Reference
	// Reason or justification for the use of this device.
	ReasonReference []types.Reference
	// Additional clinical information about the patient that may influence the request fulfilment.  For example, this may include where on the subject's body the device will be used (i.e. the target site).
	SupportingInfo []types.Reference
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// When the request transitioned to being actionable.
	AuthoredOn *types.DateTime
	// The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this DeviceRequest.
	InstantiatesUri []types.Uri
	// Whether the request is a proposal, plan, an original order or a reflex order.
	Intent types.Code
	// The timing schedule for the use of the device. The Schedule data type allows many different expressions, for example. "Every 8 hours"; "Three times a day"; "1/2 an hour before breakfast for 10 days from 23-Dec 2011:"; "15 Oct 2013, 17 Oct 2013 and 1 Nov 2013".
	Occurrence r4.Element
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Plan/proposal/order fulfilled by this request.
	BasedOn []types.Reference
	// The details of the device to be used.
	Code r4.Element
	// The individual who initiated the request and has responsibility for its activation.
	Requester *types.Reference
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// The base language in which the resource is written.
	Language *types.Code
}

func (s DeviceRequest) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
