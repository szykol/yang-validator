TAG ?= latest
TEST_RUNNER ?= gotest.tools/gotestsum@latest
SERVICE = yang-validator
PROJECT_PATH = szykol/${SERVICE}:${TAG}

build:
	go build cmd

test:
	go run ${TEST_RUNNER} ./...

lint:
	golangci-lint run ./...

coverage:
	go run ${TEST_RUNNER} -- -coverprofile=cover.out ./...
	go tool cover -html cover.out -o cover.html

container-build:
	docker build . -t ${PROJECT_PATH}

container-run:
	docker run -p 1337:1337 ${PROJECT_PATH}
