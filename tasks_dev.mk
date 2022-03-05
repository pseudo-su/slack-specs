export OPENAPI_FILENAME_COMPLETE=./dist/complete.openapi.yaml
export OPENAPI_FILENAME_MANICURED=./dist/manicured.openapi.yaml

export GO111MODULE=on

### Manage Dependencies

## Install dependencies
deps.install: testharness.go.deps.install testharness.rust.deps.install
	mkdir -p ./bin
	# install golanglint-ci into ./bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.44.2
	# install spectral into ./bin
	curl -sfL https://raw.githack.com/stoplightio/spectral/master/scripts/install.sh | sed 's#/usr/local/bin/spectral#./bin/spectral#g' | sed 's#FILENAME=\"spectral\"#FILENAME=\"spectral-linux\"#g' | sh
.PHONY: deps.install

## Update dependencies
deps.update: testharness.go.deps.update testharness.rust.deps.update
.PHONY: deps.update

## Update dependencies
deps.upgrade: testharness.go.deps.upgrade testharness.rust.deps.upgrade
.PHONY: deps.upgrade

### Codegen

## Run codegen
generate: generate.openapi testharness.go.generate testharness.rust.generate
.PHONY: generate

## Generate OpenAPI specs
generate.openapi:
	mkdir -p ./dist;

	# Generate manicured openapi spec
	go run github.com/the4thamigo-uk/conflate/conflate@v1.0.1 -data ./openapi/_template.manicured.yaml -format YAML > $(OPENAPI_FILENAME_MANICURED)

	# Generate complete openapi spec
	go run github.com/the4thamigo-uk/conflate/conflate@v1.0.1 -data ./openapi/_template.complete.yaml -format YAML > $(OPENAPI_FILENAME_COMPLETE)
.PHONY: generate.openapi

### Verify

## Verify using static analysis
verify: verify.openapi testharness.go.verify testharness.rust.verify
.PHONY: verify

## Verify using static analysis and autofix where possible
verify.fix: testharness.go.verify.fix testharness.rust.verify.fix verify.openapi
.PHONY: verify.fix

## Verify OpenAPI specs using spectral
verify.openapi:
	# Lint OpenAPI spec
	./bin/spectral lint -v --fail-severity=warn $(OPENAPI_FILENAME_MANICURED)
	./bin/spectral lint --ruleset=.spectral.relaxed.yaml -v --fail-severity=warn $(OPENAPI_FILENAME_COMPLETE)
.PHONY: verify.openapi

### Test

## Run tests
test: testharness.go.test testharness.rust.test
.PHONY: test

## Run tests and generate reports
test.report: testharness.go.test.report testharness.rust.test.report
.PHONY: test.report
