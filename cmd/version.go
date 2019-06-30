package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const Version = "v0.0.1"

var Revision = "development"

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version",
		Run: func(cmd *cobra.Command, args []string) {
			runVersionCmd()
		},
	}
}

func runVersionCmd() {
	fmt.Fprintf(os.Stdout, "version: %s (rev: %s)\n", Version, Revision)
}
