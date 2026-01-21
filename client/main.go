package main

import (
	"context"
	"grpc-simple/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(
		"127.0.0.1:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{
		Name: "Gopher",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Message)
}
