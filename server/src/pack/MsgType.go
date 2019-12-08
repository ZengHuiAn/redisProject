package pack

import (
	"fmt"
	"reflect"
)

type EPackType byte

const (
	UNDEFINED EPackType = 0
	NULL      EPackType = 100
	BOOL      EPackType = 101
	CHAR      EPackType = 102
	BYTE      EPackType = 103
	INT16     EPackType = 104
	UINT16    EPackType = 105
	INT32     EPackType = 106
	UINT32    EPackType = 107
	INT64     EPackType = 108
	UINT64    EPackType = 109
	SINGLE    EPackType = 110
	DOUBLE    EPackType = 111
	STRING    EPackType = 112
	BYTEARRAY           = 113
	ARRAY               = 114
)

func GetReflectCode(p interface{}) EPackType {
	var code = GetCodeAtCode(reflect.ValueOf(p))
	if code == UNDEFINED {
		panic(code)
	}

	return code
}

func GetCodeAtCode(p reflect.Value) EPackType {

	fmt.Println(p.Kind(),reflect.TypeOf(p.Kind()))
	var kind = p.Kind()
	switch kind {
	case reflect.Bool:
		return BOOL
	case reflect.Int8:
		return BYTE
	case reflect.Int16:
		return INT16
	case reflect.Uint16:
		return UINT16
	case reflect.Int32:
		return INT32
	case reflect.Int:
		return INT32
	case reflect.Uint32:
		return UINT32
	case reflect.Uint:
		return UINT32
	case reflect.Int64:
		return INT64
	case reflect.Uint64:
		return UINT64
	case reflect.Float32:
		return SINGLE
	case reflect.Float64:
		return  DOUBLE
	case reflect.String:
		return STRING
	default:
		if kind == reflect.Slice || kind == reflect.Array {
			switch p.Interface().(type) {
			case []byte:
				return BYTEARRAY
			}

			return ARRAY
		}
		return UNDEFINED
	}
}
