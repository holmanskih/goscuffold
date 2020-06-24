package main

import (
	"os"

	"github.com/lancer-kit/armory/log"
	"github.com/urfave/cli"

	"{{.project_name}}/cmd"
	"{{.project_name}}/config"
	"{{.project_name}}/info"
)

func main() {
	app := cli.NewApp()
	app.Usage = "A " + config.ServiceName + " service"
	app.Version = info.App.Version + ":" + info.App.Build
	app.Flags = cmd.GetFlags()
	app.Commands = cmd.GetCommands()

	if err := app.Run(os.Args); err != nil {
		log.Get().WithError(err).Errorln("failed run app")
	}
}
