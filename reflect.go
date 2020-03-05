package utils

import (
	"github.com/danhper/structomap"
	"github.com/mitchellh/mapstructure"
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

//set struct field value
//example:
// type e struct{ N string }
// e1 := &e{N:"222"}
// e1, ok := StructSetFieldValue(e1, "N", "fff")
// params v must ues ptr type
func StructSetFieldValue(v interface{}, key string, value interface{}) bool {
	t := reflect.TypeOf(v)

	var vft reflect.StructField
	var ok bool
	switch t.Kind() {
	case reflect.Ptr:
		vft, ok = reflect.TypeOf(v).Elem().FieldByName(key)
	default:
		return false
	}

	if !ok {
		return false
	}

	if vft.Type.Kind() == reflect.TypeOf(value).Kind() {
		reflect.ValueOf(v).Elem().FieldByName(key).Set(reflect.ValueOf(value))
		return true
	}

	return false
}

func Struct2Map(obj interface{}) map[string]interface{} {
	return structomap.New().UseCamelCase().Transform(obj)
}

func Struct2MapList(obj interface{}) ([]map[string]interface{}, error) {
	return structomap.New().UseCamelCase().TransformArray(obj)
}

//obj must use *point
func Map2Struct(m map[string]interface{}, obj interface{}) error {
	return mapstructure.Decode(m, obj)
}