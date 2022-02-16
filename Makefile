.SILENT:

clear:
	rm pkg/api/user*.go

protoc:
	protoc --proto_path=api/proto api/proto/*.proto --go_out=pkg/ --go-grpc_out=pkg/ --validate_out=lang=go:pkg/