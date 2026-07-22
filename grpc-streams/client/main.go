package main

import (
	"context"
	"fmt"
	"io"
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
		log.Fatalln("Error to create a grpc client ", err)
	}

	defer conn.Close()
	ctx := context.Background()

	client := pb.NewCalculatorClient(conn)

	req := &pb.FibonacciRequest{
		Number: 10,
	}

	stream, err := client.GenerateFibonacci(ctx, req)

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			log.Println("End of Stream")
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Fibonacci numer: ", res.GetNumber())

	}

	//numgen
	stream1, err := client.SendNumbers(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	for num := range 9 {
		err := stream1.Send(&pb.NumberRequest{
			Number: int32(num),
		})
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Second * 2)
	}
	res, err := stream1.CloseAndRecv()

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server response after stream: ", res.Sum)

	//bi di stream

	stream2, err := client.Chat(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	messages := []string{"hii", "hello", "How are you", "bye"}
	for _, m := range messages {
		err := stream2.Send(&pb.ChatMessage{
			Message: m,
		})
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Second)
	}
	stream2.CloseSend()
	go func() {
		for {
			msg, err := stream2.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalln(err)
			}

			log.Println("Server : ", msg)
		}
	}()
	select {}
}
