package objects

import (
	"errors"
	"fmt"
	"strings"

	"github.com/d-kuro/restart-object/pkg/logger"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type DeploymentRestarter struct {
	ClientSet kubernetes.Interface
	Namespace string
	Tag       string
	EnableSet []string
}

func NewDeploymentRestarter(clientSet kubernetes.Interface,
	namespace string, tag string, enableSet []string) *DeploymentRestarter {
	return &DeploymentRestarter{
		ClientSet: clientSet,
		Namespace: namespace,
		Tag:       tag,
		EnableSet: enableSet,
	}
}

func (d *DeploymentRestarter) List() ([]runtime.Object, error) {
	deployments, err := d.ClientSet.AppsV1().Deployments(d.Namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	objects := make(map[string]runtime.Object)
	for _, dep := range deployments.Items {
		dep := dep
		for _, container := range dep.Spec.Template.Spec.Containers {
			container := container

			logger.Logger().Info("get deployment",
				zap.String("namespace", d.Namespace),
				zap.String("deployment-name", dep.Name),
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
				zap.String("deployment-name", k),
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
				zap.String("deployment-name", enable),
				zap.String("namespace", d.Namespace),
				zap.String("image-tag", d.Tag))

			result = append(result, v)
		}
	}
	return result, nil
}

func (d *DeploymentRestarter) Restart(objects []runtime.Object) error {
	for _, obj := range objects {
		b, err := objectRestarter(obj)
		if err != nil {
			return err
		}

		switch obj := obj.(type) {
		case *appsv1.Deployment:
			result, err := d.ClientSet.AppsV1().Deployments(d.Namespace).
				Patch(obj.Name, types.StrategicMergePatchType, b)
			if err != nil {
				return err
			}
			logger.Logger().Info("restart success",
				zap.String("namespace", result.Namespace),
				zap.String("deployment-name", result.Name),
				zap.String("image-tag", d.Tag))

		case *extensionsv1beta1.Deployment:
			result, err := d.ClientSet.ExtensionsV1beta1().Deployments(d.Namespace).
				Patch(obj.Name, types.StrategicMergePatchType, b)
			if err != nil {
				return err
			}
			logger.Logger().Info("restart success",
				zap.String("namespace", result.Namespace),
				zap.String("deployment-name", result.Name),
				zap.String("image-tag", d.Tag))

		case *appsv1beta2.Deployment:
			result, err := d.ClientSet.AppsV1beta1().Deployments(d.Namespace).
				Patch(obj.Name, types.StrategicMergePatchType, b)
			if err != nil {
				return err
			}
			logger.Logger().Info("restart success",
				zap.String("namespace", result.Namespace),
				zap.String("deployment-name", result.Name),
				zap.String("image-tag", d.Tag))

		case *appsv1beta1.Deployment:
			result, err := d.ClientSet.AppsV1beta2().Deployments(d.Namespace).
				Patch(obj.Name, types.StrategicMergePatchType, b)
			if err != nil {
				return err
			}
			logger.Logger().Info("restart success",
				zap.String("namespace", result.Namespace),
				zap.String("deployment-name", result.Name),
				zap.String("image-tag", d.Tag))

		default:
			return fmt.Errorf("restarting is not supported")
		}
	}
	return nil
}
