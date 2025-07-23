package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/milindchauhan/gator/internal/config"
)

type State struct {
}

func main() {
	l := len(os.Args)
	if l < 2 {
		fmt.Println("Error : not enough arguments")
		os.Exit(1)
	}

	val := ""
	if os.Args[1] == "login" {
		if l < 3 {
			fmt.Println("Error : not enough arguments")
			os.Exit(1)
		}

		val = config.SetUser(os.Args[2])
	}

	fmt.Println(val)
}
