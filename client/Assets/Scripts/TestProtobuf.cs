using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Grpc.Core;
using Google.Protobuf;
using  System;
using NetLib;

public class TestProtobuf : MonoBehaviour
{
    private Channel channel;
    // Start is called before the first frame update
    void Start()
    {
         channel = new Channel("127.0.0.1:50051", ChannelCredentials.Insecure);
        var client = new User.UserClient(channel);

        var reply = client.SayHello(new UserRequest() {Name = "123456"});
        
        Debug.Log("Greeting: " + reply.Message);


        
//        var client = new greeter
//        GreeterServer
    }

    // Update is called once per frame
    void Update()
    {
        
    }

    private void OnDestroy()
    {
        channel.ShutdownAsync().Wait();
    }
}
