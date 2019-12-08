using System;
using Grpc.Core;
using NetLib.CommonUser;
namespace NetClient
{
    class Program
    {
        static void Main(string[] args)
        {
            Channel channel = new Channel("127.0.0.1:50051", ChannelCredentials.Insecure);

            var client = new RPC_User.RPC_UserClient(channel);
            String user = "测试一下";

            var reply = client.CreateUser(new UserRequest() { Name = user,Passwd = "123456"});
            Console.WriteLine("Greeting: " + reply.ErrorCode + reply.Uuid);

            channel.ShutdownAsync().Wait();
            Console.WriteLine("Press any key to exit...");
            Console.ReadKey();
        }
    }
}