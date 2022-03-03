export OPENAPI_FILENAME_COMPLETE=./dist/complete.openapi.yaml
export OPENAPI_FILENAME_MANICURED=./dist/manicured.openapi.yaml

export GO111MODULE=on

### Deps

## Install deps
setup:
	# install golanglint-ci into ./bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.44.2
	# install spectral into ./bin
	curl -sfL https://raw.githack.com/stoplightio/spectral/master/scripts/install.sh | sed 's#/usr/local/bin/spectral#./bin/spectral#g' | sed 's#FILENAME=\"spectral\"#FILENAME=\"spectral-linux\"#g' | sh
	# Install gobin globally
	env GO111MODULE=off go get -u github.com/myitcv/gobin
	# Download packages in go.mod file
	go mod download
.PHONY: setup

## Update deps
update-deps:
	go get -u ./...
	go mod tidy
.PHONY: update-deps

### Verify

## Static analysis
verify:
	# Lint go files
	./bin/golangci-lint run ./...
	# Lint OpenAPI spec
	./bin/spectral lint -v --fail-severity=warn $(OPENAPI_FILENAME_MANICURED)
	./bin/spectral lint -v --fail-severity=warn $(OPENAPI_FILENAME_COMPLETE)
.PHONY: verify

### Codegen

## Run codegen
generate:
	mkdir -p ./dist;

	# Generate manicured openapi spec
	gobin -m -run github.com/the4thamigo-uk/conflate/conflate -data ./openapi/_template.manicured.yaml -format YAML > $(OPENAPI_FILENAME_MANICURED)

	# Generate golang api client pkg/manicured
	mkdir -p ./pkg/manicured
	gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=pkg --generate types -o ./pkg/manicured/types.gen.go $(OPENAPI_FILENAME_MANICURED)
	gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=pkg --generate client -o ./pkg/manicured/client.gen.go $(OPENAPI_FILENAME_MANICURED)

	# Generate complete openapi spec
	gobin -m -run github.com/the4thamigo-uk/conflate/conflate -data ./openapi/_template.complete.yaml -format YAML > $(OPENAPI_FILENAME_COMPLETE)

	# Generate golang api client pkg/complete
	mkdir -p ./pkg/complete
	gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=pkg --generate types -o ./pkg/complete/types.gen.go $(OPENAPI_FILENAME_COMPLETE)
	gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=pkg --generate client -o ./pkg/complete/client.gen.go $(OPENAPI_FILENAME_COMPLETE)
.PHONY: generate

### Test

## Run tests
test: test-unit test-integration
.PHONY: test

## Run tests and generate reports
report-test: report-test-unit report-test-integration
.PHONY: report-test

## Run unit tests
test-unit:
	go test -count=1 ./cmd/... ./internal/... ./pkg/...
.PHONY: test-unit

## Run integration tests
test-integration:
	go test -count=1 ./test-suites/integration/...
.PHONY: test-integration

## Run smoke tests
test-smoke:
	go test -count=1 ./test-suites/smoke/...
.PHONY: test-smoke

## Run unit tests and generate reports
report-test-unit:
	go test -count=1 -coverprofile=reports/test-unit.out -v -p 5 ./cmd/... ./internal/... ./pkg/... | gobin -m -run github.com/apg/patter > reports/test-unit.tap
	cat reports/test-unit.tap
.PHONY: report-test-unit

## Run integration tests and generate reports
report-test-integration:
	go test -count=1 -coverprofile=reports/test-integration.out -v -p 5 ./test-suites/integration/... | gobin -m -run github.com/apg/patter > reports/test-integration.tap
	cat reports/test-integration.tap
.PHONY: report-test-integration

## Run smoke tests and generate reports
report-test-smoke:
	go test -count=1 -coverprofile=reports/test-smoke.out -v -p 5 ./test-suites/smoke/... | gobin -m -run github.com/apg/patter > reports/test-smoke.tap
	cat reports/test-smoke.tap
.PHONY: report-test-smoke
