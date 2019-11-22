package pack

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

func Pack_bool_data(value bool) []byte {
	if value == true {
		return []byte{1}
	}

	return []byte{0}
}

//// 对应c#的char
//func Pack_rune_data(value rune) []byte {
//	var buffer = bytes.NewBuffer([]byte{})
//	//var int32value = int32(value)
//	_ = binary.Write(buffer, binary.LittleEndian, value)
//	return buffer.Bytes()
//}

func Pack_Byte_data(value byte) []byte {
	return []byte{value}
}

func Pack_int16_data(value int16) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	err := binary.Write(buffer, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func Pack_uint16_data(value uint16) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	err := binary.Write(buffer, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func Pack_int32_data(value int32) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	err := binary.Write(buffer, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func Pack_uint32_data(value uint32) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	err := binary.Write(buffer, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func Pack_int64_data(value int64) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	err := binary.Write(buffer, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func Pack_uint64_data(value uint64) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	err := binary.Write(buffer, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}


	return buffer.Bytes()
}

func Pack_float_data(value float32) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	err := binary.Write(buffer, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func Pack_double_data(value float64) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	err := binary.Write(buffer, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func Pack_string_data(value string) []byte {
	//return value

	var buffer = bytes.NewBuffer([]byte{})
	var _, err = buffer.WriteString(value)
	if err != nil {
		panic(err)
	}

	var len_bytes = Pack_int32_data(int32(buffer.Len()))
	var result = append(len_bytes, buffer.Bytes()...)
	return result
}

func Pack_bytes_data(value []byte) []byte {
	var buffer = bytes.NewBuffer([]byte{})
	var err = binary.Write(buffer, binary.LittleEndian, value)
	if err != nil {
		panic(err)
	}

	var len_bytes = Pack_int32_data(int32(buffer.Len()))

	return append(len_bytes, buffer.Bytes()...)
}

func Pack_common(value interface{}) []byte {
	if value == nil {
		var nullArray = Pack_Byte_data(byte(NULL))

		return nullArray
	}

	var code = GetReflectCode(value)
	var codeArray []byte = []byte{
		byte(code),
	}

	var tempArray []byte
	switch code {
	case BOOL:
		tempArray = Pack_bool_data(value.(bool))
	case BYTE:
		tempArray = Pack_Byte_data(value.(byte))
	case INT16:
		tempArray = Pack_int16_data(value.(int16))
	case UINT16:
		tempArray = Pack_uint16_data(value.(uint16))
	case INT32:
		tempArray = Pack_int32_data(value.(int32))
	case UINT32:
		tempArray = Pack_uint32_data(value.(uint32))
	case INT64:
		tempArray = Pack_int64_data(value.(int64))
	case UINT64:
		tempArray = Pack_uint64_data(value.(uint64))
	case SINGLE:
		tempArray = Pack_float_data(value.(float32))
	case DOUBLE:
		tempArray = Pack_double_data(value.(float64))
	case STRING:
		tempArray = Pack_string_data(value.(string))
	case BYTEARRAY:
		tempArray = Pack_bytes_data(value.([]byte))
	case ARRAY:
		var tempValue = reflect.ValueOf(value)

		tempArray = append(tempArray, []byte{
			byte(tempValue.Len()),
		}...)

		for i := 0; i < tempValue.Len(); i++ {
			var itemValue = tempValue.Index(i)
			var item_bytes = Pack_common(itemValue)
			tempArray = append(tempArray, item_bytes...)
		}
	default:
		panic("error"+string(code))
	}

	return append(codeArray, tempArray...)
}

func init() {
	Pack_bool_data(true)
}
