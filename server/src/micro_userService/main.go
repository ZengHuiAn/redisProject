package main

import (
"context"
"fmt"
micro "github.com/micro/go-micro"
customProto "redisProject/build/proto"
"redisProject/src/business"
)

func ProcessEvent(ctx context.Context, event *customProto.Event) error {
	fmt.Printf("Got event %+v\n", event)
	return nil
}

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("[%v] server request: %s\n", time.Now(), req.Endpoint())
		return fn(ctx, req, rsp)
	}
}

func main() {
	service := micro.NewService(
		micro.Name("UserService"),
		micro.WrapHandler(logWrapper),
	)

	service.Init()

	customProto.RegisterUserServiceHandler(service.Server(), new(business.User))

	micro.RegisterSubscriber("events",service.Server(), ProcessEvent)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
