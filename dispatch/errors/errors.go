package errors

import (
	"fhir-toolbox/capabilities"
	"fhir-toolbox/model"
	"fmt"
)

type NotImplementedError struct {
	Interaction  string
	ResourceType string
}

func (e NotImplementedError) Error() string {
	return fmt.Sprintf("%s%s not implemented for resource type %s", e.Interaction, e.ResourceType, e.ResourceType)
}

func (e NotImplementedError) StatusCode() int {
	return 404
}

func (e NotImplementedError) OperationOutcome() model.Resource {
	return capabilities.BasicOperationOutcome{
		Issue: []capabilities.BasicOperationOutcomeIssue{
			{Severity: "fatal", Code: "not-supported"},
		},
	}
}

type UnknownResourceError struct {
	ResourceType string
}

func (e UnknownResourceError) Error() string {
	return fmt.Sprintf("unknown resource type %s", e.ResourceType)
}

func (e UnknownResourceError) StatusCode() int {
	return 404
}

func (e UnknownResourceError) OperationOutcome() model.Resource {
	return capabilities.BasicOperationOutcome{
		Issue: []capabilities.BasicOperationOutcomeIssue{
			{Severity: "fatal", Code: "not-supported"},
		},
	}
}
