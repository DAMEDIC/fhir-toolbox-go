package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// The practitioner is not available or performing this role during this period of time due to the provided reason.
type PractitionerRoleNotAvailable struct {
	// Service is not available (seasonally or for a public holiday) from this date.
	During *types.Period
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
}

func (s PractitionerRoleNotAvailable) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A collection of times the practitioner is available or performing this role at the location and/or healthcareservice.
type PractitionerRoleAvailableTime struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
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
}

func (s PractitionerRoleAvailableTime) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A specific set of Roles/Locations/specialties/services that a practitioner may perform at an organization for a period of time.
//
// Need to track services that a healthcare provider is able to provide at an organization's location, and the services that they can perform there.
type PractitionerRole struct {
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// Contact details that are specific to the role/location/service.
	Telecom []types.ContactPoint
	// A description of site availability exceptions, e.g. public holiday availability. Succinctly describing all possible exceptions to normal site availability as details in the available Times and not available Times.
	AvailabilityExceptions *types.String
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// The organization where the Practitioner performs the roles associated.
	Organization *types.Reference
	// Roles which this practitioner is authorized to perform for the organization.
	Code []types.CodeableConcept
	// Specific specialty of the practitioner.
	Specialty []types.CodeableConcept
	// The list of healthcare services that this worker provides for this role's Organization/Location(s).
	HealthcareService []types.Reference
	// The practitioner is not available or performing this role during this period of time due to the provided reason.
	NotAvailable []PractitionerRoleNotAvailable
	// Technical endpoints providing access to services operated for the practitioner with this role.
	Endpoint []types.Reference
	// The period during which the person is authorized to act as a practitioner in these role(s) for the organization.
	Period *types.Period
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// The base language in which the resource is written.
	Language *types.Code
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Business Identifiers that are specific to a role/location.
	Identifier []types.Identifier
	// Whether this practitioner role record is in active use.
	Active *types.Boolean
	// Practitioner that is able to provide the defined services for the organization.
	Practitioner *types.Reference
	// The location(s) at which this practitioner provides care.
	Location []types.Reference
	// A collection of times the practitioner is available or performing this role at the location and/or healthcareservice.
	AvailableTime []PractitionerRoleAvailableTime
}

func (s PractitionerRole) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
