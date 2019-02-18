package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/micro/go-micro"
	proto "github.com/z0sum/jebb/chat/proto"
)

type service struct {
	hub Hub
}

func (s *service) Unregister(ctx context.Context, req *proto.UnregisterRequest, res *proto.UnregisterResponse) error {

	log.Printf("[UN-REGISTERED] %+v", req.Username)

	*res = proto.UnregisterResponse{
		Connected: false,
	}

	s.hub.unregister(req.Username)

	log.Println("users: ", s.hub.list())

	return nil
}

func (s *service) Register(ctx context.Context, req *proto.RegisterRequest, res *proto.RegisterResponse) error {

	log.Printf("[REGISTERED] %+v", req.Username)

	*res = proto.RegisterResponse{
		Connected: true,
	}

	s.hub.register(req.Username)

	log.Println("users: ", s.hub.list())

	return nil
}

// func (s *service) Socket(ctx context.Context) (proto.ChatService_SocketService, error) {

// 	return Stream{}, nil
// }

func (s *service) Socket(ctx context.Context, stream proto.ChatService_SocketStream) error {

	msg, err := stream.Recv()
	if err != nil {
		return err
	}

	s.hub.link(&stream, msg.Sender)
	defer func() {
		s.hub.unlink(msg.Sender)
		log.Printf("[STREAMS after] %v", s.hub.listStreams())
	}()

	for {

		log.Printf("[STREAMS inside] %v", s.hub.listStreams())
		_, err := stream.Recv()
		if err == io.EOF || err != nil {
			break
		}

	}

	return nil
}

func main() {

	srv := micro.NewService(
		micro.Name("go.micro.srv.chat"),
		micro.Version("latest"),
	)

	srv.Init()

	h := &hub{users: make(map[string]bool), streams: make(map[string]*proto.ChatService_SocketStream)}

	proto.RegisterChatServiceHandler(srv.Server(), &service{h})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
