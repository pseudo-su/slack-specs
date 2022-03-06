export OPENAPI_FILENAME_COMPLETE=./dist/openapi.complete.yaml
export OPENAPI_FILENAME_MANICURED=./dist/openapi.manicured.yaml

export GO111MODULE=on

### Manage Dependencies

## Install dependencies
deps.install: testharness.go.deps.install testharness.rust.deps.install testharness.kotlin.deps.install testharness.dart.deps.install
	mkdir -p ./bin
	# install golanglint-ci into ./bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.44.2
	# install spectral into ./bin
	curl -sfL https://raw.githack.com/stoplightio/spectral/master/scripts/install.sh | sed 's#/usr/local/bin/spectral#./bin/spectral#g' | sed 's#FILENAME=\"spectral\"#FILENAME=\"spectral-linux\"#g' | sh
.PHONY: deps.install

## Update dependencies
deps.update: testharness.go.deps.update testharness.rust.deps.update testharness.kotlin.deps.update testharness.dart.deps.update
.PHONY: deps.update

## Update dependencies
deps.upgrade: testharness.go.deps.upgrade testharness.rust.deps.upgrade testharness.kotlin.deps.upgrade testharness.dart.deps.upgrade
.PHONY: deps.upgrade

### Codegen

## Run codegen
generate: generate.openapi testharness.go.generate testharness.rust.generate testharness.kotlin.generate testharness.dart.generate
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
verify: verify.openapi testharness.go.verify testharness.rust.verify testharness.kotlin.verify testharness.dart.verify
.PHONY: verify

## Verify using static analysis and autofix where possible
verify.fix: verify.openapi testharness.go.verify.fix testharness.rust.verify.fix verify.openapi testharness.kotlin.verify.fix testharness.dart.verify.fix
.PHONY: verify.fix

## Verify OpenAPI specs using spectral
verify.openapi:
	# Lint OpenAPI spec
	./bin/spectral lint -v --fail-severity=warn $(OPENAPI_FILENAME_MANICURED)
	./bin/spectral lint --ruleset=.spectral.relaxed.yaml -v --fail-severity=warn $(OPENAPI_FILENAME_COMPLETE)
.PHONY: verify.openapi

### Test

## Run tests
test: testharness.go.test testharness.rust.test testharness.kotlin.test testharness.dart.test
.PHONY: test

## Run tests and generate reports
test.report: testharness.go.test.report testharness.rust.test.report testharness.kotlin.test.report testharness.dart.test.report
.PHONY: test.report

## Run unit tests
test.unit: testharness.go.test.unit testharness.rust.test.unit testharness.kotlin.test.unit testharness.dart.test.unit
.PHONY: test.unit

## Run unit tests and generate reports
test.unit.report: testharness.go.test.unit.report testharness.rust.test.unit.report testharness.kotlin.test.unit.report testharness.dart.test.unit.report
.PHONY: test.unit.report

## Run unit tests
test.smoke: testharness.go.test.smoke testharness.rust.test.smoke testharness.kotlin.test.smoke testharness.dart.test.smoke
.PHONY: test.smoke

## Run smoke tests and generate reports
test.smoke.report: testharness.go.test.smoke.report testharness.rust.test.smoke.report testharness.kotlin.test.smoke.report testharness.dart.test.smoke.report
.PHONY: test.smoke.report
