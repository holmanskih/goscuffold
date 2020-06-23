package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"goscuffold/config"
	"goscuffold/templates"
)

var (
	cfgPath = flag.String("cfg", "./tmpl.config.yml", "path to templates to scaffold")
)

func main() {
	flag.Parse()

	cfg := config.ReadConfig(*cfgPath)

	p := NewProject(&cfg)
	err := p.Scaffold()
	if err != nil {
		log.Fatalf("failed to copy the scaffold project: %s", err)
	}
}

func ExecuteTemplate(name string, data interface{}) (*bytes.Buffer, error) {
	asset, err := templates.Asset(name)
	if err != nil {
		return nil, err
	}

	tpl, err := template.New("").Parse(string(asset))
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, data)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func RestoreTemplate(path, name string, data interface{}) error {
	buf, err := ExecuteTemplate(name, data)
	if err != nil {
		return err
	}

	info, err := templates.AssetInfo(name)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(path), os.FileMode(0755))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, buf.Bytes(), info.Mode())
	if err != nil {
		return err
	}

	return nil
}

type Project struct {
	cfg *config.Cfg
}

func NewProject(cfg *config.Cfg) *Project {
	return &Project{cfg: cfg}
}

// Scaffold scaffolds the bindata file templates
func (p *Project) Scaffold() error {
	for _, filePath := range templates.AssetNames() {
		file := strings.TrimPrefix(filepath.ToSlash(filePath), p.cfg.Templates.Path)

		file = filepath.Join(p.cfg.Project.Path, file)

		err := RestoreTemplate(file+".go", filePath, p.cfg.Templates.Schema)
		if err != nil {
			return fmt.Errorf("failed to restore tmpl: %s", err)
		}
	}
	return nil
}
