package gatewayRegistService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"redisProject/src/CacheServer"
	"redisProject/src/static/res"
)

// 客户端注册接口

func GetResponeBody(r io.Reader) map[string]interface{}  {
	var mesgger  map[string]interface{}
	var bf, _= ioutil.ReadAll(r)
	json.Unmarshal(bf,&mesgger)

	return  mesgger
}

type RegisterRespone struct {
	Code int
	Body map[string]interface{}
}

func RegisterGateWayServer( addr CacheServer.MicroserviceAddr) RegisterRespone  {
	var sendVariable ,_ = json.Marshal(addr)
	var body = bytes.NewBuffer(sendVariable)
	resp, err := http.Post(res.BaseRouterURL+res.ServiceRouter, "application/json", body)
	if err != nil {
		fmt.Println(err)
	}
	var message map[string]interface{} = GetResponeBody(resp.Body)
	fmt.Println(message["message"])
	if resp.StatusCode == http.StatusOK {
		fmt.Println("注册成功---------》》》")
	}
	return RegisterRespone{resp.StatusCode, message}
}

func DeletGateWayServer(name string) RegisterRespone  {

	//resp, err := http.(res.BaseRouterURL+res.ServiceRouter, "application/json", body)
	return RegisterRespone{}
}