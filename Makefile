PROTO_FILES := $(shell find proto -name '*.proto')
SWAGGER_FILES := $(shell find docs/swagger -name '*.swagger.json')

run: $(PROTO_FILES)
ifeq ($(OS),Windows_NT)
	@echo this makefile not supported on windows system, only to use on linux
else
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/favadi/protoc-go-inject-tag@latest
	@go install github.com/golang/mock/mockgen@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
	@go install github.com/go-swagger/go-swagger/cmd/swagger@latest

	@protoc -I=proto --proto_path=proto --go_out=gen/go --go-grpc_out=gen/go --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PROTO_FILES)
	@protoc -I=proto --proto_path=proto --grpc-gateway_out=gen/go --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true $(PROTO_FILES)
	@protoc -I=proto --proto_path=proto --openapiv2_out=gen/docs/swagger $(PROTO_FILES)
	@swagger mixin $(SWAGGER_FILES) -o ./gen/docs/swagger/swagger.json
	@protoc-go-inject-tag -input="./gen/go/*/*.pb.go" -remove_tag_comment
	@protoc-go-inject-tag -input="./gen/go/*/*.pb.gw.go" -remove_tag_comment
	@go generate ./gen/go/...
	@cd ./gen/go && go mod tidy
endif