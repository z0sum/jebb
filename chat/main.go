package main

import (
	"fmt"
	"log"
	"net"

	proto "github.com/z0sum/jebb/chat/proto"
	"google.golang.org/grpc"
)

 

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50056")
	if err != nil {
		log.Fatalf("Failled to listen: %v", err)
	}

	fmt.Println("-- [CHAT] 0.0.0.0:50056 --")

	s := grpc.NewServer()
	proto.RegisterChatServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
