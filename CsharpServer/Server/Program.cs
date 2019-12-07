using System;
using Common;
using Grpc.Core;
using Server.UserLogic;

namespace Server
{
    class Program
    {
        static void Main(string[] args)
        {
            
            const int Port = 50051;
            Grpc.Core.Server server = new Grpc.Core.Server
            {
                Services = { Common.User.BindService(new UserImpl()) },
                Ports = { new ServerPort("localhost", Port, ServerCredentials.Insecure) },
            };
            
            
            server.Start();

            Console.WriteLine("Greeter server listening on port " + Port);
            Console.WriteLine("Press any key to stop the server...");
            Console.ReadKey();

            server.ShutdownAsync().Wait();
            
        }
    }
}