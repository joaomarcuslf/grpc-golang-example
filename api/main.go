package main

import (
	"context"
	"fmt"
	"golang_api/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

var users = []*pb.User{
	{
		Name: "João",
		Id:   0,
	},
}

func (s *server) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("Received request:", req)

	result := "Hello " + req.GetName()

	res := &pb.HelloResponse{
		Msg: result,
	}

	return res, nil
}

func (s *server) GetAll(ctx context.Context, req *pb.UsersRequest) (*pb.UsersResponse, error) {
	log.Println("Received request:", req)

	res := &pb.UsersResponse{
		Users: users,
	}

	return res, nil
}

func (s *server) ByName(ctx context.Context, req *pb.UserByNameRequest) (*pb.UserByNameResponse, error) {
	log.Println("Received request:", req)

	for _, u := range users {
		if u.Name == req.GetName() {
			res := &pb.UserByNameResponse{
				User: u,
			}

			return res, nil
		}
	}

	return nil, fmt.Errorf("User not found")
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:5000")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterHelloServiceServer(grpcServer, &server{})
	pb.RegisterUsersServiceServer(grpcServer, &server{})

	log.Printf("Listening on %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC %v", err)
	}
}
