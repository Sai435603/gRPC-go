package main

import (
	pb "server/proto/gen"
	"time"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) GenerateFibonacci(req *pb.FibonacciRequest, stream *pb.Calculator_GenerateFibonacciServer) error {
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

}
