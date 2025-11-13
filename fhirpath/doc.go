// Package fhirpath provides an implementation of the FHIRPath v2.0.0 specification.
//
// FHIRPath is a path-based navigation and extraction language, somewhat like XPath.
// Operations are expressed in terms of the logical content of hierarchical data models,
// and support traversal, selection, and filtering of data. Its design was influenced
// by the needs for path navigation, selection, and formulation of invariants in both
// HL7 Fast Healthcare Interoperability Resources (FHIR) and HL7 Clinical Quality Language (CQL).
//
// # Basic Usage
//
// To use the FHIRPath package, you typically:
//
// 1. Parse a FHIRPath expression
// 2. Evaluate the expression against a FHIR resource
// 3. Process the result
//
// Example:
//
//	// Parse a FHIRPath expression
//	expr, err := fhirpath.Parse("Patient.name.given")
//	if err != nil {
//	    // Handle error
//	}
//
//	// Evaluate the expression against a FHIR resource
//	result, err := fhirpath.Evaluate(r4.Context(), patient, expr)
//	if err != nil {
//	    // Handle error
//	}
//
//	// Process the result
//	fmt.Println(result)
//
// # Decimal Precision
//
// The FHIRPath evaluation engine uses apd.Decimal under the hood.
// Precision of decimal operations can be set by supplying an apd.Context:
//
//	// Setup context
//	ctx := r4.Context()
//	// with defined precision for decimal operations.
//	ctx = fhirpath.WithAPDContext(ctx, apd.BaseContext.WithPrecision(100))
//
// Attention: By default the precision is set to 0.
//
// # Implementation Status
//
// This implementation supports most features of the FHIRPath specification, with the
// exception of full UCUM support (unit equality is asserted) and functions that involve code system resolution.
// For a detailed overview of the implementation status, refer to the FHIRPath section in
// the project root README.md file.
package fhirpath
