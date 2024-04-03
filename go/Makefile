.PHONY: requirements common as as-external-api

GRPC_GW_PATH := $(shell go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway)
GOOGLEAPIS_PATH := "$(GRPC_GW_PATH)/../third_party/googleapis"

all: requirements common as as-external-api

requirements:
	@go mod download
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go install github.com/golang/protobuf/protoc-gen-go

common:
	protoc -I=../protobuf -I=../protobuf/common --go_out=plugins=grpc,paths=source_relative:. common/common.proto

as:
	protoc -I=../protobuf -I=../protobuf/as --go_out=plugins=grpc,paths=source_relative:. as/as.proto

as-external-api:
	# gRPC
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/organization.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/user.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/internal.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --go_out=plugins=grpc,paths=source_relative:. as/external/api/inventory.proto

	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/organization.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/user.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/internal.proto
	protoc -I=$(GOOGLEAPIS_PATH) -I=../protobuf -I=../protobuf/as/external/api --grpc-gateway_out=paths=source_relative,logtostderr=true:. as/external/api/inventory.proto
