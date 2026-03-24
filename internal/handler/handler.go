package handler

import (
	"bytes"
	"errors"

	"github.com/sergeyreshetnyakov/static-blog-generator/internal/meta"
)

func HandlePageFile(file []byte, parseMarkdown func([]byte) ([]byte, error), parseMeta func([]byte) (meta.Meta, error)) (meta meta.Meta, content []byte, err error) {
	parts := bytes.Split(file, []byte("---"))
	if len(parts) == 1 {
		return meta, nil, errors.New("Metadata is not found")
	} else if len(parts) > 2 {
		return meta, nil, errors.New("Sign \"---\" is reserved separator. You must not use it")
	}

	meta, err = parseMeta(parts[0])
	if err != nil {
		return meta, nil, err
	}
	content, err = parseMarkdown(parts[1])
	if err != nil {
		return meta, nil, err
	}

	return meta, content, nil
}
