package cli

import (
	"fmt"
	"runtime"

	"github.com/prastuvwxyz/go-migrate/internal/template"
	"github.com/spf13/cobra"
)

var Version string

func BuildVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number.",
		RunE: runCmdE(func(cmd *cobra.Command, args []string) error {
			fmt.Printf("version: %s, built for %s\n", Version, runtime.GOOS)
			return nil
		}),
		Annotations: map[string]string{
			"group": Settings,
		},
	}
	cmd.SetUsageTemplate(template.Usage)
	return cmd
}
