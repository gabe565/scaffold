package laravel

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2/terminal"
	"os"

	"github.com/clevyr/scaffold/internal/appconfig"
	"github.com/clevyr/scaffold/templates"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "laravel",
	Short: "scaffolds a Laravel application",
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true

	appConfig := appconfig.Defaults
	if err := askQuestions(&appConfig); err != nil {
		if errors.Is(err, terminal.InterruptErr) {
			fmt.Println(err)
			return nil
		}
		return err
	}

	if err := os.Setenv("COMPOSER_MEMORY_LIMIT", "-1"); err != nil {
		return err
	}

	if err := appConfig.GenerateAppKey(); err != nil {
		return err
	}

	if err := initLaravel(appConfig); err != nil {
		return err
	}

	if err := appConfig.SetPackageName(); err != nil {
		return err
	}

	if err := templates.Laravel10BeforeComposer.Generate(appConfig); err != nil {
		return err
	}

	if err := appConfig.ComposerDeps.InstallDeps(); err != nil {
		return err
	}

	if err := appConfig.NpmDeps.InstallDeps(); err != nil {
		return err
	}

	if err := templates.Laravel20AfterComposer.Generate(appConfig); err != nil {
		return err
	}

	if err := runLaravelPint(); err != nil {
		return err
	}

	if err := appConfig.NpmDeps.Install(); err != nil {
		return err
	}

	return nil
}
