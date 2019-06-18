package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "v0.0.1"

var Revision = "development"

func NewVersionCmd(o *RestartOptions) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version",
		Run: func(cmd *cobra.Command, args []string) {
			versionCmd(o)
		},
	}
}

func versionCmd(o *RestartOptions) {
	fmt.Fprintf(o.outWriter, "version: %s (rev: %s)\n", Version, Revision)
}
