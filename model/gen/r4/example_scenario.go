package r4

import (
	"encoding/json"
	model "fhir-toolbox/model"
)

// Example of workflow instance.
type ExampleScenario struct {
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
	// An absolute URI that is used to identify this example scenario when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this example scenario is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the example scenario is stored on different servers.
	Url *Uri
	// A formal identifier that is used to identify this example scenario when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []Identifier
	// The identifier that is used to identify this version of the example scenario when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the example scenario author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *String
	// A natural language name identifying the example scenario. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name *String
	// The status of this example scenario. Enables tracking the life-cycle of the content.
	Status Code
	// A Boolean value to indicate that this example scenario is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *Boolean
	// The date  (and optionally time) when the example scenario was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the example scenario changes. (e.g. the 'content logical definition').
	Date *DateTime
	// The name of the organization or individual that published the example scenario.
	Publisher *String
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []ContactDetail
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate example scenario instances.
	UseContext []UsageContext
	// A legal or geographic region in which the example scenario is intended to be used.
	Jurisdiction []CodeableConcept
	// A copyright statement relating to the example scenario and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the example scenario.
	Copyright *Markdown
	// What the example scenario resource is created for. This should not be used to show the business purpose of the scenario itself, but the purpose of documenting a scenario.
	Purpose *Markdown
	// Actor participating in the resource.
	Actor []ExampleScenarioActor
	// Each resource and each version that is present in the workflow.
	Instance []ExampleScenarioInstance
	// Each major process - a group of operations.
	Process []ExampleScenarioProcess
	// Another nested workflow.
	Workflow []Canonical
}

func (r ExampleScenario) ResourceType() string {
	return "ExampleScenario"
}
func (r ExampleScenario) ResourceId() (string, bool) {
	if r.Id == nil {
		return "", false
	}
	if r.Id.Value == nil {
		return "", false
	}
	return *r.Id.Value, true
}

type jsonExampleScenario struct {
	ResourceType                  string                    `json:"resourceType"`
	Id                            *Id                       `json:"id,omitempty"`
	IdPrimitiveElement            *primitiveElement         `json:"_id,omitempty"`
	Meta                          *Meta                     `json:"meta,omitempty"`
	ImplicitRules                 *Uri                      `json:"implicitRules,omitempty"`
	ImplicitRulesPrimitiveElement *primitiveElement         `json:"_implicitRules,omitempty"`
	Language                      *Code                     `json:"language,omitempty"`
	LanguagePrimitiveElement      *primitiveElement         `json:"_language,omitempty"`
	Text                          *Narrative                `json:"text,omitempty"`
	Contained                     []ContainedResource       `json:"contained,omitempty"`
	Extension                     []Extension               `json:"extension,omitempty"`
	ModifierExtension             []Extension               `json:"modifierExtension,omitempty"`
	Url                           *Uri                      `json:"url,omitempty"`
	UrlPrimitiveElement           *primitiveElement         `json:"_url,omitempty"`
	Identifier                    []Identifier              `json:"identifier,omitempty"`
	Version                       *String                   `json:"version,omitempty"`
	VersionPrimitiveElement       *primitiveElement         `json:"_version,omitempty"`
	Name                          *String                   `json:"name,omitempty"`
	NamePrimitiveElement          *primitiveElement         `json:"_name,omitempty"`
	Status                        Code                      `json:"status,omitempty"`
	StatusPrimitiveElement        *primitiveElement         `json:"_status,omitempty"`
	Experimental                  *Boolean                  `json:"experimental,omitempty"`
	ExperimentalPrimitiveElement  *primitiveElement         `json:"_experimental,omitempty"`
	Date                          *DateTime                 `json:"date,omitempty"`
	DatePrimitiveElement          *primitiveElement         `json:"_date,omitempty"`
	Publisher                     *String                   `json:"publisher,omitempty"`
	PublisherPrimitiveElement     *primitiveElement         `json:"_publisher,omitempty"`
	Contact                       []ContactDetail           `json:"contact,omitempty"`
	UseContext                    []UsageContext            `json:"useContext,omitempty"`
	Jurisdiction                  []CodeableConcept         `json:"jurisdiction,omitempty"`
	Copyright                     *Markdown                 `json:"copyright,omitempty"`
	CopyrightPrimitiveElement     *primitiveElement         `json:"_copyright,omitempty"`
	Purpose                       *Markdown                 `json:"purpose,omitempty"`
	PurposePrimitiveElement       *primitiveElement         `json:"_purpose,omitempty"`
	Actor                         []ExampleScenarioActor    `json:"actor,omitempty"`
	Instance                      []ExampleScenarioInstance `json:"instance,omitempty"`
	Process                       []ExampleScenarioProcess  `json:"process,omitempty"`
	Workflow                      []Canonical               `json:"workflow,omitempty"`
	WorkflowPrimitiveElement      []*primitiveElement       `json:"_workflow,omitempty"`
}

