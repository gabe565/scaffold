package templates

import (
	"bytes"
	"embed"
	"errors"
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

type TemplateData struct {
	appconfig.AppConfig
	ExistingData string
}

func (t Template) Generate(appConfig appconfig.AppConfig) error {
	log.Println("Processing templates " + t.Name)

	functions := sprig.TxtFuncMap()
	var buf bytes.Buffer

	err := fs.WalkDir(t.Embed, ".", func(path string, d fs.DirEntry, err error) error {
		defer func() {
			buf.Reset()
		}()

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
			outputpath = strings.TrimSuffix(outputpath, ".tmpl")

			if err = os.MkdirAll(filepath.Dir(outputpath), os.ModePerm); err != nil {
				return err
			}

			existingData, err := os.ReadFile(outputpath)
			if err != nil && !errors.Is(err, os.ErrNotExist) {
				return err
			}

			d := TemplateData{
				AppConfig:    appConfig,
				ExistingData: string(bytes.Trim(existingData, "\n")),
			}

			if err = tmpl.ExecuteTemplate(&buf, basename, d); err != nil {
				return err
			}

			if len(bytes.TrimSpace(buf.Bytes())) > 0 {
				f, err := os.OpenFile(outputpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, t.GetMode(path))
				if err != nil {
					return err
				}
				defer func(f *os.File) {
					_ = f.Close()
				}(f)

				_, err = f.Write(buf.Bytes())
				if err != nil {
					return err
				}

				return f.Close()
			}
			return nil
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
