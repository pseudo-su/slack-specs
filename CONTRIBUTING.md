# Contributing guide

- [Contributing guide](#contributing-guide)
  - [Initial setup](#initial-setup)
    - [Global dependencies](#global-dependencies)
    - [Local config files](#local-config-files)
  - [Run development scripts](#run-development-scripts)
  - [Project overview](#project-overview)
  - [Updating an operation](#updating-an-operation)
    - [Updating the specs](#updating-the-specs)
    - [Adding unit tests](#adding-unit-tests)
    - [Adding smoke tests](#adding-smoke-tests)

## Initial setup

### Global dependencies

You will need to have the following global dependencies installed on your local machine

- [go](https://go.dev/doc/install)
- [cargo](https://doc.rust-lang.org/cargo/getting-started/installation.html)

### Local config files

```sh
# Copy the example .env.local file (used to override environment variables for local development)
cp docs/examples/.env.local .

# Copy the recommended vscode settings to your workspace config
cp .vscode/settings.recommended.json .vscode/settings.json
```

## Run development scripts

```sh
# Install/setup tool deps
make deps.install

# Generate OpenAPI specs and test harness client libraries
make generate

# Verify OpenAPI specs and test harnesses client librarires using static analysis tools
make verify

# Test OpenAPI specs and test harnesses client librarires and generate reports
make test
```

## Project overview

TODO

## Updating an operation

Fully implementing a new operation involves

- Updating the OpenAPI specs in `./openapi/` based on the [official slack documentation site](https://api.slack.com/methods). The OpenAPI operations should already contain reference to the specific relevant docs underneath `externalDocs`.
- Adding unit tests for the new operation in the `./test-harnessess/go` project within the `./test/unit` folder
- Adding smoke tests for the new operation in the `./test-harnessess/go` project within the `./test/smoke` folder

### Updating the specs

TODO

### Adding unit tests

TODO

### Adding smoke tests

TODO
