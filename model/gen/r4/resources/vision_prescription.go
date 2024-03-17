package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Allows for adjustment on two axis.
type VisionPrescriptionLensSpecificationPrism struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Amount of prism to compensate for eye alignment in fractional units.
	Amount types.Decimal
	// The relative base, or reference lens edge, for the prism.
	Base types.Code
}

func (s VisionPrescriptionLensSpecificationPrism) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Contain the details of  the individual lens specifications and serves as the authorization for the fullfillment by certified professionals.
type VisionPrescriptionLensSpecification struct {
	// Contact lens power measured in dioptres (0.25 units).
	Power *types.Decimal
	// Contact lens diameter measured in millimetres.
	Diameter *types.Decimal
	// Special color or pattern.
	Color *types.String
	// Brand recommendations or restrictions.
	Brand *types.String
	// The eye for which the lens specification applies.
	Eye types.Code
	// Adjustment for astigmatism measured in integer degrees.
	Axis *types.Integer
	// Power adjustment for multifocal lenses measured in dioptres (0.25 units).
	Add *types.Decimal
	// Back curvature measured in millimetres.
	BackCurve *types.Decimal
	// Identifies the type of vision correction product which is required for the patient.
	Product types.CodeableConcept
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Lens power measured in dioptres (0.25 units).
	Sphere *types.Decimal
	// Power adjustment for astigmatism measured in dioptres (0.25 units).
	Cylinder *types.Decimal
	// Notes for special requirements such as coatings and lens materials.
	Note []types.Annotation
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Allows for adjustment on two axis.
	Prism []VisionPrescriptionLensSpecificationPrism
	// The recommended maximum wear period for the lens.
	Duration *types.Quantity
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
}

func (s VisionPrescriptionLensSpecification) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// An authorization for the provision of glasses and/or contact lenses to a patient.
type VisionPrescription struct {
	// A unique identifier assigned to this vision prescription.
	Identifier []types.Identifier
	// A reference to a resource that identifies the particular occurrence of contact between patient and health care provider during which the prescription was issued.
	Encounter *types.Reference
	// The date (and perhaps time) when the prescription was written.
	DateWritten types.DateTime
	// The healthcare professional responsible for authorizing the prescription.
	Prescriber types.Reference
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// The base language in which the resource is written.
	Language *types.Code
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A resource reference to the person to whom the vision prescription applies.
	Patient types.Reference
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// The date this resource was created.
	Created types.DateTime
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// The status of the resource instance.
	Status types.Code
	// Contain the details of  the individual lens specifications and serves as the authorization for the fullfillment by certified professionals.
	LensSpecification []VisionPrescriptionLensSpecification
}

func (s VisionPrescription) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