func (r ExampleScenario) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExampleScenario) marshalJSON() jsonExampleScenario {
	m := jsonExampleScenario{}
	m.ResourceType = "ExampleScenario"
	if r.Id != nil && r.Id.Value != nil {
		m.Id = r.Id
	}
	if r.Id != nil && (r.Id.Id != nil || r.Id.Extension != nil) {
		m.IdPrimitiveElement = &primitiveElement{Id: r.Id.Id, Extension: r.Id.Extension}
	}
	m.Meta = r.Meta
	if r.ImplicitRules != nil && r.ImplicitRules.Value != nil {
		m.ImplicitRules = r.ImplicitRules
	}
	if r.ImplicitRules != nil && (r.ImplicitRules.Id != nil || r.ImplicitRules.Extension != nil) {
		m.ImplicitRulesPrimitiveElement = &primitiveElement{Id: r.ImplicitRules.Id, Extension: r.ImplicitRules.Extension}
	}
	if r.Language != nil && r.Language.Value != nil {
		m.Language = r.Language
	}
	if r.Language != nil && (r.Language.Id != nil || r.Language.Extension != nil) {
		m.LanguagePrimitiveElement = &primitiveElement{Id: r.Language.Id, Extension: r.Language.Extension}
	}
	m.Text = r.Text
	m.Contained = make([]ContainedResource, 0, len(r.Contained))
	for _, c := range r.Contained {
		m.Contained = append(m.Contained, ContainedResource{c})
	}
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Url != nil && r.Url.Value != nil {
		m.Url = r.Url
	}
	if r.Url != nil && (r.Url.Id != nil || r.Url.Extension != nil) {
		m.UrlPrimitiveElement = &primitiveElement{Id: r.Url.Id, Extension: r.Url.Extension}
	}
	m.Identifier = r.Identifier
	if r.Version != nil && r.Version.Value != nil {
		m.Version = r.Version
	}
	if r.Version != nil && (r.Version.Id != nil || r.Version.Extension != nil) {
		m.VersionPrimitiveElement = &primitiveElement{Id: r.Version.Id, Extension: r.Version.Extension}
	}
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Status.Value != nil {
		m.Status = r.Status
	}
	if r.Status.Id != nil || r.Status.Extension != nil {
		m.StatusPrimitiveElement = &primitiveElement{Id: r.Status.Id, Extension: r.Status.Extension}
	}
	if r.Experimental != nil && r.Experimental.Value != nil {
		m.Experimental = r.Experimental
	}
	if r.Experimental != nil && (r.Experimental.Id != nil || r.Experimental.Extension != nil) {
		m.ExperimentalPrimitiveElement = &primitiveElement{Id: r.Experimental.Id, Extension: r.Experimental.Extension}
	}
	if r.Date != nil && r.Date.Value != nil {
		m.Date = r.Date
	}
	if r.Date != nil && (r.Date.Id != nil || r.Date.Extension != nil) {
		m.DatePrimitiveElement = &primitiveElement{Id: r.Date.Id, Extension: r.Date.Extension}
	}
	if r.Publisher != nil && r.Publisher.Value != nil {
		m.Publisher = r.Publisher
	}
	if r.Publisher != nil && (r.Publisher.Id != nil || r.Publisher.Extension != nil) {
		m.PublisherPrimitiveElement = &primitiveElement{Id: r.Publisher.Id, Extension: r.Publisher.Extension}
	}
	m.Contact = r.Contact
	m.UseContext = r.UseContext
	m.Jurisdiction = r.Jurisdiction
	if r.Copyright != nil && r.Copyright.Value != nil {
		m.Copyright = r.Copyright
	}
	if r.Copyright != nil && (r.Copyright.Id != nil || r.Copyright.Extension != nil) {
		m.CopyrightPrimitiveElement = &primitiveElement{Id: r.Copyright.Id, Extension: r.Copyright.Extension}
	}
	if r.Purpose != nil && r.Purpose.Value != nil {
		m.Purpose = r.Purpose
	}
	if r.Purpose != nil && (r.Purpose.Id != nil || r.Purpose.Extension != nil) {
		m.PurposePrimitiveElement = &primitiveElement{Id: r.Purpose.Id, Extension: r.Purpose.Extension}
	}
	m.Actor = r.Actor
	m.Instance = r.Instance
	m.Process = r.Process
	anyWorkflowValue := false
	for _, e := range r.Workflow {
		if e.Value != nil {
			anyWorkflowValue = true
			break
		}
	}
	if anyWorkflowValue {
		m.Workflow = r.Workflow
	}
	anyWorkflowIdOrExtension := false
	for _, e := range r.Workflow {
		if e.Id != nil || e.Extension != nil {
			anyWorkflowIdOrExtension = true
			break
		}
	}
	if anyWorkflowIdOrExtension {
		m.WorkflowPrimitiveElement = make([]*primitiveElement, 0, len(r.Workflow))
		for _, e := range r.Workflow {
			if e.Id != nil || e.Extension != nil {
				m.WorkflowPrimitiveElement = append(m.WorkflowPrimitiveElement, &primitiveElement{Id: e.Id, Extension: e.Extension})
			} else {
				m.WorkflowPrimitiveElement = append(m.WorkflowPrimitiveElement, nil)
			}
		}
	}
	return m
}
func (r *ExampleScenario) UnmarshalJSON(b []byte) error {
	var m jsonExampleScenario
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExampleScenario) unmarshalJSON(m jsonExampleScenario) error {
	r.Id = m.Id
	if m.IdPrimitiveElement != nil {
		if r.Id == nil {
			r.Id = &Id{}
		}
		r.Id.Id = m.IdPrimitiveElement.Id
		r.Id.Extension = m.IdPrimitiveElement.Extension
	}
	r.Meta = m.Meta
	r.ImplicitRules = m.ImplicitRules
	if m.ImplicitRulesPrimitiveElement != nil {
		if r.ImplicitRules == nil {
			r.ImplicitRules = &Uri{}
		}
		r.ImplicitRules.Id = m.ImplicitRulesPrimitiveElement.Id
		r.ImplicitRules.Extension = m.ImplicitRulesPrimitiveElement.Extension
	}
	r.Language = m.Language
	if m.LanguagePrimitiveElement != nil {
		if r.Language == nil {
			r.Language = &Code{}
		}
		r.Language.Id = m.LanguagePrimitiveElement.Id
		r.Language.Extension = m.LanguagePrimitiveElement.Extension
	}
	r.Text = m.Text
	r.Contained = make([]model.Resource, 0, len(m.Contained))
	for _, v := range m.Contained {
		r.Contained = append(r.Contained, v.Resource)
	}
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Url = m.Url
	if m.UrlPrimitiveElement != nil {
		if r.Url == nil {
			r.Url = &Uri{}
		}
		r.Url.Id = m.UrlPrimitiveElement.Id
		r.Url.Extension = m.UrlPrimitiveElement.Extension
	}
	r.Identifier = m.Identifier
	r.Version = m.Version
	if m.VersionPrimitiveElement != nil {
		if r.Version == nil {
			r.Version = &String{}
		}
		r.Version.Id = m.VersionPrimitiveElement.Id
		r.Version.Extension = m.VersionPrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Status = m.Status
	if m.StatusPrimitiveElement != nil {
		r.Status.Id = m.StatusPrimitiveElement.Id
		r.Status.Extension = m.StatusPrimitiveElement.Extension
	}
	r.Experimental = m.Experimental
	if m.ExperimentalPrimitiveElement != nil {
		if r.Experimental == nil {
			r.Experimental = &Boolean{}
		}
		r.Experimental.Id = m.ExperimentalPrimitiveElement.Id
		r.Experimental.Extension = m.ExperimentalPrimitiveElement.Extension
	}
	r.Date = m.Date
	if m.DatePrimitiveElement != nil {
		if r.Date == nil {
			r.Date = &DateTime{}
		}
		r.Date.Id = m.DatePrimitiveElement.Id
		r.Date.Extension = m.DatePrimitiveElement.Extension
	}
	r.Publisher = m.Publisher
	if m.PublisherPrimitiveElement != nil {
		if r.Publisher == nil {
			r.Publisher = &String{}
		}
		r.Publisher.Id = m.PublisherPrimitiveElement.Id
		r.Publisher.Extension = m.PublisherPrimitiveElement.Extension
	}
	r.Contact = m.Contact
	r.UseContext = m.UseContext
	r.Jurisdiction = m.Jurisdiction
	r.Copyright = m.Copyright
	if m.CopyrightPrimitiveElement != nil {
		if r.Copyright == nil {
			r.Copyright = &Markdown{}
		}
		r.Copyright.Id = m.CopyrightPrimitiveElement.Id
		r.Copyright.Extension = m.CopyrightPrimitiveElement.Extension
	}
	r.Purpose = m.Purpose
	if m.PurposePrimitiveElement != nil {
		if r.Purpose == nil {
			r.Purpose = &Markdown{}
		}
		r.Purpose.Id = m.PurposePrimitiveElement.Id
		r.Purpose.Extension = m.PurposePrimitiveElement.Extension
	}
	r.Actor = m.Actor
	r.Instance = m.Instance
	r.Process = m.Process
	r.Workflow = m.Workflow
	for i, e := range m.WorkflowPrimitiveElement {
		if len(r.Workflow) <= i {
			r.Workflow = append(r.Workflow, Canonical{})
		}
		if e != nil {
			r.Workflow[i].Id = e.Id
			r.Workflow[i].Extension = e.Extension
		}
	}
	return nil
}
func (r ExampleScenario) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Actor participating in the resource.
type ExampleScenarioActor struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// ID or acronym of actor.
	ActorId String
	// The type of actor - person or system.
	Type Code
	// The name of the actor as shown in the page.
	Name *String
	// The description of the actor.
	Description *Markdown
}
type jsonExampleScenarioActor struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	ActorId                     String            `json:"actorId,omitempty"`
	ActorIdPrimitiveElement     *primitiveElement `json:"_actorId,omitempty"`
	Type                        Code              `json:"type,omitempty"`
	TypePrimitiveElement        *primitiveElement `json:"_type,omitempty"`
	Name                        *String           `json:"name,omitempty"`
	NamePrimitiveElement        *primitiveElement `json:"_name,omitempty"`
	Description                 *Markdown         `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
}

func (r ExampleScenarioActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExampleScenarioActor) marshalJSON() jsonExampleScenarioActor {
	m := jsonExampleScenarioActor{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.ActorId.Value != nil {
		m.ActorId = r.ActorId
	}
	if r.ActorId.Id != nil || r.ActorId.Extension != nil {
		m.ActorIdPrimitiveElement = &primitiveElement{Id: r.ActorId.Id, Extension: r.ActorId.Extension}
	}
	if r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type.Id != nil || r.Type.Extension != nil {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	return m
}
func (r *ExampleScenarioActor) UnmarshalJSON(b []byte) error {
	var m jsonExampleScenarioActor
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExampleScenarioActor) unmarshalJSON(m jsonExampleScenarioActor) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ActorId = m.ActorId
	if m.ActorIdPrimitiveElement != nil {
		r.ActorId.Id = m.ActorIdPrimitiveElement.Id
		r.ActorId.Extension = m.ActorIdPrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &Markdown{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	return nil
}
func (r ExampleScenarioActor) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Each resource and each version that is present in the workflow.
type ExampleScenarioInstance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The id of the resource for referencing.
	ResourceId String
	// The type of the resource.
	ResourceType Code
	// A short name for the resource instance.
	Name *String
	// Human-friendly description of the resource instance.
	Description *Markdown
	// A specific version of the resource.
	Version []ExampleScenarioInstanceVersion
	// Resources contained in the instance (e.g. the observations contained in a bundle).
	ContainedInstance []ExampleScenarioInstanceContainedInstance
}
type jsonExampleScenarioInstance struct {
	Id                           *string                                    `json:"id,omitempty"`
	Extension                    []Extension                                `json:"extension,omitempty"`
	ModifierExtension            []Extension                                `json:"modifierExtension,omitempty"`
	ResourceId                   String                                     `json:"resourceId,omitempty"`
	ResourceIdPrimitiveElement   *primitiveElement                          `json:"_resourceId,omitempty"`
	ResourceType                 Code                                       `json:"resourceType,omitempty"`
	ResourceTypePrimitiveElement *primitiveElement                          `json:"_resourceType,omitempty"`
	Name                         *String                                    `json:"name,omitempty"`
	NamePrimitiveElement         *primitiveElement                          `json:"_name,omitempty"`
	Description                  *Markdown                                  `json:"description,omitempty"`
	DescriptionPrimitiveElement  *primitiveElement                          `json:"_description,omitempty"`
	Version                      []ExampleScenarioInstanceVersion           `json:"version,omitempty"`
	ContainedInstance            []ExampleScenarioInstanceContainedInstance `json:"containedInstance,omitempty"`
}

func (r ExampleScenarioInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExampleScenarioInstance) marshalJSON() jsonExampleScenarioInstance {
	m := jsonExampleScenarioInstance{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.ResourceId.Value != nil {
		m.ResourceId = r.ResourceId
	}
	if r.ResourceId.Id != nil || r.ResourceId.Extension != nil {
		m.ResourceIdPrimitiveElement = &primitiveElement{Id: r.ResourceId.Id, Extension: r.ResourceId.Extension}
	}
	if r.ResourceType.Value != nil {
		m.ResourceType = r.ResourceType
	}
	if r.ResourceType.Id != nil || r.ResourceType.Extension != nil {
		m.ResourceTypePrimitiveElement = &primitiveElement{Id: r.ResourceType.Id, Extension: r.ResourceType.Extension}
	}
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Version = r.Version
	m.ContainedInstance = r.ContainedInstance
	return m
}
func (r *ExampleScenarioInstance) UnmarshalJSON(b []byte) error {
	var m jsonExampleScenarioInstance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExampleScenarioInstance) unmarshalJSON(m jsonExampleScenarioInstance) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ResourceId = m.ResourceId
	if m.ResourceIdPrimitiveElement != nil {
		r.ResourceId.Id = m.ResourceIdPrimitiveElement.Id
		r.ResourceId.Extension = m.ResourceIdPrimitiveElement.Extension
	}
	r.ResourceType = m.ResourceType
	if m.ResourceTypePrimitiveElement != nil {
		r.ResourceType.Id = m.ResourceTypePrimitiveElement.Id
		r.ResourceType.Extension = m.ResourceTypePrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &Markdown{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Version = m.Version
	r.ContainedInstance = m.ContainedInstance
	return nil
}
func (r ExampleScenarioInstance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A specific version of the resource.
type ExampleScenarioInstanceVersion struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The identifier of a specific version of a resource.
	VersionId String
	// The description of the resource version.
	Description Markdown
}
type jsonExampleScenarioInstanceVersion struct {
	Id                          *string           `json:"id,omitempty"`
	Extension                   []Extension       `json:"extension,omitempty"`
	ModifierExtension           []Extension       `json:"modifierExtension,omitempty"`
	VersionId                   String            `json:"versionId,omitempty"`
	VersionIdPrimitiveElement   *primitiveElement `json:"_versionId,omitempty"`
	Description                 Markdown          `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement `json:"_description,omitempty"`
}

func (r ExampleScenarioInstanceVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExampleScenarioInstanceVersion) marshalJSON() jsonExampleScenarioInstanceVersion {
	m := jsonExampleScenarioInstanceVersion{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.VersionId.Value != nil {
		m.VersionId = r.VersionId
	}
	if r.VersionId.Id != nil || r.VersionId.Extension != nil {
		m.VersionIdPrimitiveElement = &primitiveElement{Id: r.VersionId.Id, Extension: r.VersionId.Extension}
	}
	if r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description.Id != nil || r.Description.Extension != nil {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	return m
}
func (r *ExampleScenarioInstanceVersion) UnmarshalJSON(b []byte) error {
	var m jsonExampleScenarioInstanceVersion
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExampleScenarioInstanceVersion) unmarshalJSON(m jsonExampleScenarioInstanceVersion) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.VersionId = m.VersionId
	if m.VersionIdPrimitiveElement != nil {
		r.VersionId.Id = m.VersionIdPrimitiveElement.Id
		r.VersionId.Extension = m.VersionIdPrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	return nil
}
func (r ExampleScenarioInstanceVersion) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Resources contained in the instance (e.g. the observations contained in a bundle).
type ExampleScenarioInstanceContainedInstance struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Each resource contained in the instance.
	ResourceId String
	// A specific version of a resource contained in the instance.
	VersionId *String
}
type jsonExampleScenarioInstanceContainedInstance struct {
	Id                         *string           `json:"id,omitempty"`
	Extension                  []Extension       `json:"extension,omitempty"`
	ModifierExtension          []Extension       `json:"modifierExtension,omitempty"`
	ResourceId                 String            `json:"resourceId,omitempty"`
	ResourceIdPrimitiveElement *primitiveElement `json:"_resourceId,omitempty"`
	VersionId                  *String           `json:"versionId,omitempty"`
	VersionIdPrimitiveElement  *primitiveElement `json:"_versionId,omitempty"`
}

func (r ExampleScenarioInstanceContainedInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExampleScenarioInstanceContainedInstance) marshalJSON() jsonExampleScenarioInstanceContainedInstance {
	m := jsonExampleScenarioInstanceContainedInstance{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.ResourceId.Value != nil {
		m.ResourceId = r.ResourceId
	}
	if r.ResourceId.Id != nil || r.ResourceId.Extension != nil {
		m.ResourceIdPrimitiveElement = &primitiveElement{Id: r.ResourceId.Id, Extension: r.ResourceId.Extension}
	}
	if r.VersionId != nil && r.VersionId.Value != nil {
		m.VersionId = r.VersionId
	}
	if r.VersionId != nil && (r.VersionId.Id != nil || r.VersionId.Extension != nil) {
		m.VersionIdPrimitiveElement = &primitiveElement{Id: r.VersionId.Id, Extension: r.VersionId.Extension}
	}
	return m
}
func (r *ExampleScenarioInstanceContainedInstance) UnmarshalJSON(b []byte) error {
	var m jsonExampleScenarioInstanceContainedInstance
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExampleScenarioInstanceContainedInstance) unmarshalJSON(m jsonExampleScenarioInstanceContainedInstance) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.ResourceId = m.ResourceId
	if m.ResourceIdPrimitiveElement != nil {
		r.ResourceId.Id = m.ResourceIdPrimitiveElement.Id
		r.ResourceId.Extension = m.ResourceIdPrimitiveElement.Extension
	}
	r.VersionId = m.VersionId
	if m.VersionIdPrimitiveElement != nil {
		if r.VersionId == nil {
			r.VersionId = &String{}
		}
		r.VersionId.Id = m.VersionIdPrimitiveElement.Id
		r.VersionId.Extension = m.VersionIdPrimitiveElement.Extension
	}
	return nil
}
func (r ExampleScenarioInstanceContainedInstance) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Each major process - a group of operations.
type ExampleScenarioProcess struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The diagram title of the group of operations.
	Title String
	// A longer description of the group of operations.
	Description *Markdown
	// Description of initial status before the process starts.
	PreConditions *Markdown
	// Description of final status after the process ends.
	PostConditions *Markdown
	// Each step of the process.
	Step []ExampleScenarioProcessStep
}
type jsonExampleScenarioProcess struct {
	Id                             *string                      `json:"id,omitempty"`
	Extension                      []Extension                  `json:"extension,omitempty"`
	ModifierExtension              []Extension                  `json:"modifierExtension,omitempty"`
	Title                          String                       `json:"title,omitempty"`
	TitlePrimitiveElement          *primitiveElement            `json:"_title,omitempty"`
	Description                    *Markdown                    `json:"description,omitempty"`
	DescriptionPrimitiveElement    *primitiveElement            `json:"_description,omitempty"`
	PreConditions                  *Markdown                    `json:"preConditions,omitempty"`
	PreConditionsPrimitiveElement  *primitiveElement            `json:"_preConditions,omitempty"`
	PostConditions                 *Markdown                    `json:"postConditions,omitempty"`
	PostConditionsPrimitiveElement *primitiveElement            `json:"_postConditions,omitempty"`
	Step                           []ExampleScenarioProcessStep `json:"step,omitempty"`
}

