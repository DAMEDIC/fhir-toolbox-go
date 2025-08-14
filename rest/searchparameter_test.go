package rest_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/rest"
	"github.com/DAMEDIC/fhir-toolbox-go/utils/ptr"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// Test SearchParameter fallback capabilities when SearchParameterSearch is NOT implemented
func TestSearchParameterFallbackCapabilities(t *testing.T) {
	// Backend without SearchParameterSearch interface and without other resource search capabilities
	backend := mockBackendMinimal{}

	config := rest.DefaultConfig
	server, err := rest.NewServer[model.R4](backend, config)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	req := httptest.NewRequest("GET", "http://example.com/metadata", nil)
	req.Header.Set("Accept", "application/fhir+json")

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedBody := `{
		"date": "2024-11-28T11:25:27+01:00",
		"fhirVersion": "4.0",
		"format": [
			"xml",
			"json"
		],
		"implementation": {
			"description": "a simple FHIR service built with fhir-toolbox-go",
			"url": "http://example.com"
		},
		"kind": "instance",
		"resourceType": "CapabilityStatement",
		"rest": [
			{
				"mode": "server",
				"resource": [
					{
						"interaction": [
							{
								"code": "read"
							},
							{
								"code": "search-type"
							}
						],
						"searchParam": [
							{
								"definition": "http://example.com/SearchParameter/SearchParameter-id",
								"name": "_id",
								"type": "token"
							}
						],
						"type": "SearchParameter"
					}
				]
			}
		],
		"software": {
			"name": "fhir-toolbox-go"
		},
		"status": "active"
	}`

	assertResponse(t, "application/fhir+json", expectedBody, rr)
}

// Test SearchParameter capabilities when SearchParameterSearch IS implemented
func TestSearchParameterConcreteCapabilities(t *testing.T) {
	// Backend with SearchParameterSearch interface (inherits from minimal, not the one with Patient/Observation)
	backend := mockBackendWithSearchParameterSearchOnly{}

	config := rest.DefaultConfig
	server, err := rest.NewServer[model.R4](backend, config)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	req := httptest.NewRequest("GET", "http://example.com/metadata", nil)
	req.Header.Set("Accept", "application/fhir+json")

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedBody := `{
		"date": "2024-11-28T11:25:27+01:00",
		"fhirVersion": "4.0",
		"format": [
			"xml",
			"json"
		],
		"implementation": {
			"description": "a simple FHIR service built with fhir-toolbox-go",
			"url": "http://example.com"
		},
		"kind": "instance",
		"resourceType": "CapabilityStatement",
		"rest": [
			{
				"mode": "server",
				"resource": [
					{
						"interaction": [
							{
								"code": "search-type"
							}
						],
						"searchParam": [
							{
								"definition": "http://example.com/SearchParameter/SearchParameter-name",
								"name": "name",
								"type": "string"
							}
						],
						"type": "SearchParameter"
					}
				]
			}
		],
		"software": {
			"name": "fhir-toolbox-go"
		},
		"status": "active"
	}`

	assertResponse(t, "application/fhir+json", expectedBody, rr)
}

// Test SearchParameter read fallback when SearchParameterSearch is NOT implemented
func TestSearchParameterReadFallback(t *testing.T) {
	backend := mockBackendWithoutSearchParameterSearch{}

	config := rest.DefaultConfig
	server, err := rest.NewServer[model.R4](backend, config)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	req := httptest.NewRequest("GET", "http://example.com/SearchParameter/SearchParameter-id", nil)
	req.Header.Set("Accept", "application/fhir+json")

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedBody := `{
		"resourceType": "SearchParameter",
		"id": "SearchParameter-id",
		"url": "http://example.com/SearchParameter/SearchParameter-id",
		"name": "_id",
		"status": "active",
		"description": "Logical id of this artifact",
		"code": "_id",
		"base": ["SearchParameter"],
		"type": "token",
		"expression": "SearchParameter.id"
	}`

	assertResponse(t, "application/fhir+json", expectedBody, rr)
}

