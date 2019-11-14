package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
	customProto "redisProject/build/proto"
)

//func ProcessEvent(ctx context.Context, event *customProto.) error {
//	fmt.Printf("Got event %+v\n", event)
//	return nil
//}

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("greeter.client"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	userClient := customProto.NewUserService("UserService", service.Client())

	p := micro.NewPublisher("events", service.Client())
	ctx :=context.TODO()

	p.Publish(ctx,&customProto.Event{
		Id:"11",
	})
	// Call the greeter
	rsp, err := userClient.CreateUser(context.TODO(), &customProto.CreateUserRequest{UserName: "snake",Passwd:"123"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp)
}
