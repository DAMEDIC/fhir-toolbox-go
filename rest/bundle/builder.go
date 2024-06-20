package bundle

import (
	"fhir-toolbox/capabilities"
	"fhir-toolbox/capabilities/search"
	"fhir-toolbox/model"
	"fhir-toolbox/model/basic"
	"fmt"
	"slices"
	"strings"
)

// NewSearchBundle creates a new search bundle from the given resources and parameters.
func NewSearchBundle(
	matchResourceType string,
	result search.Result,
	usedOptions search.Options,
	searchCapabilities search.Capabilities,
	baseURL string,
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

	if result.Next != 0 {
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

func entries(matchResourceType string, resources []model.Resource, baseURL string) ([]basic.BundleEntry, capabilities.FHIRError) {
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

func entry(resource model.Resource, searchMode, baseURL string) (basic.BundleEntry, capabilities.FHIRError) {
	trimmedBaseURL := strings.TrimRight(baseURL, "/")
	resourceType := resource.ResourceType()
	resourceID, ok := resource.ResourceId()
	if !ok {
		return basic.BundleEntry{}, MissingIdError{ResourceType: resourceType}
	}

	return basic.BundleEntry{
		Resource: resource,
		FullUrl:  fmt.Sprintf("%s/%s/%s", trimmedBaseURL, resourceType, resourceID),
		Search: basic.BundleEntrySearch{
			Mode: searchMode,
		},
	}, nil
}

func relationLink(
	resourceType string,
	usedOptions search.Options,
	searchCapabilities search.Capabilities,
	baseURL string,
	cursor search.Cursor,
	count int,
) string {
	link := fmt.Sprintf("%s/%s?", strings.TrimRight(baseURL, "/"), resourceType)

	// only include includes that were actually used
	for _, include := range usedOptions.Includes {
		if slices.Contains(searchCapabilities.Includes, include) {
			link += fmt.Sprintf("_include=%s&", include)
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
			link += fmt.Sprintf("%s=", name)

			for _, or := range and {
				link += fmt.Sprintf("%s,", or)
			}

			link = link[:len(link)-1] + "&"
		}

	}

	if cursor != 0 {
		link += fmt.Sprintf("_cursor=%v&", cursor)
	}

	link += fmt.Sprintf("_count=%d&", count)

	// strip the trailing "&" or "?"
	return link[:len(link)-1]
}
