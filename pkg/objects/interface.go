package objects

import "k8s.io/apimachinery/pkg/runtime"

type ObjectRestarter interface {
	List() ([]runtime.Object, error)
	Restart(objects []runtime.Object) error
}
