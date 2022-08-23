package cmd

import (
	"github.com/clevyr/scaffold/cmd/laravel"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
	"path/filepath"
)

var Command = &cobra.Command{
	Use:     "scaffold",
	Short:   "scaffolds out an application",
	Version: buildVersion(),

	PersistentPreRunE: preRun,
}

var (
	directory   string
	versionFlag bool
)

func Execute() error {
	cmd, _, err := Command.Find(os.Args[1:])
	// Call "laravel" subcommand if no other command is given
	if err == nil && cmd.Use == Command.Use && cmd.Flags().Parse(os.Args[1:]) != pflag.ErrHelp && !versionFlag {
		args := append([]string{laravel.Command.Use}, os.Args[1:]...)
		Command.SetArgs(args)
	}

	return Command.Execute()
}

func init() {
	Command.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "version for scaffold")

	Command.PersistentFlags().StringVarP(&directory, "directory", "C", "", "Run as if the application was started in the given path.")
	err := Command.RegisterFlagCompletionFunc(
		"directory",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveFilterDirs
		})
	if err != nil {
		panic(err)
	}

	Command.AddCommand(
		laravel.Command,
	)
}

func preRun(cmd *cobra.Command, args []string) (err error) {
	cmd.SilenceUsage = true

	directory = filepath.Clean(directory)
	if directory != "." {
		if err = os.MkdirAll(directory, 0755); err != nil {
			return err
		}

		if err = os.Chdir(directory); err != nil {
			return err
		}
	}

	return nil
}
