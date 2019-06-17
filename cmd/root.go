package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

const (
	exitCodeOK  = 0
	exitCodeErr = 1
)

type writer struct {
	out io.Writer
	err io.Writer
}

func Execute(outWriter, errWriter io.Writer) int {
	w := &writer{
		out: outWriter,
		err: errWriter,
	}

	cmd := NewRootCommand()
	addCommands(cmd, w)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(errWriter, err)
		return exitCodeErr
	}
	return exitCodeOK
}

func addCommands(rootCmd *cobra.Command, w *writer) {
	rootCmd.AddCommand(
		NewVersionCmd(w),
	)
}

func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "restart-object",
		Short: "Restart Kubernetes Object",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd)
		},
	}
}

func run(cmd *cobra.Command) error {
	// print usage
	if err := cmd.Usage(); err != nil {
		return err
	}
	return nil
}
