package main

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/clevyr/installer/appconfig"
	"regexp"
)

var validationRegex, _ = regexp.Compile("^[0-9]*[kmg]$")

func askQuestions(appConfig *appconfig.AppConfig) (err error) {
	// App Name
	err = survey.AskOne(&survey.Input{
		Message: "What is the application name?",
		Default: appConfig.AppName,
	}, &appConfig.AppName, survey.WithValidator(survey.Required))
	if err != nil {
		return
	}

	// Database
	err = survey.AskOne(&survey.Select{
		Message: "Choose which main database server to configure:",
		Options: []string{"PostgreSQL", "MariaDB"},
		Default: appConfig.Database,
	}, &appConfig.Database, survey.WithValidator(survey.Required))
	if err != nil {
		return
	}

	// Enabled Modules
	appConfig.Modules.EnableSelectedDatabase(appConfig.Database)
	err = survey.AskOne(&survey.MultiSelect{
		Message: "Choose which PHP phpmodules to enable:",
		Options: appConfig.Modules.ToOptionsSlice(),
		Default: appConfig.Modules.ToDefaultSlice(),
	}, &appConfig.Modules)
	if err != nil {
		return
	}

	// Admin Gen
	err = survey.AskOne(&survey.Select{
		Message: "Choose which admin generator to include:",
		Options: []string{"None", "Nova", "Backpack"},
		Default: appConfig.AdminGen,
	}, &appConfig.AdminGen)
	if err != nil {
		return
	}

	// Max Upload Size
	err = survey.AskOne(
		&survey.Input{
			Message: "What is the maximum upload size that should be allowed?",
			Default: appConfig.MaxUploadSize,
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