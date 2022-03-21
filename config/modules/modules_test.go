package modules

import (
	"github.com/clevyr/scaffold/internal/module"
	"testing"
)

func TestModuleFiles(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		fn   func() module.Map
	}{
		{"php.yaml", Php},
		{"composer.yaml", Composer},
		{"npm.yaml", Npm},
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
