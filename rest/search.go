package rest

import (
	"context"
	"fhir-toolbox/capabilities"
	"fhir-toolbox/capabilities/search"
	"fhir-toolbox/model"
	"fhir-toolbox/model/basic"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"
)

// SearchError indicates that a search could not be performed.
type SearchError struct {
	error
}

func (e SearchError) Error() string {
	return fmt.Sprintf("processing search: %v", e.error)
}

func (e SearchError) StatusCode() int {
	return 500
}

func (e SearchError) OperationOutcome() model.Resource {
	return basic.OperationOutcome{
		Issue: []basic.OperationOutcomeIssue{
			{Severity: "fatal", Code: "processing", Diagnostics: e.Error()},
		},
	}
}

func dispatchSearch(
	context context.Context,
	backend capabilities.GenericAPI,
	resourceType string,
	parameters url.Values,
	baseURL *url.URL,
	tz *time.Location,
	maxCount,
	defaultCount int,
) (int, model.Resource) {
	searchCapabilities, err := backend.SearchCapabilities(resourceType)

	options, err := parseSearchOptions(searchCapabilities, parameters, tz, maxCount, defaultCount)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	resources, err := backend.Search(context, resourceType, options)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	bundle, err := NewSearchBundle(resourceType, resources, options, searchCapabilities, baseURL)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	return http.StatusOK, bundle
}

func parseSearchOptions(
	searchCapabilities search.Capabilities,
	params url.Values,
	tz *time.Location,
	maxCount,
	defaultCount int,
) (search.Options, capabilities.FHIRError) {
	options := search.Options{
		// need to initialize empty map
		Parameters: make(search.Parameters),
	}

	count, err := parseCount(params, maxCount, defaultCount)
	if err != nil {
		return search.Options{}, err
	}
	options.Count = count

	cursor, err := parseCursor(params)
	if err != nil {
		return search.Options{}, err
	}
	options.Cursor = cursor

	for k, v := range params {
		switch k {
		case "_include":
			options.Includes = v

		default:
			desc, ok := searchCapabilities.Parameters[k]
			if !ok {
				// only known parameters are forwarded
				continue
			}

			ands := make(search.AndList, 0, len(v))
			for _, ors := range v {
				splitStrings := strings.Split(ors, ",")

				ors := make(search.OrList, 0, len(splitStrings))
				for _, s := range splitStrings {
					value, err := parseSearchValue(desc.Type, s, tz)
					if err != nil {
						return search.Options{}, SearchError{
							error: fmt.Errorf("invalid search value for parameter %s: %w", k, err),
						}
					}
					ors = append(ors, value)
				}

				ands = append(ands, ors)
			}

			if p, ok := options.Parameters[k]; ok {
				options.Parameters[k] = append(p, ands...)
			} else {
				options.Parameters[k] = ands
			}
		}
	}

	return options, nil
}

func parseCount(params url.Values, maxCount, defaultCount int) (int, capabilities.FHIRError) {
	if count, ok := params["_count"]; ok {
		if len(count) != 1 {
			return 0, SearchError{
				error: fmt.Errorf("multiple _count parameters"),
			}
		}
		countInt, err := strconv.Atoi(count[0])
		if err != nil {
			return 0, SearchError{
				error: fmt.Errorf("invalid _count parameter: %w", err),
			}
		}
		countInt = min(countInt, maxCount)
		return countInt, nil
	}
	return defaultCount, nil
}

func parseCursor(params url.Values) (search.Cursor, capabilities.FHIRError) {
	if cursorList, ok := params["_cursor"]; ok {
		if len(cursorList) != 1 {
			return "", SearchError{
				error: fmt.Errorf("multiple _cursor parameters"),
			}
		}
		return search.Cursor(cursorList[0]), nil
	}
	return "", nil
}

func parseSearchValue(typ search.ParameterType, value string, tz *time.Location) (search.Value, error) {
	prefix := parseSearchValuePrefix(typ, value)
	if prefix != "" {
		// all prefixes have a width of 2
		value = value[2:]
	}

	valueAny, datePrecision, err := parseSearchValueAny(typ, value, tz)
	if err != nil {
		return search.Value{}, err
	}

	return search.Value{Prefix: prefix, DatePrecision: datePrecision, Value: valueAny}, nil
}

func parseSearchValuePrefix(typ search.ParameterType, value string) string {
	// all prefixes have a width of 2
	if len(value) < 2 {
		return ""
	}

	// only number, date and quantity can have prefixes
	if !slices.Contains([]search.ParameterType{search.Number, search.Date, search.Quantity}, typ) {
		return ""
	}

	if !slices.Contains(search.KnownPrefixes, value[:2]) {
		return ""
	}

	return value[:2]
}

func parseSearchValueAny(typ search.ParameterType, value string, tz *time.Location) (any, search.DatePrecision, error) {
	switch typ {
	case search.Date:
		return parseDate(value, tz)
	default:
		return value, "", nil
	}
}

func parseDate(value string, tz *time.Location) (time.Time, search.DatePrecision, error) {
	date, err := time.ParseInLocation(search.OnlyYearFormat, value, tz)
	if err == nil {
		return date, search.YearPrecision, nil
	}
	date, err = time.ParseInLocation(search.UpToMonthFormat, value, tz)
	if err == nil {
		return date, search.MonthPrecision, nil
	}
	date, err = time.ParseInLocation(search.UpToDayFormat, value, tz)
	if err == nil {
		return date, search.DayPrecision, nil
	}
	date, err = time.ParseInLocation(search.HourMinuteFormat, value, tz)
	if err == nil {
		return date, search.HourMinutePrecision, nil
	}
	date, err = time.ParseInLocation(search.FullTimeFormat, value, tz)
	if err == nil {
		return date, search.FullTimePrecision, nil
	}
	return time.Time{}, "", err
}
