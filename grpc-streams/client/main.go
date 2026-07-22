package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "client/proto/gen"
)

func main() {
	addr := "localhost:50051"

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error to create a grpc client ", err)
	}

	defer conn.Close()
	ctx := context.Background()

	calcClient := pb.NewCalculatorClient(conn)

	req := &pb.FibonacciRequest{
		Number: 10,
	}

	stream, err := calcClient.GenerateFibonacci(ctx, req)

	for {
		res, err :=  stream.Recv()
		
		if err == io.EOF {
			log.Println("End of Stream") 
			break
		}

		if err != nil {
			 log.Fatalln(err)
		}

		log.Println("Fibonacci numer: ", res.GetNumber())

	}

}
