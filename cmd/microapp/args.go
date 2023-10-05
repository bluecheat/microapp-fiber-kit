package main

import "github.com/urfave/cli/v2"

const (
	Name    = "Micro Fiber API"
	Usage   = "Micro App Starter"
	Version = "1.0.0"
)

var Flags = []cli.Flag{
	EnvFlag,
}

var EnvFlag = &cli.StringFlag{
	Name:    "env",
	Aliases: []string{"e"},
	Value:   "env.yaml",
	Usage:   "env file path",
}
