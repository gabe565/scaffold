package appconfig

import (
	"encoding/base64"
	"fmt"
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

func TestAppConfig_EnableJetstreamTeams(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input           bool
		expectedLenDiff int
	}{
		{false, 0},
		{true, 1},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("teams_%t", test.input), func(t *testing.T) {
			t.Parallel()

			appConfig := Defaults
			appConfig.JetstreamTeams = test.input
			beforeLen := len(*appConfig.ComposerDeps.Map["laravel/jetstream"].Then[0].Run)
			appConfig.EnableJetstreamTeams()
			diff := len(*appConfig.ComposerDeps.Map["laravel/jetstream"].Then[0].Run) - beforeLen
			if diff != test.expectedLenDiff {
				t.Errorf(`laravel/jetstream "--teams" flag not added. Expected len change of %d, got %d`, test.expectedLenDiff, diff)
			}
		})
	}
}
