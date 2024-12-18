// Package search contains types and helpers to work with [FHIR Search].
// You can use these to provide search capabilities for your custom implementation.
//
// Currently, only an API with cursor-based pagination is supported.
// Parameters for offset based pagination might be added eventually if there is demand.
//
// # Example
//
//	import "github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
//
//	func (b *myAPI) SearchCapabilitiesObservation() search.Capabilities {
//		// return supported search capabilities
//		return search.Capabilities{
//			Params: map[string]search.ParamDesc{
//				"_id": {Type: search.Token},
//			},
//		}
//	}
//
//	func (b *myAPI) SearchObservation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
//		// return the search result
//		return search.Result{ ... }, nil
//	}
//
// [FHIR Search]: https://hl7.org/fhir/search.html
package search

import (
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"
)

// Result contains the result of a search operation.
type Result struct {
	Resources []model.Resource
	Included  []model.Resource
	Next      Cursor
}

// Cursor is used for pagination.
//
// It references where the server shall pick up querying additional results
// for multi-page queries.
type Cursor string

// Capabilities describe what search capabilities the server provides.
//
// It can be used to derive [CapabilityStatements] which describe what a FHIR system can do.
//
// [CapabilityStatements]: https://hl7.org/fhir/capabilitystatement.html
type Capabilities struct {
	Params   map[string]ParamDesc
	Includes []string
}

// ParamDesc describes a parameter that is supported by the implementation.
type ParamDesc struct {
	Type      ParamType
	Modifiers []Modifier
}

// ParamType is the type of the parameter
type ParamType string

const (
	Number    ParamType = "number"
	Date      ParamType = "date"
	String    ParamType = "string"
	Token     ParamType = "token"
	Reference ParamType = "reference"
	Composite ParamType = "composite"
	Quantity  ParamType = "quantity"
	Uri       ParamType = "uri"
	Special   ParamType = "special"
)

// Modifier supported by an implementation.
type Modifier string

const (
	// ModifierAbove on reference, token, uri: Tests whether the value in a resource is or subsumes the supplied parameter value (is-a, or hierarchical relationships).
	ModifierAbove Modifier = "above"
	// ModifierBelow on reference, token, uri: Tests whether the value in a resource is or is subsumed by the supplied parameter value (is-a, or hierarchical relationships).
	ModifierBelow Modifier = "below"
	// ModifierCodeText on reference, token: Tests whether the textual display value in a resource (e.g., CodeableConcept.text, Coding.display, or Reference.display) matches the supplied parameter value.
	ModifierCodeText Modifier = "code-text"
	// ModifierContains on string, uri: Tests whether the value in a resource includes the supplied parameter value anywhere within the field being searched.
	ModifierContains Modifier = "contains"
	// ModifierExact on string: Tests whether the value in a resource exactly matches the supplied parameter value (the whole string, including casing and accents).
	ModifierExact Modifier = "exact"
	// ModifierIdentifier on reference: Tests whether the Reference.identifier in a resource (rather than the Reference.reference) matches the supplied parameter value.
	ModifierIdentifier Modifier = "identifier"
	// ModifierInModifier on token: Tests whether the value in a resource is a member of the supplied parameter ValueSet.
	ModifierInModifier = "in"
	// ModifierIterate Not allowed anywhere by default: The search parameter indicates an inclusion directive (_include, _revinclude) that is applied to an included resource instead of the matching resource.
	ModifierIterate Modifier = "iterate"
	// ModifierMissing on date, number, quantity, reference, string, token, uri: Tests whether the value in a resource is present (when the supplied parameter value is true) or absent (when the supplied parameter value is false).
	ModifierMissing Modifier = "missing"
	// ModifierNot on token: Tests whether the value in a resource does not match the specified parameter value. Note that this includes resources that have no value for the parameter.
	ModifierNot Modifier = "not"
	// ModifierNotIn on reference, token: Tests whether the value in a resource is not a member of the supplied parameter ValueSet.
	ModifierNotIn Modifier = "not-in"
	// ModifierOfType on token (only Identifier): Tests whether the Identifier value in a resource matches the supplied parameter value.
	ModifierOfType Modifier = "of-type"
	// ModifierText on reference, token: Tests whether the textual value in a resource (e.g., CodeableConcept.text, Coding.display, Identifier.type.text, or Reference.display) matches the supplied parameter value using basic string matching (begins with or is, case-insensitive).
	//
	//on string: The search parameter value should be processed as input to a search with advanced text handling.
	ModifierText Modifier = "text"
	// ModifierTextAdvancedModifier on reference, token: Tests whether the value in a resource matches the supplied parameter value using advanced text handling that searches text associated with the code/value - e.g., CodeableConcept.text, Coding.display, or Identifier.type.text.
	ModifierTextAdvancedModifier = "text-advanced"
	// ModifierType on reference: Tests whether the value in a resource points to a resource of the supplied parameter type. Note: a concrete ResourceType is specified as the modifier (e.g., not the literal :[type], but a value such as :Patient).
	ModifierType Modifier = "[type]"
)

