package r5

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	fhirpath "github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	model "github.com/DAMEDIC/fhir-toolbox-go/model"
	"io"
	"slices"
	"unsafe"
)

// A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident.
type NutritionOrder struct {
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
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, nor can they have their own independent transaction scope. This is allowed to be a Parameters resource if and only if it is referenced by a resource that provides context/meaning.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identifiers assigned to this order by the order sender or by the order receiver.
	Identifier []Identifier
	// The URL pointing to a FHIR-defined protocol, guideline, orderset or other definition that is adhered to in whole or in part by this NutritionOrder.
	InstantiatesCanonical []Canonical
	// The URL pointing to an externally maintained protocol, guideline, orderset or other definition that is adhered to in whole or in part by this NutritionOrder.
	InstantiatesUri []Uri
	// The URL pointing to a protocol, guideline, orderset or other definition that is adhered to in whole or in part by this NutritionOrder.
	Instantiates []Uri
	// A plan or request that is fulfilled in whole or in part by this nutrition order.
	BasedOn []Reference
	// A shared identifier common to all nutrition orders that were authorized more or less simultaneously by a single author, representing the composite or group identifier.
	GroupIdentifier *Identifier
	// The workflow status of the nutrition order/request.
	Status Code
	// Indicates the level of authority/intentionality associated with the NutrionOrder and where the request fits into the workflow chain.
	Intent Code
	// Indicates how quickly the Nutrition Order should be addressed with respect to other        requests.
	Priority *Code
	// The person or set of individuals who needs the nutrition order for an oral diet, nutritional supplement and/or enteral or formula feeding.
	Subject Reference
	// An encounter that provides additional information about the healthcare context in which this request is made.
	Encounter *Reference
	// Information to support fulfilling (i.e. dispensing or administering) of the nutrition,        for example, patient height and weight).
	SupportingInformation []Reference
	// The date and time that this nutrition order was requested.
	DateTime DateTime
	// The practitioner that holds legal responsibility for ordering the diet, nutritional supplement, or formula feedings.
	Orderer *Reference
	// The specified desired performer of the nutrition order.
	Performer []CodeableReference
	// A link to a record of allergies or intolerances  which should be included in the nutrition order.
	AllergyIntolerance []Reference
	// This modifier is used to convey order-specific modifiers about the type of food that should be given. These can be derived from patient allergies, intolerances, or preferences such as Halal, Vegan or Kosher. This modifier applies to the entire nutrition order inclusive of the oral diet, nutritional supplements and enteral formula feedings.
	FoodPreferenceModifier []CodeableConcept
	// This modifier is used to convey Order-specific modifier about the type of oral food or oral fluids that should not be given. These can be derived from patient allergies, intolerances, or preferences such as No Red Meat, No Soy or No Wheat or  Gluten-Free.  While it should not be necessary to repeat allergy or intolerance information captured in the referenced AllergyIntolerance resource in the excludeFoodModifier, this element may be used to convey additional specificity related to foods that should be eliminated from the patient’s diet for any reason.  This modifier applies to the entire nutrition order inclusive of the oral diet, nutritional supplements and enteral formula feedings.
	ExcludeFoodModifier []CodeableConcept
	// This modifier is used to convey whether a food item is allowed to be brought in by the patient and/or family.  If set to true, indicates that the receiving system does not need to supply the food item.
	OutsideFoodAllowed *Boolean
	// Diet given orally in contrast to enteral (tube) feeding.
	OralDiet *NutritionOrderOralDiet
	// Oral nutritional products given in order to add further nutritional value to the patient's diet.
	Supplement []NutritionOrderSupplement
	// Feeding provided through the gastrointestinal tract via a tube, catheter, or stoma that delivers nutrition distal to the oral cavity.
	EnteralFormula *NutritionOrderEnteralFormula
	// Comments made about the {{title}} by the requester, performer, subject or other participants.
	Note []Annotation
}

// Diet given orally in contrast to enteral (tube) feeding.
type NutritionOrderOralDiet struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The kind of diet or dietary restriction such as fiber restricted diet or diabetic diet.
	Type []CodeableConcept
	// Schedule information for an oral diet.
	Schedule *NutritionOrderOralDietSchedule
	// Class that defines the quantity and type of nutrient modifications (for example carbohydrate, fiber or sodium) required for the oral diet.
	Nutrient []NutritionOrderOralDietNutrient
	// Class that describes any texture modifications required for the patient to safely consume various types of solid foods.
	Texture []NutritionOrderOralDietTexture
	// The required consistency (e.g. honey-thick, nectar-thick, thin, thickened.) of liquids or fluids served to the patient.
	FluidConsistencyType []CodeableConcept
	// Free text or additional instructions or information pertaining to the oral diet.
	Instruction *String
}

// Schedule information for an oral diet.
type NutritionOrderOralDietSchedule struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The time period and frequency at which the diet should be given.  The diet should be given for the combination of all schedules if more than one schedule is present.
	Timing []Timing
	// Indicates whether the product is only taken when needed within a specific dosing schedule.
	AsNeeded *Boolean
	// Indicates whether the product is only taken based on a precondition for taking the product.
	AsNeededFor *CodeableConcept
}

// Class that defines the quantity and type of nutrient modifications (for example carbohydrate, fiber or sodium) required for the oral diet.
type NutritionOrderOralDietNutrient struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The nutrient that is being modified such as carbohydrate or sodium.
	Modifier *CodeableConcept
	// The quantity of the specified nutrient to include in diet.
	Amount *Quantity
}

// Class that describes any texture modifications required for the patient to safely consume various types of solid foods.
type NutritionOrderOralDietTexture struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Any texture modifications (for solid foods) that should be made, e.g. easy to chew, chopped, ground, and pureed.
	Modifier *CodeableConcept
	// The food type(s) (e.g. meats, all foods)  that the texture modification applies to.  This could be all foods types.
	FoodType *CodeableConcept
}

// Oral nutritional products given in order to add further nutritional value to the patient's diet.
type NutritionOrderSupplement struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The kind of nutritional supplement product required such as a high protein or pediatric clear liquid supplement.
	Type *CodeableReference
	// The product or brand name of the nutritional supplement such as "Acme Protein Shake".
	ProductName *String
	// Schedule information for a supplement.
	Schedule *NutritionOrderSupplementSchedule
	// The amount of the nutritional supplement to be given.
	Quantity *Quantity
	// Free text or additional instructions or information pertaining to the oral supplement.
	Instruction *String
}

// Schedule information for a supplement.
type NutritionOrderSupplementSchedule struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The time period and frequency at which the supplement should be given.  The supplement should be given for the combination of all schedules if more than one schedule is present.
	Timing []Timing
	// Indicates whether the supplement is only taken when needed within a specific dosing schedule.
	AsNeeded *Boolean
	// Indicates whether the supplement is only taken based on a precondition for taking the supplement.
	AsNeededFor *CodeableConcept
}

// Feeding provided through the gastrointestinal tract via a tube, catheter, or stoma that delivers nutrition distal to the oral cavity.
type NutritionOrderEnteralFormula struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of enteral or infant formula such as an adult standard formula with fiber or a soy-based infant formula.
	BaseFormulaType *CodeableReference
	// The product or brand name of the enteral or infant formula product such as "ACME Adult Standard Formula".
	BaseFormulaProductName *String
	// The intended type of device that is to be used for the administration of the enteral formula.
	DeliveryDevice []CodeableReference
	// Indicates modular components to be provided in addition or mixed with the base formula.
	Additive []NutritionOrderEnteralFormulaAdditive
	// The amount of energy (calories) that the formula should provide per specified volume, typically per mL or fluid oz.  For example, an infant may require a formula that provides 24 calories per fluid ounce or an adult may require an enteral formula that provides 1.5 calorie/mL.
	CaloricDensity *Quantity
	// The route or physiological path of administration into the patient's gastrointestinal  tract for purposes of providing the formula feeding, e.g. nasogastric tube.
	RouteOfAdministration *CodeableConcept
	// Formula administration instructions as structured data.  This repeating structure allows for changing the administration rate or volume over time for both bolus and continuous feeding.  An example of this would be an instruction to increase the rate of continuous feeding every 2 hours.
	Administration []NutritionOrderEnteralFormulaAdministration
	// The maximum total quantity of formula that may be administered to a subject over the period of time, e.g. 1440 mL over 24 hours.
	MaxVolumeToDeliver *Quantity
	// Free text formula administration, feeding instructions or additional instructions or information.
	AdministrationInstruction *Markdown
}

// Indicates modular components to be provided in addition or mixed with the base formula.
type NutritionOrderEnteralFormulaAdditive struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Indicates the type of modular component such as protein, carbohydrate, fat or fiber to be provided in addition to or mixed with the base formula.
	Type *CodeableReference
	// The product or brand name of the type of modular component to be added to the formula.
	ProductName *String
	// The amount of additive to be given in addition or to be mixed in with the base formula.
	Quantity *Quantity
}

// Formula administration instructions as structured data.  This repeating structure allows for changing the administration rate or volume over time for both bolus and continuous feeding.  An example of this would be an instruction to increase the rate of continuous feeding every 2 hours.
type NutritionOrderEnteralFormulaAdministration struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Schedule information for an enteral formula.
	Schedule *NutritionOrderEnteralFormulaAdministrationSchedule
	// The volume of formula to provide to the patient per the specified administration schedule.
	Quantity *Quantity
	// The rate of administration of formula via a feeding pump, e.g. 60 mL per hour, according to the specified schedule.
	Rate isNutritionOrderEnteralFormulaAdministrationRate
}
type isNutritionOrderEnteralFormulaAdministrationRate interface {
	model.Element
	isNutritionOrderEnteralFormulaAdministrationRate()
}

func (r Quantity) isNutritionOrderEnteralFormulaAdministrationRate() {}
func (r Ratio) isNutritionOrderEnteralFormulaAdministrationRate()    {}

// Schedule information for an enteral formula.
type NutritionOrderEnteralFormulaAdministrationSchedule struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and managable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The time period and frequency at which the enteral formula should be given.  The enteral formula should be given for the combination of all schedules if more than one schedule is present.
	Timing []Timing
	// Indicates whether the enteral formula is only taken when needed within a specific dosing schedule.
	AsNeeded *Boolean
	// Indicates whether the enteral formula is only taken based on a precondition for taking the enteral formula.
	AsNeededFor *CodeableConcept
}

