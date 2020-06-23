package main

import (
	"flag"
	"log"

	"goscuffold/config"
	"goscuffold/project"
)

var (
	cfgPath = flag.String("cfg", "./tmpl.config.yml", "path to templates to scaffold")
)

func main() {
	flag.Parse()

	cfg := config.ReadConfig(*cfgPath)

	p := project.NewProject(&cfg)
	err := p.Scaffold()
	if err != nil {
		log.Fatalf("failed to copy the scaffold project: %s", err)
	}
}
