package r4

import types "fhir-toolbox/model/gen/r4"

type ObservationGet interface {
	GetObservation() types.Observation
}
