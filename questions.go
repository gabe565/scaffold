package main

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/clevyr/installer/modules"
	"github.com/iancoleman/strcase"
	"regexp"
)

var validationRegex, _ = regexp.Compile("^[0-9]*[kmg]$")

func askQuestions() (appConfig AppConfig, err error) {
	// App Name
	err = survey.AskOne(&survey.Input{Message: "What is the application name?"}, &appConfig.AppName, survey.WithValidator(survey.Required))
	if err != nil {
		return
	}
	appConfig.AppSlug = strcase.ToKebab(appConfig.AppName)

	// Database
	err = survey.AskOne(&survey.Select{
		Message: "Choose which main database server to configure:",
		Options: []string{"PostgreSQL", "MariaDB"},
		Default: "PostgreSQL",
	}, &appConfig.Database, survey.WithValidator(survey.Required))
	if err != nil {
		return
	}

	// Enabled Modules
	shownModules := modules.All
	switch appConfig.Database {
	case "PostgreSQL":
		shownModules.Modules["pgsql"].Enabled = true
		break
	case "MariaDB":
		shownModules.Modules["mysql"].Enabled = true
		break
	}
	var enabledModules []string
	err = survey.AskOne(&survey.MultiSelect{
		Message: "Choose which PHP modules to enable:",
		Options: shownModules.ToOptionsSlice(),
		Default: shownModules.ToDefaultSlice(),
	}, &enabledModules)
	if err != nil {
		return
	}

	for _, module := range enabledModules {
		shownModules.Modules[module].Enabled = true
	}
	appConfig.Modules = shownModules

	// Admin Gen
	err = survey.AskOne(&survey.Select{
		Message: "Choose which admin generator to include:",
		Options: []string{"None", "Nova", "Backpack"},
		Default: "None",
	}, &appConfig.AdminGen)
	if err != nil {
		return
	}

	// Max Upload Size
	err = survey.AskOne(
		&survey.Input{
			Message: "What is the maximum upload size that should be allowed?",
			Default: "64m",
			Help: "Configures the maximum allowed upload size. " +
				"Supports the suffixes \"k\" (kilobytes), \"m\" (megabytes) and \"g\" (gigabytes).",
		},
		&appConfig.MaxUploadSize,
		survey.WithValidator(func(val interface{}) error {
			if str, ok := val.(string); !ok || !validationRegex.MatchString(str) {
				return errors.New("Make sure to enter a size followed by \"k\" (kilobytes), \"m\" (megabytes) or \"g\" (gigabytes).")
			}
			return nil
		}),
	)
	if err != nil {
		return
	}

	return
}