package modules

import (
	"testing"
)

func TestModuleFiles(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		fn   func() any
	}{
		{"php.yaml", func() any { return Php() }},
		{"composer.yaml", func() any { return Composer() }},
		{"npm.yaml", func() any { return Npm() }},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			defer func() {
				if r := recover(); r != nil {
					t.Error(r)
				}
			}()

			_ = test.fn()
		})
	}
}
