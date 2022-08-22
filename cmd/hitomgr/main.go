package main

import (
	"fmt"

	"github.com/ad9311/hitomgr/internal/cfg"
	"github.com/ad9311/hitomgr/internal/controller"
	"github.com/ad9311/hitomgr/internal/db"
	"github.com/ad9311/hitomgr/internal/render"
	"github.com/ad9311/hitomgr/internal/server"
	"github.com/ad9311/hitomgr/internal/sess"
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

	sessionData := sess.Init(config.SeverSecure)

	controller.Init(database, sessionData)

	err = render.Init(config.ServerCache, sessionData)
	if err != nil {
		fmt.Println(err)
	}

	server.Init(config, sessionData.Session)

	err = server.New().ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
