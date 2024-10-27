package model_test

import (
	"encoding/json"
	"fhir-toolbox/model/gen/r4"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swaggest/assertjson"

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

			assertjson.Equal(t, jsonIn, jsonOut)
		})
	}
}
