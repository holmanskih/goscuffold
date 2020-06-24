package cmd

import (
	"github.com/urfave/cli/v2"
)

const (
	FlagDomain     = "domain"
	FlagName       = "name"
	FlagOutputPath = "output"
	FlagGoModules  = "gomods"
)

func GetCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:   "gen",
			Usage:  "gen scaffold",
			Action: GenAction,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     FlagGoModules,
					Aliases:  []string{"p"},
					Usage:    "Initializes the go modules with module name in scaffold project",
					Required: false,
				},
				&cli.StringFlag{
					Name:     FlagOutputPath,
					Aliases:  []string{"o"},
					Usage:    "Specifies output dir to scaffold the project",
					Required: true,
					Value:    "./out",
				},
				&cli.StringFlag{
					Name:     FlagDomain,
					Aliases:  []string{"d"},
					Usage:    "Specifies project scaffold domain",
					Required: false,
					Value:    "github.com",
				},
				&cli.StringFlag{
					Name:     FlagName,
					Aliases:  []string{"n"},
					Usage:    "Specifies project scaffold name",
					Required: true,
					Value:    "scaffold/project",
				},
			},
		},
	}
}

func GetFlags() []cli.Flag {
	return []cli.Flag{}
}
