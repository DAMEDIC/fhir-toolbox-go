package testdata

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/utils"
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
	Expression string `xml:",innerxml"`
}

type FHIRPathTestOutput struct {
	OutputType string `xml:"type,attr"`
	Output     string `xml:",innerxml"`
}

func (o FHIRPathTestOutput) Children(name ...string) fhirpath.Collection {
	return nil
}

func (o FHIRPathTestOutput) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	if o.OutputType != "boolean" {
		panic("not a boolean")
	}
	b, err := strconv.ParseBool(o.Output)
	if err != nil {
		panic(err)
	}
	return utils.Ptr(fhirpath.Boolean(b)), nil
}

func (o FHIRPathTestOutput) ToString(explicit bool) (*fhirpath.String, error) {
	if o.OutputType != "string" && o.OutputType != "code" {
		panic("not a string")
	}
	return (*fhirpath.String)(&o.Output), nil
}

func (o FHIRPathTestOutput) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	if o.OutputType != "integer" {
		panic("not an integer")
	}
	i, err := strconv.Atoi(o.Output)
	if err != nil {
		panic(err)
	}
	return utils.Ptr(fhirpath.Integer(i)), nil
}

func (o FHIRPathTestOutput) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	if o.OutputType != "decimal" {
		panic("not a decimal")
	}
	d, _, err := apd.NewFromString(o.Output)
	if err != nil {
		panic(err)
	}
	return &fhirpath.Decimal{Value: d}, nil
}

func (o FHIRPathTestOutput) ToDate(explicit bool) (*fhirpath.Date, error) {
	if o.OutputType != "date" {
		panic("not a date")
	}
	d, err := fhirpath.ParseDate(o.Output)
	if err != nil {
		panic(err)
	}
	return &d, nil
}

func (o FHIRPathTestOutput) ToTime(explicit bool) (*fhirpath.Time, error) {
	if o.OutputType != "time" {
		panic("not a time")
	}
	t, err := fhirpath.ParseTime(o.Output)
	if err != nil {
		panic(err)
	}
	return &t, nil
}

func (o FHIRPathTestOutput) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	if o.OutputType != "dateTime" {
		panic("not a dateTime")
	}
	dt, err := fhirpath.ParseDateTime(o.Output)
	if err != nil {
		panic(err)
	}
	return &dt, nil
}

func (o FHIRPathTestOutput) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	if o.OutputType != "quantity" {
		panic("not a dateTime")
	}
	q, err := fhirpath.ParseQuantity(o.Output)
	if err != nil {
		panic(err)
	}
	return &q, nil
}

func (o FHIRPathTestOutput) Equal(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	e, err := o.toElement()
	if err != nil {
		return false
	}
	if e.Equal(other) {
		return true
	} else if _, ok := e.(*fhirpath.Boolean); ok {
		// the test cases seem to assume implicit conversion to booleans to check presence
		return other != nil
	} else {
		return false
	}
}

func (o FHIRPathTestOutput) Equivalent(other fhirpath.Element, _noReverseTypeConversion ...bool) bool {
	return o.Equal(other)
}

func (o FHIRPathTestOutput) TypeInfo() fhirpath.TypeInfo {
	return fhirpath.SimpleTypeInfo{
		Namespace: "System",
		Name:      strings.Title(o.Output),
	}
}

func (o FHIRPathTestOutput) String() string {
	e, err := o.toElement()
	if err != nil {
		return err.Error()
	}
	return e.String()
}

func (o FHIRPathTestOutput) toElement() (fhirpath.Element, error) {
	switch o.OutputType {
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
	case "quantity":
		return o.ToQuantity(false)
	}
	return nil, fmt.Errorf("invalid type: %s", o.OutputType)
}
