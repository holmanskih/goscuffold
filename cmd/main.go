package cmd

import (
	"github.com/urfave/cli/v2"
)

func GetCommands() []*cli.Command {
	return []*cli.Command{
		scaffoldCommand(),
	}
}

func GetFlags() []cli.Flag {
	return []cli.Flag{}
}
