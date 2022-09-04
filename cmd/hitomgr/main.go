package main

import (
	"github.com/ad9311/hitomgr/internal/api/apiv1"
	"github.com/ad9311/hitomgr/internal/cfg"
	"github.com/ad9311/hitomgr/internal/cnsl"
	"github.com/ad9311/hitomgr/internal/ctrl"
	"github.com/ad9311/hitomgr/internal/db"
	"github.com/ad9311/hitomgr/internal/render"
	"github.com/ad9311/hitomgr/internal/server"
)

type command struct {
	environment string
	load        bool
	boot        bool
	mode        string
}

const (
	environment = "-e"
)

func main() {
	cmd, err := parseArgs()
	if err != nil {
		cnsl.Error(err)
	}

	if cmd.load {
		cnsl.InitMessage()
		cmd.boot = true

		config, err := cfg.LoadConfig(cmd.environment)
		if err != nil {
			cmd.boot = false
			cnsl.Error(err)
		}

		database, err := db.New(config.DatabaseURL)
		if err != nil {
			cmd.boot = false
			cnsl.Error(err)
		}

		if err := render.SetUp(config.ServerCache); err != nil {
			cmd.boot = false
			cnsl.Error(err)
		}

		if cmd.boot {
			setupCloseHandler()

			session := server.SetUp(config.ServerPort, config.SeverSecure)
			ctrl.SetUp(database, session)
			apiv1.Setup(database)

			cnsl.ServerInfo(cmd.environment, config.ServerPort)
			cnsl.Error(server.New().ListenAndServe())
		}
	}
}
