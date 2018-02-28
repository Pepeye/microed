build:
	protoc --go_out=. service/proto/consignment/consignment.proto
	# protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/Pepeye/microed/consignments/service/proto/consignment/