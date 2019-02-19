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
	apipb "github.com/z0sum/jebb/api/proto"
	"google.golang.org/grpc"
)

type Config struct {
	Username string
	Password string
	Id       string
	Dir      string
}

const (
	cfgName = ".jebb.json"
)

func readConfig(u *user.User) (Config, error) {

	var cfg Config

	home := u.HomeDir
	path := home + "/" + os.Args[1] + cfgName

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

func authRegister(usr *user.User) (Config, error) {

	var cfg Config

	newUserText()

	validate := func(input string) error {
		if len(input) <= 3 {
			return errors.New("Username must be > 3 characters")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "username",
		Validate: validate,
	}
	uname, _ := prompt.Run()

	validate = func(input string) error {
		return nil
	}
	prompt = promptui.Prompt{
		Label:    "fileshare dir path (~/...)",
		Validate: validate,
	}
	dir, _ := prompt.Run()

	cfg, err := signup(uname)
	if err != nil {
		return cfg, err
	}

	//log.Println(cfg)

	cfg.Dir = dir

	saveConfig(usr, cfg)

	return cfg, nil

}

func signup(uname string) (Config, error) {

	cc, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	client := apipb.NewAuthServiceClient(cc)

	res, err := client.Signup(context.Background(), &apipb.SignupRequest{Username: uname})
	if err != nil {
		log.Printf("[Signup Res ERR] %v", err)
		return Config{}, err
	}

	log.Printf("[Signup Res] %v", res)

	return Config{
		Username: res.Username,
		Password: res.Password,
		Id:       res.Id,
	}, nil

}

func saveConfig(u *user.User, cfg Config) error {

	home := u.HomeDir
	path := home + "/" + os.Args[1] + cfgName

	json, err := json.Marshal(cfg)
	f, err := os.Create(path)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return err
	}
	defer f.Close()

	_, err = f.Write(json)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return err
	}

	return nil
}

func newUserText() {
	sleep(400)
	fmt.Println("> You have not been registered.")
	sleep(400)
	fmt.Println("> Please enter a username below")
	sleep(400)
}

func sleep(t int) {
	time.Sleep(time.Duration(t) * time.Millisecond)
}

func logIn(cfg Config) bool {

	cc, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	client := apipb.NewAuthServiceClient(cc)

	res, err := client.Login(
		context.Background(),
		&apipb.LoginRequest{Username: cfg.Username, Password: cfg.Password},
	)
	if err != nil {
		log.Printf("[Login Res ERR] %v", err)
		return false
	}

	log.Printf("[Login Res] %v", res)

	return res.LoggedIn

}
