package utils

import (
	"errors"
	"fmt"
	"sort"
)

/**
Like Php Array method
ArrayKeysString(slices []ArrayMap, key string) []string
ArrayKeysInt64(slices []ArrayMap, key string) []int64
ArrayKeysInt(slices []ArrayMap, key string) []int
ArrayColumns(slices []ArrayMap, key string) []interface{}
ArrayColumnValue(slices []ArrayMap, key string, index string) (map[string]interface{}, error)
ArrayColumnValues(slices []ArrayMap, key string) (map[string]ArrayMap, error)
InArrayInt64(slice []int64, value int64) int
InArrayInt(slice []int, value int) int
InArrayString(slice []string, value string) int
RemoveRepeatedInt(values []int) []int
RemoveRepeatedInt64(values []int64) []int64
RemoveRepeatedString(values []string) []string
*/
var ArrayMapKexNotExists = "key not exists"
var ArrayMapKexIsEmpty = "slice empty"

type ArrayMap map[string]interface{}

func (am ArrayMap) ExistsKey(key string) bool {
	_, ok := am[key]
	return ok
}

func ArrayKeysString(slices []ArrayMap, key string) []string {
	if len(slices) == 0 {
		return nil
	}

	res := make([]string, 0)
	for _, v := range slices {
		if v.ExistsKey(key) {
			res = append(res, Interface2String(v[key]))
		}
	}

	return res
}

func ArrayKeysInt64(slices []ArrayMap, key string) []int64 {
	if len(slices) == 0 {
		return nil
	}

	res := make([]int64, 0)
	for _, v := range slices {
		if v.ExistsKey(key) {
			res = append(res, Interface2Int64(v[key]))
		}
	}

	return res
}

func ArrayKeysInt(slices []ArrayMap, key string) []int {
	if len(slices) == 0 {
		return nil
	}

	res := make([]int, 0)
	for _, v := range slices {
		if v.ExistsKey(key) {
			res = append(res, Interface2Int(v[key]))
		}
	}

	return res
}

func ArrayColumns(slices []ArrayMap, key string) []interface{} {
	if len(slices) == 0 {
		return nil
	}

	res := make([]interface{}, 0)
	for _, v := range slices {
		if v.ExistsKey(key) {
			res = append(res, v[key])
		}
	}

	return res
}

func ArrayColumnValue(slices []ArrayMap, key string, index string) (map[string]interface{}, error) {
	if len(slices) == 0 {
		return nil, errors.New(ArrayMapKexIsEmpty)
	}

	res := make(map[string]interface{})
	for i, v := range slices {

		if !v.ExistsKey(key) {
			return nil, errors.New(fmt.Sprintf(ArrayMapKexNotExists+", key %d ", i))
		}

		if !v.ExistsKey(index) {
			return nil, errors.New(fmt.Sprintf(ArrayMapKexNotExists+", index %d ", i))
		}

		res[Interface2String(v[index])] = v[key]
	}

	return res, nil
}

func ArrayColumnValues(slices []ArrayMap, key string) (map[string]ArrayMap, error) {
	if len(slices) == 0 {
		return nil, errors.New(ArrayMapKexIsEmpty)
	}

	res := make(map[string]ArrayMap)
	for i, v := range slices {

		if !v.ExistsKey(key) {
			return nil, errors.New(fmt.Sprintf(ArrayMapKexNotExists+", key %d ", i))
		}

		res[Interface2String(v[key])] = v
	}

	return res, nil
}

func InArrayInt64(slice []int64, value int64) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return -1
}

func InArrayInt(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return -1
}

func InArrayString(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return -1
}


func RemoveRepeatedString(values []string) []string {
	res := make([]string, 0)
	sort.Strings(values)
	for i := 0; i < len(values); i++ {
		repeat := false
		for j := i + 1; j < len(values); j++ {
			if values[i] == values[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			res = append(res, values[i])
		}
	}
	return res
}


func RemoveRepeatedInt64(arr []int64) (newArr []int64) {
	newArr = make([]int64, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}


func RemoveRepeatedInt(values []int) []int {
	res := make([]int, 0)
	sort.Ints(values)
	for i := 0; i < len(values); i++ {
		repeat := false
		for j := i + 1; j < len(values); j++ {
			if values[i] == values[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			res = append(res, values[i])
		}
	}
	return res
}
