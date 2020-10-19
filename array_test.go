package utils

import "testing"

func TestRemoveRepeatedValues(t *testing.T) {
	var int64s = []int64{1, 2, 3, 4, 5, 67, 7, 2, 3, 4, 2, 1}
	RemoveRepeatedInt64(int64s)
}

func TestArrayColumns(t *testing.T) {
	s1 := []ArrayMap{map[string]interface{}{
		"key":   1,
		"value": "2",
	}}
	_ = ArrayColumns(s1, "key")
}

func TestArrayColumnValue(t *testing.T) {
	s1 := []ArrayMap{map[string]interface{}{
		"key":   1,
		"value": "2",
	}}
	_, _ = ArrayColumnValue(s1, "key", "value")
}

func TestArrayColumnValues(t *testing.T) {
	s1 := []ArrayMap{map[string]interface{}{
		"key":   1,
		"value": "2",
	}}
	_, _ = ArrayColumnValues(s1, "key")
}

func TestArrayKeysInt(t *testing.T) {
	s1 := []ArrayMap{map[string]interface{}{
		"key": 1,
	}}
	ArrayKeysInt(s1, "key")
}

func TestArrayKeysInt64(t *testing.T) {
	s1 := []ArrayMap{map[string]interface{}{
		"key": 1,
	}}
	ArrayKeysInt64(s1, "key")
}

func TestArrayKeysString(t *testing.T) {
	s1 := []ArrayMap{map[string]interface{}{
		"key": "v",
	}}
	ArrayKeysString(s1, "key")
}

func TestRemoveRepeatedInt(t *testing.T) {
	var ints = []int{1, 2, 3, 4, 5, 67, 7, 2, 3, 4, 2, 1}
	RemoveRepeatedInt(ints)
}

func TestRemoveRepeatedInt64(t *testing.T) {
	var int64s = []int64{1, 2, 3, 4, 5, 67, 7, 2, 3, 4, 2, 1}
	RemoveRepeatedInt64(int64s)
}

func TestRemoveRepeatedString(t *testing.T) {
	var ss = []string{"1", "2", "1", "2", "1"}
	RemoveRepeatedString(ss)
}

func TestInArrayInt(t *testing.T) {
	var ints = []int{1, 2, 3, 4, 5, 67, 7, 2, 3, 4, 2, 1}
	InArrayInt(ints, 0)
}

func TestInArrayInt64(t *testing.T) {
	var int64s = []int64{1, 2, 3, 4, 5, 67, 7, 2, 3, 4, 2, 1}
	InArrayInt64(int64s, 0)
}

func TestInArrayString(t *testing.T) {
	var ss = []string{"1", "2", "1", "2", "1"}
	InArrayString(ss, "1")
}