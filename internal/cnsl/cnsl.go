package cnsl

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// InitMessage ...
func InitMessage() {
	str1 := "#@#@#@#@#@#@#@#@#@#\n"
	str2 := "#@# HITO SERVER #@#\n"
	message := colorBlue + str1 + str2 + str1 + colorReset
	fmt.Fprintln(os.Stdout, message)
}

// ServerInfo ...
func ServerInfo(env string, port string) {
	message := fmt.Sprintf(
		"Starting in %s environment\nListening on http//:localhost:%s\n",
		env,
		port,
	)
	fmt.Fprintln(os.Stdout, message)
}

// Error ...
func Error(data interface{}) {
	_, file, line, _ := runtime.Caller(1)
	message := fmt.Sprintf(
		"%s[%s]: ERROR:%s %s%s:%d%s: %s",
		colorRed,
		time.Now().Format("2006-01-02 15:04:05"),
		colorReset,
		colorYellow,
		file,
		line,
		colorReset,
		data,
	)
	fmt.Fprintln(os.Stderr, message)
}

// Log ...
func Log(data interface{}) {
	message := fmt.Sprintf(
		"%s[%s]: LOG:%s %s",
		colorBlue,
		time.Now().Format("2006-01-02 15:04:05"),
		colorReset,
		data,
	)
	fmt.Fprintln(os.Stdout, message)
}

// Debug ...
func Debug(data interface{}) {
	message := fmt.Sprintf(
		"%s[%s]: DEBUG:%s %s",
		colorYellow,
		time.Now().Format("2006-01-02 15:04:05"),
		colorReset,
		data,
	)
	fmt.Fprintln(os.Stdout, message)
}

// Goodbye ...
func Goodbye() {
	fmt.Fprintln(os.Stdout, "\nGoodbye!")
}
