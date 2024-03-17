package resources

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
	types "fhir-toolbox/model/gen/r4/types"
)

// Characteristics for quantitative results of this observation.
type ObservationDefinitionQuantitativeDetails struct {
	// Number of digits after decimal separator when the results of such observations are of type Quantity.
	DecimalPrecision *types.Integer
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Customary unit used to report quantitative results of observations conforming to this ObservationDefinition.
	CustomaryUnit *types.CodeableConcept
	// SI unit used to report quantitative results of observations conforming to this ObservationDefinition.
	Unit *types.CodeableConcept
	// Factor for converting value expressed with SI unit to value expressed with customary unit.
	ConversionFactor *types.Decimal
}

func (s ObservationDefinitionQuantitativeDetails) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Multiple  ranges of results qualified by different contexts for ordinal or continuous observations conforming to this ObservationDefinition.
type ObservationDefinitionQualifiedInterval struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// May be used to represent additional information that is not part of the basic definition of the element and that modifies the understanding of the element in which it is contained and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// Codes to indicate the health context the range applies to. For example, the normal or therapeutic range.
	Context *types.CodeableConcept
	// Sex of the population the range applies to.
	Gender *types.Code
	// Text based condition for which the reference range is valid.
	Condition *types.String
	// The category of interval of values for continuous or ordinal observations conforming to this ObservationDefinition.
	Category *types.Code
	// The low and high values determining the interval. There may be only one of the two.
	Range *types.Range
	// Codes to indicate the target population this reference range applies to.
	AppliesTo []types.CodeableConcept
	// The age at which this reference range is applicable. This is a neonatal age (e.g. number of weeks at term) if the meaning says so.
	Age *types.Range
	// The gestational age to which this reference range is applicable, in the context of pregnancy.
	GestationalAge *types.Range
}

func (s ObservationDefinitionQualifiedInterval) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

// Set of definitional characteristics for a kind of observation or measurement produced or consumed by an orderable health care service.
//
// In a catalog of health-related services that use or produce observations and measurements, this resource describes the expected characteristics of these observation / measurements.
type ObservationDefinition struct {
	// May be used to represent additional information that is not part of the basic definition of the resource and that modifies the understanding of the element that contains it and/or the understanding of the containing element's descendants. Usually modifier elements provide negation or qualification. To make the use of extensions safe and manageable, there is a strict set of governance applied to the definition and use of extensions. Though any implementer is allowed to define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension. Applications processing a resource are required to check for modifier extensions.
	//
	// Modifier extensions SHALL NOT change the meaning of any elements on Resource or DomainResource (including cannot change the meaning of modifierExtension itself).
	ModifierExtension []types.Extension
	// A unique identifier assigned to this ObservationDefinition artifact.
	Identifier []types.Identifier
	// The set of abnormal coded results for the observation conforming to this ObservationDefinition.
	AbnormalCodedValueSet *types.Reference
	// The logical id of the resource, as used in the URL for the resource. Once assigned, this value never changes.
	Id *types.Id
	// The base language in which the resource is written.
	Language *types.Code
	// May be used to represent additional information that is not part of the basic definition of the resource. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []types.Extension
	// Multiple results allowed for observations conforming to this ObservationDefinition.
	MultipleResultsAllowed *types.Boolean
	// The preferred name to be used when reporting the results of observations conforming to this ObservationDefinition.
	PreferredReportName *types.String
	// Characteristics for quantitative results of this observation.
	QuantitativeDetails *ObservationDefinitionQuantitativeDetails
	// The set of valid coded results for the observations  conforming to this ObservationDefinition.
	ValidCodedValueSet *types.Reference
	// The set of normal coded results for the observations conforming to this ObservationDefinition.
	NormalCodedValueSet *types.Reference
	// The metadata about the resource. This is content that is maintained by the infrastructure. Changes to the content might not always be associated with version changes to the resource.
	Meta *types.Meta
	// A reference to a set of rules that were followed when the resource was constructed, and which must be understood when processing the content. Often, this is a reference to an implementation guide that defines the special rules along with other profiles etc.
	ImplicitRules *types.Uri
	// A code that classifies the general type of observation.
	Category []types.CodeableConcept
	// Describes what will be observed. Sometimes this is called the observation "name".
	Code types.CodeableConcept
	// The data types allowed for the value element of the instance observations conforming to this ObservationDefinition.
	PermittedDataType []types.Code
	// The set of critical coded results for the observation conforming to this ObservationDefinition.
	CriticalCodedValueSet *types.Reference
	// A human-readable narrative that contains a summary of the resource and can be used to represent the content of the resource to a human. The narrative need not encode all the structured data, but is required to contain sufficient detail to make it "clinically safe" for a human to just read the narrative. Resource definitions may define what content should be represented in the narrative to ensure clinical safety.
	Text *types.Narrative
	// These resources do not have an independent existence apart from the resource that contains them - they cannot be identified independently, and nor can they have their own independent transaction scope.
	Contained []r4.Resource
	// The method or technique used to perform the observation.
	Method *types.CodeableConcept
	// Multiple  ranges of results qualified by different contexts for ordinal or continuous observations conforming to this ObservationDefinition.
	QualifiedInterval []ObservationDefinitionQualifiedInterval
}

func (s ObservationDefinition) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
