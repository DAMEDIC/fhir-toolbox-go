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

// Parameters of a search.
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
// See https://www.hl7.org/fhir/search.html#prefix.
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
// The function is deterministic, the same `option` will always yield the same output.
//
// [result modifying parameters]: https://www.hl7.org/fhir/search.html#modifyingresults
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
