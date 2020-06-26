package cmd

import (
	"github.com/urfave/cli/v2"
)

func GetCommands() []*cli.Command {
	return []*cli.Command{
		scaffoldCommand(),
	}
}

const FlagSchemaPath = "schema"

func GetFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:   FlagSchemaPath,
			Value:  "../templates/schema.yml",
			Hidden: true,
		},
	}
}
