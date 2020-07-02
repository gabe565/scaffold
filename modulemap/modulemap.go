package modulemap

import (
	"encoding/json"
	"github.com/AlecAivazis/survey/v2/core"
	"sort"
)

type Module struct {
	Enabled bool   `json:",omitempty"`
	Hidden  bool   `json:"-"`
	Version string `json:",omitempty"`
}

func (module *Module) WriteAnswer(name string, value interface{}) error {
	module.Enabled = value.(bool)
	return nil
}

type ModuleMap map[string]*Module

func (modules ModuleMap) ToOptionsSlice() []string {
	result := make([]string, 0, len(modules))
	for i, module := range modules {
		if !module.Hidden {
			result = append(result, i)
		}
	}
	sort.Strings(result)
	return result
}

func (modules ModuleMap) ToDefaultSlice() []string {
	result := make([]string, 0, len(modules))
	for i, module := range modules {
		if module.Enabled {
			result = append(result, i)
		}
	}
	return result
}

func (modules ModuleMap) WriteAnswer(name string, value interface{}) error {
	options := value.([]core.OptionAnswer)
	for _, option := range options {
		modules[option.Value].Enabled = true
	}
	return nil
}

func (modules ModuleMap) UnmarshalJSON(data []byte) (err error) {
	tmp := make(map[string]*Module)
	err = json.Unmarshal(data, &tmp)
	for key, loadedModule := range tmp {
		if defaultModule, ok := modules[key]; ok {
			defaultModule.Enabled = loadedModule.Enabled
			defaultModule.Version = loadedModule.Version
		}
	}
	return
}