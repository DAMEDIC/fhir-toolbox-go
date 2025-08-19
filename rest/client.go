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
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4b"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r5"
)

// Compile-time interface compliance checks
var (
	_ capabilities.GenericCapabilities = (*client)(nil)
	_ capabilities.GenericCreate       = (*client)(nil)
	_ capabilities.GenericRead         = (*client)(nil)
	_ capabilities.GenericUpdate       = (*client)(nil)
	_ capabilities.GenericDelete       = (*client)(nil)
	_ capabilities.GenericSearch       = (*client)(nil)
)

// client implements a FHIR REST client that can perform all generic FHIR operations.
// This is private - use the concrete R4, R4B, R5 clients instead.
type client struct {
	baseURL    *url.URL
	httpClient *http.Client
	release    model.Release
}

// clientR4 is a FHIR R4 REST client that provides both generic and concrete APIs.
type clientR4 struct {
	*client
	capabilitiesR4.Concrete
}

// clientR4B is a FHIR R4B REST client that provides both generic and concrete APIs.
type clientR4B struct {
	*client
	capabilitiesR4B.Concrete
}

// clientR5 is a FHIR R5 REST client that provides both generic and concrete APIs.
type clientR5 struct {
	*client
	capabilitiesR5.Concrete
}

// NewClientR4 creates a new FHIR R4 REST client with the given base URL.
func NewClientR4(baseURL string, httpClient *http.Client) (*clientR4, error) {
	c, err := newClientWithRelease(baseURL, httpClient, model.R4{})
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
	c, err := newClientWithRelease(baseURL, httpClient, model.R4B{})
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
	c, err := newClientWithRelease(baseURL, httpClient, model.R5{})
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
func newClientWithRelease(baseURL string, httpClient *http.Client, release model.Release) (*client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	return &client{
		baseURL:    u,
		httpClient: httpClient,
		release:    release,
	}, nil
}

// CapabilityStatement retrieves the CapabilityStatement from the server's metadata endpoint.
func (c *client) CapabilityStatement(ctx context.Context) (basic.CapabilityStatement, error) {
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
func (c *client) Create(ctx context.Context, resource model.Resource) (model.Resource, error) {
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

	return c.parseResourceResponse(resp)
}

// Read retrieves a resource by type and ID.
func (c *client) Read(ctx context.Context, resourceType, id string) (model.Resource, error) {
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

	return c.parseResourceResponse(resp)
}

// Update updates an existing resource.
func (c *client) Update(ctx context.Context, resource model.Resource) (update.Result[model.Resource], error) {
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

	updatedResource, err := c.parseResourceResponse(resp)
	if err != nil {
		return update.Result[model.Resource]{}, fmt.Errorf("parse response: %w", err)
	}

	return update.Result[model.Resource]{
		Resource: updatedResource,
		Created:  created,
	}, nil
}

// Delete deletes a resource by type and ID.
func (c *client) Delete(ctx context.Context, resourceType, id string) error {
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
func (c *client) Search(ctx context.Context, resourceType string, parameters search.Parameters, options search.Options) (search.Result[model.Resource], error) {
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

// parseSearchResponse parses a search bundle response using the client's FHIR version.
func (c *client) parseSearchResponse(resp *http.Response) (search.Result[model.Resource], error) {
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return search.Result[model.Resource]{}, fmt.Errorf("read response body: %w", err)
	}

	// Parse the bundle using the client's FHIR release
	switch c.release.(type) {
	case model.R4:
		var bundle r4.Bundle
		if err := json.Unmarshal(body, &bundle); err != nil {
			return search.Result[model.Resource]{}, fmt.Errorf("parse R4 bundle: %w", err)
		}
		return c.parseR4SearchBundle(bundle)
	case model.R4B:
		var bundle r4b.Bundle
		if err := json.Unmarshal(body, &bundle); err != nil {
			return search.Result[model.Resource]{}, fmt.Errorf("parse R4B bundle: %w", err)
		}
		return c.parseR4BSearchBundle(bundle)
	case model.R5:
		var bundle r5.Bundle
		if err := json.Unmarshal(body, &bundle); err != nil {
			return search.Result[model.Resource]{}, fmt.Errorf("parse R5 bundle: %w", err)
		}
		return c.parseR5SearchBundle(bundle)
	}

	return search.Result[model.Resource]{}, fmt.Errorf("unsupported FHIR release: %s", c.release)
}

// parseResourceResponse parses a resource from an HTTP response using the client's FHIR version.
func (c *client) parseResourceResponse(resp *http.Response) (model.Resource, error) {
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	// Parse the resource using the client's FHIR release
	switch c.release.(type) {
	case model.R4:
		if resource := tryParseR4Resource(body); resource != nil {
			return resource, nil
		}
	case model.R4B:
		if resource := tryParseR4BResource(body); resource != nil {
			return resource, nil
		}
	case model.R5:
		if resource := tryParseR5Resource(body); resource != nil {
			return resource, nil
		}
	}

	// Peek at the resourceType field for error message
	var peek struct {
		ResourceType string `json:"resourceType"`
	}
	json.Unmarshal(body, &peek)

	return nil, fmt.Errorf("unable to parse resource of type %s for FHIR %s", peek.ResourceType, c.release)
}

// tryParseR4Resource attempts to parse a resource as R4.
func tryParseR4Resource(data []byte) model.Resource {
	var contained r4.ContainedResource
	if err := json.Unmarshal(data, &contained); err == nil && contained.Resource != nil {
		return contained.Resource
	}
	return nil
}

// tryParseR4BResource attempts to parse a resource as R4B.
func tryParseR4BResource(data []byte) model.Resource {
	var contained r4b.ContainedResource
	if err := json.Unmarshal(data, &contained); err == nil && contained.Resource != nil {
		return contained.Resource
	}
	return nil
}

// tryParseR5Resource attempts to parse a resource as R5.
func tryParseR5Resource(data []byte) model.Resource {
	var contained r5.ContainedResource
	if err := json.Unmarshal(data, &contained); err == nil && contained.Resource != nil {
		return contained.Resource
	}
	return nil
}

// parseR4SearchBundle parses an R4 Bundle into a search result.
func (c *client) parseR4SearchBundle(bundle r4.Bundle) (search.Result[model.Resource], error) {
	result := search.Result[model.Resource]{
		Resources: make([]model.Resource, 0, len(bundle.Entry)),
		Included:  make([]model.Resource, 0),
	}

	// Parse resources from bundle entries
	for _, entry := range bundle.Entry {
		if entry.Resource != nil {
			result.Resources = append(result.Resources, entry.Resource)
		}
	}

	// Extract next cursor from bundle links
	for _, link := range bundle.Link {
		if link.Relation.Value != nil && *link.Relation.Value == "next" {
			if link.Url.Value != nil {
				result.Next = search.Cursor(*link.Url.Value)
				break
			}
		}
	}

	return result, nil
}

// parseR4BSearchBundle parses an R4B Bundle into a search result.
func (c *client) parseR4BSearchBundle(bundle r4b.Bundle) (search.Result[model.Resource], error) {
	result := search.Result[model.Resource]{
		Resources: make([]model.Resource, 0, len(bundle.Entry)),
		Included:  make([]model.Resource, 0),
	}

	// Parse resources from bundle entries
	for _, entry := range bundle.Entry {
		if entry.Resource != nil {
			result.Resources = append(result.Resources, entry.Resource)
		}
	}

	// Extract next cursor from bundle links
	for _, link := range bundle.Link {
		if link.Relation.Value != nil && *link.Relation.Value == "next" {
			if link.Url.Value != nil {
				result.Next = search.Cursor(*link.Url.Value)
				break
			}
		}
	}

	return result, nil
}

// parseR5SearchBundle parses an R5 Bundle into a search result.
func (c *client) parseR5SearchBundle(bundle r5.Bundle) (search.Result[model.Resource], error) {
	result := search.Result[model.Resource]{
		Resources: make([]model.Resource, 0, len(bundle.Entry)),
		Included:  make([]model.Resource, 0),
	}

	// Parse resources from bundle entries
	for _, entry := range bundle.Entry {
		if entry.Resource != nil {
			result.Resources = append(result.Resources, entry.Resource)
		}
	}

	// Extract next cursor from bundle links
	for _, link := range bundle.Link {
		if link.Relation.Value != nil && *link.Relation.Value == "next" {
			if link.Url.Value != nil {
				result.Next = search.Cursor(*link.Url.Value)
				break
			}
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
	genericResult, err := it.client.Search(it.ctx, "", search.Params{}, search.Options{
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
