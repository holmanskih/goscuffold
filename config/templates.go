package config

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type TemplatesCfg struct {
	Path   string            `yml:"path"`
	Schema map[string]string `yml:"schema"`
}

func (c TemplatesCfg) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Path, validation.Required),
		validation.Field(&c.Schema, validation.Required),
	)
}
