export OPENAPI_FILENAME=./openapi.bundle.yaml

export GO111MODULE=on

help:
	@echo "This is a helper makefile for oapi-codegen"
	@echo "Targets:"
	@echo "    setup:              install the development dependencies"
	@echo "    generate:           rerun code generation"
	@echo "    lint:               run linters"
	@echo "    update-deps:        update all go module dependencies"
	@echo "    test-unit:          run unit tests"
	@echo "    test-integration:   run integration tests"
	@echo "    test-smoke:         run smoke tests"
	@echo "    local-ci-build:     emulate a CI build on your local machine"

.PHONY: setup
setup:
	# install golanglint-ci into ./bin
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.24.0
	# install spectral into ./bin
	curl -sfL https://raw.githack.com/stoplightio/spectral/master/scripts/install.sh | sed 's#/usr/local/bin/spectral#./bin/spectral#g' | sed 's#FILENAME=\"spectral\"#FILENAME=\"spectral-linux\"#g' | sh
	# Install gobin globally
	env GO111MODULE=off go get -u github.com/myitcv/gobin
	# Download packages in go.mod file
	go mod download

.PHONY: lint
lint:
	# Lint go files
	./bin/golangci-lint run ./...
	# Lint OpenAPI spec
	./bin/spectral lint -v --fail-severity=warn $(OPENAPI_FILENAME)

.PHONY: generate
generate:
	# Generate openapi spec
	gobin -m -run github.com/the4thamigo-uk/conflate/conflate -data ./openapi/_template.yaml -format YAML > $(OPENAPI_FILENAME)
	# Generate golang api client pkg/
	gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=pkg --generate types -o ./pkg/types.gen.go $(OPENAPI_FILENAME)
	gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=pkg --generate client -o ./pkg/client.gen.go $(OPENAPI_FILENAME)

.PHONY: update-deps
update-deps:
	go get -u ./...
	go mod tidy

.PHONY: test
test: test-unit test-integration

.PHONY: report-test
report-test: report-test-unit report-test-integration

.PHONY: test-unit
test-unit:
	go test -count=1 ./cmd/... ./internal/... ./pkg/...

.PHONY: test-integration
test-integration:
	go test -count=1 ./test-suites/integration/...

.PHONY: test-smoke
test-smoke:
	go test -count=1 ./test-suites/smoke/...

.PHONY: report-test-unit
report-test-unit:
	go test -count=1 -coverprofile=reports/test-unit.out -v -p 5 ./cmd/... ./internal/... ./pkg/... | gobin -m -run github.com/apg/patter > reports/test-unit.tap
	cat reports/test-unit.tap

.PHONY: report-test-integration
report-test-integration:
	go test -count=1 -coverprofile=reports/test-integration.out -v -p 5 ./test-suites/integration/... | gobin -m -run github.com/apg/patter > reports/test-integration.tap
	cat reports/test-integration.tap

.PHONY: report-test-smoke
report-test-smoke:
	go test -count=1 -coverprofile=reports/test-smoke.out -v -p 5 ./test-suites/smoke/... | gobin -m -run github.com/apg/patter > reports/test-smoke.tap
	cat reports/test-smoke.tap
