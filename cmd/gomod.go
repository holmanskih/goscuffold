package cmd

import (
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func InitGoModsAction(c *cli.Context) error {
	log.Println("gomods init")
	return exec.Command("go", "mod", "init", c.String(FlagGoModules)).Run()
}
