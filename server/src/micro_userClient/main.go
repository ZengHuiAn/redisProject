package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	customProto "redisProject/build/proto"
)

//func ProcessEvent(ctx context.Context, event *customProto.) error {
//	fmt.Printf("Got event %+v\n", event)
//	return nil
//}

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("[wrapper] client request to service: %s endpoint: %s\n", req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

// implements client.Wrapper as logWrapper
func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

func main() {
	// Create a new service
	service := micro.NewService(
		micro.Name("greeter.client"),
		micro.WrapClient(logWrap),
		)
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	userClient := customProto.NewUserService("UserService", service.Client())

	p := micro.NewPublisher("events", service.Client())
	ctx := context.TODO()

	p.Publish(ctx, &customProto.Event{
		Id: "11",
	})
	// Call the greeter
	rsp, err := userClient.CreateUser(context.TODO(), &customProto.CreateUserRequest{UserName: "snake", Passwd: "123"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp)
}
