package utils

import (
	"math/rand"
	"time"
)

// 获取本地时间
func CovertTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func String2Time(s string) time.Time {
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", s, loc)
	return t
}

func String2TimeD(s string) time.Time {
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02", s, loc)
	return t
}

// 时间转换 -> 天
func Time2StringD(t time.Time) string {
	return t.Format("2006-01-02")
}

// 获取当前的时间
func Time2StringByNow() string {
	t := GetNowDayFinishTime()
	return t.Format("20060102")
}

// 时间转换 -> 小时
func Time2StringH(t time.Time) string {
	return t.Format("2006-01-02 15")
}

// 时间转换 -> 分钟
func Time2StringM(t time.Time) string {
	return t.Format("2006-01-02 15:04")
}

// 时间转换 -> 秒
func Time2StringS(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// 获取秒数据
func GetSecond() int64 {
	return time.Now().Unix()
}

// 获取毫秒数据
func GetNSecond() int64 {
	return time.Now().UnixNano() / 1e6
}

// 获取纳秒数据
func GetNNSecond() int64 {
	return time.Now().UnixNano()
}

// 获取具体时间String
func GetTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当天的起始时间
func GetNowDayStartTime() time.Time {
	t := GetSecond() - 24*3600
	return CovertTime(t)
}

// 获取当天的结束时间
func GetNowDayFinishTime() time.Time {
	t := GetSecond()
	return CovertTime(t)
}

func Abs(n int64) int64 {
	if n < 0 {
		return (-1) * n
	}
	return n
}

func AbsFloat(n float64) float64 {
	if n < 0 {
		return (-1) * n
	}
	return n
}

func rapidRandom() string {
	// 生成的UUID, 参考java的UUID实现
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 16; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	result[6] &= 0x0f /* clear version        */
	result[6] |= 0x40 /* set to version 4     */
	result[8] &= 0x3f /* clear variant        */
	result[8] |= 0x80 /* set to IETF variant  */
	var msb, lsb int64
	for i := 0; i < 8; i++ {
		msb = (msb << 8) | int64(result[i]&0xff)
	}
	for i := 8; i < 16; i++ {
		lsb = (lsb << 8) | int64(result[i]&0xff)
	}
	v := Abs(msb ^ lsb)
	return Int642String(v)
}
