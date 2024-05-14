# fhir-toolbox-go

This projects aims to provide a set of packages for building FHIR services in Go.
This includes model types, as well as interfaces (and some implementations) modeling capabilities that a server provides or a client can consume.

## Features

none yet :/

## Example

Run the example in `./example` using the following command:

```sh
go run ./cmd/example https://server.fire.ly/
```

This starts a simple mock-facade that forwards all requests to the HAPI test-server.

From another terminal, run

```sh
curl 'http://localhost/Patient/547'
```

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

## Roadmap

- Extensible REST API (WIP)
  - FHIR types
  - Capabilities modeled as interfaces
  - Capability detection by runtime ~~reflecting~~ type assertion of implemented interfaces (see [Capabilities](#capabilities))
- Proof-of-conept REST backend (WIP)
- Read-only Orbis Backend (TODO)
- Read-only M-KIS Backend (TODO)

### Scope

Everything part of the FHIR specification is in scope of this project.
However, we (DAMEDIC) do not strive for feature-completion.
Instead we will only implemented what we need for building our products.
See [Contribution](#contribution) below.

## Contribution

We are happy to accept contributions to this project.
Bugfixes are always welcomed.
For more complex features we would like to see long-term commitment to maintain the contirbuted source codes.
