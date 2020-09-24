package utils

import (
	"encoding/json"
	"strconv"
)

// interface => int64
func Interface2Int64(value interface{}) (s int64) {
	switch t := value.(type) {
	case nil:
		s = 0
	case int8:
		s = int64(value.(int8))
	case uint8:
		s = int64(value.(uint8))
	case int:
		s = int64(value.(int))
	case uint:
		s = int64(value.(uint))
	case int64:
		s = value.(int64)
	case uint64:
		s = int64(value.(uint64))
	case int32:
		s = int64(value.(int32))
	case uint32:
		s = int64(value.(uint32))
	case float64:
		s = int64(value.(float64))
	case float32:
		s = int64(value.(float32))
	case bool:
		if t {
			s = 1
		} else {
			s = 0
		}
	case string:
		s = String2Int64(t, 0)
	case json.Number:
		s, _ = t.Int64()
	case json.Delim:
		s = String2Int64(t.String(), 0)
	default:
		s = 0
	}

	return s
}

// interface => string
func Interface2String(value interface{}) (s string) {
	switch t := value.(type) {
	case nil:
		s = ""
	case int8:
		s = strconv.Itoa(int(value.(int8)))
	case uint8:
		s = strconv.Itoa(int(value.(uint8)))
	case int:
		s = strconv.Itoa(t)
	case uint:
		s = strconv.Itoa(int(value.(uint)))
	case int64:
		s = strconv.FormatInt(value.(int64), 10)
	case uint64:
		s = strconv.FormatUint(value.(uint64), 10)
	case int32:
		s = strconv.FormatInt(int64(value.(int32)), 10)
	case uint32:
		s = strconv.FormatInt(int64(value.(uint32)), 10)
	case float64:
		s = strconv.FormatFloat(t, 'f', 0, 64)
	case float32:
		s = strconv.FormatFloat(float64(t), 'f', 0, 64)
	case bool:
		if t {
			s = "true"
		} else {
			s = "false"
		}
	case string:
		s = t
	case json.Number:
		s = t.String()
	case json.Delim:
		s = t.String()
	default:
		s = ""
	}

	return s
}
