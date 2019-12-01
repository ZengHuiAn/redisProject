package common

import "redisProject/src/net_struct"

type OnClientMessage interface {
	//发送数据接口
	SendMessage(data net_struct.TCPClientData) error
	// 断开客户端
	OnDispose ()
	//连接---->>>
	OnConnect ()
}

type Message struct {
	clientAddr string
	data *net_struct.TCPClientData
}

func (m *Message) Data() *net_struct.TCPClientData {
	return m.data
}

func (m *Message) ClientAddr() string {
	return m.clientAddr
}

func MakeMessage(clientAddr string,readData net_struct.TCPClientData) Message {
	return Message{clientAddr:clientAddr,data:&readData}
}

type ProtoMsgHead struct {
	EventName string
	MsgID uint32
}

func NewProtoMsgHead(eventName string, msgID uint32) ProtoMsgHead {
	return ProtoMsgHead{EventName: eventName, MsgID: msgID}
}
