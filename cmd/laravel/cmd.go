package laravel

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/clevyr/scaffold/internal/appconfig"
	"github.com/clevyr/scaffold/templates"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "laravel",
	Short: "scaffolds a Laravel application",
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) (err error) {
	cmd.SilenceUsage = true

	appConfig := appconfig.Defaults
	err = askQuestions(&appConfig)
	if err == terminal.InterruptErr {
		fmt.Println("Interrupted")
		return nil
	} else if err != nil {
		return err
	}

	err = os.Setenv("COMPOSER_MEMORY_LIMIT", "-1")
	if err != nil {
		return err
	}

	err = appConfig.GenerateAppKey()
	if err != nil {
		return err
	}

	err = initLaravel(appConfig)
	if err != nil {
		return err
	}

	err = templates.Laravel10BeforeComposer.Generate(appConfig)
	if err != nil {
		return err
	}

	err = appConfig.ComposerDeps.InstallDeps()
	if err != nil {
		return err
	}

	err = appConfig.NpmDeps.InstallDeps()
	if err != nil {
		return err
	}

	err = templates.Laravel20AfterComposer.Generate(appConfig)
	if err != nil {
		return err
	}

	err = appConfig.NpmDeps.Install()
	if err != nil {
		return err
	}

	return nil
}
