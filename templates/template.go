package templates

import (
	"embed"
	"github.com/clevyr/scaffold/internal/appconfig"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Template struct {
	Name  string
	Embed embed.FS
	Modes map[string]int32
}

func (t Template) Generate(appConfig appconfig.AppConfig) error {
	log.Println("Processing templates " + t.Name)

	functions := template.FuncMap{
		"upper": strings.ToUpper,
	}

	err := fs.WalkDir(t.Embed, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		basename := filepath.Base(path)

		if info.Mode().IsRegular() {
			tmpl, err := template.New(basename).Funcs(functions).ParseFS(t.Embed, path)
			if err != nil {
				return err
			}

			outputpath := strings.TrimPrefix(path, t.Name+"/")

			if err = os.MkdirAll(filepath.Dir(outputpath), os.ModePerm); err != nil {
				return err
			}

			mode, ok := t.Modes[path]
			if !ok {
				mode = 0644
			}

			f, err := os.OpenFile(outputpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(mode))
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
	return err
}
