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
		c.JSON(200, gin.H{
			"message": "对喽",
		})
	})
}


func init()  {
	fmt.Println("service register")
	RegisterService()
}