// Options passed to a search implementation.
type Options struct {
	Params   Params
	Includes []string
	Count    int
	Cursor   Cursor
}

type Params map[string]AndList
type AndList []OrList
type OrList []Value

// A Value to search for.
type Value struct {
	Prefix        Prefix
	Modifier      Modifier
	DatePrecision DatePrecision
	Value         any
}

// A Prefix that some parameter types can use to change search behavior.
// See https://hl7.org/fhir/search.html#prefix.
type Prefix = string

const (
	Equal          Prefix = "eq"
	NotEqual       Prefix = "ne"
	GreaterThan    Prefix = "gt"
	LessThan       Prefix = "lt"
	GreaterOrEqual Prefix = "ge"
	LessOrEqual    Prefix = "le"
	StartsAfter    Prefix = "sa"
	EndsBefore     Prefix = "eb"
)

var KnownPrefixes = []Prefix{Equal, NotEqual, GreaterThan, LessThan, GreaterOrEqual, LessOrEqual, StartsAfter, EndsBefore}

// DatePrecision represents the precision of date value.
type DatePrecision string

const (
	PrecisionYear       DatePrecision = "year"
	PrecisionMonth      DatePrecision = "month"
	PrecisionDay        DatePrecision = "day"
	PrecisionHourMinute DatePrecision = "hourMinute"
	PrecisionFullTime   DatePrecision = "time"
)

// Format strings for precision aware parsing and encoding.
const (
	DateFormatOnlyYear   = "2006"
	DateFormatUpToMonth  = "2006-01"
	DateFormatUpToDay    = "2006-01-02"
	DateFormatHourMinute = "2006-01-02T15:04Z07:00"
	DateFormatFullTime   = "2006-01-02T15:04:05.999999999Z07:00"
)

// String formats the value as string.
//
// Prefixes and date precision are taken into account.
func (v Value) String() string {
	var s string

	switch w := v.Value.(type) {
	case time.Time:
		s = formatDate(w, v.DatePrecision)
	default:
		s = fmt.Sprintf("%v", v.Value)
	}

	return v.Prefix + s
}

func formatDate(v time.Time, p DatePrecision) string {
	switch p {
	case PrecisionYear:
		return v.Format(DateFormatOnlyYear)
	case PrecisionMonth:
		return v.Format(DateFormatUpToMonth)
	case PrecisionDay:
		return v.Format(DateFormatUpToDay)
	case PrecisionHourMinute:
		return v.Format(DateFormatHourMinute)
	default:
		return v.Format(DateFormatFullTime)
	}
}

// QueryString of the search options.
//
// Search parameters are sorted alphabetically, [result modifying parameters] like `_include`
// are appended at the end.
// The function is deterministic, the same `option` input will always yield the same output.
//
// [result modifying parameters]: https://hl7.org/fhir/search.html#modifyingresults
func (o Options) QueryString() string {
	var builder strings.Builder

	builder.WriteString(o.Params.Query().Encode())

	if len(o.Includes) > 0 {
		includes := append([]string{}, o.Includes...)
		slices.Sort(includes)

		for _, include := range includes {
			if builder.Len() > 0 {
				builder.WriteByte('&')
			}
			builder.WriteString("_include=")
			builder.WriteString(url.QueryEscape(include))
		}
	}

	if o.Cursor != "" {
		if builder.Len() > 0 {
			builder.WriteByte('&')
		}
		builder.WriteString("_cursor=")
		builder.WriteString(url.QueryEscape(string(o.Cursor)))
	}

	if o.Count > 0 {
		if builder.Len() > 0 {
			builder.WriteByte('&')
		}
		builder.WriteString("_count=")
		builder.WriteString(strconv.Itoa(o.Count))
	}

	return builder.String()
}

// Query representing the search parameters.
//
// All contained values are sorted, but the returned [url.Values] is backed by a map.
// To obtain a deterministic query string you can call [url.Values.Encode], because
// it will sort the keys alphabetically.
func (p Params) Query() url.Values {
	values := url.Values{}

	for name, ands := range p {
		for _, ors := range ands {
			if len(ors) == 0 {
				continue
			}

			// we assume that all OR values use the same modifier
			modifier := ors[0].Modifier
			key := name
			if modifier != "" {
				key = fmt.Sprintf("%s:%s", name, modifier)
			}

			s := make([]string, 0, len(ors))
			for _, or := range ors {
				s = append(s, or.String())
			}
			slices.Sort(s)

			values.Add(key, strings.Join(s, ","))
			slices.Sort(values[key])
		}

	}

	return values
}

