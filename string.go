package utils

import (
	"errors"
	"regexp"
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

//验证邮箱
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 转换email中间中间的字符
func EmailToHideString(s string) string {
	sr := strings.Split(s, "@")
	if len(sr) > 1 {
		sr[0] = HideMbStr(sr[0], 3)
	}

	return strings.Join(sr, "@")
}

func HideMbStr(s string, subLen int) string {
	l := utf8.RuneCountInString(s)
	if l < 3 {
		subLen = 1
	}
	s, _ = Utf8SubStr(s, 0, subLen)
	for i := 0; i < l-subLen; i++ {
		s += "*"
	}
	return s
}

// 转换手机号中间中间的字符
func PhoneToHideString(s string) string {
	if IsValidPhoneNum(s) {
		return s[0:3] + "****" + s[7:]
	}

	return HideMbStr(s, 3)
}

// 身份证中间的字符
func IdCardToHideString(s string) string {
	l := len(s)
	if l == 18 {
		return s[0:6] + "********" + s[14:]
	}

	if l == 15 {
		return s[0:6] + "********" + s[11:]
	}

	if l > 10 {
		return s[0:2] + "********" + s[11:]
	}

	if l < 5 {
		return s
	}
	startHidden := 2
	middle := l/2 + 1
	str := s[0:startHidden]
	max := middle + (middle - startHidden - 1)
	if l%2 == 0 {
		max = max - 1
	}
	for i := startHidden; i < max; i++ {
		str += "*"
	}
	str += s[l-startHidden:]
	return str
}

var phoneRegex = regexp.MustCompile(`^((13[0-9])|(14[0-9])|(15[0-9])|(16[0-9])|(17[0-9])|(18[0-9])|(19[0-9]))\d{8}$`)

// 判断手机号是否合法
func IsValidPhoneNum(phone string) bool {
	if phoneRegex.MatchString(phone) {
		return true
	} else {
		return false
	}
}
