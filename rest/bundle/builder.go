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
	resources []model.Resource,
	usedOptions search.Options,
	searchCapabilities search.Capabilities,
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
				Url: selfRelationLink(
					matchResourceType,
					usedOptions,
					searchCapabilities,
					baseURL,
				),
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

func selfRelationLink(
	resourceType string,
	usedOptions search.Options,
	searchCapabilities search.Capabilities,
	baseURL string,
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

		for _, ors := range ands {
			link += fmt.Sprintf("%s=", name)

			for _, or := range ors {
				link += fmt.Sprintf("%s%s,", or.Prefix, or.Value)
			}

			link = link[:len(link)-1] + "&"
		}

	}

	// strip the trailing "&" or "?"
	return link[:len(link)-1]
}
