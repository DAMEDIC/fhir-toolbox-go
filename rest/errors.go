package rest

import (
	"fhir-toolbox/model"
	"fhir-toolbox/model/basic"
	"fmt"
)

type SearchError struct {
	error
}

func (e SearchError) Error() string {
	return fmt.Sprintf("processing search: %v", e.error)
}

func (e SearchError) StatusCode() int {
	return 500
}

func (e SearchError) OperationOutcome() model.Resource {
	return basic.OperationOutcome{
		Issue: []basic.OperationOutcomeIssue{
			{Severity: "fatal", Code: "processing", Diagnostics: e.Error()},
		},
	}
}
