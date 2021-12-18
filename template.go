package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/clevyr/scaffold/appconfig"
)

//go:embed templates
var templates embed.FS

func generateTemplate(appConfig appconfig.AppConfig, templateDir string) (err error) {
	fmt.Printf("Processing templates: %s\n", templateDir)

	functions := template.FuncMap{
		"upper": strings.ToUpper,
	}

	err = fs.WalkDir(templates, path.Join("templates", templateDir), func(filepath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		contents, err := templates.Open(filepath)
		if err != nil {
			return err
		}

		info, err := contents.Stat()
		if err != nil {
			return err
		}

		basename := path.Base(filepath)

		if info.Mode().IsRegular() {
			tmpl, err := template.New(basename).Funcs(functions).ParseFS(templates, filepath)
			if err != nil {
				return err
			}

			outputPath := buildOutputPath(filepath, templateDir)
			if err = os.MkdirAll(path.Dir(outputPath), os.ModePerm); err != nil {
				return err
			}

			f, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, info.Mode())
			if err != nil {
				return err
			}

			if err = tmpl.ExecuteTemplate(f, basename, appConfig); err != nil {
				_ = f.Close()
				return err
			}

			return f.Close()
		}

		return nil
	})

	return
}

func buildOutputPath(filepath string, templateDir string) string {
	templateDirIndex := strings.Index(filepath, templateDir) + len(templateDir) + 1
	return filepath[templateDirIndex:]
}
