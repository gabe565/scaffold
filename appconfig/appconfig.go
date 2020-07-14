package appconfig

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/clevyr/scaffold/appconfig/defaults"
	"github.com/clevyr/scaffold/modulemap"
)

type AppConfig struct {
	InitLaravel   bool `json:"-"`
	AppName       string
	AppSlug       string `json:"-"`
	AppKey        string
	Database      string
	PhpModules    modulemap.ModuleMap
	AdminGen      string
	ComposerDeps  modulemap.ModuleMap
	MailDev       bool
	MaxUploadSize string
}

var Defaults = AppConfig{
	Database:      "PostgreSQL",
	PhpModules:    defaults.PhpModules,
	AdminGen:      "None",
	ComposerDeps:  defaults.ComposerDeps,
	MailDev:       true,
	MaxUploadSize: "64m",
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
		break
	case "MariaDB":
		if mysqlModule, ok := (*appConfig).PhpModules["mysql"]; ok {
			mysqlModule.Enabled = true
		}
		break
	}
}

func (appConfig *AppConfig) EnableSelectedAdminGen() {
	switch appConfig.AdminGen {
	case "Nova":
		if module, ok := (*appConfig).ComposerDeps["laravel/nova"]; ok {
			module.Enabled = true
		}
		break
	case "Backpack":
		if module, ok := (*appConfig).ComposerDeps["backpack/crud"]; ok {
			module.Enabled = true
		}
		break
	}
}
