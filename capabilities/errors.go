package capabilities

import (
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/basic"
	"github.com/DAMEDIC/fhir-toolbox-go/utils"
	"strings"
)

// FHIRError is an error that can be returned by an operation.
//
// This custom error interface allows returning a HTTP StatusCode and a FHIR OperationOutcome.
type FHIRError interface {
	error
	StatusCode() int
	OperationOutcome() basic.OperationOutcome
}

// ReleaseNotSupported means that the FHIR version is not supported by the server.
type ReleaseNotSupported struct {
	Release string
}

func (e ReleaseNotSupported) Error() string {
	return fmt.Sprintf("FHIR release %s not supported", e.Release)
}

func (e ReleaseNotSupported) StatusCode() int {
	return 404
}

func (e ReleaseNotSupported) OperationOutcome() basic.OperationOutcome {
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

func (e NotImplementedError) OperationOutcome() basic.OperationOutcome {
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

func (e UnknownResourceError) OperationOutcome() basic.OperationOutcome {
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
	return fmt.Sprintf("%s is not a valid FHIR resource", e.ResourceType)
}

func (e InvalidResourceError) StatusCode() int {
	return 404
}

func (e InvalidResourceError) OperationOutcome() basic.OperationOutcome {
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

// NotFoundError means that the ResourceType with the requested ID was not found.
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

func (e NotFoundError) OperationOutcome() basic.OperationOutcome {
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

func (e SearchError) OperationOutcome() basic.OperationOutcome {
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

type joinError struct {
	errors []FHIRError
}

func (e joinError) Error() string {
	// Since Join returns nil if every value in errors is nil,
	// e.errors cannot be empty.
	if len(e.errors) == 1 {
		return e.errors[0].Error()
	}

	var b strings.Builder
	b.WriteString(e.errors[0].Error())
	for _, err := range e.errors[1:] {
		b.WriteRune('\n')
		b.WriteString(err.Error())
	}
	return b.String()
}

func (e joinError) StatusCode() int {
	maxCode := 0
	for _, err := range e.errors {
		if err.StatusCode() > maxCode {
			maxCode = err.StatusCode()
		}
	}
	if maxCode >= 400 && maxCode < 500 {
		return 400
	} else {
		return 500
	}
}

func (e joinError) OperationOutcome() basic.OperationOutcome {
	var issues []basic.OperationOutcomeIssue
	for _, err := range e.errors {
		issues = append(issues, err.OperationOutcome().Issue...)
	}
	return basic.OperationOutcome{
		Issue: issues,
	}
}

func JoinErrors(errs []FHIRError) FHIRError {
	if len(errs) == 0 {
		return nil
	}
	if len(errs) == 1 {
		return errs[0]
	}
	return joinError{errors: errs}
}
