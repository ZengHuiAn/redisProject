using System;
using System.Collections;
using System.Collections.Generic;
using System.Text;
using UnityEngine;

public class NetUnPackData
{
    public static object unpack_bool_data(ref byte[] bytes)
    {
        var result = BitConverter.ToBoolean(bytes, 0);
        bytes = splice_Bytes(bytes, 1, bytes.Length);
        return result;
    }

    public static object unpack_char_data(ref byte[] bytes)
    {
        var result = unpack_int32_data(ref bytes);
        return Convert.ToChar(result);
    }


    public static object unpack_Byte_data(ref byte[] bytes)
    {
        var result = bytes[0];
        bytes = splice_Bytes(bytes, 1, bytes.Length);
        return result;
    }

    public static object unpack_int16_data(ref byte[] bytes)
    {
        var result = BitConverter.ToInt16(bytes, 0);
        bytes = splice_Bytes(bytes, 2, bytes.Length);
        return result;
    }


    public static object unpack_uint16_data(ref byte[] bytes)
    {
        var result = BitConverter.ToUInt16(bytes, 0);
        bytes = splice_Bytes(bytes, 2, bytes.Length);
        return result;
    }


    public static object unpack_int32_data(ref byte[] bytes)
    {
        var result = BitConverter.ToInt32(bytes, 0);
        bytes = splice_Bytes(bytes, 4, bytes.Length);
        return result;
    }

    public static object unpack_uint32_data(ref byte[] bytes)
    {
        var result = BitConverter.ToUInt32(bytes, 0);
        bytes = splice_Bytes(bytes, 4, bytes.Length);
        return result;
    }

    public static object unpack_int64_data(ref byte[] bytes)
    {
        
        var result = BitConverter.ToInt64(bytes, 0);
        bytes = splice_Bytes(bytes, 8, bytes.Length);
        return result;
    }

    public static object unpack_uint64_data(ref byte[] bytes)
    {
        var result = BitConverter.ToUInt64(bytes, 0);
        bytes = splice_Bytes(bytes, 8, bytes.Length);
        return result;
    }

    public static object unpack_float_data(ref byte[] bytes)
    {
        var result = BitConverter.ToSingle(bytes, 0);
        bytes = splice_Bytes(bytes, 4, bytes.Length);
        return result;
    }


    public static object unpack_double_data(ref byte[] bytes)
    {
        var result = BitConverter.ToDouble(bytes, 0);
        bytes = splice_Bytes(bytes, 8, bytes.Length);
        return result;
    }

    public static object unpack_string_data(ref byte[] bytes)
    {
        var intValue = (Int32) unpack_int32_data(ref bytes);
        // 解析字符串 // 剪切求得字符串的长度
        var resultBytes = splice_Bytes(bytes, 0, intValue);

        var result = Encoding.UTF8.GetString(resultBytes);
        bytes = splice_Bytes(bytes, intValue, bytes.Length);

        return result;
    }

    public static object unpack_bytes_data(ref byte[] bytes)
    {
        var intValue = (Int32) unpack_int32_data(ref bytes);
        // 解析字符串
        var resultBytes = splice_Bytes(bytes, 0, intValue);

        bytes = splice_Bytes(bytes, intValue, bytes.Length);

        return resultBytes;
    }

    public static object unpack_null_data(ref byte[] bytes)
    {
        bytes = splice_Bytes(bytes, 1, bytes.Length);

        return null;
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

    public static object unpack_all(byte[] bytes)
    {
        return unpack_common(ref bytes);
    }

    public static object unpack_common(ref byte[] bytes)
    {
        byte pack_type =  (byte)unpack_Byte_data(ref bytes);

        var code = (EPackType) pack_type;
//        var code = (EPackType) unpack_Byte_data(ref bytes);
        object result;
        switch (code)
        {
            case EPackType.BOOL:
                result = unpack_bool_data(ref bytes);
                break;
            case EPackType.CHAR:
                result = unpack_char_data(ref bytes);
                break;
            case EPackType.BYTE:
                result = unpack_Byte_data(ref bytes);
                break;
            case EPackType.INT16:
                result = unpack_int16_data(ref bytes);
                break;
            case EPackType.UINT16:
                result = unpack_uint16_data(ref bytes);
                break;
            case EPackType.INT32:
                result = unpack_int32_data(ref bytes);
                break;
            case EPackType.UINT32:
                result = unpack_uint32_data(ref bytes);
                break;
            case EPackType.INT64:
                result = unpack_int64_data(ref bytes);
                break;
            case EPackType.UINT64:
                result = unpack_uint64_data(ref bytes);
                break;
            case EPackType.SINGLE:
                result = unpack_float_data(ref bytes);
                break;
            case EPackType.DOUBLE:
                result = unpack_double_data(ref bytes);
                break;
            case EPackType.STRING:
                result = unpack_string_data(ref bytes);
                break;
            case EPackType.BYTEARRAY:
                result = unpack_bytes_data(ref bytes);
                break;
            case EPackType.ARRAY:
                // 解的数组的长度
                byte by =  (byte)unpack_Byte_data(ref bytes);

                var arrayLen = (Int32) by;
                
                var composeData = new object[arrayLen];


                for (int i = 0; i < arrayLen; i++)
                {
                    var itemValue = unpack_common(ref bytes);
                    composeData[i] = itemValue;
                }

                result = composeData;
                break;
            default:
                throw new InvalidOperationException("Not supported primitive object resolver. type: " + code);
        }

        return result;
    }
}