func (r NutritionOrder) ResourceType() string {
	return "NutritionOrder"
}
func (r NutritionOrder) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}
func (r NutritionOrder) MemSize() int {
	var emptyIface any
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += r.Id.MemSize()
	}
	if r.Meta != nil {
		s += r.Meta.MemSize()
	}
	if r.ImplicitRules != nil {
		s += r.ImplicitRules.MemSize()
	}
	if r.Language != nil {
		s += r.Language.MemSize()
	}
	if r.Text != nil {
		s += r.Text.MemSize()
	}
	for _, i := range r.Contained {
		s += i.MemSize()
	}
	s += (cap(r.Contained) - len(r.Contained)) * int(unsafe.Sizeof(emptyIface))
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Identifier {
		s += i.MemSize()
	}
	s += (cap(r.Identifier) - len(r.Identifier)) * int(unsafe.Sizeof(Identifier{}))
	for _, i := range r.InstantiatesCanonical {
		s += i.MemSize()
	}
	s += (cap(r.InstantiatesCanonical) - len(r.InstantiatesCanonical)) * int(unsafe.Sizeof(Canonical{}))
	for _, i := range r.InstantiatesUri {
		s += i.MemSize()
	}
	s += (cap(r.InstantiatesUri) - len(r.InstantiatesUri)) * int(unsafe.Sizeof(Uri{}))
	for _, i := range r.Instantiates {
		s += i.MemSize()
	}
	s += (cap(r.Instantiates) - len(r.Instantiates)) * int(unsafe.Sizeof(Uri{}))
	for _, i := range r.BasedOn {
		s += i.MemSize()
	}
	s += (cap(r.BasedOn) - len(r.BasedOn)) * int(unsafe.Sizeof(Reference{}))
	if r.GroupIdentifier != nil {
		s += r.GroupIdentifier.MemSize()
	}
	s += r.Status.MemSize() - int(unsafe.Sizeof(r.Status))
	s += r.Intent.MemSize() - int(unsafe.Sizeof(r.Intent))
	if r.Priority != nil {
		s += r.Priority.MemSize()
	}
	s += r.Subject.MemSize() - int(unsafe.Sizeof(r.Subject))
	if r.Encounter != nil {
		s += r.Encounter.MemSize()
	}
	for _, i := range r.SupportingInformation {
		s += i.MemSize()
	}
	s += (cap(r.SupportingInformation) - len(r.SupportingInformation)) * int(unsafe.Sizeof(Reference{}))
	s += r.DateTime.MemSize() - int(unsafe.Sizeof(r.DateTime))
	if r.Orderer != nil {
		s += r.Orderer.MemSize()
	}
	for _, i := range r.Performer {
		s += i.MemSize()
	}
	s += (cap(r.Performer) - len(r.Performer)) * int(unsafe.Sizeof(CodeableReference{}))
	for _, i := range r.AllergyIntolerance {
		s += i.MemSize()
	}
	s += (cap(r.AllergyIntolerance) - len(r.AllergyIntolerance)) * int(unsafe.Sizeof(Reference{}))
	for _, i := range r.FoodPreferenceModifier {
		s += i.MemSize()
	}
	s += (cap(r.FoodPreferenceModifier) - len(r.FoodPreferenceModifier)) * int(unsafe.Sizeof(CodeableConcept{}))
	for _, i := range r.ExcludeFoodModifier {
		s += i.MemSize()
	}
	s += (cap(r.ExcludeFoodModifier) - len(r.ExcludeFoodModifier)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.OutsideFoodAllowed != nil {
		s += r.OutsideFoodAllowed.MemSize()
	}
	if r.OralDiet != nil {
		s += r.OralDiet.MemSize()
	}
	for _, i := range r.Supplement {
		s += i.MemSize()
	}
	s += (cap(r.Supplement) - len(r.Supplement)) * int(unsafe.Sizeof(NutritionOrderSupplement{}))
	if r.EnteralFormula != nil {
		s += r.EnteralFormula.MemSize()
	}
	for _, i := range r.Note {
		s += i.MemSize()
	}
	s += (cap(r.Note) - len(r.Note)) * int(unsafe.Sizeof(Annotation{}))
	return s
}
func (r NutritionOrderOralDiet) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Type {
		s += i.MemSize()
	}
	s += (cap(r.Type) - len(r.Type)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Schedule != nil {
		s += r.Schedule.MemSize()
	}
	for _, i := range r.Nutrient {
		s += i.MemSize()
	}
	s += (cap(r.Nutrient) - len(r.Nutrient)) * int(unsafe.Sizeof(NutritionOrderOralDietNutrient{}))
	for _, i := range r.Texture {
		s += i.MemSize()
	}
	s += (cap(r.Texture) - len(r.Texture)) * int(unsafe.Sizeof(NutritionOrderOralDietTexture{}))
	for _, i := range r.FluidConsistencyType {
		s += i.MemSize()
	}
	s += (cap(r.FluidConsistencyType) - len(r.FluidConsistencyType)) * int(unsafe.Sizeof(CodeableConcept{}))
	if r.Instruction != nil {
		s += r.Instruction.MemSize()
	}
	return s
}
func (r NutritionOrderOralDietSchedule) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Timing {
		s += i.MemSize()
	}
	s += (cap(r.Timing) - len(r.Timing)) * int(unsafe.Sizeof(Timing{}))
	if r.AsNeeded != nil {
		s += r.AsNeeded.MemSize()
	}
	if r.AsNeededFor != nil {
		s += r.AsNeededFor.MemSize()
	}
	return s
}
func (r NutritionOrderOralDietNutrient) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Modifier != nil {
		s += r.Modifier.MemSize()
	}
	if r.Amount != nil {
		s += r.Amount.MemSize()
	}
	return s
}
func (r NutritionOrderOralDietTexture) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Modifier != nil {
		s += r.Modifier.MemSize()
	}
	if r.FoodType != nil {
		s += r.FoodType.MemSize()
	}
	return s
}
func (r NutritionOrderSupplement) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Type != nil {
		s += r.Type.MemSize()
	}
	if r.ProductName != nil {
		s += r.ProductName.MemSize()
	}
	if r.Schedule != nil {
		s += r.Schedule.MemSize()
	}
	if r.Quantity != nil {
		s += r.Quantity.MemSize()
	}
	if r.Instruction != nil {
		s += r.Instruction.MemSize()
	}
	return s
}
func (r NutritionOrderSupplementSchedule) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Timing {
		s += i.MemSize()
	}
	s += (cap(r.Timing) - len(r.Timing)) * int(unsafe.Sizeof(Timing{}))
	if r.AsNeeded != nil {
		s += r.AsNeeded.MemSize()
	}
	if r.AsNeededFor != nil {
		s += r.AsNeededFor.MemSize()
	}
	return s
}
func (r NutritionOrderEnteralFormula) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.BaseFormulaType != nil {
		s += r.BaseFormulaType.MemSize()
	}
	if r.BaseFormulaProductName != nil {
		s += r.BaseFormulaProductName.MemSize()
	}
	for _, i := range r.DeliveryDevice {
		s += i.MemSize()
	}
	s += (cap(r.DeliveryDevice) - len(r.DeliveryDevice)) * int(unsafe.Sizeof(CodeableReference{}))
	for _, i := range r.Additive {
		s += i.MemSize()
	}
	s += (cap(r.Additive) - len(r.Additive)) * int(unsafe.Sizeof(NutritionOrderEnteralFormulaAdditive{}))
	if r.CaloricDensity != nil {
		s += r.CaloricDensity.MemSize()
	}
	if r.RouteOfAdministration != nil {
		s += r.RouteOfAdministration.MemSize()
	}
	for _, i := range r.Administration {
		s += i.MemSize()
	}
	s += (cap(r.Administration) - len(r.Administration)) * int(unsafe.Sizeof(NutritionOrderEnteralFormulaAdministration{}))
	if r.MaxVolumeToDeliver != nil {
		s += r.MaxVolumeToDeliver.MemSize()
	}
	if r.AdministrationInstruction != nil {
		s += r.AdministrationInstruction.MemSize()
	}
	return s
}
func (r NutritionOrderEnteralFormulaAdditive) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Type != nil {
		s += r.Type.MemSize()
	}
	if r.ProductName != nil {
		s += r.ProductName.MemSize()
	}
	if r.Quantity != nil {
		s += r.Quantity.MemSize()
	}
	return s
}
func (r NutritionOrderEnteralFormulaAdministration) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	if r.Schedule != nil {
		s += r.Schedule.MemSize()
	}
	if r.Quantity != nil {
		s += r.Quantity.MemSize()
	}
	if r.Rate != nil {
		s += r.Rate.MemSize()
	}
	return s
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) MemSize() int {
	s := int(unsafe.Sizeof(r))
	if r.Id != nil {
		s += len(*r.Id) + int(unsafe.Sizeof(*r.Id))
	}
	for _, i := range r.Extension {
		s += i.MemSize()
	}
	s += (cap(r.Extension) - len(r.Extension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.ModifierExtension {
		s += i.MemSize()
	}
	s += (cap(r.ModifierExtension) - len(r.ModifierExtension)) * int(unsafe.Sizeof(Extension{}))
	for _, i := range r.Timing {
		s += i.MemSize()
	}
	s += (cap(r.Timing) - len(r.Timing)) * int(unsafe.Sizeof(Timing{}))
	if r.AsNeeded != nil {
		s += r.AsNeeded.MemSize()
	}
	if r.AsNeededFor != nil {
		s += r.AsNeededFor.MemSize()
	}
	return s
}
func (r NutritionOrder) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r NutritionOrderOralDiet) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r NutritionOrderOralDietSchedule) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r NutritionOrderOralDietNutrient) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r NutritionOrderOralDietTexture) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r NutritionOrderSupplement) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r NutritionOrderSupplementSchedule) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r NutritionOrderEnteralFormula) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r NutritionOrderEnteralFormulaAdditive) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r NutritionOrderEnteralFormulaAdministration) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "null"
	}
	return string(buf)
}
func (r NutritionOrder) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r NutritionOrder) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\"resourceType\":\"NutritionOrder\""))
	if err != nil {
		return err
	}
	setComma := true
	if r.Id != nil && r.Id.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		p := primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_id\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Meta != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"meta\":"))
		if err != nil {
			return err
		}
		err = r.Meta.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"implicitRules\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ImplicitRules)
		if err != nil {
			return err
		}
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		p := primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_implicitRules\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && r.Language.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"language\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Language)
		if err != nil {
			return err
		}
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		p := primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_language\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Text != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"text\":"))
		if err != nil {
			return err
		}
		err = r.Text.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Contained) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"contained\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, c := range r.Contained {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = ContainedResource{c}.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Identifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"identifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Identifier {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	{
		anyValue := false
		for _, e := range r.InstantiatesCanonical {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"instantiatesCanonical\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.InstantiatesCanonical)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.InstantiatesCanonical {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_instantiatesCanonical\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.InstantiatesCanonical {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	{
		anyValue := false
		for _, e := range r.InstantiatesUri {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"instantiatesUri\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.InstantiatesUri)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.InstantiatesUri {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_instantiatesUri\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.InstantiatesUri {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	{
		anyValue := false
		for _, e := range r.Instantiates {
			if e.Value != nil {
				anyValue = true
				break
			}
		}
		if anyValue {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"instantiates\":"))
			if err != nil {
				return err
			}
			enc := json.NewEncoder(w)
			enc.SetEscapeHTML(false)
			err := enc.Encode(r.Instantiates)
			if err != nil {
				return err
			}
		}
		anyIdOrExtension := false
		for _, e := range r.Instantiates {
			if e.Id != nil || e.Extension != nil {
				anyIdOrExtension = true
				break
			}
		}
		if anyIdOrExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			_, err = w.Write([]byte("\"_instantiates\":"))
			if err != nil {
				return err
			}
			_, err = w.Write([]byte("["))
			if err != nil {
				return err
			}
			setComma = false
			for _, e := range r.Instantiates {
				if e.Id != nil || e.Extension != nil {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					p := primitiveElement{Id: e.Id, Extension: e.Extension}
					err = p.marshalJSON(w)
					if err != nil {
						return err
					}
				} else {
					if setComma {
						_, err = w.Write([]byte(","))
						if err != nil {
							return err
						}
					}
					setComma = true
					_, err = w.Write([]byte("null"))
					if err != nil {
						return err
					}
				}
			}
			_, err = w.Write([]byte("]"))
			if err != nil {
				return err
			}
		}
	}
	if len(r.BasedOn) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"basedOn\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.BasedOn {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.GroupIdentifier != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"groupIdentifier\":"))
		if err != nil {
			return err
		}
		err = r.GroupIdentifier.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"status\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Status)
		if err != nil {
			return err
		}
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		p := primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_status\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"intent\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Intent)
		if err != nil {
			return err
		}
	}
	if r.Intent.Id != nil || r.Intent.Extension != nil {
		p := primitiveElement{Id: r.Intent.Id, Extension: r.Intent.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_intent\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Priority != nil && r.Priority.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"priority\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Priority)
		if err != nil {
			return err
		}
	}
	if r.Priority != nil && (r.Priority.Id != nil || r.Priority.Extension != nil) {
		p := primitiveElement{Id: r.Priority.Id, Extension: r.Priority.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_priority\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if setComma {
		_, err = w.Write([]byte(","))
		if err != nil {
			return err
		}
	}
	setComma = true
	_, err = w.Write([]byte("\"subject\":"))
	if err != nil {
		return err
	}
	err = r.Subject.marshalJSON(w)
	if err != nil {
		return err
	}
	if r.Encounter != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"encounter\":"))
		if err != nil {
			return err
		}
		err = r.Encounter.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.SupportingInformation) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"supportingInformation\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.SupportingInformation {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	{
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"dateTime\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.DateTime)
		if err != nil {
			return err
		}
	}
	if r.DateTime.Id != nil || r.DateTime.Extension != nil {
		p := primitiveElement{Id: r.DateTime.Id, Extension: r.DateTime.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_dateTime\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Orderer != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"orderer\":"))
		if err != nil {
			return err
		}
		err = r.Orderer.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Performer) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"performer\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Performer {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.AllergyIntolerance) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"allergyIntolerance\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.AllergyIntolerance {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.FoodPreferenceModifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"foodPreferenceModifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.FoodPreferenceModifier {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ExcludeFoodModifier) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"excludeFoodModifier\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ExcludeFoodModifier {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.OutsideFoodAllowed != nil && r.OutsideFoodAllowed.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"outsideFoodAllowed\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.OutsideFoodAllowed)
		if err != nil {
			return err
		}
	}
	if r.OutsideFoodAllowed != nil && (r.OutsideFoodAllowed.Id != nil || r.OutsideFoodAllowed.Extension != nil) {
		p := primitiveElement{Id: r.OutsideFoodAllowed.Id, Extension: r.OutsideFoodAllowed.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_outsideFoodAllowed\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.OralDiet != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"oralDiet\":"))
		if err != nil {
			return err
		}
		err = r.OralDiet.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Supplement) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"supplement\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Supplement {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.EnteralFormula != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"enteralFormula\":"))
		if err != nil {
			return err
		}
		err = r.EnteralFormula.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Note) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"note\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Note {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderOralDiet) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r NutritionOrderOralDiet) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Type) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"type\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Type {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Schedule != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"schedule\":"))
		if err != nil {
			return err
		}
		err = r.Schedule.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Nutrient) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"nutrient\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Nutrient {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Texture) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"texture\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Texture {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.FluidConsistencyType) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"fluidConsistencyType\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.FluidConsistencyType {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Instruction != nil && r.Instruction.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"instruction\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Instruction)
		if err != nil {
			return err
		}
	}
	if r.Instruction != nil && (r.Instruction.Id != nil || r.Instruction.Extension != nil) {
		p := primitiveElement{Id: r.Instruction.Id, Extension: r.Instruction.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_instruction\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderOralDietSchedule) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r NutritionOrderOralDietSchedule) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Timing) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timing\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Timing {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.AsNeeded != nil && r.AsNeeded.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"asNeeded\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.AsNeeded)
		if err != nil {
			return err
		}
	}
	if r.AsNeeded != nil && (r.AsNeeded.Id != nil || r.AsNeeded.Extension != nil) {
		p := primitiveElement{Id: r.AsNeeded.Id, Extension: r.AsNeeded.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_asNeeded\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.AsNeededFor != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"asNeededFor\":"))
		if err != nil {
			return err
		}
		err = r.AsNeededFor.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderOralDietNutrient) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r NutritionOrderOralDietNutrient) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Modifier != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifier\":"))
		if err != nil {
			return err
		}
		err = r.Modifier.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Amount != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"amount\":"))
		if err != nil {
			return err
		}
		err = r.Amount.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderOralDietTexture) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r NutritionOrderOralDietTexture) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Modifier != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifier\":"))
		if err != nil {
			return err
		}
		err = r.Modifier.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.FoodType != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"foodType\":"))
		if err != nil {
			return err
		}
		err = r.FoodType.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderSupplement) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r NutritionOrderSupplement) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Type != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"type\":"))
		if err != nil {
			return err
		}
		err = r.Type.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ProductName != nil && r.ProductName.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"productName\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ProductName)
		if err != nil {
			return err
		}
	}
	if r.ProductName != nil && (r.ProductName.Id != nil || r.ProductName.Extension != nil) {
		p := primitiveElement{Id: r.ProductName.Id, Extension: r.ProductName.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_productName\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Schedule != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"schedule\":"))
		if err != nil {
			return err
		}
		err = r.Schedule.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Quantity != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"quantity\":"))
		if err != nil {
			return err
		}
		err = r.Quantity.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Instruction != nil && r.Instruction.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"instruction\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Instruction)
		if err != nil {
			return err
		}
	}
	if r.Instruction != nil && (r.Instruction.Id != nil || r.Instruction.Extension != nil) {
		p := primitiveElement{Id: r.Instruction.Id, Extension: r.Instruction.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_instruction\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderSupplementSchedule) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r NutritionOrderSupplementSchedule) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Timing) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timing\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Timing {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.AsNeeded != nil && r.AsNeeded.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"asNeeded\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.AsNeeded)
		if err != nil {
			return err
		}
	}
	if r.AsNeeded != nil && (r.AsNeeded.Id != nil || r.AsNeeded.Extension != nil) {
		p := primitiveElement{Id: r.AsNeeded.Id, Extension: r.AsNeeded.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_asNeeded\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.AsNeededFor != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"asNeededFor\":"))
		if err != nil {
			return err
		}
		err = r.AsNeededFor.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderEnteralFormula) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r NutritionOrderEnteralFormula) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.BaseFormulaType != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"baseFormulaType\":"))
		if err != nil {
			return err
		}
		err = r.BaseFormulaType.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.BaseFormulaProductName != nil && r.BaseFormulaProductName.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"baseFormulaProductName\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.BaseFormulaProductName)
		if err != nil {
			return err
		}
	}
	if r.BaseFormulaProductName != nil && (r.BaseFormulaProductName.Id != nil || r.BaseFormulaProductName.Extension != nil) {
		p := primitiveElement{Id: r.BaseFormulaProductName.Id, Extension: r.BaseFormulaProductName.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_baseFormulaProductName\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.DeliveryDevice) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"deliveryDevice\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.DeliveryDevice {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Additive) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"additive\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Additive {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.CaloricDensity != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"caloricDensity\":"))
		if err != nil {
			return err
		}
		err = r.CaloricDensity.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.RouteOfAdministration != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"routeOfAdministration\":"))
		if err != nil {
			return err
		}
		err = r.RouteOfAdministration.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if len(r.Administration) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"administration\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Administration {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.MaxVolumeToDeliver != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"maxVolumeToDeliver\":"))
		if err != nil {
			return err
		}
		err = r.MaxVolumeToDeliver.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.AdministrationInstruction != nil && r.AdministrationInstruction.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"administrationInstruction\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.AdministrationInstruction)
		if err != nil {
			return err
		}
	}
	if r.AdministrationInstruction != nil && (r.AdministrationInstruction.Id != nil || r.AdministrationInstruction.Extension != nil) {
		p := primitiveElement{Id: r.AdministrationInstruction.Id, Extension: r.AdministrationInstruction.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_administrationInstruction\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderEnteralFormulaAdditive) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r NutritionOrderEnteralFormulaAdditive) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Type != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"type\":"))
		if err != nil {
			return err
		}
		err = r.Type.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.ProductName != nil && r.ProductName.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"productName\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.ProductName)
		if err != nil {
			return err
		}
	}
	if r.ProductName != nil && (r.ProductName.Id != nil || r.ProductName.Extension != nil) {
		p := primitiveElement{Id: r.ProductName.Id, Extension: r.ProductName.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_productName\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Quantity != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"quantity\":"))
		if err != nil {
			return err
		}
		err = r.Quantity.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderEnteralFormulaAdministration) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r NutritionOrderEnteralFormulaAdministration) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.Schedule != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"schedule\":"))
		if err != nil {
			return err
		}
		err = r.Schedule.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.Quantity != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"quantity\":"))
		if err != nil {
			return err
		}
		err = r.Quantity.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	switch v := r.Rate.(type) {
	case Quantity:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"rateQuantity\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Quantity:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"rateQuantity\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case Ratio:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"rateRatio\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	case *Ratio:
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"rateRatio\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	err := r.marshalJSON(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) marshalJSON(w io.Writer) error {
	var err error
	_, err = w.Write([]byte("{"))
	if err != nil {
		return err
	}
	setComma := false
	if r.Id != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"id\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.Id)
		if err != nil {
			return err
		}
	}
	if len(r.Extension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"extension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Extension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.ModifierExtension) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"modifierExtension\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.ModifierExtension {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if len(r.Timing) > 0 {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"timing\":"))
		if err != nil {
			return err
		}
		_, err = w.Write([]byte("["))
		if err != nil {
			return err
		}
		setComma = false
		for _, e := range r.Timing {
			if setComma {
				_, err = w.Write([]byte(","))
				if err != nil {
					return err
				}
			}
			setComma = true
			err = e.marshalJSON(w)
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("]"))
		if err != nil {
			return err
		}
	}
	if r.AsNeeded != nil && r.AsNeeded.Value != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"asNeeded\":"))
		if err != nil {
			return err
		}
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		err := enc.Encode(r.AsNeeded)
		if err != nil {
			return err
		}
	}
	if r.AsNeeded != nil && (r.AsNeeded.Id != nil || r.AsNeeded.Extension != nil) {
		p := primitiveElement{Id: r.AsNeeded.Id, Extension: r.AsNeeded.Extension}
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"_asNeeded\":"))
		if err != nil {
			return err
		}
		err = p.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	if r.AsNeededFor != nil {
		if setComma {
			_, err = w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		setComma = true
		_, err = w.Write([]byte("\"asNeededFor\":"))
		if err != nil {
			return err
		}
		err = r.AsNeededFor.marshalJSON(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte("}"))
	if err != nil {
		return err
	}
	return nil
}
func (r *NutritionOrder) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return r.unmarshalJSON(d)
}
func (r *NutritionOrder) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in NutritionOrder element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in NutritionOrder element", t)
		}
		switch f {
		case "resourceType":
			_, err := d.Token()
			if err != nil {
				return err
			}
		case "id":
			var v Id
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Value = v.Value
		case "_id":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Id == nil {
				r.Id = &Id{}
			}
			r.Id.Id = v.Id
			r.Id.Extension = v.Extension
		case "meta":
			var v Meta
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Meta = &v
		case "implicitRules":
			var v Uri
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Value = v.Value
		case "_implicitRules":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ImplicitRules == nil {
				r.ImplicitRules = &Uri{}
			}
			r.ImplicitRules.Id = v.Id
			r.ImplicitRules.Extension = v.Extension
		case "language":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Value = v.Value
		case "_language":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Language == nil {
				r.Language = &Code{}
			}
			r.Language.Id = v.Id
			r.Language.Extension = v.Extension
		case "text":
			var v Narrative
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Text = &v
		case "contained":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for d.More() {
				var v ContainedResource
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Contained = append(r.Contained, v.Resource)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "identifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for d.More() {
				var v Identifier
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Identifier = append(r.Identifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "instantiatesCanonical":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for i := 0; d.More(); i++ {
				var v Canonical
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.InstantiatesCanonical) <= i {
					r.InstantiatesCanonical = append(r.InstantiatesCanonical, Canonical{})
				}
				r.InstantiatesCanonical[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "_instantiatesCanonical":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.InstantiatesCanonical) <= i {
					r.InstantiatesCanonical = append(r.InstantiatesCanonical, Canonical{})
				}
				r.InstantiatesCanonical[i].Id = v.Id
				r.InstantiatesCanonical[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "instantiatesUri":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for i := 0; d.More(); i++ {
				var v Uri
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.InstantiatesUri) <= i {
					r.InstantiatesUri = append(r.InstantiatesUri, Uri{})
				}
				r.InstantiatesUri[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "_instantiatesUri":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.InstantiatesUri) <= i {
					r.InstantiatesUri = append(r.InstantiatesUri, Uri{})
				}
				r.InstantiatesUri[i].Id = v.Id
				r.InstantiatesUri[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "instantiates":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for i := 0; d.More(); i++ {
				var v Uri
				err := d.Decode(&v)
				if err != nil {
					return err
				}
				for len(r.Instantiates) <= i {
					r.Instantiates = append(r.Instantiates, Uri{})
				}
				r.Instantiates[i].Value = v.Value
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "_instantiates":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for i := 0; d.More(); i++ {
				var v primitiveElement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				for len(r.Instantiates) <= i {
					r.Instantiates = append(r.Instantiates, Uri{})
				}
				r.Instantiates[i].Id = v.Id
				r.Instantiates[i].Extension = v.Extension
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "basedOn":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.BasedOn = append(r.BasedOn, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "groupIdentifier":
			var v Identifier
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.GroupIdentifier = &v
		case "status":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Status.Value = v.Value
		case "_status":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Status.Id = v.Id
			r.Status.Extension = v.Extension
		case "intent":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Intent.Value = v.Value
		case "_intent":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Intent.Id = v.Id
			r.Intent.Extension = v.Extension
		case "priority":
			var v Code
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Priority == nil {
				r.Priority = &Code{}
			}
			r.Priority.Value = v.Value
		case "_priority":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Priority == nil {
				r.Priority = &Code{}
			}
			r.Priority.Id = v.Id
			r.Priority.Extension = v.Extension
		case "subject":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Subject = v
		case "encounter":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Encounter = &v
		case "supportingInformation":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.SupportingInformation = append(r.SupportingInformation, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "dateTime":
			var v DateTime
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.DateTime.Value = v.Value
		case "_dateTime":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.DateTime.Id = v.Id
			r.DateTime.Extension = v.Extension
		case "orderer":
			var v Reference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Orderer = &v
		case "performer":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for d.More() {
				var v CodeableReference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Performer = append(r.Performer, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "allergyIntolerance":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for d.More() {
				var v Reference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.AllergyIntolerance = append(r.AllergyIntolerance, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "foodPreferenceModifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.FoodPreferenceModifier = append(r.FoodPreferenceModifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "excludeFoodModifier":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ExcludeFoodModifier = append(r.ExcludeFoodModifier, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "outsideFoodAllowed":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.OutsideFoodAllowed == nil {
				r.OutsideFoodAllowed = &Boolean{}
			}
			r.OutsideFoodAllowed.Value = v.Value
		case "_outsideFoodAllowed":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.OutsideFoodAllowed == nil {
				r.OutsideFoodAllowed = &Boolean{}
			}
			r.OutsideFoodAllowed.Id = v.Id
			r.OutsideFoodAllowed.Extension = v.Extension
		case "oralDiet":
			var v NutritionOrderOralDiet
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.OralDiet = &v
		case "supplement":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for d.More() {
				var v NutritionOrderSupplement
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Supplement = append(r.Supplement, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		case "enteralFormula":
			var v NutritionOrderEnteralFormula
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.EnteralFormula = &v
		case "note":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrder element", t)
			}
			for d.More() {
				var v Annotation
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrder element", t)
			}
		default:
			return fmt.Errorf("invalid field: %s in NutritionOrder", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in NutritionOrder element", t)
	}
	return nil
}
func (r *NutritionOrderOralDiet) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in NutritionOrderOralDiet element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in NutritionOrderOralDiet element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderOralDiet element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderOralDiet element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderOralDiet element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderOralDiet element", t)
			}
		case "type":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderOralDiet element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Type = append(r.Type, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderOralDiet element", t)
			}
		case "schedule":
			var v NutritionOrderOralDietSchedule
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Schedule = &v
		case "nutrient":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderOralDiet element", t)
			}
			for d.More() {
				var v NutritionOrderOralDietNutrient
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Nutrient = append(r.Nutrient, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderOralDiet element", t)
			}
		case "texture":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderOralDiet element", t)
			}
			for d.More() {
				var v NutritionOrderOralDietTexture
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Texture = append(r.Texture, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderOralDiet element", t)
			}
		case "fluidConsistencyType":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderOralDiet element", t)
			}
			for d.More() {
				var v CodeableConcept
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.FluidConsistencyType = append(r.FluidConsistencyType, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderOralDiet element", t)
			}
		case "instruction":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Instruction == nil {
				r.Instruction = &String{}
			}
			r.Instruction.Value = v.Value
		case "_instruction":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Instruction == nil {
				r.Instruction = &String{}
			}
			r.Instruction.Id = v.Id
			r.Instruction.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in NutritionOrderOralDiet", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in NutritionOrderOralDiet element", t)
	}
	return nil
}
func (r *NutritionOrderOralDietSchedule) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in NutritionOrderOralDietSchedule element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in NutritionOrderOralDietSchedule element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderOralDietSchedule element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderOralDietSchedule element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderOralDietSchedule element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderOralDietSchedule element", t)
			}
		case "timing":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderOralDietSchedule element", t)
			}
			for d.More() {
				var v Timing
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Timing = append(r.Timing, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderOralDietSchedule element", t)
			}
		case "asNeeded":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.AsNeeded == nil {
				r.AsNeeded = &Boolean{}
			}
			r.AsNeeded.Value = v.Value
		case "_asNeeded":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.AsNeeded == nil {
				r.AsNeeded = &Boolean{}
			}
			r.AsNeeded.Id = v.Id
			r.AsNeeded.Extension = v.Extension
		case "asNeededFor":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.AsNeededFor = &v
		default:
			return fmt.Errorf("invalid field: %s in NutritionOrderOralDietSchedule", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in NutritionOrderOralDietSchedule element", t)
	}
	return nil
}
func (r *NutritionOrderOralDietNutrient) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in NutritionOrderOralDietNutrient element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in NutritionOrderOralDietNutrient element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderOralDietNutrient element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderOralDietNutrient element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderOralDietNutrient element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderOralDietNutrient element", t)
			}
		case "modifier":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Modifier = &v
		case "amount":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Amount = &v
		default:
			return fmt.Errorf("invalid field: %s in NutritionOrderOralDietNutrient", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in NutritionOrderOralDietNutrient element", t)
	}
	return nil
}
func (r *NutritionOrderOralDietTexture) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in NutritionOrderOralDietTexture element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in NutritionOrderOralDietTexture element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderOralDietTexture element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderOralDietTexture element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderOralDietTexture element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderOralDietTexture element", t)
			}
		case "modifier":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Modifier = &v
		case "foodType":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.FoodType = &v
		default:
			return fmt.Errorf("invalid field: %s in NutritionOrderOralDietTexture", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in NutritionOrderOralDietTexture element", t)
	}
	return nil
}
func (r *NutritionOrderSupplement) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in NutritionOrderSupplement element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in NutritionOrderSupplement element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderSupplement element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderSupplement element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderSupplement element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderSupplement element", t)
			}
		case "type":
			var v CodeableReference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = &v
		case "productName":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ProductName == nil {
				r.ProductName = &String{}
			}
			r.ProductName.Value = v.Value
		case "_productName":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ProductName == nil {
				r.ProductName = &String{}
			}
			r.ProductName.Id = v.Id
			r.ProductName.Extension = v.Extension
		case "schedule":
			var v NutritionOrderSupplementSchedule
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Schedule = &v
		case "quantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Quantity = &v
		case "instruction":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.Instruction == nil {
				r.Instruction = &String{}
			}
			r.Instruction.Value = v.Value
		case "_instruction":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.Instruction == nil {
				r.Instruction = &String{}
			}
			r.Instruction.Id = v.Id
			r.Instruction.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in NutritionOrderSupplement", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in NutritionOrderSupplement element", t)
	}
	return nil
}
func (r *NutritionOrderSupplementSchedule) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in NutritionOrderSupplementSchedule element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in NutritionOrderSupplementSchedule element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderSupplementSchedule element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderSupplementSchedule element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderSupplementSchedule element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderSupplementSchedule element", t)
			}
		case "timing":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderSupplementSchedule element", t)
			}
			for d.More() {
				var v Timing
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Timing = append(r.Timing, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderSupplementSchedule element", t)
			}
		case "asNeeded":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.AsNeeded == nil {
				r.AsNeeded = &Boolean{}
			}
			r.AsNeeded.Value = v.Value
		case "_asNeeded":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.AsNeeded == nil {
				r.AsNeeded = &Boolean{}
			}
			r.AsNeeded.Id = v.Id
			r.AsNeeded.Extension = v.Extension
		case "asNeededFor":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.AsNeededFor = &v
		default:
			return fmt.Errorf("invalid field: %s in NutritionOrderSupplementSchedule", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in NutritionOrderSupplementSchedule element", t)
	}
	return nil
}
func (r *NutritionOrderEnteralFormula) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in NutritionOrderEnteralFormula element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in NutritionOrderEnteralFormula element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderEnteralFormula element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderEnteralFormula element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderEnteralFormula element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderEnteralFormula element", t)
			}
		case "baseFormulaType":
			var v CodeableReference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.BaseFormulaType = &v
		case "baseFormulaProductName":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.BaseFormulaProductName == nil {
				r.BaseFormulaProductName = &String{}
			}
			r.BaseFormulaProductName.Value = v.Value
		case "_baseFormulaProductName":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.BaseFormulaProductName == nil {
				r.BaseFormulaProductName = &String{}
			}
			r.BaseFormulaProductName.Id = v.Id
			r.BaseFormulaProductName.Extension = v.Extension
		case "deliveryDevice":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderEnteralFormula element", t)
			}
			for d.More() {
				var v CodeableReference
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.DeliveryDevice = append(r.DeliveryDevice, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderEnteralFormula element", t)
			}
		case "additive":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderEnteralFormula element", t)
			}
			for d.More() {
				var v NutritionOrderEnteralFormulaAdditive
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Additive = append(r.Additive, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderEnteralFormula element", t)
			}
		case "caloricDensity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.CaloricDensity = &v
		case "routeOfAdministration":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.RouteOfAdministration = &v
		case "administration":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderEnteralFormula element", t)
			}
			for d.More() {
				var v NutritionOrderEnteralFormulaAdministration
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Administration = append(r.Administration, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderEnteralFormula element", t)
			}
		case "maxVolumeToDeliver":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.MaxVolumeToDeliver = &v
		case "administrationInstruction":
			var v Markdown
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.AdministrationInstruction == nil {
				r.AdministrationInstruction = &Markdown{}
			}
			r.AdministrationInstruction.Value = v.Value
		case "_administrationInstruction":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.AdministrationInstruction == nil {
				r.AdministrationInstruction = &Markdown{}
			}
			r.AdministrationInstruction.Id = v.Id
			r.AdministrationInstruction.Extension = v.Extension
		default:
			return fmt.Errorf("invalid field: %s in NutritionOrderEnteralFormula", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in NutritionOrderEnteralFormula element", t)
	}
	return nil
}
func (r *NutritionOrderEnteralFormulaAdditive) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in NutritionOrderEnteralFormulaAdditive element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in NutritionOrderEnteralFormulaAdditive element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderEnteralFormulaAdditive element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderEnteralFormulaAdditive element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderEnteralFormulaAdditive element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderEnteralFormulaAdditive element", t)
			}
		case "type":
			var v CodeableReference
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Type = &v
		case "productName":
			var v String
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.ProductName == nil {
				r.ProductName = &String{}
			}
			r.ProductName.Value = v.Value
		case "_productName":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.ProductName == nil {
				r.ProductName = &String{}
			}
			r.ProductName.Id = v.Id
			r.ProductName.Extension = v.Extension
		case "quantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Quantity = &v
		default:
			return fmt.Errorf("invalid field: %s in NutritionOrderEnteralFormulaAdditive", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in NutritionOrderEnteralFormulaAdditive element", t)
	}
	return nil
}
func (r *NutritionOrderEnteralFormulaAdministration) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in NutritionOrderEnteralFormulaAdministration element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in NutritionOrderEnteralFormulaAdministration element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderEnteralFormulaAdministration element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderEnteralFormulaAdministration element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderEnteralFormulaAdministration element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderEnteralFormulaAdministration element", t)
			}
		case "schedule":
			var v NutritionOrderEnteralFormulaAdministrationSchedule
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Schedule = &v
		case "quantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Quantity = &v
		case "rateQuantity":
			var v Quantity
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Rate = v
		case "rateRatio":
			var v Ratio
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.Rate = v
		default:
			return fmt.Errorf("invalid field: %s in NutritionOrderEnteralFormulaAdministration", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in NutritionOrderEnteralFormulaAdministration element", t)
	}
	return nil
}
func (r *NutritionOrderEnteralFormulaAdministrationSchedule) unmarshalJSON(d *json.Decoder) error {
	t, err := d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("invalid token: %v, expected: '{' in NutritionOrderEnteralFormulaAdministrationSchedule element", t)
	}
	for d.More() {
		t, err = d.Token()
		if err != nil {
			return err
		}
		f, ok := t.(string)
		if !ok {
			return fmt.Errorf("invalid token: %v, expected: field name in NutritionOrderEnteralFormulaAdministrationSchedule element", t)
		}
		switch f {
		case "id":
			var v string
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			r.Id = &v
		case "extension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderEnteralFormulaAdministrationSchedule element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Extension = append(r.Extension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderEnteralFormulaAdministrationSchedule element", t)
			}
		case "modifierExtension":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderEnteralFormulaAdministrationSchedule element", t)
			}
			for d.More() {
				var v Extension
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.ModifierExtension = append(r.ModifierExtension, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderEnteralFormulaAdministrationSchedule element", t)
			}
		case "timing":
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim('[') {
				return fmt.Errorf("invalid token: %v, expected: '[' in NutritionOrderEnteralFormulaAdministrationSchedule element", t)
			}
			for d.More() {
				var v Timing
				err := v.unmarshalJSON(d)
				if err != nil {
					return err
				}
				r.Timing = append(r.Timing, v)
			}
			t, err = d.Token()
			if err != nil {
				return err
			}
			if t != json.Delim(']') {
				return fmt.Errorf("invalid token: %v, expected: ']' in NutritionOrderEnteralFormulaAdministrationSchedule element", t)
			}
		case "asNeeded":
			var v Boolean
			err := d.Decode(&v)
			if err != nil {
				return err
			}
			if r.AsNeeded == nil {
				r.AsNeeded = &Boolean{}
			}
			r.AsNeeded.Value = v.Value
		case "_asNeeded":
			var v primitiveElement
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			if r.AsNeeded == nil {
				r.AsNeeded = &Boolean{}
			}
			r.AsNeeded.Id = v.Id
			r.AsNeeded.Extension = v.Extension
		case "asNeededFor":
			var v CodeableConcept
			err := v.unmarshalJSON(d)
			if err != nil {
				return err
			}
			r.AsNeededFor = &v
		default:
			return fmt.Errorf("invalid field: %s in NutritionOrderEnteralFormulaAdministrationSchedule", f)
		}
	}
	t, err = d.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('}') {
		return fmt.Errorf("invalid token: %v, expected: '}' in NutritionOrderEnteralFormulaAdministrationSchedule element", t)
	}
	return nil
}
func (r NutritionOrder) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if start.Name.Local == "__contained__" {
		start.Name.Space = ""
	} else {
		start.Name.Space = "http://hl7.org/fhir"
	}
	start.Name.Local = "NutritionOrder"
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
	for _, c := range r.Contained {
		err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "contained"}})
		if err != nil {
			return err
		}
		err = e.EncodeElement(c, xml.StartElement{Name: xml.Name{Local: "__contained__"}})
		if err != nil {
			return err
		}
		err = e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "contained"}})
		if err != nil {
			return err
		}
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
	err = e.EncodeElement(r.InstantiatesCanonical, xml.StartElement{Name: xml.Name{Local: "instantiatesCanonical"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.InstantiatesUri, xml.StartElement{Name: xml.Name{Local: "instantiatesUri"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Instantiates, xml.StartElement{Name: xml.Name{Local: "instantiates"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BasedOn, xml.StartElement{Name: xml.Name{Local: "basedOn"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.GroupIdentifier, xml.StartElement{Name: xml.Name{Local: "groupIdentifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Status, xml.StartElement{Name: xml.Name{Local: "status"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Intent, xml.StartElement{Name: xml.Name{Local: "intent"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Priority, xml.StartElement{Name: xml.Name{Local: "priority"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Subject, xml.StartElement{Name: xml.Name{Local: "subject"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Encounter, xml.StartElement{Name: xml.Name{Local: "encounter"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.SupportingInformation, xml.StartElement{Name: xml.Name{Local: "supportingInformation"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DateTime, xml.StartElement{Name: xml.Name{Local: "dateTime"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Orderer, xml.StartElement{Name: xml.Name{Local: "orderer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Performer, xml.StartElement{Name: xml.Name{Local: "performer"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AllergyIntolerance, xml.StartElement{Name: xml.Name{Local: "allergyIntolerance"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FoodPreferenceModifier, xml.StartElement{Name: xml.Name{Local: "foodPreferenceModifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.ExcludeFoodModifier, xml.StartElement{Name: xml.Name{Local: "excludeFoodModifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OutsideFoodAllowed, xml.StartElement{Name: xml.Name{Local: "outsideFoodAllowed"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.OralDiet, xml.StartElement{Name: xml.Name{Local: "oralDiet"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Supplement, xml.StartElement{Name: xml.Name{Local: "supplement"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.EnteralFormula, xml.StartElement{Name: xml.Name{Local: "enteralFormula"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Note, xml.StartElement{Name: xml.Name{Local: "note"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderOralDiet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Schedule, xml.StartElement{Name: xml.Name{Local: "schedule"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Nutrient, xml.StartElement{Name: xml.Name{Local: "nutrient"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Texture, xml.StartElement{Name: xml.Name{Local: "texture"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FluidConsistencyType, xml.StartElement{Name: xml.Name{Local: "fluidConsistencyType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Instruction, xml.StartElement{Name: xml.Name{Local: "instruction"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderOralDietSchedule) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Timing, xml.StartElement{Name: xml.Name{Local: "timing"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AsNeeded, xml.StartElement{Name: xml.Name{Local: "asNeeded"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AsNeededFor, xml.StartElement{Name: xml.Name{Local: "asNeededFor"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderOralDietNutrient) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Modifier, xml.StartElement{Name: xml.Name{Local: "modifier"}})
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
func (r NutritionOrderOralDietTexture) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Modifier, xml.StartElement{Name: xml.Name{Local: "modifier"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.FoodType, xml.StartElement{Name: xml.Name{Local: "foodType"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderSupplement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ProductName, xml.StartElement{Name: xml.Name{Local: "productName"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Schedule, xml.StartElement{Name: xml.Name{Local: "schedule"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Instruction, xml.StartElement{Name: xml.Name{Local: "instruction"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderSupplementSchedule) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Timing, xml.StartElement{Name: xml.Name{Local: "timing"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AsNeeded, xml.StartElement{Name: xml.Name{Local: "asNeeded"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AsNeededFor, xml.StartElement{Name: xml.Name{Local: "asNeededFor"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderEnteralFormula) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.BaseFormulaType, xml.StartElement{Name: xml.Name{Local: "baseFormulaType"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.BaseFormulaProductName, xml.StartElement{Name: xml.Name{Local: "baseFormulaProductName"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.DeliveryDevice, xml.StartElement{Name: xml.Name{Local: "deliveryDevice"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Additive, xml.StartElement{Name: xml.Name{Local: "additive"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.CaloricDensity, xml.StartElement{Name: xml.Name{Local: "caloricDensity"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.RouteOfAdministration, xml.StartElement{Name: xml.Name{Local: "routeOfAdministration"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Administration, xml.StartElement{Name: xml.Name{Local: "administration"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.MaxVolumeToDeliver, xml.StartElement{Name: xml.Name{Local: "maxVolumeToDeliver"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AdministrationInstruction, xml.StartElement{Name: xml.Name{Local: "administrationInstruction"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderEnteralFormulaAdditive) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.ProductName, xml.StartElement{Name: xml.Name{Local: "productName"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderEnteralFormulaAdministration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Schedule, xml.StartElement{Name: xml.Name{Local: "schedule"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}})
	if err != nil {
		return err
	}
	switch v := r.Rate.(type) {
	case Quantity, *Quantity:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "rateQuantity"}})
		if err != nil {
			return err
		}
	case Ratio, *Ratio:
		err := e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: "rateRatio"}})
		if err != nil {
			return err
		}
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	err = e.EncodeElement(r.Timing, xml.StartElement{Name: xml.Name{Local: "timing"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AsNeeded, xml.StartElement{Name: xml.Name{Local: "asNeeded"}})
	if err != nil {
		return err
	}
	err = e.EncodeElement(r.AsNeededFor, xml.StartElement{Name: xml.Name{Local: "asNeededFor"}})
	if err != nil {
		return err
	}
	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}
func (r *NutritionOrder) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "instantiatesCanonical":
				var v Canonical
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.InstantiatesCanonical = append(r.InstantiatesCanonical, v)
			case "instantiatesUri":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.InstantiatesUri = append(r.InstantiatesUri, v)
			case "instantiates":
				var v Uri
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Instantiates = append(r.Instantiates, v)
			case "basedOn":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BasedOn = append(r.BasedOn, v)
			case "groupIdentifier":
				var v Identifier
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.GroupIdentifier = &v
			case "status":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Status = v
			case "intent":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Intent = v
			case "priority":
				var v Code
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Priority = &v
			case "subject":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Subject = v
			case "encounter":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Encounter = &v
			case "supportingInformation":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.SupportingInformation = append(r.SupportingInformation, v)
			case "dateTime":
				var v DateTime
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DateTime = v
			case "orderer":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Orderer = &v
			case "performer":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Performer = append(r.Performer, v)
			case "allergyIntolerance":
				var v Reference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AllergyIntolerance = append(r.AllergyIntolerance, v)
			case "foodPreferenceModifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FoodPreferenceModifier = append(r.FoodPreferenceModifier, v)
			case "excludeFoodModifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ExcludeFoodModifier = append(r.ExcludeFoodModifier, v)
			case "outsideFoodAllowed":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OutsideFoodAllowed = &v
			case "oralDiet":
				var v NutritionOrderOralDiet
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.OralDiet = &v
			case "supplement":
				var v NutritionOrderSupplement
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Supplement = append(r.Supplement, v)
			case "enteralFormula":
				var v NutritionOrderEnteralFormula
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.EnteralFormula = &v
			case "note":
				var v Annotation
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Note = append(r.Note, v)
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *NutritionOrderOralDiet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				r.Type = append(r.Type, v)
			case "schedule":
				var v NutritionOrderOralDietSchedule
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Schedule = &v
			case "nutrient":
				var v NutritionOrderOralDietNutrient
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Nutrient = append(r.Nutrient, v)
			case "texture":
				var v NutritionOrderOralDietTexture
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Texture = append(r.Texture, v)
			case "fluidConsistencyType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FluidConsistencyType = append(r.FluidConsistencyType, v)
			case "instruction":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Instruction = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *NutritionOrderOralDietSchedule) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "timing":
				var v Timing
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Timing = append(r.Timing, v)
			case "asNeeded":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AsNeeded = &v
			case "asNeededFor":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AsNeededFor = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *NutritionOrderOralDietNutrient) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "modifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Modifier = &v
			case "amount":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Amount = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *NutritionOrderOralDietTexture) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "modifier":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Modifier = &v
			case "foodType":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.FoodType = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *NutritionOrderSupplement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "productName":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductName = &v
			case "schedule":
				var v NutritionOrderSupplementSchedule
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Schedule = &v
			case "quantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = &v
			case "instruction":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Instruction = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *NutritionOrderSupplementSchedule) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "timing":
				var v Timing
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Timing = append(r.Timing, v)
			case "asNeeded":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AsNeeded = &v
			case "asNeededFor":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AsNeededFor = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *NutritionOrderEnteralFormula) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "baseFormulaType":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BaseFormulaType = &v
			case "baseFormulaProductName":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.BaseFormulaProductName = &v
			case "deliveryDevice":
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.DeliveryDevice = append(r.DeliveryDevice, v)
			case "additive":
				var v NutritionOrderEnteralFormulaAdditive
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Additive = append(r.Additive, v)
			case "caloricDensity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.CaloricDensity = &v
			case "routeOfAdministration":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.RouteOfAdministration = &v
			case "administration":
				var v NutritionOrderEnteralFormulaAdministration
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Administration = append(r.Administration, v)
			case "maxVolumeToDeliver":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.MaxVolumeToDeliver = &v
			case "administrationInstruction":
				var v Markdown
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AdministrationInstruction = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *NutritionOrderEnteralFormulaAdditive) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
				var v CodeableReference
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Type = &v
			case "productName":
				var v String
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.ProductName = &v
			case "quantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *NutritionOrderEnteralFormulaAdministration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "schedule":
				var v NutritionOrderEnteralFormulaAdministrationSchedule
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Schedule = &v
			case "quantity":
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Quantity = &v
			case "rateQuantity":
				if r.Rate != nil {
					return fmt.Errorf("multiple values for field \"Rate\"")
				}
				var v Quantity
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Rate = v
			case "rateRatio":
				if r.Rate != nil {
					return fmt.Errorf("multiple values for field \"Rate\"")
				}
				var v Ratio
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Rate = v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r *NutritionOrderEnteralFormulaAdministrationSchedule) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "timing":
				var v Timing
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.Timing = append(r.Timing, v)
			case "asNeeded":
				var v Boolean
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AsNeeded = &v
			case "asNeededFor":
				var v CodeableConcept
				err := d.DecodeElement(&v, &t)
				if err != nil {
					return err
				}
				r.AsNeededFor = &v
			}
		case xml.EndElement:
			return nil
		}
	}
}
func (r NutritionOrder) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, *r.Id)
		}
	}
	if len(name) == 0 || slices.Contains(name, "meta") {
		if r.Meta != nil {
			children = append(children, *r.Meta)
		}
	}
	if len(name) == 0 || slices.Contains(name, "implicitRules") {
		if r.ImplicitRules != nil {
			children = append(children, *r.ImplicitRules)
		}
	}
	if len(name) == 0 || slices.Contains(name, "language") {
		if r.Language != nil {
			children = append(children, *r.Language)
		}
	}
	if len(name) == 0 || slices.Contains(name, "text") {
		if r.Text != nil {
			children = append(children, *r.Text)
		}
	}
	if len(name) == 0 || slices.Contains(name, "contained") {
		for _, v := range r.Contained {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "identifier") {
		for _, v := range r.Identifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "instantiatesCanonical") {
		for _, v := range r.InstantiatesCanonical {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "instantiatesUri") {
		for _, v := range r.InstantiatesUri {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "instantiates") {
		for _, v := range r.Instantiates {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "basedOn") {
		for _, v := range r.BasedOn {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "groupIdentifier") {
		if r.GroupIdentifier != nil {
			children = append(children, *r.GroupIdentifier)
		}
	}
	if len(name) == 0 || slices.Contains(name, "status") {
		children = append(children, r.Status)
	}
	if len(name) == 0 || slices.Contains(name, "intent") {
		children = append(children, r.Intent)
	}
	if len(name) == 0 || slices.Contains(name, "priority") {
		if r.Priority != nil {
			children = append(children, *r.Priority)
		}
	}
	if len(name) == 0 || slices.Contains(name, "subject") {
		children = append(children, r.Subject)
	}
	if len(name) == 0 || slices.Contains(name, "encounter") {
		if r.Encounter != nil {
			children = append(children, *r.Encounter)
		}
	}
	if len(name) == 0 || slices.Contains(name, "supportingInformation") {
		for _, v := range r.SupportingInformation {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "dateTime") {
		children = append(children, r.DateTime)
	}
	if len(name) == 0 || slices.Contains(name, "orderer") {
		if r.Orderer != nil {
			children = append(children, *r.Orderer)
		}
	}
	if len(name) == 0 || slices.Contains(name, "performer") {
		for _, v := range r.Performer {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "allergyIntolerance") {
		for _, v := range r.AllergyIntolerance {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "foodPreferenceModifier") {
		for _, v := range r.FoodPreferenceModifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "excludeFoodModifier") {
		for _, v := range r.ExcludeFoodModifier {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "outsideFoodAllowed") {
		if r.OutsideFoodAllowed != nil {
			children = append(children, *r.OutsideFoodAllowed)
		}
	}
	if len(name) == 0 || slices.Contains(name, "oralDiet") {
		if r.OralDiet != nil {
			children = append(children, *r.OralDiet)
		}
	}
	if len(name) == 0 || slices.Contains(name, "supplement") {
		for _, v := range r.Supplement {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "enteralFormula") {
		if r.EnteralFormula != nil {
			children = append(children, *r.EnteralFormula)
		}
	}
	if len(name) == 0 || slices.Contains(name, "note") {
		for _, v := range r.Note {
			children = append(children, v)
		}
	}
	return children
}
func (r NutritionOrder) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert NutritionOrder to Boolean")
}
func (r NutritionOrder) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert NutritionOrder to String")
}
func (r NutritionOrder) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert NutritionOrder to Integer")
}
func (r NutritionOrder) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert NutritionOrder to Decimal")
}
func (r NutritionOrder) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert NutritionOrder to Date")
}
func (r NutritionOrder) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert NutritionOrder to Time")
}
func (r NutritionOrder) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert NutritionOrder to DateTime")
}
func (r NutritionOrder) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert NutritionOrder to Quantity")
}
func (r NutritionOrder) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.Id",
		}, {
			Name: "Meta",
			Type: "FHIR.Meta",
		}, {
			Name: "ImplicitRules",
			Type: "FHIR.Uri",
		}, {
			Name: "Language",
			Type: "FHIR.Code",
		}, {
			Name: "Text",
			Type: "FHIR.Narrative",
		}, {
			Name: "Contained",
			Type: "List<FHIR.>",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Identifier",
			Type: "List<FHIR.Identifier>",
		}, {
			Name: "InstantiatesCanonical",
			Type: "List<FHIR.Canonical>",
		}, {
			Name: "InstantiatesUri",
			Type: "List<FHIR.Uri>",
		}, {
			Name: "Instantiates",
			Type: "List<FHIR.Uri>",
		}, {
			Name: "BasedOn",
			Type: "List<FHIR.Reference>",
		}, {
			Name: "GroupIdentifier",
			Type: "FHIR.Identifier",
		}, {
			Name: "Status",
			Type: "FHIR.Code",
		}, {
			Name: "Intent",
			Type: "FHIR.Code",
		}, {
			Name: "Priority",
			Type: "FHIR.Code",
		}, {
			Name: "Subject",
			Type: "FHIR.Reference",
		}, {
			Name: "Encounter",
			Type: "FHIR.Reference",
		}, {
			Name: "SupportingInformation",
			Type: "List<FHIR.Reference>",
		}, {
			Name: "DateTime",
			Type: "FHIR.DateTime",
		}, {
			Name: "Orderer",
			Type: "FHIR.Reference",
		}, {
			Name: "Performer",
			Type: "List<FHIR.CodeableReference>",
		}, {
			Name: "AllergyIntolerance",
			Type: "List<FHIR.Reference>",
		}, {
			Name: "FoodPreferenceModifier",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "ExcludeFoodModifier",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "OutsideFoodAllowed",
			Type: "FHIR.Boolean",
		}, {
			Name: "OralDiet",
			Type: "FHIR.NutritionOrderOralDiet",
		}, {
			Name: "Supplement",
			Type: "List<FHIR.NutritionOrderSupplement>",
		}, {
			Name: "EnteralFormula",
			Type: "FHIR.NutritionOrderEnteralFormula",
		}, {
			Name: "Note",
			Type: "List<FHIR.Annotation>",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DomainResource",
				Namespace: "FHIR",
			},
			Name:      "NutritionOrder",
			Namespace: "FHIR",
		},
	}
}
func (r NutritionOrderOralDiet) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "type") {
		for _, v := range r.Type {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "schedule") {
		if r.Schedule != nil {
			children = append(children, *r.Schedule)
		}
	}
	if len(name) == 0 || slices.Contains(name, "nutrient") {
		for _, v := range r.Nutrient {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "texture") {
		for _, v := range r.Texture {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "fluidConsistencyType") {
		for _, v := range r.FluidConsistencyType {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "instruction") {
		if r.Instruction != nil {
			children = append(children, *r.Instruction)
		}
	}
	return children
}
func (r NutritionOrderOralDiet) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert NutritionOrderOralDiet to Boolean")
}
func (r NutritionOrderOralDiet) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert NutritionOrderOralDiet to String")
}
func (r NutritionOrderOralDiet) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert NutritionOrderOralDiet to Integer")
}
func (r NutritionOrderOralDiet) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert NutritionOrderOralDiet to Decimal")
}
func (r NutritionOrderOralDiet) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert NutritionOrderOralDiet to Date")
}
func (r NutritionOrderOralDiet) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert NutritionOrderOralDiet to Time")
}
func (r NutritionOrderOralDiet) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert NutritionOrderOralDiet to DateTime")
}
func (r NutritionOrderOralDiet) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert NutritionOrderOralDiet to Quantity")
}
func (r NutritionOrderOralDiet) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Type",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "Schedule",
			Type: "FHIR.NutritionOrderOralDietSchedule",
		}, {
			Name: "Nutrient",
			Type: "List<FHIR.NutritionOrderOralDietNutrient>",
		}, {
			Name: "Texture",
			Type: "List<FHIR.NutritionOrderOralDietTexture>",
		}, {
			Name: "FluidConsistencyType",
			Type: "List<FHIR.CodeableConcept>",
		}, {
			Name: "Instruction",
			Type: "FHIR.String",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "NutritionOrderOralDiet",
			Namespace: "FHIR",
		},
	}
}
func (r NutritionOrderOralDietSchedule) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "timing") {
		for _, v := range r.Timing {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "asNeeded") {
		if r.AsNeeded != nil {
			children = append(children, *r.AsNeeded)
		}
	}
	if len(name) == 0 || slices.Contains(name, "asNeededFor") {
		if r.AsNeededFor != nil {
			children = append(children, *r.AsNeededFor)
		}
	}
	return children
}
func (r NutritionOrderOralDietSchedule) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietSchedule to Boolean")
}
func (r NutritionOrderOralDietSchedule) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietSchedule to String")
}
func (r NutritionOrderOralDietSchedule) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietSchedule to Integer")
}
func (r NutritionOrderOralDietSchedule) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietSchedule to Decimal")
}
func (r NutritionOrderOralDietSchedule) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietSchedule to Date")
}
func (r NutritionOrderOralDietSchedule) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietSchedule to Time")
}
func (r NutritionOrderOralDietSchedule) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietSchedule to DateTime")
}
func (r NutritionOrderOralDietSchedule) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietSchedule to Quantity")
}
func (r NutritionOrderOralDietSchedule) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Timing",
			Type: "List<FHIR.Timing>",
		}, {
			Name: "AsNeeded",
			Type: "FHIR.Boolean",
		}, {
			Name: "AsNeededFor",
			Type: "FHIR.CodeableConcept",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "NutritionOrderOralDietSchedule",
			Namespace: "FHIR",
		},
	}
}
func (r NutritionOrderOralDietNutrient) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifier") {
		if r.Modifier != nil {
			children = append(children, *r.Modifier)
		}
	}
	if len(name) == 0 || slices.Contains(name, "amount") {
		if r.Amount != nil {
			children = append(children, *r.Amount)
		}
	}
	return children
}
func (r NutritionOrderOralDietNutrient) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietNutrient to Boolean")
}
func (r NutritionOrderOralDietNutrient) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietNutrient to String")
}
func (r NutritionOrderOralDietNutrient) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietNutrient to Integer")
}
func (r NutritionOrderOralDietNutrient) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietNutrient to Decimal")
}
func (r NutritionOrderOralDietNutrient) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietNutrient to Date")
}
func (r NutritionOrderOralDietNutrient) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietNutrient to Time")
}
func (r NutritionOrderOralDietNutrient) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietNutrient to DateTime")
}
func (r NutritionOrderOralDietNutrient) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietNutrient to Quantity")
}
func (r NutritionOrderOralDietNutrient) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Modifier",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Amount",
			Type: "FHIR.Quantity",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "NutritionOrderOralDietNutrient",
			Namespace: "FHIR",
		},
	}
}
func (r NutritionOrderOralDietTexture) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifier") {
		if r.Modifier != nil {
			children = append(children, *r.Modifier)
		}
	}
	if len(name) == 0 || slices.Contains(name, "foodType") {
		if r.FoodType != nil {
			children = append(children, *r.FoodType)
		}
	}
	return children
}
func (r NutritionOrderOralDietTexture) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietTexture to Boolean")
}
func (r NutritionOrderOralDietTexture) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietTexture to String")
}
func (r NutritionOrderOralDietTexture) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietTexture to Integer")
}
func (r NutritionOrderOralDietTexture) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietTexture to Decimal")
}
func (r NutritionOrderOralDietTexture) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietTexture to Date")
}
func (r NutritionOrderOralDietTexture) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietTexture to Time")
}
func (r NutritionOrderOralDietTexture) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietTexture to DateTime")
}
func (r NutritionOrderOralDietTexture) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert NutritionOrderOralDietTexture to Quantity")
}
func (r NutritionOrderOralDietTexture) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Modifier",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "FoodType",
			Type: "FHIR.CodeableConcept",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "NutritionOrderOralDietTexture",
			Namespace: "FHIR",
		},
	}
}
func (r NutritionOrderSupplement) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "type") {
		if r.Type != nil {
			children = append(children, *r.Type)
		}
	}
	if len(name) == 0 || slices.Contains(name, "productName") {
		if r.ProductName != nil {
			children = append(children, *r.ProductName)
		}
	}
	if len(name) == 0 || slices.Contains(name, "schedule") {
		if r.Schedule != nil {
			children = append(children, *r.Schedule)
		}
	}
	if len(name) == 0 || slices.Contains(name, "quantity") {
		if r.Quantity != nil {
			children = append(children, *r.Quantity)
		}
	}
	if len(name) == 0 || slices.Contains(name, "instruction") {
		if r.Instruction != nil {
			children = append(children, *r.Instruction)
		}
	}
	return children
}
func (r NutritionOrderSupplement) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert NutritionOrderSupplement to Boolean")
}
func (r NutritionOrderSupplement) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert NutritionOrderSupplement to String")
}
func (r NutritionOrderSupplement) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert NutritionOrderSupplement to Integer")
}
func (r NutritionOrderSupplement) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert NutritionOrderSupplement to Decimal")
}
func (r NutritionOrderSupplement) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert NutritionOrderSupplement to Date")
}
func (r NutritionOrderSupplement) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert NutritionOrderSupplement to Time")
}
func (r NutritionOrderSupplement) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert NutritionOrderSupplement to DateTime")
}
func (r NutritionOrderSupplement) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert NutritionOrderSupplement to Quantity")
}
func (r NutritionOrderSupplement) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Type",
			Type: "FHIR.CodeableReference",
		}, {
			Name: "ProductName",
			Type: "FHIR.String",
		}, {
			Name: "Schedule",
			Type: "FHIR.NutritionOrderSupplementSchedule",
		}, {
			Name: "Quantity",
			Type: "FHIR.Quantity",
		}, {
			Name: "Instruction",
			Type: "FHIR.String",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "NutritionOrderSupplement",
			Namespace: "FHIR",
		},
	}
}
func (r NutritionOrderSupplementSchedule) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "timing") {
		for _, v := range r.Timing {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "asNeeded") {
		if r.AsNeeded != nil {
			children = append(children, *r.AsNeeded)
		}
	}
	if len(name) == 0 || slices.Contains(name, "asNeededFor") {
		if r.AsNeededFor != nil {
			children = append(children, *r.AsNeededFor)
		}
	}
	return children
}
func (r NutritionOrderSupplementSchedule) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert NutritionOrderSupplementSchedule to Boolean")
}
func (r NutritionOrderSupplementSchedule) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert NutritionOrderSupplementSchedule to String")
}
func (r NutritionOrderSupplementSchedule) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert NutritionOrderSupplementSchedule to Integer")
}
func (r NutritionOrderSupplementSchedule) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert NutritionOrderSupplementSchedule to Decimal")
}
func (r NutritionOrderSupplementSchedule) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert NutritionOrderSupplementSchedule to Date")
}
func (r NutritionOrderSupplementSchedule) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert NutritionOrderSupplementSchedule to Time")
}
func (r NutritionOrderSupplementSchedule) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert NutritionOrderSupplementSchedule to DateTime")
}
func (r NutritionOrderSupplementSchedule) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert NutritionOrderSupplementSchedule to Quantity")
}
func (r NutritionOrderSupplementSchedule) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Timing",
			Type: "List<FHIR.Timing>",
		}, {
			Name: "AsNeeded",
			Type: "FHIR.Boolean",
		}, {
			Name: "AsNeededFor",
			Type: "FHIR.CodeableConcept",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "NutritionOrderSupplementSchedule",
			Namespace: "FHIR",
		},
	}
}
func (r NutritionOrderEnteralFormula) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "baseFormulaType") {
		if r.BaseFormulaType != nil {
			children = append(children, *r.BaseFormulaType)
		}
	}
	if len(name) == 0 || slices.Contains(name, "baseFormulaProductName") {
		if r.BaseFormulaProductName != nil {
			children = append(children, *r.BaseFormulaProductName)
		}
	}
	if len(name) == 0 || slices.Contains(name, "deliveryDevice") {
		for _, v := range r.DeliveryDevice {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "additive") {
		for _, v := range r.Additive {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "caloricDensity") {
		if r.CaloricDensity != nil {
			children = append(children, *r.CaloricDensity)
		}
	}
	if len(name) == 0 || slices.Contains(name, "routeOfAdministration") {
		if r.RouteOfAdministration != nil {
			children = append(children, *r.RouteOfAdministration)
		}
	}
	if len(name) == 0 || slices.Contains(name, "administration") {
		for _, v := range r.Administration {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "maxVolumeToDeliver") {
		if r.MaxVolumeToDeliver != nil {
			children = append(children, *r.MaxVolumeToDeliver)
		}
	}
	if len(name) == 0 || slices.Contains(name, "administrationInstruction") {
		if r.AdministrationInstruction != nil {
			children = append(children, *r.AdministrationInstruction)
		}
	}
	return children
}
func (r NutritionOrderEnteralFormula) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormula to Boolean")
}
func (r NutritionOrderEnteralFormula) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormula to String")
}
func (r NutritionOrderEnteralFormula) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormula to Integer")
}
func (r NutritionOrderEnteralFormula) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormula to Decimal")
}
func (r NutritionOrderEnteralFormula) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormula to Date")
}
func (r NutritionOrderEnteralFormula) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormula to Time")
}
func (r NutritionOrderEnteralFormula) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormula to DateTime")
}
func (r NutritionOrderEnteralFormula) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormula to Quantity")
}
func (r NutritionOrderEnteralFormula) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "BaseFormulaType",
			Type: "FHIR.CodeableReference",
		}, {
			Name: "BaseFormulaProductName",
			Type: "FHIR.String",
		}, {
			Name: "DeliveryDevice",
			Type: "List<FHIR.CodeableReference>",
		}, {
			Name: "Additive",
			Type: "List<FHIR.NutritionOrderEnteralFormulaAdditive>",
		}, {
			Name: "CaloricDensity",
			Type: "FHIR.Quantity",
		}, {
			Name: "RouteOfAdministration",
			Type: "FHIR.CodeableConcept",
		}, {
			Name: "Administration",
			Type: "List<FHIR.NutritionOrderEnteralFormulaAdministration>",
		}, {
			Name: "MaxVolumeToDeliver",
			Type: "FHIR.Quantity",
		}, {
			Name: "AdministrationInstruction",
			Type: "FHIR.Markdown",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "NutritionOrderEnteralFormula",
			Namespace: "FHIR",
		},
	}
}
func (r NutritionOrderEnteralFormulaAdditive) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "type") {
		if r.Type != nil {
			children = append(children, *r.Type)
		}
	}
	if len(name) == 0 || slices.Contains(name, "productName") {
		if r.ProductName != nil {
			children = append(children, *r.ProductName)
		}
	}
	if len(name) == 0 || slices.Contains(name, "quantity") {
		if r.Quantity != nil {
			children = append(children, *r.Quantity)
		}
	}
	return children
}
func (r NutritionOrderEnteralFormulaAdditive) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdditive to Boolean")
}
func (r NutritionOrderEnteralFormulaAdditive) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdditive to String")
}
func (r NutritionOrderEnteralFormulaAdditive) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdditive to Integer")
}
func (r NutritionOrderEnteralFormulaAdditive) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdditive to Decimal")
}
func (r NutritionOrderEnteralFormulaAdditive) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdditive to Date")
}
func (r NutritionOrderEnteralFormulaAdditive) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdditive to Time")
}
func (r NutritionOrderEnteralFormulaAdditive) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdditive to DateTime")
}
func (r NutritionOrderEnteralFormulaAdditive) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdditive to Quantity")
}
func (r NutritionOrderEnteralFormulaAdditive) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Type",
			Type: "FHIR.CodeableReference",
		}, {
			Name: "ProductName",
			Type: "FHIR.String",
		}, {
			Name: "Quantity",
			Type: "FHIR.Quantity",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "NutritionOrderEnteralFormulaAdditive",
			Namespace: "FHIR",
		},
	}
}
func (r NutritionOrderEnteralFormulaAdministration) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "schedule") {
		if r.Schedule != nil {
			children = append(children, *r.Schedule)
		}
	}
	if len(name) == 0 || slices.Contains(name, "quantity") {
		if r.Quantity != nil {
			children = append(children, *r.Quantity)
		}
	}
	if len(name) == 0 || slices.Contains(name, "rate") {
		if r.Rate != nil {
			children = append(children, r.Rate)
		}
	}
	return children
}
func (r NutritionOrderEnteralFormulaAdministration) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministration to Boolean")
}
func (r NutritionOrderEnteralFormulaAdministration) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministration to String")
}
func (r NutritionOrderEnteralFormulaAdministration) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministration to Integer")
}
func (r NutritionOrderEnteralFormulaAdministration) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministration to Decimal")
}
func (r NutritionOrderEnteralFormulaAdministration) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministration to Date")
}
func (r NutritionOrderEnteralFormulaAdministration) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministration to Time")
}
func (r NutritionOrderEnteralFormulaAdministration) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministration to DateTime")
}
func (r NutritionOrderEnteralFormulaAdministration) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministration to Quantity")
}
func (r NutritionOrderEnteralFormulaAdministration) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Schedule",
			Type: "FHIR.NutritionOrderEnteralFormulaAdministrationSchedule",
		}, {
			Name: "Quantity",
			Type: "FHIR.Quantity",
		}, {
			Name: "Rate",
			Type: "FHIR.PrimitiveElement",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "NutritionOrderEnteralFormulaAdministration",
			Namespace: "FHIR",
		},
	}
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) Children(name ...string) fhirpath.Collection {
	var children fhirpath.Collection
	if len(name) == 0 || slices.Contains(name, "id") {
		if r.Id != nil {
			children = append(children, fhirpath.String(*r.Id))
		}
	}
	if len(name) == 0 || slices.Contains(name, "extension") {
		for _, v := range r.Extension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "modifierExtension") {
		for _, v := range r.ModifierExtension {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "timing") {
		for _, v := range r.Timing {
			children = append(children, v)
		}
	}
	if len(name) == 0 || slices.Contains(name, "asNeeded") {
		if r.AsNeeded != nil {
			children = append(children, *r.AsNeeded)
		}
	}
	if len(name) == 0 || slices.Contains(name, "asNeededFor") {
		if r.AsNeededFor != nil {
			children = append(children, *r.AsNeededFor)
		}
	}
	return children
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministrationSchedule to Boolean")
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministrationSchedule to String")
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministrationSchedule to Integer")
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministrationSchedule to Decimal")
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministrationSchedule to Date")
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministrationSchedule to Time")
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministrationSchedule to DateTime")
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, errors.New("can not convert NutritionOrderEnteralFormulaAdministrationSchedule to Quantity")
}
func (r NutritionOrderEnteralFormulaAdministrationSchedule) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		Element: []fhirpath.ClassInfoElement{{
			Name: "Id",
			Type: "FHIR.string",
		}, {
			Name: "Extension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "ModifierExtension",
			Type: "List<FHIR.Extension>",
		}, {
			Name: "Timing",
			Type: "List<FHIR.Timing>",
		}, {
			Name: "AsNeeded",
			Type: "FHIR.Boolean",
		}, {
			Name: "AsNeededFor",
			Type: "FHIR.CodeableConcept",
		}},
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			BaseType: fhirpath.TypeSpecifier{
				Name:      "DataType",
				Namespace: "FHIR",
			},
			Name:      "NutritionOrderEnteralFormulaAdministrationSchedule",
			Namespace: "FHIR",
		},
	}
}
