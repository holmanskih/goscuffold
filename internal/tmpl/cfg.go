package tmpl

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	TypeHTML       = "html"
	TypeText       = "text"
	TypeEmbedded   = "embedded"
	TypeFileSystem = "filesystem"
)

type TemplateCfg struct {
	Type       string                  `yaml:"type"`        // embedded | filesystem
	PathPrefix string                  `yaml:"path_prefix"` // path to directory with templates. by default ./templates
	Specs      map[string]TemplateSpec `yaml:"specs"`       // template name => template specification
}

type TemplateSpec struct {
	File   string               `yaml:"file"`
	Type   string               `yaml:"type"`   // html | text
	Schema map[string]FieldType `yaml:"schema"` // key => key_type
}

func (s TemplateSpec) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.File, validation.Required),
		validation.Field(&s.Type, validation.Required,
			validation.In(TypeHTML, TypeText)),
		validation.Field(&s.Schema, validation.Required),
	)
}

func (cfg TemplateCfg) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.Type, validation.Required,
			validation.In(TypeEmbedded, TypeFileSystem)),
		validation.Field(&cfg.PathPrefix, validation.Required),
		validation.Field(&cfg.Specs, validation.Required),
	)
}
