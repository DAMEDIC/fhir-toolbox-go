package capabilities

import (
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/utils"
)

// FHIRError is an error that can be returned by an operation.
//
// This custom error interface allows returning a HTTP StatusCode and a FHIR OperationOutcome.
type FHIRError interface {
	error
	StatusCode() int
	OperationOutcome() model.Resource
}

// NotImplementedError means that the requested Interaction is not supported on requested ResourceType.
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
			{
				Severity:    basic.Code{Value: utils.Ptr("fatal")},
				Code:        basic.Code{Value: utils.Ptr("not-supported")},
				Diagnostics: &basic.String{Value: utils.Ptr(e.Error())},
			},
		},
	}
}

// UnknownResourceError means that the requested ResourceType is not supported.
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
			{
				Severity:    basic.Code{Value: utils.Ptr("fatal")},
				Code:        basic.Code{Value: utils.Ptr("not-supported")},
				Diagnostics: &basic.String{Value: utils.Ptr(e.Error())},
			},
		},
	}
}

// InvalidResourceError means that the requested ResourceType is not a valid FHIR resource.
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
			{
				Severity:    basic.Code{Value: utils.Ptr("fatal")},
				Code:        basic.Code{Value: utils.Ptr("processing")},
				Diagnostics: &basic.String{Value: utils.Ptr(e.Error())},
			},
		},
	}
}

// NotFoundError means that the ResourceType with requested ID was not found.
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
			{
				Severity:    basic.Code{Value: utils.Ptr("error")},
				Code:        basic.Code{Value: utils.Ptr("not-found")},
				Diagnostics: &basic.String{Value: utils.Ptr(e.Error())},
			},
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
			{
				Severity:    basic.Code{Value: utils.Ptr("fatal")},
				Code:        basic.Code{Value: utils.Ptr("processing")},
				Diagnostics: &basic.String{Value: utils.Ptr(e.Error())},
			},
		},
	}
}
