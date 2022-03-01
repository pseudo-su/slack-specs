export OPENAPI_FILENAME=./openapi.bundle.yaml

export GO111MODULE=on

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

verify:
	# Lint go files
	./bin/golangci-lint run ./...
	# Lint OpenAPI spec
	./bin/spectral lint -v --fail-severity=warn $(OPENAPI_FILENAME)
.PHONY: verify

generate:
	# Generate openapi spec
	gobin -m -run github.com/the4thamigo-uk/conflate/conflate -data ./openapi/_template.yaml -format YAML > $(OPENAPI_FILENAME)
	# Generate golang api client pkg/
	gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=pkg --generate types -o ./pkg/types.gen.go $(OPENAPI_FILENAME)
	gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=pkg --generate client -o ./pkg/client.gen.go $(OPENAPI_FILENAME)
.PHONY: generate

update-deps:
	go get -u ./...
	go mod tidy
.PHONY: update-deps

test: test-unit test-integration
.PHONY: test

report-test: report-test-unit report-test-integration
.PHONY: report-test

test-unit:
	go test -count=1 ./cmd/... ./internal/... ./pkg/...
.PHONY: test-unit

test-integration:
	go test -count=1 ./test-suites/integration/...
.PHONY: test-integration

test-smoke:
	go test -count=1 ./test-suites/smoke/...
.PHONY: test-smoke

report-test-unit:
	go test -count=1 -coverprofile=reports/test-unit.out -v -p 5 ./cmd/... ./internal/... ./pkg/... | gobin -m -run github.com/apg/patter > reports/test-unit.tap
	cat reports/test-unit.tap
.PHONY: report-test-unit

report-test-integration:
	go test -count=1 -coverprofile=reports/test-integration.out -v -p 5 ./test-suites/integration/... | gobin -m -run github.com/apg/patter > reports/test-integration.tap
	cat reports/test-integration.tap
.PHONY: report-test-integration

report-test-smoke:
	go test -count=1 -coverprofile=reports/test-smoke.out -v -p 5 ./test-suites/smoke/... | gobin -m -run github.com/apg/patter > reports/test-smoke.tap
	cat reports/test-smoke.tap
.PHONY: report-test-smoke
