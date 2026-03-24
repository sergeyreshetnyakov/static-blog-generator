package compiler

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/sergeyreshetnyakov/static-blog-generator/internal/meta"
)

type Page struct {
	Meta    meta.Meta
	Content template.HTML
}

type Compiler struct {
	TemplatesPath string
}

func New(templatesPath string, outputPath string) *Compiler {
	return &Compiler{TemplatesPath: templatesPath}
}

func (c *Compiler) NewPage(meta meta.Meta, content []byte) (*Page, error) {
	return &Page{Meta: meta, Content: template.HTML(content)}, nil
}

func (c *Compiler) CompilePage(page *Page, out io.Writer) error {
	tmpl, err := template.ParseFiles(filepath.Join(c.TemplatesPath, page.Meta.Template))
	if err != nil {
		return err
	}

	pageMap := map[string]any{
		"Title":    page.Meta.Title,
		"Date":     page.Meta.Date,
		"Template": page.Meta.Template,
		"Style":    page.Meta.Style,
		"Content":  page.Content,
	}
	err = tmpl.Execute(out, pageMap)

	return err
}
