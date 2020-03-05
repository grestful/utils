package utils


type MapReader struct {
	value map[string]interface{}
}

func NewMapperReader(m map[string]interface{}) *MapReader {
	return &MapReader{value:m}
}

func (m *MapReader) ReadWithString(key string, defaultVal string) string {
	if m.exists(key) {
		if v,ok := m.value[key].(string); ok {
			return v
		}
	}
	return defaultVal
}


func (m *MapReader) ReadWithInt64(key string, defaultVal int64) int64 {
	if m.exists(key) {
		if v,ok := m.value[key].(int64); ok {
			return v
		}
	}
	return defaultVal
}

func (m *MapReader) ReadWithUint64(key string, defaultVal uint64) uint64 {
	if m.exists(key) {
		if v,ok := m.value[key].(uint64); ok {
			return v
		}
	}
	return defaultVal
}

func (m *MapReader) ReadWithInt(key string, defaultVal int) int {
	if m.exists(key) {
		if v,ok := m.value[key].(int); ok {
			return v
		}
	}
	return defaultVal
}

func (m *MapReader) ReadWithUint8(key string, defaultVal uint8) uint8 {
	if m.exists(key) {
		if v,ok := m.value[key].(uint8); ok {
			return v
		}
	}
	return defaultVal
}

func (m *MapReader) ReadWithUint(key string, defaultVal uint) uint {
	if m.exists(key) {
		if v,ok := m.value[key].(uint); ok {
			return v
		}
	}
	return defaultVal
}

func (m *MapReader) ReadWithInt8(key string, defaultVal int8) int8 {
	if m.exists(key) {
		if v,ok := m.value[key].(int8); ok {
			return v
		}
	}
	return defaultVal
}

func (m *MapReader) ReadWithUint16(key string, defaultVal uint16) uint16 {
	if m.exists(key) {
		if v,ok := m.value[key].(uint16); ok {
			return v
		}
	}
	return defaultVal
}

func (m *MapReader) ReadWithInt32(key string, defaultVal int32) int32 {
	if m.exists(key) {
		if v,ok := m.value[key].(int32); ok {
			return v
		}
	}
	return defaultVal
}

func (m *MapReader) ReadWithBool(key string, defaultVal bool) bool {
	if m.exists(key) {
		if v,ok := m.value[key].(bool); ok {
			return v
		}
	}
	return defaultVal
}

func (m *MapReader) ReadWithInterface(key string, defaultVal interface{}) interface{} {
	if m.exists(key) {
		return m.value[key]
	}
	return defaultVal
}


func (m *MapReader) ReadWithSlice(key string) []interface{} {
	if m.exists(key) {
		if v,ok := m.value[key].([]interface{}); ok {
			return v
		}
	}
	return nil
}


func (m *MapReader) ReadWithMap(key string) map[string]interface{} {
	if m.exists(key) {
		if v,ok := m.value[key].( map[string]interface{} ); ok {
			return v
		}
	}
	return nil
}

func (m *MapReader) exists(key string) bool {
	_,ok := m.value[key]

	return ok
}