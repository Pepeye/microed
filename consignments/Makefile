build:
	protoc --go_out=plugins=grpc:. service/proto/consignment/consignment.proto
	# protoc -I ./service/proto/consignment/ ./service/proto/consignment/consignment.proto --go_out=plugins=grpc:consignment
	# protoc -I . --go_out=plugins=grpc:$(GOPATH)/src/github.com/Pepeye/microed/consignments/service/proto/consignment *.proto
	# protoc -I=$(GOPATH)/src/github.com/Pepeye/microed/consignments/service/proto/consignment --go_out=. *.proto
	# protoc --go_out=. service/proto/consignment/consignment.proto
	# protoc --go_out=plugins=grpc:. *.proto
	# protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld