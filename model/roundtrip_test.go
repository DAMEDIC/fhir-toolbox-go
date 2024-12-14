package model_test

import (
	"encoding/json"
	"encoding/xml"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4b"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r5"
	"github.com/DAMEDIC/fhir-toolbox-go/testdata/assertjson"
	"github.com/DAMEDIC/fhir-toolbox-go/testdata/assertxml"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DAMEDIC/fhir-toolbox-go/testdata"
)

var testReleases = []string{"R4", "R4B", "R5"}

func TestRoundtripJSON(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	for _, release := range testReleases {
		t.Run(release, func(t *testing.T) {
			jsonExamples := testdata.GetExamples(release, "json")

			for name, jsonIn := range jsonExamples {
				t.Run(name, func(t *testing.T) {
					t.Parallel()

					switch release {
					case "R4":
						if strings.HasSuffix(name, "-questionnaire.json") {
							t.Skip("R4 questionnaires are missing linkIds")
						}
					case "R4B":
						if strings.HasPrefix(name, "activitydefinition-") {
							t.Skip("R4B activity definitions are lacking a null in event array")
						}
						if strings.HasPrefix(name, "plandefinition-") {
							t.Skip("R4B plan definitions are lacking a null in event array")
						}
						if slices.Contains([]string{
							"codesystem-catalogType.json",
							"valueset-catalogType.json",
							"valuesets.json",
						}, name) {
							t.Skip("R4B codesystem or valueset is lacking required status")
						}
					case "R5":
						if name == "questionnaireresponse-example-f201-lifelines.json" {
							t.Skip("R5 questionnaire response is missing the questionnaire")
						}
					}

					var (
						r   model.Resource
						err error
					)
					switch release {
					case "R4":
						var r4 r4.ContainedResource
						err = json.Unmarshal(jsonIn, &r4)
						r = r4
					case "R4B":
						var r4b r4b.ContainedResource
						err = json.Unmarshal(jsonIn, &r4b)
						r = r4b
					case "R5":
						var r5 r5.ContainedResource
						err = json.Unmarshal(jsonIn, &r5)
						r = r5
					}
					assert.NoError(t, err)

					jsonOut, err := json.Marshal(r)
					assert.NoError(t, err)

					assertjson.Equal(t, string(jsonIn), string(jsonOut))
				})
			}
		})
	}
}

func TestRoundtripXML(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	for _, release := range testReleases {
		t.Run(release, func(t *testing.T) {
			xmlExamples := testdata.GetExamples(release, "xml")

			for name, xmlIn := range xmlExamples {
				t.Run(name, func(t *testing.T) {
					t.Parallel()

					switch release {
					case "R4B":
						if name == "valuesets.xml" {
							t.Skip("R4B valuesets is lacking required status")
						}
					}

					var (
						r   model.Resource
						err error
					)
					switch release {
					case "R4":
						var r4 r4.ContainedResource
						err = xml.Unmarshal(xmlIn, &r4)
						r = r4
					case "R4B":
						var r4b r4b.ContainedResource
						err = xml.Unmarshal(xmlIn, &r4b)
						r = r4b
					case "R5":
						var r5 r5.ContainedResource
						err = xml.Unmarshal(xmlIn, &r5)
						r = r5
					}
					assert.NoError(t, err)

					xmlOut, err := xml.Marshal(r)
					assert.NoError(t, err)

					assertxml.Equal(t, string(xmlIn), xml.Header+string(xmlOut))
				})
			}
		})
	}
}
