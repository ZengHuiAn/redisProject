using System;
using System.Collections;
using System.Collections.Generic;
using System.Reflection;
using JetBrains.Annotations;
using UnityEngine;

public interface IPackVlue
{
    byte[] GetData();
}
[SerializableStructAttribute]
public struct ClientHeader : IPackVlue
{
    public UInt32 Length, Flag, MessageID, ProtoType;

    public byte[] GetData()
    {
        var this_value = this;
        var t = this_value.GetType();
        List<byte> list = new List<byte>();
        foreach (FieldInfo fieldInfo in t.GetFields())
        {
            UInt32 value = (UInt32) fieldInfo.GetValue(this_value);
            var bs = BitConverter.GetBytes(value);
            list.AddRange(bs);
        }

        return list.ToArray();
    }

    public override string ToString()
    {
        return $"Length:{Length}  Flag:{Flag}  MessageID:{MessageID}  ProtoType:{ProtoType}";
    }
}

[UsedImplicitly]
public class SMessage
{
    public ClientHeader Header;
    public byte[] message ;


    private object content;

    public object Message
    {
        get
        {
            if (content == null)
            {
                if (message !=null && message.Length !=0)
                {
                    this.content = NetUnPackData.unpack_all_lua(this.message);
                }
            }

            return this.content;
        }
    }
    
    public SMessage(ClientHeader header, byte[] message)
    {
        Header = header;
        this.message = message;
    }
}


public static class PackTools
{

    public static int PACK_LENGTH = 16;
    public static byte[] PackObject(object ob)
    {
        var this_value = ob;
        var t = this_value.GetType();

        var have = false;

        foreach (var VARIABLE in t.GetCustomAttributes())
        {
            if (VARIABLE.GetType() == typeof(SerializableStructAttribute))
            {
                have = true;
                break;
            }
        }

        if (!have)
        {
            throw new Exception("VALUE NOT HAVE SerializableStructAttribute HEADER");
        }

        MethodInfo methodInfo = t.GetMethod("GetData");
        if (methodInfo == null)
        {
            return new byte[]{};
        }
        var result = methodInfo?.Invoke(ob,null);
        return  (byte[]) (result);
    }
    
    public static ClientHeader UnPackObject(byte[] bs)
    {
        ClientHeader header = new ClientHeader();
        header.Length = BitConverter.ToUInt32(bs,0);
        header.Flag = BitConverter.ToUInt32(bs,4);
        header.MessageID = BitConverter.ToUInt32(bs,8);
        header.ProtoType = BitConverter.ToUInt32(bs,12);
        return header;
    }
    
    
}
