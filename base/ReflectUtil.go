package base

import "reflect"

func GetTypeOfAndValueOf(obj interface{}) (t reflect.Type, v reflect.Value) {
	t = reflect.TypeOf(obj)
	v = reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	return
}
