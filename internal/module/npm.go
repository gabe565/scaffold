package module

import (
	"fmt"
	"sort"

	"github.com/clevyr/scaffold/internal/iexec"
)

type NpmMap struct {
	Map
}

func (NpmMap) Install() error {
	return iexec.NewBuilder("npm", "install").Run()
}

func (m NpmMap) InstallDeps() (err error) {
	var param, devParam []string

	slice := m.Slice()
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
		err = iexec.NewBuilder("npm", "install", "--save").Append(param...).Run()
		if err != nil {
			return err
		}
	}

	if len(devParam) > 0 {
		err = iexec.NewBuilder("npm", "install", "--save-dev").Append(devParam...).Run()
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
