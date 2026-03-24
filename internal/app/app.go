package app

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/sergeyreshetnyakov/static-blog-generator/internal/compiler"
	"github.com/sergeyreshetnyakov/static-blog-generator/internal/handler"
	"github.com/sergeyreshetnyakov/static-blog-generator/internal/lib/utils"
	"github.com/sergeyreshetnyakov/static-blog-generator/internal/parser"
)

type App struct {
	TemplatesPath string
	InputPath     string
	OutputPath    string
	Watcher       bool
}

func New(templatePath, inputPath, outputPath string, watcher bool) *App {
	return &App{
		TemplatesPath: templatePath,
		InputPath:     inputPath,
		OutputPath:    outputPath,
		Watcher:       watcher,
	}
}

func (a *App) Run() error {
	c := compiler.New(a.TemplatesPath, a.OutputPath)
	err := a.processFiles(c)
	return err
}

func (a *App) processFiles(compiler *compiler.Compiler) error {
	return filepath.Walk(a.InputPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("%v:%w", path, err)
		}

		if filepath.Ext(path) == ".css" {
			file, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			outDir := utils.ReplacePathRootDir(path, a.OutputPath)
			if err := os.MkdirAll(outDir, 0755); err != nil {
				return fmt.Errorf("%v:%w", path, err)
			}
			// fmt.Println(outDir)
			if err := os.WriteFile(filepath.Join(outDir, info.Name()), file, 0755); err != nil {
				return fmt.Errorf("%v:%w", path, err)
			}
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}

		file, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("%v:%w", path, err)
		}

		outDir := utils.ReplacePathRootDir(path, a.OutputPath)
		if err := os.MkdirAll(outDir, 0755); err != nil {
			return fmt.Errorf("%v:%w", path, err)
		}

		outPath := filepath.Join(outDir, strings.Replace(info.Name(), ".md", ".html", 1))

		out, err := os.Create(outPath)
		if err != nil {
			return fmt.Errorf("%v:%w", path, err)
		}

		err = processPage(file, compiler, out)
		if err != nil {
			return fmt.Errorf("%v:%w", path, err)
		}
		return nil
	})
}

func processPage(file []byte, compiler *compiler.Compiler, out io.Writer) error {
	meta, content, err := handler.HandlePageFile(file, parser.ParseMarkdown, parser.ParseMeta)

	page, err := compiler.NewPage(meta, content)
	if err != nil {
		return err
	}

	err = compiler.CompilePage(page, out)
	return err
}
