package gatewayRegistService

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"redisProject/src/CacheServer"
	"redisProject/src/static/res"
)

var cacheMapClient map[string] *CacheServer.MicroserviceAddr

func initMap()  {
	cacheMapClient = make(map[string] *CacheServer.MicroserviceAddr)
}

func RegisterService()  {
	var server = GetGateWayRegisterService()
	server.StaticMethod("POST",res.ServiceRouter, func(c *gin.Context) {
		//var buffer [] byte
		//io.ReadFull(c.Request.Body,buffer)
		body ,_ := ioutil.ReadAll(c.Request.Body)
		fmt.Println("POST",body)
		var addr = CacheServer.MicroserviceAddr{}
	 	err:= json.Unmarshal(body,&addr)
		if err !=nil {
			fmt.Println("err:======>>> json")
			c.JSON(http.StatusFailedDependency, gin.H{
				"message": "垃圾客户端，你数据结构不对",
			})
			return
		}

		if cacheMapClient[addr.Name] != nil {
			fmt.Println("err ---->> 重复注册",addr)

			c.JSON(http.StatusAlreadyReported, gin.H{
				"message": "你已经注册过了",
			})
			return
		}
		cacheMapClient[addr.Name] = &addr
		c.JSON(200, gin.H{
			"message": "对喽",
		})
	})
}

func DeletedService()  {
	var server = GetGateWayRegisterService()
	server.StaticMethod("DET",res.ServiceRouter, func(c *gin.Context) {
		//var buffer [] byte
		//io.ReadFull(c.Request.Body,buffer)
		var result = GetResponeBody(c.Request.Body)

		fmt.Println(result)

		c.JSON(200, gin.H{
			"message": "对喽",
		})
	})
}


func init()  {
	initMap()
	fmt.Println("service register")
	RegisterService()
	DeletedService()
}