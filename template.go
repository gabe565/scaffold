package main

import (
	"github.com/Masterminds/sprig"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

var templatesDir = "templates"
var outputDir = "out"

func generateTemplate(appConfig AppConfig) {
	err := filepath.Walk(templatesDir, func(inputPath string, info os.FileInfo, inErr error) (err error) {
		if inErr != nil {
			panic(inErr)
		}

		if info != nil && info.Mode().IsRegular() {
			var tmpl *template.Template
			tmpl, err = template.New(path.Base(inputPath)).Funcs(template.FuncMap(sprig.FuncMap())).ParseFiles(inputPath)
			if err != nil {
				panic(err)
			}
			outputPath := strings.Replace(inputPath, templatesDir, outputDir, 1)
			err = os.MkdirAll(path.Dir(outputPath), os.ModePerm)
			file, err := os.Create(outputPath)
			if err != nil {
				panic(err)
			}
			err = tmpl.Execute(file, appConfig)
			if err != nil {
				panic(err)
			}
			err = file.Close()
			if err != nil {
				panic(err)
			}
		}
		return
	})

	if err != nil {
		panic(err)
	}
}