// Test SearchParameter search fallback when SearchParameterSearch is NOT implemented
func TestSearchParameterSearchFallback(t *testing.T) {
	backend := mockBackendWithoutSearchParameterSearch{}

	config := rest.DefaultConfig
	server, err := rest.NewServer[model.R4](backend, config)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	req := httptest.NewRequest("GET", "http://example.com/SearchParameter?_id=SearchParameter-id", nil)
	req.Header.Set("Accept", "application/fhir+json")

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedBody := `{
		"resourceType": "Bundle",
		"type": "searchset",
		"entry": [
			{
				"fullUrl": "http://example.com/SearchParameter/SearchParameter-id",
				"resource": {
					"resourceType": "SearchParameter",
					"id": "SearchParameter-id",
					"url": "http://example.com/SearchParameter/SearchParameter-id",
					"name": "_id",
					"status": "active",
					"description": "Logical id of this artifact",
					"code": "_id",
					"base": ["SearchParameter"],
					"type": "token",
					"expression": "SearchParameter.id"
				},
				"search": {
					"mode": "match"
				}
			}
		],
		"link": [
			{
				"relation": "self",
				"url": "http://example.com/SearchParameter?_id=SearchParameter-id&_count=500"
			}
		]
	}`

	assertResponse(t, "application/fhir+json", expectedBody, rr)
}

// Test SearchParameter search with empty results when searching for unknown ID
func TestSearchParameterSearchFallbackEmptyResult(t *testing.T) {
	backend := mockBackendWithoutSearchParameterSearch{}

	config := rest.DefaultConfig
	server, err := rest.NewServer[model.R4](backend, config)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	req := httptest.NewRequest("GET", "http://example.com/SearchParameter?_id=unknown-id", nil)
	req.Header.Set("Accept", "application/fhir+json")

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedBody := `{
		"resourceType": "Bundle",
		"type": "searchset",
		"link": [
			{
				"relation": "self",
				"url": "http://example.com/SearchParameter?_id=unknown-id&_count=500"
			}
		]
	}`

	assertResponse(t, "application/fhir+json", expectedBody, rr)
}

// Mock backend with minimal capabilities (only CapabilityBase)
type mockBackendMinimal struct{}

func (m mockBackendMinimal) CapabilityBase(ctx context.Context) (basic.CapabilityStatement, error) {
	return basic.CapabilityStatement{
		Status:      basic.Code{Value: ptr.To("active")},
		Date:        basic.DateTime{Value: ptr.To("2024-11-28T11:25:27+01:00")},
		Kind:        basic.Code{Value: ptr.To("instance")},
		FhirVersion: basic.Code{Value: ptr.To("4.0")},
		Format: []basic.Code{
			{Value: ptr.To("xml")},
			{Value: ptr.To("json")},
		},
		Software: &basic.CapabilityStatementSoftware{
			Name: basic.String{Value: ptr.To("fhir-toolbox-go")},
		},
		Implementation: &basic.CapabilityStatementImplementation{
			Description: basic.String{Value: ptr.To("a simple FHIR service built with fhir-toolbox-go")},
			Url:         &basic.Url{Value: ptr.To("http://example.com")},
		},
	}, nil
}

// Mock backend WITHOUT SearchParameterSearch implementation (fallback scenario)
type mockBackendWithoutSearchParameterSearch struct{}

func (m mockBackendWithoutSearchParameterSearch) CapabilityBase(ctx context.Context) (basic.CapabilityStatement, error) {
	return basic.CapabilityStatement{
		Status:      basic.Code{Value: ptr.To("active")},
		Date:        basic.DateTime{Value: ptr.To("2024-11-28T11:25:27+01:00")},
		Kind:        basic.Code{Value: ptr.To("instance")},
		FhirVersion: basic.Code{Value: ptr.To("4.0")},
		Format: []basic.Code{
			{Value: ptr.To("xml")},
			{Value: ptr.To("json")},
		},
		Software: &basic.CapabilityStatementSoftware{
			Name: basic.String{Value: ptr.To("fhir-toolbox-go")},
		},
		Implementation: &basic.CapabilityStatementImplementation{
			Description: basic.String{Value: ptr.To("a simple FHIR service built with fhir-toolbox-go")},
			Url:         &basic.Url{Value: ptr.To("http://example.com")},
		},
	}, nil
}

// Add Patient search capabilities to generate SearchParameters
func (m mockBackendWithoutSearchParameterSearch) SearchCapabilitiesPatient(ctx context.Context) (r4.SearchCapabilities, error) {
	return r4.SearchCapabilities{
		Parameters: map[string]r4.SearchParameter{
			"_id":  {Type: r4.Code{Value: ptr.To(search.TypeToken)}},
			"date": {Type: r4.Code{Value: ptr.To(search.TypeDate)}},
			"name": {Type: r4.Code{Value: ptr.To(search.TypeString)}},
		},
	}, nil
}

// Add Patient search implementation to make it part of capabilities
func (m mockBackendWithoutSearchParameterSearch) SearchPatient(ctx context.Context, options search.Options) (search.Result[r4.Patient], error) {
	return search.Result[r4.Patient]{}, nil
}

