.PHONY: requirements external-api

GRPC_GW_PATH := $(shell go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway)
GOOGLEAPIS_PATH := "$(GRPC_GW_PATH)/../third_party/googleapis"

all: requirements external-api

requirements:
	@go mod download
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

external-api:
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/external/api --swagger_out=json_names_for_fields=true:. external/api/user.proto