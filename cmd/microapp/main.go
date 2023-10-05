package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"microapp-fiber-kit/config"
	"os"
)

// @title     MicroApp documentation
// @version   1.0.0
// @BasePath  /
func main() {
	app := &cli.App{
		Name:    Name,
		Usage:   Usage,
		Version: Version,
		Flags:   Flags,
		Action: func(c *cli.Context) error {
			// Config Load
			filename := c.String(EnvFlag.Name)
			conf := config.LoadConfigFile(filename)
			MicroApp(conf).Run()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
