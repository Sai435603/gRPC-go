// look after TLS
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "client/proto/gen"
)

func main() {
	addr := "localhost:50051"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("Did not Connect :", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := pb.AddRequest{
		A: 10,
		B: 20,
	}
	res, err := client.Add(ctx, &req)

	if err != nil {
		log.Fatalln("Could not add: ", err)
	}

	log.Println("sum : ",res.Sum)
}
