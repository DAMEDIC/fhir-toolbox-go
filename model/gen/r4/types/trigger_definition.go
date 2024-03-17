package types

import (
	"encoding/json"
	r4 "fhir-toolbox/model/gen/r4"
)

// Base StructureDefinition for TriggerDefinition Type: A description of a triggering event. Triggering events can be named events, data events, or periodic, as determined by the type element.
type TriggerDefinition struct {
	// The triggering data of the event (if this is a data trigger). If more than one data is requirement is specified, then all the data requirements must be true.
	Data []DataRequirement
	// A boolean-valued expression that is evaluated in the context of the container of the trigger definition and returns whether or not the trigger fires.
	Condition *Expression
	// Unique id for the element within a resource (for internal references). This may be any string value that does not contain spaces.
	Id *string
	// May be used to represent additional information that is not part of the basic definition of the element. To make the use of extensions safe and manageable, there is a strict set of governance  applied to the definition and use of extensions. Though any implementer can define an extension, there is a set of requirements that SHALL be met as part of the definition of the extension.
	Extension []Extension
	// The type of triggering event.
	Type Code
	// A formal name for the event. This may be an absolute URI that identifies the event formally (e.g. from a trigger registry), or a simple relative URI that identifies the event in a local context.
	Name *String
	// The timing of the event (if this is a periodic trigger).
	Timing r4.Element
}

func (s TriggerDefinition) String() string {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(buf)
}
