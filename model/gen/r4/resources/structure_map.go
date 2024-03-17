package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// A name assigned to an instance of data. The instance must be provided when the mapping is invoked.
type StructureMapGroupInput struct {
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Name for this instance of data.
	Name types.Id
	// Type for this instance of data.
	Type *types.String
	// Mode for this instance of data.
	Mode types.Code
	// Documentation for this instance of data.
	Documentation *types.String
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
}

func (s StructureMapGroupInput) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Source inputs to the mapping.
type StructureMapGroupRuleSource struct {
	// Optional field for this source.
	Element *types.String
	// Named context for field, if a field is specified.
	Variable *types.Id
	// FHIRPath expression  - must be true or the mapping engine throws an error instead of completing.
	Check *types.String
	// Type or variable this rule applies to.
	Context types.Id
	// Specified maximum cardinality for the element - a number or a "*". This is optional; if present, it acts an implicit check on the input content (* just serves as documentation; it's the default value).
	Max *types.String
	// FHIRPath expression  - must be true or the rule does not apply.
	Condition *types.String
	// A FHIRPath expression which specifies a message to put in the transform log when content matching the source rule is found.
	LogMessage *types.String
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Specified type for the element. This works as a condition on the mapping - use for polymorphic elements.
	Type *types.String
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A value to use if there is no existing value in the source object.
	DefaultValue r4.Element
	// How to handle the list mode for this element.
	ListMode *types.Code
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Specified minimum cardinality for the element. This is optional; if present, it acts an implicit check on the input content.
	Min *types.Integer
}

func (s StructureMapGroupRuleSource) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Parameters to the transform.
type StructureMapGroupRuleTargetParameter struct {
	// Parameter value - variable or literal.
	Value r4.Element
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
}

func (s StructureMapGroupRuleTargetParameter) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Content to create because of this mapping rule.
type StructureMapGroupRuleTarget struct {
	// If field is a list, how to manage the list.
	ListMode []types.Code
	// Internal rule reference for shared list items.
	ListRuleId *types.Id
	// How the data is copied / created.
	Transform *types.Code
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Type or variable this rule applies to.
	Context *types.Id
	// How to interpret the context.
	ContextType *types.Code
	// Named context for field, if desired, and a field is specified.
	Variable *types.Id
	// Parameters to the transform.
	Parameter []StructureMapGroupRuleTargetParameter
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Field to create in the context.
	Element *types.String
}

func (s StructureMapGroupRuleTarget) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Which other rules to apply in the context of this rule.
type StructureMapGroupRuleDependent struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Name of a rule or group to apply.
	Name types.Id
	// Variable to pass to the rule or group.
	Variable []types.String
}

func (s StructureMapGroupRuleDependent) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Transform Rule from source to target.
type StructureMapGroupRule struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Rules contained in this rule.
	Rule []StructureMapGroupRule
	// Documentation for this instance of data.
	Documentation *types.String
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Name of the rule for internal references.
	Name types.Id
	// Source inputs to the mapping.
	Source []StructureMapGroupRuleSource
	// Content to create because of this mapping rule.
	Target []StructureMapGroupRuleTarget
	// Which other rules to apply in the context of this rule.
	Dependent []StructureMapGroupRuleDependent
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
}

func (s StructureMapGroupRule) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Organizes the mapping into manageable chunks for human review/ease of maintenance.
type StructureMapGroup struct {
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Another group that this group adds rules to.
	Extends *types.Id
	// If this is the default rule set to apply for the source type or this combination of types.
	TypeMode types.Code
	// A name assigned to an instance of data. The instance must be provided when the mapping is invoked.
	Input []StructureMapGroupInput
	// Transform Rule from source to target.
	Rule []StructureMapGroupRule
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// A unique name for the group for the convenience of human readers.
	Name types.Id
	// Additional supporting documentation that explains the purpose of the group and the types of mappings within it.
	Documentation *types.String
}

func (s StructureMapGroup) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A structure definition used by this map. The structure definition may describe instances that are converted, or the instances that are produced.
type StructureMapStructure struct {
	// The canonical reference to the structure.
	Url types.Canonical
	// How the referenced structure is used in this mapping.
	Mode types.Code
	// The name used for this type in the map.
	Alias *types.String
	// Documentation that describes how the structure is used in the mapping.
	Documentation *types.String
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
}

func (s StructureMapStructure) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A Map of relationships between 2 structures that can be used to transform data.
type StructureMap struct {
	// The base language in which the resource is written.
	Language *types.Code
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// An absolute URI that is used to identify this structure map when it is referenced in a specification, model, design or an instance; also called its canonical identifier. This SHOULD be globally unique and SHOULD be a literal address at which at which an authoritative instance of this structure map is (or will be) published. This URL can be the target of a canonical reference. It SHALL remain the same when the structure map is stored on different servers.
	Url types.Uri
	// A formal identifier that is used to identify this structure map when it is represented in other formats, or referenced in a specification, model, design or an instance.
	Identifier []types.Identifier
	// The date  (and optionally time) when the structure map was published. The date must change when the business version changes and it must change if the status code changes. In addition, it should change when the substantive content of the structure map changes.
	Date *types.DateTime
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// A copyright statement relating to the structure map and/or its contents. Copyright statements are generally legal restrictions on the use and publishing of the structure map.
	Copyright *types.Markdown
	// Other maps used by this map (canonical URLs).
	Import []types.Canonical
	// Contact details to assist a user in finding and communicating with the publisher.
	Contact []types.ContactDetail
	// A legal or geographic region in which the structure map is intended to be used.
	Jurisdiction []types.CodeableConcept
	// Organizes the mapping into manageable chunks for human review/ease of maintenance.
	Group []StructureMapGroup
	// A Boolean value to indicate that this structure map is authored for testing purposes (or education/evaluation/marketing) and is not intended to be used for genuine usage.
	Experimental *types.Boolean
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// A natural language name identifying the structure map. This name should be usable as an identifier for the module by machine processing applications such as code generation.
	Name types.String
	// The status of this structure map. Enables tracking the life-cycle of the content.
	Status types.Code
	// The name of the organization or individual that published the structure map.
	Publisher *types.String
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// The identifier that is used to identify this version of the structure map when it is referenced in a specification, model, design or instance. This is an arbitrary value managed by the structure map author and is not expected to be globally unique. For example, it might be a timestamp (e.g. yyyymmdd) if a managed version is not available. There is also no expectation that versions can be placed in a lexicographical sequence.
	Version *types.String
	// A short, descriptive, user-friendly title for the structure map.
	Title *types.String
	// A free text natural language description of the structure map from a consumer's perspective.
	Description *types.Markdown
	// The content was developed with a focus and intent of supporting the contexts that are listed. These contexts may be general categories (gender, age, ...) or may be references to specific programs (insurance plans, studies, ...) and may be used to assist with indexing and searching for appropriate structure map instances.
	UseContext []types.UsageContext
	// Explanation of why this structure map is needed and why it has been designed as it has.
	Purpose *types.Markdown
	// A structure definition used by this map. The structure definition may describe instances that are converted, or the instances that are produced.
	Structure []StructureMapStructure
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
}

func (s StructureMap) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
