package module

import (
	"github.com/AlecAivazis/survey/v2/core"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMap_ToOptionsSlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		mmap        Map
		expectedLen int
		expectedStr string
	}{
		{"shown module", Map{"shown": &Module{}}, 1, "shown"},
		{"hidden module", Map{"hidden": &Module{Hidden: true}}, 0, ""},
	}

	for _, test := range tests {
		test := test
		t.Run("hidden module", func(t *testing.T) {
			t.Parallel()

			slice := test.mmap.ToOptionsSlice()
			if len(slice) != test.expectedLen {
				t.Errorf("options slice has invalid length. got %d, expected %d", len(slice), test.expectedLen)
			}
			if test.expectedStr != "" && slice[0] != test.expectedStr {
				t.Errorf(`first value of options slice is invalid. got "%s", expected "%s"`, slice[0], test.expectedStr)
			}
		})
	}
}

func TestMap_ToDefaultSlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		mmap        Map
		expectedLen int
		expectedStr string
	}{
		{"shown module", Map{"enabled": &Module{Enabled: true}}, 1, "enabled"},
		{"hidden module", Map{"disabled": &Module{}}, 0, ""},
	}

	for _, test := range tests {
		test := test
		t.Run("hidden module", func(t *testing.T) {
			t.Parallel()

			slice := test.mmap.ToDefaultSlice()
			if len(slice) != test.expectedLen {
				t.Errorf("defaults slice has invalid length. got %d, expected %d", len(slice), test.expectedLen)
			}
			if test.expectedStr != "" && slice[0] != test.expectedStr {
				t.Errorf(`first value of defaults slice is invalid. got "%s", expected "%s"`, slice[0], test.expectedStr)
			}
		})
	}
}

func TestMap_WriteAnswer(t *testing.T) {
	t.Parallel()

	mmap := Map{
		"first":  &Module{},
		"second": &Module{Enabled: true},
	}
	err := mmap.WriteAnswer("", []core.OptionAnswer{{Value: "first", Index: 0}})
	if err != nil {
		t.Error(err)
	}

	if !mmap["first"].Enabled {
		t.Errorf("writing a value did not enable the module. got %t, expected true", mmap["first"].Enabled)
	}
	if mmap["second"].Enabled {
		t.Errorf("writing a value did not enable other modules. got %t, expected false", mmap["second"].Enabled)
	}
}

func TestMap_UnmarshalYAML(t *testing.T) {
	t.Parallel()

	input := `test: {}`
	var mmap Map
	err := yaml.Unmarshal([]byte(input), &mmap)
	if err != nil {
		t.Error(err)
	}
	testModule, ok := mmap["test"]
	if !ok {
		t.Error("module not correctly loaded from yaml")
		return
	}
	if testModule.Name != "test" {
		t.Errorf("module name not set from yaml key")
	}
}

func TestMap_Slice(t *testing.T) {
	t.Parallel()

	mmap := Map{
		"first": &Module{},
	}
	slice := mmap.Slice()
	if len(slice) != 1 {
		t.Errorf("invalid slice length returned. got %d, expected %d", len(slice), 1)
	}
}
