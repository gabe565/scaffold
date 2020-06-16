//go:generate pkger
package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/clevyr/installer/phpmodules"
)

type AppConfig struct {
	AppName       string
	AppSlug       string
	AppKey        string
	Database      string
	Modules       phpmodules.ModuleMap
	AdminGen      string
	MaxUploadSize string
}

func main() {
	appConfig := AppConfig{}
	err := askQuestions(&appConfig)

	if err == terminal.InterruptErr {
		fmt.Println("Interrupted")
		return
	} else if err != nil {
		panic(err)
	}

	appConfig.AppKey = generateAppKey()

	if err := generateTemplate(appConfig); err != nil {
		panic(err)
	}
}
