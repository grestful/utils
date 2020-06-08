package utils

import (
	"encoding/json"
	"fmt"
	"golang.org/x/tools/container/intsets"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func TestStructSetFieldValue(t *testing.T) {

	type E struct{ N string }
	e1 := E{N: "222"}
	StructSetFieldValue(&e1, "N", "fff")
	fmt.Println(e1)

}

func TestInterfaceToInt(t *testing.T)  {
	start := time.Now()
	for i:=0;i<100000000;i++{
	//	v := BasicTypeToInt64(i)
		_ = *(*int32)(unsafe.Pointer(&i))
	}
	middle := time.Now()

//	for i:=0;i<100000000;i++{
////		v := BasicTypeToInt64(i)
//		_ = int32(i)
//	}
//	middle :=  time.Now()

	//fmt.Println(end.UnixNano() - middle.UnixNano())
	fmt.Println(middle.UnixNano() - start.UnixNano())
}

func TestMaxIntToInt(t *testing.T)  {
	i := int64(intsets.MaxInt)
	fmt.Println(i)
	fmt.Println(int32(i))
	fmt.Println(*(*int32)(unsafe.Pointer(&i)))
}


func TestUnmarshalNumber(t *testing.T)  {
	type Num struct {
		I int64 `json:"i"`
	}

	n := map[string]interface{}{
		"i":8102055647965417472,
	}
	//n := &Num{I:8102055647965417472}
	v,_ := json.Marshal(n)

	n1 := &Num{}
	UnmarshalNumber(v, n1)
	fmt.Println(reflect.TypeOf(n1.I).String())
	fmt.Println(n1.I)

}