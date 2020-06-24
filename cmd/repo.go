package cmd

import (
	"log"

	"github.com/urfave/cli/v2"
)

func InitRepositoryAction(c *cli.Context) error {
	log.Println("git init")
	//return exec.Command("git", "init", c.String(FlagGitInit)).Run()
	return nil
}
