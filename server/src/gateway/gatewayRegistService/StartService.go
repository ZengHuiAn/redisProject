package gatewayRegistService

import (
	"fmt"
	"redisProject/src/CacheServer"
)

var instance *CacheServer.HttpListenServer

func GetInstance()   {

}

func Start()  {
	if instance != nil {
		return
	}

	instance = CacheServer.CreateHttpListenServer()
}

func GetGateWayRegisterService() *CacheServer.HttpListenServer {

	if instance == nil {
		instance = CacheServer.CreateHttpListenServer()
	}

	return  instance
}

func init()  {
	fmt.Println("start service")
	Start()
}