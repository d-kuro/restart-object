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

	o := NewRestartOptions()
	cmd := NewRootCommand(o)
	addCommands(cmd)

	if err := cmd.Execute(); err != nil {
		logger.Logger().Error("Error", zap.Error(err))
		return exitCodeErr
	}
	return exitCodeOK
}

func addCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(
		NewVersionCmd(),
	)
}

func NewRootCommand(o *RestartOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "restart-object",
		Short:         "Restart Kubernetes Object",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(o)
		},
	}

	fset := cmd.Flags()
	fset.BoolVar(&o.InCluster, "in-cluster", false, "Execute for in Kubernetes cluster")

	fset.StringVar(&o.Object, "objects", "deployment", "Restart objects")
	fset.StringVar(&o.Namespace, "namespace", "default", "Namespace")
	fset.StringVar(&o.Tag, "tag", "latest", "Target to restart image tag name")

	fset.BoolVar(&o.EnableAll, "enable-all", false, "Enable all objects")
	fset.StringSliceVar(&o.Enable, "enable", []string{}, "Enable objects names")

	fset.BoolVar(&o.DisableAll, "disable-all", false, "Disable all objects")
	fset.StringSliceVar(&o.Disable, "disable", []string{}, "Disable objects names")

	return cmd
}

func NewRestartOptions() *RestartOptions {
	return &RestartOptions{}
}

func run(o *RestartOptions) error {
	enableSet, err := EnableSetBuild(o)
	if err != nil {
		return err
	}

	var f util.Factory
	if o.InCluster {
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

	r := objects.NewRestarterInitializers()
	restarter := r[o.Object](cs, o.Namespace, o.Tag, enableSet)
	objects, err := restarter.List()
	if err != nil {
		return err
	}
	if err := restarter.Restart(objects); err != nil {
		return err
	}

	return nil
}
