package main

import (
	"fmt"

	"github.com/clevyr/scaffold/appconfig"
	"github.com/clevyr/scaffold/iexec"
)

func npmInstall() (err error) {
	err = iexec.Command("npm", "install")
	return
}

func npmInstallDeps(appConfig appconfig.AppConfig) (err error) {
	var param, devParam []string

	for name, module := range appConfig.NpmDeps {
		if module.Enabled {
			var appParam string

			if module.Version == "" {
				appParam = name
			} else {
				appParam = fmt.Sprintf("%s@%s", name, module.Version)
			}

			if module.Dev {
				devParam = append(devParam, appParam)
			} else {
				param = append(param, appParam)
			}
		}
	}

	if len(param) > 0 {
		err = iexec.Command("npm", append([]string{"install", "--save"}, param...)...)
	}

	if len(devParam) > 0 {
		err = iexec.Command("npm", append([]string{"install", "--save-dev"}, devParam...)...)
	}

	for _, module := range appConfig.NpmDeps {
		for _, then := range module.Then {
			err = then.Activate()
			if err != nil {
				return err
			}
		}
	}

	return
}
