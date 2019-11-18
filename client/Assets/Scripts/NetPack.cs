using System;
using System.Collections;
using System.Collections.Generic;
using System.Reflection;
using UnityEngine;

public static class NetPack
{
    //
}


public enum EPackType
{
    UNDEFINED = 1, // 未定义
    BOOL = 101,
    CHAR = 102,
    BYTE = 103,
    INT16 = 104,
    UINT16 = 105,
    INT32 = 106,
    UINT32 = 107,
    INT64 = 108,
    UINT64 = 109,
    SINGLE = 110,
    DOUBLE = 111,
    STRING = 112,
    BYTEARRAY = 113,
    ARRAY = 114,
}

public static class PackType
{
    public static readonly Dictionary<Type, EPackType> typeToJumpCode = new Dictionary<Type, EPackType>()
    {
        {typeof(Boolean), EPackType.BOOL},
        {typeof(Char), EPackType.CHAR},
        {typeof(Byte), EPackType.BYTE},
        {typeof(Int16), EPackType.INT16},
        {typeof(UInt16), EPackType.UINT16},
        {typeof(Int32), EPackType.INT32},
        {typeof(UInt32), EPackType.UINT32},
        {typeof(Int64), EPackType.INT64},
        {typeof(UInt64), EPackType.UINT64},
        {typeof(Single), EPackType.SINGLE},
        {typeof(Double), EPackType.DOUBLE},
        {typeof(string), EPackType.STRING},
        {typeof(byte[]), EPackType.BYTEARRAY},
        {typeof(Array), EPackType.ARRAY},
    };


    public static bool IsSupportedType(Type type, TypeInfo typeInfo, object value)
    {
        if (value == null) return true;
        if (PackType.typeToJumpCode.ContainsKey(type)) return true;
        return false;
    }
}