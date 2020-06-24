package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"goscuffold/cmd"
)

func main() {
	app := cli.NewApp()
	app.Flags = cmd.GetFlags()
	app.Commands = cmd.GetCommands()

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("failed to run the app")
	}
}
