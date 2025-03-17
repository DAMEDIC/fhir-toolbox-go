package main

import (
	"context"
	"fmt"
	"github.com/DAMEDIC/fhir-toolbox-go/fhirpath"
	r4 "github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
	"github.com/DAMEDIC/fhir-toolbox-go/utils"
	"github.com/cockroachdb/apd/v3"
)

// evaluateAndPrint is a helper function that evaluates a FHIRPath expression and prints the result
// It accepts an optional custom format string for the output
func evaluateAndPrint(ctx context.Context, target fhirpath.Element, expression string, customFormat ...string) {
	result, err := fhirpath.Evaluate(ctx, target, fhirpath.MustParse(expression))
	if err != nil {
		fmt.Printf("%s => Error: %v\n", expression, err)
	} else {
		if len(customFormat) > 0 {
			fmt.Printf(customFormat[0], result)
		} else {
			fmt.Printf("%s => %v\n", expression, result)
		}
	}
}

// This example demonstrates various features of the FHIRPath package
func main() {
	// Create a sample FHIR Patient resource
	patient := r4.Patient{
		Name: []r4.HumanName{{
			Given:  []r4.String{{Value: utils.Ptr("Donald")}},
			Family: &r4.String{Value: utils.Ptr("Duck")},
		}},
		Active:        &r4.Boolean{Value: utils.Ptr(true)},
		MultipleBirth: r4.Integer{Value: utils.Ptr[int32](3)},
		Telecom: []r4.ContactPoint{
			{
				System: &r4.Code{Value: utils.Ptr("phone")},
				Value:  &r4.String{Value: utils.Ptr("555-123-4567")},
				Use:    &r4.Code{Value: utils.Ptr("home")},
			},
			{
				System: &r4.Code{Value: utils.Ptr("email")},
				Value:  &r4.String{Value: utils.Ptr("donald.duck@example.com")},
				Use:    &r4.Code{Value: utils.Ptr("work")},
			},
		},
	}

	// Basic path navigation
	fmt.Println("# Basic Path Navigation")

	evaluateAndPrint(r4.Context(), patient, "Patient.name.given")
	evaluateAndPrint(r4.Context(), patient, "Patient.name.family")
	evaluateAndPrint(r4.Context(), patient, "Patient.multipleBirth")

	// Equality and comparison
	fmt.Println("\n# Equality and Comparison")

	evaluateAndPrint(r4.Context(), patient, "Patient.name.given = 'Donald'")
	evaluateAndPrint(r4.Context(), patient, "Patient.name.given = Patient.name.family")
	evaluateAndPrint(r4.Context(), patient, "Patient.multipleBirth > 2")

	// Collection operations
	fmt.Println("\n# Collection Operations")

	evaluateAndPrint(r4.Context(), patient, "Patient.telecom.count()")
	evaluateAndPrint(r4.Context(), patient, "Patient.telecom.where(system = 'phone')")
	evaluateAndPrint(r4.Context(), patient, "Patient.telecom.select(system + ': ' + value)")

	// Boolean logic
	fmt.Println("\n# Boolean Logic")

	evaluateAndPrint(r4.Context(), patient, "Patient.active and Patient.name.exists()")
	evaluateAndPrint(r4.Context(), patient, "Patient.active or Patient.deceased")
	evaluateAndPrint(r4.Context(), patient, "Patient.active.not()")

	// String operations
	fmt.Println("\n# String Operations")

	evaluateAndPrint(r4.Context(), patient, "Patient.name.family.startsWith('D')")
	evaluateAndPrint(r4.Context(), patient, "Patient.name.family.contains('uc')")
	evaluateAndPrint(r4.Context(), patient, "Patient.name.family.replace('Duck', 'Mouse')")

	// Mathematical operations
	fmt.Println("\n# Mathematical Operations")

	evaluateAndPrint(r4.Context(), patient, "Patient.multipleBirth + 2")
	evaluateAndPrint(r4.Context(), patient, "Patient.multipleBirth * 3")
	evaluateAndPrint(r4.Context(), patient, "Patient.multipleBirth / 2")

	// Setting decimal precision
	fmt.Println("\n# Decimal Precision")
	ctx := r4.Context()
	ctx = fhirpath.WithAPDContext(ctx, apd.BaseContext.WithPrecision(10))
	evaluateAndPrint(ctx, patient, "10.0 / 3", "10.0 / 3 (precision 10): %v\n")

	// Tracing
	fmt.Println("\n# Tracing")
	ctx = fhirpath.WithTraceLogger(ctx, fhirpath.StdoutTraceLogger{})
	evaluateAndPrint(ctx, patient, "Patient.name.trace('names').given", "Result after tracing: %v\n")
}
