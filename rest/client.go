package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	capabilitiesR4 "github.com/DAMEDIC/fhir-toolbox-go/capabilities/gen/r4"
	capabilitiesR4B "github.com/DAMEDIC/fhir-toolbox-go/capabilities/gen/r4b"
	capabilitiesR5 "github.com/DAMEDIC/fhir-toolbox-go/capabilities/gen/r5"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/update"
	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
)

// Compile-time interface compliance checks
var (
	_ capabilities.GenericCapabilities = (*client[model.R4])(nil)
	_ capabilities.GenericCreate       = (*client[model.R4])(nil)
	_ capabilities.GenericRead         = (*client[model.R4])(nil)
	_ capabilities.GenericUpdate       = (*client[model.R4])(nil)
	_ capabilities.GenericDelete       = (*client[model.R4])(nil)
	_ capabilities.GenericSearch       = (*client[model.R4])(nil)
)

// client implements a FHIR REST client that can perform all generic FHIR operations.
// This is private - use the concrete R4, R4B, R5 clients instead.
type client[R model.Release] struct {
	baseURL    *url.URL
	httpClient *http.Client
}

// clientR4 is a FHIR R4 REST client that provides both generic and concrete APIs.
type clientR4 struct {
	*client[model.R4]
	capabilitiesR4.Concrete
}

// clientR4B is a FHIR R4B REST client that provides both generic and concrete APIs.
type clientR4B struct {
	*client[model.R4B]
	capabilitiesR4B.Concrete
}

// clientR5 is a FHIR R5 REST client that provides both generic and concrete APIs.
type clientR5 struct {
	*client[model.R5]
	capabilitiesR5.Concrete
}

// NewClientR4 creates a new FHIR R4 REST client with the given base URL.
func NewClientR4(baseURL string, httpClient *http.Client) (*clientR4, error) {
	c, err := newClientWithRelease[model.R4](baseURL, httpClient, model.R4{})
	if err != nil {
		return nil, err
	}
	r4Client := &clientR4{
		client: c,
		Concrete: capabilitiesR4.Concrete{
			Generic: c,
		},
	}
	return r4Client, nil
}

// NewClientR4B creates a new FHIR R4B REST client with the given base URL.
func NewClientR4B(baseURL string, httpClient *http.Client) (*clientR4B, error) {
	c, err := newClientWithRelease[model.R4B](baseURL, httpClient, model.R4B{})
	if err != nil {
		return nil, err
	}
	r4bClient := &clientR4B{
		client: c,
		Concrete: capabilitiesR4B.Concrete{
			Generic: c,
		},
	}
	return r4bClient, nil
}

// NewClientR5 creates a new FHIR R5 REST client with the given base URL.
func NewClientR5(baseURL string, httpClient *http.Client) (*clientR5, error) {
	c, err := newClientWithRelease[model.R5](baseURL, httpClient, model.R5{})
	if err != nil {
		return nil, err
	}
	r5Client := &clientR5{
		client: c,
		Concrete: capabilitiesR5.Concrete{
			Generic: c,
		},
	}
	return r5Client, nil
}

// newClientWithRelease creates a new FHIR REST client with the specified release.
func newClientWithRelease[R model.Release](baseURL string, httpClient *http.Client, release model.Release) (*client[R], error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	return &client[R]{
		baseURL:    u,
		httpClient: httpClient,
	}, nil
}

// CapabilityStatement retrieves the CapabilityStatement from the server's metadata endpoint.
func (c *client[R]) CapabilityStatement(ctx context.Context) (basic.CapabilityStatement, error) {
	u := c.baseURL.JoinPath("metadata")

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return basic.CapabilityStatement{}, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", string(FormatJSON))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return basic.CapabilityStatement{}, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return basic.CapabilityStatement{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var capability basic.CapabilityStatement
	if err := parseJSONResponse(resp, &capability); err != nil {
		return basic.CapabilityStatement{}, fmt.Errorf("parse response: %w", err)
	}

	return capability, nil
}

