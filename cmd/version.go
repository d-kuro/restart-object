package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "v0.0.1"

var Revision = "development"

func NewVersionCmd(w *writer) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version",
		Run: func(cmd *cobra.Command, args []string) {
			versionCmd(w)
		},
	}
}

func versionCmd(w *writer) {
	fmt.Fprintf(w.out, "version: %s (rev: %s)\n", Version, Revision)
}
