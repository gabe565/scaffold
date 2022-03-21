//go:build ignore

package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

type GeneratorConfig struct {
	Package      string
	TemplatesDir string
	TemplatesVar string
	OutputFile   string
}

func (config GeneratorConfig) Perms() map[string]string {
	paths := map[string]string{}
	err := filepath.Walk(config.TemplatesDir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && info.Mode() != 0644 {
			paths[path] = fmt.Sprintf("%#o", info.Mode().Perm())
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return paths
}

const OutputTemplate = `// Code generated by internal/generators/template_embed.go; DO NOT EDIT.

package {{ .Package }}

import "embed"

//go:embed "all:{{ .TemplatesDir }}"
var {{ .TemplatesVar }} embed.FS

var {{ .TemplatesVar }}Modes = map[string]int32{
{{- range $path, $perm := .Perms }}
  "{{ $path }}": {{ $perm }},
{{- end }}
}
`

func main() {
	var config GeneratorConfig
	flag.StringVar(&config.Package, "package", "main", "Package name of output template")
	flag.StringVar(&config.TemplatesDir, "templates", "templates", "Templates directory")
	flag.StringVar(&config.TemplatesVar, "var", "templates", "Generated embed variable name")
	flag.StringVar(&config.OutputFile, "output", "template_embed.go", "Output Go file path")
	flag.Parse()

	out, err := os.OpenFile(config.OutputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	tpl, err := template.New("output").Parse(OutputTemplate)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer

	err = tpl.Execute(io.Writer(&buf), config)
	if err != nil {
		panic(err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	_, err = out.Write(formatted)
	if err != nil {
		panic(err)
	}

	err = out.Close()
	if err != nil {
		panic(err)
	}
}
