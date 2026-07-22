package main

import (
	"fmt"

	"github.com/Kwasos/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = cfg.SetUser("Kris")
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cfg)
}
