package ir

type ElementMap struct {
	Map map[string]map[string]struct{}
}

type SourceFile struct {
	Name    string
	Structs []Struct
}

type Struct struct {
	Name        string
	IsPrimitive bool
	Fields      []StructField
	DocComment  string
}

type StructField struct {
	Name        string
	MarshalName string
	Type        FieldType
	Polymorph   bool
	Multiple    bool
	Optional    bool
	DocComment  string
}

type FieldType struct {
	Name    string
	Package string
}
