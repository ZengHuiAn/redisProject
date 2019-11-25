using System;
using MessagePack;
using System.Collections;
using System.Collections.Generic;
using System.Text;
using UnityEngine;


public class TestMSGPack : MonoBehaviour
{
    // Start is called before the first frame update
    void Start()
    {
//        var obj = new object[]
//        {
//            new char[]
//            {
//                'a',
//                '增',
//                '怕',
//                '%'
//            },
//        };
        LogTool.Instance.ToStringAll(BitConverter.GetBytes((Int16) (50)));
        ClientHeader header = new ClientHeader();
        SerializableStructAttribute[] objects =
            (SerializableStructAttribute[]) header.GetType()
                .GetCustomAttributes(typeof(SerializableStructAttribute), false);

        foreach (var VARIABLE in objects)
        {
            Debug.Log(VARIABLE);
        }

//
//        var bs = NetPackData.pack_all(obj);
////        
////        LogTool.Instance.ToStringAll((byte)('a'));
////        LogTool.Instance.ToStringAll(BitConverter.GetBytes(false));
////        LogTool.Instance.ToStringAll(bs);
//
//        var obj_oo = NetUnPackData.unpack_all(bs); 
//        LogTool.Instance.ToStringAll(obj_oo);
//        .Instance.LogBytes(bs);
    }

    public void testEvent()
    {
        EventManager.Instance.Call("OnLogin",null);
    }


}

