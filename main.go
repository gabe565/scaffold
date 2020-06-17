//go:generate pkger
package main

import (
	"flag"
	"fmt"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/clevyr/installer/phpmodules"
	"os"
	"path"
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

	var context string
	flag.StringVar(&context, "C", ".", "Run as if the application was started in the given path.")
	flag.Parse()

	context = path.Clean(context)
	if context != "." {
		err = os.MkdirAll(context, os.ModePerm)
		if err != nil {
			panic(err)
		}
		err = os.Chdir(context)
		if err != nil {
			panic(err)
		}
	}

	appConfig := AppConfig{}
	_ = readAppConfig(&appConfig)
	err = askQuestions(&appConfig)
	if err == terminal.InterruptErr {
		fmt.Println("Interrupted")
		return
	} else if err != nil {
		panic(err)
	}

	appConfig.AppKey, err = generateAppKey()
	if err != nil {
		panic(err)
	}

	err = generateTemplate(appConfig)
	if err != nil {
		panic(err)
	}

	err = writeAppConfig(appConfig)
	if err != nil {
		panic(err)
	}
}
