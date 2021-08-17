package main

import (
	"context"
	"golang_api/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer connection.Close()

	client := pb.NewHelloServiceClient(connection)

	request := &pb.HelloRequest{
		Name: "João",
	}

	response, err := client.Hello(context.Background(), request)

	if err != nil {
		log.Fatalf("Error during request: %v", err)
	}

	log.Println(response.Msg)
}
