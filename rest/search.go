package rest

import (
	"context"
	"fhir-toolbox/backend"
	"fhir-toolbox/capabilities"
	"fhir-toolbox/dispatch"
	"fhir-toolbox/model"
	"fhir-toolbox/rest/bundle"
	"net/http"
	"net/url"
	"strings"
)

func searchType(
	context context.Context,
	dispatch dispatch.Dispatcher,
	backend backend.Backend,
	resourceType string,
	parameters url.Values,
	baseURL string,
) (int, model.Resource) {
	searchCapabilities, err := dispatch.SearchCapabilities(backend, resourceType)

	options := parseSearchOptions(parameters)

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

func parseSearchOptions(params url.Values) capabilities.SearchOptions {
	options := capabilities.SearchOptions{
		// need to inittialize empty map
		Parameters: make(capabilities.SearchParameters),
	}

	for k, v := range params {
		switch k {
		case "_include":
			options.Includes = v

		default:
			ands := make(capabilities.SearchParameterAnd, 0, len(v))
			for _, ors := range v {
				ands = append(ands, strings.Split(ors, ","))
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
