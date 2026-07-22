package main

import (
	"context"
	"fmt"
	pb "greeter/server/proto/gen"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hey! how are you doing %s ", req.Name),
	}, nil
}

func main() {
	//create a listener
	//create a new grpc-server
	//serve that server with listener
	//implement the unimplementedgreeter

	//later we need to register the umimplemented rpc server
	// implement the unimplemented function above main..
	port := ":50051"
	cert := "cert.pem"
	key := "key.pem"
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalln("Listening Failed: ", err)
	}
	fmt.Println("gRPC Server Started and Listening at port", port)
	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalln("failed to load credentials ", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	//todo
	pb.RegisterGreeterServer(grpcServer, &server{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("Grpc Server failed: ", err)
	}
}
