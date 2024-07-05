package bundle

import (
	"fhir-toolbox/capabilities"
	"fhir-toolbox/capabilities/search"
	"fhir-toolbox/model"
	"fhir-toolbox/model/basic"
	"fmt"
	"net/url"
	"slices"
	"strings"
)

// NewSearchBundle creates a new search bundle from the given resources and parameters.
func NewSearchBundle(
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
					usedOptions.Cursor,
					usedOptions.Count,
				),
			},
		},
		Entry: entries,
	}

	if result.Next != "" {
		bundle.Link = append(bundle.Link, basic.BundleLink{
			Relation: "next",
			Url: relationLink(
				matchResourceType,
				usedOptions,
				searchCapabilities,
				baseURL,
				result.Next,
				usedOptions.Count,
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

func relationLink(
	resourceType string,
	usedOptions search.Options,
	searchCapabilities search.Capabilities,
	baseURL *url.URL,
	cursor search.Cursor,
	count int,
) string {
	path := strings.Trim(baseURL.Path, "/ ")
	link := url.URL{
		Scheme: baseURL.Scheme,
		Host:   baseURL.Host,
		Path:   fmt.Sprintf("%s/%s", path, resourceType),
	}

	// only include includes that were actually used
	for _, include := range usedOptions.Includes {
		if slices.Contains(searchCapabilities.Includes, include) {
			link.RawQuery += fmt.Sprintf("_include=%s&", include)
		}
	}

	allParams := make([]string, 0, len(usedOptions.Parameters))
	for param := range usedOptions.Parameters {
		allParams = append(allParams, param)
	}
	// sort alphabetically to make the result deterministic
	slices.Sort(allParams)

	// only include parameters that were actually used
	for _, name := range allParams {
		_, isSupportedParameter := searchCapabilities.Parameters[name]
		if !isSupportedParameter {
			continue
		}

		// this must be present because we just got the name form the map
		ands := usedOptions.Parameters[name]

		for _, and := range ands {
			link.RawQuery += fmt.Sprintf("%s=", name)

			for _, or := range and {
				link.RawQuery += fmt.Sprintf("%s,", url.QueryEscape(or.String()))
			}

			link.RawQuery = strings.TrimRight(link.RawQuery, ",")
			link.RawQuery += "&"
		}

	}

	if cursor != "" {
		link.RawQuery += fmt.Sprintf("_cursor=%s&", cursor)
	}

	link.RawQuery += fmt.Sprintf("_count=%d&", count)

	// strip the trailing "&" or "?"
	link.RawQuery = strings.TrimRight(link.RawQuery, "&?")
	return link.String()
}
