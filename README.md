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

- FHIR® model types with JSON and XML (un)marshalling
    - R4, R4B & R5

      use build tags `r4`, `r4b` or `r5` for conditional compilation if you only need runtime support for specific
      versions

    - generated from the FHIR® specification
- Extensible REST API with capabilities modeled as interfaces
    - Capability detection by runtime ~~reflection~~ type assertion (see [Capabilities](#capabilities))
        - alternatively: generic API for building wrappers
        - automatic CapabilityStatement generation
    - Interactions: `read`,  `search` (adding the remaining interactions is definitely on the agenda)
    - Cursor-based pagination
- FHIRPath evaluation
    - see [FHIRPath Implementation Status](#fhirpath-implementation-status) below

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

This starts a simple mock-facade that forwards all requests to the HAPI test-server.

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
func (a myAPI) ReadPatient(ctx context.Context, id string) (r4.Patient, capabilities.FHIRError) {}

func (a myAPI) SearchPatient(ctx context.Context, options search.Options) (search.Result, capabilities.FHIRError) {}
```

and the **generic** API:

```Go
func (a myAPI) Read(ctx context.Context, resourceType, id string) (r4.Patient, capabilities.FHIRError) {}

func (a myAPI) Search(ctx context.Context, resourceType string, options search.Options) (search.Result, capabilities.FHIRError) {}
```

You can implement your custom backend or client either way.
The **concrete** API is ideal for building custom FHIR® facades where a limited set of resources is used (see [
`./examples/mock`](./examples/mock/main.go)).
The **generic** API is better suited for e.g. building FHIR® clients (see [
`./examples/proxy`](./examples/proxy/main.go))
or standalone FHIR® servers.

#### Interoperability

The `capabilities/wrap` package provides two wrapper functions to wrap a concrete into the generic API:

```Go
genericAPI := wrap.Generic[model.R4](concreteAPI)
```

and vice versa:

```Go
concreteAPI := wrap.ConcreteR4(genericAPI)
```

## Roadmap

- proper handling of `_include` and `_revinclude`
- managed search parameters
- resource validation
- remaining interactions (`vread`, `create`, `update`, `patch`, `delete`, `history`)

## FHIRPath Implementation Status

The project includes a work-in-progress implementation for [FHIRPath v2.0.0](https://hl7.org/fhirpath/N1/).
Refer to the following table outlining the implementation statuses of different parts of the specification.

| Section of the specification             | Implementation Status                         |
|------------------------------------------|-----------------------------------------------|
| **1. Background**                        | not applicable                                |
| **2. Navigation model**                  | complete                                      |
| **3. Path selection**                    | complete                                      |
| 3.1. Collections                         | complete                                      |
| 3.2. Paths and polymorphic items         | complete                                      |
| **4. Expressions**                       | wip                                           |
| 4.1. Literals                            | complete                                      |
| 4.2. Operators                           | see "6. Operations"                           |
| 4.3. Function Invocations                | complete                                      |
| 4.4. Null and empty                      | complete                                      |
| 4.5. Singleton Evaluation of Collections | complete                                      |
| **5. Functions**                         | wip                                           |
| 5.1. Existence                           | wip                                           |
| 5.2. Filtering and projection            | wip                                           |
| 5.3. Subsetting                          | wip                                           |
| 5.4. Combining                           | wip                                           |
| 5.5. Conversion                          | wip; except `toQuantity` with unit conversion |
| 5.6. String Manipulation                 | wip                                           |
| 5.7. Math                                | wip                                           |
| 5.8. Tree navigation                     | wip                                           |
| 5.9. Utility functions                   | wip                                           |
| **6. Operations**                        | wip                                           |
| 6.1. Equality                            | complete                                      |
| 6.2. Comparison                          | complete                                      |
| 6.3. Types                               | wip                                           |
| 6.4. Collections                         | wip                                           |
| 6.5. Boolean logic                       | wip                                           |
| 6.6. Math                                | wip                                           |
| 6.7. Date/Time Arithmetic                | wip                                           |
| 6.8. Operator precedence                 | handled by ANTLR                              |
| **7. Aggregates**                        | wip                                           |
| **8. Lexical Elements**                  | handled by ANTLR                              |
| **9. Environment variables**             | complete                                      |
| **10. Types and Reflection**             | wip                                           |
| 10.1. Models                             | wip                                           |
| 10.2. Reflection                         | todo                                          |

`todo`: in scope, but has not a very high priority at the moment.

## Packages

| Package              | Description                                                                   |
|----------------------|-------------------------------------------------------------------------------|
| `model`              | Generated FHIR® model types                                                   |
| `capabilities/..`    | Interfaces modeling capabilities a server can provide or a client can consume |
| `capabilites/search` | Types and helper functions for implementing search capabilities               |
| `capabilites/wrap`   | Conversion between the concrete and generic capabilities API                  |
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
