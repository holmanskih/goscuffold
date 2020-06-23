package tmpl

import (
	"errors"
	"fmt"
	"io/ioutil"
)

type TemplateProvider struct {
	cfg TemplateCfg
}

func NewTemplateProvider(cfg TemplateCfg) *TemplateProvider {
	return &TemplateProvider{
		cfg: cfg,
	}
}

func (p *TemplateProvider) GetTmpl(specName string, data map[string]interface{}) (string, error) {
	spec, ok := p.cfg.Specs[specName]
	if !ok {
		return "", errors.New("unknown message schema")
	}

	var (
		dat []byte
		err error
	)

	tmplDirPrefix := p.cfg.PathPrefix
	tmplFileName := fmt.Sprintf("%s/%s", tmplDirPrefix, spec.File)

	switch p.cfg.Type {
	case TypeFileSystem:
		dat, err = ioutil.ReadFile(tmplFileName)
	case TypeEmbedded:
		dat, err = GetBindataAsset(tmplFileName)
	}

	if err != nil {
		return "", fmt.Errorf("can not read tmpl file: %s", err)
	}

	tmpl, err := NewTemplate(dat, spec.Schema)
	if err != nil {
		return "", fmt.Errorf("failed template initialization: %s", err)
	}
	tmpl.Data = data

	err = tmpl.Validate()
	if err != nil {
		return "", fmt.Errorf("failed template validation: %s", err)
	}

	strTmpl, err := tmpl.ExecuteToString()
	if err != nil {
		return "", fmt.Errorf("failed template execution to string: %s", err)
	}

	return strTmpl, nil
}

func (p *TemplateProvider) GetTmplSpecs() map[string]TemplateSpec {
	return p.cfg.Specs
}
