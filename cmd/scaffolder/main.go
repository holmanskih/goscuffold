package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"goscuffold/config"
	"goscuffold/templates"
)

var (
	cfgPath = flag.String("cfg", "./testdata", "path to templates to scaffold")
)

func main() {
	flag.Parse()

	cfg := config.ReadConfig(*cfgPath)

	data := map[string]interface{}{
		"project_name": cfg.Project.ProjectName(),
	}

	err := RestoreTemplate("test.go", "main.go.tpl", data)
	if err != nil {
		log.Println(err)
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
