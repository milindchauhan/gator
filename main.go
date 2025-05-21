package main

import (
	"fmt"
	"os"

	"github.com/milindchauhan/gator/internal/config"
)

func main() {
	l := len(os.Args)
	if l < 2 {
		fmt.Println("Error : not enough arguments")
		os.Exit(1)
	}

	if os.Args[1] == "login" {
		if l < 3 {
			fmt.Println("Error : not enough arguments")
			os.Exit(1)
		}

		config.SetUser(os.Args[2])
	}
}