// Add Observation search capabilities to generate SearchParameters
func (m mockBackendWithoutSearchParameterSearch) SearchCapabilitiesObservation(ctx context.Context) (r4.SearchCapabilities, error) {
	return r4.SearchCapabilities{
		Parameters: map[string]r4.SearchParameter{
			"_id": {Type: r4.Code{Value: ptr.To(search.TypeToken)}},
		},
	}, nil
}

// Add Observation search implementation to make it part of capabilities
func (m mockBackendWithoutSearchParameterSearch) SearchObservation(ctx context.Context, options search.Options) (search.Result[r4.Observation], error) {
	return search.Result[r4.Observation]{}, nil
}

// Mock backend WITH SearchParameterSearch implementation (inherits minimal base)
type mockBackendWithSearchParameterSearchOnly struct {
	mockBackendMinimal
}

func (m mockBackendWithSearchParameterSearchOnly) SearchCapabilitiesSearchParameter(ctx context.Context) (r4.SearchCapabilities, error) {
	return r4.SearchCapabilities{
		Parameters: map[string]r4.SearchParameter{
			"name": {Type: r4.Code{Value: ptr.To(search.TypeString)}},
		},
	}, nil
}

func (m mockBackendWithSearchParameterSearchOnly) SearchSearchParameter(ctx context.Context, options search.Options) (search.Result[r4.SearchParameter], error) {
	result := search.Result[r4.SearchParameter]{}

	// Return a mock SearchParameter
	mockSearchParam := r4.SearchParameter{
		Id:          &r4.Id{Value: ptr.To("custom-param")},
		Url:         r4.Uri{Value: ptr.To("http://example.com/SearchParameter/custom-param")},
		Name:        r4.String{Value: ptr.To("example")},
		Status:      r4.Code{Value: ptr.To("active")},
		Description: r4.Markdown{Value: ptr.To("Custom search parameter")},
		Code:        r4.Code{Value: ptr.To("example")},
		Base: []r4.Code{
			{Value: ptr.To("Patient")},
		},
		Type:       r4.Code{Value: ptr.To("string")},
		Expression: &r4.String{Value: ptr.To("Patient.name")},
	}

	result.Resources = []r4.SearchParameter{mockSearchParam}
	return result, nil
}

// Mock backend WITH SearchParameterSearch implementation (concrete scenario with Patient/Observation)
type mockBackendWithSearchParameterSearch struct {
	mockBackendWithoutSearchParameterSearch
}

func (m mockBackendWithSearchParameterSearch) SearchCapabilitiesSearchParameter(ctx context.Context) (r4.SearchCapabilities, error) {
	return r4.SearchCapabilities{
		Parameters: map[string]r4.SearchParameter{
			"name": {Type: r4.Code{Value: ptr.To(search.TypeString)}},
		},
	}, nil
}

func (m mockBackendWithSearchParameterSearch) SearchSearchParameter(ctx context.Context, options search.Options) (search.Result[r4.SearchParameter], error) {
	result := search.Result[r4.SearchParameter]{}

	// Return a mock SearchParameter
	mockSearchParam := r4.SearchParameter{
		Id:          &r4.Id{Value: ptr.To("custom-param")},
		Url:         r4.Uri{Value: ptr.To("http://example.com/SearchParameter/custom-param")},
		Name:        r4.String{Value: ptr.To("example")},
		Status:      r4.Code{Value: ptr.To("active")},
		Description: r4.Markdown{Value: ptr.To("Custom search parameter")},
		Code:        r4.Code{Value: ptr.To("example")},
		Base: []r4.Code{
			{Value: ptr.To("Patient")},
		},
		Type:       r4.Code{Value: ptr.To("string")},
		Expression: &r4.String{Value: ptr.To("Patient.name")},
	}

	result.Resources = []r4.SearchParameter{mockSearchParam}
	return result, nil
}

// Test reading generated SearchParameters from Patient search capabilities
func TestSearchParameterReadGeneratedFromPatient(t *testing.T) {
	backend := mockBackendWithoutSearchParameterSearch{}

	config := rest.DefaultConfig
	server, err := rest.NewServer[model.R4](backend, config)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Test reading Patient-_id SearchParameter
	req := httptest.NewRequest("GET", "http://example.com/SearchParameter/Patient-id", nil)
	req.Header.Set("Accept", "application/fhir+json")

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expectedBody := `{
		"resourceType": "SearchParameter",
		"id": "Patient-id",
		"url": "http://example.com/SearchParameter/Patient-id",
		"name": "_id",
		"status": "active",
		"description": "Search parameter _id for Patient resource",
		"code": "_id",
		"base": ["Patient"],
		"type": "token"
	}`

	assertResponse(t, "application/fhir+json", expectedBody, rr)
}

