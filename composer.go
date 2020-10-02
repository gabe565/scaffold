package main

import (
	"fmt"
	"github.com/clevyr/scaffold/appconfig"
	"github.com/clevyr/scaffold/iexec"
)

func composerRequire(appConfig appconfig.AppConfig) (err error) {
	param := []string{"require", "--ignore-platform-reqs"}
	devParam := []string{"require", "--ignore-platform-reqs", "--dev"}

	for name, module := range appConfig.ComposerDeps {
		if module.Enabled {
			var appParam string

			if module.Version == "" {
				appParam = name
			} else {
				appParam = fmt.Sprintf("%s:%s", name, module.Version)
			}

			if module.Dev {
				devParam = append(devParam, appParam)
			} else {
				param = append(param, appParam)
			}
		}
	}

	if len(devParam) > 0 {
		err = iexec.Command("composer", devParam...)
	}

	if len(param) > 0 {
		err = iexec.Command("composer", param...)
	}

	return
}