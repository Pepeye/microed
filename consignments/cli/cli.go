package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/Pepeye/microed/consignments/service/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	// setup connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Connection failed")
	}
	defer conn.Close()
	client := pb.NewShippingServiceClient(conn)

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consingment, err := parseFile(file)
	if err != nil {
		log.Fatal("Could not parse file")
	}

	r, err := client.CreateConsignment(context.Background(), consingment)
	if err != nil {
		log.Fatal("Unable to create consignment")
	}

	log.Printf("Created consignment: %v", r.Created)

	resp, err := client.GetConsignment(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatal("Unable to get consignments")
	}

	for _, c := range resp.Consignments {
		log.Println(c)
	}
}
