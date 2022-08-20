package main

import (
	"fmt"

	"github.com/ad9311/elementos_mgr/internal/cfg"
	"github.com/ad9311/elementos_mgr/internal/db"
	"github.com/ad9311/elementos_mgr/internal/server"
)

func main() {
	fmt.Println("Elementos API")

	config, err := cfg.LoadConfig("development")
	if err != nil {
		fmt.Println(err)
	}

	conn, err := db.New(config.DatabaseURL)
	if err != nil {
		fmt.Println(err)
	}

	server.SetUp(config, conn)

	err = server.New().ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
