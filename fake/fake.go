package fake

import (
	r4 "fhir-toolbox/model/gen/r4/resources"
	t4 "fhir-toolbox/model/gen/r4/types"
	"fhir-toolbox/utils"
)

type Fake struct{}

func (c *Fake) GetPatient() *r4.Patient {
	return &r4.Patient{
		Id: &t4.Id{
			Value: utils.Ptr("1234"),
		}}
}

func (c *Fake) GetObservtion() *r4.Observation {
	return &r4.Observation{
		Id: &t4.Id{
			Value: utils.Ptr("1234"),
		},
		Value: &t4.Integer{
			Value: utils.Ptr[int32](28),
		},
	}
}
