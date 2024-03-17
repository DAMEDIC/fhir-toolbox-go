package types

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
)

// Information about the base definition of the element, provided to make it unnecessary for tools to trace the deviation of the element through the derived and related profiles. When the element definition is not the original definition of an element - i.g. either in a constraint on another type, or for elements from a super type in a snap shot - then the information in provided in the element definition may be different to the base definition. On the original definition of the element, it will be same.
type ElementDefinitionBase struct {
	// The Path that identifies the base element - this matches the ElementDefinition.path for that element. Across FHIR, there is only one base definition of any element - that is, an element definition on a [StructureDefinition](structuredefinition.html#) without a StructureDefinition.base.
	Path String
	// Minimum cardinality of the base element identified by the path.
	Min UnsignedInt
	// Maximum cardinality of the base element identified by the path.
	Max String
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
}

func (s ElementDefinitionBase) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// The data type or resource that the value of this element is permitted to be.
type ElementDefinitionType struct {
	// URL of Data type or Resource that is a(or the) type used for this element. References are URLs that are relative to http://hl7.org/fhir/StructureDefinition e.g. "string" is a reference to http://hl7.org/fhir/StructureDefinition/string. Absolute URLs are only allowed in logical models.
	Code Uri
	// Identifies a profile structure or implementation Guide that applies to the datatype this element refers to. If any profiles are specified, then the content must conform to at least one of them. The URL can be a local reference - to a contained StructureDefinition, or a reference to another StructureDefinition or Implementation Guide by a canonical URL. When an implementation guide is specified, the type SHALL conform to at least one profile defined in the implementation guide.
	Profile []Canonical
	// Used when the type is "Reference" or "canonical", and identifies a profile structure or implementation Guide that applies to the target of the reference this element refers to. If any profiles are specified, then the content must conform to at least one of them. The URL can be a local reference - to a contained StructureDefinition, or a reference to another StructureDefinition or Implementation Guide by a canonical URL. When an implementation guide is specified, the target resource SHALL conform to at least one profile defined in the implementation guide.
	TargetProfile []Canonical
	// If the type is a reference to another resource, how the resource is or can be aggregated - is it a contained resource, or a reference, and if the context is a bundle, is it included in the bundle.
	Aggregation []Code
	// Whether this reference needs to be version specific or version independent, or whether either can be used.
	Versioning *Code
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
}

func (s ElementDefinitionType) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Formal constraints such as co-occurrence and other constraints that can be computationally evaluated within the context of the instance.
type ElementDefinitionConstraint struct {
	// Description of why this constraint is necessary or appropriate.
	Requirements *String
	// Identifies the impact constraint violation has on the conformance of the instance.
	Severity Code
	// Text that can be used to describe the constraint in messages identifying that the constraint has been violated.
	Human String
	// An XPath expression of constraint that can be executed to see if this constraint is met.
	Xpath *String
	// A reference to the original source of the constraint, for traceability purposes.
	Source *Canonical
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Allows identification of which elements have their cardinalities impacted by the constraint.  Will not be referenced for constraints that do not affect cardinality.
	Key Id
	// A [FHIRPath](fhirpath.html) expression of constraint that can be executed to see if this constraint is met.
	Expression *String
}

func (s ElementDefinitionConstraint) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Identifies a concept from an external specification that roughly corresponds to this element.
type ElementDefinitionMapping struct {
	// Comments that provide information about the mapping or its use.
	Comment *String
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// An internal reference to the definition of a mapping.
	Identity Id
	// Identifies the computable language in which mapping.map is expressed.
	Language *Code
	// Expresses what part of the target specification corresponds to this element.
	Map String
}

func (s ElementDefinitionMapping) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Designates which child elements are used to discriminate between the slices when processing an instance. If one or more discriminators are provided, the value of the child elements in the instance data SHALL completely distinguish which slice the element in the resource matches based on the allowed values for those elements in each of the slices.
type ElementDefinitionSlicingDiscriminator struct {
	// How the element value is interpreted when discrimination is evaluated.
	Type Code
	// A FHIRPath expression, using [the simple subset of FHIRPath](fhirpath.html#simple), that is used to identify the element on which discrimination is based.
	Path String
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
}

