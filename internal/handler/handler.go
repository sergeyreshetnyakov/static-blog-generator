package handler

import (
	"bytes"
	"errors"
)

func HandlePageFile(file []byte, parseMarkdown func([]byte) ([]byte, error), parseMeta func([]byte) (map[string]any, error)) (meta map[string]any, content []byte, err error) {
	parts := bytes.Split(file, []byte("---"))
	if len(parts) == 1 {
		return nil, nil, errors.New("Metadata is not found")
	} else if len(parts) > 2 {
		return nil, nil, errors.New("Sign \"---\" is reserved separator. You must not use it")
	}

	meta, err = parseMeta(parts[0])
	if err != nil {
		return nil, nil, err
	}
	content, err = parseMarkdown(parts[1])
	if err != nil {
		return nil, nil, err
	}

	return meta, content, nil
}
