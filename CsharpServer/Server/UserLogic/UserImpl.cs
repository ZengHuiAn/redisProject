using System;
using System.Threading.Tasks;
using Common;
using Grpc.Core;

namespace Server.UserLogic
{
    public class UserImpl : Common.User.UserBase
    {
        public override Task<UserReply> SayHello(UserRequest request, ServerCallContext context)
        {
            Console.WriteLine($"收到信息 ：{request.Name}");
            return Task.FromResult(new UserReply() {Message = $"你好,{request.Name}"});
        }
    }
}