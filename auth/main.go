package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro"
	uuid "github.com/satori/go.uuid"
	"github.com/sethvargo/go-password/password"
	proto "github.com/z0sum/jebb/auth/proto"
)

type service struct{}

func (s *service) Register(ctx context.Context, req *proto.RegisterRequest, res *proto.RegisterResponse) error {

	log.Printf("[RECEIVED MSG] %+v", *req)

	pwd, err := password.Generate(10, 5, 5, true, true)
	if err != nil {
		return err
	}

	u, err := uuid.NewV4()
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return err
	}

	*res = proto.RegisterResponse{
		Username: req.Username,
		Password: pwd,
		Id:       u.String(),
	}

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
