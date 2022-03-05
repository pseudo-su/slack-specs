# Slack specs

This projects goal is to try and create high quality OpenAPI v3 specs of [Slack's Web API](https://api.slack.com/web)

The starting point of this project was to manually extract the [OpenAPI v2 specs that slack publish](https://github.com/slackapi/slack-api-specs/tree/master/web-api) which are unsuitable for generating a client with, convert them to OpenAPI v3 yaml files, split each tag/category out into it's own folder and then striving to make manual fixes to them as required.

## How we author, compile, verify and validate the API specs

This repository uses a spec-first approach, meaning that the spec definitions are authored by contributers, it is NOT generated from any external code/systems.

Authoring such a complicated OpenAPI spec can be laborious, so in order to make the authoring process easier we do a number of things that help faciliate an acceptable authoring experience, these include:

* Split the OpenAPI spec definitions across multiple files and compiling them together as a build step using [`conflate`](https://github.com/miracl/conflate)
* Generate/Compile two versions of the Slack API spec
  * A `complete` one that all the API Operations and components that were imported as part of the original v2 OpenAPI spec
  * A `manicured` one that contains a subset of API operations that have been manually reviewed to be accurate and/or complete enough to include in (may still contain issues).
* Use [`spectral`](https://stoplight.io/open-source/spectral/) to lint the generated OpenAPI v3 specs to ensure they are valid according to the OpenAPI spec and some other desired linting rules
* Generate a "Golang API client" inside `pkg/` from the bundled spec
* Write a suite of tests in `test-suites/smoke/` using the Generated Golang client to execute against the API

## Quickstart

```sh
# Install/setup tool deps
make deps.install

# Generate
make generate

# Verify
make verify

# Test
make test-smoke
```

## TODO

* The generated API client doesn't do any request/response validation because `oapi-codegen` doesn't support it yet. Instead could use something like https://github.com/EXXETA/openapi-cop
