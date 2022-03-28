package module

import (
	"fmt"
	"github.com/clevyr/scaffold/internal/iexec"
	"sort"
)

type NpmMap struct {
	Map
}

func (NpmMap) Install() error {
	return iexec.Command("npm", "install")
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
		err = iexec.Command("npm", append([]string{"install", "--save"}, param...)...)
		if err != nil {
			return err
		}
	}

	if len(devParam) > 0 {
		err = iexec.Command("npm", append([]string{"install", "--save-dev"}, devParam...)...)
		if err != nil {
			return err
		}
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
