package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A kind of specimen with associated set of requirements.
type SpecimenDefinition struct {
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
	// A business identifier associated with the kind of specimen.
	Identifier *Identifier
	// The kind of material to be collected.
	TypeCollected *CodeableConcept
	// Preparation of the patient for specimen collection.
	PatientPreparation []CodeableConcept
	// Time aspect of specimen collection (duration or offset).
	TimeAspect *String
	// The action to be performed for collecting the specimen.
	Collection []CodeableConcept
	// Specimen conditioned in a container as expected by the testing laboratory.
	TypeTested []SpecimenDefinitionTypeTested
}

func (r SpecimenDefinition) ResourceType() string {
	return "SpecimenDefinition"
}
func (r SpecimenDefinition) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonSpecimenDefinition struct {
	ResourceType                  string                         `json:"resourceType"`
	Id                            *Id                            `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement              `json:"_id,omitempty"`
	Meta                          *Meta                          `json:"meta,omitempty"`
	ImplicitRules                 *Uri                           `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement              `json:"_implicitRules,omitempty"`
	Language                      *Code                          `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement              `json:"_language,omitempty"`
	Text                          *Narrative                     `json:"text,omitempty"`
	Contained                     []containedResource            `json:"contained,omitempty"`
	Extension                     []Extension                    `json:"extension,omitempty"`
	ModifierExtension             []Extension                    `json:"modifierExtension,omitempty"`
	Identifier                    *Identifier                    `json:"identifier,omitempty"`
	TypeCollected                 *CodeableConcept               `json:"typeCollected,omitempty"`
	PatientPreparation            []CodeableConcept              `json:"patientPreparation,omitempty"`
	TimeAspect                    *String                        `json:"timeAspect,omitempty"`
	TimeAspectPrimitiveElement    *primitiveElement              `json:"_timeAspect,omitempty"`
	Collection                    []CodeableConcept              `json:"collection,omitempty"`
	TypeTested                    []SpecimenDefinitionTypeTested `json:"typeTested,omitempty"`
}

func (r SpecimenDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SpecimenDefinition) marshalJSON() jsonSpecimenDefinition {
	m := jsonSpecimenDefinition{}
	m.ResourceType = "SpecimenDefinition"
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
	m.TypeCollected = r.TypeCollected
	m.PatientPreparation = r.PatientPreparation
	m.TimeAspect = r.TimeAspect
	if r.TimeAspect != nil && (r.TimeAspect.Id != nil || r.TimeAspect.Extension != nil) {
		m.TimeAspectPrimitiveElement = &primitiveElement{Id: r.TimeAspect.Id, Extension: r.TimeAspect.Extension}
	}
	m.Collection = r.Collection
	m.TypeTested = r.TypeTested
	return m
}
func (r *SpecimenDefinition) UnmarshalJSON(b []byte) error {
	var m jsonSpecimenDefinition
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SpecimenDefinition) unmarshalJSON(m jsonSpecimenDefinition) error {
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
	r.TypeCollected = m.TypeCollected
	r.PatientPreparation = m.PatientPreparation
	r.TimeAspect = m.TimeAspect
	if m.TimeAspectPrimitiveElement != nil {
		r.TimeAspect.Id = m.TimeAspectPrimitiveElement.Id
		r.TimeAspect.Extension = m.TimeAspectPrimitiveElement.Extension
	}
	r.Collection = m.Collection
	r.TypeTested = m.TypeTested
	return nil
}
func (r SpecimenDefinition) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Specimen conditioned in a container as expected by the testing laboratory.
type SpecimenDefinitionTypeTested struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Primary of secondary specimen.
	IsDerived *Boolean
	// The kind of specimen conditioned for testing expected by lab.
	Type *CodeableConcept
	// The preference for this type of conditioned specimen.
	Preference Code
	// The specimen's container.
	Container *SpecimenDefinitionTypeTestedContainer
	// Requirements for delivery and special handling of this kind of conditioned specimen.
	Requirement *String
	// The usual time that a specimen of this kind is retained after the ordered tests are completed, for the purpose of additional testing.
	RetentionTime *Duration
	// Criterion for rejection of the specimen in its container by the laboratory.
	RejectionCriterion []CodeableConcept
	// Set of instructions for preservation/transport of the specimen at a defined temperature interval, prior the testing process.
	Handling []SpecimenDefinitionTypeTestedHandling
}
type jsonSpecimenDefinitionTypeTested struct {
	Id                          *string                                `json:"id,omitempty"`
	Extension                   []Extension                            `json:"extension,omitempty"`
	ModifierExtension           []Extension                            `json:"modifierExtension,omitempty"`
	IsDerived                   *Boolean                               `json:"isDerived,omitempty"`
	IsDerivedPrimitiveElement   *primitiveElement                      `json:"_isDerived,omitempty"`
	Type                        *CodeableConcept                       `json:"type,omitempty"`
	Preference                  Code                                   `json:"preference,omitempty"`
	PreferencePrimitiveElement  *primitiveElement                      `json:"_preference,omitempty"`
	Container                   *SpecimenDefinitionTypeTestedContainer `json:"container,omitempty"`
	Requirement                 *String                                `json:"requirement,omitempty"`
	RequirementPrimitiveElement *primitiveElement                      `json:"_requirement,omitempty"`
	RetentionTime               *Duration                              `json:"retentionTime,omitempty"`
	RejectionCriterion          []CodeableConcept                      `json:"rejectionCriterion,omitempty"`
	Handling                    []SpecimenDefinitionTypeTestedHandling `json:"handling,omitempty"`
}

func (r SpecimenDefinitionTypeTested) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SpecimenDefinitionTypeTested) marshalJSON() jsonSpecimenDefinitionTypeTested {
	m := jsonSpecimenDefinitionTypeTested{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.IsDerived = r.IsDerived
	if r.IsDerived != nil && (r.IsDerived.Id != nil || r.IsDerived.Extension != nil) {
		m.IsDerivedPrimitiveElement = &primitiveElement{Id: r.IsDerived.Id, Extension: r.IsDerived.Extension}
	}
	m.Type = r.Type
	m.Preference = r.Preference
	if r.Preference.Id != nil || r.Preference.Extension != nil {
		m.PreferencePrimitiveElement = &primitiveElement{Id: r.Preference.Id, Extension: r.Preference.Extension}
	}
	m.Container = r.Container
	m.Requirement = r.Requirement
	if r.Requirement != nil && (r.Requirement.Id != nil || r.Requirement.Extension != nil) {
		m.RequirementPrimitiveElement = &primitiveElement{Id: r.Requirement.Id, Extension: r.Requirement.Extension}
	}
	m.RetentionTime = r.RetentionTime
	m.RejectionCriterion = r.RejectionCriterion
	m.Handling = r.Handling
	return m
}
func (r *SpecimenDefinitionTypeTested) UnmarshalJSON(b []byte) error {
	var m jsonSpecimenDefinitionTypeTested
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SpecimenDefinitionTypeTested) unmarshalJSON(m jsonSpecimenDefinitionTypeTested) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.IsDerived = m.IsDerived
	if m.IsDerivedPrimitiveElement != nil {
		r.IsDerived.Id = m.IsDerivedPrimitiveElement.Id
		r.IsDerived.Extension = m.IsDerivedPrimitiveElement.Extension
	}
	r.Type = m.Type
	r.Preference = m.Preference
	if m.PreferencePrimitiveElement != nil {
		r.Preference.Id = m.PreferencePrimitiveElement.Id
		r.Preference.Extension = m.PreferencePrimitiveElement.Extension
	}
	r.Container = m.Container
	r.Requirement = m.Requirement
	if m.RequirementPrimitiveElement != nil {
		r.Requirement.Id = m.RequirementPrimitiveElement.Id
		r.Requirement.Extension = m.RequirementPrimitiveElement.Extension
	}
	r.RetentionTime = m.RetentionTime
	r.RejectionCriterion = m.RejectionCriterion
	r.Handling = m.Handling
	return nil
}
func (r SpecimenDefinitionTypeTested) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The specimen's container.
type SpecimenDefinitionTypeTestedContainer struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The type of material of the container.
	Material *CodeableConcept
	// The type of container used to contain this kind of specimen.
	Type *CodeableConcept
	// Color of container cap.
	Cap *CodeableConcept
	// The textual description of the kind of container.
	Description *String
	// The capacity (volume or other measure) of this kind of container.
	Capacity *Quantity
	// The minimum volume to be conditioned in the container.
	MinimumVolume isSpecimenDefinitionTypeTestedContainerMinimumVolume
	// Substance introduced in the kind of container to preserve, maintain or enhance the specimen. Examples: Formalin, Citrate, EDTA.
	Additive []SpecimenDefinitionTypeTestedContainerAdditive
	// Special processing that should be applied to the container for this kind of specimen.
	Preparation *String
}
type isSpecimenDefinitionTypeTestedContainerMinimumVolume interface {
	isSpecimenDefinitionTypeTestedContainerMinimumVolume()
}

func (r Quantity) isSpecimenDefinitionTypeTestedContainerMinimumVolume() {}
func (r String) isSpecimenDefinitionTypeTestedContainerMinimumVolume()   {}

type jsonSpecimenDefinitionTypeTestedContainer struct {
	Id                                  *string                                         `json:"id,omitempty"`
	Extension                           []Extension                                     `json:"extension,omitempty"`
	ModifierExtension                   []Extension                                     `json:"modifierExtension,omitempty"`
	Material                            *CodeableConcept                                `json:"material,omitempty"`
	Type                                *CodeableConcept                                `json:"type,omitempty"`
	Cap                                 *CodeableConcept                                `json:"cap,omitempty"`
	Description                         *String                                         `json:"description,omitempty"`
	DescriptionPrimitiveElement         *primitiveElement                               `json:"_description,omitempty"`
	Capacity                            *Quantity                                       `json:"capacity,omitempty"`
	MinimumVolumeQuantity               *Quantity                                       `json:"minimumVolumeQuantity,omitempty"`
	MinimumVolumeString                 *String                                         `json:"minimumVolumeString,omitempty"`
	MinimumVolumeStringPrimitiveElement *primitiveElement                               `json:"_minimumVolumeString,omitempty"`
	Additive                            []SpecimenDefinitionTypeTestedContainerAdditive `json:"additive,omitempty"`
	Preparation                         *String                                         `json:"preparation,omitempty"`
	PreparationPrimitiveElement         *primitiveElement                               `json:"_preparation,omitempty"`
}

func (r SpecimenDefinitionTypeTestedContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SpecimenDefinitionTypeTestedContainer) marshalJSON() jsonSpecimenDefinitionTypeTestedContainer {
	m := jsonSpecimenDefinitionTypeTestedContainer{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Material = r.Material
	m.Type = r.Type
	m.Cap = r.Cap
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Capacity = r.Capacity
	switch v := r.MinimumVolume.(type) {
	case Quantity:
		m.MinimumVolumeQuantity = &v
	case *Quantity:
		m.MinimumVolumeQuantity = v
	case String:
		m.MinimumVolumeString = &v
		if v.Id != nil || v.Extension != nil {
			m.MinimumVolumeStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *String:
		m.MinimumVolumeString = v
		if v.Id != nil || v.Extension != nil {
			m.MinimumVolumeStringPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.Additive = r.Additive
	m.Preparation = r.Preparation
	if r.Preparation != nil && (r.Preparation.Id != nil || r.Preparation.Extension != nil) {
		m.PreparationPrimitiveElement = &primitiveElement{Id: r.Preparation.Id, Extension: r.Preparation.Extension}
	}
	return m
}
func (r *SpecimenDefinitionTypeTestedContainer) UnmarshalJSON(b []byte) error {
	var m jsonSpecimenDefinitionTypeTestedContainer
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SpecimenDefinitionTypeTestedContainer) unmarshalJSON(m jsonSpecimenDefinitionTypeTestedContainer) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Material = m.Material
	r.Type = m.Type
	r.Cap = m.Cap
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Capacity = m.Capacity
	if m.MinimumVolumeQuantity != nil {
		if r.MinimumVolume != nil {
			return fmt.Errorf("multiple values for field \"MinimumVolume\"")
		}
		v := m.MinimumVolumeQuantity
		r.MinimumVolume = v
	}
	if m.MinimumVolumeString != nil || m.MinimumVolumeStringPrimitiveElement != nil {
		if r.MinimumVolume != nil {
			return fmt.Errorf("multiple values for field \"MinimumVolume\"")
		}
		v := m.MinimumVolumeString
		if m.MinimumVolumeStringPrimitiveElement != nil {
			if v == nil {
				v = &String{}
			}
			v.Id = m.MinimumVolumeStringPrimitiveElement.Id
			v.Extension = m.MinimumVolumeStringPrimitiveElement.Extension
		}
		r.MinimumVolume = v
	}
	r.Additive = m.Additive
	r.Preparation = m.Preparation
	if m.PreparationPrimitiveElement != nil {
		r.Preparation.Id = m.PreparationPrimitiveElement.Id
		r.Preparation.Extension = m.PreparationPrimitiveElement.Extension
	}
	return nil
}
func (r SpecimenDefinitionTypeTestedContainer) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Substance introduced in the kind of container to preserve, maintain or enhance the specimen. Examples: Formalin, Citrate, EDTA.
type SpecimenDefinitionTypeTestedContainerAdditive struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Substance introduced in the kind of container to preserve, maintain or enhance the specimen. Examples: Formalin, Citrate, EDTA.
	Additive isSpecimenDefinitionTypeTestedContainerAdditiveAdditive
}
type isSpecimenDefinitionTypeTestedContainerAdditiveAdditive interface {
	isSpecimenDefinitionTypeTestedContainerAdditiveAdditive()
}

func (r CodeableConcept) isSpecimenDefinitionTypeTestedContainerAdditiveAdditive() {}
func (r Reference) isSpecimenDefinitionTypeTestedContainerAdditiveAdditive()       {}

type jsonSpecimenDefinitionTypeTestedContainerAdditive struct {
	Id                      *string          `json:"id,omitempty"`
	Extension               []Extension      `json:"extension,omitempty"`
	ModifierExtension       []Extension      `json:"modifierExtension,omitempty"`
	AdditiveCodeableConcept *CodeableConcept `json:"additiveCodeableConcept,omitempty"`
	AdditiveReference       *Reference       `json:"additiveReference,omitempty"`
}

func (r SpecimenDefinitionTypeTestedContainerAdditive) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SpecimenDefinitionTypeTestedContainerAdditive) marshalJSON() jsonSpecimenDefinitionTypeTestedContainerAdditive {
	m := jsonSpecimenDefinitionTypeTestedContainerAdditive{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	switch v := r.Additive.(type) {
	case CodeableConcept:
		m.AdditiveCodeableConcept = &v
	case *CodeableConcept:
		m.AdditiveCodeableConcept = v
	case Reference:
		m.AdditiveReference = &v
	case *Reference:
		m.AdditiveReference = v
	}
	return m
}
func (r *SpecimenDefinitionTypeTestedContainerAdditive) UnmarshalJSON(b []byte) error {
	var m jsonSpecimenDefinitionTypeTestedContainerAdditive
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SpecimenDefinitionTypeTestedContainerAdditive) unmarshalJSON(m jsonSpecimenDefinitionTypeTestedContainerAdditive) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	if m.AdditiveCodeableConcept != nil {
		if r.Additive != nil {
			return fmt.Errorf("multiple values for field \"Additive\"")
		}
		v := m.AdditiveCodeableConcept
		r.Additive = v
	}
	if m.AdditiveReference != nil {
		if r.Additive != nil {
			return fmt.Errorf("multiple values for field \"Additive\"")
		}
		v := m.AdditiveReference
		r.Additive = v
	}
	return nil
}
func (r SpecimenDefinitionTypeTestedContainerAdditive) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Set of instructions for preservation/transport of the specimen at a defined temperature interval, prior the testing process.
type SpecimenDefinitionTypeTestedHandling struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// It qualifies the interval of temperature, which characterizes an occurrence of handling. Conditions that are not related to temperature may be handled in the instruction element.
	TemperatureQualifier *CodeableConcept
	// The temperature interval for this set of handling instructions.
	TemperatureRange *Range
	// The maximum time interval of preservation of the specimen with these conditions.
	MaxDuration *Duration
	// Additional textual instructions for the preservation or transport of the specimen. For instance, 'Protect from light exposure'.
	Instruction *String
}
type jsonSpecimenDefinitionTypeTestedHandling struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	TemperatureQualifier        *CodeableConcept  `json:"temperatureQualifier,omitempty"`
	TemperatureRange            *Range            `json:"temperatureRange,omitempty"`
	MaxDuration                 *Duration         `json:"maxDuration,omitempty"`
	Instruction                 *String           `json:"instruction,omitempty"`
	InstructionPrimitiveElement *primitiveElement `json:"_instruction,omitempty"`
}

func (r SpecimenDefinitionTypeTestedHandling) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r SpecimenDefinitionTypeTestedHandling) marshalJSON() jsonSpecimenDefinitionTypeTestedHandling {
	m := jsonSpecimenDefinitionTypeTestedHandling{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.TemperatureQualifier = r.TemperatureQualifier
	m.TemperatureRange = r.TemperatureRange
	m.MaxDuration = r.MaxDuration
	m.Instruction = r.Instruction
	if r.Instruction != nil && (r.Instruction.Id != nil || r.Instruction.Extension != nil) {
		m.InstructionPrimitiveElement = &primitiveElement{Id: r.Instruction.Id, Extension: r.Instruction.Extension}
	}
	return m
}
func (r *SpecimenDefinitionTypeTestedHandling) UnmarshalJSON(b []byte) error {
	var m jsonSpecimenDefinitionTypeTestedHandling
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *SpecimenDefinitionTypeTestedHandling) unmarshalJSON(m jsonSpecimenDefinitionTypeTestedHandling) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.TemperatureQualifier = m.TemperatureQualifier
	r.TemperatureRange = m.TemperatureRange
	r.MaxDuration = m.MaxDuration
	r.Instruction = m.Instruction
	if m.InstructionPrimitiveElement != nil {
		r.Instruction.Id = m.InstructionPrimitiveElement.Id
		r.Instruction.Extension = m.InstructionPrimitiveElement.Extension
	}
	return nil
}
func (r SpecimenDefinitionTypeTestedHandling) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
