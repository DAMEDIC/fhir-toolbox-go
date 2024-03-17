package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// The HealthcareService is not available during this period of time due to the provided reason.
type HealthcareServiceNotAvailable struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The reason that can be presented to the user as to why this time is not available.
	Description types.String
	// Service is not available (seasonally or for a public holiday) from this date.
	During *types.Period
}

func (s HealthcareServiceNotAvailable) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Does this service have specific eligibility requirements that need to be met in order to use the service?
type HealthcareServiceEligibility struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Coded value for the eligibility.
	Code *types.CodeableConcept
	// Describes the eligibility conditions for the service.
	Comment *types.Markdown
}

func (s HealthcareServiceEligibility) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A collection of times that the Service Site is available.
type HealthcareServiceAvailableTime struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Indicates which days of the week are available between the start and end Times.
	DaysOfWeek []types.Code
	// Is this always available? (hence times are irrelevant) e.g. 24 hour service.
	AllDay *types.Boolean
	// The opening time of day. Note: If the AllDay flag is set, then this time is ignored.
	AvailableStartTime *types.Time
	// The closing time of day. Note: If the AllDay flag is set, then this time is ignored.
	AvailableEndTime *types.Time
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
}

func (s HealthcareServiceAvailableTime) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The details of a healthcare service available at a location.
type HealthcareService struct {
	// The organization that provides this healthcare service.
	ProvidedBy *types.Reference
	// Collection of characteristics (attributes).
	Characteristic []types.CodeableConcept
	// Indicates whether or not a prospective consumer will require an appointment for a particular service at a site to be provided by the Organization. Indicates if an appointment is required for access to this service.
	AppointmentRequired *types.Boolean
	// The base language in which the resource is written.
	Language *types.Code
	// External identifiers for this item.
	Identifier []types.Identifier
	// The location(s) where this healthcare service may be provided.
	Location []types.Reference
	// Any additional description of the service and/or any specific issues not covered by the other attributes, which can be displayed as further detail under the serviceName.
	Comment *types.String
	// Ways that the service accepts referrals, if this is not provided then it is implied that no referral is required.
	ReferralMethod []types.CodeableConcept
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// Identifies the broad category of service being performed or delivered.
	Category []types.CodeableConcept
	// The specific type of service that may be delivered or performed.
	Type []types.CodeableConcept
	// Some services are specifically made available in multiple languages, this property permits a directory to declare the languages this is offered in. Typically this is only provided where a service operates in communities with mixed languages used.
	Communication []types.CodeableConcept
	// The HealthcareService is not available during this period of time due to the provided reason.
	NotAvailable []HealthcareServiceNotAvailable
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Collection of specialties handled by the service site. This is more of a medical term.
	Specialty []types.CodeableConcept
	// Extra details about the service that can't be placed in the other fields.
	ExtraDetails *types.Markdown
	// Does this service have specific eligibility requirements that need to be met in order to use the service?
	Eligibility []HealthcareServiceEligibility
	// Programs that this service is applicable to.
	Program []types.CodeableConcept
	// Further description of the service as it would be presented to a consumer while searching.
	Name *types.String
	// If there is a photo/symbol associated with this HealthcareService, it may be included here to facilitate quick identification of the service in a list.
	Photo *types.Attachment
	// The location(s) that this service is available to (not where the service is provided).
	CoverageArea []types.Reference
	// The code(s) that detail the conditions under which the healthcare service is available/offered.
	ServiceProvisionCode []types.CodeableConcept
	// Technical endpoints providing access to services operated for the specific healthcare services defined at this resource.
	Endpoint []types.Reference
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// List of contacts related to this specific healthcare service.
	Telecom []types.ContactPoint
	// A collection of times that the Service Site is available.
	AvailableTime []HealthcareServiceAvailableTime
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// This flag is used to mark the record to not be used. This is not used when a center is closed for maintenance, or for holidays, the notAvailable period is to be used for this.
	Active *types.Boolean
	// A description of site availability exceptions, e.g. public holiday availability. Succinctly describing all possible exceptions to normal site availability as details in the available Times and not available Times.
	AvailabilityExceptions *types.String
}

func (s HealthcareService) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
