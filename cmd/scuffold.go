package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"

	"goscuffold/project"
)

func GenAction(c *cli.Context) error {
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
}

func execInScaffoldPath(c *cli.Context, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = c.String(FlagOutputPath)
	return cmd.Run()
}
