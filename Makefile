
protoc:
	protoc -I grpc/protos grpc/protos/*.proto --go_out=plugins=grpc:grpc/protos

run:
	go run service/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: integration-local
integration-local:
	go test -tags integration -v ./testing/integration/... -integration-local

.PHONY: integration
integration:
	go test -tags integration -v ./testing/integration/...

bin/test:
	go test -tags integration -c ./testing/integration/...
