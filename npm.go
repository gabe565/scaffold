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
	var postInstallCmds [][]string

	for _, module := range appConfig.NpmDeps {
		if module.Enabled {
			var appParam string

			if module.Version == "" {
				appParam = module.Name
			} else {
				appParam = fmt.Sprintf("%s@%s", module.Name, module.Version)
			}

			if module.Dev {
				devParam = append(devParam, appParam)
			} else {
				param = append(param, appParam)
			}

			if len(module.PostInstallCmds) > 0 {
				postInstallCmds = append(postInstallCmds, module.PostInstallCmds...)
			}
		}
	}

	if len(param) > 0 {
		err = iexec.Command("npm", append([]string{"install", "--save"}, param...)...)
	}

	if len(devParam) > 0 {
		err = iexec.Command("npm", append([]string{"install", "--save-dev"}, devParam...)...)
	}

	if len(postInstallCmds) > 0 {
		for _, cmd := range postInstallCmds {
			err = iexec.Command(cmd[0], cmd[1:]...)
		}
	}

	return
}
