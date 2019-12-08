package CacheServer

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type HttpListenServer struct {

	started bool
	routerEngine *gin.Engine
}

func  CreateHttpListenServer() *HttpListenServer {
	server := &HttpListenServer{}
	server.routerEngine = gin.Default()

	println("create success HttpListenServer ... ")
	return server
}
/*
静态注册
*/
func (server *HttpListenServer) StaticMethod(methodName string, router string ,handlers ...gin.HandlerFunc ) {
	method := strings.ToUpper(methodName)

	switch method {
	case "GET":
		server.routerEngine.GET(router,handlers...)
	case "POST":
		server.routerEngine.POST(router,handlers...)
	case "DET":
		server.routerEngine.DELETE(router,handlers...)
	}
}

func (server *HttpListenServer) Run(host string, port string) {
	server.StaticMethod("GET","/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	addr := host + ":" + port
	println( "HttpListenServer  running ",addr)
	_ = server.routerEngine.Run(addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")


}
