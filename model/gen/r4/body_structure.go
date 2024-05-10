package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// Record details about an anatomical structure.  This resource may be used when a coded concept does not provide the necessary detail needed for the use case.
type BodyStructure struct {
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
	// Identifier for this instance of the anatomical structure.
	Identifier []Identifier
	// Whether this body site is in active use.
	Active *Boolean
	// The kind of structure being represented by the body structure at `BodyStructure.location`.  This can define both normal and abnormal morphologies.
	Morphology *CodeableConcept
	// The anatomical location or region of the specimen, lesion, or body structure.
	Location *CodeableConcept
	// Qualifier to refine the anatomical location.  These include qualifiers for laterality, relative location, directionality, number, and plane.
	LocationQualifier []CodeableConcept
	// A summary, characterization or explanation of the body structure.
	Description *String
	// Image or images used to identify a location.
	Image []Attachment
	// The person to which the body site belongs.
	Patient Reference
}
type jsonBodyStructure struct {
	ResourceType                  string              `json:"resourceType"`
	Id                            *Id                 `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement   `json:"_id,omitempty"`
	Meta                          *Meta               `json:"meta,omitempty"`
	ImplicitRules                 *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                      *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement   `json:"_language,omitempty"`
	Text                          *Narrative          `json:"text,omitempty"`
	Contained                     []containedResource `json:"contained,omitempty"`
	Extension                     []Extension         `json:"extension,omitempty"`
	ModifierExtension             []Extension         `json:"modifierExtension,omitempty"`
	Identifier                    []Identifier        `json:"identifier,omitempty"`
	Active                        *Boolean            `json:"active,omitempty"`
	ActivePrimitiveElement        *primitiveElement   `json:"_active,omitempty"`
	Morphology                    *CodeableConcept    `json:"morphology,omitempty"`
	Location                      *CodeableConcept    `json:"location,omitempty"`
	LocationQualifier             []CodeableConcept   `json:"locationQualifier,omitempty"`
	Description                   *String             `json:"description,omitempty"`
	DescriptionPrimitiveElement   *primitiveElement   `json:"_description,omitempty"`
	Image                         []Attachment        `json:"image,omitempty"`
	Patient                       Reference           `json:"patient,omitempty"`
}

func (r BodyStructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r BodyStructure) marshalJSON() jsonBodyStructure {
	m := jsonBodyStructure{}
	m.ResourceType = "BodyStructure"
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
	m.Active = r.Active
	if r.Active != nil && (r.Active.Id != nil || r.Active.Extension != nil) {
		m.ActivePrimitiveElement = &primitiveElement{Id: r.Active.Id, Extension: r.Active.Extension}
	}
	m.Morphology = r.Morphology
	m.Location = r.Location
	m.LocationQualifier = r.LocationQualifier
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Image = r.Image
	m.Patient = r.Patient
	return m
}
func (r *BodyStructure) UnmarshalJSON(b []byte) error {
	var m jsonBodyStructure
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *BodyStructure) unmarshalJSON(m jsonBodyStructure) error {
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
	r.Active = m.Active
	if m.ActivePrimitiveElement != nil {
		r.Active.Id = m.ActivePrimitiveElement.Id
		r.Active.Extension = m.ActivePrimitiveElement.Extension
	}
	r.Morphology = m.Morphology
	r.Location = m.Location
	r.LocationQualifier = m.LocationQualifier
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Image = m.Image
	r.Patient = m.Patient
	return nil
}
func (r BodyStructure) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
