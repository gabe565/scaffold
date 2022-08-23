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
	return iexec.NewBuilder("composer", "install").Run()
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
		err := iexec.NewBuilder("composer", "require").Append(param...).Run()
		if err != nil {
			return err
		}
	}

	if len(devParam) > 0 {
		err = iexec.NewBuilder("composer", "require", "--dev").Append(devParam...).Run()
		if err != nil {
			return err
		}
	}

	for _, module := range slice {
		if module.Enabled {
			for _, then := range module.Then {
				if err = then.Activate(); err != nil {
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