func (s ElementDefinitionSlicingDiscriminator) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Indicates that the element is sliced into a set of alternative definitions (i.e. in a structure definition, there are multiple different constraints on a single element in the base resource). Slicing can be used in any resource that has cardinality ..* on the base resource, or any resource with a choice of types. The set of slices is any elements that come after this in the element sequence that have the same path, until a shorter path occurs (the shorter path terminates the set).
type ElementDefinitionSlicing struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Designates which child elements are used to discriminate between the slices when processing an instance. If one or more discriminators are provided, the value of the child elements in the instance data SHALL completely distinguish which slice the element in the resource matches based on the allowed values for those elements in each of the slices.
	Discriminator []ElementDefinitionSlicingDiscriminator
	// A human-readable text description of how the slicing works. If there is no discriminator, this is required to be present to provide whatever information is possible about how the slices can be differentiated.
	Description *String
	// If the matching elements have to occur in the same order as defined in the profile.
	Ordered *Boolean
	// Whether additional slices are allowed or not. When the slices are ordered, profile authors can also say that additional slices are only allowed at the end.
	Rules Code
}

func (s ElementDefinitionSlicing) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Binds to a value set if this element is coded (code, Coding, CodeableConcept, Quantity), or the data types (string, uri).
type ElementDefinitionBinding struct {
	// Indicates the degree of conformance expectations associated with this binding - that is, the degree to which the provided value set must be adhered to in the instances.
	Strength Code
	// Describes the intended use of this particular set of codes.
	Description *String
	// Refers to the value set that identifies the set of codes the binding refers to.
	ValueSet *Canonical
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
}

func (s ElementDefinitionBinding) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// A sample value for this element demonstrating the type of information that would typically be found in the element.
type ElementDefinitionExample struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// Describes the purpose of this example amoung the set of examples.
	Label String
	// The actual value for the element, which must be one of the types allowed for this element.
	Value r4.Element
}

