package main

import (
	"context"
	"grpc-simple/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloServer) SayHello(
	ctx context.Context,
	req *pb.HelloRequest,
) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &HelloServer{})

	log.Println("gRPC server listening on :50051")
	grpcServer.Serve(lis)
}
