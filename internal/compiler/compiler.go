package compiler

import (
	"html/template"
	"io"
	"path/filepath"
)

type Page struct {
	Data map[string]any
}

type Compiler struct {
	TemplatesPath string
}

func New(templatesPath string, outputPath string) *Compiler {
	return &Compiler{TemplatesPath: templatesPath}
}

func (c *Compiler) NewPage(meta map[string]any, content []byte) (*Page, error) {
	meta["Content"] = template.HTML(content)

	return &Page{Data: meta}, nil
}

func (c *Compiler) CompilePage(page *Page, out io.Writer) error {
	tmpl, err := template.ParseFiles(filepath.Join(c.TemplatesPath, "page.html"))
	if err != nil {
		return err
	}

	err = tmpl.Execute(out, page.Data)

	return err
}
