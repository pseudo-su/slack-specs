### Manage Dependencies

## Install dependencies
deps.install:
	# Download packages in go.mod file
	go mod download
.PHONY: deps.install

## Update dependencies
deps.update:
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
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.1 --package=pkg --generate types -o ./pkg/types.gen.go ../../dist/manicured.openapi.yaml
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.1 --package=pkg --generate client -o ./pkg/client.gen.go ../../dist/manicured.openapi.yaml

	# Generate golang api client pkg/complete
	mkdir -p ./pkg/complete
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.1 --package=pkg --generate types -o ./pkg/complete/types.gen.go ../../dist/complete.openapi.yaml
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.1 --package=pkg --generate client -o ./pkg/complete/client.gen.go ../../dist/complete.openapi.yaml
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
	go run github.com/joho/godotenv/cmd/godotenv@v1.4.0 -f .env.test go test -count=1 ./smoke/...
.PHONY: test.smoke

## Run smoke tests and output reports
test.smoke.report:
	mkdir -p reports
	go run github.com/joho/godotenv/cmd/godotenv@v1.4.0 -f .env.test go test -json -count=1 -covermode=count -coverprofile=reports/test.smoke.coverage.out -v -p 5 ./smoke/... > reports/test.smoke.results.json
	go tool cover -html=./reports/test.smoke.coverage.out -o=./reports/test.smoke.coverage.html
	cat reports/test.smoke.results.json
.PHONY: test.smoke.report
