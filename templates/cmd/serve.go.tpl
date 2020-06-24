package cmd

import (
	"github.com/lancer-kit/armory/log"
	"github.com/urfave/cli"

	"{{.project_name}}/config"
	"{{.project_name}}/daemons"
)

func serveAction(c *cli.Context) error {
	cfg := config.ReadConfig(c.GlobalString(FlagConfig))
	logger := log.Get().WithField("app", config.ServiceName)

	chief := daemons.InitChief(logger, cfg)
	chief.Run()
	return nil
}
