
protoc:
	protoc -I grpc/protos grpc/protos/*.proto --go_out=plugins=grpc:grpc/protos