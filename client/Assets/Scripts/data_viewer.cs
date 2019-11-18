using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class data_viewer 
{

    public static void write_ushort_len(byte[] buffer, int offest, int value)
    {
        byte[] bytes = BitConverter.GetBytes(value);

        if (!BitConverter.IsLittleEndian)
        {
            Array.Reverse(bytes);
        }
        
        Array.Copy(bytes,0,buffer,offest,bytes.Length);
    }

    public static void write_uint_len(byte[] buffer, int offest, uint value)
    {
        byte[] bytes = BitConverter.GetBytes(value);

        if (!BitConverter.IsLittleEndian)
        {
            Array.Reverse(bytes);
        }
        
        Array.Copy(bytes,0,buffer,offest,bytes.Length);
    }
    
    public static void write_bytes_len(byte[] buffer, int offest, byte[] value)
    {
        Array.Copy(value,0,buffer,offest,value.Length);
    }
}
