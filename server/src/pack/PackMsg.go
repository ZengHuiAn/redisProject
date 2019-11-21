package pack

import (
	"bytes"
	"encoding/binary"
)

func Pack_bool_data(value bool) []byte {
	if value == true {
		return []byte{1}
	}

	return []byte{0}
}

// 对应c#的char
func Pack_rune_data(value rune) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	//var int32value = int32(value)
	binary.Write(buffer, binary.LittleEndian, value)
	return buffer.Bytes()
}

func Pack_Byte_data(value byte) []byte {
	return []byte{value}
}

func Pack_int16_data(value int16) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.LittleEndian, value)
	return buffer.Bytes()
}

func Pack_uint16_data(value uint16) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.LittleEndian, value)
	return buffer.Bytes()
}

func Pack_int32_data(value int32) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.LittleEndian, value)
	return buffer.Bytes()
}


func Pack_uint32_data(value uint32) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.LittleEndian, value)
	return buffer.Bytes()
}

func Pack_int64_data(value int64) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.LittleEndian, value)
	return buffer.Bytes()
}

func Pack_uint64_data(value uint64) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.LittleEndian, value)
	return buffer.Bytes()
}

func Pack_float_data(value float32) []byte  {
	var buffer = bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.LittleEndian, value)
	return buffer.Bytes()
}

func Pack_double_data(value float64) []byte  {
	var buffer = bytes.NewBuffer([]byte{})
	binary.Write(buffer, binary.LittleEndian, value)
	return buffer.Bytes()
}

func Pack_string_data(value string)[]byte  {
	//return value
	var buffer = bytes.NewBuffer([]byte{})
	buffer.WriteString(value)
}

func init() {
	Pack_bool_data(true)
}
