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