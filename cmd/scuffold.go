package cmd

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"

	"goscuffold/project"
)

func GenAction(c *cli.Context) error {
	log.Println("scaffold project")

	p := project.NewProject(c.String(FlagDomain), c.String(FlagName), c.String(FlagOutputPath), c.String(FlagTmplPath))
	err := p.Scaffold()
	if err != nil {
		return fmt.Errorf("failed to scaffold project: %s", err)
	}
	return nil
}
