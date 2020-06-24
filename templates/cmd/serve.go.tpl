package cmd

import (
	"github.com/lancer-kit/armory/log"
	"github.com/urfave/cli"

	"{{.project_name}}/config"
)

func serveAction(c *cli.Context) error {
	cfg := config.ReadConfig(c.GlobalString(FlagConfig))
	logger := log.Get().WithField("app", config.ServiceName)

    logger.WithField("cfg", cfg).Info("Serving with cfg")
	return nil
}
