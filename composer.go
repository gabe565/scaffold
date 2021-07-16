package main

import (
	"fmt"
	"os/exec"

	"github.com/clevyr/scaffold/appconfig"
	"github.com/clevyr/scaffold/iexec"
)

func composerRequire(appConfig appconfig.AppConfig) (err error) {
	var param, devParam []string
	var postInstallCmds [][]string

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

			if len(module.PostInstallCmds) > 0 {
				postInstallCmds = append(postInstallCmds, module.PostInstallCmds...)
			}
		}
	}

	if len(param) > 0 {
		err = iexec.Command("composer", append([]string{"require"}, param...)...)
	}

	if len(devParam) > 0 {
		err = iexec.Command("composer", append([]string{"require", "--dev"}, devParam...)...)
	}

	if len(postInstallCmds) > 0 {
		for _, cmd := range postInstallCmds {
			err = iexec.Command(cmd[0], cmd[1:]...)
		}
	}

	return
}

func composerInstalled(name string) bool {
	cmd := exec.Command("composer", "show", name)
	err := cmd.Run()
	return err == nil
}
