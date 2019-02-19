package main

import (
	"context"
	"log"

	microclient "github.com/micro/go-micro/client"
	proto "github.com/z0sum/jebb/api/proto"
	authpb "github.com/z0sum/jebb/auth/proto"
)

//Server struct
type Server struct{}

func (s *Server) Signup(ctx context.Context, req *proto.SignupRequest) (*proto.SignupResponse, error) {

	log.Printf("[Signup Req] ", req)

	client := authpb.NewAuthService("go.micro.srv.auth", microclient.DefaultClient)

	res, err := client.Register(
		context.Background(),
		&authpb.RegisterRequest{Username: req.Username},
	)
	if err != nil {
		return &proto.SignupResponse{}, err
	}

	return &proto.SignupResponse{
		Username: res.Username,
		Password: res.Password,
		Id:       res.Id,
	}, nil

}

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {

	log.Printf("[Login Req] ", req)

	client := authpb.NewAuthService("go.micro.srv.auth", microclient.DefaultClient)

	res, err := client.Login(
		context.Background(),
		&authpb.LoginRequest{Username: req.Username, Password: req.Password},
	)
	if err != nil {
		return &proto.LoginResponse{LoggedIn: res.LoggedIn}, err
	}

	return &proto.LoginResponse{LoggedIn: res.LoggedIn}, nil

}
