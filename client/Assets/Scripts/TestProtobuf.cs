using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Grpc.Core;
using Google.Protobuf;
using  System;
using NetLib;
using NetLib.CommonUser;

public class TestProtobuf : MonoBehaviour
{
    
    // Start is called before the first frame update
    void Start()
    {
        RPCManager.Instance.Connect("127.0.0.1:50051");
        var request = new UserRequest() {Name = "123456"};
        var reply =  RPCManager.Instance.Protocol_CreateUser(request);
        
        Debug.Log(reply);


        
//        var client = new greeter
//        GreeterServer
    }

    // Update is called once per frame
    void Update()
    {
        
    }

    private void OnDestroy()
    {
        RPCManager.Instance.Dispose();
    }
}
