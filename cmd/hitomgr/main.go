package main

import (
	"errors"
	"fmt"

	"github.com/ad9311/hitomgr/internal/cfg"
	"github.com/ad9311/hitomgr/internal/ctrl"
	"github.com/ad9311/hitomgr/internal/db"
	"github.com/ad9311/hitomgr/internal/render"
	"github.com/ad9311/hitomgr/internal/server"
)

type command struct {
	environment string
	boot        bool
	mode        string
}

const (
	environment = "-e"
	create      = "-c"
)

func main() {
	cmd, err := parseArgs()

	switch cmd.mode {
	case create:
		if err != nil {
			fmt.Println(errors.New("database could not be created"))
		} else {
			fmt.Println("database created successfully")
		}
		break
	default:
		if err != nil {
			fmt.Println(err)
		}
		break
	}

	if cmd.boot {
		fmt.Println("HITO SERVER")

		config, err := cfg.LoadConfig(cmd.environment)
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
}
