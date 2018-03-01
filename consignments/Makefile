# Compile proto and build service
build:
	protoc --go_out=plugins=grpc:. proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment .
.PHONY: build

# Start the consignment server
run:
	docker run -d -p 50051:50051 consignment
.PHONY: run