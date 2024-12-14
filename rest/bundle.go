package rest

import (
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/utils"
	"net/url"
	"strings"
)

// MissingIdError indicates that a bundle entry is missing an id.
type MissingIdError struct {
	ResourceType string
}

func (e MissingIdError) Error() string {
	return fmt.Sprintf("missing ID for resource of type %s", e.ResourceType)
}

func (e MissingIdError) StatusCode() int {
	return 500
}

func (e MissingIdError) OperationOutcome() model.Resource {
	return r4.OperationOutcome{
		Issue: []r4.OperationOutcomeIssue{
			{
				Severity:    r4.Code{Value: utils.Ptr("fatal")},
				Code:        r4.Code{Value: utils.Ptr("processing")},
				Diagnostics: &r4.String{Value: utils.Ptr(e.Error())},
			},
		},
	}
}

// SearchBundle creates a new search bundle from the given resources and parameters.
//
// The REST server uses cursor based pagination.
// If the search results contains a `Next` cursor, a 'next' bundle link entry will be set.
func SearchBundle(
	resourceType string,
	result search.Result,
	usedOptions search.Options,
	searchCapabilities search.Capabilities,
	baseURL *url.URL,
) (r4.Bundle, capabilities.FHIRError) {
	entries, err := entries(result, baseURL)
	if err != nil {
		return r4.Bundle{}, err
	}

	bundle := r4.Bundle{
		Type: r4.Code{Value: utils.Ptr("searchset")},
		Link: []r4.BundleLink{
			{
				Relation: r4.String{Value: utils.Ptr("self")},
				Url: relationLink(
					resourceType,
					usedOptions,
					searchCapabilities,
					baseURL,
				),
			},
		},
		Entry: entries,
	}

	if result.Next != "" {
		nextOptions := usedOptions
		nextOptions.Cursor = result.Next

		bundle.Link = append(bundle.Link, r4.BundleLink{
			Relation: r4.String{Value: utils.Ptr("next")},
			Url: relationLink(
				resourceType,
				nextOptions,
				searchCapabilities,
				baseURL,
			),
		})

	}

	return bundle, nil
}

func entries(result search.Result, baseURL *url.URL) ([]r4.BundleEntry, capabilities.FHIRError) {
	entries := make([]r4.BundleEntry, 0, len(result.Resources)+len(result.Included))

	for _, r := range result.Resources {
		entry, err := entry(r, "match", baseURL)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	for _, r := range result.Included {
		entry, err := entry(r, "include", baseURL)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

func entry(resource model.Resource, searchMode string, baseURL *url.URL) (r4.BundleEntry, capabilities.FHIRError) {
	resourceType := resource.ResourceType()
	resourceID, ok := resource.ResourceId()
	if !ok {
		return r4.BundleEntry{}, MissingIdError{ResourceType: resourceType}
	}

	path := strings.Trim(baseURL.Path, "/ ")
	fullURL := url.URL{
		Scheme: baseURL.Scheme,
		Host:   baseURL.Host,
		Path:   fmt.Sprintf("%s/%s/%s", path, resourceType, resourceID),
	}

	return r4.BundleEntry{
		Resource: &resource,
		FullUrl:  &r4.Uri{Value: utils.Ptr(fullURL.String())},
		Search: &r4.BundleEntrySearch{
			Mode: &r4.Code{Value: &searchMode},
		},
	}, nil
}

// relationLink creates links to be used as `Bundle.link`.
//
// Supplied search parameters that are not supported (not included in the search capabilities)
// are removed.
func relationLink(
	resourceType string,
	options search.Options,
	searchCapabilities search.Capabilities,
	baseURL *url.URL,
) r4.Uri {
	path := strings.Trim(baseURL.Path, "/ ")
	link := url.URL{
		Scheme: baseURL.Scheme,
		Host:   baseURL.Host,
		Path:   fmt.Sprintf("%s/%s", path, resourceType),
	}

	// remove options supplied by the client, but not used/supported by the backend
	usedOptions := options
	usedOptions.Params = make(search.Params, len(options.Params))
	for key, ands := range options.Params {
		if _, ok := searchCapabilities.Params[key]; ok {
			usedOptions.Params[key] = ands
		}
	}

	link.RawQuery = usedOptions.QueryString()

	return r4.Uri{Value: utils.Ptr(link.String())}
}
