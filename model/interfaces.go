package model

// Resource is any FHIR Resource.
type Resource interface {
	Element
	ResourceType() string
	ResourceId() (string, bool)
}

// Element is any element in the FHIR model.
//
// This includes Resources, Datatypes and BackboneElements.
type Element interface {
	MemSize() int
}
