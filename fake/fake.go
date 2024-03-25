package fake

import (
	"encoding/json"
	"fhir-toolbox/model/gen/r4"
	"fhir-toolbox/utils"
)

type Fake struct{}

func (c *Fake) GetPatient() *r4.Patient {
	return &r4.Patient{
		Id: &r4.Id{
			Value: utils.Ptr("1234"),
		}}
}

func (c *Fake) GetObservtion() *r4.Observation {
	var o r4.Observation
	err := json.Unmarshal([]byte(
		`{
		  "resourceType": "Observation",
		  "id": "1234",
		  "contained": [
		    {
		      "resourceType": "Patient",
		      "id": "1234"
		    }
		  ],
		  "status": null,
		  "code": {},
		  "valueInteger": 28,
		  "_valueInteger": {
		    "extension": [
		      {
		        "url": "xxx"
		      }
		    ]
		  }
}`), &o)
	if err != nil {
		panic(err)
	}
	return &o
}
