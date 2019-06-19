package cmd

import (
	"github.com/d-kuro/restart-object/cmd/util"
	"github.com/d-kuro/restart-object/pkg/logger"
	"github.com/d-kuro/restart-object/pkg/objects"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const (
	exitCodeOK  = 0
	exitCodeErr = 1
)

type RestartOptions struct {
	Object     string
	Namespace  string
	Tag        string
	Enable     []string
	Disable    []string
	EnableAll  bool
	DisableAll bool
	InCluster  bool
}

func Execute() int {
	logger.Init(logger.Writer)

	option := NewRestartOptions()
	cmd := NewRootCommand(option)
	addCommands(cmd, option)

	if err := cmd.Execute(); err != nil {
		logger.Logger().Error("Error", zap.Error(err))
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
		Use:           "restart-object",
		Short:         "Restart Kubernetes Object",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(option)
		},
	}

	fset := cmd.Flags()
	fset.BoolVar(&option.InCluster, "in-cluster", false, "Execute for in Kubernetes cluster")

	fset.StringVar(&option.Object, "objects", "deployment", "Restart objects")
	fset.StringVar(&option.Namespace, "namespace", "default", "Namespace")
	fset.StringVar(&option.Tag, "tag", "latest", "Target to restart image tag name")

	fset.BoolVar(&option.EnableAll, "enable-all", false, "Enable all objects")
	fset.StringSliceVar(&option.Enable, "enable", []string{}, "Enable objects names")

	fset.BoolVar(&option.DisableAll, "disable-all", false, "Disable all objects")
	fset.StringSliceVar(&option.Disable, "disable", []string{}, "Disable objects names")

	return cmd
}

func NewRestartOptions() *RestartOptions {
	return &RestartOptions{}
}

func run(option *RestartOptions) error {
	enableSet, err := EnableSetBuild(option)
	if err != nil {
		return err
	}

	var f util.Factory
	if option.InCluster {
		f = util.NewInClusterFactory()
		logger.Logger().Info("execute place: in-cluster")
	} else {
		f = util.NewLocalFactory()
		logger.Logger().Info("execute place: local")
	}

	cs, err := f.ClientSet()
	if err != nil {
		return err
	}

	deployment := objects.NewDeploymentRestarter(cs, option.Namespace, option.Tag, enableSet)
	objects, err := deployment.List()
	if err != nil {
		return err
	}
	if err := deployment.Restart(objects); err != nil {
		return err
	}

	return nil
}
