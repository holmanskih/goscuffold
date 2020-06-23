package config

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ProjectDomain string

const (
	DomainGitHub = "github.com"
	DomainGitLab = "gitlab.com"
)

type GitCfg struct {
	Domain  string `yml:"domain"`
	Project string `yml:"project"`
}

func (c *ProjectCfg) URL() string {
	return fmt.Sprintf("%s/%s", c.Domain, c.Name)
}

func (c GitCfg) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Domain, validation.Required),
		validation.Field(&c.Project, validation.Required),
	)
}
