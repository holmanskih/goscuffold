package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"

	"goscuffold/project"
)

const (
	FlagDomain     = "domain"
	FlagName       = "name"
	FlagOutputPath = "output"
	FlagGoModules  = "gomods"
)

func scaffoldCommand() *cli.Command {
	cmd := &cli.Command{
		Name:  "gen",
		Usage: "gen scaffold",
		Action: func(c *cli.Context) error {
			log.Println("scaffolding project")

			p := project.NewProject(c.String(FlagDomain), c.String(FlagName), c.String(FlagOutputPath))
			err := p.Scaffold()
			if err != nil {
				return fmt.Errorf("failed to scaffold project: %s", err)
			}

			if c.String(FlagGoModules) != "" {
				log.Printf("running go mod init %s", c.String(FlagGoModules))
				err = execInScaffoldPath(c, "go", "mod", "init", c.String(FlagGoModules))
				if err != nil {
					return fmt.Errorf("failed to init go modules: %s", err)
				}

				log.Println("running go mod tidy")
				err = execInScaffoldPath(c, "go", "mod", "tidy")
				if err != nil {
					return fmt.Errorf("failed to tidy go modules: %s", err)
				}
			}
			return nil
		},
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
	}
	return cmd
}

func execInScaffoldPath(c *cli.Context, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = c.String(FlagOutputPath)
	return cmd.Run()
}
