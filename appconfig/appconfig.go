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
	JetstreamGen  string
	ComposerDeps  module.Map
	NpmDeps       module.Map
	MaxUploadSize string
}

var Defaults = AppConfig{
	Database:      "PostgreSQL",
	PhpModules:    modules.Php(),
	JetstreamGen:  "No Teams",
	ComposerDeps:  modules.Composer(),
	NpmDeps:       modules.Npm(),
	MaxUploadSize: "64m",
}

const AppKeyPrefix = "base64:"
const AppKeyBytes = 32

func (appConfig *AppConfig) GenerateAppKey() (err error) {
	if appConfig.AppKey != "" {
		return
	}
	randomBytes := make([]byte, AppKeyBytes)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return
	}
	appConfig.AppKey = fmt.Sprintf("%s%s", AppKeyPrefix, base64.StdEncoding.EncodeToString(randomBytes))
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

func (appConfig *AppConfig) EnableSelectedJetstreamGen() {
    for _, module := range (*appConfig).ComposerDeps {
        if module.Name == "laravel/jetstream" {
            switch appConfig.JetstreamGen {
            case "No Teams":
                // Do nothing
            case "With Teams":
                // Append Jetstream's post install command with the '--teams' modifier
                module.PostInstallCmds[0] = append(module.PostInstallCmds[0], "--teams")
            }
        }
    }
}
