using System;
using System.Collections;
using System.Collections.Generic;
using System.Text;
using UnityEngine;

public class NetUnPackData
{
    public static bool unpack_bool_data(ref byte[] bytes)
    {
        var result = BitConverter.ToBoolean(bytes, 0);

        return result;
    }

    public static char unpack_char_data(byte[] bytes)
    {
        var result = BitConverter.ToChar(bytes, 0);

        return result;
    }


    public static byte unpack_Byte_data(byte[] bytes)
    {
        return bytes[0];
    }

    public static Int16 unpack_int16_data(byte[] bytes)
    {
        var result = BitConverter.ToInt16(bytes, 0);

        return result;
    }


    public static UInt16 unpack_uint16_data(byte[] bytes)
    {
        var result = BitConverter.ToUInt16(bytes, 0);

        return result;
    }


    public static Int32 unpack_int32_data(byte[] bytes)
    {
        var result = BitConverter.ToInt32(bytes, 0);

        return result;
    }

    public static UInt32 unpack_uint32_data(byte[] bytes)
    {
        var result = BitConverter.ToUInt32(bytes, 0);

        return result;
    }

    public static Int64 unpack_int64_data(byte[] bytes)
    {
        var result = BitConverter.ToInt64(bytes, 0);

        return result;
    }

    public static UInt64 unpack_uint64_data(byte[] bytes)
    {
        var result = BitConverter.ToUInt64(bytes, 0);

        return result;
    }

    public static float unpack_float_data(byte[] bytes)
    {
        var result = BitConverter.ToSingle(bytes, 0);

        return result;
    }


    public static double unpack_double_data(byte[] bytes)
    {
        var result = BitConverter.ToDouble(bytes, 0);

        return result;
    }

    public static string unpack_string_data(byte[] bytes)
    {
        var intValue = unpack_int32_data(bytes);
        // 解析字符串
        var resultBytes = splice_Bytes(bytes, 4, intValue + 4);

        var result = Encoding.UTF8.GetString(resultBytes);
        return result;
    }

    public static byte[] splice_Bytes(byte[] bytes, int startIndex, int endIndex)
    {
        List<byte> result = new List<byte>();
        for (int i = startIndex; i < endIndex; i++)
        {
            result.Add(bytes[i]);
        }

        return result.ToArray();
    }


    public static byte[] unpack_bytes_data(byte[] bytes)
    {
        var intValue = unpack_int32_data(bytes);
        // 解析字符串
        var resultBytes = splice_Bytes(bytes, 4, intValue + 4);
        return resultBytes;
    }

    public static object unpack_null_data(byte[] bytes)
    {
        return null;
    }

    public static object unpack_common(byte[] bytes)
    {
        var code = (EPackType)bytes[0];
        object result;
        switch (code)
            {
                case EPackType.BOOL:
                    result = unpack_bool_data();
                    break;
                case EPackType.CHAR:
                    tempArray = pack_char_data((char) value);
                    break;
                case EPackType.BYTE:
                    tempArray = pack_Byte_data((byte) value);
                    break;
                case EPackType.INT16:
                    tempArray = pack_int16_data((Int16) value);
                    break;
                case EPackType.UINT16:
                    tempArray = pack_uint16_data((UInt16) value);
                    break;
                case EPackType.INT32:
                    tempArray = pack_int32_data((Int32) value);
                    break;
                case EPackType.UINT32:
                    tempArray = pack_uint32_data((UInt32) value);
                    break;
                case EPackType.INT64:
                    tempArray = pack_int64_data((Int64) value);
                    break;
                case EPackType.UINT64:
                    tempArray = pack_uint64_data((UInt64) value);
                    break;
                case EPackType.SINGLE:
                    tempArray = pack_float_data((Single) value);
                    break;
                case EPackType.DOUBLE:
                    tempArray = pack_double_data((Double) value);
                    break;
                case EPackType.STRING:
                    tempArray = pack_string_data((string) value);
                    break;
                case EPackType.BYTEARRAY:
                    tempArray = pack_bytes_data((byte[]) value);
                    break;
                case EPackType.ARRAY:
                    var composeData = (Array)value;
                    
                    for (int i = 0; i < composeData.Length; i++)
                    {
                        var itemValue = composeData.GetValue(i);
                        var item_bytes = pack_common(itemValue);

                        tempArray = copyBytesArray(tempArray, item_bytes);
                    }
                    break;
                default:
                    throw new InvalidOperationException("Not supported primitive object resolver. type:" + t.Name);
            }
            
            
            codeArray = copyBytesArray(codeArray, tempArray);

            if (code == EPackType.UNDEFINED)
            {
                return null;
            }
            return codeArray;
        }
}