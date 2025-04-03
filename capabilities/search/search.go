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
//			Parameters: ParameterDescriptions{
//				"_id": {Type: search.TypeToken},
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
	"github.com/cockroachdb/apd/v3"
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
	Parameters ParameterDescriptions
	Includes   []string
}

type ParameterDescriptions map[string]ParameterDescription

// ParameterDescription describes a parameter that is supported by the implementation.
type ParameterDescription struct {
	Type      ParameterType
	Modifiers []Modifier
}

// Options represents the parameters passed to a search implementation.
type Options struct {
	// Parameters defines the search parameters.
	Parameters Parameters
	// Includes specifies the related resources to include in the search results.
	Includes []string
	// Count defines the maximum number of results to return.
	Count int
	// Cursor allows for pagination of large result sets.
	Cursor Cursor
}

// Parameters defines a map of search parameters, where the key is the parameter name
// (potentially including modifiers)and the value is a slice of possible values for that parameter.
type Parameters map[ParameterKey]All

// All represents a slice of possible values for a single search parameter where each of the entry has to match.
type All []OneOf

// OneOf represents a slice of possible values for a single search parameter,
// where only one of the values has to match.
type OneOf []Value

// ParameterKey represents a key for a search parameter,
// consisting of a name and an optional modifier.
type ParameterKey struct {
	// Name is the name of the search parameter.
	Name string
	// Modifier is an optional modifier that can be applied to the search parameter,
	// such as `exact`, `contains`, `identifier`, etc.
	Modifier Modifier
}

func (p ParameterKey) String() string {
	if p.Modifier == "" {
		return p.Name
	} else {
		return fmt.Sprintf("%s:%s", p.Name, p.Modifier)
	}
}

func (p ParameterKey) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

// ParameterType is the type of the parameter
type ParameterType string

const (
	TypeNumber    ParameterType = "number"
	TypeDate      ParameterType = "date"
	TypeString    ParameterType = "string"
	TypeToken     ParameterType = "token"
	TypeReference ParameterType = "reference"
	TypeComposite ParameterType = "composite"
	TypeQuantity  ParameterType = "quantity"
	TypeUri       ParameterType = "uri"
	TypeSpecial   ParameterType = "special"
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

// A Prefix that some parameter types can use to change search behavior.
// See https://hl7.org/fhir/search.html#prefix.
type Prefix = string

const (
	PrefixEqual          Prefix = "eq"
	PrefixNotEqual       Prefix = "ne"
	PrefixGreaterThan    Prefix = "gt"
	PrefixLessThan       Prefix = "lt"
	PrefixGreaterOrEqual Prefix = "ge"
	PrefixLessOrEqual    Prefix = "le"
	PrefixStartsAfter    Prefix = "sa"
	PrefixEndsBefore     Prefix = "eb"
)

var KnownPrefixes = []Prefix{
	PrefixEqual,
	PrefixNotEqual,
	PrefixGreaterThan,
	PrefixLessThan,
	PrefixGreaterOrEqual,
	PrefixLessOrEqual,
	PrefixStartsAfter,
	PrefixEndsBefore,
}

// ParseOptions parses search options from a [url.Values] query string.
//
// Only parameters supported by the backing implementation as described
// by the passed `searchCapabilities` are used.
//
// [Result modifying parameters] are parsed into separate fields on the [Options] object.
// All other parameters are parsed into [options.Parameters].
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
			splits := strings.Split(k, ":")
			param := ParameterKey{
				Name: splits[0],
			}
			if len(splits) > 1 {
				param.Modifier = Modifier(splits[1])
			}

			desc, ok := searchCapabilities.Parameters[param.Name]
			if !ok {
				// only known parameters are forwarded
				continue
			}

			// When the :identifier modifier is used, the search value works as a token search.
			if desc.Type == TypeReference && param.Modifier == ModifierIdentifier {
				desc.Type = TypeToken
			}

			ands, err := parseSearchParam(param, v, desc, tz)
			if err != nil {
				return Options{}, err
			}

			options.Parameters[param] = ands
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

func parseSearchParam(param ParameterKey, values []string, desc ParameterDescription, tz *time.Location) (All, error) {
	if param.Modifier != "" && !slices.Contains(desc.Modifiers, param.Modifier) {
		return nil, fmt.Errorf("unsupported modifier for parameter %s, supported are: %s", param, desc.Modifiers)
	}

	ands := make(All, 0, len(values))
	for _, ors := range values {
		splitStrings := strings.Split(ors, ",")

		ors := make(OneOf, 0, len(splitStrings))
		for _, s := range splitStrings {
			value, err := parseSearchValue(desc, s, tz)
			if err != nil {
				return nil, fmt.Errorf("invalid search value for parameter %s: %w", param, err)
			}
			ors = append(ors, value)
		}

		ands = append(ands, ors)
	}
	return ands, nil
}

func parseSearchValue(desc ParameterDescription, value string, tz *time.Location) (Value, error) {
	prefix := parseSearchValuePrefix(desc.Type, value)
	if prefix != "" {
		// all prefixes have a width of 2
		value = value[2:]
	}

	switch desc.Type {
	case TypeNumber:
		dec, _, err := apd.NewFromString(value)
		return Number{
			Prefix: prefix,
			Value:  dec,
		}, err
	case TypeDate:
		date, prec, err := ParseDate(value, tz)
		return Date{
			Prefix:    prefix,
			Precision: prec,
			Value:     date,
		}, err
	case TypeString:
		return String(value), nil
	case TypeToken:
		s := strings.Split(value, "|")
		switch len(s) {
		case 1:
			return Token{
				System: nil,
				Code:   s[0],
			}, nil
		case 2:
			system, err := url.Parse(s[0])
			if err != nil {
				return nil, fmt.Errorf("invalid token system %s: %w", value, err)
			}
			return Token{
				System: system,
				Code:   s[1],
			}, nil
		default:
			return nil, fmt.Errorf("invalid token %s", value)
		}
	case TypeReference:
		// if url, there may be a version appended
		urlSplit := strings.Split(value, "|")

		url, err := url.Parse(urlSplit[0])
		if err != nil {
			return nil, fmt.Errorf("invalid reference %s: %w", value, err)
		}

		if url.Scheme != "" {
			switch len(urlSplit) {
			case 1:
				return Reference{
					Url: url,
				}, nil
			case 2:
				return Reference{
					Url:     url,
					Version: urlSplit[1],
				}, nil
			default:
				return nil, fmt.Errorf("invalid reference %s", value)

			}
		}

		// no real URL, thus local reference
		localIdSplit := strings.Split(value, "/")
		switch len(localIdSplit) {
		case 1:
			return Reference{
				Id: localIdSplit[0],
			}, nil
		case 2:
			return Reference{
				Type: localIdSplit[0],
				Id:   localIdSplit[1],
			}, nil
		case 4:
			if localIdSplit[2] != "_history" {
				return nil, fmt.Errorf("invalid reference %s, expected _history at 3rd position", value)
			}
			return Reference{
				Type:    localIdSplit[0],
				Id:      localIdSplit[1],
				Version: localIdSplit[3],
			}, nil
		default:
			return nil, fmt.Errorf("invalid reference %s", value)
		}

	case TypeComposite:
		return Composite(strings.Split(value, "$")), nil

	case TypeQuantity:
		s := strings.Split(value, "|")
		number, _, err := apd.NewFromString(s[0])
		if err != nil {
			return nil, fmt.Errorf("invalid quantity number: %w", err)
		}

		switch len(s) {
		case 1:
			return Quantity{
				Prefix: prefix,
				Value:  number,
			}, nil
		case 3:
			system, err := url.Parse(s[1])
			if err != nil {
				return nil, fmt.Errorf("invalid quantity system %s: %w", value, err)
			}
			return Quantity{
				Prefix: prefix,
				Value:  number,
				System: system,
				Code:   s[2],
			}, nil

		default:
			return nil, fmt.Errorf("invalid quantity %s", value)
		}
	case TypeUri:
		url, err := url.Parse(value)
		if err != nil {
			return nil, fmt.Errorf("invalid reference: %w", err)
		}
		return Uri{
			Value: url,
		}, nil
	case TypeSpecial:
		return Special(value), nil

	default:
		return nil, fmt.Errorf("unsupported type %s", desc.Type)
	}
}

func parseSearchValuePrefix(typ ParameterType, value string) string {
	// all prefixes have a width of 2
	if len(value) < 2 {
		return ""
	}

	// only number, date and quantity can have prefixes
	if !slices.Contains([]ParameterType{TypeNumber, TypeDate, TypeQuantity}, typ) {
		return ""
	}

	if !slices.Contains(KnownPrefixes, value[:2]) {
		return ""
	}

	return value[:2]
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
		for _, ors := range ands {
			if len(ors) == 0 {
				continue
			}

			nameWithModifier := key.Name
			if key.Modifier != "" {
				nameWithModifier = fmt.Sprintf("%s:%s", key.Name, key.Modifier)
			}

			s := make([]string, 0, len(ors))
			for _, or := range ors {
				s = append(s, or.String())
			}
			slices.Sort(s)

			values.Add(nameWithModifier, strings.Join(s, ","))
			slices.Sort(values[nameWithModifier])
		}

	}

	return values
}

