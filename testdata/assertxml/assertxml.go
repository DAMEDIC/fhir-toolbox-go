package assertxml

import (
	"bytes"
	"encoding/xml"
	"strings"
	"testing"
)

func Equal(t *testing.T, expected, actual string) {
	expectedFormatted := format(expected)
	actualFormatted := format(actual)
	if expectedFormatted != actualFormatted {
		t.Errorf("XML not equal:\nexpected: %s\nactual: %s", expectedFormatted, actualFormatted)
	}
}

func format(input string) string {
	var builder strings.Builder

	decoder := xml.NewDecoder(bytes.NewReader([]byte(input)))
	encoder := xml.NewEncoder(&builder)
	encoder.Indent("", "  ")

	for {
		t, err := decoder.Token()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		switch x := t.(type) {
		case xml.CharData:
			t = xml.CharData(bytes.TrimSpace(x))
		case xml.Comment:
			// skip comments for now
			continue
		}

		if err := encoder.EncodeToken(t); err != nil {
			panic(err)
		}
	}
	err := encoder.Flush()
	if err != nil {
		panic(err)
	}

	return builder.String()
}
