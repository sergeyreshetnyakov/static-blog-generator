package parser

import "gopkg.in/yaml.v2"

func ParseMeta(meta []byte) (map[string]any, error) {
	var unmarshaledMeta struct {
		Title   string
		Date    string
		IsIndex bool
	}

	err := yaml.Unmarshal(meta, unmarshaledMeta)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"Title":   unmarshaledMeta.Title,
		"Date":    unmarshaledMeta.Date,
		"IsIndex": unmarshaledMeta.IsIndex,
	}, nil
}
