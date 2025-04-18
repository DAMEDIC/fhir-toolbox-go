package assertjson

import (
	"bytes"
	"encoding/json"
	"testing"
)

func Equal(t *testing.T, expected, actual string) {
	expectedFormatted := format(expected)
	actualFormatted := format(actual)
	if expectedFormatted != actualFormatted {
		t.Errorf("JSON not equal:\nexpected: %s\nactual: %s", expectedFormatted, actualFormatted)
	}
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

type Value struct {
	Map    *map[string]Value
	String *string
	// to keep decimal precision when round-tripping
	Raw json.RawMessage
}

func (v Value) MarshalJSON() ([]byte, error) {
	if v.Map != nil {
		return json.Marshal(*v.Map)
	}
	if v.String != nil {
		return json.Marshal(*v.String)
	}
	return json.Marshal(v.Raw)
}

func (v *Value) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &v.Map)
	if err == nil {
		return nil
	}

	// can not handle string as raw message because of escaping
	err = json.Unmarshal(data, &v.String)
	if err == nil {
		return nil
	}

	return json.Unmarshal(data, &v.Raw)
}
