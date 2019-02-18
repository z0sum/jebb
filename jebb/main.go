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
		cfg = authRegister(usr)
	}

	header(cfg)

	prompt := promptui.Select{
		Label: "|",
		Items: []string{"fileshare", "chat"},
	}
	_, res, _ := prompt.Run()
	switch res {
	case "fileshare":
		log.Println("fileshare")
	case "chat":
		chat(cfg)
	}

}
