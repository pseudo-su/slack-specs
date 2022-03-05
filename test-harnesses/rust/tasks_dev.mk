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
test: test.smoke
.PHONY: test

## Run tests and generate reports
test.report: test.smoke.report
.PHONY: test.report

## Run smoke tests
test.smoke:
.PHONY: test.smoke

## Run smoke tests and output reports
test.smoke.report:
.PHONY: test.smoke.report
