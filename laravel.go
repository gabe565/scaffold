package main

import (
	"encoding/json"
	"fmt"
	"github.com/clevyr/scaffold/appconfig"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func initLaravel(appConfig appconfig.AppConfig) (err error) {
	if appConfig.InitLaravel {
		err = os.MkdirAll(appConfig.AppSlug, os.ModePerm)
		if err != nil {
			return
		}
		err = os.Chdir(appConfig.AppSlug)
		if err != nil {
			return
		}

		flags := []string{"create-project", "laravel/laravel", ".", "--no-install", "--no-plugins", "--no-scripts"}

		fmt.Printf("Running \"composer %s\"\n", strings.Join(flags, " "))

		cmd := exec.Command("composer", flags...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			return
		}

		var composer map[string]interface{}
		composer, err = loadComposerJson()
		if err != nil {
			return
		}

		composer["name"] = fmt.Sprintf("clevyr/%s", appConfig.AppSlug)

		if appConfig.AdminGen == "Nova" {
			composer["repositories"] = []map[string]string{
				map[string]string{
					"type": "composer",
					"url": "https://nova.laravel.com",
				},
			}
		}

		err = saveComposerJson(composer)
		if err != nil {
			return
		}
	}
	return
}

func loadComposerJson() (result map[string]interface{}, err error) {
	appConfigJson, err := ioutil.ReadFile("composer.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(appConfigJson, &result)
	if err != nil {
		return
	}
	return
}

func saveComposerJson(composer map[string]interface{}) (err error) {
	var outJson []byte
	outJson, err = json.MarshalIndent(composer, "", "    ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile("composer.json", outJson, 0644)
	return
}