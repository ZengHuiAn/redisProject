package business

import (
	"context"
	"fmt"
	"redisProject/src/tools"

	//"fmt"
	proUser "redisProject/build/proto"
	//"redisProject/src/tools"
)

type User struct{}


func (user *User) CreateUser(ctx context.Context, req *proUser.CreateUserRequest, rsp *proUser.CreateUserResponse) error {
	result, err := tools.GetRedisClient().Get("userNames").Result()
	fmt.Println("查询redis:\t", result, err, req)
	rsp.ErrorCode = 200
	rsp.UserMessage = &proUser.UserData{UserName: req.UserName, UUID: 1, ChatID: 0};

	return nil
}