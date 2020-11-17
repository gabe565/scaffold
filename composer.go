package main

import (
	"fmt"
	"github.com/clevyr/scaffold/appconfig"
	"github.com/clevyr/scaffold/iexec"
	"os/exec"
)

func composerRequire(appConfig appconfig.AppConfig) (err error) {
	var param, devParam []string

	for name, module := range appConfig.ComposerDeps {
		if module.Enabled && !composerInstalled(name) {
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
		err = iexec.Command("composer", append([]string{"require", "--dev"}, devParam...)...)
	}

	if len(param) > 0 {
		err = iexec.Command("composer", append([]string{"require"}, param...)...)
	}

	return
}

func composerInstalled(name string) bool {
	cmd := exec.Command("composer", "show", name)
	err := cmd.Run()
	return err == nil
}
