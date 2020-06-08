package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2/terminal"
	"os"
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
	appConfig, err := Ask()
	if err == terminal.InterruptErr {
		fmt.Println("Interrupted")
		os.Exit(0)
	} else if err != nil {
		panic(err)
	}

	appConfig.AppKey = generateAppKey()

	generateTemplate(appConfig)
}
