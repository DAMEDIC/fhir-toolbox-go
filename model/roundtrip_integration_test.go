package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swaggest/assertjson"

	"fhir-toolbox/generate/model"
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

			var r model.Resource
			json.Unmarshal(jsonIn, &r)

			jsonOut, err := json.Marshal(r)
			assert.NoError(t, err)

			assertjson.Equal(t, jsonIn, jsonOut)
		})
	}
}
