package utils

import "strconv"

// Itoa converts an interface to a string.
func BasicTypeToString(value interface{}) (s string) {
	switch t := value.(type) {
	case nil:
		s =  ""
	case int8:
		s =  strconv.Itoa(int(value.(int8)))
	case uint8:
		s =  strconv.Itoa(int(value.(uint8)))
	case int:
		s =  strconv.Itoa(t)
	case uint:
		s =  strconv.Itoa(int(value.(uint)))
	case int64:
		s =  strconv.FormatInt(value.(int64), 10)
	case uint64:
		s =  strconv.FormatUint(value.(uint64), 10)
	case int32:
		s =  strconv.FormatInt(int64(value.(int32)), 10)
	case uint32:
		s =  strconv.FormatInt(int64(value.(uint32)), 10)
	case float64:
		s =  strconv.FormatFloat(t, 'f', 0, 64)
	case float32:
		s =  strconv.FormatFloat(float64(t), 'f', 0, 64)
	case bool:
		if t {
			s =  "true"
		} else {
			s =  "false"
		}
	case string:
		s =  t
	default:
		s =  ""
	}

	return s
}

// Itoa converts an interface to a string.
func BasicTypeToInt64(value interface{}) (i int64) {
	switch t := value.(type) {
	case nil:
		i = 0
	case int8:
		i = int64(value.(int8))
	case uint8:
		i = int64(value.(uint8))
	case int:
		i = int64(t)
	case uint:
		i = int64(value.(uint))
	case int64:
		i = value.(int64)
	case uint64:
		i = int64(value.(uint64))
	case int32:
		i = int64(value.(int32))
	case uint32:
		i = int64(value.(uint32))
	case float64:
		i = int64(value.(float64))
	case float32:
		i = int64(value.(float32))
	case bool:
		if t {
			i = 1
		} else {
			i = 0
		}
	case string:
		i,_ = strconv.ParseInt(value.(string), 10, 64)
	default:
		i = 0
	}

	return i
}
