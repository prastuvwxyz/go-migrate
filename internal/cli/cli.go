package cli

import (
	"os"

	"github.com/spf13/cobra"
)

func runCmdE(f func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 && args[0] == "help" {
			_ = cmd.Help() // Help always returns nil.
			os.Exit(0)
		}
		return f(cmd, args)
	}
}
