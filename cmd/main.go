package cmd

import (
	"github.com/urfave/cli/v2"
)

const (
	FlagConfig = "config"

	FlagTmplPath   = "tmplPath"
	FlagDomain     = "domain"
	FlagName       = "name"
	FlagGitInit    = "git"
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
					Name:     FlagTmplPath,
					Aliases:  []string{"tmpl"},
					Usage:    "Specifies path of testdata directory to scaffold project from",
					Required: true,
					Value:    "./testdata",
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
				&cli.StringFlag{
					Name:     FlagOutputPath,
					Aliases:  []string{"o"},
					Usage:    "Specifies output dir to scaffold the project",
					Required: true,
					Value:    "./out",
				},
			},
		},
		{
			Name:    "gomod",
			Aliases: []string{"gomod"},
			Usage:   "Initializes the go modules with name in scaffold project",
			Action:  InitGoModsAction,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     FlagGoModules,
					Aliases:  []string{"m"},
					Usage:    "Initializes the go modules with module name in scaffold project",
					Required: true,
					Value:    "scaffold/project",
				},
			},
			SkipFlagParsing: false,
		},
	}
}

func GetFlags() []cli.Flag {
	return []cli.Flag{
		//&cli.StringFlag{
		//	Name:    FlagConfig,
		//	Aliases: []string{"c", "cfg"},
		//	Value:   "./config.yaml",
		//},
		//
		//&cli.StringFlag{
		//	Name:     FlagGitInit,
		//	Aliases:  []string{"repo"},
		//	Usage:    "Initializes the git repository with value and pushes init commit of scaffold project",
		//	Required: false,
		//},
	}
}
