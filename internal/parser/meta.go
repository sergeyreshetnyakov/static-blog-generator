package parser

import (
	"github.com/sergeyreshetnyakov/static-blog-generator/internal/meta"
	"gopkg.in/yaml.v2"
)

func ParseMeta(data []byte) (meta.Meta, error) {
	var unmarshaledMeta meta.Meta

	err := yaml.Unmarshal(data, &unmarshaledMeta)
	if err != nil {
		return meta.Meta{}, err
	}

	return unmarshaledMeta, nil
}
