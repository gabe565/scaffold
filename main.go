package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2/terminal"
)

type AppConfig struct {
	AppName string
	AppSlug string
	AppKey string
	Database string
	Modules map[string]bool
	AdminGen string
	MaxUploadSize string
}

func main() {
	appConfig, err := askQuestions()

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
