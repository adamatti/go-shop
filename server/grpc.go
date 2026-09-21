package main

import (
	"fmt"
	"log"
	"net"

	pb "adamatti.github.io/go-shop/grpc"
	"google.golang.org/grpc"
)

const grpcPort = 50051

func startGrpc() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterShopServer(s, &shopServer{})
	log.Printf("grpc :%d", grpcPort)
	log.Fatal(s.Serve(lis))
}
