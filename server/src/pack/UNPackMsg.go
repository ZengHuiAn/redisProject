package pack

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

func UnPack_bool_data(bts []byte) ([]byte, bool) {
	var buffer = bytes.NewBuffer(bts)
	var value, _ = buffer.ReadByte()
	return buffer.Bytes(), value == 1
}

//// 对应c#的char
//func Pack_rune_data(value rune) []byte {
//	var buffer = bytes.NewBuffer([]byte{})
//	//var int32value = int32(value)
//	_ = binary.Write(buffer, binary.LittleEndian, value)
//	return buffer.Bytes()
//}
func UnPack_Byte_data(bts []byte) ([]byte, byte) {
	var buffer = bytes.NewBuffer(bts)
	var value, err = buffer.ReadByte()
	if err != nil {
		panic(err)
	}

	return buffer.Bytes(), value
}

func UnPack_int16_data(bts []byte) ([]byte, int16) {

	var value int16
	var buffer = bytes.NewBuffer([]byte{})
	err := binary.Read(buffer, binary.LittleEndian, &value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes(), value
}

func UnPack_uint16_data(bts []byte) ([]byte, uint16) {

	var value uint16
	var buffer = bytes.NewBuffer([]byte{})
	err := binary.Read(buffer, binary.LittleEndian, &value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes(), value
}

func UnPack_int32_data(bts []byte) ([]byte, int32) {

	var value int32
	var buffer = bytes.NewBuffer(bts)
	err := binary.Read(buffer, binary.LittleEndian, &value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes(), value
}

func UnPack_uint32_data(bts []byte) ([]byte, uint32) {

	var value uint32
	var buffer = bytes.NewBuffer(bts)
	err := binary.Read(buffer, binary.LittleEndian, &value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes(), value
}

func UnPack_int64_data(bts []byte) ([]byte, int64) {

	var value int64
	var buffer = bytes.NewBuffer(bts)
	err := binary.Read(buffer, binary.LittleEndian, &value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes(), value
}

func UnPack_uint64_data(bts []byte) ([]byte, uint64) {

	var value uint64
	var buffer = bytes.NewBuffer(bts)
	err := binary.Read(buffer, binary.LittleEndian, &value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes(), value
}

func UnPack_float_data(bts []byte) ([]byte, float32) {

	var value float32
	var buffer = bytes.NewBuffer(bts)
	err := binary.Read(buffer, binary.LittleEndian, &value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes(), value
}

func UnPack_double_data(bts []byte) ([]byte, float64) {

	var value float64
	var buffer = bytes.NewBuffer(bts)
	err := binary.Read(buffer, binary.LittleEndian, &value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes(), value
}

func UnPack_string_data(bts []byte) ([]byte, string) {
	//return value
	var new_bytes, unpackLen = UnPack_int32_data(bts)
	var string_bytes = new_bytes[0:unpackLen]

	var str = *(*string)(unsafe.Pointer(&string_bytes))

	var out_bytes = new_bytes[unpackLen:]

	return out_bytes, str
}

func UnPack_bytes_data(bts []byte) ([]byte, []byte) {
	// new_bytes 裁剪得到新的字节数组
	var new_bytes, unpackLen = UnPack_int32_data(bts)
	var _bytes = new_bytes[0:unpackLen]

	var out_bytes = new_bytes[unpackLen:]

	return out_bytes, _bytes
}
func UNPack_common(bts []byte) ([]byte, interface{}) {

	var new_bytes, pack_type = UnPack_Byte_data(bts)
	var result interface{}
	var result_bytes []byte
	fmt.Println(EPackType(pack_type))
	switch EPackType(pack_type) {
	case NULL:
		result = nil
		result_bytes = new_bytes
	case BOOL:
		result_bytes, result = UnPack_bool_data(new_bytes)
	case BYTE:
		result_bytes, result = UnPack_Byte_data(new_bytes)
	case INT16:
		result_bytes, result = UnPack_int16_data(new_bytes)
	case UINT16:
		result_bytes, result = UnPack_uint16_data(new_bytes)
	case INT32:
		result_bytes, result = UnPack_int32_data(new_bytes)
	case UINT32:
		result_bytes, result = UnPack_uint32_data(new_bytes)
	case INT64:
		result_bytes, result = UnPack_int64_data(new_bytes)
	case UINT64:
		result_bytes, result = UnPack_uint64_data(new_bytes)
	case SINGLE:
		result_bytes, result = UnPack_float_data(new_bytes)
	case DOUBLE:
		result_bytes, result = UnPack_double_data(new_bytes)
	case STRING:
		result_bytes, result = UnPack_string_data(new_bytes)
	case BYTEARRAY:
		result_bytes, result = UnPack_bytes_data(new_bytes)
	case ARRAY:

		array_bytes, arraylen := UnPack_Byte_data(new_bytes)
		var tempValue []interface{} = make([]interface{}, arraylen)
		for i := 0; i < len(tempValue); i++ {
			tempBytes, itemValue := UNPack_common(array_bytes)
			tempValue[i] = itemValue
			array_bytes = tempBytes
		}

		result_bytes = array_bytes

		result = tempValue
	}

	return result_bytes, result
}
