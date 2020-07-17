# Proto parameters
SERVER_OUT := "accounty"
PROTO_FILES_INPUT_FILE := api/core.proto
PROTO_FILES_OUTPUT_PATH := pkg/
PROTO_OUT := genproto

.PHONY: all api build

all: build

pkg/users/v1/core.pb.go: api/core.proto
	@protoc -I api/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:${PROTO_FILES_OUTPUT_PATH} \
		${PROTO_FILES_INPUT_FILE}

pkg/users/v1/core.pb.gw.go: api/core.proto
	@protoc -I api/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:${PROTO_FILES_OUTPUT_PATH} \
		${PROTO_FILES_INPUT_FILE}

pkg/users/v1/core.swagger.json: api/core.proto
	@protoc -I api/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:${PROTO_FILES_OUTPUT_PATH} \
		${PROTO_FILES_INPUT_FILE}

api: pkg/users/v1/core.pb.go pkg/users/v1/core.pb.gw.go pkg/users/v1/core.swagger.json ## Auto-generate grpc go sources

clean: ## Remove previous builds
	@rm -f pkg/users/v1/core.pb.go pkg/users/v1/core.pb.gw.go pkg/users/v1/core.swagger.json $(SERVER_OUT)

build: api
	@go build -o $(SERVER_OUT)
