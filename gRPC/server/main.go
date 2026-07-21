package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "server/proto/gen"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.A + req.B
	fmt.Println("Sum is : ", sum)
	return &pb.AddResponse{
		Sum: sum,
	}, nil
}

func main() {
	port := ":50051"

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalln("Failed to listen :", err)
	}

	grpcServer := grpc.NewServer()
	//Todo
	pb.RegisterCalculatorServer(grpcServer, &server{})
	log.Println("Server is running on the port ", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("Failed to serve: ", err)
	}
}
