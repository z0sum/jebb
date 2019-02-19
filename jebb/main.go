package main

import (
	"log"
	"os/user"

	"github.com/manifoldco/promptui"
)

func main() {

	header(Config{})

	// registration ====
	usr, _ := user.Current()
	cfg, err := readConfig(usr)
	if err != nil {
		cfg, err = authRegister(usr)
		if err != nil {
			log.Printf("[ERROR] registration %v", err)
		}
	}

	res := logIn(cfg)
	if !res {
		log.Println("[err | Login]")
	}

	header(cfg)

	prompt := promptui.Select{
		Label: "SELECT",
		Items: []string{"fileshare", "chat"},
	}
	_, p, _ := prompt.Run()

	switch p {
	case "fileshare":
		Fileshare(cfg)
	case "chat":
		Chat(cfg)
	}

}
