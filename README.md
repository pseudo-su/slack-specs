# Slack specs

An experiment to try and create a tolerable authoring, linting and testing process for non-trivial OpenAPI specs.

* create a spec across multiple files in `openapi/`
* bundle the spec into a single file `bundle.openapi.yaml`
* Lint the bundled spec file
* Generate a "Golang API client" inside `pkg/` from the bundled spec
* Write a suite of tests in `test-suites/smoke/` using the Generated Golang client to execute against the API

## Quickstart

```sh
# Install/setup tool deps
make setup

# Generate
make generate

# Lint
make lint

# Test
make test-smoke
```

## TODO

* The generated API client doesn't do any request/response validation because `oapi-codegen` doesn't support it yet. Instead could use something like https://github.com/EXXETA/openapi-cop
