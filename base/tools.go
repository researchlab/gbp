package base

import (
	"errors"
	"reflect"
)

//查询Slice/Array中是否存在重复元素，存在返回true,否则返回false
func IsDuplicate(target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			for j := i + 1; j < targetValue.Len(); j++ {
				if targetValue.Index(i).Interface() == targetValue.Index(j).Interface() {
					return true, nil
				}
			}
		}
	}
	return false, errors.New("not in")
}

//查询Slice/Array/Map中是否存在某个指定的值obj,存在返回true,否则返回false
func Contains(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("not in")
}
