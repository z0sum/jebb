package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	chatpb "github.com/z0sum/jebb/chat/proto"
	"google.golang.org/grpc"
)

//Chat service
func Chat(cfg Config) {

	chatHeader(cfg)

	cc, err := grpc.Dial("localhost:50056", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	client := chatpb.NewChatServiceClient(cc)

	defer func() {
		logout(cfg, client)
	}()

	res := login(cfg, client)
	if !res {
		log.Printf("[err | LoginRes]")
		return
	}

	fmt.Println(": connected \ttrue")

	online(client)

	fmt.Println("----------------------------")

	socket(cfg, client)

}

func chatHeader(cfg Config) {
	header(cfg)
	fmt.Println("        CHAT SERVICE        ")
	fmt.Println("----------------------------")
}

func login(cfg Config, client chatpb.ChatServiceClient) bool {

	res, err := client.Login(
		context.Background(),
		&chatpb.LoginRequest{Username: cfg.Username, Password: cfg.Password},
	)
	if err != nil {
		return false
	}

	return res.LoggedIn

}

func online(client chatpb.ChatServiceClient) {

	res, err := client.Online(
		context.Background(),
		&chatpb.OnlineRequest{},
	)
	if err != nil {
		fmt.Println(": users  <none>")
		return
	}

	fmt.Println(": users ", "\t"+strings.Join(res.Users, " \n\t\t"))

}

func logout(cfg Config, client chatpb.ChatServiceClient) {

	res, err := client.Logout(
		context.Background(),
		&chatpb.LogoutRequest{Username: cfg.Username},
	)
	if err != nil || !res.LoggedOut {
		log.Println("[err | LogoutReq] failed to logout")
		return
	}

	fmt.Println(": logged out")
}

func socket(cfg Config, client chatpb.ChatServiceClient) {

	stream, err := client.Socket(context.Background())
	if err != nil {
		log.Printf("[err | Socket conn] %v", err)
		return
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	waitc := make(chan int, 1)
	defer func() {
		stream.CloseSend()
	}()

	stream.Send(&chatpb.Message{
		Sender:      cfg.Username,
		MessageType: 1,
	})

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Printf("[err | Socket Recv] %v", err)
				continue
			}
			if msg.Sender == cfg.Username {
				continue
			}
			pvt := " "
			if msg.Receiver == cfg.Username {
				pvt = "*"
			}
			fmt.Printf("\r%v%v< %v\n", pvt, msg.Sender[:2], msg.Msg)
			fmt.Print("   > ")
		}
	}()

	go func() {
		
		for {

			text, recvr := readInput()

			switch text {
			case "\\exit":
				close(waitc)
			case "\\online":
				online(client)
			default:
				err := stream.Send(&chatpb.Message{
					Msg:         text,
					MessageType: 2,
					Sender: 	cfg.Username,
					Receiver: 	recvr,
				})
				if err != nil {
					log.Printf("[err | Socket] %v", err)
				}
			}
		}
		
	}()

	<-waitc

}

func readInput() (string, string) {
	fmt.Print("   > ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	split := strings.Split(text, " ")
	receiver := ""
	content := text
	if string(split[0][0]) == "@" {
		receiver = string(split[0][1:])
		content = strings.Join(split[1:], " ")
	}
	return content, receiver
}
