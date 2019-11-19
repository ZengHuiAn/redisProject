using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class LogTool
{
    public static string BytesToString(byte[] buffer)
    {
        if (buffer == null)
        {
            return "";
        }

        string str = "";
        for (int i = 0; i < buffer.Length; i++)
        {
            str += (buffer[i] + " ");
        }

        return str;
    }

    public void LogBytes(byte[] buffer)
    {
        Debug.Log(BytesToString(buffer));
    }

    private static LogTool instance;

    public static LogTool Instance
    {
        get
        {
            if (instance == null)
            {
                instance = new LogTool();
            }

            return instance;
        }
    }

    public void ToStringAll(object value)
    {
        Debug.Log(ToStringCommon(value, ""));
    }

    public static string ToStringCommon(object value, string space = "")
    {
        if (value == null)
        {
            var nullArray = new byte[1]
            {
                (byte) EPackType.NULL,
            };
            return "null";
        }

        string str = "";
        var t = value.GetType();
        EPackType code = EPackType.UNDEFINED;
        if (PackType.typeToJumpCode.TryGetValue(t, out code) ||
            PackType.typeToJumpCode.TryGetValue(t.BaseType, out code))
        {
            switch (code)
            {
                case EPackType.BOOL:
                case EPackType.CHAR:
                case EPackType.BYTE:
                case EPackType.INT16:
                case EPackType.UINT16:
                case EPackType.INT32:
                case EPackType.UINT32:
                case EPackType.INT64:
                case EPackType.UINT64:
                case EPackType.SINGLE:
                case EPackType.DOUBLE:
                case EPackType.STRING:

                    str += (space + value + "\n");
                    break;
                case EPackType.BYTEARRAY:
//                    str += (space + BytesToString((byte[])value));
                    str += $"{space}[{BytesToString((byte[]) value)}]\n";
                    break;
                case EPackType.ARRAY:
                    str += (space + "  [\n");
                    var composeData = (Array) value;
                    for (int i = 0; i < composeData.Length; i++)
                    {
                        var itemValue = composeData.GetValue(i);
                        str += $"{space}      [{i.ToString()}] : {ToStringCommon(itemValue, space+"  ")}";
                    }

                    str += (space + "  ]\n");
                    break;
                default:
                    break;
            }
        }

        return str;
    }
}