// Create creates a new resource.
func (c *client[R]) Create(ctx context.Context, resource model.Resource) (model.Resource, error) {
	resourceType := resource.ResourceType()
	u := c.baseURL.JoinPath(resourceType)

	body, err := marshalResource(resource)
	if err != nil {
		return nil, fmt.Errorf("marshal resource: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", string(FormatJSON))
	req.Header.Set("Accept", string(FormatJSON))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return decodeResource[R](resp.Body, FormatJSON)
}

// Read retrieves a resource by type and ID.
func (c *client[R]) Read(ctx context.Context, resourceType, id string) (model.Resource, error) {
	u := c.baseURL.JoinPath(resourceType, id)

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", string(FormatJSON))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return decodeResource[R](resp.Body, FormatJSON)
}

// Update updates an existing resource.
func (c *client[R]) Update(ctx context.Context, resource model.Resource) (update.Result[model.Resource], error) {
	resourceType := resource.ResourceType()
	id, hasID := resource.ResourceId()
	if !hasID {
		return update.Result[model.Resource]{}, fmt.Errorf("resource has no ID")
	}

	u := c.baseURL.JoinPath(resourceType, id)

	body, err := marshalResource(resource)
	if err != nil {
		return update.Result[model.Resource]{}, fmt.Errorf("marshal resource: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", u.String(), body)
	if err != nil {
		return update.Result[model.Resource]{}, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", string(FormatJSON))
	req.Header.Set("Accept", string(FormatJSON))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return update.Result[model.Resource]{}, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	var created bool
	switch resp.StatusCode {
	case http.StatusOK:
		created = false
	case http.StatusCreated:
		created = true
	default:
		return update.Result[model.Resource]{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	updatedResource, err := decodeResource[R](resp.Body, FormatJSON)
	if err != nil {
		return update.Result[model.Resource]{}, fmt.Errorf("parse response: %w", err)
	}

	return update.Result[model.Resource]{
		Resource: updatedResource,
		Created:  created,
	}, nil
}

// Delete deletes a resource by type and ID.
func (c *client[R]) Delete(ctx context.Context, resourceType, id string) error {
	u := c.baseURL.JoinPath(resourceType, id)

	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// Search performs a search operation for the given resource type with the specified options.
func (c *client[R]) Search(ctx context.Context, resourceType string, parameters search.Parameters, options search.Options) (search.Result[model.Resource], error) {
	opts := options

	var u *url.URL

	// If cursor is provided, use it as the complete URL (ignoring other options)
	if opts.Cursor != "" {
		var err error
		u, err = url.Parse(string(opts.Cursor))
		if err != nil {
			return search.Result[model.Resource]{}, fmt.Errorf("invalid cursor URL: %w", err)
		}
	} else {
		// Build URL from base and resource type with search parameters
		u = c.baseURL.JoinPath(resourceType)

		// Add query parameters from search options
		queryString := search.BuildQuery(parameters, opts)
		if queryString != "" {
			u.RawQuery = queryString
		}
	}

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return search.Result[model.Resource]{}, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", string(FormatJSON))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return search.Result[model.Resource]{}, fmt.Errorf("execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return search.Result[model.Resource]{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return c.parseSearchResponse(resp)
}

// Helper functions

// marshalResource marshals a FHIR resource to JSON.
func marshalResource(resource model.Resource) (io.Reader, error) {
	data, err := json.Marshal(resource)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

// parseJSONResponse parses a JSON response into the given value.
func parseJSONResponse(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}

// parseSearchResponse parses a search bundle response using FHIRPath expressions.
func (c *client[R]) parseSearchResponse(resp *http.Response) (search.Result[model.Resource], error) {
	// Decode the bundle as a resource using the generic decode function
	bundle, err := decodeResource[R](resp.Body, FormatJSON)
	if err != nil {
		return search.Result[model.Resource]{}, fmt.Errorf("parse bundle: %w", err)
	}

	return c.parseSearchBundle(bundle)
}

// parseSearchBundle parses a Bundle into a search result using FHIRPath expressions.
func (c *client[R]) parseSearchBundle(bundle model.Resource) (search.Result[model.Resource], error) {
	// Extract all bundle entries using FHIRPath
	entryExpr := fhirpath.MustParse("entry")
	entryResults, err := fhirpath.Evaluate(context.Background(), bundle, entryExpr)
	if err != nil {
		return search.Result[model.Resource]{}, fmt.Errorf("extract bundle entries: %w", err)
	}

	var resources []model.Resource
	var included []model.Resource

	// Process each entry to determine if it's a match or include
	for _, entryElement := range entryResults {
		// Extract the resource from this entry
		resourceExpr := fhirpath.MustParse("resource")
		resourceResults, err := fhirpath.Evaluate(context.Background(), entryElement, resourceExpr)
		if err != nil {
			continue // Skip entries without resources
		}

		if len(resourceResults) == 0 {
			continue // Skip entries without resources
		}

		resource, ok := resourceResults[0].(model.Resource)
		if !ok {
			continue // Skip if not a resource
		}

		// Extract the search mode from this entry
		searchModeExpr := fhirpath.MustParse("search.mode")
		searchModeResults, err := fhirpath.Evaluate(context.Background(), entryElement, searchModeExpr)
		if err != nil || len(searchModeResults) == 0 {
			// Default to match if no search mode specified
			resources = append(resources, resource)
			continue
		}

		// Check the search mode value - handle different FHIR types
		var searchModeStr string
		if len(searchModeResults) > 0 {
			searchModeString, ok, err := searchModeResults[0].ToString(false)
			if err != nil {
				return search.Result[model.Resource]{}, fmt.Errorf("invalid search mode: %w", err)
			}
			if !ok {
				return search.Result[model.Resource]{}, fmt.Errorf("invalid search mode: %v", searchModeResults[0])
			}
			searchModeStr = string(searchModeString)
		}

		// Sort based on search mode
		switch searchModeStr {
		case "include":
			included = append(included, resource)
		case "match":
			fallthrough
		default:
			// Default to match for any other value or if mode is not specified
			resources = append(resources, resource)
		}
	}

	// Extract next link URL using FHIRPath
	nextLinkExpr := fhirpath.MustParse("link.where(relation = 'next').url")
	nextResults, err := fhirpath.Evaluate(context.Background(), bundle, nextLinkExpr)
	if err != nil {
		return search.Result[model.Resource]{}, fmt.Errorf("extract next link: %w", err)
	}

	result := search.Result[model.Resource]{
		Resources: resources,
		Included:  included,
	}

	// Set next cursor if available
	if len(nextResults) > 0 {
		if nextURL, ok := nextResults[0].(fhirpath.String); ok {
			result.Next = search.Cursor(nextURL)
		} else {
			return search.Result[model.Resource]{}, fmt.Errorf("invalid next link: %v", nextResults[0])
		}
	}

	return result, nil
}

// iterator provides pagination functionality for search results.
type iterator[T model.Resource] struct {
	ctx    context.Context
	client capabilities.GenericSearch
	result search.Result[T]
	done   bool
	first  bool
}

// Iterator creates a new iterator for paginating through search results.
// It starts with the provided initial result and uses the client to fetch subsequent pages.
func Iterator[T model.Resource](ctx context.Context, client capabilities.GenericSearch, initialResult search.Result[T]) *iterator[T] {
	return &iterator[T]{
		ctx:    ctx,
		client: client,
		result: initialResult,
		done:   false,
		first:  true,
	}
}

// Next returns the next page of search results.
// It returns io.EOF when there are no more pages available.
func (it *iterator[T]) Next() (search.Result[T], error) {
	// If we are done, return EOF
	if it.done {
		return search.Result[T]{}, io.EOF
	}

	// For the first call, return the initial result
	if it.first {
		it.first = false

		// If there is no next cursor, mark as done
		if it.result.Next == "" {
			it.done = true
		}

		return it.result, nil
	}

	// If there is no next cursor, we are done
	if it.result.Next == "" {
		it.done = true
		return search.Result[T]{}, io.EOF
	}

	// Fetch the next page using the cursor
	// The cursor contains the full URL, so we use it directly
	genericResult, err := it.client.Search(it.ctx, "", search.GenericParams{}, search.Options{
		Count:  len(it.result.Resources),
		Cursor: it.result.Next,
	})
	if err != nil {
		it.done = true
		return search.Result[T]{}, err
	}

	// Convert the generic result to our specific type
	nextResult := search.Result[T]{
		Resources: make([]T, len(genericResult.Resources)),
		Included:  genericResult.Included,
		Next:      genericResult.Next,
	}

	// Convert each resource
	for i, resource := range genericResult.Resources {
		if typedResource, ok := resource.(T); ok {
			nextResult.Resources[i] = typedResource
		}
	}

	// Update our state with the new result
	it.result = nextResult

	// If the next result has no cursor, we are done after this
	if nextResult.Next == "" {
		it.done = true
	}

	return nextResult, nil
}
