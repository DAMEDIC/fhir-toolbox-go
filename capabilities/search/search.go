// Package search contains types and helpers to work with [FHIR Search].
// You can use these to provide search capabilities for your custom implementation.
//
// Currently, only an API with cursor-based pagination is supported.
// Parameters for offset based pagination might be added eventually if there is demand.
//
// # Example
// ```Go
//
//	import "fhir-toolbox/capabilities/search"
//
//	func (b *myAPI) SearchCapabilitiesObservation() search.Capabilities {
//		   // return supported search capabilities
//		   return search.Capabilities{
//		       Parameters: map[string]search.ParameterDescription{
//	           "_id": {Type: search.Token},
//		       },
//	   }
//	}
//
//	func (b *myAPI) SearchObservation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
//		   // return the search result
//		   return search.Result{ ... }, nil
//	}
//
// ```
//
// [FHIR Search]: https://hl7.org/fhir/search.html
package search

import (
	"fhir-toolbox/model"
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"
)

// Result contains the result of a search operation.
type Result struct {
	Resources []model.Resource
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
	Parameters map[string]ParameterDescription
	Includes   []string
}

// ParameterDescription describes a parameter that is supported by the implementation.
type ParameterDescription struct {
	Type ParameterType
}

// ParameterType is the type of the parameter
type ParameterType string

const (
	Number    ParameterType = "number"
	Date                    = "date"
	String                  = "string"
	Token                   = "token"
	Reference               = "reference"
	Composite               = "composite"
	Quantity                = "quantity"
	Uri                     = "uri"
	Special                 = "special"
)

// Search Options passed to an implementation.
type Options struct {
	Parameters Parameters
	Includes   []string
	Count      int
	Cursor     Cursor
}

// Parameters of a
type Parameters map[string]AndList
type AndList []OrList
type OrList []Value

// A Value to search for.
type Value struct {
	Prefix        Prefix
	DatePrecision DatePrecision
	Value         any
}

// A Prefix that some parameter types can use to change search behavior.
// See https://hl7.org/fhir/search.html#prefix.
type Prefix = string

const (
	Equal          Prefix = "eq"
	NotEqual              = "ne"
	GreaterThan           = "gt"
	LessThan              = "lt"
	GreaterOrEqual        = "ge"
	LessOrEqual           = "le"
	StartsAfter           = "sa"
	EndsBefore            = "eb"
)

var KnownPrefixes = []Prefix{Equal, NotEqual, GreaterThan, LessThan, GreaterOrEqual, LessOrEqual, StartsAfter, EndsBefore}

// DatePrecision represents the precision of date value.
type DatePrecision string

const (
	YearPrecision       DatePrecision = "year"
	MonthPrecision                    = "month"
	DayPrecision                      = "day"
	HourMinutePrecision               = "hourMinute"
	FullTimePrecision                 = "time"
)

// Format strings for precision aware parsing and encoding.
const (
	OnlyYearFormat   = "2006"
	UpToMonthFormat  = "2006-01"
	UpToDayFormat    = "2006-01-02"
	HourMinuteFormat = "2006-01-02T15:04Z07:00"
	FullTimeFormat   = "2006-01-02T15:04:05.999999999Z07:00"
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
	case YearPrecision:
		return v.Format(OnlyYearFormat)
	case MonthPrecision:
		return v.Format(UpToMonthFormat)
	case DayPrecision:
		return v.Format(UpToDayFormat)
	case HourMinutePrecision:
		return v.Format(HourMinuteFormat)
	default:
		return v.Format(FullTimeFormat)
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

	builder.WriteString(o.Parameters.Query().Encode())

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
func (p Parameters) Query() url.Values {
	values := url.Values{}

	for key, ands := range p {
		value := make([]string, 0, len(ands))

		for _, ors := range ands {
			s := make([]string, 0, len(ors))
			for _, or := range ors {
				s = append(s, or.String())
			}
			slices.Sort(s)
			value = append(value, strings.Join(s, ","))
		}

		slices.Sort(value)
		values[key] = value
	}

	return values
}

// ParseOptions parses search options from a [url.Values] query string.
//
// Only parameters supported by the backing implementation as described
// by the passed `searchCapabilities` are used.
//
// [Result modifying parameters] are parsed into separate fields on the [Options] object.
// All other parameters are parsed into [Options.Parameters].
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
		Parameters: Parameters{},
		Count:      min(defaultCount, maxCount),
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
			desc, ok := searchCapabilities.Parameters[k]
			if !ok {
				// only known parameters are forwarded
				continue
			}

			ands, err := parseSearchParam(v, desc, tz, k)
			if err != nil {
				return Options{}, err
			}

			options.Parameters[k] = ands
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

func parseSearchParam(values []string, desc ParameterDescription, tz *time.Location, k string) (AndList, error) {
	ands := make(AndList, 0, len(values))
	for _, ors := range values {
		splitStrings := strings.Split(ors, ",")

		ors := make(OrList, 0, len(splitStrings))
		for _, s := range splitStrings {
			value, err := parseSearchValue(desc.Type, s, tz)
			if err != nil {
				return nil, fmt.Errorf("invalid search value for parameter %s: %w", k, err)
			}
			ors = append(ors, value)
		}

		ands = append(ands, ors)
	}
	return ands, nil
}

func parseSearchValue(typ ParameterType, value string, tz *time.Location) (Value, error) {
	prefix := parseSearchValuePrefix(typ, value)
	if prefix != "" {
		// all prefixes have a width of 2
		value = value[2:]
	}

	valueAny, datePrecision, err := parseSearchValueAny(typ, value, tz)
	if err != nil {
		return Value{}, err
	}

	return Value{Prefix: prefix, DatePrecision: datePrecision, Value: valueAny}, nil
}

func parseSearchValuePrefix(typ ParameterType, value string) string {
	// all prefixes have a width of 2
	if len(value) < 2 {
		return ""
	}

	// only number, date and quantity can have prefixes
	if !slices.Contains([]ParameterType{Number, Date, Quantity}, typ) {
		return ""
	}

	if !slices.Contains(KnownPrefixes, value[:2]) {
		return ""
	}

	return value[:2]
}

func parseSearchValueAny(typ ParameterType, value string, tz *time.Location) (any, DatePrecision, error) {
	switch typ {
	case Date:
		return parseDate(value, tz)
	default:
		return value, "", nil
	}
}

func parseDate(value string, tz *time.Location) (time.Time, DatePrecision, error) {
	date, err := time.ParseInLocation(OnlyYearFormat, value, tz)
	if err == nil {
		return date, YearPrecision, nil
	}
	date, err = time.ParseInLocation(UpToMonthFormat, value, tz)
	if err == nil {
		return date, MonthPrecision, nil
	}
	date, err = time.ParseInLocation(UpToDayFormat, value, tz)
	if err == nil {
		return date, DayPrecision, nil
	}
	date, err = time.ParseInLocation(HourMinuteFormat, value, tz)
	if err == nil {
		return date, HourMinutePrecision, nil
	}
	date, err = time.ParseInLocation(FullTimeFormat, value, tz)
	if err == nil {
		return date, FullTimePrecision, nil
	}
	return time.Time{}, "", err
}
