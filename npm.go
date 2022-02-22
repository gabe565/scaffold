package main

import (
	"fmt"
	"github.com/clevyr/scaffold/appconfig"
	"github.com/clevyr/scaffold/iexec"
	"sort"
)

func npmInstall() (err error) {
	err = iexec.Command("npm", "install")
	return
}

func npmInstallDeps(appConfig appconfig.AppConfig) (err error) {
	var param, devParam []string

	slice := appConfig.NpmDeps.Slice()
	sort.Sort(&slice)

	for _, module := range slice {
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
		}
	}

	if len(param) > 0 {
		err = iexec.Command("npm", append([]string{"install", "--save"}, param...)...)
	}

	if len(devParam) > 0 {
		err = iexec.Command("npm", append([]string{"install", "--save-dev"}, devParam...)...)
	}

	for _, module := range slice {
		for _, then := range module.Then {
			err = then.Activate()
			if err != nil {
				return err
			}
		}
	}

	return
}
