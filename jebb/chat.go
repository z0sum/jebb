package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	microclient "github.com/micro/go-micro/client"
	chatpb "github.com/z0sum/jebb/chat/proto"
)

func chat(user Config) {

	sig := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(sig, os.Interrupt)

	client := chatpb.NewChatService("go.micro.srv.chat", microclient.DefaultClient)

	_, err := client.Register(
		context.Background(),
		&chatpb.RegisterRequest{Username: user.Username},
	)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return
	}

	stream, err := client.Socket(context.Background())

	stream.Send(&chatpb.Message{Sender: user.Username})

	go func() {
		<-sig
		client.Unregister(
			context.Background(),
			&chatpb.UnregisterRequest{Username: user.Username},
		)
		stream.Close()
		log.Printf("[UN-REGISTERED]")
		close(done)
	}()

	log.Printf("[REGISTERED]")

	//time.Sleep(time.Second * 10)

	<-done

}
