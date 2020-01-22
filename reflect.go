package utils

import (
	"reflect"
)

func StructExistsProperty(s interface{}, key string) bool {
	t := reflect.TypeOf(s)
	if _, ok := t.FieldByName("Name"); ok {
		return true
	}

	return false
}

func SliceBuildWithInterface(v interface{}) reflect.Value {
	typ := reflect.TypeOf(v)
	sliceValue := reflect.MakeSlice(reflect.SliceOf(typ), 0, 0)

	slice := reflect.New(sliceValue.Type())
	slice.Elem().Set(sliceValue)

	return slice
}

