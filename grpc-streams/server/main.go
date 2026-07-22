package main

import (
	"fmt"
	"io"
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

func (s *server) SendNumbers(stream pb.Calculator_SendNumbersServer) error {
	sum := 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.NumberResponse{
				Sum: int32(sum),
			})
		}

		if err != nil {
			log.Fatalln(err)
		}
		log.Println(req, req.GetNumber())

		sum += int(req.GetNumber())
	}
}

func (s *server) Chat(stream pb.Calculator_ChatServer) error {
	for {
		//receiving part

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalln(err)
			return err
		}
		log.Println("Client Said ", req.GetMessage())

		//sending part
		err = stream.Send(&pb.ChatMessage{
			Message: "Server replied: " + req.GetMessage(),
		})

		if err != nil {
			log.Fatalln(err)
			return  err
		}

	}
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
