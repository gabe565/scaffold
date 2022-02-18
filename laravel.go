package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/clevyr/scaffold/appconfig"
	"github.com/clevyr/scaffold/iexec"
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

		flags := []string{"create-project", "laravel/laravel:^9.0", ".", "--no-install", "--no-plugins", "--no-scripts"}

		err = iexec.Command("composer", flags...)
		if err != nil {
			return
		}

		var composer map[string]interface{}
		composer, err = loadComposerJson()
		if err != nil {
			return
		}

		composer["name"] = fmt.Sprintf("clevyr/%s", appConfig.AppSlug)
		repositories := []map[string]string{}

		repositories = append(repositories, map[string]string{
			"type": "composer",
			"url":  "https://nova.laravel.com",
		})

		for _, module := range appConfig.ComposerDeps {
			if !module.Enabled {
				continue
			}

			if module.Name == "laravel/spark-paddle" || module.Name == "laravel/spark-stripe" {
				repositories = append(repositories, map[string]string{
					"type": "composer",
					"url":  "https://spark.laravel.com",
				})
				break
			}
		}

		composer["repositories"] = repositories

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