// Value of a search parameter, determine the concrete type by type assertion.
//
//	switch t := value.(type) {
//	case search.Number:
//	  // handle search parameter of type number
//	}
type Value interface {
	isValue()
	String() string
}

type Number struct {
	Prefix Prefix
	Value  *apd.Decimal
}

func (n Number) isValue() {}

func (n Number) String() string {
	if n.Prefix != "" {
		return fmt.Sprintf("%s%s", n.Prefix, n.Value.String())
	} else {
		return n.Value.String()
	}
}

type Date struct {
	Prefix    Prefix
	Precision DatePrecision
	Value     time.Time
}

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

func (d Date) isValue() {}

func (d Date) String() string {
	b := strings.Builder{}
	if d.Prefix != "" {
		b.WriteString(d.Prefix)
	}
	switch d.Precision {
	case PrecisionYear:
		b.WriteString(d.Value.Format(DateFormatOnlyYear))
	case PrecisionMonth:
		b.WriteString(d.Value.Format(DateFormatUpToMonth))
	case PrecisionDay:
		b.WriteString(d.Value.Format(DateFormatUpToDay))
	case PrecisionHourMinute:
		b.WriteString(d.Value.Format(DateFormatHourMinute))
	default:
		b.WriteString(d.Value.Format(DateFormatFullTime))
	}
	return b.String()
}

type String string

func (s String) isValue() {}

func (s String) String() string {
	return string(s)
}

type Token struct {
	// Go URLs can contain URIs
	System *url.URL
	Code   string
}

func (t Token) String() string {
	if t.System == nil {
		return t.Code
	} else {
		return fmt.Sprintf("%s|%s", t.System, t.Code)
	}
}

func (t Token) isValue() {}

type Reference struct {
	Modifier Modifier
	Id       string
	Type     string
	Url      *url.URL
	Version  string
}

func (r Reference) isValue() {}

func (r Reference) String() string {
	if r.Url != nil {
		b := strings.Builder{}
		b.WriteString(r.Url.String())
		if r.Version != "" {
			b.WriteRune('|')
			b.WriteString(r.Version)
		}
		return b.String()
	}

	if r.Type == "" {
		return r.Id
	}

	b := strings.Builder{}
	b.WriteString(r.Type)
	b.WriteRune('/')
	b.WriteString(r.Id)
	if r.Version != "" {
		b.WriteString("/_history/")
		b.WriteString(r.Version)
	}
	return b.String()
}

type Composite []string

func (c Composite) String() string {
	return strings.Join(c, "$")
}

func (c Composite) isValue() {}

type Quantity struct {
	Prefix Prefix
	Value  *apd.Decimal
	System *url.URL
	Code   string
}

func (q Quantity) isValue() {}

func (q Quantity) String() string {
	b := strings.Builder{}
	b.WriteString(q.Prefix)
	b.WriteString(q.Value.String())
	if q.Code != "" {
		b.WriteRune('|')
		b.WriteString(q.System.String())
		b.WriteRune('|')
		b.WriteString(q.Code)
	}
	return b.String()
}

type Uri struct {
	// Go URLs can contain URIs
	Value *url.URL
}

func (u Uri) String() string {
	return u.Value.String()
}

func (u Uri) isValue() {}

// Special string contains potential prefixes
type Special string

func (s Special) isValue() {}

func (s Special) String() string {
	return string(s)
}
