package main

import (
	"fmt"

	"github.com/ad9311/elementos_mgr/internal/cfg"
	"github.com/ad9311/elementos_mgr/internal/server"
)

func main() {
	fmt.Println("Elementos API")

	config, err := cfg.LoadConfig("development")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(config)
	}

	server.SetUp(config)

	err = server.New().ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
