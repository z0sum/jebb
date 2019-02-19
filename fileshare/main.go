package main

import (
	"fmt"
	"log"
	"net"

	proto "github.com/z0sum/jebb/fileshare/proto"
	"google.golang.org/grpc"
)

 

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50057")
	if err != nil {
		log.Fatalf("Failled to listen: %v", err)
	}

	fmt.Println("FILSHARE listening 0.0.0.0:50057...")

	s := grpc.NewServer()
	proto.RegisterFileshareServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
