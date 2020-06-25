package project

import (
	"io/ioutil"
	"log"

	"github.com/go-ozzo/ozzo-validation"

	"gopkg.in/yaml.v2"
)

type tmplDirSchema struct {
	Path string `yml:"path"`
}

func (s tmplDirSchema) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Path, validation.Required),
	)
}

type Schema struct {
	// Target defines the target directory info that is used by Scaffolder
	// to build all optional modules to its root path.
	Target tmplDirSchema `yml:"target"`

	// Modules defines an optional service directories with the same directory
	// name for mapping the directories in base directory.
	Modules map[string]tmplDirSchema `yml:"modules"`
}

func (s Schema) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Target, validation.Required),
		validation.Field(&s.Modules, validation.Required),
	)
}

func ReadSchema(path string) Schema {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("unable to read config file with path %s: %s", path, err)
	}

	config := new(Schema)
	err = yaml.Unmarshal(rawConfig, config)
	if err != nil {
		log.Fatalf("unable to unmarshal config file with raw config %s: %s", rawConfig, err)
	}

	err = config.Validate()
	if err != nil {
		log.Fatalf("invalid configuration: %s", err)
	}

	return *config
}
