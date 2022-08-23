package laravel

import (
	"encoding/json"
	"fmt"
	"github.com/clevyr/scaffold/internal/appconfig"
	"github.com/clevyr/scaffold/internal/iexec"
	"log"
	"os"
)

func initLaravel(appConfig appconfig.AppConfig) error {
	if appConfig.InitLaravel {
		if err := os.MkdirAll(appConfig.AppSlug, os.ModePerm); err != nil {
			return err
		}
		if err := os.Chdir(appConfig.AppSlug); err != nil {
			return err
		}

		flags := []string{"create-project", "laravel/laravel:9.3.2", ".", "--no-install", "--no-plugins", "--no-scripts"}

		if err := iexec.NewBuilder("composer").Append(flags...).Run(); err != nil {
			return err
		}

		composer, err := loadComposerJson()
		if err != nil {
			return err
		}

		composer["name"] = fmt.Sprintf("clevyr/%s", appConfig.AppSlug)
		var repositories []map[string]string

		// Conditionally add Spark repository
		sparkModules := []string{"laravel/spark-paddle", "laravel/spark-stripe"}
		for _, moduleName := range sparkModules {
			if module, ok := appConfig.ComposerDeps.Map[moduleName]; ok && module.Enabled {
				log.Println("Add Spark repository")
				repositories = append(repositories, map[string]string{
					"type": "composer",
					"url":  "https://spark.laravel.com",
				})
				break
			}
		}

		// Conditionally add Nova repository
		if module, ok := appConfig.ComposerDeps.Map["laravel/nova"]; ok && module.Enabled {
			log.Println("Add Nova repository")
			repositories = append(repositories, map[string]string{
				"type": "composer",
				"url":  "https://nova.laravel.com",
			})
		}

		if len(repositories) > 0 {
			composer["repositories"] = repositories
		}

		if err = saveComposerJson(composer); err != nil {
			return err
		}
	}
	return nil
}

func loadComposerJson() (result map[string]any, err error) {
	f, err := os.Open("composer.json")
	if err != nil {
		return result, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	if err = json.NewDecoder(f).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func saveComposerJson(composer map[string]any) error {
	f, err := os.OpenFile("composer.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "    ")

	if err = encoder.Encode(composer); err != nil {
		return err
	}

	if err = f.Close(); err != nil {
		return err
	}

	return nil
}
