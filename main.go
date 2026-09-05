package main

import (
	"fmt"

	"github.com/someshubham/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg.SetUser("shubham")

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cfg)
}
