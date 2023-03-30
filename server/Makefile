BINARY_NAME=crasher
BIN_PATH=bin/
LINTER=golangci-lint
PROTO_SCHEMAS_DIR=pkg/grpc/protobuf-schemas
PROTO_SRC_DIR=pkg/grpc/generated/

.ONESHELL: test build lint all proto
.PHONY: test build lint all proto

all: test build lint

test:
	@go test ./...

build:
	@go build -o $(BIN_PATH)$(BINARY_NAME) cmd/crasher/main.go

lint:
	@$(LINTER) run ./...

proto:
	@protoc --proto_path=$(PROTO_SCHEMAS_DIR)/proto --go_out=$(GOPATH)/src $(PROTO_SCHEMAS_DIR)/proto/*.proto