// ParseOptions parses search options from a [url.Values] query string.
//
// Only parameters supported by the backing implementation as described
// by the passed `searchCapabilities` are used.
//
// [Result modifying parameters] are parsed into separate fields on the [Options] object.
// All other parameters are parsed into [options.Params].
//
// [Result modifying parameters]: https://hl7.org/fhir/search.html#modifyingresults
//
// [result modifying parameters]: https://hl7.org/fhir/search.html#modifyingresults
func ParseOptions(
	searchCapabilities Capabilities,
	params url.Values,
	tz *time.Location,
	maxCount, defaultCount int,
) (Options, error) {
	options := Options{
		// Parameters is backed by a map, which must be initialized
		Params: Params{},
		Count:  min(defaultCount, maxCount),
	}

	for k, v := range params {
		switch k {
		case "_count":
			count, err := parseCount(v, maxCount)
			if err != nil {
				return Options{}, err
			}
			options.Count = count

		case "_cursor":
			cursor, err := parseCursor(v)
			if err != nil {
				return Options{}, err
			}
			options.Cursor = cursor

		case "_include":
			options.Includes = v

		// other result modifying parameters which are not supported yet:
		// https://hl7.org/fhir/search.html#modifyingresults
		case "_contained", "_elements", "_graph", "_maxresults", "_revinclude", "_score", "_summary", "_total":

		default:
			splits := strings.Split(k, ":")
			param := splits[0]
			var modifier Modifier
			if len(splits) > 1 {
				modifier = Modifier(splits[1])
			}

			desc, ok := searchCapabilities.Params[param]
			if !ok {
				// only known parameters are forwarded
				continue
			}

			ands, err := parseSearchParam(param, modifier, v, desc, tz)
			if err != nil {
				return Options{}, err
			}

			options.Params[param] = ands
		}
	}

	return options, nil
}

func parseCount(values []string, maxCount int) (int, error) {
	if len(values) != 1 {
		return 0, fmt.Errorf("multiple _count parameters")
	}
	count, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, fmt.Errorf("invalid _count parameter: %w", err)
	}
	return min(count, maxCount), nil
}

func parseCursor(values []string) (Cursor, error) {
	if len(values) != 1 {
		return "", fmt.Errorf("multiple _cursor parameters")
	}
	return Cursor(values[0]), nil
}

func parseSearchParam(param string, modifier Modifier, values []string, desc ParamDesc, tz *time.Location) (AndList, error) {
	if modifier != "" && !slices.Contains(desc.Modifiers, modifier) {
		return nil, fmt.Errorf("unsupported modifier for parameter %s:%s, supported are: %s", param, modifier, desc.Modifiers)
	}

	ands := make(AndList, 0, len(values))
	for _, ors := range values {
		splitStrings := strings.Split(ors, ",")

		ors := make(OrList, 0, len(splitStrings))
		for _, s := range splitStrings {
			value, err := parseSearchValue(desc.Type, s, modifier, tz)
			if err != nil {
				return nil, fmt.Errorf("invalid search value for parameter %s: %w", param, err)
			}
			ors = append(ors, value)
		}

		ands = append(ands, ors)
	}
	return ands, nil
}

func parseSearchValue(typ ParamType, value string, modifier Modifier, tz *time.Location) (Value, error) {
	prefix := parseSearchValuePrefix(typ, value)
	if prefix != "" {
		// all prefixes have a width of 2
		value = value[2:]
	}

	valueAny, datePrecision, err := parseSearchValueAny(typ, value, tz)
	if err != nil {
		return Value{}, err
	}

	return Value{Prefix: prefix, Modifier: modifier, DatePrecision: datePrecision, Value: valueAny}, nil
}

func parseSearchValuePrefix(typ ParamType, value string) string {
	// all prefixes have a width of 2
	if len(value) < 2 {
		return ""
	}

	// only number, date and quantity can have prefixes
	if !slices.Contains([]ParamType{Number, Date, Quantity}, typ) {
		return ""
	}

	if !slices.Contains(KnownPrefixes, value[:2]) {
		return ""
	}

	return value[:2]
}

func parseSearchValueAny(typ ParamType, value string, tz *time.Location) (any, DatePrecision, error) {
	switch typ {
	case Date:
		return ParseDate(value, tz)
	default:
		return value, "", nil
	}
}

func ParseDate(value string, tz *time.Location) (time.Time, DatePrecision, error) {
	date, err := time.ParseInLocation(DateFormatOnlyYear, value, tz)
	if err == nil {
		return date, PrecisionYear, nil
	}
	date, err = time.ParseInLocation(DateFormatUpToMonth, value, tz)
	if err == nil {
		return date, PrecisionMonth, nil
	}
	date, err = time.ParseInLocation(DateFormatUpToDay, value, tz)
	if err == nil {
		return date, PrecisionDay, nil
	}
	date, err = time.ParseInLocation(DateFormatHourMinute, value, tz)
	if err == nil {
		return date, PrecisionHourMinute, nil
	}
	date, err = time.ParseInLocation(DateFormatFullTime, value, tz)
	if err == nil {
		return date, PrecisionFullTime, nil
	}
	return time.Time{}, "", err
}
