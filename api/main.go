package main

import (
	"context"
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
	result := "Hello " + req.GetName()

	log.Println("Received request:", req.GetName())

	res := &pb.HelloResponse{
		Msg: result,
	}

	return res, nil
}

func (s *server) GetAll(ctx context.Context, req *pb.UsersRequest) (*pb.UsersResponse, error) {
	res := &pb.UsersResponse{
		Users: users,
	}

	return res, nil
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
