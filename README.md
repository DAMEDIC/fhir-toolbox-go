# fhir-toolbox-go

This projects aims to provide a set of packages for building FHIR services in Go.
This includes model types, as well as interfaces (and some implementations) modeling capabilities that a server provides or a client can consume.

## Features

- Extensible REST API with capabilties modeled as interfaces
  - FHIR model types
  - Capability detection by runtime ~~reflecting~~ type asserting implemented interfaces (see [Capabilities](#capabilities))
  - Interactions: `read`,  `search`

## Examples

You can find examples in `./examples`.
The `facade` example shows how to build custom FHIR facades using the capabilities API.
The `proxy` example uses the generic API to forward all requests to another FHIR server.

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

The client implementation in `./cmd/example` is only a proof-of-concept and does only implement the `read` interaction for a very limited set of resources.

## Capabilities

Everyting is archtectured around capabilities, represented by interfaces (e.g. `PatientSearch`).
This flexible architecture allows different use cases, such as

- building FHIR facades to legacy systems by implementing a custom backend
- using this library as a FHIR client (by leveraging a - still to be build - REST backend)

## Packages

| Package      | Description                                                                                                                                  |
| ------------ | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `model`      | Generated FHIR model types                                                                                                                   |
| `capability` | Interfaces modeling capabilities a server provides or a client can consume                                                                   |
| `dispatch`   | Dispatch from resource types to concrete implementations (e.g. `Read("Patient", 123) to ReadPatient(123)`). This is used by the REST server. |
| `rest`       | FHIR REST server implementation                                                                                                              |
| `backend`    | Implementations of capabilities for different backends; database wrappers or the REST client go here                                         |
| `generate`   | Code generation tools for generating FHIR model types, capabilites and dispatchers from the FHIR specification                               |


### Scope

Everything part of the FHIR specification is in scope of this project.
However, we (DAMEDIC) do not strive for feature-completion.
Instead we will only implemented what we need for building our products.
See [Contribution](#contribution) below.

## Contribution

We are happy to accept contributions to this project.
Bugfixes are always welcomed.
For more complex features we would like to see long-term commitment to maintain the contirbuted source codes.
