package main

import (
	"fmt"
	"redisProject/src/CacheServer"
	"redisProject/src/gateway/gatewayRegistService"
)



func main() {

	var selfAddr = CacheServer.MicroserviceAddr{
		Host:"127.0.0.1",
		Port:"8080",
		Name:"test",
	}
	var resp = gatewayRegistService.RegisterGateWayServer(selfAddr)

	fmt.Println(resp.Body)
}