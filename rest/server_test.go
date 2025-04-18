package rest_test

import (
	"context"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities"
	"github.com/DAMEDIC/fhir-toolbox-go/capabilities/search"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/rest"
	"github.com/DAMEDIC/fhir-toolbox-go/testdata/assertjson"
	"github.com/DAMEDIC/fhir-toolbox-go/testdata/assertxml"
	"github.com/DAMEDIC/fhir-toolbox-go/utils"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestCapabilityStatement(t *testing.T) {
	var tests = []struct {
		name           string
		format         string
		date           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "CapabilityStatement JSON",
			format:         "application/fhir+json",
			date:           "2024-11-28T11:25:27+01:00",
			expectedStatus: http.StatusOK,
			expectedBody: `{
			  "date": "2024-11-28T11:25:27+01:00",
			  "fhirVersion": "4.0",
			  "format": [
				"xml",
				"json"
			  ],
			  "implementation": {
				"description": "a simple FHIR service built with fhir-toolbox-go",
				"url": "http://example.com/metadata"
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
					  "searchInclude": [
						"Observation:patient"
					  ],
					  "searchParam": [
						{
						  "name": "_id",
						  "type": "token"
						}
					  ],
					  "type": "Observation"
					},
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
						  "name": "_id",
						  "type": "token"
						},
						{
						  "name": "date",
						  "type": "date"
						},
						{
						  "name": "eb8",
						  "type": "token"
						},
						{
						  "name": "eq1",
						  "type": "token"
						},
						{
						  "name": "ge5",
						  "type": "token"
						},
						{
						  "name": "gt3",
						  "type": "token"
						},
						{
						  "name": "le6",
						  "type": "token"
						},
						{
						  "name": "lt4",
						  "type": "token"
						},
						{
						  "name": "ne2",
						  "type": "token"
						},
						{
						  "name": "pre",
						  "type": "token"
						},
						{
						  "name": "sa7",
						  "type": "token"
						}
					  ],
					  "type": "Patient"
					}
				  ]
				}
			  ],
			  "software": {
				"name": "fhir-toolbox-go"
			  },
			  "status": "active"
			}`,
		},
		{
			name:           "CapabilityStatement XML",
			format:         "application/fhir+xml",
			date:           "2024-11-28T11:25:27+01:00",
			expectedStatus: http.StatusOK,
			expectedBody: `<?xml version="1.0" encoding="UTF-8"?>
				<CapabilityStatement xmlns='http://hl7.org/fhir'>
				  <status value='active'/>
				  <date value='2024-11-28T11:25:27+01:00'/>
				  <kind value='instance'/>
				  <software>
					<name value='fhir-toolbox-go'/>
				  </software>
				  <implementation>
					<description value='a simple FHIR service built with fhir-toolbox-go'/>
					<url value='http://example.com/metadata'/>
				  </implementation>
				  <fhirVersion value='4.0'/>
				  <format value='xml'/>
				  <format value='json'/>
				  <rest>
					<mode value='server'/>
					<resource>
					  <type value='Observation'/>
					  <interaction>
						<code value='search-type'/>
					  </interaction>
					  <searchInclude value='Observation:patient'/>
					  <searchParam>
						<name value='_id'/>
						<type value='token'/>
					  </searchParam>
					</resource>
					<resource>
					  <type value='Patient'/>
					  <interaction>
						<code value='read'/>
					  </interaction>
					  <interaction>
						<code value='search-type'/>
					  </interaction>
					  <searchParam>
						<name value='_id'/>
						<type value='token'/>
					  </searchParam>
					  <searchParam>
						<name value='date'/>
						<type value='date'/>
					  </searchParam>
					  <searchParam>
						<name value='eb8'/>
						<type value='token'/>
					  </searchParam>
					  <searchParam>
						<name value='eq1'/>
						<type value='token'/>
					  </searchParam>
					  <searchParam>
						<name value='ge5'/>
						<type value='token'/>
					  </searchParam>
					  <searchParam>
						<name value='gt3'/>
						<type value='token'/>
					  </searchParam>
					  <searchParam>
						<name value='le6'/>
						<type value='token'/>
					  </searchParam>
					  <searchParam>
						<name value='lt4'/>
						<type value='token'/>
					  </searchParam>
					  <searchParam>
						<name value='ne2'/>
						<type value='token'/>
					  </searchParam>
					  <searchParam>
						<name value='pre'/>
						<type value='token'/>
					  </searchParam>
					  <searchParam>
						<name value='sa7'/>
						<type value='token'/>
					  </searchParam>
					</resource>
				  </rest>
				</CapabilityStatement>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := rest.DefaultConfig
			config.Base, _ = url.Parse("http://example.com")
			config.Date = tt.date

			server, err := rest.NewServer[model.R4](mockBackend{}, config)
			if err != nil {
				t.Fatalf("Failed to create server: %v", err)
			}

			req := httptest.NewRequest("GET", "http://example.com/metadata", nil)
			req.Header.Set("Accept", tt.format)

			rr := httptest.NewRecorder()
			server.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, rr.Code)
			}

			if strings.Contains(tt.format, "json") {
				contentType := rr.Header().Get("Content-Type")
				if contentType != "application/fhir+json" {
					t.Errorf("Expected Content-Type %s, got %s", "application/fhir+json", contentType)
				}
				assertjson.Equal(t, tt.expectedBody, rr.Body.String())
			} else {
				contentType := rr.Header().Get("Content-Type")
				if contentType != "application/fhir+xml" {
					t.Errorf("Expected Content-Type %s, got %s", "application/fhir+xml", contentType)
				}
				assertxml.Equal(t, tt.expectedBody, rr.Body.String())
			}
		})
	}
}

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
				<Patient xmlns="http://hl7.org/fhir">
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
				<Patient xmlns="http://hl7.org/fhir">
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
				<Patient xmlns="http://hl7.org/fhir">
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
				<Patient xmlns="http://hl7.org/fhir">
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
			config := rest.DefaultConfig
			config.Base, _ = url.Parse("http://example.com")

			server, err := rest.NewServer[model.R4](tt.backend, config)
			if err != nil {
				t.Fatalf("Failed to create server: %v", err)
			}

			requestURL := fmt.Sprintf("http://example.com/%s/%s", tt.resourceType, tt.resourceID)
			req := httptest.NewRequest("GET", requestURL, nil)
			req.Header.Set("Accept", tt.format)

			rr := httptest.NewRecorder()
			server.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, rr.Code)
			}

			if strings.Contains(tt.format, "json") {
				contentType := rr.Header().Get("Content-Type")
				if contentType != "application/fhir+json" {
					t.Errorf("Expected Content-Type %s, got %s", "application/fhir+json", contentType)
				}
				assertjson.Equal(t, tt.expectedBody, rr.Body.String())
			} else {
				contentType := rr.Header().Get("Content-Type")
				if contentType != "application/fhir+xml" {
					t.Errorf("Expected Content-Type %s, got %s", "application/fhir+xml", contentType)
				}
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
						"url":"http://example.com/Patient?_id=1%2C2&_count=500"
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
				mockObservations: []r4.Observation{
					{Id: &r4.Id{Value: utils.Ptr("1")}, Status: r4.Code{Value: utils.Ptr("final")}},
				},
				mockObservationIncludes: []model.Resource{
					r4.Patient{Id: &r4.Id{Value: utils.Ptr("1")}},
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody: `{
				"resourceType":"Bundle",
				"type":"searchset",
				"entry":[
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
					},
					{
						"fullUrl":"http://example.com/Patient/1",
						"resource":{
							"resourceType":"Patient",
							"id":"1"
						},
						"search":{
							"mode":"include"
						}
					}
				],
				"link":[
					{
						"relation":"self",
						"url":"http://example.com/Observation?_id=1&_include=Observation%3Apatient&_count=500"
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
		{
			name:         "search with modifier",
			resourceType: "Patient",
			queryString:  "_count=1&pre:above=x",
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
						"url":"http://example.com/Patient?pre%3Aabove=x&_count=1"
					},
					{
						"relation":"next",
						"url":"http://example.com/Patient?pre%3Aabove=x&_cursor=2&_count=1"
					}
				]
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := rest.DefaultConfig
			config.Base, _ = url.Parse("http://example.com")

			server, err := rest.NewServer[model.R4](tt.backend, config)
			if err != nil {
				t.Fatalf("Failed to create server: %v", err)
			}

			requestURL := fmt.Sprintf("http://example.com/%s?%s", tt.resourceType, tt.queryString)
			req := httptest.NewRequest("GET", requestURL, nil)
			req = req.WithContext(context.WithValue(req.Context(), "t", t))

			rr := httptest.NewRecorder()
			server.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, rr.Code)
			}

			contentType := rr.Header().Get("Content-Type")
			if contentType != "application/fhir+json" {
				t.Errorf("Expected Content-Type %s, got %s", "application/fhir+json", contentType)
			}

			assertjson.Equal(t, tt.expectedBody, rr.Body.String())
		})
	}
}

