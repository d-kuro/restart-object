package objects

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
)

type Restarter interface {
	List() ([]runtime.Object, error)
	Restart(objects []runtime.Object) error
}

type NewRestarterFunc func(cs kubernetes.Interface, namespace string, tag string, enable, disable []string) Restarter

func NewRestarterInitializers() map[string]NewRestarterFunc {
	restarters := make(map[string]NewRestarterFunc, 2)
	restarters["deployment"] = NewDeploymentRestarter
	restarters["daemonset"] = NewDaemonSetRestarter
	return restarters
}
