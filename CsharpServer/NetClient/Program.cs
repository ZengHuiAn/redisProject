using System;
using Grpc.Core;
using NetLib;
namespace NetClient
{
    class Program
    {
        static void Main(string[] args)
        {
            Channel channel = new Channel("127.0.0.1:50051", ChannelCredentials.Insecure);

            var client = new User.UserClient(channel);
            String user = "测试一下";

            var reply = client.SayHello(new UserRequest() { Name = user });
            Console.WriteLine("Greeting: " + reply.Message);

            channel.ShutdownAsync().Wait();
            Console.WriteLine("Press any key to exit...");
            Console.ReadKey();
        }
    }
}