run:
	-@mkdir -p gen\go
	@protoc --proto_path=proto --go_out=gen/go --go-grpc_out=gen/go proto/*.proto