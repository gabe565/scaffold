package appconfig

import (
	"encoding/base64"
	"strings"
	"testing"
)

func TestAppConfig_GenerateAppKey(t *testing.T) {
	t.Parallel()

	appConfig := Defaults
	err := appConfig.GenerateAppKey()
	if err != nil {
		t.Error(err)
	}
}

func TestAppConfig_GenerateAppKey_Prefix(t *testing.T) {
	t.Parallel()

	appConfig := Defaults
	err := appConfig.GenerateAppKey()
	if err != nil {
		t.Error(err)
	}

	if !strings.HasPrefix(appConfig.AppKey, AppKeyPrefix) {
		t.Errorf(`invalid app key. expected prefix "%s", got "%s"`, AppKeyPrefix, appConfig.AppKey)
	}
}

func TestAppConfig_GenerateAppKey_ValidBase64(t *testing.T) {
	t.Parallel()

	appConfig := Defaults
	err := appConfig.GenerateAppKey()
	if err != nil {
		t.Error(err)
	}

	raw := strings.TrimPrefix(appConfig.AppKey, AppKeyPrefix)
	decoded, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		t.Error(err)
	}
	if string(decoded) == "" {
		t.Errorf(`missing app key`)
	}
}

func TestAppConfig_EnableSelectedDatabase(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  string
		module string
	}{
		{"PostgreSQL", "pgsql"},
		{"MariaDB", "mysql"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			appConfig := Defaults
			appConfig.Database = test.input
			appConfig.EnableSelectedDatabase()
			enabled := appConfig.PhpModules[test.module].Enabled
			if !enabled {
				t.Errorf("PHP module not enabled. Expected true, got %t", enabled)
			}
		})
	}
}

func TestAppConfig_EnableSelectedAdminGen(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  string
		module string
	}{
		{"Nova", "laravel/nova"},
		{"Backpack", "backpack/crud"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			appConfig := Defaults
			appConfig.AdminGen = test.input
			appConfig.EnableSelectedAdminGen()
			enabled := appConfig.ComposerDeps[test.module].Enabled
			if !enabled {
				t.Errorf("Composer module not enabled. Expected true, got %t", enabled)
			}
		})
	}
}
