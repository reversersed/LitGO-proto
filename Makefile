PROTO_FILES := $(shell find proto -name '*.proto')

run: $(PROTO_FILES)
ifeq ($(OS),Windows_NT)
	@echo this makefile not supported on windows system, only to use on linux
else
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/favadi/protoc-go-inject-tag@latest
	@go install github.com/golang/mock/mockgen@latest

	@protoc -I=proto --proto_path=proto --go_out=gen/go --go-grpc_out=gen/go --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PROTO_FILES)
	@protoc-go-inject-tag -input="./gen/go/*/*.pb.go" -remove_tag_comment
	@go generate ./gen/go/...
	@cd ./gen/go && go mod tidy
endif