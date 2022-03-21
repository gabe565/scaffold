package module

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"sort"

	"github.com/AlecAivazis/survey/v2/core"
)

type Map map[string]*Module

func (modules Map) ToOptionsSlice() []string {
	result := make([]string, 0, len(modules))
	for i, module := range modules {
		if !module.Hidden {
			result = append(result, i)
		}
	}
	sort.Strings(result)
	return result
}

func (modules Map) ToDefaultSlice() []string {
	result := make([]string, 0, len(modules))
	for i, module := range modules {
		if module.Enabled {
			result = append(result, i)
		}
	}
	return result
}

func (modules Map) WriteAnswer(name string, value any) error {
	// Set all to false to prevent defaults from staying enabled
	for _, module := range modules {
		module.Enabled = false
	}

	options := value.([]core.OptionAnswer)
	for _, option := range options {
		modules[option.Value].Enabled = true
	}
	return nil
}

func (modules *Map) UnmarshalYAML(value *yaml.Node) error {
	// Create raw type to decode data
	type raw Map
	err := value.Decode((*raw)(modules))
	if err != nil {
		return err
	}

	// Set names
	for key, module := range *modules {
		if module == nil {
			module = &Module{}
			(*modules)[key] = module
		}
		module.Name = key
	}
	return nil
}

func (modules Map) Slice() Slice {
	result := make(Slice, 0, len(modules))
	for _, module := range modules {
		result = append(result, module)
	}
	return result
}

var ErrInvalidModule = errors.New("unknown module")

func (modules Map) ModuleEnabled(name string) bool {
	module, ok := modules[name]
	if !ok {
		panic(fmt.Errorf("%v: %s", ErrInvalidModule, name))
	}
	return module.Enabled
}
