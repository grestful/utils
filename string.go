package utils

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 转换默认值
func String2Int(s string, defaultVal int) int {
	a, err := strconv.Atoi(s)
	if err != nil {
		return defaultVal
	}

	return a
}

func String2Int64(s string, defaultVal int64) int64 {
	a, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultVal
	}

	return a
}
func String2Int8(s string, defaultVal int8) int8 {
	a, err := strconv.Atoi(s)
	if err != nil {
		return defaultVal
	}

	return int8(a)
}

func String2Float(s string, defaultVal float64) float64 {
	a, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultVal
	}

	return a
}

func Int2String(a int) string {
	s := strconv.Itoa(a)
	return s
}

func Int642String(a int64) string {
	s := strconv.FormatInt(a, 10)
	return s
}

func Uint642String(a uint64) string {
	s := strconv.FormatUint(a, 10)
	return s
}

func Int82String(a int8) string {
	s := strconv.Itoa(int(a))
	return s
}
func String2Uint64(s string, defaultVal uint64) uint64 {
	a, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return defaultVal
	}

	return a
}

func Float2String(a float64) string {
	s := strconv.FormatFloat(a, 'f', 6, 64)
	return s
}

//[]uint8 to string
func B2S(bs []uint8) string {
	ba := make([]byte, 0)
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

func Utf8SubStr(str string, start, length int) (string, error) {
	l := utf8.RuneCountInString(str)
	if l == 0 {
		return "", errors.New("Str is empty! ")
	}

	if start < 0 {
		start = l + start
	}

	if (start + length) > l {
		return str, errors.New("Str length is less than start+length! ")
	}
	sb := new(strings.Builder)
	var pos = 0
	for _, c := range str {
		if pos >= start {
			if pos >= start+length {
				break
			}
			sb.WriteRune(c)
		}
		pos++
	}

	return sb.String(), nil
}

