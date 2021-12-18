package appconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (appConfig AppConfig) ExportToFile() (err error) {
	fmt.Printf("Persisting config to \"%s\"\n", configFilePath)
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
	err = json.Unmarshal(appConfigJson, appConfig)
	if err != nil {
		return err
	}
	fmt.Printf("Loaded previous config from \"%s\"\n", configFilePath)
	return nil
}
