OUT_FILE := ./bin/che-operator-test-harness
DOCKER_IMAGE_NAME :=quay.io/flacatus/operator
VERSION :="0.0.1"

build:
	CGO_ENABLED=0 go test -v -c -o ${OUT_FILE} ./cmd/operator_osd/che_operator_test_harness_suite_test.go

build-container:
	docker build -t $(DOCKER_IMAGE_NAME):$(VERSION) .
