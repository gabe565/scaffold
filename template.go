package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/clevyr/scaffold/appconfig"
	"github.com/markbates/pkger"
)

const templateRoot = "/templates"

func init() {
	// Include templates dir to fix dynamic paths not discovered by pkger
	_ = pkger.Include(templateRoot)
}

func generateTemplate(appConfig appconfig.AppConfig, templateDir string) (err error) {
	fmt.Printf("Processing templates: %s\n", templateDir)
	templateDir = path.Join(templateRoot, templateDir)

	functions := template.FuncMap{
		"upper": strings.ToUpper,
	}

	err = pkger.Walk(templateDir, func(filepath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode().IsRegular() {
			contents, err := readFile(filepath)
			if err != nil {
				return err
			}

			outputPath := buildOutputPath(filepath, templateDir)
			tmpl, err := template.New(outputPath).Funcs(functions).Parse(contents)
			if err != nil {
				return err
			}

			if err = os.MkdirAll(path.Dir(outputPath), os.ModePerm); err != nil {
				return err
			}

			f, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, info.Mode())
			if err != nil {
				return err
			}

			if err = tmpl.Execute(f, appConfig); err != nil {
				_ = f.Close()
				return err
			}

			return f.Close()
		}

		return nil
	})

	return
}

func readFile(filename string) (string, error) {
	f, err := pkger.Open(filename)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		f.Close()
		return "", err
	}

	if err = f.Close(); err != nil {
		return "", err
	}

	return string(b), err
}

func buildOutputPath(filepath string, templateDir string) string {
	templateDirIndex := strings.Index(filepath, templateDir) + len(templateDir) + 1
	return filepath[templateDirIndex:]
}
