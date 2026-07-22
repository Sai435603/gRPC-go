package main

import (
	"fmt"
	"log"
	"net"
	genpb "server/proto/gen"
	pb "server/proto/gen"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) GenerateFibonacci(req *pb.FibonacciRequest, stream pb.Calculator_GenerateFibonacciServer) error {
	n := req.Number
	a, b := 0, 1

	for i := 0; i < int(n); i++ {
		err := stream.Send(&pb.FibonacciResponse{
			Number: int32(a),
		})
		if err != nil {
			return err
		}
		a, b = b, a+b
		time.Sleep(time.Second * 2)
	}
	return nil
}

func main() {
	port := ":50051"

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalln("Server Listen failed ", err)
	}

	grpcServer := grpc.NewServer()

	genpb.RegisterCalculatorServer(grpcServer, &server{})
	fmt.Println("Hey server is running at port ", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("Failed to serve ", err)
	}
}
