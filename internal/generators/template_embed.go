//go:build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/huandu/xstrings"
	"go/format"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type GeneratorConfig struct {
	Package    string
	RootDir    string
	OutputFile string
}

func (config GeneratorConfig) Dirs() []DirEntry {
	paths, err := ioutil.ReadDir(config.RootDir)
	if err != nil {
		panic(err)
	}
	dirs := make([]DirEntry, 0, len(paths))
	for _, path := range paths {
		if path.IsDir() {
			dirs = append(dirs, DirEntry(filepath.Join(config.RootDir, path.Name())))
		}
	}
	return dirs
}

type DirEntry string

func (e DirEntry) Base() string {
	return filepath.Base(string(e))
}

func (e DirEntry) Slug() string {
	return xstrings.ToCamelCase(e.Base())
}

func (e DirEntry) Perms() map[string]string {
	paths := map[string]string{}
	err := filepath.Walk(string(e), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		path = strings.TrimPrefix(path, filepath.Dir(string(e))+"/")

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

import (
	"embed"
	"os"
)

{{- range .Dirs }}

//go:embed "all:{{ .Base }}"
var embed{{ .Slug }} embed.FS

var {{ .Slug }} = Template{
	Name: "{{ .Base }}",
	Embed: embed{{ .Slug }},
	Modes: map[string]os.FileMode{
	{{- range $path, $perm := .Perms }}
		"{{ $path }}": {{ $perm }},
	{{- end }}
	},
}
{{- end }}
`

func main() {
	var config GeneratorConfig
	flag.StringVar(&config.Package, "package", "templates", "Package name of output template")
	flag.StringVar(&config.RootDir, "templates", "templates", "Templates directory")
	flag.StringVar(&config.OutputFile, "output", "templates/embed.go", "Output Go file path")
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
