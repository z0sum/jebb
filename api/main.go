package main

import (
	"fmt"
	"log"
	"net"

	proto "github.com/z0sum/jebb/api/proto"
	"google.golang.org/grpc"
)

 

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50055")
	if err != nil {
		log.Fatalf("Failled to listen: %v", err)
	}

	fmt.Println("-- listening on 0.0.0.0:50055 --")

	s := grpc.NewServer()
	proto.RegisterAuthServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
