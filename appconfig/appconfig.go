package appconfig

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/clevyr/scaffold/config/modules"
	"github.com/clevyr/scaffold/module"
)

type AppConfig struct {
	InitLaravel   bool `json:"-"`
	AppName       string
	AppSlug       string `json:"-"`
	AppKey        string
	Database      string
	PhpModules    module.Map
	AdminGen      string
	ComposerDeps  module.Map
	NpmDeps       module.Map
	MaxUploadSize string
}

var Defaults = AppConfig{
	Database:      "PostgreSQL",
	PhpModules:    modules.Php(),
	AdminGen:      "Nova",
	ComposerDeps:  modules.Composer(),
	NpmDeps:       modules.Npm(),
	MaxUploadSize: "64m",
}

func (appConfig *AppConfig) GenerateAppKey() (err error) {
	if appConfig.AppKey != "" {
		return
	}
	randomBytes := make([]byte, 32)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return
	}
	appConfig.AppKey = fmt.Sprintf("base64:%s", base64.StdEncoding.EncodeToString(randomBytes))
	return
}

func (appConfig *AppConfig) EnableSelectedDatabase() {
	var name string
	switch appConfig.Database {
	case "PostgreSQL":
		name = "pgsql"
	case "MariaDB":
		name = "mysql"
	}
	if m, ok := appConfig.PhpModules[name]; ok {
		m.Enabled = true
	}
}

func (appConfig *AppConfig) EnableSelectedAdminGen() {
	var name string
	switch appConfig.AdminGen {
	case "Nova":
		name = "laravel/nova"
	case "Backpack":
		name = "backpack/crud"
	}
	if m, ok := appConfig.ComposerDeps[name]; ok {
		m.Enabled = true
	}
}
