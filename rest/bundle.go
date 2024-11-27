package rest

import (
	"fhir-toolbox/capabilities"
	"fhir-toolbox/capabilities/search"
	"fhir-toolbox/model"
	"fhir-toolbox/model/basic"
	"fmt"
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
	return basic.OperationOutcome{
		Issue: []basic.OperationOutcomeIssue{
			{Severity: "fatal", Code: "processing", Diagnostics: e.Error()},
		},
	}
}

// SearchBundle creates a new search bundle from the given resources and parameters.
//
// The REST server uses cursor based pagination.
// If the search results contains a `Next` cursor, a 'next' bundle link entry will be set.
func SearchBundle(
	matchResourceType string,
	result search.Result,
	usedOptions search.Options,
	searchCapabilities search.Capabilities,
	baseURL *url.URL,
) (basic.Bundle, capabilities.FHIRError) {
	entries, err := entries(matchResourceType, result.Resources, baseURL)
	if err != nil {
		return basic.Bundle{}, err
	}

	bundle := basic.Bundle{
		Type: "searchset",
		Link: []basic.BundleLink{
			{
				Relation: "self",
				Url: relationLink(
					matchResourceType,
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

		bundle.Link = append(bundle.Link, basic.BundleLink{
			Relation: "next",
			Url: relationLink(
				matchResourceType,
				nextOptions,
				searchCapabilities,
				baseURL,
			),
		})

	}

	return bundle, nil
}

func entries(matchResourceType string, resources []model.Resource, baseURL *url.URL) ([]basic.BundleEntry, capabilities.FHIRError) {
	entries := make([]basic.BundleEntry, 0, len(resources))

	for _, r := range resources {
		searchMode := "include"
		if matchResourceType == "*" || matchResourceType == r.ResourceType() {
			searchMode = "match"
		}

		entry, err := entry(r, searchMode, baseURL)
		if err != nil {
			return entries, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func entry(resource model.Resource, searchMode string, baseURL *url.URL) (basic.BundleEntry, capabilities.FHIRError) {
	resourceType := resource.ResourceType()
	resourceID, ok := resource.ResourceId()
	if !ok {
		return basic.BundleEntry{}, MissingIdError{ResourceType: resourceType}
	}

	path := strings.Trim(baseURL.Path, "/ ")
	fullURL := url.URL{
		Scheme: baseURL.Scheme,
		Host:   baseURL.Host,
		Path:   fmt.Sprintf("%s/%s/%s", path, resourceType, resourceID),
	}

	return basic.BundleEntry{
		Resource: resource,
		FullUrl:  fullURL.String(),
		Search: basic.BundleEntrySearch{
			Mode: searchMode,
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
) string {
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

	return link.String()
}
