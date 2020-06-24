package main

import (
	"os"

	"github.com/lancer-kit/armory/log"
	"github.com/urfave/cli/v2"

	"goscuffold/cmd"
)

func main() {
	app := cli.NewApp()
	app.Flags = cmd.GetFlags()
	app.Commands = cmd.GetCommands()

	if err := app.Run(os.Args); err != nil {
		log.Get().WithError(err).Errorln("failed run app")
	}
}
