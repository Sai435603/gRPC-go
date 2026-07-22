package main

import (
	"context"
	"fmt"
	genpb "greeter/client/proto/gen"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	//create a new grpc client
	cert := "cert.pem"
	addr := "localhost:50051"
	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalln("failed to load the credentials", err)
	}
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalln("failed to create a new grpc client", err)
	}
	defer conn.Close()

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	client := genpb.NewGreeterClient(conn)
	req := &genpb.HelloRequest{
		Name: "sai",
	}
	res, err := client.Greet(ctx, req)
	if err != nil {
		log.Fatalln("Couldn't greet")
	}
	fmt.Println("you got the greet: ", res.Message)
}
