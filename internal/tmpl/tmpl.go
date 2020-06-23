package tmpl

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"text/template"

	"goscuffold/env/templates"
)

var (
	ErrValidationEmptySchema = errors.New("empty schema field name")
	ErrValidationParseStr    = errors.New("tmpl field is not of valid string type")
	ErrValidationParseNumber = errors.New("failed to parse number field")
	ErrValidationParseURL    = errors.New("tmpl field is not of valid url type")
)

type FieldType string

const (
	String FieldType = "string"
	Number FieldType = "number"
	HTML   FieldType = "html"
	URL    FieldType = "url"
)

type Template struct {
	tmpl   *template.Template
	schema map[string]FieldType
	Data   map[string]interface{}
}

func NewTemplate(fileBody []byte, schema map[string]FieldType) (*Template, error) {
	tmpl, err := template.New("template").Parse(string(fileBody))
	if err != nil {
		return nil, fmt.Errorf("can not create a new template: %s", err)
	}

	return &Template{
		tmpl:   tmpl,
		schema: schema,
		Data:   map[string]interface{}{},
	}, nil
}

func GetBindataAsset(tmplName string) ([]byte, error) {
	asset, err := templates.Asset(tmplName)
	if err != nil {
		return nil, fmt.Errorf("failed to load bindata asset: %s", err)
	}

	return asset, nil
}

func (t *Template) Validate() error {
	for varName, varValue := range t.Data {
		validType, ok := t.schema[varName] // get the validType by varName
		if !ok {
			return ErrValidationEmptySchema
		}

		switch validType {
		case String, HTML:
			temp, ok := varValue.(string)
			if !ok {
				return ErrValidationParseStr
			}

			if temp == "" {
				return ErrValidationParseStr
			}

		case Number:
			temp := fmt.Sprintf("%v", varValue)
			_, err := strconv.ParseInt(temp, 10, 64)
			if err == nil {
				continue
			}
			_, err = strconv.ParseFloat(temp, 64)
			if err == nil {
				continue
			}
			return ErrValidationParseNumber

		case URL:
			temp, ok := varValue.(string)
			if !ok {
				return ErrValidationParseURL
			}

			if temp == "" {
				return ErrValidationParseURL
			}

			_, err := url.ParseRequestURI(temp)
			if err != nil {
				return ErrValidationParseURL
			}
		}
	}

	return nil
}

func (t *Template) Execute(w io.Writer) error {
	return t.tmpl.Execute(w, t.Data)
}

func (t *Template) ExecuteToString() (string, error) {
	var buf bytes.Buffer
	if err := t.Execute(&buf); err != nil {
		return "", err
	}

	tmpl := buf.String()
	return tmpl, nil
}
