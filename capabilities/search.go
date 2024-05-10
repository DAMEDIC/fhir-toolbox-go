package capabilities

type SearchCapabilities struct {
	Include    []string
	Parameters []string
}

type SearchParameters = map[string]SearchParameterAnd
type SearchParameterAnd = []SearchParameterOr
type SearchParameterOr = []string
