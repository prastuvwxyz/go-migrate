package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/prastuvwxyz/go-migrate/internal/cli"
	"github.com/prastuvwxyz/go-migrate/internal/term/log"
	"github.com/spf13/cobra"
)

func init() {
	color.NoColor = false
	cobra.EnableCommandSorting = false // Maintain the order in which we add commands.
}

func main() {
	cmd := buildRootCmd()
	if err := cmd.Execute(); err != nil {
		log.Errorln(err.Error())
		os.Exit(1)
	}
}

func buildRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go-migrate",
		Short: shortDescription,
		// Example: `
		// 	Displays the help menu for the "init" command.
		// 	/code $ copilot init --help`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// If we don't set a Run() function the help menu doesn't show up.
			// See https://github.com/spf13/cobra/issues/790
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.SetOut(log.OutputWriter)
	cmd.SetErr(log.DiagnosticWriter)

	// Sets version for --version flag. Version command gives more detailed
	// version information.
	cmd.Version = cli.Version
	cmd.SetVersionTemplate("go-migrate version: {{.Version}}\n")

	// "Settings" command group.
	cmd.AddCommand(cli.BuildVersionCmd())

	return cmd
}
