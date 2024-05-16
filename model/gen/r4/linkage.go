package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// Identifies two or more records (resource instances) that refer to the same real-world "occurrence".
type Linkage struct {
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
	// Indicates whether the asserted set of linkages are considered to be "in effect".
	Active *Boolean
	// Identifies the user or organization responsible for asserting the linkages as well as the user or organization who establishes the context in which the nature of each linkage is evaluated.
	Author *Reference
	// Identifies which record considered as the reference to the same real-world occurrence as well as how the items should be evaluated within the collection of linked items.
	Item []LinkageItem
}

func (r Linkage) ResourceType() string {
	return "Linkage"
}
func (r Linkage) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonLinkage struct {
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
	Active                        *Boolean            `json:"active,omitempty"`
	ActivePrimitiveElement        *primitiveElement   `json:"_active,omitempty"`
	Author                        *Reference          `json:"author,omitempty"`
	Item                          []LinkageItem       `json:"item,omitempty"`
}

func (r Linkage) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r Linkage) marshalJSON() jsonLinkage {
	m := jsonLinkage{}
	m.ResourceType = "Linkage"
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
	m.Active = r.Active
	if r.Active != nil && (r.Active.Id != nil || r.Active.Extension != nil) {
		m.ActivePrimitiveElement = &primitiveElement{Id: r.Active.Id, Extension: r.Active.Extension}
	}
	m.Author = r.Author
	m.Item = r.Item
	return m
}
func (r *Linkage) UnmarshalJSON(b []byte) error {
	var m jsonLinkage
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *Linkage) unmarshalJSON(m jsonLinkage) error {
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
	r.Active = m.Active
	if m.ActivePrimitiveElement != nil {
		r.Active.Id = m.ActivePrimitiveElement.Id
		r.Active.Extension = m.ActivePrimitiveElement.Extension
	}
	r.Author = m.Author
	r.Item = m.Item
	return nil
}
func (r Linkage) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Identifies which record considered as the reference to the same real-world occurrence as well as how the items should be evaluated within the collection of linked items.
type LinkageItem struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Distinguishes which item is "source of truth" (if any) and which items are no longer considered to be current representations.
	Type Code
	// The resource instance being linked as part of the group.
	Resource Reference
}
type jsonLinkageItem struct {
	Id                   *string           `json:"id,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Type                 Code              `json:"type,omitempty"`
	TypePrimitiveElement *primitiveElement `json:"_type,omitempty"`
	Resource             Reference         `json:"resource,omitempty"`
}

func (r LinkageItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r LinkageItem) marshalJSON() jsonLinkageItem {
	m := jsonLinkageItem{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Type = r.Type
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	m.Resource = r.Resource
	return m
}
func (r *LinkageItem) UnmarshalJSON(b []byte) error {
	var m jsonLinkageItem
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *LinkageItem) unmarshalJSON(m jsonLinkageItem) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Resource = m.Resource
	return nil
}
func (r LinkageItem) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
