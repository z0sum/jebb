package main

import (
	"context"
	"io"
	"log"

	microclient "github.com/micro/go-micro/client"
	authpb "github.com/z0sum/jebb/auth/proto"
	proto "github.com/z0sum/jebb/chat/proto"
)

type Server struct{}

/*
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Online(context.Context, *OnlineRequest) (*OnlineResponse, error)
	Socket(ChatService_SocketServer) error
*/

var users = make(map[string]bool)
var streams = make(map[string]proto.ChatService_SocketServer)

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {

	client := authpb.NewAuthService("go.micro.srv.auth", microclient.DefaultClient)

	res, err := client.Login(
		context.Background(),
		&authpb.LoginRequest{Username: req.Username, Password: req.Password},
	)
	if err != nil {
		return &proto.LoginResponse{LoggedIn: res.LoggedIn}, err
	}
	log.Println(res.LoggedIn, err)

	users[req.Username] = true

	return &proto.LoginResponse{LoggedIn: res.LoggedIn}, nil

}

func (s *Server) Online(ctx context.Context, req *proto.OnlineRequest) (*proto.OnlineResponse, error) {

	r := []string{}
	for u := range users {
		r = append(r, u)
	}

	return &proto.OnlineResponse{
		Users: r,
	}, nil

}

func (s *Server) Socket(stream proto.ChatService_SocketServer) error {

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

			switch msg.MessageType {
			case 1:
				
				log.Printf("[Link received] %+v", msg)
				
				streams[msg.Sender] = stream
			
			case 2:
				
				log.Printf("[MSG received] %+v", msg)
				
				for u, str := range streams {
					if msg.Receiver == "" || msg.Receiver == u {
						str.Send(msg)
					}
				}
			}

		}

	}()

	// go func() {
	// 	for {
	// 		msg, err := stream.Recv()
	// 		if err != nil {
	// 			log.Printf("[MSG receiving error] %v", err)
	// 			continue
	// 		}
	// 		log.Printf("[MSG received] %+v", msg)
	// 	}
	// }()

	<-waitc

	return nil

}

func (s *Server) Logout(ctx context.Context, req *proto.LogoutRequest) (*proto.LogoutResponse, error) {

	if _, ok := users[req.Username]; ok {
		delete(users, req.Username)
	}

	return &proto.LogoutResponse{
		LoggedOut: true,
	}, nil

}
