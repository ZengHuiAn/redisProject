package client

import (
	"fmt"
	"redisProject/src/common"
	"redisProject/src/eventManager"
	"redisProject/src/net_struct"
	"redisProject/src/static/res"
	"sync"
)

// 客户端链接--->>>
type ConnectManager struct {
	clients map[string]*CustomClient

	wg         sync.WaitGroup
	groupCount int
	//连接chanel
	connectChan chan *CustomClient
	// 断开连接chanel
	disConnectChan chan string
	name           string
	middleWareFunc func(msgName string, ip string, msgID uint32, data *net_struct.TCPClientData) (bool, error)
	onMsgChan      chan common.Message

	onLocalMsgChan chan common.Message
}

func defaultMiddleFunc(msgName string, msgID uint32, data *net_struct.TCPClientData) (bool, error) {
	return false, nil
}

//创建一个连接器
func MakeConnectManager(name string) *ConnectManager {
	manager := &ConnectManager{name: name}
	manager.initSetting()
	return manager
}

func (manager *ConnectManager) RegisterMiddle(middle func(msgName string, ip string, msgID uint32, data *net_struct.TCPClientData) (bool, error)) {
	manager.middleWareFunc = middle
}

func (manager *ConnectManager) Name() string {
	return manager.name
}

//当客户端异常关闭
func (manager *ConnectManager) OnCloseClient(name interface{}) {
	fmt.Println("客户端关闭?", name)
	manager.CloseClient1(name.(string))
}

//初始化设置
func (manager *ConnectManager) initSetting() {
	manager.clients = make(map[string]*CustomClient)
	manager.groupCount = 0
	manager.connectChan = make(chan *CustomClient)
	manager.disConnectChan = make(chan string)
	manager.onMsgChan = make(chan common.Message)
	manager.onLocalMsgChan = make(chan common.Message)
	//manager.wg = make(sync.WaitGroup)
	eventManager.GetEventManagerForName(res.EVENTMGR_CONNECT_Name).
		AddEventAction(res.CONTENT_NAME_EVENT_CLIENT_READ_ERROR,
			&eventManager.Event{Action: manager.OnCloseClient})
}

// 断开连接
func (manager *ConnectManager) CloseClient(names []string) {
	//
	for i := 0; i < len(names); i++ {
		manager.CloseClient1(names[i])
	}
}

// 关闭一个客户端
func (manager *ConnectManager) CloseClient1(names string) {
	manager.disConnectChan <- names
}

//运行连接携程
func (manager *ConnectManager) Run() {
	defer manager.closeAll()
	manager.wg.Add(3)
	go manager.OnConnAndClose()
	go manager.OnReadMsg()
	go manager.OnLoadReadMsg()
	manager.wg.Wait()
}

//关闭所有连接的客户端
func (manager *ConnectManager) closeAll() {
	var onCloses []string
	for k, _ := range manager.clients {
		onCloses = append(onCloses, k)
	}
}

// 注册一个连接的客户端
func (manager *ConnectManager) RegisterClient(itemClient *CustomClient) {
	manager.connectChan <- itemClient
}

// 运行客户端读写携程
func (manager *ConnectManager) addItemClient(itemClient *CustomClient) {
	name := GetRemoteAddr(*itemClient)
	manager.clients[name] = itemClient
	manager.clients[name].OnConnect()
}

func (manager *ConnectManager) ReceiveLocalMsgMany(clients []string, data net_struct.TCPClientData) {
	for _, v := range clients {
		manager.ReceiveLocalMsgDone(v, data)
	}
}

//接收本地发送方法---------->>>>
func (manager *ConnectManager) ReceiveLocalMsgDone(client string, data net_struct.TCPClientData) {
	fmt.Println("ReceiveLocalMsgDone", client)
	manager.onLocalMsgChan <- common.MakeMessage(client, data)
	fmt.Println("ReceiveLocalMsgDone success", client)
}

//读取客户端数据
func (manager *ConnectManager) receiveMsg(ip string, data net_struct.TCPClientData) {
	manager.onMsgChan <- common.MakeMessage(ip, data)
}

//读取管道数据
func (manager *ConnectManager) OnReadMsg() {
	defer manager.wg.Done()
	for {
		select {
		case value, err := <-manager.onMsgChan:
			fmt.Println(err)
			client := manager.clients[value.ClientAddr()]
			if client != nil {
				if manager.middleWareFunc == nil {
					eventManager.GetEventManagerForName(res.EVENTMGR_PROTOCOL_Name).
						CallProto(res.PROTOCOL_C2S, value.ClientAddr(), value.Data().Header.MessageID, value.Data())
				} else {
					_, err := manager.middleWareFunc(res.PROTOCOL_C2S, value.ClientAddr(), value.Data().Header.MessageID, value.Data())
					if err == nil {
						eventManager.GetEventManagerForName(res.EVENTMGR_PROTOCOL_Name).
							CallProto(res.PROTOCOL_C2S, value.ClientAddr(), value.Data().Header.MessageID, value.Data())
					} else {
						fmt.Println("中间层错误", err)
					}
				}
			}

		case value, err := <-manager.onLocalMsgChan:
			fmt.Println("前往发送,", value.ClientAddr(), err)
			client := manager.clients[value.ClientAddr()]
			if client != nil {
				// TODO   middle ware 最好做个中间发送层
				_ = client.SendMessage(value.Data())
			}
		}
	}
}

//读取管道数据
func (manager *ConnectManager) OnLoadReadMsg() {
	defer manager.wg.Done()
	for {
		select {
		case value, err := <-manager.onLocalMsgChan:
			fmt.Println("前往发送,", value.ClientAddr(), err)
			client := manager.clients[value.ClientAddr()]
			if client != nil {
				// TODO   middle ware 最好做个中间发送层
				_ = client.SendMessage(value.Data())
			}
		}
	}
}

//监听到信息时
func (manager *ConnectManager) OnConnAndClose() {
	defer manager.wg.Done()
	for {
		select {
		case value, _ := <-manager.disConnectChan:
			client := manager.clients[value]
			client.OnDispose()
		case itemClient, _ := <-manager.connectChan:
			manager.addItemClient(itemClient)
		}
	}
}

var managerInstance map[string]*ConnectManager

// 获取一个连接器
func GetManagerForName(name string) *ConnectManager {
	if _, ok := managerInstance[name]; ok == false {
		managerInstance[name] = MakeConnectManager(name)
	}
	return managerInstance[name]
}

func init() {
	managerInstance = make(map[string]*ConnectManager)
}
