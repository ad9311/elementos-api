package main

import (
	"errors"
	"os"
)

func parseArgs() (command, error) {
	args := os.Args[1:]
	cmd := command{}

	if len(args) >= 1 {
		switch args[0] {
		case environment:
			env, err := selectEnvironment(args)
			if err != nil {
				return cmd, err
			}

			cmd.environment = env
			cmd.load = true
			cmd.mode = environment

			return cmd, nil

		default:
			return cmd, errors.New("undefined flag")
		}
	}

	return cmd, errors.New("too few arguments")
}

func selectEnvironment(args []string) (string, error) {
	switch len(args) {
	case 2:
		return args[1], nil
	case 1:
		return "", errors.New("environment is missing")
	default:
		return "", errors.New("too many arguments")
	}
}
