package assertjson

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Equal(t *testing.T, expected, actual string) {
	assert.Equal(t, format(expected), format(actual))
}

func format(input string) string {
	var obj map[string]any
	err := json.Unmarshal([]byte(input), &obj)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	err = enc.Encode(obj)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