func (r ExampleScenarioProcess) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExampleScenarioProcess) marshalJSON() jsonExampleScenarioProcess {
	m := jsonExampleScenarioProcess{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Title.Value != nil {
		m.Title = r.Title
	}
	if r.Title.Id != nil || r.Title.Extension != nil {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	if r.PreConditions != nil && r.PreConditions.Value != nil {
		m.PreConditions = r.PreConditions
	}
	if r.PreConditions != nil && (r.PreConditions.Id != nil || r.PreConditions.Extension != nil) {
		m.PreConditionsPrimitiveElement = &primitiveElement{Id: r.PreConditions.Id, Extension: r.PreConditions.Extension}
	}
	if r.PostConditions != nil && r.PostConditions.Value != nil {
		m.PostConditions = r.PostConditions
	}
	if r.PostConditions != nil && (r.PostConditions.Id != nil || r.PostConditions.Extension != nil) {
		m.PostConditionsPrimitiveElement = &primitiveElement{Id: r.PostConditions.Id, Extension: r.PostConditions.Extension}
	}
	m.Step = r.Step
	return m
}
func (r *ExampleScenarioProcess) UnmarshalJSON(b []byte) error {
	var m jsonExampleScenarioProcess
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExampleScenarioProcess) unmarshalJSON(m jsonExampleScenarioProcess) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &Markdown{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.PreConditions = m.PreConditions
	if m.PreConditionsPrimitiveElement != nil {
		if r.PreConditions == nil {
			r.PreConditions = &Markdown{}
		}
		r.PreConditions.Id = m.PreConditionsPrimitiveElement.Id
		r.PreConditions.Extension = m.PreConditionsPrimitiveElement.Extension
	}
	r.PostConditions = m.PostConditions
	if m.PostConditionsPrimitiveElement != nil {
		if r.PostConditions == nil {
			r.PostConditions = &Markdown{}
		}
		r.PostConditions.Id = m.PostConditionsPrimitiveElement.Id
		r.PostConditions.Extension = m.PostConditionsPrimitiveElement.Extension
	}
	r.Step = m.Step
	return nil
}
func (r ExampleScenarioProcess) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Each step of the process.
type ExampleScenarioProcessStep struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// Nested process.
	Process []ExampleScenarioProcess
	// If there is a pause in the flow.
	Pause *Boolean
	// Each interaction or action.
	Operation *ExampleScenarioProcessStepOperation
	// Indicates an alternative step that can be taken instead of the operations on the base step in exceptional/atypical circumstances.
	Alternative []ExampleScenarioProcessStepAlternative
}
type jsonExampleScenarioProcessStep struct {
	Id                    *string                                 `json:"id,omitempty"`
	Extension             []Extension                             `json:"extension,omitempty"`
	ModifierExtension     []Extension                             `json:"modifierExtension,omitempty"`
	Process               []ExampleScenarioProcess                `json:"process,omitempty"`
	Pause                 *Boolean                                `json:"pause,omitempty"`
	PausePrimitiveElement *primitiveElement                       `json:"_pause,omitempty"`
	Operation             *ExampleScenarioProcessStepOperation    `json:"operation,omitempty"`
	Alternative           []ExampleScenarioProcessStepAlternative `json:"alternative,omitempty"`
}

func (r ExampleScenarioProcessStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExampleScenarioProcessStep) marshalJSON() jsonExampleScenarioProcessStep {
	m := jsonExampleScenarioProcessStep{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	m.Process = r.Process
	if r.Pause != nil && r.Pause.Value != nil {
		m.Pause = r.Pause
	}
	if r.Pause != nil && (r.Pause.Id != nil || r.Pause.Extension != nil) {
		m.PausePrimitiveElement = &primitiveElement{Id: r.Pause.Id, Extension: r.Pause.Extension}
	}
	m.Operation = r.Operation
	m.Alternative = r.Alternative
	return m
}
func (r *ExampleScenarioProcessStep) UnmarshalJSON(b []byte) error {
	var m jsonExampleScenarioProcessStep
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExampleScenarioProcessStep) unmarshalJSON(m jsonExampleScenarioProcessStep) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Process = m.Process
	r.Pause = m.Pause
	if m.PausePrimitiveElement != nil {
		if r.Pause == nil {
			r.Pause = &Boolean{}
		}
		r.Pause.Id = m.PausePrimitiveElement.Id
		r.Pause.Extension = m.PausePrimitiveElement.Extension
	}
	r.Operation = m.Operation
	r.Alternative = m.Alternative
	return nil
}
func (r ExampleScenarioProcessStep) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Each interaction or action.
type ExampleScenarioProcessStepOperation struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The sequential number of the interaction, e.g. 1.2.5.
	Number String
	// The type of operation - CRUD.
	Type *String
	// The human-friendly name of the interaction.
	Name *String
	// Who starts the transaction.
	Initiator *String
	// Who receives the transaction.
	Receiver *String
	// A comment to be inserted in the diagram.
	Description *Markdown
	// Whether the initiator is deactivated right after the transaction.
	InitiatorActive *Boolean
	// Whether the receiver is deactivated right after the transaction.
	ReceiverActive *Boolean
	// Each resource instance used by the initiator.
	Request *ExampleScenarioInstanceContainedInstance
	// Each resource instance used by the responder.
	Response *ExampleScenarioInstanceContainedInstance
}
type jsonExampleScenarioProcessStepOperation struct {
	Id                              *string                                   `json:"id,omitempty"`
	Extension                       []Extension                               `json:"extension,omitempty"`
	ModifierExtension               []Extension                               `json:"modifierExtension,omitempty"`
	Number                          String                                    `json:"number,omitempty"`
	NumberPrimitiveElement          *primitiveElement                         `json:"_number,omitempty"`
	Type                            *String                                   `json:"type,omitempty"`
	TypePrimitiveElement            *primitiveElement                         `json:"_type,omitempty"`
	Name                            *String                                   `json:"name,omitempty"`
	NamePrimitiveElement            *primitiveElement                         `json:"_name,omitempty"`
	Initiator                       *String                                   `json:"initiator,omitempty"`
	InitiatorPrimitiveElement       *primitiveElement                         `json:"_initiator,omitempty"`
	Receiver                        *String                                   `json:"receiver,omitempty"`
	ReceiverPrimitiveElement        *primitiveElement                         `json:"_receiver,omitempty"`
	Description                     *Markdown                                 `json:"description,omitempty"`
	DescriptionPrimitiveElement     *primitiveElement                         `json:"_description,omitempty"`
	InitiatorActive                 *Boolean                                  `json:"initiatorActive,omitempty"`
	InitiatorActivePrimitiveElement *primitiveElement                         `json:"_initiatorActive,omitempty"`
	ReceiverActive                  *Boolean                                  `json:"receiverActive,omitempty"`
	ReceiverActivePrimitiveElement  *primitiveElement                         `json:"_receiverActive,omitempty"`
	Request                         *ExampleScenarioInstanceContainedInstance `json:"request,omitempty"`
	Response                        *ExampleScenarioInstanceContainedInstance `json:"response,omitempty"`
}

func (r ExampleScenarioProcessStepOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExampleScenarioProcessStepOperation) marshalJSON() jsonExampleScenarioProcessStepOperation {
	m := jsonExampleScenarioProcessStepOperation{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Number.Value != nil {
		m.Number = r.Number
	}
	if r.Number.Id != nil || r.Number.Extension != nil {
		m.NumberPrimitiveElement = &primitiveElement{Id: r.Number.Id, Extension: r.Number.Extension}
	}
	if r.Type != nil && r.Type.Value != nil {
		m.Type = r.Type
	}
	if r.Type != nil && (r.Type.Id != nil || r.Type.Extension != nil) {
		m.TypePrimitiveElement = &primitiveElement{Id: r.Type.Id, Extension: r.Type.Extension}
	}
	if r.Name != nil && r.Name.Value != nil {
		m.Name = r.Name
	}
	if r.Name != nil && (r.Name.Id != nil || r.Name.Extension != nil) {
		m.NamePrimitiveElement = &primitiveElement{Id: r.Name.Id, Extension: r.Name.Extension}
	}
	if r.Initiator != nil && r.Initiator.Value != nil {
		m.Initiator = r.Initiator
	}
	if r.Initiator != nil && (r.Initiator.Id != nil || r.Initiator.Extension != nil) {
		m.InitiatorPrimitiveElement = &primitiveElement{Id: r.Initiator.Id, Extension: r.Initiator.Extension}
	}
	if r.Receiver != nil && r.Receiver.Value != nil {
		m.Receiver = r.Receiver
	}
	if r.Receiver != nil && (r.Receiver.Id != nil || r.Receiver.Extension != nil) {
		m.ReceiverPrimitiveElement = &primitiveElement{Id: r.Receiver.Id, Extension: r.Receiver.Extension}
	}
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	if r.InitiatorActive != nil && r.InitiatorActive.Value != nil {
		m.InitiatorActive = r.InitiatorActive
	}
	if r.InitiatorActive != nil && (r.InitiatorActive.Id != nil || r.InitiatorActive.Extension != nil) {
		m.InitiatorActivePrimitiveElement = &primitiveElement{Id: r.InitiatorActive.Id, Extension: r.InitiatorActive.Extension}
	}
	if r.ReceiverActive != nil && r.ReceiverActive.Value != nil {
		m.ReceiverActive = r.ReceiverActive
	}
	if r.ReceiverActive != nil && (r.ReceiverActive.Id != nil || r.ReceiverActive.Extension != nil) {
		m.ReceiverActivePrimitiveElement = &primitiveElement{Id: r.ReceiverActive.Id, Extension: r.ReceiverActive.Extension}
	}
	m.Request = r.Request
	m.Response = r.Response
	return m
}
func (r *ExampleScenarioProcessStepOperation) UnmarshalJSON(b []byte) error {
	var m jsonExampleScenarioProcessStepOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExampleScenarioProcessStepOperation) unmarshalJSON(m jsonExampleScenarioProcessStepOperation) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Number = m.Number
	if m.NumberPrimitiveElement != nil {
		r.Number.Id = m.NumberPrimitiveElement.Id
		r.Number.Extension = m.NumberPrimitiveElement.Extension
	}
	r.Type = m.Type
	if m.TypePrimitiveElement != nil {
		if r.Type == nil {
			r.Type = &String{}
		}
		r.Type.Id = m.TypePrimitiveElement.Id
		r.Type.Extension = m.TypePrimitiveElement.Extension
	}
	r.Name = m.Name
	if m.NamePrimitiveElement != nil {
		if r.Name == nil {
			r.Name = &String{}
		}
		r.Name.Id = m.NamePrimitiveElement.Id
		r.Name.Extension = m.NamePrimitiveElement.Extension
	}
	r.Initiator = m.Initiator
	if m.InitiatorPrimitiveElement != nil {
		if r.Initiator == nil {
			r.Initiator = &String{}
		}
		r.Initiator.Id = m.InitiatorPrimitiveElement.Id
		r.Initiator.Extension = m.InitiatorPrimitiveElement.Extension
	}
	r.Receiver = m.Receiver
	if m.ReceiverPrimitiveElement != nil {
		if r.Receiver == nil {
			r.Receiver = &String{}
		}
		r.Receiver.Id = m.ReceiverPrimitiveElement.Id
		r.Receiver.Extension = m.ReceiverPrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &Markdown{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.InitiatorActive = m.InitiatorActive
	if m.InitiatorActivePrimitiveElement != nil {
		if r.InitiatorActive == nil {
			r.InitiatorActive = &Boolean{}
		}
		r.InitiatorActive.Id = m.InitiatorActivePrimitiveElement.Id
		r.InitiatorActive.Extension = m.InitiatorActivePrimitiveElement.Extension
	}
	r.ReceiverActive = m.ReceiverActive
	if m.ReceiverActivePrimitiveElement != nil {
		if r.ReceiverActive == nil {
			r.ReceiverActive = &Boolean{}
		}
		r.ReceiverActive.Id = m.ReceiverActivePrimitiveElement.Id
		r.ReceiverActive.Extension = m.ReceiverActivePrimitiveElement.Extension
	}
	r.Request = m.Request
	r.Response = m.Response
	return nil
}
func (r ExampleScenarioProcessStepOperation) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates an alternative step that can be taken instead of the operations on the base step in exceptional/atypical circumstances.
type ExampleScenarioProcessStepAlternative struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// The label to display for the alternative that gives a sense of the circumstance in which the alternative should be invoked.
	Title String
	// A human-readable description of the alternative explaining when the alternative should occur rather than the base step.
	Description *Markdown
	// What happens in each alternative option.
	Step []ExampleScenarioProcessStep
}
type jsonExampleScenarioProcessStepAlternative struct {
	Id                          *string                      `json:"id,omitempty"`
	Extension                   []Extension                  `json:"extension,omitempty"`
	ModifierExtension           []Extension                  `json:"modifierExtension,omitempty"`
	Title                       String                       `json:"title,omitempty"`
	TitlePrimitiveElement       *primitiveElement            `json:"_title,omitempty"`
	Description                 *Markdown                    `json:"description,omitempty"`
	DescriptionPrimitiveElement *primitiveElement            `json:"_description,omitempty"`
	Step                        []ExampleScenarioProcessStep `json:"step,omitempty"`
}

func (r ExampleScenarioProcessStepAlternative) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.marshalJSON())
}
func (r ExampleScenarioProcessStepAlternative) marshalJSON() jsonExampleScenarioProcessStepAlternative {
	m := jsonExampleScenarioProcessStepAlternative{}
	m.Id = r.Id
	m.Extension = r.Extension
	m.ModifierExtension = r.ModifierExtension
	if r.Title.Value != nil {
		m.Title = r.Title
	}
	if r.Title.Id != nil || r.Title.Extension != nil {
		m.TitlePrimitiveElement = &primitiveElement{Id: r.Title.Id, Extension: r.Title.Extension}
	}
	if r.Description != nil && r.Description.Value != nil {
		m.Description = r.Description
	}
	if r.Description != nil && (r.Description.Id != nil || r.Description.Extension != nil) {
		m.DescriptionPrimitiveElement = &primitiveElement{Id: r.Description.Id, Extension: r.Description.Extension}
	}
	m.Step = r.Step
	return m
}
func (r *ExampleScenarioProcessStepAlternative) UnmarshalJSON(b []byte) error {
	var m jsonExampleScenarioProcessStepAlternative
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	return r.unmarshalJSON(m)
}
func (r *ExampleScenarioProcessStepAlternative) unmarshalJSON(m jsonExampleScenarioProcessStepAlternative) error {
	r.Id = m.Id
	r.Extension = m.Extension
	r.ModifierExtension = m.ModifierExtension
	r.Title = m.Title
	if m.TitlePrimitiveElement != nil {
		r.Title.Id = m.TitlePrimitiveElement.Id
		r.Title.Extension = m.TitlePrimitiveElement.Extension
	}
	r.Description = m.Description
	if m.DescriptionPrimitiveElement != nil {
		if r.Description == nil {
			r.Description = &Markdown{}
		}
		r.Description.Id = m.DescriptionPrimitiveElement.Id
		r.Description.Extension = m.DescriptionPrimitiveElement.Extension
	}
	r.Step = m.Step
	return nil
}
func (r ExampleScenarioProcessStepAlternative) String() string {
	buf, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
