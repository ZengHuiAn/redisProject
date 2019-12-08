
using Grpc.Core;

namespace NetLib
{
    public class meta
    {
        public static User.UserClient GetUserClient(Channel channel,UserRequest userRequest )
        {
            return new User.UserClient(channel);
        }
    }
}