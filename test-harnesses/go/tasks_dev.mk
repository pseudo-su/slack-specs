### Manage Dependencies

## Install dependencies
deps.install:
	# Download packages in go.mod file
	go mod download
.PHONY: deps.install

## Update dependencies
deps.update:
	go mod tidy
.PHONY: deps.update

## Upgrade dependencies
deps.upgrade:
	go get -u ./...
	go mod tidy
.PHONY: deps.upgrade

### Verify

## Static analysis
verify:
	../../bin/golangci-lint run ./...
.PHONY: verify

### Codegen

## Run codegen
generate:
	# Generate golang api client pkg
	mkdir -p ./pkg
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.1 --package=pkg --generate types -o ./pkg/types.gen.go ../../dist/openapi.manicured.yaml
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.1 --package=pkg --generate client -o ./pkg/client.gen.go ../../dist/openapi.manicured.yaml
.PHONY: generate

### Test

TEST_GODOTENV := go run github.com/joho/godotenv/cmd/godotenv@v1.4.0 -f ../../.env.local

## Run tests
test: test.unit test.smoke
.PHONY: test

## Run tests and generate reports
test.report: test.unit.report test.smoke.report test.report.lift
.PHONY: test.report

## Lift reports up into the project root
test.report.lift:
	mkdir -p ../../reports/go
	cp -r ./reports/ ../../reports/go
.PHONY: test.report.lift

## Run unit tests
test.unit:
	${TEST_GODOTENV} go test -v -p 5 -count=1 ./test/unit/...
.PHONY: test.unit

## Run unit tests and output reports
test.unit.report:
	mkdir -p reports
	${TEST_GODOTENV} go test -json -count=1 -covermode=count -coverprofile=reports/test.unit.coverage.out -v -p 5 ./test/unit/... > reports/test.unit.results.json
	go tool cover -html=./reports/test.unit.coverage.out -o=./reports/test.unit.coverage.html
	cat reports/test.unit.results.json
.PHONY: test.unit.report

## Run smoke tests
test.smoke:
	${TEST_GODOTENV} go test -v -p 5 -count=1 ./test/smoke/...
.PHONY: test.smoke

## Run smoke tests and output reports
test.smoke.report:
	mkdir -p reports
	${TEST_GODOTENV} go test -json -count=1 -covermode=count -coverprofile=reports/test.smoke.coverage.out -v -p 5 ./test/smoke/... > reports/test.smoke.results.json
	go tool cover -html=./reports/test.smoke.coverage.out -o=./reports/test.smoke.coverage.html
	cat reports/test.smoke.results.json
.PHONY: test.smoke.report
