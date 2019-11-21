using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System.Reflection;
using System.Text;
using UnityEngine;

public class NetPackData
{
    public static byte[] pack_bool_data(bool value)
    {
        byte[] int_bytes = BitConverter.GetBytes(value);
        
        return int_bytes;
    }    

    public static byte[] pack_char_data(char value)
    {
        byte[] int_bytes = BitConverter.GetBytes(Convert.ToInt32(value));

        return int_bytes;
    }    
    

    public static byte[] pack_Byte_data(byte value)
    {
        byte[] int_bytes = BitConverter.GetBytes(value);
        
        return int_bytes;
    }    
    
    public static byte[] pack_int16_data(Int16 value)
    {
        byte[] int_bytes = BitConverter.GetBytes(value);
        
        return int_bytes;
    }
    
    
    
    public static byte[] pack_uint16_data(UInt16 value)
    {
        byte[] int_bytes = BitConverter.GetBytes(value);
        
        return int_bytes;
    }



    public static byte[] pack_int32_data(Int32 value)
    {
        byte[] int_bytes = BitConverter.GetBytes(value);
        
        return int_bytes;
    }
    
    public static byte[] pack_uint32_data(UInt32 value)
    {
        byte[] int_bytes = BitConverter.GetBytes(value);
        
        return int_bytes;
    }
    
    public static byte[] pack_int64_data(Int64 value)
    {
        byte[] int_bytes = BitConverter.GetBytes(value);
        
        return int_bytes;
    }
    
    public static byte[] pack_uint64_data(UInt64 value)
    {
        byte[] int_bytes = BitConverter.GetBytes(value);
        
        return int_bytes;
    }
    
    public static byte[] pack_float_data(float value)
    {
        byte[] int_bytes = BitConverter.GetBytes(value);
        
        return int_bytes;
    }
    
    
    public static byte[] pack_double_data(double value)
    {
        byte[] int_bytes = BitConverter.GetBytes(value);
        
        return int_bytes;
    }
    
    public static byte[] pack_string_data(string value)
    {
        byte[] int_bytes = Encoding.UTF8.GetBytes(value);

        var len_bytes = pack_int32_data(int_bytes.Length);
        
        len_bytes = copyBytesArray(len_bytes, int_bytes);
        
        return len_bytes;
    }
    
    public static byte[] pack_bytes_data(byte[] value)
    {
        byte[] int_bytes = value;
        var len_bytes = pack_int32_data(int_bytes.Length);
        len_bytes = copyBytesArray(len_bytes, int_bytes);
        return len_bytes;
    }
    
    public static byte[] pack_null_data(byte[] value)
    {
        byte[] int_bytes = new byte[1]
        {
            1
        };
        
        return int_bytes;
    }


    public static byte[] pack_all(object value)
    {
        return pack_common(value);
    }


    public static byte[] copyBytesArray(byte[] bs1 ,byte[] bs2)
    {
        List<byte> source = new List<byte>(){};
        source.AddRange(bs1);
        source.AddRange(bs2);

        return source.ToArray();
    }

    public static byte[] pack_common(object value)
    {
        
        if (value == null)
        {
            var nullArray = new byte[1]
            {
                (byte) EPackType.NULL,
            };
            return nullArray;
        }
        var t = value.GetType();
        EPackType code = EPackType.UNDEFINED;
        if (PackType.typeToJumpCode.TryGetValue(t, out code)||PackType.typeToJumpCode.TryGetValue(t.BaseType, out code))
        {
            var codeArray = new byte[1]
            {
                (byte) code,
            };
            byte[] tempArray =  new byte[]{};
            switch (code)
            {
                case EPackType.BOOL:
                    tempArray = pack_bool_data((bool) value);
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
                    tempArray = copyBytesArray(tempArray, new byte[1]
                    {
                        (byte) composeData.Length,
                    });
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

        return null;
    }


}
