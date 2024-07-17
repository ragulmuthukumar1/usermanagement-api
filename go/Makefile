.PHONY: requirements external-api

GRPC_GW_PATH := $(shell go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway)
GOOGLEAPIS_PATH := "$(GRPC_GW_PATH)/../third_party/googleapis"

all: requirements external-api

requirements:
	@go mod download
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go install github.com/golang/protobuf/protoc-gen-go


external-api:
	# gRPC
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/external/api --go_out=plugins=grpc,paths=source_relative:. external/api/user.proto

	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. external/api/user.proto
	
