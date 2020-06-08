


# utils

golang工具包

# string
String2Int(s, default string) int

String2Int8(s, default string) int8

String2Int32(s, default string) int32

String2Int64(s, default string) int64

String2Uint(s, default string) uint

String2Uint8(s, default string) uint8

String2Uint32(s, default string) uint32

String2Uint64(s, default string) uint64

String2Float(s string, defaultVal float64) float64 

Int2String(i int) string

Int642String(i int) string

Int2String(i int) string

Float2String(f float) string

IsValidPhoneNum(s sting) bool

IdCardToHideString(s string) string

PhoneToHideString(s string) string 

VerifyEmailFormat(email string) bool

EmailToHideString(s string) string

Utf8SubStr(str string, start, length int) (string, error)

# mapper

the mapper get value through key

func name like ReadWith{type}

ReadWithString、ReadWithInt64、ReadWithSlice、ReadWithMap any more

only support basic type with go

example:
```
mp := NewMapperReader(map[string]interface{}{
    "foo":"123",
    "bar":4567
})
assert(mp.ReadWithString("bar", ""), "4567")
assert(mp.ReadWithInt64("foo", ""), 123)
```


# convert

BasicTypeToString(value interface{}) string

BasicTypeToInt64(value interface{})  int64

# utils

//use json.number

UnmarshalNumber(bt []byte, v interface{}) error 

//html tag

Marshal(v interface{}) ([]byte, error)

GetCurrentPath() string

GetLocalIP() string

# changeLog

### v1.0.1

add MapReader

### v1.0.2

add id worker

### v1.0.1

reflect

### v1.1.4

struct 2 map changed

### v1.1.5

interface to string|int64