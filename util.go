package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const configFilePath = ".clevyr-installer-config"

func generateAppKey() (result string, err error) {
	randomBytes := make([]byte, 32)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return
	}
	result = fmt.Sprintf("base64:%s", base64.StdEncoding.EncodeToString(randomBytes))
	return
}

func writeAppConfig(appConfig AppConfig) (err error) {
	var appConfigJson []byte
	appConfigJson, err = json.MarshalIndent(appConfig, "", "\t")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(configFilePath, appConfigJson, 0644)
	return
}

func readAppConfig(appConfig *AppConfig) error {
	appConfigJson, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(appConfigJson, &appConfig)
	if err != nil {
		return err
	}
	fmt.Printf("Loaded previous config from \"%s\"\n", configFilePath)
	return nil
}
