package main

import (
	"redisProject/src/gateway/gatewayRegistService"
)

func main() {
	var s =  gatewayRegistService.GetGateWayRegisterService()
	s.Run("127.0.0.1","28080")
}


