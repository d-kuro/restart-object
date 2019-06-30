package objects

import (
	"errors"
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
	EnableSet []string
}

func NewDaemonSetRestarter(clientSet kubernetes.Interface,
	namespace string, tag string, enableSet []string) Restarter {
	return &DaemonSetRestarter{
		ClientSet: clientSet,
		Namespace: namespace,
		Tag:       tag,
		EnableSet: enableSet,
	}
}

func (d *DaemonSetRestarter) List() ([]runtime.Object, error) {
	daemonsets, err := d.ClientSet.AppsV1().DaemonSets(d.Namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	objects := make(map[string]runtime.Object)
	for _, dep := range daemonsets.Items {
		dep := dep
		for _, container := range dep.Spec.Template.Spec.Containers {
			container := container

			logger.Logger().Info("get daemonset",
				zap.String("namespace", d.Namespace),
				zap.String("daemonset-name", dep.Name),
				zap.String("image", container.Image))

			tag := strings.Split(container.Image, ":")[1]
			if tag == d.Tag {
				objects[dep.Name] = &dep
			}
		}
	}

	if len(objects) == 0 {
		return nil, errors.New("restart target not found")
	}

	if len(d.EnableSet) == 0 {
		result := make([]runtime.Object, 0, len(objects))
		for k, v := range objects {
			logger.Logger().Info("restart target",
				zap.String("daemonset-name", k),
				zap.String("namespace", d.Namespace),
				zap.String("image-tag", d.Tag))

			result = append(result, v)
		}
		return result, nil
	}

	result := make([]runtime.Object, 0)
	for _, enable := range d.EnableSet {
		if v, ok := objects[enable]; ok {
			logger.Logger().Info("restart target",
				zap.String("daemonset-name", enable),
				zap.String("namespace", d.Namespace),
				zap.String("image-tag", d.Tag))

			result = append(result, v)
		}
	}
	return result, nil
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
