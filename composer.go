package main

import (
	"fmt"
	"github.com/clevyr/scaffold/appconfig"
	"github.com/clevyr/scaffold/iexec"
	"strings"
)

func composerRequire(appConfig appconfig.AppConfig) (err error) {
	dependencies := make([]string, 0, len(appConfig.ComposerDeps))
	devDependencies := make([]string, 0, len(appConfig.ComposerDeps))

	for name, module := range appConfig.ComposerDeps {
		if module.Enabled {
			var appParam string

			if module.Version == "" {
				appParam = name
			} else {
				appParam = fmt.Sprintf("%s:\"%s\"", name, module.Version)
			}

			if module.Dev {
				devDependencies = append(devDependencies, appParam)
			} else {
				dependencies = append(dependencies, appParam)
			}
		}
	}

	if len(devDependencies) > 0 {
		fmt.Printf("Running \"composer require --dev %s\"\n", strings.Join(devDependencies, " "))
		err = iexec.Command("composer", append([]string{"require", "--ignore-platform-reqs", "--dev"}, devDependencies...)...)
	}

	if len(dependencies) > 0 {
		fmt.Printf("Running \"composer require %s\"\n", strings.Join(dependencies, " "))
		err = iexec.Command("composer", append([]string{"require", "--ignore-platform-reqs"}, dependencies...)...)
	}

	return
}