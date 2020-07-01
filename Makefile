# Proto parameters
SERVER_OUT := "accounty"
PROTO_FILES_PATH := api/protobuf-spec
PROTO_OUT := genproto

.PHONY: all api build

all: build

api/users/v1/core.pb.go: api/protobuf-spec/core.proto
	@protoc -I api/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:api/ \
		api/protobuf-spec/core.proto

api/users/v1/core.pb.gw.go: api/protobuf-spec/core.proto
	@protoc -I api/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:api/ \
		api/protobuf-spec/core.proto

api/protobuf-spec/core.swagger.json: api/protobuf-spec/core.proto
	@protoc -I api/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:api/ \
		api/protobuf-spec/core.proto

api: api/users/v1/core.pb.go api/users/v1/core.pb.gw.go api/protobuf-spec/core.swagger.json ## Auto-generate grpc go sources

clean: ## Remove previous builds
	@rm -rf api/users/v1/core.pb.go api/users/v1/core.pb.gw.go api/protobuf-spec/core.swagger.json $(SERVER_OUT)

build: api
	@go build -o $(SERVER_OUT)
