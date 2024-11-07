package rest_test

import (
	"context"
	"fhir-toolbox/capabilities"
	"fhir-toolbox/capabilities/search"
	"fhir-toolbox/model"
	"fhir-toolbox/model/gen/r4"
	"fhir-toolbox/rest"
	"fhir-toolbox/testdata/assertxml"
	"fhir-toolbox/utils"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swaggest/assertjson"
)

func TestHandleRead(t *testing.T) {
	var tests = []struct {
		name           string
		format         string
		resourceType   string
		resourceID     string
		backend        any
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid JSON resource",
			format:         "application/fhir+json",
			resourceType:   "Patient",
			resourceID:     "1",
			backend:        mockBackend{mockPatients: []r4.Patient{{Id: &r4.Id{Value: utils.Ptr("1")}}}},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType": "Patient",
				"id": "1"
			}`,
		},
		{
			name:           "valid JSON resource (incomplete format I)",
			format:         "application/json",
			resourceType:   "Patient",
			resourceID:     "1",
			backend:        mockBackend{mockPatients: []r4.Patient{{Id: &r4.Id{Value: utils.Ptr("1")}}}},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType": "Patient",
				"id": "1"
			}`,
		},
		{
			name:           "valid JSON resource (incomplete format II)",
			format:         "text/json",
			resourceType:   "Patient",
			resourceID:     "1",
			backend:        mockBackend{mockPatients: []r4.Patient{{Id: &r4.Id{Value: utils.Ptr("1")}}}},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType": "Patient",
				"id": "1"
			}`,
		},
		{
			name:           "valid JSON resource (incomplete format III)",
			format:         "json",
			resourceType:   "Patient",
			resourceID:     "1",
			backend:        mockBackend{mockPatients: []r4.Patient{{Id: &r4.Id{Value: utils.Ptr("1")}}}},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType": "Patient",
				"id": "1"
			}`,
		},
		{
			name:           "valid XML resource",
			format:         "application/fhir+xml",
			resourceType:   "Patient",
			resourceID:     "1",
			backend:        mockBackend{mockPatients: []r4.Patient{{Id: &r4.Id{Value: utils.Ptr("1")}}}},
			expectedStatus: http.StatusOK,
			expectedBody: `<?xml version="1.0" encoding="UTF-8"?>
				<Patient>
					<id value="1"/>
            	</Patient>`,
		},
		{
			name:           "valid XML resource (incomplete format I)",
			format:         "application/xml",
			resourceType:   "Patient",
			resourceID:     "1",
			backend:        mockBackend{mockPatients: []r4.Patient{{Id: &r4.Id{Value: utils.Ptr("1")}}}},
			expectedStatus: http.StatusOK,
			expectedBody: `<?xml version="1.0" encoding="UTF-8"?>
				<Patient>
					<id value="1"/>
            	</Patient>`,
		},
		{
			name:           "valid XML resource (incomplete format II)",
			format:         "text/xml",
			resourceType:   "Patient",
			resourceID:     "1",
			backend:        mockBackend{mockPatients: []r4.Patient{{Id: &r4.Id{Value: utils.Ptr("1")}}}},
			expectedStatus: http.StatusOK,
			expectedBody: `<?xml version="1.0" encoding="UTF-8"?>
				<Patient>
					<id value="1"/>
            	</Patient>`,
		},
		{
			name:           "valid XML resource (incomplete format III)",
			format:         "xml",
			resourceType:   "Patient",
			resourceID:     "1",
			backend:        mockBackend{mockPatients: []r4.Patient{{Id: &r4.Id{Value: utils.Ptr("1")}}}},
			expectedStatus: http.StatusOK,
			expectedBody: `<?xml version="1.0" encoding="UTF-8"?>
				<Patient>
					<id value="1"/>
            	</Patient>`,
		},
		{
			name:           "invalid resource type",
			format:         "application/fhir+json",
			resourceType:   "UnknownType",
			resourceID:     "1",
			backend:        mockBackend{},
			expectedStatus: http.StatusNotFound,
			expectedBody: `{
				"resourceType": "OperationOutcome",
				"issue": [
					{
						"severity": "fatal",
						"code": "not-supported",
						"diagnostics":"unknown resource type UnknownType"
					}
				]
			}`,
		},
		{
			name:           "invalid resource id",
			format:         "application/fhir+json",
			resourceType:   "Patient",
			resourceID:     "unknown",
			backend:        mockBackend{},
			expectedStatus: http.StatusNotFound,
			expectedBody: `{
			    "resourceType": "OperationOutcome",
			    "issue": [
				  {
					  "severity": "error",
					  "code": "not-found",
					  "diagnostics": "Patient resource with ID unknown not found"
				  }
			    ]
			}`,
		},
		{
			name:           "invalid resource id",
			format:         "application/fhir+json",
			resourceType:   "Patient",
			resourceID:     "unknown",
			backend:        mockBackend{},
			expectedStatus: http.StatusNotFound,
			expectedBody: `{
			  	"resourceType": "OperationOutcome",
			  	"issue": [
					{
						"severity": "error",
					  	"code": "not-found",
					  	"diagnostics": "Patient resource with ID unknown not found"
					}
			  	]
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, err := rest.NewServer[model.R4](tt.backend, rest.DefaultConfig)
			assert.NoError(t, err)
			requestURL := fmt.Sprintf("/%s/%s", tt.resourceType, tt.resourceID)
			req := httptest.NewRequest("GET", requestURL, nil)
			req.Header.Set("Accept", tt.format)

			rr := httptest.NewRecorder()
			server.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			if strings.Contains(tt.format, "json") {
				assert.Equal(t, "application/fhir+json", rr.Header().Get("Content-Type"))
				assertjson.Equal(t, []byte(tt.expectedBody), rr.Body.Bytes())
			} else {
				assert.Equal(t, "application/fhir+xml", rr.Header().Get("Content-Type"))
				assertxml.Equal(t, tt.expectedBody, rr.Body.String())
			}
		})
	}
}

func TestHandleSearch(t *testing.T) {
	var tests = []struct {
		name           string
		resourceType   string
		queryString    string
		backend        any
		expectedStatus int
		expectedBody   string
	}{
		{
			name:         "valid bundle single entry",
			resourceType: "Patient",
			queryString:  "_id=1",
			backend: mockBackend{
				mockPatients: []r4.Patient{
					{Id: &r4.Id{Value: utils.Ptr("1")}},
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType":"Bundle",
				"type":"searchset",
				"entry":[
					{
						"fullUrl":"http://example.com/Patient/1",
						"resource":{
							"resourceType":"Patient",
							"id":"1"
						},
						"search":{
							"mode":"match"
						}
					}
				],
				"link":[
					{
						"relation":"self",
						"url":"http://example.com/Patient?_id=1&_count=500"
					}
				]
			}`,
		},
		{
			name:         "valid bundle two entries with or parameter",
			resourceType: "Patient",
			queryString:  "_id=1,2",
			backend: mockBackend{
				mockPatients: []r4.Patient{
					{Id: &r4.Id{Value: utils.Ptr("1")}},
					{Id: &r4.Id{Value: utils.Ptr("2")}},
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType":"Bundle",
				"type":"searchset",
				"entry":[
					{
						"fullUrl":"http://example.com/Patient/1",
						"resource":{
							"resourceType":"Patient",
							"id":"1"
						},
						"search":{
							"mode":"match"
						}
					},
					{
						"fullUrl":"http://example.com/Patient/2",
						"resource":{
							"resourceType":"Patient",
							"id":"2"
						},
						"search":{
							"mode":"match"
						}
					}
				],
				"link":[
					{
						"relation":"self",
						"url":"http://example.com/Patient?_id=1,2&_count=500"
					}
				]
			}`,
		},
		{
			name:         "valid bundle two entries with and parameter",
			resourceType: "Patient",
			queryString:  "_id=1&_id=2",
			backend: mockBackend{
				mockPatients: []r4.Patient{
					{Id: &r4.Id{Value: utils.Ptr("1")}},
					{Id: &r4.Id{Value: utils.Ptr("2")}},
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType":"Bundle",
				"type":"searchset",
				"entry":[
					{
						"fullUrl":"http://example.com/Patient/1",
						"resource":{
							"resourceType":"Patient",
							"id":"1"
						},
						"search":{
							"mode":"match"
						}
					},
					{
						"fullUrl":"http://example.com/Patient/2",
						"resource":{
							"resourceType":"Patient",
							"id":"2"
						},
						"search":{
							"mode":"match"
						}
					}
				],
				"link":[
					{
						"relation":"self",
						"url":"http://example.com/Patient?_id=1&_id=2&_count=500"
					}
				]
			}`,
		},
		{
			name:         "valid bundle with unknown parameter",
			resourceType: "Patient",
			queryString:  "_id=1&unknown=x",
			backend: mockBackend{
				mockPatients: []r4.Patient{
					{Id: &r4.Id{Value: utils.Ptr("1")}},
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType":"Bundle",
				"type":"searchset",
				"entry":[
					{
						"fullUrl":"http://example.com/Patient/1",
						"resource":{
							"resourceType":"Patient",
							"id":"1"
						},
						"search":{
							"mode":"match"
						}
					}
				],
				"link":[
					{
						"relation":"self",
						"url":"http://example.com/Patient?_id=1&_count=500"
					}
				]
			}`,
		},
		{
			name:         "valid bundle with include parameter",
			resourceType: "Observation",
			queryString:  "_include=Observation:patient&_id=1",
			backend: mockBackend{
				mockPatients: []r4.Patient{
					{Id: &r4.Id{Value: utils.Ptr("1")}},
				},
				mockObservations: []r4.Observation{
					{Id: &r4.Id{Value: utils.Ptr("1")}, Status: r4.Code{Value: utils.Ptr("final")}},
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType":"Bundle",
				"type":"searchset",
				"entry":[
					{
						"fullUrl":"http://example.com/Patient/1",
						"resource":{
							"resourceType":"Patient",
							"id":"1"
						},
						"search":{
							"mode":"include"
						}
					},
					{
						"fullUrl":"http://example.com/Observation/1",
						"resource":{
							"resourceType":"Observation",
							"id":"1",
							"status":"final",
							"code":{}
						},
						"search":{
							"mode":"match"
						}
					}
				],
				"link":[
					{
						"relation":"self",
						"url":"http://example.com/Observation?_include=Observation:patient&_id=1&_count=500"
					}
				]
			}`,
		},
		{
			name:           "invalid ResourceType",
			resourceType:   "UnknownType",
			queryString:    "_id=1",
			backend:        mockBackend{},
			expectedStatus: http.StatusNotFound,
			expectedBody: `{
				"resourceType":"OperationOutcome",
				"issue":[
					{
						"severity":"fatal",
						"code":"not-supported",
						"diagnostics":"unknown resource type UnknownType"
					}
				]
			}`,
		},
		{
			name:         "search with prefix",
			resourceType: "Patient",
			queryString:  "_id=ge1&date=ge2024-06-03",
			backend: mockBackend{
				mockPatients: []r4.Patient{
					{Id: &r4.Id{Value: utils.Ptr("1")}},
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType":"Bundle",
				"type":"searchset",
				"entry":[
					{
						"fullUrl":"http://example.com/Patient/1",
						"resource":{
							"resourceType":"Patient",
							"id":"1"
						},
						"search":{
							"mode":"match"
						}
					}
				],
				"link":[
					{
						"relation":"self",
						"url":"http://example.com/Patient?_id=ge1&date=ge2024-06-03&_count=500"
					}
				]
			}`,
		},
		{
			name:         "search with all supported prefixes",
			resourceType: "Patient",
			queryString:  "eq1=eq1&ne2=ne2&gt3=gt3&lt4=lt4&ge5=ge5&le6=le6&sa7=sa7&eb8=eb8",
			backend: mockBackend{
				mockPatients: []r4.Patient{
					{Id: &r4.Id{Value: utils.Ptr("1")}},
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType":"Bundle",
				"type":"searchset",
				"entry":[
					{
						"fullUrl":"http://example.com/Patient/1",
						"resource":{
							"resourceType":"Patient",
							"id":"1"
						},
						"search":{
							"mode":"match"
						}
					}
				],
				"link":[
					{
						"relation":"self",
						"url":"http://example.com/Patient?eb8=eb8&eq1=eq1&ge5=ge5&gt3=gt3&le6=le6&lt4=lt4&ne2=ne2&sa7=sa7&_count=500"
					}
				]
			}`,
		},
		{
			name:         "search date up to minute",
			resourceType: "Patient",
			queryString:  "date=ge2024-06-03T16:53Z",
			backend: mockBackend{
				mockPatients: []r4.Patient{
					{Id: &r4.Id{Value: utils.Ptr("1")}},
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType":"Bundle",
				"type":"searchset",
				"entry":[
					{
						"fullUrl":"http://example.com/Patient/1",
						"resource":{
							"resourceType":"Patient",
							"id":"1"
						},
						"search":{
							"mode":"match"
						}
					}
				],
				"link":[
					{
						"relation":"self",
						"url":"http://example.com/Patient?date=ge2024-06-03T16%3A53Z&_count=500"
					}
				]
			}`,
		},
		{
			name:         "search date up to second",
			resourceType: "Patient",
			queryString:  "date=ge2024-06-03T16:53:23Z",
			backend: mockBackend{
				mockPatients: []r4.Patient{
					{Id: &r4.Id{Value: utils.Ptr("1")}},
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType":"Bundle",
				"type":"searchset",
				"entry":[
					{
						"fullUrl":"http://example.com/Patient/1",
						"resource":{
							"resourceType":"Patient",
							"id":"1"
						},
						"search":{
							"mode":"match"
						}
					}
				],
				"link":[
					{
						"relation":"self",
						"url":"http://example.com/Patient?date=ge2024-06-03T16%3A53%3A23Z&_count=500"
					}
				]
			}`,
		},
		{
			name:         "search date full precision and timezone",
			resourceType: "Patient",
			queryString:  "date=ge2024-06-03T16:53:24.444%2b02:00",
			backend: mockBackend{
				mockPatients: []r4.Patient{
					{Id: &r4.Id{Value: utils.Ptr("1")}},
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType":"Bundle",
				"type":"searchset",
				"entry":[
					{
						"fullUrl":"http://example.com/Patient/1",
						"resource":{
							"resourceType":"Patient",
							"id":"1"
						},
						"search":{
							"mode":"match"
						}
					}
				],
				"link":[
					{
						"relation":"self",
						"url":"http://example.com/Patient?date=ge2024-06-03T16%3A53%3A24.444%2B02%3A00&_count=500"
					}
				]
			}`,
		},
		{
			name:         "search with count and cursor",
			resourceType: "Patient",
			queryString:  "_count=1&_cursor=2",
			backend: mockBackend{
				mockPatients: []r4.Patient{
					{Id: &r4.Id{Value: utils.Ptr("1")}},
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType":"Bundle",
				"type":"searchset",
				"entry":[
					{
						"fullUrl":"http://example.com/Patient/1",
						"resource":{
							"resourceType":"Patient",
							"id":"1"
						},
						"search":{
							"mode":"match"
						}
					}
				],
				"link":[
					{
						"relation":"self",
						"url":"http://example.com/Patient?_cursor=2&_count=1"
					}
				]
			}`,
		},
		{
			name:         "search with next link",
			resourceType: "Patient",
			queryString:  "_count=1",
			backend: mockBackend{
				nextCursor: "2",
				mockPatients: []r4.Patient{
					{Id: &r4.Id{Value: utils.Ptr("1")}},
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType":"Bundle",
				"type":"searchset",
				"entry":[
					{
						"fullUrl":"http://example.com/Patient/1",
						"resource":{
							"resourceType":"Patient",
							"id":"1"
						},
						"search":{
							"mode":"match"
						}
					}
				],
				"link":[
					{
						"relation":"self",
						"url":"http://example.com/Patient?_count=1"
					},
					{
						"relation":"next",
						"url":"http://example.com/Patient?_cursor=2&_count=1"
					}
				]
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := rest.DefaultConfig
			config.Base = "http://example.com"
			server, err := rest.NewServer[model.R4](tt.backend, config)
			assert.NoError(t, err)
			requestURL := fmt.Sprintf("http://example.com/%s?%s", tt.resourceType, tt.queryString)
			req := httptest.NewRequest("GET", requestURL, nil)
			req = req.WithContext(context.WithValue(req.Context(), "t", t))

			rr := httptest.NewRecorder()
			server.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.Equal(t, "application/fhir+json", rr.Header().Get("Content-Type"))
			assertjson.Equal(t, []byte(tt.expectedBody), rr.Body.Bytes())
		})
	}
}

type mockBackend struct {
	nextCursor       search.Cursor
	mockPatients     []r4.Patient
	mockObservations []r4.Observation
}

func (m mockBackend) ReadPatient(ctx context.Context, id string) (r4.Patient, capabilities.FHIRError) {
	if len(m.mockPatients) == 0 {
		return r4.Patient{}, capabilities.NotFoundError{ResourceType: "Patient", ID: id}
	}
	return m.mockPatients[0], nil
}

func (m mockBackend) SearchCapabilitiesPatient() search.Capabilities {
	return search.Capabilities{
		Parameters: map[string]search.Desc{
			"_id":  {Type: search.Token},
			"date": {Type: search.Date},
			"eq1":  {Type: search.Token},
			"ne2":  {Type: search.Token},
			"gt3":  {Type: search.Token},
			"lt4":  {Type: search.Token},
			"ge5":  {Type: search.Token},
			"le6":  {Type: search.Token},
			"sa7":  {Type: search.Token},
			"eb8":  {Type: search.Token},
		},
	}
}

func (m mockBackend) SearchPatient(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	t := ctx.Value("t").(*testing.T)

	if options.Parameters["_id"] != nil {
		for _, and := range options.Parameters["_id"] {
			for _, or := range and {
				assert.Emptyf(t, or.Prefix, "prefix must be empty for all except number, date, and quantity")
			}
		}
	}

	result := search.Result{}

	for _, p := range m.mockPatients {
		result.Resources = append(result.Resources, p)
	}

	if m.nextCursor != "" {
		result.Next = m.nextCursor
	}

	return result, nil
}

func (m mockBackend) SearchCapabilitiesObservation() search.Capabilities {
	return search.Capabilities{
		Parameters: map[string]search.Desc{
			"_id": {Type: search.Token},
		},
		Includes: []string{"Observation:patient"},
	}
}

func (m mockBackend) SearchObservation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	var resouces []model.Resource
	for _, p := range m.mockPatients {
		resouces = append(resouces, p)
	}
	for _, p := range m.mockObservations {
		resouces = append(resouces, p)
	}
	return search.Result{
		Resources: resouces,
	}, nil
}
