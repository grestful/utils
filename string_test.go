package utils

import (
	"fmt"
	"testing"
)

func TestCompareVersion(t *testing.T) {
	str1 := "3.9"
	str2 := "3.10"

	fmt.Println(CompareVersion(str1, str2))
	//
	//str1 := "3.2.111"
	//str2 := "3.10.222"
	//
	//fmt.Println(CompareVersion(str1, str2))
	str1 = "3.9.1173"
	str2 = "3.9.1173"
	fmt.Println(CompareVersion(str1, str2))

	str1 = "3.10.1174"
	str2 = "3.9.1173"
	fmt.Println(CompareVersion(str1, str2))

	str1 = "3.10.1174"
	str2 = "3.9"
	fmt.Println(CompareVersion(str1, str2))

	fmt.Println(CompareVersion(str2, str1))

}
