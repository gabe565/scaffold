package main

import (
	"fmt"
	"github.com/clevyr/scaffold/appconfig"
	"strings"
)

func composerRequire(appConfig appconfig.AppConfig) (err error) {
	dependencies := make([]string, 0, len(appConfig.ComposerDeps))
	for name, module := range appConfig.ComposerDeps {
		if module.Enabled {
			if module.Version != "" {
				dependencies = append(dependencies, fmt.Sprintf("%s:\"%s\"", name, module.Version))
			} else {
				dependencies = append(dependencies, name)
			}
		}
	}

	if len(dependencies) > 0 {
		fmt.Printf("Running \"composer require %s\"\n", strings.Join(dependencies, " "))
		err = interactiveCommand("composer", append([]string{"require", "--ignore-platform-reqs"}, dependencies...)...)
	}

	return
}