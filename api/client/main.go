package main

import (
	"context"
	"golang_api/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer connection.Close()

	Hello(context.Background(), connection)
	GetAllUsers(context.Background(), connection)
	GetByNameUsers(context.Background(), connection)
}

func Hello(ctx context.Context, connection *grpc.ClientConn) (*pb.HelloResponse, error) {
	client := pb.NewHelloServiceClient(connection)

	request := &pb.HelloRequest{
		Name: "João",
	}

	response, err := client.Hello(ctx, request)

	if err != nil {
		log.Fatalf("Error during request: %v", err)
	}

	log.Println(response.Msg)

	return response, err
}

func GetAllUsers(ctx context.Context, connection *grpc.ClientConn) (*pb.UsersResponse, error) {
	client := pb.NewUsersServiceClient(connection)

	request := &pb.UsersRequest{}

	response, err := client.GetAll(ctx, request)

	if err != nil {
		log.Fatalf("Error during request: %v", err)
	}

	log.Println(response.Users)

	return response, err
}

func GetByNameUsers(ctx context.Context, connection *grpc.ClientConn) (*pb.UserByNameResponse, error) {
	client := pb.NewUsersServiceClient(connection)

	request := &pb.UserByNameRequest{
		Name: "João",
	}

	response, err := client.ByName(ctx, request)

	if err != nil {
		log.Fatalf("Error during request: %v", err)
	}

	log.Println(response.User)

	return response, err
}
