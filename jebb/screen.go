package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"time"
)

func clearScreen() {
	cmd := exec.Command("clear") 
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func header(cfg Config) {

	clearScreen()

	usr, _ := user.Current()
	fmt.Println("======<[ JEBB CLI ]>=======")
	
	fmt.Println("user: ", usr.Username)
	fmt.Println("time: ", time.Now().Format(time.RFC3339))
	if cfg.Username != "" {
		fmt.Printf("login: @%v\n", cfg.Username)
	}
	fmt.Println("----------------------------")
}
