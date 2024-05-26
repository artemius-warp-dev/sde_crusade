.PHONY: all api scheduler sender integration_tests

all: api scheduler sender

api:
	@echo "Building API..."
	go build -o bin/api ./cmd/api
scheduler:
	@echo "Building Scheduler..."
	go build -o bin/scheduler ./cmd/scheduler
sender:
	@echo "Building sender..."
	go build -o bin/sender ./cmd/sender
clean:
	@echo "Cleaning up binaries..."
	rm -rf cmd/api/api cmd/scheduler/scheduler cmd/sender/sender
integration_tests:
	@echo "Running integration tests..."
	go test -v ./integration_tests
