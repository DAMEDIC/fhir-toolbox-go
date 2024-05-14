package capabilities

type SearchCapabilities struct {
	Parameters []string
	Includes   []string
}

type SearchOptions struct {
	Parameters []SearchParameters
	Includes   []string
}

type SearchParameters = map[string]SearchParameterAnd
type SearchParameterAnd = []SearchParameterOr
type SearchParameterOr = []string
