package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// Todo.
type SubstancePolymer struct {
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
	// Todo.
	Class *CodeableConcept
	// Todo.
	Geometry *CodeableConcept
	// Todo.
	CopolymerConnectivity []CodeableConcept
	// Todo.
	Modification []String
	// Todo.
	MonomerSet []SubstancePolymerMonomerSet
	// Todo.
	Repeat []SubstancePolymerRepeat
}

func (r SubstancePolymer) ResourceType() string {
	return "SubstancePolymer"
}

type jsonSubstancePolymer struct {
	ResourceType                  string                       `json:"resourceType"`
	Id                            *Id                          `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement            `json:"_id,omitempty"`
	Meta                          *Meta                        `json:"meta,omitempty"`
	ImplicitRules                 *Uri                         `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement            `json:"_implicitRules,omitempty"`
	Language                      *Code                        `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement            `json:"_language,omitempty"`
	Text                          *Narrative                   `json:"text,omitempty"`
	Contained                     []containedResource          `json:"contained,omitempty"`
	Extension                     []Extension                  `json:"extension,omitempty"`
	ModifierExtension             []Extension                  `json:"modifierExtension,omitempty"`
	Class                         *CodeableConcept             `json:"class,omitempty"`
	Geometry                      *CodeableConcept             `json:"geometry,omitempty"`
	CopolymerConnectivity         []CodeableConcept            `json:"copolymerConnectivity,omitempty"`
	Modification                  []String                     `json:"modification,omitempty"`
	ModificationPrimitiveElement  []*primitiveElement          `json:"_modification,omitempty"`
	MonomerSet                    []SubstancePolymerMonomerSet `json:"monomerSet,omitempty"`
	Repeat                        []SubstancePolymerRepeat     `json:"repeat,omitempty"`
}

func (r SubstancePolymer) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstancePolymer) marshalJSON() jsonSubstancePolymer {
	m := jsonSubstancePolymer{}
	m.ResourceType = "SubstancePolymer"
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
	m.Class = r.Class
	m.Geometry = r.Geometry
	m.CopolymerConnectivity = r.CopolymerConnectivity
	m.Modification = r.Modification
	anyModificationIdOrExtension := false
	for _, e := range r.Modification {
		if e.Id != nil || e.Extension != nil {
			anyModificationIdOrExtension = true
			break
		}
	}
	if anyModificationIdOrExtension {
		m.ModificationPrimitiveElement = make([]*primitiveElement, 0, len(r.Modification))
		for _, e := range r.Modification {
			if e.Id != nil || e.Extension != nil {
				m.ModificationPrimitiveElement = append(m.ModificationPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.ModificationPrimitiveElement = append(m.ModificationPrimitiveElement, nil)
			}
		}
	}
	m.MonomerSet = r.MonomerSet
	m.Repeat = r.Repeat
	return m
}
func (r *SubstancePolymer) UnmarshalJSON(b []byte) error {
	var m jsonSubstancePolymer
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstancePolymer) unmarshalJSON(m jsonSubstancePolymer) error {
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
	r.Class = m.Class
	r.Geometry = m.Geometry
	r.CopolymerConnectivity = m.CopolymerConnectivity
	r.Modification = m.Modification
	for i, e := range m.ModificationPrimitiveElement {
		if len(r.Modification) > i {
			r.Modification[i].Id = e.Id
			r.Modification[i].Extension = e.Extension
		} else {
			r.Modification = append(r.Modification, String{Id: e.Id, Extension: e.Extension})
		}
	}
	r.MonomerSet = m.MonomerSet
	r.Repeat = m.Repeat
	return nil
}
func (r SubstancePolymer) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Todo.
type SubstancePolymerMonomerSet struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Todo.
	RatioType *CodeableConcept
	// Todo.
	StartingMaterial []SubstancePolymerMonomerSetStartingMaterial
}
type jsonSubstancePolymerMonomerSet struct {
	Id                *string                                      `json:"id,omitempty"`
	Extension         []Extension                                  `json:"extension,omitempty"`
	ModifierExtension []Extension                                  `json:"modifierExtension,omitempty"`
	RatioType         *CodeableConcept                             `json:"ratioType,omitempty"`
	StartingMaterial  []SubstancePolymerMonomerSetStartingMaterial `json:"startingMaterial,omitempty"`
}

func (r SubstancePolymerMonomerSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstancePolymerMonomerSet) marshalJSON() jsonSubstancePolymerMonomerSet {
	m := jsonSubstancePolymerMonomerSet{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.RatioType = r.RatioType
	m.StartingMaterial = r.StartingMaterial
	return m
}
func (r *SubstancePolymerMonomerSet) UnmarshalJSON(b []byte) error {
	var m jsonSubstancePolymerMonomerSet
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstancePolymerMonomerSet) unmarshalJSON(m jsonSubstancePolymerMonomerSet) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.RatioType = m.RatioType
	r.StartingMaterial = m.StartingMaterial
	return nil
}
func (r SubstancePolymerMonomerSet) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Todo.
type SubstancePolymerMonomerSetStartingMaterial struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Todo.
	Material *CodeableConcept
	// Todo.
	Type *CodeableConcept
	// Todo.
	IsDefining *Boolean
	// Todo.
	Amount *SubstanceAmount
}
type jsonSubstancePolymerMonomerSetStartingMaterial struct {
	Id                         *string           `json:"id,omitempty"`
	Extension                  []Extension       `json:"extension,omitempty"`
	ModifierExtension          []Extension       `json:"modifierExtension,omitempty"`
	Material                   *CodeableConcept  `json:"material,omitempty"`
	Type                       *CodeableConcept  `json:"type,omitempty"`
	IsDefining                 *Boolean          `json:"isDefining,omitempty"`
	IsDefiningPrimitiveElement *primitiveElement `json:"_isDefining,omitempty"`
	Amount                     *SubstanceAmount  `json:"amount,omitempty"`
}

func (r SubstancePolymerMonomerSetStartingMaterial) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstancePolymerMonomerSetStartingMaterial) marshalJSON() jsonSubstancePolymerMonomerSetStartingMaterial {
	m := jsonSubstancePolymerMonomerSetStartingMaterial{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Material = r.Material
	m.Type = r.Type
	m.IsDefining = r.IsDefining
	if r.IsDefining != nil && (r.IsDefining.Id != nil || r.IsDefining.Extension != nil) {
		m.IsDefiningPrimitiveElement = &primitiveElement{Id: r.IsDefining.Id, Extension: r.IsDefining.Extension}
	}
	m.Amount = r.Amount
	return m
}
func (r *SubstancePolymerMonomerSetStartingMaterial) UnmarshalJSON(b []byte) error {
	var m jsonSubstancePolymerMonomerSetStartingMaterial
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstancePolymerMonomerSetStartingMaterial) unmarshalJSON(m jsonSubstancePolymerMonomerSetStartingMaterial) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Material = m.Material
	r.Type = m.Type
	r.IsDefining = m.IsDefining
	if m.IsDefiningPrimitiveElement != nil {
		r.IsDefining.Id = m.IsDefiningPrimitiveElement.Id
		r.IsDefining.Extension = m.IsDefiningPrimitiveElement.Extension
	}
	r.Amount = m.Amount
	return nil
}
func (r SubstancePolymerMonomerSetStartingMaterial) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Todo.
type SubstancePolymerRepeat struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Todo.
	NumberOfUnits *Integer
	// Todo.
	AverageMolecularFormula *String
	// Todo.
	RepeatUnitAmountType *CodeableConcept
	// Todo.
	RepeatUnit []SubstancePolymerRepeatRepeatUnit
}
type jsonSubstancePolymerRepeat struct {
	Id                                      *string                            `json:"id,omitempty"`
	Extension                               []Extension                        `json:"extension,omitempty"`
	ModifierExtension                       []Extension                        `json:"modifierExtension,omitempty"`
	NumberOfUnits                           *Integer                           `json:"numberOfUnits,omitempty"`
	NumberOfUnitsPrimitiveElement           *primitiveElement                  `json:"_numberOfUnits,omitempty"`
	AverageMolecularFormula                 *String                            `json:"averageMolecularFormula,omitempty"`
	AverageMolecularFormulaPrimitiveElement *primitiveElement                  `json:"_averageMolecularFormula,omitempty"`
	RepeatUnitAmountType                    *CodeableConcept                   `json:"repeatUnitAmountType,omitempty"`
	RepeatUnit                              []SubstancePolymerRepeatRepeatUnit `json:"repeatUnit,omitempty"`
}

func (r SubstancePolymerRepeat) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstancePolymerRepeat) marshalJSON() jsonSubstancePolymerRepeat {
	m := jsonSubstancePolymerRepeat{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.NumberOfUnits = r.NumberOfUnits
	if r.NumberOfUnits != nil && (r.NumberOfUnits.Id != nil || r.NumberOfUnits.Extension != nil) {
		m.NumberOfUnitsPrimitiveElement = &primitiveElement{Id: r.NumberOfUnits.Id, Extension: r.NumberOfUnits.Extension}
	}
	m.AverageMolecularFormula = r.AverageMolecularFormula
	if r.AverageMolecularFormula != nil && (r.AverageMolecularFormula.Id != nil || r.AverageMolecularFormula.Extension != nil) {
		m.AverageMolecularFormulaPrimitiveElement = &primitiveElement{Id: r.AverageMolecularFormula.Id, Extension: r.AverageMolecularFormula.Extension}
	}
	m.RepeatUnitAmountType = r.RepeatUnitAmountType
	m.RepeatUnit = r.RepeatUnit
	return m
}
func (r *SubstancePolymerRepeat) UnmarshalJSON(b []byte) error {
	var m jsonSubstancePolymerRepeat
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstancePolymerRepeat) unmarshalJSON(m jsonSubstancePolymerRepeat) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.NumberOfUnits = m.NumberOfUnits
	if m.NumberOfUnitsPrimitiveElement != nil {
		r.NumberOfUnits.Id = m.NumberOfUnitsPrimitiveElement.Id
		r.NumberOfUnits.Extension = m.NumberOfUnitsPrimitiveElement.Extension
	}
	r.AverageMolecularFormula = m.AverageMolecularFormula
	if m.AverageMolecularFormulaPrimitiveElement != nil {
		r.AverageMolecularFormula.Id = m.AverageMolecularFormulaPrimitiveElement.Id
		r.AverageMolecularFormula.Extension = m.AverageMolecularFormulaPrimitiveElement.Extension
	}
	r.RepeatUnitAmountType = m.RepeatUnitAmountType
	r.RepeatUnit = m.RepeatUnit
	return nil
}
func (r SubstancePolymerRepeat) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Todo.
type SubstancePolymerRepeatRepeatUnit struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Todo.
	OrientationOfPolymerisation *CodeableConcept
	// Todo.
	RepeatUnit *String
	// Todo.
	Amount *SubstanceAmount
	// Todo.
	DegreeOfPolymerisation []SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation
	// Todo.
	StructuralRepresentation []SubstancePolymerRepeatRepeatUnitStructuralRepresentation
}
type jsonSubstancePolymerRepeatRepeatUnit struct {
	Id                          *string                                                    `json:"id,omitempty"`
	Extension                   []Extension                                                `json:"extension,omitempty"`
	ModifierExtension           []Extension                                                `json:"modifierExtension,omitempty"`
	OrientationOfPolymerisation *CodeableConcept                                           `json:"orientationOfPolymerisation,omitempty"`
	RepeatUnit                  *String                                                    `json:"repeatUnit,omitempty"`
	RepeatUnitPrimitiveElement  *primitiveElement                                          `json:"_repeatUnit,omitempty"`
	Amount                      *SubstanceAmount                                           `json:"amount,omitempty"`
	DegreeOfPolymerisation      []SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation   `json:"degreeOfPolymerisation,omitempty"`
	StructuralRepresentation    []SubstancePolymerRepeatRepeatUnitStructuralRepresentation `json:"structuralRepresentation,omitempty"`
}

func (r SubstancePolymerRepeatRepeatUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstancePolymerRepeatRepeatUnit) marshalJSON() jsonSubstancePolymerRepeatRepeatUnit {
	m := jsonSubstancePolymerRepeatRepeatUnit{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.OrientationOfPolymerisation = r.OrientationOfPolymerisation
	m.RepeatUnit = r.RepeatUnit
	if r.RepeatUnit != nil && (r.RepeatUnit.Id != nil || r.RepeatUnit.Extension != nil) {
		m.RepeatUnitPrimitiveElement = &primitiveElement{Id: r.RepeatUnit.Id, Extension: r.RepeatUnit.Extension}
	}
	m.Amount = r.Amount
	m.DegreeOfPolymerisation = r.DegreeOfPolymerisation
	m.StructuralRepresentation = r.StructuralRepresentation
	return m
}
func (r *SubstancePolymerRepeatRepeatUnit) UnmarshalJSON(b []byte) error {
	var m jsonSubstancePolymerRepeatRepeatUnit
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstancePolymerRepeatRepeatUnit) unmarshalJSON(m jsonSubstancePolymerRepeatRepeatUnit) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.OrientationOfPolymerisation = m.OrientationOfPolymerisation
	r.RepeatUnit = m.RepeatUnit
	if m.RepeatUnitPrimitiveElement != nil {
		r.RepeatUnit.Id = m.RepeatUnitPrimitiveElement.Id
		r.RepeatUnit.Extension = m.RepeatUnitPrimitiveElement.Extension
	}
	r.Amount = m.Amount
	r.DegreeOfPolymerisation = m.DegreeOfPolymerisation
	r.StructuralRepresentation = m.StructuralRepresentation
	return nil
}
func (r SubstancePolymerRepeatRepeatUnit) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Todo.
type SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Todo.
	Degree *CodeableConcept
	// Todo.
	Amount *SubstanceAmount
}
type jsonSubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Degree            *CodeableConcept `json:"degree,omitempty"`
	Amount            *SubstanceAmount `json:"amount,omitempty"`
}

func (r SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation) marshalJSON() jsonSubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation {
	m := jsonSubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Degree = r.Degree
	m.Amount = r.Amount
	return m
}
func (r *SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation) UnmarshalJSON(b []byte) error {
	var m jsonSubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation) unmarshalJSON(m jsonSubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Degree = m.Degree
	r.Amount = m.Amount
	return nil
}
func (r SubstancePolymerRepeatRepeatUnitDegreeOfPolymerisation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Todo.
type SubstancePolymerRepeatRepeatUnitStructuralRepresentation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Todo.
	Type *CodeableConcept
	// Todo.
	Representation *String
	// Todo.
	Attachment *Attachment
}
type jsonSubstancePolymerRepeatRepeatUnitStructuralRepresentation struct {
	Id                             *string           `json:"id,omitempty"`
	Extension                      []Extension       `json:"extension,omitempty"`
	ModifierExtension              []Extension       `json:"modifierExtension,omitempty"`
	Type                           *CodeableConcept  `json:"type,omitempty"`
	Representation                 *String           `json:"representation,omitempty"`
	RepresentationPrimitiveElement *primitiveElement `json:"_representation,omitempty"`
	Attachment                     *Attachment       `json:"attachment,omitempty"`
}

func (r SubstancePolymerRepeatRepeatUnitStructuralRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SubstancePolymerRepeatRepeatUnitStructuralRepresentation) marshalJSON() jsonSubstancePolymerRepeatRepeatUnitStructuralRepresentation {
	m := jsonSubstancePolymerRepeatRepeatUnitStructuralRepresentation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	m.Representation = r.Representation
	if r.Representation != nil && (r.Representation.Id != nil || r.Representation.Extension != nil) {
		m.RepresentationPrimitiveElement = &primitiveElement{Id: r.Representation.Id, Extension: r.Representation.Extension}
	}
	m.Attachment = r.Attachment
	return m
}
func (r *SubstancePolymerRepeatRepeatUnitStructuralRepresentation) UnmarshalJSON(b []byte) error {
	var m jsonSubstancePolymerRepeatRepeatUnitStructuralRepresentation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SubstancePolymerRepeatRepeatUnitStructuralRepresentation) unmarshalJSON(m jsonSubstancePolymerRepeatRepeatUnitStructuralRepresentation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	r.Representation = m.Representation
	if m.RepresentationPrimitiveElement != nil {
		r.Representation.Id = m.RepresentationPrimitiveElement.Id
		r.Representation.Extension = m.RepresentationPrimitiveElement.Extension
	}
	r.Attachment = m.Attachment
	return nil
}
func (r SubstancePolymerRepeatRepeatUnitStructuralRepresentation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
