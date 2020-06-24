package project

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Project struct {
	domain  string
	name    string
	outPath string

	tmplPath string
}

func NewProject(domain, name, outPath, tmplPath string) *Project {
	return &Project{
		domain:   domain,
		name:     name,
		outPath:  outPath,
		tmplPath: tmplPath,
	}
}

// Scaffold scaffolds the bindata file tmpl
func (p *Project) Scaffold() error {
	for _, fileName := range AssetNames() {
		file := strings.TrimPrefix(filepath.ToSlash(fileName), p.tmplPath)
		file = filepath.Join(p.outPath, file)

		genPath := strings.TrimSuffix(file, filepath.Ext(file))

		schema := map[string]interface{}{
			"project_name": fmt.Sprintf("%s/%s", p.domain, p.name), // fixme when domain is null
		}
		err := RestoreTemplate(genPath, fileName, schema)
		if err != nil {
			return fmt.Errorf("failed to restore tmpl: %s", err)
		}
	}
	return nil
}

func ExecuteTemplate(name string, data interface{}) (*bytes.Buffer, error) {
	asset, err := Asset(name)
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

	info, err := AssetInfo(name)
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
