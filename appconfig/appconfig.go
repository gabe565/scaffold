package appconfig

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/clevyr/installer/phpmodules"
	"io/ioutil"
)

type AppConfig struct {
	AppName       string
	AppKey        string
	Database      string
	Modules       phpmodules.ModuleMap
	AdminGen      string
	MailDev       bool
	MaxUploadSize string
}

var Defaults = AppConfig{
	Database:      "PostgreSQL",
	Modules:       phpmodules.Defaults,
	AdminGen:      "None",
	MailDev:       true,
	MaxUploadSize: "64m",
}

const configFilePath = ".clevyr-installer-config"

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

func (appConfig AppConfig) ExportToFile() (err error) {
	var appConfigJson []byte
	appConfigJson, err = json.MarshalIndent(appConfig, "", "\t")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(configFilePath, appConfigJson, 0644)
	return
}

func (appConfig *AppConfig) ImportFromFile() error {
	appConfigJson, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(appConfigJson, &appConfig)
	if err != nil {
		return err
	}
	fmt.Printf("Loaded previous config from \"%s\"\n", configFilePath)
	return nil
}
