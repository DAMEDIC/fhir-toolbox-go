package fhirpath

import (
	"errors"
	parser "fhir-toolbox/fhirpath/parser/gen"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func evalIdentifier(tree parser.IIdentifierContext) (string, error) {
	s := tree.GetText()

	if tree.DELIMITEDIDENTIFIER() != nil {
		return unescape(s[1 : len(s)-1])
	}

	return s, nil
}

var (
	unescapeReplacer = strings.NewReplacer(
		`\'`, `'`,
		`\"`, `""`,
		`\'`, `'`,
		`\'`, `'`,
		`\'`, `'`,
	)
	escapeUnicodeRegex = regexp.MustCompile(`\\u([[:xdigit:]]){4}`)
)

func unescape(s string) (string, error) {
	unescaped := unescapeReplacer.Replace(s)

	var errs []error
	return escapeUnicodeRegex.ReplaceAllStringFunc(unescaped, func(s string) string {
		unquoted, err := strconv.Unquote(fmt.Sprintf("'%s'", s))
		if err != nil {
			errs = append(errs, err)
			return "�"
		}
		return unquoted
	}), errors.Join(errs...)
}
