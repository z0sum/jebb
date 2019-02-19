package main

import (
	"io"
	"log"

	proto "github.com/z0sum/jebb/fileshare/proto"
)

type Server struct{}

var streams map[string]proto.FileshareService_ReceiveServer

func (s *Server) Send(stream proto.FileshareService_SendServer) error {

	waitc := make(chan int)

	go func() {

		for {
			
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("[MSG receiving error] %v", err)
				continue
			}

			streams[msg.Receiver].Send(msg)

		}

	}()

	<-waitc

	return nil
}

func (s *Server) Receive(req *proto.ReceiveRequest, stream proto.FileshareService_ReceiveServer) error {

	streams[req.Username] = stream

	waitc := make(chan int)

	<-waitc

	return nil

}
