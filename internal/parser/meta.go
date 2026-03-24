package parser

import (
	"html/template"

	"gopkg.in/yaml.v2"
)

func ParseMeta(meta []byte) (map[string]any, error) {
	var unmarshaledMeta struct {
		Title    string `yaml:"title"`
		Date     string `yaml:"date"`
		Template string `yaml:"template"`
		Content  template.HTML
	}

	err := yaml.Unmarshal(meta, &unmarshaledMeta)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"Title":    unmarshaledMeta.Title,
		"Date":     unmarshaledMeta.Date,
		"Template": unmarshaledMeta.Template,
	}, nil
}
