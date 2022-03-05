### Manage Dependencies

## Install dependencies
deps.install:
.PHONY: deps.install

## Update dependencies
deps.update:
.PHONY: deps.update

## Upgrade dependencies
deps.upgrade:
.PHONY: deps.upgrade

### Verify

## Static analysis
verify:
.PHONY: verify

### Codegen

## Run codegen
generate:
.PHONY: generate

### Test

## Run tests
test: test.unit test.smoke
.PHONY: test

## Run tests and generate reports
test.report: test.unit.report test.smoke.report test.report.lift
.PHONY: test.report

## Lift reports up into the project root
test.report.lift:
	mkdir -p ../../reports/rust
	cp -r ./reports/ ../../reports/rust
.PHONY: test.report.lift

## Run unit tests
test.unit:
.PHONY: test.unit

## Run unit tests and output reports
test.unit.report:
.PHONY: test.unit.report

## Run smoke tests
test.smoke:
.PHONY: test.smoke

## Run smoke tests and output reports
test.smoke.report:
.PHONY: test.smoke.report
