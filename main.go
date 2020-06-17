//go:generate pkger
package main

import (
	"flag"
	"fmt"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/clevyr/installer/phpmodules"
)

type AppConfig struct {
	AppName       string
	AppKey        string
	Database      string
	Modules       phpmodules.ModuleMap
	AdminGen      string
	MaxUploadSize string
}

func main() {
	var err error

	context := flag.String("C", ".", "Run as if the application was started in the given path.")
	flag.Parse()

	appConfig := AppConfig{}
	err = askQuestions(&appConfig)

	if err == terminal.InterruptErr {
		fmt.Println("Interrupted")
		return
	} else if err != nil {
		panic(err)
	}

	appConfig.AppKey = generateAppKey()

	err = generateTemplate(appConfig, *context)
	if err != nil {
		panic(err)
	}
}
