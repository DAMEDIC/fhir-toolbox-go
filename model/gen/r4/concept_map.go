package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
	"fmt"
)

// A statement of relationships from one set of concepts to one or more other concepts - either concepts in code systems, or data element/data element concepts, or classes in class models.
type ConceptMap struct {
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
	// An absolute URI that is used to identify this concept map when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this concept map is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the concept map is stored on different servers.
	Url *Uri
	// A formal identifier that is used to identify this concept map when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier *Identifier
	// The identifier that is used to identify this version of the concept map when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the concept map author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the concept map. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// A short, descriptive, user-friendly title for the concept map.
	Title *String
	// The status of this concept map. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this concept map is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the concept map was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the concept map changes.
	Date *DateTime
	// The name of the organization or individual that published the concept map.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// A free text natural language description of the concept map from a consumer's perspective.
	Description *Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate concept map instances.
	UseContext []UsageContext
	// A legal or geographic region in which the concept map is intended to be used.
	Jurisdiction []CodeableConcept
	// Explanation of why this concept map is needed and why it has been designed as it has.
	Purpose *Markdown
	// A copyright statement relating to the concept map and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the concept map.
	Copyright *Markdown
	// Identifier for the source value set that contains the concepts that are being mapped and provides context for the mappings.
	Source isConceptMapSource
	// The target value set provides context for the mappings. Note that the mapping is made between concepts, not between value sets, but the value set provides important context about how the concept mapping choices are made.
	Target isConceptMapTarget
	// A group of mappings that all have the same source and target system.
	Group []ConceptMapGroup
}

func (r ConceptMap) ResourceType() string {
	return "ConceptMap"
}
func (r ConceptMap) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type isConceptMapSource interface {
	isConceptMapSource()
}

func (r Uri) isConceptMapSource()       {}
func (r Canonical) isConceptMapSource() {}

type isConceptMapTarget interface {
	isConceptMapTarget()
}

func (r Uri) isConceptMapTarget()       {}
func (r Canonical) isConceptMapTarget() {}

type jsonConceptMap struct {
	ResourceType                    string              `json:"resourceType"`
	Id                              *Id                 `json:"id,omitempty"`
	IdPrimitiveElement              *primitiveElement   `json:"_id,omitempty"`
	Meta                            *Meta               `json:"meta,omitempty"`
	ImplicitRules                   *Uri                `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement   *primitiveElement   `json:"_implicitRules,omitempty"`
	Language                        *Code               `json:"language,omitempty"`
	LanguagePrimitiveElement        *primitiveElement   `json:"_language,omitempty"`
	Text                            *Narrative          `json:"text,omitempty"`
	Contained                       []containedResource `json:"contained,omitempty"`
	Extension                       []Extension         `json:"extension,omitempty"`
	ModifierExtension               []Extension         `json:"modifierExtension,omitempty"`
	Url                             *Uri                `json:"url,omitempty"`
	UrlPrimitiveElement             *primitiveElement   `json:"_url,omitempty"`
	Identifier                      *Identifier         `json:"identifier,omitempty"`
	Version                         *String             `json:"version,omitempty"`
	VersionPrimitiveElement         *primitiveElement   `json:"_version,omitempty"`
	Name                            *String             `json:"name,omitempty"`
	NamePrimitiveElement            *primitiveElement   `json:"_name,omitempty"`
	Title                           *String             `json:"title,omitempty"`
	TitlePrimitiveElement           *primitiveElement   `json:"_title,omitempty"`
	Status                          Code                `json:"status,omitempty"`
	StatusPrimitiveElement          *primitiveElement   `json:"_status,omitempty"`
	Experimental                    *Boolean            `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement    *primitiveElement   `json:"_experimental,omitempty"`
	Date                            *DateTime           `json:"date,omitempty"`
	DatePrimitiveElement            *primitiveElement   `json:"_date,omitempty"`
	Publisher                       *String             `json:"publisher,omitempty"`
	PublisherPrimitiveElement       *primitiveElement   `json:"_publisher,omitempty"`
	Contact                         []ContactDetail     `json:"contact,omitempty"`
	Description                     *Markdown           `json:"description,omitempty"`
	DescriptionPrimitiveElement     *primitiveElement   `json:"_description,omitempty"`
	UseContext                      []UsageContext      `json:"useContext,omitempty"`
	Jurisdiction                    []CodeableConcept   `json:"jurisdiction,omitempty"`
	Purpose                         *Markdown           `json:"purpose,omitempty"`
	PurposePrimitiveElement         *primitiveElement   `json:"_purpose,omitempty"`
	Copyright                       *Markdown           `json:"copyright,omitempty"`
	CopyrightPrimitiveElement       *primitiveElement   `json:"_copyright,omitempty"`
	SourceUri                       *Uri                `json:"sourceUri,omitempty"`
	SourceUriPrimitiveElement       *primitiveElement   `json:"_sourceUri,omitempty"`
	SourceCanonical                 *Canonical          `json:"sourceCanonical,omitempty"`
	SourceCanonicalPrimitiveElement *primitiveElement   `json:"_sourceCanonical,omitempty"`
	TargetUri                       *Uri                `json:"targetUri,omitempty"`
	TargetUriPrimitiveElement       *primitiveElement   `json:"_targetUri,omitempty"`
	TargetCanonical                 *Canonical          `json:"targetCanonical,omitempty"`
	TargetCanonicalPrimitiveElement *primitiveElement   `json:"_targetCanonical,omitempty"`
	Group                           []ConceptMapGroup   `json:"group,omitempty"`
}

