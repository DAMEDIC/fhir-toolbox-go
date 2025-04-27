# fhir-toolbox-go

[![Go Reference](https://pkg.go.dev/badge/github.com/DAMEDIC/fhir-toolbox-go.svg)](https://pkg.go.dev/github.com/DAMEDIC/fhir-toolbox-go)

> FHIR® is the registered trademark of HL7® and is used with the permission of HL7®.
> Use of the FHIR trademark does not constitute endorsement of the contents of this repository by HL7®.

This project provides a set of packages for working with the HL7® FHIR® standard in Go.
You only need to implement some interfaces and get a REST implementation out-of-the box.

This includes model types and interfaces modeling capabilities that you can use to build custom FHIR® servers.


> While used in production at DAMEDIC, this project is still in its early days
> and the feature set is quite limit.
> We will add features as we require them. We welcome external contributions.

## Features

- FHIR® model types with JSON and XML (un)marshaling
    - R4, R4B & R5

      use build tags `r4`, `r4b` or `r5` for conditional compilation if you only need runtime support for specific
      versions

    - generated from the FHIR® specification
    ```Go
    var r r4.ContainedResource // container class because json.Unmarshal can not unmarshal directly into interfaces
    err := json.Unmarshal(data, &r)
    ... = r.Resource // the actual resource of type model.Resource
  ```
- Extensible REST API with capabilities modeled as interfaces
    - Capability detection by runtime ~~reflection~~ type assertion (see [Capabilities](#capabilities))
        - alternatively: generic API for building adapters
        - automatic generation fo `CapabilityStatements`
    - Interactions: `create`, `read`, `update`, `delete`, `search` (see [Roadmap](#roadmap) for the remaining interactions)
    - Cursor-based pagination
- FHIRPath evaluation
  - [FHIRPath v2.0.0](https://hl7.org/fhirpath/N1/) specification; except full UCUM support

    see [below for more information](#fhirpath)

## Getting Started

A quick "getting started" tutorial can be found in the [`./examples/demo`](./examples/demo/main.go) project.

### Other Examples

You can find more examples in [`./examples/`](./examples/o).
The [`mock`](./examples/mock/main.go) example shows how to build custom FHIR® facades on top of legacy data sources
using the capabilities API.
The [`proxy`](./examples/proxy/main.go) example uses the generic API to forward all requests to another FHIR® server.

```sh
go run ./examples/proxy https://server.fire.ly/
```

This starts a simple mock-facade that forwards all requests to a test-server.

From another terminal, run

```sh
curl 'http://localhost/Patient/547'
```

or

```sh
curl 'http://localhost/Patient?_id=547'
```

to get a bundle.

## Capabilities

Everything is designed around capabilities, represented by interfaces (e.g. `PatientSearch`).
This flexible architecture allows different use cases, such as

- building FHIR® facades to legacy systems by implementing a custom backend
- using this library as a FHIR® client (by leveraging a - still to be build - REST backend)

### Concrete vs. Generic API

The library provides two API *styles*.
The **concrete** API:

```Go
func (a myAPI) ReadPatient(ctx context.Context, id string) (r4.Patient, error) {}

func (a myAPI) SearchPatient(ctx context.Context, options search.Options) (search.Result, error) {}
```

and the **generic** API:

```Go
func (a myAPI) Read(ctx context.Context, resourceType, id string) (r4.Patient, error) {}

func (a myAPI) Search(ctx context.Context, resourceType string, options search.Options) (search.Result, error) {}
```

You can implement your custom backend or client either way.
The **concrete** API is ideal for building custom FHIR® facades where a limited set of resources is used (see [
`./examples/mock`](./examples/mock/main.go)).
The **generic** API is better suited for e.g. building FHIR® clients (see [
`./examples/proxy`](./examples/proxy/main.go))
or standalone FHIR® servers.

#### Interoperability

Wrapper structs facilitate interoperability between the generic and the concrete API.

```Go
genericAPI := capabilitiesR4.Generic{Concrete: concreteAPI}
```

and vice versa:

```Go
concreteAPI := capabilitiesR4.Concrete{Generic: genericAPI}
```

## FHIRPath

The [FHIRPath v2.0.0](https://hl7.org/fhirpath/N1/) specification is implemented with the exception of full UCUM support.
For quantity comparisons and operations, the unit is only asserted for equality.

From the additional functions defined in the FHIR specification, only
* `extension(url : string) : collection`

is implemented.

Mostly, because these require validation which is not implemented by `fhir-toolbox-go`, yet.

For a quick usaage example see  [`./examples/fhirpath`](./examples/fhirpath/main.go).

### Decimal precision

The FHIRPath evaluation engine uses [apd.Decimal](https://pkg.go.dev/github.com/cockroachdb/apd#Decimal) under the hood.
Precision of decimal operations can be set by supplying
an [apd.Context](https://pkg.go.dev/github.com/cockroachdb/apd#Context)

```Go
// Setup context
ctx := r4.Context()
// with defined precision for decimal operations.
ctx = fhirpath.WithAPDContext(ctx, apd.BaseContext.WithPrecision(100))

expr, err := fhirpath.Parse("Observation.value / 3")
if err != nil {
    // Handle error
}

// Evaluate the expression against a FHIR resource
result, err := fhirpath.Evaluate(r4.Context(), observation, expr)
if err != nil {
    // Handle error
}
```

> **Attention**: By default the precision is set to `0`.

### Testing Approach

The FHIRPath implementation is tested against the FHIRPath test suite.
Tests are downloaded on first execution and cached afterward into the `build` folder.
As the test cases XML has some inconsistencies and features not supported yet,
the tests are modified before execution in [`fhirpath/fhirpath_test.go`](fhirpath/fhirpath_test.go)

## Roadmap

- interactions
  - $operations
  - support for resource versioning (`vread`, `history`)
  - at some point `patch` and `batch/transaction`, but no priority at the moment
- constants for code systems and/or value-sets
- adapter for resolving `_include` and `_revinclude`
- validation of resources (also against profiles)

## Packages

| Package              | Description                                                                   |
|----------------------|-------------------------------------------------------------------------------|
| `model`              | Generated FHIR® model types                                                   |
| `capabilities/..`    | Interfaces modeling capabilities a server can provide or a client can consume |
| `capabilites/search` | Types and helper functions for implementing search capabilities               |
| `fhirpath`           | FHIRPath execution engine                                                     |                                                     
| `rest`               | FHIR® REST server implementation                                              |
| `testdata`           | Utils for loading test data and writing tests                                 |
| `examples`           | Examples on what you can do with this module                                  |

### Scope

Everything part of the FHIR® specification is in scope of this project.
However, we (DAMEDIC) do not strive for feature-completion.
Instead we will only implement what we need for building our products.
See [Contribution](#contribution) below.

## Contribution

We are happy to accept contributions.
Bugfixes are always welcomed.
For more elaborate features we appreciate commitment to maintain the contributed code.
