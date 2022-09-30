package laravel

import (
	"errors"
	"fmt"
	"github.com/clevyr/scaffold/internal/appconfig"
	"os"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/huandu/xstrings"
)

var validationRegex, _ = regexp.Compile("^[0-9]*[kmg]$")

func askQuestions(appConfig *appconfig.AppConfig) error {
	var err error

	// App Name
	err = survey.AskOne(&survey.Input{
		Message: "Enter the application name:",
		Default: appConfig.AppName,
		Help:    "Enter the application's name space-separated and capitalized. A slug will be generated to match.",
	}, &appConfig.AppName, survey.WithValidator(survey.Required))
	if err != nil {
		return err
	}
	appConfig.AppSlug = xstrings.ToKebabCase(appConfig.AppName)

	// Laravel Install
	_, err = os.Stat("composer.json")
	if os.IsNotExist(err) {
		wd, _ := os.Getwd()
		home, _ := os.UserHomeDir()
		wd = strings.Replace(wd, home, "~", 1)

		err = survey.AskOne(&survey.Confirm{
			Message: fmt.Sprintf("Initialize a Laravel app in \"%s/%s\"?", wd, appConfig.AppSlug),
			Default: true,
		}, &appConfig.InitLaravel)
		if err != nil {
			return err
		}
	}

	// Database
	err = survey.AskOne(&survey.Select{
		Message: "Choose default database server:",
		Options: []string{"PostgreSQL", "MariaDB"},
		Default: appConfig.Database,
	}, &appConfig.Database, survey.WithValidator(survey.Required))
	if err != nil {
		return err
	}

	// Enabled PhpModules
	appConfig.EnableSelectedDatabase()
	err = survey.AskOne(&survey.MultiSelect{
		Message: "Choose PHP modules to enable:",
		Options: appConfig.PhpModules.ToOptionsSlice(),
		Default: appConfig.PhpModules.ToDefaultSlice(),
	}, &appConfig.PhpModules)
	if err != nil {
		return err
	}

	// Composer
	err = survey.AskOne(&survey.MultiSelect{
		Message: "Choose Composer dependencies:",
		Options: appConfig.ComposerDeps.ToOptionsSlice(),
		Default: appConfig.ComposerDeps.ToDefaultSlice(),
	}, &appConfig.ComposerDeps)
	if err != nil {
		return err
	}

	// Jetstream Teams
	if appConfig.ComposerDeps.Map["laravel/jetstream"].Enabled {
		err = survey.AskOne(&survey.Confirm{
			Message: "Do you want to use Jetstream with teams?",
			Help: "'Teams' are Jetstream's built-in single layer of tenancy. " +
				"If you are unsure, then you likely don't need teams.",
			Default: appConfig.JetstreamTeams,
		}, &appConfig.JetstreamTeams)
		if err != nil {
			return err
		}
		appConfig.EnableJetstreamTeams()
	}

	err = survey.AskOne(&survey.MultiSelect{
		Message: "Choose NPM dependencies:",
		Options: appConfig.NpmDeps.ToOptionsSlice(),
		Default: appConfig.NpmDeps.ToDefaultSlice(),
	}, &appConfig.NpmDeps)
	if err != nil {
		return err
	}

	// Max Upload Size
	err = survey.AskOne(
		&survey.Input{
			Message: "File upload size limit:",
			Default: appConfig.MaxUploadSize,
			Help: "Configures the maximum allowed upload file size. " +
				"Supports suffixes \"k\" (kilobytes), \"m\" (megabytes) and \"g\" (gigabytes).",
		},
		&appConfig.MaxUploadSize,
		survey.WithValidator(func(val any) error {
			if str, ok := val.(string); !ok || !validationRegex.MatchString(str) {
				return errors.New("please enter a size followed by \"k\" (kilobytes), \"m\" (megabytes) or \"g\" (gigabytes)")
			}
			return nil
		}),
	)
	if err != nil {
		return err
	}

	err = survey.AskOne(&survey.Confirm{
		Message: "Do you want to create an initial Git commit?",
		Default: appConfig.GitCommit,
	}, &appConfig.GitCommit)
	if err != nil {
		return err
	}

	return err
}
