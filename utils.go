package utils

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Marshal(v interface{}) ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buff)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	return buff.Bytes(), err
}

//获取参数，不支持keys 0
func GetRequestKey(r *http.Request, key string) string {
	keys, ok := r.URL.Query()[key]

	if !ok || len(keys) < 1 {
		c, err := r.Cookie(key)
		if err != nil {
			return ""
		}

		return c.Value
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	return keys[0]
}

func GetCurrentPath() string {
	currentPath, _ := os.Getwd()
	return currentPath
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

func IPStringToInt(ipString string) int {
	ipSegs := strings.Split(ipString, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}

	return ipInt
}


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

func String2Uint64(s string, defaultVal uint64) uint64 {
	a, err := strconv.ParseUint(s, 10, 64)
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

func Float2String(a float64) string {
	s := strconv.FormatFloat(a, 'f', 6, 64)
	return s
}

//[]uint8 to string
func B2S(bs []uint8) string {
	ba := make([]byte,0)
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}
