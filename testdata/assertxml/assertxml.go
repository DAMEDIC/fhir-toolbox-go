package assertxml

import (
	"bytes"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Equal(t *testing.T, expected, actual string) {
	assert.Equal(t, format(expected), format(actual))
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
