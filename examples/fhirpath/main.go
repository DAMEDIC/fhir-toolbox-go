package main

import (
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	"github.com/DAMEDIC/fhir-toolbox-go/utils"
)

func main() {
	fmt.Println(fhirpath.Evaluate(
		fhirpath.FHIRContext(),
		object{"Patient"}, fhirpath.MustParse("3 is String"),
	))
}

type object struct {
	name string
}

func (o object) Type() fhirpath.TypeInfo {
	return fhirpath.ClassInfo{
		SimpleTypeInfo: fhirpath.SimpleTypeInfo{
			Namespace: "FHIR",
			Name:      o.name,
		},
	}
}

func (object) Member(s string) fhirpath.Collection {
	fmt.Println("get member", s)
	return fhirpath.Collection{object{name: s}}
}

func (object) ToBoolean(explicit bool) (*fhirpath.Boolean, error) {
	return utils.Ptr(fhirpath.Boolean(false)), fmt.Errorf("not a boolean")
}
func (object) ToString(explicit bool) (*fhirpath.String, error) {
	return nil, nil
}
func (object) ToInteger(explicit bool) (*fhirpath.Integer, error) {
	return nil, nil
}
func (object) ToDecimal(explicit bool) (*fhirpath.Decimal, error) {
	return nil, nil
}
func (object) ToDate(explicit bool) (*fhirpath.Date, error) {
	return nil, nil
}
func (object) ToTime(explicit bool) (*fhirpath.Time, error) {
	return nil, nil
}
func (object) ToDateTime(explicit bool) (*fhirpath.DateTime, error) {
	return nil, nil
}
func (object) ToQuantity(explicit bool) (*fhirpath.Quantity, error) {
	return nil, nil
}
