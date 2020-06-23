package cmd

import (
	"github.com/urfave/cli"

	"{{.project_name}}/config"
)

func GetCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "serve",
			Usage:  "starts " + config.ServiceName + " workers",
			Action: serveAction,
		},
	}
}

const FlagConfig = "config"

func GetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  FlagConfig + ", c",
			Value: "./config.yaml",
		},
	}
}
