package main

import (
	"encoding/json"
	"fmt"
	"github.com/clevyr/scaffold/appconfig"
	"github.com/clevyr/scaffold/iexec"
	"io/ioutil"
	"log"
	"os"
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

		flags := []string{"create-project", "laravel/laravel:^8.0", ".", "--no-install", "--no-plugins", "--no-scripts"}

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
		var repositories []map[string]string

		// Conditionally add Spark repository
		sparkModules := []string{"laravel/spark-paddle", "laravel/spark-stripe"}
		for _, moduleName := range sparkModules {
			if module, ok := appConfig.ComposerDeps[moduleName]; ok && module.Enabled {
				log.Println("Add Spark repository")
				repositories = append(repositories, map[string]string{
					"type": "composer",
					"url":  "https://spark.laravel.com",
				})
				break
			}
		}

		// Conditionally add Nova repository
		if module, ok := appConfig.ComposerDeps["laravel/nova"]; ok && module.Enabled {
			log.Println("Add Nova repository")
			repositories = append(repositories, map[string]string{
				"type": "composer",
				"url":  "https://nova.laravel.com",
			})
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
