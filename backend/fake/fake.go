package fake

import (
	"context"
	"encoding/json"
	"fhir-toolbox/capabilities"
	"fhir-toolbox/model/gen/r4"
)

type Fake struct{}

func (c *Fake) ReadPatient(ctx context.Context, id string) (r4.Patient, capabilities.FHIRError) {
	return r4.Patient{
		Id: &r4.Id{
			Value: &id,
		}}, nil
}

func (c *Fake) ReadObservation(ctx context.Context, id string) (r4.Observation, capabilities.FHIRError) {
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
	return o, nil
}
