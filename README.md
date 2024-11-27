# fhir-toolbox-go

This project provides a set of packages for building FHIR services in Go.

This includes model types and interfaces modeling capabilities that you can use to build custom FHIR servers.

> While used in production at DAMEDIC, this project is still in its early days
> and the feature set is quite limit.
> We will add features as we require them. We welcome external contributions.

## Features

- FHIR model types with JSON and XML (un)marshalling
- Extensible REST API with capabilities modeled as interfaces
    - Capability detection by runtime ~~reflection~~ type assertion (see [Capabilities](#capabilities))
        - alternatively: generic API for building wrappers
    - Interactions: `read`,  `search` (adding the remaining interactions is definitely on the agenda)
    - Cursor-based pagination

## Getting Started

A quick "getting started" tutorial can be found in the [`./examples/demo`](./examples/demo/main.go) project.

### Other Examples

You can find more examples in [`./examples/`](./examples/o).
The [`facade`](./examples/facade/main.go) example shows how to build custom FHIR facades on top of legacy data sources using the capabilities API.
The [`proxy`](./examples/proxy/main.go) example uses the generic API to forward all requests to another FHIR server.

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

Everything is architectured around capabilities, represented by interfaces (e.g. `PatientSearch`).
This flexible architecture allows different use cases, such as

- building FHIR facades to legacy systems by implementing a custom backend
- using this library as a FHIR client (by leveraging a - still to be build - REST backend)

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
The **concrete** API is ideal for building custom FHIR facades where a limited set of resources is used (see [
`./examples/mock`](./examples/mock/main.go)).
The **generic** API is better suited for e.g. building FHIR clients (see [`./examples/proxy`](./examples/proxy/main.go))
or standalone FHIR servers.

#### Interoperability

The `capabilities/wrap` package provides two wrapper functions to wrap a concrete into the generic API:

```Go
genericAPI := wrap.Generic[model.R4](concreteAPI)
```

and vice versa:

```Go
concreteAPI := wrap.ConcreteR4(genericAPI)
```

## Packages

| Package              | Description                                                                   |
|----------------------|-------------------------------------------------------------------------------|
| `model`              | Generated FHIR model types                                                    |
| `capabilities/..`    | Interfaces modeling capabilities a server can provide or a client can consume |
| `capabilites/search` | Types and helper functions for implementing search capabilities               |
| `capabilites/wrap`   | Conversion between the concrete and generic capabilities API                  |
| `rest`               | FHIR REST server implementation                                               |
| `testdata`           | Utils for loading test data and writing tests                                 |
| `examples`           | Examples on what you can do with this module                                  |

### Scope

Everything part of the FHIR specification is in scope of this project.
However, we (DAMEDIC) do not strive for feature-completion.
Instead we will only implement what we need for building our products.
See [Contribution](#contribution) below.

## Contribution

We are happy to accept contributions.
Bugfixes are always welcomed.
For more elaborate features we appreciate commitment to maintain the contributed code.
