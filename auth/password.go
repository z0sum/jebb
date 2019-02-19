package main

import (
	"math/rand"
	"time"
)

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nums    = "1234567890"
)

func genPassword(n int) string {

	rand.Seed(time.Now().UTC().UnixNano())
	choice := letters + nums
	count := len(choice)

	pwd := ""
	for i := 0; i < n; i++ {
		num := rand.Intn(count)
		pwd = pwd + string(choice[num])
	}

	return pwd

}
