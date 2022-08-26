package main

import (
	"fmt"

	"github.com/ad9311/hitomgr/internal/cfg"
	"github.com/ad9311/hitomgr/internal/ctrl"
	"github.com/ad9311/hitomgr/internal/db"
	"github.com/ad9311/hitomgr/internal/render"
	"github.com/ad9311/hitomgr/internal/server"
)

func main() {
	fmt.Println("HITO SERVER")

	config, err := cfg.LoadConfig("development")
	if err != nil {
		fmt.Println(err)
	}

	database, err := db.New(config.DatabaseURL)
	if err != nil {
		fmt.Println(err)
	}

	if err := render.SetUp(config.ServerCache); err != nil {
		fmt.Println(err)
	}

	session := server.SetUp(config.ServerPort, config.SeverSecure)
	ctrl.SetUp(database, session)

	fmt.Println(server.New().ListenAndServe())
}
