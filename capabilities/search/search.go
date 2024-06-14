package search

type Capabilities struct {
	Parameters map[string]ParameterDesc
	Includes   []string
}

type ParameterDesc struct {
	Type Type
}

type Type string

const (
	Number    Type = "number"
	Date      Type = "date"
	String    Type = "string"
	Token     Type = "token"
	Reference Type = "reference"
	Composite Type = "composite"
	Quantity  Type = "quantity"
	Uri       Type = "uri"
	Special   Type = "special"
)

type Options struct {
	Parameters Parameters
	Includes   []string
}

type Parameters = map[string]ParameterAndList
type ParameterAndList = []ParameterOrList
type ParameterOrList = []ParameterValue

type ParameterValue struct {
	Prefix string
	Value  string
}

type Prefix = string

const (
	Equal          Prefix = "eq"
	NotEqual       Prefix = "ne"
	GreaterThan    Prefix = "gt"
	LessThan       Prefix = "lt"
	GreaterOrEqual Prefix = "ge"
	LessOrEqual    Prefix = "le"
)

var KnownPrefixes = []Prefix{Equal, NotEqual, GreaterThan, LessThan, GreaterOrEqual, LessOrEqual}
