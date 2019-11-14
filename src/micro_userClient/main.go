package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
	proto "redisProject/build/proto"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	userClient := proto.NewUserService("User", service.Client())

	// Call the greeter
	rsp, err := userClient.CreateUser(context.TODO(), &proto.CreateUserRequest{UserName:"snake",Passwd:"123"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp)
}
