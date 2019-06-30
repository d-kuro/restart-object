package objects

import (
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
)

func PickValidObjects(objects map[string]runtime.Object, enable, disable []string) ([]runtime.Object, error) {
	if len(enable) == 0 && len(disable) == 0 {
		result := make([]runtime.Object, 0, len(objects))
		for _, v := range objects {
			result = append(result, v)
		}
		return validateReturnResult(result)

	} else if len(enable) == 0 && len(disable) == 0 {
		for _, d := range disable {
			delete(objects, d)
		}
		result := make([]runtime.Object, 0)
		for _, v := range objects {
			result = append(result, v)
		}
		return validateReturnResult(result)
	}

	result := make([]runtime.Object, 0)
	for _, enable := range enable {
		if v, ok := objects[enable]; ok {
			result = append(result, v)
		}
	}
	return validateReturnResult(result)
}

func validateReturnResult(result []runtime.Object) ([]runtime.Object, error) {
	if len(result) == 0 {
		return nil, errors.New("restart target not found")
	}
	return result, nil
}