// Test that all SearchParameter URLs in CapabilityStatement can be resolved
func TestAllCapabilityStatementSearchParameterUrlsResolvable(t *testing.T) {
	backend := mockBackendWithoutSearchParameterSearch{}

	config := rest.DefaultConfig
	server, err := rest.NewServer[model.R4](backend, config)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// First, get the CapabilityStatement
	req := httptest.NewRequest("GET", "http://example.com/metadata", nil)
	req.Header.Set("Accept", "application/fhir+json")

	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// Parse the CapabilityStatement
	var capabilityStatement map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &capabilityStatement); err != nil {
		t.Fatalf("Failed to parse CapabilityStatement JSON: %v", err)
	}

	// Extract all SearchParameter definition URLs from the CapabilityStatement
	searchParamUrls := extractSearchParameterUrls(t, capabilityStatement)

	if len(searchParamUrls) == 0 {
		t.Fatal("No SearchParameter URLs found in CapabilityStatement")
	}

	t.Logf("Found %d SearchParameter URLs in CapabilityStatement", len(searchParamUrls))

	// Test each SearchParameter URL can be resolved
	for _, definitionUrl := range searchParamUrls {
		t.Run(fmt.Sprintf("resolve_%s", getLastPathSegment(definitionUrl)), func(t *testing.T) {
			// Convert the definition URL to a relative path for the test server
			relativeUrl := convertToRelativePath(definitionUrl, "http://example.com")

			req := httptest.NewRequest("GET", relativeUrl, nil)
			req.Header.Set("Accept", "application/fhir+json")

			rr := httptest.NewRecorder()
			server.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("Failed to resolve SearchParameter at %s (converted to %s). Status: %d, Body: %s",
					definitionUrl, relativeUrl, rr.Code, rr.Body.String())
				return
			}

			// Verify it's a valid SearchParameter resource
			var searchParam map[string]interface{}
			if err := json.Unmarshal(rr.Body.Bytes(), &searchParam); err != nil {
				t.Errorf("Failed to parse SearchParameter JSON for %s: %v", definitionUrl, err)
				return
			}

			// Verify it has the expected resourceType
			if resourceType, ok := searchParam["resourceType"].(string); !ok || resourceType != "SearchParameter" {
				t.Errorf("Expected resourceType 'SearchParameter', got %v for URL %s", resourceType, definitionUrl)
				return
			}

			// Verify the URL matches the definition URL
			if urlField, ok := searchParam["url"].(string); !ok || urlField != definitionUrl {
				t.Errorf("SearchParameter URL field (%s) does not match definition URL (%s)", urlField, definitionUrl)
				return
			}

			t.Logf("Successfully resolved SearchParameter: %s", definitionUrl)
		})
	}
}

// Helper function to extract SearchParameter definition URLs from CapabilityStatement
func extractSearchParameterUrls(t *testing.T, capabilityStatement map[string]interface{}) []string {
	var urls []string

	rest, ok := capabilityStatement["rest"].([]interface{})
	if !ok {
		t.Fatal("CapabilityStatement missing 'rest' array")
	}

	for _, restItem := range rest {
		restObj, ok := restItem.(map[string]interface{})
		if !ok {
			continue
		}

		resources, ok := restObj["resource"].([]interface{})
		if !ok {
			continue
		}

		for _, resourceItem := range resources {
			resourceObj, ok := resourceItem.(map[string]interface{})
			if !ok {
				continue
			}

			searchParams, ok := resourceObj["searchParam"].([]interface{})
			if !ok {
				continue
			}

			for _, searchParamItem := range searchParams {
				searchParamObj, ok := searchParamItem.(map[string]interface{})
				if !ok {
					continue
				}

				if definition, ok := searchParamObj["definition"].(string); ok {
					urls = append(urls, definition)
				}
			}
		}
	}

	return urls
}

// Helper function to convert absolute URL to relative path for testing
func convertToRelativePath(absoluteUrl, baseUrl string) string {
	if strings.HasPrefix(absoluteUrl, baseUrl) {
		return absoluteUrl[len(baseUrl):]
	}

	// Parse the URL to get just the path
	if parsedUrl, err := url.Parse(absoluteUrl); err == nil {
		return parsedUrl.Path
	}

	return absoluteUrl
}

// Helper function to get the last path segment from a URL
func getLastPathSegment(urlStr string) string {
	if parsedUrl, err := url.Parse(urlStr); err == nil {
		segments := strings.Split(strings.Trim(parsedUrl.Path, "/"), "/")
		if len(segments) > 0 {
			return segments[len(segments)-1]
		}
	}
	return "unknown"
}