type mockBackend struct {
	nextCursor              search.Cursor
	mockPatients            []r4.Patient
	mockObservations        []r4.Observation
	mockObservationIncludes []model.Resource
}

func (m mockBackend) ReadPatient(ctx context.Context, id string) (r4.Patient, capabilities.FHIRError) {
	if len(m.mockPatients) == 0 {
		return r4.Patient{}, capabilities.NotFoundError{ResourceType: "Patient", ID: id}
	}
	return m.mockPatients[0], nil
}

func (m mockBackend) SearchCapabilitiesPatient() search.Capabilities {
	return search.Capabilities{
		Parameters: map[string]search.ParameterDescription{
			"_id":  {Type: search.TypeToken},
			"date": {Type: search.TypeDate},
			"eq1":  {Type: search.TypeToken},
			"ne2":  {Type: search.TypeToken},
			"gt3":  {Type: search.TypeToken},
			"lt4":  {Type: search.TypeToken},
			"ge5":  {Type: search.TypeToken},
			"le6":  {Type: search.TypeToken},
			"sa7":  {Type: search.TypeToken},
			"eb8":  {Type: search.TypeToken},
			"pre":  {Type: search.TypeToken, Modifiers: []search.Modifier{search.ModifierAbove}},
		},
	}
}

func (m mockBackend) SearchPatient(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
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
		Parameters: map[string]search.ParameterDescription{
			"_id": {Type: search.TypeToken},
		},
		Includes: []string{"Observation:patient"},
	}
}

func (m mockBackend) SearchObservation(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {
	var result search.Result

	for _, p := range m.mockObservations {
		result.Resources = append(result.Resources, p)
	}
	for _, p := range m.mockObservationIncludes {
		result.Included = append(result.Included, p)
	}

	return result, nil
}
