package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro"
	uuid "github.com/satori/go.uuid"
	proto "github.com/z0sum/jebb/auth/proto"
)

type service struct{}

var db = make(map[string]string)

func (s *service) Register(ctx context.Context, req *proto.RegisterRequest, res *proto.RegisterResponse) error {

	log.Printf("[req | Register] %v", req.Username)

	u, err := uuid.NewV4()
	if err != nil {
		log.Printf("[err | Register] %v", err)
		return err
	}

	pwd := genPassword(20)

	db[req.Username] = pwd

	*res = proto.RegisterResponse{
		Username: req.Username,
		Password: pwd,
		Id:       u.String(),
	}

	return nil
}

func (s *service) Login(ctx context.Context, req *proto.LoginRequest, res *proto.LoginResponse) error {

	log.Printf("[req | Login] %v", req.Username)

	if db[req.Username] != req.Password {
		*res = proto.LoginResponse{ LoggedIn: false }
		return nil
	}

	*res = proto.LoginResponse{ LoggedIn: true }

	return nil
}


func main() {

	srv := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Version("latest"),
	)

	srv.Init()

	proto.RegisterAuthServiceHandler(srv.Server(), &service{})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
