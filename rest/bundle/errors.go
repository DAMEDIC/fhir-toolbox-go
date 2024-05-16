package bundle

import (
	"fhir-toolbox/model"
	"fhir-toolbox/model/basic"
	"fmt"
)

type MissingIdError struct {
	ResourceType string
}

func (e MissingIdError) Error() string {
	return fmt.Sprintf("missing ID for resource of type %s", e.ResourceType)
}

func (e MissingIdError) StatusCode() int {
	return 500
}

func (e MissingIdError) OperationOutcome() model.Resource {
	return basic.OperationOutcome{
		Issue: []basic.OperationOutcomeIssue{
			{Severity: "fatal", Code: "processing", Diagnostics: e.Error()},
		},
	}
}
