
# generate go files from protobuf files
protoc:
	protoc -I grpc/protos grpc/protos/*.proto --go_out=plugins=grpc:grpc/protos

# run our application
.PHONY: run
run:
	go run service/main.go

# run unit tests
.PHONY: test
test:
	go test -v ./...

# run integration with local instance spun up within the test
.PHONY: integration-local
integration-local:
	go test -tags integration -v ./testing/integration/... -integration-local

# run integration tests
.PHONY: integration
integration:
	go test -tags integration -v ./testing/integration/...

# compile test binary
bin/test:
	go test -tags integration -c ./testing/integration/...
