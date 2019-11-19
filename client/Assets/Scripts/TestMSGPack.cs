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
        var obj = new object[]
        {
            "1234",
            9999,
            5.1f,
            Encoding.UTF8.GetBytes("123546"),
            
            new object[]
            {
                "1234",
                9999,
                5.1f,
            },
        };
//
        var bs = NetPackData.pack_all(obj);
        
        
        
        LogTool.Instance.ToStringAll(bs);
        
        var obj_oo = NetUnPackData.unpack_all(bs);
        LogTool.Instance.ToStringAll(obj_oo);
//        .Instance.LogBytes(bs);
    }
    
}