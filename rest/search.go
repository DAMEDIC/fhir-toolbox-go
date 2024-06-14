package rest

import (
	"context"
	"fhir-toolbox/capabilities/search"
	"fhir-toolbox/dispatch"
	"fhir-toolbox/model"
	"fhir-toolbox/rest/bundle"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

func searchType(
	context context.Context,
	dispatch dispatch.Dispatcher,
	backend Backend,
	resourceType string,
	parameters url.Values,
	baseURL string,
) (int, model.Resource) {
	searchCapabilities, err := dispatch.SearchCapabilities(backend, resourceType)

	options := parseSearchOptions(searchCapabilities, parameters)

	resources, err := dispatch.Search(context, backend, resourceType, options)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	bundle, err := bundle.NewSearchBundle(resourceType, resources, options, searchCapabilities, baseURL)
	if err != nil {
		return err.StatusCode(), err.OperationOutcome()
	}

	return http.StatusOK, bundle
}

func parseSearchOptions(searchCapabilities search.Capabilities, params url.Values) search.Options {
	options := search.Options{
		// need to initialize empty map
		Parameters: make(search.Parameters),
	}

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

			ands := make(search.ParameterAndList, 0, len(v))
			for _, ors := range v {
				splitStrings := strings.Split(ors, ",")

				ors := make(search.ParameterOrList, 0, len(splitStrings))
				for _, s := range splitStrings {
					ors = append(ors, parseSearchValue(desc.Type, s))
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

	return options
}

func parseSearchValue(typ search.Type, value string) search.ParameterValue {
	// all prefixes have a width of 2
	if len(value) < 2 {
		return search.ParameterValue{Prefix: "", Value: value}
	}

	// only number, date and quantity can have prefixes
	if !slices.Contains([]search.Type{search.Number, search.Date, search.Quantity}, typ) {
		return search.ParameterValue{Prefix: "", Value: value}
	}

	if !slices.Contains(search.KnownPrefixes, value[:2]) {
		return search.ParameterValue{Prefix: "", Value: value}
	}

	return search.ParameterValue{Prefix: value[:2], Value: value[2:]}
}