func (r ConceptMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ConceptMap) marshalJSON() jsonConceptMap {
	m := jsonConceptMap{}
	m.ResourceType = "ConceptMap"
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
	m.Url = r.Url
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	m.Identifier = r.Identifier
	m.Version = r.Version
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	m.Name = r.Name
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	m.Title = r.Title
	if r.Title != nil && (r.Title.Id != nil || r.Title.Extension != nil) {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	m.Status = r.Status
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	m.Experimental = r.Experimental
	if r.Experimental != nil && (r.Experimental.Id != nil || r.Experimental.Extension != nil) {
		m.ExperimentalPrimitiveElement = &primitiveElement{Id: r.Experimental.Id, Extension: r.Experimental.Extension}
	}
	m.Date = r.Date
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	m.Publisher = r.Publisher
	if r.Publisher != nil && (r.Publisher.Id != nil || r.Publisher.Extension != nil) {
		m.PublisherPrimitiveElement = &primitiveElement{Id: r.Publisher.Id, Extension: r.Publisher.Extension}
	}
	m.Contact = r.Contact
	m.Description = r.Description
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.UseContext = r.UseContext
	m.Jurisdiction = r.Jurisdiction
	m.Purpose = r.Purpose
	if r.Purpose != nil && (r.Purpose.Id != nil || r.Purpose.Extension != nil) {
		m.PurposePrimitiveElement = &primitiveElement{Id: r.Purpose.Id, Extension: r.Purpose.Extension}
	}
	m.Copyright = r.Copyright
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		m.CopyrightPrimitiveElement = &primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
	}
	switch v := r.Source.(type) {
	case Uri:
		m.SourceUri = &v
		if v.Id != nil || v.Extension != nil {
			m.SourceUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uri:
		m.SourceUri = v
		if v.Id != nil || v.Extension != nil {
			m.SourceUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Canonical:
		m.SourceCanonical = &v
		if v.Id != nil || v.Extension != nil {
			m.SourceCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Canonical:
		m.SourceCanonical = v
		if v.Id != nil || v.Extension != nil {
			m.SourceCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	switch v := r.Target.(type) {
	case Uri:
		m.TargetUri = &v
		if v.Id != nil || v.Extension != nil {
			m.TargetUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Uri:
		m.TargetUri = v
		if v.Id != nil || v.Extension != nil {
			m.TargetUriPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case Canonical:
		m.TargetCanonical = &v
		if v.Id != nil || v.Extension != nil {
			m.TargetCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	case *Canonical:
		m.TargetCanonical = v
		if v.Id != nil || v.Extension != nil {
			m.TargetCanonicalPrimitiveElement = &primitiveElement{Id: v.Id, Extension: v.Extension}
		}
	}
	m.Group = r.Group
	return m
}
func (r *ConceptMap) UnmarshalJSON(b []byte) error {
	var m jsonConceptMap
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ConceptMap) unmarshalJSON(m jsonConceptMap) error {
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
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Identifier = m.Identifier
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Experimental = m.Experimental
	if m.ExperimentalPrimitiveElement != nil {
		r.Experimental.Id = m.ExperimentalPrimitiveElement.Id
		r.Experimental.Extension = m.ExperimentalPrimitiveElement.Extension
	}
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Publisher = m.Publisher
	if m.PublisherPrimitiveElement != nil {
		r.Publisher.Id = m.PublisherPrimitiveElement.Id
		r.Publisher.Extension = m.PublisherPrimitiveElement.Extension
	}
	r.Contact = m.Contact
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.UseContext = m.UseContext
	r.Jurisdiction = m.Jurisdiction
	r.Purpose = m.Purpose
	if m.PurposePrimitiveElement != nil {
		r.Purpose.Id = m.PurposePrimitiveElement.Id
		r.Purpose.Extension = m.PurposePrimitiveElement.Extension
	}
	r.Copyright = m.Copyright
	if m.CopyrightPrimitiveElement != nil {
		r.Copyright.Id = m.CopyrightPrimitiveElement.Id
		r.Copyright.Extension = m.CopyrightPrimitiveElement.Extension
	}
	if m.SourceUri != nil || m.SourceUriPrimitiveElement != nil {
		if r.Source != nil {
			return fmt.Errorf("multiple values for field \"Source\"")
		}
		v := m.SourceUri
		if m.SourceUriPrimitiveElement != nil {
			if v == nil {
				v = &Uri{}
			}
			v.Id = m.SourceUriPrimitiveElement.Id
			v.Extension = m.SourceUriPrimitiveElement.Extension
		}
		r.Source = v
	}
	if m.SourceCanonical != nil || m.SourceCanonicalPrimitiveElement != nil {
		if r.Source != nil {
			return fmt.Errorf("multiple values for field \"Source\"")
		}
		v := m.SourceCanonical
		if m.SourceCanonicalPrimitiveElement != nil {
			if v == nil {
				v = &Canonical{}
			}
			v.Id = m.SourceCanonicalPrimitiveElement.Id
			v.Extension = m.SourceCanonicalPrimitiveElement.Extension
		}
		r.Source = v
	}
	if m.TargetUri != nil || m.TargetUriPrimitiveElement != nil {
		if r.Target != nil {
			return fmt.Errorf("multiple values for field \"Target\"")
		}
		v := m.TargetUri
		if m.TargetUriPrimitiveElement != nil {
			if v == nil {
				v = &Uri{}
			}
			v.Id = m.TargetUriPrimitiveElement.Id
			v.Extension = m.TargetUriPrimitiveElement.Extension
		}
		r.Target = v
	}
	if m.TargetCanonical != nil || m.TargetCanonicalPrimitiveElement != nil {
		if r.Target != nil {
			return fmt.Errorf("multiple values for field \"Target\"")
		}
		v := m.TargetCanonical
		if m.TargetCanonicalPrimitiveElement != nil {
			if v == nil {
				v = &Canonical{}
			}
			v.Id = m.TargetCanonicalPrimitiveElement.Id
			v.Extension = m.TargetCanonicalPrimitiveElement.Extension
		}
		r.Target = v
	}
	r.Group = m.Group
	return nil
}
func (r ConceptMap) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A group of mappings that all have the same source and target system.
type ConceptMapGroup struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// An absolute URI that identifies the source system where the concepts to be mapped are defined.
	Source *Uri
	// The specific version of the code system, as determined by the code system authority.
	SourceVersion *String
	// An absolute URI that identifies the target system that the concepts will be mapped to.
	Target *Uri
	// The specific version of the code system, as determined by the code system authority.
	TargetVersion *String
	// Mappings for an individual concept in the source to one or more concepts in the target.
	Element []ConceptMapGroupElement
	// What to do when there is no mapping for the source concept. "Unmapped" does not include codes that are unmatched, and the unmapped element is ignored in a code is specified to have equivalence = unmatched.
	Unmapped *ConceptMapGroupUnmapped
}
type jsonConceptMapGroup struct {
	Id                            *string                  `json:"id,omitempty"`
	Extension                     []Extension              `json:"extension,omitempty"`
	ModifierExtension             []Extension              `json:"modifierExtension,omitempty"`
	Source                        *Uri                     `json:"source,omitempty"`
	SourcePrimitiveElement        *primitiveElement        `json:"_source,omitempty"`
	SourceVersion                 *String                  `json:"sourceVersion,omitempty"`
	SourceVersionPrimitiveElement *primitiveElement        `json:"_sourceVersion,omitempty"`
	Target                        *Uri                     `json:"target,omitempty"`
	TargetPrimitiveElement        *primitiveElement        `json:"_target,omitempty"`
	TargetVersion                 *String                  `json:"targetVersion,omitempty"`
	TargetVersionPrimitiveElement *primitiveElement        `json:"_targetVersion,omitempty"`
	Element                       []ConceptMapGroupElement `json:"element,omitempty"`
	Unmapped                      *ConceptMapGroupUnmapped `json:"unmapped,omitempty"`
}

func (r ConceptMapGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ConceptMapGroup) marshalJSON() jsonConceptMapGroup {
	m := jsonConceptMapGroup{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Source = r.Source
	if r.Source != nil && (r.Source.Id != nil || r.Source.Extension != nil) {
		m.SourcePrimitiveElement = &primitiveElement{Id: r.Source.Id, Extension: r.Source.Extension}
	}
	m.SourceVersion = r.SourceVersion
	if r.SourceVersion != nil && (r.SourceVersion.Id != nil || r.SourceVersion.Extension != nil) {
		m.SourceVersionPrimitiveElement = &primitiveElement{Id: r.SourceVersion.Id, Extension: r.SourceVersion.Extension}
	}
	m.Target = r.Target
	if r.Target != nil && (r.Target.Id != nil || r.Target.Extension != nil) {
		m.TargetPrimitiveElement = &primitiveElement{Id: r.Target.Id, Extension: r.Target.Extension}
	}
	m.TargetVersion = r.TargetVersion
	if r.TargetVersion != nil && (r.TargetVersion.Id != nil || r.TargetVersion.Extension != nil) {
		m.TargetVersionPrimitiveElement = &primitiveElement{Id: r.TargetVersion.Id, Extension: r.TargetVersion.Extension}
	}
	m.Element = r.Element
	m.Unmapped = r.Unmapped
	return m
}
func (r *ConceptMapGroup) UnmarshalJSON(b []byte) error {
	var m jsonConceptMapGroup
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ConceptMapGroup) unmarshalJSON(m jsonConceptMapGroup) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Source = m.Source
	if m.SourcePrimitiveElement != nil {
		r.Source.Id = m.SourcePrimitiveElement.Id
		r.Source.Extension = m.SourcePrimitiveElement.Extension
	}
	r.SourceVersion = m.SourceVersion
	if m.SourceVersionPrimitiveElement != nil {
		r.SourceVersion.Id = m.SourceVersionPrimitiveElement.Id
		r.SourceVersion.Extension = m.SourceVersionPrimitiveElement.Extension
	}
	r.Target = m.Target
	if m.TargetPrimitiveElement != nil {
		r.Target.Id = m.TargetPrimitiveElement.Id
		r.Target.Extension = m.TargetPrimitiveElement.Extension
	}
	r.TargetVersion = m.TargetVersion
	if m.TargetVersionPrimitiveElement != nil {
		r.TargetVersion.Id = m.TargetVersionPrimitiveElement.Id
		r.TargetVersion.Extension = m.TargetVersionPrimitiveElement.Extension
	}
	r.Element = m.Element
	r.Unmapped = m.Unmapped
	return nil
}
func (r ConceptMapGroup) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Mappings for an individual concept in the source to one or more concepts in the target.
type ConceptMapGroupElement struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identity (code or path) or the element/item being mapped.
	Code *Code
	// The display for the code. The display is only provided to help editors when editing the concept map.
	Display *String
	// A concept from the target value set that this concept maps to.
	Target []ConceptMapGroupElementTarget
}
type jsonConceptMapGroupElement struct {
	Id                      *string                        `json:"id,omitempty"`
	Extension               []Extension                    `json:"extension,omitempty"`
	ModifierExtension       []Extension                    `json:"modifierExtension,omitempty"`
	Code                    *Code                          `json:"code,omitempty"`
	CodePrimitiveElement    *primitiveElement              `json:"_code,omitempty"`
	Display                 *String                        `json:"display,omitempty"`
	DisplayPrimitiveElement *primitiveElement              `json:"_display,omitempty"`
	Target                  []ConceptMapGroupElementTarget `json:"target,omitempty"`
}

func (r ConceptMapGroupElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ConceptMapGroupElement) marshalJSON() jsonConceptMapGroupElement {
	m := jsonConceptMapGroupElement{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	if r.Code != nil && (r.Code.Id != nil || r.Code.Extension != nil) {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Display = r.Display
	if r.Display != nil && (r.Display.Id != nil || r.Display.Extension != nil) {
		m.DisplayPrimitiveElement = &primitiveElement{Id: r.Display.Id, Extension: r.Display.Extension}
	}
	m.Target = r.Target
	return m
}
func (r *ConceptMapGroupElement) UnmarshalJSON(b []byte) error {
	var m jsonConceptMapGroupElement
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ConceptMapGroupElement) unmarshalJSON(m jsonConceptMapGroupElement) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Display = m.Display
	if m.DisplayPrimitiveElement != nil {
		r.Display.Id = m.DisplayPrimitiveElement.Id
		r.Display.Extension = m.DisplayPrimitiveElement.Extension
	}
	r.Target = m.Target
	return nil
}
func (r ConceptMapGroupElement) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A concept from the target value set that this concept maps to.
type ConceptMapGroupElementTarget struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Identity (code or path) or the element/item that the map refers to.
	Code *Code
	// The display for the code. The display is only provided to help editors when editing the concept map.
	Display *String
	// The equivalence between the source and target concepts (counting for the dependencies and products). The equivalence is read from target to source (e.g. the target is 'wider' than the source).
	Equivalence Code
	// A description of status/issues in mapping that conveys additional information not represented in  the structured data.
	Comment *String
	// A set of additional dependencies for this mapping to hold. This mapping is only applicable if the specified element can be resolved, and it has the specified value.
	DependsOn []ConceptMapGroupElementTargetDependsOn
	// A set of additional outcomes from this mapping to other elements. To properly execute this mapping, the specified element must be mapped to some data element or source that is in context. The mapping may still be useful without a place for the additional data elements, but the equivalence cannot be relied on.
	Product []ConceptMapGroupElementTargetDependsOn
}
type jsonConceptMapGroupElementTarget struct {
	Id                          *string                                 `json:"id,omitempty"`
	Extension                   []Extension                             `json:"extension,omitempty"`
	ModifierExtension           []Extension                             `json:"modifierExtension,omitempty"`
	Code                        *Code                                   `json:"code,omitempty"`
	CodePrimitiveElement        *primitiveElement                       `json:"_code,omitempty"`
	Display                     *String                                 `json:"display,omitempty"`
	DisplayPrimitiveElement     *primitiveElement                       `json:"_display,omitempty"`
	Equivalence                 Code                                    `json:"equivalence,omitempty"`
	EquivalencePrimitiveElement *primitiveElement                       `json:"_equivalence,omitempty"`
	Comment                     *String                                 `json:"comment,omitempty"`
	CommentPrimitiveElement     *primitiveElement                       `json:"_comment,omitempty"`
	DependsOn                   []ConceptMapGroupElementTargetDependsOn `json:"dependsOn,omitempty"`
	Product                     []ConceptMapGroupElementTargetDependsOn `json:"product,omitempty"`
}

func (r ConceptMapGroupElementTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ConceptMapGroupElementTarget) marshalJSON() jsonConceptMapGroupElementTarget {
	m := jsonConceptMapGroupElementTarget{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Code = r.Code
	if r.Code != nil && (r.Code.Id != nil || r.Code.Extension != nil) {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Display = r.Display
	if r.Display != nil && (r.Display.Id != nil || r.Display.Extension != nil) {
		m.DisplayPrimitiveElement = &primitiveElement{Id: r.Display.Id, Extension: r.Display.Extension}
	}
	m.Equivalence = r.Equivalence
	if r.Equivalence.Id != nil || r.Equivalence.Extension != nil {
		m.EquivalencePrimitiveElement = &primitiveElement{Id: r.Equivalence.Id, Extension: r.Equivalence.Extension}
	}
	m.Comment = r.Comment
	if r.Comment != nil && (r.Comment.Id != nil || r.Comment.Extension != nil) {
		m.CommentPrimitiveElement = &primitiveElement{Id: r.Comment.Id, Extension: r.Comment.Extension}
	}
	m.DependsOn = r.DependsOn
	m.Product = r.Product
	return m
}
func (r *ConceptMapGroupElementTarget) UnmarshalJSON(b []byte) error {
	var m jsonConceptMapGroupElementTarget
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ConceptMapGroupElementTarget) unmarshalJSON(m jsonConceptMapGroupElementTarget) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Display = m.Display
	if m.DisplayPrimitiveElement != nil {
		r.Display.Id = m.DisplayPrimitiveElement.Id
		r.Display.Extension = m.DisplayPrimitiveElement.Extension
	}
	r.Equivalence = m.Equivalence
	if m.EquivalencePrimitiveElement != nil {
		r.Equivalence.Id = m.EquivalencePrimitiveElement.Id
		r.Equivalence.Extension = m.EquivalencePrimitiveElement.Extension
	}
	r.Comment = m.Comment
	if m.CommentPrimitiveElement != nil {
		r.Comment.Id = m.CommentPrimitiveElement.Id
		r.Comment.Extension = m.CommentPrimitiveElement.Extension
	}
	r.DependsOn = m.DependsOn
	r.Product = m.Product
	return nil
}
func (r ConceptMapGroupElementTarget) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A set of additional dependencies for this mapping to hold. This mapping is only applicable if the specified element can be resolved, and it has the specified value.
type ConceptMapGroupElementTargetDependsOn struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// A reference to an element that holds a coded value that corresponds to a code system property. The idea is that the information model carries an element somewhere that is labeled to correspond with a code system property.
	Property Uri
	// An absolute URI that identifies the code system of the dependency code (if the source/dependency is a value set that crosses code systems).
	System *Canonical
	// Identity (code or path) or the element/item/ValueSet/text that the map depends on / refers to.
	Value String
	// The display for the code. The display is only provided to help editors when editing the concept map.
	Display *String
}
type jsonConceptMapGroupElementTargetDependsOn struct {
	Id                       *string           `json:"id,omitempty"`
	Extension                []Extension       `json:"extension,omitempty"`
	ModifierExtension        []Extension       `json:"modifierExtension,omitempty"`
	Property                 Uri               `json:"property,omitempty"`
	PropertyPrimitiveElement *primitiveElement `json:"_property,omitempty"`
	System                   *Canonical        `json:"system,omitempty"`
	SystemPrimitiveElement   *primitiveElement `json:"_system,omitempty"`
	Value                    String            `json:"value,omitempty"`
	ValuePrimitiveElement    *primitiveElement `json:"_value,omitempty"`
	Display                  *String           `json:"display,omitempty"`
	DisplayPrimitiveElement  *primitiveElement `json:"_display,omitempty"`
}

func (r ConceptMapGroupElementTargetDependsOn) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ConceptMapGroupElementTargetDependsOn) marshalJSON() jsonConceptMapGroupElementTargetDependsOn {
	m := jsonConceptMapGroupElementTargetDependsOn{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Property = r.Property
	if r.Property.Id != nil || r.Property.Extension != nil {
		m.PropertyPrimitiveElement = &primitiveElement{Id: r.Property.Id, Extension: r.Property.Extension}
	}
	m.System = r.System
	if r.System != nil && (r.System.Id != nil || r.System.Extension != nil) {
		m.SystemPrimitiveElement = &primitiveElement{Id: r.System.Id, Extension: r.System.Extension}
	}
	m.Value = r.Value
	if r.Value.Id != nil || r.Value.Extension != nil {
		m.ValuePrimitiveElement = &primitiveElement{Id: r.Value.Id, Extension: r.Value.Extension}
	}
	m.Display = r.Display
	if r.Display != nil && (r.Display.Id != nil || r.Display.Extension != nil) {
		m.DisplayPrimitiveElement = &primitiveElement{Id: r.Display.Id, Extension: r.Display.Extension}
	}
	return m
}
func (r *ConceptMapGroupElementTargetDependsOn) UnmarshalJSON(b []byte) error {
	var m jsonConceptMapGroupElementTargetDependsOn
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ConceptMapGroupElementTargetDependsOn) unmarshalJSON(m jsonConceptMapGroupElementTargetDependsOn) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Property = m.Property
	if m.PropertyPrimitiveElement != nil {
		r.Property.Id = m.PropertyPrimitiveElement.Id
		r.Property.Extension = m.PropertyPrimitiveElement.Extension
	}
	r.System = m.System
	if m.SystemPrimitiveElement != nil {
		r.System.Id = m.SystemPrimitiveElement.Id
		r.System.Extension = m.SystemPrimitiveElement.Extension
	}
	r.Value = m.Value
	if m.ValuePrimitiveElement != nil {
		r.Value.Id = m.ValuePrimitiveElement.Id
		r.Value.Extension = m.ValuePrimitiveElement.Extension
	}
	r.Display = m.Display
	if m.DisplayPrimitiveElement != nil {
		r.Display.Id = m.DisplayPrimitiveElement.Id
		r.Display.Extension = m.DisplayPrimitiveElement.Extension
	}
	return nil
}
func (r ConceptMapGroupElementTargetDependsOn) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// What to do when there is no mapping for the source concept. "Unmapped" does not include codes that are unmatched, and the unmapped element is ignored in a code is specified to have equivalence = unmatched.
type ConceptMapGroupUnmapped struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Defines which action to take if there is no match for the source concept in the target system designated for the group. One of 3 actions are possible: use the unmapped code (this is useful when doing a mapping between versions, and only a few codes have changed), use a fixed code (a default code), or alternatively, a reference to a different concept map can be provided (by canonical URL).
	Mode Code
	// The fixed code to use when the mode = 'fixed'  - all unmapped codes are mapped to a single fixed code.
	Code *Code
	// The display for the code. The display is only provided to help editors when editing the concept map.
	Display *String
	// The canonical reference to an additional ConceptMap resource instance to use for mapping if this ConceptMap resource contains no matching mapping for the source concept.
	Url *Canonical
}
type jsonConceptMapGroupUnmapped struct {
	Id                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Mode                    Code              `json:"mode,omitempty"`
	ModePrimitiveElement    *primitiveElement `json:"_mode,omitempty"`
	Code                    *Code             `json:"code,omitempty"`
	CodePrimitiveElement    *primitiveElement `json:"_code,omitempty"`
	Display                 *String           `json:"display,omitempty"`
	DisplayPrimitiveElement *primitiveElement `json:"_display,omitempty"`
	Url                     *Canonical        `json:"url,omitempty"`
	UrlPrimitiveElement     *primitiveElement `json:"_url,omitempty"`
}

func (r ConceptMapGroupUnmapped) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ConceptMapGroupUnmapped) marshalJSON() jsonConceptMapGroupUnmapped {
	m := jsonConceptMapGroupUnmapped{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Mode = r.Mode
	if r.Mode.Id != nil || r.Mode.Extension != nil {
		m.ModePrimitiveElement = &primitiveElement{Id: r.Mode.Id, Extension: r.Mode.Extension}
	}
	m.Code = r.Code
	if r.Code != nil && (r.Code.Id != nil || r.Code.Extension != nil) {
		m.CodePrimitiveElement = &primitiveElement{Id: r.Code.Id, Extension: r.Code.Extension}
	}
	m.Display = r.Display
	if r.Display != nil && (r.Display.Id != nil || r.Display.Extension != nil) {
		m.DisplayPrimitiveElement = &primitiveElement{Id: r.Display.Id, Extension: r.Display.Extension}
	}
	m.Url = r.Url
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	return m
}
func (r *ConceptMapGroupUnmapped) UnmarshalJSON(b []byte) error {
	var m jsonConceptMapGroupUnmapped
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ConceptMapGroupUnmapped) unmarshalJSON(m jsonConceptMapGroupUnmapped) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Mode = m.Mode
	if m.ModePrimitiveElement != nil {
		r.Mode.Id = m.ModePrimitiveElement.Id
		r.Mode.Extension = m.ModePrimitiveElement.Extension
	}
	r.Code = m.Code
	if m.CodePrimitiveElement != nil {
		r.Code.Id = m.CodePrimitiveElement.Id
		r.Code.Extension = m.CodePrimitiveElement.Extension
	}
	r.Display = m.Display
	if m.DisplayPrimitiveElement != nil {
		r.Display.Id = m.DisplayPrimitiveElement.Id
		r.Display.Extension = m.DisplayPrimitiveElement.Extension
	}
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	return nil
}
func (r ConceptMapGroupUnmapped) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
