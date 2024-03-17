package types

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
)

// Base StructureDefinition for UsageContext Type: Specifies clinical/business/etc. metadata that can be used to retrieve, index and/or categorize an artifact. This metadata can either be specific to the applicable population (e.g., age category, DRG) or the specific context of care (e.g., venue, care setting, provider of care).
//
// Consumers of the resource must be able to determine the intended applicability for the resource. Ideally, this information would be used programmatically to determine when and how it should be incorporated or exposed.
type UsageContext struct {
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// A code that identifies the type of context being specified by this usage context.
	Code Coding
	// A value that defines the context specified in this context of use. The interpretation of the value is defined by the code.
	Value r4.Element
}

func (s UsageContext) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
