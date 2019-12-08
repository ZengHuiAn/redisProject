using System;
using System.Collections.Generic;
using Grpc.Core;
using NetLib.CommonUser;

namespace NetLib
{
    public class RPCManager :IDisposable
    {
        private static Channel _channel;
        public bool Connected { get; }

        private ChannelCredentials InsecureInstance = ChannelCredentials.Insecure;
        private RPCManager()
        {
            Connected = false;
        }
        
        
        
        private static RPCManager instance;
        public static RPCManager Instance {
            get
            {
                if (instance == null)
                {
                    instance = new RPCManager();
                }
                return  instance;
            }
        }

        public void Connect(string addr)
        {
            _channel = new Channel(addr,InsecureInstance);
        }

        public  RPC_User.RPC_UserClient NET_RPC_UserClient => new RPC_User.RPC_UserClient(_channel);

        public UserReply Protocol_CreateUser(UserRequest userRequest)
        {
            return NET_RPC_UserClient.CreateUser(userRequest);
        }


        public void Dispose()
        {
            _channel.ShutdownAsync().Wait();
        }
    }
}