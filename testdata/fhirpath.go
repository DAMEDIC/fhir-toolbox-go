package testdata

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/cockroachdb/apd/v3"
	"io"
	"log"
	"path"
	"strconv"
	"strings"
)

func GetFHIRPathTests() FHIRPathTests {
	downloadFHRIPathTests()
	filePath := fhirPathTestsFilePath()

	log.Println("opening zip archive...")
	zip, err := zip.OpenReader(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer zip.Close()

	f, err := zip.Open("r4/tests-fhir-r4.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	testsXml, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	testsXml = bytes.ReplaceAll(testsXml, []byte("< "), []byte("&lt; "))
	testsXml = bytes.ReplaceAll(testsXml, []byte("<="), []byte("&lt;="))

	var tests FHIRPathTests
	err = xml.NewDecoder(bytes.NewReader(testsXml)).Decode(&tests)
	if err != nil {
		log.Fatal(err)
	}

	for i, g := range tests.Groups {
		for j, t := range g.Tests {
			input, err := zip.Open(path.Join("r4", "input", t.InputFile))
			if err != nil {
				log.Fatal(err)
			}
			var r r4.ContainedResource
			err = xml.NewDecoder(input).Decode(&r)
			if err != nil {
				log.Fatal(err)
			}
			t.InputResource = r.Resource
			g.Tests[j] = t
		}
		tests.Groups[i] = g
	}

	return tests
}

type FHIRPathTests struct {
	Name        string               `xml:"name,attr"`
	Description string               `xml:"description,attr"`
	Groups      []*FHIRPathTestGroup `xml:"group"`
}

type FHIRPathTestGroup struct {
	Name        string         `xml:"name,attr"`
	Description string         `xml:"description,attr"`
	Tests       []FHIRPathTest `xml:"test"`
}

type FHIRPathTest struct {
	Name          string `xml:"name,attr"`
	Description   string `xml:"description,attr"`
	InputFile     string `xml:"inputfile,attr"`
	InputResource model.Resource
	Predicate     bool                   `xml:"predicate,attr"`
	Invalid       string                 `xml:"invalid,attr"`
	Expression    FHIRPathTestExpression `xml:"expression"`
	Output        []FHIRPathTestOutput   `xml:"output"`
}

func (t FHIRPathTest) OutputCollection() fhirpath.Collection {
	var c fhirpath.Collection
	for _, o := range t.Output {
		c = append(c, o)
	}
	return c
}

type FHIRPathTestExpression struct {
	Invalid    string `xml:"invalid,attr"`
	Expression string `xml:",chardata"`
}

type FHIRPathTestOutput struct {
	Type   string `xml:"type,attr"`
	Output string `xml:",innerxml"`
}

func (o FHIRPathTestOutput) Children(name ...string) fhirpath.Collection {
	return nil
}

func (o FHIRPathTestOutput) ToBoolean(explicit bool) (v fhirpath.Boolean, ok bool, err error) {
	if o.Type != "boolean" {
		panic("not a boolean")
	}
	b, err := strconv.ParseBool(o.Output)
	if err != nil {
		panic(err)
	}
	return fhirpath.Boolean(b), true, nil
}

func (o FHIRPathTestOutput) ToString(explicit bool) (v fhirpath.String, ok bool, err error) {
	if o.Type != "string" && o.Type != "code" {
		panic("not a string")
	}
	return fhirpath.String(o.Output), true, nil
}

func (o FHIRPathTestOutput) ToInteger(explicit bool) (v fhirpath.Integer, ok bool, err error) {
	if o.Type != "integer" {
		panic("not an integer")
	}
	i, err := strconv.Atoi(o.Output)
	if err != nil {
		panic(err)
	}
	return fhirpath.Integer(i), true, nil
}

func (o FHIRPathTestOutput) ToDecimal(explicit bool) (v fhirpath.Decimal, ok bool, err error) {
	if o.Type != "decimal" {
		panic("not a decimal")
	}
	d, _, err := apd.NewFromString(o.Output)
	if err != nil {
		panic(err)
	}
	return fhirpath.Decimal{Value: d}, true, nil
}

func (o FHIRPathTestOutput) ToDate(explicit bool) (v fhirpath.Date, ok bool, err error) {
	if o.Type != "date" {
		panic("not a date")
	}
	d, err := fhirpath.ParseDate(o.Output)
	if err != nil {
		panic(err)
	}
	return d, true, nil
}

func (o FHIRPathTestOutput) ToTime(explicit bool) (v fhirpath.Time, ok bool, err error) {
	if o.Type != "time" {
		panic("not a time")
	}
	t, err := fhirpath.ParseTime(o.Output)
	if err != nil {
		panic(err)
	}
	return t, true, nil
}

func (o FHIRPathTestOutput) ToDateTime(explicit bool) (v fhirpath.DateTime, ok bool, err error) {
	if o.Type != "dateTime" {
		panic("not a dateTime")
	}
	dt, err := fhirpath.ParseDateTime(o.Output)
	if err != nil {
		panic(err)
	}
	return dt, true, nil
}

func (o FHIRPathTestOutput) ToQuantity(explicit bool) (v fhirpath.Quantity, ok bool, err error) {
	if o.Type != "Quantity" {
		panic("not a Quantity")
	}
	q, err := fhirpath.ParseQuantity(o.Output)
	if err != nil {
		panic(err)
	}
	return q, true, nil
}

func (o FHIRPathTestOutput) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) (eq bool, ok bool) {
	e, ok, err := o.toElement()
	if err != nil || !ok {
		return false, true
	}
	eq, ok = e.Equal(other)
	if eq && ok {
		return true, true
	} else if _, ok := e.(*fhirpath.Boolean); ok {
		// the test cases seem to assume implicit conversion to booleans to check presence
		return other != nil, true
	} else {
		return false, true
	}
}

func (o FHIRPathTestOutput) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	eq, _ := o.Equal(other)
	return eq
}

func (o FHIRPathTestOutput) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.SimpleTypeInfo{
		Namespace: "System",
		Name:      strings.Title(o.Output),
	}
}

func (o FHIRPathTestOutput) MarshalJSON() (out []byte, err error) {
	e, _, _ := o.toElement()
	return json.Marshal(e)
}

func (o FHIRPathTestOutput) String() string {
	e, _, _ := o.toElement()
	return e.String()
}

func (o FHIRPathTestOutput) toElement() (v fhirpath.Element, ok bool, err error) {
	switch o.Type {
	case "boolean":
		return o.ToBoolean(false)
	case "string", "code":
		return o.ToString(false)
	case "integer":
		return o.ToInteger(false)
	case "decimal":
		return o.ToDecimal(false)
	case "date":
		return o.ToDate(false)
	case "time":
		return o.ToTime(false)
	case "dateTime":
		return o.ToDateTime(false)
	case "Quantity":
		return o.ToQuantity(false)
	}
	panic(fmt.Sprintf("invalid type: %s", o.Type))
}
