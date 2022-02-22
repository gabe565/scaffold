package main

import (
	"fmt"
	"github.com/clevyr/scaffold/appconfig"
	"github.com/clevyr/scaffold/iexec"
	"os/exec"
	"sort"
)

func composerRequire(appConfig appconfig.AppConfig) (err error) {
	var param, devParam []string

	slice := appConfig.ComposerDeps.Slice()
	sort.Sort(&slice)

	for _, module := range slice {
		if module.Enabled && !composerInstalled(module.Name) {
			var appParam string

			if module.Version == "" {
				appParam = module.Name
			} else {
				appParam = fmt.Sprintf("%s:%s", module.Name, module.Version)
			}

			if module.Dev {
				devParam = append(devParam, appParam)
			} else {
				param = append(param, appParam)
			}
		}
	}

	if len(param) > 0 {
		err = iexec.Command("composer", append([]string{"require"}, param...)...)
	}

	if len(devParam) > 0 {
		err = iexec.Command("composer", append([]string{"require", "--dev"}, devParam...)...)
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

func composerInstalled(name string) bool {
	cmd := exec.Command("composer", "show", name)
	err := cmd.Run()
	return err == nil
}
