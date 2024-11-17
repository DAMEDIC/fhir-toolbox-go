package model_test

import (
	"encoding/json"
	"encoding/xml"
	"fhir-toolbox/model/gen/r4"
	"fhir-toolbox/testdata/assertjson"
	"fhir-toolbox/testdata/assertxml"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"fhir-toolbox/testdata"
)

func TestRoundtripJSONR4Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonR4Examples := testdata.GetExamples("r4", "json")

	for name, jsonIn := range jsonR4Examples {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if strings.HasSuffix(name, "-questionnaire.json") {
				t.Skip("R4 questionnaires are missing linkIds")
			}

			var r r4.ContainedResource
			err := json.Unmarshal(jsonIn, &r)
			assert.NoError(t, err)

			jsonOut, err := json.Marshal(r)
			assert.NoError(t, err)

			assertjson.Equal(t, string(jsonIn), string(jsonOut))
		})
	}
}

func TestRoundtripXMLR4Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	xmlR4Examples := testdata.GetExamples("r4", "xml")

	for name, xmlIn := range xmlR4Examples {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var r r4.ContainedResource
			err := xml.Unmarshal(xmlIn, &r)
			assert.NoError(t, err)

			xmlOut, err := xml.Marshal(r)
			assert.NoError(t, err)

			assertxml.Equal(t, string(xmlIn), xml.Header+string(xmlOut))
		})
	}
}
