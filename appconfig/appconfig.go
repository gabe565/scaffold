package appconfig

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/clevyr/scaffold/appconfig/defaults"
	"github.com/clevyr/scaffold/module"
)

type AppConfig struct {
	InitLaravel   bool `json:"-"`
	AppName       string
	AppSlug       string `json:"-"`
	AppKey        string
	Database      string
	PhpModules    module.ModuleMap
	JetstreamGen  string
	ComposerDeps  module.ModuleSlice
	NpmDeps       module.ModuleSlice
	MaxUploadSize string
}

var Defaults = AppConfig{
	Database:      "PostgreSQL",
	PhpModules:    defaults.PhpModules,
	JetstreamGen:  "No Teams",
	ComposerDeps:  defaults.ComposerDeps,
	NpmDeps:       defaults.NpmDeps,
	MaxUploadSize: "64m",
}

func (ac AppConfig) HasComposerDep(name string) bool {
	for _, module := range ac.ComposerDeps {
		if module.Name == name && module.Enabled == true {
			return true
		}
	}

	return false
}

const configFilePath = ".clevyr-scaffold-config"

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
	switch appConfig.Database {
	case "PostgreSQL":
		if pgsqlModule, ok := (*appConfig).PhpModules["pgsql"]; ok {
			pgsqlModule.Enabled = true
		}
	case "MariaDB":
		if mysqlModule, ok := (*appConfig).PhpModules["mysql"]; ok {
			mysqlModule.Enabled = true
		}
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
