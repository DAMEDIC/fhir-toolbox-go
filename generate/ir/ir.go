package ir

type ElementMap struct {
	Map map[string]map[string]struct{}
}

type SourceFile struct {
	Name    string
	Structs []*Struct
}

type Struct struct {
	Name             string
	IsResource       bool
	IsDomainResource bool
	IsPrimitive      bool
	Fields           []StructField
	DocComment       string
}

type StructField struct {
	Name          string
	MarshalName   string
	PossibleTypes []FieldType
	Polymorph     bool
	Multiple      bool
	Optional      bool
	DocComment    string
}

type FieldType struct {
	Name             string
	IsNestedResource bool
	IsPrimitive      bool
}

func FilterResources(files []SourceFile) []Struct {
	var resources []Struct
	for _, file := range files {
		for _, s := range file.Structs {
			if s.IsResource {
				resources = append(resources, *s)
			}
		}
	}
	return resources
}
