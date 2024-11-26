// Package ir contains an intermediate representation that makes code generation
// easier than directly dealing with StructureDefinitions.
package ir

// ResourceOrType is a FHIR type that can either be a resource or some other data type.
type ResourceOrType struct {
	Name        string
	FileName    string
	IsResource  bool
	IsPrimitive bool
	Structs     []Struct
}

// Struct represent a resource, type or backbone element within another resource, type or backbone element.
type Struct struct {
	Name        string
	IsResource  bool
	IsPrimitive bool
	Fields      []StructField
	DocComment  string
}

// A StructField represent child elements of a resource, type or backbone element.
type StructField struct {
	Name          string
	MarshalName   string
	PossibleTypes []FieldType
	Polymorph     bool
	Multiple      bool
	Optional      bool
	DocComment    string
}

// FieldType holds type information of a child element.
type FieldType struct {
	Name             string
	IsNestedResource bool
	IsPrimitive      bool
}

// FilterResources filters resources, types or backbone elements for only resources.
func FilterResources(t []ResourceOrType) []ResourceOrType {
	var resources []ResourceOrType
	for _, r := range t {
		if r.IsResource {
			resources = append(resources, r)
		}
	}
	return resources
}
