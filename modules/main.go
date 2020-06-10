package modules

import "sort"

type PHPModules struct {
	Modules map[string]*PHPModule
}

func (modules PHPModules) ToOptionsSlice() []string {
	result := make([]string, 0, len(modules.Modules))
	for i, module := range modules.Modules {
		if !module.Hidden {
			result = append(result, i)
		}
	}
	sort.Strings(result)
	return result
}

func (modules PHPModules) ToDefaultSlice() []string {
	result := make([]string, 0, len(modules.Modules))
	for i, module := range modules.Modules {
		if module.Enabled {
			result = append(result, i)
		}
	}
	return result
}

type PHPModule struct {
	Enabled bool
	Hidden bool
}

func (module *PHPModule) WriteAnswer(name string, value interface{}) error {
	module.Enabled = value.(bool)
	return nil
}

var (
	All = PHPModules{
		map[string]*PHPModule{
			"bcmath": {},
			"calendar": {},
			"exif": {},
			"gd": {},
			"imagick": {},
			"mosquitto": {},
			"mysql": {},
			"opcache": {Enabled: true},
			"pgsql": {},
			"redis": {Enabled: true},
			"sqlsrv": {},
			"xdebug": {Hidden: true},
			"zip": {},
		},
	}
)