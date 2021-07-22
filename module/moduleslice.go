package module

import (
	"sort"

	"github.com/AlecAivazis/survey/v2/core"
)

type ModuleSlice []*Module

func (modules ModuleSlice) WriteAnswer(name string, value interface{}) error {
	for _, module := range modules {
		module.Enabled = false
	}

	options := value.([]core.OptionAnswer)
	for _, option := range options {
		for _, module := range modules {
			if module.Name == option.Value {
				module.Enabled = true
				break
			}
		}
	}

	return nil
}

func (modules ModuleSlice) ToOptionsSlice() []string {
	result := make([]string, 0, len(modules))

	for _, module := range modules {
		if !module.Hidden {
			result = append(result, module.Name)
		}
	}

	sort.Strings(result)
	return result
}

func (modules ModuleSlice) ToDefaultSlice() []string {
	result := make([]string, 0, len(modules))

	for _, module := range modules {
		if module.Enabled {
			result = append(result, module.Name)
		}
	}
	return result
}
