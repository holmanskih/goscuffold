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

	// Optional flags that are used to scaffold custom project with some
	// defined workers api/db/
	FlagAPIService          = "api"
	FlagDBService           = "db"
	FlagSimpleWorkerService = "base_uwe"
)

func scaffoldCommand() *cli.Command {
	cmd := &cli.Command{
		Name:  "gen",
		Usage: "gen scaffold",
		Action: func(c *cli.Context) error {
			log.Println("scaffolding project")
			log.Println(c.String(FlagGoModules) != "")

			schema := project.ReadSchema(c.String(FlagSchemaPath))
			p := project.NewProject(c.String(FlagOutputPath), schema)

			var projectName string
			if c.String(FlagDomain) == "" {
				projectName = c.String(FlagName)
			} else {
				projectName = fmt.Sprintf("%s/%s", c.String(FlagDomain), c.String(FlagName))
			}

			scaffoldSchema := map[string]interface{}{
				"project_name":  projectName,
				"api":           c.Bool(FlagAPIService),
				"db":            c.Bool(FlagDBService),
				"simple_worker": c.Bool(FlagSimpleWorkerService),
				"service_name":  "service_name",
			}

			err := p.Scaffold(scaffoldSchema)
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
			&cli.BoolFlag{
				Name:     FlagAPIService,
				Usage:    "Specifies generation of optional API service logic",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     FlagDBService,
				Usage:    "Specifies generation of optional DB service logic",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     FlagSimpleWorkerService,
				Usage:    "Specifies generation of optional simple uwe worker logic",
				Required: false,
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
