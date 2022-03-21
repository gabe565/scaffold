package module

import (
	"fmt"
	"github.com/clevyr/scaffold/internal/iexec"
	"os/exec"
	"sort"
)

type ComposerMap struct {
	Map
}

func (ComposerMap) Install() error {
	return iexec.Command("composer", "install")
}

func (m ComposerMap) InstallDeps() (err error) {
	var param, devParam []string

	slice := m.Slice()
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
		if module.Enabled {
			for _, then := range module.Then {
				err = then.Activate()
				if err != nil {
					return err
				}
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
