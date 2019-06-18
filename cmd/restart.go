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

type RestartOptions struct {
	Objects    []string
	Namespace  []string
	EnableAll  bool
	Enable     []string
	DisableAll bool
	Disable    []string
	outWriter  io.Writer
	errWriter  io.Writer
}

func Execute(outWriter, errWriter io.Writer) int {
	option := NewRestartOptions(outWriter, errWriter)
	cmd := NewRootCommand(option)
	addCommands(cmd, option)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(errWriter, err)
		return exitCodeErr
	}
	return exitCodeOK
}

func addCommands(rootCmd *cobra.Command, o *RestartOptions) {
	rootCmd.AddCommand(
		NewVersionCmd(o),
	)
}

func NewRootCommand(option *RestartOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restart-object",
		Short: "Restart Kubernetes Object",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd)
		},
	}
	fset := cmd.Flags()
	fset.StringSliceVar(&option.Objects, "objects", []string{"deployment"}, "Restart objects")
	fset.StringSliceVar(&option.Namespace, "namespace", []string{}, "Namespace")

	fset.BoolVar(&option.EnableAll, "enable-all", false, "Enable all objects")
	fset.StringSliceVar(&option.Enable, "enable", []string{}, "Enable objects names")

	fset.BoolVar(&option.EnableAll, "disable-all", false, "Disable all objects")
	fset.StringSliceVar(&option.Enable, "disable", []string{}, "Disable objects names")

	return cmd
}

func NewRestartOptions(outWriter, errWriter io.Writer) *RestartOptions {
	return &RestartOptions{
		outWriter: outWriter,
		errWriter: errWriter,
	}
}

func run(cmd *cobra.Command) error {
	// print usage
	if err := cmd.Usage(); err != nil {
		return err
	}
	return nil
}
