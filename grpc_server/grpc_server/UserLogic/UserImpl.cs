﻿using System;
using System.Threading.Tasks;
using NetLib.CommonUser;
using Grpc.Core;

namespace Server.UserLogic
{
    public class UserImpl : RPC_User.RPC_UserBase
    {
        
        public override Task<UserReply> CreateUser(UserRequest request, ServerCallContext context)
        {
            Console.WriteLine($"{request.Name}  {request.Passwd}");
            return Task.FromResult(new UserReply() { ErrorCode = "200",Uuid = request.Name});
        }
        public  Task<UserReply> SayHello(UserRequest request, ServerCallContext context)
        {
            Console.WriteLine($"收到信息 ：{request.Name}");
//            return Task.FromResult(new UserReply() {Message = $"你好,{request.Name}"});
            return base.CreateUser(request, context);
        }
    }
}