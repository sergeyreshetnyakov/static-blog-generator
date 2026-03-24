package main

import (
	"flag"
	"log"

	"github.com/sergeyreshetnyakov/static-blog-generator/internal/app"
)

func main() {
	templatesPath := flag.String("templates", "templates", "path to templates")
	inputPath := flag.String("input", "", "path to templates")
	outputPath := flag.String("output", "builded", "path to templates")

	flag.Parse()

	a := app.New(*templatesPath, *inputPath, *outputPath, false)
	if err := a.Run(); err != nil {
		log.Fatalf("%v", err.Error())
	}
}
