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
            6.6d,
        };
//
        var bs = NetPackData.pack_all(obj);


//        .Instance.LogBytes(bs);
    }
    
}