func (s ElementDefinitionExample) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Base StructureDefinition for ElementDefinition Type: Captures constraints on each element within the resource, profile, or extension.
type ElementDefinition struct {
	// If true, implementations that produce or consume resources SHALL provide "support" for the element in some meaningful way.  If false, the element may be ignored and not supported. If false, whether to populate or use the data element in any way is at the discretion of the implementation.
	MustSupport *Boolean
	// Provides a complete explanation of the meaning of the data element for human readability.  For the case of elements derived from existing elements (e.g. constraints), the definition SHALL be consistent with the base definition, but convey the meaning of the element in the particular context of use of the resource. (Note: The text you are reading is specified in ElementDefinition.definition).
	Definition *Markdown
	// The maximum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	MaxValue r4.Element
	// Indicates the maximum length in characters that is permitted to be present in conformant instances and which is expected to be supported by conformant consumers that support the element.
	MaxLength *Integer
	// Specifies a value that SHALL be exactly the value  for this element in the instance. For purposes of comparison, non-significant whitespace is ignored, and all values must be an exact match (case and accent sensitive). Missing elements/attributes must also be missing.
	Fixed r4.Element
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// Information about the base definition of the element, provided to make it unnecessary for tools to trace the deviation of the element through the derived and related profiles. When the element definition is not the original definition of an element - i.g. either in a constraint on another type, or for elements from a super type in a snap shot - then the information in provided in the element definition may be different to the base definition. On the original definition of the element, it will be same.
	Base *ElementDefinitionBase
	// Identifies an element defined elsewhere in the definition whose content rules should be applied to the current element. ContentReferences bring across all the rules that are in the ElementDefinition for the element, including definitions, cardinality constraints, bindings, invariants etc.
	ContentReference *Uri
	// If true, the value of this element affects the interpretation of the element or resource that contains it, and the value of the element cannot be ignored. Typically, this is used for status, negation and qualification codes. The effect of this is that the element cannot be ignored by systems: they SHALL either recognize the element and process it, and/or a pre-determination has been made that it is not relevant to their particular system.
	IsModifier *Boolean
	// A single preferred label which is the text to display beside the element indicating its meaning or to use to prompt for the element in a user display or form.
	Label *String
	// Specifies a value that the value in the instance SHALL follow - that is, any value in the pattern must be found in the instance. Other additional values may be found too. This is effectively constraint by example.
	//
	// When pattern[x] is used to constrain a primitive, it means that the value provided in the pattern[x] must match the instance value exactly.
	//
	// When pattern[x] is used to constrain an array, it means that each element provided in the pattern[x] array must (recursively) match at least one element from the instance array.
	//
	// When pattern[x] is used to constrain a complex object, it means that each property in the pattern must be present in the complex object, and its value must recursively match -- i.e.,
	//
	// 1. If primitive: it must match exactly the pattern value
	// 2. If a complex object: it must match (recursively) the pattern value
	// 3. If an array: it must match (recursively) the pattern value.
	Pattern r4.Element
	// The minimum allowed value for the element. The value is inclusive. This is allowed for the types date, dateTime, instant, time, decimal, integer, and Quantity.
	MinValue r4.Element
	// The minimum number of times this element SHALL appear in the instance.
	Min *UnsignedInt
	// The maximum number of times this element is permitted to appear in the instance.
	Max *String
	// The value that should be used if there is no value stated in the instance (e.g. 'if not otherwise specified, the abstract is false').
	DefaultValue r4.Element
	// Whether the element should be included if a client requests a search with the parameter _summary=true.
	IsSummary *Boolean
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []Extension
	// If true, indicates that this slice definition is constraining a slice definition with the same name in an inherited profile. If false, the slice is not overriding any slice in an inherited profile. If missing, the slice might or might not be overriding a slice in an inherited profile, depending on the sliceName.
	SliceIsConstraining *Boolean
	// A code that has the same meaning as the element in a particular terminology.
	Code []Coding
	// If present, indicates that the order of the repeating element has meaning and describes what that meaning is.  If absent, it means that the order of the element has no meaning.
	OrderMeaning *String
	// Explains how that element affects the interpretation of the resource or element that contains it.
	IsModifierReason *String
	// The name of this element definition slice, when slicing is working. The name must be a token with no dots or spaces. This is a unique name referring to a specific set of constraints applied to this element, used to provide a name to different slices of the same element.
	SliceName *String
	// This element is for traceability of why the element was created and why the constraints exist as they do. This may be used to point to source materials or specifications that drove the structure of this element.
	Requirements *Markdown
	// The data type or resource that the value of this element is permitted to be.
	Type []ElementDefinitionType
	// Identifies additional names by which this element might also be known.
	Alias []String
	// Formal constraints such as co-occurrence and other constraints that can be computationally evaluated within the context of the instance.
	Constraint []ElementDefinitionConstraint
	// Identifies a concept from an external specification that roughly corresponds to this element.
	Mapping []ElementDefinitionMapping
	// Indicates that the element is sliced into a set of alternative definitions (i.e. in a structure definition, there are multiple different constraints on a single element in the base resource). Slicing can be used in any resource that has cardinality ..* on the base resource, or any resource with a choice of types. The set of slices is any elements that come after this in the element sequence that have the same path, until a shorter path occurs (the shorter path terminates the set).
	Slicing *ElementDefinitionSlicing
	// A concise description of what this element means (e.g. for use in autogenerated summaries).
	Short *String
	// Explanatory notes and implementation guidance about the data element, including notes about how to use the data properly, exceptions to proper use, etc. (Note: The text you are reading is specified in ElementDefinition.comment).
	Comment *Markdown
	// Binds to a value set if this element is coded (code, Coding, CodeableConcept, Quantity), or the data types (string, uri).
	Binding *ElementDefinitionBinding
	// The path identifies the element and is expressed as a "."-separated list of ancestor elements, beginning with the name of the resource or extension.
	Path String
	// Codes that define how this element is represented in instances, when the deviation varies from the normal case.
	Representation []Code
	// A reference to an invariant that may make additional statements about the cardinality or value in the instance.
	Condition []Id
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The Implicit meaning that is to be understood when this element is missing (e.g. 'when this element is missing, the period is ongoing').
	MeaningWhenMissing *Markdown
	// A sample value for this element demonstrating the type of information that would typically be found in the element.
	Example []ElementDefinitionExample
}

func (s ElementDefinition) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
