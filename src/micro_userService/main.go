package main

import (
	"fmt"
	micro "github.com/micro/go-micro"
	User "redisProject/build/proto"

	"redisProject/src/business"
)

func main() {
	service := micro.NewService(
		micro.Name("UserService"),
	)

	service.Init()

	User.RegisterUserServiceHandler(service.Server(), new(business.User))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
