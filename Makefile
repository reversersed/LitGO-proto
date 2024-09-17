run:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/favadi/protoc-go-inject-tag@latest
	@go install github.com/golang/mock/mockgen@latest

	@protoc --proto_path=proto --go_out=gen/go --go-grpc_out=gen/go proto/*.proto
	@protoc-go-inject-tag -input="./gen/go/*/*.pb.go" -remove_tag_comment
	@go generate ./gen/go/...