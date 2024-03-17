package r4

import types "fhir-toolbox/types/r4"

type PatientGet interface {
	GetPatient() types.Patient
}
