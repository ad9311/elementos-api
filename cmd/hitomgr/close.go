package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ad9311/hitomgr/internal/cnsl"
)

func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cnsl.Log("CTRL+C signal detected\nClosing application...")
		cnsl.Goodbye()
		os.Exit(0)
	}()
}
