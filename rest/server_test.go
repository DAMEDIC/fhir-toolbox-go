package rest_test

import (
	"context"
	"fhir-toolbox/backend"
	"fhir-toolbox/capabilities"
	"fhir-toolbox/model"
	"fhir-toolbox/model/gen/r4"
	"fhir-toolbox/rest"
	"fhir-toolbox/utils"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swaggest/assertjson"
)

type mockBackend struct {
	mockPatients []r4.Patient
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
		backend        backend.Backend
		expectedStatus int
		expectedBody   string
	}{
		{"ValidResource", "Patient", "1", mockBackend{[]r4.Patient{{Id: &r4.Id{Value: utils.Ptr("1")}}}}, http.StatusOK, `{"resourceType":"Patient","id":"1"}`},
		{"InvalidResourceType", "UnknownType", "1", mockBackend{}, http.StatusNotFound, `{"resourceType":"OperationOutcome","issue":[{"severity":"fatal","code":"not-supported","diagnostics":"unknown resource type UnknownType"}]}`},
		{"InvalidResourceID", "Patient", "unknown", mockBackend{}, http.StatusNotFound, `{"resourceType":"OperationOutcome","issue":[{"severity":"error","code":"not-found","diagnostics":"resource type Patient with ID unknown not found"}]}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := rest.NewServer[model.R4](tt.backend, rest.Config{})
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
