package modules

import (
	_ "embed"
	"github.com/clevyr/scaffold/module"
	"gopkg.in/yaml.v3"
)

//go:embed composer.yaml
var composer []byte

//go:embed npm.yaml
var npm []byte

//go:embed php.yaml
var php []byte

func unmarshalConfig(config []byte) (module.ModuleMap, error) {
	var modules module.ModuleMap
	err := yaml.Unmarshal(config, &modules)
	return modules, err
}

func Composer() module.ModuleMap {
	modules, err := unmarshalConfig(composer)
	if err != nil {
		panic(err)
	}
	return modules
}

func Npm() module.ModuleMap {
	modules, err := unmarshalConfig(npm)
	if err != nil {
		panic(err)
	}
	return modules
}

func Php() module.ModuleMap {
	modules, err := unmarshalConfig(php)
	if err != nil {
		panic(err)
	}
	return modules
}
