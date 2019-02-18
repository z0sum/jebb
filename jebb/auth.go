package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/manifoldco/promptui"
	microclient "github.com/micro/go-micro/client"
	authpb "github.com/z0sum/jebb/auth/proto"
)

type Config struct {
	Username string
	Password string
	Id       string
}

const (
	cfgName = ".jebb.json"
)

func readConfig(u *user.User) (Config, error) {

	var cfg Config

	home := u.HomeDir
	path := home + "/" + cfgName

	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil || len(b) == 0 {
		return cfg, err
	}

	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func authRegister(usr *user.User) Config {

	var cfg Config

	newUserText()

	validate := func(input string) error {
		if len(input) <= 5 {
			return errors.New("Username must be > 5 characters")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "username",
		Validate: validate,
	}
	username, _ := prompt.Run()

	client := authpb.NewAuthService("go.micro.srv.auth", microclient.DefaultClient)

	r, err := client.Register(context.Background(), &authpb.RegisterRequest{Username: username})
	if err != nil {
		return cfg
	}

	cfg = Config{
		Username: r.Username,
		Password: r.Password,
		Id:       r.Id,
	}

	saveConfig(usr, &cfg)

	return cfg

}

func saveConfig(u *user.User, cfg *Config) {

	home := u.HomeDir
	path := home + "/" + cfgName

	json, err := json.Marshal(*cfg)
	f, err := os.Create(path)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return
	}
	defer f.Close()

	_, err = f.Write(json)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return
	}

	log.Printf("[CREATED CONFIG] %+v", cfg)
}

func newUserText() {
	sleep(500)
	fmt.Println("> You have not been registered.")
	sleep(1000)
	fmt.Println("> Please enter a username below")
	sleep(1000)
}

func sleep(t int) {
	time.Sleep(time.Duration(t) * time.Millisecond)
}
