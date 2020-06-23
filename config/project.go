package config

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ProjectCfg struct {
	Domain string `yml:"domain"`
	Name   string `yml:"domain"`
}

func (c *ProjectCfg) ProjectName() string {
	return fmt.Sprintf("%s/%s", c.Domain, c.Name)
}

func (c ProjectCfg) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Domain, validation.Required, validation.In(DomainGitHub, DomainGitLab)),
		validation.Field(&c.Name, validation.Required),
	)
}
