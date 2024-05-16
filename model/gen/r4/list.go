package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// A list is a curated collection of resources.
type List struct {
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
	// Identifier for the List assigned for business purposes outside the context of FHIR.
	Identifier []Identifier
	// Indicates the current state of this list.
	Status Code
	// How this list was prepared - whether it is a working list that is suitable for being maintained on an ongoing basis, or if it represents a snapshot of a list of items from another source, or whether it is a prepared list where items may be marked as added, modified or deleted.
	Mode Code
	// A label for the list assigned by the author.
	Title *String
	// This code defines the purpose of the list - why it was created.
	Code *CodeableConcept
	// The common subject (or patient) of the resources that are in the list if there is one.
	Subject *Reference
	// The encounter that is the context in which this list was created.
	Encounter *Reference
	// The date that the list was prepared.
	Date *DateTime
	// The entity responsible for deciding what the contents of the list were. Where the list was created by a human, this is the same as the author of the list.
	Source *Reference
	// What order applies to the items in the list.
	OrderedBy *CodeableConcept
	// Comments that apply to the overall list.
	Note []Annotation
	// Entries in this list.
	Entry []ListEntry
	// If the list is empty, why the list is empty.
	EmptyReason *CodeableConcept
}

func (r List) ResourceType() string {
	return "List"
}
func (r List) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Id == nil {
		return "", false
	}
	return *r.Id.Id, true
}

type jsonList struct {
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
	Status                        Code                `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement   `json:"_status,omitempty"`
	Mode                          Code                `json:"mode,omitempty"`
	ModePrimitiveElement          *primitiveElement   `json:"_mode,omitempty"`
	Title                         *String             `json:"title,omitempty"`
	TitlePrimitiveElement         *primitiveElement   `json:"_title,omitempty"`
	Code                          *CodeableConcept    `json:"code,omitempty"`
	Subject                       *Reference          `json:"subject,omitempty"`
	Encounter                     *Reference          `json:"encounter,omitempty"`
	Date                          *DateTime           `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement   `json:"_date,omitempty"`
	Source                        *Reference          `json:"source,omitempty"`
	OrderedBy                     *CodeableConcept    `json:"orderedBy,omitempty"`
	Note                          []Annotation        `json:"note,omitempty"`
	Entry                         []ListEntry         `json:"entry,omitempty"`
	EmptyReason                   *CodeableConcept    `json:"emptyReason,omitempty"`
}

func (r List) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r List) marshalJSON() jsonList {
	m := jsonList{}
	m.ResourceType = "List"
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
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Mode = r.Mode
	if r.Mode.Id != nil || r.Mode.Extension != nil {
		m.ModePrimitiveElement = &primitiveElement{Id: r.Mode.Id, Extension: r.Mode.Extension}
	}
	m.Title = r.Title
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	m.Code = r.Code
	m.Subject = r.Subject
	m.Encounter = r.Encounter
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Source = r.Source
	m.OrderedBy = r.OrderedBy
	m.Note = r.Note
	m.Entry = r.Entry
	m.EmptyReason = r.EmptyReason
	return m
}
func (r *List) UnmarshalJSON(b []byte) error {
	var m jsonList
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *List) unmarshalJSON(m jsonList) error {
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
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Mode = m.Mode
	if m.ModePrimitiveElement != nil {
		r.Mode.Id = m.ModePrimitiveElement.Id
		r.Mode.Extension = m.ModePrimitiveElement.Extension
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Code = m.Code
	r.Subject = m.Subject
	r.Encounter = m.Encounter
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Source = m.Source
	r.OrderedBy = m.OrderedBy
	r.Note = m.Note
	r.Entry = m.Entry
	r.EmptyReason = m.EmptyReason
	return nil
}
func (r List) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Entries in this list.
type ListEntry struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The flag allows the system constructing the list to indicate the role and significance of the item in the list.
	Flag *CodeableConcept
	// True if this item is marked as deleted in the list.
	Deleted *Boolean
	// When this item was added to the list.
	Date *DateTime
	// A reference to the actual resource from which data was derived.
	Item Reference
}
type jsonListEntry struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Flag                    *CodeableConcept  `json:"flag,omitempty"`
	Deleted                 *Boolean          `json:"deleted,omitempty"`
	DeletedPrimitiveElement *primitiveElement `json:"_deleted,omitempty"`
	Date                    *DateTime         `json:"date,omitempty"`
	DatePrimitiveElement    *primitiveElement `json:"_date,omitempty"`
	Item                    Reference         `json:"item,omitempty"`
}

func (r ListEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ListEntry) marshalJSON() jsonListEntry {
	m := jsonListEntry{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Flag = r.Flag
	m.Deleted = r.Deleted
	if r.Deleted != nil && (r.Deleted.Id != nil || r.Deleted.Extension != nil) {
		m.DeletedPrimitiveElement = &primitiveElement{Id: r.Deleted.Id, Extension: r.Deleted.Extension}
	}
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Item = r.Item
	return m
}
func (r *ListEntry) UnmarshalJSON(b []byte) error {
	var m jsonListEntry
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ListEntry) unmarshalJSON(m jsonListEntry) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Flag = m.Flag
	r.Deleted = m.Deleted
	if m.DeletedPrimitiveElement != nil {
		r.Deleted.Id = m.DeletedPrimitiveElement.Id
		r.Deleted.Extension = m.DeletedPrimitiveElement.Extension
	}
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Item = m.Item
	return nil
}
func (r ListEntry) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
