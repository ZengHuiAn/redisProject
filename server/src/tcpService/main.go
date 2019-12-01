package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"redisProject/src/eventManager"
	"redisProject/src/net_struct"
	"redisProject/src/pack"
	"redisProject/src/static/res"
	"redisProject/src/tcpService/client"
	"strconv"
)

var s_instance = net_struct.ServerInstance{
	Config: &net_struct.TCPStruct{
		Name:         "ConnectServer",
		Port:         10000,
		ProtocolType: "tcp",
		Host:         "127.0.0.1",
	},
}

func main() {
	fmt.Println(s_instance.Config.GetHost() + ":" + strconv.Itoa(s_instance.Config.Port))
	var listener, err = net.Listen(s_instance.Config.ProtocolType, s_instance.Config.GetHost()+":"+strconv.Itoa(s_instance.Config.Port))

	if err != nil {
		log.Fatal("lister error ", err)
		return
	}


	defer listener.Close()

	

	eventManager.GetEventManagerForName(res.EVENTMGR_PROTOCOL_Name).AddProtoEventAction(res.PROTOCOL_C2S,101,&eventManager.Event{Action: func(args interface{}) {
		fmt.Println("收到客户端数据----->>>",pack.Encode(args.(*net_struct.TCPClientData).GetBody()))
	}})
	connectManager := client.GetManagerForName(res.CONNECT_MGR_Name)
	connectManager.RegisterMiddle(func(msgName string, msgID uint32, data *net_struct.TCPClientData) (b bool, e error) {
		return false, errors.New("测试错误------>>>")
	})
	go connectManager.Run()
	fmt.Println(fmt.Sprintf("ProtocolType %s, addr %s", listener.Addr().Network(), listener.Addr().String()))
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error ", err)
			os.Exit(1)
		}

		fmt.Println(fmt.Sprintf("message %s -> %s", conn.RemoteAddr(), conn.LocalAddr()))
		client := client.NewCustomClient(connectManager.Name(),conn)
		connectManager.RegisterClient(client)
	}

}



