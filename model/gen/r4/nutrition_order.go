package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
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
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []model.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
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
	// The workflow status of the nutrition order/request.
	Status Code
	// Indicates the level of authority/intentionality associated with the NutrionOrder and where the request fits into the workflow chain.
	Intent Code
	// The person (patient) who needs the nutrition order for an oral diet, nutritional supplement and/or enteral or formula feeding.
	Patient Reference
	// An encounter that provides additional information about the healthcare context in which this request is made.
	Encounter *Reference
	// The date and time that this nutrition order was requested.
	DateTime DateTime
	// The practitioner that holds legal responsibility for ordering the diet, nutritional supplement, or formula feedings.
	Orderer *Reference
	// A link to a record of allergies or intolerances  which should be included in the nutrition order.
	AllergyIntolerance []Reference
	// This modifier is used to convey order-specific modifiers about the type of food that should be given. These can be derived from patient allergies, intolerances, or preferences such as Halal, Vegan or Kosher. This modifier applies to the entire nutrition order inclusive of the oral diet, nutritional supplements and enteral formula feedings.
	FoodPreferenceModifier []CodeableConcept
	// This modifier is used to convey Order-specific modifier about the type of oral food or oral fluids that should not be given. These can be derived from patient allergies, intolerances, or preferences such as No Red Meat, No Soy or No Wheat or  Gluten-Free.  While it should not be necessary to repeat allergy or intolerance information captured in the referenced AllergyIntolerance resource in the excludeFoodModifier, this element may be used to convey additional specificity related to foods that should be eliminated from the patient’s diet for any reason.  This modifier applies to the entire nutrition order inclusive of the oral diet, nutritional supplements and enteral formula feedings.
	ExcludeFoodModifier []CodeableConcept
	// Diet given orally in contrast to enteral (tube) feeding.
	OralDiet *NutritionOrderOralDiet
	// Oral nutritional products given in order to add further nutritional value to the patient's diet.
	Supplement []NutritionOrderSupplement
	// Feeding provided through the gastrointestinal tract via a tube, catheter, or stoma that delivers nutrition distal to the oral cavity.
	EnteralFormula *NutritionOrderEnteralFormula
	// Comments made about the {{title}} by the requester, performer, subject or other participants.
	Note []Annotation
}
type jsonNutritionOrder struct {
	ResourceType                          string                        `json:"resourceType"`
	Id                                    *Id                           `json:"id,omitempty"`
	IdPrimitiveElement                    *primitiveElement             `json:"_id,omitempty"`
	Meta                                  *Meta                         `json:"meta,omitempty"`
	ImplicitRules                         *Uri                          `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement         *primitiveElement             `json:"_implicitRules,omitempty"`
	Language                              *Code                         `json:"language,omitempty"`
	LanguagePrimitiveElement              *primitiveElement             `json:"_language,omitempty"`
	Text                                  *Narrative                    `json:"text,omitempty"`
	Contained                             []containedResource           `json:"contained,omitempty"`
	Extension                             []Extension                   `json:"extension,omitempty"`
	ModifierExtension                     []Extension                   `json:"modifierExtension,omitempty"`
	Identifier                            []Identifier                  `json:"identifier,omitempty"`
	InstantiatesCanonical                 []Canonical                   `json:"instantiatesCanonical,omitempty"`
	InstantiatesCanonicalPrimitiveElement []*primitiveElement           `json:"_instantiatesCanonical,omitempty"`
	InstantiatesUri                       []Uri                         `json:"instantiatesUri,omitempty"`
	InstantiatesUriPrimitiveElement       []*primitiveElement           `json:"_instantiatesUri,omitempty"`
	Instantiates                          []Uri                         `json:"instantiates,omitempty"`
	InstantiatesPrimitiveElement          []*primitiveElement           `json:"_instantiates,omitempty"`
	Status                                Code                          `json:"status,omitempty"`
	StatusPrimitiveElement                *primitiveElement             `json:"_status,omitempty"`
	Intent                                Code                          `json:"intent,omitempty"`
	IntentPrimitiveElement                *primitiveElement             `json:"_intent,omitempty"`
	Patient                               Reference                     `json:"patient,omitempty"`
	Encounter                             *Reference                    `json:"encounter,omitempty"`
	DateTime                              DateTime                      `json:"dateTime,omitempty"`
	DateTimePrimitiveElement              *primitiveElement             `json:"_dateTime,omitempty"`
	Orderer                               *Reference                    `json:"orderer,omitempty"`
	AllergyIntolerance                    []Reference                   `json:"allergyIntolerance,omitempty"`
	FoodPreferenceModifier                []CodeableConcept             `json:"foodPreferenceModifier,omitempty"`
	ExcludeFoodModifier                   []CodeableConcept             `json:"excludeFoodModifier,omitempty"`
	OralDiet                              *NutritionOrderOralDiet       `json:"oralDiet,omitempty"`
	Supplement                            []NutritionOrderSupplement    `json:"supplement,omitempty"`
	EnteralFormula                        *NutritionOrderEnteralFormula `json:"enteralFormula,omitempty"`
	Note                                  []Annotation                  `json:"note,omitempty"`
}

func (r NutritionOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r NutritionOrder) marshalJSON() jsonNutritionOrder {
	m := jsonNutritionOrder{}
	m.ResourceType = "NutritionOrder"
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
	m.InstantiatesCanonical = r.InstantiatesCanonical
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
	m.InstantiatesUri = r.InstantiatesUri
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
	m.Instantiates = r.Instantiates
	anyInstantiatesIdOrExtension := false
	for _, e := range r.Instantiates {
		if e.Id != nil || e.Extension != nil {
			anyInstantiatesIdOrExtension = true
			break
		}
	}
	if anyInstantiatesIdOrExtension {
		m.InstantiatesPrimitiveElement = make([]*primitiveElement, 0, len(r.Instantiates))
		for _, e := range r.Instantiates {
			if e.Id != nil || e.Extension != nil {
				m.InstantiatesPrimitiveElement = append(m.InstantiatesPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.InstantiatesPrimitiveElement = append(m.InstantiatesPrimitiveElement, nil)
			}
		}
	}
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Intent = r.Intent
	if r.Intent.Id != nil || r.Intent.Extension != nil {
		m.IntentPrimitiveElement = &primitiveElement{Id: r.Intent.Id, Extension: r.Intent.Extension}
	}
	m.Patient = r.Patient
	m.Encounter = r.Encounter
	m.DateTime = r.DateTime
	if r.DateTime.Id != nil || r.DateTime.Extension != nil {
		m.DateTimePrimitiveElement = &primitiveElement{Id: r.DateTime.Id, Extension: r.DateTime.Extension}
	}
	m.Orderer = r.Orderer
	m.AllergyIntolerance = r.AllergyIntolerance
	m.FoodPreferenceModifier = r.FoodPreferenceModifier
	m.ExcludeFoodModifier = r.ExcludeFoodModifier
	m.OralDiet = r.OralDiet
	m.Supplement = r.Supplement
	m.EnteralFormula = r.EnteralFormula
	m.Note = r.Note
	return m
}
func (r *NutritionOrder) UnmarshalJSON(b []byte) error {
	var m jsonNutritionOrder
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *NutritionOrder) unmarshalJSON(m jsonNutritionOrder) error {
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
	r.InstantiatesCanonical = m.InstantiatesCanonical
	for i, e := range m.InstantiatesCanonicalPrimitiveElement {
		if len(r.InstantiatesCanonical) > i {
			r.InstantiatesCanonical[i].Id = e.Id
			r.InstantiatesCanonical[i].Extension = e.Extension
		} else {
			r.InstantiatesCanonical = append(r.InstantiatesCanonical, Canonical{Id: e.Id, Extension: e.Extension})
		}
	}
	r.InstantiatesUri = m.InstantiatesUri
	for i, e := range m.InstantiatesUriPrimitiveElement {
		if len(r.InstantiatesUri) > i {
			r.InstantiatesUri[i].Id = e.Id
			r.InstantiatesUri[i].Extension = e.Extension
		} else {
			r.InstantiatesUri = append(r.InstantiatesUri, Uri{Id: e.Id, Extension: e.Extension})
		}
	}
	r.Instantiates = m.Instantiates
	for i, e := range m.InstantiatesPrimitiveElement {
		if len(r.Instantiates) > i {
			r.Instantiates[i].Id = e.Id
			r.Instantiates[i].Extension = e.Extension
		} else {
			r.Instantiates = append(r.Instantiates, Uri{Id: e.Id, Extension: e.Extension})
		}
	}
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
	r.Patient = m.Patient
	r.Encounter = m.Encounter
	r.DateTime = m.DateTime
	if m.DateTimePrimitiveElement != nil {
		r.DateTime.Id = m.DateTimePrimitiveElement.Id
		r.DateTime.Extension = m.DateTimePrimitiveElement.Extension
	}
	r.Orderer = m.Orderer
	r.AllergyIntolerance = m.AllergyIntolerance
	r.FoodPreferenceModifier = m.FoodPreferenceModifier
	r.ExcludeFoodModifier = m.ExcludeFoodModifier
	r.OralDiet = m.OralDiet
	r.Supplement = m.Supplement
	r.EnteralFormula = m.EnteralFormula
	r.Note = m.Note
	return nil
}
func (r NutritionOrder) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Diet given orally in contrast to enteral (tube) feeding.
type NutritionOrderOralDiet struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The kind of diet or dietary restriction such as fiber restricted diet or diabetic diet.
	Type []CodeableConcept
	// The time period and frequency at which the diet should be given.  The diet should be given for the combination of all schedules if more than one schedule is present.
	Schedule []Timing
	// Class that defines the quantity and type of nutrient modifications (for example carbohydrate, fiber or sodium) required for the oral diet.
	Nutrient []NutritionOrderOralDietNutrient
	// Class that describes any texture modifications required for the patient to safely consume various types of solid foods.
	Texture []NutritionOrderOralDietTexture
	// The required consistency (e.g. honey-thick, nectar-thick, thin, thickened.) of liquids or fluids served to the patient.
	FluidConsistencyType []CodeableConcept
	// Free text or additional instructions or information pertaining to the oral diet.
	Instruction *String
}
type jsonNutritionOrderOralDiet struct {
	Id                          *string                          `json:"id,omitempty"`
	Extension                   []Extension                      `json:"extension,omitempty"`
	ModifierExtension           []Extension                      `json:"modifierExtension,omitempty"`
	Type                        []CodeableConcept                `json:"type,omitempty"`
	Schedule                    []Timing                         `json:"schedule,omitempty"`
	Nutrient                    []NutritionOrderOralDietNutrient `json:"nutrient,omitempty"`
	Texture                     []NutritionOrderOralDietTexture  `json:"texture,omitempty"`
	FluidConsistencyType        []CodeableConcept                `json:"fluidConsistencyType,omitempty"`
	Instruction                 *String                          `json:"instruction,omitempty"`
	InstructionPrimitiveElement *primitiveElement                `json:"_instruction,omitempty"`
}

func (r NutritionOrderOralDiet) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r NutritionOrderOralDiet) marshalJSON() jsonNutritionOrderOralDiet {
	m := jsonNutritionOrderOralDiet{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Schedule = r.Schedule
	m.Nutrient = r.Nutrient
	m.Texture = r.Texture
	m.FluidConsistencyType = r.FluidConsistencyType
	m.Instruction = r.Instruction
	if r.Instruction != nil && (r.Instruction.Id != nil || r.Instruction.Extension != nil) {
		m.InstructionPrimitiveElement = &primitiveElement{Id: r.Instruction.Id, Extension: r.Instruction.Extension}
	}
	return m
}
func (r *NutritionOrderOralDiet) UnmarshalJSON(b []byte) error {
	var m jsonNutritionOrderOralDiet
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *NutritionOrderOralDiet) unmarshalJSON(m jsonNutritionOrderOralDiet) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Schedule = m.Schedule
	r.Nutrient = m.Nutrient
	r.Texture = m.Texture
	r.FluidConsistencyType = m.FluidConsistencyType
	r.Instruction = m.Instruction
	if m.InstructionPrimitiveElement != nil {
		r.Instruction.Id = m.InstructionPrimitiveElement.Id
		r.Instruction.Extension = m.InstructionPrimitiveElement.Extension
	}
	return nil
}
func (r NutritionOrderOralDiet) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Class that defines the quantity and type of nutrient modifications (for example carbohydrate, fiber or sodium) required for the oral diet.
type NutritionOrderOralDietNutrient struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The nutrient that is being modified such as carbohydrate or sodium.
	Modifier *CodeableConcept
	// The quantity of the specified nutrient to include in diet.
	Amount *Quantity
}
type jsonNutritionOrderOralDietNutrient struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Modifier          *CodeableConcept `json:"modifier,omitempty"`
	Amount            *Quantity        `json:"amount,omitempty"`
}

func (r NutritionOrderOralDietNutrient) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r NutritionOrderOralDietNutrient) marshalJSON() jsonNutritionOrderOralDietNutrient {
	m := jsonNutritionOrderOralDietNutrient{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Modifier = r.Modifier
	m.Amount = r.Amount
	return m
}
func (r *NutritionOrderOralDietNutrient) UnmarshalJSON(b []byte) error {
	var m jsonNutritionOrderOralDietNutrient
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *NutritionOrderOralDietNutrient) unmarshalJSON(m jsonNutritionOrderOralDietNutrient) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Modifier = m.Modifier
	r.Amount = m.Amount
	return nil
}
func (r NutritionOrderOralDietNutrient) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Class that describes any texture modifications required for the patient to safely consume various types of solid foods.
type NutritionOrderOralDietTexture struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Any texture modifications (for solid foods) that should be made, e.g. easy to chew, chopped, ground, and pureed.
	Modifier *CodeableConcept
	// The food type(s) (e.g. meats, all foods)  that the texture modification applies to.  This could be all foods types.
	FoodType *CodeableConcept
}
type jsonNutritionOrderOralDietTexture struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Modifier          *CodeableConcept `json:"modifier,omitempty"`
	FoodType          *CodeableConcept `json:"foodType,omitempty"`
}

func (r NutritionOrderOralDietTexture) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r NutritionOrderOralDietTexture) marshalJSON() jsonNutritionOrderOralDietTexture {
	m := jsonNutritionOrderOralDietTexture{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Modifier = r.Modifier
	m.FoodType = r.FoodType
	return m
}
func (r *NutritionOrderOralDietTexture) UnmarshalJSON(b []byte) error {
	var m jsonNutritionOrderOralDietTexture
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *NutritionOrderOralDietTexture) unmarshalJSON(m jsonNutritionOrderOralDietTexture) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Modifier = m.Modifier
	r.FoodType = m.FoodType
	return nil
}
func (r NutritionOrderOralDietTexture) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Oral nutritional products given in order to add further nutritional value to the patient's diet.
type NutritionOrderSupplement struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The kind of nutritional supplement product required such as a high protein or pediatric clear liquid supplement.
	Type *CodeableConcept
	// The product or brand name of the nutritional supplement such as "Acme Protein Shake".
	ProductName *String
	// The time period and frequency at which the supplement(s) should be given.  The supplement should be given for the combination of all schedules if more than one schedule is present.
	Schedule []Timing
	// The amount of the nutritional supplement to be given.
	Quantity *Quantity
	// Free text or additional instructions or information pertaining to the oral supplement.
	Instruction *String
}
type jsonNutritionOrderSupplement struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	Type                        *CodeableConcept  `json:"type,omitempty"`
	ProductName                 *String           `json:"productName,omitempty"`
	ProductNamePrimitiveElement *primitiveElement `json:"_productName,omitempty"`
	Schedule                    []Timing          `json:"schedule,omitempty"`
	Quantity                    *Quantity         `json:"quantity,omitempty"`
	Instruction                 *String           `json:"instruction,omitempty"`
	InstructionPrimitiveElement *primitiveElement `json:"_instruction,omitempty"`
}

func (r NutritionOrderSupplement) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r NutritionOrderSupplement) marshalJSON() jsonNutritionOrderSupplement {
	m := jsonNutritionOrderSupplement{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.ProductName = r.ProductName
	if r.ProductName != nil && (r.ProductName.Id != nil || r.ProductName.Extension != nil) {
		m.ProductNamePrimitiveElement = &primitiveElement{Id: r.ProductName.Id, Extension: r.ProductName.Extension}
	}
	m.Schedule = r.Schedule
	m.Quantity = r.Quantity
	m.Instruction = r.Instruction
	if r.Instruction != nil && (r.Instruction.Id != nil || r.Instruction.Extension != nil) {
		m.InstructionPrimitiveElement = &primitiveElement{Id: r.Instruction.Id, Extension: r.Instruction.Extension}
	}
	return m
}
func (r *NutritionOrderSupplement) UnmarshalJSON(b []byte) error {
	var m jsonNutritionOrderSupplement
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *NutritionOrderSupplement) unmarshalJSON(m jsonNutritionOrderSupplement) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.ProductName = m.ProductName
	if m.ProductNamePrimitiveElement != nil {
		r.ProductName.Id = m.ProductNamePrimitiveElement.Id
		r.ProductName.Extension = m.ProductNamePrimitiveElement.Extension
	}
	r.Schedule = m.Schedule
	r.Quantity = m.Quantity
	r.Instruction = m.Instruction
	if m.InstructionPrimitiveElement != nil {
		r.Instruction.Id = m.InstructionPrimitiveElement.Id
		r.Instruction.Extension = m.InstructionPrimitiveElement.Extension
	}
	return nil
}
func (r NutritionOrderSupplement) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Feeding provided through the gastrointestinal tract via a tube, catheter, or stoma that delivers nutrition distal to the oral cavity.
type NutritionOrderEnteralFormula struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of enteral or infant formula such as an adult standard formula with fiber or a soy-based infant formula.
	BaseFormulaType *CodeableConcept
	// The product or brand name of the enteral or infant formula product such as "ACME Adult Standard Formula".
	BaseFormulaProductName *String
	// Indicates the type of modular component such as protein, carbohydrate, fat or fiber to be provided in addition to or mixed with the base formula.
	AdditiveType *CodeableConcept
	// The product or brand name of the type of modular component to be added to the formula.
	AdditiveProductName *String
	// The amount of energy (calories) that the formula should provide per specified volume, typically per mL or fluid oz.  For example, an infant may require a formula that provides 24 calories per fluid ounce or an adult may require an enteral formula that provides 1.5 calorie/mL.
	CaloricDensity *Quantity
	// The route or physiological path of administration into the patient's gastrointestinal  tract for purposes of providing the formula feeding, e.g. nasogastric tube.
	RouteofAdministration *CodeableConcept
	// Formula administration instructions as structured data.  This repeating structure allows for changing the administration rate or volume over time for both bolus and continuous feeding.  An example of this would be an instruction to increase the rate of continuous feeding every 2 hours.
	Administration []NutritionOrderEnteralFormulaAdministration
	// The maximum total quantity of formula that may be administered to a subject over the period of time, e.g. 1440 mL over 24 hours.
	MaxVolumeToDeliver *Quantity
	// Free text formula administration, feeding instructions or additional instructions or information.
	AdministrationInstruction *String
}
type jsonNutritionOrderEnteralFormula struct {
	Id                                        *string                                      `json:"id,omitempty"`
	Extension                                 []Extension                                  `json:"extension,omitempty"`
	ModifierExtension                         []Extension                                  `json:"modifierExtension,omitempty"`
	BaseFormulaType                           *CodeableConcept                             `json:"baseFormulaType,omitempty"`
	BaseFormulaProductName                    *String                                      `json:"baseFormulaProductName,omitempty"`
	BaseFormulaProductNamePrimitiveElement    *primitiveElement                            `json:"_baseFormulaProductName,omitempty"`
	AdditiveType                              *CodeableConcept                             `json:"additiveType,omitempty"`
	AdditiveProductName                       *String                                      `json:"additiveProductName,omitempty"`
	AdditiveProductNamePrimitiveElement       *primitiveElement                            `json:"_additiveProductName,omitempty"`
	CaloricDensity                            *Quantity                                    `json:"caloricDensity,omitempty"`
	RouteofAdministration                     *CodeableConcept                             `json:"routeofAdministration,omitempty"`
	Administration                            []NutritionOrderEnteralFormulaAdministration `json:"administration,omitempty"`
	MaxVolumeToDeliver                        *Quantity                                    `json:"maxVolumeToDeliver,omitempty"`
	AdministrationInstruction                 *String                                      `json:"administrationInstruction,omitempty"`
	AdministrationInstructionPrimitiveElement *primitiveElement                            `json:"_administrationInstruction,omitempty"`
}

func (r NutritionOrderEnteralFormula) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r NutritionOrderEnteralFormula) marshalJSON() jsonNutritionOrderEnteralFormula {
	m := jsonNutritionOrderEnteralFormula{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.BaseFormulaType = r.BaseFormulaType
	m.BaseFormulaProductName = r.BaseFormulaProductName
	if r.BaseFormulaProductName != nil && (r.BaseFormulaProductName.Id != nil || r.BaseFormulaProductName.Extension != nil) {
		m.BaseFormulaProductNamePrimitiveElement = &primitiveElement{Id: r.BaseFormulaProductName.Id, Extension: r.BaseFormulaProductName.Extension}
	}
	m.AdditiveType = r.AdditiveType
	m.AdditiveProductName = r.AdditiveProductName
	if r.AdditiveProductName != nil && (r.AdditiveProductName.Id != nil || r.AdditiveProductName.Extension != nil) {
		m.AdditiveProductNamePrimitiveElement = &primitiveElement{Id: r.AdditiveProductName.Id, Extension: r.AdditiveProductName.Extension}
	}
	m.CaloricDensity = r.CaloricDensity
	m.RouteofAdministration = r.RouteofAdministration
	m.Administration = r.Administration
	m.MaxVolumeToDeliver = r.MaxVolumeToDeliver
	m.AdministrationInstruction = r.AdministrationInstruction
	if r.AdministrationInstruction != nil && (r.AdministrationInstruction.Id != nil || r.AdministrationInstruction.Extension != nil) {
		m.AdministrationInstructionPrimitiveElement = &primitiveElement{Id: r.AdministrationInstruction.Id, Extension: r.AdministrationInstruction.Extension}
	}
	return m
}
func (r *NutritionOrderEnteralFormula) UnmarshalJSON(b []byte) error {
	var m jsonNutritionOrderEnteralFormula
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *NutritionOrderEnteralFormula) unmarshalJSON(m jsonNutritionOrderEnteralFormula) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.BaseFormulaType = m.BaseFormulaType
	r.BaseFormulaProductName = m.BaseFormulaProductName
	if m.BaseFormulaProductNamePrimitiveElement != nil {
		r.BaseFormulaProductName.Id = m.BaseFormulaProductNamePrimitiveElement.Id
		r.BaseFormulaProductName.Extension = m.BaseFormulaProductNamePrimitiveElement.Extension
	}
	r.AdditiveType = m.AdditiveType
	r.AdditiveProductName = m.AdditiveProductName
	if m.AdditiveProductNamePrimitiveElement != nil {
		r.AdditiveProductName.Id = m.AdditiveProductNamePrimitiveElement.Id
		r.AdditiveProductName.Extension = m.AdditiveProductNamePrimitiveElement.Extension
	}
	r.CaloricDensity = m.CaloricDensity
	r.RouteofAdministration = m.RouteofAdministration
	r.Administration = m.Administration
	r.MaxVolumeToDeliver = m.MaxVolumeToDeliver
	r.AdministrationInstruction = m.AdministrationInstruction
	if m.AdministrationInstructionPrimitiveElement != nil {
		r.AdministrationInstruction.Id = m.AdministrationInstructionPrimitiveElement.Id
		r.AdministrationInstruction.Extension = m.AdministrationInstructionPrimitiveElement.Extension
	}
	return nil
}
func (r NutritionOrderEnteralFormula) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Formula administration instructions as structured data.  This repeating structure allows for changing the administration rate or volume over time for both bolus and continuous feeding.  An example of this would be an instruction to increase the rate of continuous feeding every 2 hours.
type NutritionOrderEnteralFormulaAdministration struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The time period and frequency at which the enteral formula should be delivered to the patient.
	Schedule *Timing
	// The volume of formula to provide to the patient per the specified administration schedule.
	Quantity *Quantity
	// The rate of administration of formula via a feeding pump, e.g. 60 mL per hour, according to the specified schedule.
	Rate isNutritionOrderEnteralFormulaAdministrationRate
}
type isNutritionOrderEnteralFormulaAdministrationRate interface {
	isNutritionOrderEnteralFormulaAdministrationRate()
}

func (r Quantity) isNutritionOrderEnteralFormulaAdministrationRate() {}
func (r Ratio) isNutritionOrderEnteralFormulaAdministrationRate()    {}

type jsonNutritionOrderEnteralFormulaAdministration struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Schedule          *Timing     `json:"schedule,omitempty"`
	Quantity          *Quantity   `json:"quantity,omitempty"`
	RateQuantity      *Quantity   `json:"rateQuantity,omitempty"`
	RateRatio         *Ratio      `json:"rateRatio,omitempty"`
}

func (r NutritionOrderEnteralFormulaAdministration) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r NutritionOrderEnteralFormulaAdministration) marshalJSON() jsonNutritionOrderEnteralFormulaAdministration {
	m := jsonNutritionOrderEnteralFormulaAdministration{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Schedule = r.Schedule
	m.Quantity = r.Quantity
	switch v := r.Rate.(type) {
	case Quantity:
		m.RateQuantity = &v
	case *Quantity:
		m.RateQuantity = v
	case Ratio:
		m.RateRatio = &v
	case *Ratio:
		m.RateRatio = v
	}
	return m
}
func (r *NutritionOrderEnteralFormulaAdministration) UnmarshalJSON(b []byte) error {
	var m jsonNutritionOrderEnteralFormulaAdministration
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *NutritionOrderEnteralFormulaAdministration) unmarshalJSON(m jsonNutritionOrderEnteralFormulaAdministration) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Schedule = m.Schedule
	r.Quantity = m.Quantity
	if m.RateQuantity != nil {
		if r.Rate != nil {
			return fmt.Errorf("multiple values for field \"Rate\"")
		}
		v := m.RateQuantity
		r.Rate = v
	}
	if m.RateRatio != nil {
		if r.Rate != nil {
			return fmt.Errorf("multiple values for field \"Rate\"")
		}
		v := m.RateRatio
		r.Rate = v
	}
	return nil
}
func (r NutritionOrderEnteralFormulaAdministration) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
