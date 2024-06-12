package rest_test

import (
	"context"
	"fhir-toolbox/capabilities"
	"fhir-toolbox/model"
	"fhir-toolbox/model/gen/r4"
	"fhir-toolbox/rest"
	"fhir-toolbox/utils"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swaggest/assertjson"
)

type mockBackend struct {
	mockPatients     []r4.Patient
	mockObservations []r4.Observation
}

func (m mockBackend) ReadPatient(ctx context.Context, id string) (r4.Patient, capabilities.FHIRError) {
	if len(m.mockPatients) == 0 {
		return r4.Patient{}, capabilities.NotFoundError{ResourceType: "Patient", ID: id}
	}
	return m.mockPatients[0], nil
}

func TestHandleRead(t *testing.T) {
	var tests = []struct {
		name           string
		resourceType   string
		resourceID     string
		backend        any
		expectedStatus int
		expectedBody   string
	}{
		{"ValidResource", "Patient", "1", mockBackend{mockPatients: []r4.Patient{{Id: &r4.Id{Value: utils.Ptr("1")}}}}, http.StatusOK, `{"resourceType":"Patient","id":"1"}`},
		{"InvalidResourceType", "UnknownType", "1", mockBackend{}, http.StatusNotFound, `{"resourceType":"OperationOutcome","issue":[{"severity":"fatal","code":"not-supported","diagnostics":"unknown resource type UnknownType"}]}`},
		{"InvalidResourceID", "Patient", "unknown", mockBackend{}, http.StatusNotFound, `{"resourceType":"OperationOutcome","issue":[{"severity":"error","code":"not-found","diagnostics":"Patient resource with ID unknown not found"}]}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := rest.NewServer[model.R4](tt.backend, rest.DefaultConfig)
			requestURL := fmt.Sprintf("/%s/%s", tt.resourceType, tt.resourceID)
			req := httptest.NewRequest("GET", requestURL, nil)

			rr := httptest.NewRecorder()
			server.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assertjson.Equal(t, []byte(tt.expectedBody), rr.Body.Bytes())
			assert.Equal(t, "application/fhir+json", rr.Header().Get("Content-Type"))
		})
	}
}

func (m mockBackend) SearchCapabilitiesPatient() capabilities.SearchCapabilities {
	return capabilities.SearchCapabilities{
		Parameters: []string{"_id"},
	}
}

func (m mockBackend) SearchPatient(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError) {
	var resouces []model.Resource
	for _, p := range m.mockPatients {
		resouces = append(resouces, p)
	}
	return resouces, nil
}

func (m mockBackend) SearchCapabilitiesObservation() capabilities.SearchCapabilities {
	return capabilities.SearchCapabilities{
		Parameters: []string{"_id"},
		Includes:   []string{"Observation:patient"},
	}
}

func (m mockBackend) SearchObservation(ctx context.Context, options capabilities.SearchOptions) ([]model.Resource, capabilities.FHIRError) {
	var resouces []model.Resource
	for _, p := range m.mockPatients {
		resouces = append(resouces, p)
	}
	for _, p := range m.mockObservations {
		resouces = append(resouces, p)
	}
	return resouces, nil
}

func TestHandleSearchType(t *testing.T) {
	var tests = []struct {
		name           string
		resourceType   string
		QueryString    string
		config         rest.Config
		backend        any
		expectedStatus int
		expectedBody   string
	}{
		{"ValidBundleSingle", "Patient", "_id=1", rest.DefaultConfig, mockBackend{mockPatients: []r4.Patient{
			{Id: &r4.Id{Value: utils.Ptr("1")}},
		}}, http.StatusOK, `{"resourceType":"Bundle", "type":"searchset","entry":[
				{"fullUrl":"http://example.com/Patient/1","resource":{"resourceType":"Patient","id":"1"},"search":{"mode":"match"}}
			],"link":[{"relation":"self","url":"http://example.com/Patient?_id=1"}]}`},
		{"ValidBundleTwoOr", "Patient", "_id=1,2", rest.DefaultConfig, mockBackend{
			mockPatients: []r4.Patient{
				{Id: &r4.Id{Value: utils.Ptr("1")}},
				{Id: &r4.Id{Value: utils.Ptr("2")}},
			}}, http.StatusOK, `{"resourceType":"Bundle", "type":"searchset","entry":[
				{"fullUrl":"http://example.com/Patient/1","resource":{"resourceType":"Patient","id":"1"},"search":{"mode":"match"}},
				{"fullUrl":"http://example.com/Patient/2","resource":{"resourceType":"Patient","id":"2"},"search":{"mode":"match"}}
			],"link":[{"relation":"self","url":"http://example.com/Patient?_id=1,2"}]}`},
		{"ValidBundleTwoAnd", "Patient", "_id=1&_id=2", rest.DefaultConfig, mockBackend{
			mockPatients: []r4.Patient{
				{Id: &r4.Id{Value: utils.Ptr("1")}},
				{Id: &r4.Id{Value: utils.Ptr("2")}},
			}}, http.StatusOK, `{"resourceType":"Bundle", "type":"searchset","entry":[
				{"fullUrl":"http://example.com/Patient/1","resource":{"resourceType":"Patient","id":"1"},"search":{"mode":"match"}},
				{"fullUrl":"http://example.com/Patient/2","resource":{"resourceType":"Patient","id":"2"},"search":{"mode":"match"}}
			],"link":[{"relation":"self","url":"http://example.com/Patient?_id=1&_id=2"}]}`},
		{"ValidBundleUnknownParam", "Patient", "_id=1&unknown=x", rest.DefaultConfig, mockBackend{
			mockPatients: []r4.Patient{
				{Id: &r4.Id{Value: utils.Ptr("1")}},
			}}, http.StatusOK, `{"resourceType":"Bundle", "type":"searchset","entry":[
				{"fullUrl":"http://example.com/Patient/1","resource":{"resourceType":"Patient","id":"1"},"search":{"mode":"match"}}
			],"link":[{"relation":"self","url":"http://example.com/Patient?_id=1"}]}`},
		{"ValidBundleInclude", "Observation", "_include=Observation:patient&_id=1", rest.DefaultConfig, mockBackend{
			mockPatients:     []r4.Patient{{Id: &r4.Id{Value: utils.Ptr("1")}}},
			mockObservations: []r4.Observation{{Id: &r4.Id{Value: utils.Ptr("1")}, Status: r4.Code{Value: utils.Ptr("final")}}},
		}, http.StatusOK, `{"resourceType":"Bundle", "type":"searchset","entry":[
				{"fullUrl":"http://example.com/Patient/1","resource":{"resourceType":"Patient","id":"1"},"search":{"mode":"include"}},
				{"fullUrl":"http://example.com/Observation/1","resource":{"resourceType":"Observation","id":"1", "status":"final","code":{}},"search":{"mode":"match"}}
			],"link":[{"relation":"self","url":"http://example.com/Observation?_include=Observation:patient&_id=1"}]}`},
		{"ValidBundleWithBasePath", "Patient", "_id=1", rest.Config{Base: "///fhir// / "}, mockBackend{mockPatients: []r4.Patient{
			{Id: &r4.Id{Value: utils.Ptr("1")}},
		}}, http.StatusOK, `{"resourceType":"Bundle", "type":"searchset","entry":[
					{"fullUrl":"http://example.com/fhir/Patient/1","resource":{"resourceType":"Patient","id":"1"},"search":{"mode":"match"}}
				],"link":[{"relation":"self","url":"http://example.com/fhir/Patient?_id=1"}]}`},
		{"InvalidResourceType", "UnknownType", "", rest.DefaultConfig, mockBackend{}, http.StatusNotFound, `{"resourceType":"OperationOutcome","issue":[{"severity":"fatal","code":"not-supported","diagnostics":"unknown resource type UnknownType"}]}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := rest.NewServer[model.R4](tt.backend, tt.config)
			base := strings.Trim(tt.config.Base, "/ ")
			if base != "" {
				base = "/" + base
			}
			requestURL := fmt.Sprintf("%s/%s?%s", base, tt.resourceType, tt.QueryString)
			req := httptest.NewRequest("GET", requestURL, nil)

			rr := httptest.NewRecorder()
			server.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assertjson.Equal(t, []byte(tt.expectedBody), rr.Body.Bytes())
			assert.Equal(t, "application/fhir+json", rr.Header().Get("Content-Type"))
		})
	}
}
