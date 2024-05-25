.PHONY: all api scheduler sender

all: api scheduler sender

api:
	@echo "Building API..."
	go build -o cmd/api ./cmd/api
scheduler:
	@echo "Building Scheduler..."
	go build -o cmd/scheduler ./cmd/scheduler
sender:
	@echo "Building sender..."
	go build -o cmd/sender ./cmd/sender
clean:
	@echo "Cleaning up binaries..."
	rm -rf cmd/api/api cmd/scheduler/scheduler cmd/sender/sender
