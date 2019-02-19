package main

import "fmt"

func Fileshare(cfg Config) {
	fileshareHeader(cfg)
}

func fileshareHeader(cfg Config) {
	header(cfg)
	fmt.Println("         FILESHARE         ")
	fmt.Println(cfg.Dir)
	fmt.Println("----------------------------")
}
