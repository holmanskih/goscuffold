package config

import (
	"io/ioutil"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Cfg struct {
	Path string `yml:"path"`

	Templates *TemplatesCfg `yml:"templates"`
	Git       *GitCfg       `yml:"git"`
}

func (c Cfg) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Path, validation.Required),
		validation.Field(&c.Templates, validation.Required),
		validation.Field(&c.Git, validation.Required),
	)
}

func ReadConfig(path string) Cfg {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.New().WithError(err).
			WithField("path", path).
			Fatal("unable to read config file")
	}

	config := new(Cfg)
	err = yaml.Unmarshal(rawConfig, config)
	if err != nil {
		logrus.New().WithError(err).
			WithField("raw_config", rawConfig).
			Fatal("unable to unmarshal config file")
	}

	err = config.Validate()
	if err != nil {
		logrus.New().WithError(err).
			Fatal("Invalid configuration")
	}
	return *config
}
