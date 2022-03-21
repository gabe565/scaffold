package templates

import (
	"embed"
	"github.com/Masterminds/sprig"
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
	Modes map[string]os.FileMode
}

func (t Template) Generate(appConfig appconfig.AppConfig) error {
	log.Println("Processing templates " + t.Name)

	functions := sprig.TxtFuncMap()

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

			f, err := os.OpenFile(outputpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, t.GetMode(path))
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

func (t Template) GetMode(path string) os.FileMode {
	mode, ok := t.Modes[path]
	if !ok {
		return 0644
	}
	return mode
}
