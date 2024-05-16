package bundle

import (
	"fhir-toolbox/capabilities"
	"fhir-toolbox/model"
	"fhir-toolbox/model/basic"
	"fmt"
	"slices"
	"strings"
)

func NewSearchBundle(
	matchResourceType string,
	resources []model.Resource,
	searchCapabilities capabilities.SearchCapabilities,
	usedOptions capabilities.SearchOptions,
	baseURL string,
) (basic.Bundle, capabilities.FHIRError) {
	entries, err := entries(matchResourceType, resources, baseURL)
	if err != nil {
		return basic.Bundle{}, err
	}

	return basic.Bundle{
		Type: "searchset",
		Link: []basic.BundleLink{
			{
				Relation: "self",
				Url:      selfRelationLink(matchResourceType, searchCapabilities, usedOptions, baseURL),
			},
		},
		Entry: entries,
	}, nil
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

func selfRelationLink(resourceType string, searchCapabilities capabilities.SearchCapabilities, usedOptions capabilities.SearchOptions, baseURL string) string {
	link := fmt.Sprintf("%s/%s?", strings.TrimRight(baseURL, "/"), resourceType)

	// only include includes that were actually used
	for _, include := range usedOptions.Includes {
		if slices.Contains(searchCapabilities.Includes, include) {
			link += fmt.Sprintf("_include=%s&", include)
		}
	}

	// only include parameters that were actually used
	for param, ands := range usedOptions.Parameters {
		if slices.Contains(searchCapabilities.Parameters, param) {
			for _, ors := range ands {
				link += fmt.Sprintf("%s=%s&", param, strings.Join(ors, ","))
			}
		}
	}

	// strip the trailing "&" or "?"
	return link[:len(link)-1]
}
