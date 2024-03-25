package r4

import types "fhir-toolbox/model/gen/r4"

type PatientGet interface {
	GetPatient() types.Patient
}
