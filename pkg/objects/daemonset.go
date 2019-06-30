package objects

import (
	"fmt"
	"strings"

	"github.com/d-kuro/restart-object/pkg/logger"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type DaemonSetRestarter struct {
	ClientSet kubernetes.Interface
	Namespace string
	Tag       string
	Enable    []string
	Disable   []string
}

func NewDaemonSetRestarter(cs kubernetes.Interface, namespace string, tag string, enable, disable []string) Restarter {
	return &DaemonSetRestarter{
		ClientSet: cs,
		Namespace: namespace,
		Tag:       tag,
		Enable:    enable,
		Disable:   disable,
	}
}

func (d *DaemonSetRestarter) List() ([]runtime.Object, error) {
	daemonsets, err := d.ClientSet.AppsV1().DaemonSets(d.Namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	objects := make(map[string]runtime.Object)
	for _, ds := range daemonsets.Items {
		ds := ds
		for _, container := range ds.Spec.Template.Spec.Containers {
			container := container

			logger.Logger().Info("get object",
				zap.String("kind", ds.Kind),
				zap.String("apiVersion", ds.APIVersion),
				zap.String("name", ds.Name),
				zap.String("namespace", ds.Namespace),
				zap.String("image", container.Image))

			tag := strings.Split(container.Image, ":")[1]
			if tag == d.Tag {
				objects[ds.Name] = &ds
			}
		}
	}

	return PickValidObjects(objects, d.Enable, d.Disable)
}

func (d *DaemonSetRestarter) Restart(objects []runtime.Object) error {
	for _, obj := range objects {
		b, err := objectRestarter(obj)
		if err != nil {
			return err
		}

		switch obj := obj.(type) {
		case *extensionsv1beta1.DaemonSet:
			result, err := d.ClientSet.ExtensionsV1beta1().DaemonSets(d.Namespace).
				Patch(obj.Name, types.StrategicMergePatchType, b)
			if err != nil {
				return err
			}
			logger.Logger().Info("restart success",
				zap.String("namespace", result.Namespace),
				zap.String("daemonse-name", result.Name),
				zap.String("image-tag", d.Tag))

		case *appsv1.DaemonSet:
			result, err := d.ClientSet.AppsV1().DaemonSets(d.Namespace).
				Patch(obj.Name, types.StrategicMergePatchType, b)
			if err != nil {
				return err
			}
			logger.Logger().Info("restart success",
				zap.String("namespace", result.Namespace),
				zap.String("daemonse-name", result.Name),
				zap.String("image-tag", d.Tag))

		case *appsv1beta2.DaemonSet:
			result, err := d.ClientSet.AppsV1beta2().DaemonSets(d.Namespace).
				Patch(obj.Name, types.StrategicMergePatchType, b)
			if err != nil {
				return err
			}
			logger.Logger().Info("restart success",
				zap.String("namespace", result.Namespace),
				zap.String("daemonse-name", result.Name),
				zap.String("image-tag", d.Tag))

		default:
			return fmt.Errorf("restarting is not supported")
		}
	}
	return nil
}
