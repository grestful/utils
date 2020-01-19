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
