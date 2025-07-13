package rest

import (
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/utils/ptr"
	"net/url"
	"slices"
	"strings"
)

func missingIdError(resourceType string) basic.OperationOutcome {
	return basic.OperationOutcome{
		Issue: []basic.OperationOutcomeIssue{
			{
				Severity:    basic.Code{Value: ptr.To("fatal")},
				Code:        basic.Code{Value: ptr.To("exception")},
				Diagnostics: &basic.String{Value: ptr.To(fmt.Sprintf("missing id for resource of type '%s'", resourceType))},
			},
		},
	}
}

// buildSearchBundle creates a new search bundle from the given resources and parameters.
//
// The REST server uses cursor-based pagination.
// If the search results contain a `Next` cursor, a 'next' bundle link entry will be set.
func buildSearchBundle[R model.Resource](
	resourceType string,
	result search.Result[R],
	usedOptions search.Options,
	capabilityStatement basic.CapabilityStatement,
	resolveSearchParameter func(canonical string) (model.Element, error),
) (basic.Bundle, error) {
	// Extract base URL from CapabilityStatement
	baseURL, err := url.Parse(*capabilityStatement.Implementation.Url.Value)
	if err != nil {
		return basic.Bundle{}, fmt.Errorf("invalid implementation URL in CapabilityStatement: %w", err)
	}
	entries, err := entries(result, baseURL)
	if err != nil {
		return basic.Bundle{}, err
	}

	bundle := basic.Bundle{
		Type: basic.Code{Value: ptr.To("searchset")},
		Link: []basic.BundleLink{
			{
				Relation: basic.String{Value: ptr.To("self")},
				Url: relationLink(
					resourceType,
					usedOptions,
					capabilityStatement,
					resolveSearchParameter,
				),
			},
		},
		Entry: entries,
	}

	if result.Next != "" {
		nextOptions := usedOptions
		nextOptions.Cursor = result.Next

		bundle.Link = append(bundle.Link, basic.BundleLink{
			Relation: basic.String{Value: ptr.To("next")},
			Url: relationLink(
				resourceType,
				nextOptions,
				capabilityStatement,
				resolveSearchParameter,
			),
		})

	}

	return bundle, nil
}

func entries[R model.Resource](result search.Result[R], baseURL *url.URL) ([]basic.BundleEntry, error) {
	entries := make([]basic.BundleEntry, 0, len(result.Resources)+len(result.Included))

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

func entry(resource model.Resource, searchMode string, baseURL *url.URL) (basic.BundleEntry, error) {
	resourceType := resource.ResourceType()
	resourceID, ok := resource.ResourceId()
	if !ok {
		return basic.BundleEntry{}, missingIdError(resourceType)
	}

	path := strings.Trim(baseURL.Path, "/ ")
	fullURL := url.URL{
		Scheme: baseURL.Scheme,
		Host:   baseURL.Host,
		Path:   fmt.Sprintf("%s/%s/%s", path, resourceType, resourceID),
	}

	return basic.BundleEntry{
		Resource: resource,
		FullUrl:  &basic.Uri{Value: ptr.To(fullURL.String())},
		Search: &basic.BundleEntrySearch{
			Mode: &basic.Code{Value: &searchMode},
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
	capabilityStatement basic.CapabilityStatement,
	resolveSearchParameter func(canonical string) (model.Element, error),
) basic.Uri {
	// Extract base URL from CapabilityStatement
	baseURL, err := url.Parse(*capabilityStatement.Implementation.Url.Value)
	if err != nil {
		// In case of URL parsing error, return empty URL
		return basic.Uri{Value: ptr.To("")}
	}
	path := strings.Trim(baseURL.Path, "/ ")
	link := url.URL{
		Scheme: baseURL.Scheme,
		Host:   baseURL.Host,
		Path:   fmt.Sprintf("%s/%s", path, resourceType),
	}

	// remove options supplied by the client, but not used/supported by the backend
	usedOptions := options
	usedOptions.Parameters = make(map[search.ParameterKey]search.AllOf, len(options.Parameters))

	// Build search parameters map from CapabilityStatement
	searchParameters := make(map[string]search.Parameter)
	for _, rest := range capabilityStatement.Rest {
		for _, resource := range rest.Resource {
			if resource.Type.Value != nil && *resource.Type.Value == resourceType {
				for _, searchParam := range resource.SearchParam {
					if searchParam.Name.Value != nil && searchParam.Definition != nil && searchParam.Definition.Value != nil {
						paramName := *searchParam.Name.Value
						canonical := *searchParam.Definition.Value

						// Resolve the SearchParameter resource
						if param, err := resolveSearchParameter(canonical); err == nil {
							searchParameters[paramName] = param
						}
					}
				}
				break
			}
		}
	}

	for key, ands := range options.Parameters {
		p, ok := searchParameters[key.Name]
		if !ok {
			continue
		}

		fhirpathModifiers := p.Children("modifier")
		resolvedModifiers := make([]string, 0, len(fhirpathModifiers))
		for _, e := range fhirpathModifiers {
			m, ok, err := e.ToString(false)
			if !ok || err != nil {
				// should not happen as long as correct Parameter resources are used
				continue
			}
			resolvedModifiers = append(resolvedModifiers, string(m))
		}

		if key.Modifier == "" || len(resolvedModifiers) == 0 || slices.Contains(resolvedModifiers, key.Modifier) {
			usedOptions.Parameters[key] = ands
		}
	}

	link.RawQuery = usedOptions.QueryString()

	return basic.Uri{Value: ptr.To(link.String())}
}
