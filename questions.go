package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/AlecAivazis/survey/v2"
	"github.com/clevyr/scaffold/appconfig"
	"github.com/huandu/xstrings"
)

var validationRegex, _ = regexp.Compile("^[0-9]*[kmg]$")

func askQuestions(appConfig *appconfig.AppConfig) (err error) {
	// App Name
	err = survey.AskOne(&survey.Input{
		Message: "Enter the application name that should be shown to the user.",
		Default: appConfig.AppName,
		Help:    "Enter the application's name space-separated and capitalized. A slug will be generated to match.",
	}, &appConfig.AppName, survey.WithValidator(survey.Required))
	if err != nil {
		return
	}
	appConfig.AppSlug = xstrings.ToKebabCase(appConfig.AppName)

	// Laravel Install
	_, err = os.Stat("composer.json")
	if os.IsNotExist(err) {
		wd, _ := os.Getwd()

		err = survey.AskOne(&survey.Confirm{
			Message: fmt.Sprintf("composer.json not found in the current directory. Initialize a Laravel app in \"%s/%s\"?", wd, appConfig.AppSlug),
			Default: true,
		}, &appConfig.InitLaravel)
		if err != nil {
			return
		}
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

	// Enabled PhpModules
	appConfig.EnableSelectedDatabase()
	err = survey.AskOne(&survey.MultiSelect{
		Message: "Choose which PHP modules to enable:",
		Options: appConfig.PhpModules.ToOptionsSlice(),
		Default: appConfig.PhpModules.ToDefaultSlice(),
	}, &appConfig.PhpModules)
	if err != nil {
		return
	}

	// Jetstream Teams
	err = survey.AskOne(&survey.Confirm{
		Message: "Do you want to use Jetstream with teams?",
		Help: "'Teams' are Jetstream's built-in single layer of tenancy. " +
			"If you are unsure, then you likely don't need teams.",
		Default: appConfig.JetstreamTeams,
	}, &appConfig.JetstreamTeams)
	if err != nil {
		return
	}

	// Composer
	appConfig.EnableJetstreamTeams()

	err = survey.AskOne(&survey.MultiSelect{
		Message: "Choose Composer dependencies to preinstall:",
		Options: appConfig.ComposerDeps.ToOptionsSlice(),
		Default: appConfig.ComposerDeps.ToDefaultSlice(),
	}, &appConfig.ComposerDeps)
	if err != nil {
		return
	}

	err = survey.AskOne(&survey.MultiSelect{
		Message: "Choose NPM dependencies to preinstall:",
		Options: appConfig.NpmDeps.ToOptionsSlice(),
		Default: appConfig.NpmDeps.ToDefaultSlice(),
	}, &appConfig.NpmDeps)
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
