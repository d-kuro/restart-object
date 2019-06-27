package objects

import (
	"errors"
	"fmt"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/kubectl/scheme"
)

func objectRestarter(obj runtime.Object) ([]byte, error) {
	switch obj := obj.(type) {
	case *extensionsv1beta1.Deployment:
		if obj.Spec.Paused {
			return nil, errors.New("can't restart paused deployment (run rollout resume first)")
		}
		if obj.Spec.Template.ObjectMeta.Annotations == nil {
			obj.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
		}
		obj.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)
		return runtime.Encode(scheme.Codecs.LegacyCodec(extensionsv1beta1.SchemeGroupVersion), obj)

	case *appsv1.Deployment:
		if obj.Spec.Paused {
			return nil, errors.New("can't restart paused deployment (run rollout resume first)")
		}
		if obj.Spec.Template.ObjectMeta.Annotations == nil {
			obj.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
		}
		obj.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)
		return runtime.Encode(scheme.Codecs.LegacyCodec(appsv1.SchemeGroupVersion), obj)

	case *appsv1beta2.Deployment:
		if obj.Spec.Paused {
			return nil, errors.New("can't restart paused deployment (run rollout resume first)")
		}
		if obj.Spec.Template.ObjectMeta.Annotations == nil {
			obj.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
		}
		obj.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)
		return runtime.Encode(scheme.Codecs.LegacyCodec(appsv1beta2.SchemeGroupVersion), obj)

	case *appsv1beta1.Deployment:
		if obj.Spec.Paused {
			return nil, errors.New("can't restart paused deployment (run rollout resume first)")
		}
		if obj.Spec.Template.ObjectMeta.Annotations == nil {
			obj.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
		}
		obj.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)
		return runtime.Encode(scheme.Codecs.LegacyCodec(appsv1beta1.SchemeGroupVersion), obj)

	case *extensionsv1beta1.DaemonSet:
		if obj.Spec.Template.ObjectMeta.Annotations == nil {
			obj.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
		}
		obj.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)
		return runtime.Encode(scheme.Codecs.LegacyCodec(extensionsv1beta1.SchemeGroupVersion), obj)

	case *appsv1.DaemonSet:
		if obj.Spec.Template.ObjectMeta.Annotations == nil {
			obj.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
		}
		obj.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)
		return runtime.Encode(scheme.Codecs.LegacyCodec(appsv1.SchemeGroupVersion), obj)

	case *appsv1beta2.DaemonSet:
		if obj.Spec.Template.ObjectMeta.Annotations == nil {
			obj.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
		}
		obj.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)
		return runtime.Encode(scheme.Codecs.LegacyCodec(appsv1beta2.SchemeGroupVersion), obj)

	case *appsv1.StatefulSet:
		if obj.Spec.Template.ObjectMeta.Annotations == nil {
			obj.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
		}
		obj.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)
		return runtime.Encode(scheme.Codecs.LegacyCodec(appsv1.SchemeGroupVersion), obj)

	case *appsv1beta1.StatefulSet:
		if obj.Spec.Template.ObjectMeta.Annotations == nil {
			obj.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
		}
		obj.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)
		return runtime.Encode(scheme.Codecs.LegacyCodec(appsv1beta1.SchemeGroupVersion), obj)

	case *appsv1beta2.StatefulSet:
		if obj.Spec.Template.ObjectMeta.Annotations == nil {
			obj.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
		}
		obj.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)
		return runtime.Encode(scheme.Codecs.LegacyCodec(appsv1beta2.SchemeGroupVersion), obj)

	default:
		return nil, fmt.Errorf("restarting is not supported")
	}
}
