package capabilities

import (
	"fhir-toolbox/model"
	"fhir-toolbox/model/basic"
	"fmt"
)

type FHIRError interface {
	error
	StatusCode() int
	OperationOutcome() model.Resource
}

type NotImplementedError struct {
	Interaction  string
	ResourceType string
}

func (e NotImplementedError) Error() string {
	return fmt.Sprintf("%s not implemented for resource type %s", e.Interaction, e.ResourceType)
}

func (e NotImplementedError) StatusCode() int {
	return 404
}

func (e NotImplementedError) OperationOutcome() model.Resource {
	return basic.OperationOutcome{
		Issue: []basic.OperationOutcomeIssue{
			{Severity: "fatal", Code: "not-supported", Diagnostics: e.Error()},
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
	return basic.OperationOutcome{
		Issue: []basic.OperationOutcomeIssue{
			{Severity: "fatal", Code: "not-supported", Diagnostics: e.Error()},
		},
	}
}

type InvalidResourceError struct {
	ResourceType string
}

func (e InvalidResourceError) Error() string {
	return fmt.Sprintf("returned resource is not valid %s", e.ResourceType)
}

func (e InvalidResourceError) StatusCode() int {
	return 404
}

func (e InvalidResourceError) OperationOutcome() model.Resource {
	return basic.OperationOutcome{
		Issue: []basic.OperationOutcomeIssue{
			{Severity: "fatal", Code: "processing", Diagnostics: e.Error()},
		},
	}
}

type NotFoundError struct {
	ResourceType string
	ID           string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s resource with ID %s not found", e.ResourceType, e.ID)
}

func (e NotFoundError) StatusCode() int {
	return 404
}

func (e NotFoundError) OperationOutcome() model.Resource {
	return basic.OperationOutcome{
		Issue: []basic.OperationOutcomeIssue{
			{Severity: "error", Code: "not-found", Diagnostics: e.Error()},
		},
	}
}

// SearchError indicates that a search could not be performed.
type SearchError struct {
	error
}

func NewSearchError(err error) SearchError {
	return SearchError{err}
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
